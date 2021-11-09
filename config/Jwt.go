package config

type Jwt struct {
	PrivateKey      string `mapstructure:"private_key" json:"private_key" yaml:"private_key"`
	PublicKey       string `mapstructure:"public_key" json:"public_key" yaml:"public_key"`
	Iss             string `mapstructure:"iss" json:"iss" yaml:"iss"`
	Subject         string `mapstructure:"subject" json:"subject" yaml:"subject"`
	NotBeforeSecond int64  `mapstructure:"nbf_second" json:"nbf_second" yaml:"nbf_second"`
	TTL             int64  `mapstructure:"ttl" json:"ttl" yaml:"ttl"`
}
