package logic

import (
	"github.com/gin-gonic/gin"
	exception "github.com/titrxw/smart-home-server/app/Exception"
	"mime/multipart"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type AttachLogic struct {
	LogicAbstract
}

func (attachLogic *AttachLogic) GetImgAttachDir() (string, error) {
	dir := "./public/upload/img/"
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", exception.NewLogicError("创建文件夹失败")
	}

	return dir, nil
}

func (attachLogic *AttachLogic) GetAppImgAttachPath(appId string, ext string) (string, error) {
	dir, err := attachLogic.GetImgAttachDir()
	if err != nil {
		return "", err
	}

	dir = dir + appId + "/" + time.Now().Format("20060102")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return "", exception.NewLogicError("创建文件夹失败")
	}

	fileUnixName := strconv.FormatInt(time.Now().UnixMicro(), 10) + ext

	return path.Join(dir, fileUnixName), nil
}

func (attachLogic *AttachLogic) GetRemoteFileUrlByAttachPath(filePath string) string {
	return strings.NewReplacer("public/upload", "").Replace(filePath)
}

func (attachLogic *AttachLogic) SaveUploadAttach(ctx *gin.Context, file *multipart.FileHeader, fileSavePath string) error {
	return ctx.SaveUploadedFile(file, fileSavePath)
}
