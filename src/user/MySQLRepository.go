package user

import (
	"github.com/bitmyth/accounts/src/database/mysql"
	"gorm.io/gorm"
)

var Repo = NewRepo()

type RepositoryInterface interface {
	Find(out interface{}, query interface{}, args ...interface{}) error
	First(out interface{}, query interface{}, args ...interface{}) error
	Save(value interface{}) error
}

type Repository struct {
	DB *gorm.DB
}

func NewRepo() *Repository {
	return &Repository{}
}

func (ur *Repository) Init(db *gorm.DB) (*Repository, error) {
	Repo.DB = db

	return Repo, nil
}

func (ur *Repository) Bootstrap() error {
	db, err := mysql.Connect()
	if err != nil {
		return err
	}

	Repo, err = ur.Init(db)
	return err
}

func (ur *Repository) RunQuery(query func() error) error {
	err := query()

	if ur.CausedByLostConnection(err) {
		err = query()
	}

	return err
}

func (ur *Repository) CausedByLostConnection(err error) bool {
	return err != nil && err.Error() == "invalid connection"
}

func (ur *Repository) UpdateOrCreate(wheres *User) (*User, error) {
	return nil, nil
}

func (ur *Repository) Find(out interface{}, query interface{}, args ...interface{}) error {
	return ur.RunQuery(func() error {
		return ur.DB.Where(query).Find(out).Error
	})
}

func (ur *Repository) First(out interface{}, query interface{}, args ...interface{}) error {
	return ur.RunQuery(func() error {
		return ur.DB.Where(query).First(out).Error
	})
}

func (ur *Repository) Save(value interface{}) error {
	return ur.RunQuery(func() error {
		return ur.DB.Save(value).Error
	})
}
