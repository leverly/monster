package client

type TaskResponse struct {
	Message   string `json:"message"`
	ProcessID string `json:"process_id"`
}
