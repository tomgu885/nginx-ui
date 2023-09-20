package model

import (
	"fmt"
	"nginx-ui/server/settings"
	"testing"
)

func TestMain(m *testing.M) {
	confPath := "../../app.ini"
	settings.Init(confPath)
	//db.
	Init()
	m.Run()
}

func TestHello(t *testing.T) {
	fmt.Println("hello world")
	err := db.Exec("select 1+6").Error
	if err != nil {
		return
	}
}
