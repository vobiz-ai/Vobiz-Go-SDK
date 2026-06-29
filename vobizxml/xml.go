// Package vobizxml builds VobizXML call-control documents in Go.
//
// It mirrors Plivo's plivoxml (a Response root + Add* builder helpers + a
// String() serializer) and emits XML byte-identical to the Python vobizxml
// package and the Node @vobiz/sdk vobizxml namespace. The whole package lives in
// this single file so it drops cleanly into the published Go SDK as the
// github.com/vobiz-ai/Vobiz-Go-SDK/vobizxml subpackage (no internal imports to
// resolve).
//
// Attribute names are the camelCase VobizXML names, supplied via the Attr
// option. Because Go maps do not preserve insertion order, attributes are stored
// in an ordered slice so the rendered order matches the order you set them.
//
//	import "github.com/vobiz-ai/Vobiz-Go-SDK/vobizxml"
//
//	r := vobizxml.NewResponse()
//	g := r.AddGather(
//	    vobizxml.Attr("action", "https://yourapp.com/menu-choice"),
//	    vobizxml.Attr("inputType", "dtmf"), // Gather uses inputType / executionTimeout, never timeout
//	    vobizxml.Attr("numDigits", 1),
//	    vobizxml.Attr("executionTimeout", 10),
//	)
//	g.AddSpeak("Press 1 for sales, 2 for support, or 0 for an operator.")
//	r.AddSpeak("We didn't receive your input. Goodbye.")
//	r.AddHangup()
//	fmt.Println(r.String())        // pretty, with the XML declaration
//	fmt.Println(r.StringCompact()) // single line
package vobizxml

import (
	"fmt"
	"strings"
)

// XMLDeclaration is prepended to every serialized document.
const XMLDeclaration = `<?xml version="1.0" encoding="UTF-8"?>`

// indentUnit is one indentation level (4 spaces, matching the xml/*.mdx style).
const indentUnit = "    "

// attr is a single key/value attribute. Stored in an ordered slice so the
// emitted attribute order matches insertion order (Go maps reorder).
type attr struct {
	key   string
	value string
}

// Option configures an element when it is created. Options are applied in the
// order given, and Attr appends to an ordered list, so attribute order is stable.
type Option func(*element)

// Attr sets a VobizXML attribute by its camelCase name (for example
// "inputType", "executionTimeout", "numDigits", "callerId",
// "startConferenceOnEnter", "sendDigits", "audioTrack").
//
// The value is rendered as: bool -> "true"/"false"; everything else via
// fmt.Sprint. Insertion order is preserved.
func Attr(key string, value any) Option {
	return func(e *element) {
		e.attrs = append(e.attrs, attr{key: key, value: attrValue(value)})
	}
}

// attrValue renders an attribute value the way the reference builders do.
func attrValue(value any) string {
	switch v := value.(type) {
	case bool:
		if v {
			return "true"
		}
		return "false"
	case string:
		return v
	default:
		return fmt.Sprint(value)
	}
}

// escapeText escapes XML text content (& < >).
func escapeText(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	return s
}

// escapeAttr escapes an attribute value (text rules plus double quotes).
func escapeAttr(s string) string {
	return strings.ReplaceAll(escapeText(s), `"`, "&quot;")
}

func strptr(s string) *string { return &s }

// element is the internal representation of a single VobizXML element: an
// ordered attribute list, optional text content, and child elements.
type element struct {
	name     string
	content  *string // nil => no content (self-closing unless it has children)
	raw      bool    // if true, content is emitted without escaping (SSML)
	attrs    []attr
	children []*element
}

func newElement(name string, content *string, raw bool, opts []Option) *element {
	e := &element{name: name, content: content, raw: raw}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *element) openTag() string {
	if len(e.attrs) == 0 {
		return e.name
	}
	var b strings.Builder
	b.WriteString(e.name)
	for _, a := range e.attrs {
		b.WriteString(" ")
		b.WriteString(a.key)
		b.WriteString(`="`)
		b.WriteString(escapeAttr(a.value))
		b.WriteString(`"`)
	}
	return b.String()
}

func (e *element) render(level int, pretty bool) string {
	pad := ""
	if pretty {
		pad = strings.Repeat(indentUnit, level)
	}
	open := e.openTag()

	// Empty element -> self-closing.
	if len(e.children) == 0 && e.content == nil {
		return pad + "<" + open + "/>"
	}
	// Text-content element -> single line.
	if len(e.children) == 0 {
		body := *e.content
		if !e.raw {
			body = escapeText(body)
		}
		return pad + "<" + open + ">" + body + "</" + e.name + ">"
	}
	// Container element -> children indented (content, if any, is ignored).
	parts := make([]string, len(e.children))
	for i, c := range e.children {
		parts[i] = c.render(level+1, pretty)
	}
	if pretty {
		return pad + "<" + open + ">\n" + strings.Join(parts, "\n") + "\n" + pad + "</" + e.name + ">"
	}
	return "<" + open + ">" + strings.Join(parts, "") + "</" + e.name + ">"
}

func (e *element) toString(pretty bool) string {
	body := e.render(0, pretty)
	if pretty {
		return XMLDeclaration + "\n" + body
	}
	return XMLDeclaration + body
}

// node is the embeddable base shared by every element type. It carries the
// serialization helpers and the internal child-appending builders.
type node struct {
	*element
}

// String serializes the element to a pretty-printed VobizXML document (4-space
// indentation) with the XML declaration. It satisfies fmt.Stringer.
func (n node) String() string { return n.element.toString(true) }

// StringCompact serializes the element to a single-line VobizXML document with
// the XML declaration and no indentation or newlines.
func (n node) StringCompact() string { return n.element.toString(false) }

// Set adds or overrides attributes on the element after construction.
func (n node) Set(opts ...Option) {
	for _, opt := range opts {
		opt(n.element)
	}
}

// add appends a child element and returns it.
func (n node) add(child *element) *element {
	n.element.children = append(n.element.children, child)
	return child
}

// --- Internal builder helpers (shared; re-exported per container type) -------

func (n node) addSpeak(text string, opts ...Option) *Speak {
	return &Speak{node{n.add(newElement("Speak", strptr(text), false, opts))}}
}

func (n node) addSpeakSSML(ssml string, opts ...Option) *Speak {
	return &Speak{node{n.add(newElement("Speak", strptr(ssml), true, opts))}}
}

func (n node) addPlay(url string, opts ...Option) *Play {
	return &Play{node{n.add(newElement("Play", strptr(url), false, opts))}}
}

func (n node) addWait(opts ...Option) *Wait {
	return &Wait{node{n.add(newElement("Wait", nil, false, opts))}}
}

func (n node) addGather(opts ...Option) *Gather {
	return &Gather{node{n.add(newElement("Gather", nil, false, opts))}}
}

func (n node) addDial(opts ...Option) *Dial {
	return &Dial{node{n.add(newElement("Dial", nil, false, opts))}}
}

func (n node) addRecord(opts ...Option) *Record {
	return &Record{node{n.add(newElement("Record", nil, false, opts))}}
}

func (n node) addConference(room string, opts ...Option) *Conference {
	return &Conference{node{n.add(newElement("Conference", strptr(room), false, opts))}}
}

func (n node) addDTMF(digits string, opts ...Option) *DTMF {
	return &DTMF{node{n.add(newElement("DTMF", strptr(digits), false, opts))}}
}

func (n node) addRedirect(url string, opts ...Option) *Redirect {
	return &Redirect{node{n.add(newElement("Redirect", strptr(url), false, opts))}}
}

func (n node) addHangup(opts ...Option) *Hangup {
	return &Hangup{node{n.add(newElement("Hangup", nil, false, opts))}}
}

func (n node) addPreAnswer(opts ...Option) *PreAnswer {
	return &PreAnswer{node{n.add(newElement("PreAnswer", nil, false, opts))}}
}

func (n node) addStream(url string, opts ...Option) *Stream {
	return &Stream{node{n.add(newElement("Stream", strptr(url), false, opts))}}
}

func (n node) addNumber(number string, opts ...Option) *Number {
	return &Number{node{n.add(newElement("Number", strptr(number), false, opts))}}
}

func (n node) addUser(sipURI string, opts ...Option) *User {
	return &User{node{n.add(newElement("User", strptr(sipURI), false, opts))}}
}

// --- Leaf / content element types --------------------------------------------

// Speak is a <Speak> text-to-speech element.
type Speak struct{ node }

// Play is a <Play> remote audio element.
type Play struct{ node }

// Wait is a <Wait/> silent pause element.
type Wait struct{ node }

// Number is a <Number> PSTN endpoint nested in <Dial>.
type Number struct{ node }

// User is a <User> SIP endpoint nested in <Dial>.
type User struct{ node }

// Record is a <Record/> recording element.
type Record struct{ node }

// Conference is a <Conference> room element.
type Conference struct{ node }

// DTMF is a <DTMF> send-digits element.
type DTMF struct{ node }

// Redirect is a <Redirect> flow-transfer element.
type Redirect struct{ node }

// Hangup is a <Hangup/> end-call element.
type Hangup struct{ node }

// Stream is a <Stream> WebSocket audio-fork element.
type Stream struct{ node }

// --- Container element types -------------------------------------------------

// Gather is a <Gather> input-collection element. Nest Speak/Play prompts inside.
type Gather struct{ node }

// AddSpeak nests a <Speak> prompt and returns it.
func (g *Gather) AddSpeak(text string, opts ...Option) *Speak { return g.addSpeak(text, opts...) }

// AddSpeakSSML nests a <Speak> prompt with raw (unescaped) SSML content.
func (g *Gather) AddSpeakSSML(ssml string, opts ...Option) *Speak {
	return g.addSpeakSSML(ssml, opts...)
}

// AddPlay nests a <Play> prompt and returns it.
func (g *Gather) AddPlay(url string, opts ...Option) *Play { return g.addPlay(url, opts...) }

// PreAnswer is a <PreAnswer> early-media block. Nest Speak/Play/Wait inside.
type PreAnswer struct{ node }

// AddSpeak nests a <Speak> element and returns it.
func (p *PreAnswer) AddSpeak(text string, opts ...Option) *Speak { return p.addSpeak(text, opts...) }

// AddSpeakSSML nests a <Speak> element with raw (unescaped) SSML content.
func (p *PreAnswer) AddSpeakSSML(ssml string, opts ...Option) *Speak {
	return p.addSpeakSSML(ssml, opts...)
}

// AddPlay nests a <Play> element and returns it.
func (p *PreAnswer) AddPlay(url string, opts ...Option) *Play { return p.addPlay(url, opts...) }

// AddWait nests a <Wait/> element and returns it.
func (p *PreAnswer) AddWait(opts ...Option) *Wait { return p.addWait(opts...) }

// Dial is a <Dial> element bridging the caller to Number/User endpoints; it may
// also nest a Record.
type Dial struct{ node }

// AddNumber nests a <Number> PSTN endpoint and returns it.
func (d *Dial) AddNumber(number string, opts ...Option) *Number { return d.addNumber(number, opts...) }

// AddUser nests a <User> SIP endpoint and returns it.
func (d *Dial) AddUser(sipURI string, opts ...Option) *User { return d.addUser(sipURI, opts...) }

// AddRecord nests a <Record/> element (use Attr("startOnDialAnswer", true)).
func (d *Dial) AddRecord(opts ...Option) *Record { return d.addRecord(opts...) }

// Response is the <Response> root container.
type Response struct{ node }

// NewResponse creates an empty <Response> root. Build the document with its Add*
// helpers, then serialize with String() or StringCompact().
func NewResponse() *Response {
	return &Response{node{&element{name: "Response"}}}
}

// AddSpeak adds a <Speak> element and returns it.
func (r *Response) AddSpeak(text string, opts ...Option) *Speak { return r.addSpeak(text, opts...) }

// AddSpeakSSML adds a <Speak> element with raw (unescaped) SSML content.
func (r *Response) AddSpeakSSML(ssml string, opts ...Option) *Speak {
	return r.addSpeakSSML(ssml, opts...)
}

// AddPlay adds a <Play> element and returns it.
func (r *Response) AddPlay(url string, opts ...Option) *Play { return r.addPlay(url, opts...) }

// AddWait adds a <Wait/> element and returns it.
func (r *Response) AddWait(opts ...Option) *Wait { return r.addWait(opts...) }

// AddGather adds a <Gather> element and returns it for nesting prompts.
func (r *Response) AddGather(opts ...Option) *Gather { return r.addGather(opts...) }

// AddGetDigits is a Plivo-parity alias that emits <Gather>.
func (r *Response) AddGetDigits(opts ...Option) *Gather { return r.addGather(opts...) }

// AddGetInput is a Plivo-parity alias that emits <Gather>.
func (r *Response) AddGetInput(opts ...Option) *Gather { return r.addGather(opts...) }

// AddDial adds a <Dial> element and returns it for nesting Number/User/Record.
func (r *Response) AddDial(opts ...Option) *Dial { return r.addDial(opts...) }

// AddRecord adds a <Record/> element and returns it.
func (r *Response) AddRecord(opts ...Option) *Record { return r.addRecord(opts...) }

// AddConference adds a <Conference> element (room name is the text content).
func (r *Response) AddConference(room string, opts ...Option) *Conference {
	return r.addConference(room, opts...)
}

// AddDTMF adds a <DTMF> element (digits are the text content).
func (r *Response) AddDTMF(digits string, opts ...Option) *DTMF { return r.addDTMF(digits, opts...) }

// AddRedirect adds a <Redirect> element (URL is the text content).
func (r *Response) AddRedirect(url string, opts ...Option) *Redirect {
	return r.addRedirect(url, opts...)
}

// AddHangup adds a <Hangup/> element and returns it.
func (r *Response) AddHangup(opts ...Option) *Hangup { return r.addHangup(opts...) }

// AddPreAnswer adds a <PreAnswer> element and returns it for nesting early media.
func (r *Response) AddPreAnswer(opts ...Option) *PreAnswer { return r.addPreAnswer(opts...) }

// AddStream adds a <Stream> element (wss URL is the text content).
func (r *Response) AddStream(url string, opts ...Option) *Stream { return r.addStream(url, opts...) }
