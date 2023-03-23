package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/package1"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    package1Req "github.com/flipped-aurora/gin-vue-admin/server/model/package1/request"
)

type Struct2Service struct {
}

// CreateStruct2 创建Struct2记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct2Service *Struct2Service) CreateStruct2(struct2 package1.Struct2) (err error) {
	err = global.GVA_DB.Create(&struct2).Error
	return err
}

// DeleteStruct2 删除Struct2记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct2Service *Struct2Service)DeleteStruct2(struct2 package1.Struct2) (err error) {
	err = global.GVA_DB.Delete(&struct2).Error
	return err
}

// DeleteStruct2ByIds 批量删除Struct2记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct2Service *Struct2Service)DeleteStruct2ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]package1.Struct2{},"id in ?",ids.Ids).Error
	return err
}

// UpdateStruct2 更新Struct2记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct2Service *Struct2Service)UpdateStruct2(struct2 package1.Struct2) (err error) {
	err = global.GVA_DB.Save(&struct2).Error
	return err
}

// GetStruct2 根据id获取Struct2记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct2Service *Struct2Service)GetStruct2(id uint) (struct2 package1.Struct2, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&struct2).Error
	return
}

// GetStruct2InfoList 分页获取Struct2记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct2Service *Struct2Service)GetStruct2InfoList(info package1Req.Struct2Search) (list []package1.Struct2, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&package1.Struct2{})
    var struct2s []package1.Struct2
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&struct2s).Error
	return  struct2s, total, err
}
