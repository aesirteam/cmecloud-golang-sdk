package vm

import (
	"testing"

	"github.com/aesirteam/cmecloud-golang-sdk/ecloud/global"
)

func TestServerBackup(t *testing.T) {

	t.Run("OrderServerBackup", func(t *testing.T) {
		result, err := vm.OrderServerBackup()
		if err != nil {
			t.Fatal(err)
		}

		t.Log(result)
	})

	t.Run("CreateServerBackup", func(t *testing.T) {
		serverId = getServerId()

		result, err := vm.CreateServerBackup(serverId, serverBackupName)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("GetServerBackupList", func(t *testing.T) {
		result, err := vm.GetServerBackupList("", "", 0, 0)
		if err != nil {
			t.Fatal(err)
		}

		t.Log(global.Dump(result))
	})

	t.Run("RestoreServerFromBackup", func(t *testing.T) {
		serverId = getServerId()
		serverBackupId = getServerBackupId()

		err := vm.RestoreServerFromBackup(serverBackupId, serverId)
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("DeleteServerBackup", func(t *testing.T) {
		serverBackupId = getServerBackupId()

		err := vm.DeleteServerBackup(serverBackupId)
		if err != nil {
			t.Fatal(err)
		}
	})
}
