package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/Server"
)

func TestProductFlavor(t *testing.T) {

	t.Run("GetProductFlavorList", func(t *testing.T) {
		vm := ecloud.NewForConfigDie(&core.Config{
			ApiGwHost: "api-guiyang-1.cmecloud.cn",
			//ApiGwPort:     8443,
			ApiGwProtocol: "https",
		}).VM()

		result, err := vm.GetProductFlavorList(&Server.ServerSpec{Cpu: 4, Ram: 8}, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})
}
