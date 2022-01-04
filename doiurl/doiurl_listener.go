// Code generated from doiurl.g4 by ANTLR 4.9.3. DO NOT EDIT.

package doiurl // doiurl
import "github.com/antlr/antlr4/runtime/Go/antlr"

// doiurlListener is a complete listener for a parse tree produced by doiurlParser.
type doiurlListener interface {
	antlr.ParseTreeListener

	// EnterDoiuri is called when entering the doiuri production.
	EnterDoiuri(c *DoiuriContext)

	// EnterScheme is called when entering the scheme production.
	EnterScheme(c *SchemeContext)

	// EnterEncodeddoi is called when entering the encodeddoi production.
	EnterEncodeddoi(c *EncodeddoiContext)

	// EnterPrefix_ is called when entering the prefix_ production.
	EnterPrefix_(c *Prefix_Context)

	// EnterSuffix is called when entering the suffix production.
	EnterSuffix(c *SuffixContext)

	// EnterSegment is called when entering the segment production.
	EnterSegment(c *SegmentContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterFragment_ is called when entering the fragment_ production.
	EnterFragment_(c *Fragment_Context)

	// ExitDoiuri is called when exiting the doiuri production.
	ExitDoiuri(c *DoiuriContext)

	// ExitScheme is called when exiting the scheme production.
	ExitScheme(c *SchemeContext)

	// ExitEncodeddoi is called when exiting the encodeddoi production.
	ExitEncodeddoi(c *EncodeddoiContext)

	// ExitPrefix_ is called when exiting the prefix_ production.
	ExitPrefix_(c *Prefix_Context)

	// ExitSuffix is called when exiting the suffix production.
	ExitSuffix(c *SuffixContext)

	// ExitSegment is called when exiting the segment production.
	ExitSegment(c *SegmentContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitFragment_ is called when exiting the fragment_ production.
	ExitFragment_(c *Fragment_Context)
}
