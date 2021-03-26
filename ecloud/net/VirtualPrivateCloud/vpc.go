package VirtualPrivateCloud

import (
	"errors"
	"strings"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/IPSecVpn"
)

func (a *APIv2) CreateVpc(vs *global.VpcSpec) (result VpcOrderResult, err error) {
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
		"name":            vs.Name,
		"specs":           vs.Specs.String(),
		"networkName":     vs.NetworkName,
		"cidr":            vs.Cidr,
		"cidrV6":          vs.CidrV6,
		"region":          vs.Region,
		"networkTypeEnum": "VM",
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/netcenter/order/create/vpc", nil, nil, body)
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

func (a *APIv2) DeleteVpc(vpcId string) error {
	if vpcId == "" {
		return errors.New("No vpcId is available")
	}

	resp, err := a.client.NewRequest("DELETE", "/api/v2/netcenter/vpc/"+vpcId, nil, nil, nil)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

func (a *APIv2) GetVpcList(queryWord, region string, natGatewayBind bool, scale global.VpcScale, tagIds []string, page, size int) (result []VpcResult, err error) {
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
		params["tagIds"] = strings.Join(tagIds, ",")
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

	if queryWord != "" {
		for _, v := range result {
			if v.Name == queryWord {
				result = []VpcResult{v}
				break
			}
		}
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

func (a *APIv2) GetVpcNetwork(queryWord, routerId string, cidrInRange, networkIdsInRange []string, page, size int) (result []VpcNetResult, err error) {
	params := map[string]interface{}{}

	if queryWord != "" {
		params["queryWord"] = queryWord
	}

	if routerId != "" {
		params["routerId"] = routerId
	}

	if networkIdsInRange != nil && len(networkIdsInRange) > 0 {
		params["rangeInNetworkIds"] = strings.Join(networkIdsInRange, ",")
	}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
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

	if cidrInRange != nil && len(cidrInRange) > 0 {
		var obj []VpcNetResult
		for _, v := range result {
			for _, cidr := range cidrInRange {
				if len(v.Subnets) > 0 && v.Subnets[0].Cidr == cidr {
					obj = append(obj, v)
					break
				}
			}
		}

		result = obj
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
