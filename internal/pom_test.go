// TODO(bramp) Change the test data to be seperate from the real data.
package internal

import (
	"path/filepath"
	"testing"
)

const ROOT = "../grammars-v4/" // Path to grammars

func TestParseG4(t *testing.T) {
	tests := []struct {
		g4   string
		name string
	}{
		{g4: "abnf/Abnf.g4", name: "Abnf"},
		{g4: "cobol85/Cobol85.g4", name: "Cobol85"},
		{g4: "cobol85/Cobol85Preprocessor.g4", name: "Cobol85Preprocessor"},
		{g4: "kotlin/UnicodeClasses.g4", name: "UnicodeClasses"},
	}

	for _, test := range tests {
		p, err := ParseG4(filepath.Join(ROOT, test.g4))
		if err != nil {
			t.Errorf("ParseG4(%q) err = %q, want nil", test.g4, err)
			continue
		}

		if got := p.Name; got != test.name {
			t.Errorf("ParseG4(%q).Name = %q, want %q", test.g4, got, test.name)
		}
	}
}

func TestParsePom(t *testing.T) {
	tests := []struct {
		pom   string
		name  string
		lexer string
	}{
		{pom: "abnf/pom.xml", name: "Abnf", lexer: "AbnfLexer"},
	}

	for _, test := range tests {
		p, err := ParsePom(filepath.Join(ROOT, test.pom))
		if err != nil {
			t.Errorf("ParseG4(%q) err = %q, want nil", test.pom, err)
			continue
		}

		if got := p.LexerName(); got != test.lexer {
			t.Errorf("ParseG4(%q).LexerName() = %q, want %q", test.pom, got, test.lexer)
		}
	}
}
