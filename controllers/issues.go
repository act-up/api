// controllers/issues.go

package controllers

import (
            "fmt"
            "net/http"
            "database/sql"
            "strconv"
            "github.com/gin-gonic/gin"
            "github.com/act-up/api/models"
)

// Get all issues
func GetIssues(c *gin.Context) {

    db := c.MustGet("db").(*sql.DB)

    //var active_issues []*models.Issue{}
    err := models.GetAllIssues(db)//, &active_issues)

    defer db.Close()


    // If database returns an error
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Zero rows found!"})
        } else {
            c.JSON(http.StatusBadRequest, gin.H{"error": err})
        }

    } else {
        c.JSON(http.StatusOK, "fuck this shit")
    }

}


// Get an issue by ID
func GetIssue(c *gin.Context) {
    db := c.MustGet("db").(*sql.DB)
	id_str := c.Params.ByName("id")

    id, err := strconv.Atoi(id_str)
	if err != nil {
		fmt.Println(err)
	}

	var issue models.Issue
	id, err = models.GetAnIssue(db, &issue, id)

    c.JSON(http.StatusOK, id)

    defer db.Close()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	} else {
	     c.JSON(http.StatusOK, issue)
	}
}
