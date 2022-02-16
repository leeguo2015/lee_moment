package url

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiInclude(c *gin.RouterGroup) {

	c.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  400,
			"message": "",
			"data":    nil,
		})
	})
	userUrl := c.Group("/user")
	userInclude(userUrl)
	momentUrl := c.Group("/moment")
	momentInclude(momentUrl)
}
