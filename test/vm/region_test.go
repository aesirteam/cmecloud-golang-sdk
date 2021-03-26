package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestRegion(t *testing.T) {

	t.Run("GetRegionList", func(t *testing.T) {
		result, err := vm.GetRegionList("")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})
}
