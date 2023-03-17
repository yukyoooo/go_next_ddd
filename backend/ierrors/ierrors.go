package ierrors

import (
	"fmt"
)

func Wrap(errp *error, format string, args ...interface{}) {
	if *errp != nil {
		*errp = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), *errp)
	}
}

var ErrInvalidName = fmt.Errorf("名前は必須です")
var ErrInvalidDate = fmt.Errorf("日付形式が間違っています")