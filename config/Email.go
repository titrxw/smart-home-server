package config

type Email struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         string `mapstructure:"port" json:"port" yaml:"port"`
	Identify     string `mapstructure:"identify" json:"identify" yaml:"identify"`
	FromUserName string `mapstructure:"from_user_name" json:"from_user_name" yaml:"from_user_name"`
	UserName     string `mapstructure:"user_name" json:"user_name" yaml:"user_name"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
}
