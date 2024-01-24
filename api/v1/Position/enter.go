package Position

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	Position
}

var (
	positionApiService = service.ServiceGroupApp.PositionServiceGroup.Position
)
