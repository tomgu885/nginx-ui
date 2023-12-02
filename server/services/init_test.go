package services

import (
    "fmt"
    "nginx-ui/pkg/settings"
    "nginx-ui/server/internal/boot"
    "testing"
)

func TestMain(m *testing.M) {
    err := settings.Init("../../app.ini")
    if err != nil {
        fmt.Println("failed to init config:", err.Error())
        return
    }
    boot.Kernel()
    m.Run()
}
