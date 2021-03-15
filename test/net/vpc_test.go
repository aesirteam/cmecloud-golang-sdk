package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

func TestVPC(t *testing.T) {
	net := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	}).Net()

	var (
		vpcName     = "vpc99999"
		region      = "N0851-GZ-GYGZ01"
		networkName = "subnet_default"
		vpcId       string
	)

	t.Run("CreateVpc", func(t *testing.T) {
		spec := global.VpcSpec{
			Cidr:        "192.168.101.0/24",
			CidrV6:      "",
			Name:        vpcName,
			NetworkName: networkName,
			Region:      region,
		}

		var err error
		vpcId, err = net.CreateVpc(&spec)

		if err != nil {
			t.Fatal(err)
		}

		t.Logf("vpcId:%s\n", vpcId)
	})

	t.Run("GetVpcList", func(t *testing.T) {
		result, err := net.GetVpcList(false, 0, "", nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	vpc := func() VirtualPrivateCloud.VpcResult {
		vpc, err := net.GetVpcInfoByName(vpcName)
		if err != nil {
			t.Fatal(err)
		}

		return vpc
	}

	t.Run("GetVpcInfo", func(t *testing.T) {
		if vpcId == "" {
			vpcId = vpc().Id
		}

		result, err := net.GetVpcInfo(vpcId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcInfoByName", func(t *testing.T) {
		t.Log(global.Dump(vpc))
	})

	t.Run("GetVpcInfoByRouterId", func(t *testing.T) {
		routerId := vpc().RouterId

		result, err := net.GetVpcInfoByRouterId(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifyVpcInfo", func(t *testing.T) {
		if vpcId == "" {
			vpcId = vpc().Id
		}

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
		routerId := vpc().RouterId

		result, err := net.GetVpcNetwork(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcIPSecVPN", func(t *testing.T) {
		routerId := vpc().RouterId

		result, err := net.GetVpcVPN(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcNic", func(t *testing.T) {
		routerId := vpc().RouterId

		result, err := net.GetVpcNic(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("DeleteVpc", func(t *testing.T) {
		if vpcId == "" {
			vpcId = vpc().Id
		}

		err := net.DeleteVpc(vpcId)
		if err != nil {
			t.Fatal(err)
		}
	})
}
