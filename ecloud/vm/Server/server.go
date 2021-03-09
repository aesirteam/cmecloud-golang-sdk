package Server

import (
	"errors"
	"fmt"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
	json "github.com/json-iterator/go"
)

var newJson = json.Config{
	EscapeHTML:             true,
	SortMapKeys:            true,
	ValidateJsonRawMessage: true,
	TagKey:                 "newtag",
}.Froze()

func (a *APIv2) CreatServer(ss ServerSpec) (result ServerOrderResult, err error) {
	if ss.Name == "" {
		err = errors.New("No name is available")
		return
	}

	if ss.Cpu == 0 {
		err = errors.New("Cpu greater than zero")
		return
	}

	if ss.Ram == 0 {
		err = errors.New("Ram greater than zero")
		return
	}

	if ss.ImageType == 0 {
		err = errors.New("No imageType is available")
		return
	}

	if ss.Region == "" {
		err = errors.New("No region is available")
		return
	}

	if ss.Password == "" && ss.KeypairName == "" {
		err = errors.New("No password and keypairName is available")
		return
	}

	body := map[string]interface{}{
		"name":        ss.Name,
		"cpu":         ss.Cpu,
		"ram":         ss.Ram,
		"imageName":   ss.ImageType.String(),
		"vmType":      ss.VmType.String(),
		"region":      ss.Region,
		"billingType": ss.BillingType.String(),
		"dration":     ss.Dration,
		"quantity":    1,
	}

	body["disk"] = func() int {
		switch ss.OsType {
		case OS_TYPE_WINDOWS:
			return 40
		default:
			return 20
		}
	}()

	body["bootVolume"] = map[string]interface{}{
		"volumeType": ss.BootVolume.VolumeType.String(),
		"size": func() int {
			if ss.BootVolume.Size <= 0 {
				if ss.OsType == OS_TYPE_LINUX {
					return 20
				} else if ss.OsType == OS_TYPE_WINDOWS {
					return 40
				}
			}
			return ss.BootVolume.Size
		}(),
	}

	if ss.Password != "" {
		if password, err := core.RsaEncrypt([]byte(ss.Password)); err == nil {
			body["password"] = password
		}
	} else if ss.KeypairName != "" {
		body["keypairName"] = ss.KeypairName
	}

	body["networks"] = map[string]string{
		"networkId": ss.NetworkId,
	}

	var dvs []map[string]interface{}

	for _, dv := range ss.DataVolumeArray {
		if dv.Size >= 10 && dv.Size <= 32768 {
			dvs = append(dvs, map[string]interface{}{
				"resourceType": dv.ResourceType.String(),
				"size":         dv.Size,
				"isShare":      dv.IsShare,
			})
		}
	}

	if len(dvs) > 0 {
		body["dataVolume"] = dvs
	}

	if ss.BillingType == BILLING_TYPE_YEAR {
		body["billingType"] = BILLING_TYPE_MONTH.String()
		if ss.Dration == 0 {
			body["dration"] = 24
		} else if ss.Dration > 0 && ss.Dration <= 60 {
			body["dration"] = ss.Dration
		}
	}

	if ss.Quantity > 0 {
		body["quantity"] = ss.Quantity
	}

	if len(ss.SecurityGroupIds) > 0 {
		body["securityGroupIds"] = ss.SecurityGroupIds
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/server/order", nil, nil, body)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	if _err := json.UnmarshalFromString(resp.Body, &result); _err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, _err)
		return
	}

	return
}

func (a *APIv2) GetServerList(ss *ServerSpec, page, size int) (result ServerResultArray, err error) {
	params := map[string]interface{}{
		"serverTypes":  "VM",
		"productTypes": "NORMAL,AUTOSCALING,VO,CDN,PAAS_MASTER,PAAS_SLAVE,VCPE,EMR,LOGAUDIT",
		"visible":      true,
	}

	if ss.Name != "" {
		params["serverName"] = ss.Name
	}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/server/web/with/network", nil, params, nil)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	obj := json.Get([]byte(resp.Body), "content")
	if obj.LastError() != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, obj.LastError())
		return
	}

	if _err := json.UnmarshalFromString(obj.ToString(), &result); _err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, _err)
		return
	}

	return
}

func (a *APIv2) GetServerInfo(serverId string, detail bool) (result ServerResult, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	params := map[string]interface{}{
		"detail": detail,
	}

	resp, err := a.client.NewRequest("GET", "/api/server/"+serverId, nil, params, nil)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	if _err := newJson.Unmarshal([]byte(resp.Body), &result); _err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, _err)
		return
	}

	return
}

//func (a *APIv2) GetServerQuotaInfo() {
//
//}

func (a *APIv2) GetServerVNCAddress(serverId string) (result string, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	resp, err := a.client.NewRequest("GET", "/api/server/"+serverId+"/vnc", nil, nil, nil)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	result = resp.Body

	return
}

func (a *APIv2) UpdateServerName(serverId, name string) (result ServerResult, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	if name == "" {
		err = errors.New("No serverId is available")
		return
	}

	body := map[string]interface{}{
		"id":   serverId,
		"name": name,
	}

	resp, err := a.client.NewRequest("PUT", "/api/server//updateName", nil, nil, body)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	if _err := newJson.Unmarshal([]byte(resp.Body), &result); _err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, _err)
		return
	}

	return
}

func (a *APIv2) UpdateServerPassword(serverId, password string) (result ServerResult, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	if password == "" {
		err = errors.New("No password is available")
		return
	}

	body := map[string]interface{}{
		"id":       serverId,
		"password": password,
	}

	resp, err := a.client.NewRequest("PUT", "/api/server/updatePassword", nil, nil, body)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	if _err := newJson.Unmarshal([]byte(resp.Body), &result); _err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, _err)
		return
	}

	return
}

func (a *APIv2) RebootServer(serverId string) (result ServerResult, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	body := map[string]interface{}{
		"serverId": serverId,
	}

	resp, err := a.client.NewRequest("POST", "/api/server/"+serverId+"/reboot", nil, nil, body)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	if _err := newJson.Unmarshal([]byte(resp.Body), &result); _err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, _err)
		return
	}

	return
}

func (a *APIv2) StartServer(serverId string) (result ServerResult, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	body := map[string]interface{}{
		"serverId": serverId,
	}

	resp, err := a.client.NewRequest("POST", "/api/server/"+serverId+"/start", nil, nil, body)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	if _err := newJson.Unmarshal([]byte(resp.Body), &result); _err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, _err)
		return
	}

	return
}

func (a *APIv2) StopServer(serverId string) (result ServerResult, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	body := map[string]interface{}{
		"serverId": serverId,
	}

	resp, err := a.client.NewRequest("POST", "/api/server/"+serverId+"/stop", nil, nil, body)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	if _err := newJson.Unmarshal([]byte(resp.Body), &result); _err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, _err)
		return
	}

	return
}

func (a *APIv2) GetRebuildImageList(serverId string, imageType int) (result ServerRebuildImageArray, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	if imageType == 0 {
		imageType = 2
	}

	params := map[string]interface{}{
		"type": imageType,
	}

	resp, err := a.client.NewRequest("GET", "/api/server/"+serverId+"/rebuild/images", nil, params, nil)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	if _err := json.UnmarshalFromString(resp.Body, &result); _err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, _err)
		return
	}

	return
}

func (a *APIv2) RebuildServer(serverId, imageId string, adminPass, userData string) error {
	if serverId == "" {
		return errors.New("No serverId is available")
	}

	if imageId == "" {
		return errors.New("No imageId is available")
	}

	body := map[string]interface{}{
		"serverId": serverId,
		"imageId":  imageId,
	}

	if adminPass != "" {
		if password, err := core.RsaEncrypt([]byte(adminPass)); err == nil {
			body["adminPass"] = password
		}
	}

	if userData != "" {
		body["userData"] = userData
	}

	resp, err := a.client.NewRequest("PUT", "/api/v2/server/rebuild", nil, nil, body)
	if err != nil {
		return fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
	}

	return nil
}

func (a *APIv2) AttachNic(nicId, serverId string) {

}

func (a *APIv2) DetachNic(nicId, serverId string) {

}

func (a *APIv2) GetUnbindNicList(serverId string, resourceType int, page, size int) {

}
