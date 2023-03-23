package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type DfdfdfqqqService struct {
}

// CreateDfdfdfqqq 创建Dfdfdfqqq记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfqqqService *DfdfdfqqqService) CreateDfdfdfqqq(dfdfdfqqq sdfsfsfd.Dfdfdfqqq) (err error) {
	err = global.GVA_DB.Create(&dfdfdfqqq).Error
	return err
}

// DeleteDfdfdfqqq 删除Dfdfdfqqq记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfqqqService *DfdfdfqqqService)DeleteDfdfdfqqq(dfdfdfqqq sdfsfsfd.Dfdfdfqqq) (err error) {
	err = global.GVA_DB.Delete(&dfdfdfqqq).Error
	return err
}

// DeleteDfdfdfqqqByIds 批量删除Dfdfdfqqq记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfqqqService *DfdfdfqqqService)DeleteDfdfdfqqqByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdfqqq{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdfqqq 更新Dfdfdfqqq记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfqqqService *DfdfdfqqqService)UpdateDfdfdfqqq(dfdfdfqqq sdfsfsfd.Dfdfdfqqq) (err error) {
	err = global.GVA_DB.Save(&dfdfdfqqq).Error
	return err
}

// GetDfdfdfqqq 根据id获取Dfdfdfqqq记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfqqqService *DfdfdfqqqService)GetDfdfdfqqq(id uint) (dfdfdfqqq sdfsfsfd.Dfdfdfqqq, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdfqqq).Error
	return
}

// GetDfdfdfqqqInfoList 分页获取Dfdfdfqqq记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfqqqService *DfdfdfqqqService)GetDfdfdfqqqInfoList(info sdfsfsfdReq.DfdfdfqqqSearch) (list []sdfsfsfd.Dfdfdfqqq, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdfqqq{})
    var dfdfdfqqqs []sdfsfsfd.Dfdfdfqqq
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dfdfdfqqqs).Error
	return  dfdfdfqqqs, total, err
}
