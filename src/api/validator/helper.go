package validator

import (
	"errors"
)

type ValidationError struct {
	Message error
}

var (
	ErrorUserNotFound                         = errors.New("user not found")
	ErrorWorkflowNotFound                     = errors.New("workflow not found")
	ErrorWrongPassword                        = errors.New("wrong password")
	ErrorEmptyUsername                        = errors.New("empty username")
	ErrorEmptyName                            = errors.New("empty name")
	ErrorEmptyEmail                           = errors.New("empty email")
	ErrorEmptyPassword                        = errors.New("empty password")
	ErrorEmptyWorkflow                        = errors.New("empty workflow")
	ErrorEmptyTitle                           = errors.New("empty title")
	ErrorEmptyState                           = errors.New("empty state")
	ErrorEmptyPriority                        = errors.New("empty priority")
	ErrorEmptySeverity                        = errors.New("empty severity")
	ErrorEmptyDeadline                        = errors.New("empty Deadline")
	ErrorMinUsername                          = errors.New("username min 8 character")
	ErrorMinPassword                          = errors.New("password min 8 character")
	ErrorUsernameExceededLimit                = errors.New("username max 256 character")
	ErrorTitleExceededLimit                   = errors.New("title max 256 character")
	ErrorNameExceededLimit                    = errors.New("name max 256 character")
	ErrorPasswordExceededLimit                = errors.New("password max 256 character")
	ErrorUsernameFirstCharacterMustBeAplhabat = errors.New("username must start from alphabet")
	ErrorNameFirstCharacterMustBeAplhabat     = errors.New("name must start from alphabet")
	ErrorEmailWrongFormat                     = errors.New("email missing @ or . ")
	ErrorInvalidPriority                      = errors.New("invalid value for priority")
	ErrorInvalidSeverity                      = errors.New("invalid value for severity")
	ErrorInvalidDeadline                      = errors.New("invalid deadline")
	ErrorDeadlineMustBeAfterToday             = errors.New("deadline must be after today")
	FailedToDecode                            = "Failed to decode input"
	FailedToAddUser                           = "Failed to add user"
	FailedToAddTodo                           = "Failed to add todo"
	FailedToAddWorkflow                       = "Failed to add workflow"
	FailedToUpdateUserTodo                    = "Failed to update user todo"
	SuccessfullyAdded                         = "Successfully added"
	SuccessfullyLogin                         = "Logged in successfully"
	SuccessfullyLogout                        = "Logged out successfully"
	URLUsername                               = "username"
	Priority                                  = [3]string{"High", "Medium", "Low"}
	Severity                                  = [3]string{"High", "Medium", "Low"}
)

func Response(s string) map[string]string {
	result := map[string]string{}
	result["Message"] = s
	return result
}
