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

	var (
		vpcName             = "vpc99999"
		region              = "N0851-GZ-GYGZ01"
		networkName         = "subnet_test"
		networkId, subnetId string
	)

	vpc := func() VirtualPrivateCloud.VpcResult {
		vpc, err := net.GetVpcInfoByName(vpcName)
		if err != nil {
			t.Fatal(err)
		}

		return *vpc
	}()

	getNetworkId := func() (string, string) {
		if networkId == "" || subnetId == "" {
			if result, err := net.GetVpcNetwork(vpc.RouterId); err == nil {
				for _, v := range result {
					if v.Name == networkName {
						networkId, subnetId = v.Subnets[0].NetworkId, v.Subnets[0].Id
						break
					}
				}
			}
		}

		return networkId, subnetId
	}

	t.Run("CreateSubnet", func(t *testing.T) {
		subnetSpec := global.SubnetSpec{
			RouterId:    vpc.RouterId,
			NetworkName: networkName,
			Region:      region,
			Subnets: []struct {
				Cidr       string
				IpVersion  int
				SubnetName string
			}{
				{"10.18.0.0/24", 4, networkName},
			},
		}

		var err error
		networkId, err = net.CreateSubnet(&subnetSpec)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("netoworkId:%s\n", networkId)
	})

	t.Run("GetSubnetList", func(t *testing.T) {
		networkId, subnetId = getNetworkId()

		result, err := net.GetSubnetList(networkId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetSubnetInfo", func(t *testing.T) {
		networkId, subnetId = getNetworkId()

		result, err := net.GetSubnetInfo(subnetId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifySubnetName", func(t *testing.T) {
		networkId, _ = getNetworkId()

		err := net.ModifySubnetName(networkId, networkName)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteSubnet", func(t *testing.T) {
		networkId, _ = getNetworkId()

		err := net.DeleteSubnet(networkId)
		if err != nil {
			t.Fatal(err)
		}
	})
}
