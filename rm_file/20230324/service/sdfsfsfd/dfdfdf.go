package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type DfdfdfService struct {
}

// CreateDfdfdf 创建Dfdfdf记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfService *DfdfdfService) CreateDfdfdf(dfdfdf sdfsfsfd.Dfdfdf) (err error) {
	err = global.GVA_DB.Create(&dfdfdf).Error
	return err
}

// DeleteDfdfdf 删除Dfdfdf记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfService *DfdfdfService)DeleteDfdfdf(dfdfdf sdfsfsfd.Dfdfdf) (err error) {
	err = global.GVA_DB.Delete(&dfdfdf).Error
	return err
}

// DeleteDfdfdfByIds 批量删除Dfdfdf记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfService *DfdfdfService)DeleteDfdfdfByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdf{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdf 更新Dfdfdf记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfService *DfdfdfService)UpdateDfdfdf(dfdfdf sdfsfsfd.Dfdfdf) (err error) {
	err = global.GVA_DB.Save(&dfdfdf).Error
	return err
}

// GetDfdfdf 根据id获取Dfdfdf记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfService *DfdfdfService)GetDfdfdf(id uint) (dfdfdf sdfsfsfd.Dfdfdf, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdf).Error
	return
}

// GetDfdfdfInfoList 分页获取Dfdfdf记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfService *DfdfdfService)GetDfdfdfInfoList(info sdfsfsfdReq.DfdfdfSearch) (list []sdfsfsfd.Dfdfdf, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdf{})
    var dfdfdfs []sdfsfsfd.Dfdfdf
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dfdfdfs).Error
	return  dfdfdfs, total, err
}
