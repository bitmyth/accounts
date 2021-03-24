package main

import (
	"context"
	"fmt"
	"github.com/bitmyth/accounts/src/app"
	"github.com/bitmyth/accounts/src/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	_ = os.Setenv("TZ", "Asia/Shanghai")
	err := app.Bootstrap()
	println(config.RootPath)
	//gin.SetMode(gin.ReleaseMode)

	if err != nil {
		//panic(err)
	}
	app.RegisterRoutes()

	srv := app.Container.Server

	go func() {
		// service connections
		_, _ = os.Stdout.Write([]byte(fmt.Sprintf("Server listen on port %s\n", srv.Addr)))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 0*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
