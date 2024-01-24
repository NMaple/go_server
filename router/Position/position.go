package Position

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type PosiTion struct {
}

func (p *PosiTion) InitPositionRouter(Router *gin.RouterGroup) {
	PositionRouter := Router.Group("position")
	api := v1.ApiGroupApp.PositionApiGroup
	{
		PositionRouter.POST("create", api.CreatePositionApi)
		PositionRouter.POST("GetPositionList", api.GetPositionList)
		PositionRouter.DELETE("deleteposition", api.DeletePosition)
		PositionRouter.PUT("setPositionInfo", api.SetPositionInfo)
	}
}
