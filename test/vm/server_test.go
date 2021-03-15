package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

func TestServer(t *testing.T) {
	cli := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	})

	vm := cli.VM()
	net := cli.Net()

	var (
		name             = "api198994873"
		password         = "ECS@test1234"
		rebuildImageId   = "c82bb33a-69e0-4f10-b8e0-856502066384"
		vpcName          = "vpc99999"
		portName         = "port99999"
		serverId, portId string
	)

	vpc := func(name string) VirtualPrivateCloud.VpcResult {
		vpc, err := net.GetVpcInfoByName(name)
		if err != nil {
			t.Fatal(err)
		}

		return vpc
	}

	t.Run("CreateServer", func(t *testing.T) {
		spec := global.ServerSpec{
			Name: name,
			Cpu:  4,
			Ram:  8,
			// BootVolume: struct {
			// 	VolumeType global.BootVolumeType
			// 	Size       int
			// }{global.BOOT_VOLUME_PERFORMANCEOPTIMIZATION, 50},
			Password:  password,
			ImageType: global.IMAGE_BCLINUX_76_X64,
			//KeypairName:      "",
			// Networks: []struct{NetworkId string; PortId string}{
			// 	{"63ba602a-c6f6-4975-a507-017979346dee", ""},
			// },
			// SecurityGroupIds: []string{"6f2f0f0d-30a0-4ccb-bcc4-81e83e713b5e"},
			//UserData:         "",
			//VmType:           0,
			//Region:           "",
			// BillingType: global.BILLING_TYPE_YEAR,
			//Dration:          0,
			//Quantity:         0,
			// DataVolumes: []struct {
			// 	ResourceType global.DataVolumeType
			// 	Size         int
			// 	IsShare      bool
			// }{
			// 	{global.DATA_VOLUME_CAPEBS, 10, false},
			// 	{global.DATA_VOLUME_SSDEBS, 50, false},
			// },
			// OsType: global.OS_TYPE_WINDOWS,
		}

		//查询networkId
		vpc_default := vpc("vpc_default")

		if vpc_default.NetworkId == "" {
			t.Fatal("No found NetworkId")
		}

		//t.Log(vpc.Dump())
		spec.Networks = []struct {
			NetworkId string
			PortId    string
		}{
			{vpc_default.NetworkId, ""},
		}

		spec.Region = vpc_default.Region

		//步骤1：查询云主机可用规格
		product_arr, err := vm.GetProductFlavorList(&spec, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		if len(product_arr) == 0 {
			t.Fatal("No match Product")
		}

		//t.Log(product_arr.Dump())

		result, err := vm.CreatServer(&spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		//for _, s := range result.OrderExts {
		//	t.Log(s)
		//}
	})

	t.Run("GetServerList", func(t *testing.T) {
		spec := global.ServerSpec{
			Name: name,
		}

		result, err := vm.GetServerList(&spec, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		serverId = result[0].ServerId
	})

	getServerId := func() string {
		result, err := vm.GetServerList(&global.ServerSpec{Name: name}, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		return result[0].ServerId
	}

	t.Run("GetServerInfo", func(t *testing.T) {
		if serverId == "" {
			serverId = getServerId()
		}

		result, err := vm.GetServerInfo(serverId, true)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetServerVNCAddress", func(t *testing.T) {
		if serverId == "" {
			serverId = getServerId()
		}

		addr, err := vm.GetServerVNCAddress(serverId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(addr)
	})

	t.Run("ModifyServerName", func(t *testing.T) {
		if serverId == "" {
			serverId = getServerId()
		}

		result, err := vm.ModifyServerName(serverId, name)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifyServerPassword", func(t *testing.T) {
		if serverId == "" {
			serverId = getServerId()
		}

		result, err := vm.ModifyServerPassword(serverId, password)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("RebootServer", func(t *testing.T) {
		if serverId == "" {
			serverId = getServerId()
		}

		result, err := vm.RebootServer(serverId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("StartServer", func(t *testing.T) {
		if serverId == "" {
			serverId = getServerId()
		}

		result, err := vm.StartServer(serverId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("StopServer", func(t *testing.T) {
		if serverId == "" {
			serverId = getServerId()
		}

		result, err := vm.StopServer(serverId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetRebuildImageList", func(t *testing.T) {
		if serverId == "" {
			serverId = getServerId()
		}

		result, err := vm.GetRebuildImageList(serverId, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		for _, v := range result {
			if v.Name == global.IMAGE_BCLINUX_77_X64.String() {
				rebuildImageId = v.Id
				return
			}
		}
	})

	t.Run("RebuildServer", func(t *testing.T) {
		if serverId == "" {
			serverId = getServerId()
		}

		err := vm.RebuildServer(serverId, rebuildImageId, "", "")
		if err != nil {
			t.Fatal(err)
		}
	})

	getPortId := func() string {
		routerId := vpc(vpcName).RouterId

		nicList, err := net.GetVpcNic(routerId)
		if err != nil {
			t.Fatal(err)
		}

		for _, v := range nicList {
			if v.Name == portName {
				return v.FixedIps[0].PortId
			}
		}

		return ""
	}

	t.Run("AttachNic", func(t *testing.T) {
		if serverId == "" {
			serverId = getServerId()
		}

		if portId == "" {
			portId = getPortId()
		}

		result, err := vm.AttachNic(serverId, portId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("DetachNic", func(t *testing.T) {
		if serverId == "" {
			serverId = getServerId()
		}

		if portId == "" {
			portId = getPortId()
		}

		result, err := vm.DetachNic(serverId, portId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetUnbindNicList", func(t *testing.T) {
		if serverId == "" {
			serverId = getServerId()
		}

		result, err := vm.GetUnbindNicList(serverId, 0, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result)
	})
}
