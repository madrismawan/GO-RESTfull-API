package models

type Item struct {
	ItemId      uint   `gorm:"primaryKey; autoIncrement" json:"lineItemId"`
	ItemCode    string `gorm:"not null;unique; type:varchar(191)" validate:"required,min=3" json:"itemCode"`
	Description string `gorm:"not null;type:varchar(20)" validate:"required" json:"description,min:4"`
	Quantity    int    `gorm:"not null; type:int" validate:"required" json:"quantity"`
	OrderId     uint
}
