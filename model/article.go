package model

import (
	"gorm.io/gorm"
	"gorm.io/datatypes"
)

type (
	Article struct {
		gorm.Model
		Title string         `json:"title"`
		Body datatypes.JSON  `json:"body"`
		Read  bool           `json:"read"`
	}
)
