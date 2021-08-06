package middlewares

import (
    "github.com/bitmyth/accounts/src/app/auth/token"
    "github.com/bitmyth/accounts/src/app/responses"
    "github.com/bitmyth/accounts/src/user"
    "github.com/gin-gonic/gin"
    "net/http"
)

func Auth() gin.HandlerFunc {

    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")

        t, err := token.Jwt.ValidateHeader(authHeader)

        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, responses.AuthError{Code: "unauthorized", Message: err.Error()})
            return
        }

        claims := t.Claims.(*token.AuthCustomClaims)
        userID := claims.Uid

        var u user.User
        err = user.Repo.First(&u, user.User{ID: userID})
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, responses.AuthError{Code: "UserNotFound", Message: err.Error()})
        }

        c.Set("user", u)
    }
}
