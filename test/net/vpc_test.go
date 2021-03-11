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

	spec := global.VpcSpec{
		Cidr:        "192.168.101.0/24",
		CidrV6:      "",
		Name:        "vpc99999",
		NetworkName: "subnet_default",
		Region:      "N0851-GZ-GYGZ01",
	}

	getVpcInfoByName := func(name string) VirtualPrivateCloud.VpcResult {
		vpc, err := net.GetVpcInfoByName(name)
		if err != nil {
			t.Fatal(err)
		}

		return vpc
	}

	t.Run("CreateVpc", func(t *testing.T) {
		vpcId, err := net.CreateVpc(&spec)

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

	t.Run("GetVpcInfo", func(t *testing.T) {
		vpc := getVpcInfoByName(spec.Name)

		result, err := net.GetVpcInfo(vpc.Id)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcInfoByName", func(t *testing.T) {
		result := getVpcInfoByName(spec.Name)

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcInfoByRouterId", func(t *testing.T) {
		vpc := getVpcInfoByName(spec.Name)

		result, err := net.GetVpcInfoByRouterId(vpc.RouterId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifyVpcInfo", func(t *testing.T) {
		vpc := getVpcInfoByName(spec.Name)

		err := net.ModifyVpcInfo(vpc.Id, spec.Name, "ModifyVpcInfo")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetVpcFirewall", func(t *testing.T) {
		vpc := getVpcInfoByName(spec.Name)

		result, err := net.GetVpcFirewall(vpc.RouterId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result)
	})

	t.Run("GetVpcNetwork", func(t *testing.T) {
		vpc := getVpcInfoByName(spec.Name)

		result, err := net.GetVpcNetwork(vpc.RouterId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcIPSecVPN", func(t *testing.T) {
		vpc := getVpcInfoByName(spec.Name)

		result, err := net.GetVpcVPN(vpc.RouterId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpcNic", func(t *testing.T) {
		vpc := getVpcInfoByName(spec.Name)

		result, err := net.GetVpcNic(vpc.RouterId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})
}
