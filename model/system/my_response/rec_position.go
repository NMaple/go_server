package my_response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/recruiter"
)

type PositionResponse struct {
	Position recruiter.Position `json:"position"`
}
