package router

import "github.com/gin-gonic/gin"

/**
 * InitNotify 初始化电影相关路由
 */
func InitNotify(Router *gin.RouterGroup) {
	AuthGroup := Router.Group("wx")

	{
		AuthGroup.POST("Notify", Notify) // 回调接口
	}

}
