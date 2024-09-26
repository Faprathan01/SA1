package controller

import (
	"net/http"
	"PJ/backend/config"
	"PJ/backend/entity"
	"github.com/gin-gonic/gin"
)

// POST /packages
func CreatePackage(c *gin.Context) {
	var pkg entity.Package

	// Bind JSON data to package entity
	if err := c.ShouldBindJSON(&pkg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve DB connection
	db := config.DB()

	// Create Package record
	if err := db.Create(&pkg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Package created successfully", "data": pkg})
}

// GET /packages/:id
func GetPackage(c *gin.Context) {
	id := c.Param("id")
	var pkg entity.Package

	// Retrieve DB connection
	db := config.DB()

	// Fetch package record by ID
	if err := db.First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
		return
	}

	c.JSON(http.StatusOK, pkg)
}

// GET /packages
func ListPackages(c *gin.Context) {
	var packages []entity.Package

	// Retrieve DB connection
	db := config.DB()

	// Fetch all package records
	if err := db.Find(&packages).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, packages)
}

// PATCH /packages/:id
func UpdatePackage(c *gin.Context) {
	id := c.Param("id")
	var pkg entity.Package

	// Retrieve DB connection
	db := config.DB()

	// Fetch package record by ID
	if err := db.First(&pkg, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Package not found"})
		return
	}

	// Bind JSON data to existing record
	if err := c.ShouldBindJSON(&pkg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the package record
	if err := db.Save(&pkg).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package updated successfully", "data": pkg})
}

// DELETE /packages/:id
func DeletePackage(c *gin.Context) {
	id := c.Param("id")

	// Retrieve DB connection
	db := config.DB()

	// Delete the package record
	if tx := db.Delete(&entity.Package{}, id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Package not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Package deleted successfully"})
}
