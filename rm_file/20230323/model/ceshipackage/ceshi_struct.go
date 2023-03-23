// 自动生成模板CeshiStruct
package ceshipackage

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// CeshiStruct 结构体
type CeshiStruct struct {
      global.GVA_MODEL
      CeshiField  string `json:"ceshiField" form:"ceshiField" gorm:"column:ceshi_field;comment:ceshiField;"`
}


// TableName CeshiStruct 表名
func (CeshiStruct) TableName() string {
  return "ceshiStruct"
}

