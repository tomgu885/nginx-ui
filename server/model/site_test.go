package model

import (
	"fmt"
	"testing"
)

func TestGetSiteByName(t *testing.T) {
	domain := "www.baidu.com"
	site, err := GetSiteByName(domain)
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("site.id", site.ID)
}
