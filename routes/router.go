package routes

import (
	"github.com/gin-gonic/gin"
	"gosearch/utils"
	"net/http"
	"gosearch/api/v1"
)
// CORSMiddleware 设置CORS相关的响应头
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许访问的域名，*代表允许任何域名访问
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		// 设置允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		// 设置允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token")
		// 设置预检请求的有效期
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		// 如果是预检请求，则直接返回204
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		// 继续处理请求
		c.Next()
	}
}

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")
	r.Static("/images", "./images")
	// http://localhost:11333/
	r.Use(CORSMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "gotweet is running!")
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})

	/*
		前端展示页面接口
	*/
	router := r.Group("v1")
	{	
		// **用户表（users）**
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("user",v1.GetUserInfoByName)
		router.GET("users", v1.GetUsers)

		
		router.POST("/photos/upload", v1.AddPhoto)
		router.GET("search", v1.GetPhoto)
		// router.GET("users", v1.GetUsers)

		// // 文章分类信息模块
		// router.GET("category", v1.GetCate)
		// router.GET("category/:id", v1.GetCateInfo)

		// // 文章模块
		// router.GET("article", v1.GetArt)
		// router.GET("article/list/:id", v1.GetCateArt)
		// router.GET("article/info/:id", v1.GetArtInfo)

		// // 登录控制模块
		// router.POST("login", v1.Login)
		// router.POST("loginfront", v1.LoginFront)

		// // 获取个人设置信息
		// router.GET("profile/:id", v1.GetProfile)

		// // 评论模块
		// router.POST("addcomment", v1.AddComment)
		// router.GET("comment/info/:id", v1.GetComment)
		// router.GET("commentfront/:id", v1.GetCommentListFront)
		// router.GET("commentcount/:id", v1.GetCommentCount)
	}

	_ = r.Run(utils.HttpPort)

}
