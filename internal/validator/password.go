package validator

import (
	"github.com/dlclark/regexp2"
)

var passwordRegex = regexp2.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`, 0)

func IsValidPassword(password string) error {
	ok, err := passwordRegex.MatchString(password)
	if !ok || err != nil {
		return err
	}

	return nil
}
