package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type DfdfdfvvvvService struct {
}

// CreateDfdfdfvvvv 创建Dfdfdfvvvv记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfvvService *DfdfdfvvvvService) CreateDfdfdfvvvv(dfdfdfvv sdfsfsfd.Dfdfdfvvvv) (err error) {
	err = global.GVA_DB.Create(&dfdfdfvv).Error
	return err
}

// DeleteDfdfdfvvvv 删除Dfdfdfvvvv记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfvvService *DfdfdfvvvvService)DeleteDfdfdfvvvv(dfdfdfvv sdfsfsfd.Dfdfdfvvvv) (err error) {
	err = global.GVA_DB.Delete(&dfdfdfvv).Error
	return err
}

// DeleteDfdfdfvvvvByIds 批量删除Dfdfdfvvvv记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfvvService *DfdfdfvvvvService)DeleteDfdfdfvvvvByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdfvvvv{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdfvvvv 更新Dfdfdfvvvv记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfvvService *DfdfdfvvvvService)UpdateDfdfdfvvvv(dfdfdfvv sdfsfsfd.Dfdfdfvvvv) (err error) {
	err = global.GVA_DB.Save(&dfdfdfvv).Error
	return err
}

// GetDfdfdfvvvv 根据id获取Dfdfdfvvvv记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfvvService *DfdfdfvvvvService)GetDfdfdfvvvv(id uint) (dfdfdfvv sdfsfsfd.Dfdfdfvvvv, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdfvv).Error
	return
}

// GetDfdfdfvvvvInfoList 分页获取Dfdfdfvvvv记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfvvService *DfdfdfvvvvService)GetDfdfdfvvvvInfoList(info sdfsfsfdReq.DfdfdfvvvvSearch) (list []sdfsfsfd.Dfdfdfvvvv, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdfvvvv{})
    var dfdfdfvvs []sdfsfsfd.Dfdfdfvvvv
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dfdfdfvvs).Error
	return  dfdfdfvvs, total, err
}
