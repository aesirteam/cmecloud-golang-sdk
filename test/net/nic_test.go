package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/SecurityGroup"
)

func TestNic(t *testing.T) {
	cli := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	})

	net := cli.Net()

	var (
		vpcName  = "vpc99999"
		portName = "port99999"
		portId   string
	)

	vpc := func() VirtualPrivateCloud.VpcResult {
		vpc, err := net.GetVpcInfoByName(vpcName)
		if err != nil {
			t.Fatal(err)
		}

		return *vpc
	}()

	subnet := func() VirtualPrivateCloud.SubnetResult {
		result, err := net.GetSubnetList(vpc.FirstNetworkId)
		if err != nil {
			t.Fatal(err)
		}

		if len(result) == 0 {
			t.Fatalf("No match subnet by vpc: %s\n", vpcName)
		}

		return result[0]
	}()

	securityGroups := func() []SecurityGroup.SecurityGroupResult {
		result, err := cli.VM().GetSecurityGroupList("default", false, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		if len(result) == 0 {
			t.Fatal("No match securityGroup: default")
		}

		return result
	}()

	getPortId := func() string {
		if portId == "" {
			if nicList, err := net.GetVpcNic(vpc.RouterId); err == nil {
				for _, v := range nicList {
					if v.Name == portName {
						portId = v.FixedIps[0].PortId
						break
					}
				}
			}
		}

		return portId
	}

	t.Run("CreateNic", func(t *testing.T) {
		spec := global.NicSpec{
			Name:           portName,
			NetworkId:      vpc.FirstNetworkId,
			SecurityGroups: []string{securityGroups[0].Id},
			Subnets: []struct {
				IpAddress string
				SubnetId  string
			}{
				{"", subnet.Id},
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

		Ids := make([]string, len(securityGroups))
		for i, v := range securityGroups {
			Ids[i] = v.Id
		}

		err := net.ModifyNicSecurityGroup(portId, Ids)
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
