package compliance

import "fmt"

// Get retrieves a compliance document by ID.
func (s *DocumentService) Get(documentID string) (*GetDocumentResponse, error) {
	req, err := s.r.NewRequest("GET", nil, "ComplianceDocument/%s", documentID)
	if err != nil {
		return nil, err
	}
	resp := &GetDocumentResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// List returns a paginated list of compliance documents.
func (s *DocumentService) List(params DocumentListParams) (*ListDocumentResponse, error) {
	req, err := s.r.NewRequest("GET", params, "ComplianceDocument")
	if err != nil {
		return nil, err
	}
	resp := &ListDocumentResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// Create creates a new compliance document.
func (s *DocumentService) Create(params CreateDocumentParams) (*GetDocumentResponse, error) {
	uploadURL := fmt.Sprintf("%s/api/v1/Account/%s/ComplianceDocument", s.baseURL, s.authID)
	reqParams := structToStringMap(params, "file")
	req, err := newFileUploadRequest(uploadURL, reqParams, "file", params.File)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(s.authID, s.authToken)
	resp := &GetDocumentResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// Update modifies a compliance document.
func (s *DocumentService) Update(params UpdateDocumentParams) (*UpdateDocumentResponse, error) {
	uploadURL := fmt.Sprintf("%s/api/v1/Account/%s/ComplianceDocument/%s", s.baseURL, s.authID, params.ComplianceDocumentID)
	reqParams := structToStringMap(params, "file")
	req, err := newFileUploadRequest(uploadURL, reqParams, "file", params.File)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(s.authID, s.authToken)
	resp := &UpdateDocumentResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// Delete removes a compliance document.
func (s *DocumentService) Delete(documentID string) error {
	req, err := s.r.NewRequest("DELETE", nil, "ComplianceDocument/%s", documentID)
	if err != nil {
		return err
	}
	return s.r.ExecuteRequest(req, nil)
}
