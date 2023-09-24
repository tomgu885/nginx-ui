package nginx

import (
	"fmt"
	"nginx-ui/pkg/settings"
	"nginx-ui/server/model"
	"strings"
)

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

func (n NginxProxy) BuildConfig() string {
	// listen
	server := fmt.Sprintf("## site:%d lastModified: %d\n", n.Site.ID, n.Site.UpdatedAt)
	server += "server {\n"
	server += fmt.Sprintf("\tserver_name %s;\n", strings.ReplaceAll(n.Site.Domains, "\n", " "))
	server += fmt.Sprintf("\tlisten %s;\n", n.Site.HttpPorts)

	// location let's encrypt
	server += n.acmeProxy()
	server += n.proxyBlock()
	server += "}"

	if n.Site.SslCertState == model.SslCertStateOk && n.Site.SslEnable == model.StateEnable {

	}
	return server
}

func (n NginxProxy) proxyBlock() (proxy string) {
	proxy = "\tlocation / {\n"
	proxy += "\t\tproxy_http_version 1.1;\n"
	if "" == n.Site.UpstreamHost {
		proxy += "\t\tproxy_set_header Host $host;\n"
	} else {
		proxy += fmt.Sprintf("\t\tproxy_set_header Host %s;\n", n.Site.UpstreamHost)
	}
	proxy += "\t\tproxy_set_header X-Real_IP $remote_addr;\n"
	proxy += "\t\tproxy_set_header X-Forwarded-Proto $scheme;\n"
	proxy += "\t\tproxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;\n"
	proxy += "\t\tproxy_set_header Upgrade $http_upgrade;\n"
	proxy += "\t\tproxy_set_header Connection $connection_upgrade;\n"
	for _, ip := range strings.Split(n.Site.UpstreamIps, " ") {
		proxy += fmt.Sprintf("\t\tproxy_pass $cheme://%s:$port;\n", ip)
	}
	proxy += "\t}\n"
	return
}

// proxy_ssl_verify
func (n NginxProxy) sslProxyBlock() (proxy string) {
	proxy = "\tlocation / {\n"
	proxy += "\t\tproxy_ssl_verify off;\n"

	proxy += "\t}\n"
	return proxy
}

func (n NginxProxy) acmeProxy() (acme string) {
	masterUrl := settings.NginxSettings.MasterUrl
	//fmt.Println("Master:", masterUrl)
	acme = "\tlocation /.well-known/acme-challenge {\n"
	acme += "\t\tproxy_set_header Host $host;\n"
	acme += "\t\tproxy_set_header X-Real_IP $remote_addr;\n"
	acme += "\t\tproxy_set_header X-Forwarded-For $remote_addr:$remote_port;\n"
	acme += fmt.Sprintf("\t\tproxy_pass %s;\n", masterUrl)
	acme += "\t}\n"

	return
}
