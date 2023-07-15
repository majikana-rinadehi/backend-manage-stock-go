package entities

type StockCategory struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Name   string `json:"name"`
}
