package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestFloatingIP(t *testing.T) {
	var (
		fipAddress   = "36.137.45.147"
		floatingIpId = ""
	)

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

		t.Log(global.Dump(result))
	})

	t.Run("GetFloatingIpList", func(t *testing.T) {
		result, err := net.GetFloatingIpList("", "", "", false, false, false, false, nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetFloatingIpDetail", func(t *testing.T) {
		floatingIpId = getFloatingIpId(fipAddress)

		result, err := net.GetFloatingIpDetail(floatingIpId)
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
		floatingIpId = getFloatingIpId(fipAddress)
		serverPort := getServerPort()

		if serverPort == nil {
			t.Fatal("No match serverPort")
		}

		err := net.AttachFloatingIp(global.FLOATINGIP_BINDTYPE_VM, floatingIpId, serverPort["serverId"], serverPort["portId"])
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DetachFloatingIp", func(t *testing.T) {
		floatingIpId = getFloatingIpId(fipAddress)

		err := net.DetachFloatingIp(floatingIpId)
		if err != nil {
			t.Fatal(err)
		}
	})
}
