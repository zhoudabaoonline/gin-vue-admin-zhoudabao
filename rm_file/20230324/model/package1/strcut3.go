// 自动生成模板Strcut3
package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	"github.com/flipped-aurora/gin-vue-admin/server/model/globalModel")

// Strcut3 结构体
type Strcut3 struct {
      global.GVA_MODEL
      Fuwenben  *string `json:"fuwenben" form:"fuwenben" gorm:"column:fuwenben;comment:fuwenben;type:text;"`
      Multipic  globalModel.MultiPic `json:"multipic" form:"multipic" gorm:"column:multipic;comment:multipic;type:json;"`
}
// TableName Strcut3 表名
func (Strcut3) TableName() string {
      return "strcut3"
}

