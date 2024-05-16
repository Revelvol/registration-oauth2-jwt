package service

import (
	"revelvoler/registration-service/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


func SaveOrGetUserData(db *gorm.DB, user *model.User) (error) {

	// find by email and populate the user id 
	result := db.Where("email = ?", user.Email ).First(user)

	if result.Error == gorm.ErrRecordNotFound {
		// if email not found saved to db 

		//hash password first 
		hashedPasword, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
		if err!= nil {
			return err
		}

		user.Password = string(hashedPasword)
		db.Create(user)
	}

	return nil
}

func CreateUserDetail(db *gorm.DB, userDetail *model.UserDetail){
	db.Create(userDetail)
}

func GetUserDetail(db *gorm.DB, userId string) (model.UserDetail,error){
	userDetail := model.UserDetail{}
	// find by user id 
	result := db.Where("userId = ? ", userId).First(&userDetail)
	
	return userDetail, result.Error
}

func GetAllUserToken(db *gorm.DB, userId string, channel string) ([]model.UserToken, error){
	// find token by user id and channel 
	userToken := []model.UserToken{}

	result := db.Where("userId = ? ", ).Where("channel = ?", channel).Find(&userToken)
	
	return userToken, result.Error

}

func SaveOrUpdateToken(db *gorm.DB, userToken *model.UserToken) {
	db.Save(userToken)
}