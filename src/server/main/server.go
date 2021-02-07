package main

import (
    "bitmyth.com/accounts/src/app"
    "bitmyth.com/accounts/src/user/userrepo"
    "context"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    err := app.Bootstrap()
    userrepo.Get()

    if err != nil {
        panic(err)
    }
    app.RegisterRoutes()

    srv := app.Container.Server


    go func() {
        // service connections
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
