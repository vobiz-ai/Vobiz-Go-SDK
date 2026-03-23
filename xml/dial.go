package xml

// DialNumber represents nested <Number> inside <Dial>.
type DialNumber struct {
	XMLName    string `xml:"Number"`
	SendDigits string `xml:"sendDigits,attr,omitempty"`
	Number     string `xml:",chardata"`
}

// Dial corresponds to the <Dial> verb.
type Dial struct {
	XMLName        string     `xml:"Dial"`
	Action         string     `xml:"action,attr,omitempty"`
	Method         string     `xml:"method,attr,omitempty"`
	Redirect       *bool      `xml:"redirect,attr,omitempty"`
	HangupOnStar   *bool      `xml:"hangupOnStar,attr,omitempty"`
	TimeLimit      int        `xml:"timeLimit,attr,omitempty"`
	Timeout        int        `xml:"timeout,attr,omitempty"`
	CallerID       string     `xml:"callerId,attr,omitempty"`
	CallerName     string     `xml:"callerName,attr,omitempty"`
	ConfirmSound   string     `xml:"confirmSound,attr,omitempty"`
	ConfirmKey     string     `xml:"confirmKey,attr,omitempty"`
	DialMusic      string     `xml:"dialMusic,attr,omitempty"`
	CallbackURL    string     `xml:"callbackUrl,attr,omitempty"`
	CallbackMethod string     `xml:"callbackMethod,attr,omitempty"`
	Number         DialNumber `xml:"Number"`
}

func NewDial(number string) Dial {
	redir := true
	return Dial{
		XMLName:  "Dial",
		Method:   "POST",
		Redirect: &redir,
		TimeLimit: 14400,
		DialMusic: "real",
		Number:   DialNumber{XMLName: "Number", Number: number},
	}
}

func (n Dial) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) Dial(number string) *Response { return r.Add(NewDial(number)) }
