package controller

import (
	"net/http"
	"PJ/backend/config"
	"PJ/backend/entity"
	"github.com/gin-gonic/gin"
)

// POST /members
func CreateMember(c *gin.Context) {
	var member entity.Member

	// Bind JSON data to member entity
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve DB connection
	db := config.DB()

	// Validate the associated gender
	var gender entity.Gender
	if err := db.First(&gender, member.GenderID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gender not found"})
		return
	}

	// Hash the password
	hashedPassword, err := config.HashPassword(member.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create Member record
	m := entity.Member{
		FirstName: member.FirstName,
		LastName:  member.LastName,
		Email:     member.Email,
		Username:  member.Username,
		Password:  hashedPassword,
		GenderID:  member.GenderID,
		
	}

	if err := db.Create(&m).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Member created successfully", "data": m})
}

// GET /members/:id
func GetMember(c *gin.Context) {
	id := c.Param("id")
	var member entity.Member

	// Retrieve DB connection
	db := config.DB()

	// Fetch member record by ID
	if err := db.Preload("Gender").Preload("Subscriptions").First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	c.JSON(http.StatusOK, member)
}

// GET /members
func ListMembers(c *gin.Context) {
	var members []entity.Member

	// Retrieve DB connection
	db := config.DB()

	// Fetch all member records
	if err := db.Preload("Gender").Preload("Subscriptions").Find(&members).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}

// PATCH /members/:id
func UpdateMember(c *gin.Context) {
	id := c.Param("id")
	var member entity.Member

	// Retrieve DB connection
	db := config.DB()

	// Fetch member record by ID
	if err := db.First(&member, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	// Bind JSON data to existing record
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Optionally hash new password if provided
	if member.Password != "" {
		hashedPassword, err := config.HashPassword(member.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		member.Password = hashedPassword
	}

	// Update the member record
	if err := db.Save(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member updated successfully", "data": member})
}

// DELETE /members/:id
func DeleteMember(c *gin.Context) {
	id := c.Param("id")

	// Retrieve DB connection
	db := config.DB()

	// Delete the member record
	if tx := db.Delete(&entity.Member{}, id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Member not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member deleted successfully"})
}
