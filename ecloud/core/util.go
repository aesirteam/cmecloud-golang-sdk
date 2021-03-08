package core

import (
	"fmt"

	json "github.com/json-iterator/go"
)

func Dump(a interface{}) string {
	if bytes, err := json.MarshalIndent(&a, "", "  "); err == nil {
		return fmt.Sprintf("\n%s", bytes)
	}
	return ""
}
