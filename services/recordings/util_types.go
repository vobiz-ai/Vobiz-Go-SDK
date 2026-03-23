package recordings

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

type ExportParams struct {
	RecordingIDs []string `json:"recording_ids,omitempty"`
	CallbackURL  string   `json:"callback_url,omitempty"`
}

type BulkDeleteParams struct {
	RecordingIDs []string `json:"recording_ids,omitempty"`
}

type UtilityResponse = models.BaseResponse
