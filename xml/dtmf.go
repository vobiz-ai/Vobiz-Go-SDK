package xml

// DTMF corresponds to the <DTMF> verb.
type DTMF struct {
	XMLName string `xml:"DTMF"`
	Async   *bool  `xml:"async,attr,omitempty"`
	Digits  string `xml:",chardata"`
}

func NewDTMF(digits string, async bool) DTMF {
	return DTMF{XMLName: "DTMF", Digits: digits, Async: &async}
}

func (n DTMF) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) DTMF(digits string) *Response {
	v := true
	return r.Add(DTMF{XMLName: "DTMF", Digits: digits, Async: &v})
}
