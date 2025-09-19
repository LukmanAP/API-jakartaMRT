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

// GetAllStation godoc
// @Summary Daftar semua stasiun MRT
// @Description Mengambil seluruh data stasiun MRT Jakarta
// @Tags Stations
// @Accept json
// @Produce json
// @Success 200 {object} GetStationsAPIResponse
// @Failure 400 {object} ErrorAPIResponse
// @Router /stations [get]
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

// CheckScedulesByStation godoc
// @Summary Cek jadwal keberangkatan berdasarkan stasiun
// @Description Mengambil jadwal keberangkatan berikutnya untuk dua jurusan dari stasiun terpilih
// @Tags Stations
// @Accept json
// @Produce json
// @Param id path string true "Station ID (nid)"
// @Success 200 {object} GetScedulesAPIResponse
// @Failure 400 {object} ErrorAPIResponse
// @Router /stations/{id} [get]
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
