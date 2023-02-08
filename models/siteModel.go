package models

import "gorm.io/gorm"

type Site struct {
	gorm.Model

	SiteURL string
	Tag1 string
	Tag2 string
	KeyWord string
}