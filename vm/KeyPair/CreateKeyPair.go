package KeyPair

import (
	"errors"
	"github.com/aesirteam/cmecloud-golang-sdk/core"
)

type ApiKeyPair struct {
	Version string
	Client  *core.AcsClient
}

func serverPath(version string) string {
	switch version {
	case "v1":
		return "/api/keypair"
	default:
		return "/api/v2/keypair"
	}
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

	return c.Client.Post(serverPath(c.Version), nil, nil, body)
}
