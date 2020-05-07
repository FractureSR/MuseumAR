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

type baiduImageApiRelatedConfig struct {
	GetAcessTokenApi string                  `toml:"get_access_token_api"`
	Recognize        baiduImageApi_recognize `toml:"Recognize"`
	Search           baiduImageApi_search    `toml:"Search"`
}

type baiduImageApi_recognize struct {
	ApiKey    string `toml:"api_key"`
	SecretKey string `toml:"secret_key"`
	WithFace  int    `toml:"with_face"`
	Api       string `toml:"api"`
}

type baiduImageApi_search struct {
	ApiKey    string `toml:"api_key"`
	SecretKey string `toml:"secret_key"`
	PageNum   int    `toml:"page"`
	RowNum    int    `toml:"results_num"`
	Api       string `toml:"api"`
}

//package wise global variables
var (
	aWxMiniAppRelatedConfig     wxMiniAppRelatedConfig
	aBaiduImageApiRelatedConfig baiduImageApiRelatedConfig
	baiduImageApiAcessToken     string
)

func init() {
	_, err := toml.DecodeFile("../config/WxMiniAppRelated.toml", &aWxMiniAppRelatedConfig)
	if err != nil {
		log.Fatal("Fail to load Wechat Mini-application related configurations.")
	}

	_, err = toml.DecodeFile("../config/BaiduImageApiRelated.toml", &aBaiduImageApiRelatedConfig)
	if err != nil {
		log.Fatal("Fail to load Baidu Image API related configurations.")
	}
	//get and maintain access token
	err = baiduImageApi_GetAccessToken("recognize")
	if err != nil {
		log.Fatal("Fail to get Baidu Image API recognize access token.")
	}
	err = baiduImageApi_GetAccessToken("search")
	if err != nil {
		log.Fatal("Fail to get Baidu Image API search access token.")
	}
	log.Println(recognize_token)
	log.Println(search_token)
	go recognize_token.accessTokenKeeper()
	go search_token.accessTokenKeeper()
}

func (w *wxMiniAppRelatedConfig) makeAuth_code2SessionRequest(code string) string {
	request := strings.Replace(w.AAPIS.Auth_code2Session, "APPID", w.AppID, -1)
	request = strings.Replace(request, "SECRET", w.AppSecret, -1)
	return strings.Replace(request, "JSCODE", code, -1)
}

func (b *baiduImageApiRelatedConfig) makeGetAccessTokenRequest(app string) string {
	var apiKey, secretKey string
	switch app {
	case "recognize":
		apiKey = b.Recognize.ApiKey
		secretKey = b.Recognize.SecretKey
	case "search":
		apiKey = b.Search.ApiKey
		secretKey = b.Search.SecretKey
	}
	requst := strings.Replace(b.GetAcessTokenApi, "CLIENT_ID", apiKey, -1)
	return strings.Replace(requst, "CLIENT_SECRET", secretKey, -1)
}

func (b *baiduImageApiRelatedConfig) makeSearchRequest() string {
	return strings.Replace(b.Search.Api, "ACCESS_TOKEN", search_token.Token, -1)
}

func (b *baiduImageApiRelatedConfig) makeRecognizeRequest() string {
	return strings.Replace(b.Search.Api, "ACCESS_TOKEN", recognize_token.Token, -1)
}
