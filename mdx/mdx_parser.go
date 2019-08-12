// Code generated from mdx.g4 by ANTLR 4.7.2. DO NOT EDIT.

package mdx // mdx
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 57, 393,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 104, 10, 3, 3, 3, 3, 3, 5, 3, 108, 10,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 114, 10, 3, 3, 3, 5, 3, 117, 10, 3, 3,
	4, 6, 4, 120, 10, 4, 13, 4, 14, 4, 121, 3, 5, 3, 5, 5, 5, 126, 10, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 136, 10, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 146, 10, 7, 3, 7, 3, 7, 5,
	7, 150, 10, 7, 3, 8, 3, 8, 3, 8, 7, 8, 155, 10, 8, 12, 8, 14, 8, 158, 11,
	8, 3, 9, 3, 9, 3, 9, 7, 9, 163, 10, 9, 12, 9, 14, 9, 166, 11, 9, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 7,
	13, 179, 10, 13, 12, 13, 14, 13, 182, 11, 13, 3, 14, 3, 14, 5, 14, 186,
	10, 14, 3, 14, 3, 14, 5, 14, 190, 10, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 16, 5, 16, 198, 10, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	7, 17, 206, 10, 17, 12, 17, 14, 17, 209, 11, 17, 3, 18, 3, 18, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 5, 22, 220, 10, 22, 3, 22, 3, 22,
	3, 22, 3, 23, 3, 23, 3, 23, 7, 23, 228, 10, 23, 12, 23, 14, 23, 231, 11,
	23, 3, 24, 3, 24, 5, 24, 235, 10, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 27, 7, 27, 244, 10, 27, 12, 27, 14, 27, 247, 11, 27, 3, 28, 3,
	28, 3, 28, 7, 28, 252, 10, 28, 12, 28, 14, 28, 255, 11, 28, 3, 29, 3, 29,
	3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 7, 31, 266, 10, 31, 12,
	31, 14, 31, 269, 11, 31, 3, 32, 3, 32, 3, 32, 5, 32, 274, 10, 32, 3, 33,
	3, 33, 3, 33, 3, 33, 7, 33, 280, 10, 33, 12, 33, 14, 33, 283, 11, 33, 3,
	34, 3, 34, 3, 34, 7, 34, 288, 10, 34, 12, 34, 14, 34, 291, 11, 34, 3, 35,
	3, 35, 3, 35, 7, 35, 296, 10, 35, 12, 35, 14, 35, 299, 11, 35, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 5, 36, 306, 10, 36, 3, 37, 3, 37, 3, 37, 5, 37,
	311, 10, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 5,
	38, 321, 10, 38, 7, 38, 323, 10, 38, 12, 38, 14, 38, 326, 11, 38, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 335, 10, 39, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 39, 5, 39, 342, 10, 39, 3, 40, 3, 40, 3, 40, 7, 40,
	347, 10, 40, 12, 40, 14, 40, 350, 11, 40, 3, 41, 3, 41, 5, 41, 354, 10,
	41, 3, 41, 5, 41, 357, 10, 41, 3, 41, 3, 41, 5, 41, 361, 10, 41, 3, 41,
	3, 41, 3, 42, 3, 42, 7, 42, 367, 10, 42, 12, 42, 14, 42, 370, 11, 42, 3,
	43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 5, 45, 381,
	10, 45, 3, 46, 3, 46, 5, 46, 385, 10, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3,
	49, 3, 49, 3, 49, 2, 2, 50, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96,
	2, 7, 5, 2, 27, 27, 33, 33, 46, 46, 5, 2, 8, 8, 17, 17, 19, 19, 4, 2, 4,
	4, 22, 22, 6, 2, 10, 12, 14, 14, 16, 16, 18, 18, 4, 2, 29, 29, 41, 41,
	2, 392, 2, 98, 3, 2, 2, 2, 4, 103, 3, 2, 2, 2, 6, 119, 3, 2, 2, 2, 8, 125,
	3, 2, 2, 2, 10, 127, 3, 2, 2, 2, 12, 137, 3, 2, 2, 2, 14, 151, 3, 2, 2,
	2, 16, 159, 3, 2, 2, 2, 18, 167, 3, 2, 2, 2, 20, 169, 3, 2, 2, 2, 22, 173,
	3, 2, 2, 2, 24, 175, 3, 2, 2, 2, 26, 185, 3, 2, 2, 2, 28, 194, 3, 2, 2,
	2, 30, 197, 3, 2, 2, 2, 32, 202, 3, 2, 2, 2, 34, 210, 3, 2, 2, 2, 36, 212,
	3, 2, 2, 2, 38, 214, 3, 2, 2, 2, 40, 216, 3, 2, 2, 2, 42, 219, 3, 2, 2,
	2, 44, 224, 3, 2, 2, 2, 46, 234, 3, 2, 2, 2, 48, 236, 3, 2, 2, 2, 50, 238,
	3, 2, 2, 2, 52, 240, 3, 2, 2, 2, 54, 248, 3, 2, 2, 2, 56, 256, 3, 2, 2,
	2, 58, 259, 3, 2, 2, 2, 60, 262, 3, 2, 2, 2, 62, 273, 3, 2, 2, 2, 64, 275,
	3, 2, 2, 2, 66, 284, 3, 2, 2, 2, 68, 292, 3, 2, 2, 2, 70, 305, 3, 2, 2,
	2, 72, 307, 3, 2, 2, 2, 74, 314, 3, 2, 2, 2, 76, 341, 3, 2, 2, 2, 78, 343,
	3, 2, 2, 2, 80, 351, 3, 2, 2, 2, 82, 364, 3, 2, 2, 2, 84, 371, 3, 2, 2,
	2, 86, 376, 3, 2, 2, 2, 88, 380, 3, 2, 2, 2, 90, 384, 3, 2, 2, 2, 92, 386,
	3, 2, 2, 2, 94, 388, 3, 2, 2, 2, 96, 390, 3, 2, 2, 2, 98, 99, 5, 4, 3,
	2, 99, 100, 7, 2, 2, 3, 100, 3, 3, 2, 2, 2, 101, 102, 7, 50, 2, 2, 102,
	104, 5, 6, 4, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 105,
	3, 2, 2, 2, 105, 107, 7, 42, 2, 2, 106, 108, 5, 14, 8, 2, 107, 106, 3,
	2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 110, 7, 34, 2,
	2, 110, 113, 5, 36, 19, 2, 111, 112, 7, 48, 2, 2, 112, 114, 5, 40, 21,
	2, 113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 116, 3, 2, 2, 2, 115,
	117, 5, 42, 22, 2, 116, 115, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 5,
	3, 2, 2, 2, 118, 120, 5, 8, 5, 2, 119, 118, 3, 2, 2, 2, 120, 121, 3, 2,
	2, 2, 121, 119, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 7, 3, 2, 2, 2, 123,
	126, 5, 12, 7, 2, 124, 126, 5, 10, 6, 2, 125, 123, 3, 2, 2, 2, 125, 124,
	3, 2, 2, 2, 126, 9, 3, 2, 2, 2, 127, 128, 7, 44, 2, 2, 128, 129, 5, 22,
	12, 2, 129, 135, 7, 24, 2, 2, 130, 131, 7, 3, 2, 2, 131, 132, 5, 52, 27,
	2, 132, 133, 7, 3, 2, 2, 133, 136, 3, 2, 2, 2, 134, 136, 5, 52, 27, 2,
	135, 130, 3, 2, 2, 2, 135, 134, 3, 2, 2, 2, 136, 11, 3, 2, 2, 2, 137, 138,
	7, 36, 2, 2, 138, 139, 5, 18, 10, 2, 139, 145, 7, 24, 2, 2, 140, 141, 7,
	3, 2, 2, 141, 142, 5, 54, 28, 2, 142, 143, 7, 3, 2, 2, 143, 146, 3, 2,
	2, 2, 144, 146, 5, 54, 28, 2, 145, 140, 3, 2, 2, 2, 145, 144, 3, 2, 2,
	2, 146, 147, 3, 2, 2, 2, 147, 149, 7, 7, 2, 2, 148, 150, 5, 16, 9, 2, 149,
	148, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 13, 3, 2, 2, 2, 151, 156, 5,
	26, 14, 2, 152, 153, 7, 7, 2, 2, 153, 155, 5, 26, 14, 2, 154, 152, 3, 2,
	2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2,
	157, 15, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159, 164, 5, 20, 11, 2, 160,
	161, 7, 7, 2, 2, 161, 163, 5, 20, 11, 2, 162, 160, 3, 2, 2, 2, 163, 166,
	3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 17, 3, 2,
	2, 2, 166, 164, 3, 2, 2, 2, 167, 168, 5, 24, 13, 2, 168, 19, 3, 2, 2, 2,
	169, 170, 5, 88, 45, 2, 170, 171, 7, 10, 2, 2, 171, 172, 5, 54, 28, 2,
	172, 21, 3, 2, 2, 2, 173, 174, 5, 24, 13, 2, 174, 23, 3, 2, 2, 2, 175,
	180, 5, 88, 45, 2, 176, 177, 7, 9, 2, 2, 177, 179, 5, 88, 45, 2, 178, 176,
	3, 2, 2, 2, 179, 182, 3, 2, 2, 2, 180, 178, 3, 2, 2, 2, 180, 181, 3, 2,
	2, 2, 181, 25, 3, 2, 2, 2, 182, 180, 3, 2, 2, 2, 183, 184, 7, 37, 2, 2,
	184, 186, 7, 31, 2, 2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186,
	187, 3, 2, 2, 2, 187, 189, 5, 52, 27, 2, 188, 190, 5, 30, 16, 2, 189, 188,
	3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 192, 7, 39,
	2, 2, 192, 193, 5, 28, 15, 2, 193, 27, 3, 2, 2, 2, 194, 195, 5, 88, 45,
	2, 195, 29, 3, 2, 2, 2, 196, 198, 7, 29, 2, 2, 197, 196, 3, 2, 2, 2, 197,
	198, 3, 2, 2, 2, 198, 199, 3, 2, 2, 2, 199, 200, 7, 41, 2, 2, 200, 201,
	5, 32, 17, 2, 201, 31, 3, 2, 2, 2, 202, 207, 5, 34, 18, 2, 203, 204, 7,
	7, 2, 2, 204, 206, 5, 34, 18, 2, 205, 203, 3, 2, 2, 2, 206, 209, 3, 2,
	2, 2, 207, 205, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 33, 3, 2, 2, 2,
	209, 207, 3, 2, 2, 2, 210, 211, 5, 24, 13, 2, 211, 35, 3, 2, 2, 2, 212,
	213, 5, 38, 20, 2, 213, 37, 3, 2, 2, 2, 214, 215, 5, 24, 13, 2, 215, 39,
	3, 2, 2, 2, 216, 217, 5, 52, 27, 2, 217, 41, 3, 2, 2, 2, 218, 220, 7, 26,
	2, 2, 219, 218, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221, 3, 2, 2, 2,
	221, 222, 7, 41, 2, 2, 222, 223, 5, 44, 23, 2, 223, 43, 3, 2, 2, 2, 224,
	225, 5, 46, 24, 2, 225, 229, 7, 7, 2, 2, 226, 228, 5, 46, 24, 2, 227, 226,
	3, 2, 2, 2, 228, 231, 3, 2, 2, 2, 229, 227, 3, 2, 2, 2, 229, 230, 3, 2,
	2, 2, 230, 45, 3, 2, 2, 2, 231, 229, 3, 2, 2, 2, 232, 235, 5, 48, 25, 2,
	233, 235, 5, 50, 26, 2, 234, 232, 3, 2, 2, 2, 234, 233, 3, 2, 2, 2, 235,
	47, 3, 2, 2, 2, 236, 237, 9, 2, 2, 2, 237, 49, 3, 2, 2, 2, 238, 239, 5,
	88, 45, 2, 239, 51, 3, 2, 2, 2, 240, 245, 5, 54, 28, 2, 241, 242, 7, 5,
	2, 2, 242, 244, 5, 54, 28, 2, 243, 241, 3, 2, 2, 2, 244, 247, 3, 2, 2,
	2, 245, 243, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 53, 3, 2, 2, 2, 247,
	245, 3, 2, 2, 2, 248, 253, 5, 60, 31, 2, 249, 252, 5, 56, 29, 2, 250, 252,
	5, 58, 30, 2, 251, 249, 3, 2, 2, 2, 251, 250, 3, 2, 2, 2, 252, 255, 3,
	2, 2, 2, 253, 251, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 55, 3, 2, 2,
	2, 255, 253, 3, 2, 2, 2, 256, 257, 7, 49, 2, 2, 257, 258, 5, 60, 31, 2,
	258, 57, 3, 2, 2, 2, 259, 260, 7, 40, 2, 2, 260, 261, 5, 60, 31, 2, 261,
	59, 3, 2, 2, 2, 262, 267, 5, 62, 32, 2, 263, 264, 7, 23, 2, 2, 264, 266,
	5, 62, 32, 2, 265, 263, 3, 2, 2, 2, 266, 269, 3, 2, 2, 2, 267, 265, 3,
	2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 61, 3, 2, 2, 2, 269, 267, 3, 2, 2,
	2, 270, 271, 7, 38, 2, 2, 271, 274, 5, 62, 32, 2, 272, 274, 5, 64, 33,
	2, 273, 270, 3, 2, 2, 2, 273, 272, 3, 2, 2, 2, 274, 63, 3, 2, 2, 2, 275,
	281, 5, 66, 34, 2, 276, 277, 5, 86, 44, 2, 277, 278, 5, 66, 34, 2, 278,
	280, 3, 2, 2, 2, 279, 276, 3, 2, 2, 2, 280, 283, 3, 2, 2, 2, 281, 279,
	3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 65, 3, 2, 2, 2, 283, 281, 3, 2,
	2, 2, 284, 289, 5, 68, 35, 2, 285, 286, 9, 3, 2, 2, 286, 288, 5, 68, 35,
	2, 287, 285, 3, 2, 2, 2, 288, 291, 3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 289,
	290, 3, 2, 2, 2, 290, 67, 3, 2, 2, 2, 291, 289, 3, 2, 2, 2, 292, 297, 5,
	70, 36, 2, 293, 294, 9, 4, 2, 2, 294, 296, 5, 70, 36, 2, 295, 293, 3, 2,
	2, 2, 296, 299, 3, 2, 2, 2, 297, 295, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2,
	298, 69, 3, 2, 2, 2, 299, 297, 3, 2, 2, 2, 300, 301, 7, 17, 2, 2, 301,
	306, 5, 74, 38, 2, 302, 303, 7, 19, 2, 2, 303, 306, 5, 74, 38, 2, 304,
	306, 5, 74, 38, 2, 305, 300, 3, 2, 2, 2, 305, 302, 3, 2, 2, 2, 305, 304,
	3, 2, 2, 2, 306, 71, 3, 2, 2, 2, 307, 308, 5, 88, 45, 2, 308, 310, 7, 15,
	2, 2, 309, 311, 5, 78, 40, 2, 310, 309, 3, 2, 2, 2, 310, 311, 3, 2, 2,
	2, 311, 312, 3, 2, 2, 2, 312, 313, 7, 21, 2, 2, 313, 73, 3, 2, 2, 2, 314,
	324, 5, 76, 39, 2, 315, 320, 7, 9, 2, 2, 316, 321, 5, 90, 46, 2, 317, 321,
	5, 94, 48, 2, 318, 321, 5, 92, 47, 2, 319, 321, 5, 72, 37, 2, 320, 316,
	3, 2, 2, 2, 320, 317, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 320, 319, 3, 2,
	2, 2, 321, 323, 3, 2, 2, 2, 322, 315, 3, 2, 2, 2, 323, 326, 3, 2, 2, 2,
	324, 322, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2, 325, 75, 3, 2, 2, 2, 326, 324,
	3, 2, 2, 2, 327, 342, 5, 72, 37, 2, 328, 329, 7, 15, 2, 2, 329, 330, 5,
	78, 40, 2, 330, 331, 7, 21, 2, 2, 331, 342, 3, 2, 2, 2, 332, 334, 7, 13,
	2, 2, 333, 335, 5, 78, 40, 2, 334, 333, 3, 2, 2, 2, 334, 335, 3, 2, 2,
	2, 335, 336, 3, 2, 2, 2, 336, 342, 7, 20, 2, 2, 337, 342, 5, 80, 41, 2,
	338, 342, 7, 56, 2, 2, 339, 342, 7, 51, 2, 2, 340, 342, 5, 88, 45, 2, 341,
	327, 3, 2, 2, 2, 341, 328, 3, 2, 2, 2, 341, 332, 3, 2, 2, 2, 341, 337,
	3, 2, 2, 2, 341, 338, 3, 2, 2, 2, 341, 339, 3, 2, 2, 2, 341, 340, 3, 2,
	2, 2, 342, 77, 3, 2, 2, 2, 343, 348, 5, 52, 27, 2, 344, 345, 7, 7, 2, 2,
	345, 347, 5, 52, 27, 2, 346, 344, 3, 2, 2, 2, 347, 350, 3, 2, 2, 2, 348,
	346, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 79, 3, 2, 2, 2, 350, 348, 3,
	2, 2, 2, 351, 353, 7, 25, 2, 2, 352, 354, 5, 54, 28, 2, 353, 352, 3, 2,
	2, 2, 353, 354, 3, 2, 2, 2, 354, 356, 3, 2, 2, 2, 355, 357, 5, 82, 42,
	2, 356, 355, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357, 360, 3, 2, 2, 2, 358,
	359, 7, 30, 2, 2, 359, 361, 5, 54, 28, 2, 360, 358, 3, 2, 2, 2, 360, 361,
	3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 363, 7, 32, 2, 2, 363, 81, 3, 2,
	2, 2, 364, 368, 5, 84, 43, 2, 365, 367, 5, 84, 43, 2, 366, 365, 3, 2, 2,
	2, 367, 370, 3, 2, 2, 2, 368, 366, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369,
	83, 3, 2, 2, 2, 370, 368, 3, 2, 2, 2, 371, 372, 7, 47, 2, 2, 372, 373,
	5, 54, 28, 2, 373, 374, 7, 45, 2, 2, 374, 375, 5, 54, 28, 2, 375, 85, 3,
	2, 2, 2, 376, 377, 9, 5, 2, 2, 377, 87, 3, 2, 2, 2, 378, 381, 5, 90, 46,
	2, 379, 381, 5, 94, 48, 2, 380, 378, 3, 2, 2, 2, 380, 379, 3, 2, 2, 2,
	381, 89, 3, 2, 2, 2, 382, 385, 5, 96, 49, 2, 383, 385, 7, 53, 2, 2, 384,
	382, 3, 2, 2, 2, 384, 383, 3, 2, 2, 2, 385, 91, 3, 2, 2, 2, 386, 387, 7,
	54, 2, 2, 387, 93, 3, 2, 2, 2, 388, 389, 7, 55, 2, 2, 389, 95, 3, 2, 2,
	2, 390, 391, 9, 6, 2, 2, 391, 97, 3, 2, 2, 2, 42, 103, 107, 113, 116, 121,
	125, 135, 145, 149, 156, 164, 180, 185, 189, 197, 207, 219, 229, 234, 245,
	251, 253, 267, 273, 281, 289, 297, 305, 310, 320, 324, 334, 341, 348, 353,
	356, 360, 368, 380, 384,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'''", "'*'", "':'", "';'", "','", "'||'", "'.'", "'='", "'>='", "'>'",
	"'{'", "'<='", "'('", "'<'", "'-'", "'<>'", "'+'", "'}'", "')'", "'/'",
	"'AND'", "'AS'", "'CASE'", "'CELL'", "'CELL_ORDINAL'", "'CREATE'", "'DIMENSION'",
	"'ELSE'", "'EMPTY'", "'END'", "'FORMATTED_VALUE'", "'FROM'", "'GLOBAL'",
	"'MEMBER'", "'NON'", "'NOT'", "'ON'", "'OR'", "'PROPERTIES'", "'SELECT'",
	"'SESSION'", "'SET'", "'THEN'", "'VALUE'", "'WHEN'", "'WHERE'", "'XOR'",
	"'WITH'",
}
var symbolicNames = []string{
	"", "QUOTE", "ASTERISK", "COLON", "SEMICOLON", "COMMA", "CONCAT", "DOT",
	"EQ", "GE", "GT", "LBRACE", "LE", "LPAREN", "LT", "MINUS", "NE", "PLUS",
	"RBRACE", "RPAREN", "SOLIDUS", "AND", "AS", "CASE", "CELL", "CELL_ORDINAL",
	"CREATE", "DIMENSION", "ELSE", "EMPTY", "END", "FORMATTED_VALUE", "FROM",
	"GLOBAL", "MEMBER", "NON", "NOT", "ON", "OR", "PROPERTIES", "SELECT", "SESSION",
	"SET", "THEN", "VALUE", "WHEN", "WHERE", "XOR", "WITH", "NUMBER", "F",
	"ID", "AMP_QUOTED_ID", "QUOTED_ID", "STRING", "WS",
}

var ruleNames = []string{
	"mdx_statement", "select_statement", "formula_specification", "single_formula_specification",
	"set_specification", "member_specification", "axis_specification_list",
	"member_property_def_list", "member_name", "member_property_definition",
	"set_name", "compound_id", "axis_specification", "axis_name", "dim_props",
	"property_list", "property", "cube_specification", "cube_name", "slicer_specification",
	"cell_props", "cell_property_list", "cell_property", "mandatory_cell_property",
	"provider_specific_cell_property", "expression", "value_expression", "value_xor_expression",
	"value_or_expression", "term5", "term4", "term3", "term2", "term", "factor",
	"function", "value_expression_primary", "value_expression_primary0", "exp_list",
	"case_expression", "when_list", "when_clause", "comp_op", "identifier",
	"unquoted_identifier", "amp_quoted_identifier", "quoted_identifier", "keyword",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type mdxParser struct {
	*antlr.BaseParser
}

func NewmdxParser(input antlr.TokenStream) *mdxParser {
	this := new(mdxParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "mdx.g4"

	return this
}

// mdxParser tokens.
const (
	mdxParserEOF             = antlr.TokenEOF
	mdxParserQUOTE           = 1
	mdxParserASTERISK        = 2
	mdxParserCOLON           = 3
	mdxParserSEMICOLON       = 4
	mdxParserCOMMA           = 5
	mdxParserCONCAT          = 6
	mdxParserDOT             = 7
	mdxParserEQ              = 8
	mdxParserGE              = 9
	mdxParserGT              = 10
	mdxParserLBRACE          = 11
	mdxParserLE              = 12
	mdxParserLPAREN          = 13
	mdxParserLT              = 14
	mdxParserMINUS           = 15
	mdxParserNE              = 16
	mdxParserPLUS            = 17
	mdxParserRBRACE          = 18
	mdxParserRPAREN          = 19
	mdxParserSOLIDUS         = 20
	mdxParserAND             = 21
	mdxParserAS              = 22
	mdxParserCASE            = 23
	mdxParserCELL            = 24
	mdxParserCELL_ORDINAL    = 25
	mdxParserCREATE          = 26
	mdxParserDIMENSION       = 27
	mdxParserELSE            = 28
	mdxParserEMPTY           = 29
	mdxParserEND             = 30
	mdxParserFORMATTED_VALUE = 31
	mdxParserFROM            = 32
	mdxParserGLOBAL          = 33
	mdxParserMEMBER          = 34
	mdxParserNON             = 35
	mdxParserNOT             = 36
	mdxParserON              = 37
	mdxParserOR              = 38
	mdxParserPROPERTIES      = 39
	mdxParserSELECT          = 40
	mdxParserSESSION         = 41
	mdxParserSET             = 42
	mdxParserTHEN            = 43
	mdxParserVALUE           = 44
	mdxParserWHEN            = 45
	mdxParserWHERE           = 46
	mdxParserXOR             = 47
	mdxParserWITH            = 48
	mdxParserNUMBER          = 49
	mdxParserF               = 50
	mdxParserID              = 51
	mdxParserAMP_QUOTED_ID   = 52
	mdxParserQUOTED_ID       = 53
	mdxParserSTRING          = 54
	mdxParserWS              = 55
)

// mdxParser rules.
const (
	mdxParserRULE_mdx_statement                   = 0
	mdxParserRULE_select_statement                = 1
	mdxParserRULE_formula_specification           = 2
	mdxParserRULE_single_formula_specification    = 3
	mdxParserRULE_set_specification               = 4
	mdxParserRULE_member_specification            = 5
	mdxParserRULE_axis_specification_list         = 6
	mdxParserRULE_member_property_def_list        = 7
	mdxParserRULE_member_name                     = 8
	mdxParserRULE_member_property_definition      = 9
	mdxParserRULE_set_name                        = 10
	mdxParserRULE_compound_id                     = 11
	mdxParserRULE_axis_specification              = 12
	mdxParserRULE_axis_name                       = 13
	mdxParserRULE_dim_props                       = 14
	mdxParserRULE_property_list                   = 15
	mdxParserRULE_property                        = 16
	mdxParserRULE_cube_specification              = 17
	mdxParserRULE_cube_name                       = 18
	mdxParserRULE_slicer_specification            = 19
	mdxParserRULE_cell_props                      = 20
	mdxParserRULE_cell_property_list              = 21
	mdxParserRULE_cell_property                   = 22
	mdxParserRULE_mandatory_cell_property         = 23
	mdxParserRULE_provider_specific_cell_property = 24
	mdxParserRULE_expression                      = 25
	mdxParserRULE_value_expression                = 26
	mdxParserRULE_value_xor_expression            = 27
	mdxParserRULE_value_or_expression             = 28
	mdxParserRULE_term5                           = 29
	mdxParserRULE_term4                           = 30
	mdxParserRULE_term3                           = 31
	mdxParserRULE_term2                           = 32
	mdxParserRULE_term                            = 33
	mdxParserRULE_factor                          = 34
	mdxParserRULE_function                        = 35
	mdxParserRULE_value_expression_primary        = 36
	mdxParserRULE_value_expression_primary0       = 37
	mdxParserRULE_exp_list                        = 38
	mdxParserRULE_case_expression                 = 39
	mdxParserRULE_when_list                       = 40
	mdxParserRULE_when_clause                     = 41
	mdxParserRULE_comp_op                         = 42
	mdxParserRULE_identifier                      = 43
	mdxParserRULE_unquoted_identifier             = 44
	mdxParserRULE_amp_quoted_identifier           = 45
	mdxParserRULE_quoted_identifier               = 46
	mdxParserRULE_keyword                         = 47
)

// IMdx_statementContext is an interface to support dynamic dispatch.
type IMdx_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMdx_statementContext differentiates from other interfaces.
	IsMdx_statementContext()
}

type Mdx_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMdx_statementContext() *Mdx_statementContext {
	var p = new(Mdx_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_mdx_statement
	return p
}

func (*Mdx_statementContext) IsMdx_statementContext() {}

func NewMdx_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mdx_statementContext {
	var p = new(Mdx_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_mdx_statement

	return p
}

func (s *Mdx_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Mdx_statementContext) EOF() antlr.TerminalNode {
	return s.GetToken(mdxParserEOF, 0)
}

func (s *Mdx_statementContext) Select_statement() ISelect_statementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelect_statementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelect_statementContext)
}

func (s *Mdx_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mdx_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mdx_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterMdx_statement(s)
	}
}

func (s *Mdx_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitMdx_statement(s)
	}
}

func (p *mdxParser) Mdx_statement() (localctx IMdx_statementContext) {
	localctx = NewMdx_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, mdxParserRULE_mdx_statement)

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
		p.SetState(96)
		p.Select_statement()
	}

	{
		p.SetState(97)
		p.Match(mdxParserEOF)
	}

	return localctx
}

// ISelect_statementContext is an interface to support dynamic dispatch.
type ISelect_statementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelect_statementContext differentiates from other interfaces.
	IsSelect_statementContext()
}

type Select_statementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_statementContext() *Select_statementContext {
	var p = new(Select_statementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_select_statement
	return p
}

func (*Select_statementContext) IsSelect_statementContext() {}

func NewSelect_statementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_statementContext {
	var p = new(Select_statementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_select_statement

	return p
}

func (s *Select_statementContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_statementContext) SELECT() antlr.TerminalNode {
	return s.GetToken(mdxParserSELECT, 0)
}

func (s *Select_statementContext) FROM() antlr.TerminalNode {
	return s.GetToken(mdxParserFROM, 0)
}

func (s *Select_statementContext) Cube_specification() ICube_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICube_specificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICube_specificationContext)
}

func (s *Select_statementContext) WITH() antlr.TerminalNode {
	return s.GetToken(mdxParserWITH, 0)
}

func (s *Select_statementContext) Formula_specification() IFormula_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormula_specificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormula_specificationContext)
}

func (s *Select_statementContext) Axis_specification_list() IAxis_specification_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAxis_specification_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAxis_specification_listContext)
}

func (s *Select_statementContext) WHERE() antlr.TerminalNode {
	return s.GetToken(mdxParserWHERE, 0)
}

func (s *Select_statementContext) Slicer_specification() ISlicer_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISlicer_specificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISlicer_specificationContext)
}

func (s *Select_statementContext) Cell_props() ICell_propsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICell_propsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICell_propsContext)
}

func (s *Select_statementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_statementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_statementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterSelect_statement(s)
	}
}

func (s *Select_statementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitSelect_statement(s)
	}
}

func (p *mdxParser) Select_statement() (localctx ISelect_statementContext) {
	localctx = NewSelect_statementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, mdxParserRULE_select_statement)
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
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mdxParserWITH {
		{
			p.SetState(99)
			p.Match(mdxParserWITH)
		}
		{
			p.SetState(100)
			p.Formula_specification()
		}

	}
	{
		p.SetState(103)
		p.Match(mdxParserSELECT)
	}
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mdxParserLBRACE)|(1<<mdxParserLPAREN)|(1<<mdxParserMINUS)|(1<<mdxParserPLUS)|(1<<mdxParserCASE)|(1<<mdxParserDIMENSION))) != 0) || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(mdxParserNON-35))|(1<<(mdxParserNOT-35))|(1<<(mdxParserPROPERTIES-35))|(1<<(mdxParserNUMBER-35))|(1<<(mdxParserID-35))|(1<<(mdxParserQUOTED_ID-35))|(1<<(mdxParserSTRING-35)))) != 0) {
		{
			p.SetState(104)
			p.Axis_specification_list()
		}

	}
	{
		p.SetState(107)
		p.Match(mdxParserFROM)
	}
	{
		p.SetState(108)
		p.Cube_specification()
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mdxParserWHERE {
		{
			p.SetState(109)
			p.Match(mdxParserWHERE)
		}
		{
			p.SetState(110)
			p.Slicer_specification()
		}

	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mdxParserCELL || _la == mdxParserPROPERTIES {
		{
			p.SetState(113)
			p.Cell_props()
		}

	}

	return localctx
}

// IFormula_specificationContext is an interface to support dynamic dispatch.
type IFormula_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormula_specificationContext differentiates from other interfaces.
	IsFormula_specificationContext()
}

type Formula_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormula_specificationContext() *Formula_specificationContext {
	var p = new(Formula_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_formula_specification
	return p
}

func (*Formula_specificationContext) IsFormula_specificationContext() {}

func NewFormula_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Formula_specificationContext {
	var p = new(Formula_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_formula_specification

	return p
}

func (s *Formula_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Formula_specificationContext) AllSingle_formula_specification() []ISingle_formula_specificationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingle_formula_specificationContext)(nil)).Elem())
	var tst = make([]ISingle_formula_specificationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingle_formula_specificationContext)
		}
	}

	return tst
}

func (s *Formula_specificationContext) Single_formula_specification(i int) ISingle_formula_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingle_formula_specificationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingle_formula_specificationContext)
}

func (s *Formula_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Formula_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Formula_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterFormula_specification(s)
	}
}

func (s *Formula_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitFormula_specification(s)
	}
}

func (p *mdxParser) Formula_specification() (localctx IFormula_specificationContext) {
	localctx = NewFormula_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, mdxParserRULE_formula_specification)
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
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == mdxParserMEMBER || _la == mdxParserSET {
		{
			p.SetState(116)
			p.Single_formula_specification()
		}

		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISingle_formula_specificationContext is an interface to support dynamic dispatch.
type ISingle_formula_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingle_formula_specificationContext differentiates from other interfaces.
	IsSingle_formula_specificationContext()
}

type Single_formula_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_formula_specificationContext() *Single_formula_specificationContext {
	var p = new(Single_formula_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_single_formula_specification
	return p
}

func (*Single_formula_specificationContext) IsSingle_formula_specificationContext() {}

func NewSingle_formula_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_formula_specificationContext {
	var p = new(Single_formula_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_single_formula_specification

	return p
}

func (s *Single_formula_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_formula_specificationContext) Member_specification() IMember_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMember_specificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMember_specificationContext)
}

func (s *Single_formula_specificationContext) Set_specification() ISet_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_specificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_specificationContext)
}

func (s *Single_formula_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_formula_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_formula_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterSingle_formula_specification(s)
	}
}

func (s *Single_formula_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitSingle_formula_specification(s)
	}
}

func (p *mdxParser) Single_formula_specification() (localctx ISingle_formula_specificationContext) {
	localctx = NewSingle_formula_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, mdxParserRULE_single_formula_specification)

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

	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mdxParserMEMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Member_specification()
		}

	case mdxParserSET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Set_specification()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISet_specificationContext is an interface to support dynamic dispatch.
type ISet_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_specificationContext differentiates from other interfaces.
	IsSet_specificationContext()
}

type Set_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_specificationContext() *Set_specificationContext {
	var p = new(Set_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_set_specification
	return p
}

func (*Set_specificationContext) IsSet_specificationContext() {}

func NewSet_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_specificationContext {
	var p = new(Set_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_set_specification

	return p
}

func (s *Set_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_specificationContext) SET() antlr.TerminalNode {
	return s.GetToken(mdxParserSET, 0)
}

func (s *Set_specificationContext) Set_name() ISet_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_nameContext)
}

func (s *Set_specificationContext) AS() antlr.TerminalNode {
	return s.GetToken(mdxParserAS, 0)
}

func (s *Set_specificationContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(mdxParserQUOTE)
}

func (s *Set_specificationContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserQUOTE, i)
}

func (s *Set_specificationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Set_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterSet_specification(s)
	}
}

func (s *Set_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitSet_specification(s)
	}
}

func (p *mdxParser) Set_specification() (localctx ISet_specificationContext) {
	localctx = NewSet_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, mdxParserRULE_set_specification)

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
		p.Match(mdxParserSET)
	}
	{
		p.SetState(126)
		p.Set_name()
	}
	{
		p.SetState(127)
		p.Match(mdxParserAS)
	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mdxParserQUOTE:
		{
			p.SetState(128)
			p.Match(mdxParserQUOTE)
		}
		{
			p.SetState(129)
			p.Expression()
		}
		{
			p.SetState(130)
			p.Match(mdxParserQUOTE)
		}

	case mdxParserLBRACE, mdxParserLPAREN, mdxParserMINUS, mdxParserPLUS, mdxParserCASE, mdxParserDIMENSION, mdxParserNOT, mdxParserPROPERTIES, mdxParserNUMBER, mdxParserID, mdxParserQUOTED_ID, mdxParserSTRING:
		{
			p.SetState(132)
			p.Expression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMember_specificationContext is an interface to support dynamic dispatch.
type IMember_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMember_specificationContext differentiates from other interfaces.
	IsMember_specificationContext()
}

type Member_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMember_specificationContext() *Member_specificationContext {
	var p = new(Member_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_member_specification
	return p
}

func (*Member_specificationContext) IsMember_specificationContext() {}

func NewMember_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Member_specificationContext {
	var p = new(Member_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_member_specification

	return p
}

func (s *Member_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Member_specificationContext) MEMBER() antlr.TerminalNode {
	return s.GetToken(mdxParserMEMBER, 0)
}

func (s *Member_specificationContext) Member_name() IMember_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMember_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMember_nameContext)
}

func (s *Member_specificationContext) AS() antlr.TerminalNode {
	return s.GetToken(mdxParserAS, 0)
}

func (s *Member_specificationContext) COMMA() antlr.TerminalNode {
	return s.GetToken(mdxParserCOMMA, 0)
}

func (s *Member_specificationContext) AllQUOTE() []antlr.TerminalNode {
	return s.GetTokens(mdxParserQUOTE)
}

func (s *Member_specificationContext) QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserQUOTE, i)
}

func (s *Member_specificationContext) Value_expression() IValue_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValue_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValue_expressionContext)
}

func (s *Member_specificationContext) Member_property_def_list() IMember_property_def_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMember_property_def_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMember_property_def_listContext)
}

func (s *Member_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Member_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Member_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterMember_specification(s)
	}
}

func (s *Member_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitMember_specification(s)
	}
}

func (p *mdxParser) Member_specification() (localctx IMember_specificationContext) {
	localctx = NewMember_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, mdxParserRULE_member_specification)
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
		p.SetState(135)
		p.Match(mdxParserMEMBER)
	}
	{
		p.SetState(136)
		p.Member_name()
	}
	{
		p.SetState(137)
		p.Match(mdxParserAS)
	}

	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mdxParserQUOTE:
		{
			p.SetState(138)
			p.Match(mdxParserQUOTE)
		}
		{
			p.SetState(139)
			p.Value_expression()
		}
		{
			p.SetState(140)
			p.Match(mdxParserQUOTE)
		}

	case mdxParserLBRACE, mdxParserLPAREN, mdxParserMINUS, mdxParserPLUS, mdxParserCASE, mdxParserDIMENSION, mdxParserNOT, mdxParserPROPERTIES, mdxParserNUMBER, mdxParserID, mdxParserQUOTED_ID, mdxParserSTRING:
		{
			p.SetState(142)
			p.Value_expression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(145)
		p.Match(mdxParserCOMMA)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(mdxParserDIMENSION-27))|(1<<(mdxParserPROPERTIES-27))|(1<<(mdxParserID-27))|(1<<(mdxParserQUOTED_ID-27)))) != 0 {
		{
			p.SetState(146)
			p.Member_property_def_list()
		}

	}

	return localctx
}

// IAxis_specification_listContext is an interface to support dynamic dispatch.
type IAxis_specification_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAxis_specification_listContext differentiates from other interfaces.
	IsAxis_specification_listContext()
}

type Axis_specification_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAxis_specification_listContext() *Axis_specification_listContext {
	var p = new(Axis_specification_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_axis_specification_list
	return p
}

func (*Axis_specification_listContext) IsAxis_specification_listContext() {}

func NewAxis_specification_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Axis_specification_listContext {
	var p = new(Axis_specification_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_axis_specification_list

	return p
}

func (s *Axis_specification_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Axis_specification_listContext) AllAxis_specification() []IAxis_specificationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAxis_specificationContext)(nil)).Elem())
	var tst = make([]IAxis_specificationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAxis_specificationContext)
		}
	}

	return tst
}

func (s *Axis_specification_listContext) Axis_specification(i int) IAxis_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAxis_specificationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAxis_specificationContext)
}

func (s *Axis_specification_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mdxParserCOMMA)
}

func (s *Axis_specification_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserCOMMA, i)
}

func (s *Axis_specification_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Axis_specification_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Axis_specification_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterAxis_specification_list(s)
	}
}

func (s *Axis_specification_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitAxis_specification_list(s)
	}
}

func (p *mdxParser) Axis_specification_list() (localctx IAxis_specification_listContext) {
	localctx = NewAxis_specification_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, mdxParserRULE_axis_specification_list)
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
		p.SetState(149)
		p.Axis_specification()
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mdxParserCOMMA {
		{
			p.SetState(150)
			p.Match(mdxParserCOMMA)
		}
		{
			p.SetState(151)
			p.Axis_specification()
		}

		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMember_property_def_listContext is an interface to support dynamic dispatch.
type IMember_property_def_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMember_property_def_listContext differentiates from other interfaces.
	IsMember_property_def_listContext()
}

type Member_property_def_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMember_property_def_listContext() *Member_property_def_listContext {
	var p = new(Member_property_def_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_member_property_def_list
	return p
}

func (*Member_property_def_listContext) IsMember_property_def_listContext() {}

func NewMember_property_def_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Member_property_def_listContext {
	var p = new(Member_property_def_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_member_property_def_list

	return p
}

func (s *Member_property_def_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Member_property_def_listContext) AllMember_property_definition() []IMember_property_definitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMember_property_definitionContext)(nil)).Elem())
	var tst = make([]IMember_property_definitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMember_property_definitionContext)
		}
	}

	return tst
}

func (s *Member_property_def_listContext) Member_property_definition(i int) IMember_property_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMember_property_definitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMember_property_definitionContext)
}

func (s *Member_property_def_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mdxParserCOMMA)
}

func (s *Member_property_def_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserCOMMA, i)
}

func (s *Member_property_def_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Member_property_def_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Member_property_def_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterMember_property_def_list(s)
	}
}

func (s *Member_property_def_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitMember_property_def_list(s)
	}
}

func (p *mdxParser) Member_property_def_list() (localctx IMember_property_def_listContext) {
	localctx = NewMember_property_def_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, mdxParserRULE_member_property_def_list)
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
		p.SetState(157)
		p.Member_property_definition()
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mdxParserCOMMA {
		{
			p.SetState(158)
			p.Match(mdxParserCOMMA)
		}
		{
			p.SetState(159)
			p.Member_property_definition()
		}

		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMember_nameContext is an interface to support dynamic dispatch.
type IMember_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMember_nameContext differentiates from other interfaces.
	IsMember_nameContext()
}

type Member_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMember_nameContext() *Member_nameContext {
	var p = new(Member_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_member_name
	return p
}

func (*Member_nameContext) IsMember_nameContext() {}

func NewMember_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Member_nameContext {
	var p = new(Member_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_member_name

	return p
}

func (s *Member_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Member_nameContext) Compound_id() ICompound_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompound_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompound_idContext)
}

func (s *Member_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Member_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Member_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterMember_name(s)
	}
}

func (s *Member_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitMember_name(s)
	}
}

func (p *mdxParser) Member_name() (localctx IMember_nameContext) {
	localctx = NewMember_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, mdxParserRULE_member_name)

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
		p.SetState(165)
		p.Compound_id()
	}

	return localctx
}

// IMember_property_definitionContext is an interface to support dynamic dispatch.
type IMember_property_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMember_property_definitionContext differentiates from other interfaces.
	IsMember_property_definitionContext()
}

type Member_property_definitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMember_property_definitionContext() *Member_property_definitionContext {
	var p = new(Member_property_definitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_member_property_definition
	return p
}

func (*Member_property_definitionContext) IsMember_property_definitionContext() {}

func NewMember_property_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Member_property_definitionContext {
	var p = new(Member_property_definitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_member_property_definition

	return p
}

func (s *Member_property_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *Member_property_definitionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Member_property_definitionContext) EQ() antlr.TerminalNode {
	return s.GetToken(mdxParserEQ, 0)
}

func (s *Member_property_definitionContext) Value_expression() IValue_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValue_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValue_expressionContext)
}

func (s *Member_property_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Member_property_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Member_property_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterMember_property_definition(s)
	}
}

func (s *Member_property_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitMember_property_definition(s)
	}
}

func (p *mdxParser) Member_property_definition() (localctx IMember_property_definitionContext) {
	localctx = NewMember_property_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, mdxParserRULE_member_property_definition)

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
		p.SetState(167)
		p.Identifier()
	}
	{
		p.SetState(168)
		p.Match(mdxParserEQ)
	}
	{
		p.SetState(169)
		p.Value_expression()
	}

	return localctx
}

// ISet_nameContext is an interface to support dynamic dispatch.
type ISet_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_nameContext differentiates from other interfaces.
	IsSet_nameContext()
}

type Set_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_nameContext() *Set_nameContext {
	var p = new(Set_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_set_name
	return p
}

func (*Set_nameContext) IsSet_nameContext() {}

func NewSet_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_nameContext {
	var p = new(Set_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_set_name

	return p
}

func (s *Set_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_nameContext) Compound_id() ICompound_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompound_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompound_idContext)
}

func (s *Set_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterSet_name(s)
	}
}

func (s *Set_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitSet_name(s)
	}
}

func (p *mdxParser) Set_name() (localctx ISet_nameContext) {
	localctx = NewSet_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, mdxParserRULE_set_name)

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
		p.Compound_id()
	}

	return localctx
}

// ICompound_idContext is an interface to support dynamic dispatch.
type ICompound_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompound_idContext differentiates from other interfaces.
	IsCompound_idContext()
}

type Compound_idContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompound_idContext() *Compound_idContext {
	var p = new(Compound_idContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_compound_id
	return p
}

func (*Compound_idContext) IsCompound_idContext() {}

func NewCompound_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Compound_idContext {
	var p = new(Compound_idContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_compound_id

	return p
}

func (s *Compound_idContext) GetParser() antlr.Parser { return s.parser }

func (s *Compound_idContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *Compound_idContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Compound_idContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(mdxParserDOT)
}

func (s *Compound_idContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserDOT, i)
}

func (s *Compound_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Compound_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Compound_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterCompound_id(s)
	}
}

func (s *Compound_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitCompound_id(s)
	}
}

func (p *mdxParser) Compound_id() (localctx ICompound_idContext) {
	localctx = NewCompound_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, mdxParserRULE_compound_id)
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
		p.SetState(173)
		p.Identifier()
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mdxParserDOT {
		{
			p.SetState(174)
			p.Match(mdxParserDOT)
		}
		{
			p.SetState(175)
			p.Identifier()
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAxis_specificationContext is an interface to support dynamic dispatch.
type IAxis_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAxis_specificationContext differentiates from other interfaces.
	IsAxis_specificationContext()
}

type Axis_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAxis_specificationContext() *Axis_specificationContext {
	var p = new(Axis_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_axis_specification
	return p
}

func (*Axis_specificationContext) IsAxis_specificationContext() {}

func NewAxis_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Axis_specificationContext {
	var p = new(Axis_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_axis_specification

	return p
}

func (s *Axis_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Axis_specificationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Axis_specificationContext) ON() antlr.TerminalNode {
	return s.GetToken(mdxParserON, 0)
}

func (s *Axis_specificationContext) Axis_name() IAxis_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAxis_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAxis_nameContext)
}

func (s *Axis_specificationContext) NON() antlr.TerminalNode {
	return s.GetToken(mdxParserNON, 0)
}

func (s *Axis_specificationContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(mdxParserEMPTY, 0)
}

func (s *Axis_specificationContext) Dim_props() IDim_propsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDim_propsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDim_propsContext)
}

func (s *Axis_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Axis_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Axis_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterAxis_specification(s)
	}
}

func (s *Axis_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitAxis_specification(s)
	}
}

func (p *mdxParser) Axis_specification() (localctx IAxis_specificationContext) {
	localctx = NewAxis_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, mdxParserRULE_axis_specification)
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
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mdxParserNON {
		{
			p.SetState(181)
			p.Match(mdxParserNON)
		}
		{
			p.SetState(182)
			p.Match(mdxParserEMPTY)
		}

	}
	{
		p.SetState(185)
		p.Expression()
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mdxParserDIMENSION || _la == mdxParserPROPERTIES {
		{
			p.SetState(186)
			p.Dim_props()
		}

	}
	{
		p.SetState(189)
		p.Match(mdxParserON)
	}
	{
		p.SetState(190)
		p.Axis_name()
	}

	return localctx
}

// IAxis_nameContext is an interface to support dynamic dispatch.
type IAxis_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAxis_nameContext differentiates from other interfaces.
	IsAxis_nameContext()
}

type Axis_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAxis_nameContext() *Axis_nameContext {
	var p = new(Axis_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_axis_name
	return p
}

func (*Axis_nameContext) IsAxis_nameContext() {}

func NewAxis_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Axis_nameContext {
	var p = new(Axis_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_axis_name

	return p
}

func (s *Axis_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Axis_nameContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Axis_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Axis_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Axis_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterAxis_name(s)
	}
}

func (s *Axis_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitAxis_name(s)
	}
}

func (p *mdxParser) Axis_name() (localctx IAxis_nameContext) {
	localctx = NewAxis_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, mdxParserRULE_axis_name)

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
		p.Identifier()
	}

	return localctx
}

// IDim_propsContext is an interface to support dynamic dispatch.
type IDim_propsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDim_propsContext differentiates from other interfaces.
	IsDim_propsContext()
}

type Dim_propsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDim_propsContext() *Dim_propsContext {
	var p = new(Dim_propsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_dim_props
	return p
}

func (*Dim_propsContext) IsDim_propsContext() {}

func NewDim_propsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dim_propsContext {
	var p = new(Dim_propsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_dim_props

	return p
}

func (s *Dim_propsContext) GetParser() antlr.Parser { return s.parser }

func (s *Dim_propsContext) PROPERTIES() antlr.TerminalNode {
	return s.GetToken(mdxParserPROPERTIES, 0)
}

func (s *Dim_propsContext) Property_list() IProperty_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProperty_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProperty_listContext)
}

func (s *Dim_propsContext) DIMENSION() antlr.TerminalNode {
	return s.GetToken(mdxParserDIMENSION, 0)
}

func (s *Dim_propsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dim_propsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dim_propsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterDim_props(s)
	}
}

func (s *Dim_propsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitDim_props(s)
	}
}

func (p *mdxParser) Dim_props() (localctx IDim_propsContext) {
	localctx = NewDim_propsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, mdxParserRULE_dim_props)
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
	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mdxParserDIMENSION {
		{
			p.SetState(194)
			p.Match(mdxParserDIMENSION)
		}

	}
	{
		p.SetState(197)
		p.Match(mdxParserPROPERTIES)
	}
	{
		p.SetState(198)
		p.Property_list()
	}

	return localctx
}

// IProperty_listContext is an interface to support dynamic dispatch.
type IProperty_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProperty_listContext differentiates from other interfaces.
	IsProperty_listContext()
}

type Property_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProperty_listContext() *Property_listContext {
	var p = new(Property_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_property_list
	return p
}

func (*Property_listContext) IsProperty_listContext() {}

func NewProperty_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_listContext {
	var p = new(Property_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_property_list

	return p
}

func (s *Property_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Property_listContext) AllProperty() []IPropertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyContext)(nil)).Elem())
	var tst = make([]IPropertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyContext)
		}
	}

	return tst
}

func (s *Property_listContext) Property(i int) IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *Property_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mdxParserCOMMA)
}

func (s *Property_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserCOMMA, i)
}

func (s *Property_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterProperty_list(s)
	}
}

func (s *Property_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitProperty_list(s)
	}
}

func (p *mdxParser) Property_list() (localctx IProperty_listContext) {
	localctx = NewProperty_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, mdxParserRULE_property_list)
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
		p.SetState(200)
		p.Property()
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mdxParserCOMMA {
		{
			p.SetState(201)
			p.Match(mdxParserCOMMA)
		}
		{
			p.SetState(202)
			p.Property()
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = mdxParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) Compound_id() ICompound_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompound_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompound_idContext)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (p *mdxParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, mdxParserRULE_property)

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
		p.SetState(208)
		p.Compound_id()
	}

	return localctx
}

// ICube_specificationContext is an interface to support dynamic dispatch.
type ICube_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCube_specificationContext differentiates from other interfaces.
	IsCube_specificationContext()
}

type Cube_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCube_specificationContext() *Cube_specificationContext {
	var p = new(Cube_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_cube_specification
	return p
}

func (*Cube_specificationContext) IsCube_specificationContext() {}

func NewCube_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cube_specificationContext {
	var p = new(Cube_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_cube_specification

	return p
}

func (s *Cube_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Cube_specificationContext) Cube_name() ICube_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICube_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICube_nameContext)
}

func (s *Cube_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cube_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cube_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterCube_specification(s)
	}
}

func (s *Cube_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitCube_specification(s)
	}
}

func (p *mdxParser) Cube_specification() (localctx ICube_specificationContext) {
	localctx = NewCube_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, mdxParserRULE_cube_specification)

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
		p.Cube_name()
	}

	return localctx
}

// ICube_nameContext is an interface to support dynamic dispatch.
type ICube_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCube_nameContext differentiates from other interfaces.
	IsCube_nameContext()
}

type Cube_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCube_nameContext() *Cube_nameContext {
	var p = new(Cube_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_cube_name
	return p
}

func (*Cube_nameContext) IsCube_nameContext() {}

func NewCube_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cube_nameContext {
	var p = new(Cube_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_cube_name

	return p
}

func (s *Cube_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Cube_nameContext) Compound_id() ICompound_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompound_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompound_idContext)
}

func (s *Cube_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cube_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cube_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterCube_name(s)
	}
}

func (s *Cube_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitCube_name(s)
	}
}

func (p *mdxParser) Cube_name() (localctx ICube_nameContext) {
	localctx = NewCube_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, mdxParserRULE_cube_name)

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
		p.SetState(212)
		p.Compound_id()
	}

	return localctx
}

// ISlicer_specificationContext is an interface to support dynamic dispatch.
type ISlicer_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSlicer_specificationContext differentiates from other interfaces.
	IsSlicer_specificationContext()
}

type Slicer_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySlicer_specificationContext() *Slicer_specificationContext {
	var p = new(Slicer_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_slicer_specification
	return p
}

func (*Slicer_specificationContext) IsSlicer_specificationContext() {}

func NewSlicer_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Slicer_specificationContext {
	var p = new(Slicer_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_slicer_specification

	return p
}

func (s *Slicer_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Slicer_specificationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Slicer_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Slicer_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Slicer_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterSlicer_specification(s)
	}
}

func (s *Slicer_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitSlicer_specification(s)
	}
}

func (p *mdxParser) Slicer_specification() (localctx ISlicer_specificationContext) {
	localctx = NewSlicer_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, mdxParserRULE_slicer_specification)

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
		p.SetState(214)
		p.Expression()
	}

	return localctx
}

// ICell_propsContext is an interface to support dynamic dispatch.
type ICell_propsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCell_propsContext differentiates from other interfaces.
	IsCell_propsContext()
}

type Cell_propsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCell_propsContext() *Cell_propsContext {
	var p = new(Cell_propsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_cell_props
	return p
}

func (*Cell_propsContext) IsCell_propsContext() {}

func NewCell_propsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cell_propsContext {
	var p = new(Cell_propsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_cell_props

	return p
}

func (s *Cell_propsContext) GetParser() antlr.Parser { return s.parser }

func (s *Cell_propsContext) PROPERTIES() antlr.TerminalNode {
	return s.GetToken(mdxParserPROPERTIES, 0)
}

func (s *Cell_propsContext) Cell_property_list() ICell_property_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICell_property_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICell_property_listContext)
}

func (s *Cell_propsContext) CELL() antlr.TerminalNode {
	return s.GetToken(mdxParserCELL, 0)
}

func (s *Cell_propsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cell_propsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cell_propsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterCell_props(s)
	}
}

func (s *Cell_propsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitCell_props(s)
	}
}

func (p *mdxParser) Cell_props() (localctx ICell_propsContext) {
	localctx = NewCell_propsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, mdxParserRULE_cell_props)
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
	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mdxParserCELL {
		{
			p.SetState(216)
			p.Match(mdxParserCELL)
		}

	}
	{
		p.SetState(219)
		p.Match(mdxParserPROPERTIES)
	}
	{
		p.SetState(220)
		p.Cell_property_list()
	}

	return localctx
}

// ICell_property_listContext is an interface to support dynamic dispatch.
type ICell_property_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCell_property_listContext differentiates from other interfaces.
	IsCell_property_listContext()
}

type Cell_property_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCell_property_listContext() *Cell_property_listContext {
	var p = new(Cell_property_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_cell_property_list
	return p
}

func (*Cell_property_listContext) IsCell_property_listContext() {}

func NewCell_property_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cell_property_listContext {
	var p = new(Cell_property_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_cell_property_list

	return p
}

func (s *Cell_property_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Cell_property_listContext) AllCell_property() []ICell_propertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICell_propertyContext)(nil)).Elem())
	var tst = make([]ICell_propertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICell_propertyContext)
		}
	}

	return tst
}

func (s *Cell_property_listContext) Cell_property(i int) ICell_propertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICell_propertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICell_propertyContext)
}

func (s *Cell_property_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(mdxParserCOMMA, 0)
}

func (s *Cell_property_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cell_property_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cell_property_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterCell_property_list(s)
	}
}

func (s *Cell_property_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitCell_property_list(s)
	}
}

func (p *mdxParser) Cell_property_list() (localctx ICell_property_listContext) {
	localctx = NewCell_property_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, mdxParserRULE_cell_property_list)
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
		p.SetState(222)
		p.Cell_property()
	}
	{
		p.SetState(223)
		p.Match(mdxParserCOMMA)
	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-25)&-(0x1f+1)) == 0 && ((1<<uint((_la-25)))&((1<<(mdxParserCELL_ORDINAL-25))|(1<<(mdxParserDIMENSION-25))|(1<<(mdxParserFORMATTED_VALUE-25))|(1<<(mdxParserPROPERTIES-25))|(1<<(mdxParserVALUE-25))|(1<<(mdxParserID-25))|(1<<(mdxParserQUOTED_ID-25)))) != 0 {
		{
			p.SetState(224)
			p.Cell_property()
		}

		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICell_propertyContext is an interface to support dynamic dispatch.
type ICell_propertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCell_propertyContext differentiates from other interfaces.
	IsCell_propertyContext()
}

type Cell_propertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCell_propertyContext() *Cell_propertyContext {
	var p = new(Cell_propertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_cell_property
	return p
}

func (*Cell_propertyContext) IsCell_propertyContext() {}

func NewCell_propertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cell_propertyContext {
	var p = new(Cell_propertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_cell_property

	return p
}

func (s *Cell_propertyContext) GetParser() antlr.Parser { return s.parser }

func (s *Cell_propertyContext) Mandatory_cell_property() IMandatory_cell_propertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMandatory_cell_propertyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMandatory_cell_propertyContext)
}

func (s *Cell_propertyContext) Provider_specific_cell_property() IProvider_specific_cell_propertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProvider_specific_cell_propertyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProvider_specific_cell_propertyContext)
}

func (s *Cell_propertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cell_propertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cell_propertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterCell_property(s)
	}
}

func (s *Cell_propertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitCell_property(s)
	}
}

func (p *mdxParser) Cell_property() (localctx ICell_propertyContext) {
	localctx = NewCell_propertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, mdxParserRULE_cell_property)

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

	p.SetState(232)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mdxParserCELL_ORDINAL, mdxParserFORMATTED_VALUE, mdxParserVALUE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.Mandatory_cell_property()
		}

	case mdxParserDIMENSION, mdxParserPROPERTIES, mdxParserID, mdxParserQUOTED_ID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)
			p.Provider_specific_cell_property()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMandatory_cell_propertyContext is an interface to support dynamic dispatch.
type IMandatory_cell_propertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMandatory_cell_propertyContext differentiates from other interfaces.
	IsMandatory_cell_propertyContext()
}

type Mandatory_cell_propertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMandatory_cell_propertyContext() *Mandatory_cell_propertyContext {
	var p = new(Mandatory_cell_propertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_mandatory_cell_property
	return p
}

func (*Mandatory_cell_propertyContext) IsMandatory_cell_propertyContext() {}

func NewMandatory_cell_propertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mandatory_cell_propertyContext {
	var p = new(Mandatory_cell_propertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_mandatory_cell_property

	return p
}

func (s *Mandatory_cell_propertyContext) GetParser() antlr.Parser { return s.parser }

func (s *Mandatory_cell_propertyContext) CELL_ORDINAL() antlr.TerminalNode {
	return s.GetToken(mdxParserCELL_ORDINAL, 0)
}

func (s *Mandatory_cell_propertyContext) VALUE() antlr.TerminalNode {
	return s.GetToken(mdxParserVALUE, 0)
}

func (s *Mandatory_cell_propertyContext) FORMATTED_VALUE() antlr.TerminalNode {
	return s.GetToken(mdxParserFORMATTED_VALUE, 0)
}

func (s *Mandatory_cell_propertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mandatory_cell_propertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mandatory_cell_propertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterMandatory_cell_property(s)
	}
}

func (s *Mandatory_cell_propertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitMandatory_cell_property(s)
	}
}

func (p *mdxParser) Mandatory_cell_property() (localctx IMandatory_cell_propertyContext) {
	localctx = NewMandatory_cell_propertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, mdxParserRULE_mandatory_cell_property)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la-25)&-(0x1f+1)) == 0 && ((1<<uint((_la-25)))&((1<<(mdxParserCELL_ORDINAL-25))|(1<<(mdxParserFORMATTED_VALUE-25))|(1<<(mdxParserVALUE-25)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IProvider_specific_cell_propertyContext is an interface to support dynamic dispatch.
type IProvider_specific_cell_propertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProvider_specific_cell_propertyContext differentiates from other interfaces.
	IsProvider_specific_cell_propertyContext()
}

type Provider_specific_cell_propertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProvider_specific_cell_propertyContext() *Provider_specific_cell_propertyContext {
	var p = new(Provider_specific_cell_propertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_provider_specific_cell_property
	return p
}

func (*Provider_specific_cell_propertyContext) IsProvider_specific_cell_propertyContext() {}

func NewProvider_specific_cell_propertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Provider_specific_cell_propertyContext {
	var p = new(Provider_specific_cell_propertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_provider_specific_cell_property

	return p
}

func (s *Provider_specific_cell_propertyContext) GetParser() antlr.Parser { return s.parser }

func (s *Provider_specific_cell_propertyContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Provider_specific_cell_propertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Provider_specific_cell_propertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Provider_specific_cell_propertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterProvider_specific_cell_property(s)
	}
}

func (s *Provider_specific_cell_propertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitProvider_specific_cell_property(s)
	}
}

func (p *mdxParser) Provider_specific_cell_property() (localctx IProvider_specific_cell_propertyContext) {
	localctx = NewProvider_specific_cell_propertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, mdxParserRULE_provider_specific_cell_property)

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
		p.SetState(236)
		p.Identifier()
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
	p.RuleIndex = mdxParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllValue_expression() []IValue_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValue_expressionContext)(nil)).Elem())
	var tst = make([]IValue_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValue_expressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Value_expression(i int) IValue_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValue_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValue_expressionContext)
}

func (s *ExpressionContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(mdxParserCOLON)
}

func (s *ExpressionContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserCOLON, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *mdxParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, mdxParserRULE_expression)
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
		p.SetState(238)
		p.Value_expression()
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mdxParserCOLON {
		{
			p.SetState(239)
			p.Match(mdxParserCOLON)
		}
		{
			p.SetState(240)
			p.Value_expression()
		}

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IValue_expressionContext is an interface to support dynamic dispatch.
type IValue_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValue_expressionContext differentiates from other interfaces.
	IsValue_expressionContext()
}

type Value_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_expressionContext() *Value_expressionContext {
	var p = new(Value_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_value_expression
	return p
}

func (*Value_expressionContext) IsValue_expressionContext() {}

func NewValue_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_expressionContext {
	var p = new(Value_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_value_expression

	return p
}

func (s *Value_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_expressionContext) Term5() ITerm5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerm5Context)
}

func (s *Value_expressionContext) AllValue_xor_expression() []IValue_xor_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValue_xor_expressionContext)(nil)).Elem())
	var tst = make([]IValue_xor_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValue_xor_expressionContext)
		}
	}

	return tst
}

func (s *Value_expressionContext) Value_xor_expression(i int) IValue_xor_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValue_xor_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValue_xor_expressionContext)
}

func (s *Value_expressionContext) AllValue_or_expression() []IValue_or_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValue_or_expressionContext)(nil)).Elem())
	var tst = make([]IValue_or_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValue_or_expressionContext)
		}
	}

	return tst
}

func (s *Value_expressionContext) Value_or_expression(i int) IValue_or_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValue_or_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValue_or_expressionContext)
}

func (s *Value_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterValue_expression(s)
	}
}

func (s *Value_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitValue_expression(s)
	}
}

func (p *mdxParser) Value_expression() (localctx IValue_expressionContext) {
	localctx = NewValue_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, mdxParserRULE_value_expression)
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
		p.SetState(246)
		p.Term5()
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mdxParserOR || _la == mdxParserXOR {
		p.SetState(249)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case mdxParserXOR:
			{
				p.SetState(247)
				p.Value_xor_expression()
			}

		case mdxParserOR:
			{
				p.SetState(248)
				p.Value_or_expression()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IValue_xor_expressionContext is an interface to support dynamic dispatch.
type IValue_xor_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValue_xor_expressionContext differentiates from other interfaces.
	IsValue_xor_expressionContext()
}

type Value_xor_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_xor_expressionContext() *Value_xor_expressionContext {
	var p = new(Value_xor_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_value_xor_expression
	return p
}

func (*Value_xor_expressionContext) IsValue_xor_expressionContext() {}

func NewValue_xor_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_xor_expressionContext {
	var p = new(Value_xor_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_value_xor_expression

	return p
}

func (s *Value_xor_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_xor_expressionContext) XOR() antlr.TerminalNode {
	return s.GetToken(mdxParserXOR, 0)
}

func (s *Value_xor_expressionContext) Term5() ITerm5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerm5Context)
}

func (s *Value_xor_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_xor_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_xor_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterValue_xor_expression(s)
	}
}

func (s *Value_xor_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitValue_xor_expression(s)
	}
}

func (p *mdxParser) Value_xor_expression() (localctx IValue_xor_expressionContext) {
	localctx = NewValue_xor_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, mdxParserRULE_value_xor_expression)

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
		p.SetState(254)
		p.Match(mdxParserXOR)
	}
	{
		p.SetState(255)
		p.Term5()
	}

	return localctx
}

// IValue_or_expressionContext is an interface to support dynamic dispatch.
type IValue_or_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValue_or_expressionContext differentiates from other interfaces.
	IsValue_or_expressionContext()
}

type Value_or_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_or_expressionContext() *Value_or_expressionContext {
	var p = new(Value_or_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_value_or_expression
	return p
}

func (*Value_or_expressionContext) IsValue_or_expressionContext() {}

func NewValue_or_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_or_expressionContext {
	var p = new(Value_or_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_value_or_expression

	return p
}

func (s *Value_or_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_or_expressionContext) OR() antlr.TerminalNode {
	return s.GetToken(mdxParserOR, 0)
}

func (s *Value_or_expressionContext) Term5() ITerm5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerm5Context)
}

func (s *Value_or_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_or_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_or_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterValue_or_expression(s)
	}
}

func (s *Value_or_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitValue_or_expression(s)
	}
}

func (p *mdxParser) Value_or_expression() (localctx IValue_or_expressionContext) {
	localctx = NewValue_or_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, mdxParserRULE_value_or_expression)

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
		p.SetState(257)
		p.Match(mdxParserOR)
	}
	{
		p.SetState(258)
		p.Term5()
	}

	return localctx
}

// ITerm5Context is an interface to support dynamic dispatch.
type ITerm5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerm5Context differentiates from other interfaces.
	IsTerm5Context()
}

type Term5Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerm5Context() *Term5Context {
	var p = new(Term5Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_term5
	return p
}

func (*Term5Context) IsTerm5Context() {}

func NewTerm5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Term5Context {
	var p = new(Term5Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_term5

	return p
}

func (s *Term5Context) GetParser() antlr.Parser { return s.parser }

func (s *Term5Context) AllTerm4() []ITerm4Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITerm4Context)(nil)).Elem())
	var tst = make([]ITerm4Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITerm4Context)
		}
	}

	return tst
}

func (s *Term5Context) Term4(i int) ITerm4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm4Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITerm4Context)
}

func (s *Term5Context) AllAND() []antlr.TerminalNode {
	return s.GetTokens(mdxParserAND)
}

func (s *Term5Context) AND(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserAND, i)
}

func (s *Term5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Term5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Term5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterTerm5(s)
	}
}

func (s *Term5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitTerm5(s)
	}
}

func (p *mdxParser) Term5() (localctx ITerm5Context) {
	localctx = NewTerm5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, mdxParserRULE_term5)
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
		p.SetState(260)
		p.Term4()
	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mdxParserAND {
		{
			p.SetState(261)
			p.Match(mdxParserAND)
		}
		{
			p.SetState(262)
			p.Term4()
		}

		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITerm4Context is an interface to support dynamic dispatch.
type ITerm4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerm4Context differentiates from other interfaces.
	IsTerm4Context()
}

type Term4Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerm4Context() *Term4Context {
	var p = new(Term4Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_term4
	return p
}

func (*Term4Context) IsTerm4Context() {}

func NewTerm4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Term4Context {
	var p = new(Term4Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_term4

	return p
}

func (s *Term4Context) GetParser() antlr.Parser { return s.parser }

func (s *Term4Context) NOT() antlr.TerminalNode {
	return s.GetToken(mdxParserNOT, 0)
}

func (s *Term4Context) Term4() ITerm4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerm4Context)
}

func (s *Term4Context) Term3() ITerm3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerm3Context)
}

func (s *Term4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Term4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Term4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterTerm4(s)
	}
}

func (s *Term4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitTerm4(s)
	}
}

func (p *mdxParser) Term4() (localctx ITerm4Context) {
	localctx = NewTerm4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, mdxParserRULE_term4)

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

	p.SetState(271)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mdxParserNOT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.Match(mdxParserNOT)
		}
		{
			p.SetState(269)
			p.Term4()
		}

	case mdxParserLBRACE, mdxParserLPAREN, mdxParserMINUS, mdxParserPLUS, mdxParserCASE, mdxParserDIMENSION, mdxParserPROPERTIES, mdxParserNUMBER, mdxParserID, mdxParserQUOTED_ID, mdxParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)
			p.Term3()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITerm3Context is an interface to support dynamic dispatch.
type ITerm3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerm3Context differentiates from other interfaces.
	IsTerm3Context()
}

type Term3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerm3Context() *Term3Context {
	var p = new(Term3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_term3
	return p
}

func (*Term3Context) IsTerm3Context() {}

func NewTerm3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Term3Context {
	var p = new(Term3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_term3

	return p
}

func (s *Term3Context) GetParser() antlr.Parser { return s.parser }

func (s *Term3Context) AllTerm2() []ITerm2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITerm2Context)(nil)).Elem())
	var tst = make([]ITerm2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITerm2Context)
		}
	}

	return tst
}

func (s *Term3Context) Term2(i int) ITerm2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITerm2Context)
}

func (s *Term3Context) AllComp_op() []IComp_opContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComp_opContext)(nil)).Elem())
	var tst = make([]IComp_opContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComp_opContext)
		}
	}

	return tst
}

func (s *Term3Context) Comp_op(i int) IComp_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComp_opContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComp_opContext)
}

func (s *Term3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Term3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Term3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterTerm3(s)
	}
}

func (s *Term3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitTerm3(s)
	}
}

func (p *mdxParser) Term3() (localctx ITerm3Context) {
	localctx = NewTerm3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, mdxParserRULE_term3)
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
		p.Term2()
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mdxParserEQ)|(1<<mdxParserGE)|(1<<mdxParserGT)|(1<<mdxParserLE)|(1<<mdxParserLT)|(1<<mdxParserNE))) != 0 {
		{
			p.SetState(274)
			p.Comp_op()
		}
		{
			p.SetState(275)
			p.Term2()
		}

		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITerm2Context is an interface to support dynamic dispatch.
type ITerm2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerm2Context differentiates from other interfaces.
	IsTerm2Context()
}

type Term2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerm2Context() *Term2Context {
	var p = new(Term2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_term2
	return p
}

func (*Term2Context) IsTerm2Context() {}

func NewTerm2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Term2Context {
	var p = new(Term2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_term2

	return p
}

func (s *Term2Context) GetParser() antlr.Parser { return s.parser }

func (s *Term2Context) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Term2Context) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Term2Context) AllCONCAT() []antlr.TerminalNode {
	return s.GetTokens(mdxParserCONCAT)
}

func (s *Term2Context) CONCAT(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserCONCAT, i)
}

func (s *Term2Context) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(mdxParserPLUS)
}

func (s *Term2Context) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserPLUS, i)
}

func (s *Term2Context) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(mdxParserMINUS)
}

func (s *Term2Context) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserMINUS, i)
}

func (s *Term2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Term2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Term2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterTerm2(s)
	}
}

func (s *Term2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitTerm2(s)
	}
}

func (p *mdxParser) Term2() (localctx ITerm2Context) {
	localctx = NewTerm2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, mdxParserRULE_term2)
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
		p.SetState(282)
		p.Term()
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mdxParserCONCAT)|(1<<mdxParserMINUS)|(1<<mdxParserPLUS))) != 0 {
		{
			p.SetState(283)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mdxParserCONCAT)|(1<<mdxParserMINUS)|(1<<mdxParserPLUS))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(284)
			p.Term()
		}

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITermContext is an interface to support dynamic dispatch.
type ITermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTermContext differentiates from other interfaces.
	IsTermContext()
}

type TermContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTermContext() *TermContext {
	var p = new(TermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) AllFactor() []IFactorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFactorContext)(nil)).Elem())
	var tst = make([]IFactorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFactorContext)
		}
	}

	return tst
}

func (s *TermContext) Factor(i int) IFactorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFactorContext)
}

func (s *TermContext) AllSOLIDUS() []antlr.TerminalNode {
	return s.GetTokens(mdxParserSOLIDUS)
}

func (s *TermContext) SOLIDUS(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserSOLIDUS, i)
}

func (s *TermContext) AllASTERISK() []antlr.TerminalNode {
	return s.GetTokens(mdxParserASTERISK)
}

func (s *TermContext) ASTERISK(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserASTERISK, i)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *mdxParser) Term() (localctx ITermContext) {
	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, mdxParserRULE_term)
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
		p.SetState(290)
		p.Factor()
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mdxParserASTERISK || _la == mdxParserSOLIDUS {
		{
			p.SetState(291)
			_la = p.GetTokenStream().LA(1)

			if !(_la == mdxParserASTERISK || _la == mdxParserSOLIDUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(292)
			p.Factor()
		}

		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFactorContext is an interface to support dynamic dispatch.
type IFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactorContext differentiates from other interfaces.
	IsFactorContext()
}

type FactorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactorContext() *FactorContext {
	var p = new(FactorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_factor
	return p
}

func (*FactorContext) IsFactorContext() {}

func NewFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactorContext {
	var p = new(FactorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_factor

	return p
}

func (s *FactorContext) GetParser() antlr.Parser { return s.parser }

func (s *FactorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(mdxParserMINUS, 0)
}

func (s *FactorContext) Value_expression_primary() IValue_expression_primaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValue_expression_primaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValue_expression_primaryContext)
}

func (s *FactorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(mdxParserPLUS, 0)
}

func (s *FactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterFactor(s)
	}
}

func (s *FactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitFactor(s)
	}
}

func (p *mdxParser) Factor() (localctx IFactorContext) {
	localctx = NewFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, mdxParserRULE_factor)

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

	p.SetState(303)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mdxParserMINUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)
			p.Match(mdxParserMINUS)
		}
		{
			p.SetState(299)
			p.Value_expression_primary()
		}

	case mdxParserPLUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(300)
			p.Match(mdxParserPLUS)
		}
		{
			p.SetState(301)
			p.Value_expression_primary()
		}

	case mdxParserLBRACE, mdxParserLPAREN, mdxParserCASE, mdxParserDIMENSION, mdxParserPROPERTIES, mdxParserNUMBER, mdxParserID, mdxParserQUOTED_ID, mdxParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(302)
			p.Value_expression_primary()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = mdxParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(mdxParserLPAREN, 0)
}

func (s *FunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(mdxParserRPAREN, 0)
}

func (s *FunctionContext) Exp_list() IExp_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp_listContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *mdxParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, mdxParserRULE_function)
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
		p.SetState(305)
		p.Identifier()
	}
	{
		p.SetState(306)
		p.Match(mdxParserLPAREN)
	}
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mdxParserLBRACE)|(1<<mdxParserLPAREN)|(1<<mdxParserMINUS)|(1<<mdxParserPLUS)|(1<<mdxParserCASE)|(1<<mdxParserDIMENSION))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(mdxParserNOT-36))|(1<<(mdxParserPROPERTIES-36))|(1<<(mdxParserNUMBER-36))|(1<<(mdxParserID-36))|(1<<(mdxParserQUOTED_ID-36))|(1<<(mdxParserSTRING-36)))) != 0) {
		{
			p.SetState(307)
			p.Exp_list()
		}

	}
	{
		p.SetState(310)
		p.Match(mdxParserRPAREN)
	}

	return localctx
}

// IValue_expression_primaryContext is an interface to support dynamic dispatch.
type IValue_expression_primaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValue_expression_primaryContext differentiates from other interfaces.
	IsValue_expression_primaryContext()
}

type Value_expression_primaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_expression_primaryContext() *Value_expression_primaryContext {
	var p = new(Value_expression_primaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_value_expression_primary
	return p
}

func (*Value_expression_primaryContext) IsValue_expression_primaryContext() {}

func NewValue_expression_primaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_expression_primaryContext {
	var p = new(Value_expression_primaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_value_expression_primary

	return p
}

func (s *Value_expression_primaryContext) GetParser() antlr.Parser { return s.parser }

func (s *Value_expression_primaryContext) Value_expression_primary0() IValue_expression_primary0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValue_expression_primary0Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValue_expression_primary0Context)
}

func (s *Value_expression_primaryContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(mdxParserDOT)
}

func (s *Value_expression_primaryContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserDOT, i)
}

func (s *Value_expression_primaryContext) AllUnquoted_identifier() []IUnquoted_identifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnquoted_identifierContext)(nil)).Elem())
	var tst = make([]IUnquoted_identifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnquoted_identifierContext)
		}
	}

	return tst
}

func (s *Value_expression_primaryContext) Unquoted_identifier(i int) IUnquoted_identifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnquoted_identifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnquoted_identifierContext)
}

func (s *Value_expression_primaryContext) AllQuoted_identifier() []IQuoted_identifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuoted_identifierContext)(nil)).Elem())
	var tst = make([]IQuoted_identifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuoted_identifierContext)
		}
	}

	return tst
}

func (s *Value_expression_primaryContext) Quoted_identifier(i int) IQuoted_identifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuoted_identifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuoted_identifierContext)
}

func (s *Value_expression_primaryContext) AllAmp_quoted_identifier() []IAmp_quoted_identifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAmp_quoted_identifierContext)(nil)).Elem())
	var tst = make([]IAmp_quoted_identifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAmp_quoted_identifierContext)
		}
	}

	return tst
}

func (s *Value_expression_primaryContext) Amp_quoted_identifier(i int) IAmp_quoted_identifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAmp_quoted_identifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAmp_quoted_identifierContext)
}

func (s *Value_expression_primaryContext) AllFunction() []IFunctionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionContext)(nil)).Elem())
	var tst = make([]IFunctionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionContext)
		}
	}

	return tst
}

func (s *Value_expression_primaryContext) Function(i int) IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *Value_expression_primaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_expression_primaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_expression_primaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterValue_expression_primary(s)
	}
}

func (s *Value_expression_primaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitValue_expression_primary(s)
	}
}

func (p *mdxParser) Value_expression_primary() (localctx IValue_expression_primaryContext) {
	localctx = NewValue_expression_primaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, mdxParserRULE_value_expression_primary)
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
		p.SetState(312)
		p.Value_expression_primary0()
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mdxParserDOT {
		{
			p.SetState(313)
			p.Match(mdxParserDOT)
		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(314)
				p.Unquoted_identifier()
			}

		case 2:
			{
				p.SetState(315)
				p.Quoted_identifier()
			}

		case 3:
			{
				p.SetState(316)
				p.Amp_quoted_identifier()
			}

		case 4:
			{
				p.SetState(317)
				p.Function()
			}

		}

		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IValue_expression_primary0Context is an interface to support dynamic dispatch.
type IValue_expression_primary0Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValue_expression_primary0Context differentiates from other interfaces.
	IsValue_expression_primary0Context()
}

type Value_expression_primary0Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValue_expression_primary0Context() *Value_expression_primary0Context {
	var p = new(Value_expression_primary0Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_value_expression_primary0
	return p
}

func (*Value_expression_primary0Context) IsValue_expression_primary0Context() {}

func NewValue_expression_primary0Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Value_expression_primary0Context {
	var p = new(Value_expression_primary0Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_value_expression_primary0

	return p
}

func (s *Value_expression_primary0Context) GetParser() antlr.Parser { return s.parser }

func (s *Value_expression_primary0Context) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *Value_expression_primary0Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(mdxParserLPAREN, 0)
}

func (s *Value_expression_primary0Context) Exp_list() IExp_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExp_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExp_listContext)
}

func (s *Value_expression_primary0Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(mdxParserRPAREN, 0)
}

func (s *Value_expression_primary0Context) LBRACE() antlr.TerminalNode {
	return s.GetToken(mdxParserLBRACE, 0)
}

func (s *Value_expression_primary0Context) RBRACE() antlr.TerminalNode {
	return s.GetToken(mdxParserRBRACE, 0)
}

func (s *Value_expression_primary0Context) Case_expression() ICase_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICase_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICase_expressionContext)
}

func (s *Value_expression_primary0Context) STRING() antlr.TerminalNode {
	return s.GetToken(mdxParserSTRING, 0)
}

func (s *Value_expression_primary0Context) NUMBER() antlr.TerminalNode {
	return s.GetToken(mdxParserNUMBER, 0)
}

func (s *Value_expression_primary0Context) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Value_expression_primary0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Value_expression_primary0Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Value_expression_primary0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterValue_expression_primary0(s)
	}
}

func (s *Value_expression_primary0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitValue_expression_primary0(s)
	}
}

func (p *mdxParser) Value_expression_primary0() (localctx IValue_expression_primary0Context) {
	localctx = NewValue_expression_primary0Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, mdxParserRULE_value_expression_primary0)
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

	p.SetState(339)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(325)
			p.Function()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(326)
			p.Match(mdxParserLPAREN)
		}
		{
			p.SetState(327)
			p.Exp_list()
		}
		{
			p.SetState(328)
			p.Match(mdxParserRPAREN)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(330)
			p.Match(mdxParserLBRACE)
		}
		p.SetState(332)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mdxParserLBRACE)|(1<<mdxParserLPAREN)|(1<<mdxParserMINUS)|(1<<mdxParserPLUS)|(1<<mdxParserCASE)|(1<<mdxParserDIMENSION))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(mdxParserNOT-36))|(1<<(mdxParserPROPERTIES-36))|(1<<(mdxParserNUMBER-36))|(1<<(mdxParserID-36))|(1<<(mdxParserQUOTED_ID-36))|(1<<(mdxParserSTRING-36)))) != 0) {
			{
				p.SetState(331)
				p.Exp_list()
			}

		}
		{
			p.SetState(334)
			p.Match(mdxParserRBRACE)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(335)
			p.Case_expression()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(336)
			p.Match(mdxParserSTRING)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(337)
			p.Match(mdxParserNUMBER)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(338)
			p.Identifier()
		}

	}

	return localctx
}

// IExp_listContext is an interface to support dynamic dispatch.
type IExp_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExp_listContext differentiates from other interfaces.
	IsExp_listContext()
}

type Exp_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExp_listContext() *Exp_listContext {
	var p = new(Exp_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_exp_list
	return p
}

func (*Exp_listContext) IsExp_listContext() {}

func NewExp_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exp_listContext {
	var p = new(Exp_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_exp_list

	return p
}

func (s *Exp_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Exp_listContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Exp_listContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Exp_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(mdxParserCOMMA)
}

func (s *Exp_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(mdxParserCOMMA, i)
}

func (s *Exp_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exp_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exp_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterExp_list(s)
	}
}

func (s *Exp_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitExp_list(s)
	}
}

func (p *mdxParser) Exp_list() (localctx IExp_listContext) {
	localctx = NewExp_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, mdxParserRULE_exp_list)
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
		p.SetState(341)
		p.Expression()
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mdxParserCOMMA {
		{
			p.SetState(342)
			p.Match(mdxParserCOMMA)
		}
		{
			p.SetState(343)
			p.Expression()
		}

		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICase_expressionContext is an interface to support dynamic dispatch.
type ICase_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCase_expressionContext differentiates from other interfaces.
	IsCase_expressionContext()
}

type Case_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCase_expressionContext() *Case_expressionContext {
	var p = new(Case_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_case_expression
	return p
}

func (*Case_expressionContext) IsCase_expressionContext() {}

func NewCase_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_expressionContext {
	var p = new(Case_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_case_expression

	return p
}

func (s *Case_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_expressionContext) CASE() antlr.TerminalNode {
	return s.GetToken(mdxParserCASE, 0)
}

func (s *Case_expressionContext) END() antlr.TerminalNode {
	return s.GetToken(mdxParserEND, 0)
}

func (s *Case_expressionContext) AllValue_expression() []IValue_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValue_expressionContext)(nil)).Elem())
	var tst = make([]IValue_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValue_expressionContext)
		}
	}

	return tst
}

func (s *Case_expressionContext) Value_expression(i int) IValue_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValue_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValue_expressionContext)
}

func (s *Case_expressionContext) When_list() IWhen_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhen_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhen_listContext)
}

func (s *Case_expressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(mdxParserELSE, 0)
}

func (s *Case_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterCase_expression(s)
	}
}

func (s *Case_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitCase_expression(s)
	}
}

func (p *mdxParser) Case_expression() (localctx ICase_expressionContext) {
	localctx = NewCase_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, mdxParserRULE_case_expression)
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
		p.SetState(349)
		p.Match(mdxParserCASE)
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mdxParserLBRACE)|(1<<mdxParserLPAREN)|(1<<mdxParserMINUS)|(1<<mdxParserPLUS)|(1<<mdxParserCASE)|(1<<mdxParserDIMENSION))) != 0) || (((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(mdxParserNOT-36))|(1<<(mdxParserPROPERTIES-36))|(1<<(mdxParserNUMBER-36))|(1<<(mdxParserID-36))|(1<<(mdxParserQUOTED_ID-36))|(1<<(mdxParserSTRING-36)))) != 0) {
		{
			p.SetState(350)
			p.Value_expression()
		}

	}
	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mdxParserWHEN {
		{
			p.SetState(353)
			p.When_list()
		}

	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == mdxParserELSE {
		{
			p.SetState(356)
			p.Match(mdxParserELSE)
		}
		{
			p.SetState(357)
			p.Value_expression()
		}

	}
	{
		p.SetState(360)
		p.Match(mdxParserEND)
	}

	return localctx
}

// IWhen_listContext is an interface to support dynamic dispatch.
type IWhen_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhen_listContext differentiates from other interfaces.
	IsWhen_listContext()
}

type When_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhen_listContext() *When_listContext {
	var p = new(When_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_when_list
	return p
}

func (*When_listContext) IsWhen_listContext() {}

func NewWhen_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *When_listContext {
	var p = new(When_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_when_list

	return p
}

func (s *When_listContext) GetParser() antlr.Parser { return s.parser }

func (s *When_listContext) AllWhen_clause() []IWhen_clauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhen_clauseContext)(nil)).Elem())
	var tst = make([]IWhen_clauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhen_clauseContext)
		}
	}

	return tst
}

func (s *When_listContext) When_clause(i int) IWhen_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhen_clauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWhen_clauseContext)
}

func (s *When_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *When_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *When_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterWhen_list(s)
	}
}

func (s *When_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitWhen_list(s)
	}
}

func (p *mdxParser) When_list() (localctx IWhen_listContext) {
	localctx = NewWhen_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, mdxParserRULE_when_list)
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
		p.SetState(362)
		p.When_clause()
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == mdxParserWHEN {
		{
			p.SetState(363)
			p.When_clause()
		}

		p.SetState(368)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IWhen_clauseContext is an interface to support dynamic dispatch.
type IWhen_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhen_clauseContext differentiates from other interfaces.
	IsWhen_clauseContext()
}

type When_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhen_clauseContext() *When_clauseContext {
	var p = new(When_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_when_clause
	return p
}

func (*When_clauseContext) IsWhen_clauseContext() {}

func NewWhen_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *When_clauseContext {
	var p = new(When_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_when_clause

	return p
}

func (s *When_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *When_clauseContext) WHEN() antlr.TerminalNode {
	return s.GetToken(mdxParserWHEN, 0)
}

func (s *When_clauseContext) AllValue_expression() []IValue_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValue_expressionContext)(nil)).Elem())
	var tst = make([]IValue_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValue_expressionContext)
		}
	}

	return tst
}

func (s *When_clauseContext) Value_expression(i int) IValue_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValue_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValue_expressionContext)
}

func (s *When_clauseContext) THEN() antlr.TerminalNode {
	return s.GetToken(mdxParserTHEN, 0)
}

func (s *When_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *When_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *When_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterWhen_clause(s)
	}
}

func (s *When_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitWhen_clause(s)
	}
}

func (p *mdxParser) When_clause() (localctx IWhen_clauseContext) {
	localctx = NewWhen_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, mdxParserRULE_when_clause)

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
		p.SetState(369)
		p.Match(mdxParserWHEN)
	}
	{
		p.SetState(370)
		p.Value_expression()
	}
	{
		p.SetState(371)
		p.Match(mdxParserTHEN)
	}
	{
		p.SetState(372)
		p.Value_expression()
	}

	return localctx
}

// IComp_opContext is an interface to support dynamic dispatch.
type IComp_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComp_opContext differentiates from other interfaces.
	IsComp_opContext()
}

type Comp_opContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComp_opContext() *Comp_opContext {
	var p = new(Comp_opContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_comp_op
	return p
}

func (*Comp_opContext) IsComp_opContext() {}

func NewComp_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comp_opContext {
	var p = new(Comp_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_comp_op

	return p
}

func (s *Comp_opContext) GetParser() antlr.Parser { return s.parser }

func (s *Comp_opContext) EQ() antlr.TerminalNode {
	return s.GetToken(mdxParserEQ, 0)
}

func (s *Comp_opContext) NE() antlr.TerminalNode {
	return s.GetToken(mdxParserNE, 0)
}

func (s *Comp_opContext) LT() antlr.TerminalNode {
	return s.GetToken(mdxParserLT, 0)
}

func (s *Comp_opContext) GT() antlr.TerminalNode {
	return s.GetToken(mdxParserGT, 0)
}

func (s *Comp_opContext) LE() antlr.TerminalNode {
	return s.GetToken(mdxParserLE, 0)
}

func (s *Comp_opContext) GE() antlr.TerminalNode {
	return s.GetToken(mdxParserGE, 0)
}

func (s *Comp_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comp_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comp_opContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterComp_op(s)
	}
}

func (s *Comp_opContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitComp_op(s)
	}
}

func (p *mdxParser) Comp_op() (localctx IComp_opContext) {
	localctx = NewComp_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, mdxParserRULE_comp_op)
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
		p.SetState(374)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<mdxParserEQ)|(1<<mdxParserGE)|(1<<mdxParserGT)|(1<<mdxParserLE)|(1<<mdxParserLT)|(1<<mdxParserNE))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = mdxParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Unquoted_identifier() IUnquoted_identifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnquoted_identifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnquoted_identifierContext)
}

func (s *IdentifierContext) Quoted_identifier() IQuoted_identifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuoted_identifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuoted_identifierContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *mdxParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, mdxParserRULE_identifier)

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
	p.SetState(378)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mdxParserDIMENSION, mdxParserPROPERTIES, mdxParserID:
		{
			p.SetState(376)
			p.Unquoted_identifier()
		}

	case mdxParserQUOTED_ID:
		{
			p.SetState(377)
			p.Quoted_identifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnquoted_identifierContext is an interface to support dynamic dispatch.
type IUnquoted_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnquoted_identifierContext differentiates from other interfaces.
	IsUnquoted_identifierContext()
}

type Unquoted_identifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnquoted_identifierContext() *Unquoted_identifierContext {
	var p = new(Unquoted_identifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_unquoted_identifier
	return p
}

func (*Unquoted_identifierContext) IsUnquoted_identifierContext() {}

func NewUnquoted_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unquoted_identifierContext {
	var p = new(Unquoted_identifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_unquoted_identifier

	return p
}

func (s *Unquoted_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Unquoted_identifierContext) Keyword() IKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *Unquoted_identifierContext) ID() antlr.TerminalNode {
	return s.GetToken(mdxParserID, 0)
}

func (s *Unquoted_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unquoted_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unquoted_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterUnquoted_identifier(s)
	}
}

func (s *Unquoted_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitUnquoted_identifier(s)
	}
}

func (p *mdxParser) Unquoted_identifier() (localctx IUnquoted_identifierContext) {
	localctx = NewUnquoted_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, mdxParserRULE_unquoted_identifier)

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

	p.SetState(382)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case mdxParserDIMENSION, mdxParserPROPERTIES:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(380)
			p.Keyword()
		}

	case mdxParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(381)
			p.Match(mdxParserID)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAmp_quoted_identifierContext is an interface to support dynamic dispatch.
type IAmp_quoted_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAmp_quoted_identifierContext differentiates from other interfaces.
	IsAmp_quoted_identifierContext()
}

type Amp_quoted_identifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAmp_quoted_identifierContext() *Amp_quoted_identifierContext {
	var p = new(Amp_quoted_identifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_amp_quoted_identifier
	return p
}

func (*Amp_quoted_identifierContext) IsAmp_quoted_identifierContext() {}

func NewAmp_quoted_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Amp_quoted_identifierContext {
	var p = new(Amp_quoted_identifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_amp_quoted_identifier

	return p
}

func (s *Amp_quoted_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Amp_quoted_identifierContext) AMP_QUOTED_ID() antlr.TerminalNode {
	return s.GetToken(mdxParserAMP_QUOTED_ID, 0)
}

func (s *Amp_quoted_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Amp_quoted_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Amp_quoted_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterAmp_quoted_identifier(s)
	}
}

func (s *Amp_quoted_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitAmp_quoted_identifier(s)
	}
}

func (p *mdxParser) Amp_quoted_identifier() (localctx IAmp_quoted_identifierContext) {
	localctx = NewAmp_quoted_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, mdxParserRULE_amp_quoted_identifier)

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
		p.SetState(384)
		p.Match(mdxParserAMP_QUOTED_ID)
	}

	return localctx
}

// IQuoted_identifierContext is an interface to support dynamic dispatch.
type IQuoted_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuoted_identifierContext differentiates from other interfaces.
	IsQuoted_identifierContext()
}

type Quoted_identifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuoted_identifierContext() *Quoted_identifierContext {
	var p = new(Quoted_identifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_quoted_identifier
	return p
}

func (*Quoted_identifierContext) IsQuoted_identifierContext() {}

func NewQuoted_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Quoted_identifierContext {
	var p = new(Quoted_identifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_quoted_identifier

	return p
}

func (s *Quoted_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Quoted_identifierContext) QUOTED_ID() antlr.TerminalNode {
	return s.GetToken(mdxParserQUOTED_ID, 0)
}

func (s *Quoted_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Quoted_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Quoted_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterQuoted_identifier(s)
	}
}

func (s *Quoted_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitQuoted_identifier(s)
	}
}

func (p *mdxParser) Quoted_identifier() (localctx IQuoted_identifierContext) {
	localctx = NewQuoted_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, mdxParserRULE_quoted_identifier)

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
		p.SetState(386)
		p.Match(mdxParserQUOTED_ID)
	}

	return localctx
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = mdxParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = mdxParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) DIMENSION() antlr.TerminalNode {
	return s.GetToken(mdxParserDIMENSION, 0)
}

func (s *KeywordContext) PROPERTIES() antlr.TerminalNode {
	return s.GetToken(mdxParserPROPERTIES, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(mdxListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *mdxParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, mdxParserRULE_keyword)
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
		p.SetState(388)
		_la = p.GetTokenStream().LA(1)

		if !(_la == mdxParserDIMENSION || _la == mdxParserPROPERTIES) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
