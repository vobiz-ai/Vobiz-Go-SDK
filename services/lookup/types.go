package lookup

// Country holds country information.
type Country struct {
	Name string `json:"name"`
	ISO2 string `json:"iso2"`
	ISO3 string `json:"iso3"`
}

// NumberFormat holds number format variants.
type NumberFormat struct {
	E164          string `json:"e164"`
	National      string `json:"national"`
	International string `json:"international"`
	RFC3966       string `json:"rfc3966"`
}

// Carrier holds carrier information.
type Carrier struct {
	MobileCountryCode string `json:"mobile_country_code"`
	MobileNetworkCode string `json:"mobile_network_code"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Ported            string `json:"ported"`
}

// Response is the success response from the Lookup API.
type Response struct {
	ApiID       string        `json:"api_id"`
	PhoneNumber string        `json:"phone_number"`
	Country     *Country      `json:"country"`
	Format      *NumberFormat `json:"format"`
	Carrier     *Carrier      `json:"carrier"`
	ResourceURI string        `json:"resource_uri"`
}

// Params holds lookup query parameters.
type Params struct {
	Type string `url:"type"`
}

// Error is the error response from the Lookup API.
type Error struct {
	respBody  string
	ApiID     string `json:"api_id"`
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func (e *Error) Error() string { return e.respBody }
