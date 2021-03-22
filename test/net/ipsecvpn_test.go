package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

func TestVPN(t *testing.T) {
	net := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	}).Net()

	var (
		vpcName                                                                  = "vpc99999"
		vpnName                                                                  = "vpn99999"
		vpnSiteConnName                                                          = "vpn_conn_99999"
		vpnId                                                                    string
		siteConnectionId, siteConnectionIkePolicyId, siteConnectionIpsecPolicyId string
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

	getVpnId := func() string {
		if vpnId == "" {
			if result, err := net.GetIpsecVpnList("", vpc.RouterId, 0, "", 0, 0); err == nil && len(result) > 0 {
				vpnId = result[0].Id
			}
		}

		return vpnId
	}

	getVpnSiteConnectionId := func() (string, string, string) {
		if siteConnectionId == "" || siteConnectionIkePolicyId == "" || siteConnectionIpsecPolicyId == "" {
			if connList, err := net.GetIpsecVpnConnectionList("", "", vpc.FirstNetworkId, nil, 0, 0); err == nil {
				for _, v := range connList {
					if v.Name == vpnSiteConnName {
						siteConnectionId, siteConnectionIkePolicyId, siteConnectionIpsecPolicyId = v.Id, v.IkePolicyId, v.IpsecPolicyId
						break
					}
				}
			}
		}

		return siteConnectionId, siteConnectionIkePolicyId, siteConnectionIpsecPolicyId
	}

	t.Run("CreateIpsecVpn", func(t *testing.T) {
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

		spec := global.IPSecVpnSpec{
			ChargePeriod: global.BILLING_TYPE_HOUR,
			Name:         vpnName,
			RouterId:     vpc.RouterId,
			FloatingIpId: ipId,
		}

		result, err := net.CreateIpsecVpn(&spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetIpsecVpnList", func(t *testing.T) {
		result, err := net.GetIpsecVpnList("", "", 0, "", 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetIpsecVpnInfo", func(t *testing.T) {
		vpnId = getVpnId()

		result, err := net.GetIpsecVpnInfo(vpnId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetIpsecVpnQuota", func(t *testing.T) {
		result, err := net.GetIpsecVpnQuota()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("CreateConnection", func(t *testing.T) {
		vpnId = getVpnId()

		spec := global.IPSecVpnSiteConnectionSpec{
			VpnServiceId:   vpnId,
			Name:           vpnSiteConnName,
			LocalSubnetIds: subnetIds,
			PeerAddress:    "36.134.93.99",
			PeerSubnets:    []string{"192.168.0.0/24"},
			Psk:            "vpn123456789",
			IkePolicy: struct {
				AuthAlgorithm         global.VpnAuthAlgorithm
				EncryptionAlgorithm   global.VpnEncryptionAlgorithm
				Pfs                   global.VpnPfs
				Version               global.VpnIkeVersion
				Lifetime              int
				Phase1NegotiationMode global.VpnIkePhase1NegotiationMode
			}{
				global.VPN_AUTHALGORITHM_SHA1,
				global.VPN_ENCRYPTIONALGORITHM_AES128,
				global.VPN_PFS_GROUP2,
				global.VPN_IKE_VERSION_V2,
				3600,
				global.VPN_IKE_PHASE1NEGOTIATIONMODE_MAIN,
			},
			IpsecPolicy: struct {
				AuthAlgorithm       global.VpnAuthAlgorithm
				EncryptionAlgorithm global.VpnEncryptionAlgorithm
				EncapsulationMode   global.VpnEncapsulationMode
				Pfs                 global.VpnPfs
				Lifetime            int
			}{
				global.VPN_AUTHALGORITHM_SHA1,
				global.VPN_ENCRYPTIONALGORITHM_AES128,
				global.VPN_ENCAPSULATIONMODE_TUNNEL,
				global.VPN_PFS_GROUP2,
				3600,
			},
		}

		err := net.CreateIpsecVpnConnection(&spec)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetConnectionList", func(t *testing.T) {
		result, err := net.GetIpsecVpnConnectionList("", "", "", nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetConnectionInfo", func(t *testing.T) {
		siteConnectionId, siteConnectionIkePolicyId, siteConnectionIpsecPolicyId = getVpnSiteConnectionId()

		result, err := net.GetIpsecVpnConnectionInfo(siteConnectionId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetIkePolicyInfo", func(t *testing.T) {
		siteConnectionId, siteConnectionIkePolicyId, siteConnectionIpsecPolicyId = getVpnSiteConnectionId()

		result, err := net.GetIkePolicyInfo(siteConnectionId, siteConnectionIkePolicyId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetIpsecPolicyInfo", func(t *testing.T) {
		siteConnectionId, siteConnectionIkePolicyId, siteConnectionIpsecPolicyId = getVpnSiteConnectionId()

		result, err := net.GetIpsecPolicyInfo(siteConnectionId, siteConnectionIpsecPolicyId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifyConnection", func(t *testing.T) {
		siteConnectionId, _, _ = getVpnSiteConnectionId()

		spec := global.IPSecVpnSiteConnectionSpec{
			Name:           vpnSiteConnName,
			LocalSubnetIds: subnetIds,
			PeerAddress:    "36.134.93.99",
			PeerSubnets:    []string{"192.168.0.0/24"},
			Psk:            "vpn123456789",
		}

		err := net.ModifyIpsecVpnConnection(siteConnectionId, &spec)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteConnection", func(t *testing.T) {
		siteConnectionId, _, _ = getVpnSiteConnectionId()

		err := net.DeleteIpsecVpnConnection(siteConnectionId)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteIpsecVpn", func(t *testing.T) {
		vpnId = getVpnId()

		err := net.DeleteIpsecVpn(vpnId)
		if err != nil {
			t.Fatal(err)
		}
	})

}
