package brand

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

// TCRErrorDetail holds a TCR error code and message.
type TCRErrorDetail struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// Brand represents a 10DLC brand.
type Brand struct {
	BrandAlias         string            `json:"brand_alias,omitempty"`
	EntityType         string            `json:"entity_type,omitempty"`
	BrandID            string            `json:"brand_id,omitempty"`
	ProfileUUID        string            `json:"profile_uuid,omitempty"`
	FirstName          string            `json:"first_name,omitempty"`
	LastName           string            `json:"last_name,omitempty"`
	Name               string            `json:"name,omitempty"`
	CompanyName        string            `json:"company_name,omitempty"`
	BrandType          string            `json:"brand_type,omitempty"`
	Ein                string            `json:"ein,omitempty"`
	EinIssuingCountry  string            `json:"ein_issuing_country,omitempty"`
	StockSymbol        string            `json:"stock_symbol,omitempty"`
	StockExchange      string            `json:"stock_exchange,omitempty"`
	Website            string            `json:"website,omitempty"`
	Vertical           string            `json:"vertical,omitempty"`
	AltBusinessID      string            `json:"alt_business_id,omitempty"`
	AltBusinessidType  string            `json:"alt_business_id_type,omitempty"`
	RegistrationStatus string            `json:"registration_status,omitempty"`
	VettingStatus      string            `json:"vetting_status,omitempty"`
	VettingScore       int64             `json:"vetting_score,omitempty"`
	Address            Address           `json:"address,omitempty"`
	AuthorizedContact  AuthorizedContact `json:"authorized_contact,omitempty"`
	DeclinedReasons    []TCRErrorDetail  `json:"declined_reasons,omitempty"`
	CreatedAt          string            `json:"created_at,omitempty"`
}

// CreateParams holds parameters for creating a brand.
type CreateParams struct {
	BrandAlias       string  `json:"brand_alias"`
	Type             string  `json:"brand_type"`
	ProfileUUID      string  `json:"profile_uuid"`
	SecondaryVetting *string `json:"secondary_vetting,omitempty"`
	URL              string  `json:"url,omitempty"`
	Method           string  `json:"method,omitempty"`
}

// CreateResponse is the response from creating a brand.
type CreateResponse struct {
	ApiID   string `json:"api_id,omitempty"`
	BrandID string `json:"brand_id,omitempty"`
	Message string `json:"message,omitempty"`
}

// GetResponse wraps a single brand.
type GetResponse struct {
	ApiID string `json:"api_id,omitempty"`
	Brand Brand  `json:"brand,omitempty"`
}

// ListParams filters brands.
type ListParams struct {
	Type   *string `url:"type,omitempty"`
	Status *string `url:"status,omitempty"`
	Limit  int     `url:"limit,omitempty"`
	Offset int     `url:"offset,omitempty"`
}

// ListResponse wraps a paginated list of brands.
type ListResponse struct {
	ApiID string `json:"api_id,omitempty"`
	Meta  struct {
		Previous   *string
		Next       *string
		Offset     int64
		Limit      int64
		TotalCount int64 `json:"total_count"`
	} `json:"meta"`
	BrandResponse []Brand `json:"brands,omitempty"`
}

// DeleteResponse is the response from deleting a brand.
type DeleteResponse struct {
	ApiID   string `json:"api_id,omitempty"`
	BrandID string `json:"brand_id,omitempty"`
	Message string `json:"message,omitempty"`
}

// Usecase holds a brand use case.
type Usecase struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Details string `json:"details"`
}

// UsecaseResponse wraps brand use cases.
type UsecaseResponse struct {
	ApiID    string    `json:"api_id,omitempty"`
	Usecases []Usecase `json:"use_cases"`
	BrandID  string    `json:"brand_id"`
}
