package model

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ReportPayload map[string]interface{}

func (payload ReportPayload) Value() (driver.Value, error) {
	if payload == nil {
		return nil, nil
	}

	return json.Marshal(payload)
}

func (payload *ReportPayload) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("value is not []byte, value: %v", value)
	}
	if len(b) == 0 {
		return nil
	}

	return json.Unmarshal(b, &payload)
}

type DeviceReportLog struct {
	Model

	DeviceId        uint          `json:"-" gorm:"not null"`
	DeviceGatewayId uint          `json:"-" gorm:"not null;default:0"`
	DeviceType      string        `json:"device_type" gorm:"type:varchar(64);not null"`
	Source          string        `json:"source" gorm:"type:varchar(12);not null"`
	ReportName      string        `json:"report_name" gorm:"type:varchar(64);not null"`
	ReportNumber    string        `json:"report_number"  gorm:"type:varchar(64);not null"`
	ReportTime      LocalTime     `json:"report_time"`
	ReportPayload   ReportPayload `json:"report_payload" gorm:"type:varchar(500);not null;default:''"`
	ReportLevel     uint8         `json:"report_level" gorm:"not null;default:0"`
	ReportIp        string        `json:"report_ip" gorm:"type:varchar(24);not null;default:''"`
	ReportPort      string        `json:"report_port" gorm:"type:varchar(12);not null;default:''"`
	CreatedAt       LocalTime     `json:"created_at"`
}
