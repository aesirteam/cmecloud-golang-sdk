package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

func TestSubnet(t *testing.T) {
	net := ecloud.NewForConfigDie(&global.Config{
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
		networkId, err = net.CreateSubnet(&subnetSpec)
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

		result, err := net.GetSubnetInfo(networkId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifySubnet", func(t *testing.T) {
		vpc := getVpcInfoByName(vpcName)
		if networkId == "" {
			networkId = vpc.NetworkId
		}

		err := net.ModifySubnet(networkId, "subnet99")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteSubnet", func(t *testing.T) {
		vpc := getVpcInfoByName(vpcName)
		if networkId == "" {
			networkId = vpc.NetworkId
		}

		err := net.DeleteSubnet(networkId)
		if err != nil {
			t.Fatal(err)
		}
	})
}
