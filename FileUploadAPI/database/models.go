package database


type User struct {
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Username     string `gorm:"size:100;unique"`
	PasswordHash string
}

type UserInfo struct {
	Username string
	Password string
}