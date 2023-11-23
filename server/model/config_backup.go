package model

import (
	"nginx-ui/pkg/logger"
	"os"
	"path/filepath"
)

type ConfigBackup struct {
	BaseModel

	Name     string `json:"name"`
	FilePath string `json:"file_path"`
	Content  string `json:"content" gorm:"type:text"`
}

type ConfigBackupListItem struct {
	BaseModel

	Name     string `json:"name"`
	FilePath string `json:"file_path"`
}

func GetBackupList(path string) (configs []ConfigBackupListItem) {
	db.Model(&ConfigBackup{}).
		Where(&ConfigBackup{FilePath: path}).
		Find(&configs)

	return
}

func GetBackup(id int) (config ConfigBackup) {
	db.First(&config, id)

	return
}

func CreateBackup(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		logger.Error(err)
	}

	config := ConfigBackup{Name: filepath.Base(path), FilePath: path, Content: string(content)}
	result := db.Create(&config)
	if result.Error != nil {
		logger.Error(result.Error)
	}
}
