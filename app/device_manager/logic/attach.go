package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Attach struct {
	Abstract
}

func (l *Attach) GetImgAttachDir() (string, error) {
	dir := "./public/upload/img/"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", exception.NewResponseError("创建文件夹失败")
	}

	return dir, nil
}

func (l *Attach) GetAppImgAttachPath(appId string, ext string) (string, error) {
	dir, err := l.GetImgAttachDir()
	if err != nil {
		return "", err
	}

	dir = dir + appId + "/" + time.Now().Format("20060102")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", exception.NewResponseError("创建文件夹失败")
	}

	fileUnixName := strconv.FormatInt(time.Now().UnixMicro(), 10) + ext

	return path.Join(dir, fileUnixName), nil
}

func (l *Attach) GetRemoteFileUrlByAttachPath(filePath string) string {
	return strings.NewReplacer("public/upload", "").Replace(filePath)
}

func (l *Attach) SaveUploadAttach(ctx *gin.Context, file *multipart.FileHeader, fileSavePath string) error {
	return ctx.SaveUploadedFile(file, fileSavePath)
}
