// Code generated from flowmatic.g4 by ANTLR 4.9.3. DO NOT EDIT.

package flowmatic // flowmatic
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseflowmaticListener is a complete listener for a parse tree produced by flowmaticParser.
type BaseflowmaticListener struct{}

var _ flowmaticListener = &BaseflowmaticListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseflowmaticListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseflowmaticListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseflowmaticListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseflowmaticListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFlowmatic is called when production flowmatic is entered.
func (s *BaseflowmaticListener) EnterFlowmatic(ctx *FlowmaticContext) {}

// ExitFlowmatic is called when production flowmatic is exited.
func (s *BaseflowmaticListener) ExitFlowmatic(ctx *FlowmaticContext) {}

// EnterLine is called when production line is entered.
func (s *BaseflowmaticListener) EnterLine(ctx *LineContext) {}

// ExitLine is called when production line is exited.
func (s *BaseflowmaticListener) ExitLine(ctx *LineContext) {}

// EnterLabel is called when production label is entered.
func (s *BaseflowmaticListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseflowmaticListener) ExitLabel(ctx *LabelContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseflowmaticListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseflowmaticListener) ExitStatement(ctx *StatementContext) {}

// EnterHsp_statement is called when production hsp_statement is entered.
func (s *BaseflowmaticListener) EnterHsp_statement(ctx *Hsp_statementContext) {}

// ExitHsp_statement is called when production hsp_statement is exited.
func (s *BaseflowmaticListener) ExitHsp_statement(ctx *Hsp_statementContext) {}

// EnterOutput_statement is called when production output_statement is entered.
func (s *BaseflowmaticListener) EnterOutput_statement(ctx *Output_statementContext) {}

// ExitOutput_statement is called when production output_statement is exited.
func (s *BaseflowmaticListener) ExitOutput_statement(ctx *Output_statementContext) {}

// EnterInput_statement is called when production input_statement is entered.
func (s *BaseflowmaticListener) EnterInput_statement(ctx *Input_statementContext) {}

// ExitInput_statement is called when production input_statement is exited.
func (s *BaseflowmaticListener) ExitInput_statement(ctx *Input_statementContext) {}

// EnterJumpto_statement is called when production jumpto_statement is entered.
func (s *BaseflowmaticListener) EnterJumpto_statement(ctx *Jumpto_statementContext) {}

// ExitJumpto_statement is called when production jumpto_statement is exited.
func (s *BaseflowmaticListener) ExitJumpto_statement(ctx *Jumpto_statementContext) {}

// EnterReaditem_statement is called when production readitem_statement is entered.
func (s *BaseflowmaticListener) EnterReaditem_statement(ctx *Readitem_statementContext) {}

// ExitReaditem_statement is called when production readitem_statement is exited.
func (s *BaseflowmaticListener) ExitReaditem_statement(ctx *Readitem_statementContext) {}

// EnterWriteitem_statement is called when production writeitem_statement is entered.
func (s *BaseflowmaticListener) EnterWriteitem_statement(ctx *Writeitem_statementContext) {}

// ExitWriteitem_statement is called when production writeitem_statement is exited.
func (s *BaseflowmaticListener) ExitWriteitem_statement(ctx *Writeitem_statementContext) {}

// EnterTransfer_statement is called when production transfer_statement is entered.
func (s *BaseflowmaticListener) EnterTransfer_statement(ctx *Transfer_statementContext) {}

// ExitTransfer_statement is called when production transfer_statement is exited.
func (s *BaseflowmaticListener) ExitTransfer_statement(ctx *Transfer_statementContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseflowmaticListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseflowmaticListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterOtherwise_statement is called when production otherwise_statement is entered.
func (s *BaseflowmaticListener) EnterOtherwise_statement(ctx *Otherwise_statementContext) {}

// ExitOtherwise_statement is called when production otherwise_statement is exited.
func (s *BaseflowmaticListener) ExitOtherwise_statement(ctx *Otherwise_statementContext) {}

// EnterRewind_statement is called when production rewind_statement is entered.
func (s *BaseflowmaticListener) EnterRewind_statement(ctx *Rewind_statementContext) {}

// ExitRewind_statement is called when production rewind_statement is exited.
func (s *BaseflowmaticListener) ExitRewind_statement(ctx *Rewind_statementContext) {}

// EnterStop_statement is called when production stop_statement is entered.
func (s *BaseflowmaticListener) EnterStop_statement(ctx *Stop_statementContext) {}

// ExitStop_statement is called when production stop_statement is exited.
func (s *BaseflowmaticListener) ExitStop_statement(ctx *Stop_statementContext) {}

// EnterCompare_statement is called when production compare_statement is entered.
func (s *BaseflowmaticListener) EnterCompare_statement(ctx *Compare_statementContext) {}

// ExitCompare_statement is called when production compare_statement is exited.
func (s *BaseflowmaticListener) ExitCompare_statement(ctx *Compare_statementContext) {}

// EnterMove_statement is called when production move_statement is entered.
func (s *BaseflowmaticListener) EnterMove_statement(ctx *Move_statementContext) {}

// ExitMove_statement is called when production move_statement is exited.
func (s *BaseflowmaticListener) ExitMove_statement(ctx *Move_statementContext) {}

// EnterSet_statement is called when production set_statement is entered.
func (s *BaseflowmaticListener) EnterSet_statement(ctx *Set_statementContext) {}

// ExitSet_statement is called when production set_statement is exited.
func (s *BaseflowmaticListener) ExitSet_statement(ctx *Set_statementContext) {}

// EnterTest_statement is called when production test_statement is entered.
func (s *BaseflowmaticListener) EnterTest_statement(ctx *Test_statementContext) {}

// ExitTest_statement is called when production test_statement is exited.
func (s *BaseflowmaticListener) ExitTest_statement(ctx *Test_statementContext) {}

// EnterCloseout_statement is called when production closeout_statement is entered.
func (s *BaseflowmaticListener) EnterCloseout_statement(ctx *Closeout_statementContext) {}

// ExitCloseout_statement is called when production closeout_statement is exited.
func (s *BaseflowmaticListener) ExitCloseout_statement(ctx *Closeout_statementContext) {}

// EnterNum is called when production num is entered.
func (s *BaseflowmaticListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production num is exited.
func (s *BaseflowmaticListener) ExitNum(ctx *NumContext) {}

// EnterFilename is called when production filename is entered.
func (s *BaseflowmaticListener) EnterFilename(ctx *FilenameContext) {}

// ExitFilename is called when production filename is exited.
func (s *BaseflowmaticListener) ExitFilename(ctx *FilenameContext) {}

// EnterFileletter is called when production fileletter is entered.
func (s *BaseflowmaticListener) EnterFileletter(ctx *FileletterContext) {}

// ExitFileletter is called when production fileletter is exited.
func (s *BaseflowmaticListener) ExitFileletter(ctx *FileletterContext) {}

// EnterFieldname is called when production fieldname is entered.
func (s *BaseflowmaticListener) EnterFieldname(ctx *FieldnameContext) {}

// ExitFieldname is called when production fieldname is exited.
func (s *BaseflowmaticListener) ExitFieldname(ctx *FieldnameContext) {}

// EnterOperation is called when production operation is entered.
func (s *BaseflowmaticListener) EnterOperation(ctx *OperationContext) {}

// ExitOperation is called when production operation is exited.
func (s *BaseflowmaticListener) ExitOperation(ctx *OperationContext) {}

// EnterOp is called when production op is entered.
func (s *BaseflowmaticListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseflowmaticListener) ExitOp(ctx *OpContext) {}
