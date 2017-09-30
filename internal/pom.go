package internal

import (
	"bufio"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const GRAMMARS_ROOT = "grammars-v4" // TODO REMOVE THIS

// Pom represents one of the pom.xml
// TODO Rename struct to say "Project", since it represents more than the Pom.
type Pom struct {
	Name     string   // shortname of this grammar (ususally the directory name)
	LongName string   // long name, usually very similar to the shortname.
	Includes []string // list of g4 files
	Grammars []*Grammar

	// Test related info
	EntryPoint          string
	ExampleRoot         string
	Examples            []string
	CaseInsensitiveType string
}

func (p *Pom) findGrammarOfType(t string) *Grammar {
	for _, g := range p.Grammars {
		if g.Type == t {
			return g
		}
	}
	return nil
}

// ParserName returns the name of the generated Parser.
func (p *Pom) ParserName() string {
	return p.grammarParserName() + "Parser"
}

// LexerName returns the name of the generated Lexer.
func (p *Pom) LexerName() string {
	return p.grammarLexerName() + "Lexer"
}

func (p *Pom) ListenerName() string {
	return p.grammarParserName() + "Listener"
}

// FilePrefix returns the filename prefix for the generated files.
func (p *Pom) FilePrefix() string {
	return strings.ToLower(p.grammarParserName())
}

// grammarParserName returns the name parser grammar.
func (p *Pom) grammarParserName() string {
	if g := p.findGrammarOfType("PARSER"); g != nil {
		return strings.TrimSuffix(g.Name, "Parser")
	}

	if g := p.findGrammarOfType("COMBINED"); g != nil {
		return g.Name
	}

	return p.LongName
}

// grammarLexerName returns the name lexer grammar.
func (p *Pom) grammarLexerName() string {
	if g := p.findGrammarOfType("LEXER"); g != nil {
		return strings.TrimSuffix(g.Name, "Lexer")
	}

	if g := p.findGrammarOfType("COMBINED"); g != nil {
		return g.Name
	}

	return p.LongName
}

// GeneratedFilenames returns the list of generated files.
func (p *Pom) GeneratedFilenames() []string {
	// Based on the code at:
	// https://github.com/antlr/antlr4/blob/46b3aa98cc8d8b6908c2cabb64a9587b6b973e6c/tool/src/org/antlr/v4/codegen/target/GoTarget.java#L146
	var files []string
	for _, g := range p.Grammars {
		name := g.Name
		switch g.Type {
		case "PARSER":
			name = strings.TrimSuffix(name, "Parser")
			files = append(files, strings.ToLower(name)+"_parser.go")
		case "LEXER":
			name = strings.TrimSuffix(name, "Lexer")
			files = append(files, strings.ToLower(name)+"_lexer.go")
		case "COMBINED":
			files = append(files, strings.ToLower(name)+"_parser.go")
			files = append(files, strings.ToLower(name)+"_lexer.go")
		default:
			panic(fmt.Sprintf("unknown grammar type %q", g.Type))
		}
	}

	return files
}

// Grammar represents a Antlr G4 grammar file.
type Grammar struct {
	Name string // name of this grammar
	Type string // one of PARSER, LEXER or COMBINED // TODO(bramp): Change to enum.
}

func ParseG4(path string) (*Grammar, error) {
	// TODO(bramp) Use a proper antlr4 parser

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		t := ""
		if strings.HasPrefix(line, "grammar") {
			t = "COMBINED"
		} else if strings.HasPrefix(line, "lexer") {
			t = "LEXER"
		} else if strings.HasPrefix(line, "parser") {
			t = "PARSER"
		}

		if t != "" {
			parts := strings.Fields(strings.TrimSuffix(line, ";"))
			if len(parts) < 2 {
				return nil, fmt.Errorf("failed to parse grammar name: %q", line)
			}
			return &Grammar{
				Name: parts[len(parts)-1],
				Type: t,
			}, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return nil, errors.New("failed to find fields of interest in grammar")
}

func contains(haystack []string, needle string) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}
	return false
}

// ParsePom extracts information about the grammar in a very lazy way!
func ParsePom(path string) (*Pom, error) {
	dir := filepath.Dir(path)
	name := strings.TrimPrefix(strings.TrimPrefix(dir, GRAMMARS_ROOT), "/")

	p := &Pom{
		Name: name, // TODO Read name from pom
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	decoder := xml.NewDecoder(file)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "grammars", "include":
				var file string
				if err := decoder.DecodeElement(&file, &se); err != nil {
					return nil, err
				}
				file = filepath.Join(dir, file)
				if _, err := os.Stat(file); os.IsNotExist(err) {
					log.Printf("missing grammar %q referenced in %q", file, path)
				} else {
					// Ignore dups
					if contains(p.Includes, file) {
						continue
					}

					p.Includes = append(p.Includes, file)

					if g, err := ParseG4(file); err != nil {
						log.Printf("failed to parse grammar %q: %s", file, err)
					} else {
						p.Grammars = append(p.Grammars, g)
					}
				}

			case "grammarName":
				var longName string
				if err := decoder.DecodeElement(&longName, &se); err != nil {
					return nil, err
				}
				p.LongName = longName

			case "entryPoint":
				var entryPoint string
				if err := decoder.DecodeElement(&entryPoint, &se); err != nil {
					return nil, err
				}
				p.EntryPoint = entryPoint

			case "exampleFiles":
				var file string
				if err := decoder.DecodeElement(&file, &se); err != nil {
					return nil, err
				}
				p.ExampleRoot = filepath.Join(dir, file)
				examples, err := filepath.Glob(filepath.Join(p.ExampleRoot, "*"))
				if err != nil {
					return nil, err
				}
				p.Examples = examples

			case "caseInsensitiveType":
				var caseInsensitiveType string
				if err := decoder.DecodeElement(&caseInsensitiveType, &se); err != nil {
					return nil, err
				}
				p.CaseInsensitiveType = caseInsensitiveType
			}
		}
	}

	return p, nil
}
