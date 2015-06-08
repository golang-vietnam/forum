package apiErrors

type Error struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Status  int    `json:"message"`
}

var (
	USER_EXIST = Error{
		Id:      "USER_EXIST",
		Message: "This user has been exist!",
		Status:  400,
	}
	USER_EMAIL_REQUIRER = Error{
		Id:      "USER_EMAIL_REQUIRER",
		Message: "Email is required",
		Status:  400,
	}
	USER_EMAIL_INVALID = Error{
		Id:      "USER_EMAIL_INVALID",
		Message: "Email invalid",
		Status:  400,
	}
)
