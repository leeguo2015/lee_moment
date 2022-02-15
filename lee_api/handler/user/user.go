package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"lee_moment/lee_api/model"
	"net/http"
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
