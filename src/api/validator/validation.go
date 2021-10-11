package validator

import (
	"strings"
	"time"
	"todo-list/src/api/response"
)

func ValidateUsername(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return response.ErrorEmptyUsername
	}
	if len(s) > 256 {
		return response.ErrorUsernameExceededLimit
	}
	if len(s) < 8 {
		return response.ErrorMinUsername
	}
	if !((s[0] >= 65 && s[0] <= 90) || (s[0] >= 97 && s[0] <= 122)) {
		return response.ErrorUsernameFirstCharacterMustBeAplhabat
	}

	return nil
}

func ValidatePassword(s string) error {
	if len(s) == 0 {
		return response.ErrorEmptyPassword
	}
	if len(s) < 8 {
		return response.ErrorMinPassword
	}
	if len(s) > 256 {
		return response.ErrorPasswordExceededLimit
	}

	return nil
}

func ValidateEmail(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return response.ErrorEmptyEmail
	}

	if !strings.Contains(s, "@") {
		return response.ErrorEmailWrongFormat
	}

	if !strings.Contains(s, ".") {
		return response.ErrorEmailWrongFormat
	}

	return nil
}

func ValidateName(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return response.ErrorEmptyName
	}
	if !((s[0] >= 65 && s[0] <= 90) || (s[0] >= 97 && s[0] <= 122)) {
		return response.ErrorNameFirstCharacterMustBeAplhabat
	}
	if len(s) > 256 {
		return response.ErrorNameExceededLimit
	}

	return nil
}

func ValidateWorkflow(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return response.ErrorEmptyWorkflow
	}

	return nil
}

func ValidateTodoTitle(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return response.ErrorEmptyTitle
	}

	if len(s) > 256 {
		return response.ErrorTitleExceededLimit
	}

	return nil
}

func ValidateTodoState(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return response.ErrorEmptyState
	}

	return nil
}

func ValidateTodoPriority(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return response.ErrorEmptyPriority
	}
	for _, value := range response.Priority {
		if strings.EqualFold(value, s) {
			return nil
		}
	}

	return response.ErrorInvalidPriority
}

func ValidateTodoSeverity(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return response.ErrorEmptySeverity
	}
	for _, value := range response.Severity {
		if strings.EqualFold(value, s) {
			return nil
		}
	}

	return response.ErrorInvalidSeverity
}

func ValidateTodoDeadline(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return response.ErrorEmptyDeadline
	}

	deadline, err := time.Parse("01-02-2006", s)

	if err != nil {
		return response.ErrorInvalidDeadline
	}

	if deadline.Before(time.Now()) {
		return response.ErrorDeadlineMustBeAfterToday
	}

	return nil
}

func ValidateTodoId(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return response.ErrorEmptyTodoId
	}

	return nil
}
