package model

type User struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(255)"`
	Email string `gorm:"type:varchar(255)"`
}
