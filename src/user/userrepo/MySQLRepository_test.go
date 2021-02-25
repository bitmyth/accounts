package userrepo

import (
    "github.com/bitmyth/accounts/src/app"
    "github.com/bitmyth/accounts/src/hash"
    "github.com/bitmyth/accounts/src/user"
    "testing"
)

var userRepo *UserRepository
var fakeName string

func TestMain(m *testing.M) {
    _ = app.Bootstrap()

    userRepo = Get()

    fakeName = "sam"

    m.Run()
}

func TestSave(t *testing.T) {
    password, err := hash.Make([]byte("123"))

    u := &user.User{
        Credential: user.Credential{
            Name:     fakeName,
            Password: string(password),
        },
    }

    err = userRepo.Save(u)
    if err != nil {
        t.Fatalf("failed find %v", err)
    }

    t.Logf("saved user: %v", u)
}

func TestFind(t *testing.T) {
    var u []user.User

    err := userRepo.Find(&u, &user.User{Credential: user.Credential{Name: fakeName}})
    if err != nil {
        err.Error()
        t.Fatalf("failed find %v", err)
    }

    t.Logf("find: %v", u[0].Credential)
}

func TestUpdate(t *testing.T) {
    var u user.User

    err := userRepo.First(&u, &user.User{Credential: user.Credential{Name: fakeName}})
    if err != nil {
        t.Fatalf("find failed%v", err)
    }
    t.Logf("find user: %v", u)

    u.Credential.Name = "foo"
    err = userRepo.Save(&u)
    if err != nil {
        t.Fatalf("update failed %v", err)
    }

    t.Logf("updated user: %v", u)
}

