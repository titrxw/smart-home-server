package config

type Mqtt struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	UserName string `mapstructure:"user_name" json:"user_name" yaml:"user_name"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}
