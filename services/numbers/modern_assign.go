package numbers

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

type AssignParams struct {
	TrunkGroupID string `json:"trunk_group_id"`
}

// Assign assigns a number to trunk group.
func (s *Service) Assign(number, trunkID string) (*models.BaseResponse, error) {
	req, err := s.r.NewRequest("POST", AssignParams{TrunkGroupID: trunkID}, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "numbers/"+escNumber(number)+"/assign")
	resp := &models.BaseResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// Unassign removes a number from trunk group.
func (s *Service) Unassign(number string) error {
	req, err := s.r.NewRequest("DELETE", nil, "")
	if err != nil { return err }
	setModernPath(req, s.r.AuthID(), "numbers/"+escNumber(number)+"/assign")
	return s.r.ExecuteRequest(req, nil)
}
