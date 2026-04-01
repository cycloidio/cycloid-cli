// JSON shapes for Concourse-style build event payloads (SSE data lines).
package buildwatch

import "encoding/json"

type eventEnvelope struct {
	Event   string          `json:"event"`
	Version string          `json:"version"`
	EventID string          `json:"event_id"`
	Data    json.RawMessage `json:"data"`
}

type eventOrigin struct {
	ID     string `json:"id"`
	Source string `json:"source"`
}

type eventDataStatus struct {
	Status string `json:"status"`
	Time   int64  `json:"time"`
}

type eventDataLog struct {
	Time    int64       `json:"time"`
	Origin  eventOrigin `json:"origin"`
	Payload string      `json:"payload"`
}

type eventDataTask struct {
	Time       int64       `json:"time"`
	Origin     eventOrigin `json:"origin"`
	ExitStatus int         `json:"exit_status"`
}

type eventDataFinish struct {
	Time      int64       `json:"time"`
	Origin    eventOrigin `json:"origin"`
	Succeeded *bool       `json:"succeeded"`
	Name      string      `json:"name"`
}

type eventDataFinishGet struct {
	Time       int64       `json:"time"`
	Origin     eventOrigin `json:"origin"`
	ExitStatus int         `json:"exit_status"`
}
