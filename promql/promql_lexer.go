// Code generated from PromQLLexer.g4 by ANTLR 4.9.3. DO NOT EDIT.

package promql

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 44, 753,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 3, 2,
	6, 2, 89, 10, 2, 13, 2, 14, 2, 90, 3, 2, 3, 2, 6, 2, 95, 10, 2, 13, 2,
	14, 2, 96, 5, 2, 99, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 105, 10, 3, 12,
	3, 14, 3, 108, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 115, 10, 3, 12,
	3, 14, 3, 118, 11, 3, 3, 3, 5, 3, 121, 10, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5,
	30, 296, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 31, 5, 31, 696, 10, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34,
	3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 720, 10, 39, 3, 39, 3, 39,
	5, 39, 724, 10, 39, 3, 40, 6, 40, 727, 10, 40, 13, 40, 14, 40, 728, 3,
	40, 3, 40, 3, 41, 3, 41, 7, 41, 735, 10, 41, 12, 41, 14, 41, 738, 11, 41,
	3, 42, 3, 42, 7, 42, 742, 10, 42, 12, 42, 14, 42, 745, 11, 42, 3, 43, 6,
	43, 748, 10, 43, 13, 43, 14, 43, 749, 3, 43, 3, 43, 2, 2, 44, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43,
	23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61,
	32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79,
	41, 81, 42, 83, 43, 85, 44, 3, 2, 11, 3, 2, 50, 59, 4, 2, 41, 41, 94, 94,
	4, 2, 36, 36, 94, 94, 8, 2, 102, 102, 106, 106, 111, 111, 117, 117, 121,
	121, 123, 123, 6, 2, 60, 60, 67, 92, 97, 97, 99, 124, 6, 2, 50, 60, 67,
	92, 97, 97, 99, 124, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92,
	97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 823, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43,
	3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2,
	51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2,
	2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2,
	2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2,
	2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3,
	2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 3, 88, 3, 2, 2, 2, 5, 120,
	3, 2, 2, 2, 7, 122, 3, 2, 2, 2, 9, 124, 3, 2, 2, 2, 11, 126, 3, 2, 2, 2,
	13, 128, 3, 2, 2, 2, 15, 130, 3, 2, 2, 2, 17, 132, 3, 2, 2, 2, 19, 134,
	3, 2, 2, 2, 21, 138, 3, 2, 2, 2, 23, 141, 3, 2, 2, 2, 25, 148, 3, 2, 2,
	2, 27, 150, 3, 2, 2, 2, 29, 153, 3, 2, 2, 2, 31, 156, 3, 2, 2, 2, 33, 158,
	3, 2, 2, 2, 35, 160, 3, 2, 2, 2, 37, 163, 3, 2, 2, 2, 39, 166, 3, 2, 2,
	2, 41, 169, 3, 2, 2, 2, 43, 172, 3, 2, 2, 2, 45, 175, 3, 2, 2, 2, 47, 183,
	3, 2, 2, 2, 49, 186, 3, 2, 2, 2, 51, 195, 3, 2, 2, 2, 53, 206, 3, 2, 2,
	2, 55, 218, 3, 2, 2, 2, 57, 225, 3, 2, 2, 2, 59, 295, 3, 2, 2, 2, 61, 695,
	3, 2, 2, 2, 63, 697, 3, 2, 2, 2, 65, 699, 3, 2, 2, 2, 67, 701, 3, 2, 2,
	2, 69, 703, 3, 2, 2, 2, 71, 705, 3, 2, 2, 2, 73, 707, 3, 2, 2, 2, 75, 709,
	3, 2, 2, 2, 77, 723, 3, 2, 2, 2, 79, 726, 3, 2, 2, 2, 81, 732, 3, 2, 2,
	2, 83, 739, 3, 2, 2, 2, 85, 747, 3, 2, 2, 2, 87, 89, 9, 2, 2, 2, 88, 87,
	3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2,
	91, 98, 3, 2, 2, 2, 92, 94, 7, 48, 2, 2, 93, 95, 9, 2, 2, 2, 94, 93, 3,
	2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97,
	99, 3, 2, 2, 2, 98, 92, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 4, 3, 2, 2,
	2, 100, 106, 7, 41, 2, 2, 101, 105, 10, 3, 2, 2, 102, 103, 7, 94, 2, 2,
	103, 105, 11, 2, 2, 2, 104, 101, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105,
	108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 109,
	3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 121, 7, 41, 2, 2, 110, 116, 7, 36,
	2, 2, 111, 115, 10, 4, 2, 2, 112, 113, 7, 94, 2, 2, 113, 115, 11, 2, 2,
	2, 114, 111, 3, 2, 2, 2, 114, 112, 3, 2, 2, 2, 115, 118, 3, 2, 2, 2, 116,
	114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 119, 3, 2, 2, 2, 118, 116,
	3, 2, 2, 2, 119, 121, 7, 36, 2, 2, 120, 100, 3, 2, 2, 2, 120, 110, 3, 2,
	2, 2, 121, 6, 3, 2, 2, 2, 122, 123, 7, 45, 2, 2, 123, 8, 3, 2, 2, 2, 124,
	125, 7, 47, 2, 2, 125, 10, 3, 2, 2, 2, 126, 127, 7, 44, 2, 2, 127, 12,
	3, 2, 2, 2, 128, 129, 7, 49, 2, 2, 129, 14, 3, 2, 2, 2, 130, 131, 7, 39,
	2, 2, 131, 16, 3, 2, 2, 2, 132, 133, 7, 96, 2, 2, 133, 18, 3, 2, 2, 2,
	134, 135, 7, 99, 2, 2, 135, 136, 7, 112, 2, 2, 136, 137, 7, 102, 2, 2,
	137, 20, 3, 2, 2, 2, 138, 139, 7, 113, 2, 2, 139, 140, 7, 116, 2, 2, 140,
	22, 3, 2, 2, 2, 141, 142, 7, 119, 2, 2, 142, 143, 7, 112, 2, 2, 143, 144,
	7, 110, 2, 2, 144, 145, 7, 103, 2, 2, 145, 146, 7, 117, 2, 2, 146, 147,
	7, 117, 2, 2, 147, 24, 3, 2, 2, 2, 148, 149, 7, 63, 2, 2, 149, 26, 3, 2,
	2, 2, 150, 151, 7, 63, 2, 2, 151, 152, 7, 63, 2, 2, 152, 28, 3, 2, 2, 2,
	153, 154, 7, 35, 2, 2, 154, 155, 7, 63, 2, 2, 155, 30, 3, 2, 2, 2, 156,
	157, 7, 64, 2, 2, 157, 32, 3, 2, 2, 2, 158, 159, 7, 62, 2, 2, 159, 34,
	3, 2, 2, 2, 160, 161, 7, 64, 2, 2, 161, 162, 7, 63, 2, 2, 162, 36, 3, 2,
	2, 2, 163, 164, 7, 62, 2, 2, 164, 165, 7, 63, 2, 2, 165, 38, 3, 2, 2, 2,
	166, 167, 7, 63, 2, 2, 167, 168, 7, 128, 2, 2, 168, 40, 3, 2, 2, 2, 169,
	170, 7, 35, 2, 2, 170, 171, 7, 128, 2, 2, 171, 42, 3, 2, 2, 2, 172, 173,
	7, 100, 2, 2, 173, 174, 7, 123, 2, 2, 174, 44, 3, 2, 2, 2, 175, 176, 7,
	121, 2, 2, 176, 177, 7, 107, 2, 2, 177, 178, 7, 118, 2, 2, 178, 179, 7,
	106, 2, 2, 179, 180, 7, 113, 2, 2, 180, 181, 7, 119, 2, 2, 181, 182, 7,
	118, 2, 2, 182, 46, 3, 2, 2, 2, 183, 184, 7, 113, 2, 2, 184, 185, 7, 112,
	2, 2, 185, 48, 3, 2, 2, 2, 186, 187, 7, 107, 2, 2, 187, 188, 7, 105, 2,
	2, 188, 189, 7, 112, 2, 2, 189, 190, 7, 113, 2, 2, 190, 191, 7, 116, 2,
	2, 191, 192, 7, 107, 2, 2, 192, 193, 7, 112, 2, 2, 193, 194, 7, 105, 2,
	2, 194, 50, 3, 2, 2, 2, 195, 196, 7, 105, 2, 2, 196, 197, 7, 116, 2, 2,
	197, 198, 7, 113, 2, 2, 198, 199, 7, 119, 2, 2, 199, 200, 7, 114, 2, 2,
	200, 201, 7, 97, 2, 2, 201, 202, 7, 110, 2, 2, 202, 203, 7, 103, 2, 2,
	203, 204, 7, 104, 2, 2, 204, 205, 7, 118, 2, 2, 205, 52, 3, 2, 2, 2, 206,
	207, 7, 105, 2, 2, 207, 208, 7, 116, 2, 2, 208, 209, 7, 113, 2, 2, 209,
	210, 7, 119, 2, 2, 210, 211, 7, 114, 2, 2, 211, 212, 7, 97, 2, 2, 212,
	213, 7, 116, 2, 2, 213, 214, 7, 107, 2, 2, 214, 215, 7, 105, 2, 2, 215,
	216, 7, 106, 2, 2, 216, 217, 7, 118, 2, 2, 217, 54, 3, 2, 2, 2, 218, 219,
	7, 113, 2, 2, 219, 220, 7, 104, 2, 2, 220, 221, 7, 104, 2, 2, 221, 222,
	7, 117, 2, 2, 222, 223, 7, 103, 2, 2, 223, 224, 7, 118, 2, 2, 224, 56,
	3, 2, 2, 2, 225, 226, 7, 100, 2, 2, 226, 227, 7, 113, 2, 2, 227, 228, 7,
	113, 2, 2, 228, 229, 7, 110, 2, 2, 229, 58, 3, 2, 2, 2, 230, 231, 7, 117,
	2, 2, 231, 232, 7, 119, 2, 2, 232, 296, 7, 111, 2, 2, 233, 234, 7, 111,
	2, 2, 234, 235, 7, 107, 2, 2, 235, 296, 7, 112, 2, 2, 236, 237, 7, 111,
	2, 2, 237, 238, 7, 99, 2, 2, 238, 296, 7, 122, 2, 2, 239, 240, 7, 99, 2,
	2, 240, 241, 7, 120, 2, 2, 241, 296, 7, 105, 2, 2, 242, 243, 7, 105, 2,
	2, 243, 244, 7, 116, 2, 2, 244, 245, 7, 113, 2, 2, 245, 246, 7, 119, 2,
	2, 246, 296, 7, 114, 2, 2, 247, 248, 7, 117, 2, 2, 248, 249, 7, 118, 2,
	2, 249, 250, 7, 102, 2, 2, 250, 251, 7, 102, 2, 2, 251, 252, 7, 103, 2,
	2, 252, 296, 7, 120, 2, 2, 253, 254, 7, 117, 2, 2, 254, 255, 7, 118, 2,
	2, 255, 256, 7, 102, 2, 2, 256, 257, 7, 120, 2, 2, 257, 258, 7, 99, 2,
	2, 258, 296, 7, 116, 2, 2, 259, 260, 7, 101, 2, 2, 260, 261, 7, 113, 2,
	2, 261, 262, 7, 119, 2, 2, 262, 263, 7, 112, 2, 2, 263, 296, 7, 118, 2,
	2, 264, 265, 7, 101, 2, 2, 265, 266, 7, 113, 2, 2, 266, 267, 7, 119, 2,
	2, 267, 268, 7, 112, 2, 2, 268, 269, 7, 118, 2, 2, 269, 270, 7, 97, 2,
	2, 270, 271, 7, 120, 2, 2, 271, 272, 7, 99, 2, 2, 272, 273, 7, 110, 2,
	2, 273, 274, 7, 119, 2, 2, 274, 275, 7, 103, 2, 2, 275, 296, 7, 117, 2,
	2, 276, 277, 7, 100, 2, 2, 277, 278, 7, 113, 2, 2, 278, 279, 7, 118, 2,
	2, 279, 280, 7, 118, 2, 2, 280, 281, 7, 113, 2, 2, 281, 282, 7, 111, 2,
	2, 282, 296, 7, 109, 2, 2, 283, 284, 7, 118, 2, 2, 284, 285, 7, 113, 2,
	2, 285, 286, 7, 114, 2, 2, 286, 296, 7, 109, 2, 2, 287, 288, 7, 115, 2,
	2, 288, 289, 7, 119, 2, 2, 289, 290, 7, 99, 2, 2, 290, 291, 7, 112, 2,
	2, 291, 292, 7, 118, 2, 2, 292, 293, 7, 107, 2, 2, 293, 294, 7, 110, 2,
	2, 294, 296, 7, 103, 2, 2, 295, 230, 3, 2, 2, 2, 295, 233, 3, 2, 2, 2,
	295, 236, 3, 2, 2, 2, 295, 239, 3, 2, 2, 2, 295, 242, 3, 2, 2, 2, 295,
	247, 3, 2, 2, 2, 295, 253, 3, 2, 2, 2, 295, 259, 3, 2, 2, 2, 295, 264,
	3, 2, 2, 2, 295, 276, 3, 2, 2, 2, 295, 283, 3, 2, 2, 2, 295, 287, 3, 2,
	2, 2, 296, 60, 3, 2, 2, 2, 297, 298, 7, 99, 2, 2, 298, 299, 7, 100, 2,
	2, 299, 696, 7, 117, 2, 2, 300, 301, 7, 99, 2, 2, 301, 302, 7, 100, 2,
	2, 302, 303, 7, 117, 2, 2, 303, 304, 7, 103, 2, 2, 304, 305, 7, 112, 2,
	2, 305, 696, 7, 118, 2, 2, 306, 307, 7, 99, 2, 2, 307, 308, 7, 100, 2,
	2, 308, 309, 7, 117, 2, 2, 309, 310, 7, 103, 2, 2, 310, 311, 7, 112, 2,
	2, 311, 312, 7, 118, 2, 2, 312, 313, 7, 97, 2, 2, 313, 314, 7, 113, 2,
	2, 314, 315, 7, 120, 2, 2, 315, 316, 7, 103, 2, 2, 316, 317, 7, 116, 2,
	2, 317, 318, 7, 97, 2, 2, 318, 319, 7, 118, 2, 2, 319, 320, 7, 107, 2,
	2, 320, 321, 7, 111, 2, 2, 321, 696, 7, 103, 2, 2, 322, 323, 7, 101, 2,
	2, 323, 324, 7, 103, 2, 2, 324, 325, 7, 107, 2, 2, 325, 696, 7, 110, 2,
	2, 326, 327, 7, 101, 2, 2, 327, 328, 7, 106, 2, 2, 328, 329, 7, 99, 2,
	2, 329, 330, 7, 112, 2, 2, 330, 331, 7, 105, 2, 2, 331, 332, 7, 103, 2,
	2, 332, 696, 7, 117, 2, 2, 333, 334, 7, 101, 2, 2, 334, 335, 7, 110, 2,
	2, 335, 336, 7, 99, 2, 2, 336, 337, 7, 111, 2, 2, 337, 338, 7, 114, 2,
	2, 338, 339, 7, 97, 2, 2, 339, 340, 7, 111, 2, 2, 340, 341, 7, 99, 2, 2,
	341, 696, 7, 122, 2, 2, 342, 343, 7, 101, 2, 2, 343, 344, 7, 110, 2, 2,
	344, 345, 7, 99, 2, 2, 345, 346, 7, 111, 2, 2, 346, 347, 7, 114, 2, 2,
	347, 348, 7, 97, 2, 2, 348, 349, 7, 111, 2, 2, 349, 350, 7, 107, 2, 2,
	350, 696, 7, 112, 2, 2, 351, 352, 7, 102, 2, 2, 352, 353, 7, 99, 2, 2,
	353, 354, 7, 123, 2, 2, 354, 355, 7, 97, 2, 2, 355, 356, 7, 113, 2, 2,
	356, 357, 7, 104, 2, 2, 357, 358, 7, 97, 2, 2, 358, 359, 7, 111, 2, 2,
	359, 360, 7, 113, 2, 2, 360, 361, 7, 112, 2, 2, 361, 362, 7, 118, 2, 2,
	362, 696, 7, 106, 2, 2, 363, 364, 7, 102, 2, 2, 364, 365, 7, 99, 2, 2,
	365, 366, 7, 123, 2, 2, 366, 367, 7, 97, 2, 2, 367, 368, 7, 113, 2, 2,
	368, 369, 7, 104, 2, 2, 369, 370, 7, 97, 2, 2, 370, 371, 7, 121, 2, 2,
	371, 372, 7, 103, 2, 2, 372, 373, 7, 103, 2, 2, 373, 696, 7, 109, 2, 2,
	374, 375, 7, 102, 2, 2, 375, 376, 7, 99, 2, 2, 376, 377, 7, 123, 2, 2,
	377, 378, 7, 117, 2, 2, 378, 379, 7, 97, 2, 2, 379, 380, 7, 107, 2, 2,
	380, 381, 7, 112, 2, 2, 381, 382, 7, 97, 2, 2, 382, 383, 7, 111, 2, 2,
	383, 384, 7, 113, 2, 2, 384, 385, 7, 112, 2, 2, 385, 386, 7, 118, 2, 2,
	386, 696, 7, 106, 2, 2, 387, 388, 7, 102, 2, 2, 388, 389, 7, 103, 2, 2,
	389, 390, 7, 110, 2, 2, 390, 391, 7, 118, 2, 2, 391, 696, 7, 99, 2, 2,
	392, 393, 7, 102, 2, 2, 393, 394, 7, 103, 2, 2, 394, 395, 7, 116, 2, 2,
	395, 396, 7, 107, 2, 2, 396, 696, 7, 120, 2, 2, 397, 398, 7, 103, 2, 2,
	398, 399, 7, 122, 2, 2, 399, 696, 7, 114, 2, 2, 400, 401, 7, 104, 2, 2,
	401, 402, 7, 110, 2, 2, 402, 403, 7, 113, 2, 2, 403, 404, 7, 113, 2, 2,
	404, 696, 7, 116, 2, 2, 405, 406, 7, 106, 2, 2, 406, 407, 7, 107, 2, 2,
	407, 408, 7, 117, 2, 2, 408, 409, 7, 118, 2, 2, 409, 410, 7, 113, 2, 2,
	410, 411, 7, 105, 2, 2, 411, 412, 7, 116, 2, 2, 412, 413, 7, 99, 2, 2,
	413, 414, 7, 111, 2, 2, 414, 415, 7, 97, 2, 2, 415, 416, 7, 115, 2, 2,
	416, 417, 7, 119, 2, 2, 417, 418, 7, 99, 2, 2, 418, 419, 7, 112, 2, 2,
	419, 420, 7, 118, 2, 2, 420, 421, 7, 107, 2, 2, 421, 422, 7, 110, 2, 2,
	422, 696, 7, 103, 2, 2, 423, 424, 7, 106, 2, 2, 424, 425, 7, 113, 2, 2,
	425, 426, 7, 110, 2, 2, 426, 427, 7, 118, 2, 2, 427, 428, 7, 97, 2, 2,
	428, 429, 7, 121, 2, 2, 429, 430, 7, 107, 2, 2, 430, 431, 7, 112, 2, 2,
	431, 432, 7, 118, 2, 2, 432, 433, 7, 103, 2, 2, 433, 434, 7, 116, 2, 2,
	434, 696, 7, 117, 2, 2, 435, 436, 7, 106, 2, 2, 436, 437, 7, 113, 2, 2,
	437, 438, 7, 119, 2, 2, 438, 696, 7, 116, 2, 2, 439, 440, 7, 107, 2, 2,
	440, 441, 7, 102, 2, 2, 441, 442, 7, 103, 2, 2, 442, 443, 7, 110, 2, 2,
	443, 444, 7, 118, 2, 2, 444, 696, 7, 99, 2, 2, 445, 446, 7, 107, 2, 2,
	446, 447, 7, 112, 2, 2, 447, 448, 7, 101, 2, 2, 448, 449, 7, 116, 2, 2,
	449, 450, 7, 103, 2, 2, 450, 451, 7, 99, 2, 2, 451, 452, 7, 117, 2, 2,
	452, 696, 7, 103, 2, 2, 453, 454, 7, 107, 2, 2, 454, 455, 7, 116, 2, 2,
	455, 456, 7, 99, 2, 2, 456, 457, 7, 118, 2, 2, 457, 696, 7, 103, 2, 2,
	458, 459, 7, 110, 2, 2, 459, 460, 7, 99, 2, 2, 460, 461, 7, 100, 2, 2,
	461, 462, 7, 103, 2, 2, 462, 463, 7, 110, 2, 2, 463, 464, 7, 97, 2, 2,
	464, 465, 7, 108, 2, 2, 465, 466, 7, 113, 2, 2, 466, 467, 7, 107, 2, 2,
	467, 696, 7, 112, 2, 2, 468, 469, 7, 110, 2, 2, 469, 470, 7, 99, 2, 2,
	470, 471, 7, 100, 2, 2, 471, 472, 7, 103, 2, 2, 472, 473, 7, 110, 2, 2,
	473, 474, 7, 97, 2, 2, 474, 475, 7, 116, 2, 2, 475, 476, 7, 103, 2, 2,
	476, 477, 7, 114, 2, 2, 477, 478, 7, 110, 2, 2, 478, 479, 7, 99, 2, 2,
	479, 480, 7, 101, 2, 2, 480, 696, 7, 103, 2, 2, 481, 482, 7, 110, 2, 2,
	482, 696, 7, 112, 2, 2, 483, 484, 7, 110, 2, 2, 484, 485, 7, 113, 2, 2,
	485, 486, 7, 105, 2, 2, 486, 696, 7, 52, 2, 2, 487, 488, 7, 110, 2, 2,
	488, 489, 7, 113, 2, 2, 489, 490, 7, 105, 2, 2, 490, 491, 7, 51, 2, 2,
	491, 696, 7, 50, 2, 2, 492, 493, 7, 111, 2, 2, 493, 494, 7, 107, 2, 2,
	494, 495, 7, 112, 2, 2, 495, 496, 7, 119, 2, 2, 496, 497, 7, 118, 2, 2,
	497, 696, 7, 103, 2, 2, 498, 499, 7, 111, 2, 2, 499, 500, 7, 113, 2, 2,
	500, 501, 7, 112, 2, 2, 501, 502, 7, 118, 2, 2, 502, 696, 7, 106, 2, 2,
	503, 504, 7, 114, 2, 2, 504, 505, 7, 116, 2, 2, 505, 506, 7, 103, 2, 2,
	506, 507, 7, 102, 2, 2, 507, 508, 7, 107, 2, 2, 508, 509, 7, 101, 2, 2,
	509, 510, 7, 118, 2, 2, 510, 511, 7, 97, 2, 2, 511, 512, 7, 110, 2, 2,
	512, 513, 7, 107, 2, 2, 513, 514, 7, 112, 2, 2, 514, 515, 7, 103, 2, 2,
	515, 516, 7, 99, 2, 2, 516, 696, 7, 116, 2, 2, 517, 518, 7, 116, 2, 2,
	518, 519, 7, 99, 2, 2, 519, 520, 7, 118, 2, 2, 520, 696, 7, 103, 2, 2,
	521, 522, 7, 116, 2, 2, 522, 523, 7, 103, 2, 2, 523, 524, 7, 117, 2, 2,
	524, 525, 7, 103, 2, 2, 525, 526, 7, 118, 2, 2, 526, 696, 7, 117, 2, 2,
	527, 528, 7, 116, 2, 2, 528, 529, 7, 113, 2, 2, 529, 530, 7, 119, 2, 2,
	530, 531, 7, 112, 2, 2, 531, 696, 7, 102, 2, 2, 532, 533, 7, 117, 2, 2,
	533, 534, 7, 101, 2, 2, 534, 535, 7, 99, 2, 2, 535, 536, 7, 110, 2, 2,
	536, 537, 7, 99, 2, 2, 537, 696, 7, 116, 2, 2, 538, 539, 7, 117, 2, 2,
	539, 540, 7, 113, 2, 2, 540, 541, 7, 116, 2, 2, 541, 696, 7, 118, 2, 2,
	542, 543, 7, 117, 2, 2, 543, 544, 7, 113, 2, 2, 544, 545, 7, 116, 2, 2,
	545, 546, 7, 118, 2, 2, 546, 547, 7, 97, 2, 2, 547, 548, 7, 102, 2, 2,
	548, 549, 7, 103, 2, 2, 549, 550, 7, 117, 2, 2, 550, 696, 7, 101, 2, 2,
	551, 552, 7, 117, 2, 2, 552, 553, 7, 115, 2, 2, 553, 554, 7, 116, 2, 2,
	554, 696, 7, 118, 2, 2, 555, 556, 7, 118, 2, 2, 556, 557, 7, 107, 2, 2,
	557, 558, 7, 111, 2, 2, 558, 696, 7, 103, 2, 2, 559, 560, 7, 118, 2, 2,
	560, 561, 7, 107, 2, 2, 561, 562, 7, 111, 2, 2, 562, 563, 7, 103, 2, 2,
	563, 564, 7, 117, 2, 2, 564, 565, 7, 118, 2, 2, 565, 566, 7, 99, 2, 2,
	566, 567, 7, 111, 2, 2, 567, 696, 7, 114, 2, 2, 568, 569, 7, 120, 2, 2,
	569, 570, 7, 103, 2, 2, 570, 571, 7, 101, 2, 2, 571, 572, 7, 118, 2, 2,
	572, 573, 7, 113, 2, 2, 573, 696, 7, 116, 2, 2, 574, 575, 7, 123, 2, 2,
	575, 576, 7, 103, 2, 2, 576, 577, 7, 99, 2, 2, 577, 696, 7, 116, 2, 2,
	578, 579, 7, 99, 2, 2, 579, 580, 7, 120, 2, 2, 580, 581, 7, 105, 2, 2,
	581, 582, 7, 97, 2, 2, 582, 583, 7, 113, 2, 2, 583, 584, 7, 120, 2, 2,
	584, 585, 7, 103, 2, 2, 585, 586, 7, 116, 2, 2, 586, 587, 7, 97, 2, 2,
	587, 588, 7, 118, 2, 2, 588, 589, 7, 107, 2, 2, 589, 590, 7, 111, 2, 2,
	590, 696, 7, 103, 2, 2, 591, 592, 7, 111, 2, 2, 592, 593, 7, 107, 2, 2,
	593, 594, 7, 112, 2, 2, 594, 595, 7, 97, 2, 2, 595, 596, 7, 113, 2, 2,
	596, 597, 7, 120, 2, 2, 597, 598, 7, 103, 2, 2, 598, 599, 7, 116, 2, 2,
	599, 600, 7, 97, 2, 2, 600, 601, 7, 118, 2, 2, 601, 602, 7, 107, 2, 2,
	602, 603, 7, 111, 2, 2, 603, 696, 7, 103, 2, 2, 604, 605, 7, 111, 2, 2,
	605, 606, 7, 99, 2, 2, 606, 607, 7, 122, 2, 2, 607, 608, 7, 97, 2, 2, 608,
	609, 7, 113, 2, 2, 609, 610, 7, 120, 2, 2, 610, 611, 7, 103, 2, 2, 611,
	612, 7, 116, 2, 2, 612, 613, 7, 97, 2, 2, 613, 614, 7, 118, 2, 2, 614,
	615, 7, 107, 2, 2, 615, 616, 7, 111, 2, 2, 616, 696, 7, 103, 2, 2, 617,
	618, 7, 117, 2, 2, 618, 619, 7, 119, 2, 2, 619, 620, 7, 111, 2, 2, 620,
	621, 7, 97, 2, 2, 621, 622, 7, 113, 2, 2, 622, 623, 7, 120, 2, 2, 623,
	624, 7, 103, 2, 2, 624, 625, 7, 116, 2, 2, 625, 626, 7, 97, 2, 2, 626,
	627, 7, 118, 2, 2, 627, 628, 7, 107, 2, 2, 628, 629, 7, 111, 2, 2, 629,
	696, 7, 103, 2, 2, 630, 631, 7, 101, 2, 2, 631, 632, 7, 113, 2, 2, 632,
	633, 7, 119, 2, 2, 633, 634, 7, 112, 2, 2, 634, 635, 7, 118, 2, 2, 635,
	636, 7, 97, 2, 2, 636, 637, 7, 113, 2, 2, 637, 638, 7, 120, 2, 2, 638,
	639, 7, 103, 2, 2, 639, 640, 7, 116, 2, 2, 640, 641, 7, 97, 2, 2, 641,
	642, 7, 118, 2, 2, 642, 643, 7, 107, 2, 2, 643, 644, 7, 111, 2, 2, 644,
	696, 7, 103, 2, 2, 645, 646, 7, 115, 2, 2, 646, 647, 7, 119, 2, 2, 647,
	648, 7, 99, 2, 2, 648, 649, 7, 112, 2, 2, 649, 650, 7, 118, 2, 2, 650,
	651, 7, 107, 2, 2, 651, 652, 7, 110, 2, 2, 652, 653, 7, 103, 2, 2, 653,
	654, 7, 97, 2, 2, 654, 655, 7, 113, 2, 2, 655, 656, 7, 120, 2, 2, 656,
	657, 7, 103, 2, 2, 657, 658, 7, 116, 2, 2, 658, 659, 7, 97, 2, 2, 659,
	660, 7, 118, 2, 2, 660, 661, 7, 107, 2, 2, 661, 662, 7, 111, 2, 2, 662,
	696, 7, 103, 2, 2, 663, 664, 7, 117, 2, 2, 664, 665, 7, 118, 2, 2, 665,
	666, 7, 102, 2, 2, 666, 667, 7, 102, 2, 2, 667, 668, 7, 103, 2, 2, 668,
	669, 7, 120, 2, 2, 669, 670, 7, 97, 2, 2, 670, 671, 7, 113, 2, 2, 671,
	672, 7, 120, 2, 2, 672, 673, 7, 103, 2, 2, 673, 674, 7, 116, 2, 2, 674,
	675, 7, 97, 2, 2, 675, 676, 7, 118, 2, 2, 676, 677, 7, 107, 2, 2, 677,
	678, 7, 111, 2, 2, 678, 696, 7, 103, 2, 2, 679, 680, 7, 117, 2, 2, 680,
	681, 7, 118, 2, 2, 681, 682, 7, 102, 2, 2, 682, 683, 7, 120, 2, 2, 683,
	684, 7, 99, 2, 2, 684, 685, 7, 116, 2, 2, 685, 686, 7, 97, 2, 2, 686, 687,
	7, 113, 2, 2, 687, 688, 7, 120, 2, 2, 688, 689, 7, 103, 2, 2, 689, 690,
	7, 116, 2, 2, 690, 691, 7, 97, 2, 2, 691, 692, 7, 118, 2, 2, 692, 693,
	7, 107, 2, 2, 693, 694, 7, 111, 2, 2, 694, 696, 7, 103, 2, 2, 695, 297,
	3, 2, 2, 2, 695, 300, 3, 2, 2, 2, 695, 306, 3, 2, 2, 2, 695, 322, 3, 2,
	2, 2, 695, 326, 3, 2, 2, 2, 695, 333, 3, 2, 2, 2, 695, 342, 3, 2, 2, 2,
	695, 351, 3, 2, 2, 2, 695, 363, 3, 2, 2, 2, 695, 374, 3, 2, 2, 2, 695,
	387, 3, 2, 2, 2, 695, 392, 3, 2, 2, 2, 695, 397, 3, 2, 2, 2, 695, 400,
	3, 2, 2, 2, 695, 405, 3, 2, 2, 2, 695, 423, 3, 2, 2, 2, 695, 435, 3, 2,
	2, 2, 695, 439, 3, 2, 2, 2, 695, 445, 3, 2, 2, 2, 695, 453, 3, 2, 2, 2,
	695, 458, 3, 2, 2, 2, 695, 468, 3, 2, 2, 2, 695, 481, 3, 2, 2, 2, 695,
	483, 3, 2, 2, 2, 695, 487, 3, 2, 2, 2, 695, 492, 3, 2, 2, 2, 695, 498,
	3, 2, 2, 2, 695, 503, 3, 2, 2, 2, 695, 517, 3, 2, 2, 2, 695, 521, 3, 2,
	2, 2, 695, 527, 3, 2, 2, 2, 695, 532, 3, 2, 2, 2, 695, 538, 3, 2, 2, 2,
	695, 542, 3, 2, 2, 2, 695, 551, 3, 2, 2, 2, 695, 555, 3, 2, 2, 2, 695,
	559, 3, 2, 2, 2, 695, 568, 3, 2, 2, 2, 695, 574, 3, 2, 2, 2, 695, 578,
	3, 2, 2, 2, 695, 591, 3, 2, 2, 2, 695, 604, 3, 2, 2, 2, 695, 617, 3, 2,
	2, 2, 695, 630, 3, 2, 2, 2, 695, 645, 3, 2, 2, 2, 695, 663, 3, 2, 2, 2,
	695, 679, 3, 2, 2, 2, 696, 62, 3, 2, 2, 2, 697, 698, 7, 125, 2, 2, 698,
	64, 3, 2, 2, 2, 699, 700, 7, 127, 2, 2, 700, 66, 3, 2, 2, 2, 701, 702,
	7, 42, 2, 2, 702, 68, 3, 2, 2, 2, 703, 704, 7, 43, 2, 2, 704, 70, 3, 2,
	2, 2, 705, 706, 7, 93, 2, 2, 706, 72, 3, 2, 2, 2, 707, 708, 7, 95, 2, 2,
	708, 74, 3, 2, 2, 2, 709, 710, 7, 46, 2, 2, 710, 76, 3, 2, 2, 2, 711, 712,
	5, 71, 36, 2, 712, 713, 5, 79, 40, 2, 713, 714, 5, 73, 37, 2, 714, 724,
	3, 2, 2, 2, 715, 716, 5, 71, 36, 2, 716, 717, 5, 79, 40, 2, 717, 719, 7,
	60, 2, 2, 718, 720, 5, 79, 40, 2, 719, 718, 3, 2, 2, 2, 719, 720, 3, 2,
	2, 2, 720, 721, 3, 2, 2, 2, 721, 722, 5, 73, 37, 2, 722, 724, 3, 2, 2,
	2, 723, 711, 3, 2, 2, 2, 723, 715, 3, 2, 2, 2, 724, 78, 3, 2, 2, 2, 725,
	727, 9, 2, 2, 2, 726, 725, 3, 2, 2, 2, 727, 728, 3, 2, 2, 2, 728, 726,
	3, 2, 2, 2, 728, 729, 3, 2, 2, 2, 729, 730, 3, 2, 2, 2, 730, 731, 9, 5,
	2, 2, 731, 80, 3, 2, 2, 2, 732, 736, 9, 6, 2, 2, 733, 735, 9, 7, 2, 2,
	734, 733, 3, 2, 2, 2, 735, 738, 3, 2, 2, 2, 736, 734, 3, 2, 2, 2, 736,
	737, 3, 2, 2, 2, 737, 82, 3, 2, 2, 2, 738, 736, 3, 2, 2, 2, 739, 743, 9,
	8, 2, 2, 740, 742, 9, 9, 2, 2, 741, 740, 3, 2, 2, 2, 742, 745, 3, 2, 2,
	2, 743, 741, 3, 2, 2, 2, 743, 744, 3, 2, 2, 2, 744, 84, 3, 2, 2, 2, 745,
	743, 3, 2, 2, 2, 746, 748, 9, 10, 2, 2, 747, 746, 3, 2, 2, 2, 748, 749,
	3, 2, 2, 2, 749, 747, 3, 2, 2, 2, 749, 750, 3, 2, 2, 2, 750, 751, 3, 2,
	2, 2, 751, 752, 8, 43, 2, 2, 752, 86, 3, 2, 2, 2, 19, 2, 90, 96, 98, 104,
	106, 114, 116, 120, 295, 695, 719, 723, 728, 736, 743, 749, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "'+'", "'-'", "'*'", "'/'", "'%'", "'^'", "'and'", "'or'",
	"'unless'", "'='", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='", "'=~'",
	"'!~'", "'by'", "'without'", "'on'", "'ignoring'", "'group_left'", "'group_right'",
	"'offset'", "'bool'", "", "", "'{'", "'}'", "'('", "')'", "'['", "']'",
	"','",
}

var lexerSymbolicNames = []string{
	"", "NUMBER", "STRING", "ADD", "SUB", "MULT", "DIV", "MOD", "POW", "AND",
	"OR", "UNLESS", "EQ", "DEQ", "NE", "GT", "LT", "GE", "LE", "RE", "NRE",
	"BY", "WITHOUT", "ON", "IGNORING", "GROUP_LEFT", "GROUP_RIGHT", "OFFSET",
	"BOOL", "AGGREGATION_OPERATOR", "FUNCTION", "LEFT_BRACE", "RIGHT_BRACE",
	"LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACKET", "RIGHT_BRACKET", "COMMA",
	"TIME_RANGE", "DURATION", "METRIC_NAME", "LABEL_NAME", "WS",
}

var lexerRuleNames = []string{
	"NUMBER", "STRING", "ADD", "SUB", "MULT", "DIV", "MOD", "POW", "AND", "OR",
	"UNLESS", "EQ", "DEQ", "NE", "GT", "LT", "GE", "LE", "RE", "NRE", "BY",
	"WITHOUT", "ON", "IGNORING", "GROUP_LEFT", "GROUP_RIGHT", "OFFSET", "BOOL",
	"AGGREGATION_OPERATOR", "FUNCTION", "LEFT_BRACE", "RIGHT_BRACE", "LEFT_PAREN",
	"RIGHT_PAREN", "LEFT_BRACKET", "RIGHT_BRACKET", "COMMA", "TIME_RANGE",
	"DURATION", "METRIC_NAME", "LABEL_NAME", "WS",
}

type PromQLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewPromQLLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *PromQLLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPromQLLexer(input antlr.CharStream) *PromQLLexer {
	l := new(PromQLLexer)
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
	l.GrammarFileName = "PromQLLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// PromQLLexer tokens.
const (
	PromQLLexerNUMBER               = 1
	PromQLLexerSTRING               = 2
	PromQLLexerADD                  = 3
	PromQLLexerSUB                  = 4
	PromQLLexerMULT                 = 5
	PromQLLexerDIV                  = 6
	PromQLLexerMOD                  = 7
	PromQLLexerPOW                  = 8
	PromQLLexerAND                  = 9
	PromQLLexerOR                   = 10
	PromQLLexerUNLESS               = 11
	PromQLLexerEQ                   = 12
	PromQLLexerDEQ                  = 13
	PromQLLexerNE                   = 14
	PromQLLexerGT                   = 15
	PromQLLexerLT                   = 16
	PromQLLexerGE                   = 17
	PromQLLexerLE                   = 18
	PromQLLexerRE                   = 19
	PromQLLexerNRE                  = 20
	PromQLLexerBY                   = 21
	PromQLLexerWITHOUT              = 22
	PromQLLexerON                   = 23
	PromQLLexerIGNORING             = 24
	PromQLLexerGROUP_LEFT           = 25
	PromQLLexerGROUP_RIGHT          = 26
	PromQLLexerOFFSET               = 27
	PromQLLexerBOOL                 = 28
	PromQLLexerAGGREGATION_OPERATOR = 29
	PromQLLexerFUNCTION             = 30
	PromQLLexerLEFT_BRACE           = 31
	PromQLLexerRIGHT_BRACE          = 32
	PromQLLexerLEFT_PAREN           = 33
	PromQLLexerRIGHT_PAREN          = 34
	PromQLLexerLEFT_BRACKET         = 35
	PromQLLexerRIGHT_BRACKET        = 36
	PromQLLexerCOMMA                = 37
	PromQLLexerTIME_RANGE           = 38
	PromQLLexerDURATION             = 39
	PromQLLexerMETRIC_NAME          = 40
	PromQLLexerLABEL_NAME           = 41
	PromQLLexerWS                   = 42
)
