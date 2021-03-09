package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

func TestVPC(t *testing.T) {
	cli := ecloud.NewForConfigDie(&core.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	})

	t.Run("CreateVpc", func(t *testing.T) {
		result, err := cli.VPC().CreatVpc(VirtualPrivateCloud.VpcSpec{
			Name:        "vpc88888",
			NetworkName: "network1",
			Region:      "N0851-GZ-GYGZ01",
			Subnets: []VirtualPrivateCloud.SubnetSpec{
				{SubnetName: "network1", IpVersion: 4, Cidr: "192.168.18.0/24"},
			},
		})

		if err != nil {
			t.Fatal(err)
		}

		t.Log(result)
	})

	t.Run("GetVpcList", func(t *testing.T) {
		result, err := cli.VPC().GetVpcList(false, false, 0, "", nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})

	t.Run("GetVpcInfo", func(t *testing.T) {
		result, err := cli.VPC().GetVpcInfo("3b91fe26f37f4bc3b0da730ba7cbec33")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})

	t.Run("GetVpcInfoByName", func(t *testing.T) {
		result, err := cli.VPC().GetVpcInfoByName("vpc88888")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})

	t.Run("CreateSubnet", func(t *testing.T) {
		result, err := cli.VPC().CreateSubnet(VirtualPrivateCloud.VpcSpec{
			RouterId:    "0e0c53c0-111c-4e8b-822b-c4f3dcd321aa",
			NetworkName: "network1",
			Region:      "N0851-GZ-GYGZ01",
			Subnets: []VirtualPrivateCloud.SubnetSpec{
				{SubnetName: "network1", IpVersion: 4, Cidr: "10.15.0.0/24"},
			},
		})

		if err != nil {
			t.Fatal(err)
		}

		t.Log(result)
	})
}
