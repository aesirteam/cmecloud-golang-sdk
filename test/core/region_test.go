package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestRegion(t *testing.T) {

	t.Run("GetRegionList", func(t *testing.T) {
		core := ecloud.NewForConfigDie(&global.Config{
			ApiGwHost: "api-guiyang-1.cmecloud.cn",
			//ApiGwPort:     8443,
			ApiGwProtocol: "https",
		}).Core()

		result, err := core.GetRegionList()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})
}
