package response

type UserResponseModel struct {
	Id    int                 `json:"id"`
	Name  string              `json:"name"`
	Email string              `json:"email"`
	Phone *PhoneResponseModel `json:"phone,omitempty"`
}

type PhoneResponseModel struct {
	Id          int    `json:"id"`
	UserId      int    `json:"user_id"`
	CountryCode string `json:"country_code"`
	Number      string `json:"number"`
}
