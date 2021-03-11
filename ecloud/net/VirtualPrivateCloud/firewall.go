package VirtualPrivateCloud

import "errors"

//Deprecated: AddFirewallRules
func (a *APIv2) AddFirewallRules() {

}

//Deprecated: DeleteFirewallRule
func (a *APIv2) DeleteFirewallRule() {

}

//Deprecated: GetFirewallQuota
func (a *APIv2) GetFirewallQuota() {

}

//Deprecated: GetFirewallRulesList
func (a *APIv2) GetFirewallRulesList() {

}

//Deprecated: UpdateFirewallRules
func (a *APIv2) UpdateFirewallRules() {

}

//Deprecated: FirewallBindRouter
func (a *APIv2) FirewallBindRouter(firewallId, routerId string) (result string, err error) {
	if firewallId == "" {
		err = errors.New("No firewallId is available")
		return
	}

	if routerId == "" {
		err = errors.New("No routerId is available")
		return
	}

	body := map[string]interface{}{
		"firewallId": firewallId,
		"routerId":   routerId,
	}

	resp, err := a.client.NewRequest("POST", "/api/firewall/router/bind", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	result = resp.Body

	return
}

//Deprecated: FirewallUnbindRouter
func (a *APIv2) FirewallUnbindRouter() {

}
