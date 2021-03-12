// reference: https://medium.com/wesionary-team/jwt-authentication-in-golang-with-gin-63dbc0816d55
package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/errgo.v2/errors"
	"os"
	"time"
)

var s JWTService

func JWT() JWTService {
	if s != nil {
		return s
	}
	s = NewJwtService()
	return s
}

//token service
type JWTService interface {
	GenerateToken(uid uint, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
	ValidateHeader(token string) (*jwt.Token, error)
}

type AuthCustomClaims struct {
	Uid  uint `json:"uid"`
	User bool `json:"user"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

//auth-token
func NewJwtService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "Bikash",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtService) GenerateToken(uid uint, isUser bool) string {
	claims := &AuthCustomClaims{
		uid,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(encodedToken, &AuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %s", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}

func (service *jwtService) ValidateHeader(authHeader string) (*jwt.Token, error) {
	const BEARER_SCHEMA = "Bearer "
	if len(authHeader) < len(BEARER_SCHEMA) {
		return nil, errors.New("token invalid")
	}

	tokenString := authHeader[len(BEARER_SCHEMA):]
	return service.ValidateToken(tokenString)
}
