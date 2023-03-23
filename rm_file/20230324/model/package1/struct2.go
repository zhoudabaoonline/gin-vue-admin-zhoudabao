// 自动生成模板Struct2
package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	)

// Struct2 结构体
type Struct2 struct {
      global.GVA_MODEL
      Field2  *string `json:"field2" form:"field2" gorm:"column:field2;comment:field2;type:text;"`
}
// TableName Struct2 表名
func (Struct2) TableName() string {
      return "struct2"
}

