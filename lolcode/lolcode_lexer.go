// Code generated from lolcode.g4 by ANTLR 4.9.3. DO NOT EDIT.

package lolcode

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 46, 421,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3,
	41, 3, 41, 3, 42, 6, 42, 371, 10, 42, 13, 42, 14, 42, 372, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 6,
	43, 387, 10, 43, 13, 43, 14, 43, 388, 3, 43, 7, 43, 392, 10, 43, 12, 43,
	14, 43, 395, 11, 43, 3, 43, 3, 43, 7, 43, 399, 10, 43, 12, 43, 14, 43,
	402, 11, 43, 3, 43, 5, 43, 405, 10, 43, 3, 44, 3, 44, 3, 44, 3, 44, 7,
	44, 411, 10, 44, 12, 44, 14, 44, 414, 11, 44, 3, 44, 3, 44, 3, 45, 3, 45,
	3, 45, 3, 45, 2, 2, 46, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35,
	19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53,
	28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71,
	37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89,
	46, 3, 2, 5, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 36, 36, 5, 2,
	12, 12, 15, 15, 34, 34, 2, 431, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2,
	2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3,
	2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45,
	3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2,
	53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2,
	2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2,
	2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2,
	2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3,
	2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 3, 91,
	3, 2, 2, 2, 5, 95, 3, 2, 2, 2, 7, 103, 3, 2, 2, 2, 9, 112, 3, 2, 2, 2,
	11, 117, 3, 2, 2, 2, 13, 129, 3, 2, 2, 2, 15, 137, 3, 2, 2, 2, 17, 141,
	3, 2, 2, 2, 19, 145, 3, 2, 2, 2, 21, 150, 3, 2, 2, 2, 23, 155, 3, 2, 2,
	2, 25, 163, 3, 2, 2, 2, 27, 169, 3, 2, 2, 2, 29, 176, 3, 2, 2, 2, 31, 183,
	3, 2, 2, 2, 33, 187, 3, 2, 2, 2, 35, 193, 3, 2, 2, 2, 37, 200, 3, 2, 2,
	2, 39, 207, 3, 2, 2, 2, 41, 217, 3, 2, 2, 2, 43, 220, 3, 2, 2, 2, 45, 226,
	3, 2, 2, 2, 47, 238, 3, 2, 2, 2, 49, 240, 3, 2, 2, 2, 51, 250, 3, 2, 2,
	2, 53, 253, 3, 2, 2, 2, 55, 262, 3, 2, 2, 2, 57, 270, 3, 2, 2, 2, 59, 280,
	3, 2, 2, 2, 61, 289, 3, 2, 2, 2, 63, 299, 3, 2, 2, 2, 65, 306, 3, 2, 2,
	2, 67, 314, 3, 2, 2, 2, 69, 325, 3, 2, 2, 2, 71, 337, 3, 2, 2, 2, 73, 344,
	3, 2, 2, 2, 75, 349, 3, 2, 2, 2, 77, 351, 3, 2, 2, 2, 79, 358, 3, 2, 2,
	2, 81, 365, 3, 2, 2, 2, 83, 370, 3, 2, 2, 2, 85, 404, 3, 2, 2, 2, 87, 406,
	3, 2, 2, 2, 89, 417, 3, 2, 2, 2, 91, 92, 7, 74, 2, 2, 92, 93, 7, 67, 2,
	2, 93, 94, 7, 75, 2, 2, 94, 4, 3, 2, 2, 2, 95, 96, 7, 77, 2, 2, 96, 97,
	7, 86, 2, 2, 97, 98, 7, 74, 2, 2, 98, 99, 7, 90, 2, 2, 99, 100, 7, 68,
	2, 2, 100, 101, 7, 91, 2, 2, 101, 102, 7, 71, 2, 2, 102, 6, 3, 2, 2, 2,
	103, 104, 7, 75, 2, 2, 104, 105, 7, 79, 2, 2, 105, 106, 7, 34, 2, 2, 106,
	107, 7, 75, 2, 2, 107, 108, 7, 80, 2, 2, 108, 109, 7, 34, 2, 2, 109, 110,
	7, 91, 2, 2, 110, 111, 7, 84, 2, 2, 111, 8, 3, 2, 2, 2, 112, 113, 7, 89,
	2, 2, 113, 114, 7, 75, 2, 2, 114, 115, 7, 78, 2, 2, 115, 116, 7, 71, 2,
	2, 116, 10, 3, 2, 2, 2, 117, 118, 7, 75, 2, 2, 118, 119, 7, 79, 2, 2, 119,
	120, 7, 34, 2, 2, 120, 121, 7, 81, 2, 2, 121, 122, 7, 87, 2, 2, 122, 123,
	7, 86, 2, 2, 123, 124, 7, 86, 2, 2, 124, 125, 7, 67, 2, 2, 125, 126, 7,
	34, 2, 2, 126, 127, 7, 91, 2, 2, 127, 128, 7, 84, 2, 2, 128, 12, 3, 2,
	2, 2, 129, 130, 7, 75, 2, 2, 130, 131, 7, 34, 2, 2, 131, 132, 7, 74, 2,
	2, 132, 133, 7, 67, 2, 2, 133, 134, 7, 85, 2, 2, 134, 135, 7, 34, 2, 2,
	135, 136, 7, 67, 2, 2, 136, 14, 3, 2, 2, 2, 137, 138, 7, 75, 2, 2, 138,
	139, 7, 86, 2, 2, 139, 140, 7, 92, 2, 2, 140, 16, 3, 2, 2, 2, 141, 142,
	7, 68, 2, 2, 142, 143, 7, 86, 2, 2, 143, 144, 7, 89, 2, 2, 144, 18, 3,
	2, 2, 2, 145, 146, 7, 81, 2, 2, 146, 147, 7, 68, 2, 2, 147, 148, 7, 86,
	2, 2, 148, 149, 7, 89, 2, 2, 149, 20, 3, 2, 2, 2, 150, 151, 7, 86, 2, 2,
	151, 152, 7, 78, 2, 2, 152, 153, 7, 70, 2, 2, 153, 154, 7, 84, 2, 2, 154,
	22, 3, 2, 2, 2, 155, 156, 7, 88, 2, 2, 156, 157, 7, 75, 2, 2, 157, 158,
	7, 85, 2, 2, 158, 159, 7, 75, 2, 2, 159, 160, 7, 68, 2, 2, 160, 161, 7,
	78, 2, 2, 161, 162, 7, 71, 2, 2, 162, 24, 3, 2, 2, 2, 163, 164, 7, 79,
	2, 2, 164, 165, 7, 77, 2, 2, 165, 166, 7, 67, 2, 2, 166, 167, 7, 91, 2,
	2, 167, 168, 7, 65, 2, 2, 168, 26, 3, 2, 2, 2, 169, 170, 7, 81, 2, 2, 170,
	171, 7, 34, 2, 2, 171, 172, 7, 84, 2, 2, 172, 173, 7, 78, 2, 2, 173, 174,
	7, 91, 2, 2, 174, 175, 7, 65, 2, 2, 175, 28, 3, 2, 2, 2, 176, 177, 7, 91,
	2, 2, 177, 178, 7, 67, 2, 2, 178, 179, 7, 34, 2, 2, 179, 180, 7, 84, 2,
	2, 180, 181, 7, 78, 2, 2, 181, 182, 7, 91, 2, 2, 182, 30, 3, 2, 2, 2, 183,
	184, 7, 81, 2, 2, 184, 185, 7, 75, 2, 2, 185, 186, 7, 69, 2, 2, 186, 32,
	3, 2, 2, 2, 187, 188, 7, 79, 2, 2, 188, 189, 7, 71, 2, 2, 189, 190, 7,
	68, 2, 2, 190, 191, 7, 68, 2, 2, 191, 192, 7, 71, 2, 2, 192, 34, 3, 2,
	2, 2, 193, 194, 7, 80, 2, 2, 194, 195, 7, 81, 2, 2, 195, 196, 7, 34, 2,
	2, 196, 197, 7, 89, 2, 2, 197, 198, 7, 67, 2, 2, 198, 199, 7, 75, 2, 2,
	199, 36, 3, 2, 2, 2, 200, 201, 7, 73, 2, 2, 201, 202, 7, 75, 2, 2, 202,
	203, 7, 79, 2, 2, 203, 204, 7, 79, 2, 2, 204, 205, 7, 71, 2, 2, 205, 206,
	7, 74, 2, 2, 206, 38, 3, 2, 2, 2, 207, 208, 7, 74, 2, 2, 208, 209, 7, 81,
	2, 2, 209, 210, 7, 89, 2, 2, 210, 211, 7, 34, 2, 2, 211, 212, 7, 70, 2,
	2, 212, 213, 7, 87, 2, 2, 213, 214, 7, 92, 2, 2, 214, 215, 7, 34, 2, 2,
	215, 216, 7, 75, 2, 2, 216, 40, 3, 2, 2, 2, 217, 218, 7, 91, 2, 2, 218,
	219, 7, 84, 2, 2, 219, 42, 3, 2, 2, 2, 220, 221, 7, 67, 2, 2, 221, 222,
	7, 80, 2, 2, 222, 223, 7, 34, 2, 2, 223, 224, 7, 91, 2, 2, 224, 225, 7,
	84, 2, 2, 225, 44, 3, 2, 2, 2, 226, 227, 7, 75, 2, 2, 227, 228, 7, 72,
	2, 2, 228, 229, 7, 34, 2, 2, 229, 230, 7, 87, 2, 2, 230, 231, 7, 34, 2,
	2, 231, 232, 7, 85, 2, 2, 232, 233, 7, 67, 2, 2, 233, 234, 7, 91, 2, 2,
	234, 235, 7, 34, 2, 2, 235, 236, 7, 85, 2, 2, 236, 237, 7, 81, 2, 2, 237,
	46, 3, 2, 2, 2, 238, 239, 7, 84, 2, 2, 239, 48, 3, 2, 2, 2, 240, 241, 7,
	68, 2, 2, 241, 242, 7, 81, 2, 2, 242, 243, 7, 86, 2, 2, 243, 244, 7, 74,
	2, 2, 244, 245, 7, 34, 2, 2, 245, 246, 7, 85, 2, 2, 246, 247, 7, 67, 2,
	2, 247, 248, 7, 71, 2, 2, 248, 249, 7, 79, 2, 2, 249, 50, 3, 2, 2, 2, 250,
	251, 7, 67, 2, 2, 251, 252, 7, 80, 2, 2, 252, 52, 3, 2, 2, 2, 253, 254,
	7, 70, 2, 2, 254, 255, 7, 75, 2, 2, 255, 256, 7, 72, 2, 2, 256, 257, 7,
	72, 2, 2, 257, 258, 7, 84, 2, 2, 258, 259, 7, 75, 2, 2, 259, 260, 7, 80,
	2, 2, 260, 261, 7, 86, 2, 2, 261, 54, 3, 2, 2, 2, 262, 263, 7, 68, 2, 2,
	263, 264, 7, 81, 2, 2, 264, 265, 7, 86, 2, 2, 265, 266, 7, 74, 2, 2, 266,
	267, 7, 34, 2, 2, 267, 268, 7, 81, 2, 2, 268, 269, 7, 72, 2, 2, 269, 56,
	3, 2, 2, 2, 270, 271, 7, 71, 2, 2, 271, 272, 7, 75, 2, 2, 272, 273, 7,
	86, 2, 2, 273, 274, 7, 74, 2, 2, 274, 275, 7, 71, 2, 2, 275, 276, 7, 84,
	2, 2, 276, 277, 7, 34, 2, 2, 277, 278, 7, 81, 2, 2, 278, 279, 7, 72, 2,
	2, 279, 58, 3, 2, 2, 2, 280, 281, 7, 68, 2, 2, 281, 282, 7, 75, 2, 2, 282,
	283, 7, 73, 2, 2, 283, 284, 7, 73, 2, 2, 284, 285, 7, 84, 2, 2, 285, 286,
	7, 34, 2, 2, 286, 287, 7, 81, 2, 2, 287, 288, 7, 72, 2, 2, 288, 60, 3,
	2, 2, 2, 289, 290, 7, 85, 2, 2, 290, 291, 7, 79, 2, 2, 291, 292, 7, 67,
	2, 2, 292, 293, 7, 78, 2, 2, 293, 294, 7, 78, 2, 2, 294, 295, 7, 84, 2,
	2, 295, 296, 7, 34, 2, 2, 296, 297, 7, 81, 2, 2, 297, 298, 7, 72, 2, 2,
	298, 62, 3, 2, 2, 2, 299, 300, 7, 85, 2, 2, 300, 301, 7, 87, 2, 2, 301,
	302, 7, 79, 2, 2, 302, 303, 7, 34, 2, 2, 303, 304, 7, 81, 2, 2, 304, 305,
	7, 72, 2, 2, 305, 64, 3, 2, 2, 2, 306, 307, 7, 70, 2, 2, 307, 308, 7, 75,
	2, 2, 308, 309, 7, 72, 2, 2, 309, 310, 7, 72, 2, 2, 310, 311, 7, 34, 2,
	2, 311, 312, 7, 81, 2, 2, 312, 313, 7, 72, 2, 2, 313, 66, 3, 2, 2, 2, 314,
	315, 7, 82, 2, 2, 315, 316, 7, 84, 2, 2, 316, 317, 7, 81, 2, 2, 317, 318,
	7, 70, 2, 2, 318, 319, 7, 87, 2, 2, 319, 320, 7, 77, 2, 2, 320, 321, 7,
	86, 2, 2, 321, 322, 7, 34, 2, 2, 322, 323, 7, 81, 2, 2, 323, 324, 7, 72,
	2, 2, 324, 68, 3, 2, 2, 2, 325, 326, 7, 83, 2, 2, 326, 327, 7, 87, 2, 2,
	327, 328, 7, 81, 2, 2, 328, 329, 7, 85, 2, 2, 329, 330, 7, 74, 2, 2, 330,
	331, 7, 87, 2, 2, 331, 332, 7, 80, 2, 2, 332, 333, 7, 86, 2, 2, 333, 334,
	7, 34, 2, 2, 334, 335, 7, 81, 2, 2, 335, 336, 7, 72, 2, 2, 336, 70, 3,
	2, 2, 2, 337, 338, 7, 79, 2, 2, 338, 339, 7, 81, 2, 2, 339, 340, 7, 70,
	2, 2, 340, 341, 7, 34, 2, 2, 341, 342, 7, 81, 2, 2, 342, 343, 7, 72, 2,
	2, 343, 72, 3, 2, 2, 2, 344, 345, 7, 79, 2, 2, 345, 346, 7, 67, 2, 2, 346,
	347, 7, 71, 2, 2, 347, 348, 7, 77, 2, 2, 348, 74, 3, 2, 2, 2, 349, 350,
	7, 67, 2, 2, 350, 76, 3, 2, 2, 2, 351, 352, 7, 67, 2, 2, 352, 353, 7, 78,
	2, 2, 353, 354, 7, 78, 2, 2, 354, 355, 7, 34, 2, 2, 355, 356, 7, 81, 2,
	2, 356, 357, 7, 72, 2, 2, 357, 78, 3, 2, 2, 2, 358, 359, 7, 67, 2, 2, 359,
	360, 7, 80, 2, 2, 360, 361, 7, 91, 2, 2, 361, 362, 7, 34, 2, 2, 362, 363,
	7, 81, 2, 2, 363, 364, 7, 72, 2, 2, 364, 80, 3, 2, 2, 2, 365, 366, 7, 80,
	2, 2, 366, 367, 7, 81, 2, 2, 367, 368, 7, 86, 2, 2, 368, 82, 3, 2, 2, 2,
	369, 371, 9, 2, 2, 2, 370, 369, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372,
	370, 3, 2, 2, 2, 372, 373, 3, 2, 2, 2, 373, 84, 3, 2, 2, 2, 374, 375, 7,
	89, 2, 2, 375, 376, 7, 75, 2, 2, 376, 405, 7, 80, 2, 2, 377, 378, 7, 72,
	2, 2, 378, 379, 7, 67, 2, 2, 379, 380, 7, 75, 2, 2, 380, 405, 7, 78, 2,
	2, 381, 382, 7, 80, 2, 2, 382, 383, 7, 81, 2, 2, 383, 384, 7, 81, 2, 2,
	384, 405, 7, 68, 2, 2, 385, 387, 4, 50, 59, 2, 386, 385, 3, 2, 2, 2, 387,
	388, 3, 2, 2, 2, 388, 386, 3, 2, 2, 2, 388, 389, 3, 2, 2, 2, 389, 405,
	3, 2, 2, 2, 390, 392, 4, 50, 59, 2, 391, 390, 3, 2, 2, 2, 392, 395, 3,
	2, 2, 2, 393, 391, 3, 2, 2, 2, 393, 394, 3, 2, 2, 2, 394, 396, 3, 2, 2,
	2, 395, 393, 3, 2, 2, 2, 396, 400, 7, 48, 2, 2, 397, 399, 4, 50, 59, 2,
	398, 397, 3, 2, 2, 2, 399, 402, 3, 2, 2, 2, 400, 398, 3, 2, 2, 2, 400,
	401, 3, 2, 2, 2, 401, 405, 3, 2, 2, 2, 402, 400, 3, 2, 2, 2, 403, 405,
	5, 87, 44, 2, 404, 374, 3, 2, 2, 2, 404, 377, 3, 2, 2, 2, 404, 381, 3,
	2, 2, 2, 404, 386, 3, 2, 2, 2, 404, 393, 3, 2, 2, 2, 404, 403, 3, 2, 2,
	2, 405, 86, 3, 2, 2, 2, 406, 412, 7, 36, 2, 2, 407, 408, 7, 41, 2, 2, 408,
	411, 7, 36, 2, 2, 409, 411, 10, 3, 2, 2, 410, 407, 3, 2, 2, 2, 410, 409,
	3, 2, 2, 2, 411, 414, 3, 2, 2, 2, 412, 410, 3, 2, 2, 2, 412, 413, 3, 2,
	2, 2, 413, 415, 3, 2, 2, 2, 414, 412, 3, 2, 2, 2, 415, 416, 7, 36, 2, 2,
	416, 88, 3, 2, 2, 2, 417, 418, 9, 4, 2, 2, 418, 419, 3, 2, 2, 2, 419, 420,
	8, 45, 2, 2, 420, 90, 3, 2, 2, 2, 10, 2, 372, 388, 393, 400, 404, 410,
	412, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'HAI'", "'KTHXBYE'", "'IM IN YR'", "'WILE'", "'IM OUTTA YR'", "'I HAS A'",
	"'ITZ'", "'BTW'", "'OBTW'", "'TLDR'", "'VISIBLE'", "'MKAY?'", "'O RLY?'",
	"'YA RLY'", "'OIC'", "'MEBBE'", "'NO WAI'", "'GIMMEH'", "'HOW DUZ I'",
	"'YR'", "'AN YR'", "'IF U SAY SO'", "'R'", "'BOTH SAEM'", "'AN'", "'DIFFRINT'",
	"'BOTH OF'", "'EITHER OF'", "'BIGGR OF'", "'SMALLR OF'", "'SUM OF'", "'DIFF OF'",
	"'PRODUKT OF'", "'QUOSHUNT OF'", "'MOD OF'", "'MAEK'", "'A'", "'ALL OF'",
	"'ANY OF'", "'NOT'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "LABEL", "ATOM", "STRING", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "LABEL",
	"ATOM", "STRING", "WS",
}

type lolcodeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewlolcodeLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *lolcodeLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewlolcodeLexer(input antlr.CharStream) *lolcodeLexer {
	l := new(lolcodeLexer)
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
	l.GrammarFileName = "lolcode.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// lolcodeLexer tokens.
const (
	lolcodeLexerT__0   = 1
	lolcodeLexerT__1   = 2
	lolcodeLexerT__2   = 3
	lolcodeLexerT__3   = 4
	lolcodeLexerT__4   = 5
	lolcodeLexerT__5   = 6
	lolcodeLexerT__6   = 7
	lolcodeLexerT__7   = 8
	lolcodeLexerT__8   = 9
	lolcodeLexerT__9   = 10
	lolcodeLexerT__10  = 11
	lolcodeLexerT__11  = 12
	lolcodeLexerT__12  = 13
	lolcodeLexerT__13  = 14
	lolcodeLexerT__14  = 15
	lolcodeLexerT__15  = 16
	lolcodeLexerT__16  = 17
	lolcodeLexerT__17  = 18
	lolcodeLexerT__18  = 19
	lolcodeLexerT__19  = 20
	lolcodeLexerT__20  = 21
	lolcodeLexerT__21  = 22
	lolcodeLexerT__22  = 23
	lolcodeLexerT__23  = 24
	lolcodeLexerT__24  = 25
	lolcodeLexerT__25  = 26
	lolcodeLexerT__26  = 27
	lolcodeLexerT__27  = 28
	lolcodeLexerT__28  = 29
	lolcodeLexerT__29  = 30
	lolcodeLexerT__30  = 31
	lolcodeLexerT__31  = 32
	lolcodeLexerT__32  = 33
	lolcodeLexerT__33  = 34
	lolcodeLexerT__34  = 35
	lolcodeLexerT__35  = 36
	lolcodeLexerT__36  = 37
	lolcodeLexerT__37  = 38
	lolcodeLexerT__38  = 39
	lolcodeLexerT__39  = 40
	lolcodeLexerLABEL  = 41
	lolcodeLexerATOM   = 42
	lolcodeLexerSTRING = 43
	lolcodeLexerWS     = 44
)
