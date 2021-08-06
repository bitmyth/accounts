// reference: https://medium.com/wesionary-team/jwt-authentication-in-golang-with-gin-63dbc0816d55
package token

import (
    "crypto/rsa"
    "errors"
    "github.com/bitmyth/accounts/src/config"
    "github.com/dgrijalva/jwt-go"
    "time"
)

var Jwt JWTService

type JWTService interface {
    GenerateToken(uid uint, roles []string) string
    ValidateToken(token string) (*jwt.Token, error)
    ValidateHeader(token string) (*jwt.Token, error)
}

func NewJwt(pubKeyData []byte, priKeyData []byte) JWTService {
    pubKey, _ := jwt.ParseRSAPublicKeyFromPEM(pubKeyData)
    j := &jwtService{
        issuer: "Bitmyth",
    }

    if len(pubKeyData) > 0 {
        j.rsaPublicKey = pubKey
    }

    // For testing
    if len(priKeyData) > 0 {
        priKey, _ := jwt.ParseRSAPrivateKeyFromPEM(priKeyData)
        j.rsaPrivateKey = priKey
    }

    return j
}

func Bootstrap() error {
    priKeyData := config.Secret.GetString("rsa.privateKey")

    Jwt = NewJwt([]byte{}, []byte(priKeyData))
    return nil
}

type AuthCustomClaims struct {
    Uid   uint     `json:"uid"`
    Roles []string `json:"roles"`
    jwt.StandardClaims
}

type jwtService struct {
    secretKey     string
    issuer        string
    rsaPrivateKey *rsa.PrivateKey
    rsaPublicKey  *rsa.PublicKey
}

func (service *jwtService) GenerateToken(uid uint, roles []string) string {
    claims := &AuthCustomClaims{
        uid,
        roles,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
            Issuer:    service.issuer,
            IssuedAt:  time.Now().Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
    t, err := token.SignedString(service.rsaPrivateKey)

    if err != nil {
        panic(err)
    }
    return t
}

const BEARER_SCHEMA = "Bearer "

func (service *jwtService) ValidateHeader(authHeader string) (*jwt.Token, error) {
    if len(authHeader) < len(BEARER_SCHEMA) {
        return nil, errors.New("token invalid")
    }

    tokenString := authHeader[len(BEARER_SCHEMA):]

    return service.ValidateToken(tokenString)
}

func (service *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {

    return jwt.ParseWithClaims(encodedToken, &AuthCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return service.rsaPublicKey, nil
    })

}
