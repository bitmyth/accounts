package token

import (
    "github.com/bitmyth/accounts/src/config"
    "testing"
)


func TestRSASign(t *testing.T) {
    config.RootPath = config.RootPath + "/../../../../"
    println(config.RootPath)
    _ = config.Bootstrap()

    prikeyData := config.Secret.GetString("rsa.privateKey")
    jwt := NewJwt([]byte{}, []byte(prikeyData))
    r := jwt.GenerateToken(22, []string{"admin", "user"})
    t.Log(r)

    //pubkeyData, _ := ioutil.ReadFile("test/key.pub")
    //prikeyData, _ := ioutil.ReadFile("test/key")
}
