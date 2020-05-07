package webService

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type baiduImageApi_object_detect_Data struct {
	Image    string `json:"image"`
	WithFace int    `json:"with_face"`
}

type baiduImageApi_object_detect_Resp struct {
	Width            int    `json:"width"`
	Top              int    `json:"top"`
	Left             int    `json:"left"`
	Height           int    `json:"height"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func BaiduImageApi_object_detect(image string) (baiduImageApi_object_detect_Resp, error) {
	data := baiduImageApi_object_detect_Data{
		Image:    image,
		WithFace: aBaiduImageApiRelatedConfig.Recognize.WithFace,
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		return baiduImageApi_object_detect_Resp{}, err
	}

	request := aBaiduImageApiRelatedConfig.makeRecognizeRequest()
	response, err := http.Post(request, "application/x-www-form-urlencoded",
		bytes.NewBuffer(encoded))
	if err != nil {
		return baiduImageApi_object_detect_Resp{}, err
	}
	defer response.Body.Close()

	var b baiduImageApi_object_detect_Resp
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return baiduImageApi_object_detect_Resp{}, err
	}
	err = json.Unmarshal(body, &b)
	if err != nil {
		return baiduImageApi_object_detect_Resp{}, err
	}
	if b.Error != "" {
		return baiduImageApi_object_detect_Resp{}, errors.New(b.ErrorDescription)
	}

	return b, nil
}
