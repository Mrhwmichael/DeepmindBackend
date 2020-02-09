package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleQuery(c *gin.Context){
	//响应查询

	c.String(http.StatusOK, "hello, world")

}