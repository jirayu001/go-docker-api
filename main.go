package main

import (
	_ "github.com/api-server/docs"
	"github.com/api-server/handler"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	r := gin.Default()

	r.GET("/healthcheck", handler.HealthCheckHandler)

	v1 := r.Group("/auth")
	{
		customers := v1.Group("/loginWithCoulacare")
		{
			customersHandler := handler.CustomerHandler{}
			customers.POST("", customersHandler.Post)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
