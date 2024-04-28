package router

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/district/api"
)

type Group struct {
	Router
}

var GroupApp = new(Group)

type Router struct{}

func (s *Router) InitRouter(Router *gin.RouterGroup) {
	{
		// districts 表操作
		Router.POST("district/list", api.ListDistrict)
	}
}
