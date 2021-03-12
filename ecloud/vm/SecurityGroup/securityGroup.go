package SecurityGroup

import (
	"errors"
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

func (a *APIv2) GetSecurityGroupList(name string, page, size int) (result []SecurityGroupResult, err error) {
	params := map[string]interface{}{
		"types": "VM",
	}

	if name != "" {
		params["name"] = name
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

func (a *APIv2) AddSecurityGroupRules() {

}

func (a *APIv2) DeleteSecurityGroupRules() {

}

func (a *APIv2) GetSecurityGroupRules() {

}
