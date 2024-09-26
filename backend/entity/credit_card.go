package entity

import ("gorm.io/gorm"
		"time"
	)
		
type CreditCard struct {
	gorm.Model 
	NameOnCard string
	CardNumber int
	ExpiryDate time.Time
	CVV int

	// Payment ทำหน้าที่เป็น FK
	PaymentID uint
	Payment   Payment `gorm:"foreignKey:payment_id"`
}