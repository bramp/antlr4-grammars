// Code generated from DCM_2_0_grammar.g4 by ANTLR 4.9.3. DO NOT EDIT.

package dcm_2_0_grammar // DCM_2_0_grammar
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 590,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 3, 2, 7, 2,
	88, 10, 2, 12, 2, 14, 2, 91, 11, 2, 3, 2, 3, 2, 6, 2, 95, 10, 2, 13, 2,
	14, 2, 96, 3, 2, 3, 2, 3, 2, 3, 3, 5, 3, 103, 10, 3, 3, 3, 5, 3, 106, 10,
	3, 3, 3, 5, 3, 109, 10, 3, 3, 4, 6, 4, 112, 10, 4, 13, 4, 14, 4, 113, 3,
	5, 3, 5, 7, 5, 118, 10, 5, 12, 5, 14, 5, 121, 11, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 6, 9, 134, 10, 9, 13, 9,
	14, 9, 135, 3, 10, 3, 10, 3, 10, 6, 10, 141, 10, 10, 13, 10, 14, 10, 142,
	3, 10, 3, 10, 6, 10, 147, 10, 10, 13, 10, 14, 10, 148, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 6, 13, 160, 10, 13, 13, 13,
	14, 13, 161, 3, 14, 3, 14, 3, 14, 6, 14, 167, 10, 14, 13, 14, 14, 14, 168,
	3, 14, 3, 14, 6, 14, 173, 10, 14, 13, 14, 14, 14, 174, 3, 15, 3, 15, 3,
	15, 7, 15, 180, 10, 15, 12, 15, 14, 15, 183, 11, 15, 3, 15, 6, 15, 186,
	10, 15, 13, 15, 14, 15, 187, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 7, 18,
	195, 10, 18, 12, 18, 14, 18, 198, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 5, 19, 206, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20,
	213, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 6, 20, 220, 10, 20, 13,
	20, 14, 20, 221, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 229, 10, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 6, 20, 236, 10, 20, 13, 20, 14, 20,
	237, 5, 20, 240, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21,
	248, 10, 21, 3, 21, 6, 21, 251, 10, 21, 13, 21, 14, 21, 252, 3, 21, 3,
	21, 6, 21, 257, 10, 21, 13, 21, 14, 21, 258, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 5, 22, 267, 10, 22, 3, 22, 5, 22, 270, 10, 22, 3, 22, 6,
	22, 273, 10, 22, 13, 22, 14, 22, 274, 3, 22, 6, 22, 278, 10, 22, 13, 22,
	14, 22, 279, 3, 22, 3, 22, 6, 22, 284, 10, 22, 13, 22, 14, 22, 285, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 294, 10, 22, 3, 22, 5, 22,
	297, 10, 22, 3, 22, 6, 22, 300, 10, 22, 13, 22, 14, 22, 301, 3, 22, 6,
	22, 305, 10, 22, 13, 22, 14, 22, 306, 3, 22, 3, 22, 6, 22, 311, 10, 22,
	13, 22, 14, 22, 312, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 321,
	10, 22, 3, 22, 5, 22, 324, 10, 22, 3, 22, 6, 22, 327, 10, 22, 13, 22, 14,
	22, 328, 3, 22, 6, 22, 332, 10, 22, 13, 22, 14, 22, 333, 3, 22, 3, 22,
	6, 22, 338, 10, 22, 13, 22, 14, 22, 339, 5, 22, 342, 10, 22, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 351, 10, 23, 3, 23, 5, 23,
	354, 10, 23, 3, 23, 5, 23, 357, 10, 23, 3, 23, 6, 23, 360, 10, 23, 13,
	23, 14, 23, 361, 3, 23, 3, 23, 3, 23, 6, 23, 367, 10, 23, 13, 23, 14, 23,
	368, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 378, 10, 23,
	3, 23, 5, 23, 381, 10, 23, 3, 23, 5, 23, 384, 10, 23, 3, 23, 6, 23, 387,
	10, 23, 13, 23, 14, 23, 388, 3, 23, 3, 23, 3, 23, 6, 23, 394, 10, 23, 13,
	23, 14, 23, 395, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23,
	405, 10, 23, 3, 23, 5, 23, 408, 10, 23, 3, 23, 5, 23, 411, 10, 23, 3, 23,
	6, 23, 414, 10, 23, 13, 23, 14, 23, 415, 3, 23, 3, 23, 3, 23, 6, 23, 421,
	10, 23, 13, 23, 14, 23, 422, 5, 23, 425, 10, 23, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 5, 24, 433, 10, 24, 3, 24, 6, 24, 436, 10, 24, 13, 24,
	14, 24, 437, 3, 24, 3, 24, 6, 24, 442, 10, 24, 13, 24, 14, 24, 443, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 6, 25, 455,
	10, 25, 13, 25, 14, 25, 456, 3, 26, 5, 26, 460, 10, 26, 3, 26, 5, 26, 463,
	10, 26, 3, 26, 5, 26, 466, 10, 26, 3, 26, 5, 26, 469, 10, 26, 3, 27, 3,
	27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3,
	32, 3, 32, 7, 32, 495, 10, 32, 12, 32, 14, 32, 498, 11, 32, 3, 32, 3, 32,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 6, 34, 508, 10, 34, 13, 34, 14,
	34, 509, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 6, 37,
	520, 10, 37, 13, 37, 14, 37, 521, 3, 37, 3, 37, 3, 38, 3, 38, 6, 38, 528,
	10, 38, 13, 38, 14, 38, 529, 3, 38, 3, 38, 3, 38, 3, 38, 6, 38, 536, 10,
	38, 13, 38, 14, 38, 537, 3, 38, 5, 38, 541, 10, 38, 3, 39, 3, 39, 6, 39,
	545, 10, 39, 13, 39, 14, 39, 546, 3, 39, 3, 39, 3, 39, 3, 39, 6, 39, 553,
	10, 39, 13, 39, 14, 39, 554, 3, 39, 5, 39, 558, 10, 39, 3, 40, 6, 40, 561,
	10, 40, 13, 40, 14, 40, 562, 3, 40, 6, 40, 566, 10, 40, 13, 40, 14, 40,
	567, 5, 40, 570, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41, 6, 41, 576, 10, 41,
	13, 41, 14, 41, 577, 3, 42, 3, 42, 3, 42, 3, 42, 6, 42, 584, 10, 42, 13,
	42, 14, 42, 585, 3, 43, 3, 43, 3, 43, 2, 2, 44, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 2,
	4, 3, 2, 36, 37, 3, 2, 38, 39, 2, 632, 2, 89, 3, 2, 2, 2, 4, 102, 3, 2,
	2, 2, 6, 111, 3, 2, 2, 2, 8, 115, 3, 2, 2, 2, 10, 122, 3, 2, 2, 2, 12,
	126, 3, 2, 2, 2, 14, 129, 3, 2, 2, 2, 16, 131, 3, 2, 2, 2, 18, 137, 3,
	2, 2, 2, 20, 150, 3, 2, 2, 2, 22, 155, 3, 2, 2, 2, 24, 157, 3, 2, 2, 2,
	26, 163, 3, 2, 2, 2, 28, 176, 3, 2, 2, 2, 30, 189, 3, 2, 2, 2, 32, 191,
	3, 2, 2, 2, 34, 196, 3, 2, 2, 2, 36, 205, 3, 2, 2, 2, 38, 239, 3, 2, 2,
	2, 40, 241, 3, 2, 2, 2, 42, 341, 3, 2, 2, 2, 44, 424, 3, 2, 2, 2, 46, 426,
	3, 2, 2, 2, 48, 445, 3, 2, 2, 2, 50, 459, 3, 2, 2, 2, 52, 470, 3, 2, 2,
	2, 54, 474, 3, 2, 2, 2, 56, 478, 3, 2, 2, 2, 58, 482, 3, 2, 2, 2, 60, 486,
	3, 2, 2, 2, 62, 490, 3, 2, 2, 2, 64, 501, 3, 2, 2, 2, 66, 505, 3, 2, 2,
	2, 68, 513, 3, 2, 2, 2, 70, 515, 3, 2, 2, 2, 72, 517, 3, 2, 2, 2, 74, 540,
	3, 2, 2, 2, 76, 557, 3, 2, 2, 2, 78, 569, 3, 2, 2, 2, 80, 571, 3, 2, 2,
	2, 82, 579, 3, 2, 2, 2, 84, 587, 3, 2, 2, 2, 86, 88, 7, 3, 2, 2, 87, 86,
	3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2,
	90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 94, 7, 4, 2, 2, 93, 95, 7,
	3, 2, 2, 94, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96,
	97, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 99, 5, 4, 3, 2, 99, 100, 5, 34,
	18, 2, 100, 3, 3, 2, 2, 2, 101, 103, 5, 6, 4, 2, 102, 101, 3, 2, 2, 2,
	102, 103, 3, 2, 2, 2, 103, 105, 3, 2, 2, 2, 104, 106, 5, 18, 10, 2, 105,
	104, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 108, 3, 2, 2, 2, 107, 109,
	5, 26, 14, 2, 108, 107, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 5, 3, 2,
	2, 2, 110, 112, 5, 8, 5, 2, 111, 110, 3, 2, 2, 2, 112, 113, 3, 2, 2, 2,
	113, 111, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 7, 3, 2, 2, 2, 115, 119,
	5, 10, 6, 2, 116, 118, 5, 12, 7, 2, 117, 116, 3, 2, 2, 2, 118, 121, 3,
	2, 2, 2, 119, 117, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 9, 3, 2, 2, 2,
	121, 119, 3, 2, 2, 2, 122, 123, 7, 5, 2, 2, 123, 124, 5, 14, 8, 2, 124,
	125, 5, 16, 9, 2, 125, 11, 3, 2, 2, 2, 126, 127, 7, 5, 2, 2, 127, 128,
	5, 16, 9, 2, 128, 13, 3, 2, 2, 2, 129, 130, 7, 36, 2, 2, 130, 15, 3, 2,
	2, 2, 131, 133, 7, 37, 2, 2, 132, 134, 7, 3, 2, 2, 133, 132, 3, 2, 2, 2,
	134, 135, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135, 136, 3, 2, 2, 2, 136,
	17, 3, 2, 2, 2, 137, 138, 7, 6, 2, 2, 138, 140, 7, 3, 2, 2, 139, 141, 5,
	20, 11, 2, 140, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 140, 3, 2,
	2, 2, 142, 143, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 146, 7, 7, 2, 2,
	145, 147, 7, 3, 2, 2, 146, 145, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148,
	146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 19, 3, 2, 2, 2, 150, 151, 7,
	8, 2, 2, 151, 152, 7, 36, 2, 2, 152, 153, 5, 22, 12, 2, 153, 154, 5, 24,
	13, 2, 154, 21, 3, 2, 2, 2, 155, 156, 7, 37, 2, 2, 156, 23, 3, 2, 2, 2,
	157, 159, 7, 37, 2, 2, 158, 160, 7, 3, 2, 2, 159, 158, 3, 2, 2, 2, 160,
	161, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 25, 3,
	2, 2, 2, 163, 164, 7, 9, 2, 2, 164, 166, 7, 3, 2, 2, 165, 167, 5, 28, 15,
	2, 166, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 168,
	169, 3, 2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 172, 7, 7, 2, 2, 171, 173,
	7, 3, 2, 2, 172, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2, 174, 172, 3, 2,
	2, 2, 174, 175, 3, 2, 2, 2, 175, 27, 3, 2, 2, 2, 176, 177, 7, 10, 2, 2,
	177, 181, 5, 30, 16, 2, 178, 180, 5, 32, 17, 2, 179, 178, 3, 2, 2, 2, 180,
	183, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 185,
	3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 184, 186, 7, 3, 2, 2, 185, 184, 3, 2,
	2, 2, 186, 187, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2,
	188, 29, 3, 2, 2, 2, 189, 190, 7, 36, 2, 2, 190, 31, 3, 2, 2, 2, 191, 192,
	7, 36, 2, 2, 192, 33, 3, 2, 2, 2, 193, 195, 5, 36, 19, 2, 194, 193, 3,
	2, 2, 2, 195, 198, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 196, 197, 3, 2, 2,
	2, 197, 35, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 199, 206, 5, 38, 20, 2, 200,
	206, 5, 40, 21, 2, 201, 206, 5, 42, 22, 2, 202, 206, 5, 44, 23, 2, 203,
	206, 5, 46, 24, 2, 204, 206, 5, 48, 25, 2, 205, 199, 3, 2, 2, 2, 205, 200,
	3, 2, 2, 2, 205, 201, 3, 2, 2, 2, 205, 202, 3, 2, 2, 2, 205, 203, 3, 2,
	2, 2, 205, 204, 3, 2, 2, 2, 206, 37, 3, 2, 2, 2, 207, 208, 7, 11, 2, 2,
	208, 209, 7, 36, 2, 2, 209, 210, 7, 3, 2, 2, 210, 212, 5, 50, 26, 2, 211,
	213, 5, 56, 29, 2, 212, 211, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 214,
	3, 2, 2, 2, 214, 215, 7, 12, 2, 2, 215, 216, 5, 84, 43, 2, 216, 217, 7,
	3, 2, 2, 217, 219, 7, 7, 2, 2, 218, 220, 7, 3, 2, 2, 219, 218, 3, 2, 2,
	2, 220, 221, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222,
	240, 3, 2, 2, 2, 223, 224, 7, 11, 2, 2, 224, 225, 7, 36, 2, 2, 225, 226,
	7, 3, 2, 2, 226, 228, 5, 50, 26, 2, 227, 229, 5, 56, 29, 2, 228, 227, 3,
	2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 231, 7, 13, 2,
	2, 231, 232, 7, 37, 2, 2, 232, 233, 7, 3, 2, 2, 233, 235, 7, 7, 2, 2, 234,
	236, 7, 3, 2, 2, 235, 234, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 235,
	3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 240, 3, 2, 2, 2, 239, 207, 3, 2,
	2, 2, 239, 223, 3, 2, 2, 2, 240, 39, 3, 2, 2, 2, 241, 242, 7, 14, 2, 2,
	242, 243, 7, 36, 2, 2, 243, 244, 5, 68, 35, 2, 244, 245, 7, 3, 2, 2, 245,
	247, 5, 50, 26, 2, 246, 248, 5, 56, 29, 2, 247, 246, 3, 2, 2, 2, 247, 248,
	3, 2, 2, 2, 248, 250, 3, 2, 2, 2, 249, 251, 5, 74, 38, 2, 250, 249, 3,
	2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2,
	2, 253, 254, 3, 2, 2, 2, 254, 256, 7, 7, 2, 2, 255, 257, 7, 3, 2, 2, 256,
	255, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 258, 259,
	3, 2, 2, 2, 259, 41, 3, 2, 2, 2, 260, 261, 7, 15, 2, 2, 261, 262, 7, 36,
	2, 2, 262, 263, 5, 68, 35, 2, 263, 264, 7, 3, 2, 2, 264, 266, 5, 50, 26,
	2, 265, 267, 5, 52, 27, 2, 266, 265, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2,
	267, 269, 3, 2, 2, 2, 268, 270, 5, 56, 29, 2, 269, 268, 3, 2, 2, 2, 269,
	270, 3, 2, 2, 2, 270, 272, 3, 2, 2, 2, 271, 273, 5, 76, 39, 2, 272, 271,
	3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2, 274, 275, 3, 2,
	2, 2, 275, 277, 3, 2, 2, 2, 276, 278, 5, 72, 37, 2, 277, 276, 3, 2, 2,
	2, 278, 279, 3, 2, 2, 2, 279, 277, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280,
	281, 3, 2, 2, 2, 281, 283, 7, 7, 2, 2, 282, 284, 7, 3, 2, 2, 283, 282,
	3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 285, 286, 3, 2,
	2, 2, 286, 342, 3, 2, 2, 2, 287, 288, 7, 16, 2, 2, 288, 289, 7, 36, 2,
	2, 289, 290, 5, 68, 35, 2, 290, 291, 7, 3, 2, 2, 291, 293, 5, 50, 26, 2,
	292, 294, 5, 52, 27, 2, 293, 292, 3, 2, 2, 2, 293, 294, 3, 2, 2, 2, 294,
	296, 3, 2, 2, 2, 295, 297, 5, 56, 29, 2, 296, 295, 3, 2, 2, 2, 296, 297,
	3, 2, 2, 2, 297, 299, 3, 2, 2, 2, 298, 300, 5, 76, 39, 2, 299, 298, 3,
	2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 299, 3, 2, 2, 2, 301, 302, 3, 2, 2,
	2, 302, 304, 3, 2, 2, 2, 303, 305, 5, 72, 37, 2, 304, 303, 3, 2, 2, 2,
	305, 306, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307,
	308, 3, 2, 2, 2, 308, 310, 7, 7, 2, 2, 309, 311, 7, 3, 2, 2, 310, 309,
	3, 2, 2, 2, 311, 312, 3, 2, 2, 2, 312, 310, 3, 2, 2, 2, 312, 313, 3, 2,
	2, 2, 313, 342, 3, 2, 2, 2, 314, 315, 7, 17, 2, 2, 315, 316, 7, 36, 2,
	2, 316, 317, 5, 68, 35, 2, 317, 318, 7, 3, 2, 2, 318, 320, 5, 50, 26, 2,
	319, 321, 5, 52, 27, 2, 320, 319, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2, 321,
	323, 3, 2, 2, 2, 322, 324, 5, 56, 29, 2, 323, 322, 3, 2, 2, 2, 323, 324,
	3, 2, 2, 2, 324, 326, 3, 2, 2, 2, 325, 327, 5, 76, 39, 2, 326, 325, 3,
	2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2, 328, 329, 3, 2, 2,
	2, 329, 331, 3, 2, 2, 2, 330, 332, 5, 72, 37, 2, 331, 330, 3, 2, 2, 2,
	332, 333, 3, 2, 2, 2, 333, 331, 3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334,
	335, 3, 2, 2, 2, 335, 337, 7, 7, 2, 2, 336, 338, 7, 3, 2, 2, 337, 336,
	3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 337, 3, 2, 2, 2, 339, 340, 3, 2,
	2, 2, 340, 342, 3, 2, 2, 2, 341, 260, 3, 2, 2, 2, 341, 287, 3, 2, 2, 2,
	341, 314, 3, 2, 2, 2, 342, 43, 3, 2, 2, 2, 343, 344, 7, 18, 2, 2, 344,
	345, 7, 36, 2, 2, 345, 346, 5, 68, 35, 2, 346, 347, 5, 70, 36, 2, 347,
	348, 7, 3, 2, 2, 348, 350, 5, 50, 26, 2, 349, 351, 5, 52, 27, 2, 350, 349,
	3, 2, 2, 2, 350, 351, 3, 2, 2, 2, 351, 353, 3, 2, 2, 2, 352, 354, 5, 54,
	28, 2, 353, 352, 3, 2, 2, 2, 353, 354, 3, 2, 2, 2, 354, 356, 3, 2, 2, 2,
	355, 357, 5, 56, 29, 2, 356, 355, 3, 2, 2, 2, 356, 357, 3, 2, 2, 2, 357,
	359, 3, 2, 2, 2, 358, 360, 5, 76, 39, 2, 359, 358, 3, 2, 2, 2, 360, 361,
	3, 2, 2, 2, 361, 359, 3, 2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 363, 3, 2,
	2, 2, 363, 364, 5, 78, 40, 2, 364, 366, 7, 7, 2, 2, 365, 367, 7, 3, 2,
	2, 366, 365, 3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368, 366, 3, 2, 2, 2, 368,
	369, 3, 2, 2, 2, 369, 425, 3, 2, 2, 2, 370, 371, 7, 19, 2, 2, 371, 372,
	7, 36, 2, 2, 372, 373, 5, 68, 35, 2, 373, 374, 5, 70, 36, 2, 374, 375,
	7, 3, 2, 2, 375, 377, 5, 50, 26, 2, 376, 378, 5, 52, 27, 2, 377, 376, 3,
	2, 2, 2, 377, 378, 3, 2, 2, 2, 378, 380, 3, 2, 2, 2, 379, 381, 5, 54, 28,
	2, 380, 379, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 383, 3, 2, 2, 2, 382,
	384, 5, 56, 29, 2, 383, 382, 3, 2, 2, 2, 383, 384, 3, 2, 2, 2, 384, 386,
	3, 2, 2, 2, 385, 387, 5, 76, 39, 2, 386, 385, 3, 2, 2, 2, 387, 388, 3,
	2, 2, 2, 388, 386, 3, 2, 2, 2, 388, 389, 3, 2, 2, 2, 389, 390, 3, 2, 2,
	2, 390, 391, 5, 78, 40, 2, 391, 393, 7, 7, 2, 2, 392, 394, 7, 3, 2, 2,
	393, 392, 3, 2, 2, 2, 394, 395, 3, 2, 2, 2, 395, 393, 3, 2, 2, 2, 395,
	396, 3, 2, 2, 2, 396, 425, 3, 2, 2, 2, 397, 398, 7, 20, 2, 2, 398, 399,
	7, 36, 2, 2, 399, 400, 5, 68, 35, 2, 400, 401, 5, 70, 36, 2, 401, 402,
	7, 3, 2, 2, 402, 404, 5, 50, 26, 2, 403, 405, 5, 52, 27, 2, 404, 403, 3,
	2, 2, 2, 404, 405, 3, 2, 2, 2, 405, 407, 3, 2, 2, 2, 406, 408, 5, 54, 28,
	2, 407, 406, 3, 2, 2, 2, 407, 408, 3, 2, 2, 2, 408, 410, 3, 2, 2, 2, 409,
	411, 5, 56, 29, 2, 410, 409, 3, 2, 2, 2, 410, 411, 3, 2, 2, 2, 411, 413,
	3, 2, 2, 2, 412, 414, 5, 76, 39, 2, 413, 412, 3, 2, 2, 2, 414, 415, 3,
	2, 2, 2, 415, 413, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 417, 3, 2, 2,
	2, 417, 418, 5, 78, 40, 2, 418, 420, 7, 7, 2, 2, 419, 421, 7, 3, 2, 2,
	420, 419, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422, 420, 3, 2, 2, 2, 422,
	423, 3, 2, 2, 2, 423, 425, 3, 2, 2, 2, 424, 343, 3, 2, 2, 2, 424, 370,
	3, 2, 2, 2, 424, 397, 3, 2, 2, 2, 425, 45, 3, 2, 2, 2, 426, 427, 7, 21,
	2, 2, 427, 428, 7, 36, 2, 2, 428, 429, 5, 68, 35, 2, 429, 430, 7, 3, 2,
	2, 430, 432, 5, 50, 26, 2, 431, 433, 5, 52, 27, 2, 432, 431, 3, 2, 2, 2,
	432, 433, 3, 2, 2, 2, 433, 435, 3, 2, 2, 2, 434, 436, 5, 76, 39, 2, 435,
	434, 3, 2, 2, 2, 436, 437, 3, 2, 2, 2, 437, 435, 3, 2, 2, 2, 437, 438,
	3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 441, 7, 7, 2, 2, 440, 442, 7, 3,
	2, 2, 441, 440, 3, 2, 2, 2, 442, 443, 3, 2, 2, 2, 443, 441, 3, 2, 2, 2,
	443, 444, 3, 2, 2, 2, 444, 47, 3, 2, 2, 2, 445, 446, 7, 22, 2, 2, 446,
	447, 7, 36, 2, 2, 447, 448, 7, 3, 2, 2, 448, 449, 5, 50, 26, 2, 449, 450,
	7, 13, 2, 2, 450, 451, 7, 37, 2, 2, 451, 452, 7, 3, 2, 2, 452, 454, 7,
	7, 2, 2, 453, 455, 7, 3, 2, 2, 454, 453, 3, 2, 2, 2, 455, 456, 3, 2, 2,
	2, 456, 454, 3, 2, 2, 2, 456, 457, 3, 2, 2, 2, 457, 49, 3, 2, 2, 2, 458,
	460, 5, 58, 30, 2, 459, 458, 3, 2, 2, 2, 459, 460, 3, 2, 2, 2, 460, 462,
	3, 2, 2, 2, 461, 463, 5, 60, 31, 2, 462, 461, 3, 2, 2, 2, 462, 463, 3,
	2, 2, 2, 463, 465, 3, 2, 2, 2, 464, 466, 5, 62, 32, 2, 465, 464, 3, 2,
	2, 2, 465, 466, 3, 2, 2, 2, 466, 468, 3, 2, 2, 2, 467, 469, 5, 66, 34,
	2, 468, 467, 3, 2, 2, 2, 468, 469, 3, 2, 2, 2, 469, 51, 3, 2, 2, 2, 470,
	471, 7, 23, 2, 2, 471, 472, 7, 37, 2, 2, 472, 473, 7, 3, 2, 2, 473, 53,
	3, 2, 2, 2, 474, 475, 7, 24, 2, 2, 475, 476, 7, 37, 2, 2, 476, 477, 7,
	3, 2, 2, 477, 55, 3, 2, 2, 2, 478, 479, 7, 25, 2, 2, 479, 480, 7, 37, 2,
	2, 480, 481, 7, 3, 2, 2, 481, 57, 3, 2, 2, 2, 482, 483, 7, 26, 2, 2, 483,
	484, 7, 37, 2, 2, 484, 485, 7, 3, 2, 2, 485, 59, 3, 2, 2, 2, 486, 487,
	7, 27, 2, 2, 487, 488, 9, 2, 2, 2, 488, 489, 7, 3, 2, 2, 489, 61, 3, 2,
	2, 2, 490, 491, 7, 28, 2, 2, 491, 496, 5, 64, 33, 2, 492, 493, 7, 29, 2,
	2, 493, 495, 5, 64, 33, 2, 494, 492, 3, 2, 2, 2, 495, 498, 3, 2, 2, 2,
	496, 494, 3, 2, 2, 2, 496, 497, 3, 2, 2, 2, 497, 499, 3, 2, 2, 2, 498,
	496, 3, 2, 2, 2, 499, 500, 7, 3, 2, 2, 500, 63, 3, 2, 2, 2, 501, 502, 7,
	36, 2, 2, 502, 503, 7, 30, 2, 2, 503, 504, 7, 36, 2, 2, 504, 65, 3, 2,
	2, 2, 505, 507, 7, 31, 2, 2, 506, 508, 7, 36, 2, 2, 507, 506, 3, 2, 2,
	2, 508, 509, 3, 2, 2, 2, 509, 507, 3, 2, 2, 2, 509, 510, 3, 2, 2, 2, 510,
	511, 3, 2, 2, 2, 511, 512, 7, 3, 2, 2, 512, 67, 3, 2, 2, 2, 513, 514, 7,
	38, 2, 2, 514, 69, 3, 2, 2, 2, 515, 516, 7, 38, 2, 2, 516, 71, 3, 2, 2,
	2, 517, 519, 7, 12, 2, 2, 518, 520, 5, 84, 43, 2, 519, 518, 3, 2, 2, 2,
	520, 521, 3, 2, 2, 2, 521, 519, 3, 2, 2, 2, 521, 522, 3, 2, 2, 2, 522,
	523, 3, 2, 2, 2, 523, 524, 7, 3, 2, 2, 524, 73, 3, 2, 2, 2, 525, 527, 7,
	12, 2, 2, 526, 528, 5, 84, 43, 2, 527, 526, 3, 2, 2, 2, 528, 529, 3, 2,
	2, 2, 529, 527, 3, 2, 2, 2, 529, 530, 3, 2, 2, 2, 530, 531, 3, 2, 2, 2,
	531, 532, 7, 3, 2, 2, 532, 541, 3, 2, 2, 2, 533, 535, 7, 13, 2, 2, 534,
	536, 7, 37, 2, 2, 535, 534, 3, 2, 2, 2, 536, 537, 3, 2, 2, 2, 537, 535,
	3, 2, 2, 2, 537, 538, 3, 2, 2, 2, 538, 539, 3, 2, 2, 2, 539, 541, 7, 3,
	2, 2, 540, 525, 3, 2, 2, 2, 540, 533, 3, 2, 2, 2, 541, 75, 3, 2, 2, 2,
	542, 544, 7, 32, 2, 2, 543, 545, 5, 84, 43, 2, 544, 543, 3, 2, 2, 2, 545,
	546, 3, 2, 2, 2, 546, 544, 3, 2, 2, 2, 546, 547, 3, 2, 2, 2, 547, 548,
	3, 2, 2, 2, 548, 549, 7, 3, 2, 2, 549, 558, 3, 2, 2, 2, 550, 552, 7, 33,
	2, 2, 551, 553, 7, 37, 2, 2, 552, 551, 3, 2, 2, 2, 553, 554, 3, 2, 2, 2,
	554, 552, 3, 2, 2, 2, 554, 555, 3, 2, 2, 2, 555, 556, 3, 2, 2, 2, 556,
	558, 7, 3, 2, 2, 557, 542, 3, 2, 2, 2, 557, 550, 3, 2, 2, 2, 558, 77, 3,
	2, 2, 2, 559, 561, 5, 80, 41, 2, 560, 559, 3, 2, 2, 2, 561, 562, 3, 2,
	2, 2, 562, 560, 3, 2, 2, 2, 562, 563, 3, 2, 2, 2, 563, 570, 3, 2, 2, 2,
	564, 566, 5, 82, 42, 2, 565, 564, 3, 2, 2, 2, 566, 567, 3, 2, 2, 2, 567,
	565, 3, 2, 2, 2, 567, 568, 3, 2, 2, 2, 568, 570, 3, 2, 2, 2, 569, 560,
	3, 2, 2, 2, 569, 565, 3, 2, 2, 2, 570, 79, 3, 2, 2, 2, 571, 572, 7, 34,
	2, 2, 572, 573, 5, 84, 43, 2, 573, 575, 7, 3, 2, 2, 574, 576, 5, 72, 37,
	2, 575, 574, 3, 2, 2, 2, 576, 577, 3, 2, 2, 2, 577, 575, 3, 2, 2, 2, 577,
	578, 3, 2, 2, 2, 578, 81, 3, 2, 2, 2, 579, 580, 7, 35, 2, 2, 580, 581,
	7, 37, 2, 2, 581, 583, 7, 3, 2, 2, 582, 584, 5, 72, 37, 2, 583, 582, 3,
	2, 2, 2, 584, 585, 3, 2, 2, 2, 585, 583, 3, 2, 2, 2, 585, 586, 3, 2, 2,
	2, 586, 83, 3, 2, 2, 2, 587, 588, 9, 3, 2, 2, 588, 85, 3, 2, 2, 2, 81,
	89, 96, 102, 105, 108, 113, 119, 135, 142, 148, 161, 168, 174, 181, 187,
	196, 205, 212, 221, 228, 237, 239, 247, 252, 258, 266, 269, 274, 279, 285,
	293, 296, 301, 306, 312, 320, 323, 328, 333, 339, 341, 350, 353, 356, 361,
	368, 377, 380, 383, 388, 395, 404, 407, 410, 415, 422, 424, 432, 437, 443,
	456, 459, 462, 465, 468, 496, 509, 521, 529, 537, 540, 546, 554, 557, 562,
	567, 569, 577, 585,
}
var literalNames = []string{
	"", "'\n'", "'KONSERVIERUNG_FORMAT 2.0'", "'MODULKOPF'", "'FUNKTIONEN'",
	"'END'", "'FKT'", "'VARIANTENKODIERUNG'", "'KRITERIUM'", "'FESTWERT'",
	"'WERT'", "'TEXT'", "'FESTWERTEBLOCK'", "'KENNLINIE'", "'FESTKENNLINIE'",
	"'GRUPPENKENNLINIE'", "'KENNFELD'", "'FESTKENNFELD'", "'GRUPPENKENNFELD'",
	"'STUETZSTELLENVERTEILUNG'", "'TEXTSTRING'", "'EINHEIT_X'", "'EINHEIT_Y'",
	"'EINHEIT_W'", "'LANGNAME'", "'DISPLAYNAME'", "'VAR'", "','", "'='", "'FUNKTION'",
	"'ST/X'", "'ST_TX/X'", "'ST/Y'", "'ST_TX/Y'", "", "", "", "", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "NAME",
	"TEXT", "INT", "FLOAT", "MINUS", "WS", "COMMENT",
}

var ruleNames = []string{
	"konservierung", "kons_kopf", "modulkopf_info", "mod_zeile", "mod_anf_zeile",
	"mod_fort_zeile", "mod_ele_name", "mod_ele_wert", "funktionsdef", "funktionszeile",
	"fkt_version", "fkt_langname", "variantendef", "variantenkrit", "krit_name",
	"krit_wert", "kons_rumpf", "kenngroesse", "kennwert", "kennwerteblock",
	"kennlinie", "kennfeld", "gruppenstuetzstellen", "kenntext", "kgr_info",
	"einheit_x", "einheit_y", "einheit_w", "langname", "displayname", "var_abhangigkeiten",
	"var_abh", "funktionszugehorigkeit", "anzahl_x", "anzahl_y", "werteliste",
	"werteliste_kwb", "sst_liste_x", "kf_zeile_liste", "kf_zeile_liste_r",
	"kf_zeile_liste_tx", "realzahl",
}

type DCM_2_0_grammarParser struct {
	*antlr.BaseParser
}

// NewDCM_2_0_grammarParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DCM_2_0_grammarParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDCM_2_0_grammarParser(input antlr.TokenStream) *DCM_2_0_grammarParser {
	this := new(DCM_2_0_grammarParser)
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
	this.GrammarFileName = "DCM_2_0_grammar.g4"

	return this
}

// DCM_2_0_grammarParser tokens.
const (
	DCM_2_0_grammarParserEOF     = antlr.TokenEOF
	DCM_2_0_grammarParserT__0    = 1
	DCM_2_0_grammarParserT__1    = 2
	DCM_2_0_grammarParserT__2    = 3
	DCM_2_0_grammarParserT__3    = 4
	DCM_2_0_grammarParserT__4    = 5
	DCM_2_0_grammarParserT__5    = 6
	DCM_2_0_grammarParserT__6    = 7
	DCM_2_0_grammarParserT__7    = 8
	DCM_2_0_grammarParserT__8    = 9
	DCM_2_0_grammarParserT__9    = 10
	DCM_2_0_grammarParserT__10   = 11
	DCM_2_0_grammarParserT__11   = 12
	DCM_2_0_grammarParserT__12   = 13
	DCM_2_0_grammarParserT__13   = 14
	DCM_2_0_grammarParserT__14   = 15
	DCM_2_0_grammarParserT__15   = 16
	DCM_2_0_grammarParserT__16   = 17
	DCM_2_0_grammarParserT__17   = 18
	DCM_2_0_grammarParserT__18   = 19
	DCM_2_0_grammarParserT__19   = 20
	DCM_2_0_grammarParserT__20   = 21
	DCM_2_0_grammarParserT__21   = 22
	DCM_2_0_grammarParserT__22   = 23
	DCM_2_0_grammarParserT__23   = 24
	DCM_2_0_grammarParserT__24   = 25
	DCM_2_0_grammarParserT__25   = 26
	DCM_2_0_grammarParserT__26   = 27
	DCM_2_0_grammarParserT__27   = 28
	DCM_2_0_grammarParserT__28   = 29
	DCM_2_0_grammarParserT__29   = 30
	DCM_2_0_grammarParserT__30   = 31
	DCM_2_0_grammarParserT__31   = 32
	DCM_2_0_grammarParserT__32   = 33
	DCM_2_0_grammarParserNAME    = 34
	DCM_2_0_grammarParserTEXT    = 35
	DCM_2_0_grammarParserINT     = 36
	DCM_2_0_grammarParserFLOAT   = 37
	DCM_2_0_grammarParserMINUS   = 38
	DCM_2_0_grammarParserWS      = 39
	DCM_2_0_grammarParserCOMMENT = 40
)

// DCM_2_0_grammarParser rules.
const (
	DCM_2_0_grammarParserRULE_konservierung          = 0
	DCM_2_0_grammarParserRULE_kons_kopf              = 1
	DCM_2_0_grammarParserRULE_modulkopf_info         = 2
	DCM_2_0_grammarParserRULE_mod_zeile              = 3
	DCM_2_0_grammarParserRULE_mod_anf_zeile          = 4
	DCM_2_0_grammarParserRULE_mod_fort_zeile         = 5
	DCM_2_0_grammarParserRULE_mod_ele_name           = 6
	DCM_2_0_grammarParserRULE_mod_ele_wert           = 7
	DCM_2_0_grammarParserRULE_funktionsdef           = 8
	DCM_2_0_grammarParserRULE_funktionszeile         = 9
	DCM_2_0_grammarParserRULE_fkt_version            = 10
	DCM_2_0_grammarParserRULE_fkt_langname           = 11
	DCM_2_0_grammarParserRULE_variantendef           = 12
	DCM_2_0_grammarParserRULE_variantenkrit          = 13
	DCM_2_0_grammarParserRULE_krit_name              = 14
	DCM_2_0_grammarParserRULE_krit_wert              = 15
	DCM_2_0_grammarParserRULE_kons_rumpf             = 16
	DCM_2_0_grammarParserRULE_kenngroesse            = 17
	DCM_2_0_grammarParserRULE_kennwert               = 18
	DCM_2_0_grammarParserRULE_kennwerteblock         = 19
	DCM_2_0_grammarParserRULE_kennlinie              = 20
	DCM_2_0_grammarParserRULE_kennfeld               = 21
	DCM_2_0_grammarParserRULE_gruppenstuetzstellen   = 22
	DCM_2_0_grammarParserRULE_kenntext               = 23
	DCM_2_0_grammarParserRULE_kgr_info               = 24
	DCM_2_0_grammarParserRULE_einheit_x              = 25
	DCM_2_0_grammarParserRULE_einheit_y              = 26
	DCM_2_0_grammarParserRULE_einheit_w              = 27
	DCM_2_0_grammarParserRULE_langname               = 28
	DCM_2_0_grammarParserRULE_displayname            = 29
	DCM_2_0_grammarParserRULE_var_abhangigkeiten     = 30
	DCM_2_0_grammarParserRULE_var_abh                = 31
	DCM_2_0_grammarParserRULE_funktionszugehorigkeit = 32
	DCM_2_0_grammarParserRULE_anzahl_x               = 33
	DCM_2_0_grammarParserRULE_anzahl_y               = 34
	DCM_2_0_grammarParserRULE_werteliste             = 35
	DCM_2_0_grammarParserRULE_werteliste_kwb         = 36
	DCM_2_0_grammarParserRULE_sst_liste_x            = 37
	DCM_2_0_grammarParserRULE_kf_zeile_liste         = 38
	DCM_2_0_grammarParserRULE_kf_zeile_liste_r       = 39
	DCM_2_0_grammarParserRULE_kf_zeile_liste_tx      = 40
	DCM_2_0_grammarParserRULE_realzahl               = 41
)

// IKonservierungContext is an interface to support dynamic dispatch.
type IKonservierungContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKonservierungContext differentiates from other interfaces.
	IsKonservierungContext()
}

type KonservierungContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKonservierungContext() *KonservierungContext {
	var p = new(KonservierungContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_konservierung
	return p
}

func (*KonservierungContext) IsKonservierungContext() {}

func NewKonservierungContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KonservierungContext {
	var p = new(KonservierungContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_konservierung

	return p
}

func (s *KonservierungContext) GetParser() antlr.Parser { return s.parser }

func (s *KonservierungContext) Kons_kopf() IKons_kopfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKons_kopfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKons_kopfContext)
}

func (s *KonservierungContext) Kons_rumpf() IKons_rumpfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKons_rumpfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKons_rumpfContext)
}

func (s *KonservierungContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KonservierungContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KonservierungContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKonservierung(s)
	}
}

func (s *KonservierungContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKonservierung(s)
	}
}

func (p *DCM_2_0_grammarParser) Konservierung() (localctx IKonservierungContext) {
	this := p
	_ = this

	localctx = NewKonservierungContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DCM_2_0_grammarParserRULE_konservierung)
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
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DCM_2_0_grammarParserT__0 {
		{
			p.SetState(84)
			p.Match(DCM_2_0_grammarParserT__0)
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(DCM_2_0_grammarParserT__1)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
		{
			p.SetState(91)
			p.Match(DCM_2_0_grammarParserT__0)
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(96)
		p.Kons_kopf()
	}
	{
		p.SetState(97)
		p.Kons_rumpf()
	}

	return localctx
}

// IKons_kopfContext is an interface to support dynamic dispatch.
type IKons_kopfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKons_kopfContext differentiates from other interfaces.
	IsKons_kopfContext()
}

type Kons_kopfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKons_kopfContext() *Kons_kopfContext {
	var p = new(Kons_kopfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_kons_kopf
	return p
}

func (*Kons_kopfContext) IsKons_kopfContext() {}

func NewKons_kopfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kons_kopfContext {
	var p = new(Kons_kopfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_kons_kopf

	return p
}

func (s *Kons_kopfContext) GetParser() antlr.Parser { return s.parser }

func (s *Kons_kopfContext) Modulkopf_info() IModulkopf_infoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModulkopf_infoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModulkopf_infoContext)
}

func (s *Kons_kopfContext) Funktionsdef() IFunktionsdefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunktionsdefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunktionsdefContext)
}

func (s *Kons_kopfContext) Variantendef() IVariantendefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariantendefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariantendefContext)
}

func (s *Kons_kopfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kons_kopfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kons_kopfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKons_kopf(s)
	}
}

func (s *Kons_kopfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKons_kopf(s)
	}
}

func (p *DCM_2_0_grammarParser) Kons_kopf() (localctx IKons_kopfContext) {
	this := p
	_ = this

	localctx = NewKons_kopfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DCM_2_0_grammarParserRULE_kons_kopf)
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
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DCM_2_0_grammarParserT__2 {
		{
			p.SetState(99)
			p.Modulkopf_info()
		}

	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DCM_2_0_grammarParserT__3 {
		{
			p.SetState(102)
			p.Funktionsdef()
		}

	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DCM_2_0_grammarParserT__6 {
		{
			p.SetState(105)
			p.Variantendef()
		}

	}

	return localctx
}

// IModulkopf_infoContext is an interface to support dynamic dispatch.
type IModulkopf_infoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModulkopf_infoContext differentiates from other interfaces.
	IsModulkopf_infoContext()
}

type Modulkopf_infoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModulkopf_infoContext() *Modulkopf_infoContext {
	var p = new(Modulkopf_infoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_modulkopf_info
	return p
}

func (*Modulkopf_infoContext) IsModulkopf_infoContext() {}

func NewModulkopf_infoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Modulkopf_infoContext {
	var p = new(Modulkopf_infoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_modulkopf_info

	return p
}

func (s *Modulkopf_infoContext) GetParser() antlr.Parser { return s.parser }

func (s *Modulkopf_infoContext) AllMod_zeile() []IMod_zeileContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMod_zeileContext)(nil)).Elem())
	var tst = make([]IMod_zeileContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMod_zeileContext)
		}
	}

	return tst
}

func (s *Modulkopf_infoContext) Mod_zeile(i int) IMod_zeileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMod_zeileContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMod_zeileContext)
}

func (s *Modulkopf_infoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Modulkopf_infoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Modulkopf_infoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterModulkopf_info(s)
	}
}

func (s *Modulkopf_infoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitModulkopf_info(s)
	}
}

func (p *DCM_2_0_grammarParser) Modulkopf_info() (localctx IModulkopf_infoContext) {
	this := p
	_ = this

	localctx = NewModulkopf_infoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DCM_2_0_grammarParserRULE_modulkopf_info)
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
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__2 {
		{
			p.SetState(108)
			p.Mod_zeile()
		}

		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMod_zeileContext is an interface to support dynamic dispatch.
type IMod_zeileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMod_zeileContext differentiates from other interfaces.
	IsMod_zeileContext()
}

type Mod_zeileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMod_zeileContext() *Mod_zeileContext {
	var p = new(Mod_zeileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_mod_zeile
	return p
}

func (*Mod_zeileContext) IsMod_zeileContext() {}

func NewMod_zeileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mod_zeileContext {
	var p = new(Mod_zeileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_mod_zeile

	return p
}

func (s *Mod_zeileContext) GetParser() antlr.Parser { return s.parser }

func (s *Mod_zeileContext) Mod_anf_zeile() IMod_anf_zeileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMod_anf_zeileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMod_anf_zeileContext)
}

func (s *Mod_zeileContext) AllMod_fort_zeile() []IMod_fort_zeileContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMod_fort_zeileContext)(nil)).Elem())
	var tst = make([]IMod_fort_zeileContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMod_fort_zeileContext)
		}
	}

	return tst
}

func (s *Mod_zeileContext) Mod_fort_zeile(i int) IMod_fort_zeileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMod_fort_zeileContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMod_fort_zeileContext)
}

func (s *Mod_zeileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mod_zeileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mod_zeileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterMod_zeile(s)
	}
}

func (s *Mod_zeileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitMod_zeile(s)
	}
}

func (p *DCM_2_0_grammarParser) Mod_zeile() (localctx IMod_zeileContext) {
	this := p
	_ = this

	localctx = NewMod_zeileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DCM_2_0_grammarParserRULE_mod_zeile)

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
		p.SetState(113)
		p.Mod_anf_zeile()
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(114)
				p.Mod_fort_zeile()
			}

		}
		p.SetState(119)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IMod_anf_zeileContext is an interface to support dynamic dispatch.
type IMod_anf_zeileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMod_anf_zeileContext differentiates from other interfaces.
	IsMod_anf_zeileContext()
}

type Mod_anf_zeileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMod_anf_zeileContext() *Mod_anf_zeileContext {
	var p = new(Mod_anf_zeileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_mod_anf_zeile
	return p
}

func (*Mod_anf_zeileContext) IsMod_anf_zeileContext() {}

func NewMod_anf_zeileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mod_anf_zeileContext {
	var p = new(Mod_anf_zeileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_mod_anf_zeile

	return p
}

func (s *Mod_anf_zeileContext) GetParser() antlr.Parser { return s.parser }

func (s *Mod_anf_zeileContext) Mod_ele_name() IMod_ele_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMod_ele_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMod_ele_nameContext)
}

func (s *Mod_anf_zeileContext) Mod_ele_wert() IMod_ele_wertContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMod_ele_wertContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMod_ele_wertContext)
}

func (s *Mod_anf_zeileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mod_anf_zeileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mod_anf_zeileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterMod_anf_zeile(s)
	}
}

func (s *Mod_anf_zeileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitMod_anf_zeile(s)
	}
}

func (p *DCM_2_0_grammarParser) Mod_anf_zeile() (localctx IMod_anf_zeileContext) {
	this := p
	_ = this

	localctx = NewMod_anf_zeileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DCM_2_0_grammarParserRULE_mod_anf_zeile)

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
		p.SetState(120)
		p.Match(DCM_2_0_grammarParserT__2)
	}
	{
		p.SetState(121)
		p.Mod_ele_name()
	}
	{
		p.SetState(122)
		p.Mod_ele_wert()
	}

	return localctx
}

// IMod_fort_zeileContext is an interface to support dynamic dispatch.
type IMod_fort_zeileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMod_fort_zeileContext differentiates from other interfaces.
	IsMod_fort_zeileContext()
}

type Mod_fort_zeileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMod_fort_zeileContext() *Mod_fort_zeileContext {
	var p = new(Mod_fort_zeileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_mod_fort_zeile
	return p
}

func (*Mod_fort_zeileContext) IsMod_fort_zeileContext() {}

func NewMod_fort_zeileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mod_fort_zeileContext {
	var p = new(Mod_fort_zeileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_mod_fort_zeile

	return p
}

func (s *Mod_fort_zeileContext) GetParser() antlr.Parser { return s.parser }

func (s *Mod_fort_zeileContext) Mod_ele_wert() IMod_ele_wertContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMod_ele_wertContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMod_ele_wertContext)
}

func (s *Mod_fort_zeileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mod_fort_zeileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mod_fort_zeileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterMod_fort_zeile(s)
	}
}

func (s *Mod_fort_zeileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitMod_fort_zeile(s)
	}
}

func (p *DCM_2_0_grammarParser) Mod_fort_zeile() (localctx IMod_fort_zeileContext) {
	this := p
	_ = this

	localctx = NewMod_fort_zeileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DCM_2_0_grammarParserRULE_mod_fort_zeile)

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
		p.SetState(124)
		p.Match(DCM_2_0_grammarParserT__2)
	}
	{
		p.SetState(125)
		p.Mod_ele_wert()
	}

	return localctx
}

// IMod_ele_nameContext is an interface to support dynamic dispatch.
type IMod_ele_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMod_ele_nameContext differentiates from other interfaces.
	IsMod_ele_nameContext()
}

type Mod_ele_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMod_ele_nameContext() *Mod_ele_nameContext {
	var p = new(Mod_ele_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_mod_ele_name
	return p
}

func (*Mod_ele_nameContext) IsMod_ele_nameContext() {}

func NewMod_ele_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mod_ele_nameContext {
	var p = new(Mod_ele_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_mod_ele_name

	return p
}

func (s *Mod_ele_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Mod_ele_nameContext) NAME() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, 0)
}

func (s *Mod_ele_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mod_ele_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mod_ele_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterMod_ele_name(s)
	}
}

func (s *Mod_ele_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitMod_ele_name(s)
	}
}

func (p *DCM_2_0_grammarParser) Mod_ele_name() (localctx IMod_ele_nameContext) {
	this := p
	_ = this

	localctx = NewMod_ele_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, DCM_2_0_grammarParserRULE_mod_ele_name)

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
		p.SetState(127)
		p.Match(DCM_2_0_grammarParserNAME)
	}

	return localctx
}

// IMod_ele_wertContext is an interface to support dynamic dispatch.
type IMod_ele_wertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMod_ele_wertContext differentiates from other interfaces.
	IsMod_ele_wertContext()
}

type Mod_ele_wertContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMod_ele_wertContext() *Mod_ele_wertContext {
	var p = new(Mod_ele_wertContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_mod_ele_wert
	return p
}

func (*Mod_ele_wertContext) IsMod_ele_wertContext() {}

func NewMod_ele_wertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mod_ele_wertContext {
	var p = new(Mod_ele_wertContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_mod_ele_wert

	return p
}

func (s *Mod_ele_wertContext) GetParser() antlr.Parser { return s.parser }

func (s *Mod_ele_wertContext) TEXT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, 0)
}

func (s *Mod_ele_wertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mod_ele_wertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mod_ele_wertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterMod_ele_wert(s)
	}
}

func (s *Mod_ele_wertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitMod_ele_wert(s)
	}
}

func (p *DCM_2_0_grammarParser) Mod_ele_wert() (localctx IMod_ele_wertContext) {
	this := p
	_ = this

	localctx = NewMod_ele_wertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, DCM_2_0_grammarParserRULE_mod_ele_wert)
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
		p.SetState(129)
		p.Match(DCM_2_0_grammarParserTEXT)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
		{
			p.SetState(130)
			p.Match(DCM_2_0_grammarParserT__0)
		}

		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunktionsdefContext is an interface to support dynamic dispatch.
type IFunktionsdefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunktionsdefContext differentiates from other interfaces.
	IsFunktionsdefContext()
}

type FunktionsdefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunktionsdefContext() *FunktionsdefContext {
	var p = new(FunktionsdefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_funktionsdef
	return p
}

func (*FunktionsdefContext) IsFunktionsdefContext() {}

func NewFunktionsdefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunktionsdefContext {
	var p = new(FunktionsdefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_funktionsdef

	return p
}

func (s *FunktionsdefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunktionsdefContext) AllFunktionszeile() []IFunktionszeileContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunktionszeileContext)(nil)).Elem())
	var tst = make([]IFunktionszeileContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunktionszeileContext)
		}
	}

	return tst
}

func (s *FunktionsdefContext) Funktionszeile(i int) IFunktionszeileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunktionszeileContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunktionszeileContext)
}

func (s *FunktionsdefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunktionsdefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunktionsdefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterFunktionsdef(s)
	}
}

func (s *FunktionsdefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitFunktionsdef(s)
	}
}

func (p *DCM_2_0_grammarParser) Funktionsdef() (localctx IFunktionsdefContext) {
	this := p
	_ = this

	localctx = NewFunktionsdefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DCM_2_0_grammarParserRULE_funktionsdef)
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
		p.Match(DCM_2_0_grammarParserT__3)
	}
	{
		p.SetState(136)
		p.Match(DCM_2_0_grammarParserT__0)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__5 {
		{
			p.SetState(137)
			p.Funktionszeile()
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(142)
		p.Match(DCM_2_0_grammarParserT__4)
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
		{
			p.SetState(143)
			p.Match(DCM_2_0_grammarParserT__0)
		}

		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunktionszeileContext is an interface to support dynamic dispatch.
type IFunktionszeileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunktionszeileContext differentiates from other interfaces.
	IsFunktionszeileContext()
}

type FunktionszeileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunktionszeileContext() *FunktionszeileContext {
	var p = new(FunktionszeileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_funktionszeile
	return p
}

func (*FunktionszeileContext) IsFunktionszeileContext() {}

func NewFunktionszeileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunktionszeileContext {
	var p = new(FunktionszeileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_funktionszeile

	return p
}

func (s *FunktionszeileContext) GetParser() antlr.Parser { return s.parser }

func (s *FunktionszeileContext) NAME() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, 0)
}

func (s *FunktionszeileContext) Fkt_version() IFkt_versionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFkt_versionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFkt_versionContext)
}

func (s *FunktionszeileContext) Fkt_langname() IFkt_langnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFkt_langnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFkt_langnameContext)
}

func (s *FunktionszeileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunktionszeileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunktionszeileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterFunktionszeile(s)
	}
}

func (s *FunktionszeileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitFunktionszeile(s)
	}
}

func (p *DCM_2_0_grammarParser) Funktionszeile() (localctx IFunktionszeileContext) {
	this := p
	_ = this

	localctx = NewFunktionszeileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DCM_2_0_grammarParserRULE_funktionszeile)

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
		p.SetState(148)
		p.Match(DCM_2_0_grammarParserT__5)
	}
	{
		p.SetState(149)
		p.Match(DCM_2_0_grammarParserNAME)
	}
	{
		p.SetState(150)
		p.Fkt_version()
	}
	{
		p.SetState(151)
		p.Fkt_langname()
	}

	return localctx
}

// IFkt_versionContext is an interface to support dynamic dispatch.
type IFkt_versionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFkt_versionContext differentiates from other interfaces.
	IsFkt_versionContext()
}

type Fkt_versionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFkt_versionContext() *Fkt_versionContext {
	var p = new(Fkt_versionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_fkt_version
	return p
}

func (*Fkt_versionContext) IsFkt_versionContext() {}

func NewFkt_versionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fkt_versionContext {
	var p = new(Fkt_versionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_fkt_version

	return p
}

func (s *Fkt_versionContext) GetParser() antlr.Parser { return s.parser }

func (s *Fkt_versionContext) TEXT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, 0)
}

func (s *Fkt_versionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fkt_versionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fkt_versionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterFkt_version(s)
	}
}

func (s *Fkt_versionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitFkt_version(s)
	}
}

func (p *DCM_2_0_grammarParser) Fkt_version() (localctx IFkt_versionContext) {
	this := p
	_ = this

	localctx = NewFkt_versionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DCM_2_0_grammarParserRULE_fkt_version)

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
		p.SetState(153)
		p.Match(DCM_2_0_grammarParserTEXT)
	}

	return localctx
}

// IFkt_langnameContext is an interface to support dynamic dispatch.
type IFkt_langnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFkt_langnameContext differentiates from other interfaces.
	IsFkt_langnameContext()
}

type Fkt_langnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFkt_langnameContext() *Fkt_langnameContext {
	var p = new(Fkt_langnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_fkt_langname
	return p
}

func (*Fkt_langnameContext) IsFkt_langnameContext() {}

func NewFkt_langnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fkt_langnameContext {
	var p = new(Fkt_langnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_fkt_langname

	return p
}

func (s *Fkt_langnameContext) GetParser() antlr.Parser { return s.parser }

func (s *Fkt_langnameContext) TEXT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, 0)
}

func (s *Fkt_langnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fkt_langnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fkt_langnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterFkt_langname(s)
	}
}

func (s *Fkt_langnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitFkt_langname(s)
	}
}

func (p *DCM_2_0_grammarParser) Fkt_langname() (localctx IFkt_langnameContext) {
	this := p
	_ = this

	localctx = NewFkt_langnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DCM_2_0_grammarParserRULE_fkt_langname)
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
		p.Match(DCM_2_0_grammarParserTEXT)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
		{
			p.SetState(156)
			p.Match(DCM_2_0_grammarParserT__0)
		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariantendefContext is an interface to support dynamic dispatch.
type IVariantendefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariantendefContext differentiates from other interfaces.
	IsVariantendefContext()
}

type VariantendefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariantendefContext() *VariantendefContext {
	var p = new(VariantendefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_variantendef
	return p
}

func (*VariantendefContext) IsVariantendefContext() {}

func NewVariantendefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariantendefContext {
	var p = new(VariantendefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_variantendef

	return p
}

func (s *VariantendefContext) GetParser() antlr.Parser { return s.parser }

func (s *VariantendefContext) AllVariantenkrit() []IVariantenkritContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariantenkritContext)(nil)).Elem())
	var tst = make([]IVariantenkritContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariantenkritContext)
		}
	}

	return tst
}

func (s *VariantendefContext) Variantenkrit(i int) IVariantenkritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariantenkritContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariantenkritContext)
}

func (s *VariantendefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariantendefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariantendefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterVariantendef(s)
	}
}

func (s *VariantendefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitVariantendef(s)
	}
}

func (p *DCM_2_0_grammarParser) Variantendef() (localctx IVariantendefContext) {
	this := p
	_ = this

	localctx = NewVariantendefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DCM_2_0_grammarParserRULE_variantendef)
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
		p.SetState(161)
		p.Match(DCM_2_0_grammarParserT__6)
	}
	{
		p.SetState(162)
		p.Match(DCM_2_0_grammarParserT__0)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__7 {
		{
			p.SetState(163)
			p.Variantenkrit()
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(168)
		p.Match(DCM_2_0_grammarParserT__4)
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
		{
			p.SetState(169)
			p.Match(DCM_2_0_grammarParserT__0)
		}

		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariantenkritContext is an interface to support dynamic dispatch.
type IVariantenkritContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariantenkritContext differentiates from other interfaces.
	IsVariantenkritContext()
}

type VariantenkritContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariantenkritContext() *VariantenkritContext {
	var p = new(VariantenkritContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_variantenkrit
	return p
}

func (*VariantenkritContext) IsVariantenkritContext() {}

func NewVariantenkritContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariantenkritContext {
	var p = new(VariantenkritContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_variantenkrit

	return p
}

func (s *VariantenkritContext) GetParser() antlr.Parser { return s.parser }

func (s *VariantenkritContext) Krit_name() IKrit_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKrit_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKrit_nameContext)
}

func (s *VariantenkritContext) AllKrit_wert() []IKrit_wertContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKrit_wertContext)(nil)).Elem())
	var tst = make([]IKrit_wertContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKrit_wertContext)
		}
	}

	return tst
}

func (s *VariantenkritContext) Krit_wert(i int) IKrit_wertContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKrit_wertContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKrit_wertContext)
}

func (s *VariantenkritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariantenkritContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariantenkritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterVariantenkrit(s)
	}
}

func (s *VariantenkritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitVariantenkrit(s)
	}
}

func (p *DCM_2_0_grammarParser) Variantenkrit() (localctx IVariantenkritContext) {
	this := p
	_ = this

	localctx = NewVariantenkritContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DCM_2_0_grammarParserRULE_variantenkrit)
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
		p.SetState(174)
		p.Match(DCM_2_0_grammarParserT__7)
	}
	{
		p.SetState(175)
		p.Krit_name()
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DCM_2_0_grammarParserNAME {
		{
			p.SetState(176)
			p.Krit_wert()
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
		{
			p.SetState(182)
			p.Match(DCM_2_0_grammarParserT__0)
		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IKrit_nameContext is an interface to support dynamic dispatch.
type IKrit_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKrit_nameContext differentiates from other interfaces.
	IsKrit_nameContext()
}

type Krit_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKrit_nameContext() *Krit_nameContext {
	var p = new(Krit_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_krit_name
	return p
}

func (*Krit_nameContext) IsKrit_nameContext() {}

func NewKrit_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Krit_nameContext {
	var p = new(Krit_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_krit_name

	return p
}

func (s *Krit_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Krit_nameContext) NAME() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, 0)
}

func (s *Krit_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Krit_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Krit_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKrit_name(s)
	}
}

func (s *Krit_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKrit_name(s)
	}
}

func (p *DCM_2_0_grammarParser) Krit_name() (localctx IKrit_nameContext) {
	this := p
	_ = this

	localctx = NewKrit_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, DCM_2_0_grammarParserRULE_krit_name)

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
		p.SetState(187)
		p.Match(DCM_2_0_grammarParserNAME)
	}

	return localctx
}

// IKrit_wertContext is an interface to support dynamic dispatch.
type IKrit_wertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKrit_wertContext differentiates from other interfaces.
	IsKrit_wertContext()
}

type Krit_wertContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKrit_wertContext() *Krit_wertContext {
	var p = new(Krit_wertContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_krit_wert
	return p
}

func (*Krit_wertContext) IsKrit_wertContext() {}

func NewKrit_wertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Krit_wertContext {
	var p = new(Krit_wertContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_krit_wert

	return p
}

func (s *Krit_wertContext) GetParser() antlr.Parser { return s.parser }

func (s *Krit_wertContext) NAME() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, 0)
}

func (s *Krit_wertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Krit_wertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Krit_wertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKrit_wert(s)
	}
}

func (s *Krit_wertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKrit_wert(s)
	}
}

func (p *DCM_2_0_grammarParser) Krit_wert() (localctx IKrit_wertContext) {
	this := p
	_ = this

	localctx = NewKrit_wertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DCM_2_0_grammarParserRULE_krit_wert)

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
		p.Match(DCM_2_0_grammarParserNAME)
	}

	return localctx
}

// IKons_rumpfContext is an interface to support dynamic dispatch.
type IKons_rumpfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKons_rumpfContext differentiates from other interfaces.
	IsKons_rumpfContext()
}

type Kons_rumpfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKons_rumpfContext() *Kons_rumpfContext {
	var p = new(Kons_rumpfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_kons_rumpf
	return p
}

func (*Kons_rumpfContext) IsKons_rumpfContext() {}

func NewKons_rumpfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kons_rumpfContext {
	var p = new(Kons_rumpfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_kons_rumpf

	return p
}

func (s *Kons_rumpfContext) GetParser() antlr.Parser { return s.parser }

func (s *Kons_rumpfContext) AllKenngroesse() []IKenngroesseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKenngroesseContext)(nil)).Elem())
	var tst = make([]IKenngroesseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKenngroesseContext)
		}
	}

	return tst
}

func (s *Kons_rumpfContext) Kenngroesse(i int) IKenngroesseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKenngroesseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKenngroesseContext)
}

func (s *Kons_rumpfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kons_rumpfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kons_rumpfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKons_rumpf(s)
	}
}

func (s *Kons_rumpfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKons_rumpf(s)
	}
}

func (p *DCM_2_0_grammarParser) Kons_rumpf() (localctx IKons_rumpfContext) {
	this := p
	_ = this

	localctx = NewKons_rumpfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DCM_2_0_grammarParserRULE_kons_rumpf)
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
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DCM_2_0_grammarParserT__8)|(1<<DCM_2_0_grammarParserT__11)|(1<<DCM_2_0_grammarParserT__12)|(1<<DCM_2_0_grammarParserT__13)|(1<<DCM_2_0_grammarParserT__14)|(1<<DCM_2_0_grammarParserT__15)|(1<<DCM_2_0_grammarParserT__16)|(1<<DCM_2_0_grammarParserT__17)|(1<<DCM_2_0_grammarParserT__18)|(1<<DCM_2_0_grammarParserT__19))) != 0 {
		{
			p.SetState(191)
			p.Kenngroesse()
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IKenngroesseContext is an interface to support dynamic dispatch.
type IKenngroesseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKenngroesseContext differentiates from other interfaces.
	IsKenngroesseContext()
}

type KenngroesseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKenngroesseContext() *KenngroesseContext {
	var p = new(KenngroesseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_kenngroesse
	return p
}

func (*KenngroesseContext) IsKenngroesseContext() {}

func NewKenngroesseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KenngroesseContext {
	var p = new(KenngroesseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_kenngroesse

	return p
}

func (s *KenngroesseContext) GetParser() antlr.Parser { return s.parser }

func (s *KenngroesseContext) Kennwert() IKennwertContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKennwertContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKennwertContext)
}

func (s *KenngroesseContext) Kennwerteblock() IKennwerteblockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKennwerteblockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKennwerteblockContext)
}

func (s *KenngroesseContext) Kennlinie() IKennlinieContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKennlinieContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKennlinieContext)
}

func (s *KenngroesseContext) Kennfeld() IKennfeldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKennfeldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKennfeldContext)
}

func (s *KenngroesseContext) Gruppenstuetzstellen() IGruppenstuetzstellenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGruppenstuetzstellenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGruppenstuetzstellenContext)
}

func (s *KenngroesseContext) Kenntext() IKenntextContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKenntextContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKenntextContext)
}

func (s *KenngroesseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KenngroesseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KenngroesseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKenngroesse(s)
	}
}

func (s *KenngroesseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKenngroesse(s)
	}
}

func (p *DCM_2_0_grammarParser) Kenngroesse() (localctx IKenngroesseContext) {
	this := p
	_ = this

	localctx = NewKenngroesseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DCM_2_0_grammarParserRULE_kenngroesse)

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
	p.SetState(203)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DCM_2_0_grammarParserT__8:
		{
			p.SetState(197)
			p.Kennwert()
		}

	case DCM_2_0_grammarParserT__11:
		{
			p.SetState(198)
			p.Kennwerteblock()
		}

	case DCM_2_0_grammarParserT__12, DCM_2_0_grammarParserT__13, DCM_2_0_grammarParserT__14:
		{
			p.SetState(199)
			p.Kennlinie()
		}

	case DCM_2_0_grammarParserT__15, DCM_2_0_grammarParserT__16, DCM_2_0_grammarParserT__17:
		{
			p.SetState(200)
			p.Kennfeld()
		}

	case DCM_2_0_grammarParserT__18:
		{
			p.SetState(201)
			p.Gruppenstuetzstellen()
		}

	case DCM_2_0_grammarParserT__19:
		{
			p.SetState(202)
			p.Kenntext()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKennwertContext is an interface to support dynamic dispatch.
type IKennwertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKennwertContext differentiates from other interfaces.
	IsKennwertContext()
}

type KennwertContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKennwertContext() *KennwertContext {
	var p = new(KennwertContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_kennwert
	return p
}

func (*KennwertContext) IsKennwertContext() {}

func NewKennwertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KennwertContext {
	var p = new(KennwertContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_kennwert

	return p
}

func (s *KennwertContext) GetParser() antlr.Parser { return s.parser }

func (s *KennwertContext) NAME() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, 0)
}

func (s *KennwertContext) Kgr_info() IKgr_infoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKgr_infoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKgr_infoContext)
}

func (s *KennwertContext) Realzahl() IRealzahlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealzahlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealzahlContext)
}

func (s *KennwertContext) Einheit_w() IEinheit_wContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEinheit_wContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEinheit_wContext)
}

func (s *KennwertContext) TEXT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, 0)
}

func (s *KennwertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KennwertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KennwertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKennwert(s)
	}
}

func (s *KennwertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKennwert(s)
	}
}

func (p *DCM_2_0_grammarParser) Kennwert() (localctx IKennwertContext) {
	this := p
	_ = this

	localctx = NewKennwertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, DCM_2_0_grammarParserRULE_kennwert)
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

	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(205)
			p.Match(DCM_2_0_grammarParserT__8)
		}
		{
			p.SetState(206)
			p.Match(DCM_2_0_grammarParserNAME)
		}
		{
			p.SetState(207)
			p.Match(DCM_2_0_grammarParserT__0)
		}
		{
			p.SetState(208)
			p.Kgr_info()
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__22 {
			{
				p.SetState(209)
				p.Einheit_w()
			}

		}
		{
			p.SetState(212)
			p.Match(DCM_2_0_grammarParserT__9)
		}
		{
			p.SetState(213)
			p.Realzahl()
		}
		{
			p.SetState(214)
			p.Match(DCM_2_0_grammarParserT__0)
		}
		{
			p.SetState(215)
			p.Match(DCM_2_0_grammarParserT__4)
		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
			{
				p.SetState(216)
				p.Match(DCM_2_0_grammarParserT__0)
			}

			p.SetState(219)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(221)
			p.Match(DCM_2_0_grammarParserT__8)
		}
		{
			p.SetState(222)
			p.Match(DCM_2_0_grammarParserNAME)
		}
		{
			p.SetState(223)
			p.Match(DCM_2_0_grammarParserT__0)
		}
		{
			p.SetState(224)
			p.Kgr_info()
		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__22 {
			{
				p.SetState(225)
				p.Einheit_w()
			}

		}
		{
			p.SetState(228)
			p.Match(DCM_2_0_grammarParserT__10)
		}
		{
			p.SetState(229)
			p.Match(DCM_2_0_grammarParserTEXT)
		}
		{
			p.SetState(230)
			p.Match(DCM_2_0_grammarParserT__0)
		}
		{
			p.SetState(231)
			p.Match(DCM_2_0_grammarParserT__4)
		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
			{
				p.SetState(232)
				p.Match(DCM_2_0_grammarParserT__0)
			}

			p.SetState(235)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IKennwerteblockContext is an interface to support dynamic dispatch.
type IKennwerteblockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKennwerteblockContext differentiates from other interfaces.
	IsKennwerteblockContext()
}

type KennwerteblockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKennwerteblockContext() *KennwerteblockContext {
	var p = new(KennwerteblockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_kennwerteblock
	return p
}

func (*KennwerteblockContext) IsKennwerteblockContext() {}

func NewKennwerteblockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KennwerteblockContext {
	var p = new(KennwerteblockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_kennwerteblock

	return p
}

func (s *KennwerteblockContext) GetParser() antlr.Parser { return s.parser }

func (s *KennwerteblockContext) NAME() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, 0)
}

func (s *KennwerteblockContext) Anzahl_x() IAnzahl_xContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnzahl_xContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnzahl_xContext)
}

func (s *KennwerteblockContext) Kgr_info() IKgr_infoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKgr_infoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKgr_infoContext)
}

func (s *KennwerteblockContext) Einheit_w() IEinheit_wContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEinheit_wContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEinheit_wContext)
}

func (s *KennwerteblockContext) AllWerteliste_kwb() []IWerteliste_kwbContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWerteliste_kwbContext)(nil)).Elem())
	var tst = make([]IWerteliste_kwbContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWerteliste_kwbContext)
		}
	}

	return tst
}

func (s *KennwerteblockContext) Werteliste_kwb(i int) IWerteliste_kwbContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWerteliste_kwbContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWerteliste_kwbContext)
}

func (s *KennwerteblockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KennwerteblockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KennwerteblockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKennwerteblock(s)
	}
}

func (s *KennwerteblockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKennwerteblock(s)
	}
}

func (p *DCM_2_0_grammarParser) Kennwerteblock() (localctx IKennwerteblockContext) {
	this := p
	_ = this

	localctx = NewKennwerteblockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, DCM_2_0_grammarParserRULE_kennwerteblock)
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
		p.SetState(239)
		p.Match(DCM_2_0_grammarParserT__11)
	}
	{
		p.SetState(240)
		p.Match(DCM_2_0_grammarParserNAME)
	}
	{
		p.SetState(241)
		p.Anzahl_x()
	}
	{
		p.SetState(242)
		p.Match(DCM_2_0_grammarParserT__0)
	}
	{
		p.SetState(243)
		p.Kgr_info()
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DCM_2_0_grammarParserT__22 {
		{
			p.SetState(244)
			p.Einheit_w()
		}

	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__9 || _la == DCM_2_0_grammarParserT__10 {
		{
			p.SetState(247)
			p.Werteliste_kwb()
		}

		p.SetState(250)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(252)
		p.Match(DCM_2_0_grammarParserT__4)
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
		{
			p.SetState(253)
			p.Match(DCM_2_0_grammarParserT__0)
		}

		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IKennlinieContext is an interface to support dynamic dispatch.
type IKennlinieContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKennlinieContext differentiates from other interfaces.
	IsKennlinieContext()
}

type KennlinieContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKennlinieContext() *KennlinieContext {
	var p = new(KennlinieContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_kennlinie
	return p
}

func (*KennlinieContext) IsKennlinieContext() {}

func NewKennlinieContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KennlinieContext {
	var p = new(KennlinieContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_kennlinie

	return p
}

func (s *KennlinieContext) GetParser() antlr.Parser { return s.parser }

func (s *KennlinieContext) NAME() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, 0)
}

func (s *KennlinieContext) Anzahl_x() IAnzahl_xContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnzahl_xContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnzahl_xContext)
}

func (s *KennlinieContext) Kgr_info() IKgr_infoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKgr_infoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKgr_infoContext)
}

func (s *KennlinieContext) Einheit_x() IEinheit_xContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEinheit_xContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEinheit_xContext)
}

func (s *KennlinieContext) Einheit_w() IEinheit_wContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEinheit_wContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEinheit_wContext)
}

func (s *KennlinieContext) AllSst_liste_x() []ISst_liste_xContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISst_liste_xContext)(nil)).Elem())
	var tst = make([]ISst_liste_xContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISst_liste_xContext)
		}
	}

	return tst
}

func (s *KennlinieContext) Sst_liste_x(i int) ISst_liste_xContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISst_liste_xContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISst_liste_xContext)
}

func (s *KennlinieContext) AllWerteliste() []IWertelisteContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWertelisteContext)(nil)).Elem())
	var tst = make([]IWertelisteContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWertelisteContext)
		}
	}

	return tst
}

func (s *KennlinieContext) Werteliste(i int) IWertelisteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWertelisteContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWertelisteContext)
}

func (s *KennlinieContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KennlinieContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KennlinieContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKennlinie(s)
	}
}

func (s *KennlinieContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKennlinie(s)
	}
}

func (p *DCM_2_0_grammarParser) Kennlinie() (localctx IKennlinieContext) {
	this := p
	_ = this

	localctx = NewKennlinieContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, DCM_2_0_grammarParserRULE_kennlinie)
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

	switch p.GetTokenStream().LA(1) {
	case DCM_2_0_grammarParserT__12:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.Match(DCM_2_0_grammarParserT__12)
		}

		{
			p.SetState(259)
			p.Match(DCM_2_0_grammarParserNAME)
		}
		{
			p.SetState(260)
			p.Anzahl_x()
		}
		{
			p.SetState(261)
			p.Match(DCM_2_0_grammarParserT__0)
		}
		{
			p.SetState(262)
			p.Kgr_info()
		}
		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__20 {
			{
				p.SetState(263)
				p.Einheit_x()
			}

		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__22 {
			{
				p.SetState(266)
				p.Einheit_w()
			}

		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__29 || _la == DCM_2_0_grammarParserT__30 {
			{
				p.SetState(269)
				p.Sst_liste_x()
			}

			p.SetState(272)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(275)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__9 {
			{
				p.SetState(274)
				p.Werteliste()
			}

			p.SetState(277)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(279)
			p.Match(DCM_2_0_grammarParserT__4)
		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
			{
				p.SetState(280)
				p.Match(DCM_2_0_grammarParserT__0)
			}

			p.SetState(283)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case DCM_2_0_grammarParserT__13:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.Match(DCM_2_0_grammarParserT__13)
		}

		{
			p.SetState(286)
			p.Match(DCM_2_0_grammarParserNAME)
		}
		{
			p.SetState(287)
			p.Anzahl_x()
		}
		{
			p.SetState(288)
			p.Match(DCM_2_0_grammarParserT__0)
		}
		{
			p.SetState(289)
			p.Kgr_info()
		}
		p.SetState(291)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__20 {
			{
				p.SetState(290)
				p.Einheit_x()
			}

		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__22 {
			{
				p.SetState(293)
				p.Einheit_w()
			}

		}
		p.SetState(297)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__29 || _la == DCM_2_0_grammarParserT__30 {
			{
				p.SetState(296)
				p.Sst_liste_x()
			}

			p.SetState(299)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__9 {
			{
				p.SetState(301)
				p.Werteliste()
			}

			p.SetState(304)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(306)
			p.Match(DCM_2_0_grammarParserT__4)
		}
		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
			{
				p.SetState(307)
				p.Match(DCM_2_0_grammarParserT__0)
			}

			p.SetState(310)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case DCM_2_0_grammarParserT__14:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(312)
			p.Match(DCM_2_0_grammarParserT__14)
		}

		{
			p.SetState(313)
			p.Match(DCM_2_0_grammarParserNAME)
		}
		{
			p.SetState(314)
			p.Anzahl_x()
		}
		{
			p.SetState(315)
			p.Match(DCM_2_0_grammarParserT__0)
		}
		{
			p.SetState(316)
			p.Kgr_info()
		}
		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__20 {
			{
				p.SetState(317)
				p.Einheit_x()
			}

		}
		p.SetState(321)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__22 {
			{
				p.SetState(320)
				p.Einheit_w()
			}

		}
		p.SetState(324)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__29 || _la == DCM_2_0_grammarParserT__30 {
			{
				p.SetState(323)
				p.Sst_liste_x()
			}

			p.SetState(326)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__9 {
			{
				p.SetState(328)
				p.Werteliste()
			}

			p.SetState(331)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(333)
			p.Match(DCM_2_0_grammarParserT__4)
		}
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
			{
				p.SetState(334)
				p.Match(DCM_2_0_grammarParserT__0)
			}

			p.SetState(337)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKennfeldContext is an interface to support dynamic dispatch.
type IKennfeldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKennfeldContext differentiates from other interfaces.
	IsKennfeldContext()
}

type KennfeldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKennfeldContext() *KennfeldContext {
	var p = new(KennfeldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_kennfeld
	return p
}

func (*KennfeldContext) IsKennfeldContext() {}

func NewKennfeldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KennfeldContext {
	var p = new(KennfeldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_kennfeld

	return p
}

func (s *KennfeldContext) GetParser() antlr.Parser { return s.parser }

func (s *KennfeldContext) NAME() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, 0)
}

func (s *KennfeldContext) Anzahl_x() IAnzahl_xContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnzahl_xContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnzahl_xContext)
}

func (s *KennfeldContext) Anzahl_y() IAnzahl_yContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnzahl_yContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnzahl_yContext)
}

func (s *KennfeldContext) Kgr_info() IKgr_infoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKgr_infoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKgr_infoContext)
}

func (s *KennfeldContext) Kf_zeile_liste() IKf_zeile_listeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKf_zeile_listeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKf_zeile_listeContext)
}

func (s *KennfeldContext) Einheit_x() IEinheit_xContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEinheit_xContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEinheit_xContext)
}

func (s *KennfeldContext) Einheit_y() IEinheit_yContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEinheit_yContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEinheit_yContext)
}

func (s *KennfeldContext) Einheit_w() IEinheit_wContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEinheit_wContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEinheit_wContext)
}

func (s *KennfeldContext) AllSst_liste_x() []ISst_liste_xContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISst_liste_xContext)(nil)).Elem())
	var tst = make([]ISst_liste_xContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISst_liste_xContext)
		}
	}

	return tst
}

func (s *KennfeldContext) Sst_liste_x(i int) ISst_liste_xContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISst_liste_xContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISst_liste_xContext)
}

func (s *KennfeldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KennfeldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KennfeldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKennfeld(s)
	}
}

func (s *KennfeldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKennfeld(s)
	}
}

func (p *DCM_2_0_grammarParser) Kennfeld() (localctx IKennfeldContext) {
	this := p
	_ = this

	localctx = NewKennfeldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, DCM_2_0_grammarParserRULE_kennfeld)
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

	p.SetState(422)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DCM_2_0_grammarParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(341)
			p.Match(DCM_2_0_grammarParserT__15)
		}

		{
			p.SetState(342)
			p.Match(DCM_2_0_grammarParserNAME)
		}
		{
			p.SetState(343)
			p.Anzahl_x()
		}
		{
			p.SetState(344)
			p.Anzahl_y()
		}
		{
			p.SetState(345)
			p.Match(DCM_2_0_grammarParserT__0)
		}
		{
			p.SetState(346)
			p.Kgr_info()
		}
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__20 {
			{
				p.SetState(347)
				p.Einheit_x()
			}

		}
		p.SetState(351)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__21 {
			{
				p.SetState(350)
				p.Einheit_y()
			}

		}
		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__22 {
			{
				p.SetState(353)
				p.Einheit_w()
			}

		}
		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__29 || _la == DCM_2_0_grammarParserT__30 {
			{
				p.SetState(356)
				p.Sst_liste_x()
			}

			p.SetState(359)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(361)
			p.Kf_zeile_liste()
		}
		{
			p.SetState(362)
			p.Match(DCM_2_0_grammarParserT__4)
		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
			{
				p.SetState(363)
				p.Match(DCM_2_0_grammarParserT__0)
			}

			p.SetState(366)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case DCM_2_0_grammarParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(368)
			p.Match(DCM_2_0_grammarParserT__16)
		}

		{
			p.SetState(369)
			p.Match(DCM_2_0_grammarParserNAME)
		}
		{
			p.SetState(370)
			p.Anzahl_x()
		}
		{
			p.SetState(371)
			p.Anzahl_y()
		}
		{
			p.SetState(372)
			p.Match(DCM_2_0_grammarParserT__0)
		}
		{
			p.SetState(373)
			p.Kgr_info()
		}
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__20 {
			{
				p.SetState(374)
				p.Einheit_x()
			}

		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__21 {
			{
				p.SetState(377)
				p.Einheit_y()
			}

		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__22 {
			{
				p.SetState(380)
				p.Einheit_w()
			}

		}
		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__29 || _la == DCM_2_0_grammarParserT__30 {
			{
				p.SetState(383)
				p.Sst_liste_x()
			}

			p.SetState(386)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(388)
			p.Kf_zeile_liste()
		}
		{
			p.SetState(389)
			p.Match(DCM_2_0_grammarParserT__4)
		}
		p.SetState(391)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
			{
				p.SetState(390)
				p.Match(DCM_2_0_grammarParserT__0)
			}

			p.SetState(393)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case DCM_2_0_grammarParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(395)
			p.Match(DCM_2_0_grammarParserT__17)
		}

		{
			p.SetState(396)
			p.Match(DCM_2_0_grammarParserNAME)
		}
		{
			p.SetState(397)
			p.Anzahl_x()
		}
		{
			p.SetState(398)
			p.Anzahl_y()
		}
		{
			p.SetState(399)
			p.Match(DCM_2_0_grammarParserT__0)
		}
		{
			p.SetState(400)
			p.Kgr_info()
		}
		p.SetState(402)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__20 {
			{
				p.SetState(401)
				p.Einheit_x()
			}

		}
		p.SetState(405)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__21 {
			{
				p.SetState(404)
				p.Einheit_y()
			}

		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == DCM_2_0_grammarParserT__22 {
			{
				p.SetState(407)
				p.Einheit_w()
			}

		}
		p.SetState(411)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__29 || _la == DCM_2_0_grammarParserT__30 {
			{
				p.SetState(410)
				p.Sst_liste_x()
			}

			p.SetState(413)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(415)
			p.Kf_zeile_liste()
		}
		{
			p.SetState(416)
			p.Match(DCM_2_0_grammarParserT__4)
		}
		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
			{
				p.SetState(417)
				p.Match(DCM_2_0_grammarParserT__0)
			}

			p.SetState(420)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGruppenstuetzstellenContext is an interface to support dynamic dispatch.
type IGruppenstuetzstellenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGruppenstuetzstellenContext differentiates from other interfaces.
	IsGruppenstuetzstellenContext()
}

type GruppenstuetzstellenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGruppenstuetzstellenContext() *GruppenstuetzstellenContext {
	var p = new(GruppenstuetzstellenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_gruppenstuetzstellen
	return p
}

func (*GruppenstuetzstellenContext) IsGruppenstuetzstellenContext() {}

func NewGruppenstuetzstellenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GruppenstuetzstellenContext {
	var p = new(GruppenstuetzstellenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_gruppenstuetzstellen

	return p
}

func (s *GruppenstuetzstellenContext) GetParser() antlr.Parser { return s.parser }

func (s *GruppenstuetzstellenContext) NAME() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, 0)
}

func (s *GruppenstuetzstellenContext) Anzahl_x() IAnzahl_xContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnzahl_xContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnzahl_xContext)
}

func (s *GruppenstuetzstellenContext) Kgr_info() IKgr_infoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKgr_infoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKgr_infoContext)
}

func (s *GruppenstuetzstellenContext) Einheit_x() IEinheit_xContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEinheit_xContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEinheit_xContext)
}

func (s *GruppenstuetzstellenContext) AllSst_liste_x() []ISst_liste_xContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISst_liste_xContext)(nil)).Elem())
	var tst = make([]ISst_liste_xContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISst_liste_xContext)
		}
	}

	return tst
}

func (s *GruppenstuetzstellenContext) Sst_liste_x(i int) ISst_liste_xContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISst_liste_xContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISst_liste_xContext)
}

func (s *GruppenstuetzstellenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GruppenstuetzstellenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GruppenstuetzstellenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterGruppenstuetzstellen(s)
	}
}

func (s *GruppenstuetzstellenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitGruppenstuetzstellen(s)
	}
}

func (p *DCM_2_0_grammarParser) Gruppenstuetzstellen() (localctx IGruppenstuetzstellenContext) {
	this := p
	_ = this

	localctx = NewGruppenstuetzstellenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, DCM_2_0_grammarParserRULE_gruppenstuetzstellen)
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
		p.SetState(424)
		p.Match(DCM_2_0_grammarParserT__18)
	}
	{
		p.SetState(425)
		p.Match(DCM_2_0_grammarParserNAME)
	}
	{
		p.SetState(426)
		p.Anzahl_x()
	}
	{
		p.SetState(427)
		p.Match(DCM_2_0_grammarParserT__0)
	}
	{
		p.SetState(428)
		p.Kgr_info()
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DCM_2_0_grammarParserT__20 {
		{
			p.SetState(429)
			p.Einheit_x()
		}

	}
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__29 || _la == DCM_2_0_grammarParserT__30 {
		{
			p.SetState(432)
			p.Sst_liste_x()
		}

		p.SetState(435)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(437)
		p.Match(DCM_2_0_grammarParserT__4)
	}
	p.SetState(439)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
		{
			p.SetState(438)
			p.Match(DCM_2_0_grammarParserT__0)
		}

		p.SetState(441)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IKenntextContext is an interface to support dynamic dispatch.
type IKenntextContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKenntextContext differentiates from other interfaces.
	IsKenntextContext()
}

type KenntextContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKenntextContext() *KenntextContext {
	var p = new(KenntextContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_kenntext
	return p
}

func (*KenntextContext) IsKenntextContext() {}

func NewKenntextContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KenntextContext {
	var p = new(KenntextContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_kenntext

	return p
}

func (s *KenntextContext) GetParser() antlr.Parser { return s.parser }

func (s *KenntextContext) NAME() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, 0)
}

func (s *KenntextContext) Kgr_info() IKgr_infoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKgr_infoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKgr_infoContext)
}

func (s *KenntextContext) TEXT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, 0)
}

func (s *KenntextContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KenntextContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KenntextContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKenntext(s)
	}
}

func (s *KenntextContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKenntext(s)
	}
}

func (p *DCM_2_0_grammarParser) Kenntext() (localctx IKenntextContext) {
	this := p
	_ = this

	localctx = NewKenntextContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, DCM_2_0_grammarParserRULE_kenntext)
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
		p.SetState(443)
		p.Match(DCM_2_0_grammarParserT__19)
	}
	{
		p.SetState(444)
		p.Match(DCM_2_0_grammarParserNAME)
	}
	{
		p.SetState(445)
		p.Match(DCM_2_0_grammarParserT__0)
	}
	{
		p.SetState(446)
		p.Kgr_info()
	}
	{
		p.SetState(447)
		p.Match(DCM_2_0_grammarParserT__10)
	}
	{
		p.SetState(448)
		p.Match(DCM_2_0_grammarParserTEXT)
	}
	{
		p.SetState(449)
		p.Match(DCM_2_0_grammarParserT__0)
	}
	{
		p.SetState(450)
		p.Match(DCM_2_0_grammarParserT__4)
	}
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__0 {
		{
			p.SetState(451)
			p.Match(DCM_2_0_grammarParserT__0)
		}

		p.SetState(454)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IKgr_infoContext is an interface to support dynamic dispatch.
type IKgr_infoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKgr_infoContext differentiates from other interfaces.
	IsKgr_infoContext()
}

type Kgr_infoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKgr_infoContext() *Kgr_infoContext {
	var p = new(Kgr_infoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_kgr_info
	return p
}

func (*Kgr_infoContext) IsKgr_infoContext() {}

func NewKgr_infoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kgr_infoContext {
	var p = new(Kgr_infoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_kgr_info

	return p
}

func (s *Kgr_infoContext) GetParser() antlr.Parser { return s.parser }

func (s *Kgr_infoContext) Langname() ILangnameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILangnameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILangnameContext)
}

func (s *Kgr_infoContext) Displayname() IDisplaynameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDisplaynameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDisplaynameContext)
}

func (s *Kgr_infoContext) Var_abhangigkeiten() IVar_abhangigkeitenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_abhangigkeitenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_abhangigkeitenContext)
}

func (s *Kgr_infoContext) Funktionszugehorigkeit() IFunktionszugehorigkeitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunktionszugehorigkeitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunktionszugehorigkeitContext)
}

func (s *Kgr_infoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kgr_infoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kgr_infoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKgr_info(s)
	}
}

func (s *Kgr_infoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKgr_info(s)
	}
}

func (p *DCM_2_0_grammarParser) Kgr_info() (localctx IKgr_infoContext) {
	this := p
	_ = this

	localctx = NewKgr_infoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, DCM_2_0_grammarParserRULE_kgr_info)
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
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DCM_2_0_grammarParserT__23 {
		{
			p.SetState(456)
			p.Langname()
		}

	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DCM_2_0_grammarParserT__24 {
		{
			p.SetState(459)
			p.Displayname()
		}

	}
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DCM_2_0_grammarParserT__25 {
		{
			p.SetState(462)
			p.Var_abhangigkeiten()
		}

	}
	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == DCM_2_0_grammarParserT__28 {
		{
			p.SetState(465)
			p.Funktionszugehorigkeit()
		}

	}

	return localctx
}

// IEinheit_xContext is an interface to support dynamic dispatch.
type IEinheit_xContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEinheit_xContext differentiates from other interfaces.
	IsEinheit_xContext()
}

type Einheit_xContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEinheit_xContext() *Einheit_xContext {
	var p = new(Einheit_xContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_einheit_x
	return p
}

func (*Einheit_xContext) IsEinheit_xContext() {}

func NewEinheit_xContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Einheit_xContext {
	var p = new(Einheit_xContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_einheit_x

	return p
}

func (s *Einheit_xContext) GetParser() antlr.Parser { return s.parser }

func (s *Einheit_xContext) TEXT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, 0)
}

func (s *Einheit_xContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Einheit_xContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Einheit_xContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterEinheit_x(s)
	}
}

func (s *Einheit_xContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitEinheit_x(s)
	}
}

func (p *DCM_2_0_grammarParser) Einheit_x() (localctx IEinheit_xContext) {
	this := p
	_ = this

	localctx = NewEinheit_xContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, DCM_2_0_grammarParserRULE_einheit_x)

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
		p.SetState(468)
		p.Match(DCM_2_0_grammarParserT__20)
	}
	{
		p.SetState(469)
		p.Match(DCM_2_0_grammarParserTEXT)
	}
	{
		p.SetState(470)
		p.Match(DCM_2_0_grammarParserT__0)
	}

	return localctx
}

// IEinheit_yContext is an interface to support dynamic dispatch.
type IEinheit_yContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEinheit_yContext differentiates from other interfaces.
	IsEinheit_yContext()
}

type Einheit_yContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEinheit_yContext() *Einheit_yContext {
	var p = new(Einheit_yContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_einheit_y
	return p
}

func (*Einheit_yContext) IsEinheit_yContext() {}

func NewEinheit_yContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Einheit_yContext {
	var p = new(Einheit_yContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_einheit_y

	return p
}

func (s *Einheit_yContext) GetParser() antlr.Parser { return s.parser }

func (s *Einheit_yContext) TEXT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, 0)
}

func (s *Einheit_yContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Einheit_yContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Einheit_yContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterEinheit_y(s)
	}
}

func (s *Einheit_yContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitEinheit_y(s)
	}
}

func (p *DCM_2_0_grammarParser) Einheit_y() (localctx IEinheit_yContext) {
	this := p
	_ = this

	localctx = NewEinheit_yContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, DCM_2_0_grammarParserRULE_einheit_y)

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
		p.SetState(472)
		p.Match(DCM_2_0_grammarParserT__21)
	}
	{
		p.SetState(473)
		p.Match(DCM_2_0_grammarParserTEXT)
	}
	{
		p.SetState(474)
		p.Match(DCM_2_0_grammarParserT__0)
	}

	return localctx
}

// IEinheit_wContext is an interface to support dynamic dispatch.
type IEinheit_wContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEinheit_wContext differentiates from other interfaces.
	IsEinheit_wContext()
}

type Einheit_wContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEinheit_wContext() *Einheit_wContext {
	var p = new(Einheit_wContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_einheit_w
	return p
}

func (*Einheit_wContext) IsEinheit_wContext() {}

func NewEinheit_wContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Einheit_wContext {
	var p = new(Einheit_wContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_einheit_w

	return p
}

func (s *Einheit_wContext) GetParser() antlr.Parser { return s.parser }

func (s *Einheit_wContext) TEXT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, 0)
}

func (s *Einheit_wContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Einheit_wContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Einheit_wContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterEinheit_w(s)
	}
}

func (s *Einheit_wContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitEinheit_w(s)
	}
}

func (p *DCM_2_0_grammarParser) Einheit_w() (localctx IEinheit_wContext) {
	this := p
	_ = this

	localctx = NewEinheit_wContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, DCM_2_0_grammarParserRULE_einheit_w)

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
		p.SetState(476)
		p.Match(DCM_2_0_grammarParserT__22)
	}
	{
		p.SetState(477)
		p.Match(DCM_2_0_grammarParserTEXT)
	}
	{
		p.SetState(478)
		p.Match(DCM_2_0_grammarParserT__0)
	}

	return localctx
}

// ILangnameContext is an interface to support dynamic dispatch.
type ILangnameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLangnameContext differentiates from other interfaces.
	IsLangnameContext()
}

type LangnameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLangnameContext() *LangnameContext {
	var p = new(LangnameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_langname
	return p
}

func (*LangnameContext) IsLangnameContext() {}

func NewLangnameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LangnameContext {
	var p = new(LangnameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_langname

	return p
}

func (s *LangnameContext) GetParser() antlr.Parser { return s.parser }

func (s *LangnameContext) TEXT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, 0)
}

func (s *LangnameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LangnameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LangnameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterLangname(s)
	}
}

func (s *LangnameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitLangname(s)
	}
}

func (p *DCM_2_0_grammarParser) Langname() (localctx ILangnameContext) {
	this := p
	_ = this

	localctx = NewLangnameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, DCM_2_0_grammarParserRULE_langname)

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
		p.SetState(480)
		p.Match(DCM_2_0_grammarParserT__23)
	}
	{
		p.SetState(481)
		p.Match(DCM_2_0_grammarParserTEXT)
	}
	{
		p.SetState(482)
		p.Match(DCM_2_0_grammarParserT__0)
	}

	return localctx
}

// IDisplaynameContext is an interface to support dynamic dispatch.
type IDisplaynameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDisplaynameContext differentiates from other interfaces.
	IsDisplaynameContext()
}

type DisplaynameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDisplaynameContext() *DisplaynameContext {
	var p = new(DisplaynameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_displayname
	return p
}

func (*DisplaynameContext) IsDisplaynameContext() {}

func NewDisplaynameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DisplaynameContext {
	var p = new(DisplaynameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_displayname

	return p
}

func (s *DisplaynameContext) GetParser() antlr.Parser { return s.parser }

func (s *DisplaynameContext) NAME() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, 0)
}

func (s *DisplaynameContext) TEXT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, 0)
}

func (s *DisplaynameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DisplaynameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DisplaynameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterDisplayname(s)
	}
}

func (s *DisplaynameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitDisplayname(s)
	}
}

func (p *DCM_2_0_grammarParser) Displayname() (localctx IDisplaynameContext) {
	this := p
	_ = this

	localctx = NewDisplaynameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, DCM_2_0_grammarParserRULE_displayname)
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
		p.SetState(484)
		p.Match(DCM_2_0_grammarParserT__24)
	}
	{
		p.SetState(485)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DCM_2_0_grammarParserNAME || _la == DCM_2_0_grammarParserTEXT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(486)
		p.Match(DCM_2_0_grammarParserT__0)
	}

	return localctx
}

// IVar_abhangigkeitenContext is an interface to support dynamic dispatch.
type IVar_abhangigkeitenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_abhangigkeitenContext differentiates from other interfaces.
	IsVar_abhangigkeitenContext()
}

type Var_abhangigkeitenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_abhangigkeitenContext() *Var_abhangigkeitenContext {
	var p = new(Var_abhangigkeitenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_var_abhangigkeiten
	return p
}

func (*Var_abhangigkeitenContext) IsVar_abhangigkeitenContext() {}

func NewVar_abhangigkeitenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_abhangigkeitenContext {
	var p = new(Var_abhangigkeitenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_var_abhangigkeiten

	return p
}

func (s *Var_abhangigkeitenContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_abhangigkeitenContext) AllVar_abh() []IVar_abhContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVar_abhContext)(nil)).Elem())
	var tst = make([]IVar_abhContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVar_abhContext)
		}
	}

	return tst
}

func (s *Var_abhangigkeitenContext) Var_abh(i int) IVar_abhContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_abhContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVar_abhContext)
}

func (s *Var_abhangigkeitenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_abhangigkeitenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_abhangigkeitenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterVar_abhangigkeiten(s)
	}
}

func (s *Var_abhangigkeitenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitVar_abhangigkeiten(s)
	}
}

func (p *DCM_2_0_grammarParser) Var_abhangigkeiten() (localctx IVar_abhangigkeitenContext) {
	this := p
	_ = this

	localctx = NewVar_abhangigkeitenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, DCM_2_0_grammarParserRULE_var_abhangigkeiten)
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
		p.SetState(488)
		p.Match(DCM_2_0_grammarParserT__25)
	}
	{
		p.SetState(489)
		p.Var_abh()
	}
	p.SetState(494)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == DCM_2_0_grammarParserT__26 {
		{
			p.SetState(490)
			p.Match(DCM_2_0_grammarParserT__26)
		}
		{
			p.SetState(491)
			p.Var_abh()
		}

		p.SetState(496)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(497)
		p.Match(DCM_2_0_grammarParserT__0)
	}

	return localctx
}

// IVar_abhContext is an interface to support dynamic dispatch.
type IVar_abhContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_abhContext differentiates from other interfaces.
	IsVar_abhContext()
}

type Var_abhContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_abhContext() *Var_abhContext {
	var p = new(Var_abhContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_var_abh
	return p
}

func (*Var_abhContext) IsVar_abhContext() {}

func NewVar_abhContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_abhContext {
	var p = new(Var_abhContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_var_abh

	return p
}

func (s *Var_abhContext) GetParser() antlr.Parser { return s.parser }

func (s *Var_abhContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(DCM_2_0_grammarParserNAME)
}

func (s *Var_abhContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, i)
}

func (s *Var_abhContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_abhContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_abhContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterVar_abh(s)
	}
}

func (s *Var_abhContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitVar_abh(s)
	}
}

func (p *DCM_2_0_grammarParser) Var_abh() (localctx IVar_abhContext) {
	this := p
	_ = this

	localctx = NewVar_abhContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, DCM_2_0_grammarParserRULE_var_abh)

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
		p.SetState(499)
		p.Match(DCM_2_0_grammarParserNAME)
	}
	{
		p.SetState(500)
		p.Match(DCM_2_0_grammarParserT__27)
	}
	{
		p.SetState(501)
		p.Match(DCM_2_0_grammarParserNAME)
	}

	return localctx
}

// IFunktionszugehorigkeitContext is an interface to support dynamic dispatch.
type IFunktionszugehorigkeitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunktionszugehorigkeitContext differentiates from other interfaces.
	IsFunktionszugehorigkeitContext()
}

type FunktionszugehorigkeitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunktionszugehorigkeitContext() *FunktionszugehorigkeitContext {
	var p = new(FunktionszugehorigkeitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_funktionszugehorigkeit
	return p
}

func (*FunktionszugehorigkeitContext) IsFunktionszugehorigkeitContext() {}

func NewFunktionszugehorigkeitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunktionszugehorigkeitContext {
	var p = new(FunktionszugehorigkeitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_funktionszugehorigkeit

	return p
}

func (s *FunktionszugehorigkeitContext) GetParser() antlr.Parser { return s.parser }

func (s *FunktionszugehorigkeitContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(DCM_2_0_grammarParserNAME)
}

func (s *FunktionszugehorigkeitContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserNAME, i)
}

func (s *FunktionszugehorigkeitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunktionszugehorigkeitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunktionszugehorigkeitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterFunktionszugehorigkeit(s)
	}
}

func (s *FunktionszugehorigkeitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitFunktionszugehorigkeit(s)
	}
}

func (p *DCM_2_0_grammarParser) Funktionszugehorigkeit() (localctx IFunktionszugehorigkeitContext) {
	this := p
	_ = this

	localctx = NewFunktionszugehorigkeitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, DCM_2_0_grammarParserRULE_funktionszugehorigkeit)
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
		p.SetState(503)
		p.Match(DCM_2_0_grammarParserT__28)
	}
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserNAME {
		{
			p.SetState(504)
			p.Match(DCM_2_0_grammarParserNAME)
		}

		p.SetState(507)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(509)
		p.Match(DCM_2_0_grammarParserT__0)
	}

	return localctx
}

// IAnzahl_xContext is an interface to support dynamic dispatch.
type IAnzahl_xContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnzahl_xContext differentiates from other interfaces.
	IsAnzahl_xContext()
}

type Anzahl_xContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnzahl_xContext() *Anzahl_xContext {
	var p = new(Anzahl_xContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_anzahl_x
	return p
}

func (*Anzahl_xContext) IsAnzahl_xContext() {}

func NewAnzahl_xContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Anzahl_xContext {
	var p = new(Anzahl_xContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_anzahl_x

	return p
}

func (s *Anzahl_xContext) GetParser() antlr.Parser { return s.parser }

func (s *Anzahl_xContext) INT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserINT, 0)
}

func (s *Anzahl_xContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Anzahl_xContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Anzahl_xContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterAnzahl_x(s)
	}
}

func (s *Anzahl_xContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitAnzahl_x(s)
	}
}

func (p *DCM_2_0_grammarParser) Anzahl_x() (localctx IAnzahl_xContext) {
	this := p
	_ = this

	localctx = NewAnzahl_xContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, DCM_2_0_grammarParserRULE_anzahl_x)

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
		p.SetState(511)
		p.Match(DCM_2_0_grammarParserINT)
	}

	return localctx
}

// IAnzahl_yContext is an interface to support dynamic dispatch.
type IAnzahl_yContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnzahl_yContext differentiates from other interfaces.
	IsAnzahl_yContext()
}

type Anzahl_yContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnzahl_yContext() *Anzahl_yContext {
	var p = new(Anzahl_yContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_anzahl_y
	return p
}

func (*Anzahl_yContext) IsAnzahl_yContext() {}

func NewAnzahl_yContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Anzahl_yContext {
	var p = new(Anzahl_yContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_anzahl_y

	return p
}

func (s *Anzahl_yContext) GetParser() antlr.Parser { return s.parser }

func (s *Anzahl_yContext) INT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserINT, 0)
}

func (s *Anzahl_yContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Anzahl_yContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Anzahl_yContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterAnzahl_y(s)
	}
}

func (s *Anzahl_yContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitAnzahl_y(s)
	}
}

func (p *DCM_2_0_grammarParser) Anzahl_y() (localctx IAnzahl_yContext) {
	this := p
	_ = this

	localctx = NewAnzahl_yContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, DCM_2_0_grammarParserRULE_anzahl_y)

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
		p.SetState(513)
		p.Match(DCM_2_0_grammarParserINT)
	}

	return localctx
}

// IWertelisteContext is an interface to support dynamic dispatch.
type IWertelisteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWertelisteContext differentiates from other interfaces.
	IsWertelisteContext()
}

type WertelisteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWertelisteContext() *WertelisteContext {
	var p = new(WertelisteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_werteliste
	return p
}

func (*WertelisteContext) IsWertelisteContext() {}

func NewWertelisteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WertelisteContext {
	var p = new(WertelisteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_werteliste

	return p
}

func (s *WertelisteContext) GetParser() antlr.Parser { return s.parser }

func (s *WertelisteContext) AllRealzahl() []IRealzahlContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRealzahlContext)(nil)).Elem())
	var tst = make([]IRealzahlContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRealzahlContext)
		}
	}

	return tst
}

func (s *WertelisteContext) Realzahl(i int) IRealzahlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealzahlContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRealzahlContext)
}

func (s *WertelisteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WertelisteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WertelisteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterWerteliste(s)
	}
}

func (s *WertelisteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitWerteliste(s)
	}
}

func (p *DCM_2_0_grammarParser) Werteliste() (localctx IWertelisteContext) {
	this := p
	_ = this

	localctx = NewWertelisteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, DCM_2_0_grammarParserRULE_werteliste)
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
		p.SetState(515)
		p.Match(DCM_2_0_grammarParserT__9)
	}
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserINT || _la == DCM_2_0_grammarParserFLOAT {
		{
			p.SetState(516)
			p.Realzahl()
		}

		p.SetState(519)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(521)
		p.Match(DCM_2_0_grammarParserT__0)
	}

	return localctx
}

// IWerteliste_kwbContext is an interface to support dynamic dispatch.
type IWerteliste_kwbContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWerteliste_kwbContext differentiates from other interfaces.
	IsWerteliste_kwbContext()
}

type Werteliste_kwbContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWerteliste_kwbContext() *Werteliste_kwbContext {
	var p = new(Werteliste_kwbContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_werteliste_kwb
	return p
}

func (*Werteliste_kwbContext) IsWerteliste_kwbContext() {}

func NewWerteliste_kwbContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Werteliste_kwbContext {
	var p = new(Werteliste_kwbContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_werteliste_kwb

	return p
}

func (s *Werteliste_kwbContext) GetParser() antlr.Parser { return s.parser }

func (s *Werteliste_kwbContext) AllRealzahl() []IRealzahlContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRealzahlContext)(nil)).Elem())
	var tst = make([]IRealzahlContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRealzahlContext)
		}
	}

	return tst
}

func (s *Werteliste_kwbContext) Realzahl(i int) IRealzahlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealzahlContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRealzahlContext)
}

func (s *Werteliste_kwbContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(DCM_2_0_grammarParserTEXT)
}

func (s *Werteliste_kwbContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, i)
}

func (s *Werteliste_kwbContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Werteliste_kwbContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Werteliste_kwbContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterWerteliste_kwb(s)
	}
}

func (s *Werteliste_kwbContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitWerteliste_kwb(s)
	}
}

func (p *DCM_2_0_grammarParser) Werteliste_kwb() (localctx IWerteliste_kwbContext) {
	this := p
	_ = this

	localctx = NewWerteliste_kwbContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, DCM_2_0_grammarParserRULE_werteliste_kwb)
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
	p.SetState(538)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DCM_2_0_grammarParserT__9:
		{
			p.SetState(523)
			p.Match(DCM_2_0_grammarParserT__9)
		}
		p.SetState(525)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserINT || _la == DCM_2_0_grammarParserFLOAT {
			{
				p.SetState(524)
				p.Realzahl()
			}

			p.SetState(527)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(529)
			p.Match(DCM_2_0_grammarParserT__0)
		}

	case DCM_2_0_grammarParserT__10:
		{
			p.SetState(531)
			p.Match(DCM_2_0_grammarParserT__10)
		}
		p.SetState(533)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserTEXT {
			{
				p.SetState(532)
				p.Match(DCM_2_0_grammarParserTEXT)
			}

			p.SetState(535)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(537)
			p.Match(DCM_2_0_grammarParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISst_liste_xContext is an interface to support dynamic dispatch.
type ISst_liste_xContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSst_liste_xContext differentiates from other interfaces.
	IsSst_liste_xContext()
}

type Sst_liste_xContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySst_liste_xContext() *Sst_liste_xContext {
	var p = new(Sst_liste_xContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_sst_liste_x
	return p
}

func (*Sst_liste_xContext) IsSst_liste_xContext() {}

func NewSst_liste_xContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sst_liste_xContext {
	var p = new(Sst_liste_xContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_sst_liste_x

	return p
}

func (s *Sst_liste_xContext) GetParser() antlr.Parser { return s.parser }

func (s *Sst_liste_xContext) AllRealzahl() []IRealzahlContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRealzahlContext)(nil)).Elem())
	var tst = make([]IRealzahlContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRealzahlContext)
		}
	}

	return tst
}

func (s *Sst_liste_xContext) Realzahl(i int) IRealzahlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealzahlContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRealzahlContext)
}

func (s *Sst_liste_xContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(DCM_2_0_grammarParserTEXT)
}

func (s *Sst_liste_xContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, i)
}

func (s *Sst_liste_xContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sst_liste_xContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sst_liste_xContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterSst_liste_x(s)
	}
}

func (s *Sst_liste_xContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitSst_liste_x(s)
	}
}

func (p *DCM_2_0_grammarParser) Sst_liste_x() (localctx ISst_liste_xContext) {
	this := p
	_ = this

	localctx = NewSst_liste_xContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, DCM_2_0_grammarParserRULE_sst_liste_x)
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
	p.SetState(555)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DCM_2_0_grammarParserT__29:
		{
			p.SetState(540)
			p.Match(DCM_2_0_grammarParserT__29)
		}
		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserINT || _la == DCM_2_0_grammarParserFLOAT {
			{
				p.SetState(541)
				p.Realzahl()
			}

			p.SetState(544)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(546)
			p.Match(DCM_2_0_grammarParserT__0)
		}

	case DCM_2_0_grammarParserT__30:
		{
			p.SetState(548)
			p.Match(DCM_2_0_grammarParserT__30)
		}
		p.SetState(550)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserTEXT {
			{
				p.SetState(549)
				p.Match(DCM_2_0_grammarParserTEXT)
			}

			p.SetState(552)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(554)
			p.Match(DCM_2_0_grammarParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKf_zeile_listeContext is an interface to support dynamic dispatch.
type IKf_zeile_listeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKf_zeile_listeContext differentiates from other interfaces.
	IsKf_zeile_listeContext()
}

type Kf_zeile_listeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKf_zeile_listeContext() *Kf_zeile_listeContext {
	var p = new(Kf_zeile_listeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_kf_zeile_liste
	return p
}

func (*Kf_zeile_listeContext) IsKf_zeile_listeContext() {}

func NewKf_zeile_listeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kf_zeile_listeContext {
	var p = new(Kf_zeile_listeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_kf_zeile_liste

	return p
}

func (s *Kf_zeile_listeContext) GetParser() antlr.Parser { return s.parser }

func (s *Kf_zeile_listeContext) AllKf_zeile_liste_r() []IKf_zeile_liste_rContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKf_zeile_liste_rContext)(nil)).Elem())
	var tst = make([]IKf_zeile_liste_rContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKf_zeile_liste_rContext)
		}
	}

	return tst
}

func (s *Kf_zeile_listeContext) Kf_zeile_liste_r(i int) IKf_zeile_liste_rContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKf_zeile_liste_rContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKf_zeile_liste_rContext)
}

func (s *Kf_zeile_listeContext) AllKf_zeile_liste_tx() []IKf_zeile_liste_txContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKf_zeile_liste_txContext)(nil)).Elem())
	var tst = make([]IKf_zeile_liste_txContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKf_zeile_liste_txContext)
		}
	}

	return tst
}

func (s *Kf_zeile_listeContext) Kf_zeile_liste_tx(i int) IKf_zeile_liste_txContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKf_zeile_liste_txContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKf_zeile_liste_txContext)
}

func (s *Kf_zeile_listeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kf_zeile_listeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kf_zeile_listeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKf_zeile_liste(s)
	}
}

func (s *Kf_zeile_listeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKf_zeile_liste(s)
	}
}

func (p *DCM_2_0_grammarParser) Kf_zeile_liste() (localctx IKf_zeile_listeContext) {
	this := p
	_ = this

	localctx = NewKf_zeile_listeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, DCM_2_0_grammarParserRULE_kf_zeile_liste)
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
	p.SetState(567)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DCM_2_0_grammarParserT__31:
		p.SetState(558)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__31 {
			{
				p.SetState(557)
				p.Kf_zeile_liste_r()
			}

			p.SetState(560)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case DCM_2_0_grammarParserT__32:
		p.SetState(563)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__32 {
			{
				p.SetState(562)
				p.Kf_zeile_liste_tx()
			}

			p.SetState(565)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IKf_zeile_liste_rContext is an interface to support dynamic dispatch.
type IKf_zeile_liste_rContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKf_zeile_liste_rContext differentiates from other interfaces.
	IsKf_zeile_liste_rContext()
}

type Kf_zeile_liste_rContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKf_zeile_liste_rContext() *Kf_zeile_liste_rContext {
	var p = new(Kf_zeile_liste_rContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_kf_zeile_liste_r
	return p
}

func (*Kf_zeile_liste_rContext) IsKf_zeile_liste_rContext() {}

func NewKf_zeile_liste_rContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kf_zeile_liste_rContext {
	var p = new(Kf_zeile_liste_rContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_kf_zeile_liste_r

	return p
}

func (s *Kf_zeile_liste_rContext) GetParser() antlr.Parser { return s.parser }

func (s *Kf_zeile_liste_rContext) Realzahl() IRealzahlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealzahlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealzahlContext)
}

func (s *Kf_zeile_liste_rContext) AllWerteliste() []IWertelisteContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWertelisteContext)(nil)).Elem())
	var tst = make([]IWertelisteContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWertelisteContext)
		}
	}

	return tst
}

func (s *Kf_zeile_liste_rContext) Werteliste(i int) IWertelisteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWertelisteContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWertelisteContext)
}

func (s *Kf_zeile_liste_rContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kf_zeile_liste_rContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kf_zeile_liste_rContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKf_zeile_liste_r(s)
	}
}

func (s *Kf_zeile_liste_rContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKf_zeile_liste_r(s)
	}
}

func (p *DCM_2_0_grammarParser) Kf_zeile_liste_r() (localctx IKf_zeile_liste_rContext) {
	this := p
	_ = this

	localctx = NewKf_zeile_liste_rContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, DCM_2_0_grammarParserRULE_kf_zeile_liste_r)
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
		p.SetState(569)
		p.Match(DCM_2_0_grammarParserT__31)
	}
	{
		p.SetState(570)
		p.Realzahl()
	}
	{
		p.SetState(571)
		p.Match(DCM_2_0_grammarParserT__0)
	}
	p.SetState(573)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__9 {
		{
			p.SetState(572)
			p.Werteliste()
		}

		p.SetState(575)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IKf_zeile_liste_txContext is an interface to support dynamic dispatch.
type IKf_zeile_liste_txContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKf_zeile_liste_txContext differentiates from other interfaces.
	IsKf_zeile_liste_txContext()
}

type Kf_zeile_liste_txContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKf_zeile_liste_txContext() *Kf_zeile_liste_txContext {
	var p = new(Kf_zeile_liste_txContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_kf_zeile_liste_tx
	return p
}

func (*Kf_zeile_liste_txContext) IsKf_zeile_liste_txContext() {}

func NewKf_zeile_liste_txContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Kf_zeile_liste_txContext {
	var p = new(Kf_zeile_liste_txContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_kf_zeile_liste_tx

	return p
}

func (s *Kf_zeile_liste_txContext) GetParser() antlr.Parser { return s.parser }

func (s *Kf_zeile_liste_txContext) TEXT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserTEXT, 0)
}

func (s *Kf_zeile_liste_txContext) AllWerteliste() []IWertelisteContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWertelisteContext)(nil)).Elem())
	var tst = make([]IWertelisteContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWertelisteContext)
		}
	}

	return tst
}

func (s *Kf_zeile_liste_txContext) Werteliste(i int) IWertelisteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWertelisteContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWertelisteContext)
}

func (s *Kf_zeile_liste_txContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Kf_zeile_liste_txContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Kf_zeile_liste_txContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterKf_zeile_liste_tx(s)
	}
}

func (s *Kf_zeile_liste_txContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitKf_zeile_liste_tx(s)
	}
}

func (p *DCM_2_0_grammarParser) Kf_zeile_liste_tx() (localctx IKf_zeile_liste_txContext) {
	this := p
	_ = this

	localctx = NewKf_zeile_liste_txContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, DCM_2_0_grammarParserRULE_kf_zeile_liste_tx)
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
		p.SetState(577)
		p.Match(DCM_2_0_grammarParserT__32)
	}
	{
		p.SetState(578)
		p.Match(DCM_2_0_grammarParserTEXT)
	}
	{
		p.SetState(579)
		p.Match(DCM_2_0_grammarParserT__0)
	}
	p.SetState(581)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == DCM_2_0_grammarParserT__9 {
		{
			p.SetState(580)
			p.Werteliste()
		}

		p.SetState(583)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRealzahlContext is an interface to support dynamic dispatch.
type IRealzahlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealzahlContext differentiates from other interfaces.
	IsRealzahlContext()
}

type RealzahlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealzahlContext() *RealzahlContext {
	var p = new(RealzahlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DCM_2_0_grammarParserRULE_realzahl
	return p
}

func (*RealzahlContext) IsRealzahlContext() {}

func NewRealzahlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealzahlContext {
	var p = new(RealzahlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DCM_2_0_grammarParserRULE_realzahl

	return p
}

func (s *RealzahlContext) GetParser() antlr.Parser { return s.parser }

func (s *RealzahlContext) INT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserINT, 0)
}

func (s *RealzahlContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(DCM_2_0_grammarParserFLOAT, 0)
}

func (s *RealzahlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealzahlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealzahlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.EnterRealzahl(s)
	}
}

func (s *RealzahlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DCM_2_0_grammarListener); ok {
		listenerT.ExitRealzahl(s)
	}
}

func (p *DCM_2_0_grammarParser) Realzahl() (localctx IRealzahlContext) {
	this := p
	_ = this

	localctx = NewRealzahlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, DCM_2_0_grammarParserRULE_realzahl)
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
		p.SetState(585)
		_la = p.GetTokenStream().LA(1)

		if !(_la == DCM_2_0_grammarParserINT || _la == DCM_2_0_grammarParserFLOAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
