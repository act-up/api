// controllers/issues.go

package controllers

import (
            //"fmt"
            "net/http"
            "github.com/jinzhu/gorm"
            _ "strconv"
            "github.com/gin-gonic/gin"
            "github.com/act-up/api/models"
)

// Get all issues
func GetAllIssues(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)

    var active_issues []models.Issue
    err := models.GetAllIssues(db, &active_issues)

    if err != nil {
        if gorm.IsRecordNotFoundError(err) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err})
        }
    } else {
        c.JSON(http.StatusOK, active_issues)
    }

}


// Get an issue by ID
func GetAnIssue(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)
    id := c.Params.ByName("id")

    var issue models.Issue
    err := models.GetAnIssue(db, &issue, id)

    if err != nil {
        if gorm.IsRecordNotFoundError(err) {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err})
        }
    } else {
         c.JSON(http.StatusOK, issue)
    }

}
