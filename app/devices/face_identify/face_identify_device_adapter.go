package face_identify

import (
	"context"
	"fmt"
	"github.com/titrxw/smart-home-server/app/device_manager/exception"
	deviceManagerLogic "github.com/titrxw/smart-home-server/app/device_manager/logic"
	deviceManagerModel "github.com/titrxw/smart-home-server/app/device_manager/model"
	"github.com/titrxw/smart-home-server/app/devices/face_identify/logic"
	"github.com/titrxw/smart-home-server/app/devices/face_identify/model"
	deviceInterface "github.com/titrxw/smart-home-server/app/devices/interface"
	"strconv"
	"time"
)

const OperateAddModelSettingMinImgLength = 8
const DeviceOperateAddModel = "add_face_model"
const DeviceOperateDelModel = "del_face_model"
const DeviceIdentifyReport = "identify"

type DeviceAdapter struct {
	deviceInterface.Abstract
}

func (a DeviceAdapter) sendFaceModelIdentifyEmail(user *deviceManagerModel.User, device *deviceManagerModel.Device, faceModel *model.FaceModel, faceUrl string) error {
	content := fmt.Sprintf(`
	<html><div>
		<div>
			您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您的设备 %s 于 %s 提交的人脸信息:%s，
				<img src="%s"/>
			</p>
		</div>
		<div>
			<p>此邮箱为系统邮箱，请勿回复。</p>
		</div>
	</div></html>
	`, device.Name, time.Now().Format(deviceManagerModel.TimeFormat), faceModel.UserName, faceUrl)

	return deviceManagerLogic.Logic.Email.SendEmail(user.Email, content)
}

func (a DeviceAdapter) GetDeviceConfig() deviceInterface.Device {
	return deviceInterface.Device{
		Type:           deviceInterface.DeviceAppType,
		TypeName:       "face_identify",
		Name:           "识别",
		NeedGateway:    false,
		SupportOperate: []string{DeviceOperateAddModel, DeviceOperateDelModel},
		OperateDesc:    map[string]string{DeviceOperateAddModel: "添加模型", DeviceOperateDelModel: "删除模型"},
		SupportReport:  []string{DeviceIdentifyReport},
		Setting: map[string]interface{}{
			DeviceOperateAddModel: map[string]interface{}{
				"min_img_length": OperateAddModelSettingMinImgLength,
			},
		},
	}
}

func (a DeviceAdapter) BeforeTriggerOperate(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, message *deviceInterface.DeviceOperateMessage) error {
	device := deviceManagerLogic.Logic.Device.GetDeviceByDeviceId(deviceAppId)
	if message.EventType == DeviceOperateAddModel {
		if _, ok := message.Payload["user_name"]; !ok {
			return exception.NewResponseError("user_name 参数缺失")
		}
		if _, ok := message.Payload["urls"]; !ok {
			return exception.NewResponseError("urls 参数缺失")
		}
		if _, ok := message.Payload["user_name"].(string); !ok {
			return exception.NewResponseError("user_name 参数格式错误")
		}
		if _, ok := message.Payload["urls"].([]interface{}); !ok {
			return exception.NewResponseError("urls 参数格式错误")
		}
		var urlsMap model.FaceUrls
		for _, param := range message.Payload["urls"].([]interface{}) {
			switch v := param.(type) {
			case string:
				urlsMap = append(urlsMap, v)
			default:
				return exception.NewResponseError("urls 参数格式错误")
			}
		}

		if len(urlsMap) < OperateAddModelSettingMinImgLength {
			return exception.NewResponseError("模型数量不能小于" + strconv.Itoa(OperateAddModelSettingMinImgLength))
		}

		faceModel, err := logic.FaceIdentifyDeviceLogic.FaceIdentify.AddDeviceFaceModel(device, message.Payload["user_name"].(string), urlsMap)
		if err != nil {
			return err
		}

		message.Payload["label"] = faceModel.ID
	}

	if message.EventType == DeviceOperateDelModel {
		if _, ok := message.Payload["label"]; !ok {
			return exception.NewResponseError("label 参数缺失")
		}
		if _, ok := message.Payload["label"].(int64); !ok {
			return exception.NewResponseError("label 参数格式错误")
		}

		if !logic.FaceIdentifyDeviceLogic.FaceIdentify.UpdateFaceModelStatus(device, uint(message.Payload["label"].(float64)), model.FaceModelStatusDisable) {
			return exception.NewResponseError("删除模型失败")
		}
	}

	return nil
}

func (a DeviceAdapter) AfterTriggerOperate(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, message *deviceInterface.DeviceOperateMessage) error {
	return nil
}

func (a DeviceAdapter) OnOperateResponse(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, operatePayLoad map[string]interface{}, message *deviceInterface.DeviceOperateMessage) error {
	if message.EventType == DeviceOperateAddModel {
		result, err := deviceManagerLogic.Logic.DeviceOperate.IsSuccessResponse(message.Payload)
		if err != nil {
			return err
		}

		if result {
			logic.FaceIdentifyDeviceLogic.FaceIdentify.UpdateFaceModelStatus(deviceManagerLogic.Logic.Device.GetDeviceByDeviceId(deviceAppId), uint(operatePayLoad["label"].(float64)), model.FaceModelStatusEnable)
		}
	}

	return nil
}

func (a DeviceAdapter) OnReport(ctx context.Context, gatewayDeviceAppId string, deviceAppId string, deviceType string, message *deviceInterface.DeviceOperateMessage) error {
	if message.EventType == DeviceIdentifyReport {
		if _, ok := message.Payload["label"]; !ok {
			return exception.NewResponseError("label 参数缺失")
		}
		if _, ok := message.Payload["label"].(int64); !ok {
			return exception.NewResponseError("label 参数格式错误")
		}
		if _, ok := message.Payload["mat"]; !ok {
			return exception.NewResponseError("mat 参数缺失")
		}
		if _, ok := message.Payload["mat"].(string); !ok {
			return exception.NewResponseError("mat 参数格式错误")
		}

		faceModel := logic.FaceIdentifyDeviceLogic.FaceIdentify.GetByLabel(uint(message.Payload["label"].(int64)))
		if faceModel == nil {
			return exception.NewResponseError("模型不存在")
		}
		if !faceModel.IsEnable() {
			return exception.NewResponseError("模型不可用")
		}

		device := deviceManagerLogic.Logic.Device.GetDeviceByDeviceId(deviceAppId)
		//处理其他业务
		user, err := deviceManagerLogic.Logic.User.GetUserById(ctx, device.UserId)
		if err != nil {
			return exception.NewResponseError("模型对应的用户不存在")
		}

		return a.sendFaceModelIdentifyEmail(user, device, faceModel, message.Payload["mat"].(string))
	}
	return nil
}
