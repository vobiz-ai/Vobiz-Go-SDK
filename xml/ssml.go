package xml

import "encoding/xml"

// SSML renders a <Speak> node with raw inner XML content.
type SSML struct {
	XMLName  xml.Name `xml:"Speak"`
	Voice    string   `xml:"voice,attr,omitempty"`
	Language string   `xml:"language,attr,omitempty"`
	Loop     int      `xml:"loop,attr,omitempty"`
	InnerXML string   `xml:",innerxml"`
}

func NewSSML(innerXML string) SSML {
	return SSML{XMLName: xml.Name{Local: "Speak"}, Voice: "WOMAN", Language: "en-US", Loop: 1, InnerXML: innerXML}
}

func (n SSML) xmlString() (string, error) { return marshalNode(n) }

func (r *Response) SSML(innerXML string) *Response { return r.Add(NewSSML(innerXML)) }
