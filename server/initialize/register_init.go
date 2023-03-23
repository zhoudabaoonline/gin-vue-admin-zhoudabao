package initialize

// 直接引入需要的包,使用包里的Init函数初始化数据
import (
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/example"
	_ "github.com/flipped-aurora/gin-vue-admin/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
