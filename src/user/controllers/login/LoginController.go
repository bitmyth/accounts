package login

import (
    "github.com/bitmyth/accounts/src/app/auth/token"
    "github.com/bitmyth/accounts/src/app/responses"
    "github.com/bitmyth/accounts/src/app/routes"
    "github.com/bitmyth/accounts/src/user"
    "github.com/gin-gonic/gin"
)

func Login(context *gin.Context) *responses.Response {
    var u user.User
    var found user.User
    var credential user.Credential

    userRepo := user.Repo

    _ = context.BindJSON(&u)

    credential.Name = u.Name
    credential.Password = u.Password

    condition := u.Filter()

    err := userRepo.First(&found, condition)
    if err != nil {
        return responses.Json(responses.ValidationError{
            Code:    "invalid-credential",
            Message: "Invalid credential",
            Errors:  map[string]string{"name": "User Not Found"},
        })
    }

    if err := found.Authenticate(&credential); err != nil {
        return responses.Json(responses.ValidationError{
            Message: "Wrong password",
            Errors:  map[string]string{"password": "wrong password"},
        })
        return responses.Json("wrong password")
    }

    jwt := token.Jwt.GenerateToken(u.ID, []string{})
    return responses.Json(gin.H{"token": jwt, "user": found})
}

func Routes() []routes.Route {

    return []routes.Route{
        {"POST", "/v1", "/login", Login},
    }
}
