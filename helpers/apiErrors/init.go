package apiErrors

type Error struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) WithNewMessage(message string) *Error {
	return &Error{
		Id:      e.Id,
		Message: message,
		Status:  e.Status,
	}
}
func NewError(id string, message string, status int) *Error {
	return &Error{
		Id:      id,
		Message: message,
		Status:  status,
	}
}

var allError []*Error

func init() {
	allError = append(privateErrors, userErrors...)
}

func cloneError(e *Error) *Error {
	newError := *e
	return &newError
}

func ThrowError(errorId string) *Error {
	for index := range allError {
		if allError[index].Id == errorId {
			return cloneError(allError[index])
		}
	}
	panic("Error Throw Not Found")
}
