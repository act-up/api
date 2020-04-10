// config/routes.go

package config

import (
            "github.com/gin-gonic/gin"
            "github.com/jinzhu/gorm"
            "github.com/act-up/api/controllers"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

    // Declare a new router
    r := gin.Default()

    r.Use(func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    })

    g1 := r.Group("/")
    {
        // List all issues in database
        g1.GET("issues", controllers.GetAllIssues)
        g1.GET("issues/:id", controllers.GetAnIssue)
        g1.GET("statistics/:id", controllers.GetIssueStatistics)
        //g1.PUT("results/:id", controllers.UpdateResult)
        //g1.POST("suggest", controllers.CreateSuggestion)

    }

    return r

}
