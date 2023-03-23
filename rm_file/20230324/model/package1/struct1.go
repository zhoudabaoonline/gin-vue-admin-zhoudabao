// 自动生成模板Struct1
package package1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	"github.com/flipped-aurora/gin-vue-admin/server/model/globalModel")

// Struct1 结构体
type Struct1 struct {
      global.GVA_MODEL
      Field1  globalModel.MultiPic `json:"field1" form:"field1" gorm:"column:field1;comment:field1;type:json;"`
}
// TableName Struct1 表名
func (Struct1) TableName() string {
      return "struct1"
}

