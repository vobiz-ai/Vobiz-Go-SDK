package recordings

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

// Recording represents a call recording.
type Recording struct {
	AddTime                       string  `json:"add_time,omitempty"`
	CallUUID                      string  `json:"call_uuid,omitempty"`
	MonthlyRecordingStorageAmount float64 `json:"monthly_recording_storage_amount,omitempty"`
	RecordingStorageDuration      int64   `json:"recording_storage_duration,omitempty"`
	RecordingStorageRate          float64 `json:"recording_storage_rate,omitempty"`
	RecordingID                   string  `json:"recording_id,omitempty"`
	RecordingType                 string  `json:"recording_type,omitempty"`
	RecordingFormat               string  `json:"recording_format,omitempty"`
	ConferenceName                string  `json:"conference_name,omitempty"`
	RecordingURL                  string  `json:"recording_url,omitempty"`
	ResourceURI                   string  `json:"resource_uri,omitempty"`
	RecordingStartMS              string  `json:"recording_start_ms,omitempty"`
	RecordingEndMS                string  `json:"recording_end_ms,omitempty"`
	RecordingDurationMS           string  `json:"recording_duration_ms,omitempty"`
	RoundedRecordingDuration      int64   `json:"rounded_recording_duration,omitempty"`
	FromNumber                    string  `json:"from_number,omitempty"`
	ToNumber                      string  `json:"to_number,omitempty"`
}

// ListParams filters the recording list endpoint.
type ListParams struct {
	Subaccount                             string `url:"subaccount,omitempty"`
	CallUUID                               string `url:"call_uuid,omitempty"`
	AddTimeLessThan                        string `url:"add_time__lt,omitempty"`
	AddTimeGreaterThan                     string `url:"add_time__gt,omitempty"`
	AddTimeLessOrEqual                     string `url:"add_time__lte,omitempty"`
	AddTimeGreaterOrEqual                  string `url:"add_time__gte,omitempty"`
	Limit                                  int64  `url:"limit,omitempty"`
	Offset                                 int64  `url:"offset,omitempty"`
	FromNumber                             string `url:"from_number,omitempty"`
	ToNumber                               string `url:"to_number,omitempty"`
	ConferenceName                         string `url:"conference_name,omitempty"`
	MpcName                                string `url:"mpc_name,omitempty"`
	ConferenceUuid                         string `url:"conference_uuid,omitempty"`
	MpcUuid                                string `url:"mpc_uuid,omitempty"`
	RecordingStorageDurationEquals         int64  `url:"recording_storage_duration,omitempty"`
	RecordingStorageDurationLessThan       int64  `url:"recording_storage_duration__lt,omitempty"`
	RecordingStorageDurationGreaterThan    int64  `url:"recording_storage_duration__gt,omitempty"`
	RecordingStorageDurationLessOrEqual    int64  `url:"recording_storage_duration__lte,omitempty"`
	RecordingStorageDurationGreaterOrEqual int64  `url:"recording_storage_duration__gte,omitempty"`
}

// ListResponse wraps a paginated list of recordings.
type ListResponse struct {
	ApiID   string       `json:"api_id"`
	Meta    *models.Meta `json:"meta"`
	Objects []*Recording `json:"objects"`
}
