package vm

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
	"testing"
)

func TestRegion(t *testing.T) {

	t.Run("GetRegionList", func(t *testing.T) {
		cli := ecloud.NewForConfigDie(&core.Config{
			ApiGwHost: "api-guiyang-1.cmecloud.cn",
			//ApiGwPort:     8443,
			ApiGwProtocol: "https",
		})

		result, err := cli.Server().GetRegionList()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result.Dump())
	})
}
