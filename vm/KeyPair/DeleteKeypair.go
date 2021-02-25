package KeyPair

import (
	"errors"
	"github.com/aesirteam/cmecloud-golang-sdk/core"
)

func (c *ApiKeyPair) DeleteKeyPair(keyId string) (*core.AcsResponse, error) {
	if keyId == "" {
		return nil, errors.New("No KeyId is available")
	}

	return c.Client.Delete(serverPath(c.Version)+"/"+keyId, nil, nil)
}
