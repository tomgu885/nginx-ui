package api

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"nginx-ui/server/model"
)

func GetFileBackupList(c *gin.Context) {
	path := c.Query("path")
	backups := model.GetBackupList(path)

	c.JSON(http.StatusOK, gin.H{
		"backups": backups,
	})
}

func GetFileBackup(c *gin.Context) {
	id := c.Param("id")
	backup := model.GetBackup(com.StrTo(id).MustInt())

	c.JSON(http.StatusOK, backup)
}
