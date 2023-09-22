package nginx

type NginxProxy struct {
	Name       string `json:"name"`
	Domains    string `json:"domains"`
	HttpPorts  string `json:"http_ports"`  // 空格隔开
	HttpsPorts string `json:"https_ports"` //
	SslPath    string `json:"ssl_path"`
}

func (NginxProxy) BuildConfig() string {
	// listen
	server := "server {"
	server += "server_name"

	return server
}
