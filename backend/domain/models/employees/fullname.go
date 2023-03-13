package employee

import (
	"fmt"
	"regexp"

	"github.com/yukyoooo/go_next_ddd/iterrors"
)

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName string, lastName string) (_ *FullName, err error) {
	defer iterrors.Wrap(&err, "fullname.NewFullName(%s, %s)", firstName, lastName)
	fullName := new(FullName)

	if firstName == "" {
		return nil, fmt.Errorf("firstName is required")
	}
	if !ValidateName(firstName) {
		return nil, fmt.Errorf("firstName has an invalid character, letter is only")
	}
	fullName.firstName = firstName

	if lastName == "" {
		return nil, fmt.Errorf("lastName is required")
	}
	if !ValidateName(lastName) {
		return nil, fmt.Errorf("lastName has an invalid character, letter is only")
	}
	fullName.lastName = lastName

	return fullName, nil
}

func (fullName *FullName) FirstName() (firstName string) {
	return fullName.firstName
}

func (fullName *FullName) LastName() (lastName string) {
	return fullName.lastName
}

func ValidateName(value string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(value)
}
