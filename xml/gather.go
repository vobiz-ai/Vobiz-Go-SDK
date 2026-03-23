package xml

// GatherSpeak represents nested <Speak> prompt inside <Gather>.
type GatherSpeak struct {
	XMLName  string `xml:"Speak"`
	Voice    string `xml:"voice,attr,omitempty"`
	Language string `xml:"language,attr,omitempty"`
	Text     string `xml:",chardata"`
}

// Gather corresponds to the <Gather> verb.
type Gather struct {
	XMLName                       string     `xml:"Gather"`
	Action                        string     `xml:"action,attr,omitempty"`
	Method                        string     `xml:"method,attr,omitempty"`
	InputType                     string     `xml:"inputType,attr,omitempty"`
	ExecutionTimeout              int        `xml:"executionTimeout,attr,omitempty"`
	DigitEndTimeout               string     `xml:"digitEndTimeout,attr,omitempty"`
	SpeechEndTimeout              string     `xml:"speechEndTimeout,attr,omitempty"`
	FinishOnKey                   string     `xml:"finishOnKey,attr,omitempty"`
	NumDigits                     int        `xml:"numDigits,attr,omitempty"`
	SpeechModel                   string     `xml:"speechModel,attr,omitempty"`
	Language                      string     `xml:"language,attr,omitempty"`
	InterimSpeechResultsCallback  string     `xml:"interimSpeechResultsCallback,attr,omitempty"`
	Log                           *bool      `xml:"log,attr,omitempty"`
	Redirect                      *bool      `xml:"redirect,attr,omitempty"`
	ProfanityFilter               *bool      `xml:"profanityFilter,attr,omitempty"`
	Prompt                        GatherSpeak `xml:"Speak"`
}

func NewGather(prompt string) Gather {
	log := true
	redir := true
	profanity := false
	if prompt == "" {
		prompt = "Please press a digit or speak your input."
	}
	return Gather{
		XMLName:          "Gather",
		Method:           "POST",
		InputType:        "dtmf",
		ExecutionTimeout: 15,
		DigitEndTimeout:  "auto",
		SpeechEndTimeout: "auto",
		FinishOnKey:      "#",
		NumDigits:        1,
		SpeechModel:      "default",
		Language:         "en-US",
		Log:              &log,
		Redirect:         &redir,
		ProfanityFilter:  &profanity,
		Prompt: GatherSpeak{XMLName: "Speak", Voice: "WOMAN", Language: "en-US", Text: prompt},
	}
}

func (n Gather) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) Gather(prompt string) *Response { return r.Add(NewGather(prompt)) }
