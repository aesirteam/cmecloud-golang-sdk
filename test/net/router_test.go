package net

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestRouter(t *testing.T) {

	t.Run("GetRouterNetList", func(t *testing.T) {
		_, routerId = getRouterId()

		result, err := net.GetRouterNetList(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetRouterInfo", func(t *testing.T) {
		_, routerId = getRouterId()

		result, err := net.GetRouterInfo(routerId)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})
}
