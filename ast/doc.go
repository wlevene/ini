package ast

import "fmt"

type Doc struct {
	BaseNode

	// KVs   []*KVNode
	// SectS []*SetcionNode
}

func NewDoc() *Doc {
	doc := &Doc{
		// KVs:   make([]*KVNode, 0),
		// SectS: make([]*SetcionNode, 0),
	}

	return doc
}

func (doc *Doc) DumpV2() {

	// for _, item := range doc.KVs {
	// 	item.Dump(nil, 0)
	// }

	// for _, item := range doc.SectS {
	// 	item.Dump(nil, 0)
	// }
	DumpHelper(doc, nil, 0, nil, nil)
}

var KindDocNode = NewNodeKind("INIDocNode")

// Kind implements Node.Kind.
func (n *Doc) Kind() NodeKind {
	return KindDocNode
}

func (n *Doc) Type() NodeType {
	return TypeDocNode
}

func (n *Doc) IsRaw() bool {
	return false
}

func (n *Doc) Text(source []byte) []byte {
	return source
}

// Dump implements Node.Dump.
func (n *Doc) Dump(source []byte, level int) {

	m := map[string]string{
		"doc": fmt.Sprintf("%v", string(source)),
	}

	DumpHelper(n, source, level, m, nil)
}
