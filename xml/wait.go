package xml

// Wait corresponds to the <Wait> verb.
type Wait struct {
	XMLName    string `xml:"Wait"`
	Length     int    `xml:"length,attr,omitempty"`
	Silence    *bool  `xml:"silence,attr,omitempty"`
	MinSilence int    `xml:"minSilence,attr,omitempty"`
	Beep       *bool  `xml:"beep,attr,omitempty"`
}

func NewWait(length int) Wait {
	if length <= 0 {
		length = 1
	}
	return Wait{XMLName: "Wait", Length: length}
}

func (n Wait) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) Wait(length int) *Response { return r.Add(NewWait(length)) }
