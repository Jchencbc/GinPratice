package api

import (
	"ginPratice/utils"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
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

func (con FileController) FileDownload(ctx *gin.Context) {
	savePath := ctx.Query("file_path")
	_, err := os.Open(savePath)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "文件地址不存在",
		})
	} else {
		fileParts := strings.Split(savePath, "\\")
		fileName := fileParts[len(fileParts)-1]
		ctx.Header("Content-Disposition", "attachment;filename="+fileName)
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(savePath)
	}
}
