package config

type Device struct {
	Type           string                 `mapstructure:"type" json:"type" yaml:"type"`
	Name           string                 `mapstructure:"name" json:"name" yaml:"name"`
	SupportOperate []string               `mapstructure:"support_operate" json:"support_operate" yaml:"support_operate"`
	OperateDesc    map[string]string      `mapstructure:"operate_desc" json:"operate_desc" yaml:"operate_desc"`
	SupportReport  []string               `mapstructure:"support_report" json:"support_report" yaml:"support_report"`
	Setting        map[string]interface{} `mapstructure:"setting" json:"setting" yaml:"setting"`
}
