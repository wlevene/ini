package lexer

import (
	"github.com/wlevene/ini/token"
)

const (
	LineBreak = '\n'
)

type Lexer struct {
	Input         []byte
	char          byte
	position      int
	read_position int
	line          int
}

func New(input string) *Lexer {
	l := &Lexer{Input: []byte(input)}
	l.readChar()
	return l
}

func (l *Lexer) peekChar() byte {
	if l.read_position >= len(l.Input) {
		return 0
	} else {
		return l.Input[l.read_position]
	}
}

func (l *Lexer) readChar() {
	if l.read_position >= len(l.Input) {
		l.char = 0
	} else {
		l.char = l.Input[l.read_position]
	}

	if l.char == LineBreak {
		l.line++
	}

	l.position = l.read_position
	l.read_position++

}

func (l *Lexer) readLine() []byte {
	position := l.position
	for {
		if l.char == LineBreak {
			break
		}
		l.readChar()
	}

	return l.Input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' ||
		l.char == '\t' ||
		l.char == LineBreak ||
		l.char == '\r' {
		l.readChar()
	}
}

func (l *Lexer) skipspace() {
	for l.char == ' ' {
		l.readChar()
	}

	// for l.char == ' ' ||
	// 	l.char == '\t' {
	// 	l.readChar()
	// }
}

func (l *Lexer) skipline() {

	for {
		if l.char == LineBreak {
			// l.line++
			break
		}
		l.readChar()
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	var ch byte
	l.skipWhitespace()

	switch l.char {
	case '[':
		tok = newToken(token.TokenTypeLBRACKET, l.char, l.line)
	case '=':
		tok = newToken(token.TokenTypeASSIGN, l.char, l.line)
	case ']':
		tok = newToken(token.TokenTypeRBRACKET, l.char, l.line)
	case 0:
		tok.Literal = ""
		tok.Type = token.TokenTypeEOF
		return tok
	case '#':
		fallthrough
	case ';':
		tok.Line = l.line
		tok.Literal = string(l.readLine())
		tok.Type = token.TokenTypeCOMMENT

	default:
		tok.Line = l.line
		tok.Literal = string(l.readStatement())
		l.skipspace()

		ch = l.char

		if ch == '=' {
			tok.Type = token.TokenTypeKEY
		} else if ch == ']' {
			tok.Type = token.TokenTypeSECTION
		} else if ch == LineBreak {
			tok.Type = token.TokenTypeVALUE
		}
	}

	l.readChar()
	return tok
}

/*

func (s *lexer) getNextToken() token.Token {
reToken:
	ch := s.peek()
	switch {
	case isSpace(ch):
		s.skipWhitespace()
		goto reToken
	case isEOF(ch):
		return s.NewToken(token.EOF).Lit(string(s.read()))
	case ch == ';':
		return s.NewToken(token.Semicolon).Lit(string(s.read()))
	case ch == '{':
		return s.NewToken(token.BlockStart).Lit(string(s.read()))
	case ch == '}':
		return s.NewToken(token.BlockEnd).Lit(string(s.read()))
	case ch == '#':
		return s.scanComment()
	case ch == '$':
		return s.scanVariable()
	case isQuote(ch):
		return s.scanQuotedString(ch)
	default:
		return s.scanKeyword()
	}
}

*/

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return ('0' <= ch && ch <= '9') || ch == '.'
}

func newToken(tokenType token.TokenType, ch byte, line int) token.Token {
	return token.Token{Type: tokenType,
		Literal: string(ch),
		Line:    line,
	}
}

func (l *Lexer) readStatement() []byte {
	position := l.position
	for l.char != LineBreak &&
		l.char != '=' &&
		l.char != ']' &&
		l.char != ' ' {

		l.readChar()
	}
	return l.Input[position:l.position]
}

func (l *Lexer) readSection() token.Token {
	var tok token.Token
	if l.char != '[' {
		return tok
	}
	position := l.position

	for {
		if l.char != ']' {
			l.readChar()
		}

		if l.char == LineBreak {
			break
		}
	}

	return token.Token{
		Type:    token.TokenTypeSECTION,
		Literal: string(l.Input[position:l.position]),
		Line:    l.line,
	}

}

/*

func (s *lexer) read() rune {
	ch, _, err := s.reader.ReadRune()
	if err != nil {
		return rune(token.EOF)
	}

	if ch == '\n' {
		s.column = 1
		s.line++
	} else {
		s.column++
	}
	return ch
}

*/
