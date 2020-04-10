// controllers/issues.go

package controllers

import (
    "net/http"
    "github.com/jinzhu/gorm"
    "github.com/gin-gonic/gin"
    "github.com/act-up/backend/api/models"
)

// Get all issues
func GetIssues(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    var active_issues []models.Issue
    err := models.GetAllIssues(db, &active_issues)

    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	} else {
		c.JSON(http.StatusOK, active_issues)
	}

    //db.Find(&active_issues)
    //c.JSON(http.StatusOK, gin.H{"data": active_issues})

}


// Get an issue by ID
func GetIssue(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")

	var issue models.Issue
	err := models.GetAnIssue(db, &issue, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	} else {
	     c.JSON(http.StatusOK, issue)
	}
}
