package employee

import (
	"crypto/sha1"
	"fmt"

	"github.com/yukyoooo/go_next_ddd/ierrors"
)

type Password struct {
	value string
}

func NewPassword(value string) (_ *Password, err error) {
	defer ierrors.Wrap(&err, "NewPassword(%s", value)
	password := new(Password)

	if value == "" {
		return nil, fmt.Errorf("password is required")
	}

	password.value = Encrypt(value)

	return password, nil
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
