package ELB

import (
	"errors"
	"strings"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func (a *APIv2) CreateELB(es *global.ELBSpec) (result ELBOrderResult, err error) {

	if es.LoadBalanceName == "" {
		err = errors.New("No loadBalanceName is available")
		return
	}

	if es.SubnetId == "" {
		err = errors.New("No subnetId is available")
		return
	}

	body := map[string]interface{}{
		"chargePeriodEnum": strings.ToLower(es.ChargePeriod.String()),
		"loadBalanceName":  es.LoadBalanceName,
		"subnetId":         es.SubnetId,
		"flavor":           "1",
	}

	if es.FloatingIpId != "" {
		body["ipId"] = es.FloatingIpId
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/netcenter/order/create/loadbalance", nil, nil, body)
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

func (a *APIv2) GetELBList(queryWord, routerId, networkId string, fipBind, visible bool, page, size int) (result []ELBResult, err error) {
	params := map[string]interface{}{}

	if queryWord != "" {
		params["queryWord"] = queryWord
	}

	if routerId != "" {
		params["routerId"] = routerId
	}

	if networkId != "" {
		params["networkId"] = networkId
	}

	if fipBind {
		params["fipBind"] = true
	}

	if visible {
		params["visible"] = true
	}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/loadbalance", nil, params, nil)
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

func (a *APIv2) GetELBQuota() (result ELBQuotaResult, err error) {
	resp, err := a.client.NewRequest("GET", "/api/elb/quota", nil, nil, nil)
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

func (a *APIv2) DeleteELB(loadBalanceId string) error {
	resp, err := a.client.NewRequest("DELETE", "/api/v2/loadbalance/"+loadBalanceId, nil, nil, nil)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

func (a *APIv2) CreateELBListener(el *global.ELBListenerSpec) (result string, err error) {
	if el.LoadbalanceId == "" {
		err = errors.New("No loadbalanceId is available")
		return
	}

	if el.Name == "" {
		err = errors.New("No name is available")
		return
	}

	if el.ProtocolPort == 0 || el.ProtocolPort > 65535 {
		err = errors.New("No protocolPort is available")
		return
	}

	body := map[string]interface{}{
		"loadbalanceId":      el.LoadbalanceId,
		"name":               el.Name,
		"protocol":           el.Protocol.String(),
		"protocolPort":       el.ProtocolPort,
		"lbAlgorithm":        el.Algorithm.String(),
		"sessionPersistence": el.SessionPersistence.String(),
		"cookieName":         el.CookieName,
		"healthType":         el.HealthType.String(),
		"healthHttpMethod":   el.HealthHttpMethod,
		"healthUrlPath":      el.HealthUrlPath,
		"healthExpectedCode": el.HealthExpectedCode,
	}

	if el.HealthDelay == 0 {
		body["healthDelay"] = 30
	}

	if el.HealthMaxRetries == 0 {
		body["healthMaxRetries"] = 3
	}

	if el.HealthTimeout == 0 {
		body["healthTimeout"] = 5
	}

	if el.HealthUrlPath == "" {
		body["healthUrlPath"] = "/"
	}
	if el.LoadbalanceId == "" {
		err = errors.New("No loadbalanceId is available")
		return
	}
	if el.HealthExpectedCode == "" {
		body["healthExpectedCode"] = "200"
	}

	if el.MutualAuthenticationUp {
		body["mutualAuthenticationUp"] = true
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/loadbalance/listener", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	result = resp.Body

	return
}

func (a *APIv2) GetELBListenerList(loadbalanceId string) (result []ELBListenerResult, err error) {
	if loadbalanceId == "" {
		err = errors.New("No loadbalanceId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/loadbalance/listener/"+loadbalanceId+"/list", nil, nil, nil)
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

func (a *APIv2) GetELBListenerInfo(listenerId string) (result ELBListenerResult, err error) {
	if listenerId == "" {
		err = errors.New("No listenerId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/loadbalance/listener/"+listenerId, nil, nil, nil)
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

func (a *APIv2) ModifyELBListenerName(el *global.ELBListenerSpec) (err error) {
	return
}

func (a *APIv2) DeleteELBListener(listenerId string) (err error) {
	if listenerId == "" {
		err = errors.New("No listenerId is available")
		return
	}

	resp, err := a.client.NewRequest("DELETE", "/api/v2/loadbalance/listener/"+listenerId, nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) AddELBMember(ms *global.ELBMemberSpec) (result string, err error) {
	if ms.PoolId == "" {
		err = errors.New("No poolId is available")
		return
	}

	if ms.Ip == "" {
		err = errors.New("No ip is available")
		return
	}

	if ms.Port == 0 || ms.Port > 65535 {
		err = errors.New("No port is available")
		return
	}

	if ms.Weight == 0 || ms.Weight > 100 {
		err = errors.New("No weight is available")
		return
	}

	if ms.Type == 0 || ms.Type > 2 {
		err = errors.New("No type is available")
		return
	}

	if ms.VmHostId == "" {
		err = errors.New("No vmHostId is available")
		return
	}

	body := map[string]interface{}{
		"ip":       ms.Ip,
		"port":     ms.Port,
		"weight":   ms.Weight,
		"type":     ms.Type,
		"vmHostId": ms.VmHostId,
	}

	if ms.Desc != "" {
		body["description"] = ms.Desc
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/network/lbpoolmember/"+ms.PoolId+"/member", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	result = resp.Body

	return
}

func (a *APIv2) GetELBMemberList(poolId string) (result []ELBMemberResult, err error) {
	if poolId == "" {
		err = errors.New("No poolId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/network/lbpoolmember/"+poolId+"/members", nil, nil, nil)
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

func (a *APIv2) GetELBMemberInfo(poolId, memberId string) (result ELBMemberResult, err error) {
	if poolId == "" {
		err = errors.New("No poolId is available")
		return
	}

	if memberId == "" {
		err = errors.New("No memberId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/network/lbpoolmember/healthStatus/"+poolId+"/"+memberId, nil, nil, nil)
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

func (a *APIv2) DeleteELBMember(poolId, memberId string) (err error) {
	if poolId == "" {
		err = errors.New("No poolId is available")
		return
	}

	if memberId == "" {
		err = errors.New("No memberId is available")
		return
	}

	resp, err := a.client.NewRequest("DELETE", "/api/v2/network/lbpoolmember/"+poolId+"/member/"+memberId+"/delete", nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	return
}
