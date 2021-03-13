package errors

import "github.com/bitmyth/accounts/src/app/i18n"

type Err struct {
	Code    string
	Message string
}

func NewError(code string, err error) Err {
	if err != nil {
		return Err{
			Code:    code,
			Message: i18n.Translate(code) + err.Error(),
		}
	}
	return Err{
		Code:    code,
		Message: i18n.Translate(code),
	}
}

func (e Err) Error() string {
	return e.Message
}
