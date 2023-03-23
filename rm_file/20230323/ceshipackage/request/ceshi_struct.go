package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ceshipackage"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type CeshiStructSearch struct{
    ceshipackage.CeshiStruct
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
