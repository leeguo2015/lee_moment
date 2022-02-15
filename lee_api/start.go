package lee_api

import (
	"github.com/gin-gonic/gin"
	"lee_moment/lee_api/url"
	"net/http"
)

func StartServer() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": "",
			"data":    nil,
		})
	})
	//404 Handler.
	r.NoRoute(func(c *gin.Context) {
		//c.String(http.StatusNotFound, "the incorrect API route.")
		{
			c.JSON(http.StatusOK, gin.H{
				"status":  400,
				"message": "the incorrect API route.",
				"data":    nil,
			})
		}
	})
	ApiHandler := r.Group("/lee_api/vi")
	url.ApiInclude(ApiHandler)
	_ = r.Run(":8000")
}