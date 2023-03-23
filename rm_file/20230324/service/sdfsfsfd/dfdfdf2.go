package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type Dfdfdf22Service struct {
}

// CreateDfdfdf22 创建Dfdfdf22记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdf22Service *Dfdfdf22Service) CreateDfdfdf22(dfdfdf22 sdfsfsfd.Dfdfdf22) (err error) {
	err = global.GVA_DB.Create(&dfdfdf22).Error
	return err
}

// DeleteDfdfdf22 删除Dfdfdf22记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdf22Service *Dfdfdf22Service)DeleteDfdfdf22(dfdfdf22 sdfsfsfd.Dfdfdf22) (err error) {
	err = global.GVA_DB.Delete(&dfdfdf22).Error
	return err
}

// DeleteDfdfdf22ByIds 批量删除Dfdfdf22记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdf22Service *Dfdfdf22Service)DeleteDfdfdf22ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdf22{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdf22 更新Dfdfdf22记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdf22Service *Dfdfdf22Service)UpdateDfdfdf22(dfdfdf22 sdfsfsfd.Dfdfdf22) (err error) {
	err = global.GVA_DB.Save(&dfdfdf22).Error
	return err
}

// GetDfdfdf22 根据id获取Dfdfdf22记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdf22Service *Dfdfdf22Service)GetDfdfdf22(id uint) (dfdfdf22 sdfsfsfd.Dfdfdf22, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdf22).Error
	return
}

// GetDfdfdf22InfoList 分页获取Dfdfdf22记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdf22Service *Dfdfdf22Service)GetDfdfdf22InfoList(info sdfsfsfdReq.Dfdfdf22Search) (list []sdfsfsfd.Dfdfdf22, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdf22{})
    var dfdfdf22s []sdfsfsfd.Dfdfdf22
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dfdfdf22s).Error
	return  dfdfdf22s, total, err
}
