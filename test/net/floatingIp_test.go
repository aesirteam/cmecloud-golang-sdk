package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestFloatingIP(t *testing.T) {
	cli := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	})

	net := cli.Net()
	vm := cli.VM()

	var (
		serverName = "api198994873"
		fipAddress = "36.137.45.147"
		ipId       string
	)

	getIpId := func() string {
		if ipId == "" {
			if result, err := net.GetFloatingIpList(fipAddress, "", "", false, false, false, false, nil, 0, 0); err == nil && len(result) > 0 {
				ipId = result[0].Id
			}
		}

		return ipId
	}

	t.Run("CreateFloatingIp", func(t *testing.T) {
		spec := global.FloatingIpSpec{
			ChargeMode:        global.FLOATINGIP_CHARGEMODE_TRAFFIC,
			ChargePeriod:      global.BILLING_TYPE_HOUR,
			BandwidthSize:     10,
			FloatingIpAddress: fipAddress,
		}

		result, err := net.CreateFloatingIp(&spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result)
	})

	t.Run("GetFloatingIpList", func(t *testing.T) {
		result, err := net.GetFloatingIpList("", "", "", false, false, false, false, nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetFloatingIpDetail", func(t *testing.T) {
		ipId = getIpId()

		result, err := net.GetFloatingIpDetail(ipId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetFloatingIpQuotaWithK", func(t *testing.T) {
		result, err := net.GetFloatingIpQuotaWithK()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetFloatingIpQuotaWithP", func(t *testing.T) {
		result, err := net.GetFloatingIpQuotaWithP(global.FLOATINGIP_QUOTANAME_IP, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("AttachFloatingIp", func(t *testing.T) {
		ipId = getIpId()

		resources, err := vm.GetServerList(&global.ServerSpec{
			Name: serverName,
		}, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		if len(resources) == 0 {
			t.Fatalf("No match server: %s", serverName)
		}

		portDetails := resources[0].ServerPortDetailArray
		if len(portDetails) == 0 {
			t.Fatal("No match portDetail")
		}

		err = net.AttachFloatingIp(global.FLOATINGIP_BINDTYPE_VM, ipId, resources[0].ServerId, portDetails[0].PortId)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DetachFloatingIp", func(t *testing.T) {
		ipId = getIpId()

		err := net.DetachFloatingIp(ipId)
		if err != nil {
			t.Fatal(err)
		}
	})
}
