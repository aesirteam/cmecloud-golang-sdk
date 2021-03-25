package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestVPN(t *testing.T) {
	var (
		fipAddress   = "36.137.45.147"
		floatingIpId = ""
	)

	t.Run("CreateVpn", func(t *testing.T) {
		_, routerId = getRouterId()
		floatingIpId = getFloatingIpId(fipAddress)

		spec := global.IPSecVpnSpec{
			ChargePeriod: global.BILLING_TYPE_HOUR,
			Name:         vpnName,
			RouterId:     routerId,
			FloatingIpId: floatingIpId,
		}

		result, err := net.CreateIpsecVpn(&spec)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpnList", func(t *testing.T) {
		result, err := net.GetIpsecVpnList("", "", 0, "", 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpnInfo", func(t *testing.T) {
		vpnId = getVpnId()

		result, err := net.GetIpsecVpnInfo(vpnId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetVpnQuota", func(t *testing.T) {
		result, err := net.GetIpsecVpnQuota()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("CreateConnection", func(t *testing.T) {
		vpnId = getVpnId()
		subnetIds := getSubnetIds()

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
		siteConnId, siteConnIkePolicyId, siteConnIpsecPolicyId = getVpnSiteConnectionId()

		result, err := net.GetIpsecVpnConnectionInfo(siteConnId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetIkePolicyInfo", func(t *testing.T) {
		siteConnId, siteConnIkePolicyId, siteConnIpsecPolicyId = getVpnSiteConnectionId()

		result, err := net.GetIkePolicyInfo(siteConnId, siteConnIkePolicyId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetIpsecPolicyInfo", func(t *testing.T) {
		siteConnId, siteConnIkePolicyId, siteConnIpsecPolicyId = getVpnSiteConnectionId()

		result, err := net.GetIpsecPolicyInfo(siteConnId, siteConnIpsecPolicyId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("ModifyConnection", func(t *testing.T) {
		siteConnId, _, _ = getVpnSiteConnectionId()
		subnetIds := getSubnetIds()

		spec := global.IPSecVpnSiteConnectionSpec{
			Name:           vpnSiteConnName,
			LocalSubnetIds: subnetIds,
			PeerAddress:    "36.134.93.99",
			PeerSubnets:    []string{"192.168.0.0/24"},
			Psk:            "vpn123456789",
		}

		err := net.ModifyIpsecVpnConnection(siteConnId, &spec)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteConnection", func(t *testing.T) {
		siteConnId, _, _ = getVpnSiteConnectionId()

		err := net.DeleteIpsecVpnConnection(siteConnId)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteVpn", func(t *testing.T) {
		vpnId = getVpnId()

		err := net.DeleteIpsecVpn(vpnId)
		if err != nil {
			t.Fatal(err)
		}
	})

}
