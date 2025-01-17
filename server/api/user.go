package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"nginx-ui/server/model"
	"nginx-ui/server/query"
	"nginx-ui/pkg/settings"
)

func GetUsers(c *gin.Context) {
	data := model.GetUserList(c, c.Query("name"))

	c.JSON(http.StatusOK, data)
}

func GetUser(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	u := query.Auth

	user, err := u.FirstByID(id)

	if err != nil {
		ErrHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

type UserJson struct {
	Name     string `json:"name" binding:"required,max=255"`
	Password string `json:"password" binding:"max=255"`
}

func AddUser(c *gin.Context) {
	var json UserJson
	ok := BindAndValid(c, &json)
	if !ok {
		return
	}

	u := query.Auth

	pwd, err := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	if err != nil {
		ErrHandler(c, err)
		return
	}
	json.Password = string(pwd)

	user := model.Auth{
		Username:          json.Name,
		EncryptedPassword: json.Password,
	}

	err = u.Create(&user)

	if err != nil {
		ErrHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, user)

}

func EditUser(c *gin.Context) {
	userId := cast.ToInt(c.Param("id"))

	if settings.ServerSettings.Demo && userId == 1 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Prohibit changing root password in demo",
		})
		return
	}

	var json UserJson
	ok := BindAndValid(c, &json)
	if !ok {
		return
	}

	u := query.Auth
	user, err := u.FirstByID(userId)

	if err != nil {
		ErrHandler(c, err)
		return
	}
	edit := &model.Auth{
		Username: json.Name,
	}

	// encrypt password
	if json.Password != "" {
		var pwd []byte
		pwd, err = bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
		if err != nil {
			ErrHandler(c, err)
			return
		}
		edit.EncryptedPassword = string(pwd)
	}

	_, err = u.Where(u.ID.Eq(userId)).Updates(&edit)

	if err != nil {
		ErrHandler(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := cast.ToInt(c.Param("id"))

	if cast.ToInt(id) == 1 {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Prohibit deleting the default user",
		})
		return
	}

	u := query.Auth
	err := u.DeleteByID(id)
	if err != nil {
		ErrHandler(c, err)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
