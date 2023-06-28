package entity

// MessageTable 消息表
type MessageTable struct {
	Id       int    `gorm:"primary_key;auto_increment;not null"`
	Content  string `gorm:"type:varchar(255);not null"`
	NickName string `gorm:"type:varchar(255);not null"`
	Date     string `gorm:"type:dateTime;not null"`
}
