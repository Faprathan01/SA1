package controller

import (
	"net/http"
	"PJ/backend/config"
	"PJ/backend/entity"
	"github.com/gin-gonic/gin"
)

// POST /paypals
func CreatePaypal(c *gin.Context) {
	var paypal entity.Paypal

	// Bind JSON data to Paypal entity
	if err := c.ShouldBindJSON(&paypal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Retrieve DB connection
	db := config.DB()

	// Validate associated PaymentID
	if paypal.PaymentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PaymentID"})
		return
	}

	var payment entity.Payment
	if err := db.First(&payment, paypal.PaymentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	// Create Paypal record
	if err := db.Create(&paypal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Paypal record"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Paypal created successfully", "data": paypal})
}

// GET /paypals/:id
func GetPaypal(c *gin.Context) {
	id := c.Param("id")
	var paypal entity.Paypal

	// Retrieve DB connection
	db := config.DB()

	// Fetch Paypal record by ID
	if err := db.Preload("Payment").First(&paypal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paypal not found"})
		return
	}

	c.JSON(http.StatusOK, paypal)
}

// GET /paypals
func ListPaypal(c *gin.Context) {
	var paypals []entity.Paypal

	// Retrieve DB connection
	db := config.DB()

	// Fetch all Paypal records
	if err := db.Preload("Payment").Find(&paypals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Paypals"})
		return
	}

	c.JSON(http.StatusOK, paypals)
}

// PATCH /paypals/:id
func UpdatePaypal(c *gin.Context) {
	id := c.Param("id")
	var paypal entity.Paypal

	// Retrieve DB connection
	db := config.DB()

	// Fetch Paypal record by ID
	if err := db.First(&paypal, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paypal not found"})
		return
	}

	// Bind JSON data to existing record
	if err := c.ShouldBindJSON(&paypal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Optionally validate associated PaymentID
	if paypal.PaymentID > 0 {
		var payment entity.Payment
		if err := db.First(&payment, paypal.PaymentID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
			return
		}
	}

	// Update the Paypal record
	if err := db.Save(&paypal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Paypal record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Paypal updated successfully", "data": paypal})
}

// DELETE /paypals/:id
func DeletePaypal(c *gin.Context) {
	id := c.Param("id")

	// Retrieve DB connection
	db := config.DB()

	// Delete the Paypal record
	if tx := db.Delete(&entity.Paypal{}, id); tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paypal not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Paypal deleted successfully"})
}
