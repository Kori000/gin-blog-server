package ctype

import (
	"encoding/json"
)

type Role int

const (
	Admin        Role = 1  // 管理员
	User         Role = 2  // 普通用户
	Visitor      Role = 3  // 游客
	DisabledUser Role = -1 // 封停的账户
)

func (s Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Role) String() string {
	var str string
	switch s {
	case Admin:
		str = "管理员"
	case User:
		str = "普通用户"
	case Visitor:
		str = "游客"
	case DisabledUser:
		str = "封停的账户"
	default:
		str = "其他"
	}

	return str
}
