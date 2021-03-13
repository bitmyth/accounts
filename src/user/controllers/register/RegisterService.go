package register

import (
	"github.com/bitmyth/accounts/src/hash"
	"github.com/bitmyth/accounts/src/user"
)

type Service struct {
	User     user.User
	UserRepo user.RepositoryInterface
}

func NewService(user user.User, repo user.RepositoryInterface) *Service {
	return &Service{user, repo}
}

func (s Service) Do() (*user.User, error) {
	var found user.User
	condition := &user.User{
		Name: s.User.Name,
	}
	err := s.UserRepo.First(&found, condition)

	// Found existing user with the same name
	if err == nil {
		return nil, NewNameExistsError()
	}

	hashed, err := hash.Make([]byte(s.User.Password))
	if err != nil {
		return nil, NewPasswordHashFailedError(err)
	}

	u := &user.User{
		Name:     s.User.Name,
		Password: string(hashed),
	}

	err = s.UserRepo.Save(u)

	if err != nil {
		return nil, NewSaveError(err)
	}

	return u, nil
}
