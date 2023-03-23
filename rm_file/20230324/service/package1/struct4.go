package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/package1"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    package1Req "github.com/flipped-aurora/gin-vue-admin/server/model/package1/request"
)

type Struct4Service struct {
}

// CreateStruct4 创建Struct4记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct4Service *Struct4Service) CreateStruct4(struct4 package1.Struct4) (err error) {
	err = global.GVA_DB.Create(&struct4).Error
	return err
}

// DeleteStruct4 删除Struct4记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct4Service *Struct4Service)DeleteStruct4(struct4 package1.Struct4) (err error) {
	err = global.GVA_DB.Delete(&struct4).Error
	return err
}

// DeleteStruct4ByIds 批量删除Struct4记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct4Service *Struct4Service)DeleteStruct4ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]package1.Struct4{},"id in ?",ids.Ids).Error
	return err
}

// UpdateStruct4 更新Struct4记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct4Service *Struct4Service)UpdateStruct4(struct4 package1.Struct4) (err error) {
	err = global.GVA_DB.Save(&struct4).Error
	return err
}

// GetStruct4 根据id获取Struct4记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct4Service *Struct4Service)GetStruct4(id uint) (struct4 package1.Struct4, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&struct4).Error
	return
}

// GetStruct4InfoList 分页获取Struct4记录
// Author [piexlmax](https://github.com/piexlmax)
func (struct4Service *Struct4Service)GetStruct4InfoList(info package1Req.Struct4Search) (list []package1.Struct4, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&package1.Struct4{})
    var struct4s []package1.Struct4
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&struct4s).Error
	return  struct4s, total, err
}
