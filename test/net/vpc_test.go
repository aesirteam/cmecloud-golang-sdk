package net

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
	"testing"
)

func TestVPC(t *testing.T) {
	cli := ecloud.NewForConfigDie(&core.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	})

	t.Run("GetVPCList", func(t *testing.T) {
		result, err := cli.VPC().GetVPCList(false, false, 0, "", nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result.Dump())
	})

	t.Run("GetVPCInfo", func(t *testing.T) {
		result, err := cli.VPC().GetVPCInfo("3b91fe26f37f4bc3b0da730ba7cbec33")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result.Dump())
	})

	t.Run("GetVPCInfoByName", func(t *testing.T) {
		result, err := cli.VPC().GetVPCInfoByName("vpc_default")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result.Dump())
	})
}
