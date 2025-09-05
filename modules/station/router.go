package station

import (
	"jadwal-mrt/common/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/stations")

	station.GET("", func(c *gin.Context) {
		GetAllStation(c, stationService)
	})
	station.GET("/:id", func(c *gin.Context) {
		CheckScedulesByStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStation()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIresponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, response.APIresponse{
		Success: true,
		Message: "Successfully get all station",
		Data:    datas,
	})
}

func CheckScedulesByStation(c *gin.Context, service Service) {
	id := c.Param("id")

	datas, err := service.CheckScedulesByStation(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.APIresponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, response.APIresponse{
		Success: true,
		Message: "Successfully get scedules station",
		Data:    datas,
	})
}
