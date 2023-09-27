package entities

type UserAuth struct {
	UserId       int    `json:"userId"`
	AuthProvider string `json:"authProvider"`
	Uid          string `json:"uid"`
}
