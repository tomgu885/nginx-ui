package services

import "nginx-ui/server/model"

type SslService struct {
}

func (SslService) IssueCert(site model.Site) (err error) {
	// 1 check dns / 2. issue , 3. save data to db, 4. update cdn_node

	return
}

func (SslService) UpdateCert(site model.Site) (err error) {
	return
}

// https://networkbit.ch/golang-dns-lookup/
func CheckDns(domain string) (ok bool, err error) {

	return true, nil
}
