package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    sdfsfsfdReq "github.com/flipped-aurora/gin-vue-admin/server/model/sdfsfsfd/request"
)

type CeshiService struct {
}

// CreateCeshi 创建Ceshi记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshiService *CeshiService) CreateCeshi(ceshi sdfsfsfd.Ceshi) (err error) {
	err = global.GVA_DB.Create(&ceshi).Error
	return err
}

// DeleteCeshi 删除Ceshi记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshiService *CeshiService)DeleteCeshi(ceshi sdfsfsfd.Ceshi) (err error) {
	err = global.GVA_DB.Delete(&ceshi).Error
	return err
}

// DeleteCeshiByIds 批量删除Ceshi记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshiService *CeshiService)DeleteCeshiByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]sdfsfsfd.Ceshi{},"id in ?",ids.Ids).Error
	return err
}

// UpdateCeshi 更新Ceshi记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshiService *CeshiService)UpdateCeshi(ceshi sdfsfsfd.Ceshi) (err error) {
	err = global.GVA_DB.Save(&ceshi).Error
	return err
}

// GetCeshi 根据id获取Ceshi记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshiService *CeshiService)GetCeshi(id uint) (ceshi sdfsfsfd.Ceshi, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ceshi).Error
	return
}

// GetCeshiInfoList 分页获取Ceshi记录
// Author [piexlmax](https://github.com/piexlmax)
func (ceshiService *CeshiService)GetCeshiInfoList(info sdfsfsfdReq.CeshiSearch) (list []sdfsfsfd.Ceshi, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&sdfsfsfd.Ceshi{})
    var ceshis []sdfsfsfd.Ceshi
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&ceshis).Error
	return  ceshis, total, err
}
