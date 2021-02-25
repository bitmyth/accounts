package routes

import (
    "github.com/bitmyth/accounts/src/app/responses"
    "fmt"
    "github.com/gin-gonic/gin"
)

type Route struct {
    Method  string
    Prefix  string
    Url     string
    Handler HandlerFunc
}

type HandlerFunc func(context *gin.Context) *responses.Response

const GET = "GET"
const POST = "POST"
const PATCH = "PATCH"

func RegisterRoutes(router gin.IRouter, routes []Route) {
    var iRoutes gin.IRoutes

    for _, route := range routes {
        iRoutes = router

        if len(route.Prefix) > 0 {
            iRoutes = router.Group(route.Prefix)
        }

        f := func(c *gin.Context) {
            response := route.Handler(c)

            switch response.Type {
            case responses.TEXT:
                c.String(response.HttpStatusCode, fmt.Sprintf("%v", response.Content))
            case responses.JSON:
                c.JSON(response.HttpStatusCode, response.Content)
            default:
                c.JSON(response.HttpStatusCode, response.Content)
            }
        }

        switch route.Method {
        case GET:
            iRoutes.GET(route.Url, f)
        case POST:
            iRoutes.POST(route.Url, f)
        case PATCH:
            iRoutes.PATCH(route.Url, f)
        }

    }
}
