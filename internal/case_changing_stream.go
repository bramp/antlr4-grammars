package internal

import (
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// CaseChangingStream wraps an existing CharStream, but upper cases, or lower
// cases the input before it is tokenized.
//
// It would be useful if this was part of the core antlr library, but it was
// rejected in https://github.com/antlr/antlr4/pull/2046 as users can trivial
// strings.ToUpper(...) the input, or write their own version of this stream.
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
