package app

import (
	"github.com/bitmyth/accounts/src/app/auth/token"
	"github.com/bitmyth/accounts/src/app/boot"
	"github.com/bitmyth/accounts/src/app/middlewares"
	"github.com/bitmyth/accounts/src/app/routes"
	"github.com/bitmyth/accounts/src/app/version"
	"github.com/bitmyth/accounts/src/config"
	"github.com/bitmyth/accounts/src/database/mysql"
	"github.com/bitmyth/accounts/src/user"
	"github.com/bitmyth/accounts/src/user/controllers/login"
	"github.com/bitmyth/accounts/src/user/controllers/logout"
	"github.com/bitmyth/accounts/src/user/controllers/profile"
	"github.com/bitmyth/accounts/src/user/controllers/register"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type App struct {
	Server     *http.Server
	Bootstraps []boot.Bootstrap
}

var (
	Container *App
)

func init() {
	Container = New()
}

func New() *App {
	Container = &App{}
	Container.Bootstraps = []boot.Bootstrap{
		config.Bootstrap,
		mysql.Bootstrap,
		user.Repo.Bootstrap,
		token.Bootstrap,
    }

	return Container
}

func Bootstrap() error {

	for _, b := range Container.Bootstraps {
		err := b()
		if err != nil {

			i := 1
			// Retry forever
			for {
				time.Sleep(3 * time.Second)
				err = b()
				i++
				if err == nil {
					break
				} else {
					println(err.Error())
				}
			}
		}
	}

	return nil
}

func RegisterRoutes() {

	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://google.com"}
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")

	router.Use(cors.New(corsConfig))

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	routes.RegisterRoutes(router, register.Routes())
	routes.RegisterRoutes(router, login.Routes())
	routes.RegisterRoutes(router, logout.Routes())

	protected := router.Group("/")
	protected.Use(middlewares.Auth())

	routes.RegisterRoutes(protected, profile.Routes())

	version.Info{}.Append(router)
	//_ = version.Print(os.Stdout)

	port := viper.GetString("server.port")
	Container.Server = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
}
