package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

func TestRouter(t *testing.T) {
	cli := ecloud.NewForConfigDie(&core.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	})

	getVpcInfoByName := func(name string) VirtualPrivateCloud.VpcResult {
		vpc, err := cli.VPC().GetVpcInfoByName(name)
		if err != nil {
			t.Fatal(err)
		}

		return vpc
	}

	var vpcName = "vpc99999"

	t.Run("GetRouterNetList", func(t *testing.T) {
		vpc := getVpcInfoByName(vpcName)

		result, err := cli.VPC().GetRouterNetList(vpc.RouterId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})

	t.Run("GetRouterInfo", func(t *testing.T) {
		vpc := getVpcInfoByName(vpcName)

		result, err := cli.VPC().GetRouterInfo(vpc.RouterId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})
}
