package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/package1"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    package1Req "github.com/flipped-aurora/gin-vue-admin/server/model/package1/request"
)

type Struct1Service struct {
}

// CreateStruct1 创建Struct1记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct1Service *Struct1Service) CreateStruct1(struct1 package1.Struct1) (err error) {
	err = global.GVA_DB.Create(&struct1).Error
	return err
}

// DeleteStruct1 删除Struct1记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct1Service *Struct1Service)DeleteStruct1(struct1 package1.Struct1) (err error) {
	err = global.GVA_DB.Delete(&struct1).Error
	return err
}

// DeleteStruct1ByIds 批量删除Struct1记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct1Service *Struct1Service)DeleteStruct1ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]package1.Struct1{},"id in ?",ids.Ids).Error
	return err
}

// UpdateStruct1 更新Struct1记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct1Service *Struct1Service)UpdateStruct1(struct1 package1.Struct1) (err error) {
	err = global.GVA_DB.Save(&struct1).Error
	return err
}

// GetStruct1 根据id获取Struct1记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct1Service *Struct1Service)GetStruct1(id uint) (struct1 package1.Struct1, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&struct1).Error
	return
}

// GetStruct1InfoList 分页获取Struct1记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct1Service *Struct1Service)GetStruct1InfoList(info package1Req.Struct1Search) (list []package1.Struct1, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&package1.Struct1{})
    var struct1s []package1.Struct1
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&struct1s).Error
	return  struct1s, total, err
}
