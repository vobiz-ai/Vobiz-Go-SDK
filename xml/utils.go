package xml

import (
	"encoding/xml"
	"strings"
)

func marshalNode(v interface{}) (string, error) {
	b, err := xml.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func xmlEscape(s string) string {
	var b strings.Builder
	_ = xml.EscapeText(&b, []byte(s))
	return b.String()
}
