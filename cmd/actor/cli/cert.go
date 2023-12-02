package cli

import (
    "fmt"
    "github.com/spf13/cobra"
    "nginx-ui/pkg/cert"
)

var IssueCertCmd = &cobra.Command{
    Use:   "cert_issue",
    Short: "申请证书",
    Run: func(cmd *cobra.Command, args []string) {
        privateKey, fullchain, err := cert.ObtainCert(args)
        if err != nil {
            fmt.Println("failed:", err.Error())
            return
        }

        fmt.Println("private:")
        fmt.Println(string(privateKey))
        fmt.Println("fullchain")
        fmt.Println(string(fullchain))

        //payload := &cert.ConfigPayload{
        //    ServerName:      []string{"ss1l.cloud2hk.com"},
        //    ChallengeMethod: "http01",
        //    DNSCredentialID: 0,
        //}
        //logChan := make(chan string)
        //errChan := make(chan error)

    },
}
