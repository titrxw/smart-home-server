package faceIdentify

import (
	"errors"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	logic "github.com/titrxw/smart-home-server/app/Device/FaceIdentify/Logic"
	model2 "github.com/titrxw/smart-home-server/app/Device/FaceIdentify/Model"
	"github.com/titrxw/smart-home-server/app/Device/Interface"
	logic2 "github.com/titrxw/smart-home-server/app/Logic"
	model "github.com/titrxw/smart-home-server/app/Model"
	"github.com/titrxw/smart-home-server/config"
	"strconv"
)

const FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL_SETTING_MIN_IMG_LENGTH = 8
const FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL = "add_face_model"
const FACE_IDENTIFY_DEVICE_OPERATE_DEL_MODEL = "del_face_model"
const FACE_IDENTIFY_DEVICE_IDENTIFY_REPORT = "identify"

type FaceIdentifyDeviceAdapter struct {
	Interface.DeviceAdapterInterface
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) GetDeviceConfig() config.Device {
	return config.Device{
		Type:           "face_identify",
		Name:           "识别",
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

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) BeforeTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	if deviceOperateLog.OperateName == FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL {
		if _, ok := deviceOperateLog.OperatePayload["user_name"]; !ok {
			return errors.New("user_name 参数缺失")
		}
		if _, ok := deviceOperateLog.OperatePayload["urls"]; !ok {
			return errors.New("urls 参数缺失")
		}
		if _, ok := deviceOperateLog.OperatePayload["user_name"].(string); !ok {
			return errors.New("user_name 参数格式错误")
		}
		if _, ok := deviceOperateLog.OperatePayload["urls"].(model2.FaceUrls); !ok {
			return errors.New("urls 参数格式错误")
		}
		if len(deviceOperateLog.OperatePayload["urls"].(model2.FaceUrls)) < FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL_SETTING_MIN_IMG_LENGTH {
			return errors.New("模型数量不能小于" + strconv.Itoa(FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL_SETTING_MIN_IMG_LENGTH))
		}

		faceModel, err := logic.FaceIdentifyDeviceLogic.FaceIdentifyLogic.AddDeviceFaceModel(device, deviceOperateLog.OperatePayload["user_name"].(string), deviceOperateLog.OperatePayload["urls"].(model2.FaceUrls))
		if err != nil {
			return err
		}

		deviceOperateLog.OperatePayload["label"] = faceModel.ID
	}

	if deviceOperateLog.OperateName == FACE_IDENTIFY_DEVICE_OPERATE_DEL_MODEL {
		if _, ok := deviceOperateLog.OperatePayload["label"]; !ok {
			return errors.New("label 参数缺失")
		}
		if _, ok := deviceOperateLog.OperatePayload["label"].(uint); !ok {
			return errors.New("label 参数格式错误")
		}

		if !logic.FaceIdentifyDeviceLogic.FaceIdentifyLogic.UpdateFaceModelStatus(device, deviceOperateLog.OperatePayload["label"].(uint), model2.FACE_MODEL_STATUS_DISABLE) {
			return errors.New("删除模型失败")
		}
	}

	return nil
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) AfterTriggerOperate(device *model.Device, deviceOperateLog *model.DeviceOperateLog) error {
	return nil
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) OnOperateResponse(device *model.Device, deviceOperateLog *model.DeviceOperateLog, cloudEvent *cloudevents.Event) error {
	if deviceOperateLog.OperateName == FACE_IDENTIFY_DEVICE_OPERATE_ADD_MODEL {
		result, err := logic2.Logic.DeviceOperateLogic.IsSuccessResponse(deviceOperateLog.ResponsePayload)
		if err != nil {
			return err
		}

		if result {
			logic.FaceIdentifyDeviceLogic.FaceIdentifyLogic.UpdateFaceModelStatus(device, deviceOperateLog.ResponsePayload["label"].(uint), model2.FACE_MODEL_STATUS_ENABLE)
		}
	}

	return nil
}

func (faceIdentifyDeviceAdapter FaceIdentifyDeviceAdapter) OnReport(device *model.Device, deviceReportLog *model.DeviceReportLog, cloudEvent *cloudevents.Event) error {
	if deviceReportLog.ReportName == FACE_IDENTIFY_DEVICE_IDENTIFY_REPORT {
		if _, ok := deviceReportLog.ReportPayload["label"]; !ok {
			return errors.New("label 参数缺失")
		}
		if _, ok := deviceReportLog.ReportPayload["label"].(uint); !ok {
			return errors.New("label 参数格式错误")
		}

		faceModel := logic.FaceIdentifyDeviceLogic.FaceIdentifyLogic.GetByLabel(deviceReportLog.ReportPayload["label"].(uint))
		if faceModel == nil {
			return errors.New("模型不存在")
		}
		if !faceModel.IsEnable() {
			return errors.New("模型不可用")
		}

		//处理其他业务
	}
	return nil
}
