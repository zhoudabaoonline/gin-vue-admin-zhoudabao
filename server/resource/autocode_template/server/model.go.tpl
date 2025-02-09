// 自动生成模板{{.StructName}}
package {{.Package}}

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	{{ if .HasTimer }}"time"{{ end }}
      {{- if .HasMultiPic }}
	"github.com/flipped-aurora/gin-vue-admin/server/model/globalModel"
      {{- end -}}
)

// {{.StructName}} 结构体
type {{.StructName}} struct {
      global.GVA_MODEL {{- range .Fields}}
            {{- if eq .FieldType "enum" }}
      {{.FieldName}}  string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};type:enum({{.DataTypeLong}});comment:{{.Comment}};"`
            {{- else if eq .FieldType "fuwenben" }}
      {{.FieldName}}  *string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};type:text;"`
            {{- else if eq .FieldType "multipic" }}
      {{.FieldName}}  globalModel.MultiPic `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};type:json;"`
            {{- else if eq .FieldType "pic" }}
      {{.FieldName}}  *string `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};size:256;"`
            {{- else if ne .FieldType "string" }}
      {{.FieldName}}  *{{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}"`
            {{- else }}
      {{.FieldName}}  {{.FieldType}} `json:"{{.FieldJson}}" form:"{{.FieldJson}}" gorm:"column:{{.ColumnName}};comment:{{.Comment}};{{- if .DataTypeLong -}}size:{{.DataTypeLong}};{{- end -}}"`
            {{- end }}  {{- end }}
      {{- if .AutoCreateResource }}
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
      {{- end}}
}

{{- if .TableName }}
// TableName {{.StructName}} 表名
func ({{.StructName}}) TableName() string {
      return "{{.TableName}}"
}
{{ end }}
