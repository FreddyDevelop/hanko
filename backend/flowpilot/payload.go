package flowpilot

import "github.com/FreddyDevelop/hanko/backend/v2/flowpilot/jsonmanager"

type payload interface {
	jsonmanager.JSONManager
}

// newPayload creates a new instance of Payload with empty JSON data.
func newPayload() payload {
	return jsonmanager.NewJSONManager()
}
