package messages

// GetMedia retrieves a media attachment for an MMS message.
func (s *Service) GetMedia(messageUUID, mediaID string) (*MMSMedia, error) {
	req, err := s.r.NewRequest("GET", nil, "Message/%s/Media/%s", messageUUID, mediaID)
	if err != nil {
		return nil, err
	}
	resp := &MMSMedia{}
	return resp, s.r.ExecuteRequest(req, resp)
}

// ListMedia lists all media attachments for an MMS message.
func (s *Service) ListMedia(messageUUID string) ([]MMSMedia, error) {
	req, err := s.r.NewRequest("GET", nil, "Message/%s/Media", messageUUID)
	if err != nil {
		return nil, err
	}
	var resp []MMSMedia
	return resp, s.r.ExecuteRequest(req, &resp)
}

// DeleteMedia deletes a media attachment.
func (s *Service) DeleteMedia(messageUUID, mediaID string) (*MediaDeleteResponse, error) {
	req, err := s.r.NewRequest("DELETE", nil, "Message/%s/Media/%s", messageUUID, mediaID)
	if err != nil {
		return nil, err
	}
	resp := &MediaDeleteResponse{}
	return resp, s.r.ExecuteRequest(req, resp)
}
