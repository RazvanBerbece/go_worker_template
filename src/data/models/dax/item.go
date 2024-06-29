package dax

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	DisplayName string
}
