package models

import "time"

type Order struct {
	OrderId      uint      `gorm:"primaryKey; autoIncrement"`
	CustomerName string    `gorm:"not null;type:varchar(100)"`
	Items        []Item    `gorm:"foreignKey:OrderId"`
	OrderedAt    time.Time `json:"order_at"`
}
