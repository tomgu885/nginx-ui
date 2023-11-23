package model

type Cert struct {
	BaseModel
	SiteId       uint   `json:"site_id"`
	Name         string `json:"name"`
	Domains      string `json:"domains"`
	SslCertState int8   `json:"ssl_cert_state"` // ssl 证书 状态 1: 申请开始, 2: 已完成 , 3 错误
	SslKey       string `json:"ssl_key"`
	SslCer       string `json:"ssl_cer"`
	ExpiredAt    int64  `json:"expired_at"`
	Log          string `json:"log"`
}

func (Cert) TableName() string {
	return "certs"
}

func UpdateCertLog(id uint, logS string) (err error) {
	err = db.Model(Cert{}).Where("id", id).Update("log", logS).Error
	return
}

func GetAutoCertList() (list []Cert, err error) {
	err = db.Find(&list).Error
	return
}
