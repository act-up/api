// controllers/statistics.go


package controllers

import (
            "net/http"
            "github.com/jinzhu/gorm"
            "github.com/gin-gonic/gin"
            "github.com/act-up/api/models"
)

// GET: Get statistics for an issue by ID
func GetIssueStatistics(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")

	var stats models.Statistics
	err := models.GetIssueStatistics(db, &stats, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, stats)
	}
}


// PATCH: Update an issue statistics record when click event logged
func UpdateStatistics(c *gin.Context) {

    db := c.MustGet("db").(*gorm.DB)

	var stats models.Statistics
	id := c.Params.ByName("id")
	err := models.GetIssueStatistics(db, &stats, id)

	if err != nil {
		c.JSON(http.StatusNotFound, stats)
	}

	c.BindJSON(&stats)
	err = models.UpdateStatistics(db, &stats, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"data": stats})
	}
}
