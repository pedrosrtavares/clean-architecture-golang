package validator

import "github.com/dlclark/regexp2"

var emailRegex = regexp2.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, 0)

func IsValidEmail(email string) error {
	ok, err := emailRegex.MatchString(email)
	if !ok || err != nil {
		return err
	}

	return nil
}
