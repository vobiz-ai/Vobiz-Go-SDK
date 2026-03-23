package xml

// Play corresponds to the <Play> verb.
type Play struct {
	XMLName string `xml:"Play"`
	URL     string `xml:",chardata"`
}

func NewPlay(audioURL string) Play {
	return Play{XMLName: "Play", URL: audioURL}
}

func (n Play) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) Play(audioURL string) *Response { return r.Add(NewPlay(audioURL)) }
