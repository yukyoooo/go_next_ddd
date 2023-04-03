package employee

import (
	"fmt"
	"regexp"

	"github.com/yukyoooo/go_next_ddd/ierrors"
)

type FullName struct {
	FirstName string
	LastName  string
}

func NewFullName(firstName string, lastName string) (_ *FullName, err error) {
	defer ierrors.Wrap(&err, "fullname.NewFullName(%s, %s)", firstName, lastName)
	fullName := new(FullName)

	if firstName == "" {
		return nil, fmt.Errorf("firstName is required")
	}
	if !ValidateName(firstName) {
		return nil, fmt.Errorf("firstName has an invalid character, letter is only")
	}
	fullName.FirstName = firstName

	if lastName == "" {
		return nil, fmt.Errorf("lastName is required")
	}
	if !ValidateName(lastName) {
		return nil, fmt.Errorf("lastName has an invalid character, letter is only")
	}
	fullName.LastName = lastName

	return fullName, nil
}

func (fullName *FullName) GetFirstName() (firstName string) {
	return fullName.FirstName
}

func (fullName *FullName) GetLastName() (lastName string) {
	return fullName.LastName
}

func ValidateName(value string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(value)
}
