package model


type Product struct{
	ID  uint `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:varchar(255)"`
	Price uint `gorm:"type:int"`
}