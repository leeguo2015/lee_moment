package url

import (
	"github.com/gin-gonic/gin"
	"lee_moment/app/web/handler/user"
)

func userInclude(c *gin.RouterGroup) {
	baseUrl := c.Group("")
	baseUrl.GET("/:userId", user.Get)

	fansUrl := c.Group("/fans")
	fansUrl.GET("", user.Get)

}
