package webService

import (
	"log"
	"strings"

	"github.com/BurntSushi/toml"
)

type aPIS struct {
	Auth_code2Session string `toml:"auth_code2Session"`
}

type wxMiniAppRelatedConfig struct {
	AppID     string `toml:"AppID"`
	AppSecret string `toml:"AppSecret"`
	AAPIS     aPIS   `toml:"APIS"`
}

//package wise global variables
var (
	aWxMiniAppRelatedConfig wxMiniAppRelatedConfig
)

func init() {
	_, err := toml.DecodeFile("../config/WxMiniAppRelated.toml", &aWxMiniAppRelatedConfig)
	if err != nil {
		log.Fatal("Fail to load Wechat Mini-application related configurations.")
	}
	log.Println(aWxMiniAppRelatedConfig)
}

func (w *wxMiniAppRelatedConfig) makeAuth_code2SessionRequest(code string) string {
	request := strings.Replace(w.AAPIS.Auth_code2Session, "APPID", w.AppID, -1)
	request = strings.Replace(request, "SECRET", w.AppSecret, -1)
	return strings.Replace(request, "JSCODE", code, -1)
}
