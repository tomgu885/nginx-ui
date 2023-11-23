package model

import "fmt"

const (
	_KeySetting = "setting:%s"
)

func keySetting(name string) string {
	return fmt.Sprintf(_KeySetting, name)
}
