package messages

// Send sends an SMS, MMS, or WhatsApp message.
func (s *Service) Send(params SendParams) (*SendResponse, error) {
	req, err := s.r.NewRequest("POST", params, "Message")
	if err != nil {
		return nil, err
	}
	resp := &SendResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
