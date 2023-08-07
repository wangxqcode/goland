package router

import "github.com/gin-gonic/gin"

// 设置路由

func SetupRouter(mode string) *gin.Engine  {
	// set release mode
	if mode == gin.ReleaseMode{
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	// 启用日志
	r.Use()

	// 注册

	// 登录





	
}
