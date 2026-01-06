package entity

import (
	"gorm.io/gorm"
)
type Students struct {
	gorm.Model
	Fullname string	`valid:"required~Fullname is required"`
	Age      uint	`valid:"required,range(18|100)~Age must be at least 18"`
	Email    string `valid:"required~Email is required,email~Email is invalid"`
	GPA      float32 `valid:"matches(^([0-3]\\.[0-9]{2}|4\\.00)$)~GPA must be between 0.00 and 4.00"`
}