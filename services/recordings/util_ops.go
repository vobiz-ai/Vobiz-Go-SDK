package recordings

import (
	"fmt"
)

// Export exports recordings metadata/files.
func (s *Service) Export(params ExportParams) (*UtilityResponse, error) {
	req, err := s.r.NewRequest("POST", params, "")
	if err != nil { return nil, err }
	req.URL.Path = fmt.Sprintf("/api/v1/account/%s/export/recording", s.r.AuthID())
	resp := &UtilityResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// BulkDelete deletes recordings in bulk.
func (s *Service) BulkDelete(params BulkDeleteParams) (*UtilityResponse, error) {
	req, err := s.r.NewRequest("DELETE", params, "")
	if err != nil { return nil, err }
	req.URL.Path = fmt.Sprintf("/api/v1/account/%s/Recording/BulkDelete", s.r.AuthID())
	resp := &UtilityResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
