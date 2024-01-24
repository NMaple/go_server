package Position

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/recruiter"
	"gorm.io/gorm"
	"time"
)

type Position struct {
}

func (p *Position) CreatePositionService(po recruiter.Position) (positionInfo recruiter.Position, err error) {
	err = global.GVA_DB.Create(&po).Error
	return po, err
}

func (p *Position) GetPositionListService(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&recruiter.Position{})
	var positionList []recruiter.Position
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&positionList).Error
	return positionList, total, err
}

func (p *Position) Deleteposition(id int) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&recruiter.Position{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func (p *Position) SetPositionInfo(req recruiter.Position) error {
	fmt.Println(req)
	return global.GVA_DB.Model(&recruiter.Position{}).
		Select("updated_at", "name", "company_name", "requirements", "description").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at":   time.Now(),
			"name":         req.Name,
			"company_name": req.CompanyName,
			"requirements": req.Requirements,
			"description":  req.Description,
		}).Error
}
