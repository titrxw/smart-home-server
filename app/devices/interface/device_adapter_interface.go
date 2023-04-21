package _interface

import (
	"context"
)

type DeviceInterface interface {
	GetDeviceConfig() Device
	GetCtrlTopic(appId string, componentAppId string) string
	GetReportTopic(appId string) string
	GetAvailabilityTopic(appId string) string
	BeforeTriggerOperate(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, message *DeviceOperateMessage) error
	AfterTriggerOperate(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, message *DeviceOperateMessage) error
	OnOperateResponse(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, operatePayLoad map[string]interface{}, message *DeviceOperateMessage) error
	OnReport(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, message *DeviceOperateMessage) error
	PackMessage(message *DeviceOperateMessage) (string, error)
	UnPackMessage(payload string) (*DeviceOperateMessage, error)
}
