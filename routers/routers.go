package routers

import (
	"altman/middleware"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"altman/global"
	"altman/api/controller"
)

func InitRouters() {
	r := gin.Default()
	gin.ForceConsoleColor()

	r.Use(middleware.RateLimiter(time.Second, 100, 100)) // 限流阀
	r.GET("/", controller.Index)

	addr := fmt.Sprintf(":%d", global.AmConfig.Base.ListenPort)
	r.Run(addr)
}
