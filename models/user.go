package models

import (
	"jwt/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `json:"full_name" gorm:"not null" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `json:"email" gorm:"not null;uniqueIndex" form:"email" valid:"required~Your email name is required,email~Invalid email format"`
	Password string    `json:"password" gorm:"not null" form:"password" valid:"required~Your password name is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Role     string    `json:"role" gorm:"not null;default:user" form:"role" valid:"in(admin|user)"`
	Products []Product `json:"products" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
