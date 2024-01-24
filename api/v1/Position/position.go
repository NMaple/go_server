package Position

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/recruiter"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/my_response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Position struct {
}

func (p *Position) CreatePositionApi(c *gin.Context) {
	var r systemReq.CreatePosition
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.PositionVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	position := &recruiter.Position{Name: r.Name, Description: r.Description, Requirements: r.Requirements, CompanyName: r.CompanyName, UserID: r.UserID, CreateName: r.CreateName}
	positionReturn, err := positionApiService.CreatePositionService(*position)
	if err != nil {
		global.GVA_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithDetailed(my_response.PositionResponse{Position: positionReturn}, "添加失败", c)
		return
	}
	response.OkWithDetailed(my_response.PositionResponse{Position: positionReturn}, "添加成功", c)
}

func (p *Position) GetPositionList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := positionApiService.GetPositionListService(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (p *Position) DeletePosition(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = positionApiService.Deleteposition(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (p *Position) SetPositionInfo(c *gin.Context) {
	var position systemReq.ChangePositionInfo
	err := c.ShouldBindJSON(&position)
	fmt.Println(position)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(position, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = positionApiService.SetPositionInfo(recruiter.Position{
		GVA_MODEL:    global.GVA_MODEL{ID: position.ID},
		Name:         position.Name,
		CompanyName:  position.CompanyName,
		Description:  position.Description,
		Requirements: position.Requirements,
	})
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}
	response.OkWithMessage("设置成功", c)

}
