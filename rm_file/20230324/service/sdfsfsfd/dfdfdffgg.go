package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type DfdfdffgService struct {
}

// CreateDfdfdffg 创建Dfdfdffg记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdffgService *DfdfdffgService) CreateDfdfdffg(dfdfdffg sdfsfsfd.Dfdfdffg) (err error) {
	err = global.GVA_DB.Create(&dfdfdffg).Error
	return err
}

// DeleteDfdfdffg 删除Dfdfdffg记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdffgService *DfdfdffgService)DeleteDfdfdffg(dfdfdffg sdfsfsfd.Dfdfdffg) (err error) {
	err = global.GVA_DB.Delete(&dfdfdffg).Error
	return err
}

// DeleteDfdfdffgByIds 批量删除Dfdfdffg记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdffgService *DfdfdffgService)DeleteDfdfdffgByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdffg{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdffg 更新Dfdfdffg记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdffgService *DfdfdffgService)UpdateDfdfdffg(dfdfdffg sdfsfsfd.Dfdfdffg) (err error) {
	err = global.GVA_DB.Save(&dfdfdffg).Error
	return err
}

// GetDfdfdffg 根据id获取Dfdfdffg记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdffgService *DfdfdffgService)GetDfdfdffg(id uint) (dfdfdffg sdfsfsfd.Dfdfdffg, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdffg).Error
	return
}

// GetDfdfdffgInfoList 分页获取Dfdfdffg记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdffgService *DfdfdffgService)GetDfdfdffgInfoList(info sdfsfsfdReq.DfdfdffgSearch) (list []sdfsfsfd.Dfdfdffg, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdffg{})
    var dfdfdffgs []sdfsfsfd.Dfdfdffg
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dfdfdffgs).Error
	return  dfdfdffgs, total, err
}
