package internal

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
)

// PrintAllTokens consumes all the tokens from the lexer and prints them to stdout.
func PrintAllTokens(lexer antlr.Lexer) {
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%d %s\n", t.GetChannel(), t.GetText())
	}
}

// TestingErrorListener is a antlr.ErrorListener which fails a given test if a
// SyntaxError is found. The other Report errors are ignored.
type TestingErrorListener struct {
	*antlr.DefaultErrorListener

	t    *testing.T
	name string
}

func NewTestingErrorListener(t *testing.T, name string) *TestingErrorListener {
	return &TestingErrorListener{t: t, name: name}
}

func (l *TestingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.t.Errorf("%s: SyntaxError %d:%d: %s", l.name, line, column, msg)
}
