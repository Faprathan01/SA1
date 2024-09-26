package controller

import (
	"net/http"
	"PJ/backend/config"
	"PJ/backend/entity"
	"github.com/gin-gonic/gin"
)

// POST /payment
func CreatePayment(c *gin.Context) {
	var pay entity.Payment

	// bind เข้าตัวแปร payment
	if err := c.ShouldBindJSON(&pay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()
	

	// สร้าง payment
	pm := entity.Payment{
		PaymentMethodName: pay.PaymentMethodName,
		Amount: pay.Amount,
		
	}

	// บันทึก
	if err := db.Create(&pm).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created success", "data": pm})
}

// GET /payment/:id
func GetPayment(c *gin.Context) {
	ID := c.Param("id")
	var pay entity.Payment

	db := config.DB()
	results := db.First(&pay, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if pay.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	c.JSON(http.StatusOK, pay)
}

// GET /payments
func ListPayment(c *gin.Context) {

	var pays []entity.Payment

	db := config.DB()
	results := db.Find(&pays)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, pays)
}


// PATCH /payments/:id
func UpdatePayment(c *gin.Context) {
	var pay entity.Payment

	PayID := c.Param("id")

	db := config.DB()
	result := db.First(&pay, PayID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not found"})
		return
	}

	if err := c.ShouldBindJSON(&pay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&pay)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successful"})
}
// DELETE /payments/:id
func DeletePayment(c *gin.Context) {

	id := c.Param("id")
	db := config.DB()
	if tx := db.Exec("DELETE FROM payments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted successful"})

}