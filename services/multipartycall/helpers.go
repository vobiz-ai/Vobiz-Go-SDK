package multipartycall

import "log"

// MakeMPCId builds the MPC identifier string from UUID or friendly name.
func MakeMPCId(mpcUUID, friendlyName string) string {
	if mpcUUID != "" {
		return "uuid_" + mpcUUID
	}
	if friendlyName != "" {
		return "name_" + friendlyName
	}
	log.Fatal("Need to specify a mpc_uuid or name")
	return ""
}
