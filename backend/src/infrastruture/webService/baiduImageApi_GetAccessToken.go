package webService

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type accessToken struct {
	App   string
	Token string
	Alarm *time.Timer
}

type baiduImageApi_GetAccessToken_Resp struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

var (
	recognize_token accessToken
	search_token    accessToken
)

func (a *accessToken) accessTokenKeeper() {
	for {
		select {
		case <-a.Alarm.C:
			baiduImageApi_GetAccessToken(a.App)
		}
	}
}

func baiduImageApi_GetAccessToken(app string) error {
	request := aBaiduImageApiRelatedConfig.makeGetAccessTokenRequest(app)
	response, err := http.Post(request, "application/x-www-form-urlencoded",
		bytes.NewBuffer(nil))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var b baiduImageApi_GetAccessToken_Resp
	err = json.Unmarshal(body, &b)
	if err != nil {
		return err
	}
	if b.Error != "" {
		return errors.New(b.ErrorDescription)
	}

	switch app {
	case "recognize":
		recognize_token.App = app
		recognize_token.Token = b.AccessToken
		recognize_token.Alarm = time.NewTimer(
			time.Duration(b.ExpiresIn-60) * time.Second)
	case "search":
		search_token.App = app
		search_token.Token = b.AccessToken
		search_token.Alarm = time.NewTimer(
			time.Duration(b.ExpiresIn-60) * time.Second)
	}
	return nil
}
