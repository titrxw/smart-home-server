package config

type Emqx struct {
	Host      string `mapstructure:"host" json:"host" yaml:"host"`
	AppId     string `mapstructure:"app_id" json:"app_id" yaml:"app_id"`
	AppSecret string `mapstructure:"app_secret" json:"app_secret" yaml:"app_secret"`
}
