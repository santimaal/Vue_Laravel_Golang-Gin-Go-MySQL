package Routes

import (
	"net/http"
	"sanvic/Table"
	"sanvic/Thematic"
	"sanvic/User"

	"github.com/gin-gonic/gin"
)

func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// c.Header("Access-Control-Max-Age", "3001");
	// c.Header("Access-Control-Allow-Credentials", "true");

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

	api := r.Group("/api")

	// Thematic.ThematicRouting(api.Group("/thematic"))
	// Table.TableRouting(api.Group("/table"))
	// User.UserRouting(api.Group("user/"))
	// Reserve.ReserveRouting(api.Group("reserve/"))

	api.GET("table", Table.GetAllTables)
	// api.Use(User.AuthMiddleware(false))
	api.GET("table/filter", Table.GetTablesFilter)
	api.GET("table/:id", Table.GetTableByID)
	api.GET("thematic", Thematic.GetAllThematics)
	api.GET("thematic/:id", Thematic.GetThematicByID)
	api.GET("thematic/infinite", Thematic.GetThematicsInfinity)
	api.POST("user/register", User.UserRegister)
	api.POST("user/login", User.UserLogin)
	api.GET("user/:id", User.GetUserByID)

	return r
}
