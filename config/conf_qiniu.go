package config

type QiNiu struct {
	AccessKey string  `yaml:"access_key" json:"access_key"`
	SecretKey string  `yaml:"secret_key" json:"secret_key"`
	Bucket    string  `yaml:"bucket" json:"bucket"` // 存储桶名称
	CDN       string  `yaml:"cdn" json:"cdn"`       // 访问的地址前缀
	Zone      string  `yaml:"zone" json:"zone"`     // 存储的地区
	Size      float64 `yaml:"size" json:"size"`     // 存储的大小限制, 单位 MB
}
