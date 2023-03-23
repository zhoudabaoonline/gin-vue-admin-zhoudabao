package ceshipackage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ceshipackage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    ceshipackageReq "github.com/flipped-aurora/gin-vue-admin/server/model/ceshipackage/request"
)

type CeshiStructService struct {
}

// CreateCeshiStruct 创建CeshiStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshiStructService *CeshiStructService) CreateCeshiStruct(ceshiStruct ceshipackage.CeshiStruct) (err error) {
	err = global.GVA_DB.Create(&ceshiStruct).Error
	return err
}

// DeleteCeshiStruct 删除CeshiStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshiStructService *CeshiStructService)DeleteCeshiStruct(ceshiStruct ceshipackage.CeshiStruct) (err error) {
	err = global.GVA_DB.Delete(&ceshiStruct).Error
	return err
}

// DeleteCeshiStructByIds 批量删除CeshiStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshiStructService *CeshiStructService)DeleteCeshiStructByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ceshipackage.CeshiStruct{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCeshiStruct 更新CeshiStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshiStructService *CeshiStructService)UpdateCeshiStruct(ceshiStruct ceshipackage.CeshiStruct) (err error) {
	err = global.GVA_DB.Save(&ceshiStruct).Error
	return err
}

// GetCeshiStruct 根据id获取CeshiStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshiStructService *CeshiStructService)GetCeshiStruct(id uint) (ceshiStruct ceshipackage.CeshiStruct, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ceshiStruct).Error
	return
}

// GetCeshiStructInfoList 分页获取CeshiStruct记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshiStructService *CeshiStructService)GetCeshiStructInfoList(info ceshipackageReq.CeshiStructSearch) (list []ceshipackage.CeshiStruct, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&ceshipackage.CeshiStruct{})
    var ceshiStructs []ceshipackage.CeshiStruct
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&ceshiStructs).Error
	return  ceshiStructs, total, err
}
