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

func UsernameExceededLimit() error {
	return errors.New("username max 256 character")
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

func SuccessfullyAdded() string {
	return "Successfully added"
}

func SuccessfullyLogin() string {
	return "Logged in successfully"
}

func SuccessfullyLogout() string {
	return "Logged out successfully"
}

func JSONSyntaxError() string {
	return "Failed to decode"
}
