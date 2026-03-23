package messages

// Get retrieves a message by UUID.
func (s *Service) Get(messageUUID string) (*Message, error) {
	req, err := s.r.NewRequest("GET", nil, "Message/%s", messageUUID)
	if err != nil {
		return nil, err
	}
	resp := &Message{}
	return resp, s.r.ExecuteRequest(req, resp)
}
