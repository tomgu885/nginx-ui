package model

import (
    "nginx-ui/pkg/cert/dns"
)

type DnsCredential struct {
    BaseModel
    Name     string      `json:"name"`
    Config   *dns.Config `json:"config,omitempty" gorm:"serializer:json"`
    Provider string      `json:"provider"`
}
