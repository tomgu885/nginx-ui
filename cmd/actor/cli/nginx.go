package cli

import (
    "fmt"
    "github.com/go-resty/resty/v2"
    "github.com/spf13/cobra"
    model2 "nginx-ui/actor/model"
    "nginx-ui/actor/services"
    "nginx-ui/pkg/logger"
    "nginx-ui/pkg/nginx"
    "nginx-ui/pkg/settings"
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
    //PreRun: func(cmd *cobra.Command, args []string) {
    //
    //},
    RunE: func(cmd *cobra.Command, args []string) (err error) {
        url := settings.NginxSettings.MasterUrl + "/api/updates"
        client := resty.New()

        var result model2.SitePageResult
        _, err = client.R().SetResult(&result).Get(url)

        if err != nil {
            return
        }

        //fmt.Printf("resp: %s\n", resp.Body())
        fmt.Println("resultCode:", result.Code)
        fmt.Println("sites count:", len(result.Data.List))
        for _, row := range result.Data.List {
            ok, errW := services.DumpConfig(row, true)
            if errW != nil {
                logger.Errorf("writeFile: %s", errW)
                continue
            }

            logger.Infof("write :%t", ok)
        }
        //var req model.SiteListReq
        //req.WithCert = true
        //sites, total, err := model.GetSites(req)
        //fmt.Printf("total:%d\n", total)
        //for _, row := range sites {
        //    fmt.Printf("cert: ", row.Cert)
        //}
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
