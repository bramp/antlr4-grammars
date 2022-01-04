// Code generated from gff3.g4 by ANTLR 4.9.3. DO NOT EDIT.

package gff3 // gff3
import "github.com/antlr/antlr4/runtime/Go/antlr"

// Basegff3Listener is a complete listener for a parse tree produced by gff3Parser.
type Basegff3Listener struct{}

var _ gff3Listener = &Basegff3Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *Basegff3Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Basegff3Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Basegff3Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Basegff3Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *Basegff3Listener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *Basegff3Listener) ExitDocument(ctx *DocumentContext) {}

// EnterLine is called when production line is entered.
func (s *Basegff3Listener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *Basegff3Listener) ExitLine(ctx *LineContext) {}

// EnterDataline is called when production dataline is entered.
func (s *Basegff3Listener) EnterDataline(ctx *DatalineContext) {}

// ExitDataline is called when production dataline is exited.
func (s *Basegff3Listener) ExitDataline(ctx *DatalineContext) {}

// EnterAttributes is called when production attributes is entered.
func (s *Basegff3Listener) EnterAttributes(ctx *AttributesContext) {}

// ExitAttributes is called when production attributes is exited.
func (s *Basegff3Listener) ExitAttributes(ctx *AttributesContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *Basegff3Listener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *Basegff3Listener) ExitAttribute(ctx *AttributeContext) {}

// EnterSeqid is called when production seqid is entered.
func (s *Basegff3Listener) EnterSeqid(ctx *SeqidContext) {}

// ExitSeqid is called when production seqid is exited.
func (s *Basegff3Listener) ExitSeqid(ctx *SeqidContext) {}

// EnterSource is called when production source is entered.
func (s *Basegff3Listener) EnterSource(ctx *SourceContext) {}

// ExitSource is called when production source is exited.
func (s *Basegff3Listener) ExitSource(ctx *SourceContext) {}

// EnterType_ is called when production type_ is entered.
func (s *Basegff3Listener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *Basegff3Listener) ExitType_(ctx *Type_Context) {}

// EnterStart is called when production start is entered.
func (s *Basegff3Listener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *Basegff3Listener) ExitStart(ctx *StartContext) {}

// EnterEnd is called when production end is entered.
func (s *Basegff3Listener) EnterEnd(ctx *EndContext) {}

// ExitEnd is called when production end is exited.
func (s *Basegff3Listener) ExitEnd(ctx *EndContext) {}

// EnterStrand is called when production strand is entered.
func (s *Basegff3Listener) EnterStrand(ctx *StrandContext) {}

// ExitStrand is called when production strand is exited.
func (s *Basegff3Listener) ExitStrand(ctx *StrandContext) {}

// EnterScore is called when production score is entered.
func (s *Basegff3Listener) EnterScore(ctx *ScoreContext) {}

// ExitScore is called when production score is exited.
func (s *Basegff3Listener) ExitScore(ctx *ScoreContext) {}

// EnterPhase is called when production phase is entered.
func (s *Basegff3Listener) EnterPhase(ctx *PhaseContext) {}

// ExitPhase is called when production phase is exited.
func (s *Basegff3Listener) ExitPhase(ctx *PhaseContext) {}

// EnterCommentline is called when production commentline is entered.
func (s *Basegff3Listener) EnterCommentline(ctx *CommentlineContext) {}

// ExitCommentline is called when production commentline is exited.
func (s *Basegff3Listener) ExitCommentline(ctx *CommentlineContext) {}
