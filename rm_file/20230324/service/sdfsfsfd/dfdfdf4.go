package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type Dfdfdf4Service struct {
}

// CreateDfdfdf4 创建Dfdfdf4记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdf4Service *Dfdfdf4Service) CreateDfdfdf4(dfdfdf4 sdfsfsfd.Dfdfdf4) (err error) {
	err = global.GVA_DB.Create(&dfdfdf4).Error
	return err
}

// DeleteDfdfdf4 删除Dfdfdf4记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdf4Service *Dfdfdf4Service)DeleteDfdfdf4(dfdfdf4 sdfsfsfd.Dfdfdf4) (err error) {
	err = global.GVA_DB.Delete(&dfdfdf4).Error
	return err
}

// DeleteDfdfdf4ByIds 批量删除Dfdfdf4记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdf4Service *Dfdfdf4Service)DeleteDfdfdf4ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdf4{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdf4 更新Dfdfdf4记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdf4Service *Dfdfdf4Service)UpdateDfdfdf4(dfdfdf4 sdfsfsfd.Dfdfdf4) (err error) {
	err = global.GVA_DB.Save(&dfdfdf4).Error
	return err
}

// GetDfdfdf4 根据id获取Dfdfdf4记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdf4Service *Dfdfdf4Service)GetDfdfdf4(id uint) (dfdfdf4 sdfsfsfd.Dfdfdf4, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdf4).Error
	return
}

// GetDfdfdf4InfoList 分页获取Dfdfdf4记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdf4Service *Dfdfdf4Service)GetDfdfdf4InfoList(info sdfsfsfdReq.Dfdfdf4Search) (list []sdfsfsfd.Dfdfdf4, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdf4{})
    var dfdfdf4s []sdfsfsfd.Dfdfdf4
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dfdfdf4s).Error
	return  dfdfdf4s, total, err
}
