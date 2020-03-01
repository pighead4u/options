package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pighead4u/options/src/api"
)

func main() {
	router := gin.Default()
	router.GET("/title/:bourse", api.GetTitlesByBourse)
	router.GET("/content/:contract", api.GetContentsByContractAndDate)
	router.Run(":9898")
}
