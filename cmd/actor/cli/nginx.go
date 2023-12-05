package cli

import (
    "fmt"
    "github.com/spf13/cobra"
    "nginx-ui/actor/services"
    "nginx-ui/pkg/logger"
    "nginx-ui/pkg/nginx"
    "nginx-ui/server/model"
    "strconv"
)

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
    //PreRun: func(cmd *cobra.Command, args []string) {
    //
    //},
    RunE: func(cmd *cobra.Command, args []string) (err error) {
        logger.Info("config reload")
        err = services.ServerConfigReload(true, false, "cmd")
        if err != nil {
            fmt.Println("config reload failed:", err.Error())
            return
        }
        fmt.Println("config reloaded.")
        return
    },
}

var reloadCmd = &cobra.Command{
    Use:   "reload",
    Short: "软重启",
    RunE: func(cmd *cobra.Command, args []string) (err error) {
        result, err := nginx.Reload()
        if err != nil {
            return
        }
        fmt.Println("result:", result)
        return
    },
}

var restartCmd = &cobra.Command{
    Use:   "restart",
    Short: "重启 nginx 服务",
    RunE: func(cmd *cobra.Command, args []string) (err error) {
        result, err := nginx.Restart()
        if err != nil {
            logger.Errorf("restart failed: %v", err)
            return
        }
        fmt.Println("restart:", result)
        return
    },
}
