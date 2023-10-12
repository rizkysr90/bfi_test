package model

import (
	"time"

	"gorm.io/gorm"
)
type Ticket struct {
	gorm.Model
	Id            int32          `gorm:"primaryKey; autoIncrement:true"`
	Masalah         string         `json:"masalah"`
	UserId       int32         `json:"userId"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     time.Time `gorm:"index" json:"deletedAt"`
}
