package webService

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type baiduImageApi_same_hq_search_Data struct {
	Image   string `json:"image"`
	Museum  string `json:"tags"`
	PageNum string `json:"pn"`
	RowNum  string `json:"rn"`
}

type baiduImageApi_same_hq_search_Resp struct {
	Results []struct {
		Score float64 `json:"score"`
		Brief string  `json:"brief"`
	}
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

func BaiduImageApi_same_hq_search(image string, museum string) (string, error) {
	data := baiduImageApi_same_hq_search_Data{
		Image:   image,
		Museum:  museum,
		PageNum: string(aBaiduImageApiRelatedConfig.Search.PageNum),
		RowNum:  string(aBaiduImageApiRelatedConfig.Search.RowNum),
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	request := aBaiduImageApiRelatedConfig.makeSearchRequest()
	response, err := http.Post(request, "application/x-www-form-urlencoded",
		bytes.NewBuffer(encoded))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	var b baiduImageApi_same_hq_search_Resp
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &b)
	if err != nil {
		return "", err
	}
	if b.Error != "" {
		return "", errors.New(b.ErrorDescription)
	}
	if b.Results[0].Score < 0.9 {
		return "", errors.New("Not found.")
	}
	return b.Results[0].Brief, nil
}
