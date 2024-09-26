package config

import (
	"fmt"

	"PJ/backend/entity"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("sa.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {

	db.AutoMigrate(
		&entity.Member{},
		&entity.Admin{},
		&entity.Gender{},
		&entity.Payment{},
		&entity.Package{},
		&entity.Paypal{},
		&entity.PromptPay{},
		&entity.CreditCard{},
		
		
	)
	GenderMale := entity.Gender{Name: "Male"}
	GenderFemale := entity.Gender{Name: "Female"}

	db.FirstOrCreate(&GenderMale, &entity.Gender{Name: "Male"})
	db.FirstOrCreate(&GenderFemale, &entity.Gender{Name: "Female"})

	hashedPassword, _ := HashPassword("654321")
	Member := entity.Member{
		Username:  "Fah",
		Password:  hashedPassword,
		Email:     "Fah@gmail.com",
		FirstName: "Faprathan",
		LastName:  "Lertsungnoen",
		GenderID: 2,
	

	}
	hashedPasswordAd, _ := HashPassword("123456")
	Admin := entity.Admin{
		Username: "FahAdmin", 
		Password: hashedPasswordAd, 
		Email: "FahAdmin@gmail.com", 
		FirstName: "Admin",
		LastName:  "nnnnn",
		GenderID: 2,
		
    }

	Package := entity.Package{
		PackageName:  "Daily",
		Description:  "Members can access all services within the fitness center for a full day",
		Price:     "59THB/d.",
		Duration_days: "1 day" ,
		
	}


	db.FirstOrCreate(&Member, entity.Member{Email: "Fah@gmail.com"})
	db.FirstOrCreate(&Admin, entity.Admin{Email: "FahAdmin@gmail.com"})

	db.FirstOrCreate(&Package, entity.Package{PackageName: "Daily_Membership"})

}
