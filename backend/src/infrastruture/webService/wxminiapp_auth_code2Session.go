package webService

import (
	"io/ioutil"
	"net/http"
)

func Wxminiapp_auth_code2Session(code string) ([]byte, error) {
	//form the request url
	request := aWxMiniAppRelatedConfig.makeAuth_code2SessionRequest(code)
	//send the request
	response, err := http.Get(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
