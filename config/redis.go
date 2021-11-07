package config

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Db       int    `mapstructure:"db" json:"db" yaml:"db"`
}
