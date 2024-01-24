package recruiter

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Position struct {
	global.GVA_MODEL
	Name         string `json:"name"`
	Description  string `json:"description"`
	Requirements string `json:"requirements"`
	CompanyName  string `json:"company_name"`
	UserID       uint   `json:"user_id" gorm:"foreignkey:UserID"`
	CreateName   string `json:"create_name"`
}

func (Position) TableName() string {
	return "position"
}
