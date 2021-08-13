package token

import "fmt"

type TokenType string

const (
	TokenType_ILLEGAL = "ILLEGAL"
	TokenType_EOF     = "EOF"
	Ident             = "Ident"

	TokenType_SECTION  = "SECTION"
	TokenType_KEY      = "KEY"
	TokenType_VALUE    = "VALUE"
	TokenType_COMMENT  = "#"
	TokenType_COMMENT2 = "#"

	TokenType_ASSIGN   = "="
	TokenType_LBRACKET = "["
	TokenType_RBRACKET = "]"

	RETURN = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

var keywords = map[string]TokenType{
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Ident
}

func (t Token) String() string {
	return fmt.Sprintf("{Type:%s,Literal:\"%s\" Linie:%d}",
		t.Type,
		t.Literal,
		t.Line)
}

//Is check type of a token
func (t Token) Is(typ TokenType) bool {
	return t.Type == typ
}
