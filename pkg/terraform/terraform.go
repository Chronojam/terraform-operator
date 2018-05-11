package terraform

import (
	"encoding/json"
)

type Resource struct {
	Resource interface{} `json:"resource"`
}

// RenderToTerraform takes an object, and attempts to construct the appropriate terraform json from it.
func RenderToTerraform(i interface{}) ([]byte, error) {
	r := Resource{
		Resource: i,
	}
	b, err := json.Marshal(r)
	if err != nil {
		return b, err
	}

	return b, nil
}
