package media

// Media holds information about a media file.
type Media struct {
	ContentType string `json:"content_type,omitempty"`
	FileName    string `json:"file_name,omitempty"`
	MediaID     string `json:"media_id,omitempty"`
	Size        int    `json:"size,omitempty"`
	UploadTime  string `json:"upload_time,omitempty"`
	URL         string `json:"media_url,omitempty"`
}

// UploadResponse holds the result of a media upload.
type UploadResponse struct {
	ContentType  string `json:"content_type,omitempty"`
	FileName     string `json:"file_name,omitempty"`
	MediaID      string `json:"media_id,omitempty"`
	Size         int    `json:"size,omitempty"`
	UploadTime   string `json:"upload_time,omitempty"`
	URL          string `json:"url,omitempty"`
	Status       string `json:"status,omitempty"`
	StatusCode   int    `json:"status_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	ErrorCode    int    `json:"error_code,omitempty"`
}

// Meta holds pagination metadata for media lists.
type Meta struct {
	Previous   *string
	Next       *string
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset,omitempty"`
	Limit      int `json:"limit,omitempty"`
}

// UploadResponseBody wraps a list of upload results.
type UploadResponseBody struct {
	Media []UploadResponse `json:"objects"`
	ApiID string           `json:"api_id"`
}

// ListResponse wraps a paginated list of media.
type ListResponse struct {
	ApiID string  `json:"api_id"`
	Meta  Meta    `json:"meta"`
	Media []Media `json:"objects"`
}

// File holds a file path and content type for upload.
type File struct {
	FilePath    string
	ContentType string
}

// UploadParams holds files to upload.
type UploadParams struct {
	Files []File
}

// ListParams filters the media list.
type ListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}
