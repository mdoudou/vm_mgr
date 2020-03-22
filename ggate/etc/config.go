package etc

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Controller ControllerData `json:controller`
}

type ControllerData struct {
	IP string `json:"ip"`
}

func GetControllerIP() string {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	var config Config
	json.Unmarshal(file, &config)
	return config.Controller.IP
}
