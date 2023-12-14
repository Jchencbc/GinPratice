package api

import (
	"ginPratice/utils"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FileController struct {
}

func (con FileController) FileUpload(ctx *gin.Context) {
	var (
		uploadFileKey = "file"
		saveDir       = "static/files"
		saveName      = ""
		savePath      = ""
	)

	file, err := ctx.FormFile(uploadFileKey)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "上传文件失败",
		})
	}
	extstring := path.Ext(file.Filename)
	fileNameInt := time.Now().Unix()
	fileNameStr := strconv.FormatInt(fileNameInt, 10)
	saveName = fileNameStr + extstring

	//保存上传文件
	savePath = utils.Mkdir(saveDir) + "/" + saveName

	ctx.SaveUploadedFile(file, savePath)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
	})
}
