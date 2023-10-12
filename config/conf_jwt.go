package config

type JWT struct {
	Secret  string `yaml:"secret" json:"secret" default:"213"` // 秘钥
	Expires int    `yaml:"expires" json:"expires" default:"0"` // 过期时间
	Issuer  string `yaml:"issuer" json:"issuer" default:"213"` // 签发者
}
