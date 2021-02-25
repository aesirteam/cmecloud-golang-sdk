package KeyPair

import (
	"errors"
	"github.com/aesirteam/cmecloud-golang-sdk/v2/core"
)

const ApiKeyPairUri = "/api/v2/keypair"

type ApiKeyPair struct {
	Client *core.AcsClient
}

func (c *ApiKeyPair) CreateKeyPair(keyName, region string) (*core.AcsResponse, error) {
	if keyName == "" {
		return nil, errors.New("No KeyName is available")
	}

	body := core.AcsRequestBody{
		"name": keyName,
	}

	if region != "" {
		body["region"] = region
	}

	return c.Client.Post(ApiKeyPairUri, nil, nil, body)
}

func (c *ApiKeyPair) GetKeypairList(keyName string, page, size int) (*core.AcsResponse, error) {
	param := core.AcsCustomParam{}

	if keyName != "" {
		param["keyName"] = keyName
	}

	if page > 0 {
		param["page"] = page
	}

	if size > 0 {
		param["size"] = size
	}

	return c.Client.Get(ApiKeyPairUri, nil, param)
}

func (c *ApiKeyPair) DeleteKeyPair(keyId string) (*core.AcsResponse, error) {
	if keyId == "" {
		return nil, errors.New("No KeyId is available")
	}

	return c.Client.Delete(ApiKeyPairUri+"/"+keyId, nil, nil)
}
