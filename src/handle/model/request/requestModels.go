package request

type UserRequestModel struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PhoneRequestModel struct {
	UserId      string `json:"user_id"`
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
}
