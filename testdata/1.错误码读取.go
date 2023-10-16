package testdata

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"gin-blog-server/models/res"
)

const file = "models/res/err_code.json"

type ErrMap map[res.ErrorCode]string

func main() {
	byteData, err := os.ReadFile(file)
	if err != nil {
		logrus.Error(err)
		return
	}

	var errMap = ErrMap{}
	err = json.Unmarshal(byteData, &errMap)
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println(errMap[res.SettingsError])

}
