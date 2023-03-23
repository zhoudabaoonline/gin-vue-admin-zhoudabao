// 自动生成模板Ceshi
package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Ceshi 结构体
type Ceshi struct {
      global.GVA_MODEL
      顶顶顶顶  string `json:"顶顶顶顶" form:"顶顶顶顶" gorm:"column:顶顶顶顶;comment:;"`
}


// TableName Ceshi 表名
func (Ceshi) TableName() string {
  return "ceshi"
}

