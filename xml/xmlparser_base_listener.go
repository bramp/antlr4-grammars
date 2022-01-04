// Code generated from XMLParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package xml // XMLParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseXMLParserListener is a complete listener for a parse tree produced by XMLParser.
type BaseXMLParserListener struct{}

var _ XMLParserListener = &BaseXMLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseXMLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseXMLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseXMLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseXMLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseXMLParserListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseXMLParserListener) ExitDocument(ctx *DocumentContext) {}

// EnterProlog is called when production prolog is entered.
func (s *BaseXMLParserListener) EnterProlog(ctx *PrologContext) {}

// ExitProlog is called when production prolog is exited.
func (s *BaseXMLParserListener) ExitProlog(ctx *PrologContext) {}

// EnterContent is called when production content is entered.
func (s *BaseXMLParserListener) EnterContent(ctx *ContentContext) {}

// ExitContent is called when production content is exited.
func (s *BaseXMLParserListener) ExitContent(ctx *ContentContext) {}

// EnterElement is called when production element is entered.
func (s *BaseXMLParserListener) EnterElement(ctx *ElementContext) {}

// ExitElement is called when production element is exited.
func (s *BaseXMLParserListener) ExitElement(ctx *ElementContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseXMLParserListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseXMLParserListener) ExitReference(ctx *ReferenceContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseXMLParserListener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseXMLParserListener) ExitAttribute(ctx *AttributeContext) {}

// EnterChardata is called when production chardata is entered.
func (s *BaseXMLParserListener) EnterChardata(ctx *ChardataContext) {}

// ExitChardata is called when production chardata is exited.
func (s *BaseXMLParserListener) ExitChardata(ctx *ChardataContext) {}

// EnterMisc is called when production misc is entered.
func (s *BaseXMLParserListener) EnterMisc(ctx *MiscContext) {}

// ExitMisc is called when production misc is exited.
func (s *BaseXMLParserListener) ExitMisc(ctx *MiscContext) {}
