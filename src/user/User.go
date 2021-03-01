package user

import (
    "github.com/bitmyth/accounts/src/hash"
    "gorm.io/gorm"
)

// https://gorm.io/docs/models.html#embedded_struct
type User struct {
    ID       uint   `gorm:"primarykey"`
    Name     string `gorm:"index;type:varchar(100)" form:"name" binding:"required"`
    Password string `gorm:"type:varchar(100)" form:"password"`
    Email    string `gorm:"index type:varchar(100)" form:"email"`
    Phone    string `gorm:"index" form:"phone"`
    Avatar   string `gorm:"type:varchar(100)" form:"avatar"`
    gorm.Model
}

func (u *User) Credential() *Credential {
    return &Credential{
        Name:     u.Name,
        Password: u.Password,
    }

}

func (u *User) Authenticate(c *Credential) error {

    if err := u.Credential().Verify(c); err != nil {
        return err
    }
    return nil
}

func (u *User) Filter() *User {
    u.Password = ""
    return u
}

type Credential struct {
    Name     string `gorm:"index;type:varchar(100)" form:"name"`
    Password string `gorm:"type:varchar(100)" `
}

func (c *Credential) Verify(v *Credential) error {
    err := hash.Verify([]byte(c.Password), []byte(v.Password))
    if err != nil {
        return err
    }
    return nil
}
