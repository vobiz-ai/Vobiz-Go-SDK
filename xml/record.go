package xml

// Record corresponds to the <Record> verb.
type Record struct {
	XMLName             string `xml:"Record"`
	Action              string `xml:"action,attr,omitempty"`
	Method              string `xml:"method,attr,omitempty"`
	Redirect            *bool  `xml:"redirect,attr,omitempty"`
	FileFormat          string `xml:"fileFormat,attr,omitempty"`
	Timeout             int    `xml:"timeout,attr,omitempty"`
	MaxLength           int    `xml:"maxLength,attr,omitempty"`
	PlayBeep            *bool  `xml:"playBeep,attr,omitempty"`
	FinishOnKey         string `xml:"finishOnKey,attr,omitempty"`
	RecordSession       *bool  `xml:"recordSession,attr,omitempty"`
	StartOnDialAnswer   *bool  `xml:"startOnDialAnswer,attr,omitempty"`
	TranscriptionType   string `xml:"transcriptionType,attr,omitempty"`
	TranscriptionURL    string `xml:"transcriptionUrl,attr,omitempty"`
	TranscriptionMethod string `xml:"transcriptionMethod,attr,omitempty"`
	CallbackURL         string `xml:"callbackUrl,attr,omitempty"`
	CallbackMethod      string `xml:"callbackMethod,attr,omitempty"`
}

func NewRecord() Record {
	redir := false
	playBeep := true
	recSession := false
	startOnDial := false
	return Record{
		XMLName:           "Record",
		Method:            "POST",
		Redirect:          &redir,
		FileFormat:        "mp3",
		Timeout:           60,
		MaxLength:         60,
		PlayBeep:          &playBeep,
		FinishOnKey:       "#",
		RecordSession:     &recSession,
		StartOnDialAnswer: &startOnDial,
		TranscriptionType: "auto",
		TranscriptionMethod: "POST",
		CallbackMethod:      "POST",
	}
}

func (n Record) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) Record() *Response { return r.Add(NewRecord()) }
