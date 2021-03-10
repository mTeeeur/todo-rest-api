package handler

import "github.com/gin-gonic/gin"

type Handler struct {

}

func (h *Handler) Init() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/login")
		auth.POST("/register")
	}

	router.Group("/api")
	{
		lists := router.Group("/lists")
		{
			lists.POST("/")
			lists.GET("/")
			lists.GET("/:id")
			lists.PUT("/:id")
			lists.DELETE("/:id")

			items := lists.Group("/:id/items")
			{
				items.GET("/")
				items.POST("/")
				items.POST("/:item_id")
				items.PUT("/:item_id")
				items.DELETE("/:item_id")
			}
		}
	}

	return router
}
