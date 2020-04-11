// config/routes.go

package config

import (
            "github.com/gin-gonic/gin"
            _ "github.com/gin-contrib/cors"
            "github.com/jinzhu/gorm"
            "github.com/act-up/api/controllers"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

    // Declare a new router
    r := gin.Default()

    // Write custom http header to allow all origins: TODO only allow www.actup.us
    r.Use(CORS(db))

    //r.Use(cors.Default())   // use default cors policies to allow all origins

    g1 := r.Group("/")
    {
        // List all issues in database
        g1.GET("issues", controllers.GetAllIssues)
        g1.GET("issues/:id", controllers.GetAnIssue)
        g1.GET("statistics/:id", controllers.GetIssueStatistics)
        g1.PATCH("statistics/:id", controllers.UpdateStatistics)
        //g1.POST("suggest", controllers.CreateSuggestion)

    }

    return r

}


func CORS(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

        c.Set("db", db)

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
