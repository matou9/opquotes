package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"opquotes/model/base"
	"opquotes/model/market"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()
	pprof.Register(r)
	r.Use(Cors())


	api := r.Group("/api/v1", gin.Logger(),Authorize())
	{
		//以下为行情接口部分
		api.POST("tick", market.Tick)
		api.POST("ticks", market.Ticks)
		api.POST("mindata", market.Mindata)
		api.POST("bars", market.Bars)
		api.POST("getcomm",base.GetComm)
		api.POST("notice",base.Notice)
		api.POST("getoptions",base.GetOptions)
		api.POST("getoptionsbymap",base.GetOptionsByMap)
		api.POST("tickbymonth",market.TickByMonth)
		api.POST("tickunderlying",market.TickUnderlying)
		//修改密码

}
	//ws := r.Group("/ws", gin.Logger())
	//{
	//	ws.GET("/tick",websocket.WsPage)
	//}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "TEST",
		})
	})
	return r
}

//// 中间件, 顺序不能改
//r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
//r.Use(middleware.Cors())
//r.Use(middleware.CurrentUser())
//
//// 路由
//v1 := r.Group("/api/v1")
//{
//	v1.POST("ping", api.Ping)
//
//	// 用户登录
//	v1.POST("user/register", api.UserRegister)
//
//	// 用户登录
//	v1.POST("user/login", api.UserLogin)
//
//	// 需要登录保护的
//	auth := v1.Group("")
//	auth.Use(middleware.AuthRequired())
//	{
//		// User Routing
//		auth.GET("user/me", api.UserMe)
//		auth.DELETE("user/logout", api.UserLogout)
//	}
//}
