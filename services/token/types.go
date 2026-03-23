package token

// CreateParams holds parameters for creating a JWT token.
type CreateParams struct {
	Iss           string      `json:"iss,omitempty"`
	Exp           int64       `json:"exp,omitempty"`
	Nbf           int64       `json:"nbf,omitempty"`
	IncomingAllow interface{} `json:"incoming_allow,omitempty"`
	OutgoingAllow interface{} `json:"outgoing_allow,omitempty"`
	Per           interface{} `json:"per,omitempty"`
	App           string      `json:"app,omitempty"`
	Sub           string      `json:"sub,omitempty"`
}

// CreateResponse is the response from creating a token.
type CreateResponse struct {
	Token string `json:"token"`
	ApiID string `json:"api_id"`
}
