package xml

// Redirect corresponds to the <Redirect> verb.
type Redirect struct {
	XMLName string `xml:"Redirect"`
	URL     string `xml:",chardata"`
}

func NewRedirect(url string) Redirect {
	return Redirect{XMLName: "Redirect", URL: url}
}

func (n Redirect) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) Redirect(url string) *Response { return r.Add(NewRedirect(url)) }
