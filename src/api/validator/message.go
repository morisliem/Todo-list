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
	ErrorMinUsername                          = errors.New("username min 8 character")
	ErrorMinPassword                          = errors.New("password min 8 character")
	ErrorUsernameExceededLimit                = errors.New("username max 256 character")
	ErrorNameExceededLimit                    = errors.New("name max 256 character")
	ErrorPasswordExceededLimit                = errors.New("password max 256 character")
	ErrorUsernameFirstCharacterMustBeAplhabat = errors.New("username must start from alphabet")
	ErrorNameFirstCharacterMustBeAplhabat     = errors.New("name must start from alphabet")
	ErrorEmailWrongFormat                     = errors.New("email missing @ or . ")
	FailedToDecode                            = "Failed to decode input"
	FailedToAddUser                           = "Failed to add user"
	SuccessfullyAdded                         = "Successfully added"
	SuccessfullyLogin                         = "Logged in successfully"
	SuccessfullyLogout                        = "Logged out successfully"
	URLUsername                               = "username"
	FailedToAddWorkflow                       = "Failed to add workflow"
)

func Response(s string) map[string]string {
	result := map[string]string{}

	result["Message"] = s

	return result
}
