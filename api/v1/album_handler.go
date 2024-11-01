package v1

import (
	"gosearch/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"strconv"
)

/*
添加相册
1.检查数据库中是否存在该相册名

	（1.存在，则返回
	（2.不存在，继续

2.在数据库中添加一个相册记录
*/
func AddAlbum(c *gin.Context) {
	var data model.Album
	 // 获取Content-Type
	contentType := c.Request.Header.Get("Content-Type")
	// 根据Content-Type处理不同的请求体
	switch contentType {
	case "application/json":
		// JSON请求处理
		if err:= c.ShouldBindJSON(&data);err!=nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON binding error", "message": err.Error()})
			return
		}

	case "application/x-www-form-urlencoded", "multipart/form-data":
		// 表单请求处理
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Form binding error", "message": err.Error()})
			return
		}
	default:
		// 不支持的Content-Type
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported Content-Type"})
		return
	}

	// logrus.Debugln(data)
	code, err := model.CheckAlbum(data.Albumname)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, gin.H{
				"status":  code,
				"message": err.Error(),
			},
		)
	} else {
		model.CreateAlbum(&data)
		c.JSON(
			http.StatusOK, gin.H{
				"status":  code,
				"message": "",
			},
		)
	}
}

/*
	查询Param=id的相册
*/
func GetAlbumInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.JSON(
			http.StatusBadRequest, gin.H{
				"message": err.Error(),
			},
		)
	}
	data, code, err := model.GetAlbum(id)
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
/*
	查看Query=userid的用户拥有哪些相册
*/
func GetAlbumInfoByUser(c *gin.Context) {
	userid,err:= strconv.Atoi(c.Query("userid"))
	if err!=nil{
		c.JSON(
			http.StatusBadRequest, gin.H{
				"message": err.Error(),
			},
		)
	}

	data, code, err := model.GetAlbumByUser(userid)
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


