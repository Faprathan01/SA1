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

	// bind เข้าตัวแปร user
	if err := c.ShouldBindJSON(&paypal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := config.DB()


	// สร้าง paypals
	pp := entity.Paypal{
		PaypalEmail: paypal.PaypalEmail, 
		PaypalPassword: paypal.PaypalPassword,
     
		
	}
	// บันทึก
	if err := db.Create(&pp).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Paypal created successfully", "data": pp})
}

// GET /paypals/:id
func GetPaypal(c *gin.Context) {
	ID := c.Param("id")
	var paypal entity.Paypal

	db := config.DB()
	results := db.First(&paypal, ID)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	if paypal.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, paypal)
}

// GET /paypals
func ListPaypals(c *gin.Context) {
	var paypals []entity.Paypal

	db := config.DB()

	results := db.Find(&paypals)
	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, paypals)
}


// PATCH /paypals/:id
func UpdatePaypal(c *gin.Context) {
	var pay entity.Paypal

	PacID := c.Param("id")

	db := config.DB()
	result := db.First(&pay, PacID)
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

// DELETE /paypals/:id
func DeletePaypal(c *gin.Context) {
	id := c.Param("id")

	db := config.DB()

	// Delete the Paypal record
	if tx := db.Delete(&entity.Paypal{}, id); tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Paypal not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Paypal deleted successfully"})
}
