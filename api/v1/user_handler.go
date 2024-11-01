package v1

import (
	"gosearch/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

/*
添加用户
1.检查数据库中是否存在该用户名

	（1.存在，则返回
	（2.不存在，继续

2.在数据库中添加一个用户记录
*/
func AddUser(c *gin.Context) {
	var data model.User
	if err:= c.ShouldBindJSON(&data);err!=nil{
		c.JSON(
			http.StatusBadRequest, gin.H{
				"status":  400,
				"message": err.Error(),
			},
		)
		return
	}
	// logrus.Debugln(data)
	code, err := model.CheckUser(data.Username)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"status":  code,
				"message": err.Error(),
			},
		)
	} else {
		model.CreateUser(&data)
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"message": "",
			},
		)
	}
}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code, err := model.GetUser(id)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"status":  code,
				"data":    "",
				"total":   1,
				"message": err.Error(),
			},
		)
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   1,
			"message": "",
		},
	)

}

func GetUserInfoByName(c *gin.Context) {
	usename:= c.Query("username")

	data, code, err := model.GetUserByName(usename)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"status":  code,
				"data":    "",
				"total":   0,
				"message": err.Error(),
			},
		)
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   1,
			"message": "",
		},
	)

}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, err := strconv.Atoi(c.Query("pagesize"))
	if err!=nil{
		c.JSON(
			http.StatusOK, gin.H{
				"message": err.Error(),
			},
		)
	}
	pageNum, err := strconv.Atoi(c.Query("pagenum"))
	if err!=nil{
		c.JSON(
			http.StatusOK, gin.H{
				"message": err.Error(),
			},
		)
	}
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total, code, err := model.GetUsers(username, pageSize, pageNum)
	if code == 200 && err != nil {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"data":    data,
				"total":   total,
				"message": err.Error(),
			},
		)
	} else {
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"data":    data,
				"total":   total,
				"message": "",
			},
		)
	}

}
