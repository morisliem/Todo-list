package src

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

type UsernameError struct {
	Message string
}

var (
	// UserNotFound        = errors.New("user not found")
	FailedToDecode      = "Failed to decode input"
	FailedToAddUser     = "Failed to add user"
	SuccessfullyAdded   = "Successfully added"
	SuccessfullyLogin   = "Logged in successfully"
	SuccessfullyLogout  = "Logged out successfully"
	URLUsername         = "username"
	FailedToAddWorkflow = "Failed to add workflow"
)

func (err *UsernameError) Error() string {
	return fmt.Sprintf("%v", err.Message)
}

func UserNotFoundError() error {
	return errors.New("user not found")
}

func WrongPassword() error {
	return errors.New("wrong password")
}

func MinPassword() error {
	return errors.New("password min 8 character")
}

func PasswordExceededLimit() error {
	return errors.New("password max 256 character")
}

func EmptyUsername() error {
	return errors.New("empty username")
}

func MinUsername() error {
	return errors.New("username min 8 character")
}

func UsernameFirstCharacterMustBeAplhabat() error {
	return errors.New("username must start from alphabet")
}

func NameFirstCharacterMustBeAplhabat() error {
	return errors.New("name must start from alphabet")
}

func UsernameExceededLimit() error {
	return errors.New("username max 256 character")
}

func NameExceededLimit() error {
	return errors.New("name max 256 character")
}

func EmptyName() error {
	return errors.New("empty name")
}

func EmptyEmail() error {
	return errors.New("empty email")
}

func EmptyPassword() error {
	return errors.New("empty password")
}

func EmailWrongFormat() error {
	return errors.New("email missing @ or . ")
}

func EmptyWorkflow() error {
	return errors.New("empty workflow")
}

func WorkflowNotFoundError() error {
	return errors.New("workflow not found")
}

func Response(s string) map[string]string {
	result := map[string]string{}

	result["Message"] = s

	return result
}
