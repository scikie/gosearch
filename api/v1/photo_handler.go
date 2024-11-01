package v1

import (
	"gosearch/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"fmt"
	//"strconv"
)

// AddPhoto 添加照片
func AddPhoto(c *gin.Context) {
	 
	// 获取表单中的其他字段
	// album := c.PostForm("album")
	photoName := c.PostForm("photo_name")
	photoDescription := c.PostForm("photo_description")
	// 从请求中获取文件
	file, err := c.FormFile("photo")
	if err!=nil{
		logrus.Fatalln(err)
		return
	}
	// 将文件保存在服务器上
	// 这里的"./images"是保存文件的目录，需要提前创建好
	err = c.SaveUploadedFile(file, "./images/"+file.Filename)
	if err != nil {
		logrus.Fatalln(err)
		c.String(http.StatusInternalServerError, "上传失败")
		return
	}
	data :=model.Photo{
		AlbumID: 1,
		PhotoName: photoName,
		PhotoDescription: photoDescription,
		Status: true,
		IsPublic: true,
		PhotoPath: fmt.Sprintf("./images/%s",file.Filename),
	}
	err=model.InsertPhoto(data)
	if err!=nil{
		logrus.Fatalln(err)
	}
	c.String(http.StatusOK, "文件上传成功: %s", file.Filename)
}

// GetPhoto 获取照片
func GetPhoto(c *gin.Context) {
	/* 
	1.获取查询关键词
	2.通过关键词在数据库中查找
	3.根据查找的数据的情况，放回json
	*/
	searchStr:=c.Query("s")
	photos,err:=model.SelectPhoto(searchStr)
	if err!=nil{
		logrus.Fatalln(err)
	}
	// 将查询结果返回给客户端
	c.JSON(http.StatusOK, photos)
}

