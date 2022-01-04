// Code generated from bdf.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bdf // bdf
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasebdfListener is a complete listener for a parse tree produced by bdfParser.
type BasebdfListener struct{}

var _ bdfListener = &BasebdfListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasebdfListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasebdfListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasebdfListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasebdfListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFont is called when production font is entered.
func (s *BasebdfListener) EnterFont(ctx *FontContext) {}

// ExitFont is called when production font is exited.
func (s *BasebdfListener) ExitFont(ctx *FontContext) {}

// EnterStartfont is called when production startfont is entered.
func (s *BasebdfListener) EnterStartfont(ctx *StartfontContext) {}

// ExitStartfont is called when production startfont is exited.
func (s *BasebdfListener) ExitStartfont(ctx *StartfontContext) {}

// EnterFontdecl is called when production fontdecl is entered.
func (s *BasebdfListener) EnterFontdecl(ctx *FontdeclContext) {}

// ExitFontdecl is called when production fontdecl is exited.
func (s *BasebdfListener) ExitFontdecl(ctx *FontdeclContext) {}

// EnterSizedecl is called when production sizedecl is entered.
func (s *BasebdfListener) EnterSizedecl(ctx *SizedeclContext) {}

// ExitSizedecl is called when production sizedecl is exited.
func (s *BasebdfListener) ExitSizedecl(ctx *SizedeclContext) {}

// EnterFontboundingboxdecl is called when production fontboundingboxdecl is entered.
func (s *BasebdfListener) EnterFontboundingboxdecl(ctx *FontboundingboxdeclContext) {}

// ExitFontboundingboxdecl is called when production fontboundingboxdecl is exited.
func (s *BasebdfListener) ExitFontboundingboxdecl(ctx *FontboundingboxdeclContext) {}

// EnterPropertiesdecl is called when production propertiesdecl is entered.
func (s *BasebdfListener) EnterPropertiesdecl(ctx *PropertiesdeclContext) {}

// ExitPropertiesdecl is called when production propertiesdecl is exited.
func (s *BasebdfListener) ExitPropertiesdecl(ctx *PropertiesdeclContext) {}

// EnterFoundrydecl is called when production foundrydecl is entered.
func (s *BasebdfListener) EnterFoundrydecl(ctx *FoundrydeclContext) {}

// ExitFoundrydecl is called when production foundrydecl is exited.
func (s *BasebdfListener) ExitFoundrydecl(ctx *FoundrydeclContext) {}

// EnterFamilynanmedecl is called when production familynanmedecl is entered.
func (s *BasebdfListener) EnterFamilynanmedecl(ctx *FamilynanmedeclContext) {}

// ExitFamilynanmedecl is called when production familynanmedecl is exited.
func (s *BasebdfListener) ExitFamilynanmedecl(ctx *FamilynanmedeclContext) {}

// EnterWeightnamedecl is called when production weightnamedecl is entered.
func (s *BasebdfListener) EnterWeightnamedecl(ctx *WeightnamedeclContext) {}

// ExitWeightnamedecl is called when production weightnamedecl is exited.
func (s *BasebdfListener) ExitWeightnamedecl(ctx *WeightnamedeclContext) {}

// EnterSlantdecl is called when production slantdecl is entered.
func (s *BasebdfListener) EnterSlantdecl(ctx *SlantdeclContext) {}

// ExitSlantdecl is called when production slantdecl is exited.
func (s *BasebdfListener) ExitSlantdecl(ctx *SlantdeclContext) {}

// EnterSetwidthnamedecl is called when production setwidthnamedecl is entered.
func (s *BasebdfListener) EnterSetwidthnamedecl(ctx *SetwidthnamedeclContext) {}

// ExitSetwidthnamedecl is called when production setwidthnamedecl is exited.
func (s *BasebdfListener) ExitSetwidthnamedecl(ctx *SetwidthnamedeclContext) {}

// EnterAddstylenamedecl is called when production addstylenamedecl is entered.
func (s *BasebdfListener) EnterAddstylenamedecl(ctx *AddstylenamedeclContext) {}

// ExitAddstylenamedecl is called when production addstylenamedecl is exited.
func (s *BasebdfListener) ExitAddstylenamedecl(ctx *AddstylenamedeclContext) {}

// EnterPixelsizedecl is called when production pixelsizedecl is entered.
func (s *BasebdfListener) EnterPixelsizedecl(ctx *PixelsizedeclContext) {}

// ExitPixelsizedecl is called when production pixelsizedecl is exited.
func (s *BasebdfListener) ExitPixelsizedecl(ctx *PixelsizedeclContext) {}

// EnterPointsizedecl is called when production pointsizedecl is entered.
func (s *BasebdfListener) EnterPointsizedecl(ctx *PointsizedeclContext) {}

// ExitPointsizedecl is called when production pointsizedecl is exited.
func (s *BasebdfListener) ExitPointsizedecl(ctx *PointsizedeclContext) {}

// EnterResolutionxdecl is called when production resolutionxdecl is entered.
func (s *BasebdfListener) EnterResolutionxdecl(ctx *ResolutionxdeclContext) {}

// ExitResolutionxdecl is called when production resolutionxdecl is exited.
func (s *BasebdfListener) ExitResolutionxdecl(ctx *ResolutionxdeclContext) {}

// EnterResolutionydecl is called when production resolutionydecl is entered.
func (s *BasebdfListener) EnterResolutionydecl(ctx *ResolutionydeclContext) {}

// ExitResolutionydecl is called when production resolutionydecl is exited.
func (s *BasebdfListener) ExitResolutionydecl(ctx *ResolutionydeclContext) {}

// EnterSpacingdecl is called when production spacingdecl is entered.
func (s *BasebdfListener) EnterSpacingdecl(ctx *SpacingdeclContext) {}

// ExitSpacingdecl is called when production spacingdecl is exited.
func (s *BasebdfListener) ExitSpacingdecl(ctx *SpacingdeclContext) {}

// EnterAveragewidthdecl is called when production averagewidthdecl is entered.
func (s *BasebdfListener) EnterAveragewidthdecl(ctx *AveragewidthdeclContext) {}

// ExitAveragewidthdecl is called when production averagewidthdecl is exited.
func (s *BasebdfListener) ExitAveragewidthdecl(ctx *AveragewidthdeclContext) {}

// EnterCharsetregistrydecl is called when production charsetregistrydecl is entered.
func (s *BasebdfListener) EnterCharsetregistrydecl(ctx *CharsetregistrydeclContext) {}

// ExitCharsetregistrydecl is called when production charsetregistrydecl is exited.
func (s *BasebdfListener) ExitCharsetregistrydecl(ctx *CharsetregistrydeclContext) {}

// EnterCharsetencoding is called when production charsetencoding is entered.
func (s *BasebdfListener) EnterCharsetencoding(ctx *CharsetencodingContext) {}

// ExitCharsetencoding is called when production charsetencoding is exited.
func (s *BasebdfListener) ExitCharsetencoding(ctx *CharsetencodingContext) {}

// EnterFontnameregistry is called when production fontnameregistry is entered.
func (s *BasebdfListener) EnterFontnameregistry(ctx *FontnameregistryContext) {}

// ExitFontnameregistry is called when production fontnameregistry is exited.
func (s *BasebdfListener) ExitFontnameregistry(ctx *FontnameregistryContext) {}

// EnterFontascentdecl is called when production fontascentdecl is entered.
func (s *BasebdfListener) EnterFontascentdecl(ctx *FontascentdeclContext) {}

// ExitFontascentdecl is called when production fontascentdecl is exited.
func (s *BasebdfListener) ExitFontascentdecl(ctx *FontascentdeclContext) {}

// EnterCharsetcollectionsdecl is called when production charsetcollectionsdecl is entered.
func (s *BasebdfListener) EnterCharsetcollectionsdecl(ctx *CharsetcollectionsdeclContext) {}

// ExitCharsetcollectionsdecl is called when production charsetcollectionsdecl is exited.
func (s *BasebdfListener) ExitCharsetcollectionsdecl(ctx *CharsetcollectionsdeclContext) {}

// EnterFontnamedecl is called when production fontnamedecl is entered.
func (s *BasebdfListener) EnterFontnamedecl(ctx *FontnamedeclContext) {}

// ExitFontnamedecl is called when production fontnamedecl is exited.
func (s *BasebdfListener) ExitFontnamedecl(ctx *FontnamedeclContext) {}

// EnterFacenamedecl is called when production facenamedecl is entered.
func (s *BasebdfListener) EnterFacenamedecl(ctx *FacenamedeclContext) {}

// ExitFacenamedecl is called when production facenamedecl is exited.
func (s *BasebdfListener) ExitFacenamedecl(ctx *FacenamedeclContext) {}

// EnterCopyrightdecl is called when production copyrightdecl is entered.
func (s *BasebdfListener) EnterCopyrightdecl(ctx *CopyrightdeclContext) {}

// ExitCopyrightdecl is called when production copyrightdecl is exited.
func (s *BasebdfListener) ExitCopyrightdecl(ctx *CopyrightdeclContext) {}

// EnterFontdecentdecl is called when production fontdecentdecl is entered.
func (s *BasebdfListener) EnterFontdecentdecl(ctx *FontdecentdeclContext) {}

// ExitFontdecentdecl is called when production fontdecentdecl is exited.
func (s *BasebdfListener) ExitFontdecentdecl(ctx *FontdecentdeclContext) {}

// EnterFontversiondecl is called when production fontversiondecl is entered.
func (s *BasebdfListener) EnterFontversiondecl(ctx *FontversiondeclContext) {}

// ExitFontversiondecl is called when production fontversiondecl is exited.
func (s *BasebdfListener) ExitFontversiondecl(ctx *FontversiondeclContext) {}

// EnterUnderlinepositiondecl is called when production underlinepositiondecl is entered.
func (s *BasebdfListener) EnterUnderlinepositiondecl(ctx *UnderlinepositiondeclContext) {}

// ExitUnderlinepositiondecl is called when production underlinepositiondecl is exited.
func (s *BasebdfListener) ExitUnderlinepositiondecl(ctx *UnderlinepositiondeclContext) {}

// EnterUnderlinethicknessdecl is called when production underlinethicknessdecl is entered.
func (s *BasebdfListener) EnterUnderlinethicknessdecl(ctx *UnderlinethicknessdeclContext) {}

// ExitUnderlinethicknessdecl is called when production underlinethicknessdecl is exited.
func (s *BasebdfListener) ExitUnderlinethicknessdecl(ctx *UnderlinethicknessdeclContext) {}

// EnterXheightdecl is called when production xheightdecl is entered.
func (s *BasebdfListener) EnterXheightdecl(ctx *XheightdeclContext) {}

// ExitXheightdecl is called when production xheightdecl is exited.
func (s *BasebdfListener) ExitXheightdecl(ctx *XheightdeclContext) {}

// EnterCapheighdecl is called when production capheighdecl is entered.
func (s *BasebdfListener) EnterCapheighdecl(ctx *CapheighdeclContext) {}

// ExitCapheighdecl is called when production capheighdecl is exited.
func (s *BasebdfListener) ExitCapheighdecl(ctx *CapheighdeclContext) {}

// EnterRawascentdecl is called when production rawascentdecl is entered.
func (s *BasebdfListener) EnterRawascentdecl(ctx *RawascentdeclContext) {}

// ExitRawascentdecl is called when production rawascentdecl is exited.
func (s *BasebdfListener) ExitRawascentdecl(ctx *RawascentdeclContext) {}

// EnterRawdescentdecl is called when production rawdescentdecl is entered.
func (s *BasebdfListener) EnterRawdescentdecl(ctx *RawdescentdeclContext) {}

// ExitRawdescentdecl is called when production rawdescentdecl is exited.
func (s *BasebdfListener) ExitRawdescentdecl(ctx *RawdescentdeclContext) {}

// EnterNormspacedecl is called when production normspacedecl is entered.
func (s *BasebdfListener) EnterNormspacedecl(ctx *NormspacedeclContext) {}

// ExitNormspacedecl is called when production normspacedecl is exited.
func (s *BasebdfListener) ExitNormspacedecl(ctx *NormspacedeclContext) {}

// EnterRelativeweightdecl is called when production relativeweightdecl is entered.
func (s *BasebdfListener) EnterRelativeweightdecl(ctx *RelativeweightdeclContext) {}

// ExitRelativeweightdecl is called when production relativeweightdecl is exited.
func (s *BasebdfListener) ExitRelativeweightdecl(ctx *RelativeweightdeclContext) {}

// EnterRelaticesetwidthdecl is called when production relaticesetwidthdecl is entered.
func (s *BasebdfListener) EnterRelaticesetwidthdecl(ctx *RelaticesetwidthdeclContext) {}

// ExitRelaticesetwidthdecl is called when production relaticesetwidthdecl is exited.
func (s *BasebdfListener) ExitRelaticesetwidthdecl(ctx *RelaticesetwidthdeclContext) {}

// EnterFigurewidthdecl is called when production figurewidthdecl is entered.
func (s *BasebdfListener) EnterFigurewidthdecl(ctx *FigurewidthdeclContext) {}

// ExitFigurewidthdecl is called when production figurewidthdecl is exited.
func (s *BasebdfListener) ExitFigurewidthdecl(ctx *FigurewidthdeclContext) {}

// EnterAvglowercasewidthdecl is called when production avglowercasewidthdecl is entered.
func (s *BasebdfListener) EnterAvglowercasewidthdecl(ctx *AvglowercasewidthdeclContext) {}

// ExitAvglowercasewidthdecl is called when production avglowercasewidthdecl is exited.
func (s *BasebdfListener) ExitAvglowercasewidthdecl(ctx *AvglowercasewidthdeclContext) {}

// EnterAvguppercasewidthdecl is called when production avguppercasewidthdecl is entered.
func (s *BasebdfListener) EnterAvguppercasewidthdecl(ctx *AvguppercasewidthdeclContext) {}

// ExitAvguppercasewidthdecl is called when production avguppercasewidthdecl is exited.
func (s *BasebdfListener) ExitAvguppercasewidthdecl(ctx *AvguppercasewidthdeclContext) {}

// EnterCharsdecl is called when production charsdecl is entered.
func (s *BasebdfListener) EnterCharsdecl(ctx *CharsdeclContext) {}

// ExitCharsdecl is called when production charsdecl is exited.
func (s *BasebdfListener) ExitCharsdecl(ctx *CharsdeclContext) {}

// EnterChardecl is called when production chardecl is entered.
func (s *BasebdfListener) EnterChardecl(ctx *ChardeclContext) {}

// ExitChardecl is called when production chardecl is exited.
func (s *BasebdfListener) ExitChardecl(ctx *ChardeclContext) {}

// EnterEncodingdecl is called when production encodingdecl is entered.
func (s *BasebdfListener) EnterEncodingdecl(ctx *EncodingdeclContext) {}

// ExitEncodingdecl is called when production encodingdecl is exited.
func (s *BasebdfListener) ExitEncodingdecl(ctx *EncodingdeclContext) {}

// EnterSwidthdecl is called when production swidthdecl is entered.
func (s *BasebdfListener) EnterSwidthdecl(ctx *SwidthdeclContext) {}

// ExitSwidthdecl is called when production swidthdecl is exited.
func (s *BasebdfListener) ExitSwidthdecl(ctx *SwidthdeclContext) {}

// EnterDwidthdecl is called when production dwidthdecl is entered.
func (s *BasebdfListener) EnterDwidthdecl(ctx *DwidthdeclContext) {}

// ExitDwidthdecl is called when production dwidthdecl is exited.
func (s *BasebdfListener) ExitDwidthdecl(ctx *DwidthdeclContext) {}

// EnterBbxdecl is called when production bbxdecl is entered.
func (s *BasebdfListener) EnterBbxdecl(ctx *BbxdeclContext) {}

// ExitBbxdecl is called when production bbxdecl is exited.
func (s *BasebdfListener) ExitBbxdecl(ctx *BbxdeclContext) {}

// EnterBitmapdecl is called when production bitmapdecl is entered.
func (s *BasebdfListener) EnterBitmapdecl(ctx *BitmapdeclContext) {}

// ExitBitmapdecl is called when production bitmapdecl is exited.
func (s *BasebdfListener) ExitBitmapdecl(ctx *BitmapdeclContext) {}
