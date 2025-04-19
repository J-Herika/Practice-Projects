package database

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"size:100;unique"`
	PasswordHash string
}

type UserInfo struct {
	Username string
	Password string
}

type File struct {
	ID         uint
	FileName   string
	Path       string
	UserID     uint
	UploadTime time.Time
}
