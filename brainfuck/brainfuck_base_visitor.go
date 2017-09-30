// Generated from brainfuck.g4 by ANTLR 4.7.

package brainfuck // brainfuck
import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasebrainfuckVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasebrainfuckVisitor) VisitFile(ctx *FileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasebrainfuckVisitor) VisitOpcode(ctx *OpcodeContext) interface{} {
	return v.VisitChildren(ctx)
}
