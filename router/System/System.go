package system

import (
	"github.com/gin-gonic/gin"
	system "github.com/titrxw/smart-home-server/app/Controller/System"
)

type System struct {
}

func (this *System) registerRoute(router *gin.RouterGroup) {
	router.GET("/", new(system.IndexController).Index)
}
