package IPSecVpn

import (
	"errors"
	"fmt"
	"strings"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func (a *APIv2) CreateIpsecVpn(vs *global.IPSecVpnSpec) (vpnId string, err error) {

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
		"name":         vs.Name,
		"routerId":     vs.RouterId,
		"floatingipId": vs.FloatingIpId,
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/netcenter/IPSecVpn/service", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	vpnId = resp.Body

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

func (a *APIv2) ModifyIpsecVpn() {

}

func (a *APIv2) DeleteIpsecVpn(vpnId string) error {
	if vpnId == "" {
		return errors.New("No vpnId is available")
	}

	resp, err := a.client.NewRequest("DELETE", "/api/v2/netcenter/IPSecVpn/service/"+vpnId, nil, nil, nil)
	if err != nil {
		return resp.Error(err)
	}

	fmt.Println(resp.Body)

	return nil
}

func (a *APIv2) CreateIpsecVpnConnection(cs *global.IPSecVpnSiteConnectionSpec) (err error) {
	if cs.ServiceId == "" {
		err = errors.New("No serviceId is available")
		return
	}

	if cs.Name == "" {
		err = errors.New("No name is available")
		return
	}

	if cs.LocalSubnetId == "" {
		err = errors.New("No localSubnetId is available")
		return
	}

	if cs.PeerAddress == "" {
		err = errors.New("No peerAddress is available")
		return
	}

	if cs.PeerSubnetCidr == nil || len(cs.PeerSubnetCidr) == 0 {
		err = errors.New("No peerSubnetCidr is available")
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
		"serviceId":    cs.ServiceId,
	}

	endpointGroups := make([]map[string]interface{}, 2)

	endpointGroups[0] = map[string]interface{}{
		"endpoints": []string{cs.LocalSubnetId},
		"type":      "subnet",
	}

	endpointGroups[1] = map[string]interface{}{
		"endpoints": cs.PeerSubnetCidr,
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

func (a *APIv2) DeleteIpsecVpnConnection() {

}

func (a *APIv2) GetIpsecVpnConnectionList(queryWord, name, networkId string, serviceIdsInRange []string, page, size int) (result string, err error) {
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

	result = resp.Body

	return
}

func (a *APIv2) GetIpsecVpnConnectionInfo() {

}

func (a *APIv2) GetIkePolicyInfo() {

}

func (a *APIv2) GetIpsecPolicyInfo() {

}

func (a *APIv2) GetIpsecVpnQuota() {

}
