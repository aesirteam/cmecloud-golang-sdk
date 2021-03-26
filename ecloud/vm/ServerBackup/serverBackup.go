package ServerBackup

import (
	"errors"
)

func (a *APIv1) ServerBackupOrder() (result string, err error) {
	body := map[string]interface{}{
		"resourceType": "SERVERBACK_SERVICE",
	}

	resp, err := a.client.NewRequest("POST", "/api/v2/vmBackupService/order", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	result = resp.Body

	return
}

func (a *APIv1) CreateServerBackup(serverId, name string) (result ServerBackupResult, err error) {
	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	if name == "" {
		err = errors.New("No name is available")
		return
	}

	body := map[string]interface{}{
		"serverId": serverId,
		"name":     name,
	}

	resp, err := a.client.NewRequest("POST", "/api/server/backup", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv1) GetServerBackupList(serverId, name string, page, size int) (result []ServerBackupResult, err error) {
	params := map[string]interface{}{}

	if serverId != "" {
		params["serverId"] = serverId
	}

	if name != "" {
		params["name"] = name
	}

	if page > 0 {
		params["page"] = page
	}

	if size > 0 {
		params["size"] = size
	}

	resp, err := a.client.NewRequest("GET", "/api/server/backup", nil, params, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	if err = resp.UnmarshalFromContent(&result, ""); err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv1) RestoreServerFromBackup(backupId, serverId string) (err error) {
	if backupId == "" {
		err = errors.New("No backupId is available")
		return
	}

	if serverId == "" {
		err = errors.New("No serverId is available")
		return
	}

	body := map[string]interface{}{
		"id":       backupId,
		"serverId": serverId,
	}

	resp, err := a.client.NewRequest("POST", "/api/server/backup/restore", nil, nil, body)
	if err != nil {
		err = resp.Error(err)
		return
	}

	return
}

func (a *APIv1) DeleteServerBackup(backupId string) (err error) {
	if backupId == "" {
		err = errors.New("No backupId is available")
		return
	}

	resp, err := a.client.NewRequest("DELETE", "/api/server/backup/"+backupId, nil, nil, nil)
	if err != nil {
		err = resp.Error(err)
		return
	}

	return
}
