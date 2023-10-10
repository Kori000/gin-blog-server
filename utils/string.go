package utils

func TruncateStr(str string, maxNum int) string {
	if len(str) <= maxNum {
		return str
	}

	return str[:maxNum]
}
