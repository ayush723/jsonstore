package models

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	dsn := "postgres://ayush:envy@tcp(127.0.0.1:5432)/jsonstore"
	sqlDB, err := sql.Open("postgres",dsn)
	if err != nil{
		panic(err)
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		//the below automigrate is equivalent to this
		// if !db.Migrator().HasTable("user"){
		// 	db.Migrator().CreateTable(&User{})
		// }
		// if !db.Migrator().HasTable("order"){
		// 	db.Migrator().CreateTable(&Order{})
		// }

		db.AutoMigrate(&User{}, &Order{})
		return db, nil
	}

}
