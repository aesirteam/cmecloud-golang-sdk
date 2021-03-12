package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestServer(t *testing.T) {
	cli := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	})

	vm := cli.VM()

	var (
		name     = "api198994873"
		password = "ECS@test1234"
	)

	t.Run("CreatServer", func(t *testing.T) {
		spec := global.ServerSpec{
			Name: name,
			Cpu:  4,
			Ram:  8,
			//Disk:             0,
			BootVolumeSpec: global.BootVolumeSpec{
				VolumeType: global.BOOT_VOLUME_PERFORMANCEOPTIMIZATION,
				Size:       20,
			},
			Password:  password,
			ImageType: global.IMAGE_BCLINUX_76_X64,
			//KeypairName:      "",
			// NetworkSpec: global.NetworkSpec{
			// 	NetworkId: "63ba602a-c6f6-4975-a507-017979346dee",
			// },
			//SecurityGroupIds: nil,
			//UserData:         "",
			//VmType:           0,
			//Region:           "",
			// BillingType: global.BILLING_TYPE_YEAR,
			//Dration:          0,
			//Quantity:         0,
			// DataVolumes: []global.DataVolumeSpec{
			// 	{ResourceType: global.DATA_VOLUME_CAPEBS, Size: 10},
			// },
			// OsType: global.OS_TYPE_WINDOWS,
		}

		//查询networkId
		vpc, err := cli.Net().GetVpcInfoByName("vpc_default")
		if err != nil {
			t.Fatal(err)
		}

		if vpc.NetworkId == "" {
			t.Fatal("No found NetworkId")
		}

		//t.Log(vpc.Dump())
		spec.NetworkId = vpc.NetworkId
		spec.Region = vpc.Region

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

	var serverId = "eebae258-90d9-4926-aed9-3e04cf480a9a"

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

	t.Run("GetServerInfo", func(t *testing.T) {
		result, err := vm.GetServerInfo(serverId, true)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetServerVNCAddress", func(t *testing.T) {
		addr, err := vm.GetServerVNCAddress(serverId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(addr)
	})

	t.Run("ModifyServerName", func(t *testing.T) {
		result, err := vm.ModifyServerName(serverId, name)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifyServerPassword", func(t *testing.T) {
		result, err := vm.ModifyServerPassword(serverId, password)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("RebootServer", func(t *testing.T) {
		result, err := vm.RebootServer(serverId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("StartServer", func(t *testing.T) {
		result, err := vm.StartServer(serverId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("StopServer", func(t *testing.T) {
		result, err := vm.StopServer(serverId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	var imageId = "c82bb33a-69e0-4f10-b8e0-856502066384"

	t.Run("GetRebuildImageList", func(t *testing.T) {
		result, err := vm.GetRebuildImageList(serverId, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		for _, v := range result {
			if v.Name == global.IMAGE_BCLINUX_77_X64.String() {
				imageId = v.Id
				return
			}
		}
	})

	t.Run("RebuildServer", func(t *testing.T) {
		err := vm.RebuildServer(serverId, imageId, "", "")
		if err != nil {
			t.Fatal(err)
		}
	})

}
