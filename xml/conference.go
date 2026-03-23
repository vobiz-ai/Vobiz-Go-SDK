package xml

// Conference corresponds to the <Conference> verb.
type Conference struct {
	XMLName string `xml:"Conference"`
	Action  string `xml:"action,attr,omitempty"`
	Method  string `xml:"method,attr,omitempty"`
	Name    string `xml:",chardata"`
}

func NewConference(name string) Conference {
	if name == "" {
		name = "default-conference"
	}
	return Conference{XMLName: "Conference", Name: name}
}

func (n Conference) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) Conference(name string) *Response { return r.Add(NewConference(name)) }
