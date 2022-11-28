package Routes

import (
	"second-api/Controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
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
