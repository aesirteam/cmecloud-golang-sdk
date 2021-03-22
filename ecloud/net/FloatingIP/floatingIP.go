package FloatingIP

import (
	"errors"
	"strings"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func (a *APIv2) CreateFloatingIp(fs *global.FloatingIpSpec) (result FloatingIpOrderResult, err error) {
	body := map[string]interface{}{
		"chargeModeEnum":   fs.ChargeMode.String(),
		"chargePeriodEnum": strings.ToLower(fs.ChargePeriod.String()),
		"ipType":           "MOBILE",
		//"duration":         fs.Duration,
		"quantity": 1,
	}

	if fs.ChargeMode == global.FLOATINGIP_CHARGEMODE_TRAFFIC {
		body["chargePeriodEnum"] = strings.ToLower(global.BILLING_TYPE_HOUR.String())
	}

	switch fs.ChargePeriod {
	case global.BILLING_TYPE_YEAR:
		if fs.Duration == 0 {
			body["duration"] = 12
		} else if fs.Duration > 0 && fs.Duration <= 5*12 {
			body["duration"] = fs.Duration
		}
	case global.BILLING_TYPE_MONTH:
		if fs.Duration == 0 {
			body["duration"] = 1
		} else if fs.Duration > 0 && fs.Duration <= 12 {
			body["duration"] = fs.Duration
		}
	}

	if fs.Quantity > 0 {
		body["quantity"] = fs.Quantity
	}

	if fs.BandwidthSize == 0 {
		body["bandwidthSize"] = 1
	} else if fs.BandwidthSize > 1024 {
		body["bandwidthSize"] = 1024
	} else {
		body["bandwidthSize"] = fs.BandwidthSize
	}

	if fs.FloatingIpAddress != "" {
		body["floatingIpAddress"] = fs.FloatingIpAddress
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/netcenter/order/create/floatingip", nil, nil, body)
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

func (a *APIv2) GetFloatingIpList(queryWord, routerId, natGatewayId string, bound, frozen, openIpv6Converter, availableForSbw bool, tagIds []string, page, size int) (result []FloatingIpResult, err error) {
	params := map[string]interface{}{
		"ipType": "MOBILE",
	}

	if queryWord != "" {
		params["queryWord"] = queryWord
	}

	if routerId != "" {
		params["routerId"] = routerId
	}

	if natGatewayId != "" {
		params["natGatewayId"] = natGatewayId
	}

	if bound {
		params["bound"] = true
	}

	if frozen {
		params["frozen"] = true
	}

	if openIpv6Converter {
		params["openIpv6Converter"] = true
	}

	if availableForSbw {
		params["availableForSbw"] = true
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

	resp, err := a.client.NewRequest("GET", "/api/v2/floatingIp/listWithBw", nil, params, nil)
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

func (a *APIv2) GetFloatingIpDetail(ipId string) (result *FloatingIpResult, err error) {
	if ipId == "" {
		err = errors.New("No ipId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/floatingIp/getRespWithBw/"+ipId, nil, nil, nil)
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

func (a *APIv2) GetFloatingIpQuotaWithK() (result FloatingIpQuotaResult, err error) {

	resp, err := a.client.NewRequest("GET", "/api/floatingip/quota", nil, nil, nil)
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

func (a *APIv2) GetFloatingIpQuotaWithP(quotaName global.FloatingIpQuotaName, resourceId string) (result FloatingIpQuotaResult, err error) {
	params := map[string]interface{}{}

	if resourceId != "" {
		params["resourceId"] = resourceId
	}

	resp, err := a.client.NewRequest("GET", "/api/quota/"+quotaName.String(), nil, params, nil)
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

func (a *APIv2) AttachFloatingIp(bindType global.FloatingIpBindType, ipId, resourceId, portId string) error {
	if ipId == "" {
		return errors.New("No ipId is available")
	}

	if resourceId == "" {
		return errors.New("No resourceId is available")
	}

	body := map[string]interface{}{
		"type":       bindType.String(),
		"ipId":       ipId,
		"resourceId": resourceId,
	}

	if bindType == global.FLOATINGIP_BINDTYPE_VM {
		if portId == "" {
			return errors.New("No portId is available")
		}

		body["portId"] = portId
	}

	resp, err := a.client.NewRequest("PUT", "/api/v2/floatingIp/bind", nil, nil, body)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

func (a *APIv2) DetachFloatingIp(ipId string) error {
	if ipId == "" {
		return errors.New("No ipId is available")
	}

	resp, err := a.client.NewRequest("PUT", "/api/v2/floatingIp/unbind/"+ipId, nil, nil, nil)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}
