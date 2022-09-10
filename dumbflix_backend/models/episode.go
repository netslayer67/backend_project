package models

import "time"

type Episode struct {
	ID            int          `json:"id" gorm:"primary_key:auto_increment"`
	Title         string       `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string       `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	LinkFilm      string `json:"linkfilm" gorm:"type:text" form:"linkfilm"`
	CreatedAt     time.Time    `json:"-"`
	UpdatedAt     time.Time    `json:"-"`
}
