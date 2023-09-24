package model

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"nginx-ui/pkg/logger"
	"nginx-ui/pkg/utils"
	"nginx-ui/pkg/validator"
	"strconv"
	"strings"
)

type Site struct {
	Model
	Name                 string `json:"name"`
	Domains              string `json:"domains"`
	DomainCount          int    `json:"domain_count"` // 域名数量
	State                int8   `json:"state"`
	SslEnable            int8   `json:"ssl_enable"`
	SslCertState         int8   `json:"ssl_cert_state"` // ssl 证书 状态 1: 申请开始, 2: 已完成 , 3 错误
	HttpPorts            string `json:"http_ports"`
	HttpsPorts           string `json:"https_ports"`
	HttpRedirect         int8   `json:"http_redirect"`
	UpstreamPortPolicy   int8   `json:"upstream_port_policy"`   // 1: 同协议，同端口, 2: 回落到 http(80)
	UpstreamRotatePolicy int8   `json:"upstream_rotate_policy"` // 1: 轮询, 2: ip hash(暂时不实现)
	UpstreamIps          string `json:"upstream_ips"`
	UpstreamHost         string `json:"upstream_host"`
	//Path     string `json:"path"`
	//Advanced bool   `json:"advanced"`
	Cert Cert `json:"cert" gorm:"foreignKey:site_id"`
}

const (
	//ssl 证书 状态 1: 申请开始, 2: 已完成
	SslCertStatePending = 1 // 等待申请
	SslCertStateInit    = 2 // 开始申请
	SslCertStateOk      = 3 // 完成
	SslCertStateFail    = 4 // 错误
)

var (
	specialPorts = []int{110, 995}
)

func (Site) TableName() string {
	return "sites"
}

type SiteListReq struct {
	PageInfo
	Domain   string
	WithCert bool
}

type SiteCreateReq struct {
	Domains              string `json:"domains"`
	SslEnable            int8   `json:"ssl_enable"`
	HttpPorts            string `json:"http_ports" example:"多个端口以逗号隔开"`
	HttpsPorts           string `json:"https_ports" example:"多个端口以逗号隔开"`
	HttpRedirect         int8   `json:"http_redirect"`
	UpstreamPortPolicy   int8   `json:"upstream_port_policy"`   // 1: 同协议，同端口, 2: 回落到 http(80)
	UpstreamRotatePolicy int8   `json:"upstream_rotate_policy"` // 1: 轮询, 2: ip hash(暂时不实现)
	UpstreamIps          string `json:"upstream_ips"`
	UpstreamHost         string `json:"upstream_host"`
}

type SiteUpdateReq struct {
	ID uint `json:"id"`
	SiteCreateReq
}

func UpdateSite(req SiteUpdateReq) (err error) {
	row, err := GetSiteById(req.ID)
	if err != nil {
		return
	}

	if row.ID == 0 {
		return errors.New("站点不存在")
	}

	err = db.Where("id", req.ID).Updates(Site{
		Domains:              req.Domains,
		SslEnable:            req.SslEnable,
		HttpPorts:            req.HttpPorts,
		HttpsPorts:           req.HttpsPorts,
		HttpRedirect:         req.HttpRedirect,
		UpstreamPortPolicy:   req.UpstreamPortPolicy,
		UpstreamRotatePolicy: req.UpstreamRotatePolicy,
		UpstreamIps:          req.UpstreamIps,
		UpstreamHost:         req.UpstreamHost,
	}).Error

	if err != nil {
		return
	}

	// @todo 判断域名是否变化
	return
}

func CreateSite(req SiteCreateReq) (err error) {
	// 检查端口是否小于 ,检查端口重复
	req.Domains = strings.ReplaceAll(req.Domains, "\r\n", "\n")
	req.Domains = strings.TrimSpace(req.Domains)
	if "" == req.Domains {
		return errors.New("域名不能为空")
	}

	// 检查 域名正确性
	req.Domains = strings.ReplaceAll(req.Domains, " ", "\n")
	domains := strings.Split(req.Domains, "\n")
	for _, domain := range domains {
		if !validator.IsValidDomain(domain) {
			return errors.New(fmt.Sprintf("%s 不是域名", domain))
		}
	}
	exists, err := GetSiteByName(domains[0])
	if err != nil {
		return
	}

	if exists.ID > 0 {
		return errors.New("第一个域名已经存在")
	}

	req.HttpPorts = utils.RemoveDuplicatedSpace(req.HttpPorts)
	req.HttpsPorts = utils.RemoveDuplicatedSpace(req.HttpsPorts)
	if req.SslEnable == StateEnable && "" == req.HttpsPorts {
		return errors.New("https开启时必须有https端口")
	}
	if len(req.HttpPorts) == 0 {
		return errors.New("http端口不能为空,会导致无法申请ssl")
	}
	// 检查端口正确性: 是否为数字
	portsSet := []int{}
	httpPortList := strings.Split(req.HttpPorts, " ")
	for _, port := range httpPortList {
		dig, errA := strconv.Atoi(port)
		if errA != nil {
			return errors.New(fmt.Sprintf("%s 端口不是数字", port))
		}
		if utils.InArray(dig, portsSet) {
			return errors.New(fmt.Sprintf("%s 端口重复", port))
		}
		if dig < 80 {
			return errors.New("端口号不能小于80")
		}
		portsSet = append(portsSet, dig)
	}
	if len(req.HttpsPorts) > 0 {
		httpsPortList := strings.Split(req.HttpsPorts, " ")
		for _, port := range httpsPortList {
			dig, errA := strconv.Atoi(port)
			if errA != nil {
				return errors.New(fmt.Sprintf("%s 端口不是数字", port))
			}
			if utils.InArray(dig, portsSet) {
				return errors.New(fmt.Sprintf("%s 端口重复", port))
			}

			portsSet = append(portsSet, dig)
		}
	} else {
		req.SslEnable = StateDisabled
	}

	row := Site{
		Name:                 domains[0],
		Domains:              req.Domains,
		DomainCount:          len(domains),
		State:                StateEnable,
		SslEnable:            req.SslEnable,
		SslCertState:         SslCertStateInit,
		HttpPorts:            req.HttpPorts,
		HttpsPorts:           req.HttpsPorts,
		HttpRedirect:         req.HttpRedirect,
		UpstreamPortPolicy:   req.UpstreamPortPolicy,
		UpstreamRotatePolicy: req.UpstreamRotatePolicy,
		UpstreamIps:          req.UpstreamIps,
		UpstreamHost:         req.UpstreamHost,
	}

	err = db.Create(&row).Error
	if err != nil {
		logger.Errorf("db.Create(site) failed:%v", err)
		return errors.New("插入数据错误")
	}

	if StateEnable == req.SslEnable {
		ce := Cert{
			SiteId:       row.ID,
			Name:         domains[0],
			Domains:      req.Domains,
			SslKey:       "",
			SslCer:       "",
			SslCertState: SslCertStatePending,
			ExpiredAt:    0,
			Log:          "",
		}
		fmt.Println("ce", ce)
		errCe := db.Create(&ce).Error

		if errCe != nil {
			return
		}
	}

	return
}

func GetSiteById(id uint) (site Site, err error) {
	err = db.Find(&site, id).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func GetSiteByName(domain string) (s Site, err error) {
	err = db.Where("name", domain).First(&s).Error
	if err == gorm.ErrRecordNotFound {
		err = nil
	}
	return
}

func GetSites(req SiteListReq) (list []Site, total int64, err error) {
	q := db.Model(&Site{})

	if req.Domain != "" {
		q.Where("domains like ?", req.Domain)
	}

	if req.StartCreatedAt > 0 {
		q.Where("created_at >= ?", req.StartCreatedAt)
	}

	if req.EndCreatedAt > 0 {
		q.Where("created_at <= ?", req.EndCreatedAt)
	}

	err = q.Count(&total).Error
	if err != nil {
		return
	}

	offset, limit := req.GetOffsetLimit()
	if req.WithCert {
		q.Preload("Cert")
	}
	err = q.Offset(offset).Limit(limit).Order("id DESC").Find(&list).Error
	return
}
