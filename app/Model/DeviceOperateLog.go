package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type OperatePayload map[string]interface{}

func (payload OperatePayload) Value() (driver.Value, error) {
	if payload == nil {
		return nil, nil
	}

	return json.Marshal(payload)
}

func (payload *OperatePayload) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}
	if len(b) == 0 {
		return nil
	}

	return json.Unmarshal(b, &payload)
}

type DeviceOperateLog struct {
	Model

	DeviceId       uint           `json:"-" gorm:"not null"`
	DeviceType     string         `json:"device_type" gorm:"type:varchar(64);not null"`
	Source         string         `json:"source" gorm:"type:varchar(12);not null"`
	OperateName    string         `json:"operate_name" gorm:"type:varchar(64);not null"`
	OperateNumber  string         `json:"operate_number"  gorm:"type:varchar(64);not null"`
	OperateTime    LocalTime      `json:"operate_time"`
	OperatePayload OperatePayload `json:"operate_payload" gorm:"type:varchar(500);not null;default:''"`
	OperateLevel   uint8          `json:"operate_level" gorm:"not null;default:0"`

	ResponsePayload OperatePayload `json:"response_payload" gorm:"type:varchar(500);not null;default:''"`
	ResponseIp      string         `json:"response_ip" gorm:"type:varchar(24);not null;default:''"`
	ResponsePort    string         `json:"response_port" gorm:"type:varchar(12);not null;default:''"`
	ResponseTime    string         `json:"response_time" gorm:"type:varchar(24);not null;default:''"`

	CreatedAt LocalTime `json:"created_at"`
}
