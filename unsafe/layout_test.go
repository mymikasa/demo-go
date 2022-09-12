package unsafe

import (
	"fmt"
	"github.com/mymikasa/demo-go/unsafe/types"
	"testing"
	"unsafe"
)

func TestPrintFieldOffset(t *testing.T) {
	fmt.Println(unsafe.Sizeof(types.User{}))
	PrintFieldOffset(types.User{})

	fmt.Println(unsafe.Sizeof(types.UserV1{}))
	PrintFieldOffset(types.UserV1{})

	fmt.Println(unsafe.Sizeof(types.UserV2{}))
	PrintFieldOffset(types.UserV2{})
}
