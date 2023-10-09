package flag

import (
	sys_flag "flag"
	"os"

	"gin-blog-server/global"
)

type Option struct {
	DB         bool
	RemovePath string
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	re := sys_flag.String("re", "", "删除文件路径")

	// 解析命令行参数写入注册的 Flag 里
	sys_flag.Parse()
	return Option{
		DB:         *db,
		RemovePath: *re,
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

func RemoveFile(option Option) {
	return
	if option.RemovePath != "" {
		global.Log.Infof("路径是 %v", option.RemovePath)

		err := os.Remove(option.RemovePath)
		if err != nil {
			global.Log.Errorf(err.Error())
		}
	}

}
