// Code generated from pike.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pike

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 95, 641,
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
	91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 34,
	3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3,
	37, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 44, 3,
	44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48,
	3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3,
	53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56,
	3, 57, 3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3,
	61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65,
	3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 67, 3, 67, 3, 67, 3, 67, 3,
	67, 3, 67, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 69, 3, 69,
	3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 3,
	72, 3, 72, 3, 72, 3, 73, 3, 73, 3, 73, 3, 74, 3, 74, 3, 74, 3, 75, 3, 75,
	3, 75, 3, 76, 3, 76, 3, 76, 3, 77, 3, 77, 3, 78, 3, 78, 3, 78, 3, 78, 3,
	79, 3, 79, 3, 79, 3, 79, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80, 3, 80,
	3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 82, 3, 82, 3, 82, 3, 82, 3,
	82, 3, 82, 3, 82, 3, 82, 3, 83, 3, 83, 3, 83, 3, 83, 3, 83, 3, 83, 3, 83,
	3, 84, 3, 84, 3, 84, 3, 84, 3, 84, 3, 84, 3, 84, 3, 84, 3, 85, 3, 85, 3,
	85, 3, 85, 3, 85, 3, 85, 3, 86, 3, 86, 3, 86, 3, 86, 3, 86, 3, 86, 3, 86,
	3, 86, 3, 86, 3, 87, 3, 87, 3, 87, 3, 87, 3, 87, 3, 87, 3, 87, 3, 87, 3,
	87, 3, 88, 3, 88, 3, 88, 7, 88, 558, 10, 88, 12, 88, 14, 88, 561, 11, 88,
	3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3,
	88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88,
	3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 5, 88, 591, 10, 88, 3,
	89, 3, 89, 3, 90, 3, 90, 3, 91, 3, 91, 7, 91, 599, 10, 91, 12, 91, 14,
	91, 602, 11, 91, 3, 91, 3, 91, 7, 91, 606, 10, 91, 12, 91, 14, 91, 609,
	11, 91, 3, 92, 3, 92, 7, 92, 613, 10, 92, 12, 92, 14, 92, 616, 11, 92,
	3, 92, 3, 92, 3, 92, 3, 92, 7, 92, 622, 10, 92, 12, 92, 14, 92, 625, 11,
	92, 5, 92, 627, 10, 92, 3, 93, 3, 93, 7, 93, 631, 10, 93, 12, 93, 14, 93,
	634, 11, 93, 3, 93, 3, 93, 3, 94, 3, 94, 3, 94, 3, 94, 2, 2, 95, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22,
	43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31,
	61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40,
	79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49,
	97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111, 57, 113,
	58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 63, 125, 64, 127, 65, 129,
	66, 131, 67, 133, 68, 135, 69, 137, 70, 139, 71, 141, 72, 143, 73, 145,
	74, 147, 75, 149, 76, 151, 77, 153, 78, 155, 79, 157, 80, 159, 81, 161,
	82, 163, 83, 165, 84, 167, 85, 169, 86, 171, 87, 173, 88, 175, 89, 177,
	90, 179, 91, 181, 92, 183, 93, 185, 94, 187, 95, 3, 2, 6, 9, 2, 39, 40,
	44, 45, 49, 49, 62, 62, 96, 96, 126, 126, 128, 128, 5, 2, 67, 92, 97, 97,
	99, 124, 3, 2, 36, 36, 5, 2, 11, 12, 15, 15, 34, 34, 2, 663, 2, 3, 3, 2,
	2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2,
	2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3,
	2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27,
	3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2,
	35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2,
	2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2,
	2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2,
	2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3,
	2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73,
	3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2,
	81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2,
	2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2,
	2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3,
	2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2,
	111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2,
	2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125,
	3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2,
	2, 133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3,
	2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2,
	147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 2, 153, 3, 2,
	2, 2, 2, 155, 3, 2, 2, 2, 2, 157, 3, 2, 2, 2, 2, 159, 3, 2, 2, 2, 2, 161,
	3, 2, 2, 2, 2, 163, 3, 2, 2, 2, 2, 165, 3, 2, 2, 2, 2, 167, 3, 2, 2, 2,
	2, 169, 3, 2, 2, 2, 2, 171, 3, 2, 2, 2, 2, 173, 3, 2, 2, 2, 2, 175, 3,
	2, 2, 2, 2, 177, 3, 2, 2, 2, 2, 179, 3, 2, 2, 2, 2, 181, 3, 2, 2, 2, 2,
	183, 3, 2, 2, 2, 2, 185, 3, 2, 2, 2, 2, 187, 3, 2, 2, 2, 3, 189, 3, 2,
	2, 2, 5, 196, 3, 2, 2, 2, 7, 198, 3, 2, 2, 2, 9, 206, 3, 2, 2, 2, 11, 208,
	3, 2, 2, 2, 13, 210, 3, 2, 2, 2, 15, 212, 3, 2, 2, 2, 17, 214, 3, 2, 2,
	2, 19, 216, 3, 2, 2, 2, 21, 218, 3, 2, 2, 2, 23, 227, 3, 2, 2, 2, 25, 233,
	3, 2, 2, 2, 27, 235, 3, 2, 2, 2, 29, 237, 3, 2, 2, 2, 31, 244, 3, 2, 2,
	2, 33, 252, 3, 2, 2, 2, 35, 259, 3, 2, 2, 2, 37, 266, 3, 2, 2, 2, 39, 276,
	3, 2, 2, 2, 41, 283, 3, 2, 2, 2, 43, 286, 3, 2, 2, 2, 45, 291, 3, 2, 2,
	2, 47, 297, 3, 2, 2, 2, 49, 300, 3, 2, 2, 2, 51, 304, 3, 2, 2, 2, 53, 311,
	3, 2, 2, 2, 55, 316, 3, 2, 2, 2, 57, 319, 3, 2, 2, 2, 59, 327, 3, 2, 2,
	2, 61, 335, 3, 2, 2, 2, 63, 341, 3, 2, 2, 2, 65, 350, 3, 2, 2, 2, 67, 353,
	3, 2, 2, 2, 69, 356, 3, 2, 2, 2, 71, 359, 3, 2, 2, 2, 73, 362, 3, 2, 2,
	2, 75, 365, 3, 2, 2, 2, 77, 368, 3, 2, 2, 2, 79, 372, 3, 2, 2, 2, 81, 376,
	3, 2, 2, 2, 83, 379, 3, 2, 2, 2, 85, 381, 3, 2, 2, 2, 87, 384, 3, 2, 2,
	2, 89, 387, 3, 2, 2, 2, 91, 389, 3, 2, 2, 2, 93, 391, 3, 2, 2, 2, 95, 393,
	3, 2, 2, 2, 97, 396, 3, 2, 2, 2, 99, 399, 3, 2, 2, 2, 101, 401, 3, 2, 2,
	2, 103, 403, 3, 2, 2, 2, 105, 406, 3, 2, 2, 2, 107, 409, 3, 2, 2, 2, 109,
	412, 3, 2, 2, 2, 111, 415, 3, 2, 2, 2, 113, 417, 3, 2, 2, 2, 115, 419,
	3, 2, 2, 2, 117, 421, 3, 2, 2, 2, 119, 424, 3, 2, 2, 2, 121, 427, 3, 2,
	2, 2, 123, 429, 3, 2, 2, 2, 125, 431, 3, 2, 2, 2, 127, 434, 3, 2, 2, 2,
	129, 436, 3, 2, 2, 2, 131, 438, 3, 2, 2, 2, 133, 444, 3, 2, 2, 2, 135,
	450, 3, 2, 2, 2, 137, 457, 3, 2, 2, 2, 139, 464, 3, 2, 2, 2, 141, 466,
	3, 2, 2, 2, 143, 469, 3, 2, 2, 2, 145, 472, 3, 2, 2, 2, 147, 475, 3, 2,
	2, 2, 149, 478, 3, 2, 2, 2, 151, 481, 3, 2, 2, 2, 153, 484, 3, 2, 2, 2,
	155, 486, 3, 2, 2, 2, 157, 490, 3, 2, 2, 2, 159, 494, 3, 2, 2, 2, 161,
	501, 3, 2, 2, 2, 163, 507, 3, 2, 2, 2, 165, 515, 3, 2, 2, 2, 167, 522,
	3, 2, 2, 2, 169, 530, 3, 2, 2, 2, 171, 536, 3, 2, 2, 2, 173, 545, 3, 2,
	2, 2, 175, 590, 3, 2, 2, 2, 177, 592, 3, 2, 2, 2, 179, 594, 3, 2, 2, 2,
	181, 596, 3, 2, 2, 2, 183, 626, 3, 2, 2, 2, 185, 628, 3, 2, 2, 2, 187,
	637, 3, 2, 2, 2, 189, 190, 7, 107, 2, 2, 190, 191, 7, 111, 2, 2, 191, 192,
	7, 114, 2, 2, 192, 193, 7, 113, 2, 2, 193, 194, 7, 116, 2, 2, 194, 195,
	7, 118, 2, 2, 195, 4, 3, 2, 2, 2, 196, 197, 7, 61, 2, 2, 197, 6, 3, 2,
	2, 2, 198, 199, 7, 107, 2, 2, 199, 200, 7, 112, 2, 2, 200, 201, 7, 106,
	2, 2, 201, 202, 7, 103, 2, 2, 202, 203, 7, 116, 2, 2, 203, 204, 7, 107,
	2, 2, 204, 205, 7, 118, 2, 2, 205, 8, 3, 2, 2, 2, 206, 207, 7, 60, 2, 2,
	207, 10, 3, 2, 2, 2, 208, 209, 7, 42, 2, 2, 209, 12, 3, 2, 2, 2, 210, 211,
	7, 43, 2, 2, 211, 14, 3, 2, 2, 2, 212, 213, 7, 46, 2, 2, 213, 16, 3, 2,
	2, 2, 214, 215, 7, 44, 2, 2, 215, 18, 3, 2, 2, 2, 216, 217, 7, 63, 2, 2,
	217, 20, 3, 2, 2, 2, 218, 219, 7, 101, 2, 2, 219, 220, 7, 113, 2, 2, 220,
	221, 7, 112, 2, 2, 221, 222, 7, 117, 2, 2, 222, 223, 7, 118, 2, 2, 223,
	224, 7, 99, 2, 2, 224, 225, 7, 112, 2, 2, 225, 226, 7, 118, 2, 2, 226,
	22, 3, 2, 2, 2, 227, 228, 7, 101, 2, 2, 228, 229, 7, 110, 2, 2, 229, 230,
	7, 99, 2, 2, 230, 231, 7, 117, 2, 2, 231, 232, 7, 117, 2, 2, 232, 24, 3,
	2, 2, 2, 233, 234, 7, 125, 2, 2, 234, 26, 3, 2, 2, 2, 235, 236, 7, 127,
	2, 2, 236, 28, 3, 2, 2, 2, 237, 238, 7, 117, 2, 2, 238, 239, 7, 118, 2,
	2, 239, 240, 7, 99, 2, 2, 240, 241, 7, 118, 2, 2, 241, 242, 7, 107, 2,
	2, 242, 243, 7, 101, 2, 2, 243, 30, 3, 2, 2, 2, 244, 245, 7, 114, 2, 2,
	245, 246, 7, 116, 2, 2, 246, 247, 7, 107, 2, 2, 247, 248, 7, 120, 2, 2,
	248, 249, 7, 99, 2, 2, 249, 250, 7, 118, 2, 2, 250, 251, 7, 103, 2, 2,
	251, 32, 3, 2, 2, 2, 252, 253, 7, 112, 2, 2, 253, 254, 7, 113, 2, 2, 254,
	255, 7, 111, 2, 2, 255, 256, 7, 99, 2, 2, 256, 257, 7, 117, 2, 2, 257,
	258, 7, 109, 2, 2, 258, 34, 3, 2, 2, 2, 259, 260, 7, 114, 2, 2, 260, 261,
	7, 119, 2, 2, 261, 262, 7, 100, 2, 2, 262, 263, 7, 110, 2, 2, 263, 264,
	7, 107, 2, 2, 264, 265, 7, 101, 2, 2, 265, 36, 3, 2, 2, 2, 266, 267, 7,
	114, 2, 2, 267, 268, 7, 116, 2, 2, 268, 269, 7, 113, 2, 2, 269, 270, 7,
	118, 2, 2, 270, 271, 7, 103, 2, 2, 271, 272, 7, 101, 2, 2, 272, 273, 7,
	118, 2, 2, 273, 274, 7, 103, 2, 2, 274, 275, 7, 102, 2, 2, 275, 38, 3,
	2, 2, 2, 276, 277, 7, 107, 2, 2, 277, 278, 7, 112, 2, 2, 278, 279, 7, 110,
	2, 2, 279, 280, 7, 107, 2, 2, 280, 281, 7, 112, 2, 2, 281, 282, 7, 103,
	2, 2, 282, 40, 3, 2, 2, 2, 283, 284, 7, 107, 2, 2, 284, 285, 7, 104, 2,
	2, 285, 42, 3, 2, 2, 2, 286, 287, 7, 103, 2, 2, 287, 288, 7, 110, 2, 2,
	288, 289, 7, 117, 2, 2, 289, 290, 7, 103, 2, 2, 290, 44, 3, 2, 2, 2, 291,
	292, 7, 121, 2, 2, 292, 293, 7, 106, 2, 2, 293, 294, 7, 107, 2, 2, 294,
	295, 7, 110, 2, 2, 295, 296, 7, 103, 2, 2, 296, 46, 3, 2, 2, 2, 297, 298,
	7, 102, 2, 2, 298, 299, 7, 113, 2, 2, 299, 48, 3, 2, 2, 2, 300, 301, 7,
	104, 2, 2, 301, 302, 7, 113, 2, 2, 302, 303, 7, 116, 2, 2, 303, 50, 3,
	2, 2, 2, 304, 305, 7, 117, 2, 2, 305, 306, 7, 121, 2, 2, 306, 307, 7, 107,
	2, 2, 307, 308, 7, 118, 2, 2, 308, 309, 7, 101, 2, 2, 309, 310, 7, 106,
	2, 2, 310, 52, 3, 2, 2, 2, 311, 312, 7, 101, 2, 2, 312, 313, 7, 99, 2,
	2, 313, 314, 7, 117, 2, 2, 314, 315, 7, 103, 2, 2, 315, 54, 3, 2, 2, 2,
	316, 317, 7, 48, 2, 2, 317, 318, 7, 48, 2, 2, 318, 56, 3, 2, 2, 2, 319,
	320, 7, 102, 2, 2, 320, 321, 7, 103, 2, 2, 321, 322, 7, 104, 2, 2, 322,
	323, 7, 99, 2, 2, 323, 324, 7, 119, 2, 2, 324, 325, 7, 110, 2, 2, 325,
	326, 7, 118, 2, 2, 326, 58, 3, 2, 2, 2, 327, 328, 7, 104, 2, 2, 328, 329,
	7, 113, 2, 2, 329, 330, 7, 116, 2, 2, 330, 331, 7, 103, 2, 2, 331, 332,
	7, 99, 2, 2, 332, 333, 7, 101, 2, 2, 333, 334, 7, 106, 2, 2, 334, 60, 3,
	2, 2, 2, 335, 336, 7, 100, 2, 2, 336, 337, 7, 116, 2, 2, 337, 338, 7, 103,
	2, 2, 338, 339, 7, 99, 2, 2, 339, 340, 7, 109, 2, 2, 340, 62, 3, 2, 2,
	2, 341, 342, 7, 101, 2, 2, 342, 343, 7, 113, 2, 2, 343, 344, 7, 112, 2,
	2, 344, 345, 7, 118, 2, 2, 345, 346, 7, 107, 2, 2, 346, 347, 7, 112, 2,
	2, 347, 348, 7, 119, 2, 2, 348, 349, 7, 103, 2, 2, 349, 64, 3, 2, 2, 2,
	350, 351, 7, 45, 2, 2, 351, 352, 7, 63, 2, 2, 352, 66, 3, 2, 2, 2, 353,
	354, 7, 44, 2, 2, 354, 355, 7, 63, 2, 2, 355, 68, 3, 2, 2, 2, 356, 357,
	7, 49, 2, 2, 357, 358, 7, 63, 2, 2, 358, 70, 3, 2, 2, 2, 359, 360, 7, 40,
	2, 2, 360, 361, 7, 63, 2, 2, 361, 72, 3, 2, 2, 2, 362, 363, 7, 126, 2,
	2, 363, 364, 7, 63, 2, 2, 364, 74, 3, 2, 2, 2, 365, 366, 7, 96, 2, 2, 366,
	367, 7, 63, 2, 2, 367, 76, 3, 2, 2, 2, 368, 369, 7, 62, 2, 2, 369, 370,
	7, 62, 2, 2, 370, 371, 7, 63, 2, 2, 371, 78, 3, 2, 2, 2, 372, 373, 7, 64,
	2, 2, 373, 374, 7, 64, 2, 2, 374, 375, 7, 63, 2, 2, 375, 80, 3, 2, 2, 2,
	376, 377, 7, 39, 2, 2, 377, 378, 7, 63, 2, 2, 378, 82, 3, 2, 2, 2, 379,
	380, 7, 65, 2, 2, 380, 84, 3, 2, 2, 2, 381, 382, 7, 126, 2, 2, 382, 383,
	7, 126, 2, 2, 383, 86, 3, 2, 2, 2, 384, 385, 7, 40, 2, 2, 385, 386, 7,
	40, 2, 2, 386, 88, 3, 2, 2, 2, 387, 388, 7, 126, 2, 2, 388, 90, 3, 2, 2,
	2, 389, 390, 7, 96, 2, 2, 390, 92, 3, 2, 2, 2, 391, 392, 7, 40, 2, 2, 392,
	94, 3, 2, 2, 2, 393, 394, 7, 63, 2, 2, 394, 395, 7, 63, 2, 2, 395, 96,
	3, 2, 2, 2, 396, 397, 7, 35, 2, 2, 397, 398, 7, 63, 2, 2, 398, 98, 3, 2,
	2, 2, 399, 400, 7, 64, 2, 2, 400, 100, 3, 2, 2, 2, 401, 402, 7, 62, 2,
	2, 402, 102, 3, 2, 2, 2, 403, 404, 7, 64, 2, 2, 404, 405, 7, 63, 2, 2,
	405, 104, 3, 2, 2, 2, 406, 407, 7, 62, 2, 2, 407, 408, 7, 63, 2, 2, 408,
	106, 3, 2, 2, 2, 409, 410, 7, 62, 2, 2, 410, 411, 7, 62, 2, 2, 411, 108,
	3, 2, 2, 2, 412, 413, 7, 64, 2, 2, 413, 414, 7, 64, 2, 2, 414, 110, 3,
	2, 2, 2, 415, 416, 7, 45, 2, 2, 416, 112, 3, 2, 2, 2, 417, 418, 7, 49,
	2, 2, 418, 114, 3, 2, 2, 2, 419, 420, 7, 39, 2, 2, 420, 116, 3, 2, 2, 2,
	421, 422, 7, 47, 2, 2, 422, 423, 7, 47, 2, 2, 423, 118, 3, 2, 2, 2, 424,
	425, 7, 45, 2, 2, 425, 426, 7, 45, 2, 2, 426, 120, 3, 2, 2, 2, 427, 428,
	7, 128, 2, 2, 428, 122, 3, 2, 2, 2, 429, 430, 7, 47, 2, 2, 430, 124, 3,
	2, 2, 2, 431, 432, 7, 47, 2, 2, 432, 433, 7, 64, 2, 2, 433, 126, 3, 2,
	2, 2, 434, 435, 7, 93, 2, 2, 435, 128, 3, 2, 2, 2, 436, 437, 7, 95, 2,
	2, 437, 130, 3, 2, 2, 2, 438, 439, 7, 101, 2, 2, 439, 440, 7, 99, 2, 2,
	440, 441, 7, 118, 2, 2, 441, 442, 7, 101, 2, 2, 442, 443, 7, 106, 2, 2,
	443, 132, 3, 2, 2, 2, 444, 445, 7, 105, 2, 2, 445, 446, 7, 99, 2, 2, 446,
	447, 7, 119, 2, 2, 447, 448, 7, 105, 2, 2, 448, 449, 7, 103, 2, 2, 449,
	134, 3, 2, 2, 2, 450, 451, 7, 117, 2, 2, 451, 452, 7, 117, 2, 2, 452, 453,
	7, 101, 2, 2, 453, 454, 7, 99, 2, 2, 454, 455, 7, 112, 2, 2, 455, 456,
	7, 104, 2, 2, 456, 136, 3, 2, 2, 2, 457, 458, 7, 110, 2, 2, 458, 459, 7,
	99, 2, 2, 459, 460, 7, 111, 2, 2, 460, 461, 7, 100, 2, 2, 461, 462, 7,
	102, 2, 2, 462, 463, 7, 99, 2, 2, 463, 138, 3, 2, 2, 2, 464, 465, 7, 48,
	2, 2, 465, 140, 3, 2, 2, 2, 466, 467, 7, 42, 2, 2, 467, 468, 7, 125, 2,
	2, 468, 142, 3, 2, 2, 2, 469, 470, 7, 127, 2, 2, 470, 471, 7, 43, 2, 2,
	471, 144, 3, 2, 2, 2, 472, 473, 7, 42, 2, 2, 473, 474, 7, 62, 2, 2, 474,
	146, 3, 2, 2, 2, 475, 476, 7, 64, 2, 2, 476, 477, 7, 43, 2, 2, 477, 148,
	3, 2, 2, 2, 478, 479, 7, 42, 2, 2, 479, 480, 7, 93, 2, 2, 480, 150, 3,
	2, 2, 2, 481, 482, 7, 95, 2, 2, 482, 483, 7, 43, 2, 2, 483, 152, 3, 2,
	2, 2, 484, 485, 7, 66, 2, 2, 485, 154, 3, 2, 2, 2, 486, 487, 7, 48, 2,
	2, 487, 488, 7, 48, 2, 2, 488, 489, 7, 48, 2, 2, 489, 156, 3, 2, 2, 2,
	490, 491, 7, 107, 2, 2, 491, 492, 7, 112, 2, 2, 492, 493, 7, 118, 2, 2,
	493, 158, 3, 2, 2, 2, 494, 495, 7, 117, 2, 2, 495, 496, 7, 118, 2, 2, 496,
	497, 7, 116, 2, 2, 497, 498, 7, 107, 2, 2, 498, 499, 7, 112, 2, 2, 499,
	500, 7, 105, 2, 2, 500, 160, 3, 2, 2, 2, 501, 502, 7, 104, 2, 2, 502, 503,
	7, 110, 2, 2, 503, 504, 7, 113, 2, 2, 504, 505, 7, 99, 2, 2, 505, 506,
	7, 118, 2, 2, 506, 162, 3, 2, 2, 2, 507, 508, 7, 114, 2, 2, 508, 509, 7,
	116, 2, 2, 509, 510, 7, 113, 2, 2, 510, 511, 7, 105, 2, 2, 511, 512, 7,
	116, 2, 2, 512, 513, 7, 99, 2, 2, 513, 514, 7, 111, 2, 2, 514, 164, 3,
	2, 2, 2, 515, 516, 7, 113, 2, 2, 516, 517, 7, 100, 2, 2, 517, 518, 7, 108,
	2, 2, 518, 519, 7, 103, 2, 2, 519, 520, 7, 101, 2, 2, 520, 521, 7, 118,
	2, 2, 521, 166, 3, 2, 2, 2, 522, 523, 7, 111, 2, 2, 523, 524, 7, 99, 2,
	2, 524, 525, 7, 114, 2, 2, 525, 526, 7, 114, 2, 2, 526, 527, 7, 107, 2,
	2, 527, 528, 7, 112, 2, 2, 528, 529, 7, 105, 2, 2, 529, 168, 3, 2, 2, 2,
	530, 531, 7, 99, 2, 2, 531, 532, 7, 116, 2, 2, 532, 533, 7, 116, 2, 2,
	533, 534, 7, 99, 2, 2, 534, 535, 7, 123, 2, 2, 535, 170, 3, 2, 2, 2, 536,
	537, 7, 111, 2, 2, 537, 538, 7, 119, 2, 2, 538, 539, 7, 110, 2, 2, 539,
	540, 7, 118, 2, 2, 540, 541, 7, 107, 2, 2, 541, 542, 7, 117, 2, 2, 542,
	543, 7, 103, 2, 2, 543, 544, 7, 118, 2, 2, 544, 172, 3, 2, 2, 2, 545, 546,
	7, 104, 2, 2, 546, 547, 7, 119, 2, 2, 547, 548, 7, 112, 2, 2, 548, 549,
	7, 101, 2, 2, 549, 550, 7, 118, 2, 2, 550, 551, 7, 107, 2, 2, 551, 552,
	7, 113, 2, 2, 552, 553, 7, 112, 2, 2, 553, 174, 3, 2, 2, 2, 554, 559, 5,
	177, 89, 2, 555, 558, 5, 177, 89, 2, 556, 558, 5, 179, 90, 2, 557, 555,
	3, 2, 2, 2, 557, 556, 3, 2, 2, 2, 558, 561, 3, 2, 2, 2, 559, 557, 3, 2,
	2, 2, 559, 560, 3, 2, 2, 2, 560, 591, 3, 2, 2, 2, 561, 559, 3, 2, 2, 2,
	562, 591, 9, 2, 2, 2, 563, 564, 7, 62, 2, 2, 564, 591, 7, 62, 2, 2, 565,
	566, 7, 62, 2, 2, 566, 591, 7, 63, 2, 2, 567, 591, 7, 64, 2, 2, 568, 569,
	7, 64, 2, 2, 569, 591, 7, 64, 2, 2, 570, 571, 7, 64, 2, 2, 571, 591, 7,
	63, 2, 2, 572, 573, 7, 63, 2, 2, 573, 591, 7, 63, 2, 2, 574, 575, 7, 35,
	2, 2, 575, 591, 7, 63, 2, 2, 576, 591, 7, 35, 2, 2, 577, 578, 7, 42, 2,
	2, 578, 591, 7, 43, 2, 2, 579, 591, 7, 47, 2, 2, 580, 581, 7, 47, 2, 2,
	581, 591, 7, 64, 2, 2, 582, 583, 7, 47, 2, 2, 583, 584, 7, 64, 2, 2, 584,
	591, 7, 63, 2, 2, 585, 586, 7, 93, 2, 2, 586, 591, 7, 95, 2, 2, 587, 588,
	7, 93, 2, 2, 588, 589, 7, 95, 2, 2, 589, 591, 7, 63, 2, 2, 590, 554, 3,
	2, 2, 2, 590, 562, 3, 2, 2, 2, 590, 563, 3, 2, 2, 2, 590, 565, 3, 2, 2,
	2, 590, 567, 3, 2, 2, 2, 590, 568, 3, 2, 2, 2, 590, 570, 3, 2, 2, 2, 590,
	572, 3, 2, 2, 2, 590, 574, 3, 2, 2, 2, 590, 576, 3, 2, 2, 2, 590, 577,
	3, 2, 2, 2, 590, 579, 3, 2, 2, 2, 590, 580, 3, 2, 2, 2, 590, 582, 3, 2,
	2, 2, 590, 585, 3, 2, 2, 2, 590, 587, 3, 2, 2, 2, 591, 176, 3, 2, 2, 2,
	592, 593, 9, 3, 2, 2, 593, 178, 3, 2, 2, 2, 594, 595, 4, 50, 59, 2, 595,
	180, 3, 2, 2, 2, 596, 600, 5, 179, 90, 2, 597, 599, 5, 179, 90, 2, 598,
	597, 3, 2, 2, 2, 599, 602, 3, 2, 2, 2, 600, 598, 3, 2, 2, 2, 600, 601,
	3, 2, 2, 2, 601, 603, 3, 2, 2, 2, 602, 600, 3, 2, 2, 2, 603, 607, 7, 48,
	2, 2, 604, 606, 5, 179, 90, 2, 605, 604, 3, 2, 2, 2, 606, 609, 3, 2, 2,
	2, 607, 605, 3, 2, 2, 2, 607, 608, 3, 2, 2, 2, 608, 182, 3, 2, 2, 2, 609,
	607, 3, 2, 2, 2, 610, 614, 5, 179, 90, 2, 611, 613, 5, 179, 90, 2, 612,
	611, 3, 2, 2, 2, 613, 616, 3, 2, 2, 2, 614, 612, 3, 2, 2, 2, 614, 615,
	3, 2, 2, 2, 615, 627, 3, 2, 2, 2, 616, 614, 3, 2, 2, 2, 617, 618, 7, 50,
	2, 2, 618, 619, 7, 122, 2, 2, 619, 623, 3, 2, 2, 2, 620, 622, 5, 179, 90,
	2, 621, 620, 3, 2, 2, 2, 622, 625, 3, 2, 2, 2, 623, 621, 3, 2, 2, 2, 623,
	624, 3, 2, 2, 2, 624, 627, 3, 2, 2, 2, 625, 623, 3, 2, 2, 2, 626, 610,
	3, 2, 2, 2, 626, 617, 3, 2, 2, 2, 627, 184, 3, 2, 2, 2, 628, 632, 7, 36,
	2, 2, 629, 631, 10, 4, 2, 2, 630, 629, 3, 2, 2, 2, 631, 634, 3, 2, 2, 2,
	632, 630, 3, 2, 2, 2, 632, 633, 3, 2, 2, 2, 633, 635, 3, 2, 2, 2, 634,
	632, 3, 2, 2, 2, 635, 636, 7, 36, 2, 2, 636, 186, 3, 2, 2, 2, 637, 638,
	9, 5, 2, 2, 638, 639, 3, 2, 2, 2, 639, 640, 8, 94, 2, 2, 640, 188, 3, 2,
	2, 2, 12, 2, 557, 559, 590, 600, 607, 614, 623, 626, 632, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'import'", "';'", "'inherit'", "':'", "'('", "')'", "','", "'*'",
	"'='", "'constant'", "'class'", "'{'", "'}'", "'static'", "'private'",
	"'nomask'", "'public'", "'protected'", "'inline'", "'if'", "'else'", "'while'",
	"'do'", "'for'", "'switch'", "'case'", "'..'", "'default'", "'foreach'",
	"'break'", "'continue'", "'+='", "'*='", "'/='", "'&='", "'|='", "'^='",
	"'<<='", "'>>='", "'%='", "'?'", "'||'", "'&&'", "'|'", "'^'", "'&'", "'=='",
	"'!='", "'>'", "'<'", "'>='", "'<='", "'<<'", "'>>'", "'+'", "'/'", "'%'",
	"'--'", "'++'", "'~'", "'-'", "'->'", "'['", "']'", "'catch'", "'gauge'",
	"'sscanf'", "'lambda'", "'.'", "'({'", "'})'", "'(<'", "'>)'", "'(['",
	"'])'", "'@'", "'...'", "'int'", "'string'", "'float'", "'program'", "'object'",
	"'mapping'", "'array'", "'multiset'", "'function'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "IDENTIFIER",
	"LETTER", "DIGIT", "FLOAT", "NUMBER", "STRING", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "T__26", "T__27", "T__28", "T__29", "T__30", "T__31", "T__32",
	"T__33", "T__34", "T__35", "T__36", "T__37", "T__38", "T__39", "T__40",
	"T__41", "T__42", "T__43", "T__44", "T__45", "T__46", "T__47", "T__48",
	"T__49", "T__50", "T__51", "T__52", "T__53", "T__54", "T__55", "T__56",
	"T__57", "T__58", "T__59", "T__60", "T__61", "T__62", "T__63", "T__64",
	"T__65", "T__66", "T__67", "T__68", "T__69", "T__70", "T__71", "T__72",
	"T__73", "T__74", "T__75", "T__76", "T__77", "T__78", "T__79", "T__80",
	"T__81", "T__82", "T__83", "T__84", "T__85", "IDENTIFIER", "LETTER", "DIGIT",
	"FLOAT", "NUMBER", "STRING", "WS",
}

type pikeLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewpikeLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *pikeLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewpikeLexer(input antlr.CharStream) *pikeLexer {
	l := new(pikeLexer)
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
	l.GrammarFileName = "pike.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// pikeLexer tokens.
const (
	pikeLexerT__0       = 1
	pikeLexerT__1       = 2
	pikeLexerT__2       = 3
	pikeLexerT__3       = 4
	pikeLexerT__4       = 5
	pikeLexerT__5       = 6
	pikeLexerT__6       = 7
	pikeLexerT__7       = 8
	pikeLexerT__8       = 9
	pikeLexerT__9       = 10
	pikeLexerT__10      = 11
	pikeLexerT__11      = 12
	pikeLexerT__12      = 13
	pikeLexerT__13      = 14
	pikeLexerT__14      = 15
	pikeLexerT__15      = 16
	pikeLexerT__16      = 17
	pikeLexerT__17      = 18
	pikeLexerT__18      = 19
	pikeLexerT__19      = 20
	pikeLexerT__20      = 21
	pikeLexerT__21      = 22
	pikeLexerT__22      = 23
	pikeLexerT__23      = 24
	pikeLexerT__24      = 25
	pikeLexerT__25      = 26
	pikeLexerT__26      = 27
	pikeLexerT__27      = 28
	pikeLexerT__28      = 29
	pikeLexerT__29      = 30
	pikeLexerT__30      = 31
	pikeLexerT__31      = 32
	pikeLexerT__32      = 33
	pikeLexerT__33      = 34
	pikeLexerT__34      = 35
	pikeLexerT__35      = 36
	pikeLexerT__36      = 37
	pikeLexerT__37      = 38
	pikeLexerT__38      = 39
	pikeLexerT__39      = 40
	pikeLexerT__40      = 41
	pikeLexerT__41      = 42
	pikeLexerT__42      = 43
	pikeLexerT__43      = 44
	pikeLexerT__44      = 45
	pikeLexerT__45      = 46
	pikeLexerT__46      = 47
	pikeLexerT__47      = 48
	pikeLexerT__48      = 49
	pikeLexerT__49      = 50
	pikeLexerT__50      = 51
	pikeLexerT__51      = 52
	pikeLexerT__52      = 53
	pikeLexerT__53      = 54
	pikeLexerT__54      = 55
	pikeLexerT__55      = 56
	pikeLexerT__56      = 57
	pikeLexerT__57      = 58
	pikeLexerT__58      = 59
	pikeLexerT__59      = 60
	pikeLexerT__60      = 61
	pikeLexerT__61      = 62
	pikeLexerT__62      = 63
	pikeLexerT__63      = 64
	pikeLexerT__64      = 65
	pikeLexerT__65      = 66
	pikeLexerT__66      = 67
	pikeLexerT__67      = 68
	pikeLexerT__68      = 69
	pikeLexerT__69      = 70
	pikeLexerT__70      = 71
	pikeLexerT__71      = 72
	pikeLexerT__72      = 73
	pikeLexerT__73      = 74
	pikeLexerT__74      = 75
	pikeLexerT__75      = 76
	pikeLexerT__76      = 77
	pikeLexerT__77      = 78
	pikeLexerT__78      = 79
	pikeLexerT__79      = 80
	pikeLexerT__80      = 81
	pikeLexerT__81      = 82
	pikeLexerT__82      = 83
	pikeLexerT__83      = 84
	pikeLexerT__84      = 85
	pikeLexerT__85      = 86
	pikeLexerIDENTIFIER = 87
	pikeLexerLETTER     = 88
	pikeLexerDIGIT      = 89
	pikeLexerFLOAT      = 90
	pikeLexerNUMBER     = 91
	pikeLexerSTRING     = 92
	pikeLexerWS         = 93
)
