package util

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/internal/logic"
	"github.com/titrxw/smart-home-server/app/internal/model"
	errorhandler "github.com/we7coreteam/w7-rangine-go/src/core/err_handler"
	"github.com/we7coreteam/w7-rangine-go/src/http/controller"
	"path"
)

type Upload struct {
	controller.Abstract
}

func (c Upload) UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
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
		c.JsonResponseWithServerError(ctx, errorhandler.Throw("文件类型不合法", nil))
		return
	}

	path, err := logic.Logic.Attach.GetAppImgAttachPath(ctx.MustGet("app").(*model.App).AppId, extName)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	err = logic.Logic.Attach.SaveUploadAttach(ctx, file, path)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonResponseWithoutError(ctx, gin.H{
		"url": logic.Logic.Attach.GetRemoteFileUrlByAttachPath(path),
	})
}
