package entities

type Stock struct {
	Id         int    `json:"id"`
	UserId     int    `json:"userId"`
	CategoryId int    `json:"categoryId"`
	Name       string `json:"name"`
	Amount     int    `json:"amount"`
	ExpireDate string `json:"expireDate"`
}
