package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"lingotalk-exam/config/model"
)

// JSONConfig model.JSONConfig
var JSONConfig model.JSONConfig

func init() {
	goEnv := os.Getenv("GO_ENV")
	wd, _ := os.Getwd()
	if goEnv == "" {
		goEnv = "test"
		gp := os.Getenv("GOPATH")
		wd = path.Join(gp, "src/lingotalk-exam")
	}
	configFilePath := fmt.Sprintf(wd+"/config/json/config.%s.json", goEnv)
	file, ioError := ioutil.ReadFile(configFilePath)
	if ioError != nil {
		fmt.Println(ioError.Error())
	}
	JSONConfig = model.JSONConfig{}
	jsonError := json.Unmarshal(file, &JSONConfig)
	if jsonError != nil {
		fmt.Println(jsonError.Error())
	}
}
