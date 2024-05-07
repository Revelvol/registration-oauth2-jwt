package model

import "time"

type User struct {
	ID           string `gorm:"column:gpu_id;default:uuid_generate_v4();primaryKey"` // default gorm data model for auto generate id, also primary key
	GpuName      string
	Manufacturer *string   // a pointer to string, allowing for null values
	CreatedAt    time.Time //special field
	UpdatedAt    time.Time //special field
}