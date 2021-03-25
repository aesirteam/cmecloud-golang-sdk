package net

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

var cli = ecloud.NewForConfigDie(&global.Config{
	ApiGwHost: "api-guiyang-1.cmecloud.cn",
	//ApiGwPort:     8443,
	ApiGwProtocol: "https",
	// AccessKey:     "YOUR ACCESS KEY",
	// SecretKey:     "YOUR SECRET KEY",
})

var net, vm = cli.Net(), cli.VM()

var (
	serverName = "api198994873"

	vpcName                     = "vpc99999"
	subnetName                  = "subnet99999"
	subnetCidr                  = "10.18.0.0/24"
	portName                    = "port99999"
	vpcId, routerId             string
	networkId, subnetId, portId string

	vpnName                                    = "vpn99999"
	vpnSiteConnName                            = "vpn_conn99999"
	vpnId, siteConnId                          string
	siteConnIkePolicyId, siteConnIpsecPolicyId string

	floatingIpId string

	elbName                                     = "elb99999"
	listenerName                                = "elb_listener99999"
	loadBalanceId, listenerId, poolId, memberId string
)

var region = func() string {
	if result, err := vm.GetRegionList(); err == nil && len(result) > 0 {
		return result[0].Name
	}
	return ""
}()

var getRouterId = func() (string, string) {
	if vpcId == "" || routerId == "" {
		if result, err := net.GetVpcList(vpcName, region, false, 0, nil, 0, 0); err == nil && len(result) > 0 {
			vpc := &result[0]
			vpcId, routerId = vpc.Id, vpc.RouterId
		}
	}

	return vpcId, routerId
}

var getNetworkId = func() (string, string) {
	if networkId == "" || subnetId == "" {
		_, routerId = getRouterId()

		if result, err := net.GetVpcNetwork(subnetName, routerId, []string{subnetCidr}, nil, 0, 0); err == nil {
			networkId, subnetId = result[0].Subnets[0].NetworkId, result[0].Subnets[0].Id
		}
	}

	return networkId, subnetId
}

var getSubnetIds = func() []string {
	_, routerId = getRouterId()

	if result, err := net.GetRouterInfo(routerId); err == nil {
		return result.SubnetIds
	}

	return nil
}

var getPortId = func() string {
	if portId == "" {
		_, routerId = getRouterId()
		networkId, _ = getNetworkId()

		if result, err := net.GetVpcNic(routerId); err == nil {
			for _, v := range result {
				if v.NetworkId == networkId {
					portId = v.FixedIps[0].PortId
					break
				}
			}
		}

	}

	return portId
}

var getDefaultSecurityGroupIds = func() []string {
	if result, err := vm.GetSecurityGroupList("default", false, 0, 0); err == nil && len(result) > 0 {
		ids := make([]string, len(result))
		for i, v := range result {
			ids[i] = v.Id
		}
		return ids
	}

	return nil
}

var getVpnId = func() string {
	if vpnId == "" {
		_, routerId = getRouterId()

		if result, err := net.GetIpsecVpnList("", routerId, 0, "", 0, 0); err == nil && len(result) > 0 {
			vpnId = result[0].Id
		}
	}

	return vpnId
}

var getVpnSiteConnectionId = func() (string, string, string) {
	if siteConnId == "" || siteConnIkePolicyId == "" || siteConnIpsecPolicyId == "" {
		networkId, _ = getNetworkId()

		if connList, err := net.GetIpsecVpnConnectionList("", "", networkId, nil, 0, 0); err == nil {
			for _, v := range connList {
				if v.Name == vpnSiteConnName {
					siteConnId, siteConnIkePolicyId, siteConnIpsecPolicyId = v.Id, v.IkePolicyId, v.IpsecPolicyId
					break
				}
			}
		}
	}

	return siteConnId, siteConnIkePolicyId, siteConnIpsecPolicyId
}

var getFloatingIpId = func(ip string) string {
	if floatingIpId == "" {
		if result, err := net.GetFloatingIpList(ip, "", "", false, false, false, false, nil, 0, 0); err == nil && len(result) > 0 {
			floatingIpId = result[0].Id
		}
	}

	return floatingIpId
}

var getServerPort = func() map[string]string {
	networkId, _ = getNetworkId()

	if result, err := vm.GetServerList(serverName, networkId, region, 0, 0); err == nil && len(result) > 0 {
		if portDetails := result[0].ServerPortDetailArray; len(portDetails) > 0 {
			return map[string]string{
				"serverId":  result[0].ServerId,
				"portId":    portDetails[0].PortId,
				"privateIp": portDetails[0].PrivateIp,
			}
		}
	}

	return nil
}

var getLoadBalanceId = func() string {
	if loadBalanceId == "" {
		_, routerId = getRouterId()

		if result, err := net.GetELBList("", routerId, "", false, false, 0, 0); err == nil && len(result) > 0 {
			loadBalanceId = result[0].Id
		}
	}

	return loadBalanceId
}

var getListenerId = func() (string, string) {
	if listenerId == "" {
		loadBalanceId = getLoadBalanceId()

		if result, err := net.GetELBListenerList(loadBalanceId); err == nil {
			for _, v := range result {
				if v.Name == listenerName {
					listenerId, poolId = v.Id, v.PoolId
					break
				}
			}
		}
	}

	return listenerId, poolId
}

var getMemberId = func() string {
	if memberId == "" {
		_, poolId = getListenerId()

		if result, err := net.GetELBMemberList(poolId); err == nil && len(result) > 0 {
			memberId = result[0].Id
		}
	}

	return memberId
}
