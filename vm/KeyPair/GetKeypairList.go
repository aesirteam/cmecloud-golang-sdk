package KeyPair

import (
	"github.com/aesirteam/cmecloud-golang-sdk/core"
)

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

	return c.Client.Get(serverPath(c.Version), nil, param)
}
