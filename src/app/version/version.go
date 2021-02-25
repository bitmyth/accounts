package version

import (
    "github.com/gin-gonic/gin"
    "io"
    "net/http"
    "runtime"
    "text/template"
)

var (
    // Version holds the current version.
    Version = "dev"
    // Codename holds the current version codename.
    Codename = "cheddar" // beta cheese
    // BuildDate holds the build date.
    BuildTime = "I don't remember exactly"
)

var versionTemplate = `Version:      {{.Version}}
Codename:     {{.Codename}}
Go version:   {{.GoVersion}}
Built:        {{.BuildTime}}
OS/Arch:      {{.Os}}/{{.Arch}}

`

type Info struct {
    Version   string
    Codename  string
    GoVersion string
    BuildTime string
    Os        string
    Arch      string
}

func New() Info {
    return Info{
        Version:   Version,
        Codename:  Codename,
        GoVersion: runtime.Version(),
        BuildTime: BuildTime,
        Os:        runtime.GOOS,
        Arch:      runtime.GOARCH,
    }

}

// Append adds version routes on a router.
func (v Info) Append(router gin.IRouter) {

    router.GET("/api/version", func(c *gin.Context) {

        v := New()

        c.JSON(http.StatusOK, v)
    })

}

// Print write Printable version.
func Print(wr io.Writer) error {
    tmpl, err := template.New("").Parse(versionTemplate)
    if err != nil {
        return err
    }
    v := New()

    return tmpl.Execute(wr, v)
}
