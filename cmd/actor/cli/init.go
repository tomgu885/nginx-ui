package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"nginx-ui/actor/router"
	"nginx-ui/pkg/nginx"
	"nginx-ui/server/settings"
)

var configFile string
var rootCmd = &cobra.Command{
	Use:   "actor",
	Short: "节点",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		fmt.Println("PreRunE")
		err = settings.Init(configFile)

		fmt.Println("setting.SErver", settings.ServerSettings.RunMode)
		return
	},
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "监听服务",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return router.InitRouter().Run(":8080")
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

func init() {
	rootCmd.Flags().StringVarP(&configFile, "configFile", "c", "app.ini", "config file location")
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(restartCmd)
}

func Execute() (err error) {
	return rootCmd.Execute()
}
