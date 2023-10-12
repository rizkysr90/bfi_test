package model

import "gorm.io/gorm"
type User struct {
	gorm.Model
	Id            int32          `gorm:"primaryKey; autoIncrement:true"`
	Username      string         `json:"username"`
	Name 			string		`json:"name"`
	Tickets []Ticket `gorm:"foreignKey:UserId"`
}