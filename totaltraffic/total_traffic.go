package totalTraffic

type Response struct {
	Success   bool      `json:"success"`
	ErrorText string    `json:"error_text,omitempty"`
	Values    []Counter `json:"values,omitempty"`
}

type Counter struct {
	CounterName string `json:"counter_name"`
	Value       int64  `json:"value"`
	Unit        string `json:"unit"`
}
