package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestVPC(t *testing.T) {
	net := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	}).Net()

	var (
		vpcName         = "vpc99999"
		region          = "N0851-GZ-GYGZ01"
		networkName     = "subnet_default"
		vpcId, routerId string
	)

	getVpcId := func() (string, string) {
		if vpcId == "" || routerId == "" {
			if result, err := net.GetVpcInfoByName(vpcName); err == nil && result != nil {
				vpcId, routerId = result.Id, result.RouterId
			}
		}

		return vpcId, routerId
	}

	t.Run("CreateVpc", func(t *testing.T) {
		spec := global.VpcSpec{
			Cidr:        "192.168.101.0/24",
			CidrV6:      "",
			Name:        vpcName,
			NetworkName: networkName,
			Region:      region,
		}

		result, err := net.CreateVpc(&spec)

		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcList", func(t *testing.T) {
		result, err := net.GetVpcList(false, 0, "", nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcInfo", func(t *testing.T) {
		vpcId, routerId = getVpcId()

		result, err := net.GetVpcInfo(vpcId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcInfoByName", func(t *testing.T) {
		result, err := net.GetVpcInfoByName(vpcName)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcInfoByRouterId", func(t *testing.T) {
		vpcId, routerId = getVpcId()

		result, err := net.GetVpcInfoByRouterId(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifyVpcInfo", func(t *testing.T) {
		vpcId, routerId = getVpcId()

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
		_, routerId = getVpcId()

		result, err := net.GetVpcNetwork(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcIPSecVPN", func(t *testing.T) {
		_, routerId = getVpcId()

		result, err := net.GetVpcVPN(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcNic", func(t *testing.T) {
		_, routerId = getVpcId()

		result, err := net.GetVpcNic(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("DeleteVpc", func(t *testing.T) {
		vpcId, _ = getVpcId()

		err := net.DeleteVpc(vpcId)
		if err != nil {
			t.Fatal(err)
		}
	})
}
