package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestNic(t *testing.T) {

	var (
		securityGroupIds = getDefaultSecurityGroupIds()
		networkId        = ""
	)

	t.Run("CreateNic", func(t *testing.T) {
		networkId, subnetId = getNetworkId()

		spec := global.NicSpec{
			Name:           portName,
			NetworkId:      networkId,
			SecurityGroups: securityGroupIds,
			Subnets: []struct {
				IpAddress string
				SubnetId  string
			}{
				{"", subnetId},
			},
		}

		var err error
		portId, err = net.CreateNic(&spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("portId:%s\n", portId)

	})

	t.Run("GetNicDetail", func(t *testing.T) {
		portId = getPortId()

		result, err := net.GetNicDetail(portId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifyNicName", func(t *testing.T) {
		portId = getPortId()

		err := net.ModifyNicName(portId, portName)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ModifyNicSecurityGroup", func(t *testing.T) {
		portId = getPortId()

		err := net.ModifyNicSecurityGroup(portId, securityGroupIds)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteNic", func(t *testing.T) {
		portId = getPortId()

		err := net.DeleteNic(portId)
		if err != nil {
			t.Fatal(err)
		}
	})
}
