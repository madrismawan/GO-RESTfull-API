package handler

import (
	"example/main.go/helper"
	"example/main.go/models"
	"net/http"

	"github.com/go-playground/validator/v10"

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

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func (r *OrderRepo) CreateOrder(ctx *gin.Context) {
	var order models.Order
	validate = validator.New()
	ctx.ShouldBindJSON(&order)
	err := validate.Struct(order)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Bad Request", http.StatusBadRequest, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	err = r.DB.Create(&order).Error
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	response := helper.APIResponse("Successfully add new order", http.StatusCreated, "success", order)
	ctx.JSON(http.StatusCreated, response)
}

func (r *OrderRepo) GetOrder(ctx *gin.Context) {
	var orders []models.Order
	err := r.DB.Model(&models.Order{}).Preload("Items").Find(&orders).Error
	if err != nil {
		ctx.JSON(http.StatusNoContent, err)
		return
	}
	response := helper.APIResponse("Successfully get orders", http.StatusOK, "success", orders)
	ctx.JSON(http.StatusCreated, response)
}

func (r *OrderRepo) UpdateOrder(ctx *gin.Context) {
	var order models.Order
	validate = validator.New()

	id := ctx.Param("id")
	err := r.DB.First(&order, id).Error
	if err != nil {
		response := helper.APIResponse("Not found ", http.StatusNotFound, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ctx.ShouldBindJSON(&order)
	errInput := validate.Struct(order)
	if errInput != nil {
		errors := helper.FormatValidationError(errInput)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Bad Request", http.StatusBadRequest, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	errParent := r.DB.Save(&order).Error
	if errParent != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	errPivot := r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order).Error
	if errPivot != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	response := helper.APIResponse("Successfully update order", http.StatusOK, "success", order)
	ctx.JSON(http.StatusCreated, response)
}

func (r *OrderRepo) DeleteOrder(ctx *gin.Context) {
	var order models.Order
	id := ctx.Param("id")
	err := r.DB.First(&order, id).Error
	// err := r.DB.Where("order_id = ?", id).Delete(&order)
	if err != nil {
		response := helper.APIResponse("Not found ", http.StatusNotFound, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	r.DB.Model(&order).Association("Items").Clear()
	r.DB.Delete(&order)

	response := helper.APIResponse("Successfully delete order", http.StatusOK, "success", nil)
	ctx.JSON(http.StatusOK, response)
}

func (r *OrderRepo) FindById(ctx *gin.Context) {
	var order models.Order
	id := ctx.Param("id")
	err := r.DB.Preload("Items").First(&order, id).Error
	if err != nil {
		response := helper.APIResponse("Not found ", http.StatusNotFound, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIResponse("Successfully get order", http.StatusOK, "success", order)
	ctx.JSON(http.StatusCreated, response)
}
