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
	// doc.Statements = []ast.Statement{}

	// Read two tokens, so currentToken and peekToken are both set.
	p.nextToken()
	p.nextToken()

	var curSection *ast.SetcionNode
	var KV *ast.KVNode

	for !p.curTokenIs(token.TokenType_EOF) {

		if p.currentToken.Type == token.TokenType_SECTION {

			if curSection != nil {
				doc.SectS = append(doc.SectS, curSection)
			}
			curSection = nil
			curSection = &ast.SetcionNode{
				Name: p.currentToken,
			}

		} else if p.currentToken.Type == token.TokenType_KEY {
			KV = nil
			KV = &ast.KVNode{}
			KV.Key = p.currentToken

		} else if p.currentToken.Type == token.TokenType_VALUE {
			if KV != nil {
				KV.Value = p.currentToken
				if curSection != nil {
					curSection.AppendChild(curSection, KV)
				} else {
					doc.KVs = append(doc.KVs, KV)
				}
			}
		}

		p.nextToken()
	}

	if curSection != nil {
		doc.SectS = append(doc.SectS, curSection)
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
	return p.currentToken.Line == p.peekToken.Line && p.peekToken.Type != token.TokenType_EOF
}

// func (p *Parser) parseStatement() ast.Statement {

// 	var stmt ast.Statement
// 	ast.
// 	switch p.currentToken.Type {
// case token.Return:
// 	return p.parseReturnStatement()
// case token.Def:
// 	return p.parseDefMethodStatement()
// case token.Comment:
// 	return nil
// case token.While:
// 	return p.parseWhileStatement()
// case token.Class:
// 	return p.parseClassStatement()
// case token.Module:
// 	return p.parseModuleStatement()
// case token.Next:
// 	return &ast.NextStatement{BaseNode: &ast.BaseNode{Token: p.curToken}}
// case token.Break:
// 	return &ast.BreakStatement{BaseNode: &ast.BaseNode{Token: p.curToken}}
// default:
// 	exp := p.parseExpressionStatement()

// 	// If parseExpressionStatement got error exp.Expression would be nil
// 	if exp.Expression != nil {
// 		// In REPL mode everything should return a value.
// 		if p.Mode == REPLMode {
// 			exp.Expression.MarkAsExp()
// 		} else {
// 			exp.Expression.MarkAsStmt()
// 		}
// 	}
// 	}
// 	return stmt
// }
