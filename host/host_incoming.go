package host

type IncomingResponse struct {
	Success   bool              `json:"success"`
	ErrorText string            `json:"error_text,omitempty"`
	Values    []IncomingCounter `json:"values,omitempty"`
}

type IncomingCounter struct {
	Host            string `json:"host"`
	IncomingPackets int64  `json:"incoming_packets"`
	IncomingBytes   int64  `json:"incoming_bytes"`
	IncomingFlows   int64  `json:"incoming_flows"`
}
