package vm

import (
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud"
	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/core"
	"strconv"
	"testing"
	"time"
)

func TestKeypair(t *testing.T) {
	cli := ecloud.NewForConfigDie(&core.Config{
		ApiGwHost: "api-guiyang-1.cmecloud.cn",
		//ApiGwPort:     8443,
		ApiGwProtocol: "https",
	})

	var (
		name  = "kp" + strconv.Itoa(time.Now().Nanosecond())
		keyId string
	)

	t.Run("CreateKeypair", func(t *testing.T) {
		result, err := cli.Keypair().CreateKeypair(name, "")
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result.Dump())
	})

	t.Run("GetKeypairList", func(t *testing.T) {
		result, err := cli.Keypair().GetKeypairList(name, "", 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result.Dump())

		keyId = result[0].Id
	})

	t.Run("DeleteKeypair", func(t *testing.T) {
		err := cli.Keypair().DeleteKeypair(keyId)
		if err != nil {
			t.Fatal(err)
		}
	})
}
