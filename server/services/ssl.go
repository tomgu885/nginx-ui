package services

import (
    "nginx-ui/pkg/cert"
    "nginx-ui/server/model"
    "strings"
)

type SslService struct {
}

func (SslService) IssueCert(site model.Site) (err error) {
    // 1 check dns / 2. issue , 3. save data to db, 4. update cdn_node
    // @todo checkdns
    model.UpdateSiteById(site.ID, map[string]any{
        "ssl_cert_state": model.SslCertStateInit,
        "ssl_obtain_log": "",
    })

    domains := strings.Split(site.Domains, "\n")
    privateKey, fullchain, err := cert.ObtainCert(domains)
    if err != nil {
        model.UpdateSiteById(site.ID, model.Site{
            SslCertState: model.SslCertStateFail,
            SslObtainLog: "获取失败\n" + err.Error(),
        })

        return
    }

    model.UpdateSiteById(site.ID, model.Site{
        SslEnable:       model.StateEnable,
        SslCertState:    model.SslCertStateOk,
        SslFullchainCer: string(fullchain),
        SslPrivateKey:   string(privateKey),
    })
    return
}

func (SslService) UpdateCert(site model.Site) (err error) {
    return
}

// https://networkbit.ch/golang-dns-lookup/
func CheckDns(domain string) (ok bool, err error) {

    return true, nil
}
