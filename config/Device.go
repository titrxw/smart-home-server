package config

type Device struct {
	Name           string   `mapstructure:"name" json:"name" yaml:"name"`
	SupportOperate []string `mapstructure:"support_operate" json:"support_operate" yaml:"support_operate"`
}
