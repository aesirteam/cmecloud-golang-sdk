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

	spec := VirtualPrivateCloud.VpcSpec{
		Cidr:        "192.168.101.0/24",
		CidrV6:      "",
		Name:        "vpc99999",
		NetworkName: "subnet_default",
		Region:      "N0851-GZ-GYGZ01",
	}

	getVpcInfoByName := func(name string) VirtualPrivateCloud.VpcResult {
		vpc, err := cli.VPC().GetVpcInfoByName(name)
		if err != nil {
			t.Fatal(err)
		}

		return vpc
	}

	t.Run("CreateVpc", func(t *testing.T) {
		vpcId, err := cli.VPC().CreateVpc(&spec)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("vpcId:%s\n", vpcId)
	})

	t.Run("GetVpcList", func(t *testing.T) {
		result, err := cli.VPC().GetVpcList(false, false, 0, "", nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})

	t.Run("GetVpcInfo", func(t *testing.T) {
		vpc := getVpcInfoByName(spec.Name)

		result, err := cli.VPC().GetVpcInfo(vpc.Id)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})

	t.Run("GetVpcInfoByName", func(t *testing.T) {
		result := getVpcInfoByName(spec.Name)

		t.Log(core.Dump(result))
	})

	t.Run("GetVpcNetwork", func(t *testing.T) {
		vpc := getVpcInfoByName(spec.Name)

		result, err := cli.VPC().GetVpcNetwork(vpc.RouterId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})

}
