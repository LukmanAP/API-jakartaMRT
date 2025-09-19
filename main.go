package main

import (
	"jadwal-mrt/modules/station"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "jadwal-mrt/docs"
)

// @title Jadwal MRT API
// @version 1.0
// @description API untuk mendapatkan daftar stasiun dan jadwal keberangkatan MRT Jakarta.
// @host localhost:8080
// @BasePath /v1/api
func main() {
	InitiateRouter()
}

func InitiateRouter() {
	var (
		router = gin.Default()
		api    = router.Group("/v1/api")
	)

	// set swagger base path
	docs.SwaggerInfo.BasePath = "/v1/api"

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	station.Initiate(api)

	router.Run()
}
