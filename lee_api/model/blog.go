package model

import (
	"gorm.io/gorm"
)

type (
	TagType struct {
		gorm.Model
		Name string
	}
	Tags struct {
		gorm.Model
		tagType TagType `gorm:"foreignkey:TagTypeID"`
		BlogId  Blog
	}

	Blog struct {
		gorm.Model
		Uuid   string `json:"uuid" gorm:"column:uuid;type:varchar(100);"`
		UserId User   `gorm:"foreignkey:UserRefer"`
	}
)
