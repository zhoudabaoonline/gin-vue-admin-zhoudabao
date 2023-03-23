package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type DfdfdfccService struct {
}

// CreateDfdfdfcc 创建Dfdfdfcc记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfccService *DfdfdfccService) CreateDfdfdfcc(dfdfdfcc sdfsfsfd.Dfdfdfcc) (err error) {
	err = global.GVA_DB.Create(&dfdfdfcc).Error
	return err
}

// DeleteDfdfdfcc 删除Dfdfdfcc记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfccService *DfdfdfccService)DeleteDfdfdfcc(dfdfdfcc sdfsfsfd.Dfdfdfcc) (err error) {
	err = global.GVA_DB.Delete(&dfdfdfcc).Error
	return err
}

// DeleteDfdfdfccByIds 批量删除Dfdfdfcc记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfccService *DfdfdfccService)DeleteDfdfdfccByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdfcc{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdfcc 更新Dfdfdfcc记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfccService *DfdfdfccService)UpdateDfdfdfcc(dfdfdfcc sdfsfsfd.Dfdfdfcc) (err error) {
	err = global.GVA_DB.Save(&dfdfdfcc).Error
	return err
}

// GetDfdfdfcc 根据id获取Dfdfdfcc记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfccService *DfdfdfccService)GetDfdfdfcc(id uint) (dfdfdfcc sdfsfsfd.Dfdfdfcc, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdfcc).Error
	return
}

// GetDfdfdfccInfoList 分页获取Dfdfdfcc记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfccService *DfdfdfccService)GetDfdfdfccInfoList(info sdfsfsfdReq.DfdfdfccSearch) (list []sdfsfsfd.Dfdfdfcc, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdfcc{})
    var dfdfdfccs []sdfsfsfd.Dfdfdfcc
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dfdfdfccs).Error
	return  dfdfdfccs, total, err
}
