package utils

import "gin-blog-server/global"

func Contains(list []string, value string) bool {
	for _, item := range list {
		global.Log.Info(item, value)
		if item == value {
			return true
		}
	}
	return false
}
