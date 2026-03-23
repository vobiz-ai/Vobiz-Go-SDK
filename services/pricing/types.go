package pricing

// Pricing holds pricing information for a country.
type Pricing struct {
	APIID       string `json:"api_id"`
	Country     string `json:"country"`
	CountryCode int    `json:"country_code"`
	CountryISO  string `json:"country_iso"`
	Message     struct {
		Inbound struct {
			Rate string `json:"rate"`
		} `json:"inbound"`
		Outbound struct {
			Rate string `json:"rate"`
		} `json:"outbound"`
		OutboundNetworksList []struct {
			GroupName string `json:"group_name"`
			Rate      string `json:"rate"`
		} `json:"outbound_networks_list"`
	} `json:"message"`
	PhoneNumbers struct {
		Local struct {
			Rate string `json:"rate"`
		} `json:"local"`
		Tollfree struct {
			Rate string `json:"rate"`
		} `json:"tollfree"`
	} `json:"phone_numbers"`
	Voice struct {
		Inbound struct {
			IP struct {
				Rate string `json:"rate"`
			} `json:"ip"`
			Local struct {
				Rate string `json:"rate"`
			} `json:"local"`
			Tollfree struct {
				Rate string `json:"rate"`
			} `json:"tollfree"`
		} `json:"inbound"`
		Outbound struct {
			IP struct {
				Rate string `json:"rate"`
			} `json:"ip"`
			Local struct {
				Rate string `json:"rate"`
			} `json:"local"`
			Rates []struct {
				OriginationPrefix []string `json:"origination_prefix"`
				Prefix            []string `json:"prefix"`
				Rate              string   `json:"rate"`
				VoiceNetworkGroup string   `json:"voice_network_group"`
			} `json:"rates"`
			Tollfree struct {
				Rate string `json:"rate"`
			} `json:"tollfree"`
		} `json:"outbound"`
	} `json:"voice"`
}
