// controllers/issues.go

package controllers

import (
    "fmt"
    "net/http"
    "database/sql"
    "github.com/gin-gonic/gin"
    //"github.com/act-up/api/models"
)

// Get all issues
func GetIssues(c *gin.Context) {

    var count int

    db := c.MustGet("db").(*sql.DB)

    row := db.QueryRow("SELECT COUNT(*) FROM active_issues")
	err := row.Scan(&count)

    fmt.Println(err)

    defer db.Close()

   //active_issues, err := db.Query("SELECT id, first_name FROM active_issues LIMIT $1", 3)

    // If database returns an error
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No records found!"})
    } else {
        c.JSON(http.StatusOK, "hello")
    }



    /*var active_issues []models.Issue
    err := models.GetAllIssues(db, &active_issues)

    if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	} else {
		c.JSON(http.StatusOK, active_issues)
	}*/

    //db.Find(&active_issues)
    //c.JSON(http.StatusOK, gin.H{"data": active_issues})

}


// Get an issue by ID
/*func GetIssue(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")

	var issue models.Issue
	err := models.GetAnIssue(db, &issue, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	} else {
	     c.JSON(http.StatusOK, issue)
	}
}*/
