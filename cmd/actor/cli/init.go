package cli

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/spf13/cobra"
    "nginx-ui/actor/router"
    "nginx-ui/pkg/settings"
)

var configFile string
var rootCmd = &cobra.Command{
    Use:   "actor",
    Short: "节点",
    PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
        return settings.Init(configFile)
    },
}

var serverTestCmd = &cobra.Command{
    Use:   "serve_test",
    Short: "测试 ssl challenge",
    Run: func(cmd *cobra.Command, args []string) {
        r := gin.Default()

        r.GET("/", func(c *gin.Context) {
            c.String(200, "Hello from ssl challenge")
        })

        r.Run(":9180")
    },
}

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "监听服务",
    RunE: func(cmd *cobra.Command, args []string) (err error) {
        fmt.Println("start run serve")
        return router.InitRouter().Run(fmt.Sprintf(":%s", settings.NginxSettings.NodePort))
    },
}

func init() {
    rootCmd.Flags().StringVarP(&configFile, "configFile", "c", "node.ini", "config file location")
    rootCmd.AddCommand(serveCmd)
    rootCmd.AddCommand(serverTestCmd)
    rootCmd.AddCommand(restartCmd)
    rootCmd.AddCommand(reloadCmd)
    rootCmd.AddCommand(testCmd)
    rootCmd.AddCommand(configServerCmd)
    rootCmd.AddCommand(configCmd)
    rootCmd.AddCommand(hellCmd)

    rootCmd.AddCommand(IssueCertCmd)

}

func Execute() (err error) {
    return rootCmd.Execute()
}
