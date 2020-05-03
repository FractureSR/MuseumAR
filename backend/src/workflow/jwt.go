package workflow

import (
	"github.com/iris-contrib/middleware/jwt"
)

func signForToken(openid string, sessionkey string) (string, error) {
	token := jwt.NewTokenWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"openid":      openid,
			"session_key": sessionkey,
		})

	tokenString, err := token.SignedString([]byte(aWorkflowConfig.JwtSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
