package workflow

import (
	"log"

	"github.com/BurntSushi/toml"
	jwt_go "github.com/dgrijalva/jwt-go"
	"github.com/iris-contrib/middleware/jwt"
)

type workflowConfig struct {
	Jwt struct {
		JwtSecret        string `toml:"jwtsecret"`
		JwtSigningMethod string `toml:"jwtsigningmethod"`
	}
}

var (
	aWorkflowConfig  workflowConfig
	jwtAuthenticator *jwt.Middleware
)

func init() {
	_, err := toml.DecodeFile("../config/WorkflowRelated.toml", &aWorkflowConfig)
	if err != nil {
		log.Fatal("Fail to load Database related configurations.")
	}

	var signingMethod jwt_go.SigningMethod
	switch aWorkflowConfig.Jwt.JwtSigningMethod {
	case "ES256":
		signingMethod = jwt.SigningMethodES256
	case "ES384":
		signingMethod = jwt.SigningMethodES384
	case "ES512":
		signingMethod = jwt.SigningMethodES512
	case "HS256":
		signingMethod = jwt.SigningMethodHS256
	case "HS384":
		signingMethod = jwt.SigningMethodHS384
	case "HS512":
		signingMethod = jwt.SigningMethodHS512
	}

	jwtAuthenticator = jwt.New(jwt.Config{
		// Extract by "token" url parameter.
		Extractor: jwt.FromParameter("token"),
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(aWorkflowConfig.Jwt.JwtSecret), nil
		},
		SigningMethod: signingMethod,
	})
}
