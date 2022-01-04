// Code generated from orwell.g4 by ANTLR 4.9.3. DO NOT EDIT.

package orwell // orwell
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 37, 547,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 3, 2, 6, 2, 116, 10, 2, 13, 2, 14, 2,
	117, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 125, 10, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 136, 10, 5, 12, 5, 14, 5, 139,
	11, 5, 3, 6, 3, 6, 3, 6, 7, 6, 144, 10, 6, 12, 6, 14, 6, 147, 11, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 158, 10, 7, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 170, 10,
	8, 3, 9, 3, 9, 7, 9, 174, 10, 9, 12, 9, 14, 9, 177, 11, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 5, 10, 183, 10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 195, 10, 11, 3, 12, 3, 12, 3, 12,
	3, 12, 5, 12, 201, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 207, 10,
	13, 3, 14, 3, 14, 3, 14, 7, 14, 212, 10, 14, 12, 14, 14, 14, 215, 11, 14,
	5, 14, 217, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 223, 10, 15, 3,
	15, 3, 15, 5, 15, 227, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16,
	234, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 5, 17, 245, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19,
	3, 19, 7, 19, 255, 10, 19, 12, 19, 14, 19, 258, 11, 19, 3, 19, 3, 19, 3,
	20, 3, 20, 7, 20, 264, 10, 20, 12, 20, 14, 20, 267, 11, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 276, 10, 20, 3, 21, 3, 21, 6,
	21, 280, 10, 21, 13, 21, 14, 21, 281, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	5, 22, 289, 10, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 297,
	10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 303, 10, 24, 12, 24, 14, 24,
	306, 11, 24, 3, 25, 3, 25, 5, 25, 310, 10, 25, 3, 25, 5, 25, 313, 10, 25,
	3, 26, 3, 26, 3, 26, 7, 26, 318, 10, 26, 12, 26, 14, 26, 321, 11, 26, 3,
	26, 3, 26, 5, 26, 325, 10, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 6, 29, 338, 10, 29, 13, 29, 14, 29,
	339, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 346, 10, 30, 3, 31, 3, 31, 3, 31,
	3, 31, 5, 31, 352, 10, 31, 3, 32, 3, 32, 3, 32, 7, 32, 357, 10, 32, 12,
	32, 14, 32, 360, 11, 32, 5, 32, 362, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33,
	5, 33, 368, 10, 33, 3, 33, 3, 33, 5, 33, 372, 10, 33, 3, 34, 3, 34, 3,
	34, 3, 34, 5, 34, 378, 10, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 5, 35, 388, 10, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	36, 7, 36, 396, 10, 36, 12, 36, 14, 36, 399, 11, 36, 3, 36, 3, 36, 3, 37,
	3, 37, 3, 37, 3, 37, 7, 37, 407, 10, 37, 12, 37, 14, 37, 410, 11, 37, 5,
	37, 412, 10, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 420,
	10, 38, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 426, 10, 39, 3, 40, 3, 40, 3,
	40, 7, 40, 431, 10, 40, 12, 40, 14, 40, 434, 11, 40, 5, 40, 436, 10, 40,
	3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 442, 10, 41, 3, 41, 3, 41, 5, 41, 446,
	10, 41, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42, 452, 10, 42, 3, 43, 3, 43, 3,
	43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 462, 10, 43, 3, 44, 3, 44,
	3, 44, 5, 44, 467, 10, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 7,
	45, 475, 10, 45, 12, 45, 14, 45, 478, 11, 45, 3, 45, 3, 45, 3, 46, 3, 46,
	3, 46, 3, 46, 7, 46, 486, 10, 46, 12, 46, 14, 46, 489, 11, 46, 5, 46, 491,
	10, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 5, 47, 499, 10, 47, 3,
	47, 3, 47, 5, 47, 503, 10, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48,
	3, 48, 3, 48, 7, 48, 513, 10, 48, 12, 48, 14, 48, 516, 11, 48, 5, 48, 518,
	10, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 5, 49, 527, 10,
	49, 3, 50, 3, 50, 5, 50, 531, 10, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53,
	3, 53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 2,
	2, 58, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
	106, 108, 110, 112, 2, 4, 3, 2, 14, 19, 5, 2, 28, 28, 30, 30, 36, 36, 2,
	567, 2, 115, 3, 2, 2, 2, 4, 124, 3, 2, 2, 2, 6, 126, 3, 2, 2, 2, 8, 130,
	3, 2, 2, 2, 10, 140, 3, 2, 2, 2, 12, 151, 3, 2, 2, 2, 14, 169, 3, 2, 2,
	2, 16, 171, 3, 2, 2, 2, 18, 178, 3, 2, 2, 2, 20, 194, 3, 2, 2, 2, 22, 196,
	3, 2, 2, 2, 24, 206, 3, 2, 2, 2, 26, 216, 3, 2, 2, 2, 28, 226, 3, 2, 2,
	2, 30, 233, 3, 2, 2, 2, 32, 244, 3, 2, 2, 2, 34, 246, 3, 2, 2, 2, 36, 250,
	3, 2, 2, 2, 38, 275, 3, 2, 2, 2, 40, 277, 3, 2, 2, 2, 42, 288, 3, 2, 2,
	2, 44, 290, 3, 2, 2, 2, 46, 292, 3, 2, 2, 2, 48, 309, 3, 2, 2, 2, 50, 314,
	3, 2, 2, 2, 52, 326, 3, 2, 2, 2, 54, 331, 3, 2, 2, 2, 56, 335, 3, 2, 2,
	2, 58, 341, 3, 2, 2, 2, 60, 351, 3, 2, 2, 2, 62, 361, 3, 2, 2, 2, 64, 371,
	3, 2, 2, 2, 66, 377, 3, 2, 2, 2, 68, 387, 3, 2, 2, 2, 70, 389, 3, 2, 2,
	2, 72, 402, 3, 2, 2, 2, 74, 415, 3, 2, 2, 2, 76, 425, 3, 2, 2, 2, 78, 435,
	3, 2, 2, 2, 80, 445, 3, 2, 2, 2, 82, 451, 3, 2, 2, 2, 84, 461, 3, 2, 2,
	2, 86, 466, 3, 2, 2, 2, 88, 468, 3, 2, 2, 2, 90, 481, 3, 2, 2, 2, 92, 494,
	3, 2, 2, 2, 94, 506, 3, 2, 2, 2, 96, 526, 3, 2, 2, 2, 98, 530, 3, 2, 2,
	2, 100, 532, 3, 2, 2, 2, 102, 534, 3, 2, 2, 2, 104, 536, 3, 2, 2, 2, 106,
	538, 3, 2, 2, 2, 108, 540, 3, 2, 2, 2, 110, 542, 3, 2, 2, 2, 112, 544,
	3, 2, 2, 2, 114, 116, 5, 4, 3, 2, 115, 114, 3, 2, 2, 2, 116, 117, 3, 2,
	2, 2, 117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 3, 3, 2, 2, 2, 119,
	125, 5, 6, 4, 2, 120, 125, 5, 8, 5, 2, 121, 125, 5, 10, 6, 2, 122, 125,
	5, 40, 21, 2, 123, 125, 5, 46, 24, 2, 124, 119, 3, 2, 2, 2, 124, 120, 3,
	2, 2, 2, 124, 121, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 123, 3, 2, 2,
	2, 125, 5, 3, 2, 2, 2, 126, 127, 5, 14, 8, 2, 127, 128, 7, 3, 2, 2, 128,
	129, 5, 22, 12, 2, 129, 7, 3, 2, 2, 2, 130, 131, 5, 14, 8, 2, 131, 132,
	7, 4, 2, 2, 132, 137, 5, 38, 20, 2, 133, 134, 7, 5, 2, 2, 134, 136, 5,
	38, 20, 2, 135, 133, 3, 2, 2, 2, 136, 139, 3, 2, 2, 2, 137, 135, 3, 2,
	2, 2, 137, 138, 3, 2, 2, 2, 138, 9, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 140,
	145, 5, 12, 7, 2, 141, 142, 7, 6, 2, 2, 142, 144, 5, 12, 7, 2, 143, 141,
	3, 2, 2, 2, 144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2,
	2, 2, 146, 148, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148, 149, 7, 7, 2, 2,
	149, 150, 5, 22, 12, 2, 150, 11, 3, 2, 2, 2, 151, 152, 5, 112, 57, 2, 152,
	157, 7, 8, 2, 2, 153, 154, 5, 112, 57, 2, 154, 155, 5, 104, 53, 2, 155,
	158, 3, 2, 2, 2, 156, 158, 5, 102, 52, 2, 157, 153, 3, 2, 2, 2, 157, 156,
	3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 7, 9, 2, 2, 160, 13, 3, 2,
	2, 2, 161, 162, 5, 108, 55, 2, 162, 163, 5, 102, 52, 2, 163, 164, 5, 108,
	55, 2, 164, 170, 3, 2, 2, 2, 165, 166, 5, 104, 53, 2, 166, 167, 5, 108,
	55, 2, 167, 170, 3, 2, 2, 2, 168, 170, 5, 16, 9, 2, 169, 161, 3, 2, 2,
	2, 169, 165, 3, 2, 2, 2, 169, 168, 3, 2, 2, 2, 170, 15, 3, 2, 2, 2, 171,
	175, 5, 18, 10, 2, 172, 174, 5, 108, 55, 2, 173, 172, 3, 2, 2, 2, 174,
	177, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 17, 3,
	2, 2, 2, 177, 175, 3, 2, 2, 2, 178, 179, 5, 106, 54, 2, 179, 182, 7, 8,
	2, 2, 180, 183, 5, 14, 8, 2, 181, 183, 5, 20, 11, 2, 182, 180, 3, 2, 2,
	2, 182, 181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 184, 185, 7, 9, 2, 2, 185,
	19, 3, 2, 2, 2, 186, 195, 5, 104, 53, 2, 187, 195, 5, 102, 52, 2, 188,
	189, 5, 102, 52, 2, 189, 190, 5, 108, 55, 2, 190, 195, 3, 2, 2, 2, 191,
	192, 5, 108, 55, 2, 192, 193, 5, 102, 52, 2, 193, 195, 3, 2, 2, 2, 194,
	186, 3, 2, 2, 2, 194, 187, 3, 2, 2, 2, 194, 188, 3, 2, 2, 2, 194, 191,
	3, 2, 2, 2, 195, 21, 3, 2, 2, 2, 196, 200, 5, 24, 13, 2, 197, 198, 5, 102,
	52, 2, 198, 199, 5, 22, 12, 2, 199, 201, 3, 2, 2, 2, 200, 197, 3, 2, 2,
	2, 200, 201, 3, 2, 2, 2, 201, 23, 3, 2, 2, 2, 202, 203, 5, 104, 53, 2,
	203, 204, 5, 24, 13, 2, 204, 207, 3, 2, 2, 2, 205, 207, 5, 26, 14, 2, 206,
	202, 3, 2, 2, 2, 206, 205, 3, 2, 2, 2, 207, 25, 3, 2, 2, 2, 208, 217, 5,
	30, 16, 2, 209, 213, 5, 28, 15, 2, 210, 212, 5, 30, 16, 2, 211, 210, 3,
	2, 2, 2, 212, 215, 3, 2, 2, 2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2,
	2, 214, 217, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 216, 208, 3, 2, 2, 2, 216,
	209, 3, 2, 2, 2, 217, 27, 3, 2, 2, 2, 218, 227, 5, 106, 54, 2, 219, 222,
	7, 8, 2, 2, 220, 223, 5, 22, 12, 2, 221, 223, 5, 32, 17, 2, 222, 220, 3,
	2, 2, 2, 222, 221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 225, 7, 9, 2,
	2, 225, 227, 3, 2, 2, 2, 226, 218, 3, 2, 2, 2, 226, 219, 3, 2, 2, 2, 227,
	29, 3, 2, 2, 2, 228, 234, 5, 28, 15, 2, 229, 230, 5, 108, 55, 2, 230, 231,
	5, 36, 19, 2, 231, 234, 3, 2, 2, 2, 232, 234, 5, 34, 18, 2, 233, 228, 3,
	2, 2, 2, 233, 229, 3, 2, 2, 2, 233, 232, 3, 2, 2, 2, 234, 31, 3, 2, 2,
	2, 235, 236, 5, 104, 53, 2, 236, 237, 5, 102, 52, 2, 237, 245, 3, 2, 2,
	2, 238, 239, 5, 102, 52, 2, 239, 240, 5, 24, 13, 2, 240, 245, 3, 2, 2,
	2, 241, 242, 5, 24, 13, 2, 242, 243, 5, 102, 52, 2, 243, 245, 3, 2, 2,
	2, 244, 235, 3, 2, 2, 2, 244, 238, 3, 2, 2, 2, 244, 241, 3, 2, 2, 2, 245,
	33, 3, 2, 2, 2, 246, 247, 7, 10, 2, 2, 247, 248, 5, 22, 12, 2, 248, 249,
	7, 11, 2, 2, 249, 35, 3, 2, 2, 2, 250, 251, 7, 8, 2, 2, 251, 256, 5, 22,
	12, 2, 252, 253, 7, 6, 2, 2, 253, 255, 5, 22, 12, 2, 254, 252, 3, 2, 2,
	2, 255, 258, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257,
	259, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 259, 260, 7, 9, 2, 2, 260, 37, 3,
	2, 2, 2, 261, 265, 5, 110, 56, 2, 262, 264, 5, 30, 16, 2, 263, 262, 3,
	2, 2, 2, 264, 267, 3, 2, 2, 2, 265, 263, 3, 2, 2, 2, 265, 266, 3, 2, 2,
	2, 266, 276, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 268, 269, 5, 30, 16, 2,
	269, 270, 5, 102, 52, 2, 270, 271, 5, 30, 16, 2, 271, 276, 3, 2, 2, 2,
	272, 273, 5, 104, 53, 2, 273, 274, 5, 30, 16, 2, 274, 276, 3, 2, 2, 2,
	275, 261, 3, 2, 2, 2, 275, 268, 3, 2, 2, 2, 275, 272, 3, 2, 2, 2, 276,
	39, 3, 2, 2, 2, 277, 279, 5, 42, 22, 2, 278, 280, 7, 33, 2, 2, 279, 278,
	3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 279, 3, 2, 2, 2, 281, 282, 3, 2,
	2, 2, 282, 41, 3, 2, 2, 2, 283, 284, 5, 44, 23, 2, 284, 285, 7, 35, 2,
	2, 285, 289, 3, 2, 2, 2, 286, 289, 7, 12, 2, 2, 287, 289, 7, 13, 2, 2,
	288, 283, 3, 2, 2, 2, 288, 286, 3, 2, 2, 2, 288, 287, 3, 2, 2, 2, 289,
	43, 3, 2, 2, 2, 290, 291, 9, 2, 2, 2, 291, 45, 3, 2, 2, 2, 292, 293, 5,
	58, 30, 2, 293, 294, 7, 20, 2, 2, 294, 304, 5, 48, 25, 2, 295, 297, 7,
	21, 2, 2, 296, 295, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 298, 3, 2, 2,
	2, 298, 299, 5, 58, 30, 2, 299, 300, 7, 20, 2, 2, 300, 301, 5, 48, 25,
	2, 301, 303, 3, 2, 2, 2, 302, 296, 3, 2, 2, 2, 303, 306, 3, 2, 2, 2, 304,
	302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 47, 3, 2, 2, 2, 306, 304, 3,
	2, 2, 2, 307, 310, 5, 74, 38, 2, 308, 310, 5, 50, 26, 2, 309, 307, 3, 2,
	2, 2, 309, 308, 3, 2, 2, 2, 310, 312, 3, 2, 2, 2, 311, 313, 5, 56, 29,
	2, 312, 311, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 49, 3, 2, 2, 2, 314,
	319, 5, 52, 27, 2, 315, 316, 7, 20, 2, 2, 316, 318, 5, 52, 27, 2, 317,
	315, 3, 2, 2, 2, 318, 321, 3, 2, 2, 2, 319, 317, 3, 2, 2, 2, 319, 320,
	3, 2, 2, 2, 320, 324, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2, 322, 323, 7, 20,
	2, 2, 323, 325, 5, 54, 28, 2, 324, 322, 3, 2, 2, 2, 324, 325, 3, 2, 2,
	2, 325, 51, 3, 2, 2, 2, 326, 327, 5, 74, 38, 2, 327, 328, 7, 6, 2, 2, 328,
	329, 7, 22, 2, 2, 329, 330, 5, 74, 38, 2, 330, 53, 3, 2, 2, 2, 331, 332,
	5, 74, 38, 2, 332, 333, 7, 6, 2, 2, 333, 334, 7, 23, 2, 2, 334, 55, 3,
	2, 2, 2, 335, 337, 7, 24, 2, 2, 336, 338, 5, 46, 24, 2, 337, 336, 3, 2,
	2, 2, 338, 339, 3, 2, 2, 2, 339, 337, 3, 2, 2, 2, 339, 340, 3, 2, 2, 2,
	340, 57, 3, 2, 2, 2, 341, 345, 5, 60, 31, 2, 342, 343, 5, 102, 52, 2, 343,
	344, 5, 58, 30, 2, 344, 346, 3, 2, 2, 2, 345, 342, 3, 2, 2, 2, 345, 346,
	3, 2, 2, 2, 346, 59, 3, 2, 2, 2, 347, 348, 5, 104, 53, 2, 348, 349, 5,
	60, 31, 2, 349, 352, 3, 2, 2, 2, 350, 352, 5, 62, 32, 2, 351, 347, 3, 2,
	2, 2, 351, 350, 3, 2, 2, 2, 352, 61, 3, 2, 2, 2, 353, 362, 5, 66, 34, 2,
	354, 358, 5, 64, 33, 2, 355, 357, 5, 66, 34, 2, 356, 355, 3, 2, 2, 2, 357,
	360, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 362,
	3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 361, 353, 3, 2, 2, 2, 361, 354, 3, 2,
	2, 2, 362, 63, 3, 2, 2, 2, 363, 372, 5, 112, 57, 2, 364, 367, 7, 8, 2,
	2, 365, 368, 5, 58, 30, 2, 366, 368, 5, 68, 35, 2, 367, 365, 3, 2, 2, 2,
	367, 366, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 370, 7, 9, 2, 2, 370,
	372, 3, 2, 2, 2, 371, 363, 3, 2, 2, 2, 371, 364, 3, 2, 2, 2, 372, 65, 3,
	2, 2, 2, 373, 378, 5, 64, 33, 2, 374, 378, 5, 100, 51, 2, 375, 378, 5,
	70, 36, 2, 376, 378, 5, 72, 37, 2, 377, 373, 3, 2, 2, 2, 377, 374, 3, 2,
	2, 2, 377, 375, 3, 2, 2, 2, 377, 376, 3, 2, 2, 2, 378, 67, 3, 2, 2, 2,
	379, 388, 5, 104, 53, 2, 380, 388, 5, 102, 52, 2, 381, 382, 5, 102, 52,
	2, 382, 383, 5, 60, 31, 2, 383, 388, 3, 2, 2, 2, 384, 385, 5, 60, 31, 2,
	385, 386, 5, 102, 52, 2, 386, 388, 3, 2, 2, 2, 387, 379, 3, 2, 2, 2, 387,
	380, 3, 2, 2, 2, 387, 381, 3, 2, 2, 2, 387, 384, 3, 2, 2, 2, 388, 69, 3,
	2, 2, 2, 389, 390, 7, 8, 2, 2, 390, 391, 5, 58, 30, 2, 391, 392, 7, 6,
	2, 2, 392, 397, 5, 58, 30, 2, 393, 394, 7, 6, 2, 2, 394, 396, 5, 58, 30,
	2, 395, 393, 3, 2, 2, 2, 396, 399, 3, 2, 2, 2, 397, 395, 3, 2, 2, 2, 397,
	398, 3, 2, 2, 2, 398, 400, 3, 2, 2, 2, 399, 397, 3, 2, 2, 2, 400, 401,
	7, 9, 2, 2, 401, 71, 3, 2, 2, 2, 402, 411, 7, 10, 2, 2, 403, 408, 5, 58,
	30, 2, 404, 405, 7, 6, 2, 2, 405, 407, 5, 58, 30, 2, 406, 404, 3, 2, 2,
	2, 407, 410, 3, 2, 2, 2, 408, 406, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409,
	412, 3, 2, 2, 2, 410, 408, 3, 2, 2, 2, 411, 403, 3, 2, 2, 2, 411, 412,
	3, 2, 2, 2, 412, 413, 3, 2, 2, 2, 413, 414, 7, 11, 2, 2, 414, 73, 3, 2,
	2, 2, 415, 419, 5, 76, 39, 2, 416, 417, 5, 102, 52, 2, 417, 418, 5, 74,
	38, 2, 418, 420, 3, 2, 2, 2, 419, 416, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2,
	420, 75, 3, 2, 2, 2, 421, 422, 5, 104, 53, 2, 422, 423, 5, 76, 39, 2, 423,
	426, 3, 2, 2, 2, 424, 426, 5, 78, 40, 2, 425, 421, 3, 2, 2, 2, 425, 424,
	3, 2, 2, 2, 426, 77, 3, 2, 2, 2, 427, 436, 5, 82, 42, 2, 428, 432, 5, 80,
	41, 2, 429, 431, 5, 82, 42, 2, 430, 429, 3, 2, 2, 2, 431, 434, 3, 2, 2,
	2, 432, 430, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 436, 3, 2, 2, 2, 434,
	432, 3, 2, 2, 2, 435, 427, 3, 2, 2, 2, 435, 428, 3, 2, 2, 2, 436, 79, 3,
	2, 2, 2, 437, 446, 5, 112, 57, 2, 438, 441, 7, 8, 2, 2, 439, 442, 5, 74,
	38, 2, 440, 442, 5, 84, 43, 2, 441, 439, 3, 2, 2, 2, 441, 440, 3, 2, 2,
	2, 442, 443, 3, 2, 2, 2, 443, 444, 7, 9, 2, 2, 444, 446, 3, 2, 2, 2, 445,
	437, 3, 2, 2, 2, 445, 438, 3, 2, 2, 2, 446, 81, 3, 2, 2, 2, 447, 452, 5,
	80, 41, 2, 448, 452, 5, 98, 50, 2, 449, 452, 5, 88, 45, 2, 450, 452, 5,
	86, 44, 2, 451, 447, 3, 2, 2, 2, 451, 448, 3, 2, 2, 2, 451, 449, 3, 2,
	2, 2, 451, 450, 3, 2, 2, 2, 452, 83, 3, 2, 2, 2, 453, 462, 5, 104, 53,
	2, 454, 462, 5, 102, 52, 2, 455, 456, 5, 102, 52, 2, 456, 457, 5, 76, 39,
	2, 457, 462, 3, 2, 2, 2, 458, 459, 5, 76, 39, 2, 459, 460, 5, 102, 52,
	2, 460, 462, 3, 2, 2, 2, 461, 453, 3, 2, 2, 2, 461, 454, 3, 2, 2, 2, 461,
	455, 3, 2, 2, 2, 461, 458, 3, 2, 2, 2, 462, 85, 3, 2, 2, 2, 463, 467, 5,
	90, 46, 2, 464, 467, 5, 92, 47, 2, 465, 467, 5, 94, 48, 2, 466, 463, 3,
	2, 2, 2, 466, 464, 3, 2, 2, 2, 466, 465, 3, 2, 2, 2, 467, 87, 3, 2, 2,
	2, 468, 469, 7, 8, 2, 2, 469, 470, 5, 74, 38, 2, 470, 471, 7, 6, 2, 2,
	471, 476, 5, 74, 38, 2, 472, 473, 7, 6, 2, 2, 473, 475, 5, 74, 38, 2, 474,
	472, 3, 2, 2, 2, 475, 478, 3, 2, 2, 2, 476, 474, 3, 2, 2, 2, 476, 477,
	3, 2, 2, 2, 477, 479, 3, 2, 2, 2, 478, 476, 3, 2, 2, 2, 479, 480, 7, 9,
	2, 2, 480, 89, 3, 2, 2, 2, 481, 490, 7, 10, 2, 2, 482, 487, 5, 74, 38,
	2, 483, 484, 7, 6, 2, 2, 484, 486, 5, 74, 38, 2, 485, 483, 3, 2, 2, 2,
	486, 489, 3, 2, 2, 2, 487, 485, 3, 2, 2, 2, 487, 488, 3, 2, 2, 2, 488,
	491, 3, 2, 2, 2, 489, 487, 3, 2, 2, 2, 490, 482, 3, 2, 2, 2, 490, 491,
	3, 2, 2, 2, 491, 492, 3, 2, 2, 2, 492, 493, 7, 11, 2, 2, 493, 91, 3, 2,
	2, 2, 494, 495, 7, 10, 2, 2, 495, 498, 5, 74, 38, 2, 496, 497, 7, 6, 2,
	2, 497, 499, 5, 74, 38, 2, 498, 496, 3, 2, 2, 2, 498, 499, 3, 2, 2, 2,
	499, 500, 3, 2, 2, 2, 500, 502, 7, 25, 2, 2, 501, 503, 5, 74, 38, 2, 502,
	501, 3, 2, 2, 2, 502, 503, 3, 2, 2, 2, 503, 504, 3, 2, 2, 2, 504, 505,
	7, 11, 2, 2, 505, 93, 3, 2, 2, 2, 506, 507, 7, 10, 2, 2, 507, 508, 5, 74,
	38, 2, 508, 517, 7, 5, 2, 2, 509, 514, 5, 96, 49, 2, 510, 511, 7, 26, 2,
	2, 511, 513, 5, 96, 49, 2, 512, 510, 3, 2, 2, 2, 513, 516, 3, 2, 2, 2,
	514, 512, 3, 2, 2, 2, 514, 515, 3, 2, 2, 2, 515, 518, 3, 2, 2, 2, 516,
	514, 3, 2, 2, 2, 517, 509, 3, 2, 2, 2, 517, 518, 3, 2, 2, 2, 518, 519,
	3, 2, 2, 2, 519, 520, 7, 11, 2, 2, 520, 95, 3, 2, 2, 2, 521, 527, 5, 74,
	38, 2, 522, 523, 5, 58, 30, 2, 523, 524, 7, 27, 2, 2, 524, 525, 5, 74,
	38, 2, 525, 527, 3, 2, 2, 2, 526, 521, 3, 2, 2, 2, 526, 522, 3, 2, 2, 2,
	527, 97, 3, 2, 2, 2, 528, 531, 7, 29, 2, 2, 529, 531, 5, 100, 51, 2, 530,
	528, 3, 2, 2, 2, 530, 529, 3, 2, 2, 2, 531, 99, 3, 2, 2, 2, 532, 533, 9,
	3, 2, 2, 533, 101, 3, 2, 2, 2, 534, 535, 7, 33, 2, 2, 535, 103, 3, 2, 2,
	2, 536, 537, 7, 33, 2, 2, 537, 105, 3, 2, 2, 2, 538, 539, 7, 34, 2, 2,
	539, 107, 3, 2, 2, 2, 540, 541, 7, 34, 2, 2, 541, 109, 3, 2, 2, 2, 542,
	543, 7, 34, 2, 2, 543, 111, 3, 2, 2, 2, 544, 545, 7, 34, 2, 2, 545, 113,
	3, 2, 2, 2, 60, 117, 124, 137, 145, 157, 169, 175, 182, 194, 200, 206,
	213, 216, 222, 226, 233, 244, 256, 265, 275, 281, 288, 296, 304, 309, 312,
	319, 324, 339, 345, 351, 358, 361, 367, 371, 377, 387, 397, 408, 411, 419,
	425, 432, 435, 441, 445, 451, 461, 466, 476, 487, 490, 498, 502, 514, 517,
	526, 530,
}
var literalNames = []string{
	"", "'=='", "':=='", "'|'", "','", "'::'", "'('", "')'", "'['", "']'",
	"'%prefix'", "'%prefixcon'", "'%left'", "'%right'", "'%non'", "'%leftcon'",
	"'%rightcon'", "'%noncon'", "'='", "'%else'", "'if'", "'otherwise'", "'where'",
	"'..'", "';'", "'<-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "INTEGER", "FLOAT", "STRING", "ESCCHAR",
	"PRAGMA", "OP", "ID", "DIGIT", "CHAR", "WS",
}

var ruleNames = []string{
	"program", "decl", "syndecl", "condecl", "typedecl", "name", "tylhs", "tylhs1",
	"tylhsprimary", "tylhssection", "type_", "tyterm1", "tyterm2", "typrimaryname",
	"typrimary", "tysection", "tylist", "tytuple", "construct", "opdecl", "opkind",
	"assoc", "def_", "rhs", "conditional", "ifpart", "otherpart", "wherepart",
	"pat", "pat1", "pat2", "patprimaryname", "patprimary", "patsection", "pattuple",
	"patlist", "term", "term1", "term2", "primaryname", "primary", "section",
	"list_", "tuple_", "listform", "upto", "comp", "qualifier", "fliteral",
	"literal", "infix", "prefix", "tyname", "tyvar", "con", "var_",
}

type orwellParser struct {
	*antlr.BaseParser
}

// NeworwellParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *orwellParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NeworwellParser(input antlr.TokenStream) *orwellParser {
	this := new(orwellParser)
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
	this.GrammarFileName = "orwell.g4"

	return this
}

// orwellParser tokens.
const (
	orwellParserEOF     = antlr.TokenEOF
	orwellParserT__0    = 1
	orwellParserT__1    = 2
	orwellParserT__2    = 3
	orwellParserT__3    = 4
	orwellParserT__4    = 5
	orwellParserT__5    = 6
	orwellParserT__6    = 7
	orwellParserT__7    = 8
	orwellParserT__8    = 9
	orwellParserT__9    = 10
	orwellParserT__10   = 11
	orwellParserT__11   = 12
	orwellParserT__12   = 13
	orwellParserT__13   = 14
	orwellParserT__14   = 15
	orwellParserT__15   = 16
	orwellParserT__16   = 17
	orwellParserT__17   = 18
	orwellParserT__18   = 19
	orwellParserT__19   = 20
	orwellParserT__20   = 21
	orwellParserT__21   = 22
	orwellParserT__22   = 23
	orwellParserT__23   = 24
	orwellParserT__24   = 25
	orwellParserINTEGER = 26
	orwellParserFLOAT   = 27
	orwellParserSTRING  = 28
	orwellParserESCCHAR = 29
	orwellParserPRAGMA  = 30
	orwellParserOP      = 31
	orwellParserID      = 32
	orwellParserDIGIT   = 33
	orwellParserCHAR    = 34
	orwellParserWS      = 35
)

// orwellParser rules.
const (
	orwellParserRULE_program        = 0
	orwellParserRULE_decl           = 1
	orwellParserRULE_syndecl        = 2
	orwellParserRULE_condecl        = 3
	orwellParserRULE_typedecl       = 4
	orwellParserRULE_name           = 5
	orwellParserRULE_tylhs          = 6
	orwellParserRULE_tylhs1         = 7
	orwellParserRULE_tylhsprimary   = 8
	orwellParserRULE_tylhssection   = 9
	orwellParserRULE_type_          = 10
	orwellParserRULE_tyterm1        = 11
	orwellParserRULE_tyterm2        = 12
	orwellParserRULE_typrimaryname  = 13
	orwellParserRULE_typrimary      = 14
	orwellParserRULE_tysection      = 15
	orwellParserRULE_tylist         = 16
	orwellParserRULE_tytuple        = 17
	orwellParserRULE_construct      = 18
	orwellParserRULE_opdecl         = 19
	orwellParserRULE_opkind         = 20
	orwellParserRULE_assoc          = 21
	orwellParserRULE_def_           = 22
	orwellParserRULE_rhs            = 23
	orwellParserRULE_conditional    = 24
	orwellParserRULE_ifpart         = 25
	orwellParserRULE_otherpart      = 26
	orwellParserRULE_wherepart      = 27
	orwellParserRULE_pat            = 28
	orwellParserRULE_pat1           = 29
	orwellParserRULE_pat2           = 30
	orwellParserRULE_patprimaryname = 31
	orwellParserRULE_patprimary     = 32
	orwellParserRULE_patsection     = 33
	orwellParserRULE_pattuple       = 34
	orwellParserRULE_patlist        = 35
	orwellParserRULE_term           = 36
	orwellParserRULE_term1          = 37
	orwellParserRULE_term2          = 38
	orwellParserRULE_primaryname    = 39
	orwellParserRULE_primary        = 40
	orwellParserRULE_section        = 41
	orwellParserRULE_list_          = 42
	orwellParserRULE_tuple_         = 43
	orwellParserRULE_listform       = 44
	orwellParserRULE_upto           = 45
	orwellParserRULE_comp           = 46
	orwellParserRULE_qualifier      = 47
	orwellParserRULE_fliteral       = 48
	orwellParserRULE_literal        = 49
	orwellParserRULE_infix          = 50
	orwellParserRULE_prefix         = 51
	orwellParserRULE_tyname         = 52
	orwellParserRULE_tyvar          = 53
	orwellParserRULE_con            = 54
	orwellParserRULE_var_           = 55
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllDecl() []IDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclContext)(nil)).Elem())
	var tst = make([]IDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclContext)
		}
	}

	return tst
}

func (s *ProgramContext) Decl(i int) IDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *orwellParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, orwellParserRULE_program)
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
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(orwellParserT__5-6))|(1<<(orwellParserT__7-6))|(1<<(orwellParserT__9-6))|(1<<(orwellParserT__10-6))|(1<<(orwellParserT__11-6))|(1<<(orwellParserT__12-6))|(1<<(orwellParserT__13-6))|(1<<(orwellParserT__14-6))|(1<<(orwellParserT__15-6))|(1<<(orwellParserT__16-6))|(1<<(orwellParserINTEGER-6))|(1<<(orwellParserSTRING-6))|(1<<(orwellParserOP-6))|(1<<(orwellParserID-6))|(1<<(orwellParserCHAR-6)))) != 0) {
		{
			p.SetState(112)
			p.Decl()
		}

		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) Syndecl() ISyndeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISyndeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISyndeclContext)
}

func (s *DeclContext) Condecl() ICondeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondeclContext)
}

func (s *DeclContext) Typedecl() ITypedeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypedeclContext)
}

func (s *DeclContext) Opdecl() IOpdeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpdeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpdeclContext)
}

func (s *DeclContext) Def_() IDef_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDef_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDef_Context)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (p *orwellParser) Decl() (localctx IDeclContext) {
	this := p
	_ = this

	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, orwellParserRULE_decl)

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

	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Syndecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Condecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(119)
			p.Typedecl()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(120)
			p.Opdecl()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(121)
			p.Def_()
		}

	}

	return localctx
}

// ISyndeclContext is an interface to support dynamic dispatch.
type ISyndeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSyndeclContext differentiates from other interfaces.
	IsSyndeclContext()
}

type SyndeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyndeclContext() *SyndeclContext {
	var p = new(SyndeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_syndecl
	return p
}

func (*SyndeclContext) IsSyndeclContext() {}

func NewSyndeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyndeclContext {
	var p = new(SyndeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_syndecl

	return p
}

func (s *SyndeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SyndeclContext) Tylhs() ITylhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITylhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITylhsContext)
}

func (s *SyndeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *SyndeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyndeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyndeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterSyndecl(s)
	}
}

func (s *SyndeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitSyndecl(s)
	}
}

func (p *orwellParser) Syndecl() (localctx ISyndeclContext) {
	this := p
	_ = this

	localctx = NewSyndeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, orwellParserRULE_syndecl)

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
		p.Tylhs()
	}
	{
		p.SetState(125)
		p.Match(orwellParserT__0)
	}
	{
		p.SetState(126)
		p.Type_()
	}

	return localctx
}

// ICondeclContext is an interface to support dynamic dispatch.
type ICondeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCondeclContext differentiates from other interfaces.
	IsCondeclContext()
}

type CondeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondeclContext() *CondeclContext {
	var p = new(CondeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_condecl
	return p
}

func (*CondeclContext) IsCondeclContext() {}

func NewCondeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondeclContext {
	var p = new(CondeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_condecl

	return p
}

func (s *CondeclContext) GetParser() antlr.Parser { return s.parser }

func (s *CondeclContext) Tylhs() ITylhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITylhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITylhsContext)
}

func (s *CondeclContext) AllConstruct() []IConstructContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstructContext)(nil)).Elem())
	var tst = make([]IConstructContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstructContext)
		}
	}

	return tst
}

func (s *CondeclContext) Construct(i int) IConstructContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstructContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstructContext)
}

func (s *CondeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterCondecl(s)
	}
}

func (s *CondeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitCondecl(s)
	}
}

func (p *orwellParser) Condecl() (localctx ICondeclContext) {
	this := p
	_ = this

	localctx = NewCondeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, orwellParserRULE_condecl)
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
		p.SetState(128)
		p.Tylhs()
	}
	{
		p.SetState(129)
		p.Match(orwellParserT__1)
	}
	{
		p.SetState(130)
		p.Construct()
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == orwellParserT__2 {
		{
			p.SetState(131)
			p.Match(orwellParserT__2)
		}
		{
			p.SetState(132)
			p.Construct()
		}

		p.SetState(137)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypedeclContext is an interface to support dynamic dispatch.
type ITypedeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypedeclContext differentiates from other interfaces.
	IsTypedeclContext()
}

type TypedeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedeclContext() *TypedeclContext {
	var p = new(TypedeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_typedecl
	return p
}

func (*TypedeclContext) IsTypedeclContext() {}

func NewTypedeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypedeclContext {
	var p = new(TypedeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_typedecl

	return p
}

func (s *TypedeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypedeclContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *TypedeclContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TypedeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *TypedeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypedeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTypedecl(s)
	}
}

func (s *TypedeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTypedecl(s)
	}
}

func (p *orwellParser) Typedecl() (localctx ITypedeclContext) {
	this := p
	_ = this

	localctx = NewTypedeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, orwellParserRULE_typedecl)
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
		p.SetState(138)
		p.Name()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == orwellParserT__3 {
		{
			p.SetState(139)
			p.Match(orwellParserT__3)
		}
		{
			p.SetState(140)
			p.Name()
		}

		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(146)
		p.Match(orwellParserT__4)
	}
	{
		p.SetState(147)
		p.Type_()
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) AllVar_() []IVar_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVar_Context)(nil)).Elem())
	var tst = make([]IVar_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVar_Context)
		}
	}

	return tst
}

func (s *NameContext) Var_(i int) IVar_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *NameContext) Prefix() IPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *NameContext) Infix() IInfixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixContext)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *orwellParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, orwellParserRULE_name)

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
		p.Var_()
	}
	{
		p.SetState(150)
		p.Match(orwellParserT__5)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case orwellParserID:
		{
			p.SetState(151)
			p.Var_()
		}
		{
			p.SetState(152)
			p.Prefix()
		}

	case orwellParserOP:
		{
			p.SetState(154)
			p.Infix()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(157)
		p.Match(orwellParserT__6)
	}

	return localctx
}

// ITylhsContext is an interface to support dynamic dispatch.
type ITylhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTylhsContext differentiates from other interfaces.
	IsTylhsContext()
}

type TylhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTylhsContext() *TylhsContext {
	var p = new(TylhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_tylhs
	return p
}

func (*TylhsContext) IsTylhsContext() {}

func NewTylhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TylhsContext {
	var p = new(TylhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_tylhs

	return p
}

func (s *TylhsContext) GetParser() antlr.Parser { return s.parser }

func (s *TylhsContext) AllTyvar() []ITyvarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITyvarContext)(nil)).Elem())
	var tst = make([]ITyvarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITyvarContext)
		}
	}

	return tst
}

func (s *TylhsContext) Tyvar(i int) ITyvarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyvarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITyvarContext)
}

func (s *TylhsContext) Infix() IInfixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixContext)
}

func (s *TylhsContext) Prefix() IPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *TylhsContext) Tylhs1() ITylhs1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITylhs1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITylhs1Context)
}

func (s *TylhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TylhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TylhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTylhs(s)
	}
}

func (s *TylhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTylhs(s)
	}
}

func (p *orwellParser) Tylhs() (localctx ITylhsContext) {
	this := p
	_ = this

	localctx = NewTylhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, orwellParserRULE_tylhs)

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

	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			p.Tyvar()
		}
		{
			p.SetState(160)
			p.Infix()
		}
		{
			p.SetState(161)
			p.Tyvar()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(163)
			p.Prefix()
		}
		{
			p.SetState(164)
			p.Tyvar()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(166)
			p.Tylhs1()
		}

	}

	return localctx
}

// ITylhs1Context is an interface to support dynamic dispatch.
type ITylhs1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTylhs1Context differentiates from other interfaces.
	IsTylhs1Context()
}

type Tylhs1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTylhs1Context() *Tylhs1Context {
	var p = new(Tylhs1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_tylhs1
	return p
}

func (*Tylhs1Context) IsTylhs1Context() {}

func NewTylhs1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tylhs1Context {
	var p = new(Tylhs1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_tylhs1

	return p
}

func (s *Tylhs1Context) GetParser() antlr.Parser { return s.parser }

func (s *Tylhs1Context) Tylhsprimary() ITylhsprimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITylhsprimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITylhsprimaryContext)
}

func (s *Tylhs1Context) AllTyvar() []ITyvarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITyvarContext)(nil)).Elem())
	var tst = make([]ITyvarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITyvarContext)
		}
	}

	return tst
}

func (s *Tylhs1Context) Tyvar(i int) ITyvarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyvarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITyvarContext)
}

func (s *Tylhs1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tylhs1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tylhs1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTylhs1(s)
	}
}

func (s *Tylhs1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTylhs1(s)
	}
}

func (p *orwellParser) Tylhs1() (localctx ITylhs1Context) {
	this := p
	_ = this

	localctx = NewTylhs1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, orwellParserRULE_tylhs1)
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
		p.SetState(169)
		p.Tylhsprimary()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == orwellParserID {
		{
			p.SetState(170)
			p.Tyvar()
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITylhsprimaryContext is an interface to support dynamic dispatch.
type ITylhsprimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTylhsprimaryContext differentiates from other interfaces.
	IsTylhsprimaryContext()
}

type TylhsprimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTylhsprimaryContext() *TylhsprimaryContext {
	var p = new(TylhsprimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_tylhsprimary
	return p
}

func (*TylhsprimaryContext) IsTylhsprimaryContext() {}

func NewTylhsprimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TylhsprimaryContext {
	var p = new(TylhsprimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_tylhsprimary

	return p
}

func (s *TylhsprimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *TylhsprimaryContext) Tyname() ITynameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITynameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITynameContext)
}

func (s *TylhsprimaryContext) Tylhs() ITylhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITylhsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITylhsContext)
}

func (s *TylhsprimaryContext) Tylhssection() ITylhssectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITylhssectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITylhssectionContext)
}

func (s *TylhsprimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TylhsprimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TylhsprimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTylhsprimary(s)
	}
}

func (s *TylhsprimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTylhsprimary(s)
	}
}

func (p *orwellParser) Tylhsprimary() (localctx ITylhsprimaryContext) {
	this := p
	_ = this

	localctx = NewTylhsprimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, orwellParserRULE_tylhsprimary)

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
		p.Tyname()
	}
	{
		p.SetState(177)
		p.Match(orwellParserT__5)
	}
	p.SetState(180)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(178)
			p.Tylhs()
		}

	case 2:
		{
			p.SetState(179)
			p.Tylhssection()
		}

	}
	{
		p.SetState(182)
		p.Match(orwellParserT__6)
	}

	return localctx
}

// ITylhssectionContext is an interface to support dynamic dispatch.
type ITylhssectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTylhssectionContext differentiates from other interfaces.
	IsTylhssectionContext()
}

type TylhssectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTylhssectionContext() *TylhssectionContext {
	var p = new(TylhssectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_tylhssection
	return p
}

func (*TylhssectionContext) IsTylhssectionContext() {}

func NewTylhssectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TylhssectionContext {
	var p = new(TylhssectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_tylhssection

	return p
}

func (s *TylhssectionContext) GetParser() antlr.Parser { return s.parser }

func (s *TylhssectionContext) Prefix() IPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *TylhssectionContext) Infix() IInfixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixContext)
}

func (s *TylhssectionContext) Tyvar() ITyvarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyvarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyvarContext)
}

func (s *TylhssectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TylhssectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TylhssectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTylhssection(s)
	}
}

func (s *TylhssectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTylhssection(s)
	}
}

func (p *orwellParser) Tylhssection() (localctx ITylhssectionContext) {
	this := p
	_ = this

	localctx = NewTylhssectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, orwellParserRULE_tylhssection)

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

	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(184)
			p.Prefix()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(185)
			p.Infix()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(186)
			p.Infix()
		}
		{
			p.SetState(187)
			p.Tyvar()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(189)
			p.Tyvar()
		}
		{
			p.SetState(190)
			p.Infix()
		}

	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) Tyterm1() ITyterm1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyterm1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyterm1Context)
}

func (s *Type_Context) Infix() IInfixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixContext)
}

func (s *Type_Context) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *orwellParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, orwellParserRULE_type_)

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
		p.SetState(194)
		p.Tyterm1()
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(195)
			p.Infix()
		}
		{
			p.SetState(196)
			p.Type_()
		}

	}

	return localctx
}

// ITyterm1Context is an interface to support dynamic dispatch.
type ITyterm1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTyterm1Context differentiates from other interfaces.
	IsTyterm1Context()
}

type Tyterm1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTyterm1Context() *Tyterm1Context {
	var p = new(Tyterm1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_tyterm1
	return p
}

func (*Tyterm1Context) IsTyterm1Context() {}

func NewTyterm1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tyterm1Context {
	var p = new(Tyterm1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_tyterm1

	return p
}

func (s *Tyterm1Context) GetParser() antlr.Parser { return s.parser }

func (s *Tyterm1Context) Prefix() IPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *Tyterm1Context) Tyterm1() ITyterm1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyterm1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyterm1Context)
}

func (s *Tyterm1Context) Tyterm2() ITyterm2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyterm2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyterm2Context)
}

func (s *Tyterm1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tyterm1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tyterm1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTyterm1(s)
	}
}

func (s *Tyterm1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTyterm1(s)
	}
}

func (p *orwellParser) Tyterm1() (localctx ITyterm1Context) {
	this := p
	_ = this

	localctx = NewTyterm1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, orwellParserRULE_tyterm1)

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

	p.SetState(204)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case orwellParserOP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(200)
			p.Prefix()
		}
		{
			p.SetState(201)
			p.Tyterm1()
		}

	case orwellParserT__5, orwellParserT__7, orwellParserID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.Tyterm2()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITyterm2Context is an interface to support dynamic dispatch.
type ITyterm2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTyterm2Context differentiates from other interfaces.
	IsTyterm2Context()
}

type Tyterm2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTyterm2Context() *Tyterm2Context {
	var p = new(Tyterm2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_tyterm2
	return p
}

func (*Tyterm2Context) IsTyterm2Context() {}

func NewTyterm2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tyterm2Context {
	var p = new(Tyterm2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_tyterm2

	return p
}

func (s *Tyterm2Context) GetParser() antlr.Parser { return s.parser }

func (s *Tyterm2Context) AllTyprimary() []ITyprimaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITyprimaryContext)(nil)).Elem())
	var tst = make([]ITyprimaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITyprimaryContext)
		}
	}

	return tst
}

func (s *Tyterm2Context) Typrimary(i int) ITyprimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyprimaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITyprimaryContext)
}

func (s *Tyterm2Context) Typrimaryname() ITyprimarynameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyprimarynameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyprimarynameContext)
}

func (s *Tyterm2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tyterm2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tyterm2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTyterm2(s)
	}
}

func (s *Tyterm2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTyterm2(s)
	}
}

func (p *orwellParser) Tyterm2() (localctx ITyterm2Context) {
	this := p
	_ = this

	localctx = NewTyterm2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, orwellParserRULE_tyterm2)

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

	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.Typrimary()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
			p.Typrimaryname()
		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(208)
					p.Typrimary()
				}

			}
			p.SetState(213)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
		}

	}

	return localctx
}

// ITyprimarynameContext is an interface to support dynamic dispatch.
type ITyprimarynameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTyprimarynameContext differentiates from other interfaces.
	IsTyprimarynameContext()
}

type TyprimarynameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTyprimarynameContext() *TyprimarynameContext {
	var p = new(TyprimarynameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_typrimaryname
	return p
}

func (*TyprimarynameContext) IsTyprimarynameContext() {}

func NewTyprimarynameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TyprimarynameContext {
	var p = new(TyprimarynameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_typrimaryname

	return p
}

func (s *TyprimarynameContext) GetParser() antlr.Parser { return s.parser }

func (s *TyprimarynameContext) Tyname() ITynameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITynameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITynameContext)
}

func (s *TyprimarynameContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *TyprimarynameContext) Tysection() ITysectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITysectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITysectionContext)
}

func (s *TyprimarynameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TyprimarynameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TyprimarynameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTyprimaryname(s)
	}
}

func (s *TyprimarynameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTyprimaryname(s)
	}
}

func (p *orwellParser) Typrimaryname() (localctx ITyprimarynameContext) {
	this := p
	_ = this

	localctx = NewTyprimarynameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, orwellParserRULE_typrimaryname)

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

	p.SetState(224)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case orwellParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(216)
			p.Tyname()
		}

	case orwellParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(217)
			p.Match(orwellParserT__5)
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(218)
				p.Type_()
			}

		case 2:
			{
				p.SetState(219)
				p.Tysection()
			}

		}
		{
			p.SetState(222)
			p.Match(orwellParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITyprimaryContext is an interface to support dynamic dispatch.
type ITyprimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTyprimaryContext differentiates from other interfaces.
	IsTyprimaryContext()
}

type TyprimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTyprimaryContext() *TyprimaryContext {
	var p = new(TyprimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_typrimary
	return p
}

func (*TyprimaryContext) IsTyprimaryContext() {}

func NewTyprimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TyprimaryContext {
	var p = new(TyprimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_typrimary

	return p
}

func (s *TyprimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *TyprimaryContext) Typrimaryname() ITyprimarynameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyprimarynameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyprimarynameContext)
}

func (s *TyprimaryContext) Tyvar() ITyvarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyvarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyvarContext)
}

func (s *TyprimaryContext) Tytuple() ITytupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITytupleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITytupleContext)
}

func (s *TyprimaryContext) Tylist() ITylistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITylistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITylistContext)
}

func (s *TyprimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TyprimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TyprimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTyprimary(s)
	}
}

func (s *TyprimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTyprimary(s)
	}
}

func (p *orwellParser) Typrimary() (localctx ITyprimaryContext) {
	this := p
	_ = this

	localctx = NewTyprimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, orwellParserRULE_typrimary)

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

	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(226)
			p.Typrimaryname()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(227)
			p.Tyvar()
		}
		{
			p.SetState(228)
			p.Tytuple()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.Tylist()
		}

	}

	return localctx
}

// ITysectionContext is an interface to support dynamic dispatch.
type ITysectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTysectionContext differentiates from other interfaces.
	IsTysectionContext()
}

type TysectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTysectionContext() *TysectionContext {
	var p = new(TysectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_tysection
	return p
}

func (*TysectionContext) IsTysectionContext() {}

func NewTysectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TysectionContext {
	var p = new(TysectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_tysection

	return p
}

func (s *TysectionContext) GetParser() antlr.Parser { return s.parser }

func (s *TysectionContext) Prefix() IPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *TysectionContext) Infix() IInfixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixContext)
}

func (s *TysectionContext) Tyterm1() ITyterm1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyterm1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITyterm1Context)
}

func (s *TysectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TysectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TysectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTysection(s)
	}
}

func (s *TysectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTysection(s)
	}
}

func (p *orwellParser) Tysection() (localctx ITysectionContext) {
	this := p
	_ = this

	localctx = NewTysectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, orwellParserRULE_tysection)

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

	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(233)
			p.Prefix()
		}
		{
			p.SetState(234)
			p.Infix()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.Infix()
		}
		{
			p.SetState(237)
			p.Tyterm1()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(239)
			p.Tyterm1()
		}
		{
			p.SetState(240)
			p.Infix()
		}

	}

	return localctx
}

// ITylistContext is an interface to support dynamic dispatch.
type ITylistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTylistContext differentiates from other interfaces.
	IsTylistContext()
}

type TylistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTylistContext() *TylistContext {
	var p = new(TylistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_tylist
	return p
}

func (*TylistContext) IsTylistContext() {}

func NewTylistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TylistContext {
	var p = new(TylistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_tylist

	return p
}

func (s *TylistContext) GetParser() antlr.Parser { return s.parser }

func (s *TylistContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *TylistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TylistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TylistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTylist(s)
	}
}

func (s *TylistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTylist(s)
	}
}

func (p *orwellParser) Tylist() (localctx ITylistContext) {
	this := p
	_ = this

	localctx = NewTylistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, orwellParserRULE_tylist)

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
		p.SetState(244)
		p.Match(orwellParserT__7)
	}
	{
		p.SetState(245)
		p.Type_()
	}
	{
		p.SetState(246)
		p.Match(orwellParserT__8)
	}

	return localctx
}

// ITytupleContext is an interface to support dynamic dispatch.
type ITytupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTytupleContext differentiates from other interfaces.
	IsTytupleContext()
}

type TytupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTytupleContext() *TytupleContext {
	var p = new(TytupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_tytuple
	return p
}

func (*TytupleContext) IsTytupleContext() {}

func NewTytupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TytupleContext {
	var p = new(TytupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_tytuple

	return p
}

func (s *TytupleContext) GetParser() antlr.Parser { return s.parser }

func (s *TytupleContext) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *TytupleContext) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *TytupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TytupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TytupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTytuple(s)
	}
}

func (s *TytupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTytuple(s)
	}
}

func (p *orwellParser) Tytuple() (localctx ITytupleContext) {
	this := p
	_ = this

	localctx = NewTytupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, orwellParserRULE_tytuple)
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
		p.SetState(248)
		p.Match(orwellParserT__5)
	}
	{
		p.SetState(249)
		p.Type_()
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == orwellParserT__3 {
		{
			p.SetState(250)
			p.Match(orwellParserT__3)
		}
		{
			p.SetState(251)
			p.Type_()
		}

		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(257)
		p.Match(orwellParserT__6)
	}

	return localctx
}

// IConstructContext is an interface to support dynamic dispatch.
type IConstructContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstructContext differentiates from other interfaces.
	IsConstructContext()
}

type ConstructContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructContext() *ConstructContext {
	var p = new(ConstructContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_construct
	return p
}

func (*ConstructContext) IsConstructContext() {}

func NewConstructContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructContext {
	var p = new(ConstructContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_construct

	return p
}

func (s *ConstructContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructContext) Con() IConContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConContext)
}

func (s *ConstructContext) AllTyprimary() []ITyprimaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITyprimaryContext)(nil)).Elem())
	var tst = make([]ITyprimaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITyprimaryContext)
		}
	}

	return tst
}

func (s *ConstructContext) Typrimary(i int) ITyprimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITyprimaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITyprimaryContext)
}

func (s *ConstructContext) Infix() IInfixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixContext)
}

func (s *ConstructContext) Prefix() IPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *ConstructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterConstruct(s)
	}
}

func (s *ConstructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitConstruct(s)
	}
}

func (p *orwellParser) Construct() (localctx IConstructContext) {
	this := p
	_ = this

	localctx = NewConstructContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, orwellParserRULE_construct)

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

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Con()
		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(260)
					p.Typrimary()
				}

			}
			p.SetState(265)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(266)
			p.Typrimary()
		}
		{
			p.SetState(267)
			p.Infix()
		}
		{
			p.SetState(268)
			p.Typrimary()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(270)
			p.Prefix()
		}
		{
			p.SetState(271)
			p.Typrimary()
		}

	}

	return localctx
}

// IOpdeclContext is an interface to support dynamic dispatch.
type IOpdeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpdeclContext differentiates from other interfaces.
	IsOpdeclContext()
}

type OpdeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpdeclContext() *OpdeclContext {
	var p = new(OpdeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_opdecl
	return p
}

func (*OpdeclContext) IsOpdeclContext() {}

func NewOpdeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpdeclContext {
	var p = new(OpdeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_opdecl

	return p
}

func (s *OpdeclContext) GetParser() antlr.Parser { return s.parser }

func (s *OpdeclContext) Opkind() IOpkindContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpkindContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpkindContext)
}

func (s *OpdeclContext) AllOP() []antlr.TerminalNode {
	return s.GetTokens(orwellParserOP)
}

func (s *OpdeclContext) OP(i int) antlr.TerminalNode {
	return s.GetToken(orwellParserOP, i)
}

func (s *OpdeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpdeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpdeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterOpdecl(s)
	}
}

func (s *OpdeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitOpdecl(s)
	}
}

func (p *orwellParser) Opdecl() (localctx IOpdeclContext) {
	this := p
	_ = this

	localctx = NewOpdeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, orwellParserRULE_opdecl)

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
		p.SetState(275)
		p.Opkind()
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(276)
				p.Match(orwellParserOP)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(279)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IOpkindContext is an interface to support dynamic dispatch.
type IOpkindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpkindContext differentiates from other interfaces.
	IsOpkindContext()
}

type OpkindContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpkindContext() *OpkindContext {
	var p = new(OpkindContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_opkind
	return p
}

func (*OpkindContext) IsOpkindContext() {}

func NewOpkindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpkindContext {
	var p = new(OpkindContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_opkind

	return p
}

func (s *OpkindContext) GetParser() antlr.Parser { return s.parser }

func (s *OpkindContext) Assoc() IAssocContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssocContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssocContext)
}

func (s *OpkindContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(orwellParserDIGIT, 0)
}

func (s *OpkindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpkindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpkindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterOpkind(s)
	}
}

func (s *OpkindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitOpkind(s)
	}
}

func (p *orwellParser) Opkind() (localctx IOpkindContext) {
	this := p
	_ = this

	localctx = NewOpkindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, orwellParserRULE_opkind)

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

	p.SetState(286)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case orwellParserT__11, orwellParserT__12, orwellParserT__13, orwellParserT__14, orwellParserT__15, orwellParserT__16:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Assoc()
		}
		{
			p.SetState(282)
			p.Match(orwellParserDIGIT)
		}

	case orwellParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(284)
			p.Match(orwellParserT__9)
		}

	case orwellParserT__10:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(285)
			p.Match(orwellParserT__10)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssocContext is an interface to support dynamic dispatch.
type IAssocContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssocContext differentiates from other interfaces.
	IsAssocContext()
}

type AssocContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssocContext() *AssocContext {
	var p = new(AssocContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_assoc
	return p
}

func (*AssocContext) IsAssocContext() {}

func NewAssocContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssocContext {
	var p = new(AssocContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_assoc

	return p
}

func (s *AssocContext) GetParser() antlr.Parser { return s.parser }
func (s *AssocContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssocContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssocContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterAssoc(s)
	}
}

func (s *AssocContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitAssoc(s)
	}
}

func (p *orwellParser) Assoc() (localctx IAssocContext) {
	this := p
	_ = this

	localctx = NewAssocContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, orwellParserRULE_assoc)
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
		p.SetState(288)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<orwellParserT__11)|(1<<orwellParserT__12)|(1<<orwellParserT__13)|(1<<orwellParserT__14)|(1<<orwellParserT__15)|(1<<orwellParserT__16))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDef_Context is an interface to support dynamic dispatch.
type IDef_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDef_Context differentiates from other interfaces.
	IsDef_Context()
}

type Def_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDef_Context() *Def_Context {
	var p = new(Def_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_def_
	return p
}

func (*Def_Context) IsDef_Context() {}

func NewDef_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Def_Context {
	var p = new(Def_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_def_

	return p
}

func (s *Def_Context) GetParser() antlr.Parser { return s.parser }

func (s *Def_Context) AllPat() []IPatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPatContext)(nil)).Elem())
	var tst = make([]IPatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPatContext)
		}
	}

	return tst
}

func (s *Def_Context) Pat(i int) IPatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPatContext)
}

func (s *Def_Context) AllRhs() []IRhsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRhsContext)(nil)).Elem())
	var tst = make([]IRhsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRhsContext)
		}
	}

	return tst
}

func (s *Def_Context) Rhs(i int) IRhsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRhsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRhsContext)
}

func (s *Def_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Def_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterDef_(s)
	}
}

func (s *Def_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitDef_(s)
	}
}

func (p *orwellParser) Def_() (localctx IDef_Context) {
	this := p
	_ = this

	localctx = NewDef_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, orwellParserRULE_def_)
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
		p.SetState(290)
		p.Pat()
	}
	{
		p.SetState(291)
		p.Match(orwellParserT__17)
	}
	{
		p.SetState(292)
		p.Rhs()
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == orwellParserT__18 {
				{
					p.SetState(293)
					p.Match(orwellParserT__18)
				}

			}
			{
				p.SetState(296)
				p.Pat()
			}
			{
				p.SetState(297)
				p.Match(orwellParserT__17)
			}
			{
				p.SetState(298)
				p.Rhs()
			}

		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IRhsContext is an interface to support dynamic dispatch.
type IRhsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRhsContext differentiates from other interfaces.
	IsRhsContext()
}

type RhsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRhsContext() *RhsContext {
	var p = new(RhsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_rhs
	return p
}

func (*RhsContext) IsRhsContext() {}

func NewRhsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RhsContext {
	var p = new(RhsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_rhs

	return p
}

func (s *RhsContext) GetParser() antlr.Parser { return s.parser }

func (s *RhsContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *RhsContext) Conditional() IConditionalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionalContext)
}

func (s *RhsContext) Wherepart() IWherepartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWherepartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWherepartContext)
}

func (s *RhsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RhsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RhsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterRhs(s)
	}
}

func (s *RhsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitRhs(s)
	}
}

func (p *orwellParser) Rhs() (localctx IRhsContext) {
	this := p
	_ = this

	localctx = NewRhsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, orwellParserRULE_rhs)
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
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(305)
			p.Term()
		}

	case 2:
		{
			p.SetState(306)
			p.Conditional()
		}

	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == orwellParserT__21 {
		{
			p.SetState(309)
			p.Wherepart()
		}

	}

	return localctx
}

// IConditionalContext is an interface to support dynamic dispatch.
type IConditionalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionalContext differentiates from other interfaces.
	IsConditionalContext()
}

type ConditionalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalContext() *ConditionalContext {
	var p = new(ConditionalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_conditional
	return p
}

func (*ConditionalContext) IsConditionalContext() {}

func NewConditionalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalContext {
	var p = new(ConditionalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_conditional

	return p
}

func (s *ConditionalContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalContext) AllIfpart() []IIfpartContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIfpartContext)(nil)).Elem())
	var tst = make([]IIfpartContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIfpartContext)
		}
	}

	return tst
}

func (s *ConditionalContext) Ifpart(i int) IIfpartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfpartContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIfpartContext)
}

func (s *ConditionalContext) Otherpart() IOtherpartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOtherpartContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOtherpartContext)
}

func (s *ConditionalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterConditional(s)
	}
}

func (s *ConditionalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitConditional(s)
	}
}

func (p *orwellParser) Conditional() (localctx IConditionalContext) {
	this := p
	_ = this

	localctx = NewConditionalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, orwellParserRULE_conditional)
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
		p.SetState(312)
		p.Ifpart()
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(313)
				p.Match(orwellParserT__17)
			}
			{
				p.SetState(314)
				p.Ifpart()
			}

		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}
	p.SetState(322)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == orwellParserT__17 {
		{
			p.SetState(320)
			p.Match(orwellParserT__17)
		}
		{
			p.SetState(321)
			p.Otherpart()
		}

	}

	return localctx
}

// IIfpartContext is an interface to support dynamic dispatch.
type IIfpartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfpartContext differentiates from other interfaces.
	IsIfpartContext()
}

type IfpartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfpartContext() *IfpartContext {
	var p = new(IfpartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_ifpart
	return p
}

func (*IfpartContext) IsIfpartContext() {}

func NewIfpartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfpartContext {
	var p = new(IfpartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_ifpart

	return p
}

func (s *IfpartContext) GetParser() antlr.Parser { return s.parser }

func (s *IfpartContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *IfpartContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *IfpartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfpartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfpartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterIfpart(s)
	}
}

func (s *IfpartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitIfpart(s)
	}
}

func (p *orwellParser) Ifpart() (localctx IIfpartContext) {
	this := p
	_ = this

	localctx = NewIfpartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, orwellParserRULE_ifpart)

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
		p.SetState(324)
		p.Term()
	}
	{
		p.SetState(325)
		p.Match(orwellParserT__3)
	}
	{
		p.SetState(326)
		p.Match(orwellParserT__19)
	}
	{
		p.SetState(327)
		p.Term()
	}

	return localctx
}

// IOtherpartContext is an interface to support dynamic dispatch.
type IOtherpartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOtherpartContext differentiates from other interfaces.
	IsOtherpartContext()
}

type OtherpartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOtherpartContext() *OtherpartContext {
	var p = new(OtherpartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_otherpart
	return p
}

func (*OtherpartContext) IsOtherpartContext() {}

func NewOtherpartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OtherpartContext {
	var p = new(OtherpartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_otherpart

	return p
}

func (s *OtherpartContext) GetParser() antlr.Parser { return s.parser }

func (s *OtherpartContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *OtherpartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OtherpartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OtherpartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterOtherpart(s)
	}
}

func (s *OtherpartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitOtherpart(s)
	}
}

func (p *orwellParser) Otherpart() (localctx IOtherpartContext) {
	this := p
	_ = this

	localctx = NewOtherpartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, orwellParserRULE_otherpart)

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
		p.SetState(329)
		p.Term()
	}
	{
		p.SetState(330)
		p.Match(orwellParserT__3)
	}
	{
		p.SetState(331)
		p.Match(orwellParserT__20)
	}

	return localctx
}

// IWherepartContext is an interface to support dynamic dispatch.
type IWherepartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWherepartContext differentiates from other interfaces.
	IsWherepartContext()
}

type WherepartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWherepartContext() *WherepartContext {
	var p = new(WherepartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_wherepart
	return p
}

func (*WherepartContext) IsWherepartContext() {}

func NewWherepartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WherepartContext {
	var p = new(WherepartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_wherepart

	return p
}

func (s *WherepartContext) GetParser() antlr.Parser { return s.parser }

func (s *WherepartContext) AllDef_() []IDef_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDef_Context)(nil)).Elem())
	var tst = make([]IDef_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDef_Context)
		}
	}

	return tst
}

func (s *WherepartContext) Def_(i int) IDef_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDef_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDef_Context)
}

func (s *WherepartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WherepartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WherepartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterWherepart(s)
	}
}

func (s *WherepartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitWherepart(s)
	}
}

func (p *orwellParser) Wherepart() (localctx IWherepartContext) {
	this := p
	_ = this

	localctx = NewWherepartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, orwellParserRULE_wherepart)

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
		p.SetState(333)
		p.Match(orwellParserT__21)
	}
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(334)
				p.Def_()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}

	return localctx
}

// IPatContext is an interface to support dynamic dispatch.
type IPatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatContext differentiates from other interfaces.
	IsPatContext()
}

type PatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatContext() *PatContext {
	var p = new(PatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_pat
	return p
}

func (*PatContext) IsPatContext() {}

func NewPatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatContext {
	var p = new(PatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_pat

	return p
}

func (s *PatContext) GetParser() antlr.Parser { return s.parser }

func (s *PatContext) Pat1() IPat1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPat1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPat1Context)
}

func (s *PatContext) Infix() IInfixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixContext)
}

func (s *PatContext) Pat() IPatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatContext)
}

func (s *PatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterPat(s)
	}
}

func (s *PatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitPat(s)
	}
}

func (p *orwellParser) Pat() (localctx IPatContext) {
	this := p
	_ = this

	localctx = NewPatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, orwellParserRULE_pat)
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
		p.SetState(339)
		p.Pat1()
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == orwellParserOP {
		{
			p.SetState(340)
			p.Infix()
		}
		{
			p.SetState(341)
			p.Pat()
		}

	}

	return localctx
}

// IPat1Context is an interface to support dynamic dispatch.
type IPat1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPat1Context differentiates from other interfaces.
	IsPat1Context()
}

type Pat1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPat1Context() *Pat1Context {
	var p = new(Pat1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_pat1
	return p
}

func (*Pat1Context) IsPat1Context() {}

func NewPat1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pat1Context {
	var p = new(Pat1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_pat1

	return p
}

func (s *Pat1Context) GetParser() antlr.Parser { return s.parser }

func (s *Pat1Context) Prefix() IPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *Pat1Context) Pat1() IPat1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPat1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPat1Context)
}

func (s *Pat1Context) Pat2() IPat2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPat2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPat2Context)
}

func (s *Pat1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pat1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pat1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterPat1(s)
	}
}

func (s *Pat1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitPat1(s)
	}
}

func (p *orwellParser) Pat1() (localctx IPat1Context) {
	this := p
	_ = this

	localctx = NewPat1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, orwellParserRULE_pat1)

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

	p.SetState(349)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case orwellParserOP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(345)
			p.Prefix()
		}
		{
			p.SetState(346)
			p.Pat1()
		}

	case orwellParserT__5, orwellParserT__7, orwellParserINTEGER, orwellParserSTRING, orwellParserID, orwellParserCHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(348)
			p.Pat2()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPat2Context is an interface to support dynamic dispatch.
type IPat2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPat2Context differentiates from other interfaces.
	IsPat2Context()
}

type Pat2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPat2Context() *Pat2Context {
	var p = new(Pat2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_pat2
	return p
}

func (*Pat2Context) IsPat2Context() {}

func NewPat2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Pat2Context {
	var p = new(Pat2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_pat2

	return p
}

func (s *Pat2Context) GetParser() antlr.Parser { return s.parser }

func (s *Pat2Context) AllPatprimary() []IPatprimaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPatprimaryContext)(nil)).Elem())
	var tst = make([]IPatprimaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPatprimaryContext)
		}
	}

	return tst
}

func (s *Pat2Context) Patprimary(i int) IPatprimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatprimaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPatprimaryContext)
}

func (s *Pat2Context) Patprimaryname() IPatprimarynameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatprimarynameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatprimarynameContext)
}

func (s *Pat2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Pat2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Pat2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterPat2(s)
	}
}

func (s *Pat2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitPat2(s)
	}
}

func (p *orwellParser) Pat2() (localctx IPat2Context) {
	this := p
	_ = this

	localctx = NewPat2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, orwellParserRULE_pat2)
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

	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(351)
			p.Patprimary()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(352)
			p.Patprimaryname()
		}
		p.SetState(356)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(orwellParserT__5-6))|(1<<(orwellParserT__7-6))|(1<<(orwellParserINTEGER-6))|(1<<(orwellParserSTRING-6))|(1<<(orwellParserID-6))|(1<<(orwellParserCHAR-6)))) != 0 {
			{
				p.SetState(353)
				p.Patprimary()
			}

			p.SetState(358)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IPatprimarynameContext is an interface to support dynamic dispatch.
type IPatprimarynameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatprimarynameContext differentiates from other interfaces.
	IsPatprimarynameContext()
}

type PatprimarynameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatprimarynameContext() *PatprimarynameContext {
	var p = new(PatprimarynameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_patprimaryname
	return p
}

func (*PatprimarynameContext) IsPatprimarynameContext() {}

func NewPatprimarynameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatprimarynameContext {
	var p = new(PatprimarynameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_patprimaryname

	return p
}

func (s *PatprimarynameContext) GetParser() antlr.Parser { return s.parser }

func (s *PatprimarynameContext) Var_() IVar_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *PatprimarynameContext) Pat() IPatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatContext)
}

func (s *PatprimarynameContext) Patsection() IPatsectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatsectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatsectionContext)
}

func (s *PatprimarynameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatprimarynameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatprimarynameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterPatprimaryname(s)
	}
}

func (s *PatprimarynameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitPatprimaryname(s)
	}
}

func (p *orwellParser) Patprimaryname() (localctx IPatprimarynameContext) {
	this := p
	_ = this

	localctx = NewPatprimarynameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, orwellParserRULE_patprimaryname)

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

	p.SetState(369)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case orwellParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(361)
			p.Var_()
		}

	case orwellParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(362)
			p.Match(orwellParserT__5)
		}
		p.SetState(365)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(363)
				p.Pat()
			}

		case 2:
			{
				p.SetState(364)
				p.Patsection()
			}

		}
		{
			p.SetState(367)
			p.Match(orwellParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPatprimaryContext is an interface to support dynamic dispatch.
type IPatprimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatprimaryContext differentiates from other interfaces.
	IsPatprimaryContext()
}

type PatprimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatprimaryContext() *PatprimaryContext {
	var p = new(PatprimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_patprimary
	return p
}

func (*PatprimaryContext) IsPatprimaryContext() {}

func NewPatprimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatprimaryContext {
	var p = new(PatprimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_patprimary

	return p
}

func (s *PatprimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PatprimaryContext) Patprimaryname() IPatprimarynameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatprimarynameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatprimarynameContext)
}

func (s *PatprimaryContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *PatprimaryContext) Pattuple() IPattupleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPattupleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPattupleContext)
}

func (s *PatprimaryContext) Patlist() IPatlistContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatlistContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatlistContext)
}

func (s *PatprimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatprimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatprimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterPatprimary(s)
	}
}

func (s *PatprimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitPatprimary(s)
	}
}

func (p *orwellParser) Patprimary() (localctx IPatprimaryContext) {
	this := p
	_ = this

	localctx = NewPatprimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, orwellParserRULE_patprimary)

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

	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			p.Patprimaryname()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.Literal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.Pattuple()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(374)
			p.Patlist()
		}

	}

	return localctx
}

// IPatsectionContext is an interface to support dynamic dispatch.
type IPatsectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatsectionContext differentiates from other interfaces.
	IsPatsectionContext()
}

type PatsectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatsectionContext() *PatsectionContext {
	var p = new(PatsectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_patsection
	return p
}

func (*PatsectionContext) IsPatsectionContext() {}

func NewPatsectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatsectionContext {
	var p = new(PatsectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_patsection

	return p
}

func (s *PatsectionContext) GetParser() antlr.Parser { return s.parser }

func (s *PatsectionContext) Prefix() IPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *PatsectionContext) Infix() IInfixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixContext)
}

func (s *PatsectionContext) Pat1() IPat1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPat1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPat1Context)
}

func (s *PatsectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatsectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatsectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterPatsection(s)
	}
}

func (s *PatsectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitPatsection(s)
	}
}

func (p *orwellParser) Patsection() (localctx IPatsectionContext) {
	this := p
	_ = this

	localctx = NewPatsectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, orwellParserRULE_patsection)

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

	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(377)
			p.Prefix()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(378)
			p.Infix()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(379)
			p.Infix()
		}
		{
			p.SetState(380)
			p.Pat1()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(382)
			p.Pat1()
		}
		{
			p.SetState(383)
			p.Infix()
		}

	}

	return localctx
}

// IPattupleContext is an interface to support dynamic dispatch.
type IPattupleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPattupleContext differentiates from other interfaces.
	IsPattupleContext()
}

type PattupleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPattupleContext() *PattupleContext {
	var p = new(PattupleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_pattuple
	return p
}

func (*PattupleContext) IsPattupleContext() {}

func NewPattupleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PattupleContext {
	var p = new(PattupleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_pattuple

	return p
}

func (s *PattupleContext) GetParser() antlr.Parser { return s.parser }

func (s *PattupleContext) AllPat() []IPatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPatContext)(nil)).Elem())
	var tst = make([]IPatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPatContext)
		}
	}

	return tst
}

func (s *PattupleContext) Pat(i int) IPatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPatContext)
}

func (s *PattupleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PattupleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PattupleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterPattuple(s)
	}
}

func (s *PattupleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitPattuple(s)
	}
}

func (p *orwellParser) Pattuple() (localctx IPattupleContext) {
	this := p
	_ = this

	localctx = NewPattupleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, orwellParserRULE_pattuple)
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
		p.SetState(387)
		p.Match(orwellParserT__5)
	}
	{
		p.SetState(388)
		p.Pat()
	}
	{
		p.SetState(389)
		p.Match(orwellParserT__3)
	}
	{
		p.SetState(390)
		p.Pat()
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == orwellParserT__3 {
		{
			p.SetState(391)
			p.Match(orwellParserT__3)
		}
		{
			p.SetState(392)
			p.Pat()
		}

		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(398)
		p.Match(orwellParserT__6)
	}

	return localctx
}

// IPatlistContext is an interface to support dynamic dispatch.
type IPatlistContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatlistContext differentiates from other interfaces.
	IsPatlistContext()
}

type PatlistContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatlistContext() *PatlistContext {
	var p = new(PatlistContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_patlist
	return p
}

func (*PatlistContext) IsPatlistContext() {}

func NewPatlistContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatlistContext {
	var p = new(PatlistContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_patlist

	return p
}

func (s *PatlistContext) GetParser() antlr.Parser { return s.parser }

func (s *PatlistContext) AllPat() []IPatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPatContext)(nil)).Elem())
	var tst = make([]IPatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPatContext)
		}
	}

	return tst
}

func (s *PatlistContext) Pat(i int) IPatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPatContext)
}

func (s *PatlistContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatlistContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatlistContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterPatlist(s)
	}
}

func (s *PatlistContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitPatlist(s)
	}
}

func (p *orwellParser) Patlist() (localctx IPatlistContext) {
	this := p
	_ = this

	localctx = NewPatlistContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, orwellParserRULE_patlist)
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
		p.SetState(400)
		p.Match(orwellParserT__7)
	}
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(orwellParserT__5-6))|(1<<(orwellParserT__7-6))|(1<<(orwellParserINTEGER-6))|(1<<(orwellParserSTRING-6))|(1<<(orwellParserOP-6))|(1<<(orwellParserID-6))|(1<<(orwellParserCHAR-6)))) != 0 {
		{
			p.SetState(401)
			p.Pat()
		}
		p.SetState(406)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == orwellParserT__3 {
			{
				p.SetState(402)
				p.Match(orwellParserT__3)
			}
			{
				p.SetState(403)
				p.Pat()
			}

			p.SetState(408)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(411)
		p.Match(orwellParserT__8)
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
	p.RuleIndex = orwellParserRULE_term
	return p
}

func (*TermContext) IsTermContext() {}

func NewTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TermContext {
	var p = new(TermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_term

	return p
}

func (s *TermContext) GetParser() antlr.Parser { return s.parser }

func (s *TermContext) Term1() ITerm1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerm1Context)
}

func (s *TermContext) Infix() IInfixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixContext)
}

func (s *TermContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *TermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTerm(s)
	}
}

func (s *TermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTerm(s)
	}
}

func (p *orwellParser) Term() (localctx ITermContext) {
	this := p
	_ = this

	localctx = NewTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, orwellParserRULE_term)

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
		p.SetState(413)
		p.Term1()
	}
	p.SetState(417)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(414)
			p.Infix()
		}
		{
			p.SetState(415)
			p.Term()
		}

	}

	return localctx
}

// ITerm1Context is an interface to support dynamic dispatch.
type ITerm1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerm1Context differentiates from other interfaces.
	IsTerm1Context()
}

type Term1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerm1Context() *Term1Context {
	var p = new(Term1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_term1
	return p
}

func (*Term1Context) IsTerm1Context() {}

func NewTerm1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Term1Context {
	var p = new(Term1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_term1

	return p
}

func (s *Term1Context) GetParser() antlr.Parser { return s.parser }

func (s *Term1Context) Prefix() IPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *Term1Context) Term1() ITerm1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerm1Context)
}

func (s *Term1Context) Term2() ITerm2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerm2Context)
}

func (s *Term1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Term1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Term1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTerm1(s)
	}
}

func (s *Term1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTerm1(s)
	}
}

func (p *orwellParser) Term1() (localctx ITerm1Context) {
	this := p
	_ = this

	localctx = NewTerm1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, orwellParserRULE_term1)

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

	p.SetState(423)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case orwellParserOP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(419)
			p.Prefix()
		}
		{
			p.SetState(420)
			p.Term1()
		}

	case orwellParserT__5, orwellParserT__7, orwellParserINTEGER, orwellParserFLOAT, orwellParserSTRING, orwellParserID, orwellParserCHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(422)
			p.Term2()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = orwellParserRULE_term2
	return p
}

func (*Term2Context) IsTerm2Context() {}

func NewTerm2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Term2Context {
	var p = new(Term2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_term2

	return p
}

func (s *Term2Context) GetParser() antlr.Parser { return s.parser }

func (s *Term2Context) AllPrimary() []IPrimaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimaryContext)(nil)).Elem())
	var tst = make([]IPrimaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimaryContext)
		}
	}

	return tst
}

func (s *Term2Context) Primary(i int) IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *Term2Context) Primaryname() IPrimarynameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimarynameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimarynameContext)
}

func (s *Term2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Term2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Term2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTerm2(s)
	}
}

func (s *Term2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTerm2(s)
	}
}

func (p *orwellParser) Term2() (localctx ITerm2Context) {
	this := p
	_ = this

	localctx = NewTerm2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, orwellParserRULE_term2)

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

	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(425)
			p.Primary()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(426)
			p.Primaryname()
		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(427)
					p.Primary()
				}

			}
			p.SetState(432)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IPrimarynameContext is an interface to support dynamic dispatch.
type IPrimarynameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimarynameContext differentiates from other interfaces.
	IsPrimarynameContext()
}

type PrimarynameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimarynameContext() *PrimarynameContext {
	var p = new(PrimarynameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_primaryname
	return p
}

func (*PrimarynameContext) IsPrimarynameContext() {}

func NewPrimarynameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimarynameContext {
	var p = new(PrimarynameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_primaryname

	return p
}

func (s *PrimarynameContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimarynameContext) Var_() IVar_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVar_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *PrimarynameContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *PrimarynameContext) Section() ISectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISectionContext)
}

func (s *PrimarynameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimarynameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimarynameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterPrimaryname(s)
	}
}

func (s *PrimarynameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitPrimaryname(s)
	}
}

func (p *orwellParser) Primaryname() (localctx IPrimarynameContext) {
	this := p
	_ = this

	localctx = NewPrimarynameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, orwellParserRULE_primaryname)

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

	p.SetState(443)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case orwellParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)
			p.Var_()
		}

	case orwellParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(436)
			p.Match(orwellParserT__5)
		}
		p.SetState(439)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(437)
				p.Term()
			}

		case 2:
			{
				p.SetState(438)
				p.Section()
			}

		}
		{
			p.SetState(441)
			p.Match(orwellParserT__6)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) Primaryname() IPrimarynameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimarynameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimarynameContext)
}

func (s *PrimaryContext) Fliteral() IFliteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFliteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFliteralContext)
}

func (s *PrimaryContext) Tuple_() ITuple_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITuple_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITuple_Context)
}

func (s *PrimaryContext) List_() IList_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_Context)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (p *orwellParser) Primary() (localctx IPrimaryContext) {
	this := p
	_ = this

	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, orwellParserRULE_primary)

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

	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(445)
			p.Primaryname()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(446)
			p.Fliteral()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(447)
			p.Tuple_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(448)
			p.List_()
		}

	}

	return localctx
}

// ISectionContext is an interface to support dynamic dispatch.
type ISectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionContext differentiates from other interfaces.
	IsSectionContext()
}

type SectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionContext() *SectionContext {
	var p = new(SectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_section
	return p
}

func (*SectionContext) IsSectionContext() {}

func NewSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionContext {
	var p = new(SectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_section

	return p
}

func (s *SectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionContext) Prefix() IPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *SectionContext) Infix() IInfixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInfixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInfixContext)
}

func (s *SectionContext) Term1() ITerm1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerm1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerm1Context)
}

func (s *SectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterSection(s)
	}
}

func (s *SectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitSection(s)
	}
}

func (p *orwellParser) Section() (localctx ISectionContext) {
	this := p
	_ = this

	localctx = NewSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, orwellParserRULE_section)

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

	p.SetState(459)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(451)
			p.Prefix()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(452)
			p.Infix()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(453)
			p.Infix()
		}
		{
			p.SetState(454)
			p.Term1()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(456)
			p.Term1()
		}
		{
			p.SetState(457)
			p.Infix()
		}

	}

	return localctx
}

// IList_Context is an interface to support dynamic dispatch.
type IList_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_Context differentiates from other interfaces.
	IsList_Context()
}

type List_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_Context() *List_Context {
	var p = new(List_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_list_
	return p
}

func (*List_Context) IsList_Context() {}

func NewList_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_Context {
	var p = new(List_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_list_

	return p
}

func (s *List_Context) GetParser() antlr.Parser { return s.parser }

func (s *List_Context) Listform() IListformContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListformContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListformContext)
}

func (s *List_Context) Upto() IUptoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUptoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUptoContext)
}

func (s *List_Context) Comp() ICompContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompContext)
}

func (s *List_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterList_(s)
	}
}

func (s *List_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitList_(s)
	}
}

func (p *orwellParser) List_() (localctx IList_Context) {
	this := p
	_ = this

	localctx = NewList_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, orwellParserRULE_list_)

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

	p.SetState(464)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(461)
			p.Listform()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(462)
			p.Upto()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(463)
			p.Comp()
		}

	}

	return localctx
}

// ITuple_Context is an interface to support dynamic dispatch.
type ITuple_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTuple_Context differentiates from other interfaces.
	IsTuple_Context()
}

type Tuple_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTuple_Context() *Tuple_Context {
	var p = new(Tuple_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_tuple_
	return p
}

func (*Tuple_Context) IsTuple_Context() {}

func NewTuple_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tuple_Context {
	var p = new(Tuple_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_tuple_

	return p
}

func (s *Tuple_Context) GetParser() antlr.Parser { return s.parser }

func (s *Tuple_Context) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *Tuple_Context) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *Tuple_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tuple_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tuple_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTuple_(s)
	}
}

func (s *Tuple_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTuple_(s)
	}
}

func (p *orwellParser) Tuple_() (localctx ITuple_Context) {
	this := p
	_ = this

	localctx = NewTuple_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, orwellParserRULE_tuple_)
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
		p.SetState(466)
		p.Match(orwellParserT__5)
	}
	{
		p.SetState(467)
		p.Term()
	}
	{
		p.SetState(468)
		p.Match(orwellParserT__3)
	}
	{
		p.SetState(469)
		p.Term()
	}
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == orwellParserT__3 {
		{
			p.SetState(470)
			p.Match(orwellParserT__3)
		}
		{
			p.SetState(471)
			p.Term()
		}

		p.SetState(476)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(477)
		p.Match(orwellParserT__6)
	}

	return localctx
}

// IListformContext is an interface to support dynamic dispatch.
type IListformContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListformContext differentiates from other interfaces.
	IsListformContext()
}

type ListformContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListformContext() *ListformContext {
	var p = new(ListformContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_listform
	return p
}

func (*ListformContext) IsListformContext() {}

func NewListformContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListformContext {
	var p = new(ListformContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_listform

	return p
}

func (s *ListformContext) GetParser() antlr.Parser { return s.parser }

func (s *ListformContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *ListformContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *ListformContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListformContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListformContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterListform(s)
	}
}

func (s *ListformContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitListform(s)
	}
}

func (p *orwellParser) Listform() (localctx IListformContext) {
	this := p
	_ = this

	localctx = NewListformContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, orwellParserRULE_listform)
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
		p.SetState(479)
		p.Match(orwellParserT__7)
	}
	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(orwellParserT__5-6))|(1<<(orwellParserT__7-6))|(1<<(orwellParserINTEGER-6))|(1<<(orwellParserFLOAT-6))|(1<<(orwellParserSTRING-6))|(1<<(orwellParserOP-6))|(1<<(orwellParserID-6))|(1<<(orwellParserCHAR-6)))) != 0 {
		{
			p.SetState(480)
			p.Term()
		}
		p.SetState(485)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == orwellParserT__3 {
			{
				p.SetState(481)
				p.Match(orwellParserT__3)
			}
			{
				p.SetState(482)
				p.Term()
			}

			p.SetState(487)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(490)
		p.Match(orwellParserT__8)
	}

	return localctx
}

// IUptoContext is an interface to support dynamic dispatch.
type IUptoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUptoContext differentiates from other interfaces.
	IsUptoContext()
}

type UptoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUptoContext() *UptoContext {
	var p = new(UptoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_upto
	return p
}

func (*UptoContext) IsUptoContext() {}

func NewUptoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UptoContext {
	var p = new(UptoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_upto

	return p
}

func (s *UptoContext) GetParser() antlr.Parser { return s.parser }

func (s *UptoContext) AllTerm() []ITermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITermContext)(nil)).Elem())
	var tst = make([]ITermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITermContext)
		}
	}

	return tst
}

func (s *UptoContext) Term(i int) ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *UptoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UptoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UptoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterUpto(s)
	}
}

func (s *UptoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitUpto(s)
	}
}

func (p *orwellParser) Upto() (localctx IUptoContext) {
	this := p
	_ = this

	localctx = NewUptoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, orwellParserRULE_upto)
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
		p.SetState(492)
		p.Match(orwellParserT__7)
	}
	{
		p.SetState(493)
		p.Term()
	}
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == orwellParserT__3 {
		{
			p.SetState(494)
			p.Match(orwellParserT__3)
		}
		{
			p.SetState(495)
			p.Term()
		}

	}
	{
		p.SetState(498)
		p.Match(orwellParserT__22)
	}
	p.SetState(500)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(orwellParserT__5-6))|(1<<(orwellParserT__7-6))|(1<<(orwellParserINTEGER-6))|(1<<(orwellParserFLOAT-6))|(1<<(orwellParserSTRING-6))|(1<<(orwellParserOP-6))|(1<<(orwellParserID-6))|(1<<(orwellParserCHAR-6)))) != 0 {
		{
			p.SetState(499)
			p.Term()
		}

	}
	{
		p.SetState(502)
		p.Match(orwellParserT__8)
	}

	return localctx
}

// ICompContext is an interface to support dynamic dispatch.
type ICompContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompContext differentiates from other interfaces.
	IsCompContext()
}

type CompContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompContext() *CompContext {
	var p = new(CompContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_comp
	return p
}

func (*CompContext) IsCompContext() {}

func NewCompContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompContext {
	var p = new(CompContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_comp

	return p
}

func (s *CompContext) GetParser() antlr.Parser { return s.parser }

func (s *CompContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *CompContext) AllQualifier() []IQualifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQualifierContext)(nil)).Elem())
	var tst = make([]IQualifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQualifierContext)
		}
	}

	return tst
}

func (s *CompContext) Qualifier(i int) IQualifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQualifierContext)
}

func (s *CompContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterComp(s)
	}
}

func (s *CompContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitComp(s)
	}
}

func (p *orwellParser) Comp() (localctx ICompContext) {
	this := p
	_ = this

	localctx = NewCompContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, orwellParserRULE_comp)
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
		p.SetState(504)
		p.Match(orwellParserT__7)
	}
	{
		p.SetState(505)
		p.Term()
	}
	{
		p.SetState(506)
		p.Match(orwellParserT__2)
	}
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-6)&-(0x1f+1)) == 0 && ((1<<uint((_la-6)))&((1<<(orwellParserT__5-6))|(1<<(orwellParserT__7-6))|(1<<(orwellParserINTEGER-6))|(1<<(orwellParserFLOAT-6))|(1<<(orwellParserSTRING-6))|(1<<(orwellParserOP-6))|(1<<(orwellParserID-6))|(1<<(orwellParserCHAR-6)))) != 0 {
		{
			p.SetState(507)
			p.Qualifier()
		}
		p.SetState(512)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == orwellParserT__23 {
			{
				p.SetState(508)
				p.Match(orwellParserT__23)
			}
			{
				p.SetState(509)
				p.Qualifier()
			}

			p.SetState(514)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(517)
		p.Match(orwellParserT__8)
	}

	return localctx
}

// IQualifierContext is an interface to support dynamic dispatch.
type IQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifierContext differentiates from other interfaces.
	IsQualifierContext()
}

type QualifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifierContext() *QualifierContext {
	var p = new(QualifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_qualifier
	return p
}

func (*QualifierContext) IsQualifierContext() {}

func NewQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifierContext {
	var p = new(QualifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_qualifier

	return p
}

func (s *QualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifierContext) Term() ITermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITermContext)
}

func (s *QualifierContext) Pat() IPatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatContext)
}

func (s *QualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterQualifier(s)
	}
}

func (s *QualifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitQualifier(s)
	}
}

func (p *orwellParser) Qualifier() (localctx IQualifierContext) {
	this := p
	_ = this

	localctx = NewQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, orwellParserRULE_qualifier)

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

	p.SetState(524)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(519)
			p.Term()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(520)
			p.Pat()
		}
		{
			p.SetState(521)
			p.Match(orwellParserT__24)
		}
		{
			p.SetState(522)
			p.Term()
		}

	}

	return localctx
}

// IFliteralContext is an interface to support dynamic dispatch.
type IFliteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFliteralContext differentiates from other interfaces.
	IsFliteralContext()
}

type FliteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFliteralContext() *FliteralContext {
	var p = new(FliteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_fliteral
	return p
}

func (*FliteralContext) IsFliteralContext() {}

func NewFliteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FliteralContext {
	var p = new(FliteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_fliteral

	return p
}

func (s *FliteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FliteralContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(orwellParserFLOAT, 0)
}

func (s *FliteralContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *FliteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FliteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FliteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterFliteral(s)
	}
}

func (s *FliteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitFliteral(s)
	}
}

func (p *orwellParser) Fliteral() (localctx IFliteralContext) {
	this := p
	_ = this

	localctx = NewFliteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, orwellParserRULE_fliteral)

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

	p.SetState(528)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case orwellParserFLOAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(526)
			p.Match(orwellParserFLOAT)
		}

	case orwellParserINTEGER, orwellParserSTRING, orwellParserCHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(527)
			p.Literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(orwellParserINTEGER, 0)
}

func (s *LiteralContext) CHAR() antlr.TerminalNode {
	return s.GetToken(orwellParserCHAR, 0)
}

func (s *LiteralContext) STRING() antlr.TerminalNode {
	return s.GetToken(orwellParserSTRING, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *orwellParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, orwellParserRULE_literal)
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
		p.SetState(530)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-26)&-(0x1f+1)) == 0 && ((1<<uint((_la-26)))&((1<<(orwellParserINTEGER-26))|(1<<(orwellParserSTRING-26))|(1<<(orwellParserCHAR-26)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IInfixContext is an interface to support dynamic dispatch.
type IInfixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInfixContext differentiates from other interfaces.
	IsInfixContext()
}

type InfixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInfixContext() *InfixContext {
	var p = new(InfixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_infix
	return p
}

func (*InfixContext) IsInfixContext() {}

func NewInfixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InfixContext {
	var p = new(InfixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_infix

	return p
}

func (s *InfixContext) GetParser() antlr.Parser { return s.parser }

func (s *InfixContext) OP() antlr.TerminalNode {
	return s.GetToken(orwellParserOP, 0)
}

func (s *InfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InfixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterInfix(s)
	}
}

func (s *InfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitInfix(s)
	}
}

func (p *orwellParser) Infix() (localctx IInfixContext) {
	this := p
	_ = this

	localctx = NewInfixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, orwellParserRULE_infix)

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
		p.SetState(532)
		p.Match(orwellParserOP)
	}

	return localctx
}

// IPrefixContext is an interface to support dynamic dispatch.
type IPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixContext differentiates from other interfaces.
	IsPrefixContext()
}

type PrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixContext() *PrefixContext {
	var p = new(PrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_prefix
	return p
}

func (*PrefixContext) IsPrefixContext() {}

func NewPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixContext {
	var p = new(PrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_prefix

	return p
}

func (s *PrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixContext) OP() antlr.TerminalNode {
	return s.GetToken(orwellParserOP, 0)
}

func (s *PrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterPrefix(s)
	}
}

func (s *PrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitPrefix(s)
	}
}

func (p *orwellParser) Prefix() (localctx IPrefixContext) {
	this := p
	_ = this

	localctx = NewPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, orwellParserRULE_prefix)

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
		p.SetState(534)
		p.Match(orwellParserOP)
	}

	return localctx
}

// ITynameContext is an interface to support dynamic dispatch.
type ITynameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTynameContext differentiates from other interfaces.
	IsTynameContext()
}

type TynameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTynameContext() *TynameContext {
	var p = new(TynameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_tyname
	return p
}

func (*TynameContext) IsTynameContext() {}

func NewTynameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TynameContext {
	var p = new(TynameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_tyname

	return p
}

func (s *TynameContext) GetParser() antlr.Parser { return s.parser }

func (s *TynameContext) ID() antlr.TerminalNode {
	return s.GetToken(orwellParserID, 0)
}

func (s *TynameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TynameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TynameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTyname(s)
	}
}

func (s *TynameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTyname(s)
	}
}

func (p *orwellParser) Tyname() (localctx ITynameContext) {
	this := p
	_ = this

	localctx = NewTynameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, orwellParserRULE_tyname)

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
		p.SetState(536)
		p.Match(orwellParserID)
	}

	return localctx
}

// ITyvarContext is an interface to support dynamic dispatch.
type ITyvarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTyvarContext differentiates from other interfaces.
	IsTyvarContext()
}

type TyvarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTyvarContext() *TyvarContext {
	var p = new(TyvarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_tyvar
	return p
}

func (*TyvarContext) IsTyvarContext() {}

func NewTyvarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TyvarContext {
	var p = new(TyvarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_tyvar

	return p
}

func (s *TyvarContext) GetParser() antlr.Parser { return s.parser }

func (s *TyvarContext) ID() antlr.TerminalNode {
	return s.GetToken(orwellParserID, 0)
}

func (s *TyvarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TyvarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TyvarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterTyvar(s)
	}
}

func (s *TyvarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitTyvar(s)
	}
}

func (p *orwellParser) Tyvar() (localctx ITyvarContext) {
	this := p
	_ = this

	localctx = NewTyvarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, orwellParserRULE_tyvar)

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
		p.SetState(538)
		p.Match(orwellParserID)
	}

	return localctx
}

// IConContext is an interface to support dynamic dispatch.
type IConContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConContext differentiates from other interfaces.
	IsConContext()
}

type ConContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConContext() *ConContext {
	var p = new(ConContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_con
	return p
}

func (*ConContext) IsConContext() {}

func NewConContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConContext {
	var p = new(ConContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_con

	return p
}

func (s *ConContext) GetParser() antlr.Parser { return s.parser }

func (s *ConContext) ID() antlr.TerminalNode {
	return s.GetToken(orwellParserID, 0)
}

func (s *ConContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterCon(s)
	}
}

func (s *ConContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitCon(s)
	}
}

func (p *orwellParser) Con() (localctx IConContext) {
	this := p
	_ = this

	localctx = NewConContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, orwellParserRULE_con)

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
		p.SetState(540)
		p.Match(orwellParserID)
	}

	return localctx
}

// IVar_Context is an interface to support dynamic dispatch.
type IVar_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVar_Context differentiates from other interfaces.
	IsVar_Context()
}

type Var_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_Context() *Var_Context {
	var p = new(Var_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = orwellParserRULE_var_
	return p
}

func (*Var_Context) IsVar_Context() {}

func NewVar_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_Context {
	var p = new(Var_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = orwellParserRULE_var_

	return p
}

func (s *Var_Context) GetParser() antlr.Parser { return s.parser }

func (s *Var_Context) ID() antlr.TerminalNode {
	return s.GetToken(orwellParserID, 0)
}

func (s *Var_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.EnterVar_(s)
	}
}

func (s *Var_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(orwellListener); ok {
		listenerT.ExitVar_(s)
	}
}

func (p *orwellParser) Var_() (localctx IVar_Context) {
	this := p
	_ = this

	localctx = NewVar_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, orwellParserRULE_var_)

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
		p.SetState(542)
		p.Match(orwellParserID)
	}

	return localctx
}
