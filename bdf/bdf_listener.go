// Code generated from bdf.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bdf // bdf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// bdfListener is a complete listener for a parse tree produced by bdfParser.
type bdfListener interface {
	antlr.ParseTreeListener

	// EnterFont is called when entering the font production.
	EnterFont(c *FontContext)

	// EnterStartfont is called when entering the startfont production.
	EnterStartfont(c *StartfontContext)

	// EnterFontdecl is called when entering the fontdecl production.
	EnterFontdecl(c *FontdeclContext)

	// EnterSizedecl is called when entering the sizedecl production.
	EnterSizedecl(c *SizedeclContext)

	// EnterFontboundingboxdecl is called when entering the fontboundingboxdecl production.
	EnterFontboundingboxdecl(c *FontboundingboxdeclContext)

	// EnterPropertiesdecl is called when entering the propertiesdecl production.
	EnterPropertiesdecl(c *PropertiesdeclContext)

	// EnterFoundrydecl is called when entering the foundrydecl production.
	EnterFoundrydecl(c *FoundrydeclContext)

	// EnterFamilynanmedecl is called when entering the familynanmedecl production.
	EnterFamilynanmedecl(c *FamilynanmedeclContext)

	// EnterWeightnamedecl is called when entering the weightnamedecl production.
	EnterWeightnamedecl(c *WeightnamedeclContext)

	// EnterSlantdecl is called when entering the slantdecl production.
	EnterSlantdecl(c *SlantdeclContext)

	// EnterSetwidthnamedecl is called when entering the setwidthnamedecl production.
	EnterSetwidthnamedecl(c *SetwidthnamedeclContext)

	// EnterAddstylenamedecl is called when entering the addstylenamedecl production.
	EnterAddstylenamedecl(c *AddstylenamedeclContext)

	// EnterPixelsizedecl is called when entering the pixelsizedecl production.
	EnterPixelsizedecl(c *PixelsizedeclContext)

	// EnterPointsizedecl is called when entering the pointsizedecl production.
	EnterPointsizedecl(c *PointsizedeclContext)

	// EnterResolutionxdecl is called when entering the resolutionxdecl production.
	EnterResolutionxdecl(c *ResolutionxdeclContext)

	// EnterResolutionydecl is called when entering the resolutionydecl production.
	EnterResolutionydecl(c *ResolutionydeclContext)

	// EnterSpacingdecl is called when entering the spacingdecl production.
	EnterSpacingdecl(c *SpacingdeclContext)

	// EnterAveragewidthdecl is called when entering the averagewidthdecl production.
	EnterAveragewidthdecl(c *AveragewidthdeclContext)

	// EnterCharsetregistrydecl is called when entering the charsetregistrydecl production.
	EnterCharsetregistrydecl(c *CharsetregistrydeclContext)

	// EnterCharsetencoding is called when entering the charsetencoding production.
	EnterCharsetencoding(c *CharsetencodingContext)

	// EnterFontnameregistry is called when entering the fontnameregistry production.
	EnterFontnameregistry(c *FontnameregistryContext)

	// EnterFontascentdecl is called when entering the fontascentdecl production.
	EnterFontascentdecl(c *FontascentdeclContext)

	// EnterCharsetcollectionsdecl is called when entering the charsetcollectionsdecl production.
	EnterCharsetcollectionsdecl(c *CharsetcollectionsdeclContext)

	// EnterFontnamedecl is called when entering the fontnamedecl production.
	EnterFontnamedecl(c *FontnamedeclContext)

	// EnterFacenamedecl is called when entering the facenamedecl production.
	EnterFacenamedecl(c *FacenamedeclContext)

	// EnterCopyrightdecl is called when entering the copyrightdecl production.
	EnterCopyrightdecl(c *CopyrightdeclContext)

	// EnterFontdecentdecl is called when entering the fontdecentdecl production.
	EnterFontdecentdecl(c *FontdecentdeclContext)

	// EnterFontversiondecl is called when entering the fontversiondecl production.
	EnterFontversiondecl(c *FontversiondeclContext)

	// EnterUnderlinepositiondecl is called when entering the underlinepositiondecl production.
	EnterUnderlinepositiondecl(c *UnderlinepositiondeclContext)

	// EnterUnderlinethicknessdecl is called when entering the underlinethicknessdecl production.
	EnterUnderlinethicknessdecl(c *UnderlinethicknessdeclContext)

	// EnterXheightdecl is called when entering the xheightdecl production.
	EnterXheightdecl(c *XheightdeclContext)

	// EnterCapheighdecl is called when entering the capheighdecl production.
	EnterCapheighdecl(c *CapheighdeclContext)

	// EnterRawascentdecl is called when entering the rawascentdecl production.
	EnterRawascentdecl(c *RawascentdeclContext)

	// EnterRawdescentdecl is called when entering the rawdescentdecl production.
	EnterRawdescentdecl(c *RawdescentdeclContext)

	// EnterNormspacedecl is called when entering the normspacedecl production.
	EnterNormspacedecl(c *NormspacedeclContext)

	// EnterRelativeweightdecl is called when entering the relativeweightdecl production.
	EnterRelativeweightdecl(c *RelativeweightdeclContext)

	// EnterRelaticesetwidthdecl is called when entering the relaticesetwidthdecl production.
	EnterRelaticesetwidthdecl(c *RelaticesetwidthdeclContext)

	// EnterFigurewidthdecl is called when entering the figurewidthdecl production.
	EnterFigurewidthdecl(c *FigurewidthdeclContext)

	// EnterAvglowercasewidthdecl is called when entering the avglowercasewidthdecl production.
	EnterAvglowercasewidthdecl(c *AvglowercasewidthdeclContext)

	// EnterAvguppercasewidthdecl is called when entering the avguppercasewidthdecl production.
	EnterAvguppercasewidthdecl(c *AvguppercasewidthdeclContext)

	// EnterCharsdecl is called when entering the charsdecl production.
	EnterCharsdecl(c *CharsdeclContext)

	// EnterChardecl is called when entering the chardecl production.
	EnterChardecl(c *ChardeclContext)

	// EnterEncodingdecl is called when entering the encodingdecl production.
	EnterEncodingdecl(c *EncodingdeclContext)

	// EnterSwidthdecl is called when entering the swidthdecl production.
	EnterSwidthdecl(c *SwidthdeclContext)

	// EnterDwidthdecl is called when entering the dwidthdecl production.
	EnterDwidthdecl(c *DwidthdeclContext)

	// EnterBbxdecl is called when entering the bbxdecl production.
	EnterBbxdecl(c *BbxdeclContext)

	// EnterBitmapdecl is called when entering the bitmapdecl production.
	EnterBitmapdecl(c *BitmapdeclContext)

	// ExitFont is called when exiting the font production.
	ExitFont(c *FontContext)

	// ExitStartfont is called when exiting the startfont production.
	ExitStartfont(c *StartfontContext)

	// ExitFontdecl is called when exiting the fontdecl production.
	ExitFontdecl(c *FontdeclContext)

	// ExitSizedecl is called when exiting the sizedecl production.
	ExitSizedecl(c *SizedeclContext)

	// ExitFontboundingboxdecl is called when exiting the fontboundingboxdecl production.
	ExitFontboundingboxdecl(c *FontboundingboxdeclContext)

	// ExitPropertiesdecl is called when exiting the propertiesdecl production.
	ExitPropertiesdecl(c *PropertiesdeclContext)

	// ExitFoundrydecl is called when exiting the foundrydecl production.
	ExitFoundrydecl(c *FoundrydeclContext)

	// ExitFamilynanmedecl is called when exiting the familynanmedecl production.
	ExitFamilynanmedecl(c *FamilynanmedeclContext)

	// ExitWeightnamedecl is called when exiting the weightnamedecl production.
	ExitWeightnamedecl(c *WeightnamedeclContext)

	// ExitSlantdecl is called when exiting the slantdecl production.
	ExitSlantdecl(c *SlantdeclContext)

	// ExitSetwidthnamedecl is called when exiting the setwidthnamedecl production.
	ExitSetwidthnamedecl(c *SetwidthnamedeclContext)

	// ExitAddstylenamedecl is called when exiting the addstylenamedecl production.
	ExitAddstylenamedecl(c *AddstylenamedeclContext)

	// ExitPixelsizedecl is called when exiting the pixelsizedecl production.
	ExitPixelsizedecl(c *PixelsizedeclContext)

	// ExitPointsizedecl is called when exiting the pointsizedecl production.
	ExitPointsizedecl(c *PointsizedeclContext)

	// ExitResolutionxdecl is called when exiting the resolutionxdecl production.
	ExitResolutionxdecl(c *ResolutionxdeclContext)

	// ExitResolutionydecl is called when exiting the resolutionydecl production.
	ExitResolutionydecl(c *ResolutionydeclContext)

	// ExitSpacingdecl is called when exiting the spacingdecl production.
	ExitSpacingdecl(c *SpacingdeclContext)

	// ExitAveragewidthdecl is called when exiting the averagewidthdecl production.
	ExitAveragewidthdecl(c *AveragewidthdeclContext)

	// ExitCharsetregistrydecl is called when exiting the charsetregistrydecl production.
	ExitCharsetregistrydecl(c *CharsetregistrydeclContext)

	// ExitCharsetencoding is called when exiting the charsetencoding production.
	ExitCharsetencoding(c *CharsetencodingContext)

	// ExitFontnameregistry is called when exiting the fontnameregistry production.
	ExitFontnameregistry(c *FontnameregistryContext)

	// ExitFontascentdecl is called when exiting the fontascentdecl production.
	ExitFontascentdecl(c *FontascentdeclContext)

	// ExitCharsetcollectionsdecl is called when exiting the charsetcollectionsdecl production.
	ExitCharsetcollectionsdecl(c *CharsetcollectionsdeclContext)

	// ExitFontnamedecl is called when exiting the fontnamedecl production.
	ExitFontnamedecl(c *FontnamedeclContext)

	// ExitFacenamedecl is called when exiting the facenamedecl production.
	ExitFacenamedecl(c *FacenamedeclContext)

	// ExitCopyrightdecl is called when exiting the copyrightdecl production.
	ExitCopyrightdecl(c *CopyrightdeclContext)

	// ExitFontdecentdecl is called when exiting the fontdecentdecl production.
	ExitFontdecentdecl(c *FontdecentdeclContext)

	// ExitFontversiondecl is called when exiting the fontversiondecl production.
	ExitFontversiondecl(c *FontversiondeclContext)

	// ExitUnderlinepositiondecl is called when exiting the underlinepositiondecl production.
	ExitUnderlinepositiondecl(c *UnderlinepositiondeclContext)

	// ExitUnderlinethicknessdecl is called when exiting the underlinethicknessdecl production.
	ExitUnderlinethicknessdecl(c *UnderlinethicknessdeclContext)

	// ExitXheightdecl is called when exiting the xheightdecl production.
	ExitXheightdecl(c *XheightdeclContext)

	// ExitCapheighdecl is called when exiting the capheighdecl production.
	ExitCapheighdecl(c *CapheighdeclContext)

	// ExitRawascentdecl is called when exiting the rawascentdecl production.
	ExitRawascentdecl(c *RawascentdeclContext)

	// ExitRawdescentdecl is called when exiting the rawdescentdecl production.
	ExitRawdescentdecl(c *RawdescentdeclContext)

	// ExitNormspacedecl is called when exiting the normspacedecl production.
	ExitNormspacedecl(c *NormspacedeclContext)

	// ExitRelativeweightdecl is called when exiting the relativeweightdecl production.
	ExitRelativeweightdecl(c *RelativeweightdeclContext)

	// ExitRelaticesetwidthdecl is called when exiting the relaticesetwidthdecl production.
	ExitRelaticesetwidthdecl(c *RelaticesetwidthdeclContext)

	// ExitFigurewidthdecl is called when exiting the figurewidthdecl production.
	ExitFigurewidthdecl(c *FigurewidthdeclContext)

	// ExitAvglowercasewidthdecl is called when exiting the avglowercasewidthdecl production.
	ExitAvglowercasewidthdecl(c *AvglowercasewidthdeclContext)

	// ExitAvguppercasewidthdecl is called when exiting the avguppercasewidthdecl production.
	ExitAvguppercasewidthdecl(c *AvguppercasewidthdeclContext)

	// ExitCharsdecl is called when exiting the charsdecl production.
	ExitCharsdecl(c *CharsdeclContext)

	// ExitChardecl is called when exiting the chardecl production.
	ExitChardecl(c *ChardeclContext)

	// ExitEncodingdecl is called when exiting the encodingdecl production.
	ExitEncodingdecl(c *EncodingdeclContext)

	// ExitSwidthdecl is called when exiting the swidthdecl production.
	ExitSwidthdecl(c *SwidthdeclContext)

	// ExitDwidthdecl is called when exiting the dwidthdecl production.
	ExitDwidthdecl(c *DwidthdeclContext)

	// ExitBbxdecl is called when exiting the bbxdecl production.
	ExitBbxdecl(c *BbxdeclContext)

	// ExitBitmapdecl is called when exiting the bitmapdecl production.
	ExitBitmapdecl(c *BitmapdeclContext)
}
