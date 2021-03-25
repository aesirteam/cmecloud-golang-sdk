package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestELB(t *testing.T) {
	var (
		fipAddress   = "36.137.45.147"
		floatingIpId = ""
	)

	t.Run("CreateELB", func(t *testing.T) {
		floatingIpId = getFloatingIpId(fipAddress)
		subnetIds := getSubnetIds()

		if subnetIds == nil || len(subnetIds) == 0 {
			t.Fatal("No match subnetIds")
		}

		spec := global.ELBSpec{
			ChargePeriod:    global.BILLING_TYPE_HOUR,
			LoadBalanceName: elbName,
			SubnetId:        subnetIds[0],
			FloatingIpId:    floatingIpId,
		}

		result, err := net.CreateELB(&spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetELBList", func(t *testing.T) {
		result, err := net.GetELBList("", "", "", false, false, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetELBQuota", func(t *testing.T) {
		result, err := net.GetELBQuota()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("CreateListener", func(t *testing.T) {
		loadBalanceId = getLoadBalanceId()

		spec := global.ELBListenerSpec{
			LoadbalanceId: loadBalanceId,
			Name:          listenerName,
			Protocol:      global.ELB_PROTOCOL_HTTP,
			ProtocolPort:  1080,
			Algorithm:     global.ELB_ALGORITHM_ROUND_ROBIN,
			// SessionPersistence: global.ELB_PERSISTENCE_APP_COOKIE,
			// CookieName:         "JSESSION",
			HealthType: global.ELB_HEALTHCHECK_HTTP,
		}

		var err error
		listenerId, err = net.CreateELBListener(&spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("listenerId:%s\n", listenerId)
	})

	t.Run("GetListenerList", func(t *testing.T) {
		loadBalanceId = getLoadBalanceId()

		result, err := net.GetELBListenerList(loadBalanceId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetListenerInfo", func(t *testing.T) {
		listenerId, _ = getListenerId()

		result, err := net.GetELBListenerInfo(listenerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		poolId = result.PoolId
	})

	t.Run("AddMember", func(t *testing.T) {
		_, poolId = getListenerId()
		serverPort := getServerPort()

		if serverPort == nil {
			t.Fatal("No match server")
		}

		memberId, err := net.AddELBMember(&global.ELBMemberSpec{
			PoolId:   poolId,
			Ip:       serverPort["privateIp"],
			Port:     80,
			Weight:   100,
			Type:     1,
			VmHostId: serverPort["serverId"],
		})
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("memberId:%s\n", memberId)
	})

	t.Run("GetMemberList", func(t *testing.T) {
		_, poolId = getListenerId()

		result, err := net.GetELBMemberList(poolId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetMemberInfo", func(t *testing.T) {
		_, poolId = getListenerId()
		memberId = getMemberId()

		result, err := net.GetELBMemberInfo(poolId, memberId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("DeleteMember", func(t *testing.T) {
		_, poolId = getListenerId()
		memberId = getMemberId()

		err := net.DeleteELBMember(poolId, memberId)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteListener", func(t *testing.T) {
		listenerId, _ = getListenerId()
		err := net.DeleteELBListener(listenerId)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteELB", func(t *testing.T) {
		loadBalanceId = getLoadBalanceId()

		err := net.DeleteELB(loadBalanceId)
		if err != nil {
			t.Fatal(err)
		}
	})
}
