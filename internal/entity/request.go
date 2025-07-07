package entity

type Jsutrequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewJustRequest(username, email, passwrod string) *Jsutrequest {
	return &Jsutrequest{
		UserName: username,
		Email:    email,
		Password: passwrod,
	}
}
