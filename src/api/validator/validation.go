package validator

import "strings"

func ValidateUsername(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return ErrorEmptyUsername
	}
	if len(s) > 256 {
		return ErrorUsernameExceededLimit
	}
	if len(s) < 8 {
		return ErrorMinUsername
	}
	if !((s[0] >= 65 && s[0] <= 90) || (s[0] >= 97 && s[0] <= 122)) {
		return ErrorUsernameFirstCharacterMustBeAplhabat
	}

	return nil
}

func ValidatePassword(s string) error {
	if len(s) == 0 {
		return ErrorEmptyPassword
	}
	if len(s) < 8 {
		return ErrorMinPassword
	}
	if len(s) > 256 {
		return ErrorPasswordExceededLimit
	}

	return nil
}

func ValidateEmail(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return ErrorEmptyEmail
	}

	if !strings.Contains(s, "@") {
		return ErrorEmailWrongFormat
	}

	if !strings.Contains(s, ".") {
		return ErrorEmailWrongFormat
	}

	return nil
}

func ValidateName(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return ErrorEmptyName
	}
	if !((s[0] >= 65 && s[0] <= 90) || (s[0] >= 97 && s[0] <= 122)) {
		return ErrorNameFirstCharacterMustBeAplhabat
	}
	if len(s) > 256 {
		return ErrorNameExceededLimit
	}

	return nil
}

func ValidateWorkflow(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return ErrorEmptyWorkflow
	}

	return nil
}
