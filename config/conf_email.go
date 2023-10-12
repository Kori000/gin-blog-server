package config

type Email struct {
	Host             string `yaml:"host" json:"host" default:"1"`
	Port             int    `yaml:"port" json:"port" default:"0"`
	User             string `yaml:"user" json:"user" default:"23"` // 发件人邮箱
	Password         string `yaml:"password" json:"password" default:"23"`
	DefaultFromEmail string `yaml:"default_from_email" json:"default_from_email" default:"23"` // 默认的发件人名字
	UseSSL           bool   `yaml:"use_ssl" json:"use_ssl" default:"false"`                    // 是否使用 ssl
	UserTls          bool   `yaml:"user_tls" json:"user_tls" default:"false"`
}
