package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/juxuny/clipboard-sync/lib"
	"github.com/juxuny/clipboard-sync/lib/env"
	"github.com/pkg/errors"
	"time"
)

var secret = env.GetStringWithDefault(env.Key.JwtSecret, "lJZtiFGVR5ZBtV4a")

type UserInfo struct {
	UserId lib.ID `json:"userId,omitempty"`
}

type Claims struct {
	*jwt.StandardClaims
	UserInfo
}

func Parse(token string) (*Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "parse token failed")
	}
	body, ok := parsedToken.Claims.(*Claims)
	if ok {
		return body, nil
	}
	return nil, errors.Wrap(err, "invalid token")
}

func CreateToken(u UserInfo) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
		UserInfo: u,
	})
	return token.SignedString([]byte(secret))
}
