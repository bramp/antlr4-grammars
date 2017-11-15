// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

type grammarType string

const LEXER = grammarType("LEXER")
const PARSER = grammarType("PARSER")
const COMBINED = grammarType("COMBINED")

// Project represents one of language grammars defined by a pom.xml file and a set of g4 files.
type Project struct {
	FileName string // Filename of the pom.xml.

	LongName string     // Name of the grammar defined in the pom.xml
	Includes []string   // List of included g4 files
	Grammars []*Grammar // Parsed grammars

	// Test related info
	EntryPoint          string
	Examples            []string
	CaseInsensitiveType string

	FoundAntlr4MavenPlugin bool // Did we find the Antlr Maven plugin?
}

func (p *Project) findGrammarOfType(t grammarType) *Grammar {
	for _, g := range p.Grammars {
		if g.Type == t {
			return g
		}
	}
	return nil
}

func (p *Project) HasParser() bool {
	for _, g := range p.Grammars {
		if g.Type == PARSER || g.Type == COMBINED {
			return true
		}
	}
	return false
}

func (p *Project) HasLexer() bool {
	for _, g := range p.Grammars {
		if g.Type == LEXER || g.Type == COMBINED {
			return true
		}
	}
	return false
}

// ParserName returns the name of the generated Parser.
func (p *Project) ParserName() string {
	if g := p.findGrammarOfType(PARSER); g != nil {
		return strings.TrimSuffix(g.Name, "Parser") + "Parser"
	}

	if g := p.findGrammarOfType(COMBINED); g != nil {
		return g.Name + "Parser"
	}

	panic(fmt.Sprintf("%q does not contain a parser", p.FileName))
}

// LexerName returns the name of the generated Lexer.
func (p *Project) LexerName() string {
	if g := p.findGrammarOfType(LEXER); g != nil {
		return g.Name
	}

	if g := p.findGrammarOfType(COMBINED); g != nil {
		return g.Name + "Lexer"
	}

	panic(fmt.Sprintf("%q does not contain a lexer: %#v", p.FileName, p.Grammars))
}

// ListenerName returns the name of the of the generated Listener.
// See https://github.com/antlr/antlr4/blob/master/tool/src/org/antlr/v4/codegen/target/GoTarget.java#L168
func (p *Project) ListenerName() string {
	if g := p.findGrammarOfType(PARSER); g != nil {
		return g.Name + "Listener"
	}

	if g := p.findGrammarOfType(COMBINED); g != nil {
		return g.Name + "Listener"
	}

	panic(fmt.Sprintf("%q does not contain a parser", p.FileName))
}

// GeneratedFilenames returns the list of generated files.
func (p *Project) GeneratedFilenames() []string {
	// Based on the code at:
	// https://github.com/antlr/antlr4/blob/46b3aa98cc8d8b6908c2cabb64a9587b6b973e6c/tool/src/org/antlr/v4/codegen/target/GoTarget.java#L146
	var files []string
	for _, g := range p.Grammars {
		files = append(files, g.GeneratedFilenames()...)
	}
	return files
}

// Grammar represents a Antlr G4 grammar file.
type Grammar struct {
	Name     string // name of this grammar
	Filename string
	Type     grammarType // one of PARSER, LEXER or COMBINED
}

func (g *Grammar) String() string {
	return fmt.Sprintf("%s: %s", g.Type, g.Name)
}

func (g *Grammar) DependentFilenames() []string {
	var files []string
	if g.Type == "PARSER" {
		// Depend on the generated lexer
		name := strings.ToLower(strings.TrimSuffix(g.Name, "Parser"))
		files = append(files, name+"_lexer.go")
	}
	return files
}

// GeneratedFilenames returns the list of generated files.
func (g *Grammar) GeneratedFilenames() []string {
	// Based on the code at:
	// https://github.com/antlr/antlr4/blob/46b3aa98cc8d8b6908c2cabb64a9587b6b973e6c/tool/src/org/antlr/v4/codegen/target/GoTarget.java#L146
	var files []string
	switch g.Type {
	case LEXER:
		name := strings.ToLower(strings.TrimSuffix(g.Name, "Lexer"))
		files = append(files, name+"_lexer.go")

	case PARSER:
		name := strings.ToLower(g.Name)
		files = append(files, name+"_base_listener.go", name+"_listener.go")

		name = strings.ToLower(strings.TrimSuffix(g.Name, "Parser"))
		files = append(files, name+"_parser.go")

	case COMBINED:
		name := strings.ToLower(g.Name)
		files = append(files, name+"_base_listener.go", name+"_listener.go")
		files = append(files, name+"_parser.go", name+"_lexer.go")

	default:
		panic(fmt.Sprintf("unknown grammar type %q", g.Type))
	}

	return files
}

func ParseG4(path string) (*Grammar, error) {
	// TODO(bramp) Use a proper antlr4 parser

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var t grammarType

		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "grammar") {
			t = COMBINED
		} else if strings.HasPrefix(line, "lexer") {
			t = LEXER
		} else if strings.HasPrefix(line, "parser") {
			t = PARSER
		}

		if t != "" {
			if semi := strings.Index(line, ";"); semi >= 0 {
				line = line[:semi]
			}
			parts := strings.Fields(line)
			if len(parts) < 2 {
				return nil, fmt.Errorf("failed to parse grammar name: %q", line)
			}
			return &Grammar{
				Name:     parts[len(parts)-1],
				Filename: path,
				Type:     t,
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

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func (p *Project) AddGrammar(filename string) {
	// HACKS: A few hacks to the file names, to accomidate odd cases in the pom.xml
	// "Upgrade" the file to a GoTarget specific one (if it exists)
	if betterfile := strings.Replace(filename, ".g4", ".GoTarget.g4", -1); fileExists(betterfile) {
		filename = betterfile
	}

	if !fileExists(filename) {
		log.Printf("missing grammar %q", filename)
		return
	}

	// Ignore dups
	if contains(p.Includes, filename) {
		return
	}

	p.Includes = append(p.Includes, filename)

	if g, err := ParseG4(filename); err != nil {
		log.Printf("failed to parse grammar %q: %s", filename, err)
	} else {
		p.Grammars = append(p.Grammars, g)
	}
}

// ParsePom extracts information about the grammar in a very lazy way!
func ParsePom(path string) (*Project, error) {
	p := &Project{
		FileName: path,
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	dir := filepath.Dir(path)

	decoder := xml.NewDecoder(file)
	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "artifactId":
				var name string
				if err := decoder.DecodeElement(&name, &se); err != nil {
					return nil, err
				}
				if name == "antlr4-maven-plugin" {
					p.FoundAntlr4MavenPlugin = true
				}

			case "grammars", "include":
				var file string
				if err := decoder.DecodeElement(&file, &se); err != nil {
					return nil, err
				}
				p.AddGrammar(filepath.Join(dir, file))

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

				// TODO(bramp): Instead of glob'ing, recurse deeper (since some examples are nested, e.g vb6)
				examples, err := filepath.Glob(filepath.Join(dir, file, "*"))
				if err != nil {
					return nil, err
				}

				var filtered []string
				for _, example := range examples {
					if strings.HasSuffix(example, ".tree") {
						continue
					}
					if strings.HasSuffix(example, ".errors") {
						continue
					}

					filtered = append(filtered, example)
				}

				p.Examples = filtered

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
