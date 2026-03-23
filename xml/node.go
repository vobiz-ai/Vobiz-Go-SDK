package xml

// Node is the base interface for all XML response nodes.
type Node interface {
	xmlString() (string, error)
}
