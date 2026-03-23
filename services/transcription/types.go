package transcription

// GetParams filters a transcription fetch by type.
type GetParams struct {
	Type string `url:"type"`
}

// CreateParams holds the recording ID and optional callback URL.
type CreateParams struct {
	RecordingID              string `json:"-"`
	TranscriptionCallbackUrl string `json:"transcription_callback_url,omitempty"`
}

// DeleteParams holds the transcription ID to delete.
type DeleteParams struct {
	TranscriptionID string
}

// GetResponse is the response from fetching a transcription.
type GetResponse struct {
	APIID               string      `json:"api_id"`
	Cost                float64     `json:"cost"`
	Rate                float64     `json:"rate"`
	RecordingDurationMs float64     `json:"recording_duration_ms"`
	RecordingStartMs    float64     `json:"recording_start_ms"`
	Status              string      `json:"status"`
	Transcription       interface{} `json:"transcription"`
}
