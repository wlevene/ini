package ast

type Doc struct {
	KVs   []*KVNode
	SectS []*SetcionNode
}

func NewDoc() *Doc {
	doc := &Doc{
		KVs:   make([]*KVNode, 0),
		SectS: make([]*SetcionNode, 0),
	}

	return doc
}

func (doc *Doc) Dump() {
	for _, item := range doc.KVs {
		item.Dump(nil, 0)
	}

	for _, item := range doc.SectS {
		item.Dump(nil, 0)
	}
}
