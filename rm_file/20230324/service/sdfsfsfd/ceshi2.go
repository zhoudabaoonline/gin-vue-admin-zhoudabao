package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type Ceshi2Service struct {
}

// CreateCeshi2 创建Ceshi2记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshi2Service *Ceshi2Service) CreateCeshi2(ceshi2 sdfsfsfd.Ceshi2) (err error) {
	err = global.GVA_DB.Create(&ceshi2).Error
	return err
}

// DeleteCeshi2 删除Ceshi2记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshi2Service *Ceshi2Service)DeleteCeshi2(ceshi2 sdfsfsfd.Ceshi2) (err error) {
	err = global.GVA_DB.Delete(&ceshi2).Error
	return err
}

// DeleteCeshi2ByIds 批量删除Ceshi2记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshi2Service *Ceshi2Service)DeleteCeshi2ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Ceshi2{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCeshi2 更新Ceshi2记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshi2Service *Ceshi2Service)UpdateCeshi2(ceshi2 sdfsfsfd.Ceshi2) (err error) {
	err = global.GVA_DB.Save(&ceshi2).Error
	return err
}

// GetCeshi2 根据id获取Ceshi2记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshi2Service *Ceshi2Service)GetCeshi2(id uint) (ceshi2 sdfsfsfd.Ceshi2, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ceshi2).Error
	return
}

// GetCeshi2InfoList 分页获取Ceshi2记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshi2Service *Ceshi2Service)GetCeshi2InfoList(info sdfsfsfdReq.Ceshi2Search) (list []sdfsfsfd.Ceshi2, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Ceshi2{})
    var ceshi2s []sdfsfsfd.Ceshi2
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&ceshi2s).Error
	return  ceshi2s, total, err
}
