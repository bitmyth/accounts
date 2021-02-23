package login

import (
    "bitmyth.com/accounts/src/app/auth/token"
    "bitmyth.com/accounts/src/app/responses"
    "bitmyth.com/accounts/src/app/routes"
    "bitmyth.com/accounts/src/user"
    "bitmyth.com/accounts/src/user/userrepo"
    "github.com/gin-gonic/gin"
)

func Login(context *gin.Context) *responses.Response {
    var u user.User
    var found user.User
    var credential user.Credential

    userRepo := userrepo.Get()

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

    jwt := token.JWT().GenerateToken(u.ID, true)
    return responses.Json(gin.H{"token": jwt, "user": u})
}

func Routes() []routes.Route {

    return []routes.Route{
        {"POST", "/v1", "/login", Login},
    }
}
