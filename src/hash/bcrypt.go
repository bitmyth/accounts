package hash

import (
    "golang.org/x/crypto/bcrypt"
)

func Make(data []byte) ([]byte, error) {
    hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    return hash, err
}

func Verify(hashed []byte, data []byte) error {
    err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(data))
    return err
}
