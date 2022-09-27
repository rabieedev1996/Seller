package Register

type RegisterCommand struct {
	PhoneNumber int64  `json:"PhoneNumber"`
	Password    string `json:"Password"`
}
