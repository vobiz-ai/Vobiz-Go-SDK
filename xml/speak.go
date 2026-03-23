package xml

// Speak corresponds to the <Speak> verb.
type Speak struct {
	XMLName  string `xml:"Speak"`
	Voice    string `xml:"voice,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
	Loop     int    `xml:"loop,attr,omitempty"`
	Text     string `xml:",chardata"`
}

func NewSpeak(text string) Speak {
	return Speak{XMLName: "Speak", Voice: "WOMAN", Language: "en-US", Loop: 1, Text: text}
}

func (n Speak) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) Speak(text string) *Response { return r.Add(NewSpeak(text)) }
