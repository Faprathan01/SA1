package entity

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model    
	Username string    
	Password string    
	Email   string    
	FirstName string 
	LastName   string
	
	GenderID  uint     
	Gender    Gender  `gorm:"foreignKey:gender_id"`
	
	//Classes []Class `gorm:"foreignKey:AdminID"`
} 
