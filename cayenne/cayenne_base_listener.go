// Code generated from cayenne.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cayenne // cayenne
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasecayenneListener is a complete listener for a parse tree produced by cayenneParser.
type BasecayenneListener struct{}

var _ cayenneListener = &BasecayenneListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasecayenneListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasecayenneListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasecayenneListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasecayenneListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterExpr is called when production expr is entered.
func (s *BasecayenneListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BasecayenneListener) ExitExpr(ctx *ExprContext) {}

// EnterArm is called when production arm is entered.
func (s *BasecayenneListener) EnterArm(ctx *ArmContext) {}

// ExitArm is called when production arm is exited.
func (s *BasecayenneListener) ExitArm(ctx *ArmContext) {}

// EnterSign is called when production sign is entered.
func (s *BasecayenneListener) EnterSign(ctx *SignContext) {}

// ExitSign is called when production sign is exited.
func (s *BasecayenneListener) ExitSign(ctx *SignContext) {}

// EnterDefn is called when production defn is entered.
func (s *BasecayenneListener) EnterDefn(ctx *DefnContext) {}

// ExitDefn is called when production defn is exited.
func (s *BasecayenneListener) ExitDefn(ctx *DefnContext) {}

// EnterVis is called when production vis is entered.
func (s *BasecayenneListener) EnterVis(ctx *VisContext) {}

// ExitVis is called when production vis is exited.
func (s *BasecayenneListener) ExitVis(ctx *VisContext) {}

// EnterAbs_ is called when production abs_ is entered.
func (s *BasecayenneListener) EnterAbs_(ctx *Abs_Context) {}

// ExitAbs_ is called when production abs_ is exited.
func (s *BasecayenneListener) ExitAbs_(ctx *Abs_Context) {}

// EnterType_ is called when production type_ is entered.
func (s *BasecayenneListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BasecayenneListener) ExitType_(ctx *Type_Context) {}

// EnterVarid is called when production varid is entered.
func (s *BasecayenneListener) EnterVarid(ctx *VaridContext) {}

// ExitVarid is called when production varid is exited.
func (s *BasecayenneListener) ExitVarid(ctx *VaridContext) {}

// EnterConid is called when production conid is entered.
func (s *BasecayenneListener) EnterConid(ctx *ConidContext) {}

// ExitConid is called when production conid is exited.
func (s *BasecayenneListener) ExitConid(ctx *ConidContext) {}

// EnterLblid is called when production lblid is entered.
func (s *BasecayenneListener) EnterLblid(ctx *LblidContext) {}

// ExitLblid is called when production lblid is exited.
func (s *BasecayenneListener) ExitLblid(ctx *LblidContext) {}

// EnterId_ is called when production id_ is entered.
func (s *BasecayenneListener) EnterId_(ctx *Id_Context) {}

// ExitId_ is called when production id_ is exited.
func (s *BasecayenneListener) ExitId_(ctx *Id_Context) {}
