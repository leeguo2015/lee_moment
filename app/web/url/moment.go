package url

import "github.com/gin-gonic/gin"

func momentInclude(c *gin.RouterGroup) {
	momentUrl := c.Group("/moment")
	_ = momentUrl

}
