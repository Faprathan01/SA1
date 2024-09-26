package controller

import (
	"net/http"
	"PJ/backend/config"
	"PJ/backend/entity"
	"github.com/gin-gonic/gin"
)

// POST /promptpays
func CreatePromptPay(c *gin.Context) {
	var promptPay entity.PromptPay

	// Bind JSON data to PromptPay entity
	if err := c.ShouldBindJSON(&promptPay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Retrieve DB connection
	db := config.DB()

	// Validate associated PaymentID
	if promptPay.PaymentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PaymentID"})
		return
	}

	var payment entity.Payment
	if err := db.First(&payment, promptPay.PaymentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	// Create PromptPay record
	if err := db.Create(&promptPay).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create PromptPay record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "PromptPay created successfully", "data": promptPay})
}

// GET /promptpays/:id
func GetPromptPay(c *gin.Context) {
	id := c.Param("id")
	var promptPay entity.PromptPay

	// Retrieve DB connection
	db := config.DB()

	// Fetch PromptPay record by ID
	if err := db.Preload("Payment").First(&promptPay, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "PromptPay not found"})
		return
	}

	c.JSON(http.StatusOK, promptPay)
}

// GET /promptpays
func ListPromptPay(c *gin.Context) {
	var promptPays []entity.PromptPay

	// Retrieve DB connection
	db := config.DB()

	// Fetch all PromptPay records
	if err := db.Preload("Payment").Find(&promptPays).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch PromptPays"})
		return
	}

	c.JSON(http.StatusOK, promptPays)
}

// PATCH /promptpays/:id
func UpdatePromptPay(c *gin.Context) {
	id := c.Param("id")
	var promptPay entity.PromptPay

	// Retrieve DB connection
	db := config.DB()

	// Fetch PromptPay record by ID
	if err := db.First(&promptPay, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "PromptPay not found"})
		return
	}

	// Bind JSON data to existing record
	if err := c.ShouldBindJSON(&promptPay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Optionally validate associated PaymentID
	if promptPay.PaymentID > 0 {
		var payment entity.Payment
		if err := db.First(&payment, promptPay.PaymentID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
			return
		}
	}

	// Update the PromptPay record
	if err := db.Save(&promptPay).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update PromptPay record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "PromptPay updated successfully", "data": promptPay})
}

// DELETE /promptpays/:id
func DeletePromptPay(c *gin.Context) {
	id := c.Param("id")

	// Retrieve DB connection
	db := config.DB()

	// Delete the PromptPay record
	if tx := db.Delete(&entity.PromptPay{}, id); tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "PromptPay not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "PromptPay deleted successfully"})
}
