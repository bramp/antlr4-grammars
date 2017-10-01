// Generated from UCBLogo.g4 by ANTLR 4.7.

package ucb_logo // UCBLogo
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)



  import java.util.Map;
  import java.util.HashMap;


// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 183, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 3, 2, 7, 2, 30, 10, 2, 12, 2, 14, 2, 33, 11, 2, 3, 
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 41, 10, 3, 3, 4, 3, 4, 3, 4, 3, 
	4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 7, 
	6, 58, 10, 6, 12, 6, 14, 6, 61, 11, 6, 3, 7, 3, 7, 7, 7, 65, 10, 7, 12, 
	7, 14, 7, 68, 11, 7, 3, 7, 3, 7, 7, 7, 72, 10, 7, 12, 7, 14, 7, 75, 11, 
	7, 3, 7, 5, 7, 78, 10, 7, 3, 8, 3, 8, 5, 8, 82, 10, 8, 3, 9, 3, 9, 3, 9, 
	3, 9, 7, 9, 88, 10, 9, 12, 9, 14, 9, 91, 11, 9, 3, 9, 3, 9, 3, 10, 3, 10, 
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 103, 10, 11, 12, 11, 14, 
	11, 106, 11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 
	12, 126, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 
	3, 12, 7, 12, 158, 10, 12, 12, 12, 14, 12, 161, 11, 12, 3, 13, 3, 13, 3, 
	13, 7, 13, 166, 10, 13, 12, 13, 14, 13, 169, 11, 13, 3, 13, 3, 13, 3, 14, 
	3, 14, 3, 14, 7, 14, 176, 10, 14, 12, 14, 14, 14, 179, 11, 14, 3, 14, 3, 
	14, 3, 14, 2, 3, 22, 15, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 
	2, 5, 3, 2, 6, 6, 3, 2, 10, 11, 3, 2, 12, 13, 2, 204, 2, 31, 3, 2, 2, 2, 
	4, 40, 3, 2, 2, 2, 6, 42, 3, 2, 2, 2, 8, 48, 3, 2, 2, 2, 10, 54, 3, 2, 
	2, 2, 12, 77, 3, 2, 2, 2, 14, 81, 3, 2, 2, 2, 16, 83, 3, 2, 2, 2, 18, 94, 
	3, 2, 2, 2, 20, 104, 3, 2, 2, 2, 22, 125, 3, 2, 2, 2, 24, 162, 3, 2, 2, 
	2, 26, 172, 3, 2, 2, 2, 28, 30, 5, 4, 3, 2, 29, 28, 3, 2, 2, 2, 30, 33, 
	3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 34, 3, 2, 2, 2, 
	33, 31, 3, 2, 2, 2, 34, 35, 7, 2, 2, 3, 35, 3, 3, 2, 2, 2, 36, 41, 5, 6, 
	4, 2, 37, 41, 5, 8, 5, 2, 38, 41, 5, 16, 9, 2, 39, 41, 5, 18, 10, 2, 40, 
	36, 3, 2, 2, 2, 40, 37, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 39, 3, 2, 2, 
	2, 41, 5, 3, 2, 2, 2, 42, 43, 7, 5, 2, 2, 43, 44, 7, 27, 2, 2, 44, 45, 
	5, 10, 6, 2, 45, 46, 5, 12, 7, 2, 46, 47, 8, 4, 1, 2, 47, 7, 3, 2, 2, 2, 
	48, 49, 7, 7, 2, 2, 49, 50, 7, 27, 2, 2, 50, 51, 5, 10, 6, 2, 51, 52, 5, 
	12, 7, 2, 52, 53, 8, 5, 1, 2, 53, 9, 3, 2, 2, 2, 54, 59, 8, 6, 1, 2, 55, 
	56, 7, 26, 2, 2, 56, 58, 8, 6, 1, 2, 57, 55, 3, 2, 2, 2, 58, 61, 3, 2, 
	2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 11, 3, 2, 2, 2, 61, 59, 
	3, 2, 2, 2, 62, 66, 6, 7, 2, 2, 63, 65, 5, 14, 8, 2, 64, 63, 3, 2, 2, 2, 
	65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 69, 3, 
	2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 78, 7, 6, 2, 2, 70, 72, 10, 2, 2, 2, 71, 
	70, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 
	2, 74, 76, 3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 76, 78, 7, 6, 2, 2, 77, 62, 
	3, 2, 2, 2, 77, 73, 3, 2, 2, 2, 78, 13, 3, 2, 2, 2, 79, 82, 5, 16, 9, 2, 
	80, 82, 5, 18, 10, 2, 81, 79, 3, 2, 2, 2, 81, 80, 3, 2, 2, 2, 82, 15, 3, 
	2, 2, 2, 83, 84, 7, 3, 2, 2, 84, 85, 6, 9, 3, 2, 85, 89, 7, 27, 2, 2, 86, 
	88, 5, 22, 12, 2, 87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 
	2, 2, 89, 90, 3, 2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 
	7, 4, 2, 2, 93, 17, 3, 2, 2, 2, 94, 95, 6, 10, 4, 2, 95, 96, 7, 27, 2, 
	2, 96, 97, 5, 20, 11, 2, 97, 19, 3, 2, 2, 2, 98, 99, 6, 11, 5, 3, 99, 100, 
	5, 22, 12, 2, 100, 101, 8, 11, 1, 2, 101, 103, 3, 2, 2, 2, 102, 98, 3, 
	2, 2, 2, 103, 106, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 104, 105, 3, 2, 2, 
	2, 105, 107, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 107, 108, 8, 11, 1, 2, 108, 
	21, 3, 2, 2, 2, 109, 110, 8, 12, 1, 2, 110, 111, 7, 14, 2, 2, 111, 126, 
	5, 22, 12, 23, 112, 126, 5, 16, 9, 2, 113, 126, 5, 18, 10, 2, 114, 115, 
	7, 3, 2, 2, 115, 116, 5, 22, 12, 2, 116, 117, 7, 4, 2, 2, 117, 126, 3, 
	2, 2, 2, 118, 126, 5, 24, 13, 2, 119, 126, 5, 26, 14, 2, 120, 126, 7, 8, 
	2, 2, 121, 126, 7, 24, 2, 2, 122, 126, 7, 25, 2, 2, 123, 126, 7, 26, 2, 
	2, 124, 126, 7, 27, 2, 2, 125, 109, 3, 2, 2, 2, 125, 112, 3, 2, 2, 2, 125, 
	113, 3, 2, 2, 2, 125, 114, 3, 2, 2, 2, 125, 118, 3, 2, 2, 2, 125, 119, 
	3, 2, 2, 2, 125, 120, 3, 2, 2, 2, 125, 121, 3, 2, 2, 2, 125, 122, 3, 2, 
	2, 2, 125, 123, 3, 2, 2, 2, 125, 124, 3, 2, 2, 2, 126, 159, 3, 2, 2, 2, 
	127, 128, 12, 12, 2, 2, 128, 129, 7, 16, 2, 2, 129, 158, 5, 22, 12, 13, 
	130, 131, 12, 11, 2, 2, 131, 132, 7, 17, 2, 2, 132, 158, 5, 22, 12, 12, 
	133, 134, 12, 10, 2, 2, 134, 135, 7, 15, 2, 2, 135, 158, 5, 22, 12, 11, 
	136, 137, 12, 9, 2, 2, 137, 138, 7, 14, 2, 2, 138, 158, 5, 22, 12, 10, 
	139, 140, 12, 8, 2, 2, 140, 141, 7, 18, 2, 2, 141, 158, 5, 22, 12, 9, 142, 
	143, 12, 7, 2, 2, 143, 144, 7, 19, 2, 2, 144, 158, 5, 22, 12, 8, 145, 146, 
	12, 6, 2, 2, 146, 147, 7, 21, 2, 2, 147, 158, 5, 22, 12, 7, 148, 149, 12, 
	5, 2, 2, 149, 150, 7, 22, 2, 2, 150, 158, 5, 22, 12, 6, 151, 152, 12, 4, 
	2, 2, 152, 153, 7, 20, 2, 2, 153, 158, 5, 22, 12, 5, 154, 155, 12, 3, 2, 
	2, 155, 156, 7, 23, 2, 2, 156, 158, 5, 22, 12, 4, 157, 127, 3, 2, 2, 2, 
	157, 130, 3, 2, 2, 2, 157, 133, 3, 2, 2, 2, 157, 136, 3, 2, 2, 2, 157, 
	139, 3, 2, 2, 2, 157, 142, 3, 2, 2, 2, 157, 145, 3, 2, 2, 2, 157, 148, 
	3, 2, 2, 2, 157, 151, 3, 2, 2, 2, 157, 154, 3, 2, 2, 2, 158, 161, 3, 2, 
	2, 2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 23, 3, 2, 2, 2, 
	161, 159, 3, 2, 2, 2, 162, 167, 7, 10, 2, 2, 163, 166, 10, 3, 2, 2, 164, 
	166, 5, 24, 13, 2, 165, 163, 3, 2, 2, 2, 165, 164, 3, 2, 2, 2, 166, 169, 
	3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 170, 3, 2, 
	2, 2, 169, 167, 3, 2, 2, 2, 170, 171, 7, 11, 2, 2, 171, 25, 3, 2, 2, 2, 
	172, 177, 7, 12, 2, 2, 173, 176, 10, 4, 2, 2, 174, 176, 5, 26, 14, 2, 175, 
	173, 3, 2, 2, 2, 175, 174, 3, 2, 2, 2, 176, 179, 3, 2, 2, 2, 177, 175, 
	3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 180, 3, 2, 2, 2, 179, 177, 3, 2, 
	2, 2, 180, 181, 7, 13, 2, 2, 181, 27, 3, 2, 2, 2, 18, 31, 40, 59, 66, 73, 
	77, 81, 89, 104, 125, 157, 159, 165, 167, 175, 177,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "", "", "", "", "", "'{'", "'}'", "'['", "']'", "'-'", 
	"'+'", "'*'", "'/'", "'<'", "'>'", "'='", "'<='", "'>='", "'<>'",
}
var symbolicNames = []string{
	"", "", "", "TO", "END", "MACRO", "WORD", "SKIP_", "OPEN_ARRAY", "CLOSE_ARRAY", 
	"OPEN_LIST", "CLOSE_LIST", "MINUS", "PLUS", "MULT", "DIV", "LT", "GT", 
	"EQ", "LT_EQ", "GT_EQ", "NOT_EQ", "QUOTED_WORD", "NUMBER", "VARIABLE", 
	"NAME", "ANY",
}

var ruleNames = []string{
	"parse", "instruction", "procedure_def", "macro_def", "variables", "body_def", 
	"body_instruction", "procedure_call_extra_input", "procedure_call", "expressions", 
	"expression", "array", "list",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type UCBLogoParser struct {
	*antlr.BaseParser
}

func NewUCBLogoParser(input antlr.TokenStream) *UCBLogoParser {
	this := new(UCBLogoParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "UCBLogo.g4"

	return this
}



  // A Map keeping track of all procedure (and macro) names and the amount
  // of parameters each procedure expects.
  // Taken from: http://www.cs.berkeley.edu/~bh/usermanual
  final Map<String, Integer> procedures = new HashMap<String, Integer>(){{
    put("word", 2);
    put("list", 2);
    put("sentence", 2);
    put("se", 2);
    put("fput", 2);
    put("lput", 2);
    put("array", 1);
    put("mdarray", 1);
    put("listtoarray", 1);
    put("arraytolist", 1);
    put("combine", 2);
    put("reverse", 1);
    put("gensym", 0);
    put("first", 1);
    put("firsts", 1);
    put("last", 1);
    put("butfirst", 1);
    put("bf", 1);
    put("butfirsts", 1);
    put("bfs", 1);
    put("butlast", 1);
    put("bl", 1);
    put("item", 2);
    put("mditem", 2);
    put("pick", 1);
    put("remove", 2);
    put("remdup", 1);
    put("quoted", 1);
    put("setitem", 3);
    put("mdsetitem", 3);
    put(".setfirst", 2);
    put(".setbf", 2);
    put(".setitem", 3);
    put("push", 2);
    put("pop", 1);
    put("queue", 2);
    put("dequeue", 1);
    put("wordp", 1);
    put("word?", 1);
    put("listp", 1);
    put("list?", 1);
    put("arrayp", 1);
    put("array?", 1);
    put("emptyp", 1);
    put("empty?", 1);
    put("equalp", 2);
    put("equal?", 2);
    put("notequalp", 2);
    put("notequal?", 2);
    put("beforep", 2);
    put("before?", 2);
    put(".eq", 2);
    put("memberp", 2);
    put("member?", 2);
    put("substringp", 2);
    put("substring?", 2);
    put("numberp", 1);
    put("number?", 1);
    put("vbarredp", 1);
    put("vbarred?", 1);
    put("backslashedp", 1);
    put("backslashed?", 1);
    put("count", 1);
    put("ascii", 1);
    put("rawascii", 1);
    put("char", 1);
    put("member", 2);
    put("lowercase", 1);
    put("uppercase", 1);
    put("standout", 1);
    put("parse", 1);
    put("runparse", 1);
    put("print", 1);
    put("pr", 1);
    put("type", 1);
    put("show", 1);
    put("readlist", 0);
    put("rl", 0);
    put("readword", 0);
    put("rw", 0);
    put("readrawline", 0);
    put("readchar", 0);
    put("rc", 0);
    put("readchars", 1);
    put("rcs", 1);
    put("shell", 1);
    put("setprefix", 1);
    put("prefix", 0);
    put("openread", 1);
    put("openwrite", 1);
    put("openappend", 1);
    put("openupdate", 1);
    put("close", 1);
    put("allopen", 0);
    put("closeall", 0);
    put("erasefile", 1);
    put("erf", 1);
    put("dribble", 1);
    put("nodribble", 0);
    put("setread", 1);
    put("setwrite", 1);
    put("reader", 0);
    put("writer", 0);
    put("setreadpos", 1);
    put("setwritepos", 1);
    put("readpos", 0);
    put("writepos", 0);
    put("eofp", 0);
    put("eof?", 0);
    put("filep", 1);
    put("file?", 1);
    put("keyp", 0);
    put("key?", 0);
    put("cleartext", 0);
    put("ct", 0);
    put("setcursor", 1);
    put("cursor", 0);
    put("setmargins", 1);
    put("settextcolor", 2);
    put("settc", 2);
    put("increasefont", 0);
    put("decreasefont", 0);
    put("settextsize", 1);
    put("textsize", 0);
    put("setfont", 1);
    put("font", 0);
    put("sum", 2);
    put("difference", 2);
    put("minus", 1);
    put("product", 2);
    put("quotient", 2);
    put("remainder", 2);
    put("modulo", 2);
    put("int", 1);
    put("round", 1);
    put("sqrt", 1);
    put("power", 2);
    put("exp", 1);
    put("log10", 1);
    put("ln", 1);
    put("sin", 1);
    put("radsin", 1);
    put("cos", 1);
    put("radcos", 1);
    put("arctan", 1);
    put("radarctan", 1);
    put("iseq", 2);
    put("rseq", 3);
    put("lessp", 2);
    put("less?", 2);
    put("greaterp", 2);
    put("greater?", 2);
    put("lessequalp", 2);
    put("lessequal?", 2);
    put("greaterequalp", 2);
    put("greaterequal?", 2);
    put("random", 1);
    put("rerandom", 0);
    put("form", 3);
    put("bitand", 2);
    put("bitor", 2);
    put("bitxor", 2);
    put("bitnot", 1);
    put("ashift", 2);
    put("lshift", 2);
    put("and", 2);
    put("or", 2);
    put("not", 1);
    put("forward", 1);
    put("fd", 1);
    put("back", 1);
    put("bk", 1);
    put("left", 1);
    put("lt", 1);
    put("right", 1);
    put("rt", 1);
    put("setpos", 1);
    put("setxy", 2);
    put("setx", 1);
    put("sety", 1);
    put("setheading", 1);
    put("seth", 1);
    put("home", 0);
    put("arc", 2);
    put("pos", 0);
    put("xcor", 0);
    put("ycor", 0);
    put("heading", 0);
    put("towards", 1);
    put("scrunch", 0);
    put("showturtle", 0);
    put("st", 0);
    put("hideturtle", 0);
    put("ht", 0);
    put("clean", 0);
    put("clearscreen", 0);
    put("cs", 0);
    put("wrap", 0);
    put("window", 0);
    put("fence", 0);
    put("fill", 0);
    put("filled", 2);
    put("label", 1);
    put("setlabelheight", 1);
    put("textscreen", 0);
    put("ts", 0);
    put("fullscreen", 0);
    put("fs", 0);
    put("splitscreen", 0);
    put("ss", 0);
    put("setscrunch", 2);
    put("refresh", 0);
    put("norefresh", 0);
    put("shownp", 0);
    put("shown?", 0);
    put("screenmode", 0);
    put("turtlemode", 0);
    put("labelsize", 0);
    put("pendown", 0);
    put("pd", 0);
    put("penup", 0);
    put("pu", 0);
    put("penpaint", 0);
    put("ppt", 0);
    put("penerase", 0);
    put("pe", 0);
    put("penreverse", 0);
    put("px", 0);
    put("setpencolor", 1);
    put("setpc", 1);
    put("setpalette", 2);
    put("setpensize", 1);
    put("setpenpattern", 1);
    put("setpen", 1);
    put("setbackground", 1);
    put("setbg", 1);
    put("pendownp", 0);
    put("pendown?", 0);
    put("penmode", 0);
    put("pencolor", 0);
    put("pc", 0);
    put("palette", 1);
    put("pensize", 0);
    put("penpattern", 0);
    put("pen", 0);
    put("background", 0);
    put("bg", 0);
    put("savepict", 1);
    put("loadpict", 1);
    put("epspict", 1);
    put("mousepos", 0);
    put("clickpos", 0);
    put("buttonp", 0);
    put("button?", 0);
    put("button", 0);
    put("define", 2);
    put("text", 1);
    put("fulltext", 1);
    put("copydef", 2);
    put("make", 2);
    put("name", 2);
    put("local", 1);
    put("localmake", 2);
    put("thing", 1);
    put(":quoted.varname", 0);
    put("global", 1);
    put("pprop", 3);
    put("gprop", 2);
    put("remprop", 2);
    put("plist", 1);
    put("procedurep", 1);
    put("procedure?", 1);
    put("primitivep", 1);
    put("primitive?", 1);
    put("definedp", 1);
    put("defined?", 1);
    put("namep", 1);
    put("name?", 1);
    put("plistp", 1);
    put("plist?", 1);
    put("contents", 0);
    put("buried", 0);
    put("traced", 0);
    put("stepped", 0);
    put("procedures", 0);
    put("primitives", 0);
    put("names", 0);
    put("plists", 0);
    put("namelist", 1);
    put("pllist", 1);
    put("arity", 1);
    put("nodes", 0);
    put("printout", 1);
    put("po", 1);
    put("poall", 0);
    put("pops", 0);
    put("pons", 0);
    put("popls", 0);
    put("pon", 1);
    put("popl", 1);
    put("pot", 1);
    put("pots", 0);
    put("erase", 1);
    put("er", 1);
    put("erall", 0);
    put("erps", 0);
    put("erns", 0);
    put("erpls", 0);
    put("ern", 1);
    put("erpl", 1);
    put("bury", 1);
    put("buryall", 0);
    put("buryname", 1);
    put("unbury", 1);
    put("unburyall", 0);
    put("unburyname", 1);
    put("buriedp", 1);
    put("buried?", 1);
    put("trace", 1);
    put("untrace", 1);
    put("tracedp", 1);
    put("traced?", 1);
    put("step", 1);
    put("unstep", 1);
    put("steppedp", 1);
    put("stepped?", 1);
    put("edit", 1);
    put("ed", 1);
    put("editfile", 1);
    put("edall", 0);
    put("edps", 0);
    put("edns", 0);
    put("edpls", 0);
    put("edn", 1);
    put("edpl", 1);
    put("save", 1);
    put("savel", 2);
    put("load", 1);
    put("cslsload", 1);
    put("help", 1);
    put("seteditor", 1);
    put("setlibloc", 1);
    put("sethelploc", 1);
    put("setcslsloc", 1);
    put("settemploc", 1);
    put("gc", 0);
    put(".setsegmentsize", 1);
    put("run", 1);
    put("runresult", 1);
    put("repeat", 2);
    put("forever", 1);
    put("repcount", 0);
    put("if", 2);
    put("ifelse", 3);
    put("test", 1);
    put("iftrue", 1);
    put("ift", 1);
    put("iffalse", 1);
    put("iff", 1);
    put("stop", 0);
    put("output", 1);
    put("op", 1);
    put("catch", 2);
    put("throw", 1);
    put("error", 0);
    put("pause", 0);
    put("continue", 1);
    put("co", 1);
    put("wait", 1);
    put("bye", 0);
    put(".maybeoutput", 1);
    put("goto", 1);
    put("tag", 1);
    put("ignore", 1);
    put("`", 1);
    put("for", 2);
    put("do.while", 2);
    put("while", 2);
    put("do.until", 2);
    put("until", 2);
    put("case", 2);
    put("cond", 1);
    put("apply", 2);
    put("invoke", 2);
    put("foreach", 2);
    put("map", 2);
    put("map.se", 2);
    put("filter", 2);
    put("find", 2);
    put("reduce", 2);
    put("crossmap", 2);
    put("cascade", 3);
    put("cascade.2", 5);
    put("transfer", 3);
    put(".defmacro", 2);
    put("macrop", 1);
    put("macro?", 1);
    put("macroexpand", 1);
  }};

  // A flag keeping track if the parser already looked ahead to resolve user
  // defined procedures that will be stored in the 'procedures' map.
  private boolean discoveredAllProcedures = false;

  /**
   * Creates a new instance of a {@code UCBLogoParser} where
   * any user defined procedures will be resolved in an intial
   * parse.
   *
   * @param source
   *         the UCB Logo source to parse.
   */
  public UCBLogoParser(String source) {
    this(new ANTLRInputStream(source));
  }

  /**
   * Creates a new instance of a {@code UCBLogoParser} where
   * any user defined procedures will be resolved in an initial
   * parse.
   *
   * @param input
   *         the inout stream containing the UCB Logo source
   *         to parse.
   */
  public UCBLogoParser(ANTLRInputStream input) {

    this(new CommonTokenStream(new UCBLogoLexer(input)));

    // Create a lexer and parser that will resolve user defined procedures.
    UCBLogoLexer lexer = new UCBLogoLexer(input);
    UCBLogoParser parser = new UCBLogoParser(new CommonTokenStream(lexer));

    ParseTreeWalker.DEFAULT.walk(new UCBLogoBaseListener(){

      @Override
      public void enterProcedure_def(@NotNull UCBLogoParser.Procedure_defContext ctx) {
        // Yes, we found a procedure: save it in the procedures-map.
        procedures.put(ctx.NAME().getText(), ctx.variables.amount);
      }

      @Override
      public void enterMacro_def(@NotNull UCBLogoParser.Macro_defContext ctx) {
        // Yes, we found a macro: save it in the procedures-map.
        procedures.put(ctx.NAME().getText(), ctx.variables.amount);
      }
    }, parser.parse());

    // Reset the input stream after having resolved the user defined procedures.
    input.reset();

    this.discoveredAllProcedures = true;
  }

  /**
   * Returns the amount of parameters the procedure expects. Note
   * that this method will only be called after {@link #procedureNameAhead()}
   * already returned {@code true}.
   *
   * @param procedureName
   *         the name of the procedure.
   *
   * @return the amount of parameters the procedure expects.
   */
  private int amountParams(String procedureName) {
    return procedures.get(procedureName.toLowerCase());
  }

  /**
   * Returns {@code true} iff the next token in the stream is of type
   * {@code NAME} and contains the text defined in {@code name}.
   *
   * @param name
   *         the text the next token should contain.
   *
   * @returns {@code true} iff the next token in the stream is of type
   * {@code NAME} and contains the text defined in {@code name}.
   */
  private boolean nameAhead(String name) {
    Token token = _input.LT(1);
    return token.getType() == NAME && token.getText().equalsIgnoreCase(name);
  }

  /**
   * Returns {@code true} iff the next token in the inout stream is of
   * type {@code NAME} and is present in the {@code procedures}.
   *
   * @returns {@code true} iff the next token in the inout stream is of
   * type {@code NAME} and is present in the {@code procedures}.
   */
  private boolean procedureNameAhead() {
    Token token = _input.LT(1);
    return token.getType() == NAME && procedures.containsKey(token.getText().toLowerCase());
  }



// UCBLogoParser tokens.
const (
	UCBLogoParserEOF = antlr.TokenEOF
	UCBLogoParserT__0 = 1
	UCBLogoParserT__1 = 2
	UCBLogoParserTO = 3
	UCBLogoParserEND = 4
	UCBLogoParserMACRO = 5
	UCBLogoParserWORD = 6
	UCBLogoParserSKIP_ = 7
	UCBLogoParserOPEN_ARRAY = 8
	UCBLogoParserCLOSE_ARRAY = 9
	UCBLogoParserOPEN_LIST = 10
	UCBLogoParserCLOSE_LIST = 11
	UCBLogoParserMINUS = 12
	UCBLogoParserPLUS = 13
	UCBLogoParserMULT = 14
	UCBLogoParserDIV = 15
	UCBLogoParserLT = 16
	UCBLogoParserGT = 17
	UCBLogoParserEQ = 18
	UCBLogoParserLT_EQ = 19
	UCBLogoParserGT_EQ = 20
	UCBLogoParserNOT_EQ = 21
	UCBLogoParserQUOTED_WORD = 22
	UCBLogoParserNUMBER = 23
	UCBLogoParserVARIABLE = 24
	UCBLogoParserNAME = 25
	UCBLogoParserANY = 26
)

// UCBLogoParser rules.
const (
	UCBLogoParserRULE_parse = 0
	UCBLogoParserRULE_instruction = 1
	UCBLogoParserRULE_procedure_def = 2
	UCBLogoParserRULE_macro_def = 3
	UCBLogoParserRULE_variables = 4
	UCBLogoParserRULE_body_def = 5
	UCBLogoParserRULE_body_instruction = 6
	UCBLogoParserRULE_procedure_call_extra_input = 7
	UCBLogoParserRULE_procedure_call = 8
	UCBLogoParserRULE_expressions = 9
	UCBLogoParserRULE_expression = 10
	UCBLogoParserRULE_array = 11
	UCBLogoParserRULE_list = 12
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_parse
	return p
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(UCBLogoParserEOF, 0)
}

func (s *ParseContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *ParseContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitParse(s)
	}
}

func (s *ParseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitParse(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *UCBLogoParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, UCBLogoParserRULE_parse)

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
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(26)
				p.Instruction()
			}


		}
		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}
	{
		p.SetState(32)
		p.Match(UCBLogoParserEOF)
	}



	return localctx
}


// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) CopyFrom(ctx *InstructionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}




type ProcedureCallInstructionContext struct {
	*InstructionContext
}

func NewProcedureCallInstructionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ProcedureCallInstructionContext {
	var p = new(ProcedureCallInstructionContext)

	p.InstructionContext = NewEmptyInstructionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InstructionContext))

	return p
}

func (s *ProcedureCallInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureCallInstructionContext) Procedure_call() IProcedure_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedure_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcedure_callContext)
}


func (s *ProcedureCallInstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterProcedureCallInstruction(s)
	}
}

func (s *ProcedureCallInstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitProcedureCallInstruction(s)
	}
}

func (s *ProcedureCallInstructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitProcedureCallInstruction(s)

	default:
		return t.VisitChildren(s)
	}
}


type MacroDefInstructionContext struct {
	*InstructionContext
}

func NewMacroDefInstructionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MacroDefInstructionContext {
	var p = new(MacroDefInstructionContext)

	p.InstructionContext = NewEmptyInstructionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InstructionContext))

	return p
}

func (s *MacroDefInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroDefInstructionContext) Macro_def() IMacro_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMacro_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMacro_defContext)
}


func (s *MacroDefInstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterMacroDefInstruction(s)
	}
}

func (s *MacroDefInstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitMacroDefInstruction(s)
	}
}

func (s *MacroDefInstructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitMacroDefInstruction(s)

	default:
		return t.VisitChildren(s)
	}
}


type ProcedureDefInstructionContext struct {
	*InstructionContext
}

func NewProcedureDefInstructionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ProcedureDefInstructionContext {
	var p = new(ProcedureDefInstructionContext)

	p.InstructionContext = NewEmptyInstructionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InstructionContext))

	return p
}

func (s *ProcedureDefInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureDefInstructionContext) Procedure_def() IProcedure_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedure_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcedure_defContext)
}


func (s *ProcedureDefInstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterProcedureDefInstruction(s)
	}
}

func (s *ProcedureDefInstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitProcedureDefInstruction(s)
	}
}

func (s *ProcedureDefInstructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitProcedureDefInstruction(s)

	default:
		return t.VisitChildren(s)
	}
}


type ProcedureCallExtraInputInstructionContext struct {
	*InstructionContext
}

func NewProcedureCallExtraInputInstructionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ProcedureCallExtraInputInstructionContext {
	var p = new(ProcedureCallExtraInputInstructionContext)

	p.InstructionContext = NewEmptyInstructionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InstructionContext))

	return p
}

func (s *ProcedureCallExtraInputInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureCallExtraInputInstructionContext) Procedure_call_extra_input() IProcedure_call_extra_inputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedure_call_extra_inputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcedure_call_extra_inputContext)
}


func (s *ProcedureCallExtraInputInstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterProcedureCallExtraInputInstruction(s)
	}
}

func (s *ProcedureCallExtraInputInstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitProcedureCallExtraInputInstruction(s)
	}
}

func (s *ProcedureCallExtraInputInstructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitProcedureCallExtraInputInstruction(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *UCBLogoParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, UCBLogoParserRULE_instruction)

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

	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewProcedureDefInstructionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Procedure_def()
		}


	case 2:
		localctx = NewMacroDefInstructionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Macro_def()
		}


	case 3:
		localctx = NewProcedureCallExtraInputInstructionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Procedure_call_extra_input()
		}


	case 4:
		localctx = NewProcedureCallInstructionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(37)
			p.Procedure_call()
		}

	}


	return localctx
}


// IProcedure_defContext is an interface to support dynamic dispatch.
type IProcedure_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NAME returns the _NAME token.
	Get_NAME() antlr.Token 


	// Set_NAME sets the _NAME token.
	Set_NAME(antlr.Token) 


	// Get_variables returns the _variables rule contexts.
	Get_variables() IVariablesContext


	// Set_variables sets the _variables rule contexts.
	Set_variables(IVariablesContext)


	// IsProcedure_defContext differentiates from other interfaces.
	IsProcedure_defContext()
}

type Procedure_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_NAME antlr.Token
	_variables IVariablesContext 
}

func NewEmptyProcedure_defContext() *Procedure_defContext {
	var p = new(Procedure_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_procedure_def
	return p
}

func (*Procedure_defContext) IsProcedure_defContext() {}

func NewProcedure_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Procedure_defContext {
	var p = new(Procedure_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_procedure_def

	return p
}

func (s *Procedure_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Procedure_defContext) Get_NAME() antlr.Token { return s._NAME }


func (s *Procedure_defContext) Set_NAME(v antlr.Token) { s._NAME = v }


func (s *Procedure_defContext) Get_variables() IVariablesContext { return s._variables }


func (s *Procedure_defContext) Set_variables(v IVariablesContext) { s._variables = v }


func (s *Procedure_defContext) TO() antlr.TerminalNode {
	return s.GetToken(UCBLogoParserTO, 0)
}

func (s *Procedure_defContext) NAME() antlr.TerminalNode {
	return s.GetToken(UCBLogoParserNAME, 0)
}

func (s *Procedure_defContext) Variables() IVariablesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariablesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariablesContext)
}

func (s *Procedure_defContext) Body_def() IBody_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBody_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBody_defContext)
}

func (s *Procedure_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Procedure_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Procedure_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterProcedure_def(s)
	}
}

func (s *Procedure_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitProcedure_def(s)
	}
}

func (s *Procedure_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitProcedure_def(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *UCBLogoParser) Procedure_def() (localctx IProcedure_defContext) {
	localctx = NewProcedure_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, UCBLogoParserRULE_procedure_def)

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
		p.SetState(40)
		p.Match(UCBLogoParserTO)
	}
	{
		p.SetState(41)

		var _m = p.Match(UCBLogoParserNAME)

		localctx.(*Procedure_defContext)._NAME = _m
	}
	{
		p.SetState(42)

		var _x = p.Variables()


		localctx.(*Procedure_defContext)._variables = _x
	}
	{
		p.SetState(43)
		p.Body_def()
	}

	     procedures.put(localctx.(*Procedure_defContext).Get_NAME().getText(), localctx.(*Procedure_defContext).Get_variables().GetAmount());
	   



	return localctx
}


// IMacro_defContext is an interface to support dynamic dispatch.
type IMacro_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NAME returns the _NAME token.
	Get_NAME() antlr.Token 


	// Set_NAME sets the _NAME token.
	Set_NAME(antlr.Token) 


	// Get_variables returns the _variables rule contexts.
	Get_variables() IVariablesContext


	// Set_variables sets the _variables rule contexts.
	Set_variables(IVariablesContext)


	// IsMacro_defContext differentiates from other interfaces.
	IsMacro_defContext()
}

type Macro_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_NAME antlr.Token
	_variables IVariablesContext 
}

func NewEmptyMacro_defContext() *Macro_defContext {
	var p = new(Macro_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_macro_def
	return p
}

func (*Macro_defContext) IsMacro_defContext() {}

func NewMacro_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Macro_defContext {
	var p = new(Macro_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_macro_def

	return p
}

func (s *Macro_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Macro_defContext) Get_NAME() antlr.Token { return s._NAME }


func (s *Macro_defContext) Set_NAME(v antlr.Token) { s._NAME = v }


func (s *Macro_defContext) Get_variables() IVariablesContext { return s._variables }


func (s *Macro_defContext) Set_variables(v IVariablesContext) { s._variables = v }


func (s *Macro_defContext) MACRO() antlr.TerminalNode {
	return s.GetToken(UCBLogoParserMACRO, 0)
}

func (s *Macro_defContext) NAME() antlr.TerminalNode {
	return s.GetToken(UCBLogoParserNAME, 0)
}

func (s *Macro_defContext) Variables() IVariablesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariablesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariablesContext)
}

func (s *Macro_defContext) Body_def() IBody_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBody_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBody_defContext)
}

func (s *Macro_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Macro_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Macro_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterMacro_def(s)
	}
}

func (s *Macro_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitMacro_def(s)
	}
}

func (s *Macro_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitMacro_def(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *UCBLogoParser) Macro_def() (localctx IMacro_defContext) {
	localctx = NewMacro_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, UCBLogoParserRULE_macro_def)

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
		p.SetState(46)
		p.Match(UCBLogoParserMACRO)
	}
	{
		p.SetState(47)

		var _m = p.Match(UCBLogoParserNAME)

		localctx.(*Macro_defContext)._NAME = _m
	}
	{
		p.SetState(48)

		var _x = p.Variables()


		localctx.(*Macro_defContext)._variables = _x
	}
	{
		p.SetState(49)
		p.Body_def()
	}

	     procedures.put(localctx.(*Macro_defContext).Get_NAME().getText(), localctx.(*Macro_defContext).Get_variables().GetAmount());
	   



	return localctx
}


// IVariablesContext is an interface to support dynamic dispatch.
type IVariablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAmount returns the amount attribute.
	GetAmount() int


	// SetAmount sets the amount attribute.
	SetAmount(int)


	// IsVariablesContext differentiates from other interfaces.
	IsVariablesContext()
}

type VariablesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	amount int
}

func NewEmptyVariablesContext() *VariablesContext {
	var p = new(VariablesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_variables
	return p
}

func (*VariablesContext) IsVariablesContext() {}

func NewVariablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablesContext {
	var p = new(VariablesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_variables

	return p
}

func (s *VariablesContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablesContext) GetAmount() int { return s.amount }


func (s *VariablesContext) SetAmount(v int) { s.amount = v }


func (s *VariablesContext) AllVARIABLE() []antlr.TerminalNode {
	return s.GetTokens(UCBLogoParserVARIABLE)
}

func (s *VariablesContext) VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(UCBLogoParserVARIABLE, i)
}

func (s *VariablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *VariablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterVariables(s)
	}
}

func (s *VariablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitVariables(s)
	}
}

func (s *VariablesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitVariables(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *UCBLogoParser) Variables() (localctx IVariablesContext) {
	localctx = NewVariablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, UCBLogoParserRULE_variables)

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
	localctx.(*VariablesContext).SetAmount( 0)
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(53)
				p.Match(UCBLogoParserVARIABLE)
			}
			localctx.(*VariablesContext).amount++;


		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}



	return localctx
}


// IBody_defContext is an interface to support dynamic dispatch.
type IBody_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBody_defContext differentiates from other interfaces.
	IsBody_defContext()
}

type Body_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBody_defContext() *Body_defContext {
	var p = new(Body_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_body_def
	return p
}

func (*Body_defContext) IsBody_defContext() {}

func NewBody_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Body_defContext {
	var p = new(Body_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_body_def

	return p
}

func (s *Body_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Body_defContext) AllEND() []antlr.TerminalNode {
	return s.GetTokens(UCBLogoParserEND)
}

func (s *Body_defContext) END(i int) antlr.TerminalNode {
	return s.GetToken(UCBLogoParserEND, i)
}

func (s *Body_defContext) AllBody_instruction() []IBody_instructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBody_instructionContext)(nil)).Elem())
	var tst = make([]IBody_instructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBody_instructionContext)
		}
	}

	return tst
}

func (s *Body_defContext) Body_instruction(i int) IBody_instructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBody_instructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBody_instructionContext)
}

func (s *Body_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Body_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Body_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterBody_def(s)
	}
}

func (s *Body_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitBody_def(s)
	}
}

func (s *Body_defContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitBody_def(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *UCBLogoParser) Body_def() (localctx IBody_defContext) {
	localctx = NewBody_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, UCBLogoParserRULE_body_def)
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

	var _alt int

	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(60)

		if !(discoveredAllProcedures) {
			panic(antlr.NewFailedPredicateException(p, "discoveredAllProcedures", ""))
		}
		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(61)
					p.Body_instruction()
				}


			}
			p.SetState(66)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}
		{
			p.SetState(67)
			p.Match(UCBLogoParserEND)
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << UCBLogoParserT__0) | (1 << UCBLogoParserT__1) | (1 << UCBLogoParserTO) | (1 << UCBLogoParserMACRO) | (1 << UCBLogoParserWORD) | (1 << UCBLogoParserSKIP_) | (1 << UCBLogoParserOPEN_ARRAY) | (1 << UCBLogoParserCLOSE_ARRAY) | (1 << UCBLogoParserOPEN_LIST) | (1 << UCBLogoParserCLOSE_LIST) | (1 << UCBLogoParserMINUS) | (1 << UCBLogoParserPLUS) | (1 << UCBLogoParserMULT) | (1 << UCBLogoParserDIV) | (1 << UCBLogoParserLT) | (1 << UCBLogoParserGT) | (1 << UCBLogoParserEQ) | (1 << UCBLogoParserLT_EQ) | (1 << UCBLogoParserGT_EQ) | (1 << UCBLogoParserNOT_EQ) | (1 << UCBLogoParserQUOTED_WORD) | (1 << UCBLogoParserNUMBER) | (1 << UCBLogoParserVARIABLE) | (1 << UCBLogoParserNAME) | (1 << UCBLogoParserANY))) != 0) {
			p.SetState(68)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == UCBLogoParserEND  {
				p.GetErrorHandler().RecoverInline(p)
			} else {
			    p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}


			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(74)
			p.Match(UCBLogoParserEND)
		}

	}


	return localctx
}


// IBody_instructionContext is an interface to support dynamic dispatch.
type IBody_instructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBody_instructionContext differentiates from other interfaces.
	IsBody_instructionContext()
}

type Body_instructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBody_instructionContext() *Body_instructionContext {
	var p = new(Body_instructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_body_instruction
	return p
}

func (*Body_instructionContext) IsBody_instructionContext() {}

func NewBody_instructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Body_instructionContext {
	var p = new(Body_instructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_body_instruction

	return p
}

func (s *Body_instructionContext) GetParser() antlr.Parser { return s.parser }

func (s *Body_instructionContext) Procedure_call_extra_input() IProcedure_call_extra_inputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedure_call_extra_inputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcedure_call_extra_inputContext)
}

func (s *Body_instructionContext) Procedure_call() IProcedure_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedure_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcedure_callContext)
}

func (s *Body_instructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Body_instructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Body_instructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterBody_instruction(s)
	}
}

func (s *Body_instructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitBody_instruction(s)
	}
}

func (s *Body_instructionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitBody_instruction(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *UCBLogoParser) Body_instruction() (localctx IBody_instructionContext) {
	localctx = NewBody_instructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, UCBLogoParserRULE_body_instruction)

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

	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Procedure_call_extra_input()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(78)
			p.Procedure_call()
		}

	}


	return localctx
}


// IProcedure_call_extra_inputContext is an interface to support dynamic dispatch.
type IProcedure_call_extra_inputContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProcedure_call_extra_inputContext differentiates from other interfaces.
	IsProcedure_call_extra_inputContext()
}

type Procedure_call_extra_inputContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProcedure_call_extra_inputContext() *Procedure_call_extra_inputContext {
	var p = new(Procedure_call_extra_inputContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_procedure_call_extra_input
	return p
}

func (*Procedure_call_extra_inputContext) IsProcedure_call_extra_inputContext() {}

func NewProcedure_call_extra_inputContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Procedure_call_extra_inputContext {
	var p = new(Procedure_call_extra_inputContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_procedure_call_extra_input

	return p
}

func (s *Procedure_call_extra_inputContext) GetParser() antlr.Parser { return s.parser }

func (s *Procedure_call_extra_inputContext) NAME() antlr.TerminalNode {
	return s.GetToken(UCBLogoParserNAME, 0)
}

func (s *Procedure_call_extra_inputContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Procedure_call_extra_inputContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Procedure_call_extra_inputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Procedure_call_extra_inputContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Procedure_call_extra_inputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterProcedure_call_extra_input(s)
	}
}

func (s *Procedure_call_extra_inputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitProcedure_call_extra_input(s)
	}
}

func (s *Procedure_call_extra_inputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitProcedure_call_extra_input(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *UCBLogoParser) Procedure_call_extra_input() (localctx IProcedure_call_extra_inputContext) {
	localctx = NewProcedure_call_extra_inputContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, UCBLogoParserRULE_procedure_call_extra_input)

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
		p.SetState(81)
		p.Match(UCBLogoParserT__0)
	}
	p.SetState(82)

	if !(procedureNameAhead()) {
		panic(antlr.NewFailedPredicateException(p, "procedureNameAhead()", ""))
	}
	{
		p.SetState(83)
		p.Match(UCBLogoParserNAME)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(84)
				p.expression(0)
			}


		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	{
		p.SetState(90)
		p.Match(UCBLogoParserT__1)
	}



	return localctx
}


// IProcedure_callContext is an interface to support dynamic dispatch.
type IProcedure_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NAME returns the _NAME token.
	Get_NAME() antlr.Token 


	// Set_NAME sets the _NAME token.
	Set_NAME(antlr.Token) 


	// IsProcedure_callContext differentiates from other interfaces.
	IsProcedure_callContext()
}

type Procedure_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_NAME antlr.Token
}

func NewEmptyProcedure_callContext() *Procedure_callContext {
	var p = new(Procedure_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_procedure_call
	return p
}

func (*Procedure_callContext) IsProcedure_callContext() {}

func NewProcedure_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Procedure_callContext {
	var p = new(Procedure_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_procedure_call

	return p
}

func (s *Procedure_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Procedure_callContext) Get_NAME() antlr.Token { return s._NAME }


func (s *Procedure_callContext) Set_NAME(v antlr.Token) { s._NAME = v }


func (s *Procedure_callContext) NAME() antlr.TerminalNode {
	return s.GetToken(UCBLogoParserNAME, 0)
}

func (s *Procedure_callContext) Expressions() IExpressionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionsContext)
}

func (s *Procedure_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Procedure_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *Procedure_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterProcedure_call(s)
	}
}

func (s *Procedure_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitProcedure_call(s)
	}
}

func (s *Procedure_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitProcedure_call(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *UCBLogoParser) Procedure_call() (localctx IProcedure_callContext) {
	localctx = NewProcedure_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, UCBLogoParserRULE_procedure_call)

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
	p.SetState(92)

	if !(procedureNameAhead()) {
		panic(antlr.NewFailedPredicateException(p, "procedureNameAhead()", ""))
	}
	{
		p.SetState(93)

		var _m = p.Match(UCBLogoParserNAME)

		localctx.(*Procedure_callContext)._NAME = _m
	}
	{
		p.SetState(94)
		p.Expressions(localctx.(*Procedure_callContext).Get_NAME().getText(), amountParams(localctx.(*Procedure_callContext).Get_NAME().getText()))
	}



	return localctx
}


// IExpressionsContext is an interface to support dynamic dispatch.
type IExpressionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name attribute.
	GetName() String

	// GetTotal returns the total attribute.
	GetTotal() int

	// GetN returns the n attribute.
	GetN() int


	// SetName sets the name attribute.
	SetName(String)

	// SetTotal sets the total attribute.
	SetTotal(int)

	// SetN sets the n attribute.
	SetN(int)


	// IsExpressionsContext differentiates from other interfaces.
	IsExpressionsContext()
}

type ExpressionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name String
	total int
	n int// TODO = 0
}

func NewEmptyExpressionsContext() *ExpressionsContext {
	var p = new(ExpressionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_expressions
	return p
}

func (*ExpressionsContext) IsExpressionsContext() {}

func NewExpressionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, name String, total int) *ExpressionsContext {
	var p = new(ExpressionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_expressions

	p.name = name
	p.total = total

	return p
}

func (s *ExpressionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionsContext) GetName() String { return s.name }

func (s *ExpressionsContext) GetTotal() int { return s.total }

func (s *ExpressionsContext) GetN() int { return s.n }


func (s *ExpressionsContext) SetName(v String) { s.name = v }

func (s *ExpressionsContext) SetTotal(v int) { s.total = v }

func (s *ExpressionsContext) SetN(v int) { s.n = v }


func (s *ExpressionsContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionsContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterExpressions(s)
	}
}

func (s *ExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitExpressions(s)
	}
}

func (s *ExpressionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitExpressions(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *UCBLogoParser) Expressions(name String, total int) (localctx IExpressionsContext) {
	localctx = NewExpressionsContext(p, p.GetParserRuleContext(), p.GetState(), name, total)
	p.EnterRule(localctx, 18, UCBLogoParserRULE_expressions)

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
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(96)

			if !(localctx.(*ExpressionsContext).n < localctx.(*ExpressionsContext).total) {
				panic(antlr.NewFailedPredicateException(p, "$n < $total", ""))
			}
			{
				p.SetState(97)
				p.expression(0)
			}
			localctx.(*ExpressionsContext).n++;


		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	     // Make sure there are enough inputs parsed for 'name'.
	     if (localctx.(*ExpressionsContext).total > localctx.(*ExpressionsContext).n) {
	       throw new RuntimeException("not enough inputs to " + name);
	     }
	   



	return localctx
}


// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}





type NotEqualsExpressionExpressionContext struct {
	*ExpressionContext
}

func NewNotEqualsExpressionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotEqualsExpressionExpressionContext {
	var p = new(NotEqualsExpressionExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotEqualsExpressionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualsExpressionExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *NotEqualsExpressionExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *NotEqualsExpressionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterNotEqualsExpressionExpression(s)
	}
}

func (s *NotEqualsExpressionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitNotEqualsExpressionExpression(s)
	}
}

func (s *NotEqualsExpressionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitNotEqualsExpressionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ArrayExpressionContext struct {
	*ExpressionContext
}

func NewArrayExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayExpressionContext {
	var p = new(ArrayExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ArrayExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayExpressionContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}


func (s *ArrayExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterArrayExpression(s)
	}
}

func (s *ArrayExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitArrayExpression(s)
	}
}

func (s *ArrayExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitArrayExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type AdditionExpressionContext struct {
	*ExpressionContext
}

func NewAdditionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditionExpressionContext {
	var p = new(AdditionExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AdditionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditionExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AdditionExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *AdditionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterAdditionExpression(s)
	}
}

func (s *AdditionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitAdditionExpression(s)
	}
}

func (s *AdditionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitAdditionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type WordExpressionContext struct {
	*ExpressionContext
}

func NewWordExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WordExpressionContext {
	var p = new(WordExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *WordExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordExpressionContext) WORD() antlr.TerminalNode {
	return s.GetToken(UCBLogoParserWORD, 0)
}


func (s *WordExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterWordExpression(s)
	}
}

func (s *WordExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitWordExpression(s)
	}
}

func (s *WordExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitWordExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type NumberExpressionContext struct {
	*ExpressionContext
}

func NewNumberExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberExpressionContext {
	var p = new(NumberExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NumberExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(UCBLogoParserNUMBER, 0)
}


func (s *NumberExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterNumberExpression(s)
	}
}

func (s *NumberExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitNumberExpression(s)
	}
}

func (s *NumberExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitNumberExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ParensExpressionContext struct {
	*ExpressionContext
}

func NewParensExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensExpressionContext {
	var p = new(ParensExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParensExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *ParensExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterParensExpression(s)
	}
}

func (s *ParensExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitParensExpression(s)
	}
}

func (s *ParensExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitParensExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type MultiplyExpressionContext struct {
	*ExpressionContext
}

func NewMultiplyExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplyExpressionContext {
	var p = new(MultiplyExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MultiplyExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MultiplyExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *MultiplyExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterMultiplyExpression(s)
	}
}

func (s *MultiplyExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitMultiplyExpression(s)
	}
}

func (s *MultiplyExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitMultiplyExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type GreaterThanExpressionContext struct {
	*ExpressionContext
}

func NewGreaterThanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanExpressionContext {
	var p = new(GreaterThanExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *GreaterThanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *GreaterThanExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *GreaterThanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterGreaterThanExpression(s)
	}
}

func (s *GreaterThanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitGreaterThanExpression(s)
	}
}

func (s *GreaterThanExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitGreaterThanExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type VariableExpressionContext struct {
	*ExpressionContext
}

func NewVariableExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableExpressionContext {
	var p = new(VariableExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *VariableExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableExpressionContext) VARIABLE() antlr.TerminalNode {
	return s.GetToken(UCBLogoParserVARIABLE, 0)
}


func (s *VariableExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterVariableExpression(s)
	}
}

func (s *VariableExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitVariableExpression(s)
	}
}

func (s *VariableExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitVariableExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type DivideExpressionContext struct {
	*ExpressionContext
}

func NewDivideExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivideExpressionContext {
	var p = new(DivideExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *DivideExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivideExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DivideExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *DivideExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterDivideExpression(s)
	}
}

func (s *DivideExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitDivideExpression(s)
	}
}

func (s *DivideExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitDivideExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type LessThanEqualsExpressionContext struct {
	*ExpressionContext
}

func NewLessThanEqualsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanEqualsExpressionContext {
	var p = new(LessThanEqualsExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LessThanEqualsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanEqualsExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LessThanEqualsExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *LessThanEqualsExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterLessThanEqualsExpression(s)
	}
}

func (s *LessThanEqualsExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitLessThanEqualsExpression(s)
	}
}

func (s *LessThanEqualsExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitLessThanEqualsExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type GreaterThanEqualsExpressionContext struct {
	*ExpressionContext
}

func NewGreaterThanEqualsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanEqualsExpressionContext {
	var p = new(GreaterThanEqualsExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *GreaterThanEqualsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanEqualsExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *GreaterThanEqualsExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *GreaterThanEqualsExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterGreaterThanEqualsExpression(s)
	}
}

func (s *GreaterThanEqualsExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitGreaterThanEqualsExpression(s)
	}
}

func (s *GreaterThanEqualsExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitGreaterThanEqualsExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type UnaryMinusExpressionContext struct {
	*ExpressionContext
}

func NewUnaryMinusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExpressionContext {
	var p = new(UnaryMinusExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryMinusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *UnaryMinusExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterUnaryMinusExpression(s)
	}
}

func (s *UnaryMinusExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitUnaryMinusExpression(s)
	}
}

func (s *UnaryMinusExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitUnaryMinusExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type QuotedWordExpressionContext struct {
	*ExpressionContext
}

func NewQuotedWordExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuotedWordExpressionContext {
	var p = new(QuotedWordExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *QuotedWordExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedWordExpressionContext) QUOTED_WORD() antlr.TerminalNode {
	return s.GetToken(UCBLogoParserQUOTED_WORD, 0)
}


func (s *QuotedWordExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterQuotedWordExpression(s)
	}
}

func (s *QuotedWordExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitQuotedWordExpression(s)
	}
}

func (s *QuotedWordExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitQuotedWordExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type EqualsExpressionContext struct {
	*ExpressionContext
}

func NewEqualsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualsExpressionContext {
	var p = new(EqualsExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *EqualsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualsExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *EqualsExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *EqualsExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterEqualsExpression(s)
	}
}

func (s *EqualsExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitEqualsExpression(s)
	}
}

func (s *EqualsExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitEqualsExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type SubtractionExpressionContext struct {
	*ExpressionContext
}

func NewSubtractionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubtractionExpressionContext {
	var p = new(SubtractionExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *SubtractionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtractionExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *SubtractionExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *SubtractionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterSubtractionExpression(s)
	}
}

func (s *SubtractionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitSubtractionExpression(s)
	}
}

func (s *SubtractionExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitSubtractionExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ProcedureCallExpressionContext struct {
	*ExpressionContext
}

func NewProcedureCallExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ProcedureCallExpressionContext {
	var p = new(ProcedureCallExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ProcedureCallExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureCallExpressionContext) Procedure_call() IProcedure_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedure_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcedure_callContext)
}


func (s *ProcedureCallExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterProcedureCallExpression(s)
	}
}

func (s *ProcedureCallExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitProcedureCallExpression(s)
	}
}

func (s *ProcedureCallExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitProcedureCallExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type LessThanExpressionContext struct {
	*ExpressionContext
}

func NewLessThanExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanExpressionContext {
	var p = new(LessThanExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LessThanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LessThanExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}


func (s *LessThanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterLessThanExpression(s)
	}
}

func (s *LessThanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitLessThanExpression(s)
	}
}

func (s *LessThanExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitLessThanExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type ProcedureCallExtraInputContext struct {
	*ExpressionContext
}

func NewProcedureCallExtraInputContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ProcedureCallExtraInputContext {
	var p = new(ProcedureCallExtraInputContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ProcedureCallExtraInputContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProcedureCallExtraInputContext) Procedure_call_extra_input() IProcedure_call_extra_inputContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProcedure_call_extra_inputContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProcedure_call_extra_inputContext)
}


func (s *ProcedureCallExtraInputContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterProcedureCallExtraInput(s)
	}
}

func (s *ProcedureCallExtraInputContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitProcedureCallExtraInput(s)
	}
}

func (s *ProcedureCallExtraInputContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitProcedureCallExtraInput(s)

	default:
		return t.VisitChildren(s)
	}
}


type ListExpressionContext struct {
	*ExpressionContext
}

func NewListExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListExpressionContext {
	var p = new(ListExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ListExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListExpressionContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}


func (s *ListExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterListExpression(s)
	}
}

func (s *ListExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitListExpression(s)
	}
}

func (s *ListExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitListExpression(s)

	default:
		return t.VisitChildren(s)
	}
}


type NameExpressionContext struct {
	*ExpressionContext
}

func NewNameExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameExpressionContext {
	var p = new(NameExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NameExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameExpressionContext) NAME() antlr.TerminalNode {
	return s.GetToken(UCBLogoParserNAME, 0)
}


func (s *NameExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterNameExpression(s)
	}
}

func (s *NameExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitNameExpression(s)
	}
}

func (s *NameExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitNameExpression(s)

	default:
		return t.VisitChildren(s)
	}
}



func (p *UCBLogoParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *UCBLogoParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, UCBLogoParserRULE_expression, _p)

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
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		localctx = NewUnaryMinusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(108)
			p.Match(UCBLogoParserMINUS)
		}
		{
			p.SetState(109)
			p.expression(21)
		}


	case 2:
		localctx = NewProcedureCallExtraInputContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(110)
			p.Procedure_call_extra_input()
		}


	case 3:
		localctx = NewProcedureCallExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(111)
			p.Procedure_call()
		}


	case 4:
		localctx = NewParensExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(112)
			p.Match(UCBLogoParserT__0)
		}
		{
			p.SetState(113)
			p.expression(0)
		}
		{
			p.SetState(114)
			p.Match(UCBLogoParserT__1)
		}


	case 5:
		localctx = NewArrayExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(116)
			p.Array()
		}


	case 6:
		localctx = NewListExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(117)
			p.List()
		}


	case 7:
		localctx = NewWordExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(118)
			p.Match(UCBLogoParserWORD)
		}


	case 8:
		localctx = NewQuotedWordExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(119)
			p.Match(UCBLogoParserQUOTED_WORD)
		}


	case 9:
		localctx = NewNumberExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(120)
			p.Match(UCBLogoParserNUMBER)
		}


	case 10:
		localctx = NewVariableExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(121)
			p.Match(UCBLogoParserVARIABLE)
		}


	case 11:
		localctx = NewNameExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(122)
			p.Match(UCBLogoParserNAME)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(155)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplyExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, UCBLogoParserRULE_expression)
				p.SetState(125)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(126)
					p.Match(UCBLogoParserMULT)
				}
				{
					p.SetState(127)
					p.expression(11)
				}


			case 2:
				localctx = NewDivideExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, UCBLogoParserRULE_expression)
				p.SetState(128)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(129)
					p.Match(UCBLogoParserDIV)
				}
				{
					p.SetState(130)
					p.expression(10)
				}


			case 3:
				localctx = NewAdditionExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, UCBLogoParserRULE_expression)
				p.SetState(131)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(132)
					p.Match(UCBLogoParserPLUS)
				}
				{
					p.SetState(133)
					p.expression(9)
				}


			case 4:
				localctx = NewSubtractionExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, UCBLogoParserRULE_expression)
				p.SetState(134)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(135)
					p.Match(UCBLogoParserMINUS)
				}
				{
					p.SetState(136)
					p.expression(8)
				}


			case 5:
				localctx = NewLessThanExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, UCBLogoParserRULE_expression)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(138)
					p.Match(UCBLogoParserLT)
				}
				{
					p.SetState(139)
					p.expression(7)
				}


			case 6:
				localctx = NewGreaterThanExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, UCBLogoParserRULE_expression)
				p.SetState(140)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(141)
					p.Match(UCBLogoParserGT)
				}
				{
					p.SetState(142)
					p.expression(6)
				}


			case 7:
				localctx = NewLessThanEqualsExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, UCBLogoParserRULE_expression)
				p.SetState(143)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(144)
					p.Match(UCBLogoParserLT_EQ)
				}
				{
					p.SetState(145)
					p.expression(5)
				}


			case 8:
				localctx = NewGreaterThanEqualsExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, UCBLogoParserRULE_expression)
				p.SetState(146)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(147)
					p.Match(UCBLogoParserGT_EQ)
				}
				{
					p.SetState(148)
					p.expression(4)
				}


			case 9:
				localctx = NewEqualsExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, UCBLogoParserRULE_expression)
				p.SetState(149)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(150)
					p.Match(UCBLogoParserEQ)
				}
				{
					p.SetState(151)
					p.expression(3)
				}


			case 10:
				localctx = NewNotEqualsExpressionExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, UCBLogoParserRULE_expression)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(153)
					p.Match(UCBLogoParserNOT_EQ)
				}
				{
					p.SetState(154)
					p.expression(2)
				}

			}

		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}



	return localctx
}


// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllArray() []IArrayContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArrayContext)(nil)).Elem())
	var tst = make([]IArrayContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArrayContext)
		}
	}

	return tst
}

func (s *ArrayContext) Array(i int) IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *UCBLogoParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, UCBLogoParserRULE_array)
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
		p.SetState(160)
		p.Match(UCBLogoParserOPEN_ARRAY)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << UCBLogoParserT__0) | (1 << UCBLogoParserT__1) | (1 << UCBLogoParserTO) | (1 << UCBLogoParserEND) | (1 << UCBLogoParserMACRO) | (1 << UCBLogoParserWORD) | (1 << UCBLogoParserSKIP_) | (1 << UCBLogoParserOPEN_ARRAY) | (1 << UCBLogoParserOPEN_LIST) | (1 << UCBLogoParserCLOSE_LIST) | (1 << UCBLogoParserMINUS) | (1 << UCBLogoParserPLUS) | (1 << UCBLogoParserMULT) | (1 << UCBLogoParserDIV) | (1 << UCBLogoParserLT) | (1 << UCBLogoParserGT) | (1 << UCBLogoParserEQ) | (1 << UCBLogoParserLT_EQ) | (1 << UCBLogoParserGT_EQ) | (1 << UCBLogoParserNOT_EQ) | (1 << UCBLogoParserQUOTED_WORD) | (1 << UCBLogoParserNUMBER) | (1 << UCBLogoParserVARIABLE) | (1 << UCBLogoParserNAME) | (1 << UCBLogoParserANY))) != 0) {
		p.SetState(163)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case UCBLogoParserT__0, UCBLogoParserT__1, UCBLogoParserTO, UCBLogoParserEND, UCBLogoParserMACRO, UCBLogoParserWORD, UCBLogoParserSKIP_, UCBLogoParserOPEN_LIST, UCBLogoParserCLOSE_LIST, UCBLogoParserMINUS, UCBLogoParserPLUS, UCBLogoParserMULT, UCBLogoParserDIV, UCBLogoParserLT, UCBLogoParserGT, UCBLogoParserEQ, UCBLogoParserLT_EQ, UCBLogoParserGT_EQ, UCBLogoParserNOT_EQ, UCBLogoParserQUOTED_WORD, UCBLogoParserNUMBER, UCBLogoParserVARIABLE, UCBLogoParserNAME, UCBLogoParserANY:
			p.SetState(161)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == UCBLogoParserOPEN_ARRAY || _la == UCBLogoParserCLOSE_ARRAY  {
				p.GetErrorHandler().RecoverInline(p)
			} else {
			    p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}


		case UCBLogoParserOPEN_ARRAY:
			{
				p.SetState(162)
				p.Array()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(168)
		p.Match(UCBLogoParserCLOSE_ARRAY)
	}



	return localctx
}


// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = UCBLogoParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = UCBLogoParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) AllList() []IListContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IListContext)(nil)).Elem())
	var tst = make([]IListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IListContext)
		}
	}

	return tst
}

func (s *ListContext) List(i int) IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.EnterList(s)
	}
}

func (s *ListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(UCBLogoListener); ok {
		listenerT.ExitList(s)
	}
}

func (s *ListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case UCBLogoVisitor:
		return t.VisitList(s)

	default:
		return t.VisitChildren(s)
	}
}




func (p *UCBLogoParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, UCBLogoParserRULE_list)
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
		p.SetState(170)
		p.Match(UCBLogoParserOPEN_LIST)
	}
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << UCBLogoParserT__0) | (1 << UCBLogoParserT__1) | (1 << UCBLogoParserTO) | (1 << UCBLogoParserEND) | (1 << UCBLogoParserMACRO) | (1 << UCBLogoParserWORD) | (1 << UCBLogoParserSKIP_) | (1 << UCBLogoParserOPEN_ARRAY) | (1 << UCBLogoParserCLOSE_ARRAY) | (1 << UCBLogoParserOPEN_LIST) | (1 << UCBLogoParserMINUS) | (1 << UCBLogoParserPLUS) | (1 << UCBLogoParserMULT) | (1 << UCBLogoParserDIV) | (1 << UCBLogoParserLT) | (1 << UCBLogoParserGT) | (1 << UCBLogoParserEQ) | (1 << UCBLogoParserLT_EQ) | (1 << UCBLogoParserGT_EQ) | (1 << UCBLogoParserNOT_EQ) | (1 << UCBLogoParserQUOTED_WORD) | (1 << UCBLogoParserNUMBER) | (1 << UCBLogoParserVARIABLE) | (1 << UCBLogoParserNAME) | (1 << UCBLogoParserANY))) != 0) {
		p.SetState(173)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case UCBLogoParserT__0, UCBLogoParserT__1, UCBLogoParserTO, UCBLogoParserEND, UCBLogoParserMACRO, UCBLogoParserWORD, UCBLogoParserSKIP_, UCBLogoParserOPEN_ARRAY, UCBLogoParserCLOSE_ARRAY, UCBLogoParserMINUS, UCBLogoParserPLUS, UCBLogoParserMULT, UCBLogoParserDIV, UCBLogoParserLT, UCBLogoParserGT, UCBLogoParserEQ, UCBLogoParserLT_EQ, UCBLogoParserGT_EQ, UCBLogoParserNOT_EQ, UCBLogoParserQUOTED_WORD, UCBLogoParserNUMBER, UCBLogoParserVARIABLE, UCBLogoParserNAME, UCBLogoParserANY:
			p.SetState(171)
			_la = p.GetTokenStream().LA(1)

			if _la <= 0 || _la == UCBLogoParserOPEN_LIST || _la == UCBLogoParserCLOSE_LIST  {
				p.GetErrorHandler().RecoverInline(p)
			} else {
			    p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}


		case UCBLogoParserOPEN_LIST:
			{
				p.SetState(172)
				p.List()
			}



		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(178)
		p.Match(UCBLogoParserCLOSE_LIST)
	}



	return localctx
}


func (p *UCBLogoParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
			var t *Body_defContext = nil
			if localctx != nil { t = localctx.(*Body_defContext) }
			return p.Body_def_Sempred(t, predIndex)

	case 7:
			var t *Procedure_call_extra_inputContext = nil
			if localctx != nil { t = localctx.(*Procedure_call_extra_inputContext) }
			return p.Procedure_call_extra_input_Sempred(t, predIndex)

	case 8:
			var t *Procedure_callContext = nil
			if localctx != nil { t = localctx.(*Procedure_callContext) }
			return p.Procedure_call_Sempred(t, predIndex)

	case 9:
			var t *ExpressionsContext = nil
			if localctx != nil { t = localctx.(*ExpressionsContext) }
			return p.Expressions_Sempred(t, predIndex)

	case 10:
			var t *ExpressionContext = nil
			if localctx != nil { t = localctx.(*ExpressionContext) }
			return p.Expression_Sempred(t, predIndex)


	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *UCBLogoParser) Body_def_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
			return discoveredAllProcedures

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *UCBLogoParser) Procedure_call_extra_input_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
			return procedureNameAhead()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *UCBLogoParser) Procedure_call_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
			return procedureNameAhead()

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *UCBLogoParser) Expressions_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
			return localctx.(*ExpressionsContext).n < localctx.(*ExpressionsContext).total

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *UCBLogoParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 4:
			return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
			return p.Precpred(p.GetParserRuleContext(), 9)

	case 6:
			return p.Precpred(p.GetParserRuleContext(), 8)

	case 7:
			return p.Precpred(p.GetParserRuleContext(), 7)

	case 8:
			return p.Precpred(p.GetParserRuleContext(), 6)

	case 9:
			return p.Precpred(p.GetParserRuleContext(), 5)

	case 10:
			return p.Precpred(p.GetParserRuleContext(), 4)

	case 11:
			return p.Precpred(p.GetParserRuleContext(), 3)

	case 12:
			return p.Precpred(p.GetParserRuleContext(), 2)

	case 13:
			return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

