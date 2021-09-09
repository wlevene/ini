package parser

import (
	"github.com/wlevene/ini/ast"
	"github.com/wlevene/ini/lexer"
	"github.com/wlevene/ini/token"
)

// type (
// 	prefixParseFn func() ast.Expression
// 	infixParseFn  func(ast.Expression) ast.Expression
// )

type Parser struct {
	lexer        *lexer.Lexer
	errors       []string
	currentToken token.Token
	peekToken    token.Token

	// prefixParseFns map[token.TokenType]prefixParseFn
	// infixParseFns  map[token.TokenType]infixParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		lexer:  l,
		errors: []string{},
	}

	// p.prefixParseFns = make(map[token.Type]prefixParseFn)
	// p.registerPrefix(token.INTEGER, p.parseIntegerLiteral)

	// p.infixParseFns = make(map[token.Type]infixParseFn)
	// p.registerInfix(token.PLUS, p.parseInfixExpression)

	return p
}

func (p *Parser) ParseDocument() (doc *ast.Doc, err error) {

	doc = ast.NewDoc()

	// Read two tokens, so currentToken and peekToken are both set.
	p.nextToken()
	p.nextToken()

	var curSection *ast.SetcionNode
	var KV *ast.KVNode

	for !p.curTokenIs(token.TokenTypeEOF) {

		if p.currentToken.Type == token.TokenTypeSECTION {

			curSection = nil
			curSection = &ast.SetcionNode{
				Name: p.currentToken,
			}

			if curSection != nil {
				doc.AppendChild(doc, curSection)
			}

		} else if p.currentToken.Type == token.TokenTypeKEY {
			KV = nil
			KV = &ast.KVNode{}
			KV.Key = p.currentToken

		} else if p.currentToken.Type == token.TokenTypeVALUE {
			if KV != nil {
				KV.Value = p.currentToken
				if curSection != nil {
					curSection.AppendChild(curSection, KV)
				} else {
					doc.AppendChild(doc, KV)
				}
			}
		} else if p.currentToken.Type == token.TokenTypeCOMMENT {
			comment := &ast.CommentNode{
				Comment: p.currentToken,
			}
			doc.AppendChild(doc, comment)
		}

		p.nextToken()
	}

	return doc, nil
}
func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.lexer.NextToken()

}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
// 	p.prefixParseFns[tokenType] = fn
// }

// func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
// 	p.infixParseFns[tokenType] = fn
// }

func (p *Parser) peekTokenAtSameLine() bool {
	return p.currentToken.Line == p.peekToken.Line && p.peekToken.Type != token.TokenTypeEOF
}
