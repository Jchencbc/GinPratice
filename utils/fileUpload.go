package utils

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func FileUploadTools(files []*multipart.FileHeader) ([]string, error) {
	filePaths := []string{}
	saveDir := "static/files"
	for _, file := range files {
		fileNameInt := time.Now().Unix()
		fileNameStr := strconv.FormatInt(fileNameInt, 10)
		saveName := fileNameStr + file.Filename
		filePaths = append(filePaths, saveName)
		//保存上传文件
		savePath := Mkdir(saveDir) + "/" + saveName
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()
		out, err := os.Create(savePath)
		if err != nil {
			return nil, err
		}
		defer out.Close()
		_, err = io.Copy(out, src)
		return nil, err
	}
	return filePaths, nil
}
