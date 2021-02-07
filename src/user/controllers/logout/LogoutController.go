package logout

import (
    "bitmyth.com/accounts/src/app/responses"
    "bitmyth.com/accounts/src/app/routes"
    "bitmyth.com/accounts/src/user"
    "bitmyth.com/accounts/src/user/userrepo"
    "github.com/gin-gonic/gin"
)

func Logout(context *gin.Context) *responses.Response {
    var u user.User

    userRepo := userrepo.Get()

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
