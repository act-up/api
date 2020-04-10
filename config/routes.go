// config/routes.go

package config

import (
            "github.com/gin-gonic/gin"
            "database/sql"
            //"github.com/act-up/backend/api/controllers"
)

func SetupRoutes(db *sql.DB) *gin.Engine {

    // Declare a new router
    r := gin.Default()

    r.Use(func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    })

    g1 := r.Group("/")
    {
        // List all issues in database
        /*g1.GET("issues", controllers.GetIssues)
        g1.POST("issues/:id", controllers.GetIssue)
        g1.GET("results", controllers.GetResults)
        g1.GET("results/:id", controllers.GetResult)
        g1.PUT("results/:id", controllers.UpdateResult)
        g1.POST("suggest", controllers.CreateSuggestion)*/

    }

    return r

}
