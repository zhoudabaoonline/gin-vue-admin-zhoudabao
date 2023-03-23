package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type DfdfdfssssService struct {
}

// CreateDfdfdfssss 创建Dfdfdfssss记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfsssService *DfdfdfssssService) CreateDfdfdfssss(dfdfdfsss sdfsfsfd.Dfdfdfssss) (err error) {
	err = global.GVA_DB.Create(&dfdfdfsss).Error
	return err
}

// DeleteDfdfdfssss 删除Dfdfdfssss记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfsssService *DfdfdfssssService)DeleteDfdfdfssss(dfdfdfsss sdfsfsfd.Dfdfdfssss) (err error) {
	err = global.GVA_DB.Delete(&dfdfdfsss).Error
	return err
}

// DeleteDfdfdfssssByIds 批量删除Dfdfdfssss记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfsssService *DfdfdfssssService)DeleteDfdfdfssssByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Dfdfdfssss{},"id in ?",ids.Ids).Error
	return err
}

// UpdateDfdfdfssss 更新Dfdfdfssss记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfsssService *DfdfdfssssService)UpdateDfdfdfssss(dfdfdfsss sdfsfsfd.Dfdfdfssss) (err error) {
	err = global.GVA_DB.Save(&dfdfdfsss).Error
	return err
}

// GetDfdfdfssss 根据id获取Dfdfdfssss记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfsssService *DfdfdfssssService)GetDfdfdfssss(id uint) (dfdfdfsss sdfsfsfd.Dfdfdfssss, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dfdfdfsss).Error
	return
}

// GetDfdfdfssssInfoList 分页获取Dfdfdfssss记录
// Author [piexlmax](https://github.com/piexlmax)
func (dfdfdfsssService *DfdfdfssssService)GetDfdfdfssssInfoList(info sdfsfsfdReq.DfdfdfssssSearch) (list []sdfsfsfd.Dfdfdfssss, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Dfdfdfssss{})
    var dfdfdfssss []sdfsfsfd.Dfdfdfssss
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
