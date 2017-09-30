// Generated from istc.g4 by ANTLR 4.7.

package istc // istc
import "github.com/antlr/antlr4/runtime/Go/antlr"

// istcListener is a complete listener for a parse tree produced by istcParser.
type istcListener interface {
	antlr.ParseTreeListener

	// EnterIstc is called when entering the istc production.
	EnterIstc(c *IstcContext)

	// EnterRegistration is called when entering the registration production.
	EnterRegistration(c *RegistrationContext)

	// EnterYear is called when entering the year production.
	EnterYear(c *YearContext)

	// EnterWork is called when entering the work production.
	EnterWork(c *WorkContext)

	// EnterCheck is called when entering the check production.
	EnterCheck(c *CheckContext)

	// ExitIstc is called when exiting the istc production.
	ExitIstc(c *IstcContext)

	// ExitRegistration is called when exiting the registration production.
	ExitRegistration(c *RegistrationContext)

	// ExitYear is called when exiting the year production.
	ExitYear(c *YearContext)

	// ExitWork is called when exiting the work production.
	ExitWork(c *WorkContext)

	// ExitCheck is called when exiting the check production.
	ExitCheck(c *CheckContext)
}
