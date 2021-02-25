package KeyPair

import (
	"github.com/aesirteam/cmecloud-golang-sdk/core"
	json "github.com/json-iterator/go"
	"strconv"
	"testing"
	"time"
)

func TestKeyPair(t *testing.T) {
	api := &ApiKeyPair{
		Client: &core.AcsClient{
			ApiGwHost: "api-guiyang-1.cmecloud.cn",
			//ApiGwPort: 8443,
			ApiGwProtocol: "https",
		},
	}

	keyName := "keypair" + strconv.Itoa(time.Now().Nanosecond())
	keyId := ""

	t.Run("CreateKeyPair", func(t *testing.T) {
		resp, err := api.CreateKeyPair(keyName, "")

		if err != nil {
			t.Errorf("%s %s [%d]\n%s", resp.Method, resp.SignUrl, resp.StatusCode, err)
			return
		}

		t.Logf("%s %s [%d]\n%s", resp.Method, resp.SignUrl, resp.StatusCode, resp.Body)
	})

	t.Run("GetKeypairList", func(t *testing.T) {
		resp, err := api.GetKeypairList(keyName, 1, 0)

		if err != nil {
			t.Errorf("%s %s [%d]\n%s", resp.Method, resp.SignUrl, resp.StatusCode, err)
			return
		}

		//content[0].id
		if obj := json.Get([]byte(resp.Body), "content", 0, "id"); obj.LastError() == nil {
			keyId = obj.ToString()
		}

		t.Logf("%s %s [%d]\n%s", resp.Method, resp.SignUrl, resp.StatusCode, resp.Body)
	})

	t.Run("DeleteKeypair", func(t *testing.T) {
		if keyId == "" {
			t.Fatal("No keyId is available")
		}

		resp, err := api.DeleteKeyPair(keyId)

		if err != nil {
			t.Errorf("%s %s [%d]\n%s", resp.Method, resp.SignUrl, resp.StatusCode, err)
			return
		}

		t.Logf("%s %s [%d]\n%s", resp.Method, resp.SignUrl, resp.StatusCode, resp.Body)
	})
}
