package xml

import (
	"strings"
	"testing"
)

func TestResponseSpeakWaitXMLStructure(t *testing.T) {
	out := NewResponse().Speak("Hello").Wait(2).String()

	if !strings.Contains(out, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>") {
		t.Fatalf("missing xml declaration: %s", out)
	}
	if !strings.Contains(out, "<Response>") || !strings.Contains(out, "</Response>") {
		t.Fatalf("missing response tags: %s", out)
	}
	if !strings.Contains(out, "<Speak") || !strings.Contains(out, ">Hello</Speak>") {
		t.Fatalf("missing speak node: %s", out)
	}
	if !strings.Contains(out, "<Wait length=\"2\"></Wait>") {
		t.Fatalf("missing wait node: %s", out)
	}
}
