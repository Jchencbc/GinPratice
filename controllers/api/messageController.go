package api

import (
	"fmt"
	"ginPratice/utils"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
}

func (msg MessageController) MsgUp(c *gin.Context) {
	formData, _ := c.MultipartForm()
	pictures := formData.File["pictures"]
	paths, err := utils.FileUploadTools(pictures)
	if err != nil {

	}
	fmt.Println(paths) //todo
}
