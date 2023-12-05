package nginx

import (
    "fmt"
    "nginx-ui/pkg/settings"
    "nginx-ui/server/model"
    "strings"
)

// https://nginxui.com/zh_CN/guide/nginx-proxy-example.html
type NginxProxy struct {
    Site model.Site
    //Name       string `json:"name"`
    //Domains    string `json:"domains"`
    //HttpPorts  string `json:"http_ports"`  // 空格隔开
    //HttpsPorts string `json:"https_ports"` //
    //SslPath    string `json:"ssl_path"`
    //UpstreamIps string `json:"upstream_ips"` // 空格隔开
    //Upstream
}

func NewNginxProxy(site model.Site) NginxProxy {
    return NginxProxy{
        Site: site,
    }
}

func (n NginxProxy) BuildConfig() string {
    // listen
    server := fmt.Sprintf("## site:%d lastModified: %d\n", n.Site.ID, n.Site.UpdatedAt)
    server += "server {\n"
    server += fmt.Sprintf("\tserver_name %s;\n", n.Site.ServerDomains())
    server += fmt.Sprintf("\tadd_header via %s;\n", settings.NginxSettings.Via)
    for _, _port := range strings.Split(n.Site.HttpPorts, " ") {
        server += fmt.Sprintf("\tlisten %s;\n", _port)
    }

    // location let's encrypt
    server += n.acmeProxy()
    server += n.proxyBlock()

    server += "}"
    if n.Site.SslOk() {
        server += "\n# ssl\n"
        server += n.sslProxyBlock()
    }

    if n.Site.SslCertState == model.SslCertStateOk && n.Site.SslEnable == model.StateEnable {

    }
    return server
}

func (n NginxProxy) proxyBlock() (proxy string) {
    proxy = "\tlocation / {\n"
    // @todo 跳转 / redirect
    proxy += "\t\tproxy_http_version 1.1;\n"
    if "" == n.Site.UpstreamHost {
        proxy += "\t\tproxy_set_header Host $host;\n"
    } else {
        proxy += fmt.Sprintf("\t\tproxy_set_header Host %s;\n", n.Site.UpstreamHost)
    }

    proxy += fmt.Sprintf("\t\tproxy_set_header Via %s;\n", settings.NginxSettings.Via)
    proxy += "\t\tproxy_set_header X-Real_IP $remote_addr;\n"
    proxy += "\t\tproxy_set_header X-Forwarded-Proto $scheme;\n"
    proxy += "\t\tproxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;\n"
    proxy += "\t\tproxy_set_header Upgrade $http_upgrade;\n"
    proxy += "\t\tproxy_set_header Connection $connection_upgrade;\n"
    for _, ip := range strings.Split(n.Site.UpstreamIps, " ") {
        proxy += fmt.Sprintf("\t\tproxy_pass $scheme://%s:$server_port;\n", ip)
    }

    proxy += "\t}\n"
    return
}

// sslProxyBlock
// proxy_ssl_verify
// ssl_certificate /etc/letsencrypt/live/ssl3.cloud2hk.com/fullchain.pem; # managed by Certbot
//    ssl_certificate_key /etc/letsencrypt/live/ssl3.cloud2hk.com/privkey.pem; # managed by Certbot
//    include /etc/letsencrypt/options-ssl-nginx.conf; # managed by Certbot
//    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem; # managed by Certbot
// $server_port
func (n NginxProxy) sslProxyBlock() (proxy string) {
    proxy = "server {\n"
    for _, _port := range strings.Split(n.Site.HttpsPorts, " ") {
        proxy += fmt.Sprintf("\tlisten %s;\n", _port)
    }
    proxy += fmt.Sprintf("\tserver_name %s;\n", n.Site.ServerDomains())
    proxy += fmt.Sprintf("\tssl_certificate %s;\n", n.Site.FullchainFile()) // # managed by Certbot
    proxy += fmt.Sprintf("\tssl_certificate_key %s;\n", n.Site.PrivatePemFile())
    proxy += "\tinclude /etc/letsencrypt/options-ssl-nginx.conf;\n" // # managed by Certbot
    proxy += "\tssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;\n"   //# managed by Certbot
    proxy += "\tlocation / {\n"
    proxy += fmt.Sprintf("\tadd_header via %s;\n", settings.NginxSettings.Via)
    proxy += "\t\tproxy_ssl_verify off;\n"
    proxy += "\t\tproxy_http_version 1.1;\n"
    if "" == n.Site.UpstreamHost {
        proxy += "\t\tproxy_set_header Host $host;\n"
    } else {
        proxy += fmt.Sprintf("\t\tproxy_set_header Host %s;\n", n.Site.UpstreamHost)
    }

    proxy += "\t\tproxy_set_header X-Real_IP $remote_addr;\n"
    proxy += "\t\tproxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;\n"
    proxy += "\t\tproxy_set_header Upgrade $http_upgrade;\n"
    proxy += "\t\tproxy_set_header Connection $connection_upgrade;\n"

    for _, ip := range strings.Split(n.Site.UpstreamIps, " ") {
        proxy += fmt.Sprintf("\t\tproxy_pass $scheme://%s:$server_port;\n", ip)
    }

    proxy += "\t}\n" // end of location
    proxy += "}\n"   // end of server
    return proxy
}

func (n NginxProxy) acmeProxy() (acme string) {
    challengeUrl := "http://" + settings.NginxSettings.MasterIp + ":" + settings.ServerSettings.HTTPChallengePort
    //fmt.Println("Master:", masterUrl)
    acme = "\tlocation /.well-known/acme-challenge {\n"
    acme += "\t\tproxy_set_header Host $host;\n"
    acme += "\t\tproxy_set_header X-Real_IP $remote_addr;\n"
    //acme += "\t\tproxy_set_header X-Forwarded-For $re;\n"
    acme += fmt.Sprintf("\t\tproxy_pass %s;\n", challengeUrl)
    acme += "\t}\n"

    return
}
