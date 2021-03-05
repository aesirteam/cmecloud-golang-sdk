package vm

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/vm/Server"
	"strconv"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	cli := ecloud.NewForConfigDie(&core.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	})

	t.Run("CreatServer", func(t *testing.T) {
		spec := &Server.ServerSpec{
			Name: "api" + strconv.Itoa(time.Now().Nanosecond()),
			Cpu:  4,
			Ram:  8,
			//Disk:             0,
			//BootVolume:       Server.BootVolume{
			//	VolumeType: Server.BOOT_VOLUME_PERFORMANCEOPTIMIZATION,
			//	Size: 50,
			//},
			Password:  "ECS@test123",
			ImageType: Server.IMAGE_BCLINUX_76_X64,
			//KeypairName:      "",
			//Networks:         Server.Networks{
			//	NetworkId: "63ba602a-c6f6-4975-a507-017979346dee",
			//},
			//SecurityGroupIds: nil,
			//UserData:         "",
			//VmType:           0,
			//Region:           "",
			//BillingType:      BILLING_TYPE_YEAR,
			//Dration:          0,
			//Quantity:         0,
			//DataVolumeArray:  []Server.DataVolume{
			//	{ ResourceType: Server.DATA_VOLUME_CAPEBS, Size: 10 },
			//},
			//OsType:           OS_TYPE_WINDOWS,
		}

		//查询networkId
		vpc, err := cli.VPC().GetVPCInfoByName("vpc_default")
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
		product_arr, err := cli.Server().GetProductFlavorList(spec, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		if len(product_arr) == 0 {
			t.Fatal("No match Product")
		}

		//t.Log(product_arr.Dump())

		result, err := cli.Server().CreatServer(spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result.ProcedureCode, result.OrderId)

		//for _, s := range result.OrderExts {
		//	t.Log(s)
		//}
	})

	t.Run("GetServerList", func(t *testing.T) {
		spec := &Server.ServerSpec{
			Name: "api198994873",
		}

		resp, err := cli.Server().GetServerList(spec, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(resp)
	})
}
