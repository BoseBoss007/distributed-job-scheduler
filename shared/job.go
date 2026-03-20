package shared

type Job struct {
	ID         string `json:"id"`
	Command    string `json:"command"`
	Priority   string `json:"priority"`
	Status     string `json:"status"`
	Output     string `json:"output"`
	RetryCount int    `json:"retry_count"`
	MaxRetries int    `json:"max_retries"`
	CreatedAt  int64  `json:"created_at"`
}