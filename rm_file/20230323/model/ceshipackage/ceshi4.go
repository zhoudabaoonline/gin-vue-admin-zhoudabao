// 自动生成模板Ceshi4
package ceshipackage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Ceshi4 结构体
type Ceshi4 struct {
      global.GVA_MODEL
      Field45  string `json:"field45" form:"field45" gorm:"column:field45;comment:field45;"`
}


// TableName Ceshi4 表名
func (Ceshi4) TableName() string {
  return "ceshi4"
}

