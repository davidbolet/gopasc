package models

// Error represent JSON-RPC 2.0 "Error object".
type Error struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
