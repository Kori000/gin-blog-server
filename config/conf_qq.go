package config

import "fmt"

type QQ struct {
	AppID    string `yaml:"app_id" json:"app_id" default:"123"`
	Key      string `yaml:"key" json:"key" default:"123"`
	Redirect string `yaml:"redirect" json:"redirect" default:"123"` // 重定向地址
}

// 生成 Oauth url
func (q *QQ) GetPath() string {
	if q.Key == "" || q.AppID == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_uri=%s", q.AppID, q.Redirect)
}
