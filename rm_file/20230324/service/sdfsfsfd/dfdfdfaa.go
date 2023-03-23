package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type DfdfdfaaaService struct {
}

// CreateDfdfdfaaa 创建Dfdfdfaaa记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfaaService *DfdfdfaaaService) CreateDfdfdfaaa(dfdfdfaa sdfsfsfd.Dfdfdfaaa) (err error) {
	err = global.GVA_DB.Create(&dfdfdfaa).Error
	return err
}

// DeleteDfdfdfaaa 删除Dfdfdfaaa记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfaaService *DfdfdfaaaService)DeleteDfdfdfaaa(dfdfdfaa sdfsfsfd.Dfdfdfaaa) (err error) {
	err = global.GVA_DB.Delete(&dfdfdfaa).Error
	return err
}

// DeleteDfdfdfaaaByIds 批量删除Dfdfdfaaa记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfaaService *DfdfdfaaaService)DeleteDfdfdfaaaByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdfaaa{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdfaaa 更新Dfdfdfaaa记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfaaService *DfdfdfaaaService)UpdateDfdfdfaaa(dfdfdfaa sdfsfsfd.Dfdfdfaaa) (err error) {
	err = global.GVA_DB.Save(&dfdfdfaa).Error
	return err
}

// GetDfdfdfaaa 根据id获取Dfdfdfaaa记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfaaService *DfdfdfaaaService)GetDfdfdfaaa(id uint) (dfdfdfaa sdfsfsfd.Dfdfdfaaa, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdfaa).Error
	return
}

// GetDfdfdfaaaInfoList 分页获取Dfdfdfaaa记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfaaService *DfdfdfaaaService)GetDfdfdfaaaInfoList(info sdfsfsfdReq.DfdfdfaaaSearch) (list []sdfsfsfd.Dfdfdfaaa, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdfaaa{})
    var dfdfdfaas []sdfsfsfd.Dfdfdfaaa
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dfdfdfaas).Error
	return  dfdfdfaas, total, err
}
