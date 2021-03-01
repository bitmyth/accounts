package logout

import (
    "github.com/bitmyth/accounts/src/app/responses"
    "github.com/bitmyth/accounts/src/app/routes"
    "github.com/bitmyth/accounts/src/user"
    "github.com/gin-gonic/gin"
)

func Logout(context *gin.Context) *responses.Response {
    var u user.User

    userRepo := user.Repo

    condition := &user.User{
    }

    err := userRepo.First(&u, condition)
    if err != nil {
        return responses.Json("user not found")
    }

    return responses.Json(u)
}

func Routes() []routes.Route {

    return []routes.Route{
        {"POST", "/v1", "/logout", Logout},
    }
}
