package messages

// Message represents a sent or received message record.
type Message struct {
	ApiID                           string `json:"api_id,omitempty"`
	ToNumber                        string `json:"to_number,omitempty"`
	FromNumber                      string `json:"from_number,omitempty"`
	CloudRate                       string `json:"cloud_rate,omitempty"`
	MessageType                     string `json:"message_type,omitempty"`
	ResourceURI                     string `json:"resource_uri,omitempty"`
	CarrierRate                     string `json:"carrier_rate,omitempty"`
	MessageDirection                string `json:"message_direction,omitempty"`
	MessageState                    string `json:"message_state,omitempty"`
	TotalAmount                     string `json:"total_amount,omitempty"`
	MessageUUID                     string `json:"message_uuid,omitempty"`
	MessageTime                     string `json:"message_time,omitempty"`
	ErrorCode                       string `json:"error_code,omitempty"`
	ErrorMessage                    string `json:"error_message,omitempty"`
	PowerpackID                     string `json:"powerpack_id,omitempty"`
	RequesterIP                     string `json:"requester_ip,omitempty"`
	IsDomestic                      *bool  `json:"is_domestic,omitempty"`
	ReplacedSender                  string `json:"replaced_sender,omitempty"`
	TendlcCampaignID                string `json:"tendlc_campaign_id"`
	TendlcRegistrationStatus        string `json:"tendlc_registration_status"`
	DestinationCountryISO2          string `json:"destination_country_iso2"`
	ConversationID                  string `json:"conversation_id"`
	ConversationOrigin              string `json:"conversation_origin"`
	ConversationExpirationTimestamp string `json:"conversation_expiration_timestamp"`
	DLTEntityID                     string `json:"dlt_entity_id"`
	DLTTemplateID                   string `json:"dlt_template_id"`
	DLTTemplateCategory             string `json:"dlt_template_category"`
	DestinationNetwork              string `json:"destination_network"`
	CarrierFeesRate                 string `json:"carrier_fees_rate"`
	CarrierFees                     string `json:"carrier_fees"`
	Log                             string `json:"log"`
	MessageSentTime                 string `json:"message_sent_time"`
	MessageUpdatedTime              string `json:"message_updated_time"`
}

// Template is a WhatsApp message template.
type Template struct {
	Name       string               `json:"name"       validate:"required"`
	Language   *TemplateLanguage    `json:"language"   validate:"required"`
	Components []*TemplateComponent `json:"components,omitempty"`
}

// TemplateLanguage specifies the template language/policy.
type TemplateLanguage struct {
	Policy string `json:"policy" validate:"required"`
	Code   string `json:"code"   validate:"required"`
}

// TemplateComponent is a component within a WhatsApp template.
type TemplateComponent struct {
	Type       string               `json:"type"                 validate:"required"`
	SubType    string               `json:"sub_type,omitempty"`
	Index      string               `json:"index,omitempty"`
	Parameters []*TemplateParameter `json:"parameters,omitempty" validate:"required"`
}

// TemplateParameter is a parameter within a template component.
type TemplateParameter struct {
	Type     string            `json:"type"               validate:"required"`
	Text     string            `json:"text,omitempty"`
	Currency *TemplateCurrency `json:"currency,omitempty"`
	DateTime *TemplateDateTime `json:"date_time,omitempty"`
}

// TemplateCurrency holds currency info for a template parameter.
type TemplateCurrency struct {
	FallbackValue string `json:"fallback_value" validate:"required"`
	Code          string `json:"code"           validate:"required"`
	Amount1000    int64  `json:"amount_1000"    validate:"required"`
}

// TemplateDateTime holds date/time info for a template parameter.
type TemplateDateTime struct {
	FallbackValue string `json:"fallback_value" validate:"required"`
}

// Interactive is a WhatsApp interactive message payload.
type Interactive struct {
	Type   string      `json:"type"   validate:"required,oneof=button list product product_list"`
	Header interface{} `json:"header,omitempty"`
	Body   interface{} `json:"body,omitempty"`
	Footer interface{} `json:"footer,omitempty"`
	Action interface{} `json:"action" validate:"required"`
}

// Location is a WhatsApp location message payload.
type Location struct {
	Longitude float64 `json:"longitude" validate:"required"`
	Latitude  float64 `json:"latitude"  validate:"required"`
	Name      string  `json:"name,omitempty"`
	Address   string  `json:"address,omitempty"`
}

// SendParams are the parameters for sending a message.
type SendParams struct {
	Src                 string       `json:"src,omitempty"`
	Dst                 string       `json:"dst,omitempty"`
	Text                string       `json:"text,omitempty"`
	Type                string       `json:"type,omitempty"`
	URL                 string       `json:"url,omitempty"`
	Method              string       `json:"method,omitempty"`
	Trackable           bool         `json:"trackable,omitempty"`
	Log                 interface{}  `json:"log,omitempty"`
	MediaUrls           []string     `json:"media_urls,omitempty"`
	MediaIds            []string     `json:"media_ids,omitempty"`
	PowerpackUUID       string       `json:"powerpack_uuid,omitempty"`
	MessageExpiry       int          `json:"message_expiry,omitempty"`
	Template            *Template    `json:"template,omitempty"`
	Interactive         *Interactive `json:"interactive,omitempty"`
	Location            *Location    `json:"location,omitempty"`
	DLTEntityID         string       `json:"dlt_entity_id,omitempty"`
	DLTTemplateID       string       `json:"dlt_template_id,omitempty"`
	DLTTemplateCategory string       `json:"dlt_template_category,omitempty"`
}

// SendResponse is returned after sending a message.
type SendResponse struct {
	Message     string   `json:"message"`
	ApiID       string   `json:"api_id"`
	MessageUUID []string `json:"message_uuid"`
	Error       string   `json:"error"`
}

// ListParams filters the message list endpoint.
type ListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// ListResponse wraps a paginated list of messages.
type ListResponse struct {
	ApiID string `json:"api_id"`
	Meta  struct {
		Previous *string
		Next     *string
		Offset   int64
		Limit    int64
	} `json:"meta"`
	Objects []Message `json:"objects"`
}

// MMSMedia represents a media attachment for an MMS message.
type MMSMedia struct {
	ApiID       string `json:"api_id,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	MediaID     string `json:"media_id,omitempty"`
	MediaURL    string `json:"media_url,omitempty"`
	MessageUUID string `json:"message_uuid,omitempty"`
	Size        int64  `json:"size,omitempty"`
}

// MediaDeleteResponse is returned after deleting media.
type MediaDeleteResponse struct {
	Error string `json:"error,omitempty"`
}
