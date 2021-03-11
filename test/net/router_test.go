package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

func TestRouter(t *testing.T) {
	net := ecloud.NewForConfigDie(&core.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	}).Net()

	getVpcInfoByName := func(name string) VirtualPrivateCloud.VpcResult {
		vpc, err := net.GetVpcInfoByName(name)
		if err != nil {
			t.Fatal(err)
		}

		return vpc
	}

	var vpcName = "vpc99999"

	t.Run("GetRouterNetList", func(t *testing.T) {
		vpc := getVpcInfoByName(vpcName)

		result, err := net.GetRouterNetList(vpc.RouterId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})

	t.Run("GetRouterInfo", func(t *testing.T) {
		vpc := getVpcInfoByName(vpcName)

		result, err := net.GetRouterInfo(vpc.RouterId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})
}
