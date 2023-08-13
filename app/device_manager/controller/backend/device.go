package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/internal/device/manager"
	deviceInterface "github.com/titrxw/smart-home-server/app/pkg/device"
	"github.com/we7coreteam/w7-rangine-go/src/http/controller"
)

type Device struct {
	controller.Abstract
}

func (c Device) AddSupportDevice(ctx *gin.Context) {
	deviceAddRequest := deviceInterface.Device{}
	if !c.Validate(ctx, &deviceAddRequest) {
		return
	}

	manager.RegisterDevice(deviceAddRequest)

	c.JsonSuccessResponse(ctx)
}
