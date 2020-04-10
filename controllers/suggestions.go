// controllers/suggestions.go

package controllers

import (
    "net/http"
    "github.com/jinzhu/gorm"
    "github.com/gin-gonic/gin"
    "github.com/act-up/api/models"
)

// Create a Suggestion
func CreateSuggestion(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

	var suggestion models.Suggestion
	c.BindJSON(&suggestion)
	err := models.CreateASuggestion(db, &suggestion)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, suggestion)
        return
	}
}
