package user

type User struct {
	Id           string `json:"id"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Age          uint8  `json:"age"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}
