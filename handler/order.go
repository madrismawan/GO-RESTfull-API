package handler

import (
	"example/Tugas/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderRepo struct {
	DB *gorm.DB
}

func NewRepoOrder(db *gorm.DB) *OrderRepo {
	return &OrderRepo{
		DB: db,
	}
}

func (r *OrderRepo) CreateOrder(ctx *gin.Context) {
	var orderInput models.Order
	err := ctx.ShouldBindJSON(&orderInput)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	// order := models.Order{}
	// items := orderInput.Items
	// order.CustomerName = orderInput.CustomerName
	// order.OrderedAt = orderInput.OrderedAt
	// order.Items = items

	// order.Items = orderInput.Items
	// // newBranch, err := r.DB.Create(&order)

	// err := r.db.Create(&orderInput).Error
	// if err != nil {
	// 	return branch, err
	// }
	err = r.DB.Create(&orderInput).Error
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusCreated, orderInput)
}

func (r *OrderRepo) GetOrder(ctx *gin.Context) {
	var orders []models.Order
	err := r.DB.Model(&models.Order{}).Preload("Items").Find(&orders).Error
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func (r *OrderRepo) UpdateOrder(ctx *gin.Context) {

}

func (r *OrderRepo) DeleteOrder(ctx *gin.Context) {
	var order models.Order
	id := ctx.Param("id")
	err := r.DB.Delete(&order, id).Error
	r.DB.Model(&order).Association("Item").Clear()

	_ = err
}

func (r *OrderRepo) FindById(ctx *gin.Context) {
	var order models.Order

	id := ctx.Param("id")
	err := r.DB.Preload("Items").First(&order, id).Error
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	ctx.JSON(http.StatusOK, order)

}
