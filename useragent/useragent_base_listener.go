// Code generated from useragent.g4 by ANTLR 4.7.2. DO NOT EDIT.

package useragent // useragent
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseuseragentListener is a complete listener for a parse tree produced by useragentParser.
type BaseuseragentListener struct{}

var _ useragentListener = &BaseuseragentListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseuseragentListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseuseragentListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseuseragentListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseuseragentListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseuseragentListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseuseragentListener) ExitProg(ctx *ProgContext) {}

// EnterProduct is called when production product is entered.
func (s *BaseuseragentListener) EnterProduct(ctx *ProductContext) {}

// ExitProduct is called when production product is exited.
func (s *BaseuseragentListener) ExitProduct(ctx *ProductContext) {}

// EnterName is called when production name is entered.
func (s *BaseuseragentListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseuseragentListener) ExitName(ctx *NameContext) {}

// EnterVersion is called when production version is entered.
func (s *BaseuseragentListener) EnterVersion(ctx *VersionContext) {}

// ExitVersion is called when production version is exited.
func (s *BaseuseragentListener) ExitVersion(ctx *VersionContext) {}

// EnterComment is called when production comment is entered.
func (s *BaseuseragentListener) EnterComment(ctx *CommentContext) {}

// ExitComment is called when production comment is exited.
func (s *BaseuseragentListener) ExitComment(ctx *CommentContext) {}
