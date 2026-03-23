package xml

// Stream corresponds to the <Stream> verb.
type Stream struct {
	XMLName           string `xml:"Stream"`
	Bidirectional     *bool  `xml:"bidirectional,attr,omitempty"`
	StreamTimeout     int    `xml:"streamTimeout,attr,omitempty"`
	StatusCallbackURL string `xml:"statusCallbackUrl,attr,omitempty"`
	URL               string `xml:",chardata"`
}

func NewStream(streamURL string) Stream {
	bi := true
	return Stream{
		XMLName:       "Stream",
		URL:           streamURL,
		Bidirectional: &bi,
		StreamTimeout: 600,
	}
}

func (n Stream) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) Stream(streamURL string) *Response { return r.Add(NewStream(streamURL)) }
