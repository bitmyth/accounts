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
    var credential user.Credential

    userRepo := userrepo.Get()

    name := context.PostForm("name")
    phone := context.PostForm("phone")

    credential.Password = context.PostForm("password")

    for _, username := range []string{name, phone} {
        if username != "" {
            credential.Name = username
            break
        }
    }

    condition := &user.User{
        Name: name,
    }

    err := userRepo.First(&u, condition)
    if err != nil {
        return responses.Json("user not found")
    }

    if err := u.Authenticate(&credential); err != nil {
        return responses.Json("wrong password")
    }

    jwt := token.JWT().GenerateToken(u.ID, true)
    return responses.Json(jwt)
}

func Routes() []routes.Route {

    return []routes.Route{
        {"POST", "/v1", "/login", Login},
    }
}
