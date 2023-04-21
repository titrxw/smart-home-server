package _interface

const DeviceAppType uint8 = 1
const DeviceGatewayAppType uint8 = 2

type Device struct {
	Type           uint8                  `mapstructure:"type" json:"type" yaml:"type"`
	TypeName       string                 `mapstructure:"type_name" json:"type_name" yaml:"type_name"`
	Name           string                 `mapstructure:"name" json:"name" yaml:"name"`
	NeedGateway    bool                   `mapstructure:"need_gateway" json:"need_gateway" yaml:"need_gateway"`
	SupportOperate []string               `mapstructure:"support_operate" json:"support_operate" yaml:"support_operate"`
	OperateDesc    map[string]string      `mapstructure:"operate_desc" json:"operate_desc" yaml:"operate_desc"`
	SupportReport  []string               `mapstructure:"support_report" json:"support_report" yaml:"support_report"`
	Setting        map[string]interface{} `mapstructure:"setting" json:"setting" yaml:"setting"`
}
