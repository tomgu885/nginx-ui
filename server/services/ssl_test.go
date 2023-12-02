package services

import (
    "fmt"
    "nginx-ui/server/model"
    "testing"
)

func TestSslService_IssueCert(t *testing.T) {
    serv := SslService{}
    site, err := model.GetSiteById(1)
    if err != nil {
        fmt.Println("failed to get site")
        return
    }
    err = serv.IssueCert(site)
    if err != nil {
        fmt.Println("IssueCert failed")
        return
    }

    fmt.Println("success")
}
