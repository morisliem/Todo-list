package src

import "strings"

func ValidateUsername(s string) error {
	if len(s) == 0 {
		return EmptyUsername()
	}
	if len(s) > 256 {
		return UsernameExceededLimit()
	}
	if len(s) < 8 {
		return MinUsername()
	}
	if !((s[0] >= 65 && s[0] <= 90) || (s[0] >= 97 && s[0] <= 122)) {
		return UsernameFirstCharacterMustBeAplhabat()
	}

	return nil
}

func ValidatePassword(s string) error {
	if len(s) == 0 {
		return EmptyPassword()
	}
	if len(s) < 8 {
		return MinPassword()
	}
	if len(s) > 256 {
		return PasswordExceededLimit()
	}

	return nil
}

func ValidateEmail(s string) error {
	if len(s) == 0 {
		return EmptyEmail()
	}

	if !strings.Contains(s, "@") {
		return EmailWrongFormat()
	}

	if !strings.Contains(s, ".") {
		return EmailWrongFormat()
	}

	return nil
}

func ValidateName(s string) error {
	if len(s) == 0 {
		return EmptyName()
	}
	if !((s[0] >= 65 && s[0] <= 90) || (s[0] >= 97 && s[0] <= 122)) {
		return NameFirstCharacterMustBeAplhabat()
	}
	if len(s) > 256 {
		return NameExceededLimit()
	}

	return nil
}

func ValidateWorkflow(s string) error {
	if len(s) == 0 {
		return EmptyWorkflow()
	}

	return nil
}
