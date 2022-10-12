package models

import "time"

type Order struct {
	OrderId      uint      `gorm:"primaryKey; autoIncrement"`
	CustomerName string    `gorm:"not null;type:varchar(100)" validate:"required" json:"customerName" `
	OrderedAt    time.Time `validate:"required" json:"orderedAt" `
	Items        []Item    `gorm:"foreignKey:OrderId" validate:"required" json:"Items" `
}
