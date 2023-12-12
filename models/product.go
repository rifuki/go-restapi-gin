package models

type Product struct {
	Id          int32  `gorm:"primarykey" json:"id"`
	Name        string `gorm:"type:varchar(300)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}
