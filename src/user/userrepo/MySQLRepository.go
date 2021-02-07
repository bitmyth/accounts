package userrepo

import (
    "bitmyth.com/accounts/src/database/mysql"
    "bitmyth.com/accounts/src/user"
    "gorm.io/gorm"
)

var instance *UserRepository

type UserRepository struct {
    Client *gorm.DB
}

func NewUserRepository() (*UserRepository, error) {
    instance = &UserRepository{}

    db, err := mysql.DB()
    if err != nil {
        return nil, err
    }

    instance.Client = db

    db.AutoMigrate(&user.User{})

    return instance, nil
}

func Get() *UserRepository {
    if instance != nil {
        return instance
    }

    instance, err := NewUserRepository()

    if err != nil {
        panic("failed init user repository")
    }

    return instance
}

func (ur *UserRepository) RunQuery(query func() error) error {
    err := query()

    if ur.CausedByLostConnection(err) {
        err = query()
    }

    return err
}

func (ur *UserRepository) CausedByLostConnection(err error) bool {
    return err != nil && err.Error() == "invalid connection"
}

func (ur *UserRepository) UpdateOrCreate(wheres *user.User) (*user.User, error) {
    return nil, nil
}

func (ur *UserRepository) Find(out interface{}, query interface{}, args ...interface{}) error {
    return ur.RunQuery(func() error {
        return ur.Client.Where(query).Find(out).Error
    })
}

func (ur *UserRepository) First(out interface{}, query interface{}, args ...interface{}) error {
    return ur.RunQuery(func() error {
        return ur.Client.Where(query).First(out).Error
    })
}

func (ur *UserRepository) Save(value interface{}) error {
    return ur.RunQuery(func() error {
        return ur.Client.Save(value).Error
    })
}
