package cdr

// ListParams filters CDR records.
type ListParams struct {
	Limit     int    `url:"limit,omitempty"`
	Offset    int    `url:"offset,omitempty"`
	StartDate string `url:"start_date,omitempty"`
	EndDate   string `url:"end_date,omitempty"`
	PerPage   int    `url:"per_page,omitempty"`
	Page      int    `url:"page,omitempty"`
	Context   string `url:"context,omitempty"`
}

// ListResponse wraps CDR records.
type ListResponse struct {
	ApiID   string                   `json:"api_id"`
	Meta    map[string]interface{}   `json:"meta,omitempty"`
	Objects []map[string]interface{} `json:"objects,omitempty"`
}
