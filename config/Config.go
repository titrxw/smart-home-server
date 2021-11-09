package config

type Config struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Server   Server   `mapstructure:"server" json:"server" yaml:"server"`
	Database Database `mapstructure:"database" json:"database" yaml:"database"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Emqx     Emqx     `mapstructure:"emqx" json:"emqx" yaml:"emqx"`
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
}
