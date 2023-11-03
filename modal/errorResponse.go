package modal

// Error represents an error response structure containing information about the error occurred during API requests.
type Error struct {
	// Timestamp represents the timestamp when the error occurred. It can be of any type.
	Timestamp interface{} `json:"timestamp,omitempty"`
	// Status represents the HTTP status code associated with the error.
	Status int `json:"status,omitempty"`
	// Error represents the type or category of the error.
	Error string `json:"error,omitempty"`
	// Message provides a detailed description of the error.
	Message string `json:"message,omitempty"`
	// Path represents the request URL path that triggered the error.
	Path string `json:"path,omitempty"`
}
