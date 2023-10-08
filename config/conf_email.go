package config

type Email struct {
	Host             string `yaml:"host" json:"host"`
	Port             int    `yaml:"port" json:"port"`
	User             string `yaml:"user" json:"user"` // 发件人邮箱
	Password         string `yaml:"password" json:"password"`
	DefaultFromEmail string `yaml:"default_from_email" json:"default_from_email"` // 默认的发件人名字
	UseSSL           bool   `yaml:"use_ssl" json:"use_ssl"`                       // 是否使用 ssl
	UserTls          bool   `yaml:"user_tls" json:"user_tls"`                     //
}
