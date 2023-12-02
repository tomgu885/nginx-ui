package cert

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "github.com/go-acme/lego/v4/certificate"
    "github.com/go-acme/lego/v4/challenge/http01"
    "github.com/go-acme/lego/v4/lego"
    "nginx-ui/pkg/logger"
    "nginx-ui/pkg/settings"
)

func ObtainCert(domains []string) (privateBs, fullcert []byte, err error) {
    privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        return
    }
    myUser := MyUser{
        Email: settings.ServerSettings.Email,
        Key:   privateKey,
    }

    config := lego.NewConfig(&myUser)

    if settings.ServerSettings.Demo {
        config.CADirURL = "https://acme-staging-v02.api.letsencrypt.org/directory"
    }

    client, err := lego.NewClient(config)
    if err != nil {
        logger.Errorf("lego.NewClient|fail to create lego client %v", err)
        return
    }

    err = client.Challenge.SetHTTP01Provider(http01.NewProviderServer("", settings.ServerSettings.HTTPChallengePort))
    if err != nil {
        logger.Errorf("lego|fail to SetHTTP01Provider %v", err)
        return
    }

    request := certificate.ObtainRequest{
        Domains: domains,
        Bundle:  true,
    }

    certificates, err := client.Certificate.Obtain(request)
    if err != nil {
        logger.Errorf("client.Certificate.Obtain failed: %v", err)
        return
    }

    return
}

func RenewCert(domains []string) (privateBs, fullCert []byte, err error) {

    return
}
