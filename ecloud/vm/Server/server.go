package Server

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	json "github.com/json-iterator/go"
)

var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC/VpRysi0bPRLS7sbgQDJHo1MAt9/bK
+nwK5Pe3z0/O4cH5I/8kFNYy4yFsLMM+zyFvVw9C4wzjHaRcmEuF3ziJMC9PD5ufUWgfO
5nSGgZW1cmgjqnhcWJ3i+Azj72RnhKQRCn9DgJduEC9MiKfbyTICGd6FXf9cxb21nkxI7vtwIDAQAB
-----END PUBLIC KEY-----
`)

func rsaEncrypt(origData []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey) //将密钥解析成公钥实例
	if block == nil {
		return nil, errors.New("public key error")
	}

	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes) //解析pem.Decode（）返回的Block指针实例
	if err != nil {
		return nil, err
	}

	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData) //RSA算法加密
}

func (a *APIv2) CreatServer(ss *ServerSpec) (result ServerOrderResult, err error) {
	if ss.Name == "" {

	}

	if ss.Cpu == 0 {

	}

	if ss.Ram == 0 {

	}

	if ss.ImageType == 0 {

	}

	if ss.VmType == 0 {

	}

	if ss.Region == "" {

	}

	if ss.Password == "" && ss.KeypairName == "" {

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
		body["password"], _ = rsaEncrypt([]byte(ss.Password))
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

	//body["securityGroupIds"] = a.SecurityGroupIds

	//if bytes, err := json.MarshalIndent(&body, "","  "); err == nil {
	//	fmt.Printf("%s\n", bytes)
	//}

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

func (a *APIv2) GetServerList(ss *ServerSpec, page, size int) (string, error) {
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
		return "", fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
	}

	return resp.Body, nil
}

func (a *APIv2) GetServerInfo(serverId string, detail bool) {

}

//func (a *APIv2) GetServerQuotaInfo() {
//
//}

func (a *APIv2) GetServerVNCAddress(serverId string) {

}

func (a *APIv2) UpdateServerName(id, name string) {

}

func (a *APIv2) UpdateServerPassword(id, password string) {

}

func (a *APIv2) RebootServer(serverId string) {

}

func (a *APIv2) StartServer(serverId string) {

}

func (a *APIv2) StopServer(serverId string) {

}

func (a *APIv2) GetRebuildImageList(serverId string, imageType int) {

}

func (a *APIv2) RebuildServer(serverId, imageId string, adminPass, userData string) {

}
