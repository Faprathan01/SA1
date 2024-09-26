package controller

import (
	"net/http"
	"PJ/backend/config"
	"PJ/backend/entity"
	"github.com/gin-gonic/gin"
)

// POST /credit-cards
func CreateCreditCard(c *gin.Context) {
	var creditCard entity.CreditCard

	// Bind JSON data to creditCard entity
	if err := c.ShouldBindJSON(&creditCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Retrieve DB connection
	db := config.DB()

	// Validate the associated PaymentID
	if creditCard.PaymentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid PaymentID"})
		return
	}

	var payment entity.Payment
	if err := db.First(&payment, creditCard.PaymentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}

	// Create CreditCard record
	creditCard.Payment = payment
	if err := db.Create(&creditCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Credit Card"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Credit Card created successfully", "data": creditCard})
}

// GET /credit-cards/:id
func GetCreditCard(c *gin.Context) {
	id := c.Param("id")
	var creditCard entity.CreditCard

	// Retrieve DB connection
	db := config.DB()

	// Fetch credit card record by ID
	if err := db.Preload("Payment").First(&creditCard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Credit Card not found"})
		return
	}

	c.JSON(http.StatusOK, creditCard)
}

// GET /credit-cards
func ListCreditCard(c *gin.Context) {
	var creditCards []entity.CreditCard

	// Retrieve DB connection
	db := config.DB()

	// Fetch all credit card records
	if err := db.Preload("Payment").Find(&creditCards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Credit Cards"})
		return
	}

	c.JSON(http.StatusOK, creditCards)
}

// PATCH /credit-cards/:id
func UpdateCreditCard(c *gin.Context) {
	id := c.Param("id")
	var creditCard entity.CreditCard

	// Retrieve DB connection
	db := config.DB()

	// Fetch credit card record by ID
	if err := db.First(&creditCard, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Credit Card not found"})
		return
	}

	// Bind JSON data to existing record
	if err := c.ShouldBindJSON(&creditCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Optionally validate the associated PaymentID
	if creditCard.PaymentID > 0 {
		var payment entity.Payment
		if err := db.First(&payment, creditCard.PaymentID).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
			return
		}
	}

	// Update the credit card record
	if err := db.Save(&creditCard).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Credit Card"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Credit Card updated successfully", "data": creditCard})
}

// DELETE /credit-cards/:id
func DeleteCreditCard(c *gin.Context) {
	id := c.Param("id")

	// Retrieve DB connection
	db := config.DB()

	// Delete the credit card record
	if tx := db.Delete(&entity.CreditCard{}, id); tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Credit Card not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Credit Card deleted successfully"})
}
