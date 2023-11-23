package model

import (
	"context"
)

type Setting struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Content  string `json:"content"`
	DataType string
}

func (Setting) TableName() string {
	return "settings"
}

func LoadSettFromDb() {
	var items []Setting
	err := db.Find(&items).Error
	if err != nil {
		panic("fail to get items")
		return
	}

	if len(items) == 0 {
		return
	}

	for _, item := range items {
		_key := keySetting(item.Name)
		rdb.Set(context.Background(), _key, item.Content, 0)
	}

	return
}

// json, array, int, string
func GetSettingString(name string) (str string, err error) {
	_key := keySetting(name)
	str, err = rdb.Get(context.Background(), _key).Result()
	return
}
