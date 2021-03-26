package vm

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

var vm, net, storage = cli.VM(), cli.Net(), cli.Storage()

var (
	vpcName                     = "vpc99999"
	subnetName                  = "subnet99999"
	subnetCidr                  = "10.18.0.0/24"
	portName                    = "port99999"
	vpcId, routerId             string
	networkId, subnetId, portId string

	cbsVolumeName = "BS99999"
	cbsVolumeId   string

	keypairName = "kp99999"
	keypairId   string

	securityGroupName                    = "sg99999"
	securityGroupPortRange               = []int{22, 22}
	securityGroupId, securityGroupRuleId string

	serverName            = "api198994873"
	serverId, userImageId string

	serverBackupName = "backup_" + serverName
	serverBackupId   string
)

var region = func() string {
	if result, err := vm.GetRegionList(""); err == nil && len(result) > 0 {
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

var getVolumeId = func() string {
	if cbsVolumeId == "" {
		if result, err := storage.GetVolumeList(cbsVolumeName, "", false, nil, 0, 0); err == nil && len(result) > 0 {
			cbsVolumeId = result[0].Id
		}
	}

	return cbsVolumeId
}

var getKeypairId = func() string {
	if keypairId == "" {
		if result, err := vm.GetKeypairList(keypairName, region, 0, 0); err == nil && len(result) > 0 {
			keypairId = result[0].Id
		}
	}

	return keypairId
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

var getSecurityGroupId = func() string {
	if securityGroupId == "" {
		if result, err := vm.GetSecurityGroupList(securityGroupName, false, 0, 0); err == nil && len(result) > 0 {
			securityGroupId = result[0].Id
		}
	}

	return securityGroupId
}

var getSecurityGroupRuleId = func() string {
	if securityGroupRuleId == "" {
		securityGroupId = getSecurityGroupId()

		if result, err := vm.GetSecurityGroupRules(securityGroupId, ""); err == nil {
			for _, rule := range result {
				if rule.PortRangeMin == securityGroupPortRange[0] && rule.PortRangeMax == securityGroupPortRange[1] {
					securityGroupRuleId = rule.Id
					break
				}
			}
		}
	}

	return securityGroupRuleId
}

var getServerId = func() string {
	if serverId == "" {
		if result, err := vm.GetServerList(serverName, "", region, 0, 0); err == nil && len(result) > 0 {
			serverId = result[0].ServerId
		}
	}

	return serverId
}

var getUserImageId = func(name string) string {
	if userImageId == "" {
		if result, err := vm.GetUserImageList("", "", name, nil, nil, 0, 0); err == nil && len(result) > 0 {
			userImageId = result[0].Id
		}
	}

	return userImageId
}

var getServerBackupId = func() string {
	if serverBackupId == "" {
		serverId = getServerId()

		if result, err := vm.GetServerBackupList(serverId, serverBackupName, 0, 0); err == nil && len(result) > 0 {
			serverBackupId = result[0].Id
		}
	}

	return serverBackupId
}
