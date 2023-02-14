package validator

import "github.com/dlclark/regexp2"

var brazilPhoneRegex = regexp2.MustCompile(`^\+55 \d{2} \d{4,5}-\d{4}$`, 0)

func IsValidPhone(phoneNumber string) error {
	ok, err := brazilPhoneRegex.MatchString(phoneNumber)
	if !ok || err != nil {
		return err
	}

	return nil
}
