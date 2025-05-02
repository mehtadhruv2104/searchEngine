package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mehtadhruv2104/searchEngine/backend/handler"
)


func StartEngine(h *handler.Handler) (*gin.Engine){
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Range", "X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	apiRoutes := router.Group("/api")
	// set up middleware for further improvements 
	//apiRoutes.Use(Middleware)

	apiRoutes.POST("/search",h.HandleSearch)
	return router




}
