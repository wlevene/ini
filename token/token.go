package token

import "fmt"

type TokenType string

const (
	TokenTypeILLEGAL = "ILLEGAL"
	TokenTypeEOF     = "EOF"
	Ident            = "Ident"

	TokenTypeSECTION  = "SECTION"
	TokenTypeKEY      = "KEY"
	TokenTypeVALUE    = "VALUE"
	TokenTypeCOMMENT  = "#"
	TokenTypeCOMMENT2 = "#"

	TokenTypeASSIGN   = "="
	TokenTypeLBRACKET = "["
	TokenTypeRBRACKET = "]"

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
