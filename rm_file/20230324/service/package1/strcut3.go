package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/package1"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    package1Req "github.com/flipped-aurora/gin-vue-admin/server/model/package1/request"
)

type Strcut3Service struct {
}

// CreateStrcut3 创建Strcut3记录
// Author [piexlmax](https://github.com/piexlmax)
func (strcut3Service *Strcut3Service) CreateStrcut3(strcut3 package1.Strcut3) (err error) {
	err = global.GVA_DB.Create(&strcut3).Error
	return err
}

// DeleteStrcut3 删除Strcut3记录
// Author [piexlmax](https://github.com/piexlmax)
func (strcut3Service *Strcut3Service)DeleteStrcut3(strcut3 package1.Strcut3) (err error) {
	err = global.GVA_DB.Delete(&strcut3).Error
	return err
}

// DeleteStrcut3ByIds 批量删除Strcut3记录
// Author [piexlmax](https://github.com/piexlmax)
func (strcut3Service *Strcut3Service)DeleteStrcut3ByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]package1.Strcut3{},"id in ?",ids.Ids).Error
	return err
}

// UpdateStrcut3 更新Strcut3记录
// Author [piexlmax](https://github.com/piexlmax)
func (strcut3Service *Strcut3Service)UpdateStrcut3(strcut3 package1.Strcut3) (err error) {
	err = global.GVA_DB.Save(&strcut3).Error
	return err
}

// GetStrcut3 根据id获取Strcut3记录
// Author [piexlmax](https://github.com/piexlmax)
func (strcut3Service *Strcut3Service)GetStrcut3(id uint) (strcut3 package1.Strcut3, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&strcut3).Error
	return
}

// GetStrcut3InfoList 分页获取Strcut3记录
// Author [piexlmax](https://github.com/piexlmax)
func (strcut3Service *Strcut3Service)GetStrcut3InfoList(info package1Req.Strcut3Search) (list []package1.Strcut3, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&package1.Strcut3{})
    var strcut3s []package1.Strcut3
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&strcut3s).Error
	return  strcut3s, total, err
}
