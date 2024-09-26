package controller

import (
	"net/http"
	"PJ/backend/config"
	"PJ/backend/entity"
	"github.com/gin-gonic/gin"
)

// POST /payments
func CreatePayment(c *gin.Context) {
	var payment entity.Payment

	// Bind JSON data to payment entity
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()

	

	// Create Payment record
	if err := db.Create(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Payment created successfully", "data": payment})
}

// GET /payments/:id
func GetPayment(c *gin.Context) {
	id := c.Param("id")
	var payment entity.Payment

	db := config.DB()

	// Fetch payment record by ID, preload related entities
	if err := db.Preload("CreditCards").Preload("Paypals").Preload("PromptPays").Preload("Subscription").First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// GET /payments
func ListPayments(c *gin.Context) {
	var payments []entity.Payment

	db := config.DB()

	// Fetch all payment records, preload related entities
	if err := db.Preload("CreditCards").Preload("Paypals").Preload("PromptPays").Preload("Subscription").Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}

// PATCH /payments/:id
func UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	var payment entity.Payment

	db := config.DB()

	// Fetch payment record by ID
	if err := db.First(&payment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	// Bind JSON data to existing record
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	

	// Update the payment record
	if err := db.Save(&payment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment updated successfully", "data": payment})
}

// DELETE /payments/:id
func DeletePayment(c *gin.Context) {
	id := c.Param("id")

	db := config.DB()

	// Delete the payment record
	if tx := db.Delete(&entity.Payment{}, id); tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment deleted successfully"})
}
