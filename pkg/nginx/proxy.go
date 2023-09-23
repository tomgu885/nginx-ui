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
    masterUrl := settings.NginxSettings.MasterUrl
    fmt.Println("Master:", masterUrl)
    // listen
    server := "server {\n"
    server += fmt.Sprintf("\tserver_name %s;\n", strings.ReplaceAll(n.Site.Domains, "\n", " "))
    server += fmt.Sprintf("\tlisten %s;\n", n.Site.HttpPorts)

    // location let's encrypt
    server += fmt.Sprintf("\tlocation %s", masterUrl)

    server += "}"

    if n.Site.SslCertState == model.SslCertStateOk && n.Site.SslEnable == model.StateEnable {

    }
    return server
}
