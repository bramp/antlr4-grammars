// Code generated from pike.g4 by ANTLR 4.9.3. DO NOT EDIT.

package pike // pike
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 95, 592,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 3, 2, 7, 2, 106, 10, 2, 12, 2, 14,
	2, 109, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 118, 10,
	3, 3, 4, 5, 4, 121, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 5, 5, 128, 10,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 134, 10, 5, 3, 5, 3, 5, 3, 6, 5, 6, 139,
	10, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 5, 7, 149, 10, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 5, 8, 159, 10, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 7, 9, 168, 10, 9, 12, 9, 14, 9, 171,
	11, 9, 3, 10, 7, 10, 174, 10, 10, 12, 10, 14, 10, 177, 11, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 182, 10, 10, 3, 11, 5, 11, 185, 10, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 7, 12, 194, 10, 12, 12, 12, 14, 12,
	197, 11, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 5, 14, 204, 10, 14, 3,
	14, 3, 14, 5, 14, 208, 10, 14, 3, 15, 3, 15, 5, 15, 212, 10, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 7, 17, 222, 10, 17, 12,
	17, 14, 17, 225, 11, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5,
	18, 244, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 250, 10, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 5, 22, 269, 10, 22, 3, 22, 3, 22,
	5, 22, 273, 10, 22, 3, 22, 3, 22, 5, 22, 277, 10, 22, 3, 22, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	5, 24, 292, 10, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 29, 7, 29, 316, 10, 29, 12, 29, 14, 29, 319, 11,
	29, 3, 30, 3, 30, 3, 30, 7, 30, 324, 10, 30, 12, 30, 14, 30, 327, 11, 30,
	3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 337, 10,
	31, 3, 32, 3, 32, 3, 32, 7, 32, 342, 10, 32, 12, 32, 14, 32, 345, 11, 32,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 5, 33, 369, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 384, 10, 34, 3, 34,
	7, 34, 387, 10, 34, 12, 34, 14, 34, 390, 11, 34, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 402, 10, 35, 3, 35,
	3, 35, 5, 35, 406, 10, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 5,
	36, 414, 10, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 422,
	10, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 7, 38, 431, 10,
	38, 12, 38, 14, 38, 434, 11, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 39, 5, 39, 443, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 41, 3, 41, 3, 41, 7, 41, 454, 10, 41, 12, 41, 14, 41, 457, 11, 41, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44,
	3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 7, 44, 476, 10, 44, 12, 44, 14,
	44, 479, 11, 44, 5, 44, 481, 10, 44, 3, 44, 5, 44, 484, 10, 44, 3, 44,
	3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 7,
	47, 497, 10, 47, 12, 47, 14, 47, 500, 11, 47, 5, 47, 502, 10, 47, 3, 47,
	5, 47, 505, 10, 47, 3, 48, 5, 48, 508, 10, 48, 3, 48, 3, 48, 3, 49, 3,
	49, 5, 49, 514, 10, 49, 3, 49, 5, 49, 517, 10, 49, 3, 50, 3, 50, 3, 50,
	7, 50, 522, 10, 50, 12, 50, 14, 50, 525, 11, 50, 5, 50, 527, 10, 50, 3,
	50, 5, 50, 530, 10, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51,
	3, 51, 3, 51, 5, 51, 541, 10, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3,
	51, 3, 51, 5, 51, 550, 10, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 5, 51,
	557, 10, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 5, 51, 564, 10, 51, 3,
	51, 3, 51, 5, 51, 568, 10, 51, 5, 51, 570, 10, 51, 3, 51, 7, 51, 573, 10,
	51, 12, 51, 14, 51, 576, 11, 51, 3, 52, 3, 52, 3, 52, 3, 52, 7, 52, 582,
	10, 52, 12, 52, 14, 52, 585, 11, 52, 3, 52, 5, 52, 588, 10, 52, 3, 52,
	3, 52, 3, 52, 2, 2, 53, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
	28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62,
	64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98,
	100, 102, 2, 5, 3, 2, 16, 21, 4, 2, 11, 11, 34, 42, 4, 2, 10, 10, 44, 59,
	2, 639, 2, 107, 3, 2, 2, 2, 4, 117, 3, 2, 2, 2, 6, 120, 3, 2, 2, 2, 8,
	127, 3, 2, 2, 2, 10, 138, 3, 2, 2, 2, 12, 148, 3, 2, 2, 2, 14, 158, 3,
	2, 2, 2, 16, 164, 3, 2, 2, 2, 18, 175, 3, 2, 2, 2, 20, 184, 3, 2, 2, 2,
	22, 190, 3, 2, 2, 2, 24, 198, 3, 2, 2, 2, 26, 203, 3, 2, 2, 2, 28, 209,
	3, 2, 2, 2, 30, 217, 3, 2, 2, 2, 32, 219, 3, 2, 2, 2, 34, 243, 3, 2, 2,
	2, 36, 245, 3, 2, 2, 2, 38, 251, 3, 2, 2, 2, 40, 257, 3, 2, 2, 2, 42, 265,
	3, 2, 2, 2, 44, 281, 3, 2, 2, 2, 46, 287, 3, 2, 2, 2, 48, 295, 3, 2, 2,
	2, 50, 298, 3, 2, 2, 2, 52, 306, 3, 2, 2, 2, 54, 309, 3, 2, 2, 2, 56, 312,
	3, 2, 2, 2, 58, 325, 3, 2, 2, 2, 60, 330, 3, 2, 2, 2, 62, 343, 3, 2, 2,
	2, 64, 368, 3, 2, 2, 2, 66, 383, 3, 2, 2, 2, 68, 405, 3, 2, 2, 2, 70, 407,
	3, 2, 2, 2, 72, 415, 3, 2, 2, 2, 74, 423, 3, 2, 2, 2, 76, 442, 3, 2, 2,
	2, 78, 444, 3, 2, 2, 2, 80, 450, 3, 2, 2, 2, 82, 458, 3, 2, 2, 2, 84, 462,
	3, 2, 2, 2, 86, 466, 3, 2, 2, 2, 88, 487, 3, 2, 2, 2, 90, 489, 3, 2, 2,
	2, 92, 501, 3, 2, 2, 2, 94, 507, 3, 2, 2, 2, 96, 511, 3, 2, 2, 2, 98, 526,
	3, 2, 2, 2, 100, 569, 3, 2, 2, 2, 102, 577, 3, 2, 2, 2, 104, 106, 5, 4,
	3, 2, 105, 104, 3, 2, 2, 2, 106, 109, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2,
	107, 108, 3, 2, 2, 2, 108, 3, 3, 2, 2, 2, 109, 107, 3, 2, 2, 2, 110, 118,
	5, 6, 4, 2, 111, 118, 5, 8, 5, 2, 112, 118, 5, 10, 6, 2, 113, 118, 5, 12,
	7, 2, 114, 118, 5, 14, 8, 2, 115, 118, 5, 20, 11, 2, 116, 118, 5, 26, 14,
	2, 117, 110, 3, 2, 2, 2, 117, 111, 3, 2, 2, 2, 117, 112, 3, 2, 2, 2, 117,
	113, 3, 2, 2, 2, 117, 114, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 116,
	3, 2, 2, 2, 118, 5, 3, 2, 2, 2, 119, 121, 5, 30, 16, 2, 120, 119, 3, 2,
	2, 2, 120, 121, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2, 122, 123, 7, 3, 2, 2,
	123, 124, 5, 80, 41, 2, 124, 125, 7, 4, 2, 2, 125, 7, 3, 2, 2, 2, 126,
	128, 5, 30, 16, 2, 127, 126, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 129,
	3, 2, 2, 2, 129, 130, 7, 5, 2, 2, 130, 133, 5, 88, 45, 2, 131, 132, 7,
	6, 2, 2, 132, 134, 7, 89, 2, 2, 133, 131, 3, 2, 2, 2, 133, 134, 3, 2, 2,
	2, 134, 135, 3, 2, 2, 2, 135, 136, 7, 4, 2, 2, 136, 9, 3, 2, 2, 2, 137,
	139, 5, 30, 16, 2, 138, 137, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140,
	3, 2, 2, 2, 140, 141, 5, 100, 51, 2, 141, 142, 7, 89, 2, 2, 142, 143, 7,
	7, 2, 2, 143, 144, 5, 98, 50, 2, 144, 145, 7, 8, 2, 2, 145, 146, 7, 4,
	2, 2, 146, 11, 3, 2, 2, 2, 147, 149, 5, 30, 16, 2, 148, 147, 3, 2, 2, 2,
	148, 149, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 151, 5, 100, 51, 2, 151,
	152, 7, 89, 2, 2, 152, 153, 7, 7, 2, 2, 153, 154, 5, 98, 50, 2, 154, 155,
	7, 8, 2, 2, 155, 156, 5, 32, 17, 2, 156, 13, 3, 2, 2, 2, 157, 159, 5, 30,
	16, 2, 158, 157, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2,
	160, 161, 5, 100, 51, 2, 161, 162, 5, 16, 9, 2, 162, 163, 7, 4, 2, 2, 163,
	15, 3, 2, 2, 2, 164, 169, 5, 18, 10, 2, 165, 166, 7, 9, 2, 2, 166, 168,
	5, 18, 10, 2, 167, 165, 3, 2, 2, 2, 168, 171, 3, 2, 2, 2, 169, 167, 3,
	2, 2, 2, 169, 170, 3, 2, 2, 2, 170, 17, 3, 2, 2, 2, 171, 169, 3, 2, 2,
	2, 172, 174, 7, 10, 2, 2, 173, 172, 3, 2, 2, 2, 174, 177, 3, 2, 2, 2, 175,
	173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 178, 3, 2, 2, 2, 177, 175,
	3, 2, 2, 2, 178, 181, 7, 89, 2, 2, 179, 180, 7, 11, 2, 2, 180, 182, 5,
	58, 30, 2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 19, 3, 2, 2,
	2, 183, 185, 5, 30, 16, 2, 184, 183, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2,
	185, 186, 3, 2, 2, 2, 186, 187, 7, 12, 2, 2, 187, 188, 5, 22, 12, 2, 188,
	189, 7, 4, 2, 2, 189, 21, 3, 2, 2, 2, 190, 195, 5, 24, 13, 2, 191, 192,
	7, 9, 2, 2, 192, 194, 5, 24, 13, 2, 193, 191, 3, 2, 2, 2, 194, 197, 3,
	2, 2, 2, 195, 193, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2, 196, 23, 3, 2, 2,
	2, 197, 195, 3, 2, 2, 2, 198, 199, 7, 89, 2, 2, 199, 200, 7, 11, 2, 2,
	200, 201, 5, 58, 30, 2, 201, 25, 3, 2, 2, 2, 202, 204, 5, 30, 16, 2, 203,
	202, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 207,
	7, 13, 2, 2, 206, 208, 7, 4, 2, 2, 207, 206, 3, 2, 2, 2, 207, 208, 3, 2,
	2, 2, 208, 27, 3, 2, 2, 2, 209, 211, 7, 13, 2, 2, 210, 212, 7, 89, 2, 2,
	211, 210, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213,
	214, 7, 14, 2, 2, 214, 215, 5, 2, 2, 2, 215, 216, 7, 15, 2, 2, 216, 29,
	3, 2, 2, 2, 217, 218, 9, 2, 2, 2, 218, 31, 3, 2, 2, 2, 219, 223, 7, 14,
	2, 2, 220, 222, 5, 34, 18, 2, 221, 220, 3, 2, 2, 2, 222, 225, 3, 2, 2,
	2, 223, 221, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 226, 3, 2, 2, 2, 225,
	223, 3, 2, 2, 2, 226, 227, 7, 15, 2, 2, 227, 33, 3, 2, 2, 2, 228, 229,
	5, 58, 30, 2, 229, 230, 7, 4, 2, 2, 230, 244, 3, 2, 2, 2, 231, 244, 5,
	36, 19, 2, 232, 244, 5, 38, 20, 2, 233, 244, 5, 40, 21, 2, 234, 244, 5,
	42, 22, 2, 235, 244, 5, 44, 23, 2, 236, 244, 5, 46, 24, 2, 237, 244, 5,
	48, 25, 2, 238, 244, 5, 32, 17, 2, 239, 244, 5, 50, 26, 2, 240, 244, 5,
	52, 27, 2, 241, 244, 5, 54, 28, 2, 242, 244, 7, 4, 2, 2, 243, 228, 3, 2,
	2, 2, 243, 231, 3, 2, 2, 2, 243, 232, 3, 2, 2, 2, 243, 233, 3, 2, 2, 2,
	243, 234, 3, 2, 2, 2, 243, 235, 3, 2, 2, 2, 243, 236, 3, 2, 2, 2, 243,
	237, 3, 2, 2, 2, 243, 238, 3, 2, 2, 2, 243, 239, 3, 2, 2, 2, 243, 240,
	3, 2, 2, 2, 243, 241, 3, 2, 2, 2, 243, 242, 3, 2, 2, 2, 244, 35, 3, 2,
	2, 2, 245, 246, 7, 22, 2, 2, 246, 249, 5, 34, 18, 2, 247, 248, 7, 23, 2,
	2, 248, 250, 5, 34, 18, 2, 249, 247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2,
	250, 37, 3, 2, 2, 2, 251, 252, 7, 24, 2, 2, 252, 253, 7, 7, 2, 2, 253,
	254, 5, 56, 29, 2, 254, 255, 7, 8, 2, 2, 255, 256, 5, 34, 18, 2, 256, 39,
	3, 2, 2, 2, 257, 258, 7, 25, 2, 2, 258, 259, 5, 34, 18, 2, 259, 260, 5,
	38, 20, 2, 260, 261, 7, 7, 2, 2, 261, 262, 5, 56, 29, 2, 262, 263, 7, 8,
	2, 2, 263, 264, 7, 4, 2, 2, 264, 41, 3, 2, 2, 2, 265, 266, 7, 26, 2, 2,
	266, 268, 7, 7, 2, 2, 267, 269, 5, 56, 29, 2, 268, 267, 3, 2, 2, 2, 268,
	269, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 272, 7, 4, 2, 2, 271, 273,
	5, 56, 29, 2, 272, 271, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 274, 3,
	2, 2, 2, 274, 276, 7, 4, 2, 2, 275, 277, 5, 56, 29, 2, 276, 275, 3, 2,
	2, 2, 276, 277, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 279, 7, 8, 2, 2,
	279, 280, 5, 34, 18, 2, 280, 43, 3, 2, 2, 2, 281, 282, 7, 27, 2, 2, 282,
	283, 7, 7, 2, 2, 283, 284, 5, 56, 29, 2, 284, 285, 7, 8, 2, 2, 285, 286,
	5, 32, 17, 2, 286, 45, 3, 2, 2, 2, 287, 288, 7, 28, 2, 2, 288, 291, 5,
	56, 29, 2, 289, 290, 7, 29, 2, 2, 290, 292, 5, 56, 29, 2, 291, 289, 3,
	2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293, 294, 7, 6, 2,
	2, 294, 47, 3, 2, 2, 2, 295, 296, 7, 30, 2, 2, 296, 297, 7, 6, 2, 2, 297,
	49, 3, 2, 2, 2, 298, 299, 7, 31, 2, 2, 299, 300, 7, 7, 2, 2, 300, 301,
	5, 56, 29, 2, 301, 302, 7, 6, 2, 2, 302, 303, 5, 66, 34, 2, 303, 304, 7,
	8, 2, 2, 304, 305, 5, 34, 18, 2, 305, 51, 3, 2, 2, 2, 306, 307, 7, 32,
	2, 2, 307, 308, 7, 4, 2, 2, 308, 53, 3, 2, 2, 2, 309, 310, 7, 33, 2, 2,
	310, 311, 7, 4, 2, 2, 311, 55, 3, 2, 2, 2, 312, 317, 5, 58, 30, 2, 313,
	314, 7, 9, 2, 2, 314, 316, 5, 58, 30, 2, 315, 313, 3, 2, 2, 2, 316, 319,
	3, 2, 2, 2, 317, 315, 3, 2, 2, 2, 317, 318, 3, 2, 2, 2, 318, 57, 3, 2,
	2, 2, 319, 317, 3, 2, 2, 2, 320, 321, 5, 76, 39, 2, 321, 322, 9, 3, 2,
	2, 322, 324, 3, 2, 2, 2, 323, 320, 3, 2, 2, 2, 324, 327, 3, 2, 2, 2, 325,
	323, 3, 2, 2, 2, 325, 326, 3, 2, 2, 2, 326, 328, 3, 2, 2, 2, 327, 325,
	3, 2, 2, 2, 328, 329, 5, 60, 31, 2, 329, 59, 3, 2, 2, 2, 330, 336, 5, 62,
	32, 2, 331, 332, 7, 43, 2, 2, 332, 333, 5, 60, 31, 2, 333, 334, 7, 6, 2,
	2, 334, 335, 5, 60, 31, 2, 335, 337, 3, 2, 2, 2, 336, 331, 3, 2, 2, 2,
	336, 337, 3, 2, 2, 2, 337, 61, 3, 2, 2, 2, 338, 339, 5, 64, 33, 2, 339,
	340, 9, 4, 2, 2, 340, 342, 3, 2, 2, 2, 341, 338, 3, 2, 2, 2, 342, 345,
	3, 2, 2, 2, 343, 341, 3, 2, 2, 2, 343, 344, 3, 2, 2, 2, 344, 346, 3, 2,
	2, 2, 345, 343, 3, 2, 2, 2, 346, 347, 5, 64, 33, 2, 347, 63, 3, 2, 2, 2,
	348, 369, 5, 66, 34, 2, 349, 350, 7, 7, 2, 2, 350, 351, 5, 100, 51, 2,
	351, 352, 7, 8, 2, 2, 352, 353, 5, 64, 33, 2, 353, 369, 3, 2, 2, 2, 354,
	355, 7, 60, 2, 2, 355, 369, 5, 66, 34, 2, 356, 357, 7, 61, 2, 2, 357, 369,
	5, 66, 34, 2, 358, 359, 5, 66, 34, 2, 359, 360, 7, 60, 2, 2, 360, 369,
	3, 2, 2, 2, 361, 362, 5, 66, 34, 2, 362, 363, 7, 61, 2, 2, 363, 369, 3,
	2, 2, 2, 364, 365, 7, 62, 2, 2, 365, 369, 5, 64, 33, 2, 366, 367, 7, 63,
	2, 2, 367, 369, 5, 64, 33, 2, 368, 348, 3, 2, 2, 2, 368, 349, 3, 2, 2,
	2, 368, 354, 3, 2, 2, 2, 368, 356, 3, 2, 2, 2, 368, 358, 3, 2, 2, 2, 368,
	361, 3, 2, 2, 2, 368, 364, 3, 2, 2, 2, 368, 366, 3, 2, 2, 2, 369, 65, 3,
	2, 2, 2, 370, 384, 7, 94, 2, 2, 371, 384, 7, 93, 2, 2, 372, 384, 7, 92,
	2, 2, 373, 384, 5, 70, 36, 2, 374, 384, 5, 72, 37, 2, 375, 384, 5, 74,
	38, 2, 376, 384, 5, 78, 40, 2, 377, 384, 5, 28, 15, 2, 378, 384, 5, 80,
	41, 2, 379, 384, 5, 86, 44, 2, 380, 384, 5, 84, 43, 2, 381, 384, 5, 82,
	42, 2, 382, 384, 5, 90, 46, 2, 383, 370, 3, 2, 2, 2, 383, 371, 3, 2, 2,
	2, 383, 372, 3, 2, 2, 2, 383, 373, 3, 2, 2, 2, 383, 374, 3, 2, 2, 2, 383,
	375, 3, 2, 2, 2, 383, 376, 3, 2, 2, 2, 383, 377, 3, 2, 2, 2, 383, 378,
	3, 2, 2, 2, 383, 379, 3, 2, 2, 2, 383, 380, 3, 2, 2, 2, 383, 381, 3, 2,
	2, 2, 383, 382, 3, 2, 2, 2, 384, 388, 3, 2, 2, 2, 385, 387, 5, 68, 35,
	2, 386, 385, 3, 2, 2, 2, 387, 390, 3, 2, 2, 2, 388, 386, 3, 2, 2, 2, 388,
	389, 3, 2, 2, 2, 389, 67, 3, 2, 2, 2, 390, 388, 3, 2, 2, 2, 391, 392, 7,
	7, 2, 2, 392, 393, 5, 92, 47, 2, 393, 394, 7, 8, 2, 2, 394, 406, 3, 2,
	2, 2, 395, 396, 7, 64, 2, 2, 396, 406, 7, 89, 2, 2, 397, 398, 7, 65, 2,
	2, 398, 401, 5, 56, 29, 2, 399, 400, 7, 29, 2, 2, 400, 402, 5, 56, 29,
	2, 401, 399, 3, 2, 2, 2, 401, 402, 3, 2, 2, 2, 402, 403, 3, 2, 2, 2, 403,
	404, 7, 66, 2, 2, 404, 406, 3, 2, 2, 2, 405, 391, 3, 2, 2, 2, 405, 395,
	3, 2, 2, 2, 405, 397, 3, 2, 2, 2, 406, 69, 3, 2, 2, 2, 407, 413, 7, 67,
	2, 2, 408, 409, 7, 7, 2, 2, 409, 410, 5, 56, 29, 2, 410, 411, 7, 8, 2,
	2, 411, 414, 3, 2, 2, 2, 412, 414, 5, 32, 17, 2, 413, 408, 3, 2, 2, 2,
	413, 412, 3, 2, 2, 2, 414, 71, 3, 2, 2, 2, 415, 421, 7, 68, 2, 2, 416,
	417, 7, 7, 2, 2, 417, 418, 5, 56, 29, 2, 418, 419, 7, 8, 2, 2, 419, 422,
	3, 2, 2, 2, 420, 422, 5, 32, 17, 2, 421, 416, 3, 2, 2, 2, 421, 420, 3,
	2, 2, 2, 422, 73, 3, 2, 2, 2, 423, 424, 7, 69, 2, 2, 424, 425, 7, 7, 2,
	2, 425, 426, 5, 58, 30, 2, 426, 427, 7, 9, 2, 2, 427, 432, 5, 58, 30, 2,
	428, 429, 7, 9, 2, 2, 429, 431, 5, 76, 39, 2, 430, 428, 3, 2, 2, 2, 431,
	434, 3, 2, 2, 2, 432, 430, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 435,
	3, 2, 2, 2, 434, 432, 3, 2, 2, 2, 435, 436, 7, 8, 2, 2, 436, 75, 3, 2,
	2, 2, 437, 438, 7, 70, 2, 2, 438, 443, 5, 66, 34, 2, 439, 440, 5, 100,
	51, 2, 440, 441, 7, 89, 2, 2, 441, 443, 3, 2, 2, 2, 442, 437, 3, 2, 2,
	2, 442, 439, 3, 2, 2, 2, 443, 77, 3, 2, 2, 2, 444, 445, 7, 70, 2, 2, 445,
	446, 7, 7, 2, 2, 446, 447, 5, 98, 50, 2, 447, 448, 7, 8, 2, 2, 448, 449,
	5, 32, 17, 2, 449, 79, 3, 2, 2, 2, 450, 455, 7, 89, 2, 2, 451, 452, 7,
	71, 2, 2, 452, 454, 7, 89, 2, 2, 453, 451, 3, 2, 2, 2, 454, 457, 3, 2,
	2, 2, 455, 453, 3, 2, 2, 2, 455, 456, 3, 2, 2, 2, 456, 81, 3, 2, 2, 2,
	457, 455, 3, 2, 2, 2, 458, 459, 7, 72, 2, 2, 459, 460, 5, 92, 47, 2, 460,
	461, 7, 73, 2, 2, 461, 83, 3, 2, 2, 2, 462, 463, 7, 74, 2, 2, 463, 464,
	5, 92, 47, 2, 464, 465, 7, 75, 2, 2, 465, 85, 3, 2, 2, 2, 466, 480, 7,
	76, 2, 2, 467, 468, 5, 56, 29, 2, 468, 469, 7, 6, 2, 2, 469, 477, 5, 56,
	29, 2, 470, 471, 7, 9, 2, 2, 471, 472, 5, 56, 29, 2, 472, 473, 7, 6, 2,
	2, 473, 474, 5, 56, 29, 2, 474, 476, 3, 2, 2, 2, 475, 470, 3, 2, 2, 2,
	476, 479, 3, 2, 2, 2, 477, 475, 3, 2, 2, 2, 477, 478, 3, 2, 2, 2, 478,
	481, 3, 2, 2, 2, 479, 477, 3, 2, 2, 2, 480, 467, 3, 2, 2, 2, 480, 481,
	3, 2, 2, 2, 481, 483, 3, 2, 2, 2, 482, 484, 7, 9, 2, 2, 483, 482, 3, 2,
	2, 2, 483, 484, 3, 2, 2, 2, 484, 485, 3, 2, 2, 2, 485, 486, 7, 77, 2, 2,
	486, 87, 3, 2, 2, 2, 487, 488, 5, 80, 41, 2, 488, 89, 3, 2, 2, 2, 489,
	490, 7, 7, 2, 2, 490, 491, 5, 56, 29, 2, 491, 492, 7, 8, 2, 2, 492, 91,
	3, 2, 2, 2, 493, 498, 5, 94, 48, 2, 494, 495, 7, 9, 2, 2, 495, 497, 5,
	94, 48, 2, 496, 494, 3, 2, 2, 2, 497, 500, 3, 2, 2, 2, 498, 496, 3, 2,
	2, 2, 498, 499, 3, 2, 2, 2, 499, 502, 3, 2, 2, 2, 500, 498, 3, 2, 2, 2,
	501, 493, 3, 2, 2, 2, 501, 502, 3, 2, 2, 2, 502, 504, 3, 2, 2, 2, 503,
	505, 7, 9, 2, 2, 504, 503, 3, 2, 2, 2, 504, 505, 3, 2, 2, 2, 505, 93, 3,
	2, 2, 2, 506, 508, 7, 78, 2, 2, 507, 506, 3, 2, 2, 2, 507, 508, 3, 2, 2,
	2, 508, 509, 3, 2, 2, 2, 509, 510, 5, 58, 30, 2, 510, 95, 3, 2, 2, 2, 511,
	513, 5, 100, 51, 2, 512, 514, 7, 79, 2, 2, 513, 512, 3, 2, 2, 2, 513, 514,
	3, 2, 2, 2, 514, 516, 3, 2, 2, 2, 515, 517, 7, 89, 2, 2, 516, 515, 3, 2,
	2, 2, 516, 517, 3, 2, 2, 2, 517, 97, 3, 2, 2, 2, 518, 523, 5, 96, 49, 2,
	519, 520, 7, 9, 2, 2, 520, 522, 5, 96, 49, 2, 521, 519, 3, 2, 2, 2, 522,
	525, 3, 2, 2, 2, 523, 521, 3, 2, 2, 2, 523, 524, 3, 2, 2, 2, 524, 527,
	3, 2, 2, 2, 525, 523, 3, 2, 2, 2, 526, 518, 3, 2, 2, 2, 526, 527, 3, 2,
	2, 2, 527, 529, 3, 2, 2, 2, 528, 530, 7, 9, 2, 2, 529, 528, 3, 2, 2, 2,
	529, 530, 3, 2, 2, 2, 530, 99, 3, 2, 2, 2, 531, 570, 7, 80, 2, 2, 532,
	570, 7, 81, 2, 2, 533, 570, 7, 82, 2, 2, 534, 570, 7, 83, 2, 2, 535, 540,
	7, 84, 2, 2, 536, 537, 7, 7, 2, 2, 537, 538, 5, 88, 45, 2, 538, 539, 7,
	8, 2, 2, 539, 541, 3, 2, 2, 2, 540, 536, 3, 2, 2, 2, 540, 541, 3, 2, 2,
	2, 541, 570, 3, 2, 2, 2, 542, 549, 7, 85, 2, 2, 543, 544, 7, 7, 2, 2, 544,
	545, 5, 100, 51, 2, 545, 546, 7, 6, 2, 2, 546, 547, 5, 100, 51, 2, 547,
	548, 7, 8, 2, 2, 548, 550, 3, 2, 2, 2, 549, 543, 3, 2, 2, 2, 549, 550,
	3, 2, 2, 2, 550, 570, 3, 2, 2, 2, 551, 556, 7, 86, 2, 2, 552, 553, 7, 7,
	2, 2, 553, 554, 5, 100, 51, 2, 554, 555, 7, 8, 2, 2, 555, 557, 3, 2, 2,
	2, 556, 552, 3, 2, 2, 2, 556, 557, 3, 2, 2, 2, 557, 570, 3, 2, 2, 2, 558,
	563, 7, 87, 2, 2, 559, 560, 7, 7, 2, 2, 560, 561, 5, 100, 51, 2, 561, 562,
	7, 8, 2, 2, 562, 564, 3, 2, 2, 2, 563, 559, 3, 2, 2, 2, 563, 564, 3, 2,
	2, 2, 564, 570, 3, 2, 2, 2, 565, 567, 7, 88, 2, 2, 566, 568, 5, 102, 52,
	2, 567, 566, 3, 2, 2, 2, 567, 568, 3, 2, 2, 2, 568, 570, 3, 2, 2, 2, 569,
	531, 3, 2, 2, 2, 569, 532, 3, 2, 2, 2, 569, 533, 3, 2, 2, 2, 569, 534,
	3, 2, 2, 2, 569, 535, 3, 2, 2, 2, 569, 542, 3, 2, 2, 2, 569, 551, 3, 2,
	2, 2, 569, 558, 3, 2, 2, 2, 569, 565, 3, 2, 2, 2, 570, 574, 3, 2, 2, 2,
	571, 573, 7, 10, 2, 2, 572, 571, 3, 2, 2, 2, 573, 576, 3, 2, 2, 2, 574,
	572, 3, 2, 2, 2, 574, 575, 3, 2, 2, 2, 575, 101, 3, 2, 2, 2, 576, 574,
	3, 2, 2, 2, 577, 578, 7, 7, 2, 2, 578, 583, 5, 100, 51, 2, 579, 580, 7,
	9, 2, 2, 580, 582, 5, 100, 51, 2, 581, 579, 3, 2, 2, 2, 582, 585, 3, 2,
	2, 2, 583, 581, 3, 2, 2, 2, 583, 584, 3, 2, 2, 2, 584, 587, 3, 2, 2, 2,
	585, 583, 3, 2, 2, 2, 586, 588, 7, 79, 2, 2, 587, 586, 3, 2, 2, 2, 587,
	588, 3, 2, 2, 2, 588, 589, 3, 2, 2, 2, 589, 590, 7, 8, 2, 2, 590, 103,
	3, 2, 2, 2, 60, 107, 117, 120, 127, 133, 138, 148, 158, 169, 175, 181,
	184, 195, 203, 207, 211, 223, 243, 249, 268, 272, 276, 291, 317, 325, 336,
	343, 368, 383, 388, 401, 405, 413, 421, 432, 442, 455, 477, 480, 483, 498,
	501, 504, 507, 513, 516, 523, 526, 529, 540, 549, 556, 563, 567, 569, 574,
	583, 587,
}
var literalNames = []string{
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
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "IDENTIFIER",
	"LETTER", "DIGIT", "FLOAT", "NUMBER", "STRING", "WS",
}

var ruleNames = []string{
	"program", "definition", "impo", "inheritance", "function_declaration",
	"function_definition", "variables", "variable_names", "variable_name",
	"constant", "constant_names", "constant_name", "class_def", "class_implementation",
	"modifiers", "block", "statement", "cond", "while_stmt", "do_while_stmt",
	"for_stmt", "switch_stmt", "case_stmt", "default_stmt", "foreach_stmt",
	"break_stmt", "continue_stmt", "expression", "expression2", "expression3",
	"expression4", "expression5", "expression6", "extension", "catch_", "gauge",
	"sscanf", "lvalue", "lambda", "constant_identifier", "array", "multiset",
	"mapping", "program_specifier", "parenthesis", "expression_list", "splice_expression",
	"argument", "arguments", "type_", "function_type",
}

type pikeParser struct {
	*antlr.BaseParser
}

// NewpikeParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *pikeParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewpikeParser(input antlr.TokenStream) *pikeParser {
	this := new(pikeParser)
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
	this.GrammarFileName = "pike.g4"

	return this
}

// pikeParser tokens.
const (
	pikeParserEOF        = antlr.TokenEOF
	pikeParserT__0       = 1
	pikeParserT__1       = 2
	pikeParserT__2       = 3
	pikeParserT__3       = 4
	pikeParserT__4       = 5
	pikeParserT__5       = 6
	pikeParserT__6       = 7
	pikeParserT__7       = 8
	pikeParserT__8       = 9
	pikeParserT__9       = 10
	pikeParserT__10      = 11
	pikeParserT__11      = 12
	pikeParserT__12      = 13
	pikeParserT__13      = 14
	pikeParserT__14      = 15
	pikeParserT__15      = 16
	pikeParserT__16      = 17
	pikeParserT__17      = 18
	pikeParserT__18      = 19
	pikeParserT__19      = 20
	pikeParserT__20      = 21
	pikeParserT__21      = 22
	pikeParserT__22      = 23
	pikeParserT__23      = 24
	pikeParserT__24      = 25
	pikeParserT__25      = 26
	pikeParserT__26      = 27
	pikeParserT__27      = 28
	pikeParserT__28      = 29
	pikeParserT__29      = 30
	pikeParserT__30      = 31
	pikeParserT__31      = 32
	pikeParserT__32      = 33
	pikeParserT__33      = 34
	pikeParserT__34      = 35
	pikeParserT__35      = 36
	pikeParserT__36      = 37
	pikeParserT__37      = 38
	pikeParserT__38      = 39
	pikeParserT__39      = 40
	pikeParserT__40      = 41
	pikeParserT__41      = 42
	pikeParserT__42      = 43
	pikeParserT__43      = 44
	pikeParserT__44      = 45
	pikeParserT__45      = 46
	pikeParserT__46      = 47
	pikeParserT__47      = 48
	pikeParserT__48      = 49
	pikeParserT__49      = 50
	pikeParserT__50      = 51
	pikeParserT__51      = 52
	pikeParserT__52      = 53
	pikeParserT__53      = 54
	pikeParserT__54      = 55
	pikeParserT__55      = 56
	pikeParserT__56      = 57
	pikeParserT__57      = 58
	pikeParserT__58      = 59
	pikeParserT__59      = 60
	pikeParserT__60      = 61
	pikeParserT__61      = 62
	pikeParserT__62      = 63
	pikeParserT__63      = 64
	pikeParserT__64      = 65
	pikeParserT__65      = 66
	pikeParserT__66      = 67
	pikeParserT__67      = 68
	pikeParserT__68      = 69
	pikeParserT__69      = 70
	pikeParserT__70      = 71
	pikeParserT__71      = 72
	pikeParserT__72      = 73
	pikeParserT__73      = 74
	pikeParserT__74      = 75
	pikeParserT__75      = 76
	pikeParserT__76      = 77
	pikeParserT__77      = 78
	pikeParserT__78      = 79
	pikeParserT__79      = 80
	pikeParserT__80      = 81
	pikeParserT__81      = 82
	pikeParserT__82      = 83
	pikeParserT__83      = 84
	pikeParserT__84      = 85
	pikeParserT__85      = 86
	pikeParserIDENTIFIER = 87
	pikeParserLETTER     = 88
	pikeParserDIGIT      = 89
	pikeParserFLOAT      = 90
	pikeParserNUMBER     = 91
	pikeParserSTRING     = 92
	pikeParserWS         = 93
)

// pikeParser rules.
const (
	pikeParserRULE_program              = 0
	pikeParserRULE_definition           = 1
	pikeParserRULE_impo                 = 2
	pikeParserRULE_inheritance          = 3
	pikeParserRULE_function_declaration = 4
	pikeParserRULE_function_definition  = 5
	pikeParserRULE_variables            = 6
	pikeParserRULE_variable_names       = 7
	pikeParserRULE_variable_name        = 8
	pikeParserRULE_constant             = 9
	pikeParserRULE_constant_names       = 10
	pikeParserRULE_constant_name        = 11
	pikeParserRULE_class_def            = 12
	pikeParserRULE_class_implementation = 13
	pikeParserRULE_modifiers            = 14
	pikeParserRULE_block                = 15
	pikeParserRULE_statement            = 16
	pikeParserRULE_cond                 = 17
	pikeParserRULE_while_stmt           = 18
	pikeParserRULE_do_while_stmt        = 19
	pikeParserRULE_for_stmt             = 20
	pikeParserRULE_switch_stmt          = 21
	pikeParserRULE_case_stmt            = 22
	pikeParserRULE_default_stmt         = 23
	pikeParserRULE_foreach_stmt         = 24
	pikeParserRULE_break_stmt           = 25
	pikeParserRULE_continue_stmt        = 26
	pikeParserRULE_expression           = 27
	pikeParserRULE_expression2          = 28
	pikeParserRULE_expression3          = 29
	pikeParserRULE_expression4          = 30
	pikeParserRULE_expression5          = 31
	pikeParserRULE_expression6          = 32
	pikeParserRULE_extension            = 33
	pikeParserRULE_catch_               = 34
	pikeParserRULE_gauge                = 35
	pikeParserRULE_sscanf               = 36
	pikeParserRULE_lvalue               = 37
	pikeParserRULE_lambda               = 38
	pikeParserRULE_constant_identifier  = 39
	pikeParserRULE_array                = 40
	pikeParserRULE_multiset             = 41
	pikeParserRULE_mapping              = 42
	pikeParserRULE_program_specifier    = 43
	pikeParserRULE_parenthesis          = 44
	pikeParserRULE_expression_list      = 45
	pikeParserRULE_splice_expression    = 46
	pikeParserRULE_argument             = 47
	pikeParserRULE_arguments            = 48
	pikeParserRULE_type_                = 49
	pikeParserRULE_function_type        = 50
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
	p.RuleIndex = pikeParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *ProgramContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *pikeParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, pikeParserRULE_program)
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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pikeParserT__0)|(1<<pikeParserT__2)|(1<<pikeParserT__9)|(1<<pikeParserT__10)|(1<<pikeParserT__13)|(1<<pikeParserT__14)|(1<<pikeParserT__15)|(1<<pikeParserT__16)|(1<<pikeParserT__17)|(1<<pikeParserT__18))) != 0) || (((_la-78)&-(0x1f+1)) == 0 && ((1<<uint((_la-78)))&((1<<(pikeParserT__77-78))|(1<<(pikeParserT__78-78))|(1<<(pikeParserT__79-78))|(1<<(pikeParserT__80-78))|(1<<(pikeParserT__81-78))|(1<<(pikeParserT__82-78))|(1<<(pikeParserT__83-78))|(1<<(pikeParserT__84-78))|(1<<(pikeParserT__85-78)))) != 0) {
		{
			p.SetState(102)
			p.Definition()
		}

		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) Impo() IImpoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImpoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImpoContext)
}

func (s *DefinitionContext) Inheritance() IInheritanceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInheritanceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInheritanceContext)
}

func (s *DefinitionContext) Function_declaration() IFunction_declarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_declarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_declarationContext)
}

func (s *DefinitionContext) Function_definition() IFunction_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_definitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_definitionContext)
}

func (s *DefinitionContext) Variables() IVariablesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariablesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariablesContext)
}

func (s *DefinitionContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *DefinitionContext) Class_def() IClass_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClass_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClass_defContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *pikeParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, pikeParserRULE_definition)

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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Impo()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Inheritance()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Function_declaration()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.Function_definition()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(112)
			p.Variables()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(113)
			p.Constant()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(114)
			p.Class_def()
		}

	}

	return localctx
}

// IImpoContext is an interface to support dynamic dispatch.
type IImpoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImpoContext differentiates from other interfaces.
	IsImpoContext()
}

type ImpoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImpoContext() *ImpoContext {
	var p = new(ImpoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_impo
	return p
}

func (*ImpoContext) IsImpoContext() {}

func NewImpoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImpoContext {
	var p = new(ImpoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_impo

	return p
}

func (s *ImpoContext) GetParser() antlr.Parser { return s.parser }

func (s *ImpoContext) Constant_identifier() IConstant_identifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstant_identifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstant_identifierContext)
}

func (s *ImpoContext) Modifiers() IModifiersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifiersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifiersContext)
}

func (s *ImpoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImpoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImpoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterImpo(s)
	}
}

func (s *ImpoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitImpo(s)
	}
}

func (p *pikeParser) Impo() (localctx IImpoContext) {
	this := p
	_ = this

	localctx = NewImpoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, pikeParserRULE_impo)
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
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pikeParserT__13)|(1<<pikeParserT__14)|(1<<pikeParserT__15)|(1<<pikeParserT__16)|(1<<pikeParserT__17)|(1<<pikeParserT__18))) != 0 {
		{
			p.SetState(117)
			p.Modifiers()
		}

	}
	{
		p.SetState(120)
		p.Match(pikeParserT__0)
	}
	{
		p.SetState(121)
		p.Constant_identifier()
	}
	{
		p.SetState(122)
		p.Match(pikeParserT__1)
	}

	return localctx
}

// IInheritanceContext is an interface to support dynamic dispatch.
type IInheritanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInheritanceContext differentiates from other interfaces.
	IsInheritanceContext()
}

type InheritanceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInheritanceContext() *InheritanceContext {
	var p = new(InheritanceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_inheritance
	return p
}

func (*InheritanceContext) IsInheritanceContext() {}

func NewInheritanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InheritanceContext {
	var p = new(InheritanceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_inheritance

	return p
}

func (s *InheritanceContext) GetParser() antlr.Parser { return s.parser }

func (s *InheritanceContext) Program_specifier() IProgram_specifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgram_specifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgram_specifierContext)
}

func (s *InheritanceContext) Modifiers() IModifiersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifiersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifiersContext)
}

func (s *InheritanceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pikeParserIDENTIFIER, 0)
}

func (s *InheritanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InheritanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InheritanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterInheritance(s)
	}
}

func (s *InheritanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitInheritance(s)
	}
}

func (p *pikeParser) Inheritance() (localctx IInheritanceContext) {
	this := p
	_ = this

	localctx = NewInheritanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, pikeParserRULE_inheritance)
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
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pikeParserT__13)|(1<<pikeParserT__14)|(1<<pikeParserT__15)|(1<<pikeParserT__16)|(1<<pikeParserT__17)|(1<<pikeParserT__18))) != 0 {
		{
			p.SetState(124)
			p.Modifiers()
		}

	}
	{
		p.SetState(127)
		p.Match(pikeParserT__2)
	}
	{
		p.SetState(128)
		p.Program_specifier()
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__3 {
		{
			p.SetState(129)
			p.Match(pikeParserT__3)
		}
		{
			p.SetState(130)
			p.Match(pikeParserIDENTIFIER)
		}

	}
	{
		p.SetState(133)
		p.Match(pikeParserT__1)
	}

	return localctx
}

// IFunction_declarationContext is an interface to support dynamic dispatch.
type IFunction_declarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_declarationContext differentiates from other interfaces.
	IsFunction_declarationContext()
}

type Function_declarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_declarationContext() *Function_declarationContext {
	var p = new(Function_declarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_function_declaration
	return p
}

func (*Function_declarationContext) IsFunction_declarationContext() {}

func NewFunction_declarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_declarationContext {
	var p = new(Function_declarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_function_declaration

	return p
}

func (s *Function_declarationContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_declarationContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Function_declarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pikeParserIDENTIFIER, 0)
}

func (s *Function_declarationContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *Function_declarationContext) Modifiers() IModifiersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifiersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifiersContext)
}

func (s *Function_declarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_declarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_declarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterFunction_declaration(s)
	}
}

func (s *Function_declarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitFunction_declaration(s)
	}
}

func (p *pikeParser) Function_declaration() (localctx IFunction_declarationContext) {
	this := p
	_ = this

	localctx = NewFunction_declarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, pikeParserRULE_function_declaration)
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
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pikeParserT__13)|(1<<pikeParserT__14)|(1<<pikeParserT__15)|(1<<pikeParserT__16)|(1<<pikeParserT__17)|(1<<pikeParserT__18))) != 0 {
		{
			p.SetState(135)
			p.Modifiers()
		}

	}
	{
		p.SetState(138)
		p.Type_()
	}
	{
		p.SetState(139)
		p.Match(pikeParserIDENTIFIER)
	}
	{
		p.SetState(140)
		p.Match(pikeParserT__4)
	}
	{
		p.SetState(141)
		p.Arguments()
	}
	{
		p.SetState(142)
		p.Match(pikeParserT__5)
	}
	{
		p.SetState(143)
		p.Match(pikeParserT__1)
	}

	return localctx
}

// IFunction_definitionContext is an interface to support dynamic dispatch.
type IFunction_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_definitionContext differentiates from other interfaces.
	IsFunction_definitionContext()
}

type Function_definitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_definitionContext() *Function_definitionContext {
	var p = new(Function_definitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_function_definition
	return p
}

func (*Function_definitionContext) IsFunction_definitionContext() {}

func NewFunction_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_definitionContext {
	var p = new(Function_definitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_function_definition

	return p
}

func (s *Function_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_definitionContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Function_definitionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pikeParserIDENTIFIER, 0)
}

func (s *Function_definitionContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *Function_definitionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Function_definitionContext) Modifiers() IModifiersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifiersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifiersContext)
}

func (s *Function_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterFunction_definition(s)
	}
}

func (s *Function_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitFunction_definition(s)
	}
}

func (p *pikeParser) Function_definition() (localctx IFunction_definitionContext) {
	this := p
	_ = this

	localctx = NewFunction_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, pikeParserRULE_function_definition)
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
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pikeParserT__13)|(1<<pikeParserT__14)|(1<<pikeParserT__15)|(1<<pikeParserT__16)|(1<<pikeParserT__17)|(1<<pikeParserT__18))) != 0 {
		{
			p.SetState(145)
			p.Modifiers()
		}

	}
	{
		p.SetState(148)
		p.Type_()
	}
	{
		p.SetState(149)
		p.Match(pikeParserIDENTIFIER)
	}
	{
		p.SetState(150)
		p.Match(pikeParserT__4)
	}
	{
		p.SetState(151)
		p.Arguments()
	}
	{
		p.SetState(152)
		p.Match(pikeParserT__5)
	}
	{
		p.SetState(153)
		p.Block()
	}

	return localctx
}

// IVariablesContext is an interface to support dynamic dispatch.
type IVariablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariablesContext differentiates from other interfaces.
	IsVariablesContext()
}

type VariablesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariablesContext() *VariablesContext {
	var p = new(VariablesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_variables
	return p
}

func (*VariablesContext) IsVariablesContext() {}

func NewVariablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablesContext {
	var p = new(VariablesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_variables

	return p
}

func (s *VariablesContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablesContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VariablesContext) Variable_names() IVariable_namesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariable_namesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariable_namesContext)
}

func (s *VariablesContext) Modifiers() IModifiersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifiersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifiersContext)
}

func (s *VariablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterVariables(s)
	}
}

func (s *VariablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitVariables(s)
	}
}

func (p *pikeParser) Variables() (localctx IVariablesContext) {
	this := p
	_ = this

	localctx = NewVariablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, pikeParserRULE_variables)
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
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pikeParserT__13)|(1<<pikeParserT__14)|(1<<pikeParserT__15)|(1<<pikeParserT__16)|(1<<pikeParserT__17)|(1<<pikeParserT__18))) != 0 {
		{
			p.SetState(155)
			p.Modifiers()
		}

	}
	{
		p.SetState(158)
		p.Type_()
	}
	{
		p.SetState(159)
		p.Variable_names()
	}
	{
		p.SetState(160)
		p.Match(pikeParserT__1)
	}

	return localctx
}

// IVariable_namesContext is an interface to support dynamic dispatch.
type IVariable_namesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariable_namesContext differentiates from other interfaces.
	IsVariable_namesContext()
}

type Variable_namesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariable_namesContext() *Variable_namesContext {
	var p = new(Variable_namesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_variable_names
	return p
}

func (*Variable_namesContext) IsVariable_namesContext() {}

func NewVariable_namesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_namesContext {
	var p = new(Variable_namesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_variable_names

	return p
}

func (s *Variable_namesContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_namesContext) AllVariable_name() []IVariable_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariable_nameContext)(nil)).Elem())
	var tst = make([]IVariable_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariable_nameContext)
		}
	}

	return tst
}

func (s *Variable_namesContext) Variable_name(i int) IVariable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariable_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariable_nameContext)
}

func (s *Variable_namesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_namesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Variable_namesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterVariable_names(s)
	}
}

func (s *Variable_namesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitVariable_names(s)
	}
}

func (p *pikeParser) Variable_names() (localctx IVariable_namesContext) {
	this := p
	_ = this

	localctx = NewVariable_namesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, pikeParserRULE_variable_names)
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
		p.SetState(162)
		p.Variable_name()
	}
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pikeParserT__6 {
		{
			p.SetState(163)
			p.Match(pikeParserT__6)
		}
		{
			p.SetState(164)
			p.Variable_name()
		}

		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVariable_nameContext is an interface to support dynamic dispatch.
type IVariable_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariable_nameContext differentiates from other interfaces.
	IsVariable_nameContext()
}

type Variable_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariable_nameContext() *Variable_nameContext {
	var p = new(Variable_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_variable_name
	return p
}

func (*Variable_nameContext) IsVariable_nameContext() {}

func NewVariable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_nameContext {
	var p = new(Variable_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_variable_name

	return p
}

func (s *Variable_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pikeParserIDENTIFIER, 0)
}

func (s *Variable_nameContext) Expression2() IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *Variable_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Variable_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterVariable_name(s)
	}
}

func (s *Variable_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitVariable_name(s)
	}
}

func (p *pikeParser) Variable_name() (localctx IVariable_nameContext) {
	this := p
	_ = this

	localctx = NewVariable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, pikeParserRULE_variable_name)
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
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pikeParserT__7 {
		{
			p.SetState(170)
			p.Match(pikeParserT__7)
		}

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(176)
		p.Match(pikeParserIDENTIFIER)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__8 {
		{
			p.SetState(177)
			p.Match(pikeParserT__8)
		}
		{
			p.SetState(178)
			p.Expression2()
		}

	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) Constant_names() IConstant_namesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstant_namesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstant_namesContext)
}

func (s *ConstantContext) Modifiers() IModifiersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifiersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifiersContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (p *pikeParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, pikeParserRULE_constant)
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
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pikeParserT__13)|(1<<pikeParserT__14)|(1<<pikeParserT__15)|(1<<pikeParserT__16)|(1<<pikeParserT__17)|(1<<pikeParserT__18))) != 0 {
		{
			p.SetState(181)
			p.Modifiers()
		}

	}
	{
		p.SetState(184)
		p.Match(pikeParserT__9)
	}
	{
		p.SetState(185)
		p.Constant_names()
	}
	{
		p.SetState(186)
		p.Match(pikeParserT__1)
	}

	return localctx
}

// IConstant_namesContext is an interface to support dynamic dispatch.
type IConstant_namesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstant_namesContext differentiates from other interfaces.
	IsConstant_namesContext()
}

type Constant_namesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstant_namesContext() *Constant_namesContext {
	var p = new(Constant_namesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_constant_names
	return p
}

func (*Constant_namesContext) IsConstant_namesContext() {}

func NewConstant_namesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Constant_namesContext {
	var p = new(Constant_namesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_constant_names

	return p
}

func (s *Constant_namesContext) GetParser() antlr.Parser { return s.parser }

func (s *Constant_namesContext) AllConstant_name() []IConstant_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstant_nameContext)(nil)).Elem())
	var tst = make([]IConstant_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstant_nameContext)
		}
	}

	return tst
}

func (s *Constant_namesContext) Constant_name(i int) IConstant_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstant_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstant_nameContext)
}

func (s *Constant_namesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constant_namesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Constant_namesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterConstant_names(s)
	}
}

func (s *Constant_namesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitConstant_names(s)
	}
}

func (p *pikeParser) Constant_names() (localctx IConstant_namesContext) {
	this := p
	_ = this

	localctx = NewConstant_namesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, pikeParserRULE_constant_names)
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
		p.SetState(188)
		p.Constant_name()
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pikeParserT__6 {
		{
			p.SetState(189)
			p.Match(pikeParserT__6)
		}
		{
			p.SetState(190)
			p.Constant_name()
		}

		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConstant_nameContext is an interface to support dynamic dispatch.
type IConstant_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstant_nameContext differentiates from other interfaces.
	IsConstant_nameContext()
}

type Constant_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstant_nameContext() *Constant_nameContext {
	var p = new(Constant_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_constant_name
	return p
}

func (*Constant_nameContext) IsConstant_nameContext() {}

func NewConstant_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Constant_nameContext {
	var p = new(Constant_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_constant_name

	return p
}

func (s *Constant_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Constant_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pikeParserIDENTIFIER, 0)
}

func (s *Constant_nameContext) Expression2() IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *Constant_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constant_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Constant_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterConstant_name(s)
	}
}

func (s *Constant_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitConstant_name(s)
	}
}

func (p *pikeParser) Constant_name() (localctx IConstant_nameContext) {
	this := p
	_ = this

	localctx = NewConstant_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, pikeParserRULE_constant_name)

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
		p.SetState(196)
		p.Match(pikeParserIDENTIFIER)
	}
	{
		p.SetState(197)
		p.Match(pikeParserT__8)
	}
	{
		p.SetState(198)
		p.Expression2()
	}

	return localctx
}

// IClass_defContext is an interface to support dynamic dispatch.
type IClass_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClass_defContext differentiates from other interfaces.
	IsClass_defContext()
}

type Class_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClass_defContext() *Class_defContext {
	var p = new(Class_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_class_def
	return p
}

func (*Class_defContext) IsClass_defContext() {}

func NewClass_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Class_defContext {
	var p = new(Class_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_class_def

	return p
}

func (s *Class_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Class_defContext) Modifiers() IModifiersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModifiersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModifiersContext)
}

func (s *Class_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Class_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Class_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterClass_def(s)
	}
}

func (s *Class_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitClass_def(s)
	}
}

func (p *pikeParser) Class_def() (localctx IClass_defContext) {
	this := p
	_ = this

	localctx = NewClass_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, pikeParserRULE_class_def)
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
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pikeParserT__13)|(1<<pikeParserT__14)|(1<<pikeParserT__15)|(1<<pikeParserT__16)|(1<<pikeParserT__17)|(1<<pikeParserT__18))) != 0 {
		{
			p.SetState(200)
			p.Modifiers()
		}

	}
	{
		p.SetState(203)
		p.Match(pikeParserT__10)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__1 {
		{
			p.SetState(204)
			p.Match(pikeParserT__1)
		}

	}

	return localctx
}

// IClass_implementationContext is an interface to support dynamic dispatch.
type IClass_implementationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClass_implementationContext differentiates from other interfaces.
	IsClass_implementationContext()
}

type Class_implementationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClass_implementationContext() *Class_implementationContext {
	var p = new(Class_implementationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_class_implementation
	return p
}

func (*Class_implementationContext) IsClass_implementationContext() {}

func NewClass_implementationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Class_implementationContext {
	var p = new(Class_implementationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_class_implementation

	return p
}

func (s *Class_implementationContext) GetParser() antlr.Parser { return s.parser }

func (s *Class_implementationContext) Program() IProgramContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgramContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgramContext)
}

func (s *Class_implementationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pikeParserIDENTIFIER, 0)
}

func (s *Class_implementationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Class_implementationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Class_implementationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterClass_implementation(s)
	}
}

func (s *Class_implementationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitClass_implementation(s)
	}
}

func (p *pikeParser) Class_implementation() (localctx IClass_implementationContext) {
	this := p
	_ = this

	localctx = NewClass_implementationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, pikeParserRULE_class_implementation)
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
		p.SetState(207)
		p.Match(pikeParserT__10)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserIDENTIFIER {
		{
			p.SetState(208)
			p.Match(pikeParserIDENTIFIER)
		}

	}
	{
		p.SetState(211)
		p.Match(pikeParserT__11)
	}
	{
		p.SetState(212)
		p.Program()
	}
	{
		p.SetState(213)
		p.Match(pikeParserT__12)
	}

	return localctx
}

// IModifiersContext is an interface to support dynamic dispatch.
type IModifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModifiersContext differentiates from other interfaces.
	IsModifiersContext()
}

type ModifiersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModifiersContext() *ModifiersContext {
	var p = new(ModifiersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_modifiers
	return p
}

func (*ModifiersContext) IsModifiersContext() {}

func NewModifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModifiersContext {
	var p = new(ModifiersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_modifiers

	return p
}

func (s *ModifiersContext) GetParser() antlr.Parser { return s.parser }
func (s *ModifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModifiersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterModifiers(s)
	}
}

func (s *ModifiersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitModifiers(s)
	}
}

func (p *pikeParser) Modifiers() (localctx IModifiersContext) {
	this := p
	_ = this

	localctx = NewModifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, pikeParserRULE_modifiers)
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
		p.SetState(215)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pikeParserT__13)|(1<<pikeParserT__14)|(1<<pikeParserT__15)|(1<<pikeParserT__16)|(1<<pikeParserT__17)|(1<<pikeParserT__18))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = pikeParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

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

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *pikeParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, pikeParserRULE_block)
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
		p.SetState(217)
		p.Match(pikeParserT__11)
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<pikeParserT__1)|(1<<pikeParserT__4)|(1<<pikeParserT__10)|(1<<pikeParserT__11)|(1<<pikeParserT__19)|(1<<pikeParserT__21)|(1<<pikeParserT__22)|(1<<pikeParserT__23)|(1<<pikeParserT__24)|(1<<pikeParserT__25)|(1<<pikeParserT__27)|(1<<pikeParserT__28)|(1<<pikeParserT__29)|(1<<pikeParserT__30))) != 0) || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(pikeParserT__57-58))|(1<<(pikeParserT__58-58))|(1<<(pikeParserT__59-58))|(1<<(pikeParserT__60-58))|(1<<(pikeParserT__64-58))|(1<<(pikeParserT__65-58))|(1<<(pikeParserT__66-58))|(1<<(pikeParserT__67-58))|(1<<(pikeParserT__69-58))|(1<<(pikeParserT__71-58))|(1<<(pikeParserT__73-58))|(1<<(pikeParserT__77-58))|(1<<(pikeParserT__78-58))|(1<<(pikeParserT__79-58))|(1<<(pikeParserT__80-58))|(1<<(pikeParserT__81-58))|(1<<(pikeParserT__82-58))|(1<<(pikeParserT__83-58))|(1<<(pikeParserT__84-58))|(1<<(pikeParserT__85-58))|(1<<(pikeParserIDENTIFIER-58)))) != 0) || (((_la-90)&-(0x1f+1)) == 0 && ((1<<uint((_la-90)))&((1<<(pikeParserFLOAT-90))|(1<<(pikeParserNUMBER-90))|(1<<(pikeParserSTRING-90)))) != 0) {
		{
			p.SetState(218)
			p.Statement()
		}

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(224)
		p.Match(pikeParserT__12)
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
	p.RuleIndex = pikeParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Expression2() IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *StatementContext) Cond() ICondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondContext)
}

func (s *StatementContext) While_stmt() IWhile_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *StatementContext) Do_while_stmt() IDo_while_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDo_while_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDo_while_stmtContext)
}

func (s *StatementContext) For_stmt() IFor_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_stmtContext)
}

func (s *StatementContext) Switch_stmt() ISwitch_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitch_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitch_stmtContext)
}

func (s *StatementContext) Case_stmt() ICase_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICase_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICase_stmtContext)
}

func (s *StatementContext) Default_stmt() IDefault_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefault_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefault_stmtContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) Foreach_stmt() IForeach_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForeach_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForeach_stmtContext)
}

func (s *StatementContext) Break_stmt() IBreak_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreak_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreak_stmtContext)
}

func (s *StatementContext) Continue_stmt() IContinue_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinue_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinue_stmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *pikeParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, pikeParserRULE_statement)

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

	p.SetState(241)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pikeParserT__4, pikeParserT__10, pikeParserT__57, pikeParserT__58, pikeParserT__59, pikeParserT__60, pikeParserT__64, pikeParserT__65, pikeParserT__66, pikeParserT__67, pikeParserT__69, pikeParserT__71, pikeParserT__73, pikeParserT__77, pikeParserT__78, pikeParserT__79, pikeParserT__80, pikeParserT__81, pikeParserT__82, pikeParserT__83, pikeParserT__84, pikeParserT__85, pikeParserIDENTIFIER, pikeParserFLOAT, pikeParserNUMBER, pikeParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(226)
			p.Expression2()
		}
		{
			p.SetState(227)
			p.Match(pikeParserT__1)
		}

	case pikeParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.Cond()
		}

	case pikeParserT__21:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(230)
			p.While_stmt()
		}

	case pikeParserT__22:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(231)
			p.Do_while_stmt()
		}

	case pikeParserT__23:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(232)
			p.For_stmt()
		}

	case pikeParserT__24:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(233)
			p.Switch_stmt()
		}

	case pikeParserT__25:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(234)
			p.Case_stmt()
		}

	case pikeParserT__27:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(235)
			p.Default_stmt()
		}

	case pikeParserT__11:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(236)
			p.Block()
		}

	case pikeParserT__28:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(237)
			p.Foreach_stmt()
		}

	case pikeParserT__29:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(238)
			p.Break_stmt()
		}

	case pikeParserT__30:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(239)
			p.Continue_stmt()
		}

	case pikeParserT__1:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(240)
			p.Match(pikeParserT__1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICondContext is an interface to support dynamic dispatch.
type ICondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCondContext differentiates from other interfaces.
	IsCondContext()
}

type CondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondContext() *CondContext {
	var p = new(CondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_cond
	return p
}

func (*CondContext) IsCondContext() {}

func NewCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondContext {
	var p = new(CondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_cond

	return p
}

func (s *CondContext) GetParser() antlr.Parser { return s.parser }

func (s *CondContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *CondContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterCond(s)
	}
}

func (s *CondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitCond(s)
	}
}

func (p *pikeParser) Cond() (localctx ICondContext) {
	this := p
	_ = this

	localctx = NewCondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, pikeParserRULE_cond)

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
		p.SetState(243)
		p.Match(pikeParserT__19)
	}
	{
		p.SetState(244)
		p.Statement()
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(245)
			p.Match(pikeParserT__20)
		}
		{
			p.SetState(246)
			p.Statement()
		}

	}

	return localctx
}

// IWhile_stmtContext is an interface to support dynamic dispatch.
type IWhile_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhile_stmtContext differentiates from other interfaces.
	IsWhile_stmtContext()
}

type While_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_stmtContext() *While_stmtContext {
	var p = new(While_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_while_stmt
	return p
}

func (*While_stmtContext) IsWhile_stmtContext() {}

func NewWhile_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_stmtContext {
	var p = new(While_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_while_stmt

	return p
}

func (s *While_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *While_stmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *While_stmtContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *While_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *While_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterWhile_stmt(s)
	}
}

func (s *While_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitWhile_stmt(s)
	}
}

func (p *pikeParser) While_stmt() (localctx IWhile_stmtContext) {
	this := p
	_ = this

	localctx = NewWhile_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, pikeParserRULE_while_stmt)

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
		p.SetState(249)
		p.Match(pikeParserT__21)
	}
	{
		p.SetState(250)
		p.Match(pikeParserT__4)
	}
	{
		p.SetState(251)
		p.Expression()
	}
	{
		p.SetState(252)
		p.Match(pikeParserT__5)
	}
	{
		p.SetState(253)
		p.Statement()
	}

	return localctx
}

// IDo_while_stmtContext is an interface to support dynamic dispatch.
type IDo_while_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDo_while_stmtContext differentiates from other interfaces.
	IsDo_while_stmtContext()
}

type Do_while_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDo_while_stmtContext() *Do_while_stmtContext {
	var p = new(Do_while_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_do_while_stmt
	return p
}

func (*Do_while_stmtContext) IsDo_while_stmtContext() {}

func NewDo_while_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Do_while_stmtContext {
	var p = new(Do_while_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_do_while_stmt

	return p
}

func (s *Do_while_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Do_while_stmtContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Do_while_stmtContext) While_stmt() IWhile_stmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhile_stmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhile_stmtContext)
}

func (s *Do_while_stmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Do_while_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Do_while_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Do_while_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterDo_while_stmt(s)
	}
}

func (s *Do_while_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitDo_while_stmt(s)
	}
}

func (p *pikeParser) Do_while_stmt() (localctx IDo_while_stmtContext) {
	this := p
	_ = this

	localctx = NewDo_while_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, pikeParserRULE_do_while_stmt)

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
		p.SetState(255)
		p.Match(pikeParserT__22)
	}
	{
		p.SetState(256)
		p.Statement()
	}
	{
		p.SetState(257)
		p.While_stmt()
	}
	{
		p.SetState(258)
		p.Match(pikeParserT__4)
	}
	{
		p.SetState(259)
		p.Expression()
	}
	{
		p.SetState(260)
		p.Match(pikeParserT__5)
	}
	{
		p.SetState(261)
		p.Match(pikeParserT__1)
	}

	return localctx
}

// IFor_stmtContext is an interface to support dynamic dispatch.
type IFor_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_stmtContext differentiates from other interfaces.
	IsFor_stmtContext()
}

type For_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_stmtContext() *For_stmtContext {
	var p = new(For_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_for_stmt
	return p
}

func (*For_stmtContext) IsFor_stmtContext() {}

func NewFor_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_stmtContext {
	var p = new(For_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_for_stmt

	return p
}

func (s *For_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *For_stmtContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *For_stmtContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *For_stmtContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *For_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterFor_stmt(s)
	}
}

func (s *For_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitFor_stmt(s)
	}
}

func (p *pikeParser) For_stmt() (localctx IFor_stmtContext) {
	this := p
	_ = this

	localctx = NewFor_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, pikeParserRULE_for_stmt)
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
		p.SetState(263)
		p.Match(pikeParserT__23)
	}
	{
		p.SetState(264)
		p.Match(pikeParserT__4)
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__4 || _la == pikeParserT__10 || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(pikeParserT__57-58))|(1<<(pikeParserT__58-58))|(1<<(pikeParserT__59-58))|(1<<(pikeParserT__60-58))|(1<<(pikeParserT__64-58))|(1<<(pikeParserT__65-58))|(1<<(pikeParserT__66-58))|(1<<(pikeParserT__67-58))|(1<<(pikeParserT__69-58))|(1<<(pikeParserT__71-58))|(1<<(pikeParserT__73-58))|(1<<(pikeParserT__77-58))|(1<<(pikeParserT__78-58))|(1<<(pikeParserT__79-58))|(1<<(pikeParserT__80-58))|(1<<(pikeParserT__81-58))|(1<<(pikeParserT__82-58))|(1<<(pikeParserT__83-58))|(1<<(pikeParserT__84-58))|(1<<(pikeParserT__85-58))|(1<<(pikeParserIDENTIFIER-58)))) != 0) || (((_la-90)&-(0x1f+1)) == 0 && ((1<<uint((_la-90)))&((1<<(pikeParserFLOAT-90))|(1<<(pikeParserNUMBER-90))|(1<<(pikeParserSTRING-90)))) != 0) {
		{
			p.SetState(265)
			p.Expression()
		}

	}
	{
		p.SetState(268)
		p.Match(pikeParserT__1)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__4 || _la == pikeParserT__10 || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(pikeParserT__57-58))|(1<<(pikeParserT__58-58))|(1<<(pikeParserT__59-58))|(1<<(pikeParserT__60-58))|(1<<(pikeParserT__64-58))|(1<<(pikeParserT__65-58))|(1<<(pikeParserT__66-58))|(1<<(pikeParserT__67-58))|(1<<(pikeParserT__69-58))|(1<<(pikeParserT__71-58))|(1<<(pikeParserT__73-58))|(1<<(pikeParserT__77-58))|(1<<(pikeParserT__78-58))|(1<<(pikeParserT__79-58))|(1<<(pikeParserT__80-58))|(1<<(pikeParserT__81-58))|(1<<(pikeParserT__82-58))|(1<<(pikeParserT__83-58))|(1<<(pikeParserT__84-58))|(1<<(pikeParserT__85-58))|(1<<(pikeParserIDENTIFIER-58)))) != 0) || (((_la-90)&-(0x1f+1)) == 0 && ((1<<uint((_la-90)))&((1<<(pikeParserFLOAT-90))|(1<<(pikeParserNUMBER-90))|(1<<(pikeParserSTRING-90)))) != 0) {
		{
			p.SetState(269)
			p.Expression()
		}

	}
	{
		p.SetState(272)
		p.Match(pikeParserT__1)
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__4 || _la == pikeParserT__10 || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(pikeParserT__57-58))|(1<<(pikeParserT__58-58))|(1<<(pikeParserT__59-58))|(1<<(pikeParserT__60-58))|(1<<(pikeParserT__64-58))|(1<<(pikeParserT__65-58))|(1<<(pikeParserT__66-58))|(1<<(pikeParserT__67-58))|(1<<(pikeParserT__69-58))|(1<<(pikeParserT__71-58))|(1<<(pikeParserT__73-58))|(1<<(pikeParserT__77-58))|(1<<(pikeParserT__78-58))|(1<<(pikeParserT__79-58))|(1<<(pikeParserT__80-58))|(1<<(pikeParserT__81-58))|(1<<(pikeParserT__82-58))|(1<<(pikeParserT__83-58))|(1<<(pikeParserT__84-58))|(1<<(pikeParserT__85-58))|(1<<(pikeParserIDENTIFIER-58)))) != 0) || (((_la-90)&-(0x1f+1)) == 0 && ((1<<uint((_la-90)))&((1<<(pikeParserFLOAT-90))|(1<<(pikeParserNUMBER-90))|(1<<(pikeParserSTRING-90)))) != 0) {
		{
			p.SetState(273)
			p.Expression()
		}

	}
	{
		p.SetState(276)
		p.Match(pikeParserT__5)
	}
	{
		p.SetState(277)
		p.Statement()
	}

	return localctx
}

// ISwitch_stmtContext is an interface to support dynamic dispatch.
type ISwitch_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitch_stmtContext differentiates from other interfaces.
	IsSwitch_stmtContext()
}

type Switch_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_stmtContext() *Switch_stmtContext {
	var p = new(Switch_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_switch_stmt
	return p
}

func (*Switch_stmtContext) IsSwitch_stmtContext() {}

func NewSwitch_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_stmtContext {
	var p = new(Switch_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_switch_stmt

	return p
}

func (s *Switch_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_stmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Switch_stmtContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Switch_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Switch_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterSwitch_stmt(s)
	}
}

func (s *Switch_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitSwitch_stmt(s)
	}
}

func (p *pikeParser) Switch_stmt() (localctx ISwitch_stmtContext) {
	this := p
	_ = this

	localctx = NewSwitch_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, pikeParserRULE_switch_stmt)

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
		p.SetState(279)
		p.Match(pikeParserT__24)
	}
	{
		p.SetState(280)
		p.Match(pikeParserT__4)
	}
	{
		p.SetState(281)
		p.Expression()
	}
	{
		p.SetState(282)
		p.Match(pikeParserT__5)
	}
	{
		p.SetState(283)
		p.Block()
	}

	return localctx
}

// ICase_stmtContext is an interface to support dynamic dispatch.
type ICase_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCase_stmtContext differentiates from other interfaces.
	IsCase_stmtContext()
}

type Case_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCase_stmtContext() *Case_stmtContext {
	var p = new(Case_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_case_stmt
	return p
}

func (*Case_stmtContext) IsCase_stmtContext() {}

func NewCase_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Case_stmtContext {
	var p = new(Case_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_case_stmt

	return p
}

func (s *Case_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Case_stmtContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Case_stmtContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Case_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Case_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Case_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterCase_stmt(s)
	}
}

func (s *Case_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitCase_stmt(s)
	}
}

func (p *pikeParser) Case_stmt() (localctx ICase_stmtContext) {
	this := p
	_ = this

	localctx = NewCase_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, pikeParserRULE_case_stmt)
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
		p.SetState(285)
		p.Match(pikeParserT__25)
	}
	{
		p.SetState(286)
		p.Expression()
	}
	p.SetState(289)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__26 {
		{
			p.SetState(287)
			p.Match(pikeParserT__26)
		}
		{
			p.SetState(288)
			p.Expression()
		}

	}
	{
		p.SetState(291)
		p.Match(pikeParserT__3)
	}

	return localctx
}

// IDefault_stmtContext is an interface to support dynamic dispatch.
type IDefault_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefault_stmtContext differentiates from other interfaces.
	IsDefault_stmtContext()
}

type Default_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_stmtContext() *Default_stmtContext {
	var p = new(Default_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_default_stmt
	return p
}

func (*Default_stmtContext) IsDefault_stmtContext() {}

func NewDefault_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_stmtContext {
	var p = new(Default_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_default_stmt

	return p
}

func (s *Default_stmtContext) GetParser() antlr.Parser { return s.parser }
func (s *Default_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Default_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterDefault_stmt(s)
	}
}

func (s *Default_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitDefault_stmt(s)
	}
}

func (p *pikeParser) Default_stmt() (localctx IDefault_stmtContext) {
	this := p
	_ = this

	localctx = NewDefault_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, pikeParserRULE_default_stmt)

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
		p.SetState(293)
		p.Match(pikeParserT__27)
	}
	{
		p.SetState(294)
		p.Match(pikeParserT__3)
	}

	return localctx
}

// IForeach_stmtContext is an interface to support dynamic dispatch.
type IForeach_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForeach_stmtContext differentiates from other interfaces.
	IsForeach_stmtContext()
}

type Foreach_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForeach_stmtContext() *Foreach_stmtContext {
	var p = new(Foreach_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_foreach_stmt
	return p
}

func (*Foreach_stmtContext) IsForeach_stmtContext() {}

func NewForeach_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Foreach_stmtContext {
	var p = new(Foreach_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_foreach_stmt

	return p
}

func (s *Foreach_stmtContext) GetParser() antlr.Parser { return s.parser }

func (s *Foreach_stmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Foreach_stmtContext) Expression6() IExpression6Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression6Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression6Context)
}

func (s *Foreach_stmtContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *Foreach_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Foreach_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Foreach_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterForeach_stmt(s)
	}
}

func (s *Foreach_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitForeach_stmt(s)
	}
}

func (p *pikeParser) Foreach_stmt() (localctx IForeach_stmtContext) {
	this := p
	_ = this

	localctx = NewForeach_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, pikeParserRULE_foreach_stmt)

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
		p.SetState(296)
		p.Match(pikeParserT__28)
	}
	{
		p.SetState(297)
		p.Match(pikeParserT__4)
	}
	{
		p.SetState(298)
		p.Expression()
	}
	{
		p.SetState(299)
		p.Match(pikeParserT__3)
	}
	{
		p.SetState(300)
		p.Expression6()
	}
	{
		p.SetState(301)
		p.Match(pikeParserT__5)
	}
	{
		p.SetState(302)
		p.Statement()
	}

	return localctx
}

// IBreak_stmtContext is an interface to support dynamic dispatch.
type IBreak_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreak_stmtContext differentiates from other interfaces.
	IsBreak_stmtContext()
}

type Break_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_stmtContext() *Break_stmtContext {
	var p = new(Break_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_break_stmt
	return p
}

func (*Break_stmtContext) IsBreak_stmtContext() {}

func NewBreak_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_stmtContext {
	var p = new(Break_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_break_stmt

	return p
}

func (s *Break_stmtContext) GetParser() antlr.Parser { return s.parser }
func (s *Break_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterBreak_stmt(s)
	}
}

func (s *Break_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitBreak_stmt(s)
	}
}

func (p *pikeParser) Break_stmt() (localctx IBreak_stmtContext) {
	this := p
	_ = this

	localctx = NewBreak_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, pikeParserRULE_break_stmt)

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
		p.SetState(304)
		p.Match(pikeParserT__29)
	}
	{
		p.SetState(305)
		p.Match(pikeParserT__1)
	}

	return localctx
}

// IContinue_stmtContext is an interface to support dynamic dispatch.
type IContinue_stmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinue_stmtContext differentiates from other interfaces.
	IsContinue_stmtContext()
}

type Continue_stmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinue_stmtContext() *Continue_stmtContext {
	var p = new(Continue_stmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_continue_stmt
	return p
}

func (*Continue_stmtContext) IsContinue_stmtContext() {}

func NewContinue_stmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Continue_stmtContext {
	var p = new(Continue_stmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_continue_stmt

	return p
}

func (s *Continue_stmtContext) GetParser() antlr.Parser { return s.parser }
func (s *Continue_stmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Continue_stmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Continue_stmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterContinue_stmt(s)
	}
}

func (s *Continue_stmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitContinue_stmt(s)
	}
}

func (p *pikeParser) Continue_stmt() (localctx IContinue_stmtContext) {
	this := p
	_ = this

	localctx = NewContinue_stmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, pikeParserRULE_continue_stmt)

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
		p.SetState(307)
		p.Match(pikeParserT__30)
	}
	{
		p.SetState(308)
		p.Match(pikeParserT__1)
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
	p.RuleIndex = pikeParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllExpression2() []IExpression2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression2Context)(nil)).Elem())
	var tst = make([]IExpression2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression2Context)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression2(i int) IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *pikeParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, pikeParserRULE_expression)

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
		p.SetState(310)
		p.Expression2()
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(311)
				p.Match(pikeParserT__6)
			}
			{
				p.SetState(312)
				p.Expression2()
			}

		}
		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

// IExpression2Context is an interface to support dynamic dispatch.
type IExpression2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression2Context differentiates from other interfaces.
	IsExpression2Context()
}

type Expression2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression2Context() *Expression2Context {
	var p = new(Expression2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_expression2
	return p
}

func (*Expression2Context) IsExpression2Context() {}

func NewExpression2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression2Context {
	var p = new(Expression2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_expression2

	return p
}

func (s *Expression2Context) GetParser() antlr.Parser { return s.parser }

func (s *Expression2Context) Expression3() IExpression3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression3Context)
}

func (s *Expression2Context) AllLvalue() []ILvalueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILvalueContext)(nil)).Elem())
	var tst = make([]ILvalueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILvalueContext)
		}
	}

	return tst
}

func (s *Expression2Context) Lvalue(i int) ILvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILvalueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *Expression2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterExpression2(s)
	}
}

func (s *Expression2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitExpression2(s)
	}
}

func (p *pikeParser) Expression2() (localctx IExpression2Context) {
	this := p
	_ = this

	localctx = NewExpression2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, pikeParserRULE_expression2)
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
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(318)
				p.Lvalue()
			}
			{
				p.SetState(319)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-9)&-(0x1f+1)) == 0 && ((1<<uint((_la-9)))&((1<<(pikeParserT__8-9))|(1<<(pikeParserT__31-9))|(1<<(pikeParserT__32-9))|(1<<(pikeParserT__33-9))|(1<<(pikeParserT__34-9))|(1<<(pikeParserT__35-9))|(1<<(pikeParserT__36-9))|(1<<(pikeParserT__37-9))|(1<<(pikeParserT__38-9))|(1<<(pikeParserT__39-9)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}
	{
		p.SetState(326)
		p.Expression3()
	}

	return localctx
}

// IExpression3Context is an interface to support dynamic dispatch.
type IExpression3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression3Context differentiates from other interfaces.
	IsExpression3Context()
}

type Expression3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression3Context() *Expression3Context {
	var p = new(Expression3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_expression3
	return p
}

func (*Expression3Context) IsExpression3Context() {}

func NewExpression3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression3Context {
	var p = new(Expression3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_expression3

	return p
}

func (s *Expression3Context) GetParser() antlr.Parser { return s.parser }

func (s *Expression3Context) Expression4() IExpression4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression4Context)
}

func (s *Expression3Context) AllExpression3() []IExpression3Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression3Context)(nil)).Elem())
	var tst = make([]IExpression3Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression3Context)
		}
	}

	return tst
}

func (s *Expression3Context) Expression3(i int) IExpression3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression3Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpression3Context)
}

func (s *Expression3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterExpression3(s)
	}
}

func (s *Expression3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitExpression3(s)
	}
}

func (p *pikeParser) Expression3() (localctx IExpression3Context) {
	this := p
	_ = this

	localctx = NewExpression3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, pikeParserRULE_expression3)
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
		p.SetState(328)
		p.Expression4()
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__40 {
		{
			p.SetState(329)
			p.Match(pikeParserT__40)
		}
		{
			p.SetState(330)
			p.Expression3()
		}
		{
			p.SetState(331)
			p.Match(pikeParserT__3)
		}
		{
			p.SetState(332)
			p.Expression3()
		}

	}

	return localctx
}

// IExpression4Context is an interface to support dynamic dispatch.
type IExpression4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression4Context differentiates from other interfaces.
	IsExpression4Context()
}

type Expression4Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression4Context() *Expression4Context {
	var p = new(Expression4Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_expression4
	return p
}

func (*Expression4Context) IsExpression4Context() {}

func NewExpression4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression4Context {
	var p = new(Expression4Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_expression4

	return p
}

func (s *Expression4Context) GetParser() antlr.Parser { return s.parser }

func (s *Expression4Context) AllExpression5() []IExpression5Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression5Context)(nil)).Elem())
	var tst = make([]IExpression5Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression5Context)
		}
	}

	return tst
}

func (s *Expression4Context) Expression5(i int) IExpression5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression5Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpression5Context)
}

func (s *Expression4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterExpression4(s)
	}
}

func (s *Expression4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitExpression4(s)
	}
}

func (p *pikeParser) Expression4() (localctx IExpression4Context) {
	this := p
	_ = this

	localctx = NewExpression4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, pikeParserRULE_expression4)
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
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(336)
				p.Expression5()
			}
			{
				p.SetState(337)
				_la = p.GetTokenStream().LA(1)

				if !(_la == pikeParserT__7 || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(pikeParserT__41-42))|(1<<(pikeParserT__42-42))|(1<<(pikeParserT__43-42))|(1<<(pikeParserT__44-42))|(1<<(pikeParserT__45-42))|(1<<(pikeParserT__46-42))|(1<<(pikeParserT__47-42))|(1<<(pikeParserT__48-42))|(1<<(pikeParserT__49-42))|(1<<(pikeParserT__50-42))|(1<<(pikeParserT__51-42))|(1<<(pikeParserT__52-42))|(1<<(pikeParserT__53-42))|(1<<(pikeParserT__54-42))|(1<<(pikeParserT__55-42))|(1<<(pikeParserT__56-42)))) != 0)) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}
	{
		p.SetState(344)
		p.Expression5()
	}

	return localctx
}

// IExpression5Context is an interface to support dynamic dispatch.
type IExpression5Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression5Context differentiates from other interfaces.
	IsExpression5Context()
}

type Expression5Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression5Context() *Expression5Context {
	var p = new(Expression5Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_expression5
	return p
}

func (*Expression5Context) IsExpression5Context() {}

func NewExpression5Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression5Context {
	var p = new(Expression5Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_expression5

	return p
}

func (s *Expression5Context) GetParser() antlr.Parser { return s.parser }

func (s *Expression5Context) Expression6() IExpression6Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression6Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression6Context)
}

func (s *Expression5Context) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Expression5Context) Expression5() IExpression5Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression5Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression5Context)
}

func (s *Expression5Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression5Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression5Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterExpression5(s)
	}
}

func (s *Expression5Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitExpression5(s)
	}
}

func (p *pikeParser) Expression5() (localctx IExpression5Context) {
	this := p
	_ = this

	localctx = NewExpression5Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, pikeParserRULE_expression5)

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

	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(346)
			p.Expression6()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(347)
			p.Match(pikeParserT__4)
		}
		{
			p.SetState(348)
			p.Type_()
		}
		{
			p.SetState(349)
			p.Match(pikeParserT__5)
		}
		{
			p.SetState(350)
			p.Expression5()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(352)
			p.Match(pikeParserT__57)
		}
		{
			p.SetState(353)
			p.Expression6()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(354)
			p.Match(pikeParserT__58)
		}
		{
			p.SetState(355)
			p.Expression6()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(356)
			p.Expression6()
		}
		{
			p.SetState(357)
			p.Match(pikeParserT__57)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(359)
			p.Expression6()
		}
		{
			p.SetState(360)
			p.Match(pikeParserT__58)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(362)
			p.Match(pikeParserT__59)
		}
		{
			p.SetState(363)
			p.Expression5()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(364)
			p.Match(pikeParserT__60)
		}
		{
			p.SetState(365)
			p.Expression5()
		}

	}

	return localctx
}

// IExpression6Context is an interface to support dynamic dispatch.
type IExpression6Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression6Context differentiates from other interfaces.
	IsExpression6Context()
}

type Expression6Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression6Context() *Expression6Context {
	var p = new(Expression6Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_expression6
	return p
}

func (*Expression6Context) IsExpression6Context() {}

func NewExpression6Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression6Context {
	var p = new(Expression6Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_expression6

	return p
}

func (s *Expression6Context) GetParser() antlr.Parser { return s.parser }

func (s *Expression6Context) STRING() antlr.TerminalNode {
	return s.GetToken(pikeParserSTRING, 0)
}

func (s *Expression6Context) NUMBER() antlr.TerminalNode {
	return s.GetToken(pikeParserNUMBER, 0)
}

func (s *Expression6Context) FLOAT() antlr.TerminalNode {
	return s.GetToken(pikeParserFLOAT, 0)
}

func (s *Expression6Context) Catch_() ICatch_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICatch_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICatch_Context)
}

func (s *Expression6Context) Gauge() IGaugeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGaugeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGaugeContext)
}

func (s *Expression6Context) Sscanf() ISscanfContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISscanfContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISscanfContext)
}

func (s *Expression6Context) Lambda() ILambdaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILambdaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILambdaContext)
}

func (s *Expression6Context) Class_implementation() IClass_implementationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClass_implementationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClass_implementationContext)
}

func (s *Expression6Context) Constant_identifier() IConstant_identifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstant_identifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstant_identifierContext)
}

func (s *Expression6Context) Mapping() IMappingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMappingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMappingContext)
}

func (s *Expression6Context) Multiset() IMultisetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultisetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultisetContext)
}

func (s *Expression6Context) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *Expression6Context) Parenthesis() IParenthesisContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParenthesisContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParenthesisContext)
}

func (s *Expression6Context) AllExtension() []IExtensionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExtensionContext)(nil)).Elem())
	var tst = make([]IExtensionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExtensionContext)
		}
	}

	return tst
}

func (s *Expression6Context) Extension(i int) IExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtensionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExtensionContext)
}

func (s *Expression6Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression6Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression6Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterExpression6(s)
	}
}

func (s *Expression6Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitExpression6(s)
	}
}

func (p *pikeParser) Expression6() (localctx IExpression6Context) {
	this := p
	_ = this

	localctx = NewExpression6Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, pikeParserRULE_expression6)
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
	p.SetState(381)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pikeParserSTRING:
		{
			p.SetState(368)
			p.Match(pikeParserSTRING)
		}

	case pikeParserNUMBER:
		{
			p.SetState(369)
			p.Match(pikeParserNUMBER)
		}

	case pikeParserFLOAT:
		{
			p.SetState(370)
			p.Match(pikeParserFLOAT)
		}

	case pikeParserT__64:
		{
			p.SetState(371)
			p.Catch_()
		}

	case pikeParserT__65:
		{
			p.SetState(372)
			p.Gauge()
		}

	case pikeParserT__66:
		{
			p.SetState(373)
			p.Sscanf()
		}

	case pikeParserT__67:
		{
			p.SetState(374)
			p.Lambda()
		}

	case pikeParserT__10:
		{
			p.SetState(375)
			p.Class_implementation()
		}

	case pikeParserIDENTIFIER:
		{
			p.SetState(376)
			p.Constant_identifier()
		}

	case pikeParserT__73:
		{
			p.SetState(377)
			p.Mapping()
		}

	case pikeParserT__71:
		{
			p.SetState(378)
			p.Multiset()
		}

	case pikeParserT__69:
		{
			p.SetState(379)
			p.Array()
		}

	case pikeParserT__4:
		{
			p.SetState(380)
			p.Parenthesis()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pikeParserT__4 || _la == pikeParserT__61 || _la == pikeParserT__62 {
		{
			p.SetState(383)
			p.Extension()
		}

		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExtensionContext is an interface to support dynamic dispatch.
type IExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtensionContext differentiates from other interfaces.
	IsExtensionContext()
}

type ExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionContext() *ExtensionContext {
	var p = new(ExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_extension
	return p
}

func (*ExtensionContext) IsExtensionContext() {}

func NewExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionContext {
	var p = new(ExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_extension

	return p
}

func (s *ExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionContext) Expression_list() IExpression_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *ExtensionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pikeParserIDENTIFIER, 0)
}

func (s *ExtensionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExtensionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterExtension(s)
	}
}

func (s *ExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitExtension(s)
	}
}

func (p *pikeParser) Extension() (localctx IExtensionContext) {
	this := p
	_ = this

	localctx = NewExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, pikeParserRULE_extension)
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

	p.SetState(403)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pikeParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(389)
			p.Match(pikeParserT__4)
		}
		{
			p.SetState(390)
			p.Expression_list()
		}
		{
			p.SetState(391)
			p.Match(pikeParserT__5)
		}

	case pikeParserT__61:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(393)
			p.Match(pikeParserT__61)
		}
		{
			p.SetState(394)
			p.Match(pikeParserIDENTIFIER)
		}

	case pikeParserT__62:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(395)
			p.Match(pikeParserT__62)
		}
		{
			p.SetState(396)
			p.Expression()
		}
		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == pikeParserT__26 {
			{
				p.SetState(397)
				p.Match(pikeParserT__26)
			}
			{
				p.SetState(398)
				p.Expression()
			}

		}
		{
			p.SetState(401)
			p.Match(pikeParserT__63)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICatch_Context is an interface to support dynamic dispatch.
type ICatch_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCatch_Context differentiates from other interfaces.
	IsCatch_Context()
}

type Catch_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCatch_Context() *Catch_Context {
	var p = new(Catch_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_catch_
	return p
}

func (*Catch_Context) IsCatch_Context() {}

func NewCatch_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Catch_Context {
	var p = new(Catch_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_catch_

	return p
}

func (s *Catch_Context) GetParser() antlr.Parser { return s.parser }

func (s *Catch_Context) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Catch_Context) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *Catch_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Catch_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Catch_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterCatch_(s)
	}
}

func (s *Catch_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitCatch_(s)
	}
}

func (p *pikeParser) Catch_() (localctx ICatch_Context) {
	this := p
	_ = this

	localctx = NewCatch_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, pikeParserRULE_catch_)

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
		p.SetState(405)
		p.Match(pikeParserT__64)
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pikeParserT__4:
		{
			p.SetState(406)
			p.Match(pikeParserT__4)
		}
		{
			p.SetState(407)
			p.Expression()
		}
		{
			p.SetState(408)
			p.Match(pikeParserT__5)
		}

	case pikeParserT__11:
		{
			p.SetState(410)
			p.Block()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGaugeContext is an interface to support dynamic dispatch.
type IGaugeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGaugeContext differentiates from other interfaces.
	IsGaugeContext()
}

type GaugeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGaugeContext() *GaugeContext {
	var p = new(GaugeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_gauge
	return p
}

func (*GaugeContext) IsGaugeContext() {}

func NewGaugeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GaugeContext {
	var p = new(GaugeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_gauge

	return p
}

func (s *GaugeContext) GetParser() antlr.Parser { return s.parser }

func (s *GaugeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GaugeContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *GaugeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GaugeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GaugeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterGauge(s)
	}
}

func (s *GaugeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitGauge(s)
	}
}

func (p *pikeParser) Gauge() (localctx IGaugeContext) {
	this := p
	_ = this

	localctx = NewGaugeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, pikeParserRULE_gauge)

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
		p.Match(pikeParserT__65)
	}
	p.SetState(419)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pikeParserT__4:
		{
			p.SetState(414)
			p.Match(pikeParserT__4)
		}
		{
			p.SetState(415)
			p.Expression()
		}
		{
			p.SetState(416)
			p.Match(pikeParserT__5)
		}

	case pikeParserT__11:
		{
			p.SetState(418)
			p.Block()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISscanfContext is an interface to support dynamic dispatch.
type ISscanfContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSscanfContext differentiates from other interfaces.
	IsSscanfContext()
}

type SscanfContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySscanfContext() *SscanfContext {
	var p = new(SscanfContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_sscanf
	return p
}

func (*SscanfContext) IsSscanfContext() {}

func NewSscanfContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SscanfContext {
	var p = new(SscanfContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_sscanf

	return p
}

func (s *SscanfContext) GetParser() antlr.Parser { return s.parser }

func (s *SscanfContext) AllExpression2() []IExpression2Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpression2Context)(nil)).Elem())
	var tst = make([]IExpression2Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpression2Context)
		}
	}

	return tst
}

func (s *SscanfContext) Expression2(i int) IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *SscanfContext) AllLvalue() []ILvalueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILvalueContext)(nil)).Elem())
	var tst = make([]ILvalueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILvalueContext)
		}
	}

	return tst
}

func (s *SscanfContext) Lvalue(i int) ILvalueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILvalueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *SscanfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SscanfContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SscanfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterSscanf(s)
	}
}

func (s *SscanfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitSscanf(s)
	}
}

func (p *pikeParser) Sscanf() (localctx ISscanfContext) {
	this := p
	_ = this

	localctx = NewSscanfContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, pikeParserRULE_sscanf)
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
		p.SetState(421)
		p.Match(pikeParserT__66)
	}
	{
		p.SetState(422)
		p.Match(pikeParserT__4)
	}
	{
		p.SetState(423)
		p.Expression2()
	}
	{
		p.SetState(424)
		p.Match(pikeParserT__6)
	}
	{
		p.SetState(425)
		p.Expression2()
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pikeParserT__6 {
		{
			p.SetState(426)
			p.Match(pikeParserT__6)
		}
		{
			p.SetState(427)
			p.Lvalue()
		}

		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(433)
		p.Match(pikeParserT__5)
	}

	return localctx
}

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_lvalue
	return p
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) Expression6() IExpression6Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression6Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression6Context)
}

func (s *LvalueContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *LvalueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pikeParserIDENTIFIER, 0)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterLvalue(s)
	}
}

func (s *LvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitLvalue(s)
	}
}

func (p *pikeParser) Lvalue() (localctx ILvalueContext) {
	this := p
	_ = this

	localctx = NewLvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, pikeParserRULE_lvalue)

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

	p.SetState(440)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pikeParserT__67:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)
			p.Match(pikeParserT__67)
		}
		{
			p.SetState(436)
			p.Expression6()
		}

	case pikeParserT__77, pikeParserT__78, pikeParserT__79, pikeParserT__80, pikeParserT__81, pikeParserT__82, pikeParserT__83, pikeParserT__84, pikeParserT__85:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(437)
			p.Type_()
		}
		{
			p.SetState(438)
			p.Match(pikeParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_lambda
	return p
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *LambdaContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterLambda(s)
	}
}

func (s *LambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitLambda(s)
	}
}

func (p *pikeParser) Lambda() (localctx ILambdaContext) {
	this := p
	_ = this

	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, pikeParserRULE_lambda)

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
		p.SetState(442)
		p.Match(pikeParserT__67)
	}
	{
		p.SetState(443)
		p.Match(pikeParserT__4)
	}
	{
		p.SetState(444)
		p.Arguments()
	}
	{
		p.SetState(445)
		p.Match(pikeParserT__5)
	}
	{
		p.SetState(446)
		p.Block()
	}

	return localctx
}

// IConstant_identifierContext is an interface to support dynamic dispatch.
type IConstant_identifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstant_identifierContext differentiates from other interfaces.
	IsConstant_identifierContext()
}

type Constant_identifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstant_identifierContext() *Constant_identifierContext {
	var p = new(Constant_identifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_constant_identifier
	return p
}

func (*Constant_identifierContext) IsConstant_identifierContext() {}

func NewConstant_identifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Constant_identifierContext {
	var p = new(Constant_identifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_constant_identifier

	return p
}

func (s *Constant_identifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Constant_identifierContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(pikeParserIDENTIFIER)
}

func (s *Constant_identifierContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(pikeParserIDENTIFIER, i)
}

func (s *Constant_identifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constant_identifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Constant_identifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterConstant_identifier(s)
	}
}

func (s *Constant_identifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitConstant_identifier(s)
	}
}

func (p *pikeParser) Constant_identifier() (localctx IConstant_identifierContext) {
	this := p
	_ = this

	localctx = NewConstant_identifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, pikeParserRULE_constant_identifier)
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
		p.SetState(448)
		p.Match(pikeParserIDENTIFIER)
	}
	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pikeParserT__68 {
		{
			p.SetState(449)
			p.Match(pikeParserT__68)
		}
		{
			p.SetState(450)
			p.Match(pikeParserIDENTIFIER)
		}

		p.SetState(455)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) Expression_list() IExpression_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitArray(s)
	}
}

func (p *pikeParser) Array() (localctx IArrayContext) {
	this := p
	_ = this

	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, pikeParserRULE_array)

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
		p.SetState(456)
		p.Match(pikeParserT__69)
	}
	{
		p.SetState(457)
		p.Expression_list()
	}
	{
		p.SetState(458)
		p.Match(pikeParserT__70)
	}

	return localctx
}

// IMultisetContext is an interface to support dynamic dispatch.
type IMultisetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultisetContext differentiates from other interfaces.
	IsMultisetContext()
}

type MultisetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultisetContext() *MultisetContext {
	var p = new(MultisetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_multiset
	return p
}

func (*MultisetContext) IsMultisetContext() {}

func NewMultisetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultisetContext {
	var p = new(MultisetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_multiset

	return p
}

func (s *MultisetContext) GetParser() antlr.Parser { return s.parser }

func (s *MultisetContext) Expression_list() IExpression_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *MultisetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultisetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultisetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterMultiset(s)
	}
}

func (s *MultisetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitMultiset(s)
	}
}

func (p *pikeParser) Multiset() (localctx IMultisetContext) {
	this := p
	_ = this

	localctx = NewMultisetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, pikeParserRULE_multiset)

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
		p.SetState(460)
		p.Match(pikeParserT__71)
	}
	{
		p.SetState(461)
		p.Expression_list()
	}
	{
		p.SetState(462)
		p.Match(pikeParserT__72)
	}

	return localctx
}

// IMappingContext is an interface to support dynamic dispatch.
type IMappingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMappingContext differentiates from other interfaces.
	IsMappingContext()
}

type MappingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMappingContext() *MappingContext {
	var p = new(MappingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_mapping
	return p
}

func (*MappingContext) IsMappingContext() {}

func NewMappingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MappingContext {
	var p = new(MappingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_mapping

	return p
}

func (s *MappingContext) GetParser() antlr.Parser { return s.parser }

func (s *MappingContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MappingContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MappingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MappingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MappingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterMapping(s)
	}
}

func (s *MappingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitMapping(s)
	}
}

func (p *pikeParser) Mapping() (localctx IMappingContext) {
	this := p
	_ = this

	localctx = NewMappingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, pikeParserRULE_mapping)
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
		p.SetState(464)
		p.Match(pikeParserT__73)
	}
	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__4 || _la == pikeParserT__10 || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(pikeParserT__57-58))|(1<<(pikeParserT__58-58))|(1<<(pikeParserT__59-58))|(1<<(pikeParserT__60-58))|(1<<(pikeParserT__64-58))|(1<<(pikeParserT__65-58))|(1<<(pikeParserT__66-58))|(1<<(pikeParserT__67-58))|(1<<(pikeParserT__69-58))|(1<<(pikeParserT__71-58))|(1<<(pikeParserT__73-58))|(1<<(pikeParserT__77-58))|(1<<(pikeParserT__78-58))|(1<<(pikeParserT__79-58))|(1<<(pikeParserT__80-58))|(1<<(pikeParserT__81-58))|(1<<(pikeParserT__82-58))|(1<<(pikeParserT__83-58))|(1<<(pikeParserT__84-58))|(1<<(pikeParserT__85-58))|(1<<(pikeParserIDENTIFIER-58)))) != 0) || (((_la-90)&-(0x1f+1)) == 0 && ((1<<uint((_la-90)))&((1<<(pikeParserFLOAT-90))|(1<<(pikeParserNUMBER-90))|(1<<(pikeParserSTRING-90)))) != 0) {
		{
			p.SetState(465)
			p.Expression()
		}
		{
			p.SetState(466)
			p.Match(pikeParserT__3)
		}
		{
			p.SetState(467)
			p.Expression()
		}
		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(468)
					p.Match(pikeParserT__6)
				}
				{
					p.SetState(469)
					p.Expression()
				}
				{
					p.SetState(470)
					p.Match(pikeParserT__3)
				}
				{
					p.SetState(471)
					p.Expression()
				}

			}
			p.SetState(477)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
		}

	}
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__6 {
		{
			p.SetState(480)
			p.Match(pikeParserT__6)
		}

	}
	{
		p.SetState(483)
		p.Match(pikeParserT__74)
	}

	return localctx
}

// IProgram_specifierContext is an interface to support dynamic dispatch.
type IProgram_specifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgram_specifierContext differentiates from other interfaces.
	IsProgram_specifierContext()
}

type Program_specifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgram_specifierContext() *Program_specifierContext {
	var p = new(Program_specifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_program_specifier
	return p
}

func (*Program_specifierContext) IsProgram_specifierContext() {}

func NewProgram_specifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Program_specifierContext {
	var p = new(Program_specifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_program_specifier

	return p
}

func (s *Program_specifierContext) GetParser() antlr.Parser { return s.parser }

func (s *Program_specifierContext) Constant_identifier() IConstant_identifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstant_identifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstant_identifierContext)
}

func (s *Program_specifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Program_specifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Program_specifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterProgram_specifier(s)
	}
}

func (s *Program_specifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitProgram_specifier(s)
	}
}

func (p *pikeParser) Program_specifier() (localctx IProgram_specifierContext) {
	this := p
	_ = this

	localctx = NewProgram_specifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, pikeParserRULE_program_specifier)

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
		p.SetState(485)
		p.Constant_identifier()
	}

	return localctx
}

// IParenthesisContext is an interface to support dynamic dispatch.
type IParenthesisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParenthesisContext differentiates from other interfaces.
	IsParenthesisContext()
}

type ParenthesisContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParenthesisContext() *ParenthesisContext {
	var p = new(ParenthesisContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_parenthesis
	return p
}

func (*ParenthesisContext) IsParenthesisContext() {}

func NewParenthesisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParenthesisContext {
	var p = new(ParenthesisContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_parenthesis

	return p
}

func (s *ParenthesisContext) GetParser() antlr.Parser { return s.parser }

func (s *ParenthesisContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterParenthesis(s)
	}
}

func (s *ParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitParenthesis(s)
	}
}

func (p *pikeParser) Parenthesis() (localctx IParenthesisContext) {
	this := p
	_ = this

	localctx = NewParenthesisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, pikeParserRULE_parenthesis)

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
		p.SetState(487)
		p.Match(pikeParserT__4)
	}
	{
		p.SetState(488)
		p.Expression()
	}
	{
		p.SetState(489)
		p.Match(pikeParserT__5)
	}

	return localctx
}

// IExpression_listContext is an interface to support dynamic dispatch.
type IExpression_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression_listContext differentiates from other interfaces.
	IsExpression_listContext()
}

type Expression_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_listContext() *Expression_listContext {
	var p = new(Expression_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_expression_list
	return p
}

func (*Expression_listContext) IsExpression_listContext() {}

func NewExpression_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_listContext {
	var p = new(Expression_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_expression_list

	return p
}

func (s *Expression_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_listContext) AllSplice_expression() []ISplice_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISplice_expressionContext)(nil)).Elem())
	var tst = make([]ISplice_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISplice_expressionContext)
		}
	}

	return tst
}

func (s *Expression_listContext) Splice_expression(i int) ISplice_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISplice_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISplice_expressionContext)
}

func (s *Expression_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterExpression_list(s)
	}
}

func (s *Expression_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitExpression_list(s)
	}
}

func (p *pikeParser) Expression_list() (localctx IExpression_listContext) {
	this := p
	_ = this

	localctx = NewExpression_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, pikeParserRULE_expression_list)
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
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__4 || _la == pikeParserT__10 || (((_la-58)&-(0x1f+1)) == 0 && ((1<<uint((_la-58)))&((1<<(pikeParserT__57-58))|(1<<(pikeParserT__58-58))|(1<<(pikeParserT__59-58))|(1<<(pikeParserT__60-58))|(1<<(pikeParserT__64-58))|(1<<(pikeParserT__65-58))|(1<<(pikeParserT__66-58))|(1<<(pikeParserT__67-58))|(1<<(pikeParserT__69-58))|(1<<(pikeParserT__71-58))|(1<<(pikeParserT__73-58))|(1<<(pikeParserT__75-58))|(1<<(pikeParserT__77-58))|(1<<(pikeParserT__78-58))|(1<<(pikeParserT__79-58))|(1<<(pikeParserT__80-58))|(1<<(pikeParserT__81-58))|(1<<(pikeParserT__82-58))|(1<<(pikeParserT__83-58))|(1<<(pikeParserT__84-58))|(1<<(pikeParserT__85-58))|(1<<(pikeParserIDENTIFIER-58)))) != 0) || (((_la-90)&-(0x1f+1)) == 0 && ((1<<uint((_la-90)))&((1<<(pikeParserFLOAT-90))|(1<<(pikeParserNUMBER-90))|(1<<(pikeParserSTRING-90)))) != 0) {
		{
			p.SetState(491)
			p.Splice_expression()
		}
		p.SetState(496)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(492)
					p.Match(pikeParserT__6)
				}
				{
					p.SetState(493)
					p.Splice_expression()
				}

			}
			p.SetState(498)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext())
		}

	}
	p.SetState(502)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__6 {
		{
			p.SetState(501)
			p.Match(pikeParserT__6)
		}

	}

	return localctx
}

// ISplice_expressionContext is an interface to support dynamic dispatch.
type ISplice_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSplice_expressionContext differentiates from other interfaces.
	IsSplice_expressionContext()
}

type Splice_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySplice_expressionContext() *Splice_expressionContext {
	var p = new(Splice_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_splice_expression
	return p
}

func (*Splice_expressionContext) IsSplice_expressionContext() {}

func NewSplice_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Splice_expressionContext {
	var p = new(Splice_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_splice_expression

	return p
}

func (s *Splice_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Splice_expressionContext) Expression2() IExpression2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression2Context)
}

func (s *Splice_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Splice_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Splice_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterSplice_expression(s)
	}
}

func (s *Splice_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitSplice_expression(s)
	}
}

func (p *pikeParser) Splice_expression() (localctx ISplice_expressionContext) {
	this := p
	_ = this

	localctx = NewSplice_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, pikeParserRULE_splice_expression)
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
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__75 {
		{
			p.SetState(504)
			p.Match(pikeParserT__75)
		}

	}
	{
		p.SetState(507)
		p.Expression2()
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *ArgumentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(pikeParserIDENTIFIER, 0)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (p *pikeParser) Argument() (localctx IArgumentContext) {
	this := p
	_ = this

	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, pikeParserRULE_argument)
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
		p.SetState(509)
		p.Type_()
	}
	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__76 {
		{
			p.SetState(510)
			p.Match(pikeParserT__76)
		}

	}
	p.SetState(514)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserIDENTIFIER {
		{
			p.SetState(513)
			p.Match(pikeParserIDENTIFIER)
		}

	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *pikeParser) Arguments() (localctx IArgumentsContext) {
	this := p
	_ = this

	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, pikeParserRULE_arguments)
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
	p.SetState(524)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-78)&-(0x1f+1)) == 0 && ((1<<uint((_la-78)))&((1<<(pikeParserT__77-78))|(1<<(pikeParserT__78-78))|(1<<(pikeParserT__79-78))|(1<<(pikeParserT__80-78))|(1<<(pikeParserT__81-78))|(1<<(pikeParserT__82-78))|(1<<(pikeParserT__83-78))|(1<<(pikeParserT__84-78))|(1<<(pikeParserT__85-78)))) != 0 {
		{
			p.SetState(516)
			p.Argument()
		}
		p.SetState(521)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(517)
					p.Match(pikeParserT__6)
				}
				{
					p.SetState(518)
					p.Argument()
				}

			}
			p.SetState(523)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext())
		}

	}
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__6 {
		{
			p.SetState(526)
			p.Match(pikeParserT__6)
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
	p.RuleIndex = pikeParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) Program_specifier() IProgram_specifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProgram_specifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProgram_specifierContext)
}

func (s *Type_Context) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *Type_Context) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Type_Context) Function_type() IFunction_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_typeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *pikeParser) Type_() (localctx IType_Context) {
	this := p
	_ = this

	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, pikeParserRULE_type_)
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
	p.SetState(567)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case pikeParserT__77:
		{
			p.SetState(529)
			p.Match(pikeParserT__77)
		}

	case pikeParserT__78:
		{
			p.SetState(530)
			p.Match(pikeParserT__78)
		}

	case pikeParserT__79:
		{
			p.SetState(531)
			p.Match(pikeParserT__79)
		}

	case pikeParserT__80:
		{
			p.SetState(532)
			p.Match(pikeParserT__80)
		}

	case pikeParserT__81:
		{
			p.SetState(533)
			p.Match(pikeParserT__81)
		}
		p.SetState(538)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == pikeParserT__4 {
			{
				p.SetState(534)
				p.Match(pikeParserT__4)
			}
			{
				p.SetState(535)
				p.Program_specifier()
			}
			{
				p.SetState(536)
				p.Match(pikeParserT__5)
			}

		}

	case pikeParserT__82:
		{
			p.SetState(540)
			p.Match(pikeParserT__82)
		}
		p.SetState(547)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == pikeParserT__4 {
			{
				p.SetState(541)
				p.Match(pikeParserT__4)
			}
			{
				p.SetState(542)
				p.Type_()
			}
			{
				p.SetState(543)
				p.Match(pikeParserT__3)
			}
			{
				p.SetState(544)
				p.Type_()
			}
			{
				p.SetState(545)
				p.Match(pikeParserT__5)
			}

		}

	case pikeParserT__83:
		{
			p.SetState(549)
			p.Match(pikeParserT__83)
		}
		p.SetState(554)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == pikeParserT__4 {
			{
				p.SetState(550)
				p.Match(pikeParserT__4)
			}
			{
				p.SetState(551)
				p.Type_()
			}
			{
				p.SetState(552)
				p.Match(pikeParserT__5)
			}

		}

	case pikeParserT__84:
		{
			p.SetState(556)
			p.Match(pikeParserT__84)
		}
		p.SetState(561)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == pikeParserT__4 {
			{
				p.SetState(557)
				p.Match(pikeParserT__4)
			}
			{
				p.SetState(558)
				p.Type_()
			}
			{
				p.SetState(559)
				p.Match(pikeParserT__5)
			}

		}

	case pikeParserT__85:
		{
			p.SetState(563)
			p.Match(pikeParserT__85)
		}
		p.SetState(565)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == pikeParserT__4 {
			{
				p.SetState(564)
				p.Function_type()
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(569)
				p.Match(pikeParserT__7)
			}

		}
		p.SetState(574)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())
	}

	return localctx
}

// IFunction_typeContext is an interface to support dynamic dispatch.
type IFunction_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_typeContext differentiates from other interfaces.
	IsFunction_typeContext()
}

type Function_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_typeContext() *Function_typeContext {
	var p = new(Function_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = pikeParserRULE_function_type
	return p
}

func (*Function_typeContext) IsFunction_typeContext() {}

func NewFunction_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_typeContext {
	var p = new(Function_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = pikeParserRULE_function_type

	return p
}

func (s *Function_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_typeContext) AllType_() []IType_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_Context)(nil)).Elem())
	var tst = make([]IType_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_Context)
		}
	}

	return tst
}

func (s *Function_typeContext) Type_(i int) IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *Function_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.EnterFunction_type(s)
	}
}

func (s *Function_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(pikeListener); ok {
		listenerT.ExitFunction_type(s)
	}
}

func (p *pikeParser) Function_type() (localctx IFunction_typeContext) {
	this := p
	_ = this

	localctx = NewFunction_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, pikeParserRULE_function_type)
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
		p.SetState(575)
		p.Match(pikeParserT__4)
	}
	{
		p.SetState(576)
		p.Type_()
	}
	p.SetState(581)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == pikeParserT__6 {
		{
			p.SetState(577)
			p.Match(pikeParserT__6)
		}
		{
			p.SetState(578)
			p.Type_()
		}

		p.SetState(583)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(585)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == pikeParserT__76 {
		{
			p.SetState(584)
			p.Match(pikeParserT__76)
		}

	}
	{
		p.SetState(587)
		p.Match(pikeParserT__5)
	}

	return localctx
}
