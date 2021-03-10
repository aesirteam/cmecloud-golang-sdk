package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

func TestSubnet(t *testing.T) {
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

	var (
		vpcName   = "vpc99999"
		region    = "N0851-GZ-GYGZ01"
		networkId string
	)

	t.Run("CreateSubnet", func(t *testing.T) {
		vpc := getVpcInfoByName(vpcName)

		subnetSpec := VirtualPrivateCloud.SubnetSpec{
			RouterId:    vpc.RouterId,
			NetworkName: "subnet1",
			Region:      region,
			Subnets: []struct {
				Cidr       string
				IpVersion  int
				SubnetName string
			}{
				{"10.18.0.0/24", 4, "subnet1"},
			},
		}

		var err error
		networkId, err = cli.VPC().CreateSubnet(&subnetSpec)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("netoworkId:%s\n", networkId)
	})

	t.Run("GetSubnetInfo", func(t *testing.T) {
		vpc := getVpcInfoByName(vpcName)
		if networkId == "" {
			networkId = vpc.NetworkId
		}

		result, err := cli.VPC().GetSubnetInfo(networkId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(core.Dump(result))
	})

	t.Run("ModifySubnet", func(t *testing.T) {
		vpc := getVpcInfoByName(vpcName)
		if networkId == "" {
			networkId = vpc.NetworkId
		}

		err := cli.VPC().ModifySubnet(networkId, "subnet99")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteSubnet", func(t *testing.T) {
		vpc := getVpcInfoByName(vpcName)
		if networkId == "" {
			networkId = vpc.NetworkId
		}

		err := cli.VPC().DeleteSubnet(networkId)
		if err != nil {
			t.Fatal(err)
		}
	})
}