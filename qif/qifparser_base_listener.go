// Code generated from qifParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package qif // qifParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseqifParserListener is a complete listener for a parse tree produced by qifParser.
type BaseqifParserListener struct{}

var _ qifParserListener = &BaseqifParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseqifParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseqifParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseqifParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseqifParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQif is called when production qif is entered.
func (s *BaseqifParserListener) EnterQif(ctx *QifContext) {}

// ExitQif is called when production qif is exited.
func (s *BaseqifParserListener) ExitQif(ctx *QifContext) {}

// EnterRecord is called when production record is entered.
func (s *BaseqifParserListener) EnterRecord(ctx *RecordContext) {}

// ExitRecord is called when production record is exited.
func (s *BaseqifParserListener) ExitRecord(ctx *RecordContext) {}

// EnterRecordtype is called when production recordtype is entered.
func (s *BaseqifParserListener) EnterRecordtype(ctx *RecordtypeContext) {}

// ExitRecordtype is called when production recordtype is exited.
func (s *BaseqifParserListener) ExitRecordtype(ctx *RecordtypeContext) {}

// EnterDate is called when production date is entered.
func (s *BaseqifParserListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BaseqifParserListener) ExitDate(ctx *DateContext) {}

// EnterTotal is called when production total is entered.
func (s *BaseqifParserListener) EnterTotal(ctx *TotalContext) {}

// ExitTotal is called when production total is exited.
func (s *BaseqifParserListener) ExitTotal(ctx *TotalContext) {}

// EnterCheck is called when production check is entered.
func (s *BaseqifParserListener) EnterCheck(ctx *CheckContext) {}

// ExitCheck is called when production check is exited.
func (s *BaseqifParserListener) ExitCheck(ctx *CheckContext) {}

// EnterState_ is called when production state_ is entered.
func (s *BaseqifParserListener) EnterState_(ctx *State_Context) {}

// ExitState_ is called when production state_ is exited.
func (s *BaseqifParserListener) ExitState_(ctx *State_Context) {}

// EnterMemo is called when production memo is entered.
func (s *BaseqifParserListener) EnterMemo(ctx *MemoContext) {}

// ExitMemo is called when production memo is exited.
func (s *BaseqifParserListener) ExitMemo(ctx *MemoContext) {}

// EnterPayee is called when production payee is entered.
func (s *BaseqifParserListener) EnterPayee(ctx *PayeeContext) {}

// ExitPayee is called when production payee is exited.
func (s *BaseqifParserListener) ExitPayee(ctx *PayeeContext) {}

// EnterAccountorcategory is called when production accountorcategory is entered.
func (s *BaseqifParserListener) EnterAccountorcategory(ctx *AccountorcategoryContext) {}

// ExitAccountorcategory is called when production accountorcategory is exited.
func (s *BaseqifParserListener) ExitAccountorcategory(ctx *AccountorcategoryContext) {}

// EnterAccount is called when production account is entered.
func (s *BaseqifParserListener) EnterAccount(ctx *AccountContext) {}

// ExitAccount is called when production account is exited.
func (s *BaseqifParserListener) ExitAccount(ctx *AccountContext) {}

// EnterCategory is called when production category is entered.
func (s *BaseqifParserListener) EnterCategory(ctx *CategoryContext) {}

// ExitCategory is called when production category is exited.
func (s *BaseqifParserListener) ExitCategory(ctx *CategoryContext) {}

// EnterEor is called when production eor is entered.
func (s *BaseqifParserListener) EnterEor(ctx *EorContext) {}

// ExitEor is called when production eor is exited.
func (s *BaseqifParserListener) ExitEor(ctx *EorContext) {}
