package device

import (
	"context"
	"github.com/titrxw/smart-home-server/app/internal/http"
	"github.com/titrxw/smart-home-server/app/pkg/exception"
)

func IsSuccessResponse(operatePayload map[string]interface{}) (bool, error) {
	if _, ok := operatePayload["status"]; !ok {
		return false, exception.NewResponseError("status 参数缺失")
	}
	if _, ok := operatePayload["status"].(string); !ok {
		return false, exception.NewResponseError("status 参数格式错误")
	}

	if operatePayload["status"] == "success" {
		return true, nil
	}

	return false, nil
}

func TriggerOperate(ctx context.Context, serverUrl string, appId string, appSecret string, UserId uint, deviceAppId string, operateType string, payload map[string]interface{}) error {
	pushData := map[string]interface{}{
		"user_id":         UserId,
		"device_appid":    deviceAppId,
		"operate_type":    operateType,
		"operate_payload": payload,
	}
	_, err := http.PostWithAppSignByWWWForm(ctx, appId, appSecret, serverUrl, pushData, nil)
	return err
}
