package config

type JWT struct {
	Secret  string `yaml:"secret" json:"secret"`   // 秘钥
	Expires int    `yaml:"expires" json:"expires"` // 过期时间
	Issuer  string `yaml:"issuer" json:"issuer"`   // 签发者
}
