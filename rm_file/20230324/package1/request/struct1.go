package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/package1"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type Struct1Search struct{
    package1.Struct1
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
