package util

import (
	"github.com/gin-gonic/gin"
	base "github.com/titrxw/smart-home-server/app/Controller/Base"
	logic "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
	"path"
)

type UploadController struct {
	base.ControllerAbstract
}

func (uploadController UploadController) UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		uploadController.JsonResponseWithServerError(ctx, err)
		return
	}

	extName := path.Ext(file.Filename)
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		uploadController.JsonResponseWithServerError(ctx, "文件类型不合法")
		return
	}

	path, err := logic.Logic.AppLogic.GetAppAttachPath(ctx.MustGet("app").(*model.App).AppId, extName)
	if err != nil {
		uploadController.JsonResponseWithServerError(ctx, err)
		return
	}

	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		uploadController.JsonResponseWithServerError(ctx, err)
		return
	}

	uploadController.JsonResponseWithoutError(ctx, gin.H{
		"url": logic.Logic.AppLogic.GetRemoteFileUrlByAttachPath(path),
	})
}
