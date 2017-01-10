package paypal

type Error struct {
	Name            string        `json:"name"`
	DebugId         string        `json:"debug_id"`
	Message         string        `json:"message"`
	InformationLink string        `json:"information_link"`
	Details         []ErrorDetail `json:"details"`
}

type ErrorDetail struct {
	Field string `json:"field"`
	Issue string `json:"issue"`
}
