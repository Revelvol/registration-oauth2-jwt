package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id              string    `gorm:"column:user;default:uuid_generate_v4();primaryKey"` // default gorm data model for auto generate id, also primary key
	CreatedAt       time.Time //special field
	UpdatedAt       time.Time //special field
	Email           string    `gorm:"uniqueIndex"`
	Password        string
	IsDetailFilled  bool
	UserDetail      UserDetail        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	UserToken       []UserToken       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	EmailValidation []EmailValidation `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

type UserDetail struct {
	Id             string `gorm:"column:user_detail_id;default:uuid_generate_v4();primaryKey"` // default gorm data model for auto generate id, also primary key
	UserId         string
	CreatedAt      time.Time //special field
	UpdatedAt      time.Time //special field
	FullName       string
	Address        string
	PhoneNumber    string
	ProfilePicPath string
}

type UserToken struct {
	Id           string `gorm:"column:user_token;default:uuid_generate_v4();primaryKey"` // default gorm data model for auto generate id, also primary key
	UserId       string
	CreatedAt    time.Time //special field
	UpdatedAt    time.Time //special field
	AccessToken  sql.NullString
	RefreshToken sql.NullString
	Channel      sql.NullString
}

type EmailValidation struct {
	Id                 string `gorm:"column:email_validation;default:uuid_generate_v4();primaryKey"`
	UserId             string
	CreatedAt          time.Time //special field
	UpdatedAt          time.Time //special field
	ExpiredAt          time.Time
	Status             string
	VerificationString string
}
