package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mehtadhruv2104/searchEngine/backend/handler"
)


func StartEngine(h *handler.Handler) (*gin.Engine){
	router := gin.Default()
	apiRoutes := router.Group("/api")
	// set up middleware for further improvements 
	//apiRoutes.Use(Middleware)

	apiRoutes.POST("/search",h.HandleSearch)
	return router



	
}
