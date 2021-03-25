package VirtualPrivateCloud

import (
	"errors"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func (a *APIv2) CreateNic(ns *global.NicSpec) (portId string, err error) {
	if ns.Name == "" {
		err = errors.New("No name is available")
		return
	}

	if ns.NetworkId == "" {
		err = errors.New("No networkId is available")
		return
	}

	if ns.SecurityGroups == nil || len(ns.SecurityGroups) == 0 {
		err = errors.New("No securityGroups is available")
		return
	}

	if len(ns.Subnets) == 0 {
		err = errors.New("No subnets is available")
		return
	}

	body := map[string]interface{}{
		"name":      ns.Name,
		"networkId": ns.NetworkId,
		"sgIds":     ns.SecurityGroups,
		"type":      "vm",
	}

	dvs := make([]map[string]string, len(ns.Subnets))

	for i, v := range ns.Subnets {
		dvs[i] = map[string]string{
			"ipAddress": v.IpAddress,
			"subnetId":  v.SubnetId,
		}
	}

	body["ips"] = dvs

	resp, err := a.client.NewRequest("POST", "/api/v2/netcenter/port", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	portId = resp.Body

	return
}

func (a *APIv2) DeleteNic(portId string) error {
	if portId == "" {
		return errors.New("No portId is available")
	}

	resp, err := a.client.NewRequest("DELETE", "/api/v2/netcenter/port/portId/"+portId, nil, nil, nil)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

func (a *APIv2) GetNicDetail(portId string) (result *NicResult, err error) {
	if portId == "" {
		err = errors.New("No portId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/netcenter/port/"+portId+"/PortDetailResp", nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, "newtag"); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) ModifyNicName(portId, portName string) error {
	if portId == "" {
		return errors.New("No portId is available")
	}

	if portName == "" {
		return errors.New("No portName is available")
	}

	body := map[string]interface{}{
		"id":   portId,
		"name": portName,
	}

	resp, err := a.client.NewRequest("PUT", "/api/v2/netcenter/port/portName", nil, nil, body)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

func (a *APIv2) ModifyNicSecurityGroup(portId string, securityGroupIds []string) error {
	if portId == "" {
		return errors.New("No portId is available")
	}

	if len(securityGroupIds) == 0 {
		return errors.New("No securityGroupIds is available")
	}

	body := map[string]interface{}{
		"id":               portId,
		"securityGroupIds": securityGroupIds,
	}

	resp, err := a.client.NewRequest("PUT", "/api/v2/netcenter/port/portSecurityGroups", nil, nil, body)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}
