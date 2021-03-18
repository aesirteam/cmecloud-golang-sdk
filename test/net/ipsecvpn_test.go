package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/IPSecVpn"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

func TestVPN(t *testing.T) {
	net := ecloud.NewForConfigDie(&global.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	}).Net()

	var (
		vpcName         = "vpc99999"
		vpnName         = "vpn99999"
		vpnSiteConnName = "vpn_conn_99999"
		vpnId           string
	)

	vpc := func() VirtualPrivateCloud.VpcResult {
		vpc, err := net.GetVpcInfoByName(vpcName)
		if err != nil {
			t.Fatal(err)
		}

		return vpc
	}()

	subnetIds := func() []string {
		result, err := net.GetSubnetInfo(vpc.NetworkId)
		if err != nil {
			t.Fatal(err)
		}

		subnetIds := make([]string, len(result))

		for i, v := range result {
			subnetIds[i] = v.Id
		}

		return subnetIds
	}()

	getFloatingIpId := func() string {
		result, err := net.GetFloatingIpList("", "", "", false, false, false, false, nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		for _, v := range result {
			if !v.Bound && v.Status == global.FLOATINGIP_BINDSTATUS_UNBOUND.String() {
				return v.Id
			}
		}

		return ""
	}

	getVpnId := func() string {
		result, err := net.GetIpsecVpnList("", vpc.RouterId, 0, "", 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		return result[0].Id
	}

	vpnSiteConnection := func() *IPSecVpn.VpnConnectionResult {
		connList, err := net.GetIpsecVpnConnectionList("", "", "", nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		for _, v := range connList {
			if v.Name == vpnSiteConnName {
				return &v
			}
		}

		return nil
	}

	t.Run("CreateIpsecVpn", func(t *testing.T) {
		ipId := getFloatingIpId()

		spec := global.IPSecVpnSpec{
			Name:         vpnName,
			RouterId:     vpc.RouterId,
			FloatingIpId: ipId,
		}

		var err error
		vpnId, err = net.CreateIpsecVpn(&spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Logf("vpnId:%s\n", vpnId)
	})

	t.Run("GetIpsecVpnList", func(t *testing.T) {
		result, err := net.GetIpsecVpnList("", "", 0, "", 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetIpsecVpnInfo", func(t *testing.T) {
		if vpnId == "" {
			vpnId = getVpnId()
		}

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

	t.Run("CreateIpsecVpnConnection", func(t *testing.T) {
		if vpnId == "" {
			vpnId = getVpnId()
		}

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

	t.Run("GetIpsecVpnConnectionList", func(t *testing.T) {
		result, err := net.GetIpsecVpnConnectionList("", "", "", nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetIpsecVpnConnectionInfo", func(t *testing.T) {
		siteConnection := vpnSiteConnection()
		if siteConnection == nil {
			t.Fatal("No match siteConnection")
		}

		result, err := net.GetIpsecVpnConnectionInfo(siteConnection.Id)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetIkePolicyInfo", func(t *testing.T) {
		siteConnection := vpnSiteConnection()
		if siteConnection == nil {
			t.Fatal("No match siteConnection")
		}

		result, err := net.GetIkePolicyInfo(siteConnection.Id, siteConnection.IkePolicyId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetIpsecPolicyInfo", func(t *testing.T) {
		siteConnection := vpnSiteConnection()
		if siteConnection == nil {
			t.Fatal("No match siteConnection")
		}

		result, err := net.GetIpsecPolicyInfo(siteConnection.Id, siteConnection.IpsecPolicyId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifyIpsecVpnConnection", func(t *testing.T) {
		siteConnection := vpnSiteConnection()
		if siteConnection == nil {
			t.Fatal("No match siteConnection")
		}

		spec := global.IPSecVpnSiteConnectionSpec{
			Name:           vpnSiteConnName,
			LocalSubnetIds: subnetIds,
			PeerAddress:    "36.134.93.99",
			PeerSubnets:    []string{"192.168.0.0/24"},
			Psk:            "vpn123456789",
		}

		err := net.ModifyIpsecVpnConnection(siteConnection.Id, &spec)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteIpsecVpnConnection", func(t *testing.T) {
		siteConnection := vpnSiteConnection()
		if siteConnection == nil {
			t.Fatal("No match siteConnection")
		}

		err := net.DeleteIpsecVpnConnection(siteConnection.Id)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteIpsecVpn", func(t *testing.T) {
		if vpnId == "" {
			vpnId = getVpnId()
		}

		err := net.DeleteIpsecVpn(vpnId)
		if err != nil {
			t.Fatal(err)
		}
	})

}
