package apiErrors

type apiError struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *apiError) Error() string {
	return e.Message
}

func (e *apiError) WithNewMessage(message string) *apiError {
	return &apiError{
		Id:      e.Id,
		Message: message,
		Status:  e.Status,
	}
}
func NewError(id string, message string, status int) *apiError {
	return &apiError{
		Id:      id,
		Message: message,
		Status:  status,
	}
}

var allError []apiError

func init() {
	allError = append(privateErrors, userErrors...)
}

func cloneError(e *apiError) *apiError {
	newError := *e
	return &newError
}

// Use for Error API
func GetErrorById(errorId string) *apiError {
	for index := range allError {
		if allError[index].Id == errorId {
			return cloneError(&allError[index])
		}
	}
	return nil
}

func ThrowError(errorId string) *apiError {
	if err := GetErrorById(errorId); err != nil {
		return err
	}
	panic("Error To Throw Not Defined")
}

func ParseError(err error) *apiError {
	if parseError, ok := err.(*apiError); ok {
		return parseError
	}
	return nil
}
