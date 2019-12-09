package kohhi

import (
	"encoding/json"
	"kohhiapi/data"
	"kohhiapi/types"
)

// GetProfile is get profile for kohhi.
func GetProfile() (kohhiProfile types.KohhiProfile, err error) {
	err = json.Unmarshal(data.KohhiProfileJSON, &kohhiProfile)
	return kohhiProfile, err
}
