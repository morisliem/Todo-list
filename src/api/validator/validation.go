package validator

import (
	"strings"
	"time"
)

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

func ValidateTodoTitle(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return ErrorEmptyTitle
	}

	if len(s) > 256 {
		return ErrorTitleExceededLimit
	}

	return nil
}

func ValidateTodoState(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return ErrorEmptyState
	}

	return nil
}

func ValidateTodoPriority(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return ErrorEmptyPriority
	}
	for _, value := range Priority {
		if strings.EqualFold(value, s) {
			return nil
		}
	}

	return ErrorInvalidPriority
}

func ValidateTodoSeverity(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return ErrorEmptySeverity
	}
	for _, value := range Severity {
		if strings.EqualFold(value, s) {
			return nil
		}
	}

	return ErrorInvalidSeverity
}

func ValidateTodoDeadline(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return ErrorEmptyDeadline
	}

	deadline, err := time.Parse("01-02-2006", s)

	if err != nil {
		return ErrorInvalidDeadline
	}

	if deadline.Before(time.Now()) {
		return ErrorDeadlineMustBeAfterToday
	}

	return nil
}

func ValidateTodoId(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return ErrorEmptyTodoId
	}

	return nil
}
