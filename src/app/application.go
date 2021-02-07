package app

import (
    "bitmyth.com/accounts/src/app/boot"
    "bitmyth.com/accounts/src/app/middlewares"
    "bitmyth.com/accounts/src/app/routes"
    "bitmyth.com/accounts/src/config"
    "bitmyth.com/accounts/src/database/mysql"
    "bitmyth.com/accounts/src/user/controllers/login"
    "bitmyth.com/accounts/src/user/controllers/logout"
    "bitmyth.com/accounts/src/user/controllers/profile"
    "bitmyth.com/accounts/src/user/controllers/register"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

type App struct {
    Server     *http.Server
    Bootstraps []boot.Bootstrap
}

var Container *App

func init() {
    Container = New()
}

func New() *App {
    Container = &App{}
    Container.Bootstraps = []boot.Bootstrap{
        mysql.Bootstrap,
    }

    return Container
}

func Bootstrap() error {
    config.Read()

    for _, b := range Container.Bootstraps {
        err := b()
        if err != nil {
            return err
        }
    }

    return nil
}

func RegisterRoutes() {

    router := gin.Default()
    //router.Use(middlewares.Auth())

    router.GET("/", func(c *gin.Context) {
        time.Sleep(5 * time.Second)
        c.String(http.StatusOK, "Welcome Gin Server")
    })

    routes.RegisterRoutes(router, register.Routes())
    routes.RegisterRoutes(router, login.Routes())
    routes.RegisterRoutes(router, logout.Routes())


    protected:=router.Group("/")
    protected.Use(middlewares.Auth())

    routes.RegisterRoutes(protected, profile.Routes())

    Container.Server = &http.Server{
        Addr:    ":8080",
        Handler: router,
    }
}
