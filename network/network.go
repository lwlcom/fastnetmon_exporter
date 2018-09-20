package network

type Response struct {
	Success   bool      `json:"success"`
	ErrorText string    `json:"error_text,omitempty"`
	Values    []Counter `json:"values,omitempty"`
}

type Counter struct {
	NetworkName     string `json:"network_name"`
	IncomingPackets int64  `json:"incoming_packets"`
	OutgoingPackets int64  `json:"outgoing_packets"`
	IncomingBytes   int64  `json:"incoming_bytes"`
	OutgoingBytes   int64  `json:"outgoing_bytes"`
}
