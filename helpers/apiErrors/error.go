package apiErrors

type Error struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *Error) Error() string {
	return e.Message
}
