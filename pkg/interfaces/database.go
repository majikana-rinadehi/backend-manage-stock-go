package interfaces

import "gorm.io/gorm"

type Database interface {
	Connect() error
	Disconnect() error
	GetDB() *gorm.DB
}
