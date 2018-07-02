package notifier

import "time"

//EventType holds the information regarding an event
type EventType struct {
	Event     string                 `json:"event"`
	ClientID  string                 `json:"client_id"`
	Principal string                 `json:"principal"`
	Target    Target                 `json:"target"`
	Detail    map[string]interface{} `json:"detail"`
	CreatedAt time.Time
}

//Target target of an event
type Target struct {
	ResourceGroup string `json:"resource_group"`
	ResourceARN   string `json:"resource_arn"`
}
