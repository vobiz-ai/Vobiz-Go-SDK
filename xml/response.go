package xml

import "strings"

// Response is the root Vobiz XML response builder.
type Response struct {
	nodes []Node
}

// NewResponse creates a new empty XML response.
func NewResponse() *Response {
	return &Response{nodes: make([]Node, 0, 8)}
}

// Add appends a generic node to the response.
func (r *Response) Add(n Node) *Response {
	if n != nil {
		r.nodes = append(r.nodes, n)
	}
	return r
}

// String renders the XML response as a UTF-8 XML document.
func (r *Response) String() string {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	b.WriteString("\n<Response>")
	for _, n := range r.nodes {
		s, err := n.xmlString()
		if err != nil {
			continue
		}
		b.WriteString("\n  ")
		b.WriteString(s)
	}
	b.WriteString("\n</Response>")
	return b.String()
}

// Bytes renders the XML response as bytes.
func (r *Response) Bytes() []byte {
	return []byte(r.String())
}
