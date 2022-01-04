// Code generated from itn.g4 by ANTLR 4.9.3. DO NOT EDIT.

package itn // itn
import "github.com/antlr/antlr4/runtime/Go/antlr"

// itnListener is a complete listener for a parse tree produced by itnParser.
type itnListener interface {
	antlr.ParseTreeListener

	// EnterItinerary is called when entering the itinerary production.
	EnterItinerary(c *ItineraryContext)

	// EnterLine is called when entering the line production.
	EnterLine(c *LineContext)

	// EnterLongitude is called when entering the longitude production.
	EnterLongitude(c *LongitudeContext)

	// EnterLatitude is called when entering the latitude production.
	EnterLatitude(c *LatitudeContext)

	// EnterDescr is called when entering the descr production.
	EnterDescr(c *DescrContext)

	// EnterFlag is called when entering the flag production.
	EnterFlag(c *FlagContext)

	// ExitItinerary is called when exiting the itinerary production.
	ExitItinerary(c *ItineraryContext)

	// ExitLine is called when exiting the line production.
	ExitLine(c *LineContext)

	// ExitLongitude is called when exiting the longitude production.
	ExitLongitude(c *LongitudeContext)

	// ExitLatitude is called when exiting the latitude production.
	ExitLatitude(c *LatitudeContext)

	// ExitDescr is called when exiting the descr production.
	ExitDescr(c *DescrContext)

	// ExitFlag is called when exiting the flag production.
	ExitFlag(c *FlagContext)
}
