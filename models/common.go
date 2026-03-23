// Package models contains shared request/response types used across all services.
package models

// Meta holds pagination metadata returned by list endpoints.
type Meta struct {
	Previous   *string `json:"previous"`
	Next       *string `json:"next"`
	TotalCount int64   `json:"total_count"`
	Offset     int64   `json:"offset"`
	Limit      int64   `json:"limit"`
}

// BaseListParams is the common pagination input for list operations.
type BaseListParams struct {
	Limit  int `json:"limit,omitempty"  url:"limit,omitempty"`
	Offset int `json:"offset,omitempty" url:"offset,omitempty"`
}

// BaseResponse is the minimal envelope returned by mutating operations.
type BaseResponse struct {
	ApiID   string `json:"api_id"  url:"api_id"`
	Message string `json:"message" url:"message"`
}

// BaseListResponse is the minimal envelope returned by list operations.
type BaseListResponse struct {
	ApiID string `json:"api_id" url:"api_id"`
	Meta  Meta   `json:"meta"   url:"meta"`
}

// Address is a reusable postal address used in compliance/brand/profile resources.
type Address struct {
	Street     string `json:"street"      validate:"max=100"`
	City       string `json:"city"        validate:"max=100"`
	State      string `json:"state"       validate:"max=20"`
	PostalCode string `json:"postal_code" validate:"max=10"`
	Country    string `json:"country"     validate:"max=2"`
}

// AuthorizedContact is a reusable contact used in compliance/brand/profile resources.
type AuthorizedContact struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"     validate:"max=16"`
	Email     string `json:"email,omitempty"     validate:"max=100"`
	Title     string `json:"title,omitempty"`
	Seniority string `json:"seniority,omitempty"`
}

// TCRErrorDetail holds a TCR decline reason code and message.
type TCRErrorDetail struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
