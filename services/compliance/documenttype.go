package compliance

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

// Get retrieves a document type by ID.
func (s *DocumentTypeService) Get(docTypeID string) (*GetDocumentTypeResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "ComplianceDocumentType/%s", docTypeID)
	if err != nil {
		return nil, err
	}
	resp := &GetDocumentTypeResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// List returns a paginated list of document types.
func (s *DocumentTypeService) List(params models.BaseListParams) (*ListDocumentTypeResponse, error) {
	req, err := s.r.NewRequest("GET", params, "ComplianceDocumentType")
	if err != nil {
		return nil, err
	}
	resp := &ListDocumentTypeResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
