package applications

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

// Application represents a Vobiz application.
type Application struct {
	FallbackMethod      string `json:"fallback_method,omitempty"`
	DefaultApp          bool   `json:"default_app,omitempty"`
	AppName             string `json:"app_name,omitempty"`
	ProductionApp       bool   `json:"production_app,omitempty"`
	AppID               string `json:"app_id,omitempty"`
	HangupURL           string `json:"hangup_url,omitempty"`
	AnswerURL           string `json:"answer_url,omitempty"`
	MessageURL          string `json:"message_url,omitempty"`
	ResourceURI         string `json:"resource_uri,omitempty"`
	HangupMethod        string `json:"hangup_method,omitempty"`
	MessageMethod       string `json:"message_method,omitempty"`
	FallbackAnswerURL   string `json:"fallback_answer_url,omitempty"`
	AnswerMethod        string `json:"answer_method,omitempty"`
	ApiID               string `json:"api_id,omitempty"`
	LogIncomingMessages bool   `json:"log_incoming_messages,omitempty"`
	PublicURI           bool   `json:"public_uri,omitempty"`
	DefaultNumberApp    bool   `json:"default_number_app,omitempty"`
	DefaultEndpointApp  bool   `json:"default_endpoint_app,omitempty"`
}

// CreateParams holds parameters for creating an application.
type CreateParams = Application

// UpdateParams holds parameters for updating an application.
type UpdateParams = Application

// CreateResponse is the response from creating an application.
type CreateResponse struct {
	Message string `json:"message"`
	ApiID   string `json:"api_id"`
	AppID   string `json:"app_id"`
}

// UpdateResponse is the response from updating an application.
type UpdateResponse = models.BaseResponse

// ListParams filters the application list.
type ListParams struct {
	Subaccount string `url:"subaccount,omitempty"`
	AppName    string `url:"app_name,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
}

// ListResponse wraps a paginated list of applications.
type ListResponse struct {
	models.BaseListResponse
	Objects []Application `json:"objects"`
}

// DeleteParams holds optional parameters for deleting an application.
type DeleteParams struct {
	Cascade                bool   `json:"cascade"`
	NewEndpointApplication string `json:"new_endpoint_application,omitempty"`
}
