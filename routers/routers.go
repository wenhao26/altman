package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"altman/global"
	"altman/api/controller"
)

func InitRouters() {
	r := gin.Default()

	r.GET("/", controller.Index)

	addr := fmt.Sprintf(":%d", global.AmConfig.Base.ListenPort)
	r.Run(addr)
}
