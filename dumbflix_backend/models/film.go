package models

import "time"

type Film struct {
	ID            int    `json:"id" gorm:"primary_key:auto_increment"`
	Title         string `json:"title" form:"title" gorm:"type: varchar(255)"`
	ThumbnailFilm string `json:"thumbnailfilm" form:"thumbnailfilm" gorm:"type: varchar(255)"`
	Image         string `json:"image" form:"image" gorm:"type: varchar(255)"`
	Description   string `json:"description" gorm:"type:text" form:"desc"`
	Year          int    `json:"year" gorm:"type: int"`
	// Category      CategoryResponse `json:"category"`
	CategoryID int       `json:"category_id" form:"category_id"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
