// 自动生成模板Ceshi2
package sdfsfsfd

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Ceshi2 结构体
type Ceshi2 struct {
      global.GVA_MODEL
      Field1  string `json:"field1" form:"field1" gorm:"column:field1;comment:field1;"`
}


// TableName Ceshi2 表名
func (Ceshi2) TableName() string {
  return "ceshi2"
}

