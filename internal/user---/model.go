package user

type User struct {
	Id           uint64 `json:"id"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Age          uint8  `json:"age"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}
