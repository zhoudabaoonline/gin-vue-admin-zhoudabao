package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type DfdfdfzzService struct {
}

// CreateDfdfdfzz 创建Dfdfdfzz记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfzzzzService *DfdfdfzzService) CreateDfdfdfzz(dfdfdfzzzz sdfsfsfd.Dfdfdfzz) (err error) {
	err = global.GVA_DB.Create(&dfdfdfzzzz).Error
	return err
}

// DeleteDfdfdfzz 删除Dfdfdfzz记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfzzzzService *DfdfdfzzService)DeleteDfdfdfzz(dfdfdfzzzz sdfsfsfd.Dfdfdfzz) (err error) {
	err = global.GVA_DB.Delete(&dfdfdfzzzz).Error
	return err
}

// DeleteDfdfdfzzByIds 批量删除Dfdfdfzz记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfzzzzService *DfdfdfzzService)DeleteDfdfdfzzByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdfzz{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdfzz 更新Dfdfdfzz记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfzzzzService *DfdfdfzzService)UpdateDfdfdfzz(dfdfdfzzzz sdfsfsfd.Dfdfdfzz) (err error) {
	err = global.GVA_DB.Save(&dfdfdfzzzz).Error
	return err
}

// GetDfdfdfzz 根据id获取Dfdfdfzz记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfzzzzService *DfdfdfzzService)GetDfdfdfzz(id uint) (dfdfdfzzzz sdfsfsfd.Dfdfdfzz, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdfzzzz).Error
	return
}

// GetDfdfdfzzInfoList 分页获取Dfdfdfzz记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfzzzzService *DfdfdfzzService)GetDfdfdfzzInfoList(info sdfsfsfdReq.DfdfdfzzSearch) (list []sdfsfsfd.Dfdfdfzz, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdfzz{})
    var dfdfdfzzzzs []sdfsfsfd.Dfdfdfzz
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dfdfdfzzzzs).Error
	return  dfdfdfzzzzs, total, err
}
