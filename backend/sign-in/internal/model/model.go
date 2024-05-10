package model

import "time"

type User struct {
	ID           string `gorm:"column:gpu_id;default:uuid_generate_v4();primaryKey"` // default gorm data model for auto generate id, also primary key
	CreatedAt    time.Time //special field
	UpdatedAt    time.Time //special field
}

type UserDetail struct {
	ID           string `gorm:"column:gpu_id;default:uuid_generate_v4();primaryKey"` // default gorm data model for auto generate id, also primary key
	CreatedAt    time.Time //special field
	UpdatedAt    time.Time //special field
}

type UserToken struct {
	ID           string `gorm:"column:gpu_id;default:uuid_generate_v4();primaryKey"` // default gorm data model for auto generate id, also primary key
	CreatedAt    time.Time //special field
	UpdatedAt    time.Time //special field
}