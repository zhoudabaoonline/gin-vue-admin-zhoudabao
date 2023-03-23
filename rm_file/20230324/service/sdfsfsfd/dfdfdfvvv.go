package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type DfdfdfvvvService struct {
}

// CreateDfdfdfvvv 创建Dfdfdfvvv记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfvvvvService *DfdfdfvvvService) CreateDfdfdfvvv(dfdfdfvvvv sdfsfsfd.Dfdfdfvvv) (err error) {
	err = global.GVA_DB.Create(&dfdfdfvvvv).Error
	return err
}

// DeleteDfdfdfvvv 删除Dfdfdfvvv记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfvvvvService *DfdfdfvvvService)DeleteDfdfdfvvv(dfdfdfvvvv sdfsfsfd.Dfdfdfvvv) (err error) {
	err = global.GVA_DB.Delete(&dfdfdfvvvv).Error
	return err
}

// DeleteDfdfdfvvvByIds 批量删除Dfdfdfvvv记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfvvvvService *DfdfdfvvvService)DeleteDfdfdfvvvByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdfvvv{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdfvvv 更新Dfdfdfvvv记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfvvvvService *DfdfdfvvvService)UpdateDfdfdfvvv(dfdfdfvvvv sdfsfsfd.Dfdfdfvvv) (err error) {
	err = global.GVA_DB.Save(&dfdfdfvvvv).Error
	return err
}

// GetDfdfdfvvv 根据id获取Dfdfdfvvv记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfvvvvService *DfdfdfvvvService)GetDfdfdfvvv(id uint) (dfdfdfvvvv sdfsfsfd.Dfdfdfvvv, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdfvvvv).Error
	return
}

// GetDfdfdfvvvInfoList 分页获取Dfdfdfvvv记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfvvvvService *DfdfdfvvvService)GetDfdfdfvvvInfoList(info sdfsfsfdReq.DfdfdfvvvSearch) (list []sdfsfsfd.Dfdfdfvvv, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdfvvv{})
    var dfdfdfvvvvs []sdfsfsfd.Dfdfdfvvv
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dfdfdfvvvvs).Error
	return  dfdfdfvvvvs, total, err
}
