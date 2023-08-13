package device

import (
	"github.com/gin-gonic/gin"
	"github.com/titrxw/smart-home-server/app/device_manager/controller/frontend/frontend"
	"github.com/titrxw/smart-home-server/app/internal/logic"
)

type ReportLog struct {
	frontend.Abstract
}

type ReportPageRequest struct {
	DeviceId uint `form:"device_id" binding:"required,id"`
	Page     uint `form:"page" binding:"required,page"`
	PageSize uint `form:"page_size" binding:"required,page"`
}

type ReportDetailRequest struct {
	DeviceId     uint   `form:"device_id" binding:"required,id"`
	ReportNumber string `form:"report_number" binding:"required"`
}

func (c ReportLog) DeviceReportLog(ctx *gin.Context) {
	deviceReportPageRequest := ReportPageRequest{}
	if !c.Validate(ctx, &deviceReportPageRequest) {
		return
	}

	device, err := logic.Logic.Device.GetUserDeviceById(c.GetUserId(ctx), deviceReportPageRequest.DeviceId)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	pageData := logic.Logic.DeviceReport.GetDeviceReports(device, deviceReportPageRequest.Page, deviceReportPageRequest.PageSize)

	c.JsonResponseWithoutError(ctx, pageData)
}

func (c ReportLog) ReportDetail(ctx *gin.Context) {
	deviceReportDetailRequest := ReportDetailRequest{}
	if !c.Validate(ctx, &deviceReportDetailRequest) {
		return
	}

	device, err := logic.Logic.Device.GetUserDeviceById(c.GetUserId(ctx), deviceReportDetailRequest.DeviceId)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	reportLog, err := logic.Logic.DeviceReport.GetDeviceReportLogByNumber(device, deviceReportDetailRequest.ReportNumber)
	if err != nil {
		c.JsonResponseWithServerError(ctx, err)
		return
	}

	c.JsonResponseWithoutError(ctx, reportLog)
}
