// Code generated from fol.g4 by ANTLR 4.9.3. DO NOT EDIT.

package fol // fol
import "github.com/antlr/antlr4/runtime/Go/antlr"

// folListener is a complete listener for a parse tree produced by folParser.
type folListener interface {
	antlr.ParseTreeListener

	// EnterCondition is called when entering the condition production.
	EnterCondition(c *ConditionContext)

	// EnterFormula is called when entering the formula production.
	EnterFormula(c *FormulaContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterBin_connective is called when entering the bin_connective production.
	EnterBin_connective(c *Bin_connectiveContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterPred_constant is called when entering the pred_constant production.
	EnterPred_constant(c *Pred_constantContext)

	// EnterInd_constant is called when entering the ind_constant production.
	EnterInd_constant(c *Ind_constantContext)

	// EnterFunc_constant is called when entering the func_constant production.
	EnterFunc_constant(c *Func_constantContext)

	// EnterSeparator is called when entering the separator production.
	EnterSeparator(c *SeparatorContext)

	// ExitCondition is called when exiting the condition production.
	ExitCondition(c *ConditionContext)

	// ExitFormula is called when exiting the formula production.
	ExitFormula(c *FormulaContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitBin_connective is called when exiting the bin_connective production.
	ExitBin_connective(c *Bin_connectiveContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitPred_constant is called when exiting the pred_constant production.
	ExitPred_constant(c *Pred_constantContext)

	// ExitInd_constant is called when exiting the ind_constant production.
	ExitInd_constant(c *Ind_constantContext)

	// ExitFunc_constant is called when exiting the func_constant production.
	ExitFunc_constant(c *Func_constantContext)

	// ExitSeparator is called when exiting the separator production.
	ExitSeparator(c *SeparatorContext)
}
