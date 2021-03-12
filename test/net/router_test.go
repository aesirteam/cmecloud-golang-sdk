package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

func TestRouter(t *testing.T) {
	net := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	}).Net()

	var vpcName = "vpc99999"

	vpc := func() VirtualPrivateCloud.VpcResult {
		vpc, err := net.GetVpcInfoByName(vpcName)
		if err != nil {
			t.Fatal(err)
		}

		return vpc
	}()

	t.Run("GetRouterNetList", func(t *testing.T) {
		result, err := net.GetRouterNetList(vpc.RouterId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetRouterInfo", func(t *testing.T) {
		result, err := net.GetRouterInfo(vpc.RouterId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})
}
