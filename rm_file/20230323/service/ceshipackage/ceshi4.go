package ceshipackage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ceshipackage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ceshipackageReq "github.com/flipped-aurora/gin-vue-admin/server/model/ceshipackage/request"
)

type Ceshi4Service struct {
}

// CreateCeshi4 创建Ceshi4记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshi4Service *Ceshi4Service) CreateCeshi4(ceshi4 ceshipackage.Ceshi4) (err error) {
	err = global.GVA_DB.Create(&ceshi4).Error
	return err
}

// DeleteCeshi4 删除Ceshi4记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshi4Service *Ceshi4Service)DeleteCeshi4(ceshi4 ceshipackage.Ceshi4) (err error) {
	err = global.GVA_DB.Delete(&ceshi4).Error
	return err
}

// DeleteCeshi4ByIds 批量删除Ceshi4记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshi4Service *Ceshi4Service)DeleteCeshi4ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ceshipackage.Ceshi4{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCeshi4 更新Ceshi4记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshi4Service *Ceshi4Service)UpdateCeshi4(ceshi4 ceshipackage.Ceshi4) (err error) {
	err = global.GVA_DB.Save(&ceshi4).Error
	return err
}

// GetCeshi4 根据id获取Ceshi4记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshi4Service *Ceshi4Service)GetCeshi4(id uint) (ceshi4 ceshipackage.Ceshi4, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ceshi4).Error
	return
}

// GetCeshi4InfoList 分页获取Ceshi4记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshi4Service *Ceshi4Service)GetCeshi4InfoList(info ceshipackageReq.Ceshi4Search) (list []ceshipackage.Ceshi4, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ceshipackage.Ceshi4{})
    var ceshi4s []ceshipackage.Ceshi4
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&ceshi4s).Error
	return  ceshi4s, total, err
}
