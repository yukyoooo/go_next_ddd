package employee

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/yukyoooo/go_next_ddd/ierrors"
)

type Email struct {
	value string
}

func (e Email) Value() string {
	return e.value
}

func NewEmail(value string) (_ *Email, err error) {
	defer ierrors.Wrap(&err, "email.NewEmail(%s", value)
	email := new(Email)

	if value == "" {
		return nil, fmt.Errorf("email is required")
	}

	if !validateEmail(value) {
		return nil, fmt.Errorf("wrong format")
	}
	email.value = value

	return email, nil
}

func validateEmail(value string) bool {
	return govalidator.IsEmail(value)
}
