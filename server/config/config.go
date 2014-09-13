package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Pass       string `json:"pass"`
	PublicPath string `json:"public_path"`
}

var C Config

func init() {
	conf, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(conf, &C)
}
