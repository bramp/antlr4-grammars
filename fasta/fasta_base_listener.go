// Code generated from fasta.g4 by ANTLR 4.9.3. DO NOT EDIT.

package fasta // fasta
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasefastaListener is a complete listener for a parse tree produced by fastaParser.
type BasefastaListener struct{}

var _ fastaListener = &BasefastaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasefastaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasefastaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasefastaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasefastaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSequence is called when production sequence is entered.
func (s *BasefastaListener) EnterSequence(ctx *SequenceContext) {}

// ExitSequence is called when production sequence is exited.
func (s *BasefastaListener) ExitSequence(ctx *SequenceContext) {}

// EnterSection is called when production section is entered.
func (s *BasefastaListener) EnterSection(ctx *SectionContext) {}

// ExitSection is called when production section is exited.
func (s *BasefastaListener) ExitSection(ctx *SectionContext) {}

// EnterSequencelines is called when production sequencelines is entered.
func (s *BasefastaListener) EnterSequencelines(ctx *SequencelinesContext) {}

// ExitSequencelines is called when production sequencelines is exited.
func (s *BasefastaListener) ExitSequencelines(ctx *SequencelinesContext) {}

// EnterDescriptionline is called when production descriptionline is entered.
func (s *BasefastaListener) EnterDescriptionline(ctx *DescriptionlineContext) {}

// ExitDescriptionline is called when production descriptionline is exited.
func (s *BasefastaListener) ExitDescriptionline(ctx *DescriptionlineContext) {}

// EnterCommentline is called when production commentline is entered.
func (s *BasefastaListener) EnterCommentline(ctx *CommentlineContext) {}

// ExitCommentline is called when production commentline is exited.
func (s *BasefastaListener) ExitCommentline(ctx *CommentlineContext) {}
