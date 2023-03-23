// 自动生成模板Struct4
package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	)

// Struct4 结构体
type Struct4 struct {
      global.GVA_MODEL
      Field4  *string `json:"field4" form:"field4" gorm:"column:field4;comment:field4;type:text;"`
}
// TableName Struct4 表名
func (Struct4) TableName() string {
      return "struct4"
}

