package workflow

import (
	"encoding/json"
	"log"

	"github.com/kataras/iris/v12"

	"MuseumAR_Backend/infrastruture/database"
	"MuseumAR_Backend/infrastruture/webService"
)

type Wxminiapp_auth_code2Session_Resp struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

/*
	the work flow of the sign in process
	1. extract parameters from request
	2. send request to wx service, deal with the result
	3. deal with database update
	4. sign for a token
	5. return information to user
*/
func SignIn(ctx iris.Context) {
	//extract parameters
	code := ctx.URLParam("code")
	if code == "" {
		ctx.StatusCode(400)
		return
	}

	//send request to wx, auth the code
	respBytes, err := webService.Wxminiapp_auth_code2Session(code)
	if err != nil {
		log.Fatal(err)
		ctx.StatusCode(500)
		return
	}
	//decode the response
	var w Wxminiapp_auth_code2Session_Resp
	err = json.Unmarshal(respBytes, &w)
	if err != nil {
		log.Fatal(err)
		ctx.StatusCode(500)
		return
	}

	switch w.ErrCode {
	case -1:
		log.Fatal(err)
		ctx.StatusCode(500)
		return
	case 40029:
		fallthrough
	case 45011:
	default:
		ctx.StatusCode(400)
		return
	case 0:
	}
	//database update
	db := database.Get()
	rows, err := db.Query([]string{"*"},
		[]string{"users"}, "u_openid = '"+w.OpenID+"'", nil)
	if err != nil {
		log.Fatal(err)
		ctx.StatusCode(500)
		return
	}
	if rows.Next() == false {
		err = db.Insert("users", []string{"u_openid"}, []string{"'" + w.OpenID + "'"}, nil)
		if err != nil {
			log.Fatal(err)
			ctx.StatusCode(500)
			return
		}
	}
	//sign for a token
	token, err := signForToken(w.OpenID, w.SessionKey)
	if err != nil {
		log.Fatal(err)
		ctx.StatusCode(500)
		return
	}

	ctx.JSON(iris.Map{
		"openid": w.OpenID,
		"token":  token,
	})
}
