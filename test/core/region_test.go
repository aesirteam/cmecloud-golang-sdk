package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
)

func TestRegion(t *testing.T) {

	t.Run("GetRegionList", func(t *testing.T) {
		Core := ecloud.NewForConfigDie(&core.Config{
			ApiGwHost: "api-guiyang-1.cmecloud.cn",
			//ApiGwPort:     8443,
			ApiGwProtocol: "https",
		}).Core()

		result, err := Core.GetRegionList()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})
}
