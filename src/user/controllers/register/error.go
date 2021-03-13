package register

import (
	"github.com/bitmyth/accounts/src/app/errors"
	"github.com/bitmyth/accounts/src/app/i18n/locale"
)

type NameExistsError struct {
	errors.Err
}

func NewNameExistsError() error {
	return NameExistsError{
		errors.NewError(locale.NameExist, nil),
	}
}

type PasswordHashFailedError struct {
	errors.Err
}

func NewPasswordHashFailedError(err error) error {
	return PasswordHashFailedError{
		errors.NewError(locale.PasswordHashFailed, err),
	}
}

type SaveError struct {
	errors.Err
}

func NewSaveError(err error) error {
	return SaveError{
		errors.NewError(locale.SaveFailed, err),
	}
}
