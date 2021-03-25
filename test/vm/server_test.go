package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestServer(t *testing.T) {

	var (
		password       = "ECS@test1234"
		rebuildImageId = "c82bb33a-69e0-4f10-b8e0-856502066384"
	)

	t.Run("CreateServer", func(t *testing.T) {
		networkId, _ = getNetworkId()
		// securityGroupIds := getDefaultSecurityGroupIds()

		spec := global.ServerSpec{
			Name: serverName,
			Cpu:  4,
			Ram:  8,
			// BootVolume: struct {
			// 	VolumeType global.BootVolumeType
			// 	Size       int
			// }{global.BOOT_VOLUME_PERFORMANCEOPTIMIZATION, 50},
			Password:  password,
			ImageName: global.IMAGE_BCLINUX_76_X64.String(),
			// KeypairName:      keypairName,
			Networks: &struct {
				NetworkId string
				PortId    string
			}{
				networkId, "",
			},
			// SecurityGroupIds: securityGroupIds,
			// VmType: global.VM_TYPE_COMMON,
			Region: region,
			// BillingType: global.BILLING_TYPE_YEAR,
			//Duration: 0,
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

		//步骤1：查询云主机可用规格
		product_arr, err := vm.GetProductFlavorList(spec.Cpu, spec.Ram, spec.OsType, spec.VmType, 0, 0)
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
		result, err := vm.GetServerList("", "", "", 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		serverId = result[0].ServerId
	})

	t.Run("GetServerInfo", func(t *testing.T) {
		serverId = getServerId()

		result, err := vm.GetServerInfo(serverId, true)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetServerVNCAddress", func(t *testing.T) {
		serverId = getServerId()

		addr, err := vm.GetServerVNCAddress(serverId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(addr)
	})

	t.Run("ModifyServerName", func(t *testing.T) {
		serverId = getServerId()

		result, err := vm.ModifyServerName(serverId, serverName)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifyServerPassword", func(t *testing.T) {
		serverId = getServerId()

		result, err := vm.ModifyServerPassword(serverId, password)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("RebootServer", func(t *testing.T) {
		serverId = getServerId()

		result, err := vm.RebootServer(serverId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("StartServer", func(t *testing.T) {
		serverId = getServerId()

		result, err := vm.StartServer(serverId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("StopServer", func(t *testing.T) {
		serverId = getServerId()

		result, err := vm.StopServer(serverId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetRebuildImageList", func(t *testing.T) {
		serverId = getServerId()

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
		serverId = getServerId()

		err := vm.RebuildServer(serverId, rebuildImageId, "", "")
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("AttachNic", func(t *testing.T) {
		serverId = getServerId()
		portId = getPortId()

		result, err := vm.AttachNic(serverId, portId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("DetachNic", func(t *testing.T) {
		serverId = getServerId()
		portId = getPortId()

		result, err := vm.DetachNic(serverId, portId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetUnbindNicList", func(t *testing.T) {
		serverId = getServerId()

		result, err := vm.GetUnbindNicList(serverId, 0, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})
}
