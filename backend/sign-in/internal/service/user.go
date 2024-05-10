package userService

import (
	"revelvoler/registration-service/internal/model"

	"gorm.io/gorm"
)


func SaveOrGetUserData(db *gorm.DB,  user *model.User) {
	// find by email 

	// if email found return 

	// if not saved and return 
}

func GetUserDetail(db *gorm.DB,  userDetail *model.UserDetail){
	// find by user id 

	// return user detail 
}

func saveUserToken(db *gorm.DB, useToken *model.UserToken){
	// find token by user id and channel 

	// return user token detail

}

func updateExistingToken(db *gorm.DB, useToken *model.UserToken) {
	// find token by user id and channel 


	// update curent with now acces token and stuff
}