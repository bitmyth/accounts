package user

import (
    "github.com/bitmyth/accounts/src/config"
    "github.com/bitmyth/accounts/src/database/mysql"
    "github.com/bitmyth/accounts/src/hash"
    "testing"
)

var userRepo *Repository
var fakeName string

func TestMain(m *testing.M) {
    println(config.RootPath)
    _ = config.Bootstrap()
    _ = mysql.Bootstrap()
    _ = Repo.Bootstrap()

    userRepo = Repo

    fakeName = "sam"

    m.Run()
}

func TestSave(t *testing.T) {
    password, err := hash.Make([]byte("123"))

    u := &User{
        Name:     fakeName,
        Password: string(password),
    }

    err = userRepo.Save(u)
    if err != nil {
        t.Fatalf("failed find %v", err)
    }

    t.Logf("saved user: %v", u)
}

func TestFind(t *testing.T) {
    var u []User

    err := userRepo.Find(&u, &User{Name: fakeName})
    if err != nil {
        err.Error()
        t.Fatalf("failed find %v", err)
    }

    t.Logf("find: %v", u[0])
}

func TestUpdate(t *testing.T) {
    var u User

    err := userRepo.First(&u, &User{Name: fakeName})
    if err != nil {
        t.Fatalf("find failed%v", err)
    }
    t.Logf("find user: %v", u)

    u.Name = "foo"
    err = userRepo.Save(&u)
    if err != nil {
        t.Fatalf("update failed %v", err)
    }

    t.Logf("updated user: %v", u)
}
