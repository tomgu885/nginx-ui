package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"nginx-ui/pkg/nginx"
	"strings"
)

var testCmd = &cobra.Command{
	Use:   "nginx_test",
	Short: "test nginx config",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		pid := nginx.GetNginxPIDPath()
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
