package register

import (
	"github.com/bitmyth/accounts/src/app/auth/token"
	"github.com/bitmyth/accounts/src/app/responses"
	"github.com/bitmyth/accounts/src/app/routes"
	"github.com/bitmyth/accounts/src/user"
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
	err := context.BindJSON(&req)
	if err != nil {
		return responses.Json(err.Error())
	}

	if req.Password != req.PasswordConfirm {
		return responses.Json(responses.ValidationError{
			Message: "Wrong password confirmation",
			Errors:  map[string]string{"passwordConfirm": "wrong password confirmation"},
		})
	}

	service := NewService(req.User, user.Repo)

	registered, err := service.Do()

	if err != nil {

		switch err.(type) {
		case NameExistsError:
			return responses.Json(responses.ValidationError{
				Code:    err.(NameExistsError).Code,
				Message: err.(NameExistsError).Message,
				Errors:  map[string]string{"name": err.Error()},
			})
		case PasswordHashFailedError:
			return responses.Json("failed hashing password")
		case SaveError:
			return responses.Json(err)
		}
	}

	jwt := token.Jwt.GenerateToken(registered.ID, []string{})

	return responses.Json(gin.H{"token": jwt, "user": registered.Filter()})
}

func Routes() []routes.Route {

	return []routes.Route{
		{"POST", "/v1", "/register", Register},
	}
}
