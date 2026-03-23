package numbers

import "github.com/vobiz/all-vobiz-sdk/go-sdk/models"

// PurchaseFromInventory buys a number from inventory.
func (s *Service) PurchaseFromInventory(params PurchaseFromInventoryParams) (*models.BaseResponse, error) {
	req, err := s.r.NewRequest("POST", params, "")
	if err != nil { return nil, err }
	setModernPath(req, s.r.AuthID(), "numbers/purchase-from-inventory")
	resp := &models.BaseResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
