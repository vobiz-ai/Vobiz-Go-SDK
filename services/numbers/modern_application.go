package numbers

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

// LinkApplication links a number to an application.
func (s *Service) LinkApplication(number string, params LinkApplicationParams) (*models.BaseResponse, error) {
	req, err := s.r.NewRequest("POST", params, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "numbers/"+escNumber(number)+"/application")
	resp := &models.BaseResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// UnlinkApplication unlinks a number from its application.
func (s *Service) UnlinkApplication(number string) error {
	req, err := s.r.NewRequest("DELETE", nil, "")
	if err != nil { return err }
	setModernPath(req, s.r.AuthID(), "numbers/"+escNumber(number)+"/application")
	return s.r.ExecuteRequest(req, nil)
}
