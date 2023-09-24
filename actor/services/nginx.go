package services

import (
	"nginx-ui/pkg/nginx"
	"nginx-ui/server/model"
)

func ServerConfig(id uint) (conf string, err error) {
	cfg := nginx.NginxProxy{}
	cfg.Site, err = model.GetSiteById(uint(id))

	if err != nil {
		return
	}

	conf = cfg.BuildConfig()

	return
}

func ServerUpdate() {

}
