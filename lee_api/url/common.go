package url

import (
	"github.com/gin-gonic/gin"
)

func ApiInclude(c *gin.RouterGroup) {
	userUrl := c.Group("/user")
	userInclude(userUrl)
	momentUrl := c.Group("/moment")
	momentInclude(momentUrl)

}
