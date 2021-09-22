package main

import (
	"net/http"

	"gorm.io/gorm"
)

//DB stores the database session information.
type DBClient struct{
	db *gorm.DB
}

//UserResponse is the response to be send back for user
type UserResponse struct{
	User models.User `json:"user"`
	Data interface{} `json:"data"`
}

//GetUserByFirstName fetches the original URL for the given encoded(short) string
func (driver *DBClient) GetUserByFirstName(w http.ResponseWriter, r *http.Request){

}
func main(){

}