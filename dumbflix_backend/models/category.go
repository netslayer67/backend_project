package models

import "time"

type Category struct {
	ID       int       `json:"id" gorm:"primary_key:auto_increment"`
	Name     string    `json:"name" gorm:"type:varchar(255)"`
	CreateAt time.Time `json:"-"`
	UpdateAt time.Time `json:"-"`
}
