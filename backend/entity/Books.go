package entity

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model 
	Title string  `valid:"alpha,stringlength(3|100)"`
	Price float64 `valid:"range(50|500)"`
	Code  string  `valid:"matches(^[BK][0-9]{6}$)"`
}

func (b Books) Validate() error {
	_, err := govalidator.ValidateStruct(b)
	return err
}
