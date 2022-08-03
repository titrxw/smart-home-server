package device

import (
	"github.com/gin-gonic/gin"
	base "github.com/titrxw/smart-home-server/app/Controller/Base"
	frontend "github.com/titrxw/smart-home-server/app/Controller/Frontend/Frontend"
	logic "github.com/titrxw/smart-home-server/app/Logic"
)

type DeviceReportLogController struct {
	frontend.ControllerAbstract
}

type DeviceReportDetailRequest struct {
	base.RequestAbstract
	DeviceId     uint   `form:"device_id" binding:"required,id"`
	ReportNumber string `form:"report_number" binding:"required"`
}

type DeviceReportPageRequest struct {
	base.RequestAbstract
	DeviceId uint `form:"device_id" binding:"required,id"`
	Page     uint `form:"page" binding:"required,page"`
	PageSize uint `form:"page_size" binding:"required,page"`
}

func (deviceReportController DeviceReportLogController) ReportDetail(ctx *gin.Context) {
	deviceReportDetailRequest := DeviceReportDetailRequest{}
	if !deviceReportController.ValidateFormPost(ctx, &deviceReportDetailRequest) {
		return
	}

	device, err := logic.Logic.DeviceLogic.GetUserDeviceById(deviceReportController.GetUserId(ctx), deviceReportDetailRequest.DeviceId)
	if err != nil {
		deviceReportController.JsonResponseWithServerError(ctx, err)
		return
	}

	reportLog, err := logic.Logic.DeviceReportLogic.GetDeviceReportLogByNumber(device, deviceReportDetailRequest.ReportNumber)
	if err != nil {
		deviceReportController.JsonResponseWithServerError(ctx, err)
		return
	}

	deviceReportController.JsonResponseWithoutError(ctx, reportLog)
}

func (deviceReportController DeviceReportLogController) DeviceReportLog(ctx *gin.Context) {
	deviceReportPageRequest := DeviceReportPageRequest{}
	if !deviceReportController.ValidateFormPost(ctx, &deviceReportPageRequest) {
		return
	}

	device, err := logic.Logic.DeviceLogic.GetUserDeviceById(deviceReportController.GetUserId(ctx), deviceReportPageRequest.DeviceId)
	if err != nil {
		deviceReportController.JsonResponseWithServerError(ctx, err)
		return
	}

	pageData := logic.Logic.DeviceReportLogic.GetDeviceReports(device, deviceReportPageRequest.Page, deviceReportPageRequest.PageSize)

	deviceReportController.JsonResponseWithoutError(ctx, pageData)
}
