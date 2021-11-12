package user

import (
	"fmt"
	"lee_moment/app/web/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	fmt.Println(c.Param("userId"))
	selectUser := model.User{}
	c.JSON(http.StatusOK, gin.H{
		"status":  400,
		"message": "",
		"data":    selectUser,
	})
}

func Post(c *gin.Context) {
	c.String(http.StatusOK, "User!")
}
