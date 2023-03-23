package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type DfdfdfyyyyService struct {
}

// CreateDfdfdfyyyy 创建Dfdfdfyyyy记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfyyyService *DfdfdfyyyyService) CreateDfdfdfyyyy(dfdfdfyyy sdfsfsfd.Dfdfdfyyyy) (err error) {
	err = global.GVA_DB.Create(&dfdfdfyyy).Error
	return err
}

// DeleteDfdfdfyyyy 删除Dfdfdfyyyy记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfyyyService *DfdfdfyyyyService)DeleteDfdfdfyyyy(dfdfdfyyy sdfsfsfd.Dfdfdfyyyy) (err error) {
	err = global.GVA_DB.Delete(&dfdfdfyyy).Error
	return err
}

// DeleteDfdfdfyyyyByIds 批量删除Dfdfdfyyyy记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfyyyService *DfdfdfyyyyService)DeleteDfdfdfyyyyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdfyyyy{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdfyyyy 更新Dfdfdfyyyy记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfyyyService *DfdfdfyyyyService)UpdateDfdfdfyyyy(dfdfdfyyy sdfsfsfd.Dfdfdfyyyy) (err error) {
	err = global.GVA_DB.Save(&dfdfdfyyy).Error
	return err
}

// GetDfdfdfyyyy 根据id获取Dfdfdfyyyy记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfyyyService *DfdfdfyyyyService)GetDfdfdfyyyy(id uint) (dfdfdfyyy sdfsfsfd.Dfdfdfyyyy, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdfyyy).Error
	return
}

// GetDfdfdfyyyyInfoList 分页获取Dfdfdfyyyy记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfyyyService *DfdfdfyyyyService)GetDfdfdfyyyyInfoList(info sdfsfsfdReq.DfdfdfyyyySearch) (list []sdfsfsfd.Dfdfdfyyyy, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdfyyyy{})
    var dfdfdfyyys []sdfsfsfd.Dfdfdfyyyy
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dfdfdfyyys).Error
	return  dfdfdfyyys, total, err
}
