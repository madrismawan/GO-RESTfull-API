package models

type Item struct {
	ItemId      uint   `gorm:"primaryKey; autoIncrement"`
	ItemCode    string `gorm:"not null;unique; type:varchar(191)"`
	Description string `gorm:"not null;type:varchar(20)"`
	Quantity    int    `gorm:"not null; type:int"`
	OrderId     uint
}
