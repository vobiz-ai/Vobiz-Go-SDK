package campaign

// CampaignAttributes holds boolean campaign attributes.
type CampaignAttributes struct {
	EmbeddedLink       bool `json:"embedded_link"`
	EmbeddedPhone      bool `json:"embedded_phone"`
	AgeGated           bool `json:"age_gated"`
	DirectLending      bool `json:"direct_lending"`
	SubscriberOptin    bool `json:"subscriber_optin"`
	SubscriberOptout   bool `json:"subscriber_optout"`
	SubscriberHelp     bool `json:"subscriber_help"`
	AffiliateMarketing bool `json:"affiliate_marketing"`
}

// OperatorDetail holds MNO-specific metadata.
type OperatorDetail struct {
	BrandTier string `json:"brand_tier,omitempty"`
	TPM       int    `json:"tpm,omitempty"`
}

// MnoMetadata holds per-operator metadata.
type MnoMetadata struct {
	ATandT          OperatorDetail `json:"AT&T,omitempty"`
	TMobile         OperatorDetail `json:"T-Mobile,omitempty"`
	VerizonWireless OperatorDetail `json:"Verizon Wireless,omitempty"`
	USCellular      OperatorDetail `json:"US Cellular,omitempty"`
}

// Campaign represents a 10DLC campaign.
type Campaign struct {
	BrandID                string             `json:"brand_id,omitempty"`
	CampaignID             string             `json:"campaign_id,omitempty"`
	MnoMetadata            MnoMetadata        `json:"mno_metadata,omitempty"`
	ResellerID             string             `json:"reseller_id,omitempty"`
	Usecase                string             `json:"usecase,omitempty"`
	SubUsecase             string             `json:"sub_usecase,omitempty"`
	RegistrationStatus     string             `json:"registration_status,omitempty"`
	MessageFlow            string             `json:"message_flow,omitempty"`
	HelpMessage            string             `json:"help_message,omitempty"`
	OptinKeywords          string             `json:"optin_keywords,omitempty"`
	OptinMessage           string             `json:"optin_message,omitempty"`
	OptoutKeywords         string             `json:"optout_keywords,omitempty"`
	OptoutMessage          string             `json:"optout_message,omitempty"`
	HelpKeywords           string             `json:"help_keywords,omitempty"`
	SampleMessage1         string             `json:"sample1,omitempty"`
	SampleMessage2         string             `json:"sample2,omitempty"`
	SampleMessage3         string             `json:"sample3,omitempty"`
	SampleMessage4         string             `json:"sample4,omitempty"`
	SampleMessage5         string             `json:"sample5,omitempty"`
	TermsAndConditionsLink string             `json:"terms_and_conditions_link,omitempty"`
	PrivacyPolicyLink      string             `json:"privacy_policy_link,omitempty"`
	CampaignDescription    string             `json:"description,omitempty"`
	CampaignAttributes     CampaignAttributes `json:"campaign_attributes,omitempty"`
	CreatedAt              string             `json:"created_at,omitempty"`
	CampaignSource         string             `json:"campaign_source,omitempty"`
	ErrorCode              string             `json:"error_code,omitempty"`
	ErrorReason            string             `json:"error_reason,omitempty"`
	Vertical               string             `json:"vertical,omitempty"`
	CampaignAlias          string             `json:"campaign_alias,omitempty"`
	ErrorDescription       string             `json:"error_description,omitempty"`
}

// CreateParams holds parameters for creating a campaign.
type CreateParams struct {
	BrandID                string    `json:"brand_id"`
	CampaignAlias          *string   `json:"campaign_alias,omitempty"`
	Vertical               string    `json:"vertical"`
	Usecase                string    `json:"usecase"`
	ResellerID             string    `json:"reseller_id,omitempty"`
	SubUsecases            *[]string `json:"sub_usecases"`
	Description            string    `json:"description,omitempty"`
	EmbeddedLink           *bool     `json:"embedded_link,omitempty"`
	EmbeddedPhone          *bool     `json:"embedded_phone,omitempty"`
	AgeGated               *bool     `json:"age_gated,omitempty"`
	DirectLending          *bool     `json:"direct_lending,omitempty"`
	SubscriberOptin        bool      `json:"subscriber_optin"`
	SubscriberOptout       bool      `json:"subscriber_optout"`
	SubscriberHelp         bool      `json:"subscriber_help"`
	AffiliateMarketing     bool      `json:"affiliate_marketing"`
	Sample1                *string   `json:"sample1"`
	Sample2                *string   `json:"sample2,omitempty"`
	Sample3                *string   `json:"sample3,omitempty"`
	Sample4                *string   `json:"sample4,omitempty"`
	Sample5                *string   `json:"sample5,omitempty"`
	URL                    string    `json:"url,omitempty"`
	Method                 string    `json:"method,omitempty"`
	MessageFlow            string    `json:"message_flow,omitempty"`
	HelpMessage            string    `json:"help_message,omitempty"`
	OptinKeywords          string    `json:"optin_keywords,omitempty"`
	OptinMessage           string    `json:"optin_message,omitempty"`
	OptoutKeywords         string    `json:"optout_keywords,omitempty"`
	OptoutMessage          string    `json:"optout_message,omitempty"`
	HelpKeywords           string    `json:"help_keywords,omitempty"`
	TermsAndConditionsLink string    `json:"terms_and_conditions_link,omitempty"`
	PrivacyPolicyLink      string    `json:"privacy_policy_link,omitempty"`
}

// ImportParams holds parameters for importing a campaign.
type ImportParams struct {
	CampaignID    string `json:"campaign_id,omitempty"`
	CampaignAlias string `json:"campaign_alias,omitempty"`
	URL           string `json:"url,omitempty"`
	Method        string `json:"method,omitempty"`
}

// UpdateParams holds parameters for updating a campaign.
type UpdateParams struct {
	ResellerID             string `json:"reseller_id,omitempty"`
	Description            string `json:"description,omitempty"`
	Sample1                string `json:"sample1"`
	Sample2                string `json:"sample2,omitempty"`
	Sample3                string `json:"sample3,omitempty"`
	Sample4                string `json:"sample4,omitempty"`
	Sample5                string `json:"sample5,omitempty"`
	MessageFlow            string `json:"message_flow,omitempty"`
	HelpMessage            string `json:"help_message,omitempty"`
	OptinKeywords          string `json:"optin_keywords,omitempty"`
	OptinMessage           string `json:"optin_message,omitempty"`
	OptoutKeywords         string `json:"optout_keywords,omitempty"`
	OptoutMessage          string `json:"optout_message,omitempty"`
	HelpKeywords           string `json:"help_keywords,omitempty"`
	TermsAndConditionsLink string `json:"terms_and_conditions_link,omitempty"`
	PrivacyPolicyLink      string `json:"privacy_policy_link,omitempty"`
}

// ListParams filters campaigns.
type ListParams struct {
	BrandID            *string `url:"brand_id,omitempty"`
	Usecase            *string `url:"usecase,omitempty"`
	RegistrationStatus *string `url:"registration_status,omitempty"`
	CampaignSource     *string `url:"campaign_source,omitempty"`
	Limit              int     `url:"limit,omitempty"`
	Offset             int     `url:"offset,omitempty"`
}

// CreateResponse is the response from creating/importing a campaign.
type CreateResponse struct {
	ApiID      string `json:"api_id,omitempty"`
	CampaignID string `json:"campaign_id"`
	Message    string `json:"message,omitempty"`
}

// GetResponse wraps a single campaign.
type GetResponse struct {
	ApiID    string   `json:"api_id"`
	Campaign Campaign `json:"campaign"`
}

// ListResponse wraps a paginated list of campaigns.
type ListResponse struct {
	ApiID string `json:"api_id,omitempty"`
	Meta  struct {
		Previous   *string
		Next       *string
		Offset     int64
		Limit      int64
		TotalCount int64 `json:"total_count"`
	} `json:"meta"`
	CampaignResponse []Campaign `json:"campaigns,omitempty"`
}

// DeleteResponse is the response from deleting a campaign.
type DeleteResponse struct {
	ApiID      string `json:"api_id"`
	CampaignID string `json:"campaign_id,omitempty"`
	Message    string `json:"message,omitempty"`
}

// NumberLinkParams holds parameters for linking numbers to a campaign.
type NumberLinkParams struct {
	Numbers []string `json:"numbers,omitempty"`
	URL     string   `json:"url,omitempty"`
	Method  string   `json:"method,omitempty"`
}

// NumberUnlinkParams holds parameters for unlinking a number.
type NumberUnlinkParams struct {
	URL    string `json:"url,omitempty"`
	Method string `json:"method,omitempty"`
}

// NumbersGetParams filters campaign numbers.
type NumbersGetParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// NumberLinkUnlinkResponse is the response from linking/unlinking numbers.
type NumberLinkUnlinkResponse struct {
	ApiID   string `json:"api_id"`
	Message string `json:"message,omitempty"`
}

// CampaignNumber holds a number and its status.
type CampaignNumber struct {
	Number string `json:"number"`
	Status string `json:"status"`
}

// NumberGetResponse wraps campaign numbers.
type NumberGetResponse struct {
	ApiID                 string           `json:"api_id"`
	CampaignID            string           `json:"campaign_id"`
	CampaignAlias         string           `json:"campaign_alias"`
	CampaignUseCase       string           `json:"usecase"`
	CampaignNumbers       []CampaignNumber `json:"phone_numbers"`
	CampaignNumberSummary map[string]int   `json:"phone_numbers_summary,omitempty"`
	NumberPoolLimit       int              `json:"number_pool_limit,omitempty"`
}
