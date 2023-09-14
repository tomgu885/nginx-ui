package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nginx-ui/server/internal/translation"
)

func GetTranslation(c *gin.Context) {
	code := c.Param("code")

	c.JSON(http.StatusOK, translation.GetTranslation(code))
}
