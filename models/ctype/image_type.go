package ctype

import (
	"encoding/json"
)

type ImageType int

const (
	Local ImageType = 1 //  本地服务器
	QiNiu ImageType = 2 //  七牛云
)

// 将数据序列化成 json 编码
func (s ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s ImageType) String() string {
	var str string
	switch s {
	case Local:
		str = "服务器"
	case QiNiu:
		str = "七牛云"
	default:
		str = "其他"
	}

	return str
}
