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
		vpcName = "vpc99999"
		vpnName = "vpn99999"
		vpnId   string
	)

	vpc := func() VirtualPrivateCloud.VpcResult {
		vpc, err := net.GetVpcInfoByName(vpcName)
		if err != nil {
			t.Fatal(err)
		}

		return vpc
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

	t.Run("CreateIpsecVpnConnection", func(t *testing.T) {
		if vpnId == "" {
			vpnId = getVpnId()
		}

		subnet, err := net.GetSubnetInfo(vpc.NetworkId)
		if err != nil {
			t.Fatal(err)
		}

		spec := global.IPSecVpnSiteConnectionSpec{
			Name:           "vpn_conn_99999",
			LocalSubnetId:  subnet[0].Id,
			PeerAddress:    "117.187.201.180",
			PeerSubnetCidr: []string{"10.204.80.0/24", "10.204.62.0/24"},
			ServiceId:      vpnId,
			Psk:            "vpn12345",
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
				global.VPN_IKE_PHASE1NEGOTIATIONMODE_AGGRESSIVE,
			},
			IpsecPolicy: struct {
				AuthAlgorithm       global.VpnAuthAlgorithm
				EncryptionAlgorithm global.VpnEncryptionAlgorithm
				EncapsulationMode   global.VpnEncapsulationMode
				Pfs                 global.VpnPfs
				Lifetime            int
			}{
				global.VPN_AUTHALGORITHM_SHA256,
				global.VPN_ENCRYPTIONALGORITHM_DES3,
				global.VPN_ENCAPSULATIONMODE_TUNNEL,
				global.VPN_PFS_GROUP2,
				3600,
			},
		}

		err = net.CreateIpsecVpnConnection(&spec)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("GetIpsecVpnConnectionList", func(t *testing.T) {
		result, err := net.GetIpsecVpnConnectionList("", "", "", nil, 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result)
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
