package flag

import (
	sys_flag "flag"
)

type Option struct {
	DB bool
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")

	// 解析命令行参数吸入注册的 Flag 里
	sys_flag.Parse()
	return Option{
		DB: *db,
	}
}

func IsWebStop(option Option) bool {
	if option.DB {
		return true
	} else {
		return false
	}
}

func SwitchOption(option Option) {
	if option.DB {
		MakeMigrations()
	}
}
