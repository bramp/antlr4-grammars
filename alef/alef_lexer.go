// Code generated from alef.g4 by ANTLR 4.9.3. DO NOT EDIT.

package alef

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 101, 641,
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
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4, 96, 9, 96,
	4, 97, 9, 97, 4, 98, 9, 98, 4, 99, 9, 99, 4, 100, 9, 100, 4, 101, 9, 101,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30,
	3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34, 3,
	34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37,
	3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3,
	40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 44,
	3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3,
	48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 51, 3, 51,
	3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53, 3,
	53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55, 3, 55,
	3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3,
	57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 59,
	3, 59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3,
	61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62,
	3, 62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 64, 3,
	64, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66,
	3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3,
	68, 3, 68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70, 3, 71,
	3, 71, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3,
	73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 74, 3, 74, 3, 74, 3, 74, 3, 75, 3, 75,
	3, 75, 3, 75, 3, 76, 3, 76, 3, 76, 3, 76, 3, 76, 3, 77, 3, 77, 3, 77, 3,
	77, 3, 77, 3, 77, 3, 78, 3, 78, 3, 78, 3, 78, 3, 78, 3, 78, 3, 78, 3, 79,
	3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 3, 80, 3, 80, 3, 80, 3, 80, 3,
	80, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 82, 3, 82, 3, 82,
	3, 82, 3, 82, 3, 82, 3, 82, 3, 83, 3, 83, 3, 83, 3, 83, 3, 83, 3, 84, 3,
	84, 3, 84, 3, 84, 3, 84, 3, 84, 3, 85, 3, 85, 3, 85, 3, 85, 3, 85, 3, 85,
	3, 85, 3, 85, 3, 86, 3, 86, 3, 86, 3, 86, 3, 86, 3, 86, 3, 86, 3, 87, 3,
	87, 3, 87, 3, 87, 3, 87, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 89,
	3, 89, 3, 89, 3, 89, 3, 89, 3, 89, 3, 89, 3, 89, 3, 90, 3, 90, 3, 90, 3,
	90, 3, 90, 3, 90, 3, 91, 3, 91, 3, 91, 3, 91, 3, 91, 3, 91, 3, 92, 3, 92,
	3, 92, 3, 92, 3, 92, 3, 93, 3, 93, 3, 93, 3, 93, 3, 93, 3, 93, 3, 94, 3,
	94, 3, 94, 3, 94, 3, 94, 3, 94, 3, 95, 3, 95, 3, 95, 3, 95, 3, 95, 3, 95,
	3, 95, 3, 95, 3, 96, 3, 96, 7, 96, 588, 10, 96, 12, 96, 14, 96, 591, 11,
	96, 3, 97, 3, 97, 7, 97, 595, 10, 97, 12, 97, 14, 97, 598, 11, 97, 3, 97,
	3, 97, 3, 98, 3, 98, 7, 98, 604, 10, 98, 12, 98, 14, 98, 607, 11, 98, 3,
	98, 3, 98, 3, 99, 6, 99, 612, 10, 99, 13, 99, 14, 99, 613, 3, 99, 3, 99,
	7, 99, 618, 10, 99, 12, 99, 14, 99, 621, 11, 99, 5, 99, 623, 10, 99, 3,
	99, 3, 99, 6, 99, 627, 10, 99, 13, 99, 14, 99, 628, 5, 99, 631, 10, 99,
	3, 100, 3, 100, 3, 101, 6, 101, 636, 10, 101, 13, 101, 14, 101, 637, 3,
	101, 3, 101, 2, 2, 102, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35,
	19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53,
	28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71,
	37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89,
	46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54,
	107, 55, 109, 56, 111, 57, 113, 58, 115, 59, 117, 60, 119, 61, 121, 62,
	123, 63, 125, 64, 127, 65, 129, 66, 131, 67, 133, 68, 135, 69, 137, 70,
	139, 71, 141, 72, 143, 73, 145, 74, 147, 75, 149, 76, 151, 77, 153, 78,
	155, 79, 157, 80, 159, 81, 161, 82, 163, 83, 165, 84, 167, 85, 169, 86,
	171, 87, 173, 88, 175, 89, 177, 90, 179, 91, 181, 92, 183, 93, 185, 94,
	187, 95, 189, 96, 191, 97, 193, 98, 195, 99, 197, 100, 199, 2, 201, 101,
	3, 2, 8, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99,
	124, 3, 2, 36, 36, 3, 2, 41, 41, 3, 2, 50, 59, 5, 2, 11, 12, 15, 15, 34,
	34, 2, 648, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9,
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
	3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2,
	2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3,
	2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2,
	123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2,
	2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137,
	3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2,
	2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3,
	2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 157, 3, 2, 2, 2, 2,
	159, 3, 2, 2, 2, 2, 161, 3, 2, 2, 2, 2, 163, 3, 2, 2, 2, 2, 165, 3, 2,
	2, 2, 2, 167, 3, 2, 2, 2, 2, 169, 3, 2, 2, 2, 2, 171, 3, 2, 2, 2, 2, 173,
	3, 2, 2, 2, 2, 175, 3, 2, 2, 2, 2, 177, 3, 2, 2, 2, 2, 179, 3, 2, 2, 2,
	2, 181, 3, 2, 2, 2, 2, 183, 3, 2, 2, 2, 2, 185, 3, 2, 2, 2, 2, 187, 3,
	2, 2, 2, 2, 189, 3, 2, 2, 2, 2, 191, 3, 2, 2, 2, 2, 193, 3, 2, 2, 2, 2,
	195, 3, 2, 2, 2, 2, 197, 3, 2, 2, 2, 2, 201, 3, 2, 2, 2, 3, 203, 3, 2,
	2, 2, 5, 205, 3, 2, 2, 2, 7, 207, 3, 2, 2, 2, 9, 209, 3, 2, 2, 2, 11, 211,
	3, 2, 2, 2, 13, 213, 3, 2, 2, 2, 15, 215, 3, 2, 2, 2, 17, 217, 3, 2, 2,
	2, 19, 219, 3, 2, 2, 2, 21, 221, 3, 2, 2, 2, 23, 223, 3, 2, 2, 2, 25, 225,
	3, 2, 2, 2, 27, 227, 3, 2, 2, 2, 29, 229, 3, 2, 2, 2, 31, 231, 3, 2, 2,
	2, 33, 233, 3, 2, 2, 2, 35, 235, 3, 2, 2, 2, 37, 237, 3, 2, 2, 2, 39, 240,
	3, 2, 2, 2, 41, 243, 3, 2, 2, 2, 43, 245, 3, 2, 2, 2, 45, 247, 3, 2, 2,
	2, 47, 250, 3, 2, 2, 2, 49, 253, 3, 2, 2, 2, 51, 256, 3, 2, 2, 2, 53, 259,
	3, 2, 2, 2, 55, 261, 3, 2, 2, 2, 57, 263, 3, 2, 2, 2, 59, 265, 3, 2, 2,
	2, 61, 268, 3, 2, 2, 2, 63, 271, 3, 2, 2, 2, 65, 274, 3, 2, 2, 2, 67, 277,
	3, 2, 2, 2, 69, 280, 3, 2, 2, 2, 71, 283, 3, 2, 2, 2, 73, 286, 3, 2, 2,
	2, 75, 289, 3, 2, 2, 2, 77, 292, 3, 2, 2, 2, 79, 296, 3, 2, 2, 2, 81, 300,
	3, 2, 2, 2, 83, 303, 3, 2, 2, 2, 85, 306, 3, 2, 2, 2, 87, 309, 3, 2, 2,
	2, 89, 312, 3, 2, 2, 2, 91, 315, 3, 2, 2, 2, 93, 318, 3, 2, 2, 2, 95, 320,
	3, 2, 2, 2, 97, 322, 3, 2, 2, 2, 99, 325, 3, 2, 2, 2, 101, 327, 3, 2, 2,
	2, 103, 331, 3, 2, 2, 2, 105, 336, 3, 2, 2, 2, 107, 342, 3, 2, 2, 2, 109,
	346, 3, 2, 2, 2, 111, 353, 3, 2, 2, 2, 113, 359, 3, 2, 2, 2, 115, 364,
	3, 2, 2, 2, 117, 369, 3, 2, 2, 2, 119, 374, 3, 2, 2, 2, 121, 380, 3, 2,
	2, 2, 123, 389, 3, 2, 2, 2, 125, 397, 3, 2, 2, 2, 127, 400, 3, 2, 2, 2,
	129, 405, 3, 2, 2, 2, 131, 410, 3, 2, 2, 2, 133, 417, 3, 2, 2, 2, 135,
	423, 3, 2, 2, 2, 137, 427, 3, 2, 2, 2, 139, 432, 3, 2, 2, 2, 141, 435,
	3, 2, 2, 2, 143, 439, 3, 2, 2, 2, 145, 446, 3, 2, 2, 2, 147, 451, 3, 2,
	2, 2, 149, 455, 3, 2, 2, 2, 151, 459, 3, 2, 2, 2, 153, 464, 3, 2, 2, 2,
	155, 470, 3, 2, 2, 2, 157, 477, 3, 2, 2, 2, 159, 484, 3, 2, 2, 2, 161,
	489, 3, 2, 2, 2, 163, 496, 3, 2, 2, 2, 165, 503, 3, 2, 2, 2, 167, 508,
	3, 2, 2, 2, 169, 514, 3, 2, 2, 2, 171, 522, 3, 2, 2, 2, 173, 529, 3, 2,
	2, 2, 175, 534, 3, 2, 2, 2, 177, 540, 3, 2, 2, 2, 179, 548, 3, 2, 2, 2,
	181, 554, 3, 2, 2, 2, 183, 560, 3, 2, 2, 2, 185, 565, 3, 2, 2, 2, 187,
	571, 3, 2, 2, 2, 189, 577, 3, 2, 2, 2, 191, 585, 3, 2, 2, 2, 193, 592,
	3, 2, 2, 2, 195, 601, 3, 2, 2, 2, 197, 611, 3, 2, 2, 2, 199, 632, 3, 2,
	2, 2, 201, 635, 3, 2, 2, 2, 203, 204, 7, 61, 2, 2, 204, 4, 3, 2, 2, 2,
	205, 206, 7, 42, 2, 2, 206, 6, 3, 2, 2, 2, 207, 208, 7, 43, 2, 2, 208,
	8, 3, 2, 2, 2, 209, 210, 7, 48, 2, 2, 210, 10, 3, 2, 2, 2, 211, 212, 7,
	125, 2, 2, 212, 12, 3, 2, 2, 2, 213, 214, 7, 127, 2, 2, 214, 14, 3, 2,
	2, 2, 215, 216, 7, 93, 2, 2, 216, 16, 3, 2, 2, 2, 217, 218, 7, 95, 2, 2,
	218, 18, 3, 2, 2, 2, 219, 220, 7, 46, 2, 2, 220, 20, 3, 2, 2, 2, 221, 222,
	7, 63, 2, 2, 222, 22, 3, 2, 2, 2, 223, 224, 7, 44, 2, 2, 224, 24, 3, 2,
	2, 2, 225, 226, 7, 35, 2, 2, 226, 26, 3, 2, 2, 2, 227, 228, 7, 60, 2, 2,
	228, 28, 3, 2, 2, 2, 229, 230, 7, 49, 2, 2, 230, 30, 3, 2, 2, 2, 231, 232,
	7, 39, 2, 2, 232, 32, 3, 2, 2, 2, 233, 234, 7, 45, 2, 2, 234, 34, 3, 2,
	2, 2, 235, 236, 7, 47, 2, 2, 236, 36, 3, 2, 2, 2, 237, 238, 7, 64, 2, 2,
	238, 239, 7, 64, 2, 2, 239, 38, 3, 2, 2, 2, 240, 241, 7, 62, 2, 2, 241,
	242, 7, 62, 2, 2, 242, 40, 3, 2, 2, 2, 243, 244, 7, 62, 2, 2, 244, 42,
	3, 2, 2, 2, 245, 246, 7, 64, 2, 2, 246, 44, 3, 2, 2, 2, 247, 248, 7, 62,
	2, 2, 248, 249, 7, 63, 2, 2, 249, 46, 3, 2, 2, 2, 250, 251, 7, 64, 2, 2,
	251, 252, 7, 63, 2, 2, 252, 48, 3, 2, 2, 2, 253, 254, 7, 63, 2, 2, 254,
	255, 7, 63, 2, 2, 255, 50, 3, 2, 2, 2, 256, 257, 7, 35, 2, 2, 257, 258,
	7, 63, 2, 2, 258, 52, 3, 2, 2, 2, 259, 260, 7, 40, 2, 2, 260, 54, 3, 2,
	2, 2, 261, 262, 7, 96, 2, 2, 262, 56, 3, 2, 2, 2, 263, 264, 7, 126, 2,
	2, 264, 58, 3, 2, 2, 2, 265, 266, 7, 40, 2, 2, 266, 267, 7, 40, 2, 2, 267,
	60, 3, 2, 2, 2, 268, 269, 7, 126, 2, 2, 269, 270, 7, 126, 2, 2, 270, 62,
	3, 2, 2, 2, 271, 272, 7, 60, 2, 2, 272, 273, 7, 63, 2, 2, 273, 64, 3, 2,
	2, 2, 274, 275, 7, 62, 2, 2, 275, 276, 7, 47, 2, 2, 276, 66, 3, 2, 2, 2,
	277, 278, 7, 45, 2, 2, 278, 279, 7, 63, 2, 2, 279, 68, 3, 2, 2, 2, 280,
	281, 7, 47, 2, 2, 281, 282, 7, 63, 2, 2, 282, 70, 3, 2, 2, 2, 283, 284,
	7, 44, 2, 2, 284, 285, 7, 63, 2, 2, 285, 72, 3, 2, 2, 2, 286, 287, 7, 49,
	2, 2, 287, 288, 7, 63, 2, 2, 288, 74, 3, 2, 2, 2, 289, 290, 7, 39, 2, 2,
	290, 291, 7, 63, 2, 2, 291, 76, 3, 2, 2, 2, 292, 293, 7, 64, 2, 2, 293,
	294, 7, 64, 2, 2, 294, 295, 7, 63, 2, 2, 295, 78, 3, 2, 2, 2, 296, 297,
	7, 62, 2, 2, 297, 298, 7, 62, 2, 2, 298, 299, 7, 63, 2, 2, 299, 80, 3,
	2, 2, 2, 300, 301, 7, 40, 2, 2, 301, 302, 7, 63, 2, 2, 302, 82, 3, 2, 2,
	2, 303, 304, 7, 126, 2, 2, 304, 305, 7, 63, 2, 2, 305, 84, 3, 2, 2, 2,
	306, 307, 7, 96, 2, 2, 307, 308, 7, 63, 2, 2, 308, 86, 3, 2, 2, 2, 309,
	310, 7, 60, 2, 2, 310, 311, 7, 60, 2, 2, 311, 88, 3, 2, 2, 2, 312, 313,
	7, 47, 2, 2, 313, 314, 7, 47, 2, 2, 314, 90, 3, 2, 2, 2, 315, 316, 7, 45,
	2, 2, 316, 317, 7, 45, 2, 2, 317, 92, 3, 2, 2, 2, 318, 319, 7, 128, 2,
	2, 319, 94, 3, 2, 2, 2, 320, 321, 7, 65, 2, 2, 321, 96, 3, 2, 2, 2, 322,
	323, 7, 47, 2, 2, 323, 324, 7, 64, 2, 2, 324, 98, 3, 2, 2, 2, 325, 326,
	7, 38, 2, 2, 326, 100, 3, 2, 2, 2, 327, 328, 7, 99, 2, 2, 328, 329, 7,
	102, 2, 2, 329, 330, 7, 118, 2, 2, 330, 102, 3, 2, 2, 2, 331, 332, 7, 99,
	2, 2, 332, 333, 7, 105, 2, 2, 333, 334, 7, 105, 2, 2, 334, 335, 7, 116,
	2, 2, 335, 104, 3, 2, 2, 2, 336, 337, 7, 99, 2, 2, 337, 338, 7, 110, 2,
	2, 338, 339, 7, 110, 2, 2, 339, 340, 7, 113, 2, 2, 340, 341, 7, 101, 2,
	2, 341, 106, 3, 2, 2, 2, 342, 343, 7, 99, 2, 2, 343, 344, 7, 110, 2, 2,
	344, 345, 7, 118, 2, 2, 345, 108, 3, 2, 2, 2, 346, 347, 7, 100, 2, 2, 347,
	348, 7, 103, 2, 2, 348, 349, 7, 101, 2, 2, 349, 350, 7, 113, 2, 2, 350,
	351, 7, 111, 2, 2, 351, 352, 7, 103, 2, 2, 352, 110, 3, 2, 2, 2, 353, 354,
	7, 100, 2, 2, 354, 355, 7, 116, 2, 2, 355, 356, 7, 103, 2, 2, 356, 357,
	7, 99, 2, 2, 357, 358, 7, 109, 2, 2, 358, 112, 3, 2, 2, 2, 359, 360, 7,
	100, 2, 2, 360, 361, 7, 123, 2, 2, 361, 362, 7, 118, 2, 2, 362, 363, 7,
	103, 2, 2, 363, 114, 3, 2, 2, 2, 364, 365, 7, 101, 2, 2, 365, 366, 7, 99,
	2, 2, 366, 367, 7, 117, 2, 2, 367, 368, 7, 103, 2, 2, 368, 116, 3, 2, 2,
	2, 369, 370, 7, 101, 2, 2, 370, 371, 7, 106, 2, 2, 371, 372, 7, 99, 2,
	2, 372, 373, 7, 112, 2, 2, 373, 118, 3, 2, 2, 2, 374, 375, 7, 101, 2, 2,
	375, 376, 7, 106, 2, 2, 376, 377, 7, 103, 2, 2, 377, 378, 7, 101, 2, 2,
	378, 379, 7, 109, 2, 2, 379, 120, 3, 2, 2, 2, 380, 381, 7, 101, 2, 2, 381,
	382, 7, 113, 2, 2, 382, 383, 7, 112, 2, 2, 383, 384, 7, 118, 2, 2, 384,
	385, 7, 107, 2, 2, 385, 386, 7, 112, 2, 2, 386, 387, 7, 119, 2, 2, 387,
	388, 7, 103, 2, 2, 388, 122, 3, 2, 2, 2, 389, 390, 7, 102, 2, 2, 390, 391,
	7, 103, 2, 2, 391, 392, 7, 104, 2, 2, 392, 393, 7, 99, 2, 2, 393, 394,
	7, 119, 2, 2, 394, 395, 7, 110, 2, 2, 395, 396, 7, 118, 2, 2, 396, 124,
	3, 2, 2, 2, 397, 398, 7, 102, 2, 2, 398, 399, 7, 113, 2, 2, 399, 126, 3,
	2, 2, 2, 400, 401, 7, 103, 2, 2, 401, 402, 7, 110, 2, 2, 402, 403, 7, 117,
	2, 2, 403, 404, 7, 103, 2, 2, 404, 128, 3, 2, 2, 2, 405, 406, 7, 103, 2,
	2, 406, 407, 7, 112, 2, 2, 407, 408, 7, 119, 2, 2, 408, 409, 7, 111, 2,
	2, 409, 130, 3, 2, 2, 2, 410, 411, 7, 103, 2, 2, 411, 412, 7, 122, 2, 2,
	412, 413, 7, 118, 2, 2, 413, 414, 7, 103, 2, 2, 414, 415, 7, 116, 2, 2,
	415, 416, 7, 112, 2, 2, 416, 132, 3, 2, 2, 2, 417, 418, 7, 104, 2, 2, 418,
	419, 7, 110, 2, 2, 419, 420, 7, 113, 2, 2, 420, 421, 7, 99, 2, 2, 421,
	422, 7, 118, 2, 2, 422, 134, 3, 2, 2, 2, 423, 424, 7, 104, 2, 2, 424, 425,
	7, 113, 2, 2, 425, 426, 7, 116, 2, 2, 426, 136, 3, 2, 2, 2, 427, 428, 7,
	105, 2, 2, 428, 429, 7, 113, 2, 2, 429, 430, 7, 118, 2, 2, 430, 431, 7,
	113, 2, 2, 431, 138, 3, 2, 2, 2, 432, 433, 7, 107, 2, 2, 433, 434, 7, 104,
	2, 2, 434, 140, 3, 2, 2, 2, 435, 436, 7, 107, 2, 2, 436, 437, 7, 112, 2,
	2, 437, 438, 7, 118, 2, 2, 438, 142, 3, 2, 2, 2, 439, 440, 7, 107, 2, 2,
	440, 441, 7, 112, 2, 2, 441, 442, 7, 118, 2, 2, 442, 443, 7, 103, 2, 2,
	443, 444, 7, 116, 2, 2, 444, 445, 7, 112, 2, 2, 445, 144, 3, 2, 2, 2, 446,
	447, 7, 110, 2, 2, 447, 448, 7, 107, 2, 2, 448, 449, 7, 112, 2, 2, 449,
	450, 7, 118, 2, 2, 450, 146, 3, 2, 2, 2, 451, 452, 7, 112, 2, 2, 452, 453,
	7, 107, 2, 2, 453, 454, 7, 110, 2, 2, 454, 148, 3, 2, 2, 2, 455, 456, 7,
	114, 2, 2, 456, 457, 7, 99, 2, 2, 457, 458, 7, 116, 2, 2, 458, 150, 3,
	2, 2, 2, 459, 460, 7, 114, 2, 2, 460, 461, 7, 116, 2, 2, 461, 462, 7, 113,
	2, 2, 462, 463, 7, 101, 2, 2, 463, 152, 3, 2, 2, 2, 464, 465, 7, 116, 2,
	2, 465, 466, 7, 99, 2, 2, 466, 467, 7, 107, 2, 2, 467, 468, 7, 117, 2,
	2, 468, 469, 7, 103, 2, 2, 469, 154, 3, 2, 2, 2, 470, 471, 7, 116, 2, 2,
	471, 472, 7, 103, 2, 2, 472, 473, 7, 117, 2, 2, 473, 474, 7, 101, 2, 2,
	474, 475, 7, 119, 2, 2, 475, 476, 7, 103, 2, 2, 476, 156, 3, 2, 2, 2, 477,
	478, 7, 116, 2, 2, 478, 479, 7, 103, 2, 2, 479, 480, 7, 118, 2, 2, 480,
	481, 7, 119, 2, 2, 481, 482, 7, 116, 2, 2, 482, 483, 7, 112, 2, 2, 483,
	158, 3, 2, 2, 2, 484, 485, 7, 117, 2, 2, 485, 486, 7, 107, 2, 2, 486, 487,
	7, 112, 2, 2, 487, 488, 7, 118, 2, 2, 488, 160, 3, 2, 2, 2, 489, 490, 7,
	117, 2, 2, 490, 491, 7, 107, 2, 2, 491, 492, 7, 124, 2, 2, 492, 493, 7,
	103, 2, 2, 493, 494, 7, 113, 2, 2, 494, 495, 7, 104, 2, 2, 495, 162, 3,
	2, 2, 2, 496, 497, 7, 117, 2, 2, 497, 498, 7, 121, 2, 2, 498, 499, 7, 107,
	2, 2, 499, 500, 7, 118, 2, 2, 500, 501, 7, 101, 2, 2, 501, 502, 7, 106,
	2, 2, 502, 164, 3, 2, 2, 2, 503, 504, 7, 118, 2, 2, 504, 505, 7, 99, 2,
	2, 505, 506, 7, 117, 2, 2, 506, 507, 7, 109, 2, 2, 507, 166, 3, 2, 2, 2,
	508, 509, 7, 118, 2, 2, 509, 510, 7, 119, 2, 2, 510, 511, 7, 114, 2, 2,
	511, 512, 7, 110, 2, 2, 512, 513, 7, 103, 2, 2, 513, 168, 3, 2, 2, 2, 514,
	515, 7, 118, 2, 2, 515, 516, 7, 123, 2, 2, 516, 517, 7, 114, 2, 2, 517,
	518, 7, 103, 2, 2, 518, 519, 7, 102, 2, 2, 519, 520, 7, 103, 2, 2, 520,
	521, 7, 104, 2, 2, 521, 170, 3, 2, 2, 2, 522, 523, 7, 118, 2, 2, 523, 524,
	7, 123, 2, 2, 524, 525, 7, 114, 2, 2, 525, 526, 7, 103, 2, 2, 526, 527,
	7, 113, 2, 2, 527, 528, 7, 104, 2, 2, 528, 172, 3, 2, 2, 2, 529, 530, 7,
	119, 2, 2, 530, 531, 7, 107, 2, 2, 531, 532, 7, 112, 2, 2, 532, 533, 7,
	118, 2, 2, 533, 174, 3, 2, 2, 2, 534, 535, 7, 119, 2, 2, 535, 536, 7, 110,
	2, 2, 536, 537, 7, 107, 2, 2, 537, 538, 7, 112, 2, 2, 538, 539, 7, 118,
	2, 2, 539, 176, 3, 2, 2, 2, 540, 541, 7, 119, 2, 2, 541, 542, 7, 112, 2,
	2, 542, 543, 7, 99, 2, 2, 543, 544, 7, 110, 2, 2, 544, 545, 7, 110, 2,
	2, 545, 546, 7, 113, 2, 2, 546, 547, 7, 101, 2, 2, 547, 178, 3, 2, 2, 2,
	548, 549, 7, 119, 2, 2, 549, 550, 7, 112, 2, 2, 550, 551, 7, 107, 2, 2,
	551, 552, 7, 113, 2, 2, 552, 553, 7, 112, 2, 2, 553, 180, 3, 2, 2, 2, 554,
	555, 7, 119, 2, 2, 555, 556, 7, 117, 2, 2, 556, 557, 7, 107, 2, 2, 557,
	558, 7, 112, 2, 2, 558, 559, 7, 118, 2, 2, 559, 182, 3, 2, 2, 2, 560, 561,
	7, 120, 2, 2, 561, 562, 7, 113, 2, 2, 562, 563, 7, 107, 2, 2, 563, 564,
	7, 102, 2, 2, 564, 184, 3, 2, 2, 2, 565, 566, 7, 121, 2, 2, 566, 567, 7,
	106, 2, 2, 567, 568, 7, 107, 2, 2, 568, 569, 7, 110, 2, 2, 569, 570, 7,
	103, 2, 2, 570, 186, 3, 2, 2, 2, 571, 572, 7, 124, 2, 2, 572, 573, 7, 103,
	2, 2, 573, 574, 7, 116, 2, 2, 574, 575, 7, 113, 2, 2, 575, 576, 7, 122,
	2, 2, 576, 188, 3, 2, 2, 2, 577, 578, 7, 114, 2, 2, 578, 579, 7, 116, 2,
	2, 579, 580, 7, 107, 2, 2, 580, 581, 7, 120, 2, 2, 581, 582, 7, 99, 2,
	2, 582, 583, 7, 118, 2, 2, 583, 584, 7, 103, 2, 2, 584, 190, 3, 2, 2, 2,
	585, 589, 9, 2, 2, 2, 586, 588, 9, 3, 2, 2, 587, 586, 3, 2, 2, 2, 588,
	591, 3, 2, 2, 2, 589, 587, 3, 2, 2, 2, 589, 590, 3, 2, 2, 2, 590, 192,
	3, 2, 2, 2, 591, 589, 3, 2, 2, 2, 592, 596, 7, 36, 2, 2, 593, 595, 10,
	4, 2, 2, 594, 593, 3, 2, 2, 2, 595, 598, 3, 2, 2, 2, 596, 594, 3, 2, 2,
	2, 596, 597, 3, 2, 2, 2, 597, 599, 3, 2, 2, 2, 598, 596, 3, 2, 2, 2, 599,
	600, 7, 36, 2, 2, 600, 194, 3, 2, 2, 2, 601, 605, 7, 41, 2, 2, 602, 604,
	10, 5, 2, 2, 603, 602, 3, 2, 2, 2, 604, 607, 3, 2, 2, 2, 605, 603, 3, 2,
	2, 2, 605, 606, 3, 2, 2, 2, 606, 608, 3, 2, 2, 2, 607, 605, 3, 2, 2, 2,
	608, 609, 7, 41, 2, 2, 609, 196, 3, 2, 2, 2, 610, 612, 5, 199, 100, 2,
	611, 610, 3, 2, 2, 2, 612, 613, 3, 2, 2, 2, 613, 611, 3, 2, 2, 2, 613,
	614, 3, 2, 2, 2, 614, 622, 3, 2, 2, 2, 615, 619, 7, 48, 2, 2, 616, 618,
	5, 199, 100, 2, 617, 616, 3, 2, 2, 2, 618, 621, 3, 2, 2, 2, 619, 617, 3,
	2, 2, 2, 619, 620, 3, 2, 2, 2, 620, 623, 3, 2, 2, 2, 621, 619, 3, 2, 2,
	2, 622, 615, 3, 2, 2, 2, 622, 623, 3, 2, 2, 2, 623, 630, 3, 2, 2, 2, 624,
	626, 7, 103, 2, 2, 625, 627, 5, 199, 100, 2, 626, 625, 3, 2, 2, 2, 627,
	628, 3, 2, 2, 2, 628, 626, 3, 2, 2, 2, 628, 629, 3, 2, 2, 2, 629, 631,
	3, 2, 2, 2, 630, 624, 3, 2, 2, 2, 630, 631, 3, 2, 2, 2, 631, 198, 3, 2,
	2, 2, 632, 633, 9, 6, 2, 2, 633, 200, 3, 2, 2, 2, 634, 636, 9, 7, 2, 2,
	635, 634, 3, 2, 2, 2, 636, 637, 3, 2, 2, 2, 637, 635, 3, 2, 2, 2, 637,
	638, 3, 2, 2, 2, 638, 639, 3, 2, 2, 2, 639, 640, 8, 101, 2, 2, 640, 202,
	3, 2, 2, 2, 12, 2, 589, 596, 605, 613, 619, 622, 628, 630, 637, 3, 8, 2,
	2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'('", "')'", "'.'", "'{'", "'}'", "'['", "']'", "','", "'='",
	"'*'", "'!'", "':'", "'/'", "'%'", "'+'", "'-'", "'>>'", "'<<'", "'<'",
	"'>'", "'<='", "'>='", "'=='", "'!='", "'&'", "'^'", "'|'", "'&&'", "'||'",
	"':='", "'<-'", "'+='", "'-='", "'*='", "'/='", "'%='", "'>>='", "'<<='",
	"'&='", "'|='", "'^='", "'::'", "'--'", "'++'", "'~'", "'?'", "'->'", "'$'",
	"'adt'", "'aggr'", "'alloc'", "'alt'", "'become'", "'break'", "'byte'",
	"'case'", "'chan'", "'check'", "'continue'", "'default'", "'do'", "'else'",
	"'enum'", "'extern'", "'float'", "'for'", "'goto'", "'if'", "'int'", "'intern'",
	"'lint'", "'nil'", "'par'", "'proc'", "'raise'", "'rescue'", "'return'",
	"'sint'", "'sizeof'", "'switch'", "'task'", "'tuple'", "'typedef'", "'typeof'",
	"'uint'", "'ulint'", "'unalloc'", "'union'", "'usint'", "'void'", "'while'",
	"'zerox'", "'private'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "ADT", "AGGR",
	"ALLOC", "ALT", "BECOME", "BREAK", "BYTE", "CASE", "CHAN", "CHECK", "CONTINUE",
	"DEFAULT", "DO", "ELSE", "ENUM", "EXTERN", "FLOAT", "FOR", "GOTO", "IF",
	"INT", "INTERN", "LINT", "NIL", "PAR", "PROC", "RAISE", "RESCUE", "RETURN",
	"SINT", "SIZEOF", "SWITCH", "TASK", "TUPLE", "TYPEDEF", "TYPEOF", "UINT",
	"ULINT", "UNALLOC", "UNION", "USINT", "VOID", "WHILE", "ZEROX", "PRIVATE",
	"IDENTIFIER", "STRING_CONST", "CONSTANT", "ARITHMETIC_CONST", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
	"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
	"ADT", "AGGR", "ALLOC", "ALT", "BECOME", "BREAK", "BYTE", "CASE", "CHAN",
	"CHECK", "CONTINUE", "DEFAULT", "DO", "ELSE", "ENUM", "EXTERN", "FLOAT",
	"FOR", "GOTO", "IF", "INT", "INTERN", "LINT", "NIL", "PAR", "PROC", "RAISE",
	"RESCUE", "RETURN", "SINT", "SIZEOF", "SWITCH", "TASK", "TUPLE", "TYPEDEF",
	"TYPEOF", "UINT", "ULINT", "UNALLOC", "UNION", "USINT", "VOID", "WHILE",
	"ZEROX", "PRIVATE", "IDENTIFIER", "STRING_CONST", "CONSTANT", "ARITHMETIC_CONST",
	"DIGIT", "WS",
}

type alefLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewalefLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *alefLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewalefLexer(input antlr.CharStream) *alefLexer {
	l := new(alefLexer)
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
	l.GrammarFileName = "alef.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// alefLexer tokens.
const (
	alefLexerT__0             = 1
	alefLexerT__1             = 2
	alefLexerT__2             = 3
	alefLexerT__3             = 4
	alefLexerT__4             = 5
	alefLexerT__5             = 6
	alefLexerT__6             = 7
	alefLexerT__7             = 8
	alefLexerT__8             = 9
	alefLexerT__9             = 10
	alefLexerT__10            = 11
	alefLexerT__11            = 12
	alefLexerT__12            = 13
	alefLexerT__13            = 14
	alefLexerT__14            = 15
	alefLexerT__15            = 16
	alefLexerT__16            = 17
	alefLexerT__17            = 18
	alefLexerT__18            = 19
	alefLexerT__19            = 20
	alefLexerT__20            = 21
	alefLexerT__21            = 22
	alefLexerT__22            = 23
	alefLexerT__23            = 24
	alefLexerT__24            = 25
	alefLexerT__25            = 26
	alefLexerT__26            = 27
	alefLexerT__27            = 28
	alefLexerT__28            = 29
	alefLexerT__29            = 30
	alefLexerT__30            = 31
	alefLexerT__31            = 32
	alefLexerT__32            = 33
	alefLexerT__33            = 34
	alefLexerT__34            = 35
	alefLexerT__35            = 36
	alefLexerT__36            = 37
	alefLexerT__37            = 38
	alefLexerT__38            = 39
	alefLexerT__39            = 40
	alefLexerT__40            = 41
	alefLexerT__41            = 42
	alefLexerT__42            = 43
	alefLexerT__43            = 44
	alefLexerT__44            = 45
	alefLexerT__45            = 46
	alefLexerT__46            = 47
	alefLexerT__47            = 48
	alefLexerT__48            = 49
	alefLexerADT              = 50
	alefLexerAGGR             = 51
	alefLexerALLOC            = 52
	alefLexerALT              = 53
	alefLexerBECOME           = 54
	alefLexerBREAK            = 55
	alefLexerBYTE             = 56
	alefLexerCASE             = 57
	alefLexerCHAN             = 58
	alefLexerCHECK            = 59
	alefLexerCONTINUE         = 60
	alefLexerDEFAULT          = 61
	alefLexerDO               = 62
	alefLexerELSE             = 63
	alefLexerENUM             = 64
	alefLexerEXTERN           = 65
	alefLexerFLOAT            = 66
	alefLexerFOR              = 67
	alefLexerGOTO             = 68
	alefLexerIF               = 69
	alefLexerINT              = 70
	alefLexerINTERN           = 71
	alefLexerLINT             = 72
	alefLexerNIL              = 73
	alefLexerPAR              = 74
	alefLexerPROC             = 75
	alefLexerRAISE            = 76
	alefLexerRESCUE           = 77
	alefLexerRETURN           = 78
	alefLexerSINT             = 79
	alefLexerSIZEOF           = 80
	alefLexerSWITCH           = 81
	alefLexerTASK             = 82
	alefLexerTUPLE            = 83
	alefLexerTYPEDEF          = 84
	alefLexerTYPEOF           = 85
	alefLexerUINT             = 86
	alefLexerULINT            = 87
	alefLexerUNALLOC          = 88
	alefLexerUNION            = 89
	alefLexerUSINT            = 90
	alefLexerVOID             = 91
	alefLexerWHILE            = 92
	alefLexerZEROX            = 93
	alefLexerPRIVATE          = 94
	alefLexerIDENTIFIER       = 95
	alefLexerSTRING_CONST     = 96
	alefLexerCONSTANT         = 97
	alefLexerARITHMETIC_CONST = 98
	alefLexerWS               = 99
)