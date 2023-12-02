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
        payload := &cert.ConfigPayload{
            ServerName:      []string{"ss1l.cloud2hk.com"},
            ChallengeMethod: "http01",
            DNSCredentialID: 0,
        }
        logChan := make(chan string)
        errChan := make(chan error)
        go cert.IssueCert(payload, logChan, errChan)

        for msg := range logChan {
            fmt.Println("msg:", msg)
        }
    },
}
