package IPSecVpn

import (
	"errors"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func (a *APIv2) CreateIpsecVpn() {

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

func (a *APIv2) CreateIpsecVpnConnection() {

}

func (a *APIv2) DeleteIpsecVpnConnection() {

}

func (a *APIv2) GetIpsecVpnConnectionList() {

}

func (a *APIv2) GetIpsecVpnConnectionInfo() {

}

func (a *APIv2) GetIkePolicyInfo() {

}

func (a *APIv2) GetIpsecPolicyInfo() {

}

func (a *APIv2) GetIpsecVpnQuota() {

}
