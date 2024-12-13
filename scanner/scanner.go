package scanner

import (
	"fmt"

	"github.com/wlevene/ini/lexer"
	"github.com/wlevene/ini/parser"
)

type Scanner struct {
	src []byte // source
	l   *lexer.Lexer
	p   *parser.Parser
}

func New(buf []byte) *Scanner {
	s := &Scanner{
		src: buf,
	}

	return s
}

func (this *Scanner) Scan() {
	this.l = lexer.New(string(this.src))
	this.p = parser.New(this.l)

	doc, err := this.p.ParseDocument()

	if err != nil {
		fmt.Println("err:", err)
	}
	doc.DumpV2()
}
