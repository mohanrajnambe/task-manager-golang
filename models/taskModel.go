package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID          uint
	Title       string
	Description string
	Status      string
}
