package cli

import (
    "github.com/spf13/cobra"
    "nginx-ui/actor/router"
    "nginx-ui/server/settings"
)

var configFile string
var rootCmd = &cobra.Command{
    Use:   "actor",
    Short: "节点",
    PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
        return settings.Init(configFile)
    },
}

var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "监听服务",
    RunE: func(cmd *cobra.Command, args []string) (err error) {
        return router.InitRouter().Run(":8080")
    },
}

func init() {
    rootCmd.Flags().StringVarP(&configFile, "configFile", "c", "app.ini", "config file location")
    rootCmd.AddCommand(serveCmd)
    rootCmd.AddCommand(restartCmd)
    rootCmd.AddCommand(testCmd)
    rootCmd.AddCommand()
}

func Execute() (err error) {
    return rootCmd.Execute()
}
