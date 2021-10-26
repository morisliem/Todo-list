package response

import (
	"errors"
)

type DataStoreError struct {
	Message string
}

type ServerInternalError struct {
	Message string
}

type BadInputError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

func (d *DataStoreError) Error() string {
	return d.Message
}

func (d *ServerInternalError) Error() string {
	return d.Message
}

func (d *BadInputError) Error() string {
	return d.Message
}

func (d *NotFoundError) Error() string {
	return d.Message
}

var (
	ErrorUserNotFound                         = errors.New("user not found")
	ErrorTodoNotFound                         = errors.New("todo not found")
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
	ErrorEmptyTodoId                          = errors.New("empty TodoId")
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
	ErrorInternalServer                       = errors.New("internal server error")
	ErrorFailedToDecode                       = errors.New("failed to decode input")
	ErrorToParsePicture                       = errors.New("failed to parse input")
	ErrorToRetrieveFile                       = errors.New("failed to retrieve file")
	ErrorToCreateTempFile                     = errors.New("failed to create temporary file")
	ErrorToSaveFile                           = errors.New("failed to save file")
	ErrorFailedToAddUser                      = errors.New("failed to add user")
	ErrorFailedToAddTodo                      = errors.New("failed to add todo")
	ErrorFailedToAddPict                      = errors.New("failed to add picture")
	ErrorFailedToAddWorkflow                  = errors.New("failed to add workflow")
	ErrorFailedToUpdateUserTodo               = errors.New("failed to update user todo")
	ErrorWorkflowNotExist                     = errors.New("workflow does not exist")
	SuccessfullyAdded                         = "Successfully added"
	SuccessfullyUpdated                       = "Successfully updated"
	SuccessfullyLogin                         = "Logged in successfully"
	SuccessfullyLogout                        = "Logged out successfully"
	URLUsername                               = "username"
	URLUTodoId                                = "todoId"
	Priority                                  = [3]string{"High", "Medium", "Low"}
	Severity                                  = [3]string{"High", "Medium", "Low"}
)

func Response(s string) map[string]string {
	result := map[string]string{}
	result["Message"] = s
	return result
}
