package Routes

import (
	"second-api/Controllers"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}
func SetupRouter() *gin.Engine {

	r := gin.Default()
	r.Use(CORS)

	grp1 := r.Group("/api")
	{
		grp1.GET("table", Controllers.GetTables)
		grp1.POST("table", Controllers.CreateTable)
		grp1.GET("table/:id", Controllers.GetTableByID)
		grp1.PUT("table/:id", Controllers.UpdateTable)
		grp1.DELETE("table/:id", Controllers.DeleteTable)
		grp1.GET("thematic", Controllers.GetThematics)
		grp1.POST("thematic", Controllers.CreateThematic)
		grp1.GET("thematic/:id", Controllers.GetThematicByID)
		grp1.PUT("thematic/:id", Controllers.UpdateThematic)
		grp1.DELETE("thematic/:id", Controllers.DeleteThematic)
	}
	return r
}
