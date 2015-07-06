package apiErrors

type apiError struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
	Detail  string `json:"detail"`
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

// Use for test, parse to ApiError...
type ApiError apiError

// Use for Error API
var ApiErrors []apiError

func init() {
	ApiErrors = append(privateErrors, userErrors...)
	ApiErrors = append(ApiErrors, apiErrorErrors...)
	ApiErrors = append(ApiErrors, authErrors...)
}

func cloneError(e *apiError) *apiError {
	newError := *e
	return &newError
}

// Use for Error API
func FindErrorById(errorId string) *apiError {
	for index := range ApiErrors {
		if ApiErrors[index].Id == errorId {
			return cloneError(&ApiErrors[index])
		}
	}
	return nil
}

func ThrowError(errorId string) *apiError {
	if err := FindErrorById(errorId); err != nil {
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
