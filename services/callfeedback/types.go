package callfeedback

// CreateParams holds parameters for submitting call feedback.
type CreateParams struct {
	CallUUID string      `json:"call_uuid"`
	Notes    string      `json:"notes"`
	Rating   interface{} `json:"rating"`
	Issues   []string    `json:"issues"`
}

// CreateResponse is the response from submitting call feedback.
type CreateResponse struct {
	ApiID   string `json:"api_id"`
	Message string `json:"message"`
	Status  string `json:"status"`
}
