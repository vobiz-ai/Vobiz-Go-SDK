package compliance

import (
	"time"

	"github.com/vobiz/all-vobiz-sdk/go-sdk/models"
)

// ---- Document types --------------------------------------------------------

// DocumentMeta holds document metadata fields.
type DocumentMeta struct {
	LastName                     string `json:"last_name,omitempty"`
	FirstName                    string `json:"first_name,omitempty"`
	DateOfBirth                  string `json:"date_of_birth,omitempty"`
	AddressLine1                 string `json:"address_line_1,omitempty"`
	AddressLine2                 string `json:"address_line_2,omitempty"`
	City                         string `json:"city,omitempty"`
	Country                      string `json:"country,omitempty"`
	PostalCode                   string `json:"postal_code,omitempty"`
	UniqueIdentificationNumber   string `json:"unique_identification_number,omitempty"`
	Nationality                  string `json:"nationality,omitempty"`
	PlaceOfBirth                 string `json:"place_of_birth,omitempty"`
	DateOfIssue                  string `json:"date_of_issue,omitempty"`
	DateOfExpiration             string `json:"date_of_expiration,omitempty"`
	TypeOfUtility                string `json:"type_of_utility,omitempty"`
	BillingId                    string `json:"billing_id,omitempty"`
	BillingDate                  string `json:"billing_date,omitempty"`
	BusinessName                 string `json:"business_name,omitempty"`
	TypeOfId                     string `json:"type_of_id,omitempty"`
	SupportPhoneNumber           string `json:"support_phone_number,omitempty"`
	SupportEmail                 string `json:"support_email,omitempty"`
	AuthorizedRepresentativeName string `json:"authorized_representative_name,omitempty"`
	BillDate                     string `json:"bill_date,omitempty"`
	BillId                       string `json:"bill_id,omitempty"`
	UseCaseDescription           string `json:"use_case_description,omitempty"`
}

// GetDocumentResponse is the response from fetching a compliance document.
type GetDocumentResponse struct {
	APIID           string       `json:"api_id"`
	DocumentID      string       `json:"document_id"`
	EndUserID       string       `json:"end_user_id"`
	DocumentTypeID  string       `json:"document_type_id"`
	Alias           string       `json:"alias"`
	FileName        string       `json:"file_name,omitempty"`
	MetaInformation DocumentMeta `json:"meta_information"`
	CreatedAt       string       `json:"created_at"`
}

// DocumentListParams filters compliance documents.
type DocumentListParams struct {
	Limit          int    `url:"limit,omitempty"`
	Offset         int    `url:"offset,omitempty"`
	EndUserID      string `url:"end_user_id,omitempty"`
	DocumentTypeID string `url:"document_type_id,omitempty"`
	Alias          string `url:"alias,omitempty"`
}

// ListDocumentResponse wraps a paginated list of compliance documents.
type ListDocumentResponse struct {
	APIID   string `json:"api_id"`
	Meta    struct {
		Limit      int         `json:"limit"`
		Next       string      `json:"next"`
		Offset     int         `json:"offset"`
		Previous   interface{} `json:"previous"`
		TotalCount int         `json:"total_count"`
	} `json:"meta"`
	Objects []struct {
		CreatedAt            time.Time    `json:"created_at"`
		ComplianceDocumentID string       `json:"compliance_document_id"`
		Alias                string       `json:"alias"`
		MetaInformation      DocumentMeta `json:"meta_information"`
		File                 string       `json:"file,omitempty"`
		EndUserID            string       `json:"end_user_id"`
		DocumentTypeID       string       `json:"document_type_id"`
	} `json:"objects"`
}

// CreateDocumentParams holds parameters for creating a compliance document.
type CreateDocumentParams struct {
	File           string `json:"file,omitempty"`
	EndUserID      string `json:"end_user_id,omitempty"`
	DocumentTypeID string `json:"document_type_id,omitempty"`
	Alias          string `json:"alias,omitempty"`
	DocumentMeta
}

// UpdateDocumentParams holds parameters for updating a compliance document.
type UpdateDocumentParams struct {
	ComplianceDocumentID string `json:"compliance_document_id"`
	CreateDocumentParams
}

// UpdateDocumentResponse is the response from updating a compliance document.
type UpdateDocumentResponse = models.BaseResponse

// ---- DocumentType types ----------------------------------------------------

// DocumentTypeField holds a field definition for a document type.
type DocumentTypeField struct {
	FieldName    string   `json:"field_name"`
	FieldType    string   `json:"field_type"`
	FriendlyName string   `json:"friendly_name"`
	HelpText     string   `json:"help_text,omitempty"`
	MaxLength    int      `json:"max_length,omitempty"`
	MinLength    int      `json:"min_length,omitempty"`
	Format       string   `json:"format,omitempty"`
	Enums        []string `json:"enums,omitempty"`
}

// GetDocumentTypeResponse is the response from fetching a document type.
type GetDocumentTypeResponse struct {
	APIID          string              `json:"api_id"`
	CreatedAt      time.Time           `json:"created_at"`
	Description    string              `json:"description"`
	DocumentName   string              `json:"document_name"`
	DocumentTypeID string              `json:"document_type_id"`
	Information    []DocumentTypeField `json:"information"`
	ProofRequired  interface{}         `json:"proof_required"`
}

// ListDocumentTypeResponse wraps a paginated list of document types.
type ListDocumentTypeResponse struct {
	APIID   string `json:"api_id"`
	Meta    struct {
		Limit      int         `json:"limit"`
		Next       interface{} `json:"next"`
		Offset     int         `json:"offset"`
		Previous   interface{} `json:"previous"`
		TotalCount int         `json:"total_count"`
	} `json:"meta"`
	Objects []struct {
		CreatedAt      time.Time           `json:"created_at"`
		Description    string              `json:"description"`
		DocumentName   string              `json:"document_name"`
		DocumentTypeID string              `json:"document_type_id"`
		Information    []DocumentTypeField `json:"information"`
		ProofRequired  interface{}         `json:"proof_required"`
	} `json:"objects"`
}

// ---- Requirement types -----------------------------------------------------

// AcceptableDocument holds a document type reference.
type AcceptableDocument struct {
	DocumentTypeID   string `json:"document_type_id"`
	DocumentTypeName string `json:"document_type_name"`
}

// AcceptableDocumentType holds a document type category.
type AcceptableDocumentType struct {
	Name                string               `json:"name"`
	Scope               string               `json:"scope"`
	AcceptableDocuments []AcceptableDocument `json:"acceptable_documents"`
}

// GetRequirementResponse is the response from fetching a compliance requirement.
type GetRequirementResponse struct {
	APIID                   string                   `json:"api_id"`
	ComplianceRequirementID string                   `json:"compliance_requirement_id"`
	CountryIso2             string                   `json:"country_iso2"`
	NumberType              string                   `json:"number_type"`
	EndUserType             string                   `json:"end_user_type"`
	AcceptableDocumentTypes []AcceptableDocumentType `json:"acceptable_document_types"`
}

// ListRequirementParams filters compliance requirements.
type ListRequirementParams struct {
	CountryIso2 string `url:"country_iso2,omitempty"`
	NumberType  string `url:"number_type,omitempty"`
	EndUserType string `url:"end_user_type,omitempty"`
	PhoneNumber string `url:"phone_number,omitempty"`
}

// ---- Application types -----------------------------------------------------

// ApplicationDocument holds a document reference in an application.
type ApplicationDocument struct {
	DocumentID       string `json:"document_id"`
	Name             string `json:"name"`
	DocumentName     string `json:"document_name,omitempty"`
	DocumentTypeID   string `json:"document_type_id,omitempty"`
	DocumentTypeName string `json:"document_type_name,omitempty"`
	Scope            string `json:"scope,omitempty"`
}

// ApplicationResponse is the response from fetching/creating a compliance application.
type ApplicationResponse struct {
	APIID                   string                `json:"api_id"`
	CreatedAt               time.Time             `json:"created_at"`
	ComplianceApplicationID string                `json:"compliance_application_id"`
	Alias                   string                `json:"alias"`
	Status                  string                `json:"status"`
	EndUserType             string                `json:"end_user_type"`
	CountryIso2             string                `json:"country_iso2"`
	NumberType              string                `json:"number_type"`
	EndUserID               string                `json:"end_user_id"`
	ComplianceRequirementID string                `json:"compliance_requirement_id"`
	Documents               []ApplicationDocument `json:"documents"`
	RejectionReason         string                `json:"rejection_reason,omitempty"`
}

// CreateApplicationParams holds parameters for creating a compliance application.
type CreateApplicationParams struct {
	ComplianceRequirementId string   `json:"compliance_requirement_id,omitempty"`
	EndUserId               string   `json:"end_user_id,omitempty"`
	DocumentIds             []string `json:"document_ids,omitempty"`
	Alias                   string   `json:"alias,omitempty"`
	EndUserType             string   `json:"end_user_type,omitempty"`
	CountryIso2             string   `json:"country_iso2,omitempty"`
	NumberType              string   `json:"number_type,omitempty"`
}

// UpdateApplicationParams holds parameters for updating a compliance application.
type UpdateApplicationParams struct {
	ComplianceApplicationId string   `json:"compliance_application_id"`
	DocumentIds             []string `json:"document_ids"`
}

// UpdateApplicationResponse is the response from updating a compliance application.
type UpdateApplicationResponse = models.BaseResponse

// ApplicationListParams filters compliance applications.
type ApplicationListParams struct {
	Limit       int    `url:"limit,omitempty"`
	Offset      int    `url:"offset,omitempty"`
	EndUserID   string `url:"end_user_id,omitempty"`
	Country     string `url:"country,omitempty"`
	NumberType  string `url:"number_type,omitempty"`
	EndUserType string `url:"end_user_type,omitempty"`
	Alias       string `url:"alias,omitempty"`
	Status      string `url:"status,omitempty"`
}

// ListApplicationResponse wraps a paginated list of compliance applications.
type ListApplicationResponse struct {
	APIID   string `json:"api_id"`
	Meta    struct {
		Limit      int         `json:"limit"`
		Next       string      `json:"next"`
		Offset     int         `json:"offset"`
		Previous   interface{} `json:"previous"`
		TotalCount int         `json:"total_count"`
	} `json:"meta"`
	Objects []ApplicationResponse `json:"objects"`
}
