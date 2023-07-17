package entities

import (
	"github.com/majikana-rinadehi/backend-manage-stock-go/util"
)

type Stock struct {
	Id         int    `json:"id"`
	UserId     int    `json:"userId"`
	CategoryId int    `json:"categoryId"`
	Name       string `json:"name"`
	Amount     int    `json:"amount"`
	ExpireDate string `json:"expireDate"`
}

func (s Stock) String() string {
	return util.CustomStringer(s)
}
