package faceIdentify

import (
	"context"
	"fmt"
	logic "github.com/titrxw/smart-home-server/app/Device/FaceIdentify/Logic"
	model2 "github.com/titrxw/smart-home-server/app/Device/FaceIdentify/Model"
	"github.com/titrxw/smart-home-server/app/Device/Interface"
	exception "github.com/titrxw/smart-home-server/app/Exception"
	logic2 "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
	"github.com/titrxw/smart-home-server/config"
	"strconv"
	"time"
)

const FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL_SETTING_MIN_IMG_LENGTH = 8
const FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL = "add_face_model"
const FACE_IDENTIFY_DEVICE_OPERATE_DEL_MODEL = "del_face_model"
const FACE_IDENTIFY_DEVICE_IDENTIFY_REPORT = "identify"

type FaceIdentifyDeviceAdapter struct {
	Interface.DeviceAdapterAbstract
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) sendFaceModelIdentifyEmail(user *model.User, device *model.Device, faceModel *model2.FaceModel, faceUrl string) error {
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
	`, device.Name, time.Now().Format(model.TimeFormat), faceModel.UserName, faceUrl)

	return logic2.Logic.EmailLogic.SendEmail(user.Email, content)
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) GetDeviceConfig() config.Device {
	return config.Device{
		Type:           model.DEVICE_APP_TYPE,
		TypeName:       "face_identify",
		Name:           "识别",
		NeedGateway:    false,
		SupportOperate: []string{FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL, FACE_IDENTIFY_DEVICE_OPERATE_DEL_MODEL},
		OperateDesc:    map[string]string{FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL: "添加模型", FACE_IDENTIFY_DEVICE_OPERATE_DEL_MODEL: "删除模型"},
		SupportReport:  []string{FACE_IDENTIFY_DEVICE_IDENTIFY_REPORT},
		Setting: map[string]interface{}{
			FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL: map[string]interface{}{
				"min_img_length": FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL_SETTING_MIN_IMG_LENGTH,
			},
		},
	}
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) BeforeTriggerOperate(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	if deviceOperateLog.OperateName == FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL {
		if _, ok := deviceOperateLog.OperatePayload["user_name"]; !ok {
			return exception.NewArgsError("user_name 参数缺失")
		}
		if _, ok := deviceOperateLog.OperatePayload["urls"]; !ok {
			return exception.NewArgsError("urls 参数缺失")
		}
		if _, ok := deviceOperateLog.OperatePayload["user_name"].(string); !ok {
			return exception.NewArgsError("user_name 参数格式错误")
		}
		if _, ok := deviceOperateLog.OperatePayload["urls"].([]interface{}); !ok {
			return exception.NewArgsError("urls 参数格式错误")
		}
		var urlsMap model2.FaceUrls
		for _, param := range deviceOperateLog.OperatePayload["urls"].([]interface{}) {
			switch v := param.(type) {
			case string:
				urlsMap = append(urlsMap, v)
			default:
				return exception.NewArgsError("urls 参数格式错误")
			}
		}

		if len(urlsMap) < FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL_SETTING_MIN_IMG_LENGTH {
			return exception.NewLogicError("模型数量不能小于" + strconv.Itoa(FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL_SETTING_MIN_IMG_LENGTH))
		}

		faceModel, err := logic.FaceIdentifyDeviceLogic.FaceIdentifyLogic.AddDeviceFaceModel(device, deviceOperateLog.OperatePayload["user_name"].(string), urlsMap)
		if err != nil {
			return err
		}

		deviceOperateLog.OperatePayload["label"] = faceModel.ID
	}

	if deviceOperateLog.OperateName == FACE_IDENTIFY_DEVICE_OPERATE_DEL_MODEL {
		if _, ok := deviceOperateLog.OperatePayload["label"]; !ok {
			return exception.NewArgsError("label 参数缺失")
		}
		if _, ok := deviceOperateLog.OperatePayload["label"].(int64); !ok {
			return exception.NewArgsError("label 参数格式错误")
		}

		if !logic.FaceIdentifyDeviceLogic.FaceIdentifyLogic.UpdateFaceModelStatus(device, uint(deviceOperateLog.OperatePayload["label"].(float64)), model2.FACE_MODEL_STATUS_DISABLE) {
			return exception.NewLogicError("删除模型失败")
		}
	}

	return nil
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) AfterTriggerOperate(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) OnOperateResponse(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceOperateLog *model.DeviceOperateLog, message *model.IotMessage) error {
	if deviceOperateLog.OperateName == FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL {
		result, err := logic2.Logic.DeviceOperateLogic.IsSuccessResponse(deviceOperateLog.ResponsePayload)
		if err != nil {
			return err
		}

		if result {
			logic.FaceIdentifyDeviceLogic.FaceIdentifyLogic.UpdateFaceModelStatus(device, uint(deviceOperateLog.OperatePayload["label"].(float64)), model2.FACE_MODEL_STATUS_ENABLE)
		}
	}

	return nil
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) OnReport(ctx context.Context, gatewayDevice *model.Device, device *model.Device, deviceReportLog *model.DeviceReportLog, message *model.IotMessage) error {
	if deviceReportLog.ReportName == FACE_IDENTIFY_DEVICE_IDENTIFY_REPORT {
		if _, ok := deviceReportLog.ReportPayload["label"]; !ok {
			return exception.NewArgsError("label 参数缺失")
		}
		if _, ok := deviceReportLog.ReportPayload["label"].(int64); !ok {
			return exception.NewArgsError("label 参数格式错误")
		}
		if _, ok := deviceReportLog.ReportPayload["mat"]; !ok {
			return exception.NewArgsError("mat 参数缺失")
		}
		if _, ok := deviceReportLog.ReportPayload["mat"].(string); !ok {
			return exception.NewArgsError("mat 参数格式错误")
		}

		faceModel := logic.FaceIdentifyDeviceLogic.FaceIdentifyLogic.GetByLabel(uint(deviceReportLog.ReportPayload["label"].(int64)))
		if faceModel == nil {
			return exception.NewLogicError("模型不存在")
		}
		if !faceModel.IsEnable() {
			return exception.NewLogicError("模型不可用")
		}

		//处理其他业务
		user, err := logic2.Logic.UserLogic.GetUserById(ctx, device.UserId)
		if err != nil {
			return exception.NewLogicError("模型对应的用户不存在")
		}

		return faceIdentifyDeviceAdapter.sendFaceModelIdentifyEmail(user, device, faceModel, deviceReportLog.ReportPayload["mat"].(string))
	}
	return nil
}
