package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
)

func TestVPN(t *testing.T) {
	net := ecloud.NewForConfigDie(&core.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	}).Net()

	// getVpcInfoByName := func(name string) VirtualPrivateCloud.VpcResult {
	// 	vpc, err := net.GetVpcInfoByName(name)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	return vpc
	// }
	t.Run("GetIPSecVpnList", func(t *testing.T) {
		result, err := net.GetIpsecVpnList("", "", 0, "", 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})

	t.Run("GetIPSecVpnInfo", func(t *testing.T) {
		result, err := net.GetIpsecVpnInfo("71e74bbb-e0c0-4950-85dd-ef7703dc1dd6")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})
}
