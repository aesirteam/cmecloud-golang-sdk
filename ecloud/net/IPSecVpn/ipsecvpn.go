package IPSecVpn

import (
	"errors"
	"strings"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func (a *APIv2) CreateIpsecVpn(vs *global.IPSecVpnSpec) (result VpnOrderResult, err error) {

	if vs.Name == "" {
		err = errors.New("No name is available")
		return
	}

	if vs.RouterId == "" {
		err = errors.New("No routerId is available")
		return
	}

	if vs.FloatingIpId == "" {
		err = errors.New("No floatingIpId is available")
		return
	}

	body := map[string]interface{}{
		"chargePeriodEnum": strings.ToLower(vs.ChargePeriod.String()),
		"quantity":         1,
		"name":             vs.Name,
		"routerId":         vs.RouterId,
		"floatingipId":     vs.FloatingIpId,
	}

	switch vs.ChargePeriod {
	case global.BILLING_TYPE_YEAR:
		if vs.Duration == 0 {
			body["duration"] = 12
		} else if vs.Duration > 0 && vs.Duration <= 5*12 {
			body["duration"] = vs.Duration
		}
	case global.BILLING_TYPE_MONTH:
		if vs.Duration == 0 {
			body["duration"] = 1
		} else if vs.Duration > 0 && vs.Duration <= 12 {
			body["duration"] = vs.Duration
		}
	}

	if vs.Quantity > 0 {
		body["quantity"] = vs.Quantity
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/netcenter/order/create/ipsecvpn", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) GetIpsecVpnList(vpcName, routerId string, scale global.VpcScale, region string, page, size int) (result []VpnResult, err error) {
	params := map[string]interface{}{
		"vpcScale": scale.String(),
		"visible":  true,
	}

	if vpcName != "" {
		params["vpcName"] = vpcName
	}

	if routerId != "" {
		params["routerId"] = routerId
	}

	if region != "" {
		params["region"] = region
	}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/IPSecVpn/service/IPSecVpnServiceResps", nil, params, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) GetIpsecVpnInfo(vpnId string) (result VpnResult, err error) {
	if vpnId == "" {
		err = errors.New("No vpnId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/IPSecVpn/service/"+vpnId+"/IPSecVpnServiceDetailResp", nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

// func (a *APIv2) ModifyIpsecVpn() {

// }

func (a *APIv2) GetIpsecVpnQuota() (result VpnQuotaResult, err error) {
	resp, err := a.client.NewRequest("GET", "/api/vpn/vpnService/quota", nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) DeleteIpsecVpn(vpnId string) error {
	if vpnId == "" {
		return errors.New("No vpnId is available")
	}

	resp, err := a.client.NewRequest("DELETE", "/api/v2/netcenter/IPSecVpn/service/"+vpnId, nil, nil, nil)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

func (a *APIv2) CreateIpsecVpnConnection(cs *global.IPSecVpnSiteConnectionSpec) (err error) {
	if cs.VpnServiceId == "" {
		err = errors.New("No vpnServiceId is available")
		return
	}

	if cs.Name == "" {
		err = errors.New("No name is available")
		return
	}

	if cs.LocalSubnetIds == nil || len(cs.LocalSubnetIds) == 0 {
		err = errors.New("No localSubnetIds is available")
		return
	}

	if cs.PeerAddress == "" {
		err = errors.New("No peerAddress is available")
		return
	}

	if cs.PeerSubnets == nil || len(cs.PeerSubnets) == 0 {
		err = errors.New("No peerSubnets is available")
		return
	}

	if cs.Psk == "" {
		err = errors.New("No psk is available")
		return
	}

	if cs.IkePolicy.Lifetime == 0 {
		cs.IkePolicy.Lifetime = 3600
	}

	if cs.IpsecPolicy.Lifetime == 0 {
		cs.IpsecPolicy.Lifetime = 3600
	}

	body := map[string]interface{}{
		"bandwithSize": 0,
		"localId":      "",
		"name":         cs.Name,
		"peerAddress":  cs.PeerAddress,
		"peerId":       "",
		"psk":          cs.Psk,
		"serviceId":    cs.VpnServiceId,
	}

	endpointGroups := make([]map[string]interface{}, 2)

	endpointGroups[0] = map[string]interface{}{
		"endpoints": cs.LocalSubnetIds,
		"type":      "subnet",
	}

	endpointGroups[1] = map[string]interface{}{
		"endpoints": cs.PeerSubnets,
		"type":      "cidr",
	}

	body["endpointGroups"] = endpointGroups

	body["ikePolicy"] = map[string]interface{}{
		"authAlgorithm":       cs.IkePolicy.AuthAlgorithm.String(),
		"encryptionAlgorithm": cs.IkePolicy.EncryptionAlgorithm.String(),
		"lifetime":            cs.IkePolicy.Lifetime,
		"pfs":                 cs.IkePolicy.Pfs.String(),
		"phase1NegotiationMode": func() string {
			if cs.IkePolicy.Version == global.VPN_IKE_VERSION_V2 {
				return global.VPN_IKE_PHASE1NEGOTIATIONMODE_MAIN.String()
			}
			return cs.IkePolicy.Phase1NegotiationMode.String()
		}(),
		"version": cs.IkePolicy.Version.String(),
	}

	body["ipsecPolicy"] = map[string]interface{}{
		"authAlgorithm":       cs.IpsecPolicy.AuthAlgorithm.String(),
		"encryptionAlgorithm": cs.IpsecPolicy.EncryptionAlgorithm.String(),
		"encapsulationMode":   cs.IpsecPolicy.EncapsulationMode.String(),
		"lifetime":            cs.IpsecPolicy.Lifetime,
		"pfs":                 cs.IpsecPolicy.Pfs.String(),
		"transformProtocol":   "ESP",
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/network/IPSecVpn/siteConnection", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) GetIpsecVpnConnectionList(queryWord, name, networkId string, serviceIdsInRange []string, page, size int) (result []VpnConnectionResult, err error) {
	params := map[string]interface{}{}

	if queryWord != "" {
		params["queryWord"] = queryWord
	}

	if name != "" {
		params["name"] = name
	}

	if networkId != "" {
		params["networkId"] = networkId
	}

	if serviceIdsInRange != nil && len(serviceIdsInRange) > 0 {
		params["serviceIdsInRange"] = strings.Join(serviceIdsInRange, ",")
	}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/network/IPSecVpn/siteConnection/IPSecVpnSiteConnectionResps", nil, params, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) GetIpsecVpnConnectionInfo(siteConnectionId string) (result VpnConnectionResult, err error) {
	if siteConnectionId == "" {
		err = errors.New("No siteConnectionId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/network/IPSecVpn/siteConnection/"+siteConnectionId+"/IPSecVpnSiteConnectionResp", nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) GetIkePolicyInfo(siteConnectionId, ikePolicyId string) (result VpnConnectionPolicyResult, err error) {
	if siteConnectionId == "" {
		err = errors.New("No siteConnectionId is available")
		return
	}

	if ikePolicyId == "" {
		err = errors.New("No ikePolicyId is available")
		return
	}

	params := map[string]interface{}{
		"siteConnectionId": siteConnectionId,
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/IPSecVpn/IkePolicy/"+ikePolicyId, nil, params, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) GetIpsecPolicyInfo(siteConnectionId, ipsecPolicyId string) (result VpnConnectionPolicyResult, err error) {
	if siteConnectionId == "" {
		err = errors.New("No siteConnectionId is available")
		return
	}

	if ipsecPolicyId == "" {
		err = errors.New("No ipsecPolicyId is available")
		return
	}

	params := map[string]interface{}{
		"siteConnectionId": siteConnectionId,
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/IPSecVpn/IPSecPolicy/"+ipsecPolicyId, nil, params, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) ModifyIpsecVpnConnection(siteConnectionId string, cs *global.IPSecVpnSiteConnectionSpec) (err error) {
	if siteConnectionId == "" {
		err = errors.New("No siteConnectionId is available")
		return
	}

	if cs.Name == "" {
		err = errors.New("No name is available")
		return
	}

	if cs.LocalSubnetIds == nil || len(cs.LocalSubnetIds) == 0 {
		err = errors.New("No localSubnetIds is available")
		return
	}

	if cs.PeerAddress == "" {
		err = errors.New("No peerAddress is available")
		return
	}

	if cs.PeerSubnets == nil || len(cs.PeerSubnets) == 0 {
		err = errors.New("No peerSubnets is available")
		return
	}

	if cs.Psk == "" {
		err = errors.New("No psk is available")
		return
	}

	body := map[string]interface{}{
		"bandwithSize": 0,
		"localId":      "",
		"name":         cs.Name,
		"peerAddress":  cs.PeerAddress,
		"peerId":       "",
		"psk":          cs.Psk,
	}

	endpointGroups := make([]map[string]interface{}, 2)

	endpointGroups[0] = map[string]interface{}{
		"endpoints": cs.LocalSubnetIds,
		"type":      "subnet",
	}

	endpointGroups[1] = map[string]interface{}{
		"endpoints": cs.PeerSubnets,
		"type":      "cidr",
	}

	body["endpointGroups"] = endpointGroups

	resp, err := a.client.NewRequest("PUT", "/api/v2/network/IPSecVpn/siteConnection/"+siteConnectionId, nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) DeleteIpsecVpnConnection(siteConnectionId string) (err error) {
	if siteConnectionId == "" {
		err = errors.New("No siteConnectionId is available")
		return
	}

	resp, err := a.client.NewRequest("DELETE", "/api/v2/network/IPSecVpn/siteConnection/"+siteConnectionId, nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	return
}
