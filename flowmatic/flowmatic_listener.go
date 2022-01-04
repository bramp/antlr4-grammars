// Code generated from flowmatic.g4 by ANTLR 4.9.3. DO NOT EDIT.

package flowmatic // flowmatic
import "github.com/antlr/antlr4/runtime/Go/antlr"

// flowmaticListener is a complete listener for a parse tree produced by flowmaticParser.
type flowmaticListener interface {
	antlr.ParseTreeListener

	// EnterFlowmatic is called when entering the flowmatic production.
	EnterFlowmatic(c *FlowmaticContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterHsp_statement is called when entering the hsp_statement production.
	EnterHsp_statement(c *Hsp_statementContext)

	// EnterOutput_statement is called when entering the output_statement production.
	EnterOutput_statement(c *Output_statementContext)

	// EnterInput_statement is called when entering the input_statement production.
	EnterInput_statement(c *Input_statementContext)

	// EnterJumpto_statement is called when entering the jumpto_statement production.
	EnterJumpto_statement(c *Jumpto_statementContext)

	// EnterReaditem_statement is called when entering the readitem_statement production.
	EnterReaditem_statement(c *Readitem_statementContext)

	// EnterWriteitem_statement is called when entering the writeitem_statement production.
	EnterWriteitem_statement(c *Writeitem_statementContext)

	// EnterTransfer_statement is called when entering the transfer_statement production.
	EnterTransfer_statement(c *Transfer_statementContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterOtherwise_statement is called when entering the otherwise_statement production.
	EnterOtherwise_statement(c *Otherwise_statementContext)

	// EnterRewind_statement is called when entering the rewind_statement production.
	EnterRewind_statement(c *Rewind_statementContext)

	// EnterStop_statement is called when entering the stop_statement production.
	EnterStop_statement(c *Stop_statementContext)

	// EnterCompare_statement is called when entering the compare_statement production.
	EnterCompare_statement(c *Compare_statementContext)

	// EnterMove_statement is called when entering the move_statement production.
	EnterMove_statement(c *Move_statementContext)

	// EnterSet_statement is called when entering the set_statement production.
	EnterSet_statement(c *Set_statementContext)

	// EnterTest_statement is called when entering the test_statement production.
	EnterTest_statement(c *Test_statementContext)

	// EnterCloseout_statement is called when entering the closeout_statement production.
	EnterCloseout_statement(c *Closeout_statementContext)

	// EnterNum is called when entering the num production.
	EnterNum(c *NumContext)

	// EnterFilename is called when entering the filename production.
	EnterFilename(c *FilenameContext)

	// EnterFileletter is called when entering the fileletter production.
	EnterFileletter(c *FileletterContext)

	// EnterFieldname is called when entering the fieldname production.
	EnterFieldname(c *FieldnameContext)

	// EnterOperation is called when entering the operation production.
	EnterOperation(c *OperationContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// ExitFlowmatic is called when exiting the flowmatic production.
	ExitFlowmatic(c *FlowmaticContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitHsp_statement is called when exiting the hsp_statement production.
	ExitHsp_statement(c *Hsp_statementContext)

	// ExitOutput_statement is called when exiting the output_statement production.
	ExitOutput_statement(c *Output_statementContext)

	// ExitInput_statement is called when exiting the input_statement production.
	ExitInput_statement(c *Input_statementContext)

	// ExitJumpto_statement is called when exiting the jumpto_statement production.
	ExitJumpto_statement(c *Jumpto_statementContext)

	// ExitReaditem_statement is called when exiting the readitem_statement production.
	ExitReaditem_statement(c *Readitem_statementContext)

	// ExitWriteitem_statement is called when exiting the writeitem_statement production.
	ExitWriteitem_statement(c *Writeitem_statementContext)

	// ExitTransfer_statement is called when exiting the transfer_statement production.
	ExitTransfer_statement(c *Transfer_statementContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitOtherwise_statement is called when exiting the otherwise_statement production.
	ExitOtherwise_statement(c *Otherwise_statementContext)

	// ExitRewind_statement is called when exiting the rewind_statement production.
	ExitRewind_statement(c *Rewind_statementContext)

	// ExitStop_statement is called when exiting the stop_statement production.
	ExitStop_statement(c *Stop_statementContext)

	// ExitCompare_statement is called when exiting the compare_statement production.
	ExitCompare_statement(c *Compare_statementContext)

	// ExitMove_statement is called when exiting the move_statement production.
	ExitMove_statement(c *Move_statementContext)

	// ExitSet_statement is called when exiting the set_statement production.
	ExitSet_statement(c *Set_statementContext)

	// ExitTest_statement is called when exiting the test_statement production.
	ExitTest_statement(c *Test_statementContext)

	// ExitCloseout_statement is called when exiting the closeout_statement production.
	ExitCloseout_statement(c *Closeout_statementContext)

	// ExitNum is called when exiting the num production.
	ExitNum(c *NumContext)

	// ExitFilename is called when exiting the filename production.
	ExitFilename(c *FilenameContext)

	// ExitFileletter is called when exiting the fileletter production.
	ExitFileletter(c *FileletterContext)

	// ExitFieldname is called when exiting the fieldname production.
	ExitFieldname(c *FieldnameContext)

	// ExitOperation is called when exiting the operation production.
	ExitOperation(c *OperationContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)
}
