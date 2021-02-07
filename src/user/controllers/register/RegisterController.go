package register

import (
    "bitmyth.com/accounts/src/app/responses"
    "bitmyth.com/accounts/src/app/routes"
    "bitmyth.com/accounts/src/hash"
    "bitmyth.com/accounts/src/user"
    "bitmyth.com/accounts/src/user/userrepo"
    "github.com/gin-gonic/gin"
)

func Register(context *gin.Context) *responses.Response {
    name := context.PostForm("name")

    password := context.PostForm("password")
    hashed, err := hash.Make([]byte(password))
    if err != nil {
        return responses.Json("failed hashing password")
    }

    user := &user.User{
        Name:     name,
        Password: string(hashed),
    }

    userRepo := userrepo.Get()
    err = userRepo.Save(user)

    if err != nil {
        return responses.Json(err)
    }

    return responses.Json(user)
}

func Routes() []routes.Route {

    return []routes.Route{
        {"POST", "/v1", "/register", Register},
    }
}
