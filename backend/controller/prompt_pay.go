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

	// สร้าง promptpays
	prom := entity.PromptPay{
		PromptpayNumber: promptPay.PromptpayNumber,
			
	}

	// บันทึก
	if err := db.Create(&prom).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "PromptPay created successfully", "data": prom})
}

// GET /promptpays/:id
func GetPromptPay(c *gin.Context) {
	ID := c.Param("id")
	var promptPay entity.PromptPay

	db := config.DB()
	results := db.First(&promptPay, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if promptPay.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, promptPay)
}
// GET /promptpays
func ListPromptpay(c *gin.Context) {

	var promptpay []entity.PromptPay

	db := config.DB()
	results := db.Find(&promptpay)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, promptpay)
}

// PATCH /promptpays/:id
func UpdatePromptpay(c *gin.Context) {
	var Promptpay entity.PromptPay

	PromID := c.Param("id")

	db := config.DB()
	result := db.First(&Promptpay, PromID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&Promptpay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&Promptpay)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}

// DELETE /promptpays/:id
func DeletePromtpay(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM Promtpay WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}
