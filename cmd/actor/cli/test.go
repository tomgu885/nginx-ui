package cli

import (
    "fmt"
    "github.com/spf13/cobra"
    "nginx-ui/pkg/settings"
)

var hellCmd = &cobra.Command{
    Use:   "hello",
    Short: "just output hello",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("hello from cdn node")
    },
}

var testCmd = &cobra.Command{
    Use:   "test",
    Short: "开发测试",
    Run: func(cmd *cobra.Command, args []string) {
        //pid := nginx.GetNginxPIDPath()
        //fmt.Println("pid", pid)
        //result, err := nginx.TestConf()
        //if err != nil {
        //    fmt.Println("err:", err)
        //    return
        //}
        //fmt.Println("test result:", result)
        //if strings.Contains(result, "[warn]") {
        //    fmt.Println("config failed")
        //}
        nodeList := settings.GetNodeList()
        for _, url := range nodeList {
            fmt.Printf("node:%s\n", url)
        }

        nodeIps := settings.GetNodeIps()
        for _, ip := range nodeIps {
            fmt.Println("ip:", ip)
        }

        fmt.Printf("masterIp:|%s|\n", settings.NginxSettings.MasterIp)
        fmt.Printf("NodePort:|%s|\n", settings.NginxSettings.NodePort)
    },
}
