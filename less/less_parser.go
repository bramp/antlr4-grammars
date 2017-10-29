// Generated from LessParser.g4 by ANTLR 4.7.

package less // LessParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 145, 339,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 3, 2, 7, 2, 74, 10, 2, 12, 2, 14, 2,
	77, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 85, 10, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 5, 4, 91, 10, 4, 3, 5, 6, 5, 94, 10, 5, 13, 5, 14, 5, 95,
	3, 5, 5, 5, 99, 10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 5, 8, 111, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8,
	119, 10, 8, 3, 9, 3, 9, 3, 9, 5, 9, 124, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 7, 10, 131, 10, 10, 12, 10, 14, 10, 134, 11, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 145, 10, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 152, 10, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 14, 3, 14, 3, 14, 5, 14, 161, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 15, 5, 15, 169, 10, 15, 3, 16, 3, 16, 3, 16, 7, 16, 174, 10,
	16, 12, 16, 14, 16, 177, 11, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 7, 18, 188, 10, 18, 12, 18, 14, 18, 191, 11, 18,
	3, 18, 5, 18, 194, 10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 7, 19, 203, 10, 19, 12, 19, 14, 19, 206, 11, 19, 5, 19, 208, 10, 19,
	3, 19, 5, 19, 211, 10, 19, 3, 19, 3, 19, 5, 19, 215, 10, 19, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 5, 21, 224, 10, 21, 3, 22, 3, 22,
	3, 22, 5, 22, 229, 10, 22, 3, 22, 3, 22, 5, 22, 233, 10, 22, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 23, 7, 23, 240, 10, 23, 12, 23, 14, 23, 243, 11, 23,
	3, 24, 6, 24, 246, 10, 24, 13, 24, 14, 24, 247, 3, 24, 3, 24, 3, 24, 7,
	24, 253, 10, 24, 12, 24, 14, 24, 256, 11, 24, 3, 24, 7, 24, 259, 10, 24,
	12, 24, 14, 24, 262, 11, 24, 3, 24, 5, 24, 265, 10, 24, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 5, 25, 272, 10, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 286, 10, 27, 3,
	28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 7, 30, 294, 10, 30, 12, 30, 14,
	30, 297, 11, 30, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 303, 10, 30, 12, 30,
	14, 30, 306, 11, 30, 5, 30, 308, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 5, 31, 315, 10, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34,
	3, 34, 3, 34, 7, 34, 326, 10, 34, 12, 34, 14, 34, 329, 11, 34, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 36, 3, 36, 5, 36, 337, 10, 36, 3, 36, 2, 2, 37, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 2, 10, 3, 2,
	26, 30, 4, 2, 19, 19, 54, 54, 6, 2, 14, 14, 16, 16, 32, 33, 35, 35, 3,
	2, 55, 56, 4, 2, 17, 17, 25, 25, 4, 2, 14, 15, 26, 26, 3, 2, 35, 37, 4,
	2, 55, 55, 145, 145, 2, 351, 2, 75, 3, 2, 2, 2, 4, 84, 3, 2, 2, 2, 6, 90,
	3, 2, 2, 2, 8, 93, 3, 2, 2, 2, 10, 100, 3, 2, 2, 2, 12, 102, 3, 2, 2, 2,
	14, 118, 3, 2, 2, 2, 16, 120, 3, 2, 2, 2, 18, 127, 3, 2, 2, 2, 20, 144,
	3, 2, 2, 2, 22, 151, 3, 2, 2, 2, 24, 153, 3, 2, 2, 2, 26, 157, 3, 2, 2,
	2, 28, 168, 3, 2, 2, 2, 30, 170, 3, 2, 2, 2, 32, 178, 3, 2, 2, 2, 34, 181,
	3, 2, 2, 2, 36, 197, 3, 2, 2, 2, 38, 218, 3, 2, 2, 2, 40, 223, 3, 2, 2,
	2, 42, 225, 3, 2, 2, 2, 44, 236, 3, 2, 2, 2, 46, 245, 3, 2, 2, 2, 48, 266,
	3, 2, 2, 2, 50, 275, 3, 2, 2, 2, 52, 285, 3, 2, 2, 2, 54, 287, 3, 2, 2,
	2, 56, 289, 3, 2, 2, 2, 58, 307, 3, 2, 2, 2, 60, 314, 3, 2, 2, 2, 62, 316,
	3, 2, 2, 2, 64, 318, 3, 2, 2, 2, 66, 322, 3, 2, 2, 2, 68, 330, 3, 2, 2,
	2, 70, 334, 3, 2, 2, 2, 72, 74, 5, 4, 3, 2, 73, 72, 3, 2, 2, 2, 74, 77,
	3, 2, 2, 2, 75, 73, 3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 3, 3, 2, 2, 2,
	77, 75, 3, 2, 2, 2, 78, 85, 5, 26, 14, 2, 79, 85, 5, 32, 17, 2, 80, 81,
	5, 24, 13, 2, 81, 82, 7, 18, 2, 2, 82, 85, 3, 2, 2, 2, 83, 85, 5, 36, 19,
	2, 84, 78, 3, 2, 2, 2, 84, 79, 3, 2, 2, 2, 84, 80, 3, 2, 2, 2, 84, 83,
	3, 2, 2, 2, 85, 5, 3, 2, 2, 2, 86, 87, 7, 22, 2, 2, 87, 91, 5, 6, 4, 2,
	88, 89, 7, 22, 2, 2, 89, 91, 7, 55, 2, 2, 90, 86, 3, 2, 2, 2, 90, 88, 3,
	2, 2, 2, 91, 7, 3, 2, 2, 2, 92, 94, 5, 14, 8, 2, 93, 92, 3, 2, 2, 2, 94,
	95, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 98, 3, 2, 2,
	2, 97, 99, 5, 12, 7, 2, 98, 97, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 9,
	3, 2, 2, 2, 100, 101, 9, 2, 2, 2, 101, 11, 3, 2, 2, 2, 102, 103, 5, 10,
	6, 2, 103, 104, 5, 8, 5, 2, 104, 13, 3, 2, 2, 2, 105, 119, 5, 70, 36, 2,
	106, 119, 5, 58, 30, 2, 107, 108, 5, 58, 30, 2, 108, 110, 7, 8, 2, 2, 109,
	111, 5, 66, 34, 2, 110, 109, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 112,
	3, 2, 2, 2, 112, 113, 7, 9, 2, 2, 113, 119, 3, 2, 2, 2, 114, 119, 7, 58,
	2, 2, 115, 119, 7, 56, 2, 2, 116, 119, 5, 68, 35, 2, 117, 119, 5, 6, 4,
	2, 118, 105, 3, 2, 2, 2, 118, 106, 3, 2, 2, 2, 118, 107, 3, 2, 2, 2, 118,
	114, 3, 2, 2, 2, 118, 115, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 117,
	3, 2, 2, 2, 119, 15, 3, 2, 2, 2, 120, 121, 7, 62, 2, 2, 121, 123, 7, 8,
	2, 2, 122, 124, 5, 66, 34, 2, 123, 122, 3, 2, 2, 2, 123, 124, 3, 2, 2,
	2, 124, 125, 3, 2, 2, 2, 125, 126, 7, 9, 2, 2, 126, 17, 3, 2, 2, 2, 127,
	132, 5, 20, 11, 2, 128, 129, 9, 3, 2, 2, 129, 131, 5, 20, 11, 2, 130, 128,
	3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2,
	2, 2, 133, 19, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 135, 136, 7, 8, 2, 2,
	136, 137, 5, 22, 12, 2, 137, 138, 7, 9, 2, 2, 138, 145, 3, 2, 2, 2, 139,
	140, 7, 53, 2, 2, 140, 141, 7, 8, 2, 2, 141, 142, 5, 22, 12, 2, 142, 143,
	7, 9, 2, 2, 143, 145, 3, 2, 2, 2, 144, 135, 3, 2, 2, 2, 144, 139, 3, 2,
	2, 2, 145, 21, 3, 2, 2, 2, 146, 147, 5, 8, 5, 2, 147, 148, 9, 4, 2, 2,
	148, 149, 5, 8, 5, 2, 149, 152, 3, 2, 2, 2, 150, 152, 5, 8, 5, 2, 151,
	146, 3, 2, 2, 2, 151, 150, 3, 2, 2, 2, 152, 23, 3, 2, 2, 2, 153, 154, 5,
	6, 4, 2, 154, 155, 7, 17, 2, 2, 155, 156, 5, 66, 34, 2, 156, 25, 3, 2,
	2, 2, 157, 158, 7, 40, 2, 2, 158, 160, 5, 28, 15, 2, 159, 161, 5, 30, 16,
	2, 160, 159, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162,
	163, 7, 18, 2, 2, 163, 27, 3, 2, 2, 2, 164, 169, 7, 56, 2, 2, 165, 166,
	7, 39, 2, 2, 166, 167, 7, 142, 2, 2, 167, 169, 7, 141, 2, 2, 168, 164,
	3, 2, 2, 2, 168, 165, 3, 2, 2, 2, 169, 29, 3, 2, 2, 2, 170, 175, 7, 55,
	2, 2, 171, 172, 7, 19, 2, 2, 172, 174, 7, 55, 2, 2, 173, 171, 3, 2, 2,
	2, 174, 177, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176,
	31, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 178, 179, 5, 44, 23, 2, 179, 180,
	5, 34, 18, 2, 180, 33, 3, 2, 2, 2, 181, 189, 7, 10, 2, 2, 182, 183, 5,
	64, 33, 2, 183, 184, 7, 18, 2, 2, 184, 188, 3, 2, 2, 2, 185, 188, 5, 4,
	3, 2, 186, 188, 5, 42, 22, 2, 187, 182, 3, 2, 2, 2, 187, 185, 3, 2, 2,
	2, 187, 186, 3, 2, 2, 2, 188, 191, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 189,
	190, 3, 2, 2, 2, 190, 193, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 192, 194,
	5, 64, 33, 2, 193, 192, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 195, 3,
	2, 2, 2, 195, 196, 7, 11, 2, 2, 196, 35, 3, 2, 2, 2, 197, 198, 5, 44, 23,
	2, 198, 207, 7, 8, 2, 2, 199, 204, 5, 40, 21, 2, 200, 201, 7, 18, 2, 2,
	201, 203, 5, 40, 21, 2, 202, 200, 3, 2, 2, 2, 203, 206, 3, 2, 2, 2, 204,
	202, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206, 204,
	3, 2, 2, 2, 207, 199, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 210, 3, 2,
	2, 2, 209, 211, 7, 6, 2, 2, 210, 209, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2,
	211, 212, 3, 2, 2, 2, 212, 214, 7, 9, 2, 2, 213, 215, 5, 38, 20, 2, 214,
	213, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 217,
	5, 34, 18, 2, 217, 37, 3, 2, 2, 2, 218, 219, 7, 52, 2, 2, 219, 220, 5,
	18, 10, 2, 220, 39, 3, 2, 2, 2, 221, 224, 5, 6, 4, 2, 222, 224, 5, 24,
	13, 2, 223, 221, 3, 2, 2, 2, 223, 222, 3, 2, 2, 2, 224, 41, 3, 2, 2, 2,
	225, 226, 5, 46, 24, 2, 226, 228, 7, 8, 2, 2, 227, 229, 5, 66, 34, 2, 228,
	227, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 232,
	7, 9, 2, 2, 231, 233, 7, 43, 2, 2, 232, 231, 3, 2, 2, 2, 232, 233, 3, 2,
	2, 2, 233, 234, 3, 2, 2, 2, 234, 235, 7, 18, 2, 2, 235, 43, 3, 2, 2, 2,
	236, 241, 5, 46, 24, 2, 237, 238, 7, 19, 2, 2, 238, 240, 5, 46, 24, 2,
	239, 237, 3, 2, 2, 2, 240, 243, 3, 2, 2, 2, 241, 239, 3, 2, 2, 2, 241,
	242, 3, 2, 2, 2, 242, 45, 3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 244, 246, 5,
	52, 27, 2, 245, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2, 247, 245, 3, 2,
	2, 2, 247, 248, 3, 2, 2, 2, 248, 254, 3, 2, 2, 2, 249, 250, 5, 54, 28,
	2, 250, 251, 5, 52, 27, 2, 251, 253, 3, 2, 2, 2, 252, 249, 3, 2, 2, 2,
	253, 256, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255,
	260, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 257, 259, 5, 48, 25, 2, 258, 257,
	3, 2, 2, 2, 259, 262, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 260, 261, 3, 2,
	2, 2, 261, 264, 3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 263, 265, 5, 50, 26,
	2, 264, 263, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 47, 3, 2, 2, 2, 266,
	267, 7, 12, 2, 2, 267, 271, 7, 55, 2, 2, 268, 269, 5, 56, 29, 2, 269, 270,
	9, 5, 2, 2, 270, 272, 3, 2, 2, 2, 271, 268, 3, 2, 2, 2, 271, 272, 3, 2,
	2, 2, 272, 273, 3, 2, 2, 2, 273, 274, 7, 13, 2, 2, 274, 49, 3, 2, 2, 2,
	275, 276, 9, 6, 2, 2, 276, 277, 7, 55, 2, 2, 277, 51, 3, 2, 2, 2, 278,
	286, 5, 58, 30, 2, 279, 280, 7, 24, 2, 2, 280, 286, 5, 58, 30, 2, 281,
	282, 7, 20, 2, 2, 282, 286, 5, 58, 30, 2, 283, 286, 7, 23, 2, 2, 284, 286,
	7, 27, 2, 2, 285, 278, 3, 2, 2, 2, 285, 279, 3, 2, 2, 2, 285, 281, 3, 2,
	2, 2, 285, 283, 3, 2, 2, 2, 285, 284, 3, 2, 2, 2, 286, 53, 3, 2, 2, 2,
	287, 288, 9, 7, 2, 2, 288, 55, 3, 2, 2, 2, 289, 290, 9, 8, 2, 2, 290, 57,
	3, 2, 2, 2, 291, 295, 7, 55, 2, 2, 292, 294, 5, 60, 31, 2, 293, 292, 3,
	2, 2, 2, 294, 297, 3, 2, 2, 2, 295, 293, 3, 2, 2, 2, 295, 296, 3, 2, 2,
	2, 296, 308, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 298, 299, 7, 7, 2, 2, 299,
	300, 5, 62, 32, 2, 300, 304, 7, 11, 2, 2, 301, 303, 5, 60, 31, 2, 302,
	301, 3, 2, 2, 2, 303, 306, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 304, 305,
	3, 2, 2, 2, 305, 308, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 307, 291, 3, 2,
	2, 2, 307, 298, 3, 2, 2, 2, 308, 59, 3, 2, 2, 2, 309, 310, 7, 144, 2, 2,
	310, 311, 5, 62, 32, 2, 311, 312, 7, 11, 2, 2, 312, 315, 3, 2, 2, 2, 313,
	315, 7, 145, 2, 2, 314, 309, 3, 2, 2, 2, 314, 313, 3, 2, 2, 2, 315, 61,
	3, 2, 2, 2, 316, 317, 9, 9, 2, 2, 317, 63, 3, 2, 2, 2, 318, 319, 5, 58,
	30, 2, 319, 320, 7, 17, 2, 2, 320, 321, 5, 66, 34, 2, 321, 65, 3, 2, 2,
	2, 322, 327, 5, 8, 5, 2, 323, 324, 7, 19, 2, 2, 324, 326, 5, 8, 5, 2, 325,
	323, 3, 2, 2, 2, 326, 329, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 327, 328,
	3, 2, 2, 2, 328, 67, 3, 2, 2, 2, 329, 327, 3, 2, 2, 2, 330, 331, 7, 39,
	2, 2, 331, 332, 7, 142, 2, 2, 332, 333, 7, 141, 2, 2, 333, 69, 3, 2, 2,
	2, 334, 336, 7, 57, 2, 2, 335, 337, 7, 5, 2, 2, 336, 335, 3, 2, 2, 2, 336,
	337, 3, 2, 2, 2, 337, 71, 3, 2, 2, 2, 39, 75, 84, 90, 95, 98, 110, 118,
	123, 132, 144, 151, 160, 168, 175, 187, 189, 193, 204, 207, 210, 214, 223,
	228, 232, 241, 247, 254, 260, 264, 271, 285, 295, 304, 307, 314, 327, 336,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'null'", "'in'", "", "'...'", "", "'('", "')'", "'{'", "'}'", "'['",
	"']'", "'>'", "'~'", "'<'", "':'", "';'", "','", "'.'", "'$'", "'@'", "'&'",
	"'#'", "'::'", "'+'", "'*'", "'/'", "'-'", "", "'=='", "'>='", "'<='",
	"'!='", "'='", "'|='", "'~='", "'url'", "", "'@import'", "'@media'", "':extend'",
	"'!important'", "'@arguments'", "'@rest'", "'reference'", "'inline'", "'less'",
	"'css'", "'once'", "'multiple'", "'when'", "'not'", "'and'", "", "", "",
	"", "", "", "", "", "'color'", "'convert'", "'data-uri'", "'default'",
	"'unit'", "'get-unit'", "'svg-gradient'", "'escape'", "'e'", "", "'replace'",
	"'length'", "'extract'", "'ceil'", "'floor'", "'percentage'", "'round'",
	"'sqrt'", "'abs'", "'sin'", "'asin'", "'cos'", "'acos'", "'tan'", "'atan'",
	"'pi'", "'pow'", "'mod'", "'min'", "'max'", "'isnumber'", "'isstring'",
	"'iscolor'", "'iskeyword'", "'isurl'", "'ispixel'", "'isem'", "'ispercentage'",
	"'isunit'", "'rgb'", "'rgba'", "'argb'", "'hsl'", "'hsla'", "'hsv'", "'hsva'",
	"'hue'", "'saturation'", "'lightness'", "'hsvhue'", "'hsvsaturation'",
	"'hsvvalue'", "'red'", "'green'", "'blue'", "'alpha'", "'luma'", "'luminance'",
	"'saturate'", "'desaturate'", "'lighten'", "'darken'", "'fadein'", "'fadeout'",
	"'fade'", "'spin'", "'mix'", "'greyscale'", "'contrast'", "'multiply'",
	"'screen'", "'overlay'", "'softlight'", "'hardlight'", "'difference'",
	"'exclusion'", "'average'", "'negation'",
}
var symbolicNames = []string{
	"", "NULL", "IN", "Unit", "Ellipsis", "InterpolationStart", "LPAREN", "RPAREN",
	"BlockStart", "BlockEnd", "LBRACK", "RBRACK", "GT", "TIL", "LT", "COLON",
	"SEMI", "COMMA", "DOT", "DOLLAR", "AT", "PARENTREF", "HASH", "COLONCOLON",
	"PLUS", "TIMES", "DIV", "MINUS", "PERC", "EQEQ", "GTEQ", "LTEQ", "NOTEQ",
	"EQ", "PIPE_EQ", "TILD_EQ", "URL", "UrlStart", "IMPORT", "MEDIA", "EXTEND",
	"IMPORTANT", "ARGUMENTS", "REST", "REFERENCE", "INLINE", "LESS", "CSS",
	"ONCE", "MULTIPLE", "WHEN", "NOT", "AND", "Identifier", "StringLiteral",
	"Number", "Color", "WS", "SL_COMMENT", "COMMENT", "FUNCTION_NAME", "COLOR",
	"CONVERT", "DATA_URI", "DEFAULT", "UNIT", "GET_UNIT", "SVG_GRADIENT", "ESCAPE",
	"E", "FORMAT", "REPLACE", "LENGTH", "EXTRACT", "CEIL", "FLOOR", "PERCENTAGE",
	"ROUND", "SQRT", "ABS", "SIN", "ASIN", "COS", "ACOS", "TAN", "ATAN", "PI",
	"POW", "MOD", "MIN", "MAX", "ISNUMBER", "ISSTRING", "ISCOLOR", "ISKEYWORD",
	"ISURL", "ISPIXEL", "ISEM", "ISPERCENTAGE", "ISUNIT", "RGB", "RGBA", "ARGB",
	"HSL", "HSLA", "HSV", "HSVA", "HUE", "SATURATION", "LIGHTNESS", "HSVHUE",
	"HSVSATURATION", "HSVVALUE", "RED", "GREEN", "BLUE", "ALPHA", "LUMA", "LUMINANCE",
	"SATURATE", "DESATURATE", "LIGHTEN", "DARKEN", "FADEIN", "FADEOUT", "FADE",
	"SPIN", "MIX", "GREYSCALE", "CONTRAST", "MULTIPLY", "SCREEN", "OVERLAY",
	"SOFTLIGHT", "HARDLIGHT", "DIFFERENCE", "EXCLUSION", "AVERAGE", "NEGATION",
	"UrlEnd", "Url", "SPACE", "InterpolationStartAfter", "IdentifierAfter",
}

var ruleNames = []string{
	"stylesheet", "statement", "variableName", "commandStatement", "mathCharacter",
	"mathStatement", "expression", "function", "conditions", "condition", "conditionStatement",
	"variableDeclaration", "importDeclaration", "referenceUrl", "mediaTypes",
	"ruleset", "block", "mixinDefinition", "mixinGuard", "mixinDefinitionParam",
	"mixinReference", "selectors", "selector", "attrib", "pseudo", "element",
	"selectorPrefix", "attribRelate", "identifier", "identifierPart", "identifierVariableName",
	"property", "values", "url", "measurement",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type LessParser struct {
	*antlr.BaseParser
}

func NewLessParser(input antlr.TokenStream) *LessParser {
	this := new(LessParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "LessParser.g4"

	return this
}

// LessParser tokens.
const (
	LessParserEOF                     = antlr.TokenEOF
	LessParserNULL                    = 1
	LessParserIN                      = 2
	LessParserUnit                    = 3
	LessParserEllipsis                = 4
	LessParserInterpolationStart      = 5
	LessParserLPAREN                  = 6
	LessParserRPAREN                  = 7
	LessParserBlockStart              = 8
	LessParserBlockEnd                = 9
	LessParserLBRACK                  = 10
	LessParserRBRACK                  = 11
	LessParserGT                      = 12
	LessParserTIL                     = 13
	LessParserLT                      = 14
	LessParserCOLON                   = 15
	LessParserSEMI                    = 16
	LessParserCOMMA                   = 17
	LessParserDOT                     = 18
	LessParserDOLLAR                  = 19
	LessParserAT                      = 20
	LessParserPARENTREF               = 21
	LessParserHASH                    = 22
	LessParserCOLONCOLON              = 23
	LessParserPLUS                    = 24
	LessParserTIMES                   = 25
	LessParserDIV                     = 26
	LessParserMINUS                   = 27
	LessParserPERC                    = 28
	LessParserEQEQ                    = 29
	LessParserGTEQ                    = 30
	LessParserLTEQ                    = 31
	LessParserNOTEQ                   = 32
	LessParserEQ                      = 33
	LessParserPIPE_EQ                 = 34
	LessParserTILD_EQ                 = 35
	LessParserURL                     = 36
	LessParserUrlStart                = 37
	LessParserIMPORT                  = 38
	LessParserMEDIA                   = 39
	LessParserEXTEND                  = 40
	LessParserIMPORTANT               = 41
	LessParserARGUMENTS               = 42
	LessParserREST                    = 43
	LessParserREFERENCE               = 44
	LessParserINLINE                  = 45
	LessParserLESS                    = 46
	LessParserCSS                     = 47
	LessParserONCE                    = 48
	LessParserMULTIPLE                = 49
	LessParserWHEN                    = 50
	LessParserNOT                     = 51
	LessParserAND                     = 52
	LessParserIdentifier              = 53
	LessParserStringLiteral           = 54
	LessParserNumber                  = 55
	LessParserColor                   = 56
	LessParserWS                      = 57
	LessParserSL_COMMENT              = 58
	LessParserCOMMENT                 = 59
	LessParserFUNCTION_NAME           = 60
	LessParserCOLOR                   = 61
	LessParserCONVERT                 = 62
	LessParserDATA_URI                = 63
	LessParserDEFAULT                 = 64
	LessParserUNIT                    = 65
	LessParserGET_UNIT                = 66
	LessParserSVG_GRADIENT            = 67
	LessParserESCAPE                  = 68
	LessParserE                       = 69
	LessParserFORMAT                  = 70
	LessParserREPLACE                 = 71
	LessParserLENGTH                  = 72
	LessParserEXTRACT                 = 73
	LessParserCEIL                    = 74
	LessParserFLOOR                   = 75
	LessParserPERCENTAGE              = 76
	LessParserROUND                   = 77
	LessParserSQRT                    = 78
	LessParserABS                     = 79
	LessParserSIN                     = 80
	LessParserASIN                    = 81
	LessParserCOS                     = 82
	LessParserACOS                    = 83
	LessParserTAN                     = 84
	LessParserATAN                    = 85
	LessParserPI                      = 86
	LessParserPOW                     = 87
	LessParserMOD                     = 88
	LessParserMIN                     = 89
	LessParserMAX                     = 90
	LessParserISNUMBER                = 91
	LessParserISSTRING                = 92
	LessParserISCOLOR                 = 93
	LessParserISKEYWORD               = 94
	LessParserISURL                   = 95
	LessParserISPIXEL                 = 96
	LessParserISEM                    = 97
	LessParserISPERCENTAGE            = 98
	LessParserISUNIT                  = 99
	LessParserRGB                     = 100
	LessParserRGBA                    = 101
	LessParserARGB                    = 102
	LessParserHSL                     = 103
	LessParserHSLA                    = 104
	LessParserHSV                     = 105
	LessParserHSVA                    = 106
	LessParserHUE                     = 107
	LessParserSATURATION              = 108
	LessParserLIGHTNESS               = 109
	LessParserHSVHUE                  = 110
	LessParserHSVSATURATION           = 111
	LessParserHSVVALUE                = 112
	LessParserRED                     = 113
	LessParserGREEN                   = 114
	LessParserBLUE                    = 115
	LessParserALPHA                   = 116
	LessParserLUMA                    = 117
	LessParserLUMINANCE               = 118
	LessParserSATURATE                = 119
	LessParserDESATURATE              = 120
	LessParserLIGHTEN                 = 121
	LessParserDARKEN                  = 122
	LessParserFADEIN                  = 123
	LessParserFADEOUT                 = 124
	LessParserFADE                    = 125
	LessParserSPIN                    = 126
	LessParserMIX                     = 127
	LessParserGREYSCALE               = 128
	LessParserCONTRAST                = 129
	LessParserMULTIPLY                = 130
	LessParserSCREEN                  = 131
	LessParserOVERLAY                 = 132
	LessParserSOFTLIGHT               = 133
	LessParserHARDLIGHT               = 134
	LessParserDIFFERENCE              = 135
	LessParserEXCLUSION               = 136
	LessParserAVERAGE                 = 137
	LessParserNEGATION                = 138
	LessParserUrlEnd                  = 139
	LessParserUrl                     = 140
	LessParserSPACE                   = 141
	LessParserInterpolationStartAfter = 142
	LessParserIdentifierAfter         = 143
)

// LessParser rules.
const (
	LessParserRULE_stylesheet             = 0
	LessParserRULE_statement              = 1
	LessParserRULE_variableName           = 2
	LessParserRULE_commandStatement       = 3
	LessParserRULE_mathCharacter          = 4
	LessParserRULE_mathStatement          = 5
	LessParserRULE_expression             = 6
	LessParserRULE_function               = 7
	LessParserRULE_conditions             = 8
	LessParserRULE_condition              = 9
	LessParserRULE_conditionStatement     = 10
	LessParserRULE_variableDeclaration    = 11
	LessParserRULE_importDeclaration      = 12
	LessParserRULE_referenceUrl           = 13
	LessParserRULE_mediaTypes             = 14
	LessParserRULE_ruleset                = 15
	LessParserRULE_block                  = 16
	LessParserRULE_mixinDefinition        = 17
	LessParserRULE_mixinGuard             = 18
	LessParserRULE_mixinDefinitionParam   = 19
	LessParserRULE_mixinReference         = 20
	LessParserRULE_selectors              = 21
	LessParserRULE_selector               = 22
	LessParserRULE_attrib                 = 23
	LessParserRULE_pseudo                 = 24
	LessParserRULE_element                = 25
	LessParserRULE_selectorPrefix         = 26
	LessParserRULE_attribRelate           = 27
	LessParserRULE_identifier             = 28
	LessParserRULE_identifierPart         = 29
	LessParserRULE_identifierVariableName = 30
	LessParserRULE_property               = 31
	LessParserRULE_values                 = 32
	LessParserRULE_url                    = 33
	LessParserRULE_measurement            = 34
)

// IStylesheetContext is an interface to support dynamic dispatch.
type IStylesheetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStylesheetContext differentiates from other interfaces.
	IsStylesheetContext()
}

type StylesheetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStylesheetContext() *StylesheetContext {
	var p = new(StylesheetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_stylesheet
	return p
}

func (*StylesheetContext) IsStylesheetContext() {}

func NewStylesheetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StylesheetContext {
	var p = new(StylesheetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_stylesheet

	return p
}

func (s *StylesheetContext) GetParser() antlr.Parser { return s.parser }

func (s *StylesheetContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StylesheetContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StylesheetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StylesheetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StylesheetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterStylesheet(s)
	}
}

func (s *StylesheetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitStylesheet(s)
	}
}

func (p *LessParser) Stylesheet() (localctx IStylesheetContext) {
	localctx = NewStylesheetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LessParserRULE_stylesheet)
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
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LessParserInterpolationStart)|(1<<LessParserDOT)|(1<<LessParserAT)|(1<<LessParserPARENTREF)|(1<<LessParserHASH)|(1<<LessParserTIMES))) != 0) || _la == LessParserIMPORT || _la == LessParserIdentifier {
		{
			p.SetState(70)
			p.Statement()
		}

		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ImportDeclaration() IImportDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportDeclarationContext)
}

func (s *StatementContext) Ruleset() IRulesetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRulesetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRulesetContext)
}

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) MixinDefinition() IMixinDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMixinDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMixinDefinitionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *LessParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LessParserRULE_statement)

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

	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(76)
			p.ImportDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Ruleset()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.VariableDeclaration()
		}
		{
			p.SetState(79)
			p.Match(LessParserSEMI)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(81)
			p.MixinDefinition()
		}

	}

	return localctx
}

// IVariableNameContext is an interface to support dynamic dispatch.
type IVariableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableNameContext differentiates from other interfaces.
	IsVariableNameContext()
}

type VariableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableNameContext() *VariableNameContext {
	var p = new(VariableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_variableName
	return p
}

func (*VariableNameContext) IsVariableNameContext() {}

func NewVariableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableNameContext {
	var p = new(VariableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_variableName

	return p
}

func (s *VariableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableNameContext) AT() antlr.TerminalNode {
	return s.GetToken(LessParserAT, 0)
}

func (s *VariableNameContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *VariableNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LessParserIdentifier, 0)
}

func (s *VariableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterVariableName(s)
	}
}

func (s *VariableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitVariableName(s)
	}
}

func (p *LessParser) VariableName() (localctx IVariableNameContext) {
	localctx = NewVariableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LessParserRULE_variableName)

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

	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(84)
			p.Match(LessParserAT)
		}
		{
			p.SetState(85)
			p.VariableName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(86)
			p.Match(LessParserAT)
		}
		{
			p.SetState(87)
			p.Match(LessParserIdentifier)
		}

	}

	return localctx
}

// ICommandStatementContext is an interface to support dynamic dispatch.
type ICommandStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandStatementContext differentiates from other interfaces.
	IsCommandStatementContext()
}

type CommandStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandStatementContext() *CommandStatementContext {
	var p = new(CommandStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_commandStatement
	return p
}

func (*CommandStatementContext) IsCommandStatementContext() {}

func NewCommandStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandStatementContext {
	var p = new(CommandStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_commandStatement

	return p
}

func (s *CommandStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandStatementContext) MathStatement() IMathStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathStatementContext)
}

func (s *CommandStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CommandStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CommandStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterCommandStatement(s)
	}
}

func (s *CommandStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitCommandStatement(s)
	}
}

func (p *LessParser) CommandStatement() (localctx ICommandStatementContext) {
	localctx = NewCommandStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LessParserRULE_commandStatement)
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
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LessParserInterpolationStart || _la == LessParserAT || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(LessParserUrlStart-37))|(1<<(LessParserIdentifier-37))|(1<<(LessParserStringLiteral-37))|(1<<(LessParserNumber-37))|(1<<(LessParserColor-37)))) != 0) {
		{
			p.SetState(90)
			p.Expression()
		}

		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LessParserPLUS)|(1<<LessParserTIMES)|(1<<LessParserDIV)|(1<<LessParserMINUS)|(1<<LessParserPERC))) != 0 {
		{
			p.SetState(95)
			p.MathStatement()
		}

	}

	return localctx
}

// IMathCharacterContext is an interface to support dynamic dispatch.
type IMathCharacterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathCharacterContext differentiates from other interfaces.
	IsMathCharacterContext()
}

type MathCharacterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathCharacterContext() *MathCharacterContext {
	var p = new(MathCharacterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mathCharacter
	return p
}

func (*MathCharacterContext) IsMathCharacterContext() {}

func NewMathCharacterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathCharacterContext {
	var p = new(MathCharacterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mathCharacter

	return p
}

func (s *MathCharacterContext) GetParser() antlr.Parser { return s.parser }

func (s *MathCharacterContext) TIMES() antlr.TerminalNode {
	return s.GetToken(LessParserTIMES, 0)
}

func (s *MathCharacterContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LessParserPLUS, 0)
}

func (s *MathCharacterContext) DIV() antlr.TerminalNode {
	return s.GetToken(LessParserDIV, 0)
}

func (s *MathCharacterContext) MINUS() antlr.TerminalNode {
	return s.GetToken(LessParserMINUS, 0)
}

func (s *MathCharacterContext) PERC() antlr.TerminalNode {
	return s.GetToken(LessParserPERC, 0)
}

func (s *MathCharacterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathCharacterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathCharacterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMathCharacter(s)
	}
}

func (s *MathCharacterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMathCharacter(s)
	}
}

func (p *LessParser) MathCharacter() (localctx IMathCharacterContext) {
	localctx = NewMathCharacterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LessParserRULE_mathCharacter)
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
	p.SetState(98)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LessParserPLUS)|(1<<LessParserTIMES)|(1<<LessParserDIV)|(1<<LessParserMINUS)|(1<<LessParserPERC))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IMathStatementContext is an interface to support dynamic dispatch.
type IMathStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathStatementContext differentiates from other interfaces.
	IsMathStatementContext()
}

type MathStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathStatementContext() *MathStatementContext {
	var p = new(MathStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mathStatement
	return p
}

func (*MathStatementContext) IsMathStatementContext() {}

func NewMathStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathStatementContext {
	var p = new(MathStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mathStatement

	return p
}

func (s *MathStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MathStatementContext) MathCharacter() IMathCharacterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathCharacterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathCharacterContext)
}

func (s *MathStatementContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *MathStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMathStatement(s)
	}
}

func (s *MathStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMathStatement(s)
	}
}

func (p *LessParser) MathStatement() (localctx IMathStatementContext) {
	localctx = NewMathStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LessParserRULE_mathStatement)

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
		p.SetState(100)
		p.MathCharacter()
	}
	{
		p.SetState(101)
		p.CommandStatement()
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
	p.RuleIndex = LessParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Measurement() IMeasurementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeasurementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeasurementContext)
}

func (s *ExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserLPAREN, 0)
}

func (s *ExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserRPAREN, 0)
}

func (s *ExpressionContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *ExpressionContext) Color() antlr.TerminalNode {
	return s.GetToken(LessParserColor, 0)
}

func (s *ExpressionContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(LessParserStringLiteral, 0)
}

func (s *ExpressionContext) Url() IUrlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUrlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUrlContext)
}

func (s *ExpressionContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LessParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LessParserRULE_expression)
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

	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)
			p.Measurement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Identifier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(105)
			p.Identifier()
		}
		{
			p.SetState(106)
			p.Match(LessParserLPAREN)
		}
		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LessParserInterpolationStart || _la == LessParserAT || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(LessParserUrlStart-37))|(1<<(LessParserIdentifier-37))|(1<<(LessParserStringLiteral-37))|(1<<(LessParserNumber-37))|(1<<(LessParserColor-37)))) != 0) {
			{
				p.SetState(107)
				p.Values()
			}

		}
		{
			p.SetState(110)
			p.Match(LessParserRPAREN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(112)
			p.Match(LessParserColor)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(113)
			p.Match(LessParserStringLiteral)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(114)
			p.Url()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(115)
			p.VariableName()
		}

	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FUNCTION_NAME() antlr.TerminalNode {
	return s.GetToken(LessParserFUNCTION_NAME, 0)
}

func (s *FunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserLPAREN, 0)
}

func (s *FunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserRPAREN, 0)
}

func (s *FunctionContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *LessParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LessParserRULE_function)
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
		p.SetState(118)
		p.Match(LessParserFUNCTION_NAME)
	}
	{
		p.SetState(119)
		p.Match(LessParserLPAREN)
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserInterpolationStart || _la == LessParserAT || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(LessParserUrlStart-37))|(1<<(LessParserIdentifier-37))|(1<<(LessParserStringLiteral-37))|(1<<(LessParserNumber-37))|(1<<(LessParserColor-37)))) != 0) {
		{
			p.SetState(120)
			p.Values()
		}

	}
	{
		p.SetState(123)
		p.Match(LessParserRPAREN)
	}

	return localctx
}

// IConditionsContext is an interface to support dynamic dispatch.
type IConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionsContext differentiates from other interfaces.
	IsConditionsContext()
}

type ConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsContext() *ConditionsContext {
	var p = new(ConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_conditions
	return p
}

func (*ConditionsContext) IsConditionsContext() {}

func NewConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsContext {
	var p = new(ConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_conditions

	return p
}

func (s *ConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsContext) AllCondition() []IConditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionContext)(nil)).Elem())
	var tst = make([]IConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionContext)
		}
	}

	return tst
}

func (s *ConditionsContext) Condition(i int) IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionsContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(LessParserAND)
}

func (s *ConditionsContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(LessParserAND, i)
}

func (s *ConditionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LessParserCOMMA)
}

func (s *ConditionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LessParserCOMMA, i)
}

func (s *ConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterConditions(s)
	}
}

func (s *ConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitConditions(s)
	}
}

func (p *LessParser) Conditions() (localctx IConditionsContext) {
	localctx = NewConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LessParserRULE_conditions)
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
		p.Condition()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LessParserCOMMA || _la == LessParserAND {
		p.SetState(126)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LessParserCOMMA || _la == LessParserAND) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(127)
			p.Condition()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserLPAREN, 0)
}

func (s *ConditionContext) ConditionStatement() IConditionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionStatementContext)
}

func (s *ConditionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserRPAREN, 0)
}

func (s *ConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LessParserNOT, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *LessParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LessParserRULE_condition)

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

	switch p.GetTokenStream().LA(1) {
	case LessParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)
			p.Match(LessParserLPAREN)
		}
		{
			p.SetState(134)
			p.ConditionStatement()
		}
		{
			p.SetState(135)
			p.Match(LessParserRPAREN)
		}

	case LessParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Match(LessParserNOT)
		}
		{
			p.SetState(138)
			p.Match(LessParserLPAREN)
		}
		{
			p.SetState(139)
			p.ConditionStatement()
		}
		{
			p.SetState(140)
			p.Match(LessParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConditionStatementContext is an interface to support dynamic dispatch.
type IConditionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionStatementContext differentiates from other interfaces.
	IsConditionStatementContext()
}

type ConditionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionStatementContext() *ConditionStatementContext {
	var p = new(ConditionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_conditionStatement
	return p
}

func (*ConditionStatementContext) IsConditionStatementContext() {}

func NewConditionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionStatementContext {
	var p = new(ConditionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_conditionStatement

	return p
}

func (s *ConditionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionStatementContext) AllCommandStatement() []ICommandStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem())
	var tst = make([]ICommandStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandStatementContext)
		}
	}

	return tst
}

func (s *ConditionStatementContext) CommandStatement(i int) ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *ConditionStatementContext) EQ() antlr.TerminalNode {
	return s.GetToken(LessParserEQ, 0)
}

func (s *ConditionStatementContext) LT() antlr.TerminalNode {
	return s.GetToken(LessParserLT, 0)
}

func (s *ConditionStatementContext) GT() antlr.TerminalNode {
	return s.GetToken(LessParserGT, 0)
}

func (s *ConditionStatementContext) GTEQ() antlr.TerminalNode {
	return s.GetToken(LessParserGTEQ, 0)
}

func (s *ConditionStatementContext) LTEQ() antlr.TerminalNode {
	return s.GetToken(LessParserLTEQ, 0)
}

func (s *ConditionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterConditionStatement(s)
	}
}

func (s *ConditionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitConditionStatement(s)
	}
}

func (p *LessParser) ConditionStatement() (localctx IConditionStatementContext) {
	localctx = NewConditionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LessParserRULE_conditionStatement)
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

	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.CommandStatement()
		}
		p.SetState(145)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-12)&-(0x1f+1)) == 0 && ((1<<uint((_la-12)))&((1<<(LessParserGT-12))|(1<<(LessParserLT-12))|(1<<(LessParserGTEQ-12))|(1<<(LessParserLTEQ-12))|(1<<(LessParserEQ-12)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(146)
			p.CommandStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.CommandStatement()
		}

	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *VariableDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(LessParserCOLON, 0)
}

func (s *VariableDeclarationContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *LessParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LessParserRULE_variableDeclaration)

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
		p.SetState(151)
		p.VariableName()
	}
	{
		p.SetState(152)
		p.Match(LessParserCOLON)
	}
	{
		p.SetState(153)
		p.Values()
	}

	return localctx
}

// IImportDeclarationContext is an interface to support dynamic dispatch.
type IImportDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclarationContext differentiates from other interfaces.
	IsImportDeclarationContext()
}

type ImportDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclarationContext() *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_importDeclaration
	return p
}

func (*ImportDeclarationContext) IsImportDeclarationContext() {}

func NewImportDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_importDeclaration

	return p
}

func (s *ImportDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclarationContext) ReferenceUrl() IReferenceUrlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceUrlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferenceUrlContext)
}

func (s *ImportDeclarationContext) MediaTypes() IMediaTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMediaTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMediaTypesContext)
}

func (s *ImportDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitImportDeclaration(s)
	}
}

func (p *LessParser) ImportDeclaration() (localctx IImportDeclarationContext) {
	localctx = NewImportDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LessParserRULE_importDeclaration)
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
		p.SetState(155)
		p.Match(LessParserIMPORT)
	}
	{
		p.SetState(156)
		p.ReferenceUrl()
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserIdentifier {
		{
			p.SetState(157)
			p.MediaTypes()
		}

	}
	{
		p.SetState(160)
		p.Match(LessParserSEMI)
	}

	return localctx
}

// IReferenceUrlContext is an interface to support dynamic dispatch.
type IReferenceUrlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferenceUrlContext differentiates from other interfaces.
	IsReferenceUrlContext()
}

type ReferenceUrlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceUrlContext() *ReferenceUrlContext {
	var p = new(ReferenceUrlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_referenceUrl
	return p
}

func (*ReferenceUrlContext) IsReferenceUrlContext() {}

func NewReferenceUrlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceUrlContext {
	var p = new(ReferenceUrlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_referenceUrl

	return p
}

func (s *ReferenceUrlContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceUrlContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(LessParserStringLiteral, 0)
}

func (s *ReferenceUrlContext) UrlStart() antlr.TerminalNode {
	return s.GetToken(LessParserUrlStart, 0)
}

func (s *ReferenceUrlContext) Url() antlr.TerminalNode {
	return s.GetToken(LessParserUrl, 0)
}

func (s *ReferenceUrlContext) UrlEnd() antlr.TerminalNode {
	return s.GetToken(LessParserUrlEnd, 0)
}

func (s *ReferenceUrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceUrlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceUrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterReferenceUrl(s)
	}
}

func (s *ReferenceUrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitReferenceUrl(s)
	}
}

func (p *LessParser) ReferenceUrl() (localctx IReferenceUrlContext) {
	localctx = NewReferenceUrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LessParserRULE_referenceUrl)

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

	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LessParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(162)
			p.Match(LessParserStringLiteral)
		}

	case LessParserUrlStart:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.Match(LessParserUrlStart)
		}
		{
			p.SetState(164)
			p.Match(LessParserUrl)
		}
		{
			p.SetState(165)
			p.Match(LessParserUrlEnd)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMediaTypesContext is an interface to support dynamic dispatch.
type IMediaTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMediaTypesContext differentiates from other interfaces.
	IsMediaTypesContext()
}

type MediaTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMediaTypesContext() *MediaTypesContext {
	var p = new(MediaTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mediaTypes
	return p
}

func (*MediaTypesContext) IsMediaTypesContext() {}

func NewMediaTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MediaTypesContext {
	var p = new(MediaTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mediaTypes

	return p
}

func (s *MediaTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *MediaTypesContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(LessParserIdentifier)
}

func (s *MediaTypesContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(LessParserIdentifier, i)
}

func (s *MediaTypesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LessParserCOMMA)
}

func (s *MediaTypesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LessParserCOMMA, i)
}

func (s *MediaTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MediaTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MediaTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMediaTypes(s)
	}
}

func (s *MediaTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMediaTypes(s)
	}
}

func (p *LessParser) MediaTypes() (localctx IMediaTypesContext) {
	localctx = NewMediaTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LessParserRULE_mediaTypes)
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
		p.SetState(168)
		p.Match(LessParserIdentifier)
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LessParserCOMMA {
		{
			p.SetState(169)
			p.Match(LessParserCOMMA)
		}
		{
			p.SetState(170)
			p.Match(LessParserIdentifier)
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRulesetContext is an interface to support dynamic dispatch.
type IRulesetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulesetContext differentiates from other interfaces.
	IsRulesetContext()
}

type RulesetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulesetContext() *RulesetContext {
	var p = new(RulesetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_ruleset
	return p
}

func (*RulesetContext) IsRulesetContext() {}

func NewRulesetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulesetContext {
	var p = new(RulesetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_ruleset

	return p
}

func (s *RulesetContext) GetParser() antlr.Parser { return s.parser }

func (s *RulesetContext) Selectors() ISelectorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorsContext)
}

func (s *RulesetContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RulesetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulesetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulesetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterRuleset(s)
	}
}

func (s *RulesetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitRuleset(s)
	}
}

func (p *LessParser) Ruleset() (localctx IRulesetContext) {
	localctx = NewRulesetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LessParserRULE_ruleset)

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
		p.SetState(176)
		p.Selectors()
	}
	{
		p.SetState(177)
		p.Block()
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) BlockStart() antlr.TerminalNode {
	return s.GetToken(LessParserBlockStart, 0)
}

func (s *BlockContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(LessParserBlockEnd, 0)
}

func (s *BlockContext) AllProperty() []IPropertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyContext)(nil)).Elem())
	var tst = make([]IPropertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyContext)
		}
	}

	return tst
}

func (s *BlockContext) Property(i int) IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) AllMixinReference() []IMixinReferenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMixinReferenceContext)(nil)).Elem())
	var tst = make([]IMixinReferenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMixinReferenceContext)
		}
	}

	return tst
}

func (s *BlockContext) MixinReference(i int) IMixinReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMixinReferenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMixinReferenceContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *LessParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LessParserRULE_block)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(LessParserBlockStart)
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(185)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(180)
					p.Property()
				}
				{
					p.SetState(181)
					p.Match(LessParserSEMI)
				}

			case 2:
				{
					p.SetState(183)
					p.Statement()
				}

			case 3:
				{
					p.SetState(184)
					p.MixinReference()
				}

			}

		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserInterpolationStart || _la == LessParserIdentifier {
		{
			p.SetState(190)
			p.Property()
		}

	}
	{
		p.SetState(193)
		p.Match(LessParserBlockEnd)
	}

	return localctx
}

// IMixinDefinitionContext is an interface to support dynamic dispatch.
type IMixinDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMixinDefinitionContext differentiates from other interfaces.
	IsMixinDefinitionContext()
}

type MixinDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixinDefinitionContext() *MixinDefinitionContext {
	var p = new(MixinDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mixinDefinition
	return p
}

func (*MixinDefinitionContext) IsMixinDefinitionContext() {}

func NewMixinDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MixinDefinitionContext {
	var p = new(MixinDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mixinDefinition

	return p
}

func (s *MixinDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MixinDefinitionContext) Selectors() ISelectorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorsContext)
}

func (s *MixinDefinitionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserLPAREN, 0)
}

func (s *MixinDefinitionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserRPAREN, 0)
}

func (s *MixinDefinitionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *MixinDefinitionContext) AllMixinDefinitionParam() []IMixinDefinitionParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMixinDefinitionParamContext)(nil)).Elem())
	var tst = make([]IMixinDefinitionParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMixinDefinitionParamContext)
		}
	}

	return tst
}

func (s *MixinDefinitionContext) MixinDefinitionParam(i int) IMixinDefinitionParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMixinDefinitionParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMixinDefinitionParamContext)
}

func (s *MixinDefinitionContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(LessParserEllipsis, 0)
}

func (s *MixinDefinitionContext) MixinGuard() IMixinGuardContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMixinGuardContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMixinGuardContext)
}

func (s *MixinDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MixinDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MixinDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMixinDefinition(s)
	}
}

func (s *MixinDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMixinDefinition(s)
	}
}

func (p *LessParser) MixinDefinition() (localctx IMixinDefinitionContext) {
	localctx = NewMixinDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LessParserRULE_mixinDefinition)
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
		p.SetState(195)
		p.Selectors()
	}
	{
		p.SetState(196)
		p.Match(LessParserLPAREN)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserAT {
		{
			p.SetState(197)
			p.MixinDefinitionParam()
		}
		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LessParserSEMI {
			{
				p.SetState(198)
				p.Match(LessParserSEMI)
			}
			{
				p.SetState(199)
				p.MixinDefinitionParam()
			}

			p.SetState(204)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserEllipsis {
		{
			p.SetState(207)
			p.Match(LessParserEllipsis)
		}

	}
	{
		p.SetState(210)
		p.Match(LessParserRPAREN)
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserWHEN {
		{
			p.SetState(211)
			p.MixinGuard()
		}

	}
	{
		p.SetState(214)
		p.Block()
	}

	return localctx
}

// IMixinGuardContext is an interface to support dynamic dispatch.
type IMixinGuardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMixinGuardContext differentiates from other interfaces.
	IsMixinGuardContext()
}

type MixinGuardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixinGuardContext() *MixinGuardContext {
	var p = new(MixinGuardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mixinGuard
	return p
}

func (*MixinGuardContext) IsMixinGuardContext() {}

func NewMixinGuardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MixinGuardContext {
	var p = new(MixinGuardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mixinGuard

	return p
}

func (s *MixinGuardContext) GetParser() antlr.Parser { return s.parser }

func (s *MixinGuardContext) WHEN() antlr.TerminalNode {
	return s.GetToken(LessParserWHEN, 0)
}

func (s *MixinGuardContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *MixinGuardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MixinGuardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MixinGuardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMixinGuard(s)
	}
}

func (s *MixinGuardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMixinGuard(s)
	}
}

func (p *LessParser) MixinGuard() (localctx IMixinGuardContext) {
	localctx = NewMixinGuardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LessParserRULE_mixinGuard)

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
		p.Match(LessParserWHEN)
	}
	{
		p.SetState(217)
		p.Conditions()
	}

	return localctx
}

// IMixinDefinitionParamContext is an interface to support dynamic dispatch.
type IMixinDefinitionParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMixinDefinitionParamContext differentiates from other interfaces.
	IsMixinDefinitionParamContext()
}

type MixinDefinitionParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixinDefinitionParamContext() *MixinDefinitionParamContext {
	var p = new(MixinDefinitionParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mixinDefinitionParam
	return p
}

func (*MixinDefinitionParamContext) IsMixinDefinitionParamContext() {}

func NewMixinDefinitionParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MixinDefinitionParamContext {
	var p = new(MixinDefinitionParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mixinDefinitionParam

	return p
}

func (s *MixinDefinitionParamContext) GetParser() antlr.Parser { return s.parser }

func (s *MixinDefinitionParamContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *MixinDefinitionParamContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *MixinDefinitionParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MixinDefinitionParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MixinDefinitionParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMixinDefinitionParam(s)
	}
}

func (s *MixinDefinitionParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMixinDefinitionParam(s)
	}
}

func (p *LessParser) MixinDefinitionParam() (localctx IMixinDefinitionParamContext) {
	localctx = NewMixinDefinitionParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LessParserRULE_mixinDefinitionParam)

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

	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.VariableName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.VariableDeclaration()
		}

	}

	return localctx
}

// IMixinReferenceContext is an interface to support dynamic dispatch.
type IMixinReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMixinReferenceContext differentiates from other interfaces.
	IsMixinReferenceContext()
}

type MixinReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixinReferenceContext() *MixinReferenceContext {
	var p = new(MixinReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mixinReference
	return p
}

func (*MixinReferenceContext) IsMixinReferenceContext() {}

func NewMixinReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MixinReferenceContext {
	var p = new(MixinReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mixinReference

	return p
}

func (s *MixinReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *MixinReferenceContext) Selector() ISelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *MixinReferenceContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserLPAREN, 0)
}

func (s *MixinReferenceContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserRPAREN, 0)
}

func (s *MixinReferenceContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *MixinReferenceContext) IMPORTANT() antlr.TerminalNode {
	return s.GetToken(LessParserIMPORTANT, 0)
}

func (s *MixinReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MixinReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MixinReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMixinReference(s)
	}
}

func (s *MixinReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMixinReference(s)
	}
}

func (p *LessParser) MixinReference() (localctx IMixinReferenceContext) {
	localctx = NewMixinReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LessParserRULE_mixinReference)
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
		p.SetState(223)
		p.Selector()
	}
	{
		p.SetState(224)
		p.Match(LessParserLPAREN)
	}
	p.SetState(226)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserInterpolationStart || _la == LessParserAT || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(LessParserUrlStart-37))|(1<<(LessParserIdentifier-37))|(1<<(LessParserStringLiteral-37))|(1<<(LessParserNumber-37))|(1<<(LessParserColor-37)))) != 0) {
		{
			p.SetState(225)
			p.Values()
		}

	}
	{
		p.SetState(228)
		p.Match(LessParserRPAREN)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserIMPORTANT {
		{
			p.SetState(229)
			p.Match(LessParserIMPORTANT)
		}

	}
	{
		p.SetState(232)
		p.Match(LessParserSEMI)
	}

	return localctx
}

// ISelectorsContext is an interface to support dynamic dispatch.
type ISelectorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorsContext differentiates from other interfaces.
	IsSelectorsContext()
}

type SelectorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorsContext() *SelectorsContext {
	var p = new(SelectorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_selectors
	return p
}

func (*SelectorsContext) IsSelectorsContext() {}

func NewSelectorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorsContext {
	var p = new(SelectorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_selectors

	return p
}

func (s *SelectorsContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorsContext) AllSelector() []ISelectorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectorContext)(nil)).Elem())
	var tst = make([]ISelectorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectorContext)
		}
	}

	return tst
}

func (s *SelectorsContext) Selector(i int) ISelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *SelectorsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LessParserCOMMA)
}

func (s *SelectorsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LessParserCOMMA, i)
}

func (s *SelectorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterSelectors(s)
	}
}

func (s *SelectorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitSelectors(s)
	}
}

func (p *LessParser) Selectors() (localctx ISelectorsContext) {
	localctx = NewSelectorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LessParserRULE_selectors)
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
		p.SetState(234)
		p.Selector()
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LessParserCOMMA {
		{
			p.SetState(235)
			p.Match(LessParserCOMMA)
		}
		{
			p.SetState(236)
			p.Selector()
		}

		p.SetState(241)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_selector
	return p
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *SelectorContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *SelectorContext) AllSelectorPrefix() []ISelectorPrefixContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectorPrefixContext)(nil)).Elem())
	var tst = make([]ISelectorPrefixContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectorPrefixContext)
		}
	}

	return tst
}

func (s *SelectorContext) SelectorPrefix(i int) ISelectorPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorPrefixContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectorPrefixContext)
}

func (s *SelectorContext) AllAttrib() []IAttribContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttribContext)(nil)).Elem())
	var tst = make([]IAttribContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttribContext)
		}
	}

	return tst
}

func (s *SelectorContext) Attrib(i int) IAttribContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttribContext)
}

func (s *SelectorContext) Pseudo() IPseudoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPseudoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPseudoContext)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (p *LessParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LessParserRULE_selector)
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
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LessParserInterpolationStart)|(1<<LessParserDOT)|(1<<LessParserPARENTREF)|(1<<LessParserHASH)|(1<<LessParserTIMES))) != 0) || _la == LessParserIdentifier {
		{
			p.SetState(242)
			p.Element()
		}

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(252)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LessParserGT)|(1<<LessParserTIL)|(1<<LessParserPLUS))) != 0 {
		{
			p.SetState(247)
			p.SelectorPrefix()
		}
		{
			p.SetState(248)
			p.Element()
		}

		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LessParserLBRACK {
		{
			p.SetState(255)
			p.Attrib()
		}

		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserCOLON || _la == LessParserCOLONCOLON {
		{
			p.SetState(261)
			p.Pseudo()
		}

	}

	return localctx
}

// IAttribContext is an interface to support dynamic dispatch.
type IAttribContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttribContext differentiates from other interfaces.
	IsAttribContext()
}

type AttribContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribContext() *AttribContext {
	var p = new(AttribContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_attrib
	return p
}

func (*AttribContext) IsAttribContext() {}

func NewAttribContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttribContext {
	var p = new(AttribContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_attrib

	return p
}

func (s *AttribContext) GetParser() antlr.Parser { return s.parser }

func (s *AttribContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(LessParserIdentifier)
}

func (s *AttribContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(LessParserIdentifier, i)
}

func (s *AttribContext) AttribRelate() IAttribRelateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribRelateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttribRelateContext)
}

func (s *AttribContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(LessParserStringLiteral, 0)
}

func (s *AttribContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttribContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterAttrib(s)
	}
}

func (s *AttribContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitAttrib(s)
	}
}

func (p *LessParser) Attrib() (localctx IAttribContext) {
	localctx = NewAttribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LessParserRULE_attrib)
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
		p.SetState(264)
		p.Match(LessParserLBRACK)
	}
	{
		p.SetState(265)
		p.Match(LessParserIdentifier)
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(LessParserEQ-33))|(1<<(LessParserPIPE_EQ-33))|(1<<(LessParserTILD_EQ-33)))) != 0 {
		{
			p.SetState(266)
			p.AttribRelate()
		}
		p.SetState(267)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LessParserIdentifier || _la == LessParserStringLiteral) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}
	{
		p.SetState(271)
		p.Match(LessParserRBRACK)
	}

	return localctx
}

// IPseudoContext is an interface to support dynamic dispatch.
type IPseudoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPseudoContext differentiates from other interfaces.
	IsPseudoContext()
}

type PseudoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPseudoContext() *PseudoContext {
	var p = new(PseudoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_pseudo
	return p
}

func (*PseudoContext) IsPseudoContext() {}

func NewPseudoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PseudoContext {
	var p = new(PseudoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_pseudo

	return p
}

func (s *PseudoContext) GetParser() antlr.Parser { return s.parser }

func (s *PseudoContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LessParserIdentifier, 0)
}

func (s *PseudoContext) COLON() antlr.TerminalNode {
	return s.GetToken(LessParserCOLON, 0)
}

func (s *PseudoContext) COLONCOLON() antlr.TerminalNode {
	return s.GetToken(LessParserCOLONCOLON, 0)
}

func (s *PseudoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PseudoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PseudoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterPseudo(s)
	}
}

func (s *PseudoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitPseudo(s)
	}
}

func (p *LessParser) Pseudo() (localctx IPseudoContext) {
	localctx = NewPseudoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LessParserRULE_pseudo)
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
	p.SetState(273)
	_la = p.GetTokenStream().LA(1)

	if !(_la == LessParserCOLON || _la == LessParserCOLONCOLON) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(274)
		p.Match(LessParserIdentifier)
	}

	return localctx
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_element
	return p
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ElementContext) PARENTREF() antlr.TerminalNode {
	return s.GetToken(LessParserPARENTREF, 0)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitElement(s)
	}
}

func (p *LessParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LessParserRULE_element)

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

	p.SetState(283)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LessParserInterpolationStart, LessParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.Identifier()
		}

	case LessParserHASH:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.Match(LessParserHASH)
		}
		{
			p.SetState(278)
			p.Identifier()
		}

	case LessParserDOT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(279)
			p.Match(LessParserDOT)
		}
		{
			p.SetState(280)
			p.Identifier()
		}

	case LessParserPARENTREF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(281)
			p.Match(LessParserPARENTREF)
		}

	case LessParserTIMES:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(282)
			p.Match(LessParserTIMES)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISelectorPrefixContext is an interface to support dynamic dispatch.
type ISelectorPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorPrefixContext differentiates from other interfaces.
	IsSelectorPrefixContext()
}

type SelectorPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorPrefixContext() *SelectorPrefixContext {
	var p = new(SelectorPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_selectorPrefix
	return p
}

func (*SelectorPrefixContext) IsSelectorPrefixContext() {}

func NewSelectorPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorPrefixContext {
	var p = new(SelectorPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_selectorPrefix

	return p
}

func (s *SelectorPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorPrefixContext) GT() antlr.TerminalNode {
	return s.GetToken(LessParserGT, 0)
}

func (s *SelectorPrefixContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LessParserPLUS, 0)
}

func (s *SelectorPrefixContext) TIL() antlr.TerminalNode {
	return s.GetToken(LessParserTIL, 0)
}

func (s *SelectorPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterSelectorPrefix(s)
	}
}

func (s *SelectorPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitSelectorPrefix(s)
	}
}

func (p *LessParser) SelectorPrefix() (localctx ISelectorPrefixContext) {
	localctx = NewSelectorPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LessParserRULE_selectorPrefix)
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
	p.SetState(285)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LessParserGT)|(1<<LessParserTIL)|(1<<LessParserPLUS))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IAttribRelateContext is an interface to support dynamic dispatch.
type IAttribRelateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttribRelateContext differentiates from other interfaces.
	IsAttribRelateContext()
}

type AttribRelateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribRelateContext() *AttribRelateContext {
	var p = new(AttribRelateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_attribRelate
	return p
}

func (*AttribRelateContext) IsAttribRelateContext() {}

func NewAttribRelateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttribRelateContext {
	var p = new(AttribRelateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_attribRelate

	return p
}

func (s *AttribRelateContext) GetParser() antlr.Parser { return s.parser }
func (s *AttribRelateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribRelateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttribRelateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterAttribRelate(s)
	}
}

func (s *AttribRelateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitAttribRelate(s)
	}
}

func (p *LessParser) AttribRelate() (localctx IAttribRelateContext) {
	localctx = NewAttribRelateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, LessParserRULE_attribRelate)
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
	p.SetState(287)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(LessParserEQ-33))|(1<<(LessParserPIPE_EQ-33))|(1<<(LessParserTILD_EQ-33)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LessParserIdentifier, 0)
}

func (s *IdentifierContext) AllIdentifierPart() []IIdentifierPartContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierPartContext)(nil)).Elem())
	var tst = make([]IIdentifierPartContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierPartContext)
		}
	}

	return tst
}

func (s *IdentifierContext) IdentifierPart(i int) IIdentifierPartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierPartContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierPartContext)
}

func (s *IdentifierContext) InterpolationStart() antlr.TerminalNode {
	return s.GetToken(LessParserInterpolationStart, 0)
}

func (s *IdentifierContext) IdentifierVariableName() IIdentifierVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierVariableNameContext)
}

func (s *IdentifierContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(LessParserBlockEnd, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *LessParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, LessParserRULE_identifier)
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

	p.SetState(305)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LessParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(289)
			p.Match(LessParserIdentifier)
		}
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LessParserInterpolationStartAfter || _la == LessParserIdentifierAfter {
			{
				p.SetState(290)
				p.IdentifierPart()
			}

			p.SetState(295)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case LessParserInterpolationStart:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(296)
			p.Match(LessParserInterpolationStart)
		}
		{
			p.SetState(297)
			p.IdentifierVariableName()
		}
		{
			p.SetState(298)
			p.Match(LessParserBlockEnd)
		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LessParserInterpolationStartAfter || _la == LessParserIdentifierAfter {
			{
				p.SetState(299)
				p.IdentifierPart()
			}

			p.SetState(304)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifierPartContext is an interface to support dynamic dispatch.
type IIdentifierPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierPartContext differentiates from other interfaces.
	IsIdentifierPartContext()
}

type IdentifierPartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierPartContext() *IdentifierPartContext {
	var p = new(IdentifierPartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_identifierPart
	return p
}

func (*IdentifierPartContext) IsIdentifierPartContext() {}

func NewIdentifierPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierPartContext {
	var p = new(IdentifierPartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_identifierPart

	return p
}

func (s *IdentifierPartContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierPartContext) InterpolationStartAfter() antlr.TerminalNode {
	return s.GetToken(LessParserInterpolationStartAfter, 0)
}

func (s *IdentifierPartContext) IdentifierVariableName() IIdentifierVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierVariableNameContext)
}

func (s *IdentifierPartContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(LessParserBlockEnd, 0)
}

func (s *IdentifierPartContext) IdentifierAfter() antlr.TerminalNode {
	return s.GetToken(LessParserIdentifierAfter, 0)
}

func (s *IdentifierPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierPartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterIdentifierPart(s)
	}
}

func (s *IdentifierPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitIdentifierPart(s)
	}
}

func (p *LessParser) IdentifierPart() (localctx IIdentifierPartContext) {
	localctx = NewIdentifierPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, LessParserRULE_identifierPart)

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

	p.SetState(312)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LessParserInterpolationStartAfter:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(307)
			p.Match(LessParserInterpolationStartAfter)
		}
		{
			p.SetState(308)
			p.IdentifierVariableName()
		}
		{
			p.SetState(309)
			p.Match(LessParserBlockEnd)
		}

	case LessParserIdentifierAfter:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(311)
			p.Match(LessParserIdentifierAfter)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifierVariableNameContext is an interface to support dynamic dispatch.
type IIdentifierVariableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierVariableNameContext differentiates from other interfaces.
	IsIdentifierVariableNameContext()
}

type IdentifierVariableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierVariableNameContext() *IdentifierVariableNameContext {
	var p = new(IdentifierVariableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_identifierVariableName
	return p
}

func (*IdentifierVariableNameContext) IsIdentifierVariableNameContext() {}

func NewIdentifierVariableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierVariableNameContext {
	var p = new(IdentifierVariableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_identifierVariableName

	return p
}

func (s *IdentifierVariableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierVariableNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LessParserIdentifier, 0)
}

func (s *IdentifierVariableNameContext) IdentifierAfter() antlr.TerminalNode {
	return s.GetToken(LessParserIdentifierAfter, 0)
}

func (s *IdentifierVariableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierVariableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierVariableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterIdentifierVariableName(s)
	}
}

func (s *IdentifierVariableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitIdentifierVariableName(s)
	}
}

func (p *LessParser) IdentifierVariableName() (localctx IIdentifierVariableNameContext) {
	localctx = NewIdentifierVariableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, LessParserRULE_identifierVariableName)
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
	p.SetState(314)
	_la = p.GetTokenStream().LA(1)

	if !(_la == LessParserIdentifier || _la == LessParserIdentifierAfter) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PropertyContext) COLON() antlr.TerminalNode {
	return s.GetToken(LessParserCOLON, 0)
}

func (s *PropertyContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (p *LessParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, LessParserRULE_property)

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
		p.SetState(316)
		p.Identifier()
	}
	{
		p.SetState(317)
		p.Match(LessParserCOLON)
	}
	{
		p.SetState(318)
		p.Values()
	}

	return localctx
}

// IValuesContext is an interface to support dynamic dispatch.
type IValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValuesContext differentiates from other interfaces.
	IsValuesContext()
}

type ValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuesContext() *ValuesContext {
	var p = new(ValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_values
	return p
}

func (*ValuesContext) IsValuesContext() {}

func NewValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuesContext {
	var p = new(ValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_values

	return p
}

func (s *ValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuesContext) AllCommandStatement() []ICommandStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem())
	var tst = make([]ICommandStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandStatementContext)
		}
	}

	return tst
}

func (s *ValuesContext) CommandStatement(i int) ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *ValuesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LessParserCOMMA)
}

func (s *ValuesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LessParserCOMMA, i)
}

func (s *ValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterValues(s)
	}
}

func (s *ValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitValues(s)
	}
}

func (p *LessParser) Values() (localctx IValuesContext) {
	localctx = NewValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, LessParserRULE_values)
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
		p.SetState(320)
		p.CommandStatement()
	}
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LessParserCOMMA {
		{
			p.SetState(321)
			p.Match(LessParserCOMMA)
		}
		{
			p.SetState(322)
			p.CommandStatement()
		}

		p.SetState(327)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUrlContext is an interface to support dynamic dispatch.
type IUrlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUrlContext differentiates from other interfaces.
	IsUrlContext()
}

type UrlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrlContext() *UrlContext {
	var p = new(UrlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_url
	return p
}

func (*UrlContext) IsUrlContext() {}

func NewUrlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UrlContext {
	var p = new(UrlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_url

	return p
}

func (s *UrlContext) GetParser() antlr.Parser { return s.parser }

func (s *UrlContext) UrlStart() antlr.TerminalNode {
	return s.GetToken(LessParserUrlStart, 0)
}

func (s *UrlContext) Url() antlr.TerminalNode {
	return s.GetToken(LessParserUrl, 0)
}

func (s *UrlContext) UrlEnd() antlr.TerminalNode {
	return s.GetToken(LessParserUrlEnd, 0)
}

func (s *UrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UrlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterUrl(s)
	}
}

func (s *UrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitUrl(s)
	}
}

func (p *LessParser) Url() (localctx IUrlContext) {
	localctx = NewUrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, LessParserRULE_url)

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
		p.SetState(328)
		p.Match(LessParserUrlStart)
	}
	{
		p.SetState(329)
		p.Match(LessParserUrl)
	}
	{
		p.SetState(330)
		p.Match(LessParserUrlEnd)
	}

	return localctx
}

// IMeasurementContext is an interface to support dynamic dispatch.
type IMeasurementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeasurementContext differentiates from other interfaces.
	IsMeasurementContext()
}

type MeasurementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeasurementContext() *MeasurementContext {
	var p = new(MeasurementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_measurement
	return p
}

func (*MeasurementContext) IsMeasurementContext() {}

func NewMeasurementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MeasurementContext {
	var p = new(MeasurementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_measurement

	return p
}

func (s *MeasurementContext) GetParser() antlr.Parser { return s.parser }

func (s *MeasurementContext) Number() antlr.TerminalNode {
	return s.GetToken(LessParserNumber, 0)
}

func (s *MeasurementContext) Unit() antlr.TerminalNode {
	return s.GetToken(LessParserUnit, 0)
}

func (s *MeasurementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MeasurementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MeasurementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMeasurement(s)
	}
}

func (s *MeasurementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMeasurement(s)
	}
}

func (p *LessParser) Measurement() (localctx IMeasurementContext) {
	localctx = NewMeasurementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, LessParserRULE_measurement)
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
		p.SetState(332)
		p.Match(LessParserNumber)
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserUnit {
		{
			p.SetState(333)
			p.Match(LessParserUnit)
		}

	}

	return localctx
}
