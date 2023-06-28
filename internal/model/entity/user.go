package entity

// UserTable 用户表
type UserTable struct {
	Id       int    `gorm:"primary_key;auto_increment;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Nickname string `gorm:"type:varchar(255);not null"`
}
