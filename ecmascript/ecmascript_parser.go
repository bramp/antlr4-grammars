// Generated from ECMAScript.g4 by ANTLR 4.7.

package ecmascript // ECMAScript
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "strings"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 105, 622,
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
	9, 55, 4, 56, 9, 56, 3, 2, 5, 2, 114, 10, 2, 3, 2, 3, 2, 3, 3, 6, 3, 119,
	10, 3, 13, 3, 14, 3, 120, 3, 4, 3, 4, 5, 4, 125, 10, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 5, 5, 142, 10, 5, 3, 6, 3, 6, 5, 6, 146, 10, 6, 3, 6, 3, 6, 3, 7, 6,
	7, 151, 10, 7, 13, 7, 14, 7, 152, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	9, 7, 9, 162, 10, 9, 12, 9, 14, 9, 165, 11, 9, 3, 10, 3, 10, 5, 10, 169,
	10, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 187, 10, 14, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 206, 10, 15, 3, 15, 3,
	15, 5, 15, 210, 10, 15, 3, 15, 3, 15, 5, 15, 214, 10, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 224, 10, 15, 3, 15, 3,
	15, 5, 15, 228, 10, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 5, 15, 250, 10, 15, 3, 16, 3, 16, 5, 16, 254, 10, 16,
	3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 260, 10, 17, 3, 17, 3, 17, 3, 18, 3,
	18, 5, 18, 266, 10, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 5, 21, 284,
	10, 21, 3, 21, 3, 21, 5, 21, 288, 10, 21, 5, 21, 290, 10, 21, 3, 21, 3,
	21, 3, 22, 6, 22, 295, 10, 22, 13, 22, 14, 22, 296, 3, 23, 3, 23, 3, 23,
	3, 23, 5, 23, 303, 10, 23, 3, 24, 3, 24, 3, 24, 5, 24, 308, 10, 24, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5,
	27, 331, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29,
	3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 349, 10,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 7, 32, 359,
	10, 32, 12, 32, 14, 32, 362, 11, 32, 3, 33, 5, 33, 365, 10, 33, 3, 34,
	3, 34, 5, 34, 369, 10, 34, 3, 34, 5, 34, 372, 10, 34, 3, 34, 5, 34, 375,
	10, 34, 3, 34, 3, 34, 3, 35, 5, 35, 380, 10, 35, 3, 35, 3, 35, 3, 35, 5,
	35, 385, 10, 35, 3, 35, 7, 35, 388, 10, 35, 12, 35, 14, 35, 391, 11, 35,
	3, 36, 6, 36, 394, 10, 36, 13, 36, 14, 36, 395, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 5, 37, 403, 10, 37, 3, 37, 3, 37, 5, 37, 407, 10, 37, 3, 38,
	3, 38, 3, 38, 7, 38, 412, 10, 38, 12, 38, 14, 38, 415, 11, 38, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 436, 10, 39, 3,
	40, 3, 40, 3, 40, 5, 40, 441, 10, 40, 3, 41, 3, 41, 3, 42, 3, 42, 5, 42,
	447, 10, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 7, 43, 454, 10, 43, 12,
	43, 14, 43, 457, 11, 43, 3, 44, 3, 44, 3, 44, 7, 44, 462, 10, 44, 12, 44,
	14, 44, 465, 11, 44, 3, 45, 3, 45, 3, 45, 5, 45, 470, 10, 45, 3, 45, 3,
	45, 5, 45, 474, 10, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 5, 45, 484, 10, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5,
	45, 513, 10, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 7, 45, 580, 10, 45, 12, 45, 14, 45, 583,
	11, 45, 3, 46, 3, 46, 3, 47, 3, 47, 5, 47, 589, 10, 47, 3, 48, 3, 48, 3,
	49, 3, 49, 5, 49, 595, 10, 49, 3, 50, 3, 50, 3, 50, 5, 50, 600, 10, 50,
	3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3,
	54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 5, 55, 618, 10, 55, 3, 56, 3, 56,
	3, 56, 2, 3, 88, 57, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
	66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100,
	102, 104, 106, 108, 110, 2, 13, 3, 2, 23, 25, 3, 2, 19, 20, 3, 2, 26, 28,
	3, 2, 29, 32, 3, 2, 33, 36, 3, 2, 42, 52, 5, 2, 3, 3, 53, 54, 101, 101,
	3, 2, 55, 57, 3, 2, 53, 54, 3, 2, 58, 83, 3, 2, 84, 99, 2, 675, 2, 113,
	3, 2, 2, 2, 4, 118, 3, 2, 2, 2, 6, 124, 3, 2, 2, 2, 8, 141, 3, 2, 2, 2,
	10, 143, 3, 2, 2, 2, 12, 150, 3, 2, 2, 2, 14, 154, 3, 2, 2, 2, 16, 158,
	3, 2, 2, 2, 18, 166, 3, 2, 2, 2, 20, 170, 3, 2, 2, 2, 22, 173, 3, 2, 2,
	2, 24, 175, 3, 2, 2, 2, 26, 179, 3, 2, 2, 2, 28, 249, 3, 2, 2, 2, 30, 251,
	3, 2, 2, 2, 32, 257, 3, 2, 2, 2, 34, 263, 3, 2, 2, 2, 36, 269, 3, 2, 2,
	2, 38, 275, 3, 2, 2, 2, 40, 281, 3, 2, 2, 2, 42, 294, 3, 2, 2, 2, 44, 298,
	3, 2, 2, 2, 46, 304, 3, 2, 2, 2, 48, 309, 3, 2, 2, 2, 50, 313, 3, 2, 2,
	2, 52, 330, 3, 2, 2, 2, 54, 332, 3, 2, 2, 2, 56, 338, 3, 2, 2, 2, 58, 341,
	3, 2, 2, 2, 60, 344, 3, 2, 2, 2, 62, 355, 3, 2, 2, 2, 64, 364, 3, 2, 2,
	2, 66, 366, 3, 2, 2, 2, 68, 379, 3, 2, 2, 2, 70, 393, 3, 2, 2, 2, 72, 406,
	3, 2, 2, 2, 74, 408, 3, 2, 2, 2, 76, 435, 3, 2, 2, 2, 78, 440, 3, 2, 2,
	2, 80, 442, 3, 2, 2, 2, 82, 444, 3, 2, 2, 2, 84, 450, 3, 2, 2, 2, 86, 458,
	3, 2, 2, 2, 88, 512, 3, 2, 2, 2, 90, 584, 3, 2, 2, 2, 92, 588, 3, 2, 2,
	2, 94, 590, 3, 2, 2, 2, 96, 594, 3, 2, 2, 2, 98, 599, 3, 2, 2, 2, 100,
	601, 3, 2, 2, 2, 102, 603, 3, 2, 2, 2, 104, 605, 3, 2, 2, 2, 106, 609,
	3, 2, 2, 2, 108, 617, 3, 2, 2, 2, 110, 619, 3, 2, 2, 2, 112, 114, 5, 4,
	3, 2, 113, 112, 3, 2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2,
	115, 116, 7, 2, 2, 3, 116, 3, 3, 2, 2, 2, 117, 119, 5, 6, 4, 2, 118, 117,
	3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2,
	2, 2, 121, 5, 3, 2, 2, 2, 122, 125, 5, 8, 5, 2, 123, 125, 5, 60, 31, 2,
	124, 122, 3, 2, 2, 2, 124, 123, 3, 2, 2, 2, 125, 7, 3, 2, 2, 2, 126, 142,
	5, 10, 6, 2, 127, 142, 5, 14, 8, 2, 128, 142, 5, 22, 12, 2, 129, 142, 5,
	24, 13, 2, 130, 142, 5, 26, 14, 2, 131, 142, 5, 28, 15, 2, 132, 142, 5,
	30, 16, 2, 133, 142, 5, 32, 17, 2, 134, 142, 5, 34, 18, 2, 135, 142, 5,
	36, 19, 2, 136, 142, 5, 48, 25, 2, 137, 142, 5, 38, 20, 2, 138, 142, 5,
	50, 26, 2, 139, 142, 5, 52, 27, 2, 140, 142, 5, 58, 30, 2, 141, 126, 3,
	2, 2, 2, 141, 127, 3, 2, 2, 2, 141, 128, 3, 2, 2, 2, 141, 129, 3, 2, 2,
	2, 141, 130, 3, 2, 2, 2, 141, 131, 3, 2, 2, 2, 141, 132, 3, 2, 2, 2, 141,
	133, 3, 2, 2, 2, 141, 134, 3, 2, 2, 2, 141, 135, 3, 2, 2, 2, 141, 136,
	3, 2, 2, 2, 141, 137, 3, 2, 2, 2, 141, 138, 3, 2, 2, 2, 141, 139, 3, 2,
	2, 2, 141, 140, 3, 2, 2, 2, 142, 9, 3, 2, 2, 2, 143, 145, 7, 9, 2, 2, 144,
	146, 5, 12, 7, 2, 145, 144, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 147,
	3, 2, 2, 2, 147, 148, 7, 10, 2, 2, 148, 11, 3, 2, 2, 2, 149, 151, 5, 8,
	5, 2, 150, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2,
	152, 153, 3, 2, 2, 2, 153, 13, 3, 2, 2, 2, 154, 155, 7, 65, 2, 2, 155,
	156, 5, 16, 9, 2, 156, 157, 5, 108, 55, 2, 157, 15, 3, 2, 2, 2, 158, 163,
	5, 18, 10, 2, 159, 160, 7, 12, 2, 2, 160, 162, 5, 18, 10, 2, 161, 159,
	3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2,
	2, 2, 164, 17, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 166, 168, 7, 100, 2, 2,
	167, 169, 5, 20, 11, 2, 168, 167, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169,
	19, 3, 2, 2, 2, 170, 171, 7, 13, 2, 2, 171, 172, 5, 88, 45, 2, 172, 21,
	3, 2, 2, 2, 173, 174, 7, 11, 2, 2, 174, 23, 3, 2, 2, 2, 175, 176, 6, 13,
	2, 2, 176, 177, 5, 86, 44, 2, 177, 178, 5, 108, 55, 2, 178, 25, 3, 2, 2,
	2, 179, 180, 7, 79, 2, 2, 180, 181, 7, 7, 2, 2, 181, 182, 5, 86, 44, 2,
	182, 183, 7, 8, 2, 2, 183, 186, 5, 8, 5, 2, 184, 185, 7, 63, 2, 2, 185,
	187, 5, 8, 5, 2, 186, 184, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2, 187, 27, 3,
	2, 2, 2, 188, 189, 7, 59, 2, 2, 189, 190, 5, 8, 5, 2, 190, 191, 7, 73,
	2, 2, 191, 192, 7, 7, 2, 2, 192, 193, 5, 86, 44, 2, 193, 194, 7, 8, 2,
	2, 194, 195, 5, 108, 55, 2, 195, 250, 3, 2, 2, 2, 196, 197, 7, 73, 2, 2,
	197, 198, 7, 7, 2, 2, 198, 199, 5, 86, 44, 2, 199, 200, 7, 8, 2, 2, 200,
	201, 5, 8, 5, 2, 201, 250, 3, 2, 2, 2, 202, 203, 7, 71, 2, 2, 203, 205,
	7, 7, 2, 2, 204, 206, 5, 86, 44, 2, 205, 204, 3, 2, 2, 2, 205, 206, 3,
	2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 209, 7, 11, 2, 2, 208, 210, 5, 86,
	44, 2, 209, 208, 3, 2, 2, 2, 209, 210, 3, 2, 2, 2, 210, 211, 3, 2, 2, 2,
	211, 213, 7, 11, 2, 2, 212, 214, 5, 86, 44, 2, 213, 212, 3, 2, 2, 2, 213,
	214, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2, 215, 216, 7, 8, 2, 2, 216, 250,
	5, 8, 5, 2, 217, 218, 7, 71, 2, 2, 218, 219, 7, 7, 2, 2, 219, 220, 7, 65,
	2, 2, 220, 221, 5, 16, 9, 2, 221, 223, 7, 11, 2, 2, 222, 224, 5, 86, 44,
	2, 223, 222, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225,
	227, 7, 11, 2, 2, 226, 228, 5, 86, 44, 2, 227, 226, 3, 2, 2, 2, 227, 228,
	3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 230, 7, 8, 2, 2, 230, 231, 5, 8,
	5, 2, 231, 250, 3, 2, 2, 2, 232, 233, 7, 71, 2, 2, 233, 234, 7, 7, 2, 2,
	234, 235, 5, 88, 45, 2, 235, 236, 7, 82, 2, 2, 236, 237, 5, 86, 44, 2,
	237, 238, 7, 8, 2, 2, 238, 239, 5, 8, 5, 2, 239, 250, 3, 2, 2, 2, 240,
	241, 7, 71, 2, 2, 241, 242, 7, 7, 2, 2, 242, 243, 7, 65, 2, 2, 243, 244,
	5, 18, 10, 2, 244, 245, 7, 82, 2, 2, 245, 246, 5, 86, 44, 2, 246, 247,
	7, 8, 2, 2, 247, 248, 5, 8, 5, 2, 248, 250, 3, 2, 2, 2, 249, 188, 3, 2,
	2, 2, 249, 196, 3, 2, 2, 2, 249, 202, 3, 2, 2, 2, 249, 217, 3, 2, 2, 2,
	249, 232, 3, 2, 2, 2, 249, 240, 3, 2, 2, 2, 250, 29, 3, 2, 2, 2, 251, 253,
	7, 70, 2, 2, 252, 254, 7, 100, 2, 2, 253, 252, 3, 2, 2, 2, 253, 254, 3,
	2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 256, 5, 108, 55, 2, 256, 31, 3, 2,
	2, 2, 257, 259, 7, 58, 2, 2, 258, 260, 7, 100, 2, 2, 259, 258, 3, 2, 2,
	2, 259, 260, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261, 262, 5, 108, 55, 2,
	262, 33, 3, 2, 2, 2, 263, 265, 7, 68, 2, 2, 264, 266, 5, 86, 44, 2, 265,
	264, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266, 267, 3, 2, 2, 2, 267, 268,
	5, 108, 55, 2, 268, 35, 3, 2, 2, 2, 269, 270, 7, 77, 2, 2, 270, 271, 7,
	7, 2, 2, 271, 272, 5, 86, 44, 2, 272, 273, 7, 8, 2, 2, 273, 274, 5, 8,
	5, 2, 274, 37, 3, 2, 2, 2, 275, 276, 7, 72, 2, 2, 276, 277, 7, 7, 2, 2,
	277, 278, 5, 86, 44, 2, 278, 279, 7, 8, 2, 2, 279, 280, 5, 40, 21, 2, 280,
	39, 3, 2, 2, 2, 281, 283, 7, 9, 2, 2, 282, 284, 5, 42, 22, 2, 283, 282,
	3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 289, 3, 2, 2, 2, 285, 287, 5, 46,
	24, 2, 286, 288, 5, 42, 22, 2, 287, 286, 3, 2, 2, 2, 287, 288, 3, 2, 2,
	2, 288, 290, 3, 2, 2, 2, 289, 285, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290,
	291, 3, 2, 2, 2, 291, 292, 7, 10, 2, 2, 292, 41, 3, 2, 2, 2, 293, 295,
	5, 44, 23, 2, 294, 293, 3, 2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 294, 3,
	2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 43, 3, 2, 2, 2, 298, 299, 7, 62, 2,
	2, 299, 300, 5, 86, 44, 2, 300, 302, 7, 15, 2, 2, 301, 303, 5, 12, 7, 2,
	302, 301, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303, 45, 3, 2, 2, 2, 304, 305,
	7, 78, 2, 2, 305, 307, 7, 15, 2, 2, 306, 308, 5, 12, 7, 2, 307, 306, 3,
	2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 47, 3, 2, 2, 2, 309, 310, 7, 100, 2,
	2, 310, 311, 7, 15, 2, 2, 311, 312, 5, 8, 5, 2, 312, 49, 3, 2, 2, 2, 313,
	314, 7, 80, 2, 2, 314, 315, 5, 86, 44, 2, 315, 316, 5, 108, 55, 2, 316,
	51, 3, 2, 2, 2, 317, 318, 7, 83, 2, 2, 318, 319, 5, 10, 6, 2, 319, 320,
	5, 54, 28, 2, 320, 331, 3, 2, 2, 2, 321, 322, 7, 83, 2, 2, 322, 323, 5,
	10, 6, 2, 323, 324, 5, 56, 29, 2, 324, 331, 3, 2, 2, 2, 325, 326, 7, 83,
	2, 2, 326, 327, 5, 10, 6, 2, 327, 328, 5, 54, 28, 2, 328, 329, 5, 56, 29,
	2, 329, 331, 3, 2, 2, 2, 330, 317, 3, 2, 2, 2, 330, 321, 3, 2, 2, 2, 330,
	325, 3, 2, 2, 2, 331, 53, 3, 2, 2, 2, 332, 333, 7, 66, 2, 2, 333, 334,
	7, 7, 2, 2, 334, 335, 7, 100, 2, 2, 335, 336, 7, 8, 2, 2, 336, 337, 5,
	10, 6, 2, 337, 55, 3, 2, 2, 2, 338, 339, 7, 67, 2, 2, 339, 340, 5, 10,
	6, 2, 340, 57, 3, 2, 2, 2, 341, 342, 7, 74, 2, 2, 342, 343, 5, 108, 55,
	2, 343, 59, 3, 2, 2, 2, 344, 345, 7, 75, 2, 2, 345, 346, 7, 100, 2, 2,
	346, 348, 7, 7, 2, 2, 347, 349, 5, 62, 32, 2, 348, 347, 3, 2, 2, 2, 348,
	349, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350, 351, 7, 8, 2, 2, 351, 352,
	7, 9, 2, 2, 352, 353, 5, 64, 33, 2, 353, 354, 7, 10, 2, 2, 354, 61, 3,
	2, 2, 2, 355, 360, 7, 100, 2, 2, 356, 357, 7, 12, 2, 2, 357, 359, 7, 100,
	2, 2, 358, 356, 3, 2, 2, 2, 359, 362, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2,
	360, 361, 3, 2, 2, 2, 361, 63, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 363, 365,
	5, 4, 3, 2, 364, 363, 3, 2, 2, 2, 364, 365, 3, 2, 2, 2, 365, 65, 3, 2,
	2, 2, 366, 368, 7, 5, 2, 2, 367, 369, 5, 68, 35, 2, 368, 367, 3, 2, 2,
	2, 368, 369, 3, 2, 2, 2, 369, 371, 3, 2, 2, 2, 370, 372, 7, 12, 2, 2, 371,
	370, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 374, 3, 2, 2, 2, 373, 375,
	5, 70, 36, 2, 374, 373, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375, 376, 3,
	2, 2, 2, 376, 377, 7, 6, 2, 2, 377, 67, 3, 2, 2, 2, 378, 380, 5, 70, 36,
	2, 379, 378, 3, 2, 2, 2, 379, 380, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381,
	389, 5, 88, 45, 2, 382, 384, 7, 12, 2, 2, 383, 385, 5, 70, 36, 2, 384,
	383, 3, 2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 386, 3, 2, 2, 2, 386, 388,
	5, 88, 45, 2, 387, 382, 3, 2, 2, 2, 388, 391, 3, 2, 2, 2, 389, 387, 3,
	2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 69, 3, 2, 2, 2, 391, 389, 3, 2, 2,
	2, 392, 394, 7, 12, 2, 2, 393, 392, 3, 2, 2, 2, 394, 395, 3, 2, 2, 2, 395,
	393, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 71, 3, 2, 2, 2, 397, 398, 7,
	9, 2, 2, 398, 407, 7, 10, 2, 2, 399, 400, 7, 9, 2, 2, 400, 402, 5, 74,
	38, 2, 401, 403, 7, 12, 2, 2, 402, 401, 3, 2, 2, 2, 402, 403, 3, 2, 2,
	2, 403, 404, 3, 2, 2, 2, 404, 405, 7, 10, 2, 2, 405, 407, 3, 2, 2, 2, 406,
	397, 3, 2, 2, 2, 406, 399, 3, 2, 2, 2, 407, 73, 3, 2, 2, 2, 408, 413, 5,
	76, 39, 2, 409, 410, 7, 12, 2, 2, 410, 412, 5, 76, 39, 2, 411, 409, 3,
	2, 2, 2, 412, 415, 3, 2, 2, 2, 413, 411, 3, 2, 2, 2, 413, 414, 3, 2, 2,
	2, 414, 75, 3, 2, 2, 2, 415, 413, 3, 2, 2, 2, 416, 417, 5, 78, 40, 2, 417,
	418, 7, 15, 2, 2, 418, 419, 5, 88, 45, 2, 419, 436, 3, 2, 2, 2, 420, 421,
	5, 104, 53, 2, 421, 422, 7, 7, 2, 2, 422, 423, 7, 8, 2, 2, 423, 424, 7,
	9, 2, 2, 424, 425, 5, 64, 33, 2, 425, 426, 7, 10, 2, 2, 426, 436, 3, 2,
	2, 2, 427, 428, 5, 106, 54, 2, 428, 429, 7, 7, 2, 2, 429, 430, 5, 80, 41,
	2, 430, 431, 7, 8, 2, 2, 431, 432, 7, 9, 2, 2, 432, 433, 5, 64, 33, 2,
	433, 434, 7, 10, 2, 2, 434, 436, 3, 2, 2, 2, 435, 416, 3, 2, 2, 2, 435,
	420, 3, 2, 2, 2, 435, 427, 3, 2, 2, 2, 436, 77, 3, 2, 2, 2, 437, 441, 5,
	96, 49, 2, 438, 441, 7, 101, 2, 2, 439, 441, 5, 94, 48, 2, 440, 437, 3,
	2, 2, 2, 440, 438, 3, 2, 2, 2, 440, 439, 3, 2, 2, 2, 441, 79, 3, 2, 2,
	2, 442, 443, 7, 100, 2, 2, 443, 81, 3, 2, 2, 2, 444, 446, 7, 7, 2, 2, 445,
	447, 5, 84, 43, 2, 446, 445, 3, 2, 2, 2, 446, 447, 3, 2, 2, 2, 447, 448,
	3, 2, 2, 2, 448, 449, 7, 8, 2, 2, 449, 83, 3, 2, 2, 2, 450, 455, 5, 88,
	45, 2, 451, 452, 7, 12, 2, 2, 452, 454, 5, 88, 45, 2, 453, 451, 3, 2, 2,
	2, 454, 457, 3, 2, 2, 2, 455, 453, 3, 2, 2, 2, 455, 456, 3, 2, 2, 2, 456,
	85, 3, 2, 2, 2, 457, 455, 3, 2, 2, 2, 458, 463, 5, 88, 45, 2, 459, 460,
	7, 12, 2, 2, 460, 462, 5, 88, 45, 2, 461, 459, 3, 2, 2, 2, 462, 465, 3,
	2, 2, 2, 463, 461, 3, 2, 2, 2, 463, 464, 3, 2, 2, 2, 464, 87, 3, 2, 2,
	2, 465, 463, 3, 2, 2, 2, 466, 467, 8, 45, 1, 2, 467, 469, 7, 75, 2, 2,
	468, 470, 7, 100, 2, 2, 469, 468, 3, 2, 2, 2, 469, 470, 3, 2, 2, 2, 470,
	471, 3, 2, 2, 2, 471, 473, 7, 7, 2, 2, 472, 474, 5, 62, 32, 2, 473, 472,
	3, 2, 2, 2, 473, 474, 3, 2, 2, 2, 474, 475, 3, 2, 2, 2, 475, 476, 7, 8,
	2, 2, 476, 477, 7, 9, 2, 2, 477, 478, 5, 64, 33, 2, 478, 479, 7, 10, 2,
	2, 479, 513, 3, 2, 2, 2, 480, 481, 7, 64, 2, 2, 481, 483, 5, 88, 45, 2,
	482, 484, 5, 82, 42, 2, 483, 482, 3, 2, 2, 2, 483, 484, 3, 2, 2, 2, 484,
	513, 3, 2, 2, 2, 485, 486, 7, 81, 2, 2, 486, 513, 5, 88, 45, 32, 487, 488,
	7, 69, 2, 2, 488, 513, 5, 88, 45, 31, 489, 490, 7, 61, 2, 2, 490, 513,
	5, 88, 45, 30, 491, 492, 7, 17, 2, 2, 492, 513, 5, 88, 45, 29, 493, 494,
	7, 18, 2, 2, 494, 513, 5, 88, 45, 28, 495, 496, 7, 19, 2, 2, 496, 513,
	5, 88, 45, 27, 497, 498, 7, 20, 2, 2, 498, 513, 5, 88, 45, 26, 499, 500,
	7, 21, 2, 2, 500, 513, 5, 88, 45, 25, 501, 502, 7, 22, 2, 2, 502, 513,
	5, 88, 45, 24, 503, 513, 7, 76, 2, 2, 504, 513, 7, 100, 2, 2, 505, 513,
	5, 92, 47, 2, 506, 513, 5, 66, 34, 2, 507, 513, 5, 72, 37, 2, 508, 509,
	7, 7, 2, 2, 509, 510, 5, 86, 44, 2, 510, 511, 7, 8, 2, 2, 511, 513, 3,
	2, 2, 2, 512, 466, 3, 2, 2, 2, 512, 480, 3, 2, 2, 2, 512, 485, 3, 2, 2,
	2, 512, 487, 3, 2, 2, 2, 512, 489, 3, 2, 2, 2, 512, 491, 3, 2, 2, 2, 512,
	493, 3, 2, 2, 2, 512, 495, 3, 2, 2, 2, 512, 497, 3, 2, 2, 2, 512, 499,
	3, 2, 2, 2, 512, 501, 3, 2, 2, 2, 512, 503, 3, 2, 2, 2, 512, 504, 3, 2,
	2, 2, 512, 505, 3, 2, 2, 2, 512, 506, 3, 2, 2, 2, 512, 507, 3, 2, 2, 2,
	512, 508, 3, 2, 2, 2, 513, 581, 3, 2, 2, 2, 514, 515, 12, 23, 2, 2, 515,
	516, 9, 2, 2, 2, 516, 580, 5, 88, 45, 24, 517, 518, 12, 22, 2, 2, 518,
	519, 9, 3, 2, 2, 519, 580, 5, 88, 45, 23, 520, 521, 12, 21, 2, 2, 521,
	522, 9, 4, 2, 2, 522, 580, 5, 88, 45, 22, 523, 524, 12, 20, 2, 2, 524,
	525, 9, 5, 2, 2, 525, 580, 5, 88, 45, 21, 526, 527, 12, 19, 2, 2, 527,
	528, 7, 60, 2, 2, 528, 580, 5, 88, 45, 20, 529, 530, 12, 18, 2, 2, 530,
	531, 7, 82, 2, 2, 531, 580, 5, 88, 45, 19, 532, 533, 12, 17, 2, 2, 533,
	534, 9, 6, 2, 2, 534, 580, 5, 88, 45, 18, 535, 536, 12, 16, 2, 2, 536,
	537, 7, 37, 2, 2, 537, 580, 5, 88, 45, 17, 538, 539, 12, 15, 2, 2, 539,
	540, 7, 38, 2, 2, 540, 580, 5, 88, 45, 16, 541, 542, 12, 14, 2, 2, 542,
	543, 7, 39, 2, 2, 543, 580, 5, 88, 45, 15, 544, 545, 12, 13, 2, 2, 545,
	546, 7, 40, 2, 2, 546, 580, 5, 88, 45, 14, 547, 548, 12, 12, 2, 2, 548,
	549, 7, 41, 2, 2, 549, 580, 5, 88, 45, 13, 550, 551, 12, 11, 2, 2, 551,
	552, 7, 14, 2, 2, 552, 553, 5, 88, 45, 2, 553, 554, 7, 15, 2, 2, 554, 555,
	5, 88, 45, 12, 555, 580, 3, 2, 2, 2, 556, 557, 12, 10, 2, 2, 557, 558,
	7, 13, 2, 2, 558, 580, 5, 88, 45, 11, 559, 560, 12, 9, 2, 2, 560, 561,
	5, 90, 46, 2, 561, 562, 5, 88, 45, 10, 562, 580, 3, 2, 2, 2, 563, 564,
	12, 38, 2, 2, 564, 565, 7, 5, 2, 2, 565, 566, 5, 86, 44, 2, 566, 567, 7,
	6, 2, 2, 567, 580, 3, 2, 2, 2, 568, 569, 12, 37, 2, 2, 569, 570, 7, 16,
	2, 2, 570, 580, 5, 96, 49, 2, 571, 572, 12, 36, 2, 2, 572, 580, 5, 82,
	42, 2, 573, 574, 12, 34, 2, 2, 574, 575, 6, 45, 22, 2, 575, 580, 7, 17,
	2, 2, 576, 577, 12, 33, 2, 2, 577, 578, 6, 45, 24, 2, 578, 580, 7, 18,
	2, 2, 579, 514, 3, 2, 2, 2, 579, 517, 3, 2, 2, 2, 579, 520, 3, 2, 2, 2,
	579, 523, 3, 2, 2, 2, 579, 526, 3, 2, 2, 2, 579, 529, 3, 2, 2, 2, 579,
	532, 3, 2, 2, 2, 579, 535, 3, 2, 2, 2, 579, 538, 3, 2, 2, 2, 579, 541,
	3, 2, 2, 2, 579, 544, 3, 2, 2, 2, 579, 547, 3, 2, 2, 2, 579, 550, 3, 2,
	2, 2, 579, 556, 3, 2, 2, 2, 579, 559, 3, 2, 2, 2, 579, 563, 3, 2, 2, 2,
	579, 568, 3, 2, 2, 2, 579, 571, 3, 2, 2, 2, 579, 573, 3, 2, 2, 2, 579,
	576, 3, 2, 2, 2, 580, 583, 3, 2, 2, 2, 581, 579, 3, 2, 2, 2, 581, 582,
	3, 2, 2, 2, 582, 89, 3, 2, 2, 2, 583, 581, 3, 2, 2, 2, 584, 585, 9, 7,
	2, 2, 585, 91, 3, 2, 2, 2, 586, 589, 9, 8, 2, 2, 587, 589, 5, 94, 48, 2,
	588, 586, 3, 2, 2, 2, 588, 587, 3, 2, 2, 2, 589, 93, 3, 2, 2, 2, 590, 591,
	9, 9, 2, 2, 591, 95, 3, 2, 2, 2, 592, 595, 7, 100, 2, 2, 593, 595, 5, 98,
	50, 2, 594, 592, 3, 2, 2, 2, 594, 593, 3, 2, 2, 2, 595, 97, 3, 2, 2, 2,
	596, 600, 5, 100, 51, 2, 597, 600, 5, 102, 52, 2, 598, 600, 9, 10, 2, 2,
	599, 596, 3, 2, 2, 2, 599, 597, 3, 2, 2, 2, 599, 598, 3, 2, 2, 2, 600,
	99, 3, 2, 2, 2, 601, 602, 9, 11, 2, 2, 602, 101, 3, 2, 2, 2, 603, 604,
	9, 12, 2, 2, 604, 103, 3, 2, 2, 2, 605, 606, 6, 53, 25, 2, 606, 607, 7,
	100, 2, 2, 607, 608, 5, 78, 40, 2, 608, 105, 3, 2, 2, 2, 609, 610, 6, 54,
	26, 2, 610, 611, 7, 100, 2, 2, 611, 612, 5, 78, 40, 2, 612, 107, 3, 2,
	2, 2, 613, 618, 7, 11, 2, 2, 614, 618, 7, 2, 2, 3, 615, 618, 6, 55, 27,
	2, 616, 618, 6, 55, 28, 2, 617, 613, 3, 2, 2, 2, 617, 614, 3, 2, 2, 2,
	617, 615, 3, 2, 2, 2, 617, 616, 3, 2, 2, 2, 618, 109, 3, 2, 2, 2, 619,
	620, 7, 2, 2, 3, 620, 111, 3, 2, 2, 2, 55, 113, 120, 124, 141, 145, 152,
	163, 168, 186, 205, 209, 213, 223, 227, 249, 253, 259, 265, 283, 287, 289,
	296, 302, 307, 330, 348, 360, 364, 368, 371, 374, 379, 384, 389, 395, 402,
	406, 413, 435, 440, 446, 455, 463, 469, 473, 483, 512, 579, 581, 588, 594,
	599, 617,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "'['", "']'", "'('", "')'", "'{'", "'}'", "';'", "','", "'='",
	"'?'", "':'", "'.'", "'++'", "'--'", "'+'", "'-'", "'~'", "'!'", "'*'",
	"'/'", "'%'", "'>>'", "'<<'", "'>>>'", "'<'", "'>'", "'<='", "'>='", "'=='",
	"'!='", "'==='", "'!=='", "'&'", "'^'", "'|'", "'&&'", "'||'", "'*='",
	"'/='", "'%='", "'+='", "'-='", "'<<='", "'>>='", "'>>>='", "'&='", "'^='",
	"'|='", "'null'", "", "", "", "", "'break'", "'do'", "'instanceof'", "'typeof'",
	"'case'", "'else'", "'new'", "'var'", "'catch'", "'finally'", "'return'",
	"'void'", "'continue'", "'for'", "'switch'", "'while'", "'debugger'", "'function'",
	"'this'", "'with'", "'default'", "'if'", "'throw'", "'delete'", "'in'",
	"'try'", "'class'", "'enum'", "'extends'", "'super'", "'const'", "'export'",
	"'import'",
}
var symbolicNames = []string{
	"", "RegularExpressionLiteral", "LineTerminator", "OpenBracket", "CloseBracket",
	"OpenParen", "CloseParen", "OpenBrace", "CloseBrace", "SemiColon", "Comma",
	"Assign", "QuestionMark", "Colon", "Dot", "PlusPlus", "MinusMinus", "Plus",
	"Minus", "BitNot", "Not", "Multiply", "Divide", "Modulus", "RightShiftArithmetic",
	"LeftShiftArithmetic", "RightShiftLogical", "LessThan", "MoreThan", "LessThanEquals",
	"GreaterThanEquals", "Equals", "NotEquals", "IdentityEquals", "IdentityNotEquals",
	"BitAnd", "BitXOr", "BitOr", "And", "Or", "MultiplyAssign", "DivideAssign",
	"ModulusAssign", "PlusAssign", "MinusAssign", "LeftShiftArithmeticAssign",
	"RightShiftArithmeticAssign", "RightShiftLogicalAssign", "BitAndAssign",
	"BitXorAssign", "BitOrAssign", "NullLiteral", "BooleanLiteral", "DecimalLiteral",
	"HexIntegerLiteral", "OctalIntegerLiteral", "Break", "Do", "Instanceof",
	"Typeof", "Case", "Else", "New", "Var", "Catch", "Finally", "Return", "Void",
	"Continue", "For", "Switch", "While", "Debugger", "Function", "This", "With",
	"Default", "If", "Throw", "Delete", "In", "Try", "Class", "Enum", "Extends",
	"Super", "Const", "Export", "Import", "Implements", "Let", "Private", "Public",
	"Interface", "Package", "Protected", "Static", "Yield", "Identifier", "StringLiteral",
	"WhiteSpaces", "MultiLineComment", "SingleLineComment", "UnexpectedCharacter",
}

var ruleNames = []string{
	"program", "sourceElements", "sourceElement", "statement", "block", "statementList",
	"variableStatement", "variableDeclarationList", "variableDeclaration",
	"initialiser", "voidStatement", "expressionStatement", "ifStatement", "iterationStatement",
	"continueStatement", "breakStatement", "returnStatement", "withStatement",
	"switchStatement", "caseBlock", "caseClauses", "caseClause", "defaultClause",
	"labelledStatement", "throwStatement", "tryStatement", "catchProduction",
	"finallyProduction", "debuggerStatement", "functionDeclaration", "formalParameterList",
	"functionBody", "arrayLiteral", "elementList", "elision", "objectLiteral",
	"propertyNameAndValueList", "propertyAssignment", "propertyName", "propertySetParameterList",
	"arguments", "argumentList", "expressionSequence", "singleExpression",
	"assignmentOperator", "literal", "numericLiteral", "identifierName", "reservedWord",
	"keyword", "futureReservedWord", "getter", "setter", "eos", "eof",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ECMAScriptParser struct {
	*antlr.BaseParser
}

func NewECMAScriptParser(input antlr.TokenStream) *ECMAScriptParser {
	this := new(ECMAScriptParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "ECMAScript.g4"

	return this
}

// here returns `true` iff on the current index of the parser's
// token stream a token of the given `type` exists on the
// `HIDDEN` channel.
//
// Args:
//  type (int): the type of the token on the `HIDDEN` channel
//      to check.
//
//  Returns:
//      `True` iff on the current index of the parser's
//      token stream a token of the given `type` exists on the
//      `HIDDEN` channel.
func (p *ECMAScriptParser) here(tokenType int) bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	// Check if the token resides on the HIDDEN channel and if it is of the
	// provided tokenType.
	return (ahead.GetChannel() == antlr.LexerHidden) && (ahead.GetTokenType() == tokenType)
}

// lineTerminatorAhead returns `true` iff on the current index of the parser's
// token stream a token exists on the `HIDDEN` channel which
// either is a line terminator, or is a multi line comment that
// contains a line terminator.
func (p *ECMAScriptParser) lineTerminatorAhead() bool {
	// Get the token ahead of the current index.
	possibleIndexEosToken := p.GetCurrentToken().GetTokenIndex() - 1
	ahead := p.GetTokenStream().Get(possibleIndexEosToken)

	if ahead.GetChannel() != antlr.LexerHidden {
		// We're only interested in tokens on the HIDDEN channel.
		return false
	}

	if ahead.GetTokenType() == ECMAScriptParserLineTerminator {
		// There is definitely a line terminator ahead.
		return true
	}

	if ahead.GetTokenType() == ECMAScriptParserWhiteSpaces {
		// Get the token ahead of the current whitespaces.
		possibleIndexEosToken = p.GetCurrentToken().GetTokenIndex() - 2
		ahead = p.GetTokenStream().Get(possibleIndexEosToken)
	}

	// Get the token's text and type.
	text := ahead.GetText()
	tokenType := ahead.GetTokenType()

	// Check if the token is, or contains a line terminator.
	if tokenType == ECMAScriptParserMultiLineComment && strings.ContainsAny(text, "\r\n") {
		return true
	}

	return tokenType == ECMAScriptParserLineTerminator
}

// ECMAScriptParser tokens.
const (
	ECMAScriptParserEOF                        = antlr.TokenEOF
	ECMAScriptParserRegularExpressionLiteral   = 1
	ECMAScriptParserLineTerminator             = 2
	ECMAScriptParserOpenBracket                = 3
	ECMAScriptParserCloseBracket               = 4
	ECMAScriptParserOpenParen                  = 5
	ECMAScriptParserCloseParen                 = 6
	ECMAScriptParserOpenBrace                  = 7
	ECMAScriptParserCloseBrace                 = 8
	ECMAScriptParserSemiColon                  = 9
	ECMAScriptParserComma                      = 10
	ECMAScriptParserAssign                     = 11
	ECMAScriptParserQuestionMark               = 12
	ECMAScriptParserColon                      = 13
	ECMAScriptParserDot                        = 14
	ECMAScriptParserPlusPlus                   = 15
	ECMAScriptParserMinusMinus                 = 16
	ECMAScriptParserPlus                       = 17
	ECMAScriptParserMinus                      = 18
	ECMAScriptParserBitNot                     = 19
	ECMAScriptParserNot                        = 20
	ECMAScriptParserMultiply                   = 21
	ECMAScriptParserDivide                     = 22
	ECMAScriptParserModulus                    = 23
	ECMAScriptParserRightShiftArithmetic       = 24
	ECMAScriptParserLeftShiftArithmetic        = 25
	ECMAScriptParserRightShiftLogical          = 26
	ECMAScriptParserLessThan                   = 27
	ECMAScriptParserMoreThan                   = 28
	ECMAScriptParserLessThanEquals             = 29
	ECMAScriptParserGreaterThanEquals          = 30
	ECMAScriptParserEquals                     = 31
	ECMAScriptParserNotEquals                  = 32
	ECMAScriptParserIdentityEquals             = 33
	ECMAScriptParserIdentityNotEquals          = 34
	ECMAScriptParserBitAnd                     = 35
	ECMAScriptParserBitXOr                     = 36
	ECMAScriptParserBitOr                      = 37
	ECMAScriptParserAnd                        = 38
	ECMAScriptParserOr                         = 39
	ECMAScriptParserMultiplyAssign             = 40
	ECMAScriptParserDivideAssign               = 41
	ECMAScriptParserModulusAssign              = 42
	ECMAScriptParserPlusAssign                 = 43
	ECMAScriptParserMinusAssign                = 44
	ECMAScriptParserLeftShiftArithmeticAssign  = 45
	ECMAScriptParserRightShiftArithmeticAssign = 46
	ECMAScriptParserRightShiftLogicalAssign    = 47
	ECMAScriptParserBitAndAssign               = 48
	ECMAScriptParserBitXorAssign               = 49
	ECMAScriptParserBitOrAssign                = 50
	ECMAScriptParserNullLiteral                = 51
	ECMAScriptParserBooleanLiteral             = 52
	ECMAScriptParserDecimalLiteral             = 53
	ECMAScriptParserHexIntegerLiteral          = 54
	ECMAScriptParserOctalIntegerLiteral        = 55
	ECMAScriptParserBreak                      = 56
	ECMAScriptParserDo                         = 57
	ECMAScriptParserInstanceof                 = 58
	ECMAScriptParserTypeof                     = 59
	ECMAScriptParserCase                       = 60
	ECMAScriptParserElse                       = 61
	ECMAScriptParserNew                        = 62
	ECMAScriptParserVar                        = 63
	ECMAScriptParserCatch                      = 64
	ECMAScriptParserFinally                    = 65
	ECMAScriptParserReturn                     = 66
	ECMAScriptParserVoid                       = 67
	ECMAScriptParserContinue                   = 68
	ECMAScriptParserFor                        = 69
	ECMAScriptParserSwitch                     = 70
	ECMAScriptParserWhile                      = 71
	ECMAScriptParserDebugger                   = 72
	ECMAScriptParserFunction                   = 73
	ECMAScriptParserThis                       = 74
	ECMAScriptParserWith                       = 75
	ECMAScriptParserDefault                    = 76
	ECMAScriptParserIf                         = 77
	ECMAScriptParserThrow                      = 78
	ECMAScriptParserDelete                     = 79
	ECMAScriptParserIn                         = 80
	ECMAScriptParserTry                        = 81
	ECMAScriptParserClass                      = 82
	ECMAScriptParserEnum                       = 83
	ECMAScriptParserExtends                    = 84
	ECMAScriptParserSuper                      = 85
	ECMAScriptParserConst                      = 86
	ECMAScriptParserExport                     = 87
	ECMAScriptParserImport                     = 88
	ECMAScriptParserImplements                 = 89
	ECMAScriptParserLet                        = 90
	ECMAScriptParserPrivate                    = 91
	ECMAScriptParserPublic                     = 92
	ECMAScriptParserInterface                  = 93
	ECMAScriptParserPackage                    = 94
	ECMAScriptParserProtected                  = 95
	ECMAScriptParserStatic                     = 96
	ECMAScriptParserYield                      = 97
	ECMAScriptParserIdentifier                 = 98
	ECMAScriptParserStringLiteral              = 99
	ECMAScriptParserWhiteSpaces                = 100
	ECMAScriptParserMultiLineComment           = 101
	ECMAScriptParserSingleLineComment          = 102
	ECMAScriptParserUnexpectedCharacter        = 103
)

// ECMAScriptParser rules.
const (
	ECMAScriptParserRULE_program                  = 0
	ECMAScriptParserRULE_sourceElements           = 1
	ECMAScriptParserRULE_sourceElement            = 2
	ECMAScriptParserRULE_statement                = 3
	ECMAScriptParserRULE_block                    = 4
	ECMAScriptParserRULE_statementList            = 5
	ECMAScriptParserRULE_variableStatement        = 6
	ECMAScriptParserRULE_variableDeclarationList  = 7
	ECMAScriptParserRULE_variableDeclaration      = 8
	ECMAScriptParserRULE_initialiser              = 9
	ECMAScriptParserRULE_voidStatement            = 10
	ECMAScriptParserRULE_expressionStatement      = 11
	ECMAScriptParserRULE_ifStatement              = 12
	ECMAScriptParserRULE_iterationStatement       = 13
	ECMAScriptParserRULE_continueStatement        = 14
	ECMAScriptParserRULE_breakStatement           = 15
	ECMAScriptParserRULE_returnStatement          = 16
	ECMAScriptParserRULE_withStatement            = 17
	ECMAScriptParserRULE_switchStatement          = 18
	ECMAScriptParserRULE_caseBlock                = 19
	ECMAScriptParserRULE_caseClauses              = 20
	ECMAScriptParserRULE_caseClause               = 21
	ECMAScriptParserRULE_defaultClause            = 22
	ECMAScriptParserRULE_labelledStatement        = 23
	ECMAScriptParserRULE_throwStatement           = 24
	ECMAScriptParserRULE_tryStatement             = 25
	ECMAScriptParserRULE_catchProduction          = 26
	ECMAScriptParserRULE_finallyProduction        = 27
	ECMAScriptParserRULE_debuggerStatement        = 28
	ECMAScriptParserRULE_functionDeclaration      = 29
	ECMAScriptParserRULE_formalParameterList      = 30
	ECMAScriptParserRULE_functionBody             = 31
	ECMAScriptParserRULE_arrayLiteral             = 32
	ECMAScriptParserRULE_elementList              = 33
	ECMAScriptParserRULE_elision                  = 34
	ECMAScriptParserRULE_objectLiteral            = 35
	ECMAScriptParserRULE_propertyNameAndValueList = 36
	ECMAScriptParserRULE_propertyAssignment       = 37
	ECMAScriptParserRULE_propertyName             = 38
	ECMAScriptParserRULE_propertySetParameterList = 39
	ECMAScriptParserRULE_arguments                = 40
	ECMAScriptParserRULE_argumentList             = 41
	ECMAScriptParserRULE_expressionSequence       = 42
	ECMAScriptParserRULE_singleExpression         = 43
	ECMAScriptParserRULE_assignmentOperator       = 44
	ECMAScriptParserRULE_literal                  = 45
	ECMAScriptParserRULE_numericLiteral           = 46
	ECMAScriptParserRULE_identifierName           = 47
	ECMAScriptParserRULE_reservedWord             = 48
	ECMAScriptParserRULE_keyword                  = 49
	ECMAScriptParserRULE_futureReservedWord       = 50
	ECMAScriptParserRULE_getter                   = 51
	ECMAScriptParserRULE_setter                   = 52
	ECMAScriptParserRULE_eos                      = 53
	ECMAScriptParserRULE_eof                      = 54
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
	p.RuleIndex = ECMAScriptParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserEOF, 0)
}

func (s *ProgramContext) SourceElements() ISourceElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceElementsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *ECMAScriptParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ECMAScriptParserRULE_program)

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
	p.SetState(111)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(110)
			p.SourceElements()
		}

	}
	{
		p.SetState(113)
		p.Match(ECMAScriptParserEOF)
	}

	return localctx
}

// ISourceElementsContext is an interface to support dynamic dispatch.
type ISourceElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceElementsContext differentiates from other interfaces.
	IsSourceElementsContext()
}

type SourceElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceElementsContext() *SourceElementsContext {
	var p = new(SourceElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_sourceElements
	return p
}

func (*SourceElementsContext) IsSourceElementsContext() {}

func NewSourceElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceElementsContext {
	var p = new(SourceElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_sourceElements

	return p
}

func (s *SourceElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceElementsContext) AllSourceElement() []ISourceElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISourceElementContext)(nil)).Elem())
	var tst = make([]ISourceElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISourceElementContext)
		}
	}

	return tst
}

func (s *SourceElementsContext) SourceElement(i int) ISourceElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISourceElementContext)
}

func (s *SourceElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterSourceElements(s)
	}
}

func (s *SourceElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitSourceElements(s)
	}
}

func (p *ECMAScriptParser) SourceElements() (localctx ISourceElementsContext) {
	localctx = NewSourceElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ECMAScriptParserRULE_sourceElements)

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
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(115)
				p.SourceElement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// ISourceElementContext is an interface to support dynamic dispatch.
type ISourceElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceElementContext differentiates from other interfaces.
	IsSourceElementContext()
}

type SourceElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceElementContext() *SourceElementContext {
	var p = new(SourceElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_sourceElement
	return p
}

func (*SourceElementContext) IsSourceElementContext() {}

func NewSourceElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceElementContext {
	var p = new(SourceElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_sourceElement

	return p
}

func (s *SourceElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceElementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SourceElementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *SourceElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterSourceElement(s)
	}
}

func (s *SourceElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitSourceElement(s)
	}
}

func (p *ECMAScriptParser) SourceElement() (localctx ISourceElementContext) {
	localctx = NewSourceElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ECMAScriptParserRULE_sourceElement)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(120)
			p.Statement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(121)
			p.FunctionDeclaration()
		}

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
	p.RuleIndex = ECMAScriptParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) VariableStatement() IVariableStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableStatementContext)
}

func (s *StatementContext) VoidStatement() IVoidStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVoidStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVoidStatementContext)
}

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) IterationStatement() IIterationStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIterationStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIterationStatementContext)
}

func (s *StatementContext) ContinueStatement() IContinueStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinueStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinueStatementContext)
}

func (s *StatementContext) BreakStatement() IBreakStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) WithStatement() IWithStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWithStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWithStatementContext)
}

func (s *StatementContext) LabelledStatement() ILabelledStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelledStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelledStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementContext) ThrowStatement() IThrowStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThrowStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThrowStatementContext)
}

func (s *StatementContext) TryStatement() ITryStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITryStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITryStatementContext)
}

func (s *StatementContext) DebuggerStatement() IDebuggerStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDebuggerStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDebuggerStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *ECMAScriptParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ECMAScriptParserRULE_statement)

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

	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Block()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.VariableStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(126)
			p.VoidStatement()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(127)
			p.ExpressionStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(128)
			p.IfStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(129)
			p.IterationStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(130)
			p.ContinueStatement()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(131)
			p.BreakStatement()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(132)
			p.ReturnStatement()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(133)
			p.WithStatement()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(134)
			p.LabelledStatement()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(135)
			p.SwitchStatement()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(136)
			p.ThrowStatement()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(137)
			p.TryStatement()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(138)
			p.DebuggerStatement()
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
	p.RuleIndex = ECMAScriptParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *ECMAScriptParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ECMAScriptParserRULE_block)

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
		p.SetState(141)
		p.Match(ECMAScriptParserOpenBrace)
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(142)
			p.StatementList()
		}

	}
	{
		p.SetState(145)
		p.Match(ECMAScriptParserCloseBrace)
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (p *ECMAScriptParser) StatementList() (localctx IStatementListContext) {
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ECMAScriptParserRULE_statementList)

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
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(147)
				p.Statement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IVariableStatementContext is an interface to support dynamic dispatch.
type IVariableStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableStatementContext differentiates from other interfaces.
	IsVariableStatementContext()
}

type VariableStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableStatementContext() *VariableStatementContext {
	var p = new(VariableStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_variableStatement
	return p
}

func (*VariableStatementContext) IsVariableStatementContext() {}

func NewVariableStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableStatementContext {
	var p = new(VariableStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_variableStatement

	return p
}

func (s *VariableStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableStatementContext) Var() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserVar, 0)
}

func (s *VariableStatementContext) VariableDeclarationList() IVariableDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationListContext)
}

func (s *VariableStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *VariableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterVariableStatement(s)
	}
}

func (s *VariableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitVariableStatement(s)
	}
}

func (p *ECMAScriptParser) VariableStatement() (localctx IVariableStatementContext) {
	localctx = NewVariableStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ECMAScriptParserRULE_variableStatement)

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
		p.SetState(152)
		p.Match(ECMAScriptParserVar)
	}
	{
		p.SetState(153)
		p.VariableDeclarationList()
	}
	{
		p.SetState(154)
		p.Eos()
	}

	return localctx
}

// IVariableDeclarationListContext is an interface to support dynamic dispatch.
type IVariableDeclarationListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationListContext differentiates from other interfaces.
	IsVariableDeclarationListContext()
}

type VariableDeclarationListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationListContext() *VariableDeclarationListContext {
	var p = new(VariableDeclarationListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_variableDeclarationList
	return p
}

func (*VariableDeclarationListContext) IsVariableDeclarationListContext() {}

func NewVariableDeclarationListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationListContext {
	var p = new(VariableDeclarationListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_variableDeclarationList

	return p
}

func (s *VariableDeclarationListContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationListContext) AllVariableDeclaration() []IVariableDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem())
	var tst = make([]IVariableDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDeclarationContext)
		}
	}

	return tst
}

func (s *VariableDeclarationListContext) VariableDeclaration(i int) IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *VariableDeclarationListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterVariableDeclarationList(s)
	}
}

func (s *VariableDeclarationListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitVariableDeclarationList(s)
	}
}

func (p *ECMAScriptParser) VariableDeclarationList() (localctx IVariableDeclarationListContext) {
	localctx = NewVariableDeclarationListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ECMAScriptParserRULE_variableDeclarationList)

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
		p.SetState(156)
		p.VariableDeclaration()
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(157)
				p.Match(ECMAScriptParserComma)
			}
			{
				p.SetState(158)
				p.VariableDeclaration()
			}

		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

// IVariableDeclarationContext is an interface to support dynamic dispatch.
type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDeclarationContext differentiates from other interfaces.
	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *VariableDeclarationContext) Initialiser() IInitialiserContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitialiserContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitialiserContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *ECMAScriptParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ECMAScriptParserRULE_variableDeclaration)

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
		p.SetState(164)
		p.Match(ECMAScriptParserIdentifier)
	}
	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(165)
			p.Initialiser()
		}

	}

	return localctx
}

// IInitialiserContext is an interface to support dynamic dispatch.
type IInitialiserContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitialiserContext differentiates from other interfaces.
	IsInitialiserContext()
}

type InitialiserContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitialiserContext() *InitialiserContext {
	var p = new(InitialiserContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_initialiser
	return p
}

func (*InitialiserContext) IsInitialiserContext() {}

func NewInitialiserContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitialiserContext {
	var p = new(InitialiserContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_initialiser

	return p
}

func (s *InitialiserContext) GetParser() antlr.Parser { return s.parser }

func (s *InitialiserContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *InitialiserContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitialiserContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitialiserContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterInitialiser(s)
	}
}

func (s *InitialiserContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitInitialiser(s)
	}
}

func (p *ECMAScriptParser) Initialiser() (localctx IInitialiserContext) {
	localctx = NewInitialiserContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ECMAScriptParserRULE_initialiser)

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
		p.SetState(168)
		p.Match(ECMAScriptParserAssign)
	}
	{
		p.SetState(169)
		p.singleExpression(0)
	}

	return localctx
}

// IVoidStatementContext is an interface to support dynamic dispatch.
type IVoidStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVoidStatementContext differentiates from other interfaces.
	IsVoidStatementContext()
}

type VoidStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVoidStatementContext() *VoidStatementContext {
	var p = new(VoidStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_voidStatement
	return p
}

func (*VoidStatementContext) IsVoidStatementContext() {}

func NewVoidStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VoidStatementContext {
	var p = new(VoidStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_voidStatement

	return p
}

func (s *VoidStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *VoidStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserSemiColon, 0)
}

func (s *VoidStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoidStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VoidStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterVoidStatement(s)
	}
}

func (s *VoidStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitVoidStatement(s)
	}
}

func (p *ECMAScriptParser) VoidStatement() (localctx IVoidStatementContext) {
	localctx = NewVoidStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ECMAScriptParserRULE_voidStatement)

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
		p.Match(ECMAScriptParserSemiColon)
	}

	return localctx
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_expressionStatement
	return p
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ExpressionStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (p *ECMAScriptParser) ExpressionStatement() (localctx IExpressionStatementContext) {
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ECMAScriptParserRULE_expressionStatement)

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

	if !((p.GetInputStream().LA(1) != ECMAScriptParserOpenBrace) && (p.GetInputStream().LA(1) != ECMAScriptParserFunction)) {
		panic(antlr.NewFailedPredicateException(p, "(p.GetInputStream().LA(1) != ECMAScriptParserOpenBrace) && (p.GetInputStream().LA(1) != ECMAScriptParserFunction)", ""))
	}
	{
		p.SetState(174)
		p.ExpressionSequence()
	}
	{
		p.SetState(175)
		p.Eos()
	}

	return localctx
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_ifStatement
	return p
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) If() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIf, 0)
}

func (s *IfStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *IfStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserElse, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *ECMAScriptParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ECMAScriptParserRULE_ifStatement)

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
		p.SetState(177)
		p.Match(ECMAScriptParserIf)
	}
	{
		p.SetState(178)
		p.Match(ECMAScriptParserOpenParen)
	}
	{
		p.SetState(179)
		p.ExpressionSequence()
	}
	{
		p.SetState(180)
		p.Match(ECMAScriptParserCloseParen)
	}
	{
		p.SetState(181)
		p.Statement()
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(182)
			p.Match(ECMAScriptParserElse)
		}
		{
			p.SetState(183)
			p.Statement()
		}

	}

	return localctx
}

// IIterationStatementContext is an interface to support dynamic dispatch.
type IIterationStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIterationStatementContext differentiates from other interfaces.
	IsIterationStatementContext()
}

type IterationStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIterationStatementContext() *IterationStatementContext {
	var p = new(IterationStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_iterationStatement
	return p
}

func (*IterationStatementContext) IsIterationStatementContext() {}

func NewIterationStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IterationStatementContext {
	var p = new(IterationStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_iterationStatement

	return p
}

func (s *IterationStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IterationStatementContext) CopyFrom(ctx *IterationStatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *IterationStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IterationStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DoStatementContext struct {
	*IterationStatementContext
}

func NewDoStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoStatementContext {
	var p = new(DoStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *DoStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoStatementContext) Do() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDo, 0)
}

func (s *DoStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DoStatementContext) While() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserWhile, 0)
}

func (s *DoStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *DoStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *DoStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterDoStatement(s)
	}
}

func (s *DoStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitDoStatement(s)
	}
}

type ForVarStatementContext struct {
	*IterationStatementContext
}

func NewForVarStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForVarStatementContext {
	var p = new(ForVarStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *ForVarStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForVarStatementContext) For() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFor, 0)
}

func (s *ForVarStatementContext) Var() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserVar, 0)
}

func (s *ForVarStatementContext) VariableDeclarationList() IVariableDeclarationListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationListContext)
}

func (s *ForVarStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForVarStatementContext) AllExpressionSequence() []IExpressionSequenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem())
	var tst = make([]IExpressionSequenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionSequenceContext)
		}
	}

	return tst
}

func (s *ForVarStatementContext) ExpressionSequence(i int) IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ForVarStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterForVarStatement(s)
	}
}

func (s *ForVarStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitForVarStatement(s)
	}
}

type ForVarInStatementContext struct {
	*IterationStatementContext
}

func NewForVarInStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForVarInStatementContext {
	var p = new(ForVarInStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *ForVarInStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForVarInStatementContext) For() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFor, 0)
}

func (s *ForVarInStatementContext) Var() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserVar, 0)
}

func (s *ForVarInStatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *ForVarInStatementContext) In() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIn, 0)
}

func (s *ForVarInStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ForVarInStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForVarInStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterForVarInStatement(s)
	}
}

func (s *ForVarInStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitForVarInStatement(s)
	}
}

type WhileStatementContext struct {
	*IterationStatementContext
}

func NewWhileStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserWhile, 0)
}

func (s *WhileStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *WhileStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

type ForStatementContext struct {
	*IterationStatementContext
}

func NewForStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForStatementContext {
	var p = new(ForStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) For() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFor, 0)
}

func (s *ForStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForStatementContext) AllExpressionSequence() []IExpressionSequenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem())
	var tst = make([]IExpressionSequenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionSequenceContext)
		}
	}

	return tst
}

func (s *ForStatementContext) ExpressionSequence(i int) IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitForStatement(s)
	}
}

type ForInStatementContext struct {
	*IterationStatementContext
}

func NewForInStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForInStatementContext {
	var p = new(ForInStatementContext)

	p.IterationStatementContext = NewEmptyIterationStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*IterationStatementContext))

	return p
}

func (s *ForInStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForInStatementContext) For() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFor, 0)
}

func (s *ForInStatementContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ForInStatementContext) In() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIn, 0)
}

func (s *ForInStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ForInStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForInStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterForInStatement(s)
	}
}

func (s *ForInStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitForInStatement(s)
	}
}

func (p *ECMAScriptParser) IterationStatement() (localctx IIterationStatementContext) {
	localctx = NewIterationStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ECMAScriptParserRULE_iterationStatement)
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

	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDoStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(186)
			p.Match(ECMAScriptParserDo)
		}
		{
			p.SetState(187)
			p.Statement()
		}
		{
			p.SetState(188)
			p.Match(ECMAScriptParserWhile)
		}
		{
			p.SetState(189)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(190)
			p.ExpressionSequence()
		}
		{
			p.SetState(191)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(192)
			p.Eos()
		}

	case 2:
		localctx = NewWhileStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.Match(ECMAScriptParserWhile)
		}
		{
			p.SetState(195)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(196)
			p.ExpressionSequence()
		}
		{
			p.SetState(197)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(198)
			p.Statement()
		}

	case 3:
		localctx = NewForStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(200)
			p.Match(ECMAScriptParserFor)
		}
		{
			p.SetState(201)
			p.Match(ECMAScriptParserOpenParen)
		}
		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRegularExpressionLiteral)|(1<<ECMAScriptParserOpenBracket)|(1<<ECMAScriptParserOpenParen)|(1<<ECMAScriptParserOpenBrace)|(1<<ECMAScriptParserPlusPlus)|(1<<ECMAScriptParserMinusMinus)|(1<<ECMAScriptParserPlus)|(1<<ECMAScriptParserMinus)|(1<<ECMAScriptParserBitNot)|(1<<ECMAScriptParserNot))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(ECMAScriptParserNullLiteral-51))|(1<<(ECMAScriptParserBooleanLiteral-51))|(1<<(ECMAScriptParserDecimalLiteral-51))|(1<<(ECMAScriptParserHexIntegerLiteral-51))|(1<<(ECMAScriptParserOctalIntegerLiteral-51))|(1<<(ECMAScriptParserTypeof-51))|(1<<(ECMAScriptParserNew-51))|(1<<(ECMAScriptParserVoid-51))|(1<<(ECMAScriptParserFunction-51))|(1<<(ECMAScriptParserThis-51))|(1<<(ECMAScriptParserDelete-51)))) != 0) || _la == ECMAScriptParserIdentifier || _la == ECMAScriptParserStringLiteral {
			{
				p.SetState(202)
				p.ExpressionSequence()
			}

		}
		{
			p.SetState(205)
			p.Match(ECMAScriptParserSemiColon)
		}
		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRegularExpressionLiteral)|(1<<ECMAScriptParserOpenBracket)|(1<<ECMAScriptParserOpenParen)|(1<<ECMAScriptParserOpenBrace)|(1<<ECMAScriptParserPlusPlus)|(1<<ECMAScriptParserMinusMinus)|(1<<ECMAScriptParserPlus)|(1<<ECMAScriptParserMinus)|(1<<ECMAScriptParserBitNot)|(1<<ECMAScriptParserNot))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(ECMAScriptParserNullLiteral-51))|(1<<(ECMAScriptParserBooleanLiteral-51))|(1<<(ECMAScriptParserDecimalLiteral-51))|(1<<(ECMAScriptParserHexIntegerLiteral-51))|(1<<(ECMAScriptParserOctalIntegerLiteral-51))|(1<<(ECMAScriptParserTypeof-51))|(1<<(ECMAScriptParserNew-51))|(1<<(ECMAScriptParserVoid-51))|(1<<(ECMAScriptParserFunction-51))|(1<<(ECMAScriptParserThis-51))|(1<<(ECMAScriptParserDelete-51)))) != 0) || _la == ECMAScriptParserIdentifier || _la == ECMAScriptParserStringLiteral {
			{
				p.SetState(206)
				p.ExpressionSequence()
			}

		}
		{
			p.SetState(209)
			p.Match(ECMAScriptParserSemiColon)
		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRegularExpressionLiteral)|(1<<ECMAScriptParserOpenBracket)|(1<<ECMAScriptParserOpenParen)|(1<<ECMAScriptParserOpenBrace)|(1<<ECMAScriptParserPlusPlus)|(1<<ECMAScriptParserMinusMinus)|(1<<ECMAScriptParserPlus)|(1<<ECMAScriptParserMinus)|(1<<ECMAScriptParserBitNot)|(1<<ECMAScriptParserNot))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(ECMAScriptParserNullLiteral-51))|(1<<(ECMAScriptParserBooleanLiteral-51))|(1<<(ECMAScriptParserDecimalLiteral-51))|(1<<(ECMAScriptParserHexIntegerLiteral-51))|(1<<(ECMAScriptParserOctalIntegerLiteral-51))|(1<<(ECMAScriptParserTypeof-51))|(1<<(ECMAScriptParserNew-51))|(1<<(ECMAScriptParserVoid-51))|(1<<(ECMAScriptParserFunction-51))|(1<<(ECMAScriptParserThis-51))|(1<<(ECMAScriptParserDelete-51)))) != 0) || _la == ECMAScriptParserIdentifier || _la == ECMAScriptParserStringLiteral {
			{
				p.SetState(210)
				p.ExpressionSequence()
			}

		}
		{
			p.SetState(213)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(214)
			p.Statement()
		}

	case 4:
		localctx = NewForVarStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(215)
			p.Match(ECMAScriptParserFor)
		}
		{
			p.SetState(216)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(217)
			p.Match(ECMAScriptParserVar)
		}
		{
			p.SetState(218)
			p.VariableDeclarationList()
		}
		{
			p.SetState(219)
			p.Match(ECMAScriptParserSemiColon)
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRegularExpressionLiteral)|(1<<ECMAScriptParserOpenBracket)|(1<<ECMAScriptParserOpenParen)|(1<<ECMAScriptParserOpenBrace)|(1<<ECMAScriptParserPlusPlus)|(1<<ECMAScriptParserMinusMinus)|(1<<ECMAScriptParserPlus)|(1<<ECMAScriptParserMinus)|(1<<ECMAScriptParserBitNot)|(1<<ECMAScriptParserNot))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(ECMAScriptParserNullLiteral-51))|(1<<(ECMAScriptParserBooleanLiteral-51))|(1<<(ECMAScriptParserDecimalLiteral-51))|(1<<(ECMAScriptParserHexIntegerLiteral-51))|(1<<(ECMAScriptParserOctalIntegerLiteral-51))|(1<<(ECMAScriptParserTypeof-51))|(1<<(ECMAScriptParserNew-51))|(1<<(ECMAScriptParserVoid-51))|(1<<(ECMAScriptParserFunction-51))|(1<<(ECMAScriptParserThis-51))|(1<<(ECMAScriptParserDelete-51)))) != 0) || _la == ECMAScriptParserIdentifier || _la == ECMAScriptParserStringLiteral {
			{
				p.SetState(220)
				p.ExpressionSequence()
			}

		}
		{
			p.SetState(223)
			p.Match(ECMAScriptParserSemiColon)
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRegularExpressionLiteral)|(1<<ECMAScriptParserOpenBracket)|(1<<ECMAScriptParserOpenParen)|(1<<ECMAScriptParserOpenBrace)|(1<<ECMAScriptParserPlusPlus)|(1<<ECMAScriptParserMinusMinus)|(1<<ECMAScriptParserPlus)|(1<<ECMAScriptParserMinus)|(1<<ECMAScriptParserBitNot)|(1<<ECMAScriptParserNot))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(ECMAScriptParserNullLiteral-51))|(1<<(ECMAScriptParserBooleanLiteral-51))|(1<<(ECMAScriptParserDecimalLiteral-51))|(1<<(ECMAScriptParserHexIntegerLiteral-51))|(1<<(ECMAScriptParserOctalIntegerLiteral-51))|(1<<(ECMAScriptParserTypeof-51))|(1<<(ECMAScriptParserNew-51))|(1<<(ECMAScriptParserVoid-51))|(1<<(ECMAScriptParserFunction-51))|(1<<(ECMAScriptParserThis-51))|(1<<(ECMAScriptParserDelete-51)))) != 0) || _la == ECMAScriptParserIdentifier || _la == ECMAScriptParserStringLiteral {
			{
				p.SetState(224)
				p.ExpressionSequence()
			}

		}
		{
			p.SetState(227)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(228)
			p.Statement()
		}

	case 5:
		localctx = NewForInStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(230)
			p.Match(ECMAScriptParserFor)
		}
		{
			p.SetState(231)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(232)
			p.singleExpression(0)
		}
		{
			p.SetState(233)
			p.Match(ECMAScriptParserIn)
		}
		{
			p.SetState(234)
			p.ExpressionSequence()
		}
		{
			p.SetState(235)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(236)
			p.Statement()
		}

	case 6:
		localctx = NewForVarInStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(238)
			p.Match(ECMAScriptParserFor)
		}
		{
			p.SetState(239)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(240)
			p.Match(ECMAScriptParserVar)
		}
		{
			p.SetState(241)
			p.VariableDeclaration()
		}
		{
			p.SetState(242)
			p.Match(ECMAScriptParserIn)
		}
		{
			p.SetState(243)
			p.ExpressionSequence()
		}
		{
			p.SetState(244)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(245)
			p.Statement()
		}

	}

	return localctx
}

// IContinueStatementContext is an interface to support dynamic dispatch.
type IContinueStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinueStatementContext differentiates from other interfaces.
	IsContinueStatementContext()
}

type ContinueStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStatementContext() *ContinueStatementContext {
	var p = new(ContinueStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_continueStatement
	return p
}

func (*ContinueStatementContext) IsContinueStatementContext() {}

func NewContinueStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_continueStatement

	return p
}

func (s *ContinueStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserContinue, 0)
}

func (s *ContinueStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ContinueStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (p *ECMAScriptParser) ContinueStatement() (localctx IContinueStatementContext) {
	localctx = NewContinueStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ECMAScriptParserRULE_continueStatement)

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
		p.Match(ECMAScriptParserContinue)
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(250)
			p.Match(ECMAScriptParserIdentifier)
		}

	}
	{
		p.SetState(253)
		p.Eos()
	}

	return localctx
}

// IBreakStatementContext is an interface to support dynamic dispatch.
type IBreakStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreakStatementContext differentiates from other interfaces.
	IsBreakStatementContext()
}

type BreakStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStatementContext() *BreakStatementContext {
	var p = new(BreakStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_breakStatement
	return p
}

func (*BreakStatementContext) IsBreakStatementContext() {}

func NewBreakStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_breakStatement

	return p
}

func (s *BreakStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStatementContext) Break() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBreak, 0)
}

func (s *BreakStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *BreakStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (p *ECMAScriptParser) BreakStatement() (localctx IBreakStatementContext) {
	localctx = NewBreakStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ECMAScriptParserRULE_breakStatement)

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
		p.Match(ECMAScriptParserBreak)
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(256)
			p.Match(ECMAScriptParserIdentifier)
		}

	}
	{
		p.SetState(259)
		p.Eos()
	}

	return localctx
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_returnStatement
	return p
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserReturn, 0)
}

func (s *ReturnStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ReturnStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *ECMAScriptParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ECMAScriptParserRULE_returnStatement)

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
		p.SetState(261)
		p.Match(ECMAScriptParserReturn)
	}
	p.SetState(263)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(262)
			p.ExpressionSequence()
		}

	}
	{
		p.SetState(265)
		p.Eos()
	}

	return localctx
}

// IWithStatementContext is an interface to support dynamic dispatch.
type IWithStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWithStatementContext differentiates from other interfaces.
	IsWithStatementContext()
}

type WithStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWithStatementContext() *WithStatementContext {
	var p = new(WithStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_withStatement
	return p
}

func (*WithStatementContext) IsWithStatementContext() {}

func NewWithStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WithStatementContext {
	var p = new(WithStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_withStatement

	return p
}

func (s *WithStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WithStatementContext) With() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserWith, 0)
}

func (s *WithStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *WithStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WithStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WithStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WithStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterWithStatement(s)
	}
}

func (s *WithStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitWithStatement(s)
	}
}

func (p *ECMAScriptParser) WithStatement() (localctx IWithStatementContext) {
	localctx = NewWithStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ECMAScriptParserRULE_withStatement)

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
		p.SetState(267)
		p.Match(ECMAScriptParserWith)
	}
	{
		p.SetState(268)
		p.Match(ECMAScriptParserOpenParen)
	}
	{
		p.SetState(269)
		p.ExpressionSequence()
	}
	{
		p.SetState(270)
		p.Match(ECMAScriptParserCloseParen)
	}
	{
		p.SetState(271)
		p.Statement()
	}

	return localctx
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_switchStatement
	return p
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) Switch() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserSwitch, 0)
}

func (s *SwitchStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *SwitchStatementContext) CaseBlock() ICaseBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICaseBlockContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (p *ECMAScriptParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ECMAScriptParserRULE_switchStatement)

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
		p.Match(ECMAScriptParserSwitch)
	}
	{
		p.SetState(274)
		p.Match(ECMAScriptParserOpenParen)
	}
	{
		p.SetState(275)
		p.ExpressionSequence()
	}
	{
		p.SetState(276)
		p.Match(ECMAScriptParserCloseParen)
	}
	{
		p.SetState(277)
		p.CaseBlock()
	}

	return localctx
}

// ICaseBlockContext is an interface to support dynamic dispatch.
type ICaseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseBlockContext differentiates from other interfaces.
	IsCaseBlockContext()
}

type CaseBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseBlockContext() *CaseBlockContext {
	var p = new(CaseBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_caseBlock
	return p
}

func (*CaseBlockContext) IsCaseBlockContext() {}

func NewCaseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseBlockContext {
	var p = new(CaseBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_caseBlock

	return p
}

func (s *CaseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseBlockContext) AllCaseClauses() []ICaseClausesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseClausesContext)(nil)).Elem())
	var tst = make([]ICaseClausesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseClausesContext)
		}
	}

	return tst
}

func (s *CaseBlockContext) CaseClauses(i int) ICaseClausesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseClausesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseClausesContext)
}

func (s *CaseBlockContext) DefaultClause() IDefaultClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultClauseContext)
}

func (s *CaseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterCaseBlock(s)
	}
}

func (s *CaseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitCaseBlock(s)
	}
}

func (p *ECMAScriptParser) CaseBlock() (localctx ICaseBlockContext) {
	localctx = NewCaseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ECMAScriptParserRULE_caseBlock)
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
		p.SetState(279)
		p.Match(ECMAScriptParserOpenBrace)
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ECMAScriptParserCase {
		{
			p.SetState(280)
			p.CaseClauses()
		}

	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ECMAScriptParserDefault {
		{
			p.SetState(283)
			p.DefaultClause()
		}
		p.SetState(285)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ECMAScriptParserCase {
			{
				p.SetState(284)
				p.CaseClauses()
			}

		}

	}
	{
		p.SetState(289)
		p.Match(ECMAScriptParserCloseBrace)
	}

	return localctx
}

// ICaseClausesContext is an interface to support dynamic dispatch.
type ICaseClausesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseClausesContext differentiates from other interfaces.
	IsCaseClausesContext()
}

type CaseClausesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseClausesContext() *CaseClausesContext {
	var p = new(CaseClausesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_caseClauses
	return p
}

func (*CaseClausesContext) IsCaseClausesContext() {}

func NewCaseClausesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseClausesContext {
	var p = new(CaseClausesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_caseClauses

	return p
}

func (s *CaseClausesContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseClausesContext) AllCaseClause() []ICaseClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseClauseContext)(nil)).Elem())
	var tst = make([]ICaseClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseClauseContext)
		}
	}

	return tst
}

func (s *CaseClausesContext) CaseClause(i int) ICaseClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseClauseContext)
}

func (s *CaseClausesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseClausesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseClausesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterCaseClauses(s)
	}
}

func (s *CaseClausesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitCaseClauses(s)
	}
}

func (p *ECMAScriptParser) CaseClauses() (localctx ICaseClausesContext) {
	localctx = NewCaseClausesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ECMAScriptParserRULE_caseClauses)
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
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ECMAScriptParserCase {
		{
			p.SetState(291)
			p.CaseClause()
		}

		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICaseClauseContext is an interface to support dynamic dispatch.
type ICaseClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCaseClauseContext differentiates from other interfaces.
	IsCaseClauseContext()
}

type CaseClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseClauseContext() *CaseClauseContext {
	var p = new(CaseClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_caseClause
	return p
}

func (*CaseClauseContext) IsCaseClauseContext() {}

func NewCaseClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseClauseContext {
	var p = new(CaseClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_caseClause

	return p
}

func (s *CaseClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseClauseContext) Case() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCase, 0)
}

func (s *CaseClauseContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *CaseClauseContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *CaseClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterCaseClause(s)
	}
}

func (s *CaseClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitCaseClause(s)
	}
}

func (p *ECMAScriptParser) CaseClause() (localctx ICaseClauseContext) {
	localctx = NewCaseClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ECMAScriptParserRULE_caseClause)

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
		p.Match(ECMAScriptParserCase)
	}
	{
		p.SetState(297)
		p.ExpressionSequence()
	}
	{
		p.SetState(298)
		p.Match(ECMAScriptParserColon)
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(299)
			p.StatementList()
		}

	}

	return localctx
}

// IDefaultClauseContext is an interface to support dynamic dispatch.
type IDefaultClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultClauseContext differentiates from other interfaces.
	IsDefaultClauseContext()
}

type DefaultClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultClauseContext() *DefaultClauseContext {
	var p = new(DefaultClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_defaultClause
	return p
}

func (*DefaultClauseContext) IsDefaultClauseContext() {}

func NewDefaultClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultClauseContext {
	var p = new(DefaultClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_defaultClause

	return p
}

func (s *DefaultClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultClauseContext) Default() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDefault, 0)
}

func (s *DefaultClauseContext) StatementList() IStatementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *DefaultClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterDefaultClause(s)
	}
}

func (s *DefaultClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitDefaultClause(s)
	}
}

func (p *ECMAScriptParser) DefaultClause() (localctx IDefaultClauseContext) {
	localctx = NewDefaultClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ECMAScriptParserRULE_defaultClause)

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
		p.SetState(302)
		p.Match(ECMAScriptParserDefault)
	}
	{
		p.SetState(303)
		p.Match(ECMAScriptParserColon)
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(304)
			p.StatementList()
		}

	}

	return localctx
}

// ILabelledStatementContext is an interface to support dynamic dispatch.
type ILabelledStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelledStatementContext differentiates from other interfaces.
	IsLabelledStatementContext()
}

type LabelledStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelledStatementContext() *LabelledStatementContext {
	var p = new(LabelledStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_labelledStatement
	return p
}

func (*LabelledStatementContext) IsLabelledStatementContext() {}

func NewLabelledStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelledStatementContext {
	var p = new(LabelledStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_labelledStatement

	return p
}

func (s *LabelledStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelledStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *LabelledStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *LabelledStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelledStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelledStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterLabelledStatement(s)
	}
}

func (s *LabelledStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitLabelledStatement(s)
	}
}

func (p *ECMAScriptParser) LabelledStatement() (localctx ILabelledStatementContext) {
	localctx = NewLabelledStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ECMAScriptParserRULE_labelledStatement)

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
		p.Match(ECMAScriptParserIdentifier)
	}
	{
		p.SetState(308)
		p.Match(ECMAScriptParserColon)
	}
	{
		p.SetState(309)
		p.Statement()
	}

	return localctx
}

// IThrowStatementContext is an interface to support dynamic dispatch.
type IThrowStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThrowStatementContext differentiates from other interfaces.
	IsThrowStatementContext()
}

type ThrowStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThrowStatementContext() *ThrowStatementContext {
	var p = new(ThrowStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_throwStatement
	return p
}

func (*ThrowStatementContext) IsThrowStatementContext() {}

func NewThrowStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThrowStatementContext {
	var p = new(ThrowStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_throwStatement

	return p
}

func (s *ThrowStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ThrowStatementContext) Throw() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserThrow, 0)
}

func (s *ThrowStatementContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ThrowStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ThrowStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThrowStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThrowStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterThrowStatement(s)
	}
}

func (s *ThrowStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitThrowStatement(s)
	}
}

func (p *ECMAScriptParser) ThrowStatement() (localctx IThrowStatementContext) {
	localctx = NewThrowStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ECMAScriptParserRULE_throwStatement)

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
		p.SetState(311)
		p.Match(ECMAScriptParserThrow)
	}
	{
		p.SetState(312)
		p.ExpressionSequence()
	}
	{
		p.SetState(313)
		p.Eos()
	}

	return localctx
}

// ITryStatementContext is an interface to support dynamic dispatch.
type ITryStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTryStatementContext differentiates from other interfaces.
	IsTryStatementContext()
}

type TryStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTryStatementContext() *TryStatementContext {
	var p = new(TryStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_tryStatement
	return p
}

func (*TryStatementContext) IsTryStatementContext() {}

func NewTryStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TryStatementContext {
	var p = new(TryStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_tryStatement

	return p
}

func (s *TryStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *TryStatementContext) Try() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserTry, 0)
}

func (s *TryStatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *TryStatementContext) CatchProduction() ICatchProductionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICatchProductionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICatchProductionContext)
}

func (s *TryStatementContext) FinallyProduction() IFinallyProductionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFinallyProductionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFinallyProductionContext)
}

func (s *TryStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TryStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TryStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterTryStatement(s)
	}
}

func (s *TryStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitTryStatement(s)
	}
}

func (p *ECMAScriptParser) TryStatement() (localctx ITryStatementContext) {
	localctx = NewTryStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ECMAScriptParserRULE_tryStatement)

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

	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(315)
			p.Match(ECMAScriptParserTry)
		}
		{
			p.SetState(316)
			p.Block()
		}
		{
			p.SetState(317)
			p.CatchProduction()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(319)
			p.Match(ECMAScriptParserTry)
		}
		{
			p.SetState(320)
			p.Block()
		}
		{
			p.SetState(321)
			p.FinallyProduction()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(323)
			p.Match(ECMAScriptParserTry)
		}
		{
			p.SetState(324)
			p.Block()
		}
		{
			p.SetState(325)
			p.CatchProduction()
		}
		{
			p.SetState(326)
			p.FinallyProduction()
		}

	}

	return localctx
}

// ICatchProductionContext is an interface to support dynamic dispatch.
type ICatchProductionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCatchProductionContext differentiates from other interfaces.
	IsCatchProductionContext()
}

type CatchProductionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCatchProductionContext() *CatchProductionContext {
	var p = new(CatchProductionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_catchProduction
	return p
}

func (*CatchProductionContext) IsCatchProductionContext() {}

func NewCatchProductionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CatchProductionContext {
	var p = new(CatchProductionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_catchProduction

	return p
}

func (s *CatchProductionContext) GetParser() antlr.Parser { return s.parser }

func (s *CatchProductionContext) Catch() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCatch, 0)
}

func (s *CatchProductionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *CatchProductionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CatchProductionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CatchProductionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CatchProductionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterCatchProduction(s)
	}
}

func (s *CatchProductionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitCatchProduction(s)
	}
}

func (p *ECMAScriptParser) CatchProduction() (localctx ICatchProductionContext) {
	localctx = NewCatchProductionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ECMAScriptParserRULE_catchProduction)

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
		p.SetState(330)
		p.Match(ECMAScriptParserCatch)
	}
	{
		p.SetState(331)
		p.Match(ECMAScriptParserOpenParen)
	}
	{
		p.SetState(332)
		p.Match(ECMAScriptParserIdentifier)
	}
	{
		p.SetState(333)
		p.Match(ECMAScriptParserCloseParen)
	}
	{
		p.SetState(334)
		p.Block()
	}

	return localctx
}

// IFinallyProductionContext is an interface to support dynamic dispatch.
type IFinallyProductionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFinallyProductionContext differentiates from other interfaces.
	IsFinallyProductionContext()
}

type FinallyProductionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFinallyProductionContext() *FinallyProductionContext {
	var p = new(FinallyProductionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_finallyProduction
	return p
}

func (*FinallyProductionContext) IsFinallyProductionContext() {}

func NewFinallyProductionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FinallyProductionContext {
	var p = new(FinallyProductionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_finallyProduction

	return p
}

func (s *FinallyProductionContext) GetParser() antlr.Parser { return s.parser }

func (s *FinallyProductionContext) Finally() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFinally, 0)
}

func (s *FinallyProductionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FinallyProductionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FinallyProductionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FinallyProductionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterFinallyProduction(s)
	}
}

func (s *FinallyProductionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitFinallyProduction(s)
	}
}

func (p *ECMAScriptParser) FinallyProduction() (localctx IFinallyProductionContext) {
	localctx = NewFinallyProductionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ECMAScriptParserRULE_finallyProduction)

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
		p.SetState(336)
		p.Match(ECMAScriptParserFinally)
	}
	{
		p.SetState(337)
		p.Block()
	}

	return localctx
}

// IDebuggerStatementContext is an interface to support dynamic dispatch.
type IDebuggerStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDebuggerStatementContext differentiates from other interfaces.
	IsDebuggerStatementContext()
}

type DebuggerStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDebuggerStatementContext() *DebuggerStatementContext {
	var p = new(DebuggerStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_debuggerStatement
	return p
}

func (*DebuggerStatementContext) IsDebuggerStatementContext() {}

func NewDebuggerStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DebuggerStatementContext {
	var p = new(DebuggerStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_debuggerStatement

	return p
}

func (s *DebuggerStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DebuggerStatementContext) Debugger() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDebugger, 0)
}

func (s *DebuggerStatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *DebuggerStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DebuggerStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DebuggerStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterDebuggerStatement(s)
	}
}

func (s *DebuggerStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitDebuggerStatement(s)
	}
}

func (p *ECMAScriptParser) DebuggerStatement() (localctx IDebuggerStatementContext) {
	localctx = NewDebuggerStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ECMAScriptParserRULE_debuggerStatement)

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
		p.Match(ECMAScriptParserDebugger)
	}
	{
		p.SetState(340)
		p.Eos()
	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) Function() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFunction, 0)
}

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionDeclarationContext) FormalParameterList() IFormalParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (p *ECMAScriptParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ECMAScriptParserRULE_functionDeclaration)
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
		p.SetState(342)
		p.Match(ECMAScriptParserFunction)
	}
	{
		p.SetState(343)
		p.Match(ECMAScriptParserIdentifier)
	}
	{
		p.SetState(344)
		p.Match(ECMAScriptParserOpenParen)
	}
	p.SetState(346)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ECMAScriptParserIdentifier {
		{
			p.SetState(345)
			p.FormalParameterList()
		}

	}
	{
		p.SetState(348)
		p.Match(ECMAScriptParserCloseParen)
	}
	{
		p.SetState(349)
		p.Match(ECMAScriptParserOpenBrace)
	}
	{
		p.SetState(350)
		p.FunctionBody()
	}
	{
		p.SetState(351)
		p.Match(ECMAScriptParserCloseBrace)
	}

	return localctx
}

// IFormalParameterListContext is an interface to support dynamic dispatch.
type IFormalParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFormalParameterListContext differentiates from other interfaces.
	IsFormalParameterListContext()
}

type FormalParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFormalParameterListContext() *FormalParameterListContext {
	var p = new(FormalParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_formalParameterList
	return p
}

func (*FormalParameterListContext) IsFormalParameterListContext() {}

func NewFormalParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormalParameterListContext {
	var p = new(FormalParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_formalParameterList

	return p
}

func (s *FormalParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *FormalParameterListContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(ECMAScriptParserIdentifier)
}

func (s *FormalParameterListContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, i)
}

func (s *FormalParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormalParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormalParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterFormalParameterList(s)
	}
}

func (s *FormalParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitFormalParameterList(s)
	}
}

func (p *ECMAScriptParser) FormalParameterList() (localctx IFormalParameterListContext) {
	localctx = NewFormalParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ECMAScriptParserRULE_formalParameterList)
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
		p.SetState(353)
		p.Match(ECMAScriptParserIdentifier)
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ECMAScriptParserComma {
		{
			p.SetState(354)
			p.Match(ECMAScriptParserComma)
		}
		{
			p.SetState(355)
			p.Match(ECMAScriptParserIdentifier)
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunctionBodyContext is an interface to support dynamic dispatch.
type IFunctionBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionBodyContext differentiates from other interfaces.
	IsFunctionBodyContext()
}

type FunctionBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionBodyContext() *FunctionBodyContext {
	var p = new(FunctionBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_functionBody
	return p
}

func (*FunctionBodyContext) IsFunctionBodyContext() {}

func NewFunctionBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionBodyContext {
	var p = new(FunctionBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_functionBody

	return p
}

func (s *FunctionBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionBodyContext) SourceElements() ISourceElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceElementsContext)
}

func (s *FunctionBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterFunctionBody(s)
	}
}

func (s *FunctionBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitFunctionBody(s)
	}
}

func (p *ECMAScriptParser) FunctionBody() (localctx IFunctionBodyContext) {
	localctx = NewFunctionBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ECMAScriptParserRULE_functionBody)

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
	p.SetState(362)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(361)
			p.SourceElements()
		}

	}

	return localctx
}

// IArrayLiteralContext is an interface to support dynamic dispatch.
type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayLiteralContext differentiates from other interfaces.
	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_arrayLiteral
	return p
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) ElementList() IElementListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementListContext)
}

func (s *ArrayLiteralContext) Elision() IElisionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElisionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElisionContext)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (p *ECMAScriptParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ECMAScriptParserRULE_arrayLiteral)
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
		p.SetState(364)
		p.Match(ECMAScriptParserOpenBracket)
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(365)
			p.ElementList()
		}

	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(368)
			p.Match(ECMAScriptParserComma)
		}

	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ECMAScriptParserComma {
		{
			p.SetState(371)
			p.Elision()
		}

	}
	{
		p.SetState(374)
		p.Match(ECMAScriptParserCloseBracket)
	}

	return localctx
}

// IElementListContext is an interface to support dynamic dispatch.
type IElementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementListContext differentiates from other interfaces.
	IsElementListContext()
}

type ElementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementListContext() *ElementListContext {
	var p = new(ElementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_elementList
	return p
}

func (*ElementListContext) IsElementListContext() {}

func NewElementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementListContext {
	var p = new(ElementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_elementList

	return p
}

func (s *ElementListContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementListContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ElementListContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ElementListContext) AllElision() []IElisionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElisionContext)(nil)).Elem())
	var tst = make([]IElisionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElisionContext)
		}
	}

	return tst
}

func (s *ElementListContext) Elision(i int) IElisionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElisionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElisionContext)
}

func (s *ElementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterElementList(s)
	}
}

func (s *ElementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitElementList(s)
	}
}

func (p *ECMAScriptParser) ElementList() (localctx IElementListContext) {
	localctx = NewElementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ECMAScriptParserRULE_elementList)
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
	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ECMAScriptParserComma {
		{
			p.SetState(376)
			p.Elision()
		}

	}
	{
		p.SetState(379)
		p.singleExpression(0)
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(380)
				p.Match(ECMAScriptParserComma)
			}
			p.SetState(382)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == ECMAScriptParserComma {
				{
					p.SetState(381)
					p.Elision()
				}

			}
			{
				p.SetState(384)
				p.singleExpression(0)
			}

		}
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}

	return localctx
}

// IElisionContext is an interface to support dynamic dispatch.
type IElisionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElisionContext differentiates from other interfaces.
	IsElisionContext()
}

type ElisionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElisionContext() *ElisionContext {
	var p = new(ElisionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_elision
	return p
}

func (*ElisionContext) IsElisionContext() {}

func NewElisionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElisionContext {
	var p = new(ElisionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_elision

	return p
}

func (s *ElisionContext) GetParser() antlr.Parser { return s.parser }
func (s *ElisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElisionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterElision(s)
	}
}

func (s *ElisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitElision(s)
	}
}

func (p *ECMAScriptParser) Elision() (localctx IElisionContext) {
	localctx = NewElisionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ECMAScriptParserRULE_elision)
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
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ECMAScriptParserComma {
		{
			p.SetState(390)
			p.Match(ECMAScriptParserComma)
		}

		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IObjectLiteralContext is an interface to support dynamic dispatch.
type IObjectLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectLiteralContext differentiates from other interfaces.
	IsObjectLiteralContext()
}

type ObjectLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectLiteralContext() *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_objectLiteral
	return p
}

func (*ObjectLiteralContext) IsObjectLiteralContext() {}

func NewObjectLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_objectLiteral

	return p
}

func (s *ObjectLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectLiteralContext) PropertyNameAndValueList() IPropertyNameAndValueListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameAndValueListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameAndValueListContext)
}

func (s *ObjectLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitObjectLiteral(s)
	}
}

func (p *ECMAScriptParser) ObjectLiteral() (localctx IObjectLiteralContext) {
	localctx = NewObjectLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ECMAScriptParserRULE_objectLiteral)
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

	p.SetState(404)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(395)
			p.Match(ECMAScriptParserOpenBrace)
		}
		{
			p.SetState(396)
			p.Match(ECMAScriptParserCloseBrace)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(397)
			p.Match(ECMAScriptParserOpenBrace)
		}
		{
			p.SetState(398)
			p.PropertyNameAndValueList()
		}
		p.SetState(400)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ECMAScriptParserComma {
			{
				p.SetState(399)
				p.Match(ECMAScriptParserComma)
			}

		}
		{
			p.SetState(402)
			p.Match(ECMAScriptParserCloseBrace)
		}

	}

	return localctx
}

// IPropertyNameAndValueListContext is an interface to support dynamic dispatch.
type IPropertyNameAndValueListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyNameAndValueListContext differentiates from other interfaces.
	IsPropertyNameAndValueListContext()
}

type PropertyNameAndValueListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyNameAndValueListContext() *PropertyNameAndValueListContext {
	var p = new(PropertyNameAndValueListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_propertyNameAndValueList
	return p
}

func (*PropertyNameAndValueListContext) IsPropertyNameAndValueListContext() {}

func NewPropertyNameAndValueListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameAndValueListContext {
	var p = new(PropertyNameAndValueListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_propertyNameAndValueList

	return p
}

func (s *PropertyNameAndValueListContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameAndValueListContext) AllPropertyAssignment() []IPropertyAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyAssignmentContext)(nil)).Elem())
	var tst = make([]IPropertyAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyAssignmentContext)
		}
	}

	return tst
}

func (s *PropertyNameAndValueListContext) PropertyAssignment(i int) IPropertyAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyAssignmentContext)
}

func (s *PropertyNameAndValueListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameAndValueListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameAndValueListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPropertyNameAndValueList(s)
	}
}

func (s *PropertyNameAndValueListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPropertyNameAndValueList(s)
	}
}

func (p *ECMAScriptParser) PropertyNameAndValueList() (localctx IPropertyNameAndValueListContext) {
	localctx = NewPropertyNameAndValueListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ECMAScriptParserRULE_propertyNameAndValueList)

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
		p.SetState(406)
		p.PropertyAssignment()
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(407)
				p.Match(ECMAScriptParserComma)
			}
			{
				p.SetState(408)
				p.PropertyAssignment()
			}

		}
		p.SetState(413)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}

	return localctx
}

// IPropertyAssignmentContext is an interface to support dynamic dispatch.
type IPropertyAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyAssignmentContext differentiates from other interfaces.
	IsPropertyAssignmentContext()
}

type PropertyAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyAssignmentContext() *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_propertyAssignment
	return p
}

func (*PropertyAssignmentContext) IsPropertyAssignmentContext() {}

func NewPropertyAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyAssignmentContext {
	var p = new(PropertyAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_propertyAssignment

	return p
}

func (s *PropertyAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyAssignmentContext) CopyFrom(ctx *PropertyAssignmentContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PropertyAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropertyExpressionAssignmentContext struct {
	*PropertyAssignmentContext
}

func NewPropertyExpressionAssignmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyExpressionAssignmentContext {
	var p = new(PropertyExpressionAssignmentContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertyExpressionAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyExpressionAssignmentContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *PropertyExpressionAssignmentContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PropertyExpressionAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPropertyExpressionAssignment(s)
	}
}

func (s *PropertyExpressionAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPropertyExpressionAssignment(s)
	}
}

type PropertySetterContext struct {
	*PropertyAssignmentContext
}

func NewPropertySetterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertySetterContext {
	var p = new(PropertySetterContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertySetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertySetterContext) Setter() ISetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetterContext)
}

func (s *PropertySetterContext) PropertySetParameterList() IPropertySetParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertySetParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertySetParameterListContext)
}

func (s *PropertySetterContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *PropertySetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPropertySetter(s)
	}
}

func (s *PropertySetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPropertySetter(s)
	}
}

type PropertyGetterContext struct {
	*PropertyAssignmentContext
}

func NewPropertyGetterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropertyGetterContext {
	var p = new(PropertyGetterContext)

	p.PropertyAssignmentContext = NewEmptyPropertyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PropertyAssignmentContext))

	return p
}

func (s *PropertyGetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyGetterContext) Getter() IGetterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGetterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGetterContext)
}

func (s *PropertyGetterContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *PropertyGetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPropertyGetter(s)
	}
}

func (s *PropertyGetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPropertyGetter(s)
	}
}

func (p *ECMAScriptParser) PropertyAssignment() (localctx IPropertyAssignmentContext) {
	localctx = NewPropertyAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ECMAScriptParserRULE_propertyAssignment)

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

	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPropertyExpressionAssignmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(414)
			p.PropertyName()
		}
		{
			p.SetState(415)
			p.Match(ECMAScriptParserColon)
		}
		{
			p.SetState(416)
			p.singleExpression(0)
		}

	case 2:
		localctx = NewPropertyGetterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(418)
			p.Getter()
		}
		{
			p.SetState(419)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(420)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(421)
			p.Match(ECMAScriptParserOpenBrace)
		}
		{
			p.SetState(422)
			p.FunctionBody()
		}
		{
			p.SetState(423)
			p.Match(ECMAScriptParserCloseBrace)
		}

	case 3:
		localctx = NewPropertySetterContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(425)
			p.Setter()
		}
		{
			p.SetState(426)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(427)
			p.PropertySetParameterList()
		}
		{
			p.SetState(428)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(429)
			p.Match(ECMAScriptParserOpenBrace)
		}
		{
			p.SetState(430)
			p.FunctionBody()
		}
		{
			p.SetState(431)
			p.Match(ECMAScriptParserCloseBrace)
		}

	}

	return localctx
}

// IPropertyNameContext is an interface to support dynamic dispatch.
type IPropertyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyNameContext differentiates from other interfaces.
	IsPropertyNameContext()
}

type PropertyNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyNameContext() *PropertyNameContext {
	var p = new(PropertyNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_propertyName
	return p
}

func (*PropertyNameContext) IsPropertyNameContext() {}

func NewPropertyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyNameContext {
	var p = new(PropertyNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_propertyName

	return p
}

func (s *PropertyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyNameContext) IdentifierName() IIdentifierNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierNameContext)
}

func (s *PropertyNameContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserStringLiteral, 0)
}

func (s *PropertyNameContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *PropertyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPropertyName(s)
	}
}

func (s *PropertyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPropertyName(s)
	}
}

func (p *ECMAScriptParser) PropertyName() (localctx IPropertyNameContext) {
	localctx = NewPropertyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ECMAScriptParserRULE_propertyName)

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

	p.SetState(438)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ECMAScriptParserNullLiteral, ECMAScriptParserBooleanLiteral, ECMAScriptParserBreak, ECMAScriptParserDo, ECMAScriptParserInstanceof, ECMAScriptParserTypeof, ECMAScriptParserCase, ECMAScriptParserElse, ECMAScriptParserNew, ECMAScriptParserVar, ECMAScriptParserCatch, ECMAScriptParserFinally, ECMAScriptParserReturn, ECMAScriptParserVoid, ECMAScriptParserContinue, ECMAScriptParserFor, ECMAScriptParserSwitch, ECMAScriptParserWhile, ECMAScriptParserDebugger, ECMAScriptParserFunction, ECMAScriptParserThis, ECMAScriptParserWith, ECMAScriptParserDefault, ECMAScriptParserIf, ECMAScriptParserThrow, ECMAScriptParserDelete, ECMAScriptParserIn, ECMAScriptParserTry, ECMAScriptParserClass, ECMAScriptParserEnum, ECMAScriptParserExtends, ECMAScriptParserSuper, ECMAScriptParserConst, ECMAScriptParserExport, ECMAScriptParserImport, ECMAScriptParserImplements, ECMAScriptParserLet, ECMAScriptParserPrivate, ECMAScriptParserPublic, ECMAScriptParserInterface, ECMAScriptParserPackage, ECMAScriptParserProtected, ECMAScriptParserStatic, ECMAScriptParserYield, ECMAScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(435)
			p.IdentifierName()
		}

	case ECMAScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(436)
			p.Match(ECMAScriptParserStringLiteral)
		}

	case ECMAScriptParserDecimalLiteral, ECMAScriptParserHexIntegerLiteral, ECMAScriptParserOctalIntegerLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(437)
			p.NumericLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPropertySetParameterListContext is an interface to support dynamic dispatch.
type IPropertySetParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertySetParameterListContext differentiates from other interfaces.
	IsPropertySetParameterListContext()
}

type PropertySetParameterListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertySetParameterListContext() *PropertySetParameterListContext {
	var p = new(PropertySetParameterListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_propertySetParameterList
	return p
}

func (*PropertySetParameterListContext) IsPropertySetParameterListContext() {}

func NewPropertySetParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertySetParameterListContext {
	var p = new(PropertySetParameterListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_propertySetParameterList

	return p
}

func (s *PropertySetParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertySetParameterListContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *PropertySetParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertySetParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertySetParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPropertySetParameterList(s)
	}
}

func (s *PropertySetParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPropertySetParameterList(s)
	}
}

func (p *ECMAScriptParser) PropertySetParameterList() (localctx IPropertySetParameterListContext) {
	localctx = NewPropertySetParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ECMAScriptParserRULE_propertySetParameterList)

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
		p.SetState(440)
		p.Match(ECMAScriptParserIdentifier)
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
	p.RuleIndex = ECMAScriptParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) ArgumentList() IArgumentListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (p *ECMAScriptParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ECMAScriptParserRULE_arguments)
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
		p.SetState(442)
		p.Match(ECMAScriptParserOpenParen)
	}
	p.SetState(444)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRegularExpressionLiteral)|(1<<ECMAScriptParserOpenBracket)|(1<<ECMAScriptParserOpenParen)|(1<<ECMAScriptParserOpenBrace)|(1<<ECMAScriptParserPlusPlus)|(1<<ECMAScriptParserMinusMinus)|(1<<ECMAScriptParserPlus)|(1<<ECMAScriptParserMinus)|(1<<ECMAScriptParserBitNot)|(1<<ECMAScriptParserNot))) != 0) || (((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(ECMAScriptParserNullLiteral-51))|(1<<(ECMAScriptParserBooleanLiteral-51))|(1<<(ECMAScriptParserDecimalLiteral-51))|(1<<(ECMAScriptParserHexIntegerLiteral-51))|(1<<(ECMAScriptParserOctalIntegerLiteral-51))|(1<<(ECMAScriptParserTypeof-51))|(1<<(ECMAScriptParserNew-51))|(1<<(ECMAScriptParserVoid-51))|(1<<(ECMAScriptParserFunction-51))|(1<<(ECMAScriptParserThis-51))|(1<<(ECMAScriptParserDelete-51)))) != 0) || _la == ECMAScriptParserIdentifier || _la == ECMAScriptParserStringLiteral {
		{
			p.SetState(443)
			p.ArgumentList()
		}

	}
	{
		p.SetState(446)
		p.Match(ECMAScriptParserCloseParen)
	}

	return localctx
}

// IArgumentListContext is an interface to support dynamic dispatch.
type IArgumentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentListContext differentiates from other interfaces.
	IsArgumentListContext()
}

type ArgumentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_argumentList
	return p
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ArgumentListContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (p *ECMAScriptParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ECMAScriptParserRULE_argumentList)
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
		p.singleExpression(0)
	}
	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ECMAScriptParserComma {
		{
			p.SetState(449)
			p.Match(ECMAScriptParserComma)
		}
		{
			p.SetState(450)
			p.singleExpression(0)
		}

		p.SetState(455)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExpressionSequenceContext is an interface to support dynamic dispatch.
type IExpressionSequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionSequenceContext differentiates from other interfaces.
	IsExpressionSequenceContext()
}

type ExpressionSequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionSequenceContext() *ExpressionSequenceContext {
	var p = new(ExpressionSequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_expressionSequence
	return p
}

func (*ExpressionSequenceContext) IsExpressionSequenceContext() {}

func NewExpressionSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionSequenceContext {
	var p = new(ExpressionSequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_expressionSequence

	return p
}

func (s *ExpressionSequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionSequenceContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionSequenceContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ExpressionSequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionSequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionSequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterExpressionSequence(s)
	}
}

func (s *ExpressionSequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitExpressionSequence(s)
	}
}

func (p *ECMAScriptParser) ExpressionSequence() (localctx IExpressionSequenceContext) {
	localctx = NewExpressionSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ECMAScriptParserRULE_expressionSequence)

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
		p.SetState(456)
		p.singleExpression(0)
	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(457)
				p.Match(ECMAScriptParserComma)
			}
			{
				p.SetState(458)
				p.singleExpression(0)
			}

		}
		p.SetState(463)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext())
	}

	return localctx
}

// ISingleExpressionContext is an interface to support dynamic dispatch.
type ISingleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleExpressionContext differentiates from other interfaces.
	IsSingleExpressionContext()
}

type SingleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleExpressionContext() *SingleExpressionContext {
	var p = new(SingleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_singleExpression
	return p
}

func (*SingleExpressionContext) IsSingleExpressionContext() {}

func NewSingleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExpressionContext {
	var p = new(SingleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_singleExpression

	return p
}

func (s *SingleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExpressionContext) CopyFrom(ctx *SingleExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SingleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TernaryExpressionContext struct {
	*SingleExpressionContext
}

func NewTernaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpressionContext {
	var p = new(TernaryExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *TernaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *TernaryExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *TernaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterTernaryExpression(s)
	}
}

func (s *TernaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitTernaryExpression(s)
	}
}

type LogicalAndExpressionContext struct {
	*SingleExpressionContext
}

func NewLogicalAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalAndExpressionContext {
	var p = new(LogicalAndExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LogicalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalAndExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *LogicalAndExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *LogicalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterLogicalAndExpression(s)
	}
}

func (s *LogicalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitLogicalAndExpression(s)
	}
}

type PreIncrementExpressionContext struct {
	*SingleExpressionContext
}

func NewPreIncrementExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreIncrementExpressionContext {
	var p = new(PreIncrementExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PreIncrementExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreIncrementExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PreIncrementExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPreIncrementExpression(s)
	}
}

func (s *PreIncrementExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPreIncrementExpression(s)
	}
}

type ObjectLiteralExpressionContext struct {
	*SingleExpressionContext
}

func NewObjectLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectLiteralExpressionContext {
	var p = new(ObjectLiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ObjectLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralExpressionContext) ObjectLiteral() IObjectLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectLiteralContext)
}

func (s *ObjectLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterObjectLiteralExpression(s)
	}
}

func (s *ObjectLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitObjectLiteralExpression(s)
	}
}

type InExpressionContext struct {
	*SingleExpressionContext
}

func NewInExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InExpressionContext {
	var p = new(InExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *InExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *InExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *InExpressionContext) In() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIn, 0)
}

func (s *InExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterInExpression(s)
	}
}

func (s *InExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitInExpression(s)
	}
}

type LogicalOrExpressionContext struct {
	*SingleExpressionContext
}

func NewLogicalOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalOrExpressionContext {
	var p = new(LogicalOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LogicalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *LogicalOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *LogicalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterLogicalOrExpression(s)
	}
}

func (s *LogicalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitLogicalOrExpression(s)
	}
}

type NotExpressionContext struct {
	*SingleExpressionContext
}

func NewNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExpressionContext {
	var p = new(NotExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *NotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *NotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterNotExpression(s)
	}
}

func (s *NotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitNotExpression(s)
	}
}

type PreDecreaseExpressionContext struct {
	*SingleExpressionContext
}

func NewPreDecreaseExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PreDecreaseExpressionContext {
	var p = new(PreDecreaseExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PreDecreaseExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PreDecreaseExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PreDecreaseExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPreDecreaseExpression(s)
	}
}

func (s *PreDecreaseExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPreDecreaseExpression(s)
	}
}

type ArgumentsExpressionContext struct {
	*SingleExpressionContext
}

func NewArgumentsExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArgumentsExpressionContext {
	var p = new(ArgumentsExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ArgumentsExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *ArgumentsExpressionContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *ArgumentsExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterArgumentsExpression(s)
	}
}

func (s *ArgumentsExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitArgumentsExpression(s)
	}
}

type ThisExpressionContext struct {
	*SingleExpressionContext
}

func NewThisExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ThisExpressionContext {
	var p = new(ThisExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ThisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThisExpressionContext) This() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserThis, 0)
}

func (s *ThisExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterThisExpression(s)
	}
}

func (s *ThisExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitThisExpression(s)
	}
}

type FunctionExpressionContext struct {
	*SingleExpressionContext
}

func NewFunctionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionExpressionContext {
	var p = new(FunctionExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *FunctionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionExpressionContext) Function() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFunction, 0)
}

func (s *FunctionExpressionContext) FunctionBody() IFunctionBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionBodyContext)
}

func (s *FunctionExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *FunctionExpressionContext) FormalParameterList() IFormalParameterListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormalParameterListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormalParameterListContext)
}

func (s *FunctionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterFunctionExpression(s)
	}
}

func (s *FunctionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitFunctionExpression(s)
	}
}

type UnaryMinusExpressionContext struct {
	*SingleExpressionContext
}

func NewUnaryMinusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryMinusExpressionContext {
	var p = new(UnaryMinusExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *UnaryMinusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryMinusExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *UnaryMinusExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterUnaryMinusExpression(s)
	}
}

func (s *UnaryMinusExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitUnaryMinusExpression(s)
	}
}

type AssignmentExpressionContext struct {
	*SingleExpressionContext
}

func NewAssignmentExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *AssignmentExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

type PostDecreaseExpressionContext struct {
	*SingleExpressionContext
}

func NewPostDecreaseExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostDecreaseExpressionContext {
	var p = new(PostDecreaseExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PostDecreaseExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostDecreaseExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PostDecreaseExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPostDecreaseExpression(s)
	}
}

func (s *PostDecreaseExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPostDecreaseExpression(s)
	}
}

type TypeofExpressionContext struct {
	*SingleExpressionContext
}

func NewTypeofExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeofExpressionContext {
	var p = new(TypeofExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *TypeofExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeofExpressionContext) Typeof() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserTypeof, 0)
}

func (s *TypeofExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *TypeofExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterTypeofExpression(s)
	}
}

func (s *TypeofExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitTypeofExpression(s)
	}
}

type InstanceofExpressionContext struct {
	*SingleExpressionContext
}

func NewInstanceofExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InstanceofExpressionContext {
	var p = new(InstanceofExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *InstanceofExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstanceofExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *InstanceofExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *InstanceofExpressionContext) Instanceof() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserInstanceof, 0)
}

func (s *InstanceofExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterInstanceofExpression(s)
	}
}

func (s *InstanceofExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitInstanceofExpression(s)
	}
}

type UnaryPlusExpressionContext struct {
	*SingleExpressionContext
}

func NewUnaryPlusExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryPlusExpressionContext {
	var p = new(UnaryPlusExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *UnaryPlusExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryPlusExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *UnaryPlusExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterUnaryPlusExpression(s)
	}
}

func (s *UnaryPlusExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitUnaryPlusExpression(s)
	}
}

type DeleteExpressionContext struct {
	*SingleExpressionContext
}

func NewDeleteExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeleteExpressionContext {
	var p = new(DeleteExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *DeleteExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteExpressionContext) Delete() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDelete, 0)
}

func (s *DeleteExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *DeleteExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterDeleteExpression(s)
	}
}

func (s *DeleteExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitDeleteExpression(s)
	}
}

type EqualityExpressionContext struct {
	*SingleExpressionContext
}

func NewEqualityExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualityExpressionContext {
	var p = new(EqualityExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *EqualityExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *EqualityExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *EqualityExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterEqualityExpression(s)
	}
}

func (s *EqualityExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitEqualityExpression(s)
	}
}

type BitXOrExpressionContext struct {
	*SingleExpressionContext
}

func NewBitXOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitXOrExpressionContext {
	var p = new(BitXOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitXOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitXOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitXOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitXOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBitXOrExpression(s)
	}
}

func (s *BitXOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBitXOrExpression(s)
	}
}

type MultiplicativeExpressionContext struct {
	*SingleExpressionContext
}

func NewMultiplicativeExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

type BitShiftExpressionContext struct {
	*SingleExpressionContext
}

func NewBitShiftExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitShiftExpressionContext {
	var p = new(BitShiftExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitShiftExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitShiftExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitShiftExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitShiftExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBitShiftExpression(s)
	}
}

func (s *BitShiftExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBitShiftExpression(s)
	}
}

type ParenthesizedExpressionContext struct {
	*SingleExpressionContext
}

func NewParenthesizedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesizedExpressionContext {
	var p = new(ParenthesizedExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ParenthesizedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesizedExpressionContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *ParenthesizedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterParenthesizedExpression(s)
	}
}

func (s *ParenthesizedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitParenthesizedExpression(s)
	}
}

type AdditiveExpressionContext struct {
	*SingleExpressionContext
}

func NewAdditiveExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

type RelationalExpressionContext struct {
	*SingleExpressionContext
}

func NewRelationalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *RelationalExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

type PostIncrementExpressionContext struct {
	*SingleExpressionContext
}

func NewPostIncrementExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PostIncrementExpressionContext {
	var p = new(PostIncrementExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *PostIncrementExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostIncrementExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *PostIncrementExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterPostIncrementExpression(s)
	}
}

func (s *PostIncrementExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitPostIncrementExpression(s)
	}
}

type BitNotExpressionContext struct {
	*SingleExpressionContext
}

func NewBitNotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitNotExpressionContext {
	var p = new(BitNotExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitNotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitNotExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitNotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBitNotExpression(s)
	}
}

func (s *BitNotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBitNotExpression(s)
	}
}

type NewExpressionContext struct {
	*SingleExpressionContext
}

func NewNewExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NewExpressionContext {
	var p = new(NewExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *NewExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NewExpressionContext) New() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserNew, 0)
}

func (s *NewExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *NewExpressionContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *NewExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterNewExpression(s)
	}
}

func (s *NewExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitNewExpression(s)
	}
}

type LiteralExpressionContext struct {
	*SingleExpressionContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

type ArrayLiteralExpressionContext struct {
	*SingleExpressionContext
}

func NewArrayLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayLiteralExpressionContext {
	var p = new(ArrayLiteralExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *ArrayLiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralExpressionContext) ArrayLiteral() IArrayLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ArrayLiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterArrayLiteralExpression(s)
	}
}

func (s *ArrayLiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitArrayLiteralExpression(s)
	}
}

type MemberDotExpressionContext struct {
	*SingleExpressionContext
}

func NewMemberDotExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberDotExpressionContext {
	var p = new(MemberDotExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MemberDotExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberDotExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *MemberDotExpressionContext) IdentifierName() IIdentifierNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierNameContext)
}

func (s *MemberDotExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterMemberDotExpression(s)
	}
}

func (s *MemberDotExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitMemberDotExpression(s)
	}
}

type MemberIndexExpressionContext struct {
	*SingleExpressionContext
}

func NewMemberIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberIndexExpressionContext {
	var p = new(MemberIndexExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *MemberIndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberIndexExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *MemberIndexExpressionContext) ExpressionSequence() IExpressionSequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionSequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionSequenceContext)
}

func (s *MemberIndexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterMemberIndexExpression(s)
	}
}

func (s *MemberIndexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitMemberIndexExpression(s)
	}
}

type IdentifierExpressionContext struct {
	*SingleExpressionContext
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

type BitAndExpressionContext struct {
	*SingleExpressionContext
}

func NewBitAndExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitAndExpressionContext {
	var p = new(BitAndExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitAndExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitAndExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBitAndExpression(s)
	}
}

func (s *BitAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBitAndExpression(s)
	}
}

type BitOrExpressionContext struct {
	*SingleExpressionContext
}

func NewBitOrExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BitOrExpressionContext {
	var p = new(BitOrExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *BitOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BitOrExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *BitOrExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *BitOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterBitOrExpression(s)
	}
}

func (s *BitOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitBitOrExpression(s)
	}
}

type AssignmentOperatorExpressionContext struct {
	*SingleExpressionContext
}

func NewAssignmentOperatorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentOperatorExpressionContext {
	var p = new(AssignmentOperatorExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *AssignmentOperatorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorExpressionContext) AllSingleExpression() []ISingleExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem())
	var tst = make([]ISingleExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISingleExpressionContext)
		}
	}

	return tst
}

func (s *AssignmentOperatorExpressionContext) SingleExpression(i int) ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *AssignmentOperatorExpressionContext) AssignmentOperator() IAssignmentOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentOperatorContext)
}

func (s *AssignmentOperatorExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterAssignmentOperatorExpression(s)
	}
}

func (s *AssignmentOperatorExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitAssignmentOperatorExpression(s)
	}
}

type VoidExpressionContext struct {
	*SingleExpressionContext
}

func NewVoidExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VoidExpressionContext {
	var p = new(VoidExpressionContext)

	p.SingleExpressionContext = NewEmptySingleExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SingleExpressionContext))

	return p
}

func (s *VoidExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VoidExpressionContext) Void() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserVoid, 0)
}

func (s *VoidExpressionContext) SingleExpression() ISingleExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleExpressionContext)
}

func (s *VoidExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterVoidExpression(s)
	}
}

func (s *VoidExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitVoidExpression(s)
	}
}

func (p *ECMAScriptParser) SingleExpression() (localctx ISingleExpressionContext) {
	return p.singleExpression(0)
}

func (p *ECMAScriptParser) singleExpression(_p int) (localctx ISingleExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSingleExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISingleExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 86
	p.EnterRecursionRule(localctx, 86, ECMAScriptParserRULE_singleExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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
	p.SetState(510)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ECMAScriptParserFunction:
		localctx = NewFunctionExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(465)
			p.Match(ECMAScriptParserFunction)
		}
		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ECMAScriptParserIdentifier {
			{
				p.SetState(466)
				p.Match(ECMAScriptParserIdentifier)
			}

		}
		{
			p.SetState(469)
			p.Match(ECMAScriptParserOpenParen)
		}
		p.SetState(471)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ECMAScriptParserIdentifier {
			{
				p.SetState(470)
				p.FormalParameterList()
			}

		}
		{
			p.SetState(473)
			p.Match(ECMAScriptParserCloseParen)
		}
		{
			p.SetState(474)
			p.Match(ECMAScriptParserOpenBrace)
		}
		{
			p.SetState(475)
			p.FunctionBody()
		}
		{
			p.SetState(476)
			p.Match(ECMAScriptParserCloseBrace)
		}

	case ECMAScriptParserNew:
		localctx = NewNewExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(478)
			p.Match(ECMAScriptParserNew)
		}
		{
			p.SetState(479)
			p.singleExpression(0)
		}
		p.SetState(481)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(480)
				p.Arguments()
			}

		}

	case ECMAScriptParserDelete:
		localctx = NewDeleteExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(483)
			p.Match(ECMAScriptParserDelete)
		}
		{
			p.SetState(484)
			p.singleExpression(30)
		}

	case ECMAScriptParserVoid:
		localctx = NewVoidExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(485)
			p.Match(ECMAScriptParserVoid)
		}
		{
			p.SetState(486)
			p.singleExpression(29)
		}

	case ECMAScriptParserTypeof:
		localctx = NewTypeofExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(487)
			p.Match(ECMAScriptParserTypeof)
		}
		{
			p.SetState(488)
			p.singleExpression(28)
		}

	case ECMAScriptParserPlusPlus:
		localctx = NewPreIncrementExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(489)
			p.Match(ECMAScriptParserPlusPlus)
		}
		{
			p.SetState(490)
			p.singleExpression(27)
		}

	case ECMAScriptParserMinusMinus:
		localctx = NewPreDecreaseExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(491)
			p.Match(ECMAScriptParserMinusMinus)
		}
		{
			p.SetState(492)
			p.singleExpression(26)
		}

	case ECMAScriptParserPlus:
		localctx = NewUnaryPlusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(493)
			p.Match(ECMAScriptParserPlus)
		}
		{
			p.SetState(494)
			p.singleExpression(25)
		}

	case ECMAScriptParserMinus:
		localctx = NewUnaryMinusExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(495)
			p.Match(ECMAScriptParserMinus)
		}
		{
			p.SetState(496)
			p.singleExpression(24)
		}

	case ECMAScriptParserBitNot:
		localctx = NewBitNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(497)
			p.Match(ECMAScriptParserBitNot)
		}
		{
			p.SetState(498)
			p.singleExpression(23)
		}

	case ECMAScriptParserNot:
		localctx = NewNotExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(499)
			p.Match(ECMAScriptParserNot)
		}
		{
			p.SetState(500)
			p.singleExpression(22)
		}

	case ECMAScriptParserThis:
		localctx = NewThisExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(501)
			p.Match(ECMAScriptParserThis)
		}

	case ECMAScriptParserIdentifier:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(502)
			p.Match(ECMAScriptParserIdentifier)
		}

	case ECMAScriptParserRegularExpressionLiteral, ECMAScriptParserNullLiteral, ECMAScriptParserBooleanLiteral, ECMAScriptParserDecimalLiteral, ECMAScriptParserHexIntegerLiteral, ECMAScriptParserOctalIntegerLiteral, ECMAScriptParserStringLiteral:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(503)
			p.Literal()
		}

	case ECMAScriptParserOpenBracket:
		localctx = NewArrayLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(504)
			p.ArrayLiteral()
		}

	case ECMAScriptParserOpenBrace:
		localctx = NewObjectLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(505)
			p.ObjectLiteral()
		}

	case ECMAScriptParserOpenParen:
		localctx = NewParenthesizedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(506)
			p.Match(ECMAScriptParserOpenParen)
		}
		{
			p.SetState(507)
			p.ExpressionSequence()
		}
		{
			p.SetState(508)
			p.Match(ECMAScriptParserCloseParen)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(579)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(577)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplicativeExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(512)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				p.SetState(513)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserMultiply)|(1<<ECMAScriptParserDivide)|(1<<ECMAScriptParserModulus))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(514)
					p.singleExpression(22)
				}

			case 2:
				localctx = NewAdditiveExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(515)

				if !(p.Precpred(p.GetParserRuleContext(), 20)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 20)", ""))
				}
				p.SetState(516)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ECMAScriptParserPlus || _la == ECMAScriptParserMinus) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(517)
					p.singleExpression(21)
				}

			case 3:
				localctx = NewBitShiftExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(518)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				p.SetState(519)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserRightShiftArithmetic)|(1<<ECMAScriptParserLeftShiftArithmetic)|(1<<ECMAScriptParserRightShiftLogical))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(520)
					p.singleExpression(20)
				}

			case 4:
				localctx = NewRelationalExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(521)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				p.SetState(522)
				_la = p.GetTokenStream().LA(1)

				if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ECMAScriptParserLessThan)|(1<<ECMAScriptParserMoreThan)|(1<<ECMAScriptParserLessThanEquals)|(1<<ECMAScriptParserGreaterThanEquals))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(523)
					p.singleExpression(19)
				}

			case 5:
				localctx = NewInstanceofExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(524)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(525)
					p.Match(ECMAScriptParserInstanceof)
				}
				{
					p.SetState(526)
					p.singleExpression(18)
				}

			case 6:
				localctx = NewInExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(527)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(528)
					p.Match(ECMAScriptParserIn)
				}
				{
					p.SetState(529)
					p.singleExpression(17)
				}

			case 7:
				localctx = NewEqualityExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(530)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				p.SetState(531)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(ECMAScriptParserEquals-31))|(1<<(ECMAScriptParserNotEquals-31))|(1<<(ECMAScriptParserIdentityEquals-31))|(1<<(ECMAScriptParserIdentityNotEquals-31)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
				{
					p.SetState(532)
					p.singleExpression(16)
				}

			case 8:
				localctx = NewBitAndExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(533)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(534)
					p.Match(ECMAScriptParserBitAnd)
				}
				{
					p.SetState(535)
					p.singleExpression(15)
				}

			case 9:
				localctx = NewBitXOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(536)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(537)
					p.Match(ECMAScriptParserBitXOr)
				}
				{
					p.SetState(538)
					p.singleExpression(14)
				}

			case 10:
				localctx = NewBitOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(539)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(540)
					p.Match(ECMAScriptParserBitOr)
				}
				{
					p.SetState(541)
					p.singleExpression(13)
				}

			case 11:
				localctx = NewLogicalAndExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(542)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(543)
					p.Match(ECMAScriptParserAnd)
				}
				{
					p.SetState(544)
					p.singleExpression(12)
				}

			case 12:
				localctx = NewLogicalOrExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(545)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(546)
					p.Match(ECMAScriptParserOr)
				}
				{
					p.SetState(547)
					p.singleExpression(11)
				}

			case 13:
				localctx = NewTernaryExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(548)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(549)
					p.Match(ECMAScriptParserQuestionMark)
				}
				{
					p.SetState(550)
					p.singleExpression(0)
				}
				{
					p.SetState(551)
					p.Match(ECMAScriptParserColon)
				}
				{
					p.SetState(552)
					p.singleExpression(10)
				}

			case 14:
				localctx = NewAssignmentExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(554)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(555)
					p.Match(ECMAScriptParserAssign)
				}
				{
					p.SetState(556)
					p.singleExpression(9)
				}

			case 15:
				localctx = NewAssignmentOperatorExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(557)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(558)
					p.AssignmentOperator()
				}
				{
					p.SetState(559)
					p.singleExpression(8)
				}

			case 16:
				localctx = NewMemberIndexExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(561)

				if !(p.Precpred(p.GetParserRuleContext(), 36)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 36)", ""))
				}
				{
					p.SetState(562)
					p.Match(ECMAScriptParserOpenBracket)
				}
				{
					p.SetState(563)
					p.ExpressionSequence()
				}
				{
					p.SetState(564)
					p.Match(ECMAScriptParserCloseBracket)
				}

			case 17:
				localctx = NewMemberDotExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(566)

				if !(p.Precpred(p.GetParserRuleContext(), 35)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 35)", ""))
				}
				{
					p.SetState(567)
					p.Match(ECMAScriptParserDot)
				}
				{
					p.SetState(568)
					p.IdentifierName()
				}

			case 18:
				localctx = NewArgumentsExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(569)

				if !(p.Precpred(p.GetParserRuleContext(), 34)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 34)", ""))
				}
				{
					p.SetState(570)
					p.Arguments()
				}

			case 19:
				localctx = NewPostIncrementExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(571)

				if !(p.Precpred(p.GetParserRuleContext(), 32)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 32)", ""))
				}
				p.SetState(572)

				if !(!p.here(ECMAScriptParserLineTerminator)) {
					panic(antlr.NewFailedPredicateException(p, "!p.here(ECMAScriptParserLineTerminator)", ""))
				}
				{
					p.SetState(573)
					p.Match(ECMAScriptParserPlusPlus)
				}

			case 20:
				localctx = NewPostDecreaseExpressionContext(p, NewSingleExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ECMAScriptParserRULE_singleExpression)
				p.SetState(574)

				if !(p.Precpred(p.GetParserRuleContext(), 31)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 31)", ""))
				}
				p.SetState(575)

				if !(!p.here(ECMAScriptParserLineTerminator)) {
					panic(antlr.NewFailedPredicateException(p, "!p.here(ECMAScriptParserLineTerminator)", ""))
				}
				{
					p.SetState(576)
					p.Match(ECMAScriptParserMinusMinus)
				}

			}

		}
		p.SetState(581)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
	}

	return localctx
}

// IAssignmentOperatorContext is an interface to support dynamic dispatch.
type IAssignmentOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentOperatorContext differentiates from other interfaces.
	IsAssignmentOperatorContext()
}

type AssignmentOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentOperatorContext() *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_assignmentOperator
	return p
}

func (*AssignmentOperatorContext) IsAssignmentOperatorContext() {}

func NewAssignmentOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentOperatorContext {
	var p = new(AssignmentOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_assignmentOperator

	return p
}

func (s *AssignmentOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *AssignmentOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterAssignmentOperator(s)
	}
}

func (s *AssignmentOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitAssignmentOperator(s)
	}
}

func (p *ECMAScriptParser) AssignmentOperator() (localctx IAssignmentOperatorContext) {
	localctx = NewAssignmentOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ECMAScriptParserRULE_assignmentOperator)
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
	p.SetState(582)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(ECMAScriptParserMultiplyAssign-40))|(1<<(ECMAScriptParserDivideAssign-40))|(1<<(ECMAScriptParserModulusAssign-40))|(1<<(ECMAScriptParserPlusAssign-40))|(1<<(ECMAScriptParserMinusAssign-40))|(1<<(ECMAScriptParserLeftShiftArithmeticAssign-40))|(1<<(ECMAScriptParserRightShiftArithmeticAssign-40))|(1<<(ECMAScriptParserRightShiftLogicalAssign-40))|(1<<(ECMAScriptParserBitAndAssign-40))|(1<<(ECMAScriptParserBitXorAssign-40))|(1<<(ECMAScriptParserBitOrAssign-40)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
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
	p.RuleIndex = ECMAScriptParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserNullLiteral, 0)
}

func (s *LiteralContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBooleanLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserStringLiteral, 0)
}

func (s *LiteralContext) RegularExpressionLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserRegularExpressionLiteral, 0)
}

func (s *LiteralContext) NumericLiteral() INumericLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *ECMAScriptParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ECMAScriptParserRULE_literal)
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

	p.SetState(586)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ECMAScriptParserRegularExpressionLiteral, ECMAScriptParserNullLiteral, ECMAScriptParserBooleanLiteral, ECMAScriptParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(584)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ECMAScriptParserRegularExpressionLiteral || _la == ECMAScriptParserNullLiteral || _la == ECMAScriptParserBooleanLiteral || _la == ECMAScriptParserStringLiteral) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	case ECMAScriptParserDecimalLiteral, ECMAScriptParserHexIntegerLiteral, ECMAScriptParserOctalIntegerLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(585)
			p.NumericLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_numericLiteral
	return p
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) DecimalLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDecimalLiteral, 0)
}

func (s *NumericLiteralContext) HexIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserHexIntegerLiteral, 0)
}

func (s *NumericLiteralContext) OctalIntegerLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserOctalIntegerLiteral, 0)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (p *ECMAScriptParser) NumericLiteral() (localctx INumericLiteralContext) {
	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ECMAScriptParserRULE_numericLiteral)
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
	p.SetState(588)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(ECMAScriptParserDecimalLiteral-53))|(1<<(ECMAScriptParserHexIntegerLiteral-53))|(1<<(ECMAScriptParserOctalIntegerLiteral-53)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IIdentifierNameContext is an interface to support dynamic dispatch.
type IIdentifierNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierNameContext differentiates from other interfaces.
	IsIdentifierNameContext()
}

type IdentifierNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierNameContext() *IdentifierNameContext {
	var p = new(IdentifierNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_identifierName
	return p
}

func (*IdentifierNameContext) IsIdentifierNameContext() {}

func NewIdentifierNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierNameContext {
	var p = new(IdentifierNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_identifierName

	return p
}

func (s *IdentifierNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *IdentifierNameContext) ReservedWord() IReservedWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReservedWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReservedWordContext)
}

func (s *IdentifierNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterIdentifierName(s)
	}
}

func (s *IdentifierNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitIdentifierName(s)
	}
}

func (p *ECMAScriptParser) IdentifierName() (localctx IIdentifierNameContext) {
	localctx = NewIdentifierNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ECMAScriptParserRULE_identifierName)

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

	p.SetState(592)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ECMAScriptParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(590)
			p.Match(ECMAScriptParserIdentifier)
		}

	case ECMAScriptParserNullLiteral, ECMAScriptParserBooleanLiteral, ECMAScriptParserBreak, ECMAScriptParserDo, ECMAScriptParserInstanceof, ECMAScriptParserTypeof, ECMAScriptParserCase, ECMAScriptParserElse, ECMAScriptParserNew, ECMAScriptParserVar, ECMAScriptParserCatch, ECMAScriptParserFinally, ECMAScriptParserReturn, ECMAScriptParserVoid, ECMAScriptParserContinue, ECMAScriptParserFor, ECMAScriptParserSwitch, ECMAScriptParserWhile, ECMAScriptParserDebugger, ECMAScriptParserFunction, ECMAScriptParserThis, ECMAScriptParserWith, ECMAScriptParserDefault, ECMAScriptParserIf, ECMAScriptParserThrow, ECMAScriptParserDelete, ECMAScriptParserIn, ECMAScriptParserTry, ECMAScriptParserClass, ECMAScriptParserEnum, ECMAScriptParserExtends, ECMAScriptParserSuper, ECMAScriptParserConst, ECMAScriptParserExport, ECMAScriptParserImport, ECMAScriptParserImplements, ECMAScriptParserLet, ECMAScriptParserPrivate, ECMAScriptParserPublic, ECMAScriptParserInterface, ECMAScriptParserPackage, ECMAScriptParserProtected, ECMAScriptParserStatic, ECMAScriptParserYield:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(591)
			p.ReservedWord()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReservedWordContext is an interface to support dynamic dispatch.
type IReservedWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReservedWordContext differentiates from other interfaces.
	IsReservedWordContext()
}

type ReservedWordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedWordContext() *ReservedWordContext {
	var p = new(ReservedWordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_reservedWord
	return p
}

func (*ReservedWordContext) IsReservedWordContext() {}

func NewReservedWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedWordContext {
	var p = new(ReservedWordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_reservedWord

	return p
}

func (s *ReservedWordContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedWordContext) Keyword() IKeywordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *ReservedWordContext) FutureReservedWord() IFutureReservedWordContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFutureReservedWordContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFutureReservedWordContext)
}

func (s *ReservedWordContext) NullLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserNullLiteral, 0)
}

func (s *ReservedWordContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBooleanLiteral, 0)
}

func (s *ReservedWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedWordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedWordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterReservedWord(s)
	}
}

func (s *ReservedWordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitReservedWord(s)
	}
}

func (p *ECMAScriptParser) ReservedWord() (localctx IReservedWordContext) {
	localctx = NewReservedWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, ECMAScriptParserRULE_reservedWord)
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

	p.SetState(597)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ECMAScriptParserBreak, ECMAScriptParserDo, ECMAScriptParserInstanceof, ECMAScriptParserTypeof, ECMAScriptParserCase, ECMAScriptParserElse, ECMAScriptParserNew, ECMAScriptParserVar, ECMAScriptParserCatch, ECMAScriptParserFinally, ECMAScriptParserReturn, ECMAScriptParserVoid, ECMAScriptParserContinue, ECMAScriptParserFor, ECMAScriptParserSwitch, ECMAScriptParserWhile, ECMAScriptParserDebugger, ECMAScriptParserFunction, ECMAScriptParserThis, ECMAScriptParserWith, ECMAScriptParserDefault, ECMAScriptParserIf, ECMAScriptParserThrow, ECMAScriptParserDelete, ECMAScriptParserIn, ECMAScriptParserTry:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(594)
			p.Keyword()
		}

	case ECMAScriptParserClass, ECMAScriptParserEnum, ECMAScriptParserExtends, ECMAScriptParserSuper, ECMAScriptParserConst, ECMAScriptParserExport, ECMAScriptParserImport, ECMAScriptParserImplements, ECMAScriptParserLet, ECMAScriptParserPrivate, ECMAScriptParserPublic, ECMAScriptParserInterface, ECMAScriptParserPackage, ECMAScriptParserProtected, ECMAScriptParserStatic, ECMAScriptParserYield:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(595)
			p.FutureReservedWord()
		}

	case ECMAScriptParserNullLiteral, ECMAScriptParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(596)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ECMAScriptParserNullLiteral || _la == ECMAScriptParserBooleanLiteral) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = ECMAScriptParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) Break() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserBreak, 0)
}

func (s *KeywordContext) Do() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDo, 0)
}

func (s *KeywordContext) Instanceof() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserInstanceof, 0)
}

func (s *KeywordContext) Typeof() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserTypeof, 0)
}

func (s *KeywordContext) Case() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCase, 0)
}

func (s *KeywordContext) Else() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserElse, 0)
}

func (s *KeywordContext) New() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserNew, 0)
}

func (s *KeywordContext) Var() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserVar, 0)
}

func (s *KeywordContext) Catch() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserCatch, 0)
}

func (s *KeywordContext) Finally() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFinally, 0)
}

func (s *KeywordContext) Return() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserReturn, 0)
}

func (s *KeywordContext) Void() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserVoid, 0)
}

func (s *KeywordContext) Continue() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserContinue, 0)
}

func (s *KeywordContext) For() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFor, 0)
}

func (s *KeywordContext) Switch() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserSwitch, 0)
}

func (s *KeywordContext) While() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserWhile, 0)
}

func (s *KeywordContext) Debugger() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDebugger, 0)
}

func (s *KeywordContext) Function() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserFunction, 0)
}

func (s *KeywordContext) This() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserThis, 0)
}

func (s *KeywordContext) With() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserWith, 0)
}

func (s *KeywordContext) Default() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDefault, 0)
}

func (s *KeywordContext) If() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIf, 0)
}

func (s *KeywordContext) Throw() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserThrow, 0)
}

func (s *KeywordContext) Delete() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserDelete, 0)
}

func (s *KeywordContext) In() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIn, 0)
}

func (s *KeywordContext) Try() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserTry, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitKeyword(s)
	}
}

func (p *ECMAScriptParser) Keyword() (localctx IKeywordContext) {
	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, ECMAScriptParserRULE_keyword)
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
	p.SetState(599)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-56)&-(0x1f+1)) == 0 && ((1<<uint((_la-56)))&((1<<(ECMAScriptParserBreak-56))|(1<<(ECMAScriptParserDo-56))|(1<<(ECMAScriptParserInstanceof-56))|(1<<(ECMAScriptParserTypeof-56))|(1<<(ECMAScriptParserCase-56))|(1<<(ECMAScriptParserElse-56))|(1<<(ECMAScriptParserNew-56))|(1<<(ECMAScriptParserVar-56))|(1<<(ECMAScriptParserCatch-56))|(1<<(ECMAScriptParserFinally-56))|(1<<(ECMAScriptParserReturn-56))|(1<<(ECMAScriptParserVoid-56))|(1<<(ECMAScriptParserContinue-56))|(1<<(ECMAScriptParserFor-56))|(1<<(ECMAScriptParserSwitch-56))|(1<<(ECMAScriptParserWhile-56))|(1<<(ECMAScriptParserDebugger-56))|(1<<(ECMAScriptParserFunction-56))|(1<<(ECMAScriptParserThis-56))|(1<<(ECMAScriptParserWith-56))|(1<<(ECMAScriptParserDefault-56))|(1<<(ECMAScriptParserIf-56))|(1<<(ECMAScriptParserThrow-56))|(1<<(ECMAScriptParserDelete-56))|(1<<(ECMAScriptParserIn-56))|(1<<(ECMAScriptParserTry-56)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IFutureReservedWordContext is an interface to support dynamic dispatch.
type IFutureReservedWordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFutureReservedWordContext differentiates from other interfaces.
	IsFutureReservedWordContext()
}

type FutureReservedWordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFutureReservedWordContext() *FutureReservedWordContext {
	var p = new(FutureReservedWordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_futureReservedWord
	return p
}

func (*FutureReservedWordContext) IsFutureReservedWordContext() {}

func NewFutureReservedWordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FutureReservedWordContext {
	var p = new(FutureReservedWordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_futureReservedWord

	return p
}

func (s *FutureReservedWordContext) GetParser() antlr.Parser { return s.parser }

func (s *FutureReservedWordContext) Class() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserClass, 0)
}

func (s *FutureReservedWordContext) Enum() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserEnum, 0)
}

func (s *FutureReservedWordContext) Extends() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserExtends, 0)
}

func (s *FutureReservedWordContext) Super() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserSuper, 0)
}

func (s *FutureReservedWordContext) Const() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserConst, 0)
}

func (s *FutureReservedWordContext) Export() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserExport, 0)
}

func (s *FutureReservedWordContext) Import() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserImport, 0)
}

func (s *FutureReservedWordContext) Implements() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserImplements, 0)
}

func (s *FutureReservedWordContext) Let() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserLet, 0)
}

func (s *FutureReservedWordContext) Private() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserPrivate, 0)
}

func (s *FutureReservedWordContext) Public() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserPublic, 0)
}

func (s *FutureReservedWordContext) Interface() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserInterface, 0)
}

func (s *FutureReservedWordContext) Package() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserPackage, 0)
}

func (s *FutureReservedWordContext) Protected() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserProtected, 0)
}

func (s *FutureReservedWordContext) Static() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserStatic, 0)
}

func (s *FutureReservedWordContext) Yield() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserYield, 0)
}

func (s *FutureReservedWordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FutureReservedWordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FutureReservedWordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterFutureReservedWord(s)
	}
}

func (s *FutureReservedWordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitFutureReservedWord(s)
	}
}

func (p *ECMAScriptParser) FutureReservedWord() (localctx IFutureReservedWordContext) {
	localctx = NewFutureReservedWordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, ECMAScriptParserRULE_futureReservedWord)
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
	p.SetState(601)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-82)&-(0x1f+1)) == 0 && ((1<<uint((_la-82)))&((1<<(ECMAScriptParserClass-82))|(1<<(ECMAScriptParserEnum-82))|(1<<(ECMAScriptParserExtends-82))|(1<<(ECMAScriptParserSuper-82))|(1<<(ECMAScriptParserConst-82))|(1<<(ECMAScriptParserExport-82))|(1<<(ECMAScriptParserImport-82))|(1<<(ECMAScriptParserImplements-82))|(1<<(ECMAScriptParserLet-82))|(1<<(ECMAScriptParserPrivate-82))|(1<<(ECMAScriptParserPublic-82))|(1<<(ECMAScriptParserInterface-82))|(1<<(ECMAScriptParserPackage-82))|(1<<(ECMAScriptParserProtected-82))|(1<<(ECMAScriptParserStatic-82))|(1<<(ECMAScriptParserYield-82)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IGetterContext is an interface to support dynamic dispatch.
type IGetterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGetterContext differentiates from other interfaces.
	IsGetterContext()
}

type GetterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGetterContext() *GetterContext {
	var p = new(GetterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_getter
	return p
}

func (*GetterContext) IsGetterContext() {}

func NewGetterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GetterContext {
	var p = new(GetterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_getter

	return p
}

func (s *GetterContext) GetParser() antlr.Parser { return s.parser }

func (s *GetterContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *GetterContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *GetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterGetter(s)
	}
}

func (s *GetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitGetter(s)
	}
}

func (p *ECMAScriptParser) Getter() (localctx IGetterContext) {
	localctx = NewGetterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, ECMAScriptParserRULE_getter)

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
	p.SetState(603)

	if !(p.GetTokenStream().LT(1).GetText() == "get") {
		panic(antlr.NewFailedPredicateException(p, "p.GetTokenStream().LT(1).GetText() == \"get\"", ""))
	}
	{
		p.SetState(604)
		p.Match(ECMAScriptParserIdentifier)
	}
	{
		p.SetState(605)
		p.PropertyName()
	}

	return localctx
}

// ISetterContext is an interface to support dynamic dispatch.
type ISetterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetterContext differentiates from other interfaces.
	IsSetterContext()
}

type SetterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetterContext() *SetterContext {
	var p = new(SetterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_setter
	return p
}

func (*SetterContext) IsSetterContext() {}

func NewSetterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetterContext {
	var p = new(SetterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_setter

	return p
}

func (s *SetterContext) GetParser() antlr.Parser { return s.parser }

func (s *SetterContext) Identifier() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserIdentifier, 0)
}

func (s *SetterContext) PropertyName() IPropertyNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyNameContext)
}

func (s *SetterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterSetter(s)
	}
}

func (s *SetterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitSetter(s)
	}
}

func (p *ECMAScriptParser) Setter() (localctx ISetterContext) {
	localctx = NewSetterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, ECMAScriptParserRULE_setter)

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
	p.SetState(607)

	if !(p.GetTokenStream().LT(1).GetText() == "set") {
		panic(antlr.NewFailedPredicateException(p, "p.GetTokenStream().LT(1).GetText() == \"set\"", ""))
	}
	{
		p.SetState(608)
		p.Match(ECMAScriptParserIdentifier)
	}
	{
		p.SetState(609)
		p.PropertyName()
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserSemiColon, 0)
}

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserEOF, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitEos(s)
	}
}

func (p *ECMAScriptParser) Eos() (localctx IEosContext) {
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, ECMAScriptParserRULE_eos)

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

	p.SetState(615)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(611)
			p.Match(ECMAScriptParserSemiColon)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(612)
			p.Match(ECMAScriptParserEOF)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(613)

		if !(p.lineTerminatorAhead()) {
			panic(antlr.NewFailedPredicateException(p, "p.lineTerminatorAhead()", ""))
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(614)

		if !(p.GetTokenStream().LT(1).GetTokenType() == ECMAScriptParserCloseBrace) {
			panic(antlr.NewFailedPredicateException(p, "p.GetTokenStream().LT(1).GetTokenType() == ECMAScriptParserCloseBrace", ""))
		}

	}

	return localctx
}

// IEofContext is an interface to support dynamic dispatch.
type IEofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEofContext differentiates from other interfaces.
	IsEofContext()
}

type EofContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEofContext() *EofContext {
	var p = new(EofContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ECMAScriptParserRULE_eof
	return p
}

func (*EofContext) IsEofContext() {}

func NewEofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EofContext {
	var p = new(EofContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ECMAScriptParserRULE_eof

	return p
}

func (s *EofContext) GetParser() antlr.Parser { return s.parser }

func (s *EofContext) EOF() antlr.TerminalNode {
	return s.GetToken(ECMAScriptParserEOF, 0)
}

func (s *EofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EofContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.EnterEof(s)
	}
}

func (s *EofContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ECMAScriptListener); ok {
		listenerT.ExitEof(s)
	}
}

func (p *ECMAScriptParser) Eof() (localctx IEofContext) {
	localctx = NewEofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, ECMAScriptParserRULE_eof)

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
		p.SetState(617)
		p.Match(ECMAScriptParserEOF)
	}

	return localctx
}

func (p *ECMAScriptParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 11:
		var t *ExpressionStatementContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionStatementContext)
		}
		return p.ExpressionStatement_Sempred(t, predIndex)

	case 43:
		var t *SingleExpressionContext = nil
		if localctx != nil {
			t = localctx.(*SingleExpressionContext)
		}
		return p.SingleExpression_Sempred(t, predIndex)

	case 51:
		var t *GetterContext = nil
		if localctx != nil {
			t = localctx.(*GetterContext)
		}
		return p.Getter_Sempred(t, predIndex)

	case 52:
		var t *SetterContext = nil
		if localctx != nil {
			t = localctx.(*SetterContext)
		}
		return p.Setter_Sempred(t, predIndex)

	case 53:
		var t *EosContext = nil
		if localctx != nil {
			t = localctx.(*EosContext)
		}
		return p.Eos_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ECMAScriptParser) ExpressionStatement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return (p.GetInputStream().LA(1) != ECMAScriptParserOpenBrace) && (p.GetInputStream().LA(1) != ECMAScriptParserFunction)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ECMAScriptParser) SingleExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 21)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 20)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 14:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 15:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 16:
		return p.Precpred(p.GetParserRuleContext(), 36)

	case 17:
		return p.Precpred(p.GetParserRuleContext(), 35)

	case 18:
		return p.Precpred(p.GetParserRuleContext(), 34)

	case 19:
		return p.Precpred(p.GetParserRuleContext(), 32)

	case 20:
		return !p.here(ECMAScriptParserLineTerminator)

	case 21:
		return p.Precpred(p.GetParserRuleContext(), 31)

	case 22:
		return !p.here(ECMAScriptParserLineTerminator)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ECMAScriptParser) Getter_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 23:
		return p.GetTokenStream().LT(1).GetText() == "get"

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ECMAScriptParser) Setter_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 24:
		return p.GetTokenStream().LT(1).GetText() == "set"

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *ECMAScriptParser) Eos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 25:
		return p.lineTerminatorAhead()

	case 26:
		return p.GetTokenStream().LT(1).GetTokenType() == ECMAScriptParserCloseBrace

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
