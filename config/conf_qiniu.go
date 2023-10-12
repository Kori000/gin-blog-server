package config

type QiNiu struct {
	Enable    bool    `yaml:"enable" json:"enable" default:"true"` // 是否启用
	AccessKey string  `yaml:"access_key" json:"access_key" default:"DSiVMvhK1idyOSeVpu4XBu0WR5vDKjNjosvGGGsA"`
	SecretKey string  `yaml:"secret_key" json:"secret_key" default:"oDJjQW4psIwzUvA4q-XyVSs1kBrrXtzlPWE5HeAR"`
	Bucket    string  `yaml:"bucket" json:"bucket" default:"kori-foreign"` // 存储桶名称
	CDN       string  `yaml:"cdn" json:"cdn" default:"123"`                // 访问的地址前缀
	Zone      string  `yaml:"zone" json:"zone" default:"213"`              // 存储的地区
	Size      float64 `yaml:"size" json:"size" default:"0"`                // 存储的大小限制, 单位 MB
}
