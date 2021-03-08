package KeyPair

import (
	"errors"
	"fmt"

	json "github.com/json-iterator/go"
)

func (a *APIv2) CreateKeypair(name, region string) (result KeypairResult, err error) {
	if name == "" {
		err = errors.New("No keyName is available")
		return
	}

	body := map[string]interface{}{
		"name": name,
	}

	if region != "" {
		body["region"] = region
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/keypair", nil, nil, body)
	if err != nil {
		err = fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
		return
	}

	result = KeypairResult{
		Name:   resp.Body,
		Region: region,
	}

	return
}

func (a *APIv2) GetKeypairList(name, region string, page, size int) (result KeypairResultArray, err error) {
	params := map[string]interface{}{}

	if name != "" {
		params["keyName"] = name
	}

	if region != "" {
		params["region"] = name
	}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	resp, err := a.client.NewRequest("GET", "/api/v2/keypair", nil, params, nil)
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

func (a *APIv2) DeleteKeypair(keyId string) error {
	if keyId == "" {
		return errors.New("No keyId is available")
	}

	resp, err := a.client.NewRequest("DELETE", "/api/v2/keypair/"+keyId, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("%s %s [%d] %s", resp.Method, resp.SignUrl, resp.StatusCode, err)
	}

	return nil
}
