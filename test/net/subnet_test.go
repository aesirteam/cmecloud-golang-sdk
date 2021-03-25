package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestSubnet(t *testing.T) {

	var (
		networkId, routerId = "", ""
	)

	t.Run("CreateSubnet", func(t *testing.T) {
		_, routerId = getRouterId()

		subnetSpec := global.SubnetSpec{
			RouterId:    routerId,
			NetworkName: subnetName,
			Region:      region,
			Subnets: []struct {
				Cidr       string
				IpVersion  int
				SubnetName string
			}{
				{subnetCidr, 4, subnetName},
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
		networkId, _ = getNetworkId()

		result, err := net.GetSubnetList(networkId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetSubnetInfo", func(t *testing.T) {
		_, subnetId = getNetworkId()

		result, err := net.GetSubnetInfo(subnetId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifySubnetName", func(t *testing.T) {
		networkId, _ = getNetworkId()

		err := net.ModifySubnetName(networkId, subnetName)
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
