package device

const DeviceAppType uint8 = 1
const DeviceGatewayAppType uint8 = 2

type Device struct {
	Type              uint8                  `mapstructure:"type" form:"type" json:"type" yaml:"type" binding:"required"`
	TypeName          string                 `mapstructure:"type_name" form:"type_name" json:"type_name" yaml:"type_name" binding:"required"`
	Name              string                 `mapstructure:"name" form:"name" json:"name" yaml:"name" binding:"required"`
	NeedGateway       bool                   `mapstructure:"need_gateway" form:"need_gateway" json:"need_gateway" yaml:"need_gateway" binding:"required"`
	SupportOperate    []string               `mapstructure:"support_operate" form:"support_operate" json:"support_operate" yaml:"support_operate"`
	OperateDesc       map[string]string      `mapstructure:"operate_desc" form:"operate_desc" json:"operate_desc" yaml:"operate_desc"`
	SupportReport     []string               `mapstructure:"support_report" form:"support_report" json:"support_report" yaml:"support_report"`
	Setting           map[string]interface{} `mapstructure:"setting" form:"setting" json:"setting" yaml:"setting"`
	CtrlTopic         string                 `mapstructure:"ctrl_topic" form:"ctrl_topic" json:"ctrl_topic" yaml:"ctrl_topic"`
	ReportTopic       string                 `mapstructure:"report_topic" form:"report_topic" json:"report_topic" yaml:"report_topic"`
	AvailabilityTopic string                 `mapstructure:"availability_topic" form:"availability_topic" json:"availability_topic" yaml:"availability_topic"`
}
