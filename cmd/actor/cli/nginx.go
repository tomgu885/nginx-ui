package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"nginx-ui/actor/services"
	"nginx-ui/pkg/nginx"
	"nginx-ui/server/model"
	"strconv"
	"strings"
)

var testCmd = &cobra.Command{
	Use:   "nginx_test",
	Short: "test nginx config",
	PreRun: func(cmd *cobra.Command, args []string) {
		model.Init()
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		pid := nginx.GetNginxPIDPath()
		fmt.Println("pid", pid)
		result, err := nginx.TestConf()
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println("test result:", result)
		if strings.Contains(result, "[warn]") {
			fmt.Println("config failed")
		}
		return
	},
}

var configServerCmd = &cobra.Command{
	Use:   "config_server",
	Short: "一个服务的配置",
	PreRun: func(cmd *cobra.Command, args []string) {
		model.Init()
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if len(args) == 0 {
			fmt.Println("请输入站点id")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return
		}

		conf, err := services.ServerConfig(uint(id))
		if err != nil {
			return
		}
		fmt.Println("conf")
		fmt.Println(conf)
		return
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "加载配置",
	PreRun: func(cmd *cobra.Command, args []string) {
		model.Init()
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		var req model.SiteListReq
		req.WithCert = true
		sites, total, err := model.GetSites(req)
		fmt.Printf("total:%d\n", total)
		for _, row := range sites {
			fmt.Printf("cert: ", row.Cert)
		}
		return
	},
}

var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "重启 nginx 服务",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		result, err := nginx.Reload()
		if err != nil {
			return
		}
		fmt.Println("restart...", result)
		return
	},
}
