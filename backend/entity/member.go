package entity

import "gorm.io/gorm"
		


type Member struct {
	gorm.Model 
	FirstName string
	LastName string
	Email string
	Username string
	Password  string 
	GenderID  uint 
	Age string
	

	
    // 1 member สามารถมีชำระเงินหลายครั่ง
	Payments [] Payment `gorm:"foreignKey:member_id "`
	
	
}