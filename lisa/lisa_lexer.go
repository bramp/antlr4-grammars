// Code generated from lisa.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lisa

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 50, 318,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27,
	3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3,
	40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 7, 45, 280,
	10, 45, 12, 45, 14, 45, 283, 11, 45, 3, 46, 3, 46, 7, 46, 287, 10, 46,
	12, 46, 14, 46, 290, 11, 46, 3, 47, 3, 47, 7, 47, 294, 10, 47, 12, 47,
	14, 47, 297, 11, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 7, 48, 305,
	10, 48, 12, 48, 14, 48, 308, 11, 48, 3, 48, 3, 48, 3, 49, 6, 49, 313, 10,
	49, 13, 49, 14, 49, 314, 3, 49, 3, 49, 3, 295, 2, 50, 3, 3, 5, 4, 7, 5,
	9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63,
	33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81,
	42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 3,
	2, 8, 3, 2, 51, 59, 3, 2, 50, 59, 4, 2, 67, 92, 99, 124, 6, 2, 50, 59,
	67, 92, 97, 97, 99, 124, 4, 2, 12, 12, 15, 15, 5, 2, 11, 12, 15, 15, 34,
	34, 2, 322, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
	3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2,
	17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2,
	2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2,
	2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2,
	2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3,
	2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55,
	3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2,
	63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2,
	2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2,
	2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2,
	2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3,
	2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 3, 99, 3, 2, 2, 2, 5, 107,
	3, 2, 2, 2, 7, 109, 3, 2, 2, 2, 9, 111, 3, 2, 2, 2, 11, 113, 3, 2, 2, 2,
	13, 117, 3, 2, 2, 2, 15, 121, 3, 2, 2, 2, 17, 125, 3, 2, 2, 2, 19, 131,
	3, 2, 2, 2, 21, 136, 3, 2, 2, 2, 23, 143, 3, 2, 2, 2, 25, 151, 3, 2, 2,
	2, 27, 157, 3, 2, 2, 2, 29, 166, 3, 2, 2, 2, 31, 175, 3, 2, 2, 2, 33, 177,
	3, 2, 2, 2, 35, 179, 3, 2, 2, 2, 37, 181, 3, 2, 2, 2, 39, 188, 3, 2, 2,
	2, 41, 198, 3, 2, 2, 2, 43, 201, 3, 2, 2, 2, 45, 203, 3, 2, 2, 2, 47, 206,
	3, 2, 2, 2, 49, 209, 3, 2, 2, 2, 51, 212, 3, 2, 2, 2, 53, 215, 3, 2, 2,
	2, 55, 218, 3, 2, 2, 2, 57, 221, 3, 2, 2, 2, 59, 223, 3, 2, 2, 2, 61, 226,
	3, 2, 2, 2, 63, 229, 3, 2, 2, 2, 65, 232, 3, 2, 2, 2, 67, 235, 3, 2, 2,
	2, 69, 237, 3, 2, 2, 2, 71, 239, 3, 2, 2, 2, 73, 241, 3, 2, 2, 2, 75, 243,
	3, 2, 2, 2, 77, 245, 3, 2, 2, 2, 79, 247, 3, 2, 2, 2, 81, 252, 3, 2, 2,
	2, 83, 260, 3, 2, 2, 2, 85, 265, 3, 2, 2, 2, 87, 271, 3, 2, 2, 2, 89, 277,
	3, 2, 2, 2, 91, 284, 3, 2, 2, 2, 93, 291, 3, 2, 2, 2, 95, 300, 3, 2, 2,
	2, 97, 312, 3, 2, 2, 2, 99, 100, 7, 102, 2, 2, 100, 101, 7, 103, 2, 2,
	101, 102, 7, 101, 2, 2, 102, 103, 7, 110, 2, 2, 103, 104, 7, 99, 2, 2,
	104, 105, 7, 116, 2, 2, 105, 106, 7, 103, 2, 2, 106, 4, 3, 2, 2, 2, 107,
	108, 7, 125, 2, 2, 108, 6, 3, 2, 2, 2, 109, 110, 7, 127, 2, 2, 110, 8,
	3, 2, 2, 2, 111, 112, 7, 61, 2, 2, 112, 10, 3, 2, 2, 2, 113, 114, 7, 107,
	2, 2, 114, 115, 7, 112, 2, 2, 115, 116, 7, 118, 2, 2, 116, 12, 3, 2, 2,
	2, 117, 118, 7, 102, 2, 2, 118, 119, 7, 104, 2, 2, 119, 120, 7, 99, 2,
	2, 120, 14, 3, 2, 2, 2, 121, 122, 7, 112, 2, 2, 122, 123, 7, 104, 2, 2,
	123, 124, 7, 99, 2, 2, 124, 16, 3, 2, 2, 2, 125, 126, 7, 116, 2, 2, 126,
	127, 7, 103, 2, 2, 127, 128, 7, 105, 2, 2, 128, 129, 7, 103, 2, 2, 129,
	130, 7, 122, 2, 2, 130, 18, 3, 2, 2, 2, 131, 132, 7, 100, 2, 2, 132, 133,
	7, 113, 2, 2, 133, 134, 7, 113, 2, 2, 134, 135, 7, 110, 2, 2, 135, 20,
	3, 2, 2, 2, 136, 137, 7, 117, 2, 2, 137, 138, 7, 118, 2, 2, 138, 139, 7,
	116, 2, 2, 139, 140, 7, 107, 2, 2, 140, 141, 7, 112, 2, 2, 141, 142, 7,
	105, 2, 2, 142, 22, 3, 2, 2, 2, 143, 144, 7, 114, 2, 2, 144, 145, 7, 116,
	2, 2, 145, 146, 7, 113, 2, 2, 146, 147, 7, 105, 2, 2, 147, 148, 7, 116,
	2, 2, 148, 149, 7, 99, 2, 2, 149, 150, 7, 111, 2, 2, 150, 24, 3, 2, 2,
	2, 151, 152, 7, 100, 2, 2, 152, 153, 7, 116, 2, 2, 153, 154, 7, 103, 2,
	2, 154, 155, 7, 99, 2, 2, 155, 156, 7, 109, 2, 2, 156, 26, 3, 2, 2, 2,
	157, 158, 7, 101, 2, 2, 158, 159, 7, 113, 2, 2, 159, 160, 7, 112, 2, 2,
	160, 161, 7, 118, 2, 2, 161, 162, 7, 107, 2, 2, 162, 163, 7, 112, 2, 2,
	163, 164, 7, 119, 2, 2, 164, 165, 7, 103, 2, 2, 165, 28, 3, 2, 2, 2, 166,
	167, 7, 105, 2, 2, 167, 168, 7, 103, 2, 2, 168, 169, 7, 112, 2, 2, 169,
	170, 7, 103, 2, 2, 170, 171, 7, 116, 2, 2, 171, 172, 7, 99, 2, 2, 172,
	173, 7, 118, 2, 2, 173, 174, 7, 103, 2, 2, 174, 30, 3, 2, 2, 2, 175, 176,
	7, 42, 2, 2, 176, 32, 3, 2, 2, 2, 177, 178, 7, 46, 2, 2, 178, 34, 3, 2,
	2, 2, 179, 180, 7, 43, 2, 2, 180, 36, 3, 2, 2, 2, 181, 182, 7, 116, 2,
	2, 182, 183, 7, 99, 2, 2, 183, 184, 7, 112, 2, 2, 184, 185, 7, 102, 2,
	2, 185, 186, 7, 113, 2, 2, 186, 187, 7, 111, 2, 2, 187, 38, 3, 2, 2, 2,
	188, 189, 7, 103, 2, 2, 189, 190, 7, 112, 2, 2, 190, 191, 7, 119, 2, 2,
	191, 192, 7, 111, 2, 2, 192, 193, 7, 103, 2, 2, 193, 194, 7, 116, 2, 2,
	194, 195, 7, 99, 2, 2, 195, 196, 7, 118, 2, 2, 196, 197, 7, 103, 2, 2,
	197, 40, 3, 2, 2, 2, 198, 199, 7, 107, 2, 2, 199, 200, 7, 104, 2, 2, 200,
	42, 3, 2, 2, 2, 201, 202, 7, 63, 2, 2, 202, 44, 3, 2, 2, 2, 203, 204, 7,
	47, 2, 2, 204, 205, 7, 45, 2, 2, 205, 46, 3, 2, 2, 2, 206, 207, 7, 44,
	2, 2, 207, 208, 7, 63, 2, 2, 208, 48, 3, 2, 2, 2, 209, 210, 7, 49, 2, 2,
	210, 211, 7, 63, 2, 2, 211, 50, 3, 2, 2, 2, 212, 213, 7, 45, 2, 2, 213,
	214, 7, 63, 2, 2, 214, 52, 3, 2, 2, 2, 215, 216, 7, 126, 2, 2, 216, 217,
	7, 126, 2, 2, 217, 54, 3, 2, 2, 2, 218, 219, 7, 40, 2, 2, 219, 220, 7,
	40, 2, 2, 220, 56, 3, 2, 2, 2, 221, 222, 7, 35, 2, 2, 222, 58, 3, 2, 2,
	2, 223, 224, 7, 62, 2, 2, 224, 225, 7, 63, 2, 2, 225, 60, 3, 2, 2, 2, 226,
	227, 7, 64, 2, 2, 227, 228, 7, 63, 2, 2, 228, 62, 3, 2, 2, 2, 229, 230,
	7, 63, 2, 2, 230, 231, 7, 63, 2, 2, 231, 64, 3, 2, 2, 2, 232, 233, 7, 35,
	2, 2, 233, 234, 7, 63, 2, 2, 234, 66, 3, 2, 2, 2, 235, 236, 7, 64, 2, 2,
	236, 68, 3, 2, 2, 2, 237, 238, 7, 47, 2, 2, 238, 70, 3, 2, 2, 2, 239, 240,
	7, 45, 2, 2, 240, 72, 3, 2, 2, 2, 241, 242, 7, 44, 2, 2, 242, 74, 3, 2,
	2, 2, 243, 244, 7, 49, 2, 2, 244, 76, 3, 2, 2, 2, 245, 246, 7, 39, 2, 2,
	246, 78, 3, 2, 2, 2, 247, 248, 7, 112, 2, 2, 248, 249, 7, 103, 2, 2, 249,
	250, 7, 122, 2, 2, 250, 251, 7, 118, 2, 2, 251, 80, 3, 2, 2, 2, 252, 253,
	7, 106, 2, 2, 253, 254, 7, 99, 2, 2, 254, 255, 7, 117, 2, 2, 255, 256,
	7, 112, 2, 2, 256, 257, 7, 103, 2, 2, 257, 258, 7, 122, 2, 2, 258, 259,
	7, 118, 2, 2, 259, 82, 3, 2, 2, 2, 260, 261, 7, 86, 2, 2, 261, 262, 7,
	84, 2, 2, 262, 263, 7, 87, 2, 2, 263, 264, 7, 71, 2, 2, 264, 84, 3, 2,
	2, 2, 265, 266, 7, 72, 2, 2, 266, 267, 7, 67, 2, 2, 267, 268, 7, 78, 2,
	2, 268, 269, 7, 85, 2, 2, 269, 270, 7, 71, 2, 2, 270, 86, 3, 2, 2, 2, 271,
	272, 7, 89, 2, 2, 272, 273, 7, 74, 2, 2, 273, 274, 7, 75, 2, 2, 274, 275,
	7, 78, 2, 2, 275, 276, 7, 71, 2, 2, 276, 88, 3, 2, 2, 2, 277, 281, 9, 2,
	2, 2, 278, 280, 9, 3, 2, 2, 279, 278, 3, 2, 2, 2, 280, 283, 3, 2, 2, 2,
	281, 279, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 90, 3, 2, 2, 2, 283, 281,
	3, 2, 2, 2, 284, 288, 9, 4, 2, 2, 285, 287, 9, 5, 2, 2, 286, 285, 3, 2,
	2, 2, 287, 290, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2,
	289, 92, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 291, 295, 7, 36, 2, 2, 292,
	294, 11, 2, 2, 2, 293, 292, 3, 2, 2, 2, 294, 297, 3, 2, 2, 2, 295, 296,
	3, 2, 2, 2, 295, 293, 3, 2, 2, 2, 296, 298, 3, 2, 2, 2, 297, 295, 3, 2,
	2, 2, 298, 299, 7, 36, 2, 2, 299, 94, 3, 2, 2, 2, 300, 301, 7, 49, 2, 2,
	301, 302, 7, 49, 2, 2, 302, 306, 3, 2, 2, 2, 303, 305, 10, 6, 2, 2, 304,
	303, 3, 2, 2, 2, 305, 308, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306, 307,
	3, 2, 2, 2, 307, 309, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 309, 310, 8, 48,
	2, 2, 310, 96, 3, 2, 2, 2, 311, 313, 9, 7, 2, 2, 312, 311, 3, 2, 2, 2,
	313, 314, 3, 2, 2, 2, 314, 312, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315,
	316, 3, 2, 2, 2, 316, 317, 8, 49, 2, 2, 317, 98, 3, 2, 2, 2, 8, 2, 281,
	288, 295, 306, 314, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'declare'", "'{'", "'}'", "';'", "'int'", "'dfa'", "'nfa'", "'regex'",
	"'bool'", "'string'", "'program'", "'break'", "'continue'", "'generate'",
	"'('", "','", "')'", "'random'", "'enumerate'", "'if'", "'='", "'-+'",
	"'*='", "'/='", "'+='", "'||'", "'&&'", "'!'", "'<='", "'>='", "'=='",
	"'!='", "'>'", "'-'", "'+'", "'*'", "'/'", "'%'", "'next'", "'hasnext'",
	"'TRUE'", "'FALSE'", "'WHILE'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "TRUE", "FALSE", "WHILE", "INTEGER", "ID", "STRINGLITERAL",
	"COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "TRUE",
	"FALSE", "WHILE", "INTEGER", "ID", "STRINGLITERAL", "COMMENT", "WS",
}

type lisaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewlisaLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *lisaLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewlisaLexer(input antlr.CharStream) *lisaLexer {
	l := new(lisaLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "lisa.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// lisaLexer tokens.
const (
	lisaLexerT__0          = 1
	lisaLexerT__1          = 2
	lisaLexerT__2          = 3
	lisaLexerT__3          = 4
	lisaLexerT__4          = 5
	lisaLexerT__5          = 6
	lisaLexerT__6          = 7
	lisaLexerT__7          = 8
	lisaLexerT__8          = 9
	lisaLexerT__9          = 10
	lisaLexerT__10         = 11
	lisaLexerT__11         = 12
	lisaLexerT__12         = 13
	lisaLexerT__13         = 14
	lisaLexerT__14         = 15
	lisaLexerT__15         = 16
	lisaLexerT__16         = 17
	lisaLexerT__17         = 18
	lisaLexerT__18         = 19
	lisaLexerT__19         = 20
	lisaLexerT__20         = 21
	lisaLexerT__21         = 22
	lisaLexerT__22         = 23
	lisaLexerT__23         = 24
	lisaLexerT__24         = 25
	lisaLexerT__25         = 26
	lisaLexerT__26         = 27
	lisaLexerT__27         = 28
	lisaLexerT__28         = 29
	lisaLexerT__29         = 30
	lisaLexerT__30         = 31
	lisaLexerT__31         = 32
	lisaLexerT__32         = 33
	lisaLexerT__33         = 34
	lisaLexerT__34         = 35
	lisaLexerT__35         = 36
	lisaLexerT__36         = 37
	lisaLexerT__37         = 38
	lisaLexerT__38         = 39
	lisaLexerT__39         = 40
	lisaLexerTRUE          = 41
	lisaLexerFALSE         = 42
	lisaLexerWHILE         = 43
	lisaLexerINTEGER       = 44
	lisaLexerID            = 45
	lisaLexerSTRINGLITERAL = 46
	lisaLexerCOMMENT       = 47
	lisaLexerWS            = 48
)
