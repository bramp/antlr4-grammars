// Code generated from snobol.g4 by ANTLR 4.9.3. DO NOT EDIT.

package snobol

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 57, 475,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6,
	3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3,
	37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42,
	3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3,
	47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 7, 49, 351, 10, 49, 12, 49,
	14, 49, 354, 11, 49, 3, 49, 3, 49, 3, 50, 3, 50, 7, 50, 360, 10, 50, 12,
	50, 14, 50, 363, 11, 50, 3, 50, 3, 50, 3, 51, 3, 51, 7, 51, 369, 10, 51,
	12, 51, 14, 51, 372, 11, 51, 3, 52, 5, 52, 375, 10, 52, 3, 52, 6, 52, 378,
	10, 52, 13, 52, 14, 52, 379, 3, 53, 5, 53, 383, 10, 53, 3, 53, 6, 53, 386,
	10, 53, 13, 53, 14, 53, 387, 3, 53, 3, 53, 6, 53, 392, 10, 53, 13, 53,
	14, 53, 393, 5, 53, 396, 10, 53, 3, 53, 3, 53, 7, 53, 400, 10, 53, 12,
	53, 14, 53, 403, 11, 53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57,
	3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 60, 3, 60, 3, 61, 3, 61, 3, 62, 3,
	62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67,
	3, 68, 3, 68, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3,
	73, 3, 73, 3, 74, 3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 77, 3, 77, 3, 78,
	3, 78, 3, 79, 3, 79, 3, 80, 3, 80, 7, 80, 459, 10, 80, 12, 80, 14, 80,
	462, 11, 80, 3, 81, 6, 81, 465, 10, 81, 13, 81, 14, 81, 466, 3, 82, 6,
	82, 470, 10, 82, 13, 82, 14, 82, 471, 3, 82, 3, 82, 2, 2, 83, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43,
	23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61,
	32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79,
	41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97,
	50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 2, 109, 2, 111, 2, 113, 2,
	115, 2, 117, 2, 119, 2, 121, 2, 123, 2, 125, 2, 127, 2, 129, 2, 131, 2,
	133, 2, 135, 2, 137, 2, 139, 2, 141, 2, 143, 2, 145, 2, 147, 2, 149, 2,
	151, 2, 153, 2, 155, 2, 157, 2, 159, 55, 161, 56, 163, 57, 3, 2, 35, 5,
	2, 12, 12, 15, 15, 36, 36, 5, 2, 12, 12, 15, 15, 41, 41, 4, 2, 67, 92,
	99, 124, 5, 2, 50, 59, 67, 92, 99, 124, 4, 2, 45, 45, 47, 47, 4, 2, 71,
	71, 103, 103, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69,
	101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73,
	105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76,
	108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79,
	111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82,
	114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85,
	117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88,
	120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91,
	123, 123, 4, 2, 92, 92, 124, 124, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34,
	34, 2, 461, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
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
	2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101,
	3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 159, 3, 2, 2, 2,
	2, 161, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2, 3, 165, 3, 2, 2, 2, 5, 167, 3,
	2, 2, 2, 7, 169, 3, 2, 2, 2, 9, 175, 3, 2, 2, 2, 11, 182, 3, 2, 2, 2, 13,
	185, 3, 2, 2, 2, 15, 188, 3, 2, 2, 2, 17, 191, 3, 2, 2, 2, 19, 194, 3,
	2, 2, 2, 21, 197, 3, 2, 2, 2, 23, 200, 3, 2, 2, 2, 25, 208, 3, 2, 2, 2,
	27, 212, 3, 2, 2, 2, 29, 217, 3, 2, 2, 2, 31, 222, 3, 2, 2, 2, 33, 226,
	3, 2, 2, 2, 35, 230, 3, 2, 2, 2, 37, 233, 3, 2, 2, 2, 39, 239, 3, 2, 2,
	2, 41, 243, 3, 2, 2, 2, 43, 247, 3, 2, 2, 2, 45, 252, 3, 2, 2, 2, 47, 260,
	3, 2, 2, 2, 49, 265, 3, 2, 2, 2, 51, 273, 3, 2, 2, 2, 53, 278, 3, 2, 2,
	2, 55, 283, 3, 2, 2, 2, 57, 289, 3, 2, 2, 2, 59, 297, 3, 2, 2, 2, 61, 303,
	3, 2, 2, 2, 63, 308, 3, 2, 2, 2, 65, 314, 3, 2, 2, 2, 67, 316, 3, 2, 2,
	2, 69, 318, 3, 2, 2, 2, 71, 320, 3, 2, 2, 2, 73, 322, 3, 2, 2, 2, 75, 324,
	3, 2, 2, 2, 77, 326, 3, 2, 2, 2, 79, 328, 3, 2, 2, 2, 81, 330, 3, 2, 2,
	2, 83, 332, 3, 2, 2, 2, 85, 334, 3, 2, 2, 2, 87, 336, 3, 2, 2, 2, 89, 338,
	3, 2, 2, 2, 91, 340, 3, 2, 2, 2, 93, 342, 3, 2, 2, 2, 95, 344, 3, 2, 2,
	2, 97, 348, 3, 2, 2, 2, 99, 357, 3, 2, 2, 2, 101, 366, 3, 2, 2, 2, 103,
	374, 3, 2, 2, 2, 105, 382, 3, 2, 2, 2, 107, 404, 3, 2, 2, 2, 109, 406,
	3, 2, 2, 2, 111, 408, 3, 2, 2, 2, 113, 410, 3, 2, 2, 2, 115, 412, 3, 2,
	2, 2, 117, 414, 3, 2, 2, 2, 119, 416, 3, 2, 2, 2, 121, 418, 3, 2, 2, 2,
	123, 420, 3, 2, 2, 2, 125, 422, 3, 2, 2, 2, 127, 424, 3, 2, 2, 2, 129,
	426, 3, 2, 2, 2, 131, 428, 3, 2, 2, 2, 133, 430, 3, 2, 2, 2, 135, 432,
	3, 2, 2, 2, 137, 434, 3, 2, 2, 2, 139, 436, 3, 2, 2, 2, 141, 438, 3, 2,
	2, 2, 143, 440, 3, 2, 2, 2, 145, 442, 3, 2, 2, 2, 147, 444, 3, 2, 2, 2,
	149, 446, 3, 2, 2, 2, 151, 448, 3, 2, 2, 2, 153, 450, 3, 2, 2, 2, 155,
	452, 3, 2, 2, 2, 157, 454, 3, 2, 2, 2, 159, 456, 3, 2, 2, 2, 161, 464,
	3, 2, 2, 2, 163, 469, 3, 2, 2, 2, 165, 166, 7, 93, 2, 2, 166, 4, 3, 2,
	2, 2, 167, 168, 7, 95, 2, 2, 168, 6, 3, 2, 2, 2, 169, 170, 7, 107, 2, 2,
	170, 171, 7, 102, 2, 2, 171, 172, 7, 103, 2, 2, 172, 173, 7, 112, 2, 2,
	173, 174, 7, 118, 2, 2, 174, 8, 3, 2, 2, 2, 175, 176, 7, 102, 2, 2, 176,
	177, 7, 107, 2, 2, 177, 178, 7, 104, 2, 2, 178, 179, 7, 104, 2, 2, 179,
	180, 7, 103, 2, 2, 180, 181, 7, 116, 2, 2, 181, 10, 3, 2, 2, 2, 182, 183,
	7, 103, 2, 2, 183, 184, 7, 115, 2, 2, 184, 12, 3, 2, 2, 2, 185, 186, 7,
	112, 2, 2, 186, 187, 7, 103, 2, 2, 187, 14, 3, 2, 2, 2, 188, 189, 7, 105,
	2, 2, 189, 190, 7, 103, 2, 2, 190, 16, 3, 2, 2, 2, 191, 192, 7, 105, 2,
	2, 192, 193, 7, 118, 2, 2, 193, 18, 3, 2, 2, 2, 194, 195, 7, 110, 2, 2,
	195, 196, 7, 103, 2, 2, 196, 20, 3, 2, 2, 2, 197, 198, 7, 110, 2, 2, 198,
	199, 7, 118, 2, 2, 199, 22, 3, 2, 2, 2, 200, 201, 7, 107, 2, 2, 201, 202,
	7, 112, 2, 2, 202, 203, 7, 118, 2, 2, 203, 204, 7, 103, 2, 2, 204, 205,
	7, 105, 2, 2, 205, 206, 7, 103, 2, 2, 206, 207, 7, 116, 2, 2, 207, 24,
	3, 2, 2, 2, 208, 209, 7, 110, 2, 2, 209, 210, 7, 105, 2, 2, 210, 211, 7,
	118, 2, 2, 211, 26, 3, 2, 2, 2, 212, 213, 7, 99, 2, 2, 213, 214, 7, 118,
	2, 2, 214, 215, 7, 99, 2, 2, 215, 216, 7, 112, 2, 2, 216, 28, 3, 2, 2,
	2, 217, 218, 7, 101, 2, 2, 218, 219, 7, 106, 2, 2, 219, 220, 7, 113, 2,
	2, 220, 221, 7, 114, 2, 2, 221, 30, 3, 2, 2, 2, 222, 223, 7, 101, 2, 2,
	223, 224, 7, 113, 2, 2, 224, 225, 7, 117, 2, 2, 225, 32, 3, 2, 2, 2, 226,
	227, 7, 103, 2, 2, 227, 228, 7, 122, 2, 2, 228, 229, 7, 114, 2, 2, 229,
	34, 3, 2, 2, 2, 230, 231, 7, 110, 2, 2, 231, 232, 7, 112, 2, 2, 232, 36,
	3, 2, 2, 2, 233, 234, 7, 116, 2, 2, 234, 235, 7, 103, 2, 2, 235, 236, 7,
	111, 2, 2, 236, 237, 7, 102, 2, 2, 237, 238, 7, 116, 2, 2, 238, 38, 3,
	2, 2, 2, 239, 240, 7, 117, 2, 2, 240, 241, 7, 107, 2, 2, 241, 242, 7, 112,
	2, 2, 242, 40, 3, 2, 2, 2, 243, 244, 7, 118, 2, 2, 244, 245, 7, 99, 2,
	2, 245, 246, 7, 112, 2, 2, 246, 42, 3, 2, 2, 2, 247, 248, 7, 102, 2, 2,
	248, 249, 7, 119, 2, 2, 249, 250, 7, 114, 2, 2, 250, 251, 7, 110, 2, 2,
	251, 44, 3, 2, 2, 2, 252, 253, 7, 116, 2, 2, 253, 254, 7, 103, 2, 2, 254,
	255, 7, 120, 2, 2, 255, 256, 7, 103, 2, 2, 256, 257, 7, 116, 2, 2, 257,
	258, 7, 117, 2, 2, 258, 259, 7, 103, 2, 2, 259, 46, 3, 2, 2, 2, 260, 261,
	7, 102, 2, 2, 261, 262, 7, 99, 2, 2, 262, 263, 7, 118, 2, 2, 263, 264,
	7, 103, 2, 2, 264, 48, 3, 2, 2, 2, 265, 266, 7, 116, 2, 2, 266, 267, 7,
	103, 2, 2, 267, 268, 7, 114, 2, 2, 268, 269, 7, 110, 2, 2, 269, 270, 7,
	99, 2, 2, 270, 271, 7, 101, 2, 2, 271, 272, 7, 103, 2, 2, 272, 50, 3, 2,
	2, 2, 273, 274, 7, 117, 2, 2, 274, 275, 7, 107, 2, 2, 275, 276, 7, 124,
	2, 2, 276, 277, 7, 103, 2, 2, 277, 52, 3, 2, 2, 2, 278, 279, 7, 118, 2,
	2, 279, 280, 7, 116, 2, 2, 280, 281, 7, 107, 2, 2, 281, 282, 7, 111, 2,
	2, 282, 54, 3, 2, 2, 2, 283, 284, 7, 99, 2, 2, 284, 285, 7, 116, 2, 2,
	285, 286, 7, 116, 2, 2, 286, 287, 7, 99, 2, 2, 287, 288, 7, 123, 2, 2,
	288, 56, 3, 2, 2, 2, 289, 290, 7, 101, 2, 2, 290, 291, 7, 113, 2, 2, 291,
	292, 7, 112, 2, 2, 292, 293, 7, 120, 2, 2, 293, 294, 7, 103, 2, 2, 294,
	295, 7, 116, 2, 2, 295, 296, 7, 118, 2, 2, 296, 58, 3, 2, 2, 2, 297, 298,
	7, 118, 2, 2, 298, 299, 7, 99, 2, 2, 299, 300, 7, 100, 2, 2, 300, 301,
	7, 110, 2, 2, 301, 302, 7, 103, 2, 2, 302, 60, 3, 2, 2, 2, 303, 304, 7,
	117, 2, 2, 304, 305, 7, 113, 2, 2, 305, 306, 7, 116, 2, 2, 306, 307, 7,
	118, 2, 2, 307, 62, 3, 2, 2, 2, 308, 309, 7, 100, 2, 2, 309, 310, 7, 116,
	2, 2, 310, 311, 7, 103, 2, 2, 311, 312, 7, 99, 2, 2, 312, 313, 7, 109,
	2, 2, 313, 64, 3, 2, 2, 2, 314, 315, 7, 104, 2, 2, 315, 66, 3, 2, 2, 2,
	316, 317, 7, 72, 2, 2, 317, 68, 3, 2, 2, 2, 318, 319, 7, 117, 2, 2, 319,
	70, 3, 2, 2, 2, 320, 321, 7, 85, 2, 2, 321, 72, 3, 2, 2, 2, 322, 323, 7,
	46, 2, 2, 323, 74, 3, 2, 2, 2, 324, 325, 7, 42, 2, 2, 325, 76, 3, 2, 2,
	2, 326, 327, 7, 43, 2, 2, 327, 78, 3, 2, 2, 2, 328, 329, 7, 40, 2, 2, 329,
	80, 3, 2, 2, 2, 330, 331, 7, 45, 2, 2, 331, 82, 3, 2, 2, 2, 332, 333, 7,
	47, 2, 2, 333, 84, 3, 2, 2, 2, 334, 335, 7, 44, 2, 2, 335, 86, 3, 2, 2,
	2, 336, 337, 7, 49, 2, 2, 337, 88, 3, 2, 2, 2, 338, 339, 7, 96, 2, 2, 339,
	90, 3, 2, 2, 2, 340, 341, 7, 63, 2, 2, 341, 92, 3, 2, 2, 2, 342, 343, 7,
	60, 2, 2, 343, 94, 3, 2, 2, 2, 344, 345, 7, 71, 2, 2, 345, 346, 7, 80,
	2, 2, 346, 347, 7, 70, 2, 2, 347, 96, 3, 2, 2, 2, 348, 352, 7, 36, 2, 2,
	349, 351, 10, 2, 2, 2, 350, 349, 3, 2, 2, 2, 351, 354, 3, 2, 2, 2, 352,
	350, 3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 355, 3, 2, 2, 2, 354, 352,
	3, 2, 2, 2, 355, 356, 7, 36, 2, 2, 356, 98, 3, 2, 2, 2, 357, 361, 7, 41,
	2, 2, 358, 360, 10, 3, 2, 2, 359, 358, 3, 2, 2, 2, 360, 363, 3, 2, 2, 2,
	361, 359, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 364, 3, 2, 2, 2, 363,
	361, 3, 2, 2, 2, 364, 365, 7, 41, 2, 2, 365, 100, 3, 2, 2, 2, 366, 370,
	9, 4, 2, 2, 367, 369, 9, 5, 2, 2, 368, 367, 3, 2, 2, 2, 369, 372, 3, 2,
	2, 2, 370, 368, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2, 371, 102, 3, 2, 2, 2,
	372, 370, 3, 2, 2, 2, 373, 375, 9, 6, 2, 2, 374, 373, 3, 2, 2, 2, 374,
	375, 3, 2, 2, 2, 375, 377, 3, 2, 2, 2, 376, 378, 4, 50, 59, 2, 377, 376,
	3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 377, 3, 2, 2, 2, 379, 380, 3, 2,
	2, 2, 380, 104, 3, 2, 2, 2, 381, 383, 9, 6, 2, 2, 382, 381, 3, 2, 2, 2,
	382, 383, 3, 2, 2, 2, 383, 385, 3, 2, 2, 2, 384, 386, 4, 50, 59, 2, 385,
	384, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 385, 3, 2, 2, 2, 387, 388,
	3, 2, 2, 2, 388, 395, 3, 2, 2, 2, 389, 391, 7, 48, 2, 2, 390, 392, 4, 50,
	59, 2, 391, 390, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2, 393, 391, 3, 2, 2, 2,
	393, 394, 3, 2, 2, 2, 394, 396, 3, 2, 2, 2, 395, 389, 3, 2, 2, 2, 395,
	396, 3, 2, 2, 2, 396, 401, 3, 2, 2, 2, 397, 398, 9, 7, 2, 2, 398, 400,
	5, 105, 53, 2, 399, 397, 3, 2, 2, 2, 400, 403, 3, 2, 2, 2, 401, 399, 3,
	2, 2, 2, 401, 402, 3, 2, 2, 2, 402, 106, 3, 2, 2, 2, 403, 401, 3, 2, 2,
	2, 404, 405, 9, 8, 2, 2, 405, 108, 3, 2, 2, 2, 406, 407, 9, 9, 2, 2, 407,
	110, 3, 2, 2, 2, 408, 409, 9, 10, 2, 2, 409, 112, 3, 2, 2, 2, 410, 411,
	9, 11, 2, 2, 411, 114, 3, 2, 2, 2, 412, 413, 9, 7, 2, 2, 413, 116, 3, 2,
	2, 2, 414, 415, 9, 12, 2, 2, 415, 118, 3, 2, 2, 2, 416, 417, 9, 13, 2,
	2, 417, 120, 3, 2, 2, 2, 418, 419, 9, 14, 2, 2, 419, 122, 3, 2, 2, 2, 420,
	421, 9, 15, 2, 2, 421, 124, 3, 2, 2, 2, 422, 423, 9, 16, 2, 2, 423, 126,
	3, 2, 2, 2, 424, 425, 9, 17, 2, 2, 425, 128, 3, 2, 2, 2, 426, 427, 9, 18,
	2, 2, 427, 130, 3, 2, 2, 2, 428, 429, 9, 19, 2, 2, 429, 132, 3, 2, 2, 2,
	430, 431, 9, 20, 2, 2, 431, 134, 3, 2, 2, 2, 432, 433, 9, 21, 2, 2, 433,
	136, 3, 2, 2, 2, 434, 435, 9, 22, 2, 2, 435, 138, 3, 2, 2, 2, 436, 437,
	9, 23, 2, 2, 437, 140, 3, 2, 2, 2, 438, 439, 9, 24, 2, 2, 439, 142, 3,
	2, 2, 2, 440, 441, 9, 25, 2, 2, 441, 144, 3, 2, 2, 2, 442, 443, 9, 26,
	2, 2, 443, 146, 3, 2, 2, 2, 444, 445, 9, 27, 2, 2, 445, 148, 3, 2, 2, 2,
	446, 447, 9, 28, 2, 2, 447, 150, 3, 2, 2, 2, 448, 449, 9, 29, 2, 2, 449,
	152, 3, 2, 2, 2, 450, 451, 9, 30, 2, 2, 451, 154, 3, 2, 2, 2, 452, 453,
	9, 31, 2, 2, 453, 156, 3, 2, 2, 2, 454, 455, 9, 32, 2, 2, 455, 158, 3,
	2, 2, 2, 456, 460, 7, 44, 2, 2, 457, 459, 10, 33, 2, 2, 458, 457, 3, 2,
	2, 2, 459, 462, 3, 2, 2, 2, 460, 458, 3, 2, 2, 2, 460, 461, 3, 2, 2, 2,
	461, 160, 3, 2, 2, 2, 462, 460, 3, 2, 2, 2, 463, 465, 9, 33, 2, 2, 464,
	463, 3, 2, 2, 2, 465, 466, 3, 2, 2, 2, 466, 464, 3, 2, 2, 2, 466, 467,
	3, 2, 2, 2, 467, 162, 3, 2, 2, 2, 468, 470, 9, 34, 2, 2, 469, 468, 3, 2,
	2, 2, 470, 471, 3, 2, 2, 2, 471, 469, 3, 2, 2, 2, 471, 472, 3, 2, 2, 2,
	472, 473, 3, 2, 2, 2, 473, 474, 8, 82, 2, 2, 474, 164, 3, 2, 2, 2, 16,
	2, 352, 361, 370, 374, 379, 382, 387, 393, 395, 401, 460, 466, 471, 3,
	8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'['", "']'", "'ident'", "'differ'", "'eq'", "'ne'", "'ge'", "'gt'",
	"'le'", "'lt'", "'integer'", "'lgt'", "'atan'", "'chop'", "'cos'", "'exp'",
	"'ln'", "'remdr'", "'sin'", "'tan'", "'dupl'", "'reverse'", "'date'", "'replace'",
	"'size'", "'trim'", "'array'", "'convert'", "'table'", "'sort'", "'break'",
	"'f'", "'F'", "'s'", "'S'", "','", "'('", "')'", "'&'", "'+'", "'-'", "'*'",
	"'/'", "'^'", "'='", "':'", "'END'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"COMMA", "LPAREN", "RPAREN", "AMP", "PLUS", "MINUS", "TIMES", "DIV", "POW",
	"EQ", "COLON", "END", "STRINGLITERAL1", "STRINGLITERAL2", "STRING", "INTEGER",
	"REAL", "COMMENT", "EOL", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "COMMA", "LPAREN", "RPAREN", "AMP", "PLUS", "MINUS",
	"TIMES", "DIV", "POW", "EQ", "COLON", "END", "STRINGLITERAL1", "STRINGLITERAL2",
	"STRING", "INTEGER", "REAL", "A", "B", "C", "D", "E", "F", "G", "H", "I",
	"J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X",
	"Y", "Z", "COMMENT", "EOL", "WS",
}

type snobolLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewsnobolLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *snobolLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewsnobolLexer(input antlr.CharStream) *snobolLexer {
	l := new(snobolLexer)
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
	l.GrammarFileName = "snobol.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// snobolLexer tokens.
const (
	snobolLexerT__0           = 1
	snobolLexerT__1           = 2
	snobolLexerT__2           = 3
	snobolLexerT__3           = 4
	snobolLexerT__4           = 5
	snobolLexerT__5           = 6
	snobolLexerT__6           = 7
	snobolLexerT__7           = 8
	snobolLexerT__8           = 9
	snobolLexerT__9           = 10
	snobolLexerT__10          = 11
	snobolLexerT__11          = 12
	snobolLexerT__12          = 13
	snobolLexerT__13          = 14
	snobolLexerT__14          = 15
	snobolLexerT__15          = 16
	snobolLexerT__16          = 17
	snobolLexerT__17          = 18
	snobolLexerT__18          = 19
	snobolLexerT__19          = 20
	snobolLexerT__20          = 21
	snobolLexerT__21          = 22
	snobolLexerT__22          = 23
	snobolLexerT__23          = 24
	snobolLexerT__24          = 25
	snobolLexerT__25          = 26
	snobolLexerT__26          = 27
	snobolLexerT__27          = 28
	snobolLexerT__28          = 29
	snobolLexerT__29          = 30
	snobolLexerT__30          = 31
	snobolLexerT__31          = 32
	snobolLexerT__32          = 33
	snobolLexerT__33          = 34
	snobolLexerT__34          = 35
	snobolLexerCOMMA          = 36
	snobolLexerLPAREN         = 37
	snobolLexerRPAREN         = 38
	snobolLexerAMP            = 39
	snobolLexerPLUS           = 40
	snobolLexerMINUS          = 41
	snobolLexerTIMES          = 42
	snobolLexerDIV            = 43
	snobolLexerPOW            = 44
	snobolLexerEQ             = 45
	snobolLexerCOLON          = 46
	snobolLexerEND            = 47
	snobolLexerSTRINGLITERAL1 = 48
	snobolLexerSTRINGLITERAL2 = 49
	snobolLexerSTRING         = 50
	snobolLexerINTEGER        = 51
	snobolLexerREAL           = 52
	snobolLexerCOMMENT        = 53
	snobolLexerEOL            = 54
	snobolLexerWS             = 55
)
