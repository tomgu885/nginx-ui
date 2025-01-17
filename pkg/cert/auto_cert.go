package cert

import (
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "github.com/go-acme/lego/v4/certificate"
    "github.com/go-acme/lego/v4/challenge/http01"
    "github.com/go-acme/lego/v4/lego"
    "github.com/go-acme/lego/v4/registration"
    "nginx-ui/pkg/logger"
    "nginx-ui/pkg/settings"
)

func ObtainCert(domains []string) (privateKey, fullchain []byte, err error) {
    _privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    if err != nil {
        return
    }
    logger.Infof("settings.ServerSettings.Email: %s", settings.ServerSettings.Email)
    myUser := MyUser{
        Email: settings.ServerSettings.Email,
        Key:   _privateKey,
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

    _, err = client.Registration.Register(registration.RegisterOptions{TermsOfServiceAgreed: true})
    if err != nil {
        logger.Errorf("failed to register", err.Error())
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

    //fmt.Println("certificates", string(certificates.Certificate))
    fullchain = certificates.Certificate
    privateKey = certificates.PrivateKey
    return
}

func RenewCert(domains []string) (privateBs, fullCert []byte, err error) {

    return
}
