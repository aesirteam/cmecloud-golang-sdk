package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestKeypair(t *testing.T) {

	t.Run("CreateKeypair", func(t *testing.T) {
		result, err := vm.CreateKeypair(keypairName, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result)
	})

	t.Run("GetKeypairList", func(t *testing.T) {
		result, err := vm.GetKeypairList("", "", 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("DeleteKeypair", func(t *testing.T) {
		keypairId = getKeypairId()

		err := vm.DeleteKeypair(keypairId)
		if err != nil {
			t.Fatal(err)
		}
	})
}
