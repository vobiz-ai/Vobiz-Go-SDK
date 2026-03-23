package numbers

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

// Number represents an account phone number.
type Number struct {
	Alias                              string `json:"alias,omitempty"`
	VoiceEnabled                       bool   `json:"voice_enabled,omitempty"`
	SMSEnabled                         bool   `json:"sms_enabled,omitempty"`
	MMSEnabled                         bool   `json:"mms_enabled,omitempty"`
	Description                        string `json:"description,omitempty"`
	VobizNumber                        bool   `json:"vobiz_number,omitempty"`
	City                               string `json:"city,omitempty"`
	Country                            string `json:"country,omitempty"`
	Carrier                            string `json:"carrier,omitempty"`
	Number                             string `json:"number,omitempty"`
	NumberType                         string `json:"number_type,omitempty"`
	MonthlyRentalRate                  string `json:"monthly_rental_rate,omitempty"`
	Application                        string `json:"application,omitempty"`
	RenewalDate                        string `json:"renewal_date,omitempty"`
	CNAMLookup                         string `json:"cnam_lookup,omitempty"`
	AddedOn                            string `json:"added_on,omitempty"`
	ResourceURI                        string `json:"resource_uri,omitempty"`
	VoiceRate                          string `json:"voice_rate,omitempty"`
	SMSRate                            string `json:"sms_rate,omitempty"`
	MMSRate                            string `json:"mms_rate,omitempty"`
	TendlcCampaignID                   string `json:"tendlc_campaign_id,omitempty"`
	TendlcRegistrationStatus           string `json:"tendlc_registration_status,omitempty"`
	TollFreeSMSVerification            string `json:"toll_free_sms_verification,omitempty"`
	TollFreeSMSVerificationID          string `json:"toll_free_sms_verification_id,omitempty"`
	TollFreeSMSVerificationOrderStatus string `json:"toll_free_sms_verification_order_status,omitempty"`
}

// CreateParams holds parameters for adding a number.
type CreateParams struct {
	Numbers    string `json:"numbers,omitempty"`
	Carrier    string `json:"carrier,omitempty"`
	Region     string `json:"region,omitempty"`
	NumberType string `json:"number_type,omitempty"`
	AppID      string `json:"app_id,omitempty"`
	Subaccount string `json:"subaccount,omitempty"`
}

// CreateResponse is the response from adding a number.
type CreateResponse = models.BaseResponse

// UpdateParams holds parameters for updating a number.
type UpdateParams struct {
	AppID      string `json:"app_id,omitempty"`
	Subaccount string `json:"subaccount,omitempty"`
	Alias      string `json:"alias,omitempty"`
	CNAMLookup string `json:"cnam_lookup,omitempty"`
}

// UpdateResponse is the response from updating a number.
type UpdateResponse = models.BaseResponse

// ListParams filters the number list.
type ListParams struct {
	NumberType                         string `url:"number_type,omitempty"`
	NumberStartsWith                   string `url:"number_startswith,omitempty"`
	Subaccount                         string `url:"subaccount,omitempty"`
	RenewalDate                        string `url:"renewal_date,omitempty"`
	RenewalDateLt                      string `url:"renewal_date__lt,omitempty"`
	RenewalDateLte                     string `url:"renewal_date__lte,omitempty"`
	RenewalDateGt                      string `url:"renewal_date__gt,omitempty"`
	RenewalDateGte                     string `url:"renewal_date__gte,omitempty"`
	Services                           string `url:"services,omitempty"`
	Alias                              string `url:"alias,omitempty"`
	Limit                              int64  `url:"limit,omitempty"`
	Offset                             int64  `url:"offset,omitempty"`
	TendlcCampaignID                   string `url:"tendlc_campaign_id,omitempty"`
	TendlcRegistrationStatus           string `url:"tendlc_registration_status,omitempty"`
	TollFreeSMSVerification            string `url:"toll_free_sms_verification,omitempty"`
	CNAMLookup                         string `url:"cnam_lookup,omitempty"`
	TollFreeSMSVerificationOrderStatus string `url:"toll_free_sms_verification_order_status,omitempty"`
}

// ListResponse wraps a paginated list of numbers.
type ListResponse struct {
	ApiID   string       `json:"api_id"`
	Meta    *models.Meta `json:"meta"`
	Objects []*Number    `json:"objects"`
}

// PhoneNumber represents an available phone number from the catalog.
type PhoneNumber struct {
	Country           string `json:"country"`
	City              string `json:"city"`
	Lata              int    `json:"lata"`
	MonthlyRentalRate string `json:"monthly_rental_rate"`
	Number            string `json:"number"`
	Type              string `json:"type"`
	Prefix            string `json:"prefix"`
	RateCenter        string `json:"rate_center"`
	Region            string `json:"region"`
	ResourceURI       string `json:"resource_uri"`
	Restriction       string `json:"restriction"`
	RestrictionText   string `json:"restriction_text"`
	SetupRate         string `json:"setup_rate"`
	SmsEnabled        bool   `json:"sms_enabled"`
	SmsRate           string `json:"sms_rate"`
	MmsEnabled        bool   `json:"mms_enabled"`
	MmsRate           string `json:"mms_rate"`
	VoiceEnabled      bool   `json:"voice_enabled"`
	VoiceRate         string `json:"voice_rate"`
	FallbackNumber    string `json:"fallback_number,omitempty"`
	HAEnabled         bool   `json:"ha_enabled,omitempty"`
}

// PhoneNumberCreateParams holds parameters for buying a phone number.
type PhoneNumberCreateParams struct {
	AppID      string `json:"app_id,omitempty"`
	CNAMLookup string `json:"cnam_lookup,omitempty"`
	HAEnable   *bool  `json:"ha_enable,omitempty"`
}

// PhoneNumberCreateResponse is the response from buying a phone number.
type PhoneNumberCreateResponse struct {
	APIID   string `json:"api_id"`
	Message string `json:"message"`
	Numbers []struct {
		Number         string `json:"number"`
		Status         string `json:"status"`
		FallbackNumber string `json:"fallback_number,omitempty"`
	} `json:"numbers"`
	Status string `json:"status"`
}

// PhoneNumberListParams filters the phone number catalog.
type PhoneNumberListParams struct {
	CountryISO string `url:"country_iso,omitempty"`
	Type       string `url:"type,omitempty"`
	Pattern    string `url:"pattern,omitempty"`
	Region     string `url:"region,omitempty"`
	Services   string `url:"services,omitempty"`
	LATA       string `url:"lata,omitempty"`
	RateCenter string `url:"rate_center,omitempty"`
	City       string `url:"city,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Offset     int    `url:"offset,omitempty"`
}

// PhoneNumberListResponse wraps a paginated list of available phone numbers.
type PhoneNumberListResponse struct {
	ApiID   string         `json:"api_id"`
	Meta    *models.Meta   `json:"meta"`
	Objects []*PhoneNumber `json:"objects"`
}
