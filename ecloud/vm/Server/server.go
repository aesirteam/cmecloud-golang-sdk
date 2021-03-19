package Server

import (
	"errors"
	"strconv"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/net/VirtualPrivateCloud"
)

func (a *APIv2) CreatServer(ss *global.ServerSpec) (result ServerOrderResult, err error) {
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

	if ss.ImageName == "" {
		err = errors.New("No imageName is available")
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

	if ss.Networks == nil {
		err = errors.New("No networks is available")
		return
	}

	body := map[string]interface{}{
		"name":        ss.Name,
		"cpu":         ss.Cpu,
		"ram":         ss.Ram,
		"imageName":   ss.ImageName,
		"vmType":      ss.VmType.String(),
		"region":      ss.Region,
		"billingType": ss.BillingType.String(),
		"quantity":    1,
	}

	body["disk"] = func() int {
		switch ss.OsType {
		case global.OS_TYPE_WINDOWS:
			return 40
		default:
			return 20
		}
	}()

	body["bootVolume"] = map[string]interface{}{
		"volumeType": ss.BootVolume.VolumeType.String(),
		"size": func() int {
			if ss.BootVolume.Size <= 0 {
				if ss.OsType == global.OS_TYPE_LINUX {
					return 20
				} else if ss.OsType == global.OS_TYPE_WINDOWS {
					return 40
				}
			}
			return ss.BootVolume.Size
		}(),
	}

	if ss.Password != "" {
		if password, err := global.RsaEncrypt([]byte(ss.Password)); err == nil {
			body["password"] = password
		}
	} else if ss.KeypairName != "" {
		body["keypairName"] = ss.KeypairName
	}

	body["networks"] = map[string]string{
		"networkId": ss.Networks.NetworkId,
		"portId":    ss.Networks.PortId,
	}

	dvs := make([]map[string]interface{}, 0, 15)

	for _, dv := range ss.DataVolumes {
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

	switch ss.BillingType {
	case global.BILLING_TYPE_YEAR:
		body["billingType"] = global.BILLING_TYPE_MONTH.String()
		if ss.Duration == 0 {
			body["dration"] = 12
		} else if ss.Duration > 0 && ss.Duration <= 5*12 {
			body["dration"] = ss.Duration
		}
	case global.BILLING_TYPE_MONTH:
		if ss.Duration == 0 {
			body["dration"] = 1
		} else if ss.Duration > 0 && ss.Duration <= 12 {
			body["dration"] = ss.Duration
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
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) GetServerList(ss *global.ServerSpec, page, size int) (result []ServerResult, err error) {
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
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) GetServerInfo(serverId string, detail bool) (result *ServerResult, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	params := map[string]interface{}{
		"detail": detail,
	}

	resp, err := a.client.NewRequest("GET", "/api/server/"+serverId, nil, params, nil)
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
		err = resp.Error(err)
		return
	}

	result = resp.Body

	return
}

func (a *APIv2) ModifyServerName(serverId, name string) (result ServerResult, err error) {
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
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, "newtag"); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) ModifyServerPassword(serverId, password string) (result ServerResult, err error) {
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
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, "newtag"); err != nil {
		err = resp.Error(err)
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
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, "newtag"); err != nil {
		err = resp.Error(err)
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
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, "newtag"); err != nil {
		err = resp.Error(err)
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
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, "newtag"); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv2) GetRebuildImageList(serverId string, imageType int) (result []ServerRebuildImageResult, err error) {
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
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
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
		if password, err := global.RsaEncrypt([]byte(adminPass)); err == nil {
			body["adminPass"] = password
		}
	}

	if userData != "" {
		body["userData"] = userData
	}

	resp, err := a.client.NewRequest("PUT", "/api/v2/server/rebuild", nil, nil, body)
	if err != nil {
		return resp.Error(err)
	}

	return nil
}

func (a *APIv2) AttachNic(serverId, portId string) (result VirtualPrivateCloud.NicResult, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	if portId == "" {
		err = errors.New("No portId is available")
		return
	}

	body := map[string]interface{}{
		"id":       portId,
		"serverId": serverId,
	}

	resp, err := a.client.NewRequest("POST", "/api/port/attach", nil, nil, body)
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

func (a *APIv2) DetachNic(serverId, portId string) (result VirtualPrivateCloud.NicResult, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	if portId == "" {
		err = errors.New("No portId is available")
		return
	}

	body := map[string]interface{}{
		"id":       portId,
		"serverId": serverId,
	}

	resp, err := a.client.NewRequest("POST", "/api/port/detach", nil, nil, body)
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

func (a *APIv2) GetUnbindNicList(serverId string, resourceType int, page, size int) (result string, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	params := map[string]interface{}{}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	resp, err := a.client.NewRequest("GET", "/api/server/unbindNic/"+serverId+"/"+strconv.Itoa(resourceType), nil, params, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	result = resp.Body

	return
}
