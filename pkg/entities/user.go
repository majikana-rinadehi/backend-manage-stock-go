package entities

type User struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	MailAddress string `json:"mailAddress"`
}
