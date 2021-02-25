package register

import (
    "github.com/bitmyth/accounts/src/app/auth/token"
    "github.com/bitmyth/accounts/src/app/responses"
    "github.com/bitmyth/accounts/src/app/routes"
    "github.com/bitmyth/accounts/src/hash"
    "github.com/bitmyth/accounts/src/user"
    "github.com/bitmyth/accounts/src/user/userrepo"
    "github.com/gin-gonic/gin"
)

func Register(context *gin.Context) *responses.Response {
    //name := context.PostForm("name")
    //password := context.PostForm("password")

    type RegisterForm struct {
        user.User
        PasswordConfirm string `form:"passwordConfirm"`
    }
    var req RegisterForm
    _ = context.BindJSON(&req)

    userRepo := userrepo.Get()

    condition := &user.User{
        Name: req.Name,
    }

    var found user.User
    err := userRepo.First(&found, condition)

    // Found existing user with the same name
    if err == nil {
        return responses.Json(responses.ValidationError{
            Code:    "invalid-name",
            Message: "Invalid name",
            Errors:  map[string]string{"name": "name exists"},
        })
    }

    if req.Password != req.PasswordConfirm {
        return responses.Json(responses.ValidationError{
            Message: "Wrong password confirmation",
            Errors:  map[string]string{"passwordConfirm": "wrong password confirmation"},
        })
    }

    hashed, err := hash.Make([]byte(req.User.Password))
    if err != nil {
        return responses.Json("failed hashing password")
    }

    u := &user.User{
        Name:     req.Name,
        Password: string(hashed),
    }

    err = userRepo.Save(u)

    if err != nil {
        return responses.Json(err)
    }

    jwt := token.JWT().GenerateToken(u.ID, true)

    return responses.Json(gin.H{"token": jwt, "user": u.Filter()})
}

func Routes() []routes.Route {

    return []routes.Route{
        {"POST", "/v1", "/register", Register},
    }
}
