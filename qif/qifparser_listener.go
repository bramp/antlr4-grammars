// Code generated from qifParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package qif // qifParser
import "github.com/antlr/antlr4/runtime/Go/antlr"

// qifParserListener is a complete listener for a parse tree produced by qifParser.
type qifParserListener interface {
	antlr.ParseTreeListener

	// EnterQif is called when entering the qif production.
	EnterQif(c *QifContext)

	// EnterRecord is called when entering the record production.
	EnterRecord(c *RecordContext)

	// EnterRecordtype is called when entering the recordtype production.
	EnterRecordtype(c *RecordtypeContext)

	// EnterDate is called when entering the date production.
	EnterDate(c *DateContext)

	// EnterTotal is called when entering the total production.
	EnterTotal(c *TotalContext)

	// EnterCheck is called when entering the check production.
	EnterCheck(c *CheckContext)

	// EnterState_ is called when entering the state_ production.
	EnterState_(c *State_Context)

	// EnterMemo is called when entering the memo production.
	EnterMemo(c *MemoContext)

	// EnterPayee is called when entering the payee production.
	EnterPayee(c *PayeeContext)

	// EnterAccountorcategory is called when entering the accountorcategory production.
	EnterAccountorcategory(c *AccountorcategoryContext)

	// EnterAccount is called when entering the account production.
	EnterAccount(c *AccountContext)

	// EnterCategory is called when entering the category production.
	EnterCategory(c *CategoryContext)

	// EnterEor is called when entering the eor production.
	EnterEor(c *EorContext)

	// ExitQif is called when exiting the qif production.
	ExitQif(c *QifContext)

	// ExitRecord is called when exiting the record production.
	ExitRecord(c *RecordContext)

	// ExitRecordtype is called when exiting the recordtype production.
	ExitRecordtype(c *RecordtypeContext)

	// ExitDate is called when exiting the date production.
	ExitDate(c *DateContext)

	// ExitTotal is called when exiting the total production.
	ExitTotal(c *TotalContext)

	// ExitCheck is called when exiting the check production.
	ExitCheck(c *CheckContext)

	// ExitState_ is called when exiting the state_ production.
	ExitState_(c *State_Context)

	// ExitMemo is called when exiting the memo production.
	ExitMemo(c *MemoContext)

	// ExitPayee is called when exiting the payee production.
	ExitPayee(c *PayeeContext)

	// ExitAccountorcategory is called when exiting the accountorcategory production.
	ExitAccountorcategory(c *AccountorcategoryContext)

	// ExitAccount is called when exiting the account production.
	ExitAccount(c *AccountContext)

	// ExitCategory is called when exiting the category production.
	ExitCategory(c *CategoryContext)

	// ExitEor is called when exiting the eor production.
	ExitEor(c *EorContext)
}
