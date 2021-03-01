package profile

import (
    "fmt"
    "github.com/bitmyth/accounts/src/app/responses"
    "github.com/bitmyth/accounts/src/app/routes"
    "github.com/bitmyth/accounts/src/user"
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
    fmt.Printf("%v", err)
    fmt.Printf("%v", modifiedUser)
    u.Name = modifiedUser.Name

    err = user.Repo.Save(u)
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
