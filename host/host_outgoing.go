package host

type OutgoingResponse struct {
	Success   bool              `json:"success"`
	ErrorText string            `json:"error_text,omitempty"`
	Values    []OutgoingCounter `json:"values,omitempty"`
}

type OutgoingCounter struct {
	Host            string `json:"host"`
	OutgoingPackets int64  `json:"outgoing_packets"`
	OutgoingBytes   int64  `json:"outgoing_bytes"`
	OutgoingFlows   int64  `json:"outgoing_flows"`
}
