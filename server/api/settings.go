package api

import "github.com/gin-gonic/gin"

//
func GetSettings(c *gin.Context) {

}

func CreateSetting(c *gin.Context) {

}

// DeleteSetting
// DELETE /api/setting/:id
func DeleteSetting(c *gin.Context) {

}

//func GetSettings(c *gin.Context) {
//	c.JSON(http.StatusOK, gin.H{
//		"server": settings.ServerSettings,
//		"nginx":  settings.NginxSettings,
//		"openai": settings.OpenAISettings,
//	})
//}
//
//func SaveSettings(c *gin.Context) {
//	var json struct {
//		Server settings.Server `json:"server"`
//		Nginx  settings.Nginx  `json:"nginx"`
//		Openai settings.OpenAI `json:"openai"`
//	}
//
//	if !BindAndValid(c, &json) {
//		return
//	}
//
//	settings.ServerSettings = json.Server
//	settings.NginxSettings = json.Nginx
//	settings.OpenAISettings = json.Openai
//
//	settings.ReflectFrom()
//
//	err := settings.Save()
//	if err != nil {
//		ErrHandler(c, err)
//		return
//	}
//
//	GetSettings(c)
//}
