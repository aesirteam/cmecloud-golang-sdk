package VirtualPrivateCloud

import (
	"errors"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/IPSecVpn"
)

func (a *APIv2) CreateVpc(vs *global.VpcSpec) (vpcId string, err error) {
	if vs.Cidr == "" {
		err = errors.New("No cidr is available")
		return
	}

	if vs.Name == "" {
		err = errors.New("No name is available")
		return
	}

	if vs.NetworkName == "" {
		err = errors.New("No networkName is available")
		return
	}

	if vs.Region == "" {
		err = errors.New("No region is available")
		return
	}

	body := map[string]interface{}{
		"cidr":            vs.Cidr,
		"cidrV6":          vs.CidrV6,
		"name":            vs.Name,
		"networkName":     vs.NetworkName,
		"region":          vs.Region,
		"networkTypeEnum": "VM",
		"specs":           vs.Specs.String(),
	}

	//fmt.Println(core.Dump(body))

	resp, err := a.client.NewRequest("POST", "/api/v2/netcenter/vpc", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	vpcId = resp.Body

	return
}

func (a *APIv2) DeleteVpc(vpcId string) {

}

func (a *APIv2) GetVpcList(natGatewayBind bool, scale global.VpcScale, region string, tagIds []string, page, size int) (result []VpcResult, err error) {
	params := map[string]interface{}{
		"scale":   scale.String(),
		"visible": true,
	}

	if natGatewayBind {
		params["natGatewayBind"] = true
	}

	if region != "" {
		params["region"] = region
	}

	if tagIds != nil && len(tagIds) > 0 {
		params["tagIds"] = tagIds
	}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/vpc", nil, params, nil)
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

func (a *APIv2) GetVpcInfo(vpcId string) (result VpcResult, err error) {
	if vpcId == "" {
		err = errors.New("No vpcId is available")
		return
	}

	params := map[string]interface{}{
		"visible": true,
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/vpc/"+vpcId, nil, params, nil)
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

func (a *APIv2) GetVpcInfoByName(name string) (result VpcResult, err error) {
	arr, err := a.GetVpcList(false, 0, "", nil, 0, 0)
	if err != nil {
		return
	}

	for _, vpc := range arr {
		if vpc.Name == name {
			return a.GetVpcInfo(vpc.Id)
		}
	}

	err = errors.New("No match vpc info")
	return
}

func (a *APIv2) GetVpcInfoByRouterId(routerId string) (result VpcResult, err error) {
	if routerId == "" {
		err = errors.New("No routerId is available")
		return
	}

	params := map[string]interface{}{
		"visible": true,
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/vpc/router/"+routerId, nil, params, nil)
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

func (a *APIv2) ModifyVpcInfo(vpcId, name, desc string) error {
	if vpcId == "" {
		return errors.New("No vpcId is available")
	}

	if name == "" {
		return errors.New("No name is available")
	}

	body := map[string]interface{}{
		"id":   vpcId,
		"name": name,
	}

	if desc != "" {
		body["description"] = desc
	}

	resp, err := a.client.NewRequest("PUT", "/api/v2/vpc", nil, nil, body)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

func (a *APIv2) GetVpcFirewall(routerId string) (result string, err error) {
	if routerId == "" {
		err = errors.New("No routerId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/vpc/"+routerId+"/firewall", nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	result = resp.Body

	return
}

func (a *APIv2) GetVpcNetwork(routerId string) (result []VpcNetResult, err error) {
	if routerId == "" {
		err = errors.New("No routerId is available")
		return
	}

	params := map[string]interface{}{
		"routerId": routerId,
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/network/NetworkResps", nil, params, nil)
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

func (a *APIv2) GetVpcVPN(routerId string) ([]IPSecVpn.VpnResult, error) {
	if routerId == "" {
		return nil, errors.New("No routerId is available")
	}

	vpn := IPSecVpn.New(a.client)

	return vpn.GetIpsecVpnList("", routerId, 0, "", 0, 0)
}

func (a *APIv2) GetVpcNic(routerId string) (result []NicResult, err error) {
	if routerId == "" {
		err = errors.New("No routerId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/vpc/"+routerId+"/nic", nil, nil, nil)
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
