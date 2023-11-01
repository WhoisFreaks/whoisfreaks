package modal

type Error struct {
	Timestamp interface{} `json:"timestamp,omitempty"`
	Status    int         `json:"status,omitempty"`
	Error     string      `json:"error,omitempty"`
	Message   string      `json:"message,omitempty"`
	Path      string      `json:"path,omitempty"`
}
