package profile

// Address holds a physical address.
type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

// AuthorizedContact holds contact information.
type AuthorizedContact struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Email     string `json:"email,omitempty"`
	Title     string `json:"title,omitempty"`
	Seniority string `json:"seniority,omitempty"`
}

// Profile represents a Vobiz compliance profile.
type Profile struct {
	ProfileUUID          string            `json:"profile_uuid,omitempty"`
	ProfileAlias         string            `json:"profile_alias,omitempty"`
	ProfileType          string            `json:"profile_type,omitempty"`
	PrimaryProfile       string            `json:"primary_profile,omitempty"`
	CustomerType         string            `json:"customer_type,omitempty"`
	EntityType           string            `json:"entity_type,omitempty"`
	CompanyName          string            `json:"company_name,omitempty"`
	Ein                  string            `json:"ein,omitempty"`
	EinIssuingCountry    string            `json:"ein_issuing_country,omitempty"`
	Address              Address           `json:"address,omitempty"`
	StockSymbol          string            `json:"stock_symbol,omitempty"`
	StockExchange        string            `json:"stock_exchange,omitempty"`
	Website              string            `json:"website,omitempty"`
	Vertical             string            `json:"vertical,omitempty"`
	AltBusinessID        string            `json:"alt_business_id,omitempty"`
	AltBusinessidType    string            `json:"alt_business_id_type,omitempty"`
	VobizSubaccount      string            `json:"vobiz_subaccount,omitempty"`
	AuthorizedContact    AuthorizedContact `json:"authorized_contact,omitempty"`
	BusinessContactEmail string            `json:"business_contact_email,omitempty"`
	CreatedAt            string            `json:"created_at,omitempty"`
}

// CreateParams holds parameters for creating a profile.
type CreateParams struct {
	ProfileAlias         string             `json:"profile_alias"`
	CustomerType         string             `json:"customer_type"`
	EntityType           string             `json:"entity_type"`
	CompanyName          string             `json:"company_name"`
	Ein                  string             `json:"ein"`
	EinIssuingCountry    string             `json:"ein_issuing_country"`
	Address              *Address           `json:"address"`
	StockSymbol          string             `json:"stock_symbol"`
	StockExchange        string             `json:"stock_exchange"`
	Website              string             `json:"website"`
	Vertical             string             `json:"vertical"`
	AltBusinessID        string             `json:"alt_business_id"`
	AltBusinessidType    string             `json:"alt_business_id_type"`
	VobizSubaccount      string             `json:"vobiz_subaccount"`
	AuthorizedContact    *AuthorizedContact `json:"authorized_contact"`
	BusinessContactEmail string             `json:"business_contact_email,omitempty"`
}

// UpdateParams holds parameters for updating a profile.
type UpdateParams struct {
	EntityType           string             `json:"entity_type"`
	CompanyName          string             `json:"company_name"`
	Address              *Address           `json:"address"`
	Website              string             `json:"website"`
	Vertical             string             `json:"vertical"`
	AuthorizedContact    *AuthorizedContact `json:"authorized_contact"`
	BusinessContactEmail string             `json:"business_contact_email,omitempty"`
	Ein                  string             `json:"ein,omitempty"`
	EinIssuingCountry    string             `json:"ein_issuing_country,omitempty"`
	AltBusinessID        string             `json:"alt_business_id,omitempty"`
	AltBusinessidType    string             `json:"alt_business_id_type,omitempty"`
}

// CreateResponse is the response from creating a profile.
type CreateResponse struct {
	ApiID       string `json:"api_id"`
	ProfileUUID string `json:"profile_uuid,omitempty"`
	Message     string `json:"message,omitempty"`
}

// GetResponse wraps a single profile.
type GetResponse struct {
	ApiID   string  `json:"api_id"`
	Profile Profile `json:"profile"`
}

// ListParams filters profiles.
type ListParams struct {
	Limit  int `url:"limit,omitempty"`
	Offset int `url:"offset,omitempty"`
}

// ListResponse wraps a paginated list of profiles.
type ListResponse struct {
	ApiID string `json:"api_id"`
	Meta  struct {
		Previous *string
		Next     *string
		Offset   int64
		Limit    int64
	} `json:"meta"`
	ProfileResponse []Profile `json:"profiles"`
}

// DeleteResponse is the response from deleting a profile.
type DeleteResponse struct {
	ApiID   string `json:"api_id"`
	Message string `json:"message,omitempty"`
}
