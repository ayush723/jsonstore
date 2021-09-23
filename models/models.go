package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	Orders []Order
	Data   string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Order struct {
	gorm.Model
	User User
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

//GORM creates tables with plural names.
func (User) TableName() string {
	return "user"
}

func (Order) TableName() string {
	return "order"
}

func InitDB() (*gorm.DB, error) {
	var err error
	
	db, err := gorm.Open("postgres","postgres://ayush:envy@localhost/jsonstore?sslmode=disable")
	if err != nil {
		return nil, err
	}
	
	//creates tables 
	db.AutoMigrate(&User{}, &Order{})
		return db, nil
	

}
