package profile

import (
    "bitmyth.com/accounts/src/app/responses"
    "bitmyth.com/accounts/src/app/routes"
    "bitmyth.com/accounts/src/user"
    "bitmyth.com/accounts/src/user/userrepo"
    "fmt"
    "github.com/gin-gonic/gin"
)

func Show(context *gin.Context) *responses.Response {
    u, ok := context.Get("user")
    if !ok {
        return responses.Json(responses.Error{Code: "NotFound", Message: "user not found"})
    }

    filtered := u.(user.User)
    // Mask password
    filtered.Filter()

    return responses.Json(filtered)
}

func Update(context *gin.Context) *responses.Response {
    cu, _ := context.Get("user")
    u := cu.(user.User)

    var modifiedUser user.User
    err := context.Bind(&modifiedUser)
    fmt.Printf("%v",err)
    fmt.Printf("%v",modifiedUser)
    u.Name = modifiedUser.Name

    err = userrepo.Get().Save(u)
    if err != nil {
        return responses.Json(responses.Error{Code: "NotFound", Message: "user not found"})
    }

    return responses.Json(u.Filter())
}

func Routes() []routes.Route {

    return []routes.Route{
        {routes.GET, "/v1", "/profile", Show},
        {routes.POST, "/v1", "/profile", Update},
    }
}
