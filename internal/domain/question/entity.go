package question

import "gorm.io/gorm"

type Entity struct {
	gorm.Model
	Description string
}
