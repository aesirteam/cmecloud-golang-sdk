package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestProductFlavor(t *testing.T) {

	t.Run("GetProductFlavorList", func(t *testing.T) {
		result, err := vm.GetProductFlavorList(4, 8, global.OS_TYPE_LINUX, global.VM_TYPE_COMMON, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})
}
