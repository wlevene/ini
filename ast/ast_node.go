package ast

import (
	"fmt"

	"github.com/wlevene/ini/token"
)

type SetcionNode struct {
	BaseNode
	Name token.Token
}

var KindSetcion = NewNodeKind("Section")

// Kind implements Node.Kind.
func (n *SetcionNode) Kind() NodeKind {
	return KindSetcion
}

func (n *SetcionNode) Type() NodeType {
	return TypeDocNode
}

func (n *SetcionNode) IsRaw() bool {
	return false
}

// Text implements Node.Text.
func (n *SetcionNode) Text(source []byte) []byte {
	return []byte(n.Name.Literal)
}

// Dump implements Node.Dump.
func (n *SetcionNode) Dump(source []byte, level int) {

	if n == nil {
		return
	}

	name := n.Name.Literal
	m := map[string]string{
		"Section": fmt.Sprintf("[%v]", name),
		"Line":    fmt.Sprintf("%v", n.Name.Line),
	}
	DumpHelper(n, source, level, m, nil)
}

type KVNode struct {
	BaseNode
	Key   token.Token
	Value token.Token
	Index int
	Line  int
}

var KindKVNode = NewNodeKind("KVNode")

// Kind implements Node.Kind.
func (n *KVNode) Kind() NodeKind {
	return KindKVNode
}

func (n *KVNode) Type() NodeType {
	return TypeDocNode
}

func (n *KVNode) IsRaw() bool {
	return false
}

func (n *KVNode) Text(source []byte) []byte {
	return source
}

// Dump implements Node.Dump.
func (n *KVNode) Dump(source []byte, level int) {

	m := map[string]string{
		"Key":   fmt.Sprintf("%v", n.Key.Literal),
		"Value": fmt.Sprintf("%v", n.Value.Literal),
		"Line":  fmt.Sprintf("%v", n.Key.Line),
	}

	DumpHelper(n, source, level, m, nil)
}

type CommentNode struct {
	BaseNode
	Comment token.Token
}

var KindCommentNode = NewNodeKind("CommentNode")

// Kind implements Node.Kind.
func (n *CommentNode) Kind() NodeKind {
	return KindCommentNode
}

func (n *CommentNode) Type() NodeType {
	return TypeDocNode
}

func (n *CommentNode) IsRaw() bool {
	return false
}

func (n *CommentNode) Text(source []byte) []byte {
	return source
}

// Dump implements Node.Dump.
func (n *CommentNode) Dump(source []byte, level int) {

	m := map[string]string{
		"Comment": fmt.Sprintf("%v", n.Comment.Literal),
		"Line":    fmt.Sprintf("%v", n.Comment.Line),
	}

	DumpHelper(n, source, level, m, nil)
}
