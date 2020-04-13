package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pighead4u/options/src/api"
	"github.com/unrolled/secure"
)

func main() {
	router := gin.Default()

	router.Use(TlsHandler())
	router.GET("/title/:bourse", api.GetTitlesByBourse)
	router.GET("/content/:contract", api.GetContentsByContractAndDate)
	router.Run(":9898")
}

func TlsHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        secureMiddleware := secure.New(secure.Options{
            SSLRedirect: true,
            SSLHost:     "localhost:80",
        })
        err := secureMiddleware.Process(c.Writer, c.Request)

        // If there was an error, do not continue.
        if err != nil {
            return
        }

        c.Next()
    }
}
