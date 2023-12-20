package api

import (
	"encoding/json"
	"ginPratice/middlewares"
	"ginPratice/models"
	"ginPratice/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
}

func (msg MessageController) MsgUp(c *gin.Context) {
	db := utils.InitDB()
	formData, _ := c.MultipartForm()
	pictures := formData.File["pictures"]
	message := formData.Value["message"]
	paths, err := utils.FileUploadTools(pictures)
	if err != nil {
		utils.Fail(c, utils.InvalidArgs)
	}
	pathsSave := strings.Join(paths, "&")
	user, _ := c.Get("user")
	claims := user.(*middlewares.UserClaims)
	messages := models.Message{
		UserID:   int(claims.ID),
		Message:  message[0],
		Pictures: pathsSave,
	}
	db.Create(&messages)
	messagesByte, _ := json.Marshal(messages)
	utils.OkWithData(c, utils.OK, string(messagesByte))
}
