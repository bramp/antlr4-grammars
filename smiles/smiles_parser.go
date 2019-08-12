// Code generated from smiles.g4 by ANTLR 4.7.2. DO NOT EDIT.

package smiles // smiles
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 213, 170,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 48, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 56, 10, 6,
	3, 6, 3, 6, 5, 6, 60, 10, 6, 3, 6, 5, 6, 63, 10, 6, 3, 6, 5, 6, 66, 10,
	6, 3, 6, 5, 6, 69, 10, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 5, 7, 76, 10, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12,
	5, 12, 89, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 5, 13, 99, 10, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 5, 16,
	107, 10, 16, 3, 16, 3, 16, 5, 16, 111, 10, 16, 3, 16, 3, 16, 3, 16, 5,
	16, 116, 10, 16, 3, 17, 3, 17, 7, 17, 120, 10, 17, 12, 17, 14, 17, 123,
	11, 17, 3, 17, 7, 17, 126, 10, 17, 12, 17, 14, 17, 129, 11, 17, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 5, 18, 145, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 159, 10, 19, 12,
	19, 14, 19, 162, 11, 19, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 168, 10, 20,
	3, 20, 2, 3, 36, 21, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 2, 8, 3, 2, 4, 13, 3, 2, 14, 19, 4, 2, 4, 13, 22, 122,
	4, 2, 15, 19, 123, 124, 3, 2, 125, 193, 4, 2, 194, 194, 198, 203, 2, 178,
	2, 40, 3, 2, 2, 2, 4, 47, 3, 2, 2, 2, 6, 49, 3, 2, 2, 2, 8, 51, 3, 2, 2,
	2, 10, 53, 3, 2, 2, 2, 12, 75, 3, 2, 2, 2, 14, 77, 3, 2, 2, 2, 16, 79,
	3, 2, 2, 2, 18, 81, 3, 2, 2, 2, 20, 83, 3, 2, 2, 2, 22, 88, 3, 2, 2, 2,
	24, 98, 3, 2, 2, 2, 26, 100, 3, 2, 2, 2, 28, 103, 3, 2, 2, 2, 30, 115,
	3, 2, 2, 2, 32, 117, 3, 2, 2, 2, 34, 144, 3, 2, 2, 2, 36, 146, 3, 2, 2,
	2, 38, 167, 3, 2, 2, 2, 40, 41, 5, 36, 19, 2, 41, 42, 5, 38, 20, 2, 42,
	3, 3, 2, 2, 2, 43, 48, 5, 10, 6, 2, 44, 48, 5, 6, 4, 2, 45, 48, 5, 8, 5,
	2, 46, 48, 7, 3, 2, 2, 47, 43, 3, 2, 2, 2, 47, 44, 3, 2, 2, 2, 47, 45,
	3, 2, 2, 2, 47, 46, 3, 2, 2, 2, 48, 5, 3, 2, 2, 2, 49, 50, 9, 2, 2, 2,
	50, 7, 3, 2, 2, 2, 51, 52, 9, 3, 2, 2, 52, 9, 3, 2, 2, 2, 53, 55, 7, 20,
	2, 2, 54, 56, 5, 14, 8, 2, 55, 54, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56,
	57, 3, 2, 2, 2, 57, 59, 5, 12, 7, 2, 58, 60, 5, 20, 11, 2, 59, 58, 3, 2,
	2, 2, 59, 60, 3, 2, 2, 2, 60, 62, 3, 2, 2, 2, 61, 63, 5, 22, 12, 2, 62,
	61, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2, 64, 66, 5, 24,
	13, 2, 65, 64, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68, 3, 2, 2, 2, 67,
	69, 5, 26, 14, 2, 68, 67, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 70, 3, 2,
	2, 2, 70, 71, 7, 21, 2, 2, 71, 11, 3, 2, 2, 2, 72, 76, 5, 16, 9, 2, 73,
	76, 5, 18, 10, 2, 74, 76, 7, 3, 2, 2, 75, 72, 3, 2, 2, 2, 75, 73, 3, 2,
	2, 2, 75, 74, 3, 2, 2, 2, 76, 13, 3, 2, 2, 2, 77, 78, 7, 212, 2, 2, 78,
	15, 3, 2, 2, 2, 79, 80, 9, 4, 2, 2, 80, 17, 3, 2, 2, 2, 81, 82, 9, 5, 2,
	2, 82, 19, 3, 2, 2, 2, 83, 84, 9, 6, 2, 2, 84, 21, 3, 2, 2, 2, 85, 89,
	7, 22, 2, 2, 86, 87, 7, 22, 2, 2, 87, 89, 7, 211, 2, 2, 88, 85, 3, 2, 2,
	2, 88, 86, 3, 2, 2, 2, 89, 23, 3, 2, 2, 2, 90, 99, 7, 194, 2, 2, 91, 92,
	7, 194, 2, 2, 92, 99, 7, 211, 2, 2, 93, 99, 7, 195, 2, 2, 94, 95, 7, 195,
	2, 2, 95, 99, 7, 211, 2, 2, 96, 99, 7, 196, 2, 2, 97, 99, 7, 197, 2, 2,
	98, 90, 3, 2, 2, 2, 98, 91, 3, 2, 2, 2, 98, 93, 3, 2, 2, 2, 98, 94, 3,
	2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 97, 3, 2, 2, 2, 99, 25, 3, 2, 2, 2, 100,
	101, 7, 198, 2, 2, 101, 102, 7, 212, 2, 2, 102, 27, 3, 2, 2, 2, 103, 104,
	9, 7, 2, 2, 104, 29, 3, 2, 2, 2, 105, 107, 5, 28, 15, 2, 106, 105, 3, 2,
	2, 2, 106, 107, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 116, 7, 211, 2,
	2, 109, 111, 5, 28, 15, 2, 110, 109, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2,
	111, 112, 3, 2, 2, 2, 112, 113, 7, 204, 2, 2, 113, 114, 7, 211, 2, 2, 114,
	116, 7, 211, 2, 2, 115, 106, 3, 2, 2, 2, 115, 110, 3, 2, 2, 2, 116, 31,
	3, 2, 2, 2, 117, 121, 5, 4, 3, 2, 118, 120, 5, 30, 16, 2, 119, 118, 3,
	2, 2, 2, 120, 123, 3, 2, 2, 2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2,
	2, 122, 127, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 124, 126, 5, 34, 18, 2,
	125, 124, 3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127,
	128, 3, 2, 2, 2, 128, 33, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130, 131, 7,
	205, 2, 2, 131, 132, 5, 36, 19, 2, 132, 133, 7, 206, 2, 2, 133, 145, 3,
	2, 2, 2, 134, 135, 7, 205, 2, 2, 135, 136, 5, 28, 15, 2, 136, 137, 5, 36,
	19, 2, 137, 138, 7, 206, 2, 2, 138, 145, 3, 2, 2, 2, 139, 140, 7, 205,
	2, 2, 140, 141, 7, 207, 2, 2, 141, 142, 5, 36, 19, 2, 142, 143, 7, 206,
	2, 2, 143, 145, 3, 2, 2, 2, 144, 130, 3, 2, 2, 2, 144, 134, 3, 2, 2, 2,
	144, 139, 3, 2, 2, 2, 145, 35, 3, 2, 2, 2, 146, 147, 8, 19, 1, 2, 147,
	148, 5, 32, 17, 2, 148, 160, 3, 2, 2, 2, 149, 150, 12, 5, 2, 2, 150, 159,
	5, 32, 17, 2, 151, 152, 12, 4, 2, 2, 152, 153, 5, 28, 15, 2, 153, 154,
	5, 32, 17, 2, 154, 159, 3, 2, 2, 2, 155, 156, 12, 3, 2, 2, 156, 157, 7,
	207, 2, 2, 157, 159, 5, 32, 17, 2, 158, 149, 3, 2, 2, 2, 158, 151, 3, 2,
	2, 2, 158, 155, 3, 2, 2, 2, 159, 162, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2,
	160, 161, 3, 2, 2, 2, 161, 37, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 163, 164,
	7, 210, 2, 2, 164, 168, 7, 213, 2, 2, 165, 168, 7, 208, 2, 2, 166, 168,
	7, 209, 2, 2, 167, 163, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 166, 3,
	2, 2, 2, 168, 39, 3, 2, 2, 2, 20, 47, 55, 59, 62, 65, 68, 75, 88, 98, 106,
	110, 115, 121, 127, 144, 158, 160, 167,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'*'", "'B'", "'C'", "'N'", "'O'", "'S'", "'P'", "'F'", "'Cl'", "'Br'",
	"'I'", "'b'", "'c'", "'n'", "'o'", "'s'", "'p'", "'['", "']'", "'H'", "'He'",
	"'Li'", "'Be'", "'Ne'", "'Na'", "'Mg'", "'Al'", "'Si'", "'Ar'", "'K'",
	"'Ca'", "'Sc'", "'Ti'", "'V'", "'Cr'", "'Mn'", "'Fe'", "'Co'", "'Ni'",
	"'Cu'", "'Zn'", "'Ga'", "'Ge'", "'As'", "'Se'", "'Kr'", "'Rb'", "'Sr'",
	"'Y'", "'Zr'", "'Nb'", "'Mo'", "'Tc'", "'Ru'", "'Rh'", "'Pd'", "'Ag'",
	"'Cd'", "'In'", "'Sn'", "'Sb'", "'Te'", "'Xe'", "'Cs'", "'Ba'", "'Hf'",
	"'Ta'", "'W'", "'Re'", "'Os'", "'Ir'", "'Pt'", "'Au'", "'Hg'", "'Tl'",
	"'Pb'", "'Bi'", "'Po'", "'At'", "'Rn'", "'Fr'", "'Ra'", "'Rf'", "'Db'",
	"'Sg'", "'Bh'", "'Hs'", "'Mt'", "'Ds'", "'Rg'", "'La'", "'Ce'", "'Pr'",
	"'Nd'", "'Pm'", "'Sm'", "'Eu'", "'Gd'", "'Tb'", "'Dy'", "'Ho'", "'Er'",
	"'Tm'", "'Yb'", "'Lu'", "'Ac'", "'Th'", "'Pa'", "'U'", "'Np'", "'Pu'",
	"'Am'", "'Cm'", "'Bk'", "'Cf'", "'Es'", "'Fm'", "'Md'", "'No'", "'Lr'",
	"'se'", "'as'", "'@'", "'@@'", "'@TH1'", "'@TH2'", "'@AL1'", "'@AL2'",
	"'@SP1'", "'@SP2'", "'@SP3'", "'@TB1'", "'@TB2'", "'@TB3'", "'@TB4'", "'@TB5'",
	"'@TB6'", "'@TB7'", "'@TB8'", "'@TB9'", "'@TB10'", "'@TB11'", "'@TB12'",
	"'@TB13'", "'@TB14'", "'@TB15'", "'@TB16'", "'@TB17'", "'@TB18'", "'@TB19'",
	"'@TB20'", "'@TB21'", "'@TB22'", "'@TB23'", "'@TB24'", "'@TB25'", "'@TB26'",
	"'@TB27'", "'@TB28'", "'@TB29'", "'@TB30'", "'@OH1'", "'@OH2'", "'@OH3'",
	"'@OH4'", "'@OH5'", "'@OH6'", "'@OH7'", "'@OH8'", "'@OH9'", "'@OH10'",
	"'@OH11'", "'@OH12'", "'@OH13'", "'@OH14'", "'@OH15'", "'@OH16'", "'@OH17'",
	"'@OH18'", "'@OH19'", "'@OH20'", "'@OH21'", "'@OH22'", "'@OH23'", "'@OH24'",
	"'@OH25'", "'@OH26'", "'@OH27'", "'@OH28'", "'@OH29'", "'@OH30'", "'-'",
	"'+'", "'--'", "'++'", "':'", "'='", "'#'", "'$'", "'/'", "'\\'", "'%'",
	"'('", "')'", "'.'", "'\r'", "'\n'", "' '", "", "", "'\t'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "DOT", "LINEFEED", "CARRIAGE_RETURN", "SPACE",
	"DIGIT", "NUMBER", "TAB",
}

var ruleNames = []string{
	"smiles", "atom", "aliphatic_organic", "aromatic_organic", "bracket_atom",
	"symbol", "isotope", "element_symbols", "aromatic_symbols", "chiral", "hcount",
	"charge", "class_", "bond", "ringbond", "branched_atom", "branch", "chain",
	"terminator",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type smilesParser struct {
	*antlr.BaseParser
}

func NewsmilesParser(input antlr.TokenStream) *smilesParser {
	this := new(smilesParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "smiles.g4"

	return this
}

// smilesParser tokens.
const (
	smilesParserEOF             = antlr.TokenEOF
	smilesParserT__0            = 1
	smilesParserT__1            = 2
	smilesParserT__2            = 3
	smilesParserT__3            = 4
	smilesParserT__4            = 5
	smilesParserT__5            = 6
	smilesParserT__6            = 7
	smilesParserT__7            = 8
	smilesParserT__8            = 9
	smilesParserT__9            = 10
	smilesParserT__10           = 11
	smilesParserT__11           = 12
	smilesParserT__12           = 13
	smilesParserT__13           = 14
	smilesParserT__14           = 15
	smilesParserT__15           = 16
	smilesParserT__16           = 17
	smilesParserT__17           = 18
	smilesParserT__18           = 19
	smilesParserT__19           = 20
	smilesParserT__20           = 21
	smilesParserT__21           = 22
	smilesParserT__22           = 23
	smilesParserT__23           = 24
	smilesParserT__24           = 25
	smilesParserT__25           = 26
	smilesParserT__26           = 27
	smilesParserT__27           = 28
	smilesParserT__28           = 29
	smilesParserT__29           = 30
	smilesParserT__30           = 31
	smilesParserT__31           = 32
	smilesParserT__32           = 33
	smilesParserT__33           = 34
	smilesParserT__34           = 35
	smilesParserT__35           = 36
	smilesParserT__36           = 37
	smilesParserT__37           = 38
	smilesParserT__38           = 39
	smilesParserT__39           = 40
	smilesParserT__40           = 41
	smilesParserT__41           = 42
	smilesParserT__42           = 43
	smilesParserT__43           = 44
	smilesParserT__44           = 45
	smilesParserT__45           = 46
	smilesParserT__46           = 47
	smilesParserT__47           = 48
	smilesParserT__48           = 49
	smilesParserT__49           = 50
	smilesParserT__50           = 51
	smilesParserT__51           = 52
	smilesParserT__52           = 53
	smilesParserT__53           = 54
	smilesParserT__54           = 55
	smilesParserT__55           = 56
	smilesParserT__56           = 57
	smilesParserT__57           = 58
	smilesParserT__58           = 59
	smilesParserT__59           = 60
	smilesParserT__60           = 61
	smilesParserT__61           = 62
	smilesParserT__62           = 63
	smilesParserT__63           = 64
	smilesParserT__64           = 65
	smilesParserT__65           = 66
	smilesParserT__66           = 67
	smilesParserT__67           = 68
	smilesParserT__68           = 69
	smilesParserT__69           = 70
	smilesParserT__70           = 71
	smilesParserT__71           = 72
	smilesParserT__72           = 73
	smilesParserT__73           = 74
	smilesParserT__74           = 75
	smilesParserT__75           = 76
	smilesParserT__76           = 77
	smilesParserT__77           = 78
	smilesParserT__78           = 79
	smilesParserT__79           = 80
	smilesParserT__80           = 81
	smilesParserT__81           = 82
	smilesParserT__82           = 83
	smilesParserT__83           = 84
	smilesParserT__84           = 85
	smilesParserT__85           = 86
	smilesParserT__86           = 87
	smilesParserT__87           = 88
	smilesParserT__88           = 89
	smilesParserT__89           = 90
	smilesParserT__90           = 91
	smilesParserT__91           = 92
	smilesParserT__92           = 93
	smilesParserT__93           = 94
	smilesParserT__94           = 95
	smilesParserT__95           = 96
	smilesParserT__96           = 97
	smilesParserT__97           = 98
	smilesParserT__98           = 99
	smilesParserT__99           = 100
	smilesParserT__100          = 101
	smilesParserT__101          = 102
	smilesParserT__102          = 103
	smilesParserT__103          = 104
	smilesParserT__104          = 105
	smilesParserT__105          = 106
	smilesParserT__106          = 107
	smilesParserT__107          = 108
	smilesParserT__108          = 109
	smilesParserT__109          = 110
	smilesParserT__110          = 111
	smilesParserT__111          = 112
	smilesParserT__112          = 113
	smilesParserT__113          = 114
	smilesParserT__114          = 115
	smilesParserT__115          = 116
	smilesParserT__116          = 117
	smilesParserT__117          = 118
	smilesParserT__118          = 119
	smilesParserT__119          = 120
	smilesParserT__120          = 121
	smilesParserT__121          = 122
	smilesParserT__122          = 123
	smilesParserT__123          = 124
	smilesParserT__124          = 125
	smilesParserT__125          = 126
	smilesParserT__126          = 127
	smilesParserT__127          = 128
	smilesParserT__128          = 129
	smilesParserT__129          = 130
	smilesParserT__130          = 131
	smilesParserT__131          = 132
	smilesParserT__132          = 133
	smilesParserT__133          = 134
	smilesParserT__134          = 135
	smilesParserT__135          = 136
	smilesParserT__136          = 137
	smilesParserT__137          = 138
	smilesParserT__138          = 139
	smilesParserT__139          = 140
	smilesParserT__140          = 141
	smilesParserT__141          = 142
	smilesParserT__142          = 143
	smilesParserT__143          = 144
	smilesParserT__144          = 145
	smilesParserT__145          = 146
	smilesParserT__146          = 147
	smilesParserT__147          = 148
	smilesParserT__148          = 149
	smilesParserT__149          = 150
	smilesParserT__150          = 151
	smilesParserT__151          = 152
	smilesParserT__152          = 153
	smilesParserT__153          = 154
	smilesParserT__154          = 155
	smilesParserT__155          = 156
	smilesParserT__156          = 157
	smilesParserT__157          = 158
	smilesParserT__158          = 159
	smilesParserT__159          = 160
	smilesParserT__160          = 161
	smilesParserT__161          = 162
	smilesParserT__162          = 163
	smilesParserT__163          = 164
	smilesParserT__164          = 165
	smilesParserT__165          = 166
	smilesParserT__166          = 167
	smilesParserT__167          = 168
	smilesParserT__168          = 169
	smilesParserT__169          = 170
	smilesParserT__170          = 171
	smilesParserT__171          = 172
	smilesParserT__172          = 173
	smilesParserT__173          = 174
	smilesParserT__174          = 175
	smilesParserT__175          = 176
	smilesParserT__176          = 177
	smilesParserT__177          = 178
	smilesParserT__178          = 179
	smilesParserT__179          = 180
	smilesParserT__180          = 181
	smilesParserT__181          = 182
	smilesParserT__182          = 183
	smilesParserT__183          = 184
	smilesParserT__184          = 185
	smilesParserT__185          = 186
	smilesParserT__186          = 187
	smilesParserT__187          = 188
	smilesParserT__188          = 189
	smilesParserT__189          = 190
	smilesParserT__190          = 191
	smilesParserT__191          = 192
	smilesParserT__192          = 193
	smilesParserT__193          = 194
	smilesParserT__194          = 195
	smilesParserT__195          = 196
	smilesParserT__196          = 197
	smilesParserT__197          = 198
	smilesParserT__198          = 199
	smilesParserT__199          = 200
	smilesParserT__200          = 201
	smilesParserT__201          = 202
	smilesParserT__202          = 203
	smilesParserT__203          = 204
	smilesParserDOT             = 205
	smilesParserLINEFEED        = 206
	smilesParserCARRIAGE_RETURN = 207
	smilesParserSPACE           = 208
	smilesParserDIGIT           = 209
	smilesParserNUMBER          = 210
	smilesParserTAB             = 211
)

// smilesParser rules.
const (
	smilesParserRULE_smiles            = 0
	smilesParserRULE_atom              = 1
	smilesParserRULE_aliphatic_organic = 2
	smilesParserRULE_aromatic_organic  = 3
	smilesParserRULE_bracket_atom      = 4
	smilesParserRULE_symbol            = 5
	smilesParserRULE_isotope           = 6
	smilesParserRULE_element_symbols   = 7
	smilesParserRULE_aromatic_symbols  = 8
	smilesParserRULE_chiral            = 9
	smilesParserRULE_hcount            = 10
	smilesParserRULE_charge            = 11
	smilesParserRULE_class_            = 12
	smilesParserRULE_bond              = 13
	smilesParserRULE_ringbond          = 14
	smilesParserRULE_branched_atom     = 15
	smilesParserRULE_branch            = 16
	smilesParserRULE_chain             = 17
	smilesParserRULE_terminator        = 18
)

// ISmilesContext is an interface to support dynamic dispatch.
type ISmilesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSmilesContext differentiates from other interfaces.
	IsSmilesContext()
}

type SmilesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySmilesContext() *SmilesContext {
	var p = new(SmilesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_smiles
	return p
}

func (*SmilesContext) IsSmilesContext() {}

func NewSmilesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SmilesContext {
	var p = new(SmilesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_smiles

	return p
}

func (s *SmilesContext) GetParser() antlr.Parser { return s.parser }

func (s *SmilesContext) Chain() IChainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChainContext)
}

func (s *SmilesContext) Terminator() ITerminatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerminatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *SmilesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SmilesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SmilesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterSmiles(s)
	}
}

func (s *SmilesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitSmiles(s)
	}
}

func (p *smilesParser) Smiles() (localctx ISmilesContext) {
	localctx = NewSmilesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, smilesParserRULE_smiles)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.chain(0)
	}
	{
		p.SetState(39)
		p.Terminator()
	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Bracket_atom() IBracket_atomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBracket_atomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBracket_atomContext)
}

func (s *AtomContext) Aliphatic_organic() IAliphatic_organicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliphatic_organicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliphatic_organicContext)
}

func (s *AtomContext) Aromatic_organic() IAromatic_organicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAromatic_organicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAromatic_organicContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *smilesParser) Atom() (localctx IAtomContext) {
	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, smilesParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case smilesParserT__17:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Bracket_atom()
		}

	case smilesParserT__1, smilesParserT__2, smilesParserT__3, smilesParserT__4, smilesParserT__5, smilesParserT__6, smilesParserT__7, smilesParserT__8, smilesParserT__9, smilesParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Aliphatic_organic()
		}

	case smilesParserT__11, smilesParserT__12, smilesParserT__13, smilesParserT__14, smilesParserT__15, smilesParserT__16:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.Aromatic_organic()
		}

	case smilesParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(44)
			p.Match(smilesParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAliphatic_organicContext is an interface to support dynamic dispatch.
type IAliphatic_organicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliphatic_organicContext differentiates from other interfaces.
	IsAliphatic_organicContext()
}

type Aliphatic_organicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliphatic_organicContext() *Aliphatic_organicContext {
	var p = new(Aliphatic_organicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_aliphatic_organic
	return p
}

func (*Aliphatic_organicContext) IsAliphatic_organicContext() {}

func NewAliphatic_organicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aliphatic_organicContext {
	var p = new(Aliphatic_organicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_aliphatic_organic

	return p
}

func (s *Aliphatic_organicContext) GetParser() antlr.Parser { return s.parser }
func (s *Aliphatic_organicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aliphatic_organicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Aliphatic_organicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterAliphatic_organic(s)
	}
}

func (s *Aliphatic_organicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitAliphatic_organic(s)
	}
}

func (p *smilesParser) Aliphatic_organic() (localctx IAliphatic_organicContext) {
	localctx = NewAliphatic_organicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, smilesParserRULE_aliphatic_organic)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<smilesParserT__1)|(1<<smilesParserT__2)|(1<<smilesParserT__3)|(1<<smilesParserT__4)|(1<<smilesParserT__5)|(1<<smilesParserT__6)|(1<<smilesParserT__7)|(1<<smilesParserT__8)|(1<<smilesParserT__9)|(1<<smilesParserT__10))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAromatic_organicContext is an interface to support dynamic dispatch.
type IAromatic_organicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAromatic_organicContext differentiates from other interfaces.
	IsAromatic_organicContext()
}

type Aromatic_organicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAromatic_organicContext() *Aromatic_organicContext {
	var p = new(Aromatic_organicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_aromatic_organic
	return p
}

func (*Aromatic_organicContext) IsAromatic_organicContext() {}

func NewAromatic_organicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aromatic_organicContext {
	var p = new(Aromatic_organicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_aromatic_organic

	return p
}

func (s *Aromatic_organicContext) GetParser() antlr.Parser { return s.parser }
func (s *Aromatic_organicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aromatic_organicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Aromatic_organicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterAromatic_organic(s)
	}
}

func (s *Aromatic_organicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitAromatic_organic(s)
	}
}

func (p *smilesParser) Aromatic_organic() (localctx IAromatic_organicContext) {
	localctx = NewAromatic_organicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, smilesParserRULE_aromatic_organic)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<smilesParserT__11)|(1<<smilesParserT__12)|(1<<smilesParserT__13)|(1<<smilesParserT__14)|(1<<smilesParserT__15)|(1<<smilesParserT__16))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBracket_atomContext is an interface to support dynamic dispatch.
type IBracket_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBracket_atomContext differentiates from other interfaces.
	IsBracket_atomContext()
}

type Bracket_atomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracket_atomContext() *Bracket_atomContext {
	var p = new(Bracket_atomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_bracket_atom
	return p
}

func (*Bracket_atomContext) IsBracket_atomContext() {}

func NewBracket_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bracket_atomContext {
	var p = new(Bracket_atomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_bracket_atom

	return p
}

func (s *Bracket_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Bracket_atomContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Bracket_atomContext) Isotope() IIsotopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsotopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsotopeContext)
}

func (s *Bracket_atomContext) Chiral() IChiralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChiralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChiralContext)
}

func (s *Bracket_atomContext) Hcount() IHcountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHcountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHcountContext)
}

func (s *Bracket_atomContext) Charge() IChargeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChargeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChargeContext)
}

func (s *Bracket_atomContext) Class_() IClass_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClass_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClass_Context)
}

func (s *Bracket_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bracket_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bracket_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterBracket_atom(s)
	}
}

func (s *Bracket_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitBracket_atom(s)
	}
}

func (p *smilesParser) Bracket_atom() (localctx IBracket_atomContext) {
	localctx = NewBracket_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, smilesParserRULE_bracket_atom)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(51)
		p.Match(smilesParserT__17)
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == smilesParserNUMBER {
		{
			p.SetState(52)
			p.Isotope()
		}

	}
	{
		p.SetState(55)
		p.Symbol()
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la-123)&-(0x1f+1)) == 0 && ((1<<uint((_la-123)))&((1<<(smilesParserT__122-123))|(1<<(smilesParserT__123-123))|(1<<(smilesParserT__124-123))|(1<<(smilesParserT__125-123))|(1<<(smilesParserT__126-123))|(1<<(smilesParserT__127-123))|(1<<(smilesParserT__128-123))|(1<<(smilesParserT__129-123))|(1<<(smilesParserT__130-123))|(1<<(smilesParserT__131-123))|(1<<(smilesParserT__132-123))|(1<<(smilesParserT__133-123))|(1<<(smilesParserT__134-123))|(1<<(smilesParserT__135-123))|(1<<(smilesParserT__136-123))|(1<<(smilesParserT__137-123))|(1<<(smilesParserT__138-123))|(1<<(smilesParserT__139-123))|(1<<(smilesParserT__140-123))|(1<<(smilesParserT__141-123))|(1<<(smilesParserT__142-123))|(1<<(smilesParserT__143-123))|(1<<(smilesParserT__144-123))|(1<<(smilesParserT__145-123))|(1<<(smilesParserT__146-123))|(1<<(smilesParserT__147-123))|(1<<(smilesParserT__148-123))|(1<<(smilesParserT__149-123))|(1<<(smilesParserT__150-123))|(1<<(smilesParserT__151-123))|(1<<(smilesParserT__152-123))|(1<<(smilesParserT__153-123)))) != 0) || (((_la-155)&-(0x1f+1)) == 0 && ((1<<uint((_la-155)))&((1<<(smilesParserT__154-155))|(1<<(smilesParserT__155-155))|(1<<(smilesParserT__156-155))|(1<<(smilesParserT__157-155))|(1<<(smilesParserT__158-155))|(1<<(smilesParserT__159-155))|(1<<(smilesParserT__160-155))|(1<<(smilesParserT__161-155))|(1<<(smilesParserT__162-155))|(1<<(smilesParserT__163-155))|(1<<(smilesParserT__164-155))|(1<<(smilesParserT__165-155))|(1<<(smilesParserT__166-155))|(1<<(smilesParserT__167-155))|(1<<(smilesParserT__168-155))|(1<<(smilesParserT__169-155))|(1<<(smilesParserT__170-155))|(1<<(smilesParserT__171-155))|(1<<(smilesParserT__172-155))|(1<<(smilesParserT__173-155))|(1<<(smilesParserT__174-155))|(1<<(smilesParserT__175-155))|(1<<(smilesParserT__176-155))|(1<<(smilesParserT__177-155))|(1<<(smilesParserT__178-155))|(1<<(smilesParserT__179-155))|(1<<(smilesParserT__180-155))|(1<<(smilesParserT__181-155))|(1<<(smilesParserT__182-155))|(1<<(smilesParserT__183-155))|(1<<(smilesParserT__184-155))|(1<<(smilesParserT__185-155)))) != 0) || (((_la-187)&-(0x1f+1)) == 0 && ((1<<uint((_la-187)))&((1<<(smilesParserT__186-187))|(1<<(smilesParserT__187-187))|(1<<(smilesParserT__188-187))|(1<<(smilesParserT__189-187))|(1<<(smilesParserT__190-187)))) != 0) {
		{
			p.SetState(56)
			p.Chiral()
		}

	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == smilesParserT__19 {
		{
			p.SetState(59)
			p.Hcount()
		}

	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-192)&-(0x1f+1)) == 0 && ((1<<uint((_la-192)))&((1<<(smilesParserT__191-192))|(1<<(smilesParserT__192-192))|(1<<(smilesParserT__193-192))|(1<<(smilesParserT__194-192)))) != 0 {
		{
			p.SetState(62)
			p.Charge()
		}

	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == smilesParserT__195 {
		{
			p.SetState(65)
			p.Class_()
		}

	}
	{
		p.SetState(68)
		p.Match(smilesParserT__18)
	}

	return localctx
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_symbol
	return p
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) Element_symbols() IElement_symbolsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElement_symbolsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElement_symbolsContext)
}

func (s *SymbolContext) Aromatic_symbols() IAromatic_symbolsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAromatic_symbolsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAromatic_symbolsContext)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitSymbol(s)
	}
}

func (p *smilesParser) Symbol() (localctx ISymbolContext) {
	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, smilesParserRULE_symbol)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case smilesParserT__1, smilesParserT__2, smilesParserT__3, smilesParserT__4, smilesParserT__5, smilesParserT__6, smilesParserT__7, smilesParserT__8, smilesParserT__9, smilesParserT__10, smilesParserT__19, smilesParserT__20, smilesParserT__21, smilesParserT__22, smilesParserT__23, smilesParserT__24, smilesParserT__25, smilesParserT__26, smilesParserT__27, smilesParserT__28, smilesParserT__29, smilesParserT__30, smilesParserT__31, smilesParserT__32, smilesParserT__33, smilesParserT__34, smilesParserT__35, smilesParserT__36, smilesParserT__37, smilesParserT__38, smilesParserT__39, smilesParserT__40, smilesParserT__41, smilesParserT__42, smilesParserT__43, smilesParserT__44, smilesParserT__45, smilesParserT__46, smilesParserT__47, smilesParserT__48, smilesParserT__49, smilesParserT__50, smilesParserT__51, smilesParserT__52, smilesParserT__53, smilesParserT__54, smilesParserT__55, smilesParserT__56, smilesParserT__57, smilesParserT__58, smilesParserT__59, smilesParserT__60, smilesParserT__61, smilesParserT__62, smilesParserT__63, smilesParserT__64, smilesParserT__65, smilesParserT__66, smilesParserT__67, smilesParserT__68, smilesParserT__69, smilesParserT__70, smilesParserT__71, smilesParserT__72, smilesParserT__73, smilesParserT__74, smilesParserT__75, smilesParserT__76, smilesParserT__77, smilesParserT__78, smilesParserT__79, smilesParserT__80, smilesParserT__81, smilesParserT__82, smilesParserT__83, smilesParserT__84, smilesParserT__85, smilesParserT__86, smilesParserT__87, smilesParserT__88, smilesParserT__89, smilesParserT__90, smilesParserT__91, smilesParserT__92, smilesParserT__93, smilesParserT__94, smilesParserT__95, smilesParserT__96, smilesParserT__97, smilesParserT__98, smilesParserT__99, smilesParserT__100, smilesParserT__101, smilesParserT__102, smilesParserT__103, smilesParserT__104, smilesParserT__105, smilesParserT__106, smilesParserT__107, smilesParserT__108, smilesParserT__109, smilesParserT__110, smilesParserT__111, smilesParserT__112, smilesParserT__113, smilesParserT__114, smilesParserT__115, smilesParserT__116, smilesParserT__117, smilesParserT__118, smilesParserT__119:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Element_symbols()
		}

	case smilesParserT__12, smilesParserT__13, smilesParserT__14, smilesParserT__15, smilesParserT__16, smilesParserT__120, smilesParserT__121:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(71)
			p.Aromatic_symbols()
		}

	case smilesParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(72)
			p.Match(smilesParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIsotopeContext is an interface to support dynamic dispatch.
type IIsotopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsotopeContext differentiates from other interfaces.
	IsIsotopeContext()
}

type IsotopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsotopeContext() *IsotopeContext {
	var p = new(IsotopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_isotope
	return p
}

func (*IsotopeContext) IsIsotopeContext() {}

func NewIsotopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsotopeContext {
	var p = new(IsotopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_isotope

	return p
}

func (s *IsotopeContext) GetParser() antlr.Parser { return s.parser }

func (s *IsotopeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(smilesParserNUMBER, 0)
}

func (s *IsotopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsotopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsotopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterIsotope(s)
	}
}

func (s *IsotopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitIsotope(s)
	}
}

func (p *smilesParser) Isotope() (localctx IIsotopeContext) {
	localctx = NewIsotopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, smilesParserRULE_isotope)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(smilesParserNUMBER)
	}

	return localctx
}

// IElement_symbolsContext is an interface to support dynamic dispatch.
type IElement_symbolsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElement_symbolsContext differentiates from other interfaces.
	IsElement_symbolsContext()
}

type Element_symbolsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElement_symbolsContext() *Element_symbolsContext {
	var p = new(Element_symbolsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_element_symbols
	return p
}

func (*Element_symbolsContext) IsElement_symbolsContext() {}

func NewElement_symbolsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Element_symbolsContext {
	var p = new(Element_symbolsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_element_symbols

	return p
}

func (s *Element_symbolsContext) GetParser() antlr.Parser { return s.parser }
func (s *Element_symbolsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Element_symbolsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Element_symbolsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterElement_symbols(s)
	}
}

func (s *Element_symbolsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitElement_symbols(s)
	}
}

func (p *smilesParser) Element_symbols() (localctx IElement_symbolsContext) {
	localctx = NewElement_symbolsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, smilesParserRULE_element_symbols)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<smilesParserT__1)|(1<<smilesParserT__2)|(1<<smilesParserT__3)|(1<<smilesParserT__4)|(1<<smilesParserT__5)|(1<<smilesParserT__6)|(1<<smilesParserT__7)|(1<<smilesParserT__8)|(1<<smilesParserT__9)|(1<<smilesParserT__10)|(1<<smilesParserT__19)|(1<<smilesParserT__20)|(1<<smilesParserT__21)|(1<<smilesParserT__22)|(1<<smilesParserT__23)|(1<<smilesParserT__24)|(1<<smilesParserT__25)|(1<<smilesParserT__26)|(1<<smilesParserT__27)|(1<<smilesParserT__28)|(1<<smilesParserT__29)|(1<<smilesParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(smilesParserT__31-32))|(1<<(smilesParserT__32-32))|(1<<(smilesParserT__33-32))|(1<<(smilesParserT__34-32))|(1<<(smilesParserT__35-32))|(1<<(smilesParserT__36-32))|(1<<(smilesParserT__37-32))|(1<<(smilesParserT__38-32))|(1<<(smilesParserT__39-32))|(1<<(smilesParserT__40-32))|(1<<(smilesParserT__41-32))|(1<<(smilesParserT__42-32))|(1<<(smilesParserT__43-32))|(1<<(smilesParserT__44-32))|(1<<(smilesParserT__45-32))|(1<<(smilesParserT__46-32))|(1<<(smilesParserT__47-32))|(1<<(smilesParserT__48-32))|(1<<(smilesParserT__49-32))|(1<<(smilesParserT__50-32))|(1<<(smilesParserT__51-32))|(1<<(smilesParserT__52-32))|(1<<(smilesParserT__53-32))|(1<<(smilesParserT__54-32))|(1<<(smilesParserT__55-32))|(1<<(smilesParserT__56-32))|(1<<(smilesParserT__57-32))|(1<<(smilesParserT__58-32))|(1<<(smilesParserT__59-32))|(1<<(smilesParserT__60-32))|(1<<(smilesParserT__61-32))|(1<<(smilesParserT__62-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(smilesParserT__63-64))|(1<<(smilesParserT__64-64))|(1<<(smilesParserT__65-64))|(1<<(smilesParserT__66-64))|(1<<(smilesParserT__67-64))|(1<<(smilesParserT__68-64))|(1<<(smilesParserT__69-64))|(1<<(smilesParserT__70-64))|(1<<(smilesParserT__71-64))|(1<<(smilesParserT__72-64))|(1<<(smilesParserT__73-64))|(1<<(smilesParserT__74-64))|(1<<(smilesParserT__75-64))|(1<<(smilesParserT__76-64))|(1<<(smilesParserT__77-64))|(1<<(smilesParserT__78-64))|(1<<(smilesParserT__79-64))|(1<<(smilesParserT__80-64))|(1<<(smilesParserT__81-64))|(1<<(smilesParserT__82-64))|(1<<(smilesParserT__83-64))|(1<<(smilesParserT__84-64))|(1<<(smilesParserT__85-64))|(1<<(smilesParserT__86-64))|(1<<(smilesParserT__87-64))|(1<<(smilesParserT__88-64))|(1<<(smilesParserT__89-64))|(1<<(smilesParserT__90-64))|(1<<(smilesParserT__91-64))|(1<<(smilesParserT__92-64))|(1<<(smilesParserT__93-64))|(1<<(smilesParserT__94-64)))) != 0) || (((_la-96)&-(0x1f+1)) == 0 && ((1<<uint((_la-96)))&((1<<(smilesParserT__95-96))|(1<<(smilesParserT__96-96))|(1<<(smilesParserT__97-96))|(1<<(smilesParserT__98-96))|(1<<(smilesParserT__99-96))|(1<<(smilesParserT__100-96))|(1<<(smilesParserT__101-96))|(1<<(smilesParserT__102-96))|(1<<(smilesParserT__103-96))|(1<<(smilesParserT__104-96))|(1<<(smilesParserT__105-96))|(1<<(smilesParserT__106-96))|(1<<(smilesParserT__107-96))|(1<<(smilesParserT__108-96))|(1<<(smilesParserT__109-96))|(1<<(smilesParserT__110-96))|(1<<(smilesParserT__111-96))|(1<<(smilesParserT__112-96))|(1<<(smilesParserT__113-96))|(1<<(smilesParserT__114-96))|(1<<(smilesParserT__115-96))|(1<<(smilesParserT__116-96))|(1<<(smilesParserT__117-96))|(1<<(smilesParserT__118-96))|(1<<(smilesParserT__119-96)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAromatic_symbolsContext is an interface to support dynamic dispatch.
type IAromatic_symbolsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAromatic_symbolsContext differentiates from other interfaces.
	IsAromatic_symbolsContext()
}

type Aromatic_symbolsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAromatic_symbolsContext() *Aromatic_symbolsContext {
	var p = new(Aromatic_symbolsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_aromatic_symbols
	return p
}

func (*Aromatic_symbolsContext) IsAromatic_symbolsContext() {}

func NewAromatic_symbolsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aromatic_symbolsContext {
	var p = new(Aromatic_symbolsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_aromatic_symbols

	return p
}

func (s *Aromatic_symbolsContext) GetParser() antlr.Parser { return s.parser }
func (s *Aromatic_symbolsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aromatic_symbolsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Aromatic_symbolsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterAromatic_symbols(s)
	}
}

func (s *Aromatic_symbolsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitAromatic_symbols(s)
	}
}

func (p *smilesParser) Aromatic_symbols() (localctx IAromatic_symbolsContext) {
	localctx = NewAromatic_symbolsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, smilesParserRULE_aromatic_symbols)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<smilesParserT__12)|(1<<smilesParserT__13)|(1<<smilesParserT__14)|(1<<smilesParserT__15)|(1<<smilesParserT__16))) != 0) || _la == smilesParserT__120 || _la == smilesParserT__121) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IChiralContext is an interface to support dynamic dispatch.
type IChiralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChiralContext differentiates from other interfaces.
	IsChiralContext()
}

type ChiralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChiralContext() *ChiralContext {
	var p = new(ChiralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_chiral
	return p
}

func (*ChiralContext) IsChiralContext() {}

func NewChiralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChiralContext {
	var p = new(ChiralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_chiral

	return p
}

func (s *ChiralContext) GetParser() antlr.Parser { return s.parser }
func (s *ChiralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChiralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChiralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterChiral(s)
	}
}

func (s *ChiralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitChiral(s)
	}
}

func (p *smilesParser) Chiral() (localctx IChiralContext) {
	localctx = NewChiralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, smilesParserRULE_chiral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-123)&-(0x1f+1)) == 0 && ((1<<uint((_la-123)))&((1<<(smilesParserT__122-123))|(1<<(smilesParserT__123-123))|(1<<(smilesParserT__124-123))|(1<<(smilesParserT__125-123))|(1<<(smilesParserT__126-123))|(1<<(smilesParserT__127-123))|(1<<(smilesParserT__128-123))|(1<<(smilesParserT__129-123))|(1<<(smilesParserT__130-123))|(1<<(smilesParserT__131-123))|(1<<(smilesParserT__132-123))|(1<<(smilesParserT__133-123))|(1<<(smilesParserT__134-123))|(1<<(smilesParserT__135-123))|(1<<(smilesParserT__136-123))|(1<<(smilesParserT__137-123))|(1<<(smilesParserT__138-123))|(1<<(smilesParserT__139-123))|(1<<(smilesParserT__140-123))|(1<<(smilesParserT__141-123))|(1<<(smilesParserT__142-123))|(1<<(smilesParserT__143-123))|(1<<(smilesParserT__144-123))|(1<<(smilesParserT__145-123))|(1<<(smilesParserT__146-123))|(1<<(smilesParserT__147-123))|(1<<(smilesParserT__148-123))|(1<<(smilesParserT__149-123))|(1<<(smilesParserT__150-123))|(1<<(smilesParserT__151-123))|(1<<(smilesParserT__152-123))|(1<<(smilesParserT__153-123)))) != 0) || (((_la-155)&-(0x1f+1)) == 0 && ((1<<uint((_la-155)))&((1<<(smilesParserT__154-155))|(1<<(smilesParserT__155-155))|(1<<(smilesParserT__156-155))|(1<<(smilesParserT__157-155))|(1<<(smilesParserT__158-155))|(1<<(smilesParserT__159-155))|(1<<(smilesParserT__160-155))|(1<<(smilesParserT__161-155))|(1<<(smilesParserT__162-155))|(1<<(smilesParserT__163-155))|(1<<(smilesParserT__164-155))|(1<<(smilesParserT__165-155))|(1<<(smilesParserT__166-155))|(1<<(smilesParserT__167-155))|(1<<(smilesParserT__168-155))|(1<<(smilesParserT__169-155))|(1<<(smilesParserT__170-155))|(1<<(smilesParserT__171-155))|(1<<(smilesParserT__172-155))|(1<<(smilesParserT__173-155))|(1<<(smilesParserT__174-155))|(1<<(smilesParserT__175-155))|(1<<(smilesParserT__176-155))|(1<<(smilesParserT__177-155))|(1<<(smilesParserT__178-155))|(1<<(smilesParserT__179-155))|(1<<(smilesParserT__180-155))|(1<<(smilesParserT__181-155))|(1<<(smilesParserT__182-155))|(1<<(smilesParserT__183-155))|(1<<(smilesParserT__184-155))|(1<<(smilesParserT__185-155)))) != 0) || (((_la-187)&-(0x1f+1)) == 0 && ((1<<uint((_la-187)))&((1<<(smilesParserT__186-187))|(1<<(smilesParserT__187-187))|(1<<(smilesParserT__188-187))|(1<<(smilesParserT__189-187))|(1<<(smilesParserT__190-187)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IHcountContext is an interface to support dynamic dispatch.
type IHcountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHcountContext differentiates from other interfaces.
	IsHcountContext()
}

type HcountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHcountContext() *HcountContext {
	var p = new(HcountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_hcount
	return p
}

func (*HcountContext) IsHcountContext() {}

func NewHcountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HcountContext {
	var p = new(HcountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_hcount

	return p
}

func (s *HcountContext) GetParser() antlr.Parser { return s.parser }

func (s *HcountContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(smilesParserDIGIT, 0)
}

func (s *HcountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HcountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HcountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterHcount(s)
	}
}

func (s *HcountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitHcount(s)
	}
}

func (p *smilesParser) Hcount() (localctx IHcountContext) {
	localctx = NewHcountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, smilesParserRULE_hcount)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Match(smilesParserT__19)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.Match(smilesParserT__19)
		}
		{
			p.SetState(85)
			p.Match(smilesParserDIGIT)
		}

	}

	return localctx
}

// IChargeContext is an interface to support dynamic dispatch.
type IChargeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChargeContext differentiates from other interfaces.
	IsChargeContext()
}

type ChargeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChargeContext() *ChargeContext {
	var p = new(ChargeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_charge
	return p
}

func (*ChargeContext) IsChargeContext() {}

func NewChargeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChargeContext {
	var p = new(ChargeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_charge

	return p
}

func (s *ChargeContext) GetParser() antlr.Parser { return s.parser }

func (s *ChargeContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(smilesParserDIGIT, 0)
}

func (s *ChargeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChargeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChargeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterCharge(s)
	}
}

func (s *ChargeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitCharge(s)
	}
}

func (p *smilesParser) Charge() (localctx IChargeContext) {
	localctx = NewChargeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, smilesParserRULE_charge)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(smilesParserT__191)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Match(smilesParserT__191)
		}
		{
			p.SetState(90)
			p.Match(smilesParserDIGIT)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(91)
			p.Match(smilesParserT__192)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(92)
			p.Match(smilesParserT__192)
		}
		{
			p.SetState(93)
			p.Match(smilesParserDIGIT)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)
			p.Match(smilesParserT__193)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(95)
			p.Match(smilesParserT__194)
		}

	}

	return localctx
}

// IClass_Context is an interface to support dynamic dispatch.
type IClass_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClass_Context differentiates from other interfaces.
	IsClass_Context()
}

type Class_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClass_Context() *Class_Context {
	var p = new(Class_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_class_
	return p
}

func (*Class_Context) IsClass_Context() {}

func NewClass_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Class_Context {
	var p = new(Class_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_class_

	return p
}

func (s *Class_Context) GetParser() antlr.Parser { return s.parser }

func (s *Class_Context) NUMBER() antlr.TerminalNode {
	return s.GetToken(smilesParserNUMBER, 0)
}

func (s *Class_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Class_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Class_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterClass_(s)
	}
}

func (s *Class_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitClass_(s)
	}
}

func (p *smilesParser) Class_() (localctx IClass_Context) {
	localctx = NewClass_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, smilesParserRULE_class_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(smilesParserT__195)
	}
	{
		p.SetState(99)
		p.Match(smilesParserNUMBER)
	}

	return localctx
}

// IBondContext is an interface to support dynamic dispatch.
type IBondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBondContext differentiates from other interfaces.
	IsBondContext()
}

type BondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBondContext() *BondContext {
	var p = new(BondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_bond
	return p
}

func (*BondContext) IsBondContext() {}

func NewBondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BondContext {
	var p = new(BondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_bond

	return p
}

func (s *BondContext) GetParser() antlr.Parser { return s.parser }
func (s *BondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterBond(s)
	}
}

func (s *BondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitBond(s)
	}
}

func (p *smilesParser) Bond() (localctx IBondContext) {
	localctx = NewBondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, smilesParserRULE_bond)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-192)&-(0x1f+1)) == 0 && ((1<<uint((_la-192)))&((1<<(smilesParserT__191-192))|(1<<(smilesParserT__195-192))|(1<<(smilesParserT__196-192))|(1<<(smilesParserT__197-192))|(1<<(smilesParserT__198-192))|(1<<(smilesParserT__199-192))|(1<<(smilesParserT__200-192)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRingbondContext is an interface to support dynamic dispatch.
type IRingbondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRingbondContext differentiates from other interfaces.
	IsRingbondContext()
}

type RingbondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRingbondContext() *RingbondContext {
	var p = new(RingbondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_ringbond
	return p
}

func (*RingbondContext) IsRingbondContext() {}

func NewRingbondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RingbondContext {
	var p = new(RingbondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_ringbond

	return p
}

func (s *RingbondContext) GetParser() antlr.Parser { return s.parser }

func (s *RingbondContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(smilesParserDIGIT)
}

func (s *RingbondContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(smilesParserDIGIT, i)
}

func (s *RingbondContext) Bond() IBondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBondContext)
}

func (s *RingbondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RingbondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RingbondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterRingbond(s)
	}
}

func (s *RingbondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitRingbond(s)
	}
}

func (p *smilesParser) Ringbond() (localctx IRingbondContext) {
	localctx = NewRingbondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, smilesParserRULE_ringbond)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-192)&-(0x1f+1)) == 0 && ((1<<uint((_la-192)))&((1<<(smilesParserT__191-192))|(1<<(smilesParserT__195-192))|(1<<(smilesParserT__196-192))|(1<<(smilesParserT__197-192))|(1<<(smilesParserT__198-192))|(1<<(smilesParserT__199-192))|(1<<(smilesParserT__200-192)))) != 0 {
			{
				p.SetState(103)
				p.Bond()
			}

		}
		{
			p.SetState(106)
			p.Match(smilesParserDIGIT)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-192)&-(0x1f+1)) == 0 && ((1<<uint((_la-192)))&((1<<(smilesParserT__191-192))|(1<<(smilesParserT__195-192))|(1<<(smilesParserT__196-192))|(1<<(smilesParserT__197-192))|(1<<(smilesParserT__198-192))|(1<<(smilesParserT__199-192))|(1<<(smilesParserT__200-192)))) != 0 {
			{
				p.SetState(107)
				p.Bond()
			}

		}
		{
			p.SetState(110)
			p.Match(smilesParserT__201)
		}
		{
			p.SetState(111)
			p.Match(smilesParserDIGIT)
		}
		{
			p.SetState(112)
			p.Match(smilesParserDIGIT)
		}

	}

	return localctx
}

// IBranched_atomContext is an interface to support dynamic dispatch.
type IBranched_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBranched_atomContext differentiates from other interfaces.
	IsBranched_atomContext()
}

type Branched_atomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranched_atomContext() *Branched_atomContext {
	var p = new(Branched_atomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_branched_atom
	return p
}

func (*Branched_atomContext) IsBranched_atomContext() {}

func NewBranched_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Branched_atomContext {
	var p = new(Branched_atomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_branched_atom

	return p
}

func (s *Branched_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Branched_atomContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *Branched_atomContext) AllRingbond() []IRingbondContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRingbondContext)(nil)).Elem())
	var tst = make([]IRingbondContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRingbondContext)
		}
	}

	return tst
}

func (s *Branched_atomContext) Ringbond(i int) IRingbondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRingbondContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRingbondContext)
}

func (s *Branched_atomContext) AllBranch() []IBranchContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBranchContext)(nil)).Elem())
	var tst = make([]IBranchContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBranchContext)
		}
	}

	return tst
}

func (s *Branched_atomContext) Branch(i int) IBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *Branched_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Branched_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Branched_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterBranched_atom(s)
	}
}

func (s *Branched_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitBranched_atom(s)
	}
}

func (p *smilesParser) Branched_atom() (localctx IBranched_atomContext) {
	localctx = NewBranched_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, smilesParserRULE_branched_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Atom()
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(116)
				p.Ringbond()
			}

		}
		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(122)
				p.Branch()
			}

		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IBranchContext is an interface to support dynamic dispatch.
type IBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBranchContext differentiates from other interfaces.
	IsBranchContext()
}

type BranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranchContext() *BranchContext {
	var p = new(BranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_branch
	return p
}

func (*BranchContext) IsBranchContext() {}

func NewBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchContext {
	var p = new(BranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_branch

	return p
}

func (s *BranchContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchContext) Chain() IChainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChainContext)
}

func (s *BranchContext) Bond() IBondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBondContext)
}

func (s *BranchContext) DOT() antlr.TerminalNode {
	return s.GetToken(smilesParserDOT, 0)
}

func (s *BranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterBranch(s)
	}
}

func (s *BranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitBranch(s)
	}
}

func (p *smilesParser) Branch() (localctx IBranchContext) {
	localctx = NewBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, smilesParserRULE_branch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(128)
			p.Match(smilesParserT__202)
		}
		{
			p.SetState(129)
			p.chain(0)
		}
		{
			p.SetState(130)
			p.Match(smilesParserT__203)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Match(smilesParserT__202)
		}
		{
			p.SetState(133)
			p.Bond()
		}
		{
			p.SetState(134)
			p.chain(0)
		}
		{
			p.SetState(135)
			p.Match(smilesParserT__203)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(137)
			p.Match(smilesParserT__202)
		}
		{
			p.SetState(138)
			p.Match(smilesParserDOT)
		}
		{
			p.SetState(139)
			p.chain(0)
		}
		{
			p.SetState(140)
			p.Match(smilesParserT__203)
		}

	}

	return localctx
}

// IChainContext is an interface to support dynamic dispatch.
type IChainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChainContext differentiates from other interfaces.
	IsChainContext()
}

type ChainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChainContext() *ChainContext {
	var p = new(ChainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_chain
	return p
}

func (*ChainContext) IsChainContext() {}

func NewChainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChainContext {
	var p = new(ChainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_chain

	return p
}

func (s *ChainContext) GetParser() antlr.Parser { return s.parser }

func (s *ChainContext) Branched_atom() IBranched_atomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranched_atomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBranched_atomContext)
}

func (s *ChainContext) Chain() IChainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChainContext)
}

func (s *ChainContext) Bond() IBondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBondContext)
}

func (s *ChainContext) DOT() antlr.TerminalNode {
	return s.GetToken(smilesParserDOT, 0)
}

func (s *ChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterChain(s)
	}
}

func (s *ChainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitChain(s)
	}
}

func (p *smilesParser) Chain() (localctx IChainContext) {
	return p.chain(0)
}

func (p *smilesParser) chain(_p int) (localctx IChainContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewChainContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IChainContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, smilesParserRULE_chain, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.Branched_atom()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(156)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewChainContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, smilesParserRULE_chain)
				p.SetState(147)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(148)
					p.Branched_atom()
				}

			case 2:
				localctx = NewChainContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, smilesParserRULE_chain)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(150)
					p.Bond()
				}
				{
					p.SetState(151)
					p.Branched_atom()
				}

			case 3:
				localctx = NewChainContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, smilesParserRULE_chain)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(154)
					p.Match(smilesParserDOT)
				}
				{
					p.SetState(155)
					p.Branched_atom()
				}

			}

		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// ITerminatorContext is an interface to support dynamic dispatch.
type ITerminatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerminatorContext differentiates from other interfaces.
	IsTerminatorContext()
}

type TerminatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminatorContext() *TerminatorContext {
	var p = new(TerminatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_terminator
	return p
}

func (*TerminatorContext) IsTerminatorContext() {}

func NewTerminatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerminatorContext {
	var p = new(TerminatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_terminator

	return p
}

func (s *TerminatorContext) GetParser() antlr.Parser { return s.parser }

func (s *TerminatorContext) SPACE() antlr.TerminalNode {
	return s.GetToken(smilesParserSPACE, 0)
}

func (s *TerminatorContext) TAB() antlr.TerminalNode {
	return s.GetToken(smilesParserTAB, 0)
}

func (s *TerminatorContext) LINEFEED() antlr.TerminalNode {
	return s.GetToken(smilesParserLINEFEED, 0)
}

func (s *TerminatorContext) CARRIAGE_RETURN() antlr.TerminalNode {
	return s.GetToken(smilesParserCARRIAGE_RETURN, 0)
}

func (s *TerminatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerminatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterTerminator(s)
	}
}

func (s *TerminatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitTerminator(s)
	}
}

func (p *smilesParser) Terminator() (localctx ITerminatorContext) {
	localctx = NewTerminatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, smilesParserRULE_terminator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case smilesParserSPACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Match(smilesParserSPACE)
		}
		{
			p.SetState(162)
			p.Match(smilesParserTAB)
		}

	case smilesParserLINEFEED:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.Match(smilesParserLINEFEED)
		}

	case smilesParserCARRIAGE_RETURN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(164)
			p.Match(smilesParserCARRIAGE_RETURN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *smilesParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 17:
		var t *ChainContext = nil
		if localctx != nil {
			t = localctx.(*ChainContext)
		}
		return p.Chain_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *smilesParser) Chain_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
