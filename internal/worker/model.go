package worker

type Workrer struct {
	Id           string `json:"id"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Age          uint8  `json:"age"`
	Experieons   uint8  `json:"exp"`
	Number       string `json:"number"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
}
