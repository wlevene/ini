package lexer

import (
	"fmt"
	"testing"
)

func TestLexer(t *testing.T) {
	var ch byte
	ch = ']'

	fmt.Println(isLetter(ch) || isDigit(ch))

}
