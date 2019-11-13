package controller

type Account struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

type Message struct {
	Message string `json:"message"`
}
