package xml

// Hangup corresponds to the <Hangup> verb.
type Hangup struct {
	XMLName  string `xml:"Hangup"`
	Reason   string `xml:"reason,attr,omitempty"`
	Schedule string `xml:"schedule,attr,omitempty"`
}

func NewHangup() Hangup { return Hangup{XMLName: "Hangup"} }

func (n Hangup) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) Hangup() *Response { return r.Add(NewHangup()) }
