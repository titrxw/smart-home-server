package config

type Session struct {
	DbConnection string `mapstructure:"db_connection" json:"db_connection" yaml:"db_connection"`
}
