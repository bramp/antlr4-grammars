package internal

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"testing"
	"unicode"
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

	t *testing.T
}

func NewTestingErrorListener(t *testing.T) *TestingErrorListener {
	return &TestingErrorListener{t: t}
}

func (l *TestingErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.t.Errorf("SyntaxError %d:%d: %s", line, column, msg)
}

// CaseChangingStream wraps an existing CharStream, but upper cases, or lower cases
// the input before it is tokenized.
type CaseChangingStream struct {
	antlr.CharStream

	upper bool
}

func NewCaseChangingStream(in antlr.CharStream, upper bool) *CaseChangingStream {
	return &CaseChangingStream{
		in, upper,
	}
}

func (is *CaseChangingStream) LA(offset int) int {
	in := is.CharStream.LA(offset)
	if in < 0 {
		// Such as antlr.TokenEOF which is -1
		return in
	}
	if is.upper {
		return int(unicode.ToUpper(rune(in)))
	}
	return int(unicode.ToLower(rune(in)))
}

/*
func (is *CaseChangingStream) convertString(in string) string {
	fmt.Printf("here\n")
	if is.upper {
		return strings.ToUpper(in)
	}
	return strings.ToLower(in)
}

func (is *CaseChangingStream) GetText(start int, stop int) string {
	return is.convertString(is.CharStream.GetText(start, stop))
}

func (is *CaseChangingStream) GetTextFromTokens(start, end antlr.Token) string {
	return is.convertString(is.CharStream.GetTextFromTokens(start, end))
}

func (is *CaseChangingStream) GetTextFromInterval(i *antlr.Interval) string {
	return is.convertString(is.CharStream.GetTextFromInterval(i))
}
*/
