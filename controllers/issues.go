// controllers/issues.go

package controllers

import (
            "net/http"
            "database/sql"
            "github.com/gin-gonic/gin"
            "github.com/act-up/api/models"
)

// Get all issues
func GetIssues(c *gin.Context) {

    db := c.MustGet("db").(*sql.DB)

    //var active_issues []*models.Issue{}
    active_issues, err := models.GetAllIssues(db)//, &active_issues)

    defer db.Close()


    // If database returns an error
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Zero rows found!"})
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err})
        }

    } else {
        c.JSON(http.StatusOK, active_issues)
    }

}


// Get an issue by ID
func GetIssue(c *gin.Context) {
    db := c.MustGet("db").(*sql.DB)
	id := c.Params.ByName("id")

	var issue models.Issue
	err := models.GetAnIssue(db, &issue, id)

    defer db.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
	     c.JSON(http.StatusOK, issue)
	}
}
