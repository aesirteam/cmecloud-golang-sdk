package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

func TestELB(t *testing.T) {
	cli := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	})

	net := cli.Net()

	var (
		vpcName      = "vpc99999"
		elbName      = "elb99999"
		listenerName = "listener99999"

		loadBalanceId, listenerId, poolId string
	)

	vpc := func() VirtualPrivateCloud.VpcResult {
		vpc, err := net.GetVpcInfoByName(vpcName)
		if err != nil {
			t.Fatal(err)
		}

		return *vpc
	}()

	subnetIds := func() []string {
		result, err := net.GetSubnetList(vpc.FirstNetworkId)
		if err != nil {
			t.Fatal(err)
		}

		subnetIds := make([]string, len(result))

		for i, v := range result {
			subnetIds[i] = v.Id
		}

		return subnetIds
	}()

	getLoadBalanceId := func() string {
		if loadBalanceId == "" {
			if result, err := net.GetELBList(elbName, "", "", false, false, 0, 0); err == nil && len(result) > 0 {
				loadBalanceId = result[0].Id
			}
		}
		return loadBalanceId
	}

	getListenerId := func() string {
		if listenerId == "" {
			loadBalanceId = getLoadBalanceId()
			if result, err := net.GetELBListenerList(loadBalanceId); err == nil {
				for _, v := range result {
					if v.Name == listenerName {
						listenerId = v.Id
						break
					}
				}
			}
		}

		return listenerId
	}

	getPoolId := func() string {
		if poolId == "" {
			listenerId = getListenerId()
			if result, err := net.GetELBListenerInfo(listenerId); err == nil {
				poolId = result.PoolId
			}
		}

		return poolId
	}

	t.Run("CreateELB", func(t *testing.T) {
		var (
			err  error
			ipId string
		)

		fips, err := net.GetFloatingIpList("", "", "", false, false, false, false, nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		for _, v := range fips {
			if !v.Bound && v.Status == global.FLOATINGIP_BINDSTATUS_UNBOUND.String() {
				ipId = v.Id
				break
			}
		}

		spec := global.ELBSpec{
			ChargePeriod:    global.BILLING_TYPE_HOUR,
			LoadBalanceName: elbName,
			SubnetId:        subnetIds[0],
			FloatingIpId:    ipId,
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
			ProtocolPort:  1443,
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
		listenerId = getListenerId()

		result, err := net.GetELBListenerInfo(listenerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))

		poolId = result.PoolId
	})

	t.Run("", func(t *testing.T) {
		poolId = getPoolId()
	})

	t.Run("DeleteListener", func(t *testing.T) {
		listenerId = getListenerId()

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
