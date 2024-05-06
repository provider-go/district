package district

import (
	"github.com/gin-gonic/gin"
	"github.com/provider-go/district/global"
	"github.com/provider-go/district/router"
	"github.com/provider-go/pkg/types"
)

type Plugin struct{}

func CreatePlugin() *Plugin {
	return &Plugin{}
}

func CreatePluginAndDB(instance types.PluginNeedInstance) *Plugin {
	global.DB = instance.Mysql
	return &Plugin{}
}

func (*Plugin) Register(group *gin.RouterGroup) {
	router.GroupApp.InitRouter(group)
}

func (*Plugin) RouterPath() string {
	return "district"
}
