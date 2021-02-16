package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	routerV1_Auth := r.Group("api/v1")
	routerV1_Auth.Use(middleware.JwtToken())
	{
		// 用户模块的路由接口
		routerV1_Auth.PUT("user/:id", v1.EditUser)
		routerV1_Auth.DELETE("user/:id", v1.DeleteUser)

		// 分类模块的路由接口
		routerV1_Auth.POST("category/addCategory", v1.AddCategory)
		routerV1_Auth.PUT("category/:id", v1.EditCategory)
		routerV1_Auth.DELETE("category/:id", v1.DeleteCategory)

		// 文章模块的路由接口
		routerV1_Auth.POST("article/addArticle", v1.AddArticle)
		routerV1_Auth.PUT("article/:id", v1.EditArticle)
		routerV1_Auth.DELETE("article/:id", v1.DeleteArticle)

	}

	routerV1_No_Auth := r.Group("api/v1")
	{
		// 用户模块的路由接口
		routerV1_No_Auth.POST("user/addUser", v1.AddUser)
		routerV1_No_Auth.GET("user/queryUsers", v1.GetUsers)

		// 分类模块的路由接口
		routerV1_No_Auth.GET("category/queryCategory", v1.GetCategory)

		// 文章模块的路由接口
		routerV1_No_Auth.GET("article/queryArticle", v1.GetArticle)
		routerV1_No_Auth.GET("article/queryCategoryArticle/:id", v1.GetArtCategory) //查询分类下所有文章
		routerV1_No_Auth.GET("article/info/:id", v1.GetArtInfo)                     // 查询单个文章
		routerV1_No_Auth.POST("login", v1.Login)
	}

	r.Run(utils.HttpPort)
}
