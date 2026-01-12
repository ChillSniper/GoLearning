package controller

import (
	"myDemo/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义接收前端参数的结构体
// json 表示接收前端传过来的数据，而 binding 表示校验标签
type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var req UserRequest
	// 参数绑定与校验
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误: " + err.Error()})
		return
	}

	if err := service.Register(req.Username, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

func Login(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	token, err := service.Login(req.Username, req.Password)
	if err != nil {
		// 统一返回 401 未授权
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": token})
}
