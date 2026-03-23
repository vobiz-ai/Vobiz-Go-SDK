package xml

// PreAnswer corresponds to the <PreAnswer> verb.
type PreAnswer struct {
	XMLName string `xml:"PreAnswer"`
}

func NewPreAnswer() PreAnswer { return PreAnswer{XMLName: "PreAnswer"} }

func (n PreAnswer) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) PreAnswer() *Response { return r.Add(NewPreAnswer()) }
