package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestProductFlavor(t *testing.T) {

	t.Run("GetProductFlavorList", func(t *testing.T) {
		vm := ecloud.NewForConfigDie(&global.Config{
			ApiGwHost: "api-guiyang-1.cmecloud.cn",
			//ApiGwPort:     8443,
			ApiGwProtocol: "https",
		}).VM()

		result, err := vm.GetProductFlavorList(&global.ServerSpec{Cpu: 4, Ram: 8}, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})
}
