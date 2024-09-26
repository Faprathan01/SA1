package controller

import (
	"net/http"
	"PJ/backend/config"
	"PJ/backend/entity"
	"github.com/gin-gonic/gin"
)
// POST /Creditcard
func CreateCreditcard(c *gin.Context) {
	var creditCard entity.CreditCard

	// Bind JSON data to PromptPay entity
	if err := c.ShouldBindJSON(&creditCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// Retrieve DB connection
	db := config.DB()

	// สร้าง Creditcard
	cc := entity.CreditCard{
		NameOnCard: creditCard.NameOnCard,
		CardNumber: creditCard.CardNumber,
		ExpiryDate: creditCard.ExpiryDate,
		CVV:        creditCard.CVV,
	}

	// บันทึก
	if err := db.Create(&cc).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "PromptPay created successfully", "data": cc})
}
// GET /Creditcard/:id
func GetCreditcard(c *gin.Context) {
	ID := c.Param("id")
	var creditCard entity.CreditCard

	db := config.DB()
	results := db.First(&creditCard, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if creditCard.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, creditCard)
}

// GET /Creditcard
func ListCreditcard(c *gin.Context) {

	var creditCard []entity.CreditCard

	db := config.DB()
	results := db.Find(&creditCard)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, creditCard)
}

// PATCH /Creditcards/:id
func UpdateCreditcard(c *gin.Context) {
	var creditCard entity.CreditCard

	ccID := c.Param("id")

	db := config.DB()
	result := db.First(&creditCard, ccID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&creditCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&creditCard)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}

// DELETE /Creditcards/:id
func DeleteCreditcard(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM Creditcard WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}