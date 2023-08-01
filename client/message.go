package client

type AddTaskResponse struct {
	Message   string `json:"message"`
	ProcessID string `json:"process_id"`
}

// taskstatus result request
type TaskStatusRequest struct {
	ProcessID string `json:"process_id"`
}

type TaskResultResponse struct {
	Message      string       `json:"message"`
	ResponseData ResponseData `json:"response_data"`
}

func (t TaskResultResponse) Status() string {
	return t.ResponseData.Status
}

func (t TaskResultResponse) GetErrMessage() string {
	return t.ResponseData.GetErrMessage()
}

func (t TaskResultResponse) GetOutput() []string {
	return t.ResponseData.GetOutput()
}

func (t TaskResultResponse) GetText() string {
	return t.ResponseData.GetText()
}

const (
	TASK_STATUS_QUEUE     = "IN_QUEUE"
	TASK_STATUS_PROGRESS  = "IN_PROGRESS"
	TASK_STATUS_COMPLETED = "COMPLETED"
	TASK_STATUS_FAILED    = "FAILED"
)

type ResponseData struct {
	ProcessID  string      `json:"process_id"`
	Status     string      `json:"status"`
	Result     interface{} `json:"result"`
	CreditUsed int         `json:"credit_used"`
	Overage    int         `json:"overage"`
}

func (d ResponseData) GetErrMessage() string {
	if d.Status == TASK_STATUS_FAILED {
		return d.Result.(map[string]interface{})["errorMessage"].(string)
	}
	return ""
}

func (d ResponseData) GetOutput() []string {
	if d.Status == TASK_STATUS_COMPLETED {
		slice := d.Result.(map[string]interface{})["output"].([]interface{})
		stringSlice := make([]string, len(slice))
		for i, v := range slice {
			if str, ok := v.(string); ok {
				stringSlice[i] = str
			}
		}
		return stringSlice
	}
	return nil
}

func (d ResponseData) GetText() string {
	if d.Status == TASK_STATUS_COMPLETED {
		return d.Result.(map[string]interface{})["text"].(string)
	}
	return ""
}
