package system

import (
	"fmt"
	"strings"
	"testing"
)

func Test_injectionCode(t *testing.T) {
	Init("system")
	injectionCodeMeta := strings.Builder{}

	fmt.Println("顶顶顶顶顶顶顶")
	err := injectionCode("hahahahahahahah", &injectionCodeMeta)

	fmt.Println(err)
}
