package SecurityGroup

import (
	"errors"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func (a *APIv2) CreateSecurityGroup(name, desc string) (result SecurityGroupResult, err error) {
	if name == "" {
		err = errors.New("No name is available")
		return
	}

	body := map[string]interface{}{
		"name": name,
	}

	if desc != "" {
		body["description"] = desc
	}

	resp, err := a.client.NewRequest("POST", "/api/securitygroup", nil, nil, body)
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

func (a *APIv2) DeleteSecurityGroup(id string) error {
	if id == "" {
		return errors.New("No id is available")
	}

	resp, err := a.client.NewRequest("DELETE", "/api/securitygroup/"+id, nil, nil, nil)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

func (a *APIv2) GetSecurityGroupList(name string, ruleShow bool, page, size int) (result []SecurityGroupResult, err error) {
	params := map[string]interface{}{
		"types": "VM",
	}

	if name != "" {
		params["name"] = name
	}

	if ruleShow {
		params["ruleShow"] = true
	}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	resp, err := a.client.NewRequest("GET", "/api/securitygroup", nil, params, nil)
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

func (a *APIv2) ModifySecurityGroup(id, name, desc string) (result SecurityGroupResult, err error) {
	if id == "" {
		err = errors.New("No networkId is available")
		return
	}

	if name == "" {
		err = errors.New("No name is available")
		return
	}

	body := map[string]interface{}{
		"id":   id,
		"name": name,
	}

	if desc != "" {
		body["description"] = desc
	}

	resp, err := a.client.NewRequest("PUT", "/api/securitygroup", nil, nil, body)
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

func (a *APIv2) AddSecurityGroupRules(sg *global.SecurityGroupRuleSpec) (result SecurityGroupRuleResult, err error) {
	if sg.SecurityGroupId == "" {
		err = errors.New("No securityGroupId is available")
		return
	}

	if sg.MinPortRange == 0 && sg.MaxPortRange == 0 {
		err = errors.New("No minPortRange and maxPortRange is available")
		return
	} else if sg.MinPortRange == 0 && sg.MaxPortRange > 0 {
		if sg.MaxPortRange > 65535 {
			err = errors.New("MaxPortRange less than 65535")
			return
		}
		sg.MinPortRange = sg.MaxPortRange
	} else if sg.MaxPortRange == 0 && sg.MinPortRange > 0 {
		if sg.MinPortRange > 65535 {
			err = errors.New("MinPortRange less than 65535")
			return
		}
		sg.MaxPortRange = sg.MinPortRange
	}

	if sg.RemoteType == global.SECURITYGROUP_REMOTETYPE_SECURITYGROUP && sg.RemoteSecurityGroupId == "" {
		err = errors.New("No remoteSecurityGroupId is available")
		return
	}

	body := map[string]interface{}{
		"securityGroupId": sg.SecurityGroupId,
		"protocol":        sg.Protocol.String(),
		"direction":       sg.Direction.String(),
	}

	body["portRangeMin"] = sg.MinPortRange
	body["portRangeMax"] = sg.MaxPortRange

	switch sg.RemoteType {
	case global.SECURITYGROUP_REMOTETYPE_CIDR:
		if sg.RemoteIpPrefix == "" {
			body["remoteIpPrefix"] = "0.0.0.0/0"
		} else {
			body["remoteIpPrefix"] = sg.RemoteIpPrefix
		}
	case global.SECURITYGROUP_REMOTETYPE_SECURITYGROUP:
		body["remoteGroupId"] = sg.RemoteSecurityGroupId
	}

	resp, err := a.client.NewRequest("POST", "/api/securitygrouprule", nil, nil, body)
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

func (a *APIv2) GetSecurityGroupRules(securityGroupId, ruleId string) (result []SecurityGroupRuleResult, err error) {
	params := map[string]interface{}{}

	if securityGroupId != "" {
		params["securityGroupId"] = securityGroupId
	}

	if ruleId != "" {
		params["id"] = ruleId
	}

	resp, err := a.client.NewRequest("GET", "/api/securitygrouprule", nil, params, nil)
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

func (a *APIv2) DeleteSecurityGroupRules(ruleId string) error {
	if ruleId == "" {
		return errors.New("No ruleId is available")
	}

	resp, err := a.client.NewRequest("DELETE", "/api/securitygrouprule/"+ruleId, nil, nil, nil)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}
