// Code generated from bdf.g4 by ANTLR 4.9.3. DO NOT EDIT.

package bdf // bdf
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 54, 314,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 6, 3, 108, 10, 3, 13,
	3, 14, 3, 109, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 164, 10, 7, 12, 7, 14, 7,
	167, 11, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3,
	35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38,
	3, 38, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3,
	42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 7, 43,
	283, 10, 43, 12, 43, 14, 43, 286, 11, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3,
	44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 7, 48, 309, 10, 48, 12, 48, 14,
	48, 312, 11, 48, 3, 48, 2, 2, 49, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58,
	60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94,
	2, 2, 2, 312, 2, 96, 3, 2, 2, 2, 4, 99, 3, 2, 2, 2, 6, 113, 3, 2, 2, 2,
	8, 116, 3, 2, 2, 2, 10, 121, 3, 2, 2, 2, 12, 127, 3, 2, 2, 2, 14, 170,
	3, 2, 2, 2, 16, 173, 3, 2, 2, 2, 18, 176, 3, 2, 2, 2, 20, 179, 3, 2, 2,
	2, 22, 182, 3, 2, 2, 2, 24, 185, 3, 2, 2, 2, 26, 188, 3, 2, 2, 2, 28, 191,
	3, 2, 2, 2, 30, 194, 3, 2, 2, 2, 32, 197, 3, 2, 2, 2, 34, 200, 3, 2, 2,
	2, 36, 203, 3, 2, 2, 2, 38, 206, 3, 2, 2, 2, 40, 209, 3, 2, 2, 2, 42, 212,
	3, 2, 2, 2, 44, 215, 3, 2, 2, 2, 46, 218, 3, 2, 2, 2, 48, 221, 3, 2, 2,
	2, 50, 224, 3, 2, 2, 2, 52, 227, 3, 2, 2, 2, 54, 230, 3, 2, 2, 2, 56, 233,
	3, 2, 2, 2, 58, 236, 3, 2, 2, 2, 60, 239, 3, 2, 2, 2, 62, 242, 3, 2, 2,
	2, 64, 245, 3, 2, 2, 2, 66, 248, 3, 2, 2, 2, 68, 251, 3, 2, 2, 2, 70, 254,
	3, 2, 2, 2, 72, 257, 3, 2, 2, 2, 74, 260, 3, 2, 2, 2, 76, 263, 3, 2, 2,
	2, 78, 266, 3, 2, 2, 2, 80, 269, 3, 2, 2, 2, 82, 272, 3, 2, 2, 2, 84, 275,
	3, 2, 2, 2, 86, 289, 3, 2, 2, 2, 88, 292, 3, 2, 2, 2, 90, 296, 3, 2, 2,
	2, 92, 300, 3, 2, 2, 2, 94, 306, 3, 2, 2, 2, 96, 97, 5, 4, 3, 2, 97, 98,
	7, 2, 2, 3, 98, 3, 3, 2, 2, 2, 99, 100, 7, 3, 2, 2, 100, 107, 7, 52, 2,
	2, 101, 108, 5, 6, 4, 2, 102, 108, 5, 8, 5, 2, 103, 108, 5, 10, 6, 2, 104,
	108, 5, 12, 7, 2, 105, 108, 5, 82, 42, 2, 106, 108, 5, 84, 43, 2, 107,
	101, 3, 2, 2, 2, 107, 102, 3, 2, 2, 2, 107, 103, 3, 2, 2, 2, 107, 104,
	3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 106, 3, 2, 2, 2, 108, 109, 3, 2,
	2, 2, 109, 107, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2,
	111, 112, 7, 4, 2, 2, 112, 5, 3, 2, 2, 2, 113, 114, 7, 5, 2, 2, 114, 115,
	7, 52, 2, 2, 115, 7, 3, 2, 2, 2, 116, 117, 7, 6, 2, 2, 117, 118, 7, 52,
	2, 2, 118, 119, 7, 52, 2, 2, 119, 120, 7, 52, 2, 2, 120, 9, 3, 2, 2, 2,
	121, 122, 7, 7, 2, 2, 122, 123, 7, 52, 2, 2, 123, 124, 7, 52, 2, 2, 124,
	125, 7, 52, 2, 2, 125, 126, 7, 52, 2, 2, 126, 11, 3, 2, 2, 2, 127, 128,
	7, 8, 2, 2, 128, 165, 7, 52, 2, 2, 129, 164, 5, 44, 23, 2, 130, 164, 5,
	54, 28, 2, 131, 164, 5, 14, 8, 2, 132, 164, 5, 16, 9, 2, 133, 164, 5, 18,
	10, 2, 134, 164, 5, 20, 11, 2, 135, 164, 5, 22, 12, 2, 136, 164, 5, 24,
	13, 2, 137, 164, 5, 26, 14, 2, 138, 164, 5, 28, 15, 2, 139, 164, 5, 30,
	16, 2, 140, 164, 5, 32, 17, 2, 141, 164, 5, 34, 18, 2, 142, 164, 5, 36,
	19, 2, 143, 164, 5, 38, 20, 2, 144, 164, 5, 40, 21, 2, 145, 164, 5, 42,
	22, 2, 146, 164, 5, 46, 24, 2, 147, 164, 5, 48, 25, 2, 148, 164, 5, 50,
	26, 2, 149, 164, 5, 52, 27, 2, 150, 164, 5, 56, 29, 2, 151, 164, 5, 58,
	30, 2, 152, 164, 5, 60, 31, 2, 153, 164, 5, 62, 32, 2, 154, 164, 5, 64,
	33, 2, 155, 164, 5, 66, 34, 2, 156, 164, 5, 68, 35, 2, 157, 164, 5, 70,
	36, 2, 158, 164, 5, 72, 37, 2, 159, 164, 5, 74, 38, 2, 160, 164, 5, 76,
	39, 2, 161, 164, 5, 78, 40, 2, 162, 164, 5, 80, 41, 2, 163, 129, 3, 2,
	2, 2, 163, 130, 3, 2, 2, 2, 163, 131, 3, 2, 2, 2, 163, 132, 3, 2, 2, 2,
	163, 133, 3, 2, 2, 2, 163, 134, 3, 2, 2, 2, 163, 135, 3, 2, 2, 2, 163,
	136, 3, 2, 2, 2, 163, 137, 3, 2, 2, 2, 163, 138, 3, 2, 2, 2, 163, 139,
	3, 2, 2, 2, 163, 140, 3, 2, 2, 2, 163, 141, 3, 2, 2, 2, 163, 142, 3, 2,
	2, 2, 163, 143, 3, 2, 2, 2, 163, 144, 3, 2, 2, 2, 163, 145, 3, 2, 2, 2,
	163, 146, 3, 2, 2, 2, 163, 147, 3, 2, 2, 2, 163, 148, 3, 2, 2, 2, 163,
	149, 3, 2, 2, 2, 163, 150, 3, 2, 2, 2, 163, 151, 3, 2, 2, 2, 163, 152,
	3, 2, 2, 2, 163, 153, 3, 2, 2, 2, 163, 154, 3, 2, 2, 2, 163, 155, 3, 2,
	2, 2, 163, 156, 3, 2, 2, 2, 163, 157, 3, 2, 2, 2, 163, 158, 3, 2, 2, 2,
	163, 159, 3, 2, 2, 2, 163, 160, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163,
	162, 3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166,
	3, 2, 2, 2, 166, 168, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 168, 169, 7, 9,
	2, 2, 169, 13, 3, 2, 2, 2, 170, 171, 7, 10, 2, 2, 171, 172, 7, 53, 2, 2,
	172, 15, 3, 2, 2, 2, 173, 174, 7, 11, 2, 2, 174, 175, 7, 53, 2, 2, 175,
	17, 3, 2, 2, 2, 176, 177, 7, 12, 2, 2, 177, 178, 7, 53, 2, 2, 178, 19,
	3, 2, 2, 2, 179, 180, 7, 13, 2, 2, 180, 181, 7, 53, 2, 2, 181, 21, 3, 2,
	2, 2, 182, 183, 7, 14, 2, 2, 183, 184, 7, 53, 2, 2, 184, 23, 3, 2, 2, 2,
	185, 186, 7, 15, 2, 2, 186, 187, 7, 53, 2, 2, 187, 25, 3, 2, 2, 2, 188,
	189, 7, 16, 2, 2, 189, 190, 7, 52, 2, 2, 190, 27, 3, 2, 2, 2, 191, 192,
	7, 17, 2, 2, 192, 193, 7, 52, 2, 2, 193, 29, 3, 2, 2, 2, 194, 195, 7, 18,
	2, 2, 195, 196, 7, 52, 2, 2, 196, 31, 3, 2, 2, 2, 197, 198, 7, 19, 2, 2,
	198, 199, 7, 52, 2, 2, 199, 33, 3, 2, 2, 2, 200, 201, 7, 20, 2, 2, 201,
	202, 7, 53, 2, 2, 202, 35, 3, 2, 2, 2, 203, 204, 7, 21, 2, 2, 204, 205,
	7, 52, 2, 2, 205, 37, 3, 2, 2, 2, 206, 207, 7, 22, 2, 2, 207, 208, 7, 53,
	2, 2, 208, 39, 3, 2, 2, 2, 209, 210, 7, 23, 2, 2, 210, 211, 7, 53, 2, 2,
	211, 41, 3, 2, 2, 2, 212, 213, 7, 24, 2, 2, 213, 214, 7, 53, 2, 2, 214,
	43, 3, 2, 2, 2, 215, 216, 7, 25, 2, 2, 216, 217, 7, 52, 2, 2, 217, 45,
	3, 2, 2, 2, 218, 219, 7, 26, 2, 2, 219, 220, 7, 53, 2, 2, 220, 47, 3, 2,
	2, 2, 221, 222, 7, 27, 2, 2, 222, 223, 7, 53, 2, 2, 223, 49, 3, 2, 2, 2,
	224, 225, 7, 28, 2, 2, 225, 226, 7, 53, 2, 2, 226, 51, 3, 2, 2, 2, 227,
	228, 7, 29, 2, 2, 228, 229, 7, 53, 2, 2, 229, 53, 3, 2, 2, 2, 230, 231,
	7, 30, 2, 2, 231, 232, 7, 52, 2, 2, 232, 55, 3, 2, 2, 2, 233, 234, 7, 31,
	2, 2, 234, 235, 7, 53, 2, 2, 235, 57, 3, 2, 2, 2, 236, 237, 7, 32, 2, 2,
	237, 238, 7, 52, 2, 2, 238, 59, 3, 2, 2, 2, 239, 240, 7, 33, 2, 2, 240,
	241, 7, 52, 2, 2, 241, 61, 3, 2, 2, 2, 242, 243, 7, 34, 2, 2, 243, 244,
	7, 52, 2, 2, 244, 63, 3, 2, 2, 2, 245, 246, 7, 35, 2, 2, 246, 247, 7, 52,
	2, 2, 247, 65, 3, 2, 2, 2, 248, 249, 7, 36, 2, 2, 249, 250, 7, 52, 2, 2,
	250, 67, 3, 2, 2, 2, 251, 252, 7, 37, 2, 2, 252, 253, 7, 52, 2, 2, 253,
	69, 3, 2, 2, 2, 254, 255, 7, 38, 2, 2, 255, 256, 7, 52, 2, 2, 256, 71,
	3, 2, 2, 2, 257, 258, 7, 39, 2, 2, 258, 259, 7, 52, 2, 2, 259, 73, 3, 2,
	2, 2, 260, 261, 7, 40, 2, 2, 261, 262, 7, 52, 2, 2, 262, 75, 3, 2, 2, 2,
	263, 264, 7, 41, 2, 2, 264, 265, 7, 52, 2, 2, 265, 77, 3, 2, 2, 2, 266,
	267, 7, 42, 2, 2, 267, 268, 7, 52, 2, 2, 268, 79, 3, 2, 2, 2, 269, 270,
	7, 43, 2, 2, 270, 271, 7, 52, 2, 2, 271, 81, 3, 2, 2, 2, 272, 273, 7, 44,
	2, 2, 273, 274, 7, 52, 2, 2, 274, 83, 3, 2, 2, 2, 275, 276, 7, 45, 2, 2,
	276, 284, 7, 52, 2, 2, 277, 283, 5, 86, 44, 2, 278, 283, 5, 88, 45, 2,
	279, 283, 5, 90, 46, 2, 280, 283, 5, 92, 47, 2, 281, 283, 5, 94, 48, 2,
	282, 277, 3, 2, 2, 2, 282, 278, 3, 2, 2, 2, 282, 279, 3, 2, 2, 2, 282,
	280, 3, 2, 2, 2, 282, 281, 3, 2, 2, 2, 283, 286, 3, 2, 2, 2, 284, 282,
	3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 287, 3, 2, 2, 2, 286, 284, 3, 2,
	2, 2, 287, 288, 7, 46, 2, 2, 288, 85, 3, 2, 2, 2, 289, 290, 7, 47, 2, 2,
	290, 291, 7, 52, 2, 2, 291, 87, 3, 2, 2, 2, 292, 293, 7, 48, 2, 2, 293,
	294, 7, 52, 2, 2, 294, 295, 7, 52, 2, 2, 295, 89, 3, 2, 2, 2, 296, 297,
	7, 49, 2, 2, 297, 298, 7, 52, 2, 2, 298, 299, 7, 52, 2, 2, 299, 91, 3,
	2, 2, 2, 300, 301, 7, 50, 2, 2, 301, 302, 7, 52, 2, 2, 302, 303, 7, 52,
	2, 2, 303, 304, 7, 52, 2, 2, 304, 305, 7, 52, 2, 2, 305, 93, 3, 2, 2, 2,
	306, 310, 7, 51, 2, 2, 307, 309, 7, 52, 2, 2, 308, 307, 3, 2, 2, 2, 309,
	312, 3, 2, 2, 2, 310, 308, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 95, 3,
	2, 2, 2, 312, 310, 3, 2, 2, 2, 9, 107, 109, 163, 165, 282, 284, 310,
}
var literalNames = []string{
	"", "'STARTFONT'", "'ENDFONT'", "'FONT'", "'SIZE'", "'FONTBOUNDINGBOX'",
	"'STARTPROPERTIES'", "'ENDPROPERTIES'", "'FOUNDRY'", "'FAMILY_NAME'", "'WEIGHT_NAME'",
	"'SLANT'", "'SETWIDTH_NAME'", "'ADD_STYLE_NAME'", "'PIXEL_SIZE'", "'POINT_SIZE'",
	"'RESOLUTION_X'", "'RESOLUTION_Y'", "'SPACING'", "'AVERAGE_WIDTH'", "'CHARSET_REGISTRY'",
	"'CHARSET_ENCODING'", "'FONTNAME_REGISTRY'", "'FONT_ASCENT'", "'CHARSET_COLLECTIONS'",
	"'FONT_NAME'", "'FACE_NAME'", "'COPYRIGHT'", "'FONT_DESCENT'", "'FONT_VERSION'",
	"'UNDERLINE_POSITION'", "'UNDERLINE_THICKNESS'", "'X_HEIGHT'", "'CAP_HEIGHT'",
	"'RAW_ASCENT'", "'RAW_DESCENT'", "'NORM_SPACE'", "'RELATIVE_WEIGHT'", "'RELATIVE_SETWIDTH'",
	"'FIGURE_WIDTH'", "'AVG_LOWERCASE_WIDTH'", "'AVG_UPPERCASE_WIDTH'", "'CHARS'",
	"'STARTCHAR'", "'ENDCHAR'", "'ENCODING'", "'SWIDTH'", "'DWIDTH'", "'BBX'",
	"'BITMAP'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "ARG", "QUOTEDSTRING",
	"WS",
}

var ruleNames = []string{
	"font", "startfont", "fontdecl", "sizedecl", "fontboundingboxdecl", "propertiesdecl",
	"foundrydecl", "familynanmedecl", "weightnamedecl", "slantdecl", "setwidthnamedecl",
	"addstylenamedecl", "pixelsizedecl", "pointsizedecl", "resolutionxdecl",
	"resolutionydecl", "spacingdecl", "averagewidthdecl", "charsetregistrydecl",
	"charsetencoding", "fontnameregistry", "fontascentdecl", "charsetcollectionsdecl",
	"fontnamedecl", "facenamedecl", "copyrightdecl", "fontdecentdecl", "fontversiondecl",
	"underlinepositiondecl", "underlinethicknessdecl", "xheightdecl", "capheighdecl",
	"rawascentdecl", "rawdescentdecl", "normspacedecl", "relativeweightdecl",
	"relaticesetwidthdecl", "figurewidthdecl", "avglowercasewidthdecl", "avguppercasewidthdecl",
	"charsdecl", "chardecl", "encodingdecl", "swidthdecl", "dwidthdecl", "bbxdecl",
	"bitmapdecl",
}

type bdfParser struct {
	*antlr.BaseParser
}

// NewbdfParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *bdfParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewbdfParser(input antlr.TokenStream) *bdfParser {
	this := new(bdfParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "bdf.g4"

	return this
}

// bdfParser tokens.
const (
	bdfParserEOF          = antlr.TokenEOF
	bdfParserT__0         = 1
	bdfParserT__1         = 2
	bdfParserT__2         = 3
	bdfParserT__3         = 4
	bdfParserT__4         = 5
	bdfParserT__5         = 6
	bdfParserT__6         = 7
	bdfParserT__7         = 8
	bdfParserT__8         = 9
	bdfParserT__9         = 10
	bdfParserT__10        = 11
	bdfParserT__11        = 12
	bdfParserT__12        = 13
	bdfParserT__13        = 14
	bdfParserT__14        = 15
	bdfParserT__15        = 16
	bdfParserT__16        = 17
	bdfParserT__17        = 18
	bdfParserT__18        = 19
	bdfParserT__19        = 20
	bdfParserT__20        = 21
	bdfParserT__21        = 22
	bdfParserT__22        = 23
	bdfParserT__23        = 24
	bdfParserT__24        = 25
	bdfParserT__25        = 26
	bdfParserT__26        = 27
	bdfParserT__27        = 28
	bdfParserT__28        = 29
	bdfParserT__29        = 30
	bdfParserT__30        = 31
	bdfParserT__31        = 32
	bdfParserT__32        = 33
	bdfParserT__33        = 34
	bdfParserT__34        = 35
	bdfParserT__35        = 36
	bdfParserT__36        = 37
	bdfParserT__37        = 38
	bdfParserT__38        = 39
	bdfParserT__39        = 40
	bdfParserT__40        = 41
	bdfParserT__41        = 42
	bdfParserT__42        = 43
	bdfParserT__43        = 44
	bdfParserT__44        = 45
	bdfParserT__45        = 46
	bdfParserT__46        = 47
	bdfParserT__47        = 48
	bdfParserT__48        = 49
	bdfParserARG          = 50
	bdfParserQUOTEDSTRING = 51
	bdfParserWS           = 52
)

// bdfParser rules.
const (
	bdfParserRULE_font                   = 0
	bdfParserRULE_startfont              = 1
	bdfParserRULE_fontdecl               = 2
	bdfParserRULE_sizedecl               = 3
	bdfParserRULE_fontboundingboxdecl    = 4
	bdfParserRULE_propertiesdecl         = 5
	bdfParserRULE_foundrydecl            = 6
	bdfParserRULE_familynanmedecl        = 7
	bdfParserRULE_weightnamedecl         = 8
	bdfParserRULE_slantdecl              = 9
	bdfParserRULE_setwidthnamedecl       = 10
	bdfParserRULE_addstylenamedecl       = 11
	bdfParserRULE_pixelsizedecl          = 12
	bdfParserRULE_pointsizedecl          = 13
	bdfParserRULE_resolutionxdecl        = 14
	bdfParserRULE_resolutionydecl        = 15
	bdfParserRULE_spacingdecl            = 16
	bdfParserRULE_averagewidthdecl       = 17
	bdfParserRULE_charsetregistrydecl    = 18
	bdfParserRULE_charsetencoding        = 19
	bdfParserRULE_fontnameregistry       = 20
	bdfParserRULE_fontascentdecl         = 21
	bdfParserRULE_charsetcollectionsdecl = 22
	bdfParserRULE_fontnamedecl           = 23
	bdfParserRULE_facenamedecl           = 24
	bdfParserRULE_copyrightdecl          = 25
	bdfParserRULE_fontdecentdecl         = 26
	bdfParserRULE_fontversiondecl        = 27
	bdfParserRULE_underlinepositiondecl  = 28
	bdfParserRULE_underlinethicknessdecl = 29
	bdfParserRULE_xheightdecl            = 30
	bdfParserRULE_capheighdecl           = 31
	bdfParserRULE_rawascentdecl          = 32
	bdfParserRULE_rawdescentdecl         = 33
	bdfParserRULE_normspacedecl          = 34
	bdfParserRULE_relativeweightdecl     = 35
	bdfParserRULE_relaticesetwidthdecl   = 36
	bdfParserRULE_figurewidthdecl        = 37
	bdfParserRULE_avglowercasewidthdecl  = 38
	bdfParserRULE_avguppercasewidthdecl  = 39
	bdfParserRULE_charsdecl              = 40
	bdfParserRULE_chardecl               = 41
	bdfParserRULE_encodingdecl           = 42
	bdfParserRULE_swidthdecl             = 43
	bdfParserRULE_dwidthdecl             = 44
	bdfParserRULE_bbxdecl                = 45
	bdfParserRULE_bitmapdecl             = 46
)

// IFontContext is an interface to support dynamic dispatch.
type IFontContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFontContext differentiates from other interfaces.
	IsFontContext()
}

type FontContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFontContext() *FontContext {
	var p = new(FontContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_font
	return p
}

func (*FontContext) IsFontContext() {}

func NewFontContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FontContext {
	var p = new(FontContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_font

	return p
}

func (s *FontContext) GetParser() antlr.Parser { return s.parser }

func (s *FontContext) Startfont() IStartfontContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStartfontContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStartfontContext)
}

func (s *FontContext) EOF() antlr.TerminalNode {
	return s.GetToken(bdfParserEOF, 0)
}

func (s *FontContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FontContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FontContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterFont(s)
	}
}

func (s *FontContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitFont(s)
	}
}

func (p *bdfParser) Font() (localctx IFontContext) {
	this := p
	_ = this

	localctx = NewFontContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, bdfParserRULE_font)

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
		p.SetState(94)
		p.Startfont()
	}
	{
		p.SetState(95)
		p.Match(bdfParserEOF)
	}

	return localctx
}

// IStartfontContext is an interface to support dynamic dispatch.
type IStartfontContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartfontContext differentiates from other interfaces.
	IsStartfontContext()
}

type StartfontContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartfontContext() *StartfontContext {
	var p = new(StartfontContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_startfont
	return p
}

func (*StartfontContext) IsStartfontContext() {}

func NewStartfontContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartfontContext {
	var p = new(StartfontContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_startfont

	return p
}

func (s *StartfontContext) GetParser() antlr.Parser { return s.parser }

func (s *StartfontContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *StartfontContext) AllFontdecl() []IFontdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFontdeclContext)(nil)).Elem())
	var tst = make([]IFontdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFontdeclContext)
		}
	}

	return tst
}

func (s *StartfontContext) Fontdecl(i int) IFontdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFontdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFontdeclContext)
}

func (s *StartfontContext) AllSizedecl() []ISizedeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISizedeclContext)(nil)).Elem())
	var tst = make([]ISizedeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISizedeclContext)
		}
	}

	return tst
}

func (s *StartfontContext) Sizedecl(i int) ISizedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISizedeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISizedeclContext)
}

func (s *StartfontContext) AllFontboundingboxdecl() []IFontboundingboxdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFontboundingboxdeclContext)(nil)).Elem())
	var tst = make([]IFontboundingboxdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFontboundingboxdeclContext)
		}
	}

	return tst
}

func (s *StartfontContext) Fontboundingboxdecl(i int) IFontboundingboxdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFontboundingboxdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFontboundingboxdeclContext)
}

func (s *StartfontContext) AllPropertiesdecl() []IPropertiesdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertiesdeclContext)(nil)).Elem())
	var tst = make([]IPropertiesdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertiesdeclContext)
		}
	}

	return tst
}

func (s *StartfontContext) Propertiesdecl(i int) IPropertiesdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertiesdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertiesdeclContext)
}

func (s *StartfontContext) AllCharsdecl() []ICharsdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharsdeclContext)(nil)).Elem())
	var tst = make([]ICharsdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharsdeclContext)
		}
	}

	return tst
}

func (s *StartfontContext) Charsdecl(i int) ICharsdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharsdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharsdeclContext)
}

func (s *StartfontContext) AllChardecl() []IChardeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IChardeclContext)(nil)).Elem())
	var tst = make([]IChardeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IChardeclContext)
		}
	}

	return tst
}

func (s *StartfontContext) Chardecl(i int) IChardeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChardeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IChardeclContext)
}

func (s *StartfontContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartfontContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartfontContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterStartfont(s)
	}
}

func (s *StartfontContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitStartfont(s)
	}
}

func (p *bdfParser) Startfont() (localctx IStartfontContext) {
	this := p
	_ = this

	localctx = NewStartfontContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, bdfParserRULE_startfont)
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
		p.SetState(97)
		p.Match(bdfParserT__0)
	}
	{
		p.SetState(98)
		p.Match(bdfParserARG)
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<bdfParserT__2)|(1<<bdfParserT__3)|(1<<bdfParserT__4)|(1<<bdfParserT__5))) != 0) || _la == bdfParserT__41 || _la == bdfParserT__42 {
		p.SetState(105)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case bdfParserT__2:
			{
				p.SetState(99)
				p.Fontdecl()
			}

		case bdfParserT__3:
			{
				p.SetState(100)
				p.Sizedecl()
			}

		case bdfParserT__4:
			{
				p.SetState(101)
				p.Fontboundingboxdecl()
			}

		case bdfParserT__5:
			{
				p.SetState(102)
				p.Propertiesdecl()
			}

		case bdfParserT__41:
			{
				p.SetState(103)
				p.Charsdecl()
			}

		case bdfParserT__42:
			{
				p.SetState(104)
				p.Chardecl()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(109)
		p.Match(bdfParserT__1)
	}

	return localctx
}

// IFontdeclContext is an interface to support dynamic dispatch.
type IFontdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFontdeclContext differentiates from other interfaces.
	IsFontdeclContext()
}

type FontdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFontdeclContext() *FontdeclContext {
	var p = new(FontdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_fontdecl
	return p
}

func (*FontdeclContext) IsFontdeclContext() {}

func NewFontdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FontdeclContext {
	var p = new(FontdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_fontdecl

	return p
}

func (s *FontdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FontdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *FontdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FontdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FontdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterFontdecl(s)
	}
}

func (s *FontdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitFontdecl(s)
	}
}

func (p *bdfParser) Fontdecl() (localctx IFontdeclContext) {
	this := p
	_ = this

	localctx = NewFontdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, bdfParserRULE_fontdecl)

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
		p.SetState(111)
		p.Match(bdfParserT__2)
	}
	{
		p.SetState(112)
		p.Match(bdfParserARG)
	}

	return localctx
}

// ISizedeclContext is an interface to support dynamic dispatch.
type ISizedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSizedeclContext differentiates from other interfaces.
	IsSizedeclContext()
}

type SizedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySizedeclContext() *SizedeclContext {
	var p = new(SizedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_sizedecl
	return p
}

func (*SizedeclContext) IsSizedeclContext() {}

func NewSizedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SizedeclContext {
	var p = new(SizedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_sizedecl

	return p
}

func (s *SizedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SizedeclContext) AllARG() []antlr.TerminalNode {
	return s.GetTokens(bdfParserARG)
}

func (s *SizedeclContext) ARG(i int) antlr.TerminalNode {
	return s.GetToken(bdfParserARG, i)
}

func (s *SizedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SizedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterSizedecl(s)
	}
}

func (s *SizedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitSizedecl(s)
	}
}

func (p *bdfParser) Sizedecl() (localctx ISizedeclContext) {
	this := p
	_ = this

	localctx = NewSizedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, bdfParserRULE_sizedecl)

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
		p.SetState(114)
		p.Match(bdfParserT__3)
	}
	{
		p.SetState(115)
		p.Match(bdfParserARG)
	}
	{
		p.SetState(116)
		p.Match(bdfParserARG)
	}
	{
		p.SetState(117)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IFontboundingboxdeclContext is an interface to support dynamic dispatch.
type IFontboundingboxdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFontboundingboxdeclContext differentiates from other interfaces.
	IsFontboundingboxdeclContext()
}

type FontboundingboxdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFontboundingboxdeclContext() *FontboundingboxdeclContext {
	var p = new(FontboundingboxdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_fontboundingboxdecl
	return p
}

func (*FontboundingboxdeclContext) IsFontboundingboxdeclContext() {}

func NewFontboundingboxdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FontboundingboxdeclContext {
	var p = new(FontboundingboxdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_fontboundingboxdecl

	return p
}

func (s *FontboundingboxdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FontboundingboxdeclContext) AllARG() []antlr.TerminalNode {
	return s.GetTokens(bdfParserARG)
}

func (s *FontboundingboxdeclContext) ARG(i int) antlr.TerminalNode {
	return s.GetToken(bdfParserARG, i)
}

func (s *FontboundingboxdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FontboundingboxdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FontboundingboxdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterFontboundingboxdecl(s)
	}
}

func (s *FontboundingboxdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitFontboundingboxdecl(s)
	}
}

func (p *bdfParser) Fontboundingboxdecl() (localctx IFontboundingboxdeclContext) {
	this := p
	_ = this

	localctx = NewFontboundingboxdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, bdfParserRULE_fontboundingboxdecl)

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
		p.SetState(119)
		p.Match(bdfParserT__4)
	}
	{
		p.SetState(120)
		p.Match(bdfParserARG)
	}
	{
		p.SetState(121)
		p.Match(bdfParserARG)
	}
	{
		p.SetState(122)
		p.Match(bdfParserARG)
	}
	{
		p.SetState(123)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IPropertiesdeclContext is an interface to support dynamic dispatch.
type IPropertiesdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertiesdeclContext differentiates from other interfaces.
	IsPropertiesdeclContext()
}

type PropertiesdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertiesdeclContext() *PropertiesdeclContext {
	var p = new(PropertiesdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_propertiesdecl
	return p
}

func (*PropertiesdeclContext) IsPropertiesdeclContext() {}

func NewPropertiesdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertiesdeclContext {
	var p = new(PropertiesdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_propertiesdecl

	return p
}

func (s *PropertiesdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertiesdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *PropertiesdeclContext) AllFontascentdecl() []IFontascentdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFontascentdeclContext)(nil)).Elem())
	var tst = make([]IFontascentdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFontascentdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Fontascentdecl(i int) IFontascentdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFontascentdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFontascentdeclContext)
}

func (s *PropertiesdeclContext) AllFontdecentdecl() []IFontdecentdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFontdecentdeclContext)(nil)).Elem())
	var tst = make([]IFontdecentdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFontdecentdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Fontdecentdecl(i int) IFontdecentdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFontdecentdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFontdecentdeclContext)
}

func (s *PropertiesdeclContext) AllFoundrydecl() []IFoundrydeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFoundrydeclContext)(nil)).Elem())
	var tst = make([]IFoundrydeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFoundrydeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Foundrydecl(i int) IFoundrydeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFoundrydeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFoundrydeclContext)
}

func (s *PropertiesdeclContext) AllFamilynanmedecl() []IFamilynanmedeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFamilynanmedeclContext)(nil)).Elem())
	var tst = make([]IFamilynanmedeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFamilynanmedeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Familynanmedecl(i int) IFamilynanmedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFamilynanmedeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFamilynanmedeclContext)
}

func (s *PropertiesdeclContext) AllWeightnamedecl() []IWeightnamedeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWeightnamedeclContext)(nil)).Elem())
	var tst = make([]IWeightnamedeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWeightnamedeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Weightnamedecl(i int) IWeightnamedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWeightnamedeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWeightnamedeclContext)
}

func (s *PropertiesdeclContext) AllSlantdecl() []ISlantdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISlantdeclContext)(nil)).Elem())
	var tst = make([]ISlantdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISlantdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Slantdecl(i int) ISlantdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISlantdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISlantdeclContext)
}

func (s *PropertiesdeclContext) AllSetwidthnamedecl() []ISetwidthnamedeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISetwidthnamedeclContext)(nil)).Elem())
	var tst = make([]ISetwidthnamedeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISetwidthnamedeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Setwidthnamedecl(i int) ISetwidthnamedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetwidthnamedeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISetwidthnamedeclContext)
}

func (s *PropertiesdeclContext) AllAddstylenamedecl() []IAddstylenamedeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAddstylenamedeclContext)(nil)).Elem())
	var tst = make([]IAddstylenamedeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAddstylenamedeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Addstylenamedecl(i int) IAddstylenamedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddstylenamedeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAddstylenamedeclContext)
}

func (s *PropertiesdeclContext) AllPixelsizedecl() []IPixelsizedeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPixelsizedeclContext)(nil)).Elem())
	var tst = make([]IPixelsizedeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPixelsizedeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Pixelsizedecl(i int) IPixelsizedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPixelsizedeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPixelsizedeclContext)
}

func (s *PropertiesdeclContext) AllPointsizedecl() []IPointsizedeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPointsizedeclContext)(nil)).Elem())
	var tst = make([]IPointsizedeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPointsizedeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Pointsizedecl(i int) IPointsizedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointsizedeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPointsizedeclContext)
}

func (s *PropertiesdeclContext) AllResolutionxdecl() []IResolutionxdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IResolutionxdeclContext)(nil)).Elem())
	var tst = make([]IResolutionxdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IResolutionxdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Resolutionxdecl(i int) IResolutionxdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResolutionxdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IResolutionxdeclContext)
}

func (s *PropertiesdeclContext) AllResolutionydecl() []IResolutionydeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IResolutionydeclContext)(nil)).Elem())
	var tst = make([]IResolutionydeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IResolutionydeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Resolutionydecl(i int) IResolutionydeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IResolutionydeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IResolutionydeclContext)
}

func (s *PropertiesdeclContext) AllSpacingdecl() []ISpacingdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISpacingdeclContext)(nil)).Elem())
	var tst = make([]ISpacingdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISpacingdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Spacingdecl(i int) ISpacingdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpacingdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISpacingdeclContext)
}

func (s *PropertiesdeclContext) AllAveragewidthdecl() []IAveragewidthdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAveragewidthdeclContext)(nil)).Elem())
	var tst = make([]IAveragewidthdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAveragewidthdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Averagewidthdecl(i int) IAveragewidthdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAveragewidthdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAveragewidthdeclContext)
}

func (s *PropertiesdeclContext) AllCharsetregistrydecl() []ICharsetregistrydeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharsetregistrydeclContext)(nil)).Elem())
	var tst = make([]ICharsetregistrydeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharsetregistrydeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Charsetregistrydecl(i int) ICharsetregistrydeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharsetregistrydeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharsetregistrydeclContext)
}

func (s *PropertiesdeclContext) AllCharsetencoding() []ICharsetencodingContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharsetencodingContext)(nil)).Elem())
	var tst = make([]ICharsetencodingContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharsetencodingContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Charsetencoding(i int) ICharsetencodingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharsetencodingContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharsetencodingContext)
}

func (s *PropertiesdeclContext) AllFontnameregistry() []IFontnameregistryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFontnameregistryContext)(nil)).Elem())
	var tst = make([]IFontnameregistryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFontnameregistryContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Fontnameregistry(i int) IFontnameregistryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFontnameregistryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFontnameregistryContext)
}

func (s *PropertiesdeclContext) AllCharsetcollectionsdecl() []ICharsetcollectionsdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICharsetcollectionsdeclContext)(nil)).Elem())
	var tst = make([]ICharsetcollectionsdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICharsetcollectionsdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Charsetcollectionsdecl(i int) ICharsetcollectionsdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICharsetcollectionsdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICharsetcollectionsdeclContext)
}

func (s *PropertiesdeclContext) AllFontnamedecl() []IFontnamedeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFontnamedeclContext)(nil)).Elem())
	var tst = make([]IFontnamedeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFontnamedeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Fontnamedecl(i int) IFontnamedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFontnamedeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFontnamedeclContext)
}

func (s *PropertiesdeclContext) AllFacenamedecl() []IFacenamedeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFacenamedeclContext)(nil)).Elem())
	var tst = make([]IFacenamedeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFacenamedeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Facenamedecl(i int) IFacenamedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFacenamedeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFacenamedeclContext)
}

func (s *PropertiesdeclContext) AllCopyrightdecl() []ICopyrightdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICopyrightdeclContext)(nil)).Elem())
	var tst = make([]ICopyrightdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICopyrightdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Copyrightdecl(i int) ICopyrightdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICopyrightdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICopyrightdeclContext)
}

func (s *PropertiesdeclContext) AllFontversiondecl() []IFontversiondeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFontversiondeclContext)(nil)).Elem())
	var tst = make([]IFontversiondeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFontversiondeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Fontversiondecl(i int) IFontversiondeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFontversiondeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFontversiondeclContext)
}

func (s *PropertiesdeclContext) AllUnderlinepositiondecl() []IUnderlinepositiondeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnderlinepositiondeclContext)(nil)).Elem())
	var tst = make([]IUnderlinepositiondeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnderlinepositiondeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Underlinepositiondecl(i int) IUnderlinepositiondeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnderlinepositiondeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnderlinepositiondeclContext)
}

func (s *PropertiesdeclContext) AllUnderlinethicknessdecl() []IUnderlinethicknessdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnderlinethicknessdeclContext)(nil)).Elem())
	var tst = make([]IUnderlinethicknessdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnderlinethicknessdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Underlinethicknessdecl(i int) IUnderlinethicknessdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnderlinethicknessdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnderlinethicknessdeclContext)
}

func (s *PropertiesdeclContext) AllXheightdecl() []IXheightdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IXheightdeclContext)(nil)).Elem())
	var tst = make([]IXheightdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IXheightdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Xheightdecl(i int) IXheightdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IXheightdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IXheightdeclContext)
}

func (s *PropertiesdeclContext) AllCapheighdecl() []ICapheighdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICapheighdeclContext)(nil)).Elem())
	var tst = make([]ICapheighdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICapheighdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Capheighdecl(i int) ICapheighdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICapheighdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICapheighdeclContext)
}

func (s *PropertiesdeclContext) AllRawascentdecl() []IRawascentdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRawascentdeclContext)(nil)).Elem())
	var tst = make([]IRawascentdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRawascentdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Rawascentdecl(i int) IRawascentdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRawascentdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRawascentdeclContext)
}

func (s *PropertiesdeclContext) AllRawdescentdecl() []IRawdescentdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRawdescentdeclContext)(nil)).Elem())
	var tst = make([]IRawdescentdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRawdescentdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Rawdescentdecl(i int) IRawdescentdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRawdescentdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRawdescentdeclContext)
}

func (s *PropertiesdeclContext) AllNormspacedecl() []INormspacedeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INormspacedeclContext)(nil)).Elem())
	var tst = make([]INormspacedeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INormspacedeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Normspacedecl(i int) INormspacedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INormspacedeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INormspacedeclContext)
}

func (s *PropertiesdeclContext) AllRelativeweightdecl() []IRelativeweightdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelativeweightdeclContext)(nil)).Elem())
	var tst = make([]IRelativeweightdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelativeweightdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Relativeweightdecl(i int) IRelativeweightdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelativeweightdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelativeweightdeclContext)
}

func (s *PropertiesdeclContext) AllRelaticesetwidthdecl() []IRelaticesetwidthdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelaticesetwidthdeclContext)(nil)).Elem())
	var tst = make([]IRelaticesetwidthdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelaticesetwidthdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Relaticesetwidthdecl(i int) IRelaticesetwidthdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelaticesetwidthdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelaticesetwidthdeclContext)
}

func (s *PropertiesdeclContext) AllFigurewidthdecl() []IFigurewidthdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFigurewidthdeclContext)(nil)).Elem())
	var tst = make([]IFigurewidthdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFigurewidthdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Figurewidthdecl(i int) IFigurewidthdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFigurewidthdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFigurewidthdeclContext)
}

func (s *PropertiesdeclContext) AllAvglowercasewidthdecl() []IAvglowercasewidthdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAvglowercasewidthdeclContext)(nil)).Elem())
	var tst = make([]IAvglowercasewidthdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAvglowercasewidthdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Avglowercasewidthdecl(i int) IAvglowercasewidthdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAvglowercasewidthdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAvglowercasewidthdeclContext)
}

func (s *PropertiesdeclContext) AllAvguppercasewidthdecl() []IAvguppercasewidthdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAvguppercasewidthdeclContext)(nil)).Elem())
	var tst = make([]IAvguppercasewidthdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAvguppercasewidthdeclContext)
		}
	}

	return tst
}

func (s *PropertiesdeclContext) Avguppercasewidthdecl(i int) IAvguppercasewidthdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAvguppercasewidthdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAvguppercasewidthdeclContext)
}

func (s *PropertiesdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertiesdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertiesdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterPropertiesdecl(s)
	}
}

func (s *PropertiesdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitPropertiesdecl(s)
	}
}

func (p *bdfParser) Propertiesdecl() (localctx IPropertiesdeclContext) {
	this := p
	_ = this

	localctx = NewPropertiesdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, bdfParserRULE_propertiesdecl)
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
		p.SetState(125)
		p.Match(bdfParserT__5)
	}
	{
		p.SetState(126)
		p.Match(bdfParserARG)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<bdfParserT__7)|(1<<bdfParserT__8)|(1<<bdfParserT__9)|(1<<bdfParserT__10)|(1<<bdfParserT__11)|(1<<bdfParserT__12)|(1<<bdfParserT__13)|(1<<bdfParserT__14)|(1<<bdfParserT__15)|(1<<bdfParserT__16)|(1<<bdfParserT__17)|(1<<bdfParserT__18)|(1<<bdfParserT__19)|(1<<bdfParserT__20)|(1<<bdfParserT__21)|(1<<bdfParserT__22)|(1<<bdfParserT__23)|(1<<bdfParserT__24)|(1<<bdfParserT__25)|(1<<bdfParserT__26)|(1<<bdfParserT__27)|(1<<bdfParserT__28)|(1<<bdfParserT__29)|(1<<bdfParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(bdfParserT__31-32))|(1<<(bdfParserT__32-32))|(1<<(bdfParserT__33-32))|(1<<(bdfParserT__34-32))|(1<<(bdfParserT__35-32))|(1<<(bdfParserT__36-32))|(1<<(bdfParserT__37-32))|(1<<(bdfParserT__38-32))|(1<<(bdfParserT__39-32))|(1<<(bdfParserT__40-32)))) != 0) {
		p.SetState(161)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case bdfParserT__22:
			{
				p.SetState(127)
				p.Fontascentdecl()
			}

		case bdfParserT__27:
			{
				p.SetState(128)
				p.Fontdecentdecl()
			}

		case bdfParserT__7:
			{
				p.SetState(129)
				p.Foundrydecl()
			}

		case bdfParserT__8:
			{
				p.SetState(130)
				p.Familynanmedecl()
			}

		case bdfParserT__9:
			{
				p.SetState(131)
				p.Weightnamedecl()
			}

		case bdfParserT__10:
			{
				p.SetState(132)
				p.Slantdecl()
			}

		case bdfParserT__11:
			{
				p.SetState(133)
				p.Setwidthnamedecl()
			}

		case bdfParserT__12:
			{
				p.SetState(134)
				p.Addstylenamedecl()
			}

		case bdfParserT__13:
			{
				p.SetState(135)
				p.Pixelsizedecl()
			}

		case bdfParserT__14:
			{
				p.SetState(136)
				p.Pointsizedecl()
			}

		case bdfParserT__15:
			{
				p.SetState(137)
				p.Resolutionxdecl()
			}

		case bdfParserT__16:
			{
				p.SetState(138)
				p.Resolutionydecl()
			}

		case bdfParserT__17:
			{
				p.SetState(139)
				p.Spacingdecl()
			}

		case bdfParserT__18:
			{
				p.SetState(140)
				p.Averagewidthdecl()
			}

		case bdfParserT__19:
			{
				p.SetState(141)
				p.Charsetregistrydecl()
			}

		case bdfParserT__20:
			{
				p.SetState(142)
				p.Charsetencoding()
			}

		case bdfParserT__21:
			{
				p.SetState(143)
				p.Fontnameregistry()
			}

		case bdfParserT__23:
			{
				p.SetState(144)
				p.Charsetcollectionsdecl()
			}

		case bdfParserT__24:
			{
				p.SetState(145)
				p.Fontnamedecl()
			}

		case bdfParserT__25:
			{
				p.SetState(146)
				p.Facenamedecl()
			}

		case bdfParserT__26:
			{
				p.SetState(147)
				p.Copyrightdecl()
			}

		case bdfParserT__28:
			{
				p.SetState(148)
				p.Fontversiondecl()
			}

		case bdfParserT__29:
			{
				p.SetState(149)
				p.Underlinepositiondecl()
			}

		case bdfParserT__30:
			{
				p.SetState(150)
				p.Underlinethicknessdecl()
			}

		case bdfParserT__31:
			{
				p.SetState(151)
				p.Xheightdecl()
			}

		case bdfParserT__32:
			{
				p.SetState(152)
				p.Capheighdecl()
			}

		case bdfParserT__33:
			{
				p.SetState(153)
				p.Rawascentdecl()
			}

		case bdfParserT__34:
			{
				p.SetState(154)
				p.Rawdescentdecl()
			}

		case bdfParserT__35:
			{
				p.SetState(155)
				p.Normspacedecl()
			}

		case bdfParserT__36:
			{
				p.SetState(156)
				p.Relativeweightdecl()
			}

		case bdfParserT__37:
			{
				p.SetState(157)
				p.Relaticesetwidthdecl()
			}

		case bdfParserT__38:
			{
				p.SetState(158)
				p.Figurewidthdecl()
			}

		case bdfParserT__39:
			{
				p.SetState(159)
				p.Avglowercasewidthdecl()
			}

		case bdfParserT__40:
			{
				p.SetState(160)
				p.Avguppercasewidthdecl()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(166)
		p.Match(bdfParserT__6)
	}

	return localctx
}

// IFoundrydeclContext is an interface to support dynamic dispatch.
type IFoundrydeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFoundrydeclContext differentiates from other interfaces.
	IsFoundrydeclContext()
}

type FoundrydeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFoundrydeclContext() *FoundrydeclContext {
	var p = new(FoundrydeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_foundrydecl
	return p
}

func (*FoundrydeclContext) IsFoundrydeclContext() {}

func NewFoundrydeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FoundrydeclContext {
	var p = new(FoundrydeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_foundrydecl

	return p
}

func (s *FoundrydeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FoundrydeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *FoundrydeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FoundrydeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FoundrydeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterFoundrydecl(s)
	}
}

func (s *FoundrydeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitFoundrydecl(s)
	}
}

func (p *bdfParser) Foundrydecl() (localctx IFoundrydeclContext) {
	this := p
	_ = this

	localctx = NewFoundrydeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, bdfParserRULE_foundrydecl)

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
		p.SetState(168)
		p.Match(bdfParserT__7)
	}
	{
		p.SetState(169)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// IFamilynanmedeclContext is an interface to support dynamic dispatch.
type IFamilynanmedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFamilynanmedeclContext differentiates from other interfaces.
	IsFamilynanmedeclContext()
}

type FamilynanmedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFamilynanmedeclContext() *FamilynanmedeclContext {
	var p = new(FamilynanmedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_familynanmedecl
	return p
}

func (*FamilynanmedeclContext) IsFamilynanmedeclContext() {}

func NewFamilynanmedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FamilynanmedeclContext {
	var p = new(FamilynanmedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_familynanmedecl

	return p
}

func (s *FamilynanmedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FamilynanmedeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *FamilynanmedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FamilynanmedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FamilynanmedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterFamilynanmedecl(s)
	}
}

func (s *FamilynanmedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitFamilynanmedecl(s)
	}
}

func (p *bdfParser) Familynanmedecl() (localctx IFamilynanmedeclContext) {
	this := p
	_ = this

	localctx = NewFamilynanmedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, bdfParserRULE_familynanmedecl)

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
		p.SetState(171)
		p.Match(bdfParserT__8)
	}
	{
		p.SetState(172)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// IWeightnamedeclContext is an interface to support dynamic dispatch.
type IWeightnamedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWeightnamedeclContext differentiates from other interfaces.
	IsWeightnamedeclContext()
}

type WeightnamedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWeightnamedeclContext() *WeightnamedeclContext {
	var p = new(WeightnamedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_weightnamedecl
	return p
}

func (*WeightnamedeclContext) IsWeightnamedeclContext() {}

func NewWeightnamedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WeightnamedeclContext {
	var p = new(WeightnamedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_weightnamedecl

	return p
}

func (s *WeightnamedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *WeightnamedeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *WeightnamedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WeightnamedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WeightnamedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterWeightnamedecl(s)
	}
}

func (s *WeightnamedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitWeightnamedecl(s)
	}
}

func (p *bdfParser) Weightnamedecl() (localctx IWeightnamedeclContext) {
	this := p
	_ = this

	localctx = NewWeightnamedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, bdfParserRULE_weightnamedecl)

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
		p.SetState(174)
		p.Match(bdfParserT__9)
	}
	{
		p.SetState(175)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// ISlantdeclContext is an interface to support dynamic dispatch.
type ISlantdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSlantdeclContext differentiates from other interfaces.
	IsSlantdeclContext()
}

type SlantdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySlantdeclContext() *SlantdeclContext {
	var p = new(SlantdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_slantdecl
	return p
}

func (*SlantdeclContext) IsSlantdeclContext() {}

func NewSlantdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SlantdeclContext {
	var p = new(SlantdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_slantdecl

	return p
}

func (s *SlantdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SlantdeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *SlantdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SlantdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SlantdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterSlantdecl(s)
	}
}

func (s *SlantdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitSlantdecl(s)
	}
}

func (p *bdfParser) Slantdecl() (localctx ISlantdeclContext) {
	this := p
	_ = this

	localctx = NewSlantdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, bdfParserRULE_slantdecl)

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
		p.SetState(177)
		p.Match(bdfParserT__10)
	}
	{
		p.SetState(178)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// ISetwidthnamedeclContext is an interface to support dynamic dispatch.
type ISetwidthnamedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetwidthnamedeclContext differentiates from other interfaces.
	IsSetwidthnamedeclContext()
}

type SetwidthnamedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetwidthnamedeclContext() *SetwidthnamedeclContext {
	var p = new(SetwidthnamedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_setwidthnamedecl
	return p
}

func (*SetwidthnamedeclContext) IsSetwidthnamedeclContext() {}

func NewSetwidthnamedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetwidthnamedeclContext {
	var p = new(SetwidthnamedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_setwidthnamedecl

	return p
}

func (s *SetwidthnamedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SetwidthnamedeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *SetwidthnamedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetwidthnamedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetwidthnamedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterSetwidthnamedecl(s)
	}
}

func (s *SetwidthnamedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitSetwidthnamedecl(s)
	}
}

func (p *bdfParser) Setwidthnamedecl() (localctx ISetwidthnamedeclContext) {
	this := p
	_ = this

	localctx = NewSetwidthnamedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, bdfParserRULE_setwidthnamedecl)

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
		p.SetState(180)
		p.Match(bdfParserT__11)
	}
	{
		p.SetState(181)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// IAddstylenamedeclContext is an interface to support dynamic dispatch.
type IAddstylenamedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddstylenamedeclContext differentiates from other interfaces.
	IsAddstylenamedeclContext()
}

type AddstylenamedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddstylenamedeclContext() *AddstylenamedeclContext {
	var p = new(AddstylenamedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_addstylenamedecl
	return p
}

func (*AddstylenamedeclContext) IsAddstylenamedeclContext() {}

func NewAddstylenamedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddstylenamedeclContext {
	var p = new(AddstylenamedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_addstylenamedecl

	return p
}

func (s *AddstylenamedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *AddstylenamedeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *AddstylenamedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddstylenamedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddstylenamedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterAddstylenamedecl(s)
	}
}

func (s *AddstylenamedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitAddstylenamedecl(s)
	}
}

func (p *bdfParser) Addstylenamedecl() (localctx IAddstylenamedeclContext) {
	this := p
	_ = this

	localctx = NewAddstylenamedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, bdfParserRULE_addstylenamedecl)

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
		p.SetState(183)
		p.Match(bdfParserT__12)
	}
	{
		p.SetState(184)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// IPixelsizedeclContext is an interface to support dynamic dispatch.
type IPixelsizedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPixelsizedeclContext differentiates from other interfaces.
	IsPixelsizedeclContext()
}

type PixelsizedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPixelsizedeclContext() *PixelsizedeclContext {
	var p = new(PixelsizedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_pixelsizedecl
	return p
}

func (*PixelsizedeclContext) IsPixelsizedeclContext() {}

func NewPixelsizedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PixelsizedeclContext {
	var p = new(PixelsizedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_pixelsizedecl

	return p
}

func (s *PixelsizedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PixelsizedeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *PixelsizedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PixelsizedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PixelsizedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterPixelsizedecl(s)
	}
}

func (s *PixelsizedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitPixelsizedecl(s)
	}
}

func (p *bdfParser) Pixelsizedecl() (localctx IPixelsizedeclContext) {
	this := p
	_ = this

	localctx = NewPixelsizedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, bdfParserRULE_pixelsizedecl)

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
		p.SetState(186)
		p.Match(bdfParserT__13)
	}
	{
		p.SetState(187)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IPointsizedeclContext is an interface to support dynamic dispatch.
type IPointsizedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointsizedeclContext differentiates from other interfaces.
	IsPointsizedeclContext()
}

type PointsizedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointsizedeclContext() *PointsizedeclContext {
	var p = new(PointsizedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_pointsizedecl
	return p
}

func (*PointsizedeclContext) IsPointsizedeclContext() {}

func NewPointsizedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointsizedeclContext {
	var p = new(PointsizedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_pointsizedecl

	return p
}

func (s *PointsizedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PointsizedeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *PointsizedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointsizedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointsizedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterPointsizedecl(s)
	}
}

func (s *PointsizedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitPointsizedecl(s)
	}
}

func (p *bdfParser) Pointsizedecl() (localctx IPointsizedeclContext) {
	this := p
	_ = this

	localctx = NewPointsizedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, bdfParserRULE_pointsizedecl)

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
		p.SetState(189)
		p.Match(bdfParserT__14)
	}
	{
		p.SetState(190)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IResolutionxdeclContext is an interface to support dynamic dispatch.
type IResolutionxdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResolutionxdeclContext differentiates from other interfaces.
	IsResolutionxdeclContext()
}

type ResolutionxdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResolutionxdeclContext() *ResolutionxdeclContext {
	var p = new(ResolutionxdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_resolutionxdecl
	return p
}

func (*ResolutionxdeclContext) IsResolutionxdeclContext() {}

func NewResolutionxdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResolutionxdeclContext {
	var p = new(ResolutionxdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_resolutionxdecl

	return p
}

func (s *ResolutionxdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ResolutionxdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *ResolutionxdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResolutionxdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResolutionxdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterResolutionxdecl(s)
	}
}

func (s *ResolutionxdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitResolutionxdecl(s)
	}
}

func (p *bdfParser) Resolutionxdecl() (localctx IResolutionxdeclContext) {
	this := p
	_ = this

	localctx = NewResolutionxdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, bdfParserRULE_resolutionxdecl)

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
		p.SetState(192)
		p.Match(bdfParserT__15)
	}
	{
		p.SetState(193)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IResolutionydeclContext is an interface to support dynamic dispatch.
type IResolutionydeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsResolutionydeclContext differentiates from other interfaces.
	IsResolutionydeclContext()
}

type ResolutionydeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResolutionydeclContext() *ResolutionydeclContext {
	var p = new(ResolutionydeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_resolutionydecl
	return p
}

func (*ResolutionydeclContext) IsResolutionydeclContext() {}

func NewResolutionydeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResolutionydeclContext {
	var p = new(ResolutionydeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_resolutionydecl

	return p
}

func (s *ResolutionydeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ResolutionydeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *ResolutionydeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResolutionydeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResolutionydeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterResolutionydecl(s)
	}
}

func (s *ResolutionydeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitResolutionydecl(s)
	}
}

func (p *bdfParser) Resolutionydecl() (localctx IResolutionydeclContext) {
	this := p
	_ = this

	localctx = NewResolutionydeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, bdfParserRULE_resolutionydecl)

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
		p.SetState(195)
		p.Match(bdfParserT__16)
	}
	{
		p.SetState(196)
		p.Match(bdfParserARG)
	}

	return localctx
}

// ISpacingdeclContext is an interface to support dynamic dispatch.
type ISpacingdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpacingdeclContext differentiates from other interfaces.
	IsSpacingdeclContext()
}

type SpacingdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpacingdeclContext() *SpacingdeclContext {
	var p = new(SpacingdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_spacingdecl
	return p
}

func (*SpacingdeclContext) IsSpacingdeclContext() {}

func NewSpacingdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpacingdeclContext {
	var p = new(SpacingdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_spacingdecl

	return p
}

func (s *SpacingdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SpacingdeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *SpacingdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpacingdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpacingdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterSpacingdecl(s)
	}
}

func (s *SpacingdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitSpacingdecl(s)
	}
}

func (p *bdfParser) Spacingdecl() (localctx ISpacingdeclContext) {
	this := p
	_ = this

	localctx = NewSpacingdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, bdfParserRULE_spacingdecl)

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
		p.SetState(198)
		p.Match(bdfParserT__17)
	}
	{
		p.SetState(199)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// IAveragewidthdeclContext is an interface to support dynamic dispatch.
type IAveragewidthdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAveragewidthdeclContext differentiates from other interfaces.
	IsAveragewidthdeclContext()
}

type AveragewidthdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAveragewidthdeclContext() *AveragewidthdeclContext {
	var p = new(AveragewidthdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_averagewidthdecl
	return p
}

func (*AveragewidthdeclContext) IsAveragewidthdeclContext() {}

func NewAveragewidthdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AveragewidthdeclContext {
	var p = new(AveragewidthdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_averagewidthdecl

	return p
}

func (s *AveragewidthdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *AveragewidthdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *AveragewidthdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AveragewidthdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AveragewidthdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterAveragewidthdecl(s)
	}
}

func (s *AveragewidthdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitAveragewidthdecl(s)
	}
}

func (p *bdfParser) Averagewidthdecl() (localctx IAveragewidthdeclContext) {
	this := p
	_ = this

	localctx = NewAveragewidthdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, bdfParserRULE_averagewidthdecl)

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
		p.SetState(201)
		p.Match(bdfParserT__18)
	}
	{
		p.SetState(202)
		p.Match(bdfParserARG)
	}

	return localctx
}

// ICharsetregistrydeclContext is an interface to support dynamic dispatch.
type ICharsetregistrydeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharsetregistrydeclContext differentiates from other interfaces.
	IsCharsetregistrydeclContext()
}

type CharsetregistrydeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharsetregistrydeclContext() *CharsetregistrydeclContext {
	var p = new(CharsetregistrydeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_charsetregistrydecl
	return p
}

func (*CharsetregistrydeclContext) IsCharsetregistrydeclContext() {}

func NewCharsetregistrydeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharsetregistrydeclContext {
	var p = new(CharsetregistrydeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_charsetregistrydecl

	return p
}

func (s *CharsetregistrydeclContext) GetParser() antlr.Parser { return s.parser }

func (s *CharsetregistrydeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *CharsetregistrydeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharsetregistrydeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharsetregistrydeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterCharsetregistrydecl(s)
	}
}

func (s *CharsetregistrydeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitCharsetregistrydecl(s)
	}
}

func (p *bdfParser) Charsetregistrydecl() (localctx ICharsetregistrydeclContext) {
	this := p
	_ = this

	localctx = NewCharsetregistrydeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, bdfParserRULE_charsetregistrydecl)

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
		p.SetState(204)
		p.Match(bdfParserT__19)
	}
	{
		p.SetState(205)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// ICharsetencodingContext is an interface to support dynamic dispatch.
type ICharsetencodingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharsetencodingContext differentiates from other interfaces.
	IsCharsetencodingContext()
}

type CharsetencodingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharsetencodingContext() *CharsetencodingContext {
	var p = new(CharsetencodingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_charsetencoding
	return p
}

func (*CharsetencodingContext) IsCharsetencodingContext() {}

func NewCharsetencodingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharsetencodingContext {
	var p = new(CharsetencodingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_charsetencoding

	return p
}

func (s *CharsetencodingContext) GetParser() antlr.Parser { return s.parser }

func (s *CharsetencodingContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *CharsetencodingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharsetencodingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharsetencodingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterCharsetencoding(s)
	}
}

func (s *CharsetencodingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitCharsetencoding(s)
	}
}

func (p *bdfParser) Charsetencoding() (localctx ICharsetencodingContext) {
	this := p
	_ = this

	localctx = NewCharsetencodingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, bdfParserRULE_charsetencoding)

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
		p.SetState(207)
		p.Match(bdfParserT__20)
	}
	{
		p.SetState(208)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// IFontnameregistryContext is an interface to support dynamic dispatch.
type IFontnameregistryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFontnameregistryContext differentiates from other interfaces.
	IsFontnameregistryContext()
}

type FontnameregistryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFontnameregistryContext() *FontnameregistryContext {
	var p = new(FontnameregistryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_fontnameregistry
	return p
}

func (*FontnameregistryContext) IsFontnameregistryContext() {}

func NewFontnameregistryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FontnameregistryContext {
	var p = new(FontnameregistryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_fontnameregistry

	return p
}

func (s *FontnameregistryContext) GetParser() antlr.Parser { return s.parser }

func (s *FontnameregistryContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *FontnameregistryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FontnameregistryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FontnameregistryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterFontnameregistry(s)
	}
}

func (s *FontnameregistryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitFontnameregistry(s)
	}
}

func (p *bdfParser) Fontnameregistry() (localctx IFontnameregistryContext) {
	this := p
	_ = this

	localctx = NewFontnameregistryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, bdfParserRULE_fontnameregistry)

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
		p.SetState(210)
		p.Match(bdfParserT__21)
	}
	{
		p.SetState(211)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// IFontascentdeclContext is an interface to support dynamic dispatch.
type IFontascentdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFontascentdeclContext differentiates from other interfaces.
	IsFontascentdeclContext()
}

type FontascentdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFontascentdeclContext() *FontascentdeclContext {
	var p = new(FontascentdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_fontascentdecl
	return p
}

func (*FontascentdeclContext) IsFontascentdeclContext() {}

func NewFontascentdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FontascentdeclContext {
	var p = new(FontascentdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_fontascentdecl

	return p
}

func (s *FontascentdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FontascentdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *FontascentdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FontascentdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FontascentdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterFontascentdecl(s)
	}
}

func (s *FontascentdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitFontascentdecl(s)
	}
}

func (p *bdfParser) Fontascentdecl() (localctx IFontascentdeclContext) {
	this := p
	_ = this

	localctx = NewFontascentdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, bdfParserRULE_fontascentdecl)

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
		p.SetState(213)
		p.Match(bdfParserT__22)
	}
	{
		p.SetState(214)
		p.Match(bdfParserARG)
	}

	return localctx
}

// ICharsetcollectionsdeclContext is an interface to support dynamic dispatch.
type ICharsetcollectionsdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharsetcollectionsdeclContext differentiates from other interfaces.
	IsCharsetcollectionsdeclContext()
}

type CharsetcollectionsdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharsetcollectionsdeclContext() *CharsetcollectionsdeclContext {
	var p = new(CharsetcollectionsdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_charsetcollectionsdecl
	return p
}

func (*CharsetcollectionsdeclContext) IsCharsetcollectionsdeclContext() {}

func NewCharsetcollectionsdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharsetcollectionsdeclContext {
	var p = new(CharsetcollectionsdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_charsetcollectionsdecl

	return p
}

func (s *CharsetcollectionsdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *CharsetcollectionsdeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *CharsetcollectionsdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharsetcollectionsdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharsetcollectionsdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterCharsetcollectionsdecl(s)
	}
}

func (s *CharsetcollectionsdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitCharsetcollectionsdecl(s)
	}
}

func (p *bdfParser) Charsetcollectionsdecl() (localctx ICharsetcollectionsdeclContext) {
	this := p
	_ = this

	localctx = NewCharsetcollectionsdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, bdfParserRULE_charsetcollectionsdecl)

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
		p.SetState(216)
		p.Match(bdfParserT__23)
	}
	{
		p.SetState(217)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// IFontnamedeclContext is an interface to support dynamic dispatch.
type IFontnamedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFontnamedeclContext differentiates from other interfaces.
	IsFontnamedeclContext()
}

type FontnamedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFontnamedeclContext() *FontnamedeclContext {
	var p = new(FontnamedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_fontnamedecl
	return p
}

func (*FontnamedeclContext) IsFontnamedeclContext() {}

func NewFontnamedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FontnamedeclContext {
	var p = new(FontnamedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_fontnamedecl

	return p
}

func (s *FontnamedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FontnamedeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *FontnamedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FontnamedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FontnamedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterFontnamedecl(s)
	}
}

func (s *FontnamedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitFontnamedecl(s)
	}
}

func (p *bdfParser) Fontnamedecl() (localctx IFontnamedeclContext) {
	this := p
	_ = this

	localctx = NewFontnamedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, bdfParserRULE_fontnamedecl)

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
		p.SetState(219)
		p.Match(bdfParserT__24)
	}
	{
		p.SetState(220)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// IFacenamedeclContext is an interface to support dynamic dispatch.
type IFacenamedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFacenamedeclContext differentiates from other interfaces.
	IsFacenamedeclContext()
}

type FacenamedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFacenamedeclContext() *FacenamedeclContext {
	var p = new(FacenamedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_facenamedecl
	return p
}

func (*FacenamedeclContext) IsFacenamedeclContext() {}

func NewFacenamedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FacenamedeclContext {
	var p = new(FacenamedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_facenamedecl

	return p
}

func (s *FacenamedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FacenamedeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *FacenamedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FacenamedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FacenamedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterFacenamedecl(s)
	}
}

func (s *FacenamedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitFacenamedecl(s)
	}
}

func (p *bdfParser) Facenamedecl() (localctx IFacenamedeclContext) {
	this := p
	_ = this

	localctx = NewFacenamedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, bdfParserRULE_facenamedecl)

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
		p.SetState(222)
		p.Match(bdfParserT__25)
	}
	{
		p.SetState(223)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// ICopyrightdeclContext is an interface to support dynamic dispatch.
type ICopyrightdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCopyrightdeclContext differentiates from other interfaces.
	IsCopyrightdeclContext()
}

type CopyrightdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCopyrightdeclContext() *CopyrightdeclContext {
	var p = new(CopyrightdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_copyrightdecl
	return p
}

func (*CopyrightdeclContext) IsCopyrightdeclContext() {}

func NewCopyrightdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CopyrightdeclContext {
	var p = new(CopyrightdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_copyrightdecl

	return p
}

func (s *CopyrightdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *CopyrightdeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *CopyrightdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CopyrightdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CopyrightdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterCopyrightdecl(s)
	}
}

func (s *CopyrightdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitCopyrightdecl(s)
	}
}

func (p *bdfParser) Copyrightdecl() (localctx ICopyrightdeclContext) {
	this := p
	_ = this

	localctx = NewCopyrightdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, bdfParserRULE_copyrightdecl)

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
		p.SetState(225)
		p.Match(bdfParserT__26)
	}
	{
		p.SetState(226)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// IFontdecentdeclContext is an interface to support dynamic dispatch.
type IFontdecentdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFontdecentdeclContext differentiates from other interfaces.
	IsFontdecentdeclContext()
}

type FontdecentdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFontdecentdeclContext() *FontdecentdeclContext {
	var p = new(FontdecentdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_fontdecentdecl
	return p
}

func (*FontdecentdeclContext) IsFontdecentdeclContext() {}

func NewFontdecentdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FontdecentdeclContext {
	var p = new(FontdecentdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_fontdecentdecl

	return p
}

func (s *FontdecentdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FontdecentdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *FontdecentdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FontdecentdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FontdecentdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterFontdecentdecl(s)
	}
}

func (s *FontdecentdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitFontdecentdecl(s)
	}
}

func (p *bdfParser) Fontdecentdecl() (localctx IFontdecentdeclContext) {
	this := p
	_ = this

	localctx = NewFontdecentdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, bdfParserRULE_fontdecentdecl)

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
		p.SetState(228)
		p.Match(bdfParserT__27)
	}
	{
		p.SetState(229)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IFontversiondeclContext is an interface to support dynamic dispatch.
type IFontversiondeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFontversiondeclContext differentiates from other interfaces.
	IsFontversiondeclContext()
}

type FontversiondeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFontversiondeclContext() *FontversiondeclContext {
	var p = new(FontversiondeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_fontversiondecl
	return p
}

func (*FontversiondeclContext) IsFontversiondeclContext() {}

func NewFontversiondeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FontversiondeclContext {
	var p = new(FontversiondeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_fontversiondecl

	return p
}

func (s *FontversiondeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FontversiondeclContext) QUOTEDSTRING() antlr.TerminalNode {
	return s.GetToken(bdfParserQUOTEDSTRING, 0)
}

func (s *FontversiondeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FontversiondeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FontversiondeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterFontversiondecl(s)
	}
}

func (s *FontversiondeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitFontversiondecl(s)
	}
}

func (p *bdfParser) Fontversiondecl() (localctx IFontversiondeclContext) {
	this := p
	_ = this

	localctx = NewFontversiondeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, bdfParserRULE_fontversiondecl)

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
		p.SetState(231)
		p.Match(bdfParserT__28)
	}
	{
		p.SetState(232)
		p.Match(bdfParserQUOTEDSTRING)
	}

	return localctx
}

// IUnderlinepositiondeclContext is an interface to support dynamic dispatch.
type IUnderlinepositiondeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnderlinepositiondeclContext differentiates from other interfaces.
	IsUnderlinepositiondeclContext()
}

type UnderlinepositiondeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnderlinepositiondeclContext() *UnderlinepositiondeclContext {
	var p = new(UnderlinepositiondeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_underlinepositiondecl
	return p
}

func (*UnderlinepositiondeclContext) IsUnderlinepositiondeclContext() {}

func NewUnderlinepositiondeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnderlinepositiondeclContext {
	var p = new(UnderlinepositiondeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_underlinepositiondecl

	return p
}

func (s *UnderlinepositiondeclContext) GetParser() antlr.Parser { return s.parser }

func (s *UnderlinepositiondeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *UnderlinepositiondeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnderlinepositiondeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnderlinepositiondeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterUnderlinepositiondecl(s)
	}
}

func (s *UnderlinepositiondeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitUnderlinepositiondecl(s)
	}
}

func (p *bdfParser) Underlinepositiondecl() (localctx IUnderlinepositiondeclContext) {
	this := p
	_ = this

	localctx = NewUnderlinepositiondeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, bdfParserRULE_underlinepositiondecl)

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
		p.SetState(234)
		p.Match(bdfParserT__29)
	}
	{
		p.SetState(235)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IUnderlinethicknessdeclContext is an interface to support dynamic dispatch.
type IUnderlinethicknessdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnderlinethicknessdeclContext differentiates from other interfaces.
	IsUnderlinethicknessdeclContext()
}

type UnderlinethicknessdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnderlinethicknessdeclContext() *UnderlinethicknessdeclContext {
	var p = new(UnderlinethicknessdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_underlinethicknessdecl
	return p
}

func (*UnderlinethicknessdeclContext) IsUnderlinethicknessdeclContext() {}

func NewUnderlinethicknessdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnderlinethicknessdeclContext {
	var p = new(UnderlinethicknessdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_underlinethicknessdecl

	return p
}

func (s *UnderlinethicknessdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *UnderlinethicknessdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *UnderlinethicknessdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnderlinethicknessdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnderlinethicknessdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterUnderlinethicknessdecl(s)
	}
}

func (s *UnderlinethicknessdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitUnderlinethicknessdecl(s)
	}
}

func (p *bdfParser) Underlinethicknessdecl() (localctx IUnderlinethicknessdeclContext) {
	this := p
	_ = this

	localctx = NewUnderlinethicknessdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, bdfParserRULE_underlinethicknessdecl)

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
		p.SetState(237)
		p.Match(bdfParserT__30)
	}
	{
		p.SetState(238)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IXheightdeclContext is an interface to support dynamic dispatch.
type IXheightdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsXheightdeclContext differentiates from other interfaces.
	IsXheightdeclContext()
}

type XheightdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyXheightdeclContext() *XheightdeclContext {
	var p = new(XheightdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_xheightdecl
	return p
}

func (*XheightdeclContext) IsXheightdeclContext() {}

func NewXheightdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *XheightdeclContext {
	var p = new(XheightdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_xheightdecl

	return p
}

func (s *XheightdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *XheightdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *XheightdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XheightdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *XheightdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterXheightdecl(s)
	}
}

func (s *XheightdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitXheightdecl(s)
	}
}

func (p *bdfParser) Xheightdecl() (localctx IXheightdeclContext) {
	this := p
	_ = this

	localctx = NewXheightdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, bdfParserRULE_xheightdecl)

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
		p.SetState(240)
		p.Match(bdfParserT__31)
	}
	{
		p.SetState(241)
		p.Match(bdfParserARG)
	}

	return localctx
}

// ICapheighdeclContext is an interface to support dynamic dispatch.
type ICapheighdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCapheighdeclContext differentiates from other interfaces.
	IsCapheighdeclContext()
}

type CapheighdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCapheighdeclContext() *CapheighdeclContext {
	var p = new(CapheighdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_capheighdecl
	return p
}

func (*CapheighdeclContext) IsCapheighdeclContext() {}

func NewCapheighdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CapheighdeclContext {
	var p = new(CapheighdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_capheighdecl

	return p
}

func (s *CapheighdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *CapheighdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *CapheighdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CapheighdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CapheighdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterCapheighdecl(s)
	}
}

func (s *CapheighdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitCapheighdecl(s)
	}
}

func (p *bdfParser) Capheighdecl() (localctx ICapheighdeclContext) {
	this := p
	_ = this

	localctx = NewCapheighdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, bdfParserRULE_capheighdecl)

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
		p.SetState(243)
		p.Match(bdfParserT__32)
	}
	{
		p.SetState(244)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IRawascentdeclContext is an interface to support dynamic dispatch.
type IRawascentdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRawascentdeclContext differentiates from other interfaces.
	IsRawascentdeclContext()
}

type RawascentdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRawascentdeclContext() *RawascentdeclContext {
	var p = new(RawascentdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_rawascentdecl
	return p
}

func (*RawascentdeclContext) IsRawascentdeclContext() {}

func NewRawascentdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RawascentdeclContext {
	var p = new(RawascentdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_rawascentdecl

	return p
}

func (s *RawascentdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *RawascentdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *RawascentdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawascentdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RawascentdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterRawascentdecl(s)
	}
}

func (s *RawascentdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitRawascentdecl(s)
	}
}

func (p *bdfParser) Rawascentdecl() (localctx IRawascentdeclContext) {
	this := p
	_ = this

	localctx = NewRawascentdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, bdfParserRULE_rawascentdecl)

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
		p.SetState(246)
		p.Match(bdfParserT__33)
	}
	{
		p.SetState(247)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IRawdescentdeclContext is an interface to support dynamic dispatch.
type IRawdescentdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRawdescentdeclContext differentiates from other interfaces.
	IsRawdescentdeclContext()
}

type RawdescentdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRawdescentdeclContext() *RawdescentdeclContext {
	var p = new(RawdescentdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_rawdescentdecl
	return p
}

func (*RawdescentdeclContext) IsRawdescentdeclContext() {}

func NewRawdescentdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RawdescentdeclContext {
	var p = new(RawdescentdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_rawdescentdecl

	return p
}

func (s *RawdescentdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *RawdescentdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *RawdescentdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RawdescentdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RawdescentdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterRawdescentdecl(s)
	}
}

func (s *RawdescentdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitRawdescentdecl(s)
	}
}

func (p *bdfParser) Rawdescentdecl() (localctx IRawdescentdeclContext) {
	this := p
	_ = this

	localctx = NewRawdescentdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, bdfParserRULE_rawdescentdecl)

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
		p.SetState(249)
		p.Match(bdfParserT__34)
	}
	{
		p.SetState(250)
		p.Match(bdfParserARG)
	}

	return localctx
}

// INormspacedeclContext is an interface to support dynamic dispatch.
type INormspacedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNormspacedeclContext differentiates from other interfaces.
	IsNormspacedeclContext()
}

type NormspacedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNormspacedeclContext() *NormspacedeclContext {
	var p = new(NormspacedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_normspacedecl
	return p
}

func (*NormspacedeclContext) IsNormspacedeclContext() {}

func NewNormspacedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NormspacedeclContext {
	var p = new(NormspacedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_normspacedecl

	return p
}

func (s *NormspacedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *NormspacedeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *NormspacedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NormspacedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NormspacedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterNormspacedecl(s)
	}
}

func (s *NormspacedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitNormspacedecl(s)
	}
}

func (p *bdfParser) Normspacedecl() (localctx INormspacedeclContext) {
	this := p
	_ = this

	localctx = NewNormspacedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, bdfParserRULE_normspacedecl)

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
		p.SetState(252)
		p.Match(bdfParserT__35)
	}
	{
		p.SetState(253)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IRelativeweightdeclContext is an interface to support dynamic dispatch.
type IRelativeweightdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelativeweightdeclContext differentiates from other interfaces.
	IsRelativeweightdeclContext()
}

type RelativeweightdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelativeweightdeclContext() *RelativeweightdeclContext {
	var p = new(RelativeweightdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_relativeweightdecl
	return p
}

func (*RelativeweightdeclContext) IsRelativeweightdeclContext() {}

func NewRelativeweightdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelativeweightdeclContext {
	var p = new(RelativeweightdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_relativeweightdecl

	return p
}

func (s *RelativeweightdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *RelativeweightdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *RelativeweightdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelativeweightdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelativeweightdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterRelativeweightdecl(s)
	}
}

func (s *RelativeweightdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitRelativeweightdecl(s)
	}
}

func (p *bdfParser) Relativeweightdecl() (localctx IRelativeweightdeclContext) {
	this := p
	_ = this

	localctx = NewRelativeweightdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, bdfParserRULE_relativeweightdecl)

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
		p.SetState(255)
		p.Match(bdfParserT__36)
	}
	{
		p.SetState(256)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IRelaticesetwidthdeclContext is an interface to support dynamic dispatch.
type IRelaticesetwidthdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelaticesetwidthdeclContext differentiates from other interfaces.
	IsRelaticesetwidthdeclContext()
}

type RelaticesetwidthdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelaticesetwidthdeclContext() *RelaticesetwidthdeclContext {
	var p = new(RelaticesetwidthdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_relaticesetwidthdecl
	return p
}

func (*RelaticesetwidthdeclContext) IsRelaticesetwidthdeclContext() {}

func NewRelaticesetwidthdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelaticesetwidthdeclContext {
	var p = new(RelaticesetwidthdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_relaticesetwidthdecl

	return p
}

func (s *RelaticesetwidthdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *RelaticesetwidthdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *RelaticesetwidthdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelaticesetwidthdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelaticesetwidthdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterRelaticesetwidthdecl(s)
	}
}

func (s *RelaticesetwidthdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitRelaticesetwidthdecl(s)
	}
}

func (p *bdfParser) Relaticesetwidthdecl() (localctx IRelaticesetwidthdeclContext) {
	this := p
	_ = this

	localctx = NewRelaticesetwidthdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, bdfParserRULE_relaticesetwidthdecl)

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
		p.SetState(258)
		p.Match(bdfParserT__37)
	}
	{
		p.SetState(259)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IFigurewidthdeclContext is an interface to support dynamic dispatch.
type IFigurewidthdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFigurewidthdeclContext differentiates from other interfaces.
	IsFigurewidthdeclContext()
}

type FigurewidthdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFigurewidthdeclContext() *FigurewidthdeclContext {
	var p = new(FigurewidthdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_figurewidthdecl
	return p
}

func (*FigurewidthdeclContext) IsFigurewidthdeclContext() {}

func NewFigurewidthdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FigurewidthdeclContext {
	var p = new(FigurewidthdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_figurewidthdecl

	return p
}

func (s *FigurewidthdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FigurewidthdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *FigurewidthdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FigurewidthdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FigurewidthdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterFigurewidthdecl(s)
	}
}

func (s *FigurewidthdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitFigurewidthdecl(s)
	}
}

func (p *bdfParser) Figurewidthdecl() (localctx IFigurewidthdeclContext) {
	this := p
	_ = this

	localctx = NewFigurewidthdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, bdfParserRULE_figurewidthdecl)

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
		p.SetState(261)
		p.Match(bdfParserT__38)
	}
	{
		p.SetState(262)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IAvglowercasewidthdeclContext is an interface to support dynamic dispatch.
type IAvglowercasewidthdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAvglowercasewidthdeclContext differentiates from other interfaces.
	IsAvglowercasewidthdeclContext()
}

type AvglowercasewidthdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAvglowercasewidthdeclContext() *AvglowercasewidthdeclContext {
	var p = new(AvglowercasewidthdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_avglowercasewidthdecl
	return p
}

func (*AvglowercasewidthdeclContext) IsAvglowercasewidthdeclContext() {}

func NewAvglowercasewidthdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AvglowercasewidthdeclContext {
	var p = new(AvglowercasewidthdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_avglowercasewidthdecl

	return p
}

func (s *AvglowercasewidthdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *AvglowercasewidthdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *AvglowercasewidthdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AvglowercasewidthdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AvglowercasewidthdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterAvglowercasewidthdecl(s)
	}
}

func (s *AvglowercasewidthdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitAvglowercasewidthdecl(s)
	}
}

func (p *bdfParser) Avglowercasewidthdecl() (localctx IAvglowercasewidthdeclContext) {
	this := p
	_ = this

	localctx = NewAvglowercasewidthdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, bdfParserRULE_avglowercasewidthdecl)

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
		p.SetState(264)
		p.Match(bdfParserT__39)
	}
	{
		p.SetState(265)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IAvguppercasewidthdeclContext is an interface to support dynamic dispatch.
type IAvguppercasewidthdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAvguppercasewidthdeclContext differentiates from other interfaces.
	IsAvguppercasewidthdeclContext()
}

type AvguppercasewidthdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAvguppercasewidthdeclContext() *AvguppercasewidthdeclContext {
	var p = new(AvguppercasewidthdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_avguppercasewidthdecl
	return p
}

func (*AvguppercasewidthdeclContext) IsAvguppercasewidthdeclContext() {}

func NewAvguppercasewidthdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AvguppercasewidthdeclContext {
	var p = new(AvguppercasewidthdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_avguppercasewidthdecl

	return p
}

func (s *AvguppercasewidthdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *AvguppercasewidthdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *AvguppercasewidthdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AvguppercasewidthdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AvguppercasewidthdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterAvguppercasewidthdecl(s)
	}
}

func (s *AvguppercasewidthdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitAvguppercasewidthdecl(s)
	}
}

func (p *bdfParser) Avguppercasewidthdecl() (localctx IAvguppercasewidthdeclContext) {
	this := p
	_ = this

	localctx = NewAvguppercasewidthdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, bdfParserRULE_avguppercasewidthdecl)

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
		p.SetState(267)
		p.Match(bdfParserT__40)
	}
	{
		p.SetState(268)
		p.Match(bdfParserARG)
	}

	return localctx
}

// ICharsdeclContext is an interface to support dynamic dispatch.
type ICharsdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCharsdeclContext differentiates from other interfaces.
	IsCharsdeclContext()
}

type CharsdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCharsdeclContext() *CharsdeclContext {
	var p = new(CharsdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_charsdecl
	return p
}

func (*CharsdeclContext) IsCharsdeclContext() {}

func NewCharsdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CharsdeclContext {
	var p = new(CharsdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_charsdecl

	return p
}

func (s *CharsdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *CharsdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *CharsdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharsdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CharsdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterCharsdecl(s)
	}
}

func (s *CharsdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitCharsdecl(s)
	}
}

func (p *bdfParser) Charsdecl() (localctx ICharsdeclContext) {
	this := p
	_ = this

	localctx = NewCharsdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, bdfParserRULE_charsdecl)

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
		p.SetState(270)
		p.Match(bdfParserT__41)
	}
	{
		p.SetState(271)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IChardeclContext is an interface to support dynamic dispatch.
type IChardeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChardeclContext differentiates from other interfaces.
	IsChardeclContext()
}

type ChardeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChardeclContext() *ChardeclContext {
	var p = new(ChardeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_chardecl
	return p
}

func (*ChardeclContext) IsChardeclContext() {}

func NewChardeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChardeclContext {
	var p = new(ChardeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_chardecl

	return p
}

func (s *ChardeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ChardeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *ChardeclContext) AllEncodingdecl() []IEncodingdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEncodingdeclContext)(nil)).Elem())
	var tst = make([]IEncodingdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEncodingdeclContext)
		}
	}

	return tst
}

func (s *ChardeclContext) Encodingdecl(i int) IEncodingdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEncodingdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEncodingdeclContext)
}

func (s *ChardeclContext) AllSwidthdecl() []ISwidthdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISwidthdeclContext)(nil)).Elem())
	var tst = make([]ISwidthdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISwidthdeclContext)
		}
	}

	return tst
}

func (s *ChardeclContext) Swidthdecl(i int) ISwidthdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwidthdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISwidthdeclContext)
}

func (s *ChardeclContext) AllDwidthdecl() []IDwidthdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDwidthdeclContext)(nil)).Elem())
	var tst = make([]IDwidthdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDwidthdeclContext)
		}
	}

	return tst
}

func (s *ChardeclContext) Dwidthdecl(i int) IDwidthdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDwidthdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDwidthdeclContext)
}

func (s *ChardeclContext) AllBbxdecl() []IBbxdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBbxdeclContext)(nil)).Elem())
	var tst = make([]IBbxdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBbxdeclContext)
		}
	}

	return tst
}

func (s *ChardeclContext) Bbxdecl(i int) IBbxdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBbxdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBbxdeclContext)
}

func (s *ChardeclContext) AllBitmapdecl() []IBitmapdeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBitmapdeclContext)(nil)).Elem())
	var tst = make([]IBitmapdeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBitmapdeclContext)
		}
	}

	return tst
}

func (s *ChardeclContext) Bitmapdecl(i int) IBitmapdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBitmapdeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBitmapdeclContext)
}

func (s *ChardeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChardeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChardeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterChardecl(s)
	}
}

func (s *ChardeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitChardecl(s)
	}
}

func (p *bdfParser) Chardecl() (localctx IChardeclContext) {
	this := p
	_ = this

	localctx = NewChardeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, bdfParserRULE_chardecl)
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
		p.SetState(273)
		p.Match(bdfParserT__42)
	}
	{
		p.SetState(274)
		p.Match(bdfParserARG)
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(bdfParserT__44-45))|(1<<(bdfParserT__45-45))|(1<<(bdfParserT__46-45))|(1<<(bdfParserT__47-45))|(1<<(bdfParserT__48-45)))) != 0 {
		p.SetState(280)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case bdfParserT__44:
			{
				p.SetState(275)
				p.Encodingdecl()
			}

		case bdfParserT__45:
			{
				p.SetState(276)
				p.Swidthdecl()
			}

		case bdfParserT__46:
			{
				p.SetState(277)
				p.Dwidthdecl()
			}

		case bdfParserT__47:
			{
				p.SetState(278)
				p.Bbxdecl()
			}

		case bdfParserT__48:
			{
				p.SetState(279)
				p.Bitmapdecl()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(285)
		p.Match(bdfParserT__43)
	}

	return localctx
}

// IEncodingdeclContext is an interface to support dynamic dispatch.
type IEncodingdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEncodingdeclContext differentiates from other interfaces.
	IsEncodingdeclContext()
}

type EncodingdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEncodingdeclContext() *EncodingdeclContext {
	var p = new(EncodingdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_encodingdecl
	return p
}

func (*EncodingdeclContext) IsEncodingdeclContext() {}

func NewEncodingdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EncodingdeclContext {
	var p = new(EncodingdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_encodingdecl

	return p
}

func (s *EncodingdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *EncodingdeclContext) ARG() antlr.TerminalNode {
	return s.GetToken(bdfParserARG, 0)
}

func (s *EncodingdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EncodingdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EncodingdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterEncodingdecl(s)
	}
}

func (s *EncodingdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitEncodingdecl(s)
	}
}

func (p *bdfParser) Encodingdecl() (localctx IEncodingdeclContext) {
	this := p
	_ = this

	localctx = NewEncodingdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, bdfParserRULE_encodingdecl)

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
		p.SetState(287)
		p.Match(bdfParserT__44)
	}
	{
		p.SetState(288)
		p.Match(bdfParserARG)
	}

	return localctx
}

// ISwidthdeclContext is an interface to support dynamic dispatch.
type ISwidthdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwidthdeclContext differentiates from other interfaces.
	IsSwidthdeclContext()
}

type SwidthdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwidthdeclContext() *SwidthdeclContext {
	var p = new(SwidthdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_swidthdecl
	return p
}

func (*SwidthdeclContext) IsSwidthdeclContext() {}

func NewSwidthdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwidthdeclContext {
	var p = new(SwidthdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_swidthdecl

	return p
}

func (s *SwidthdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SwidthdeclContext) AllARG() []antlr.TerminalNode {
	return s.GetTokens(bdfParserARG)
}

func (s *SwidthdeclContext) ARG(i int) antlr.TerminalNode {
	return s.GetToken(bdfParserARG, i)
}

func (s *SwidthdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwidthdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwidthdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterSwidthdecl(s)
	}
}

func (s *SwidthdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitSwidthdecl(s)
	}
}

func (p *bdfParser) Swidthdecl() (localctx ISwidthdeclContext) {
	this := p
	_ = this

	localctx = NewSwidthdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, bdfParserRULE_swidthdecl)

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
		p.SetState(290)
		p.Match(bdfParserT__45)
	}
	{
		p.SetState(291)
		p.Match(bdfParserARG)
	}
	{
		p.SetState(292)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IDwidthdeclContext is an interface to support dynamic dispatch.
type IDwidthdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDwidthdeclContext differentiates from other interfaces.
	IsDwidthdeclContext()
}

type DwidthdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDwidthdeclContext() *DwidthdeclContext {
	var p = new(DwidthdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_dwidthdecl
	return p
}

func (*DwidthdeclContext) IsDwidthdeclContext() {}

func NewDwidthdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DwidthdeclContext {
	var p = new(DwidthdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_dwidthdecl

	return p
}

func (s *DwidthdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DwidthdeclContext) AllARG() []antlr.TerminalNode {
	return s.GetTokens(bdfParserARG)
}

func (s *DwidthdeclContext) ARG(i int) antlr.TerminalNode {
	return s.GetToken(bdfParserARG, i)
}

func (s *DwidthdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DwidthdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DwidthdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterDwidthdecl(s)
	}
}

func (s *DwidthdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitDwidthdecl(s)
	}
}

func (p *bdfParser) Dwidthdecl() (localctx IDwidthdeclContext) {
	this := p
	_ = this

	localctx = NewDwidthdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, bdfParserRULE_dwidthdecl)

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
		p.SetState(294)
		p.Match(bdfParserT__46)
	}
	{
		p.SetState(295)
		p.Match(bdfParserARG)
	}
	{
		p.SetState(296)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IBbxdeclContext is an interface to support dynamic dispatch.
type IBbxdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBbxdeclContext differentiates from other interfaces.
	IsBbxdeclContext()
}

type BbxdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBbxdeclContext() *BbxdeclContext {
	var p = new(BbxdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_bbxdecl
	return p
}

func (*BbxdeclContext) IsBbxdeclContext() {}

func NewBbxdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BbxdeclContext {
	var p = new(BbxdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_bbxdecl

	return p
}

func (s *BbxdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *BbxdeclContext) AllARG() []antlr.TerminalNode {
	return s.GetTokens(bdfParserARG)
}

func (s *BbxdeclContext) ARG(i int) antlr.TerminalNode {
	return s.GetToken(bdfParserARG, i)
}

func (s *BbxdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BbxdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BbxdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterBbxdecl(s)
	}
}

func (s *BbxdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitBbxdecl(s)
	}
}

func (p *bdfParser) Bbxdecl() (localctx IBbxdeclContext) {
	this := p
	_ = this

	localctx = NewBbxdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, bdfParserRULE_bbxdecl)

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
		p.SetState(298)
		p.Match(bdfParserT__47)
	}
	{
		p.SetState(299)
		p.Match(bdfParserARG)
	}
	{
		p.SetState(300)
		p.Match(bdfParserARG)
	}
	{
		p.SetState(301)
		p.Match(bdfParserARG)
	}
	{
		p.SetState(302)
		p.Match(bdfParserARG)
	}

	return localctx
}

// IBitmapdeclContext is an interface to support dynamic dispatch.
type IBitmapdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBitmapdeclContext differentiates from other interfaces.
	IsBitmapdeclContext()
}

type BitmapdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBitmapdeclContext() *BitmapdeclContext {
	var p = new(BitmapdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = bdfParserRULE_bitmapdecl
	return p
}

func (*BitmapdeclContext) IsBitmapdeclContext() {}

func NewBitmapdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BitmapdeclContext {
	var p = new(BitmapdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = bdfParserRULE_bitmapdecl

	return p
}

func (s *BitmapdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *BitmapdeclContext) AllARG() []antlr.TerminalNode {
	return s.GetTokens(bdfParserARG)
}

func (s *BitmapdeclContext) ARG(i int) antlr.TerminalNode {
	return s.GetToken(bdfParserARG, i)
}

func (s *BitmapdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitmapdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BitmapdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.EnterBitmapdecl(s)
	}
}

func (s *BitmapdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(bdfListener); ok {
		listenerT.ExitBitmapdecl(s)
	}
}

func (p *bdfParser) Bitmapdecl() (localctx IBitmapdeclContext) {
	this := p
	_ = this

	localctx = NewBitmapdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, bdfParserRULE_bitmapdecl)
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
		p.SetState(304)
		p.Match(bdfParserT__48)
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == bdfParserARG {
		{
			p.SetState(305)
			p.Match(bdfParserARG)
		}

		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
