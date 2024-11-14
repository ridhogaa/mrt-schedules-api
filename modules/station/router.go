package station

import (
	"github.com/gin-gonic/gin"
	"mrt-schedules-api/common/response"
	"net/http"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()
	station := router.Group("/stations")
	station.GET("", func(c *gin.Context) {
		GetAllStation(c, stationService)
	})
	station.GET(":id", func(c *gin.Context) {
		GetScheduleByStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	data, err := service.GetAllStation()
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: err.Error(),
			Success: false,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		Message: "Successfully Get All Stations",
		Success: true,
		Data:    data,
	})
}

func GetScheduleByStation(c *gin.Context, service Service) {
	id := c.Param("id")

	data, err := service.GetScheduleByStation(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: err.Error(),
			Success: false,
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.BaseResponse{
		Message: "Successfully Get Schedules by Station",
		Success: true,
		Data:    data,
	})
}
