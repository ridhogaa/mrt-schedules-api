package main

import (
	"github.com/gin-gonic/gin"
	"mrt-schedules-api/modules/station"
)

func main() {
	InitiateRouter()
}

func InitiateRouter() {
	var (
		router = gin.Default()
		api    = router.Group("api/v1/")
	)

	station.Initiate(api)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
