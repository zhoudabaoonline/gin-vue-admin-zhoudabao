package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type DfdfdffdfService struct {
}

// CreateDfdfdffdf 创建Dfdfdffdf记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfsssService *DfdfdffdfService) CreateDfdfdffdf(dfdfdfsss sdfsfsfd.Dfdfdffdf) (err error) {
	err = global.GVA_DB.Create(&dfdfdfsss).Error
	return err
}

// DeleteDfdfdffdf 删除Dfdfdffdf记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfsssService *DfdfdffdfService)DeleteDfdfdffdf(dfdfdfsss sdfsfsfd.Dfdfdffdf) (err error) {
	err = global.GVA_DB.Delete(&dfdfdfsss).Error
	return err
}

// DeleteDfdfdffdfByIds 批量删除Dfdfdffdf记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfsssService *DfdfdffdfService)DeleteDfdfdffdfByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdffdf{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdffdf 更新Dfdfdffdf记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfsssService *DfdfdffdfService)UpdateDfdfdffdf(dfdfdfsss sdfsfsfd.Dfdfdffdf) (err error) {
	err = global.GVA_DB.Save(&dfdfdfsss).Error
	return err
}

// GetDfdfdffdf 根据id获取Dfdfdffdf记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfsssService *DfdfdffdfService)GetDfdfdffdf(id uint) (dfdfdfsss sdfsfsfd.Dfdfdffdf, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdfsss).Error
	return
}

// GetDfdfdffdfInfoList 分页获取Dfdfdffdf记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfsssService *DfdfdffdfService)GetDfdfdffdfInfoList(info sdfsfsfdReq.DfdfdffdfSearch) (list []sdfsfsfd.Dfdfdffdf, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdffdf{})
    var dfdfdfssss []sdfsfsfd.Dfdfdffdf
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&dfdfdfssss).Error
	return  dfdfdfssss, total, err
}
