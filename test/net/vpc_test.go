package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestVPC(t *testing.T) {

	t.Run("CreateVpc", func(t *testing.T) {
		spec := global.VpcSpec{
			Cidr:        "192.168.101.0/24",
			CidrV6:      "",
			Name:        vpcName,
			NetworkName: subnetName,
			Region:      region,
		}

		result, err := net.CreateVpc(&spec)

		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcList", func(t *testing.T) {
		result, err := net.GetVpcList("", "", false, 0, nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcInfo", func(t *testing.T) {
		vpcId, _ = getRouterId()

		result, err := net.GetVpcInfo(vpcId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcInfoByRouterId", func(t *testing.T) {
		_, routerId = getRouterId()

		result, err := net.GetVpcInfoByRouterId(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifyVpcInfo", func(t *testing.T) {
		vpcId, _ = getRouterId()

		err := net.ModifyVpcInfo(vpcId, vpcName, "ModifyVpcInfo")
		if err != nil {
			t.Fatal(err)
		}
	})

	// t.Run("GetVpcFirewall", func(t *testing.T) {
	// 	result, err := net.GetVpcFirewall(vpc.RouterId)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	t.Log(result)
	// })

	t.Run("GetVpcNetwork", func(t *testing.T) {
		_, routerId = getRouterId()

		result, err := net.GetVpcNetwork("", routerId, nil, nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcIPSecVPN", func(t *testing.T) {
		_, routerId = getRouterId()

		result, err := net.GetVpcVPN(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcNic", func(t *testing.T) {
		_, routerId = getRouterId()

		result, err := net.GetVpcNic(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("DeleteVpc", func(t *testing.T) {
		vpcId, _ = getRouterId()

		err := net.DeleteVpc(vpcId)
		if err != nil {
			t.Fatal(err)
		}
	})
}
