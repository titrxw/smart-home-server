package config

type Config struct {
	Server    Server            `mapstructure:"server" json:"server" yaml:"server"`
	Mqtt      Mqtt              `mapstructure:"mqtt" json:"mqtt" yaml:"mqtt"`
	Session   Session           `mapstructure:"session" json:"session" yaml:"session"`
	Common    Common            `mapstructure:"common" json:"common" yaml:"common"`
	Emqx      Emqx              `mapstructure:"emqx" json:"emqx" yaml:"emqx"`
	Jwt       Jwt               `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	DeviceMap map[string]Device `mapstructure:"device" json:"device" yaml:"device"`
}
