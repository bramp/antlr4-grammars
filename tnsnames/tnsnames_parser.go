// Generated from tnsnamesParser.g4 by ANTLR 4.7.

package tnsnames // tnsnamesParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 86, 784,
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
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 70, 4,
	71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75, 4, 76,
	9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4, 81, 9,
	81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86, 9, 86,
	4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 3, 2, 3, 2, 3,
	2, 7, 2, 184, 10, 2, 12, 2, 14, 2, 187, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3,
	5, 3, 193, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	6, 5, 204, 10, 5, 13, 5, 14, 5, 205, 5, 5, 208, 10, 5, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 6, 6, 215, 10, 6, 13, 6, 14, 6, 216, 5, 6, 219, 10, 6, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 7, 7, 226, 10, 7, 12, 7, 14, 7, 229, 11, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 6, 8, 235, 10, 8, 13, 8, 14, 8, 236, 5, 8, 239, 10,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 245, 10, 9, 3, 9, 6, 9, 248, 10, 9, 13,
	9, 14, 9, 249, 3, 9, 5, 9, 253, 10, 9, 3, 9, 3, 9, 3, 10, 6, 10, 258, 10,
	10, 13, 10, 14, 10, 259, 3, 11, 3, 11, 3, 11, 5, 11, 265, 10, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 5, 12, 271, 10, 12, 3, 12, 3, 12, 6, 12, 275, 10,
	12, 13, 12, 14, 12, 276, 5, 12, 279, 10, 12, 3, 12, 5, 12, 282, 10, 12,
	3, 12, 3, 12, 5, 12, 286, 10, 12, 3, 12, 3, 12, 3, 13, 6, 13, 291, 10,
	13, 13, 13, 14, 13, 292, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 307, 10, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 373,
	10, 25, 3, 25, 6, 25, 376, 10, 25, 13, 25, 14, 25, 377, 3, 25, 5, 25, 381,
	10, 25, 3, 25, 3, 25, 3, 26, 6, 26, 386, 10, 26, 13, 26, 14, 26, 387, 3,
	27, 3, 27, 3, 27, 5, 27, 393, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 418, 10, 31,
	3, 31, 3, 31, 3, 32, 6, 32, 423, 10, 32, 13, 32, 14, 32, 424, 3, 33, 3,
	33, 5, 33, 429, 10, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 436,
	10, 34, 3, 35, 3, 35, 3, 36, 6, 36, 441, 10, 36, 13, 36, 14, 36, 442, 3,
	37, 3, 37, 3, 37, 5, 37, 448, 10, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3,
	40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 6, 41, 472, 10, 41, 13, 41,
	14, 41, 473, 3, 41, 5, 41, 477, 10, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3,
	44, 6, 44, 484, 10, 44, 13, 44, 14, 44, 485, 3, 45, 3, 45, 5, 45, 490,
	10, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47,
	3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 6, 49, 507, 10, 49, 13, 49, 14,
	49, 508, 3, 50, 3, 50, 5, 50, 513, 10, 50, 3, 51, 3, 51, 3, 51, 3, 51,
	3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3,
	54, 6, 54, 530, 10, 54, 13, 54, 14, 54, 531, 3, 55, 3, 55, 3, 55, 5, 55,
	537, 10, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3,
	57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 59,
	3, 59, 3, 60, 6, 60, 560, 10, 60, 13, 60, 14, 60, 561, 3, 61, 3, 61, 3,
	61, 3, 61, 5, 61, 568, 10, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62, 3, 62,
	3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 64, 3,
	64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66,
	3, 66, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 67, 3, 68, 6, 68, 605, 10,
	68, 13, 68, 14, 68, 606, 3, 69, 3, 69, 5, 69, 611, 10, 69, 3, 70, 3, 70,
	3, 70, 3, 70, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3,
	72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 72, 3, 73, 6, 73, 632, 10, 73, 13, 73,
	14, 73, 633, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 3,
	74, 5, 74, 645, 10, 74, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 7, 75,
	653, 10, 75, 12, 75, 14, 75, 656, 11, 75, 3, 75, 3, 75, 3, 76, 3, 76, 3,
	76, 3, 76, 3, 76, 3, 76, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 3, 77, 7, 77,
	672, 10, 77, 12, 77, 14, 77, 675, 11, 77, 3, 77, 3, 77, 3, 78, 3, 78, 3,
	78, 3, 78, 3, 78, 3, 78, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79, 7, 79,
	691, 10, 79, 12, 79, 14, 79, 694, 11, 79, 3, 79, 3, 79, 3, 80, 3, 80, 3,
	80, 3, 80, 3, 80, 3, 80, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81, 3, 81,
	5, 81, 711, 10, 81, 3, 81, 3, 81, 3, 81, 7, 81, 716, 10, 81, 12, 81, 14,
	81, 719, 11, 81, 3, 81, 3, 81, 3, 82, 3, 82, 3, 82, 3, 82, 3, 82, 3, 82,
	3, 83, 3, 83, 3, 83, 3, 83, 3, 83, 3, 83, 3, 84, 6, 84, 736, 10, 84, 13,
	84, 14, 84, 737, 3, 85, 3, 85, 3, 85, 3, 85, 3, 85, 5, 85, 745, 10, 85,
	3, 86, 3, 86, 3, 86, 3, 86, 3, 86, 3, 86, 3, 87, 3, 87, 3, 87, 3, 87, 3,
	87, 3, 87, 7, 87, 759, 10, 87, 12, 87, 14, 87, 762, 11, 87, 3, 87, 3, 87,
	3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 89, 3, 89, 3, 89, 3, 89, 3,
	89, 3, 89, 3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 2, 2, 91, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
	42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76,
	78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110,
	112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140,
	142, 144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166, 168, 170,
	172, 174, 176, 178, 2, 7, 3, 2, 23, 25, 3, 2, 23, 24, 3, 2, 29, 31, 3,
	2, 70, 72, 3, 2, 74, 75, 2, 781, 2, 185, 3, 2, 2, 2, 4, 188, 3, 2, 2, 2,
	6, 194, 3, 2, 2, 2, 8, 198, 3, 2, 2, 2, 10, 209, 3, 2, 2, 2, 12, 222, 3,
	2, 2, 2, 14, 238, 3, 2, 2, 2, 16, 240, 3, 2, 2, 2, 18, 257, 3, 2, 2, 2,
	20, 264, 3, 2, 2, 2, 22, 266, 3, 2, 2, 2, 24, 290, 3, 2, 2, 2, 26, 306,
	3, 2, 2, 2, 28, 308, 3, 2, 2, 2, 30, 314, 3, 2, 2, 2, 32, 320, 3, 2, 2,
	2, 34, 326, 3, 2, 2, 2, 36, 332, 3, 2, 2, 2, 38, 338, 3, 2, 2, 2, 40, 344,
	3, 2, 2, 2, 42, 350, 3, 2, 2, 2, 44, 356, 3, 2, 2, 2, 46, 362, 3, 2, 2,
	2, 48, 368, 3, 2, 2, 2, 50, 385, 3, 2, 2, 2, 52, 392, 3, 2, 2, 2, 54, 394,
	3, 2, 2, 2, 56, 400, 3, 2, 2, 2, 58, 406, 3, 2, 2, 2, 60, 412, 3, 2, 2,
	2, 62, 422, 3, 2, 2, 2, 64, 428, 3, 2, 2, 2, 66, 435, 3, 2, 2, 2, 68, 437,
	3, 2, 2, 2, 70, 440, 3, 2, 2, 2, 72, 447, 3, 2, 2, 2, 74, 449, 3, 2, 2,
	2, 76, 455, 3, 2, 2, 2, 78, 461, 3, 2, 2, 2, 80, 476, 3, 2, 2, 2, 82, 478,
	3, 2, 2, 2, 84, 480, 3, 2, 2, 2, 86, 483, 3, 2, 2, 2, 88, 489, 3, 2, 2,
	2, 90, 491, 3, 2, 2, 2, 92, 497, 3, 2, 2, 2, 94, 503, 3, 2, 2, 2, 96, 506,
	3, 2, 2, 2, 98, 512, 3, 2, 2, 2, 100, 514, 3, 2, 2, 2, 102, 520, 3, 2,
	2, 2, 104, 526, 3, 2, 2, 2, 106, 529, 3, 2, 2, 2, 108, 536, 3, 2, 2, 2,
	110, 538, 3, 2, 2, 2, 112, 544, 3, 2, 2, 2, 114, 550, 3, 2, 2, 2, 116,
	556, 3, 2, 2, 2, 118, 559, 3, 2, 2, 2, 120, 567, 3, 2, 2, 2, 122, 569,
	3, 2, 2, 2, 124, 575, 3, 2, 2, 2, 126, 581, 3, 2, 2, 2, 128, 587, 3, 2,
	2, 2, 130, 593, 3, 2, 2, 2, 132, 597, 3, 2, 2, 2, 134, 604, 3, 2, 2, 2,
	136, 610, 3, 2, 2, 2, 138, 612, 3, 2, 2, 2, 140, 618, 3, 2, 2, 2, 142,
	624, 3, 2, 2, 2, 144, 631, 3, 2, 2, 2, 146, 644, 3, 2, 2, 2, 148, 646,
	3, 2, 2, 2, 150, 659, 3, 2, 2, 2, 152, 665, 3, 2, 2, 2, 154, 678, 3, 2,
	2, 2, 156, 684, 3, 2, 2, 2, 158, 697, 3, 2, 2, 2, 160, 703, 3, 2, 2, 2,
	162, 722, 3, 2, 2, 2, 164, 728, 3, 2, 2, 2, 166, 735, 3, 2, 2, 2, 168,
	744, 3, 2, 2, 2, 170, 746, 3, 2, 2, 2, 172, 752, 3, 2, 2, 2, 174, 765,
	3, 2, 2, 2, 176, 771, 3, 2, 2, 2, 178, 777, 3, 2, 2, 2, 180, 184, 5, 4,
	3, 2, 181, 184, 5, 6, 4, 2, 182, 184, 5, 8, 5, 2, 183, 180, 3, 2, 2, 2,
	183, 181, 3, 2, 2, 2, 183, 182, 3, 2, 2, 2, 184, 187, 3, 2, 2, 2, 185,
	183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 3, 3, 2, 2, 2, 187, 185, 3,
	2, 2, 2, 188, 189, 5, 12, 7, 2, 189, 192, 7, 7, 2, 2, 190, 193, 5, 16,
	9, 2, 191, 193, 5, 22, 12, 2, 192, 190, 3, 2, 2, 2, 192, 191, 3, 2, 2,
	2, 193, 5, 3, 2, 2, 2, 194, 195, 7, 58, 2, 2, 195, 196, 7, 81, 2, 2, 196,
	197, 7, 82, 2, 2, 197, 7, 3, 2, 2, 2, 198, 199, 5, 14, 8, 2, 199, 207,
	7, 7, 2, 2, 200, 208, 5, 10, 6, 2, 201, 208, 5, 48, 25, 2, 202, 204, 5,
	60, 31, 2, 203, 202, 3, 2, 2, 2, 204, 205, 3, 2, 2, 2, 205, 203, 3, 2,
	2, 2, 205, 206, 3, 2, 2, 2, 206, 208, 3, 2, 2, 2, 207, 200, 3, 2, 2, 2,
	207, 201, 3, 2, 2, 2, 207, 203, 3, 2, 2, 2, 208, 9, 3, 2, 2, 2, 209, 210,
	7, 3, 2, 2, 210, 211, 7, 14, 2, 2, 211, 218, 7, 7, 2, 2, 212, 219, 5, 48,
	25, 2, 213, 215, 5, 60, 31, 2, 214, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2,
	2, 216, 214, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217, 219, 3, 2, 2, 2, 218,
	212, 3, 2, 2, 2, 218, 214, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221,
	7, 4, 2, 2, 221, 11, 3, 2, 2, 2, 222, 227, 5, 14, 8, 2, 223, 224, 7, 9,
	2, 2, 224, 226, 5, 14, 8, 2, 225, 223, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2,
	227, 225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 13, 3, 2, 2, 2, 229, 227,
	3, 2, 2, 2, 230, 239, 7, 79, 2, 2, 231, 234, 7, 79, 2, 2, 232, 233, 7,
	8, 2, 2, 233, 235, 7, 79, 2, 2, 234, 232, 3, 2, 2, 2, 235, 236, 3, 2, 2,
	2, 236, 234, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 239, 3, 2, 2, 2, 238,
	230, 3, 2, 2, 2, 238, 231, 3, 2, 2, 2, 239, 15, 3, 2, 2, 2, 240, 241, 7,
	3, 2, 2, 241, 242, 7, 13, 2, 2, 242, 244, 7, 7, 2, 2, 243, 245, 5, 18,
	10, 2, 244, 243, 3, 2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 247, 3, 2, 2, 2,
	246, 248, 5, 22, 12, 2, 247, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249,
	247, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250, 252, 3, 2, 2, 2, 251, 253,
	5, 18, 10, 2, 252, 251, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 254, 3,
	2, 2, 2, 254, 255, 7, 4, 2, 2, 255, 17, 3, 2, 2, 2, 256, 258, 5, 20, 11,
	2, 257, 256, 3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 259,
	260, 3, 2, 2, 2, 260, 19, 3, 2, 2, 2, 261, 265, 5, 54, 28, 2, 262, 265,
	5, 56, 29, 2, 263, 265, 5, 58, 30, 2, 264, 261, 3, 2, 2, 2, 264, 262, 3,
	2, 2, 2, 264, 263, 3, 2, 2, 2, 265, 21, 3, 2, 2, 2, 266, 267, 7, 3, 2,
	2, 267, 268, 7, 14, 2, 2, 268, 270, 7, 7, 2, 2, 269, 271, 5, 24, 13, 2,
	270, 269, 3, 2, 2, 2, 270, 271, 3, 2, 2, 2, 271, 278, 3, 2, 2, 2, 272,
	279, 5, 48, 25, 2, 273, 275, 5, 60, 31, 2, 274, 273, 3, 2, 2, 2, 275, 276,
	3, 2, 2, 2, 276, 274, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 279, 3, 2,
	2, 2, 278, 272, 3, 2, 2, 2, 278, 274, 3, 2, 2, 2, 279, 281, 3, 2, 2, 2,
	280, 282, 5, 24, 13, 2, 281, 280, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282,
	283, 3, 2, 2, 2, 283, 285, 5, 142, 72, 2, 284, 286, 5, 24, 13, 2, 285,
	284, 3, 2, 2, 2, 285, 286, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 288,
	7, 4, 2, 2, 288, 23, 3, 2, 2, 2, 289, 291, 5, 26, 14, 2, 290, 289, 3, 2,
	2, 2, 291, 292, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2,
	293, 25, 3, 2, 2, 2, 294, 307, 5, 28, 15, 2, 295, 307, 5, 54, 28, 2, 296,
	307, 5, 56, 29, 2, 297, 307, 5, 30, 16, 2, 298, 307, 5, 32, 17, 2, 299,
	307, 5, 34, 18, 2, 300, 307, 5, 58, 30, 2, 301, 307, 5, 36, 19, 2, 302,
	307, 5, 38, 20, 2, 303, 307, 5, 40, 21, 2, 304, 307, 5, 42, 22, 2, 305,
	307, 5, 44, 23, 2, 306, 294, 3, 2, 2, 2, 306, 295, 3, 2, 2, 2, 306, 296,
	3, 2, 2, 2, 306, 297, 3, 2, 2, 2, 306, 298, 3, 2, 2, 2, 306, 299, 3, 2,
	2, 2, 306, 300, 3, 2, 2, 2, 306, 301, 3, 2, 2, 2, 306, 302, 3, 2, 2, 2,
	306, 303, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306, 305, 3, 2, 2, 2, 307,
	27, 3, 2, 2, 2, 308, 309, 7, 3, 2, 2, 309, 310, 7, 36, 2, 2, 310, 311,
	7, 7, 2, 2, 311, 312, 7, 37, 2, 2, 312, 313, 7, 4, 2, 2, 313, 29, 3, 2,
	2, 2, 314, 315, 7, 3, 2, 2, 315, 316, 7, 38, 2, 2, 316, 317, 7, 7, 2, 2,
	317, 318, 7, 27, 2, 2, 318, 319, 7, 4, 2, 2, 319, 31, 3, 2, 2, 2, 320,
	321, 7, 3, 2, 2, 321, 322, 7, 39, 2, 2, 322, 323, 7, 7, 2, 2, 323, 324,
	7, 27, 2, 2, 324, 325, 7, 4, 2, 2, 325, 33, 3, 2, 2, 2, 326, 327, 7, 3,
	2, 2, 327, 328, 7, 40, 2, 2, 328, 329, 7, 7, 2, 2, 329, 330, 7, 27, 2,
	2, 330, 331, 7, 4, 2, 2, 331, 35, 3, 2, 2, 2, 332, 333, 7, 3, 2, 2, 333,
	334, 7, 43, 2, 2, 334, 335, 7, 7, 2, 2, 335, 336, 7, 79, 2, 2, 336, 337,
	7, 4, 2, 2, 337, 37, 3, 2, 2, 2, 338, 339, 7, 3, 2, 2, 339, 340, 7, 53,
	2, 2, 340, 341, 7, 7, 2, 2, 341, 342, 5, 46, 24, 2, 342, 343, 7, 4, 2,
	2, 343, 39, 3, 2, 2, 2, 344, 345, 7, 3, 2, 2, 345, 346, 7, 55, 2, 2, 346,
	347, 7, 7, 2, 2, 347, 348, 7, 27, 2, 2, 348, 349, 7, 4, 2, 2, 349, 41,
	3, 2, 2, 2, 350, 351, 7, 3, 2, 2, 351, 352, 7, 56, 2, 2, 352, 353, 7, 7,
	2, 2, 353, 354, 7, 27, 2, 2, 354, 355, 7, 4, 2, 2, 355, 43, 3, 2, 2, 2,
	356, 357, 7, 3, 2, 2, 357, 358, 7, 57, 2, 2, 358, 359, 7, 7, 2, 2, 359,
	360, 7, 27, 2, 2, 360, 361, 7, 4, 2, 2, 361, 45, 3, 2, 2, 2, 362, 363,
	7, 3, 2, 2, 363, 364, 7, 54, 2, 2, 364, 365, 7, 7, 2, 2, 365, 366, 7, 59,
	2, 2, 366, 367, 7, 4, 2, 2, 367, 47, 3, 2, 2, 2, 368, 369, 7, 3, 2, 2,
	369, 370, 7, 15, 2, 2, 370, 372, 7, 7, 2, 2, 371, 373, 5, 50, 26, 2, 372,
	371, 3, 2, 2, 2, 372, 373, 3, 2, 2, 2, 373, 375, 3, 2, 2, 2, 374, 376,
	5, 60, 31, 2, 375, 374, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 375, 3,
	2, 2, 2, 377, 378, 3, 2, 2, 2, 378, 380, 3, 2, 2, 2, 379, 381, 5, 50, 26,
	2, 380, 379, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382,
	383, 7, 4, 2, 2, 383, 49, 3, 2, 2, 2, 384, 386, 5, 52, 27, 2, 385, 384,
	3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 385, 3, 2, 2, 2, 387, 388, 3, 2,
	2, 2, 388, 51, 3, 2, 2, 2, 389, 393, 5, 54, 28, 2, 390, 393, 5, 56, 29,
	2, 391, 393, 5, 58, 30, 2, 392, 389, 3, 2, 2, 2, 392, 390, 3, 2, 2, 2,
	392, 391, 3, 2, 2, 2, 393, 53, 3, 2, 2, 2, 394, 395, 7, 3, 2, 2, 395, 396,
	7, 33, 2, 2, 396, 397, 7, 7, 2, 2, 397, 398, 9, 2, 2, 2, 398, 399, 7, 4,
	2, 2, 399, 55, 3, 2, 2, 2, 400, 401, 7, 3, 2, 2, 401, 402, 7, 32, 2, 2,
	402, 403, 7, 7, 2, 2, 403, 404, 9, 2, 2, 2, 404, 405, 7, 4, 2, 2, 405,
	57, 3, 2, 2, 2, 406, 407, 7, 3, 2, 2, 407, 408, 7, 41, 2, 2, 408, 409,
	7, 7, 2, 2, 409, 410, 9, 3, 2, 2, 410, 411, 7, 4, 2, 2, 411, 59, 3, 2,
	2, 2, 412, 413, 7, 3, 2, 2, 413, 414, 7, 16, 2, 2, 414, 415, 7, 7, 2, 2,
	415, 417, 5, 66, 34, 2, 416, 418, 5, 62, 32, 2, 417, 416, 3, 2, 2, 2, 417,
	418, 3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 419, 420, 7, 4, 2, 2, 420, 61, 3,
	2, 2, 2, 421, 423, 5, 64, 33, 2, 422, 421, 3, 2, 2, 2, 423, 424, 3, 2,
	2, 2, 424, 422, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 63, 3, 2, 2, 2,
	426, 429, 5, 34, 18, 2, 427, 429, 5, 32, 17, 2, 428, 426, 3, 2, 2, 2, 428,
	427, 3, 2, 2, 2, 429, 65, 3, 2, 2, 2, 430, 436, 5, 68, 35, 2, 431, 436,
	5, 84, 43, 2, 432, 436, 5, 94, 48, 2, 433, 436, 5, 104, 53, 2, 434, 436,
	5, 116, 59, 2, 435, 430, 3, 2, 2, 2, 435, 431, 3, 2, 2, 2, 435, 432, 3,
	2, 2, 2, 435, 433, 3, 2, 2, 2, 435, 434, 3, 2, 2, 2, 436, 67, 3, 2, 2,
	2, 437, 438, 5, 70, 36, 2, 438, 69, 3, 2, 2, 2, 439, 441, 5, 72, 37, 2,
	440, 439, 3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 440, 3, 2, 2, 2, 442,
	443, 3, 2, 2, 2, 443, 71, 3, 2, 2, 2, 444, 448, 5, 74, 38, 2, 445, 448,
	5, 76, 39, 2, 446, 448, 5, 78, 40, 2, 447, 444, 3, 2, 2, 2, 447, 445, 3,
	2, 2, 2, 447, 446, 3, 2, 2, 2, 448, 73, 3, 2, 2, 2, 449, 450, 7, 3, 2,
	2, 450, 451, 7, 19, 2, 2, 451, 452, 7, 7, 2, 2, 452, 453, 5, 80, 41, 2,
	453, 454, 7, 4, 2, 2, 454, 75, 3, 2, 2, 2, 455, 456, 7, 3, 2, 2, 456, 457,
	7, 20, 2, 2, 457, 458, 7, 7, 2, 2, 458, 459, 5, 82, 42, 2, 459, 460, 7,
	4, 2, 2, 460, 77, 3, 2, 2, 2, 461, 462, 7, 3, 2, 2, 462, 463, 7, 17, 2,
	2, 463, 464, 7, 7, 2, 2, 464, 465, 7, 18, 2, 2, 465, 466, 7, 4, 2, 2, 466,
	79, 3, 2, 2, 2, 467, 477, 7, 79, 2, 2, 468, 471, 7, 79, 2, 2, 469, 470,
	7, 8, 2, 2, 470, 472, 7, 79, 2, 2, 471, 469, 3, 2, 2, 2, 472, 473, 3, 2,
	2, 2, 473, 471, 3, 2, 2, 2, 473, 474, 3, 2, 2, 2, 474, 477, 3, 2, 2, 2,
	475, 477, 7, 22, 2, 2, 476, 467, 3, 2, 2, 2, 476, 468, 3, 2, 2, 2, 476,
	475, 3, 2, 2, 2, 477, 81, 3, 2, 2, 2, 478, 479, 7, 27, 2, 2, 479, 83, 3,
	2, 2, 2, 480, 481, 5, 86, 44, 2, 481, 85, 3, 2, 2, 2, 482, 484, 5, 88,
	45, 2, 483, 482, 3, 2, 2, 2, 484, 485, 3, 2, 2, 2, 485, 483, 3, 2, 2, 2,
	485, 486, 3, 2, 2, 2, 486, 87, 3, 2, 2, 2, 487, 490, 5, 90, 46, 2, 488,
	490, 5, 92, 47, 2, 489, 487, 3, 2, 2, 2, 489, 488, 3, 2, 2, 2, 490, 89,
	3, 2, 2, 2, 491, 492, 7, 3, 2, 2, 492, 493, 7, 17, 2, 2, 493, 494, 7, 7,
	2, 2, 494, 495, 7, 45, 2, 2, 495, 496, 7, 4, 2, 2, 496, 91, 3, 2, 2, 2,
	497, 498, 7, 3, 2, 2, 498, 499, 7, 44, 2, 2, 499, 500, 7, 7, 2, 2, 500,
	501, 7, 79, 2, 2, 501, 502, 7, 4, 2, 2, 502, 93, 3, 2, 2, 2, 503, 504,
	5, 96, 49, 2, 504, 95, 3, 2, 2, 2, 505, 507, 5, 98, 50, 2, 506, 505, 3,
	2, 2, 2, 507, 508, 3, 2, 2, 2, 508, 506, 3, 2, 2, 2, 508, 509, 3, 2, 2,
	2, 509, 97, 3, 2, 2, 2, 510, 513, 5, 100, 51, 2, 511, 513, 5, 102, 52,
	2, 512, 510, 3, 2, 2, 2, 512, 511, 3, 2, 2, 2, 513, 99, 3, 2, 2, 2, 514,
	515, 7, 3, 2, 2, 515, 516, 7, 17, 2, 2, 516, 517, 7, 7, 2, 2, 517, 518,
	7, 46, 2, 2, 518, 519, 7, 4, 2, 2, 519, 101, 3, 2, 2, 2, 520, 521, 7, 3,
	2, 2, 521, 522, 7, 42, 2, 2, 522, 523, 7, 7, 2, 2, 523, 524, 7, 79, 2,
	2, 524, 525, 7, 4, 2, 2, 525, 103, 3, 2, 2, 2, 526, 527, 5, 106, 54, 2,
	527, 105, 3, 2, 2, 2, 528, 530, 5, 108, 55, 2, 529, 528, 3, 2, 2, 2, 530,
	531, 3, 2, 2, 2, 531, 529, 3, 2, 2, 2, 531, 532, 3, 2, 2, 2, 532, 107,
	3, 2, 2, 2, 533, 537, 5, 110, 56, 2, 534, 537, 5, 112, 57, 2, 535, 537,
	5, 114, 58, 2, 536, 533, 3, 2, 2, 2, 536, 534, 3, 2, 2, 2, 536, 535, 3,
	2, 2, 2, 537, 109, 3, 2, 2, 2, 538, 539, 7, 3, 2, 2, 539, 540, 7, 17, 2,
	2, 540, 541, 7, 7, 2, 2, 541, 542, 7, 47, 2, 2, 542, 543, 7, 4, 2, 2, 543,
	111, 3, 2, 2, 2, 544, 545, 7, 3, 2, 2, 545, 546, 7, 67, 2, 2, 546, 547,
	7, 7, 2, 2, 547, 548, 7, 79, 2, 2, 548, 549, 7, 4, 2, 2, 549, 113, 3, 2,
	2, 2, 550, 551, 7, 3, 2, 2, 551, 552, 7, 49, 2, 2, 552, 553, 7, 7, 2, 2,
	553, 554, 7, 79, 2, 2, 554, 555, 7, 4, 2, 2, 555, 115, 3, 2, 2, 2, 556,
	557, 5, 118, 60, 2, 557, 117, 3, 2, 2, 2, 558, 560, 5, 120, 61, 2, 559,
	558, 3, 2, 2, 2, 560, 561, 3, 2, 2, 2, 561, 559, 3, 2, 2, 2, 561, 562,
	3, 2, 2, 2, 562, 119, 3, 2, 2, 2, 563, 568, 5, 122, 62, 2, 564, 568, 5,
	124, 63, 2, 565, 568, 5, 126, 64, 2, 566, 568, 5, 128, 65, 2, 567, 563,
	3, 2, 2, 2, 567, 564, 3, 2, 2, 2, 567, 565, 3, 2, 2, 2, 567, 566, 3, 2,
	2, 2, 568, 121, 3, 2, 2, 2, 569, 570, 7, 3, 2, 2, 570, 571, 7, 17, 2, 2,
	571, 572, 7, 7, 2, 2, 572, 573, 7, 48, 2, 2, 573, 574, 7, 4, 2, 2, 574,
	123, 3, 2, 2, 2, 575, 576, 7, 3, 2, 2, 576, 577, 7, 50, 2, 2, 577, 578,
	7, 7, 2, 2, 578, 579, 7, 79, 2, 2, 579, 580, 7, 4, 2, 2, 580, 125, 3, 2,
	2, 2, 581, 582, 7, 3, 2, 2, 582, 583, 7, 51, 2, 2, 583, 584, 7, 7, 2, 2,
	584, 585, 7, 79, 2, 2, 585, 586, 7, 4, 2, 2, 586, 127, 3, 2, 2, 2, 587,
	588, 7, 3, 2, 2, 588, 589, 7, 52, 2, 2, 589, 590, 7, 7, 2, 2, 590, 591,
	5, 130, 66, 2, 591, 592, 7, 4, 2, 2, 592, 129, 3, 2, 2, 2, 593, 594, 7,
	11, 2, 2, 594, 595, 5, 132, 67, 2, 595, 596, 7, 11, 2, 2, 596, 131, 3,
	2, 2, 2, 597, 598, 7, 3, 2, 2, 598, 599, 7, 14, 2, 2, 599, 600, 7, 7, 2,
	2, 600, 601, 5, 134, 68, 2, 601, 602, 7, 4, 2, 2, 602, 133, 3, 2, 2, 2,
	603, 605, 5, 136, 69, 2, 604, 603, 3, 2, 2, 2, 605, 606, 3, 2, 2, 2, 606,
	604, 3, 2, 2, 2, 606, 607, 3, 2, 2, 2, 607, 135, 3, 2, 2, 2, 608, 611,
	5, 138, 70, 2, 609, 611, 5, 140, 71, 2, 610, 608, 3, 2, 2, 2, 610, 609,
	3, 2, 2, 2, 611, 137, 3, 2, 2, 2, 612, 613, 7, 3, 2, 2, 613, 614, 7, 21,
	2, 2, 614, 615, 7, 7, 2, 2, 615, 616, 7, 23, 2, 2, 616, 617, 7, 4, 2, 2,
	617, 139, 3, 2, 2, 2, 618, 619, 7, 3, 2, 2, 619, 620, 7, 16, 2, 2, 620,
	621, 7, 7, 2, 2, 621, 622, 5, 122, 62, 2, 622, 623, 7, 4, 2, 2, 623, 141,
	3, 2, 2, 2, 624, 625, 7, 3, 2, 2, 625, 626, 7, 12, 2, 2, 626, 627, 7, 7,
	2, 2, 627, 628, 5, 144, 73, 2, 628, 629, 7, 4, 2, 2, 629, 143, 3, 2, 2,
	2, 630, 632, 5, 146, 74, 2, 631, 630, 3, 2, 2, 2, 632, 633, 3, 2, 2, 2,
	633, 631, 3, 2, 2, 2, 633, 634, 3, 2, 2, 2, 634, 145, 3, 2, 2, 2, 635,
	645, 5, 148, 75, 2, 636, 645, 5, 150, 76, 2, 637, 645, 5, 152, 77, 2, 638,
	645, 5, 154, 78, 2, 639, 645, 5, 156, 79, 2, 640, 645, 5, 158, 80, 2, 641,
	645, 5, 160, 81, 2, 642, 645, 5, 162, 82, 2, 643, 645, 5, 164, 83, 2, 644,
	635, 3, 2, 2, 2, 644, 636, 3, 2, 2, 2, 644, 637, 3, 2, 2, 2, 644, 638,
	3, 2, 2, 2, 644, 639, 3, 2, 2, 2, 644, 640, 3, 2, 2, 2, 644, 641, 3, 2,
	2, 2, 644, 642, 3, 2, 2, 2, 644, 643, 3, 2, 2, 2, 645, 147, 3, 2, 2, 2,
	646, 647, 7, 3, 2, 2, 647, 648, 7, 60, 2, 2, 648, 649, 7, 7, 2, 2, 649,
	654, 7, 79, 2, 2, 650, 651, 7, 8, 2, 2, 651, 653, 7, 79, 2, 2, 652, 650,
	3, 2, 2, 2, 653, 656, 3, 2, 2, 2, 654, 652, 3, 2, 2, 2, 654, 655, 3, 2,
	2, 2, 655, 657, 3, 2, 2, 2, 656, 654, 3, 2, 2, 2, 657, 658, 7, 4, 2, 2,
	658, 149, 3, 2, 2, 2, 659, 660, 7, 3, 2, 2, 660, 661, 7, 61, 2, 2, 661,
	662, 7, 7, 2, 2, 662, 663, 7, 79, 2, 2, 663, 664, 7, 4, 2, 2, 664, 151,
	3, 2, 2, 2, 665, 666, 7, 3, 2, 2, 666, 667, 7, 62, 2, 2, 667, 668, 7, 7,
	2, 2, 668, 673, 7, 79, 2, 2, 669, 670, 7, 8, 2, 2, 670, 672, 7, 79, 2,
	2, 671, 669, 3, 2, 2, 2, 672, 675, 3, 2, 2, 2, 673, 671, 3, 2, 2, 2, 673,
	674, 3, 2, 2, 2, 674, 676, 3, 2, 2, 2, 675, 673, 3, 2, 2, 2, 676, 677,
	7, 4, 2, 2, 677, 153, 3, 2, 2, 2, 678, 679, 7, 3, 2, 2, 679, 680, 7, 63,
	2, 2, 680, 681, 7, 7, 2, 2, 681, 682, 5, 166, 84, 2, 682, 683, 7, 4, 2,
	2, 683, 155, 3, 2, 2, 2, 684, 685, 7, 3, 2, 2, 685, 686, 7, 64, 2, 2, 686,
	687, 7, 7, 2, 2, 687, 692, 7, 79, 2, 2, 688, 689, 7, 8, 2, 2, 689, 691,
	7, 79, 2, 2, 690, 688, 3, 2, 2, 2, 691, 694, 3, 2, 2, 2, 692, 690, 3, 2,
	2, 2, 692, 693, 3, 2, 2, 2, 693, 695, 3, 2, 2, 2, 694, 692, 3, 2, 2, 2,
	695, 696, 7, 4, 2, 2, 696, 157, 3, 2, 2, 2, 697, 698, 7, 3, 2, 2, 698,
	699, 7, 65, 2, 2, 699, 700, 7, 7, 2, 2, 700, 701, 7, 28, 2, 2, 701, 702,
	7, 4, 2, 2, 702, 159, 3, 2, 2, 2, 703, 704, 7, 3, 2, 2, 704, 705, 7, 66,
	2, 2, 705, 710, 7, 7, 2, 2, 706, 707, 7, 5, 2, 2, 707, 708, 7, 8, 2, 2,
	708, 709, 7, 79, 2, 2, 709, 711, 7, 6, 2, 2, 710, 706, 3, 2, 2, 2, 710,
	711, 3, 2, 2, 2, 711, 712, 3, 2, 2, 2, 712, 717, 7, 79, 2, 2, 713, 714,
	7, 8, 2, 2, 714, 716, 7, 79, 2, 2, 715, 713, 3, 2, 2, 2, 716, 719, 3, 2,
	2, 2, 717, 715, 3, 2, 2, 2, 717, 718, 3, 2, 2, 2, 718, 720, 3, 2, 2, 2,
	719, 717, 3, 2, 2, 2, 720, 721, 7, 4, 2, 2, 721, 161, 3, 2, 2, 2, 722,
	723, 7, 3, 2, 2, 723, 724, 7, 67, 2, 2, 724, 725, 7, 7, 2, 2, 725, 726,
	9, 4, 2, 2, 726, 727, 7, 4, 2, 2, 727, 163, 3, 2, 2, 2, 728, 729, 7, 3,
	2, 2, 729, 730, 7, 34, 2, 2, 730, 731, 7, 7, 2, 2, 731, 732, 7, 35, 2,
	2, 732, 733, 7, 4, 2, 2, 733, 165, 3, 2, 2, 2, 734, 736, 5, 168, 85, 2,
	735, 734, 3, 2, 2, 2, 736, 737, 3, 2, 2, 2, 737, 735, 3, 2, 2, 2, 737,
	738, 3, 2, 2, 2, 738, 167, 3, 2, 2, 2, 739, 745, 5, 170, 86, 2, 740, 745,
	5, 172, 87, 2, 741, 745, 5, 174, 88, 2, 742, 745, 5, 176, 89, 2, 743, 745,
	5, 178, 90, 2, 744, 739, 3, 2, 2, 2, 744, 740, 3, 2, 2, 2, 744, 741, 3,
	2, 2, 2, 744, 742, 3, 2, 2, 2, 744, 743, 3, 2, 2, 2, 745, 169, 3, 2, 2,
	2, 746, 747, 7, 3, 2, 2, 747, 748, 7, 69, 2, 2, 748, 749, 7, 7, 2, 2, 749,
	750, 9, 5, 2, 2, 750, 751, 7, 4, 2, 2, 751, 171, 3, 2, 2, 2, 752, 753,
	7, 3, 2, 2, 753, 754, 7, 68, 2, 2, 754, 755, 7, 7, 2, 2, 755, 760, 7, 79,
	2, 2, 756, 757, 7, 8, 2, 2, 757, 759, 7, 79, 2, 2, 758, 756, 3, 2, 2, 2,
	759, 762, 3, 2, 2, 2, 760, 758, 3, 2, 2, 2, 760, 761, 3, 2, 2, 2, 761,
	763, 3, 2, 2, 2, 762, 760, 3, 2, 2, 2, 763, 764, 7, 4, 2, 2, 764, 173,
	3, 2, 2, 2, 765, 766, 7, 3, 2, 2, 766, 767, 7, 73, 2, 2, 767, 768, 7, 7,
	2, 2, 768, 769, 9, 6, 2, 2, 769, 770, 7, 4, 2, 2, 770, 175, 3, 2, 2, 2,
	771, 772, 7, 3, 2, 2, 772, 773, 7, 76, 2, 2, 773, 774, 7, 7, 2, 2, 774,
	775, 7, 27, 2, 2, 775, 776, 7, 4, 2, 2, 776, 177, 3, 2, 2, 2, 777, 778,
	7, 3, 2, 2, 778, 779, 7, 77, 2, 2, 779, 780, 7, 7, 2, 2, 780, 781, 7, 27,
	2, 2, 781, 782, 7, 4, 2, 2, 782, 179, 3, 2, 2, 2, 57, 183, 185, 192, 205,
	207, 216, 218, 227, 236, 238, 244, 249, 252, 259, 264, 270, 276, 278, 281,
	285, 292, 306, 372, 377, 380, 387, 392, 417, 424, 428, 435, 442, 447, 473,
	476, 485, 489, 508, 512, 531, 536, 561, 567, 606, 610, 633, 644, 654, 673,
	692, 710, 717, 737, 744, 760,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'['", "']'", "", "'.'", "','", "'\"'", "'''",
}
var symbolicNames = []string{
	"", "L_PAREN", "R_PAREN", "L_SQUARE", "R_SQUARE", "EQUAL", "DOT", "COMMA",
	"D_QUOTE", "S_QUOTE", "CONNECT_DATA", "DESCRIPTION_LIST", "DESCRIPTION",
	"ADDRESS_LIST", "ADDRESS", "PROTOCOL", "TCP", "HOST", "PORT", "LOCAL",
	"IP", "YES_NO", "ON_OFF", "TRUE_FALSE", "COMMENT", "INT", "OK", "DEDICATED",
	"SHARED", "POOLED", "LOAD_BALANCE", "FAILOVER", "UR", "UR_A", "ENABLE",
	"BROKEN", "SDU", "RECV_BUF", "SEND_BUF", "SOURCE_ROUTE", "SERVICE", "SERVICE_TYPE",
	"KEY", "IPC", "SPX", "NMP", "BEQ", "PIPE", "PROGRAM", "ARGV0", "ARGS",
	"SECURITY", "SSL_CERT", "CONN_TIMEOUT", "RETRY_COUNT", "TCT", "IFILE",
	"DQ_STRING", "SERVICE_NAME", "SID", "INSTANCE_NAME", "FAILOVER_MODE", "GLOBAL_NAME",
	"HS", "RDB_DATABASE", "SERVER", "BACKUP", "TYPE", "SESSION", "SELECT",
	"NONE", "METHOD", "BASIC", "PRECONNECT", "RETRIES", "DELAY", "QUAD", "ID",
	"WS", "I_EQUAL", "I_STRING", "ISQ_STRING", "IUQ_STRING", "I_WS", "I_COMMENT",
}

var ruleNames = []string{
	"tnsnames", "tns_entry", "ifile", "lsnr_entry", "lsnr_description", "alias_list",
	"alias", "description_list", "dl_params", "dl_parameter", "description",
	"d_params", "d_parameter", "d_enable", "d_sdu", "d_recv_buf", "d_send_buf",
	"d_service_type", "d_security", "d_conn_timeout", "d_retry_count", "d_tct",
	"ds_parameter", "address_list", "al_params", "al_parameter", "al_failover",
	"al_load_balance", "al_source_route", "address", "a_params", "a_parameter",
	"protocol_info", "tcp_protocol", "tcp_params", "tcp_parameter", "tcp_host",
	"tcp_port", "tcp_tcp", "host", "port", "ipc_protocol", "ipc_params", "ipc_parameter",
	"ipc_ipc", "ipc_key", "spx_protocol", "spx_params", "spx_parameter", "spx_spx",
	"spx_service", "nmp_protocol", "nmp_params", "nmp_parameter", "nmp_nmp",
	"nmp_server", "nmp_pipe", "beq_protocol", "beq_params", "beq_parameter",
	"beq_beq", "beq_program", "beq_argv0", "beq_args", "ba_parameter", "ba_description",
	"bad_params", "bad_parameter", "bad_local", "bad_address", "connect_data",
	"cd_params", "cd_parameter", "cd_service_name", "cd_sid", "cd_instance_name",
	"cd_failover_mode", "cd_global_name", "cd_hs", "cd_rdb_database", "cd_server",
	"cd_ur", "fo_params", "fo_parameter", "fo_type", "fo_backup", "fo_method",
	"fo_retries", "fo_delay",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type tnsnamesParser struct {
	*antlr.BaseParser
}

func NewtnsnamesParser(input antlr.TokenStream) *tnsnamesParser {
	this := new(tnsnamesParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "tnsnamesParser.g4"

	return this
}

// tnsnamesParser tokens.
const (
	tnsnamesParserEOF              = antlr.TokenEOF
	tnsnamesParserL_PAREN          = 1
	tnsnamesParserR_PAREN          = 2
	tnsnamesParserL_SQUARE         = 3
	tnsnamesParserR_SQUARE         = 4
	tnsnamesParserEQUAL            = 5
	tnsnamesParserDOT              = 6
	tnsnamesParserCOMMA            = 7
	tnsnamesParserD_QUOTE          = 8
	tnsnamesParserS_QUOTE          = 9
	tnsnamesParserCONNECT_DATA     = 10
	tnsnamesParserDESCRIPTION_LIST = 11
	tnsnamesParserDESCRIPTION      = 12
	tnsnamesParserADDRESS_LIST     = 13
	tnsnamesParserADDRESS          = 14
	tnsnamesParserPROTOCOL         = 15
	tnsnamesParserTCP              = 16
	tnsnamesParserHOST             = 17
	tnsnamesParserPORT             = 18
	tnsnamesParserLOCAL            = 19
	tnsnamesParserIP               = 20
	tnsnamesParserYES_NO           = 21
	tnsnamesParserON_OFF           = 22
	tnsnamesParserTRUE_FALSE       = 23
	tnsnamesParserCOMMENT          = 24
	tnsnamesParserINT              = 25
	tnsnamesParserOK               = 26
	tnsnamesParserDEDICATED        = 27
	tnsnamesParserSHARED           = 28
	tnsnamesParserPOOLED           = 29
	tnsnamesParserLOAD_BALANCE     = 30
	tnsnamesParserFAILOVER         = 31
	tnsnamesParserUR               = 32
	tnsnamesParserUR_A             = 33
	tnsnamesParserENABLE           = 34
	tnsnamesParserBROKEN           = 35
	tnsnamesParserSDU              = 36
	tnsnamesParserRECV_BUF         = 37
	tnsnamesParserSEND_BUF         = 38
	tnsnamesParserSOURCE_ROUTE     = 39
	tnsnamesParserSERVICE          = 40
	tnsnamesParserSERVICE_TYPE     = 41
	tnsnamesParserKEY              = 42
	tnsnamesParserIPC              = 43
	tnsnamesParserSPX              = 44
	tnsnamesParserNMP              = 45
	tnsnamesParserBEQ              = 46
	tnsnamesParserPIPE             = 47
	tnsnamesParserPROGRAM          = 48
	tnsnamesParserARGV0            = 49
	tnsnamesParserARGS             = 50
	tnsnamesParserSECURITY         = 51
	tnsnamesParserSSL_CERT         = 52
	tnsnamesParserCONN_TIMEOUT     = 53
	tnsnamesParserRETRY_COUNT      = 54
	tnsnamesParserTCT              = 55
	tnsnamesParserIFILE            = 56
	tnsnamesParserDQ_STRING        = 57
	tnsnamesParserSERVICE_NAME     = 58
	tnsnamesParserSID              = 59
	tnsnamesParserINSTANCE_NAME    = 60
	tnsnamesParserFAILOVER_MODE    = 61
	tnsnamesParserGLOBAL_NAME      = 62
	tnsnamesParserHS               = 63
	tnsnamesParserRDB_DATABASE     = 64
	tnsnamesParserSERVER           = 65
	tnsnamesParserBACKUP           = 66
	tnsnamesParserTYPE             = 67
	tnsnamesParserSESSION          = 68
	tnsnamesParserSELECT           = 69
	tnsnamesParserNONE             = 70
	tnsnamesParserMETHOD           = 71
	tnsnamesParserBASIC            = 72
	tnsnamesParserPRECONNECT       = 73
	tnsnamesParserRETRIES          = 74
	tnsnamesParserDELAY            = 75
	tnsnamesParserQUAD             = 76
	tnsnamesParserID               = 77
	tnsnamesParserWS               = 78
	tnsnamesParserI_EQUAL          = 79
	tnsnamesParserI_STRING         = 80
	tnsnamesParserISQ_STRING       = 81
	tnsnamesParserIUQ_STRING       = 82
	tnsnamesParserI_WS             = 83
	tnsnamesParserI_COMMENT        = 84
)

// tnsnamesParser rules.
const (
	tnsnamesParserRULE_tnsnames         = 0
	tnsnamesParserRULE_tns_entry        = 1
	tnsnamesParserRULE_ifile            = 2
	tnsnamesParserRULE_lsnr_entry       = 3
	tnsnamesParserRULE_lsnr_description = 4
	tnsnamesParserRULE_alias_list       = 5
	tnsnamesParserRULE_alias            = 6
	tnsnamesParserRULE_description_list = 7
	tnsnamesParserRULE_dl_params        = 8
	tnsnamesParserRULE_dl_parameter     = 9
	tnsnamesParserRULE_description      = 10
	tnsnamesParserRULE_d_params         = 11
	tnsnamesParserRULE_d_parameter      = 12
	tnsnamesParserRULE_d_enable         = 13
	tnsnamesParserRULE_d_sdu            = 14
	tnsnamesParserRULE_d_recv_buf       = 15
	tnsnamesParserRULE_d_send_buf       = 16
	tnsnamesParserRULE_d_service_type   = 17
	tnsnamesParserRULE_d_security       = 18
	tnsnamesParserRULE_d_conn_timeout   = 19
	tnsnamesParserRULE_d_retry_count    = 20
	tnsnamesParserRULE_d_tct            = 21
	tnsnamesParserRULE_ds_parameter     = 22
	tnsnamesParserRULE_address_list     = 23
	tnsnamesParserRULE_al_params        = 24
	tnsnamesParserRULE_al_parameter     = 25
	tnsnamesParserRULE_al_failover      = 26
	tnsnamesParserRULE_al_load_balance  = 27
	tnsnamesParserRULE_al_source_route  = 28
	tnsnamesParserRULE_address          = 29
	tnsnamesParserRULE_a_params         = 30
	tnsnamesParserRULE_a_parameter      = 31
	tnsnamesParserRULE_protocol_info    = 32
	tnsnamesParserRULE_tcp_protocol     = 33
	tnsnamesParserRULE_tcp_params       = 34
	tnsnamesParserRULE_tcp_parameter    = 35
	tnsnamesParserRULE_tcp_host         = 36
	tnsnamesParserRULE_tcp_port         = 37
	tnsnamesParserRULE_tcp_tcp          = 38
	tnsnamesParserRULE_host             = 39
	tnsnamesParserRULE_port             = 40
	tnsnamesParserRULE_ipc_protocol     = 41
	tnsnamesParserRULE_ipc_params       = 42
	tnsnamesParserRULE_ipc_parameter    = 43
	tnsnamesParserRULE_ipc_ipc          = 44
	tnsnamesParserRULE_ipc_key          = 45
	tnsnamesParserRULE_spx_protocol     = 46
	tnsnamesParserRULE_spx_params       = 47
	tnsnamesParserRULE_spx_parameter    = 48
	tnsnamesParserRULE_spx_spx          = 49
	tnsnamesParserRULE_spx_service      = 50
	tnsnamesParserRULE_nmp_protocol     = 51
	tnsnamesParserRULE_nmp_params       = 52
	tnsnamesParserRULE_nmp_parameter    = 53
	tnsnamesParserRULE_nmp_nmp          = 54
	tnsnamesParserRULE_nmp_server       = 55
	tnsnamesParserRULE_nmp_pipe         = 56
	tnsnamesParserRULE_beq_protocol     = 57
	tnsnamesParserRULE_beq_params       = 58
	tnsnamesParserRULE_beq_parameter    = 59
	tnsnamesParserRULE_beq_beq          = 60
	tnsnamesParserRULE_beq_program      = 61
	tnsnamesParserRULE_beq_argv0        = 62
	tnsnamesParserRULE_beq_args         = 63
	tnsnamesParserRULE_ba_parameter     = 64
	tnsnamesParserRULE_ba_description   = 65
	tnsnamesParserRULE_bad_params       = 66
	tnsnamesParserRULE_bad_parameter    = 67
	tnsnamesParserRULE_bad_local        = 68
	tnsnamesParserRULE_bad_address      = 69
	tnsnamesParserRULE_connect_data     = 70
	tnsnamesParserRULE_cd_params        = 71
	tnsnamesParserRULE_cd_parameter     = 72
	tnsnamesParserRULE_cd_service_name  = 73
	tnsnamesParserRULE_cd_sid           = 74
	tnsnamesParserRULE_cd_instance_name = 75
	tnsnamesParserRULE_cd_failover_mode = 76
	tnsnamesParserRULE_cd_global_name   = 77
	tnsnamesParserRULE_cd_hs            = 78
	tnsnamesParserRULE_cd_rdb_database  = 79
	tnsnamesParserRULE_cd_server        = 80
	tnsnamesParserRULE_cd_ur            = 81
	tnsnamesParserRULE_fo_params        = 82
	tnsnamesParserRULE_fo_parameter     = 83
	tnsnamesParserRULE_fo_type          = 84
	tnsnamesParserRULE_fo_backup        = 85
	tnsnamesParserRULE_fo_method        = 86
	tnsnamesParserRULE_fo_retries       = 87
	tnsnamesParserRULE_fo_delay         = 88
)

// ITnsnamesContext is an interface to support dynamic dispatch.
type ITnsnamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTnsnamesContext differentiates from other interfaces.
	IsTnsnamesContext()
}

type TnsnamesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTnsnamesContext() *TnsnamesContext {
	var p = new(TnsnamesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_tnsnames
	return p
}

func (*TnsnamesContext) IsTnsnamesContext() {}

func NewTnsnamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TnsnamesContext {
	var p = new(TnsnamesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_tnsnames

	return p
}

func (s *TnsnamesContext) GetParser() antlr.Parser { return s.parser }

func (s *TnsnamesContext) AllTns_entry() []ITns_entryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITns_entryContext)(nil)).Elem())
	var tst = make([]ITns_entryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITns_entryContext)
		}
	}

	return tst
}

func (s *TnsnamesContext) Tns_entry(i int) ITns_entryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITns_entryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITns_entryContext)
}

func (s *TnsnamesContext) AllIfile() []IIfileContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIfileContext)(nil)).Elem())
	var tst = make([]IIfileContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIfileContext)
		}
	}

	return tst
}

func (s *TnsnamesContext) Ifile(i int) IIfileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfileContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIfileContext)
}

func (s *TnsnamesContext) AllLsnr_entry() []ILsnr_entryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILsnr_entryContext)(nil)).Elem())
	var tst = make([]ILsnr_entryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILsnr_entryContext)
		}
	}

	return tst
}

func (s *TnsnamesContext) Lsnr_entry(i int) ILsnr_entryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILsnr_entryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILsnr_entryContext)
}

func (s *TnsnamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TnsnamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TnsnamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterTnsnames(s)
	}
}

func (s *TnsnamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitTnsnames(s)
	}
}

func (s *TnsnamesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitTnsnames(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Tnsnames() (localctx ITnsnamesContext) {
	localctx = NewTnsnamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tnsnamesParserRULE_tnsnames)
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
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tnsnamesParserIFILE || _la == tnsnamesParserID {
		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(178)
				p.Tns_entry()
			}

		case 2:
			{
				p.SetState(179)
				p.Ifile()
			}

		case 3:
			{
				p.SetState(180)
				p.Lsnr_entry()
			}

		}

		p.SetState(185)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITns_entryContext is an interface to support dynamic dispatch.
type ITns_entryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTns_entryContext differentiates from other interfaces.
	IsTns_entryContext()
}

type Tns_entryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTns_entryContext() *Tns_entryContext {
	var p = new(Tns_entryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_tns_entry
	return p
}

func (*Tns_entryContext) IsTns_entryContext() {}

func NewTns_entryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tns_entryContext {
	var p = new(Tns_entryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_tns_entry

	return p
}

func (s *Tns_entryContext) GetParser() antlr.Parser { return s.parser }

func (s *Tns_entryContext) Alias_list() IAlias_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlias_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlias_listContext)
}

func (s *Tns_entryContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Tns_entryContext) Description_list() IDescription_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescription_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescription_listContext)
}

func (s *Tns_entryContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *Tns_entryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tns_entryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tns_entryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterTns_entry(s)
	}
}

func (s *Tns_entryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitTns_entry(s)
	}
}

func (s *Tns_entryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitTns_entry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Tns_entry() (localctx ITns_entryContext) {
	localctx = NewTns_entryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tnsnamesParserRULE_tns_entry)

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
		p.SetState(186)
		p.Alias_list()
	}
	{
		p.SetState(187)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(188)
			p.Description_list()
		}

	case 2:
		{
			p.SetState(189)
			p.Description()
		}

	}

	return localctx
}

// IIfileContext is an interface to support dynamic dispatch.
type IIfileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfileContext differentiates from other interfaces.
	IsIfileContext()
}

type IfileContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfileContext() *IfileContext {
	var p = new(IfileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_ifile
	return p
}

func (*IfileContext) IsIfileContext() {}

func NewIfileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfileContext {
	var p = new(IfileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_ifile

	return p
}

func (s *IfileContext) GetParser() antlr.Parser { return s.parser }

func (s *IfileContext) IFILE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserIFILE, 0)
}

func (s *IfileContext) I_EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserI_EQUAL, 0)
}

func (s *IfileContext) I_STRING() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserI_STRING, 0)
}

func (s *IfileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterIfile(s)
	}
}

func (s *IfileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitIfile(s)
	}
}

func (s *IfileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitIfile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Ifile() (localctx IIfileContext) {
	localctx = NewIfileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tnsnamesParserRULE_ifile)

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
		p.SetState(192)
		p.Match(tnsnamesParserIFILE)
	}
	{
		p.SetState(193)
		p.Match(tnsnamesParserI_EQUAL)
	}
	{
		p.SetState(194)
		p.Match(tnsnamesParserI_STRING)
	}

	return localctx
}

// ILsnr_entryContext is an interface to support dynamic dispatch.
type ILsnr_entryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLsnr_entryContext differentiates from other interfaces.
	IsLsnr_entryContext()
}

type Lsnr_entryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLsnr_entryContext() *Lsnr_entryContext {
	var p = new(Lsnr_entryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_lsnr_entry
	return p
}

func (*Lsnr_entryContext) IsLsnr_entryContext() {}

func NewLsnr_entryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lsnr_entryContext {
	var p = new(Lsnr_entryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_lsnr_entry

	return p
}

func (s *Lsnr_entryContext) GetParser() antlr.Parser { return s.parser }

func (s *Lsnr_entryContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *Lsnr_entryContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Lsnr_entryContext) Lsnr_description() ILsnr_descriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILsnr_descriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILsnr_descriptionContext)
}

func (s *Lsnr_entryContext) Address_list() IAddress_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddress_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddress_listContext)
}

func (s *Lsnr_entryContext) AllAddress() []IAddressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAddressContext)(nil)).Elem())
	var tst = make([]IAddressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAddressContext)
		}
	}

	return tst
}

func (s *Lsnr_entryContext) Address(i int) IAddressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAddressContext)
}

func (s *Lsnr_entryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lsnr_entryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lsnr_entryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterLsnr_entry(s)
	}
}

func (s *Lsnr_entryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitLsnr_entry(s)
	}
}

func (s *Lsnr_entryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitLsnr_entry(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Lsnr_entry() (localctx ILsnr_entryContext) {
	localctx = NewLsnr_entryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tnsnamesParserRULE_lsnr_entry)
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
		p.SetState(196)
		p.Alias()
	}
	{
		p.SetState(197)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(198)
			p.Lsnr_description()
		}

	case 2:
		{
			p.SetState(199)
			p.Address_list()
		}

	case 3:
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == tnsnamesParserL_PAREN {
			{
				p.SetState(200)
				p.Address()
			}

			p.SetState(203)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// ILsnr_descriptionContext is an interface to support dynamic dispatch.
type ILsnr_descriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLsnr_descriptionContext differentiates from other interfaces.
	IsLsnr_descriptionContext()
}

type Lsnr_descriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLsnr_descriptionContext() *Lsnr_descriptionContext {
	var p = new(Lsnr_descriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_lsnr_description
	return p
}

func (*Lsnr_descriptionContext) IsLsnr_descriptionContext() {}

func NewLsnr_descriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Lsnr_descriptionContext {
	var p = new(Lsnr_descriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_lsnr_description

	return p
}

func (s *Lsnr_descriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *Lsnr_descriptionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Lsnr_descriptionContext) DESCRIPTION() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDESCRIPTION, 0)
}

func (s *Lsnr_descriptionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Lsnr_descriptionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Lsnr_descriptionContext) Address_list() IAddress_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddress_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddress_listContext)
}

func (s *Lsnr_descriptionContext) AllAddress() []IAddressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAddressContext)(nil)).Elem())
	var tst = make([]IAddressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAddressContext)
		}
	}

	return tst
}

func (s *Lsnr_descriptionContext) Address(i int) IAddressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAddressContext)
}

func (s *Lsnr_descriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Lsnr_descriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Lsnr_descriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterLsnr_description(s)
	}
}

func (s *Lsnr_descriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitLsnr_description(s)
	}
}

func (s *Lsnr_descriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitLsnr_description(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Lsnr_description() (localctx ILsnr_descriptionContext) {
	localctx = NewLsnr_descriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tnsnamesParserRULE_lsnr_description)
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
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(208)
		p.Match(tnsnamesParserDESCRIPTION)
	}
	{
		p.SetState(209)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(210)
			p.Address_list()
		}

	case 2:
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == tnsnamesParserL_PAREN {
			{
				p.SetState(211)
				p.Address()
			}

			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(218)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IAlias_listContext is an interface to support dynamic dispatch.
type IAlias_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlias_listContext differentiates from other interfaces.
	IsAlias_listContext()
}

type Alias_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlias_listContext() *Alias_listContext {
	var p = new(Alias_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_alias_list
	return p
}

func (*Alias_listContext) IsAlias_listContext() {}

func NewAlias_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Alias_listContext {
	var p = new(Alias_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_alias_list

	return p
}

func (s *Alias_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Alias_listContext) AllAlias() []IAliasContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAliasContext)(nil)).Elem())
	var tst = make([]IAliasContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAliasContext)
		}
	}

	return tst
}

func (s *Alias_listContext) Alias(i int) IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *Alias_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserCOMMA)
}

func (s *Alias_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserCOMMA, i)
}

func (s *Alias_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Alias_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Alias_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterAlias_list(s)
	}
}

func (s *Alias_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitAlias_list(s)
	}
}

func (s *Alias_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitAlias_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Alias_list() (localctx IAlias_listContext) {
	localctx = NewAlias_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tnsnamesParserRULE_alias_list)
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
		p.SetState(220)
		p.Alias()
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tnsnamesParserCOMMA {
		{
			p.SetState(221)
			p.Match(tnsnamesParserCOMMA)
		}
		{
			p.SetState(222)
			p.Alias()
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserID)
}

func (s *AliasContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, i)
}

func (s *AliasContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserDOT)
}

func (s *AliasContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDOT, i)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (s *AliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tnsnamesParserRULE_alias)
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

	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.Match(tnsnamesParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(229)
			p.Match(tnsnamesParserID)
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == tnsnamesParserDOT {
			{
				p.SetState(230)
				p.Match(tnsnamesParserDOT)
			}
			{
				p.SetState(231)
				p.Match(tnsnamesParserID)
			}

			p.SetState(234)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IDescription_listContext is an interface to support dynamic dispatch.
type IDescription_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescription_listContext differentiates from other interfaces.
	IsDescription_listContext()
}

type Description_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescription_listContext() *Description_listContext {
	var p = new(Description_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_description_list
	return p
}

func (*Description_listContext) IsDescription_listContext() {}

func NewDescription_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Description_listContext {
	var p = new(Description_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_description_list

	return p
}

func (s *Description_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Description_listContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Description_listContext) DESCRIPTION_LIST() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDESCRIPTION_LIST, 0)
}

func (s *Description_listContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Description_listContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Description_listContext) AllDl_params() []IDl_paramsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDl_paramsContext)(nil)).Elem())
	var tst = make([]IDl_paramsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDl_paramsContext)
		}
	}

	return tst
}

func (s *Description_listContext) Dl_params(i int) IDl_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDl_paramsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDl_paramsContext)
}

func (s *Description_listContext) AllDescription() []IDescriptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDescriptionContext)(nil)).Elem())
	var tst = make([]IDescriptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDescriptionContext)
		}
	}

	return tst
}

func (s *Description_listContext) Description(i int) IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *Description_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Description_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Description_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterDescription_list(s)
	}
}

func (s *Description_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitDescription_list(s)
	}
}

func (s *Description_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitDescription_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Description_list() (localctx IDescription_listContext) {
	localctx = NewDescription_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tnsnamesParserRULE_description_list)
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
		p.SetState(238)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(239)
		p.Match(tnsnamesParserDESCRIPTION_LIST)
	}
	{
		p.SetState(240)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(241)
			p.Dl_params()
		}

	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(244)
				p.Description()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tnsnamesParserL_PAREN {
		{
			p.SetState(249)
			p.Dl_params()
		}

	}
	{
		p.SetState(252)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IDl_paramsContext is an interface to support dynamic dispatch.
type IDl_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDl_paramsContext differentiates from other interfaces.
	IsDl_paramsContext()
}

type Dl_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDl_paramsContext() *Dl_paramsContext {
	var p = new(Dl_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_dl_params
	return p
}

func (*Dl_paramsContext) IsDl_paramsContext() {}

func NewDl_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dl_paramsContext {
	var p = new(Dl_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_dl_params

	return p
}

func (s *Dl_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Dl_paramsContext) AllDl_parameter() []IDl_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDl_parameterContext)(nil)).Elem())
	var tst = make([]IDl_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDl_parameterContext)
		}
	}

	return tst
}

func (s *Dl_paramsContext) Dl_parameter(i int) IDl_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDl_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDl_parameterContext)
}

func (s *Dl_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dl_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dl_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterDl_params(s)
	}
}

func (s *Dl_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitDl_params(s)
	}
}

func (s *Dl_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitDl_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Dl_params() (localctx IDl_paramsContext) {
	localctx = NewDl_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, tnsnamesParserRULE_dl_params)

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
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(254)
				p.Dl_parameter()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}

	return localctx
}

// IDl_parameterContext is an interface to support dynamic dispatch.
type IDl_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDl_parameterContext differentiates from other interfaces.
	IsDl_parameterContext()
}

type Dl_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDl_parameterContext() *Dl_parameterContext {
	var p = new(Dl_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_dl_parameter
	return p
}

func (*Dl_parameterContext) IsDl_parameterContext() {}

func NewDl_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dl_parameterContext {
	var p = new(Dl_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_dl_parameter

	return p
}

func (s *Dl_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Dl_parameterContext) Al_failover() IAl_failoverContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAl_failoverContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAl_failoverContext)
}

func (s *Dl_parameterContext) Al_load_balance() IAl_load_balanceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAl_load_balanceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAl_load_balanceContext)
}

func (s *Dl_parameterContext) Al_source_route() IAl_source_routeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAl_source_routeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAl_source_routeContext)
}

func (s *Dl_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dl_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dl_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterDl_parameter(s)
	}
}

func (s *Dl_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitDl_parameter(s)
	}
}

func (s *Dl_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitDl_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Dl_parameter() (localctx IDl_parameterContext) {
	localctx = NewDl_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, tnsnamesParserRULE_dl_parameter)

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

	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(259)
			p.Al_failover()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.Al_load_balance()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(261)
			p.Al_source_route()
		}

	}

	return localctx
}

// IDescriptionContext is an interface to support dynamic dispatch.
type IDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionContext differentiates from other interfaces.
	IsDescriptionContext()
}

type DescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionContext() *DescriptionContext {
	var p = new(DescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_description
	return p
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *DescriptionContext) DESCRIPTION() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDESCRIPTION, 0)
}

func (s *DescriptionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *DescriptionContext) Connect_data() IConnect_dataContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnect_dataContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnect_dataContext)
}

func (s *DescriptionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *DescriptionContext) Address_list() IAddress_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddress_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAddress_listContext)
}

func (s *DescriptionContext) AllD_params() []ID_paramsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ID_paramsContext)(nil)).Elem())
	var tst = make([]ID_paramsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ID_paramsContext)
		}
	}

	return tst
}

func (s *DescriptionContext) D_params(i int) ID_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_paramsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ID_paramsContext)
}

func (s *DescriptionContext) AllAddress() []IAddressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAddressContext)(nil)).Elem())
	var tst = make([]IAddressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAddressContext)
		}
	}

	return tst
}

func (s *DescriptionContext) Address(i int) IAddressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAddressContext)
}

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterDescription(s)
	}
}

func (s *DescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitDescription(s)
	}
}

func (s *DescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Description() (localctx IDescriptionContext) {
	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, tnsnamesParserRULE_description)
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
		p.SetState(264)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(265)
		p.Match(tnsnamesParserDESCRIPTION)
	}
	{
		p.SetState(266)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(267)
			p.D_params()
		}

	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(270)
			p.Address_list()
		}

	case 2:
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(271)
					p.Address()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
		}

	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(278)
			p.D_params()
		}

	}
	{
		p.SetState(281)
		p.Connect_data()
	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tnsnamesParserL_PAREN {
		{
			p.SetState(282)
			p.D_params()
		}

	}
	{
		p.SetState(285)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ID_paramsContext is an interface to support dynamic dispatch.
type ID_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsD_paramsContext differentiates from other interfaces.
	IsD_paramsContext()
}

type D_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyD_paramsContext() *D_paramsContext {
	var p = new(D_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_d_params
	return p
}

func (*D_paramsContext) IsD_paramsContext() {}

func NewD_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *D_paramsContext {
	var p = new(D_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_d_params

	return p
}

func (s *D_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *D_paramsContext) AllD_parameter() []ID_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ID_parameterContext)(nil)).Elem())
	var tst = make([]ID_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ID_parameterContext)
		}
	}

	return tst
}

func (s *D_paramsContext) D_parameter(i int) ID_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ID_parameterContext)
}

func (s *D_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *D_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *D_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterD_params(s)
	}
}

func (s *D_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitD_params(s)
	}
}

func (s *D_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitD_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) D_params() (localctx ID_paramsContext) {
	localctx = NewD_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, tnsnamesParserRULE_d_params)

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
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(287)
				p.D_parameter()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// ID_parameterContext is an interface to support dynamic dispatch.
type ID_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsD_parameterContext differentiates from other interfaces.
	IsD_parameterContext()
}

type D_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyD_parameterContext() *D_parameterContext {
	var p = new(D_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_d_parameter
	return p
}

func (*D_parameterContext) IsD_parameterContext() {}

func NewD_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *D_parameterContext {
	var p = new(D_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_d_parameter

	return p
}

func (s *D_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *D_parameterContext) D_enable() ID_enableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_enableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ID_enableContext)
}

func (s *D_parameterContext) Al_failover() IAl_failoverContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAl_failoverContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAl_failoverContext)
}

func (s *D_parameterContext) Al_load_balance() IAl_load_balanceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAl_load_balanceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAl_load_balanceContext)
}

func (s *D_parameterContext) D_sdu() ID_sduContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_sduContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ID_sduContext)
}

func (s *D_parameterContext) D_recv_buf() ID_recv_bufContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_recv_bufContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ID_recv_bufContext)
}

func (s *D_parameterContext) D_send_buf() ID_send_bufContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_send_bufContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ID_send_bufContext)
}

func (s *D_parameterContext) Al_source_route() IAl_source_routeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAl_source_routeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAl_source_routeContext)
}

func (s *D_parameterContext) D_service_type() ID_service_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_service_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ID_service_typeContext)
}

func (s *D_parameterContext) D_security() ID_securityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_securityContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ID_securityContext)
}

func (s *D_parameterContext) D_conn_timeout() ID_conn_timeoutContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_conn_timeoutContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ID_conn_timeoutContext)
}

func (s *D_parameterContext) D_retry_count() ID_retry_countContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_retry_countContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ID_retry_countContext)
}

func (s *D_parameterContext) D_tct() ID_tctContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_tctContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ID_tctContext)
}

func (s *D_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *D_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *D_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterD_parameter(s)
	}
}

func (s *D_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitD_parameter(s)
	}
}

func (s *D_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitD_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) D_parameter() (localctx ID_parameterContext) {
	localctx = NewD_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, tnsnamesParserRULE_d_parameter)

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

	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(292)
			p.D_enable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(293)
			p.Al_failover()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(294)
			p.Al_load_balance()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(295)
			p.D_sdu()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(296)
			p.D_recv_buf()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(297)
			p.D_send_buf()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(298)
			p.Al_source_route()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(299)
			p.D_service_type()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(300)
			p.D_security()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(301)
			p.D_conn_timeout()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(302)
			p.D_retry_count()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(303)
			p.D_tct()
		}

	}

	return localctx
}

// ID_enableContext is an interface to support dynamic dispatch.
type ID_enableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsD_enableContext differentiates from other interfaces.
	IsD_enableContext()
}

type D_enableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyD_enableContext() *D_enableContext {
	var p = new(D_enableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_d_enable
	return p
}

func (*D_enableContext) IsD_enableContext() {}

func NewD_enableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *D_enableContext {
	var p = new(D_enableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_d_enable

	return p
}

func (s *D_enableContext) GetParser() antlr.Parser { return s.parser }

func (s *D_enableContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *D_enableContext) ENABLE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserENABLE, 0)
}

func (s *D_enableContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *D_enableContext) BROKEN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserBROKEN, 0)
}

func (s *D_enableContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *D_enableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *D_enableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *D_enableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterD_enable(s)
	}
}

func (s *D_enableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitD_enable(s)
	}
}

func (s *D_enableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitD_enable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) D_enable() (localctx ID_enableContext) {
	localctx = NewD_enableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, tnsnamesParserRULE_d_enable)

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
		p.SetState(306)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(307)
		p.Match(tnsnamesParserENABLE)
	}
	{
		p.SetState(308)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(309)
		p.Match(tnsnamesParserBROKEN)
	}
	{
		p.SetState(310)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ID_sduContext is an interface to support dynamic dispatch.
type ID_sduContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsD_sduContext differentiates from other interfaces.
	IsD_sduContext()
}

type D_sduContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyD_sduContext() *D_sduContext {
	var p = new(D_sduContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_d_sdu
	return p
}

func (*D_sduContext) IsD_sduContext() {}

func NewD_sduContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *D_sduContext {
	var p = new(D_sduContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_d_sdu

	return p
}

func (s *D_sduContext) GetParser() antlr.Parser { return s.parser }

func (s *D_sduContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *D_sduContext) SDU() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSDU, 0)
}

func (s *D_sduContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *D_sduContext) INT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserINT, 0)
}

func (s *D_sduContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *D_sduContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *D_sduContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *D_sduContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterD_sdu(s)
	}
}

func (s *D_sduContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitD_sdu(s)
	}
}

func (s *D_sduContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitD_sdu(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) D_sdu() (localctx ID_sduContext) {
	localctx = NewD_sduContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, tnsnamesParserRULE_d_sdu)

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
		p.SetState(312)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(313)
		p.Match(tnsnamesParserSDU)
	}
	{
		p.SetState(314)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(315)
		p.Match(tnsnamesParserINT)
	}
	{
		p.SetState(316)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ID_recv_bufContext is an interface to support dynamic dispatch.
type ID_recv_bufContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsD_recv_bufContext differentiates from other interfaces.
	IsD_recv_bufContext()
}

type D_recv_bufContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyD_recv_bufContext() *D_recv_bufContext {
	var p = new(D_recv_bufContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_d_recv_buf
	return p
}

func (*D_recv_bufContext) IsD_recv_bufContext() {}

func NewD_recv_bufContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *D_recv_bufContext {
	var p = new(D_recv_bufContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_d_recv_buf

	return p
}

func (s *D_recv_bufContext) GetParser() antlr.Parser { return s.parser }

func (s *D_recv_bufContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *D_recv_bufContext) RECV_BUF() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserRECV_BUF, 0)
}

func (s *D_recv_bufContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *D_recv_bufContext) INT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserINT, 0)
}

func (s *D_recv_bufContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *D_recv_bufContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *D_recv_bufContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *D_recv_bufContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterD_recv_buf(s)
	}
}

func (s *D_recv_bufContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitD_recv_buf(s)
	}
}

func (s *D_recv_bufContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitD_recv_buf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) D_recv_buf() (localctx ID_recv_bufContext) {
	localctx = NewD_recv_bufContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, tnsnamesParserRULE_d_recv_buf)

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
		p.SetState(318)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(319)
		p.Match(tnsnamesParserRECV_BUF)
	}
	{
		p.SetState(320)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(321)
		p.Match(tnsnamesParserINT)
	}
	{
		p.SetState(322)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ID_send_bufContext is an interface to support dynamic dispatch.
type ID_send_bufContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsD_send_bufContext differentiates from other interfaces.
	IsD_send_bufContext()
}

type D_send_bufContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyD_send_bufContext() *D_send_bufContext {
	var p = new(D_send_bufContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_d_send_buf
	return p
}

func (*D_send_bufContext) IsD_send_bufContext() {}

func NewD_send_bufContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *D_send_bufContext {
	var p = new(D_send_bufContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_d_send_buf

	return p
}

func (s *D_send_bufContext) GetParser() antlr.Parser { return s.parser }

func (s *D_send_bufContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *D_send_bufContext) SEND_BUF() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSEND_BUF, 0)
}

func (s *D_send_bufContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *D_send_bufContext) INT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserINT, 0)
}

func (s *D_send_bufContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *D_send_bufContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *D_send_bufContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *D_send_bufContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterD_send_buf(s)
	}
}

func (s *D_send_bufContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitD_send_buf(s)
	}
}

func (s *D_send_bufContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitD_send_buf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) D_send_buf() (localctx ID_send_bufContext) {
	localctx = NewD_send_bufContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, tnsnamesParserRULE_d_send_buf)

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
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(325)
		p.Match(tnsnamesParserSEND_BUF)
	}
	{
		p.SetState(326)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(327)
		p.Match(tnsnamesParserINT)
	}
	{
		p.SetState(328)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ID_service_typeContext is an interface to support dynamic dispatch.
type ID_service_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsD_service_typeContext differentiates from other interfaces.
	IsD_service_typeContext()
}

type D_service_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyD_service_typeContext() *D_service_typeContext {
	var p = new(D_service_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_d_service_type
	return p
}

func (*D_service_typeContext) IsD_service_typeContext() {}

func NewD_service_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *D_service_typeContext {
	var p = new(D_service_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_d_service_type

	return p
}

func (s *D_service_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *D_service_typeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *D_service_typeContext) SERVICE_TYPE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSERVICE_TYPE, 0)
}

func (s *D_service_typeContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *D_service_typeContext) ID() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, 0)
}

func (s *D_service_typeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *D_service_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *D_service_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *D_service_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterD_service_type(s)
	}
}

func (s *D_service_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitD_service_type(s)
	}
}

func (s *D_service_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitD_service_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) D_service_type() (localctx ID_service_typeContext) {
	localctx = NewD_service_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, tnsnamesParserRULE_d_service_type)

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
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(331)
		p.Match(tnsnamesParserSERVICE_TYPE)
	}
	{
		p.SetState(332)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(333)
		p.Match(tnsnamesParserID)
	}
	{
		p.SetState(334)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ID_securityContext is an interface to support dynamic dispatch.
type ID_securityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsD_securityContext differentiates from other interfaces.
	IsD_securityContext()
}

type D_securityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyD_securityContext() *D_securityContext {
	var p = new(D_securityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_d_security
	return p
}

func (*D_securityContext) IsD_securityContext() {}

func NewD_securityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *D_securityContext {
	var p = new(D_securityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_d_security

	return p
}

func (s *D_securityContext) GetParser() antlr.Parser { return s.parser }

func (s *D_securityContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *D_securityContext) SECURITY() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSECURITY, 0)
}

func (s *D_securityContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *D_securityContext) Ds_parameter() IDs_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDs_parameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDs_parameterContext)
}

func (s *D_securityContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *D_securityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *D_securityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *D_securityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterD_security(s)
	}
}

func (s *D_securityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitD_security(s)
	}
}

func (s *D_securityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitD_security(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) D_security() (localctx ID_securityContext) {
	localctx = NewD_securityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, tnsnamesParserRULE_d_security)

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
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(337)
		p.Match(tnsnamesParserSECURITY)
	}
	{
		p.SetState(338)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(339)
		p.Ds_parameter()
	}
	{
		p.SetState(340)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ID_conn_timeoutContext is an interface to support dynamic dispatch.
type ID_conn_timeoutContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsD_conn_timeoutContext differentiates from other interfaces.
	IsD_conn_timeoutContext()
}

type D_conn_timeoutContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyD_conn_timeoutContext() *D_conn_timeoutContext {
	var p = new(D_conn_timeoutContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_d_conn_timeout
	return p
}

func (*D_conn_timeoutContext) IsD_conn_timeoutContext() {}

func NewD_conn_timeoutContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *D_conn_timeoutContext {
	var p = new(D_conn_timeoutContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_d_conn_timeout

	return p
}

func (s *D_conn_timeoutContext) GetParser() antlr.Parser { return s.parser }

func (s *D_conn_timeoutContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *D_conn_timeoutContext) CONN_TIMEOUT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserCONN_TIMEOUT, 0)
}

func (s *D_conn_timeoutContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *D_conn_timeoutContext) INT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserINT, 0)
}

func (s *D_conn_timeoutContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *D_conn_timeoutContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *D_conn_timeoutContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *D_conn_timeoutContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterD_conn_timeout(s)
	}
}

func (s *D_conn_timeoutContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitD_conn_timeout(s)
	}
}

func (s *D_conn_timeoutContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitD_conn_timeout(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) D_conn_timeout() (localctx ID_conn_timeoutContext) {
	localctx = NewD_conn_timeoutContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, tnsnamesParserRULE_d_conn_timeout)

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
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(343)
		p.Match(tnsnamesParserCONN_TIMEOUT)
	}
	{
		p.SetState(344)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(345)
		p.Match(tnsnamesParserINT)
	}
	{
		p.SetState(346)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ID_retry_countContext is an interface to support dynamic dispatch.
type ID_retry_countContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsD_retry_countContext differentiates from other interfaces.
	IsD_retry_countContext()
}

type D_retry_countContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyD_retry_countContext() *D_retry_countContext {
	var p = new(D_retry_countContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_d_retry_count
	return p
}

func (*D_retry_countContext) IsD_retry_countContext() {}

func NewD_retry_countContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *D_retry_countContext {
	var p = new(D_retry_countContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_d_retry_count

	return p
}

func (s *D_retry_countContext) GetParser() antlr.Parser { return s.parser }

func (s *D_retry_countContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *D_retry_countContext) RETRY_COUNT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserRETRY_COUNT, 0)
}

func (s *D_retry_countContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *D_retry_countContext) INT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserINT, 0)
}

func (s *D_retry_countContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *D_retry_countContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *D_retry_countContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *D_retry_countContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterD_retry_count(s)
	}
}

func (s *D_retry_countContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitD_retry_count(s)
	}
}

func (s *D_retry_countContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitD_retry_count(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) D_retry_count() (localctx ID_retry_countContext) {
	localctx = NewD_retry_countContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, tnsnamesParserRULE_d_retry_count)

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
		p.SetState(348)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(349)
		p.Match(tnsnamesParserRETRY_COUNT)
	}
	{
		p.SetState(350)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(351)
		p.Match(tnsnamesParserINT)
	}
	{
		p.SetState(352)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ID_tctContext is an interface to support dynamic dispatch.
type ID_tctContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsD_tctContext differentiates from other interfaces.
	IsD_tctContext()
}

type D_tctContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyD_tctContext() *D_tctContext {
	var p = new(D_tctContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_d_tct
	return p
}

func (*D_tctContext) IsD_tctContext() {}

func NewD_tctContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *D_tctContext {
	var p = new(D_tctContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_d_tct

	return p
}

func (s *D_tctContext) GetParser() antlr.Parser { return s.parser }

func (s *D_tctContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *D_tctContext) TCT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserTCT, 0)
}

func (s *D_tctContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *D_tctContext) INT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserINT, 0)
}

func (s *D_tctContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *D_tctContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *D_tctContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *D_tctContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterD_tct(s)
	}
}

func (s *D_tctContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitD_tct(s)
	}
}

func (s *D_tctContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitD_tct(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) D_tct() (localctx ID_tctContext) {
	localctx = NewD_tctContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, tnsnamesParserRULE_d_tct)

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
		p.SetState(354)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(355)
		p.Match(tnsnamesParserTCT)
	}
	{
		p.SetState(356)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(357)
		p.Match(tnsnamesParserINT)
	}
	{
		p.SetState(358)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IDs_parameterContext is an interface to support dynamic dispatch.
type IDs_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDs_parameterContext differentiates from other interfaces.
	IsDs_parameterContext()
}

type Ds_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDs_parameterContext() *Ds_parameterContext {
	var p = new(Ds_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_ds_parameter
	return p
}

func (*Ds_parameterContext) IsDs_parameterContext() {}

func NewDs_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ds_parameterContext {
	var p = new(Ds_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_ds_parameter

	return p
}

func (s *Ds_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Ds_parameterContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Ds_parameterContext) SSL_CERT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSSL_CERT, 0)
}

func (s *Ds_parameterContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Ds_parameterContext) DQ_STRING() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDQ_STRING, 0)
}

func (s *Ds_parameterContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Ds_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ds_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ds_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterDs_parameter(s)
	}
}

func (s *Ds_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitDs_parameter(s)
	}
}

func (s *Ds_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitDs_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Ds_parameter() (localctx IDs_parameterContext) {
	localctx = NewDs_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, tnsnamesParserRULE_ds_parameter)

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
		p.SetState(360)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(361)
		p.Match(tnsnamesParserSSL_CERT)
	}
	{
		p.SetState(362)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(363)
		p.Match(tnsnamesParserDQ_STRING)
	}
	{
		p.SetState(364)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IAddress_listContext is an interface to support dynamic dispatch.
type IAddress_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddress_listContext differentiates from other interfaces.
	IsAddress_listContext()
}

type Address_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddress_listContext() *Address_listContext {
	var p = new(Address_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_address_list
	return p
}

func (*Address_listContext) IsAddress_listContext() {}

func NewAddress_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Address_listContext {
	var p = new(Address_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_address_list

	return p
}

func (s *Address_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Address_listContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Address_listContext) ADDRESS_LIST() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserADDRESS_LIST, 0)
}

func (s *Address_listContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Address_listContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Address_listContext) AllAl_params() []IAl_paramsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAl_paramsContext)(nil)).Elem())
	var tst = make([]IAl_paramsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAl_paramsContext)
		}
	}

	return tst
}

func (s *Address_listContext) Al_params(i int) IAl_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAl_paramsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAl_paramsContext)
}

func (s *Address_listContext) AllAddress() []IAddressContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAddressContext)(nil)).Elem())
	var tst = make([]IAddressContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAddressContext)
		}
	}

	return tst
}

func (s *Address_listContext) Address(i int) IAddressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAddressContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAddressContext)
}

func (s *Address_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Address_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Address_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterAddress_list(s)
	}
}

func (s *Address_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitAddress_list(s)
	}
}

func (s *Address_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitAddress_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Address_list() (localctx IAddress_listContext) {
	localctx = NewAddress_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, tnsnamesParserRULE_address_list)
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
		p.SetState(366)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(367)
		p.Match(tnsnamesParserADDRESS_LIST)
	}
	{
		p.SetState(368)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(369)
			p.Al_params()
		}

	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(372)
				p.Address()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tnsnamesParserL_PAREN {
		{
			p.SetState(377)
			p.Al_params()
		}

	}
	{
		p.SetState(380)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IAl_paramsContext is an interface to support dynamic dispatch.
type IAl_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAl_paramsContext differentiates from other interfaces.
	IsAl_paramsContext()
}

type Al_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAl_paramsContext() *Al_paramsContext {
	var p = new(Al_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_al_params
	return p
}

func (*Al_paramsContext) IsAl_paramsContext() {}

func NewAl_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Al_paramsContext {
	var p = new(Al_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_al_params

	return p
}

func (s *Al_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Al_paramsContext) AllAl_parameter() []IAl_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAl_parameterContext)(nil)).Elem())
	var tst = make([]IAl_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAl_parameterContext)
		}
	}

	return tst
}

func (s *Al_paramsContext) Al_parameter(i int) IAl_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAl_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAl_parameterContext)
}

func (s *Al_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Al_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Al_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterAl_params(s)
	}
}

func (s *Al_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitAl_params(s)
	}
}

func (s *Al_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitAl_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Al_params() (localctx IAl_paramsContext) {
	localctx = NewAl_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, tnsnamesParserRULE_al_params)

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
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(382)
				p.Al_parameter()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// IAl_parameterContext is an interface to support dynamic dispatch.
type IAl_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAl_parameterContext differentiates from other interfaces.
	IsAl_parameterContext()
}

type Al_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAl_parameterContext() *Al_parameterContext {
	var p = new(Al_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_al_parameter
	return p
}

func (*Al_parameterContext) IsAl_parameterContext() {}

func NewAl_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Al_parameterContext {
	var p = new(Al_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_al_parameter

	return p
}

func (s *Al_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Al_parameterContext) Al_failover() IAl_failoverContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAl_failoverContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAl_failoverContext)
}

func (s *Al_parameterContext) Al_load_balance() IAl_load_balanceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAl_load_balanceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAl_load_balanceContext)
}

func (s *Al_parameterContext) Al_source_route() IAl_source_routeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAl_source_routeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAl_source_routeContext)
}

func (s *Al_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Al_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Al_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterAl_parameter(s)
	}
}

func (s *Al_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitAl_parameter(s)
	}
}

func (s *Al_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitAl_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Al_parameter() (localctx IAl_parameterContext) {
	localctx = NewAl_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, tnsnamesParserRULE_al_parameter)

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

	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(387)
			p.Al_failover()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(388)
			p.Al_load_balance()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(389)
			p.Al_source_route()
		}

	}

	return localctx
}

// IAl_failoverContext is an interface to support dynamic dispatch.
type IAl_failoverContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAl_failoverContext differentiates from other interfaces.
	IsAl_failoverContext()
}

type Al_failoverContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAl_failoverContext() *Al_failoverContext {
	var p = new(Al_failoverContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_al_failover
	return p
}

func (*Al_failoverContext) IsAl_failoverContext() {}

func NewAl_failoverContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Al_failoverContext {
	var p = new(Al_failoverContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_al_failover

	return p
}

func (s *Al_failoverContext) GetParser() antlr.Parser { return s.parser }

func (s *Al_failoverContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Al_failoverContext) FAILOVER() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserFAILOVER, 0)
}

func (s *Al_failoverContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Al_failoverContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Al_failoverContext) YES_NO() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserYES_NO, 0)
}

func (s *Al_failoverContext) ON_OFF() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserON_OFF, 0)
}

func (s *Al_failoverContext) TRUE_FALSE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserTRUE_FALSE, 0)
}

func (s *Al_failoverContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Al_failoverContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Al_failoverContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterAl_failover(s)
	}
}

func (s *Al_failoverContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitAl_failover(s)
	}
}

func (s *Al_failoverContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitAl_failover(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Al_failover() (localctx IAl_failoverContext) {
	localctx = NewAl_failoverContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, tnsnamesParserRULE_al_failover)
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
		p.SetState(392)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(393)
		p.Match(tnsnamesParserFAILOVER)
	}
	{
		p.SetState(394)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(395)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tnsnamesParserYES_NO)|(1<<tnsnamesParserON_OFF)|(1<<tnsnamesParserTRUE_FALSE))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(396)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IAl_load_balanceContext is an interface to support dynamic dispatch.
type IAl_load_balanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAl_load_balanceContext differentiates from other interfaces.
	IsAl_load_balanceContext()
}

type Al_load_balanceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAl_load_balanceContext() *Al_load_balanceContext {
	var p = new(Al_load_balanceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_al_load_balance
	return p
}

func (*Al_load_balanceContext) IsAl_load_balanceContext() {}

func NewAl_load_balanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Al_load_balanceContext {
	var p = new(Al_load_balanceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_al_load_balance

	return p
}

func (s *Al_load_balanceContext) GetParser() antlr.Parser { return s.parser }

func (s *Al_load_balanceContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Al_load_balanceContext) LOAD_BALANCE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserLOAD_BALANCE, 0)
}

func (s *Al_load_balanceContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Al_load_balanceContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Al_load_balanceContext) YES_NO() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserYES_NO, 0)
}

func (s *Al_load_balanceContext) ON_OFF() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserON_OFF, 0)
}

func (s *Al_load_balanceContext) TRUE_FALSE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserTRUE_FALSE, 0)
}

func (s *Al_load_balanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Al_load_balanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Al_load_balanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterAl_load_balance(s)
	}
}

func (s *Al_load_balanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitAl_load_balance(s)
	}
}

func (s *Al_load_balanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitAl_load_balance(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Al_load_balance() (localctx IAl_load_balanceContext) {
	localctx = NewAl_load_balanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, tnsnamesParserRULE_al_load_balance)
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
		p.SetState(398)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(399)
		p.Match(tnsnamesParserLOAD_BALANCE)
	}
	{
		p.SetState(400)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(401)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tnsnamesParserYES_NO)|(1<<tnsnamesParserON_OFF)|(1<<tnsnamesParserTRUE_FALSE))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(402)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IAl_source_routeContext is an interface to support dynamic dispatch.
type IAl_source_routeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAl_source_routeContext differentiates from other interfaces.
	IsAl_source_routeContext()
}

type Al_source_routeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAl_source_routeContext() *Al_source_routeContext {
	var p = new(Al_source_routeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_al_source_route
	return p
}

func (*Al_source_routeContext) IsAl_source_routeContext() {}

func NewAl_source_routeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Al_source_routeContext {
	var p = new(Al_source_routeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_al_source_route

	return p
}

func (s *Al_source_routeContext) GetParser() antlr.Parser { return s.parser }

func (s *Al_source_routeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Al_source_routeContext) SOURCE_ROUTE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSOURCE_ROUTE, 0)
}

func (s *Al_source_routeContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Al_source_routeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Al_source_routeContext) YES_NO() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserYES_NO, 0)
}

func (s *Al_source_routeContext) ON_OFF() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserON_OFF, 0)
}

func (s *Al_source_routeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Al_source_routeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Al_source_routeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterAl_source_route(s)
	}
}

func (s *Al_source_routeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitAl_source_route(s)
	}
}

func (s *Al_source_routeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitAl_source_route(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Al_source_route() (localctx IAl_source_routeContext) {
	localctx = NewAl_source_routeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, tnsnamesParserRULE_al_source_route)
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
		p.SetState(404)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(405)
		p.Match(tnsnamesParserSOURCE_ROUTE)
	}
	{
		p.SetState(406)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(407)
	_la = p.GetTokenStream().LA(1)

	if !(_la == tnsnamesParserYES_NO || _la == tnsnamesParserON_OFF) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(408)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IAddressContext is an interface to support dynamic dispatch.
type IAddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAddressContext differentiates from other interfaces.
	IsAddressContext()
}

type AddressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAddressContext() *AddressContext {
	var p = new(AddressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_address
	return p
}

func (*AddressContext) IsAddressContext() {}

func NewAddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AddressContext {
	var p = new(AddressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_address

	return p
}

func (s *AddressContext) GetParser() antlr.Parser { return s.parser }

func (s *AddressContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *AddressContext) ADDRESS() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserADDRESS, 0)
}

func (s *AddressContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *AddressContext) Protocol_info() IProtocol_infoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtocol_infoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProtocol_infoContext)
}

func (s *AddressContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *AddressContext) A_params() IA_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IA_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IA_paramsContext)
}

func (s *AddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AddressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterAddress(s)
	}
}

func (s *AddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitAddress(s)
	}
}

func (s *AddressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitAddress(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Address() (localctx IAddressContext) {
	localctx = NewAddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, tnsnamesParserRULE_address)
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
		p.SetState(410)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(411)
		p.Match(tnsnamesParserADDRESS)
	}
	{
		p.SetState(412)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(413)
		p.Protocol_info()
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tnsnamesParserL_PAREN {
		{
			p.SetState(414)
			p.A_params()
		}

	}
	{
		p.SetState(417)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IA_paramsContext is an interface to support dynamic dispatch.
type IA_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsA_paramsContext differentiates from other interfaces.
	IsA_paramsContext()
}

type A_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_paramsContext() *A_paramsContext {
	var p = new(A_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_a_params
	return p
}

func (*A_paramsContext) IsA_paramsContext() {}

func NewA_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_paramsContext {
	var p = new(A_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_a_params

	return p
}

func (s *A_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *A_paramsContext) AllA_parameter() []IA_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IA_parameterContext)(nil)).Elem())
	var tst = make([]IA_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IA_parameterContext)
		}
	}

	return tst
}

func (s *A_paramsContext) A_parameter(i int) IA_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IA_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IA_parameterContext)
}

func (s *A_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterA_params(s)
	}
}

func (s *A_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitA_params(s)
	}
}

func (s *A_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitA_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) A_params() (localctx IA_paramsContext) {
	localctx = NewA_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, tnsnamesParserRULE_a_params)
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
	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tnsnamesParserL_PAREN {
		{
			p.SetState(419)
			p.A_parameter()
		}

		p.SetState(422)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IA_parameterContext is an interface to support dynamic dispatch.
type IA_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsA_parameterContext differentiates from other interfaces.
	IsA_parameterContext()
}

type A_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyA_parameterContext() *A_parameterContext {
	var p = new(A_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_a_parameter
	return p
}

func (*A_parameterContext) IsA_parameterContext() {}

func NewA_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *A_parameterContext {
	var p = new(A_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_a_parameter

	return p
}

func (s *A_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *A_parameterContext) D_send_buf() ID_send_bufContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_send_bufContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ID_send_bufContext)
}

func (s *A_parameterContext) D_recv_buf() ID_recv_bufContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ID_recv_bufContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ID_recv_bufContext)
}

func (s *A_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *A_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *A_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterA_parameter(s)
	}
}

func (s *A_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitA_parameter(s)
	}
}

func (s *A_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitA_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) A_parameter() (localctx IA_parameterContext) {
	localctx = NewA_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, tnsnamesParserRULE_a_parameter)

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

	p.SetState(426)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(424)
			p.D_send_buf()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(425)
			p.D_recv_buf()
		}

	}

	return localctx
}

// IProtocol_infoContext is an interface to support dynamic dispatch.
type IProtocol_infoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProtocol_infoContext differentiates from other interfaces.
	IsProtocol_infoContext()
}

type Protocol_infoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtocol_infoContext() *Protocol_infoContext {
	var p = new(Protocol_infoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_protocol_info
	return p
}

func (*Protocol_infoContext) IsProtocol_infoContext() {}

func NewProtocol_infoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Protocol_infoContext {
	var p = new(Protocol_infoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_protocol_info

	return p
}

func (s *Protocol_infoContext) GetParser() antlr.Parser { return s.parser }

func (s *Protocol_infoContext) Tcp_protocol() ITcp_protocolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITcp_protocolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITcp_protocolContext)
}

func (s *Protocol_infoContext) Ipc_protocol() IIpc_protocolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpc_protocolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpc_protocolContext)
}

func (s *Protocol_infoContext) Spx_protocol() ISpx_protocolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpx_protocolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpx_protocolContext)
}

func (s *Protocol_infoContext) Nmp_protocol() INmp_protocolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INmp_protocolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INmp_protocolContext)
}

func (s *Protocol_infoContext) Beq_protocol() IBeq_protocolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeq_protocolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeq_protocolContext)
}

func (s *Protocol_infoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Protocol_infoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Protocol_infoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterProtocol_info(s)
	}
}

func (s *Protocol_infoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitProtocol_info(s)
	}
}

func (s *Protocol_infoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitProtocol_info(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Protocol_info() (localctx IProtocol_infoContext) {
	localctx = NewProtocol_infoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, tnsnamesParserRULE_protocol_info)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(428)
			p.Tcp_protocol()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(429)
			p.Ipc_protocol()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(430)
			p.Spx_protocol()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(431)
			p.Nmp_protocol()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(432)
			p.Beq_protocol()
		}

	}

	return localctx
}

// ITcp_protocolContext is an interface to support dynamic dispatch.
type ITcp_protocolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTcp_protocolContext differentiates from other interfaces.
	IsTcp_protocolContext()
}

type Tcp_protocolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTcp_protocolContext() *Tcp_protocolContext {
	var p = new(Tcp_protocolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_tcp_protocol
	return p
}

func (*Tcp_protocolContext) IsTcp_protocolContext() {}

func NewTcp_protocolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tcp_protocolContext {
	var p = new(Tcp_protocolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_tcp_protocol

	return p
}

func (s *Tcp_protocolContext) GetParser() antlr.Parser { return s.parser }

func (s *Tcp_protocolContext) Tcp_params() ITcp_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITcp_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITcp_paramsContext)
}

func (s *Tcp_protocolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tcp_protocolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tcp_protocolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterTcp_protocol(s)
	}
}

func (s *Tcp_protocolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitTcp_protocol(s)
	}
}

func (s *Tcp_protocolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitTcp_protocol(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Tcp_protocol() (localctx ITcp_protocolContext) {
	localctx = NewTcp_protocolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, tnsnamesParserRULE_tcp_protocol)

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
		p.SetState(435)
		p.Tcp_params()
	}

	return localctx
}

// ITcp_paramsContext is an interface to support dynamic dispatch.
type ITcp_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTcp_paramsContext differentiates from other interfaces.
	IsTcp_paramsContext()
}

type Tcp_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTcp_paramsContext() *Tcp_paramsContext {
	var p = new(Tcp_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_tcp_params
	return p
}

func (*Tcp_paramsContext) IsTcp_paramsContext() {}

func NewTcp_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tcp_paramsContext {
	var p = new(Tcp_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_tcp_params

	return p
}

func (s *Tcp_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Tcp_paramsContext) AllTcp_parameter() []ITcp_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITcp_parameterContext)(nil)).Elem())
	var tst = make([]ITcp_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITcp_parameterContext)
		}
	}

	return tst
}

func (s *Tcp_paramsContext) Tcp_parameter(i int) ITcp_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITcp_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITcp_parameterContext)
}

func (s *Tcp_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tcp_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tcp_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterTcp_params(s)
	}
}

func (s *Tcp_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitTcp_params(s)
	}
}

func (s *Tcp_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitTcp_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Tcp_params() (localctx ITcp_paramsContext) {
	localctx = NewTcp_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, tnsnamesParserRULE_tcp_params)

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
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(437)
				p.Tcp_parameter()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}

	return localctx
}

// ITcp_parameterContext is an interface to support dynamic dispatch.
type ITcp_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTcp_parameterContext differentiates from other interfaces.
	IsTcp_parameterContext()
}

type Tcp_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTcp_parameterContext() *Tcp_parameterContext {
	var p = new(Tcp_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_tcp_parameter
	return p
}

func (*Tcp_parameterContext) IsTcp_parameterContext() {}

func NewTcp_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tcp_parameterContext {
	var p = new(Tcp_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_tcp_parameter

	return p
}

func (s *Tcp_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Tcp_parameterContext) Tcp_host() ITcp_hostContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITcp_hostContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITcp_hostContext)
}

func (s *Tcp_parameterContext) Tcp_port() ITcp_portContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITcp_portContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITcp_portContext)
}

func (s *Tcp_parameterContext) Tcp_tcp() ITcp_tcpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITcp_tcpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITcp_tcpContext)
}

func (s *Tcp_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tcp_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tcp_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterTcp_parameter(s)
	}
}

func (s *Tcp_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitTcp_parameter(s)
	}
}

func (s *Tcp_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitTcp_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Tcp_parameter() (localctx ITcp_parameterContext) {
	localctx = NewTcp_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, tnsnamesParserRULE_tcp_parameter)

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

	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(442)
			p.Tcp_host()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(443)
			p.Tcp_port()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(444)
			p.Tcp_tcp()
		}

	}

	return localctx
}

// ITcp_hostContext is an interface to support dynamic dispatch.
type ITcp_hostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTcp_hostContext differentiates from other interfaces.
	IsTcp_hostContext()
}

type Tcp_hostContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTcp_hostContext() *Tcp_hostContext {
	var p = new(Tcp_hostContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_tcp_host
	return p
}

func (*Tcp_hostContext) IsTcp_hostContext() {}

func NewTcp_hostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tcp_hostContext {
	var p = new(Tcp_hostContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_tcp_host

	return p
}

func (s *Tcp_hostContext) GetParser() antlr.Parser { return s.parser }

func (s *Tcp_hostContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Tcp_hostContext) HOST() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserHOST, 0)
}

func (s *Tcp_hostContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Tcp_hostContext) Host() IHostContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHostContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHostContext)
}

func (s *Tcp_hostContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Tcp_hostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tcp_hostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tcp_hostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterTcp_host(s)
	}
}

func (s *Tcp_hostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitTcp_host(s)
	}
}

func (s *Tcp_hostContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitTcp_host(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Tcp_host() (localctx ITcp_hostContext) {
	localctx = NewTcp_hostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, tnsnamesParserRULE_tcp_host)

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
		p.SetState(447)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(448)
		p.Match(tnsnamesParserHOST)
	}
	{
		p.SetState(449)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(450)
		p.Host()
	}
	{
		p.SetState(451)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ITcp_portContext is an interface to support dynamic dispatch.
type ITcp_portContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTcp_portContext differentiates from other interfaces.
	IsTcp_portContext()
}

type Tcp_portContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTcp_portContext() *Tcp_portContext {
	var p = new(Tcp_portContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_tcp_port
	return p
}

func (*Tcp_portContext) IsTcp_portContext() {}

func NewTcp_portContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tcp_portContext {
	var p = new(Tcp_portContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_tcp_port

	return p
}

func (s *Tcp_portContext) GetParser() antlr.Parser { return s.parser }

func (s *Tcp_portContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Tcp_portContext) PORT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserPORT, 0)
}

func (s *Tcp_portContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Tcp_portContext) Port() IPortContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPortContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPortContext)
}

func (s *Tcp_portContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Tcp_portContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tcp_portContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tcp_portContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterTcp_port(s)
	}
}

func (s *Tcp_portContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitTcp_port(s)
	}
}

func (s *Tcp_portContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitTcp_port(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Tcp_port() (localctx ITcp_portContext) {
	localctx = NewTcp_portContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, tnsnamesParserRULE_tcp_port)

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
		p.SetState(453)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(454)
		p.Match(tnsnamesParserPORT)
	}
	{
		p.SetState(455)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(456)
		p.Port()
	}
	{
		p.SetState(457)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ITcp_tcpContext is an interface to support dynamic dispatch.
type ITcp_tcpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTcp_tcpContext differentiates from other interfaces.
	IsTcp_tcpContext()
}

type Tcp_tcpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTcp_tcpContext() *Tcp_tcpContext {
	var p = new(Tcp_tcpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_tcp_tcp
	return p
}

func (*Tcp_tcpContext) IsTcp_tcpContext() {}

func NewTcp_tcpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tcp_tcpContext {
	var p = new(Tcp_tcpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_tcp_tcp

	return p
}

func (s *Tcp_tcpContext) GetParser() antlr.Parser { return s.parser }

func (s *Tcp_tcpContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Tcp_tcpContext) PROTOCOL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserPROTOCOL, 0)
}

func (s *Tcp_tcpContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Tcp_tcpContext) TCP() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserTCP, 0)
}

func (s *Tcp_tcpContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Tcp_tcpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tcp_tcpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tcp_tcpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterTcp_tcp(s)
	}
}

func (s *Tcp_tcpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitTcp_tcp(s)
	}
}

func (s *Tcp_tcpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitTcp_tcp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Tcp_tcp() (localctx ITcp_tcpContext) {
	localctx = NewTcp_tcpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, tnsnamesParserRULE_tcp_tcp)

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
		p.SetState(459)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(460)
		p.Match(tnsnamesParserPROTOCOL)
	}
	{
		p.SetState(461)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(462)
		p.Match(tnsnamesParserTCP)
	}
	{
		p.SetState(463)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IHostContext is an interface to support dynamic dispatch.
type IHostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHostContext differentiates from other interfaces.
	IsHostContext()
}

type HostContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHostContext() *HostContext {
	var p = new(HostContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_host
	return p
}

func (*HostContext) IsHostContext() {}

func NewHostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HostContext {
	var p = new(HostContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_host

	return p
}

func (s *HostContext) GetParser() antlr.Parser { return s.parser }

func (s *HostContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserID)
}

func (s *HostContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, i)
}

func (s *HostContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserDOT)
}

func (s *HostContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDOT, i)
}

func (s *HostContext) IP() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserIP, 0)
}

func (s *HostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterHost(s)
	}
}

func (s *HostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitHost(s)
	}
}

func (s *HostContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitHost(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Host() (localctx IHostContext) {
	localctx = NewHostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, tnsnamesParserRULE_host)
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

	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(465)
			p.Match(tnsnamesParserID)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(466)
			p.Match(tnsnamesParserID)
		}
		p.SetState(469)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == tnsnamesParserDOT {
			{
				p.SetState(467)
				p.Match(tnsnamesParserDOT)
			}
			{
				p.SetState(468)
				p.Match(tnsnamesParserID)
			}

			p.SetState(471)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(473)
			p.Match(tnsnamesParserIP)
		}

	}

	return localctx
}

// IPortContext is an interface to support dynamic dispatch.
type IPortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPortContext differentiates from other interfaces.
	IsPortContext()
}

type PortContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortContext() *PortContext {
	var p = new(PortContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_port
	return p
}

func (*PortContext) IsPortContext() {}

func NewPortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortContext {
	var p = new(PortContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_port

	return p
}

func (s *PortContext) GetParser() antlr.Parser { return s.parser }

func (s *PortContext) INT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserINT, 0)
}

func (s *PortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterPort(s)
	}
}

func (s *PortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitPort(s)
	}
}

func (s *PortContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitPort(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Port() (localctx IPortContext) {
	localctx = NewPortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, tnsnamesParserRULE_port)

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
		p.Match(tnsnamesParserINT)
	}

	return localctx
}

// IIpc_protocolContext is an interface to support dynamic dispatch.
type IIpc_protocolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpc_protocolContext differentiates from other interfaces.
	IsIpc_protocolContext()
}

type Ipc_protocolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpc_protocolContext() *Ipc_protocolContext {
	var p = new(Ipc_protocolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_ipc_protocol
	return p
}

func (*Ipc_protocolContext) IsIpc_protocolContext() {}

func NewIpc_protocolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipc_protocolContext {
	var p = new(Ipc_protocolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_ipc_protocol

	return p
}

func (s *Ipc_protocolContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipc_protocolContext) Ipc_params() IIpc_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpc_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpc_paramsContext)
}

func (s *Ipc_protocolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipc_protocolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipc_protocolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterIpc_protocol(s)
	}
}

func (s *Ipc_protocolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitIpc_protocol(s)
	}
}

func (s *Ipc_protocolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitIpc_protocol(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Ipc_protocol() (localctx IIpc_protocolContext) {
	localctx = NewIpc_protocolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, tnsnamesParserRULE_ipc_protocol)

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
		p.SetState(478)
		p.Ipc_params()
	}

	return localctx
}

// IIpc_paramsContext is an interface to support dynamic dispatch.
type IIpc_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpc_paramsContext differentiates from other interfaces.
	IsIpc_paramsContext()
}

type Ipc_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpc_paramsContext() *Ipc_paramsContext {
	var p = new(Ipc_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_ipc_params
	return p
}

func (*Ipc_paramsContext) IsIpc_paramsContext() {}

func NewIpc_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipc_paramsContext {
	var p = new(Ipc_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_ipc_params

	return p
}

func (s *Ipc_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipc_paramsContext) AllIpc_parameter() []IIpc_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIpc_parameterContext)(nil)).Elem())
	var tst = make([]IIpc_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIpc_parameterContext)
		}
	}

	return tst
}

func (s *Ipc_paramsContext) Ipc_parameter(i int) IIpc_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpc_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIpc_parameterContext)
}

func (s *Ipc_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipc_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipc_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterIpc_params(s)
	}
}

func (s *Ipc_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitIpc_params(s)
	}
}

func (s *Ipc_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitIpc_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Ipc_params() (localctx IIpc_paramsContext) {
	localctx = NewIpc_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, tnsnamesParserRULE_ipc_params)

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
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(480)
				p.Ipc_parameter()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(483)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext())
	}

	return localctx
}

// IIpc_parameterContext is an interface to support dynamic dispatch.
type IIpc_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpc_parameterContext differentiates from other interfaces.
	IsIpc_parameterContext()
}

type Ipc_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpc_parameterContext() *Ipc_parameterContext {
	var p = new(Ipc_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_ipc_parameter
	return p
}

func (*Ipc_parameterContext) IsIpc_parameterContext() {}

func NewIpc_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipc_parameterContext {
	var p = new(Ipc_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_ipc_parameter

	return p
}

func (s *Ipc_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipc_parameterContext) Ipc_ipc() IIpc_ipcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpc_ipcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpc_ipcContext)
}

func (s *Ipc_parameterContext) Ipc_key() IIpc_keyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIpc_keyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIpc_keyContext)
}

func (s *Ipc_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipc_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipc_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterIpc_parameter(s)
	}
}

func (s *Ipc_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitIpc_parameter(s)
	}
}

func (s *Ipc_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitIpc_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Ipc_parameter() (localctx IIpc_parameterContext) {
	localctx = NewIpc_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, tnsnamesParserRULE_ipc_parameter)

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

	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(485)
			p.Ipc_ipc()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(486)
			p.Ipc_key()
		}

	}

	return localctx
}

// IIpc_ipcContext is an interface to support dynamic dispatch.
type IIpc_ipcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpc_ipcContext differentiates from other interfaces.
	IsIpc_ipcContext()
}

type Ipc_ipcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpc_ipcContext() *Ipc_ipcContext {
	var p = new(Ipc_ipcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_ipc_ipc
	return p
}

func (*Ipc_ipcContext) IsIpc_ipcContext() {}

func NewIpc_ipcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipc_ipcContext {
	var p = new(Ipc_ipcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_ipc_ipc

	return p
}

func (s *Ipc_ipcContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipc_ipcContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Ipc_ipcContext) PROTOCOL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserPROTOCOL, 0)
}

func (s *Ipc_ipcContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Ipc_ipcContext) IPC() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserIPC, 0)
}

func (s *Ipc_ipcContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Ipc_ipcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipc_ipcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipc_ipcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterIpc_ipc(s)
	}
}

func (s *Ipc_ipcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitIpc_ipc(s)
	}
}

func (s *Ipc_ipcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitIpc_ipc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Ipc_ipc() (localctx IIpc_ipcContext) {
	localctx = NewIpc_ipcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, tnsnamesParserRULE_ipc_ipc)

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
		p.SetState(489)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(490)
		p.Match(tnsnamesParserPROTOCOL)
	}
	{
		p.SetState(491)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(492)
		p.Match(tnsnamesParserIPC)
	}
	{
		p.SetState(493)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IIpc_keyContext is an interface to support dynamic dispatch.
type IIpc_keyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIpc_keyContext differentiates from other interfaces.
	IsIpc_keyContext()
}

type Ipc_keyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIpc_keyContext() *Ipc_keyContext {
	var p = new(Ipc_keyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_ipc_key
	return p
}

func (*Ipc_keyContext) IsIpc_keyContext() {}

func NewIpc_keyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ipc_keyContext {
	var p = new(Ipc_keyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_ipc_key

	return p
}

func (s *Ipc_keyContext) GetParser() antlr.Parser { return s.parser }

func (s *Ipc_keyContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Ipc_keyContext) KEY() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserKEY, 0)
}

func (s *Ipc_keyContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Ipc_keyContext) ID() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, 0)
}

func (s *Ipc_keyContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Ipc_keyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ipc_keyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ipc_keyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterIpc_key(s)
	}
}

func (s *Ipc_keyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitIpc_key(s)
	}
}

func (s *Ipc_keyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitIpc_key(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Ipc_key() (localctx IIpc_keyContext) {
	localctx = NewIpc_keyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, tnsnamesParserRULE_ipc_key)

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
		p.SetState(495)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(496)
		p.Match(tnsnamesParserKEY)
	}
	{
		p.SetState(497)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(498)
		p.Match(tnsnamesParserID)
	}
	{
		p.SetState(499)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ISpx_protocolContext is an interface to support dynamic dispatch.
type ISpx_protocolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpx_protocolContext differentiates from other interfaces.
	IsSpx_protocolContext()
}

type Spx_protocolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpx_protocolContext() *Spx_protocolContext {
	var p = new(Spx_protocolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_spx_protocol
	return p
}

func (*Spx_protocolContext) IsSpx_protocolContext() {}

func NewSpx_protocolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Spx_protocolContext {
	var p = new(Spx_protocolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_spx_protocol

	return p
}

func (s *Spx_protocolContext) GetParser() antlr.Parser { return s.parser }

func (s *Spx_protocolContext) Spx_params() ISpx_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpx_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpx_paramsContext)
}

func (s *Spx_protocolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Spx_protocolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Spx_protocolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterSpx_protocol(s)
	}
}

func (s *Spx_protocolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitSpx_protocol(s)
	}
}

func (s *Spx_protocolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitSpx_protocol(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Spx_protocol() (localctx ISpx_protocolContext) {
	localctx = NewSpx_protocolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, tnsnamesParserRULE_spx_protocol)

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
		p.SetState(501)
		p.Spx_params()
	}

	return localctx
}

// ISpx_paramsContext is an interface to support dynamic dispatch.
type ISpx_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpx_paramsContext differentiates from other interfaces.
	IsSpx_paramsContext()
}

type Spx_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpx_paramsContext() *Spx_paramsContext {
	var p = new(Spx_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_spx_params
	return p
}

func (*Spx_paramsContext) IsSpx_paramsContext() {}

func NewSpx_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Spx_paramsContext {
	var p = new(Spx_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_spx_params

	return p
}

func (s *Spx_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Spx_paramsContext) AllSpx_parameter() []ISpx_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISpx_parameterContext)(nil)).Elem())
	var tst = make([]ISpx_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISpx_parameterContext)
		}
	}

	return tst
}

func (s *Spx_paramsContext) Spx_parameter(i int) ISpx_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpx_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISpx_parameterContext)
}

func (s *Spx_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Spx_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Spx_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterSpx_params(s)
	}
}

func (s *Spx_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitSpx_params(s)
	}
}

func (s *Spx_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitSpx_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Spx_params() (localctx ISpx_paramsContext) {
	localctx = NewSpx_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, tnsnamesParserRULE_spx_params)

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
	p.SetState(504)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(503)
				p.Spx_parameter()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(506)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext())
	}

	return localctx
}

// ISpx_parameterContext is an interface to support dynamic dispatch.
type ISpx_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpx_parameterContext differentiates from other interfaces.
	IsSpx_parameterContext()
}

type Spx_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpx_parameterContext() *Spx_parameterContext {
	var p = new(Spx_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_spx_parameter
	return p
}

func (*Spx_parameterContext) IsSpx_parameterContext() {}

func NewSpx_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Spx_parameterContext {
	var p = new(Spx_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_spx_parameter

	return p
}

func (s *Spx_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Spx_parameterContext) Spx_spx() ISpx_spxContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpx_spxContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpx_spxContext)
}

func (s *Spx_parameterContext) Spx_service() ISpx_serviceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpx_serviceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpx_serviceContext)
}

func (s *Spx_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Spx_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Spx_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterSpx_parameter(s)
	}
}

func (s *Spx_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitSpx_parameter(s)
	}
}

func (s *Spx_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitSpx_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Spx_parameter() (localctx ISpx_parameterContext) {
	localctx = NewSpx_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, tnsnamesParserRULE_spx_parameter)

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

	p.SetState(510)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(508)
			p.Spx_spx()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(509)
			p.Spx_service()
		}

	}

	return localctx
}

// ISpx_spxContext is an interface to support dynamic dispatch.
type ISpx_spxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpx_spxContext differentiates from other interfaces.
	IsSpx_spxContext()
}

type Spx_spxContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpx_spxContext() *Spx_spxContext {
	var p = new(Spx_spxContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_spx_spx
	return p
}

func (*Spx_spxContext) IsSpx_spxContext() {}

func NewSpx_spxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Spx_spxContext {
	var p = new(Spx_spxContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_spx_spx

	return p
}

func (s *Spx_spxContext) GetParser() antlr.Parser { return s.parser }

func (s *Spx_spxContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Spx_spxContext) PROTOCOL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserPROTOCOL, 0)
}

func (s *Spx_spxContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Spx_spxContext) SPX() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSPX, 0)
}

func (s *Spx_spxContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Spx_spxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Spx_spxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Spx_spxContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterSpx_spx(s)
	}
}

func (s *Spx_spxContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitSpx_spx(s)
	}
}

func (s *Spx_spxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitSpx_spx(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Spx_spx() (localctx ISpx_spxContext) {
	localctx = NewSpx_spxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, tnsnamesParserRULE_spx_spx)

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
		p.SetState(512)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(513)
		p.Match(tnsnamesParserPROTOCOL)
	}
	{
		p.SetState(514)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(515)
		p.Match(tnsnamesParserSPX)
	}
	{
		p.SetState(516)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ISpx_serviceContext is an interface to support dynamic dispatch.
type ISpx_serviceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpx_serviceContext differentiates from other interfaces.
	IsSpx_serviceContext()
}

type Spx_serviceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpx_serviceContext() *Spx_serviceContext {
	var p = new(Spx_serviceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_spx_service
	return p
}

func (*Spx_serviceContext) IsSpx_serviceContext() {}

func NewSpx_serviceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Spx_serviceContext {
	var p = new(Spx_serviceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_spx_service

	return p
}

func (s *Spx_serviceContext) GetParser() antlr.Parser { return s.parser }

func (s *Spx_serviceContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Spx_serviceContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSERVICE, 0)
}

func (s *Spx_serviceContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Spx_serviceContext) ID() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, 0)
}

func (s *Spx_serviceContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Spx_serviceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Spx_serviceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Spx_serviceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterSpx_service(s)
	}
}

func (s *Spx_serviceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitSpx_service(s)
	}
}

func (s *Spx_serviceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitSpx_service(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Spx_service() (localctx ISpx_serviceContext) {
	localctx = NewSpx_serviceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, tnsnamesParserRULE_spx_service)

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
		p.SetState(518)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(519)
		p.Match(tnsnamesParserSERVICE)
	}
	{
		p.SetState(520)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(521)
		p.Match(tnsnamesParserID)
	}
	{
		p.SetState(522)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// INmp_protocolContext is an interface to support dynamic dispatch.
type INmp_protocolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNmp_protocolContext differentiates from other interfaces.
	IsNmp_protocolContext()
}

type Nmp_protocolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNmp_protocolContext() *Nmp_protocolContext {
	var p = new(Nmp_protocolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_nmp_protocol
	return p
}

func (*Nmp_protocolContext) IsNmp_protocolContext() {}

func NewNmp_protocolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nmp_protocolContext {
	var p = new(Nmp_protocolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_nmp_protocol

	return p
}

func (s *Nmp_protocolContext) GetParser() antlr.Parser { return s.parser }

func (s *Nmp_protocolContext) Nmp_params() INmp_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INmp_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INmp_paramsContext)
}

func (s *Nmp_protocolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nmp_protocolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nmp_protocolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterNmp_protocol(s)
	}
}

func (s *Nmp_protocolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitNmp_protocol(s)
	}
}

func (s *Nmp_protocolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitNmp_protocol(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Nmp_protocol() (localctx INmp_protocolContext) {
	localctx = NewNmp_protocolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, tnsnamesParserRULE_nmp_protocol)

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
		p.SetState(524)
		p.Nmp_params()
	}

	return localctx
}

// INmp_paramsContext is an interface to support dynamic dispatch.
type INmp_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNmp_paramsContext differentiates from other interfaces.
	IsNmp_paramsContext()
}

type Nmp_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNmp_paramsContext() *Nmp_paramsContext {
	var p = new(Nmp_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_nmp_params
	return p
}

func (*Nmp_paramsContext) IsNmp_paramsContext() {}

func NewNmp_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nmp_paramsContext {
	var p = new(Nmp_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_nmp_params

	return p
}

func (s *Nmp_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Nmp_paramsContext) AllNmp_parameter() []INmp_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INmp_parameterContext)(nil)).Elem())
	var tst = make([]INmp_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INmp_parameterContext)
		}
	}

	return tst
}

func (s *Nmp_paramsContext) Nmp_parameter(i int) INmp_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INmp_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INmp_parameterContext)
}

func (s *Nmp_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nmp_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nmp_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterNmp_params(s)
	}
}

func (s *Nmp_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitNmp_params(s)
	}
}

func (s *Nmp_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitNmp_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Nmp_params() (localctx INmp_paramsContext) {
	localctx = NewNmp_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, tnsnamesParserRULE_nmp_params)

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
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(526)
				p.Nmp_parameter()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(529)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}

	return localctx
}

// INmp_parameterContext is an interface to support dynamic dispatch.
type INmp_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNmp_parameterContext differentiates from other interfaces.
	IsNmp_parameterContext()
}

type Nmp_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNmp_parameterContext() *Nmp_parameterContext {
	var p = new(Nmp_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_nmp_parameter
	return p
}

func (*Nmp_parameterContext) IsNmp_parameterContext() {}

func NewNmp_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nmp_parameterContext {
	var p = new(Nmp_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_nmp_parameter

	return p
}

func (s *Nmp_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Nmp_parameterContext) Nmp_nmp() INmp_nmpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INmp_nmpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INmp_nmpContext)
}

func (s *Nmp_parameterContext) Nmp_server() INmp_serverContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INmp_serverContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INmp_serverContext)
}

func (s *Nmp_parameterContext) Nmp_pipe() INmp_pipeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INmp_pipeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INmp_pipeContext)
}

func (s *Nmp_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nmp_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nmp_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterNmp_parameter(s)
	}
}

func (s *Nmp_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitNmp_parameter(s)
	}
}

func (s *Nmp_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitNmp_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Nmp_parameter() (localctx INmp_parameterContext) {
	localctx = NewNmp_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, tnsnamesParserRULE_nmp_parameter)

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

	p.SetState(534)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(531)
			p.Nmp_nmp()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(532)
			p.Nmp_server()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(533)
			p.Nmp_pipe()
		}

	}

	return localctx
}

// INmp_nmpContext is an interface to support dynamic dispatch.
type INmp_nmpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNmp_nmpContext differentiates from other interfaces.
	IsNmp_nmpContext()
}

type Nmp_nmpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNmp_nmpContext() *Nmp_nmpContext {
	var p = new(Nmp_nmpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_nmp_nmp
	return p
}

func (*Nmp_nmpContext) IsNmp_nmpContext() {}

func NewNmp_nmpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nmp_nmpContext {
	var p = new(Nmp_nmpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_nmp_nmp

	return p
}

func (s *Nmp_nmpContext) GetParser() antlr.Parser { return s.parser }

func (s *Nmp_nmpContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Nmp_nmpContext) PROTOCOL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserPROTOCOL, 0)
}

func (s *Nmp_nmpContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Nmp_nmpContext) NMP() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserNMP, 0)
}

func (s *Nmp_nmpContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Nmp_nmpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nmp_nmpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nmp_nmpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterNmp_nmp(s)
	}
}

func (s *Nmp_nmpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitNmp_nmp(s)
	}
}

func (s *Nmp_nmpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitNmp_nmp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Nmp_nmp() (localctx INmp_nmpContext) {
	localctx = NewNmp_nmpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, tnsnamesParserRULE_nmp_nmp)

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
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(537)
		p.Match(tnsnamesParserPROTOCOL)
	}
	{
		p.SetState(538)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(539)
		p.Match(tnsnamesParserNMP)
	}
	{
		p.SetState(540)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// INmp_serverContext is an interface to support dynamic dispatch.
type INmp_serverContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNmp_serverContext differentiates from other interfaces.
	IsNmp_serverContext()
}

type Nmp_serverContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNmp_serverContext() *Nmp_serverContext {
	var p = new(Nmp_serverContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_nmp_server
	return p
}

func (*Nmp_serverContext) IsNmp_serverContext() {}

func NewNmp_serverContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nmp_serverContext {
	var p = new(Nmp_serverContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_nmp_server

	return p
}

func (s *Nmp_serverContext) GetParser() antlr.Parser { return s.parser }

func (s *Nmp_serverContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Nmp_serverContext) SERVER() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSERVER, 0)
}

func (s *Nmp_serverContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Nmp_serverContext) ID() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, 0)
}

func (s *Nmp_serverContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Nmp_serverContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nmp_serverContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nmp_serverContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterNmp_server(s)
	}
}

func (s *Nmp_serverContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitNmp_server(s)
	}
}

func (s *Nmp_serverContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitNmp_server(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Nmp_server() (localctx INmp_serverContext) {
	localctx = NewNmp_serverContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, tnsnamesParserRULE_nmp_server)

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
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(543)
		p.Match(tnsnamesParserSERVER)
	}
	{
		p.SetState(544)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(545)
		p.Match(tnsnamesParserID)
	}
	{
		p.SetState(546)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// INmp_pipeContext is an interface to support dynamic dispatch.
type INmp_pipeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNmp_pipeContext differentiates from other interfaces.
	IsNmp_pipeContext()
}

type Nmp_pipeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNmp_pipeContext() *Nmp_pipeContext {
	var p = new(Nmp_pipeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_nmp_pipe
	return p
}

func (*Nmp_pipeContext) IsNmp_pipeContext() {}

func NewNmp_pipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nmp_pipeContext {
	var p = new(Nmp_pipeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_nmp_pipe

	return p
}

func (s *Nmp_pipeContext) GetParser() antlr.Parser { return s.parser }

func (s *Nmp_pipeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Nmp_pipeContext) PIPE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserPIPE, 0)
}

func (s *Nmp_pipeContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Nmp_pipeContext) ID() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, 0)
}

func (s *Nmp_pipeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Nmp_pipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nmp_pipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nmp_pipeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterNmp_pipe(s)
	}
}

func (s *Nmp_pipeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitNmp_pipe(s)
	}
}

func (s *Nmp_pipeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitNmp_pipe(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Nmp_pipe() (localctx INmp_pipeContext) {
	localctx = NewNmp_pipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, tnsnamesParserRULE_nmp_pipe)

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
		p.SetState(548)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(549)
		p.Match(tnsnamesParserPIPE)
	}
	{
		p.SetState(550)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(551)
		p.Match(tnsnamesParserID)
	}
	{
		p.SetState(552)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IBeq_protocolContext is an interface to support dynamic dispatch.
type IBeq_protocolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBeq_protocolContext differentiates from other interfaces.
	IsBeq_protocolContext()
}

type Beq_protocolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeq_protocolContext() *Beq_protocolContext {
	var p = new(Beq_protocolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_beq_protocol
	return p
}

func (*Beq_protocolContext) IsBeq_protocolContext() {}

func NewBeq_protocolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Beq_protocolContext {
	var p = new(Beq_protocolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_beq_protocol

	return p
}

func (s *Beq_protocolContext) GetParser() antlr.Parser { return s.parser }

func (s *Beq_protocolContext) Beq_params() IBeq_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeq_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeq_paramsContext)
}

func (s *Beq_protocolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Beq_protocolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Beq_protocolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBeq_protocol(s)
	}
}

func (s *Beq_protocolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBeq_protocol(s)
	}
}

func (s *Beq_protocolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBeq_protocol(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Beq_protocol() (localctx IBeq_protocolContext) {
	localctx = NewBeq_protocolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, tnsnamesParserRULE_beq_protocol)

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
		p.SetState(554)
		p.Beq_params()
	}

	return localctx
}

// IBeq_paramsContext is an interface to support dynamic dispatch.
type IBeq_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBeq_paramsContext differentiates from other interfaces.
	IsBeq_paramsContext()
}

type Beq_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeq_paramsContext() *Beq_paramsContext {
	var p = new(Beq_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_beq_params
	return p
}

func (*Beq_paramsContext) IsBeq_paramsContext() {}

func NewBeq_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Beq_paramsContext {
	var p = new(Beq_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_beq_params

	return p
}

func (s *Beq_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Beq_paramsContext) AllBeq_parameter() []IBeq_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBeq_parameterContext)(nil)).Elem())
	var tst = make([]IBeq_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBeq_parameterContext)
		}
	}

	return tst
}

func (s *Beq_paramsContext) Beq_parameter(i int) IBeq_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeq_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBeq_parameterContext)
}

func (s *Beq_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Beq_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Beq_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBeq_params(s)
	}
}

func (s *Beq_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBeq_params(s)
	}
}

func (s *Beq_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBeq_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Beq_params() (localctx IBeq_paramsContext) {
	localctx = NewBeq_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, tnsnamesParserRULE_beq_params)

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
	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(556)
				p.Beq_parameter()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(559)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext())
	}

	return localctx
}

// IBeq_parameterContext is an interface to support dynamic dispatch.
type IBeq_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBeq_parameterContext differentiates from other interfaces.
	IsBeq_parameterContext()
}

type Beq_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeq_parameterContext() *Beq_parameterContext {
	var p = new(Beq_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_beq_parameter
	return p
}

func (*Beq_parameterContext) IsBeq_parameterContext() {}

func NewBeq_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Beq_parameterContext {
	var p = new(Beq_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_beq_parameter

	return p
}

func (s *Beq_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Beq_parameterContext) Beq_beq() IBeq_beqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeq_beqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeq_beqContext)
}

func (s *Beq_parameterContext) Beq_program() IBeq_programContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeq_programContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeq_programContext)
}

func (s *Beq_parameterContext) Beq_argv0() IBeq_argv0Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeq_argv0Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeq_argv0Context)
}

func (s *Beq_parameterContext) Beq_args() IBeq_argsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeq_argsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeq_argsContext)
}

func (s *Beq_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Beq_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Beq_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBeq_parameter(s)
	}
}

func (s *Beq_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBeq_parameter(s)
	}
}

func (s *Beq_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBeq_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Beq_parameter() (localctx IBeq_parameterContext) {
	localctx = NewBeq_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, tnsnamesParserRULE_beq_parameter)

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

	p.SetState(565)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(561)
			p.Beq_beq()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(562)
			p.Beq_program()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(563)
			p.Beq_argv0()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(564)
			p.Beq_args()
		}

	}

	return localctx
}

// IBeq_beqContext is an interface to support dynamic dispatch.
type IBeq_beqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBeq_beqContext differentiates from other interfaces.
	IsBeq_beqContext()
}

type Beq_beqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeq_beqContext() *Beq_beqContext {
	var p = new(Beq_beqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_beq_beq
	return p
}

func (*Beq_beqContext) IsBeq_beqContext() {}

func NewBeq_beqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Beq_beqContext {
	var p = new(Beq_beqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_beq_beq

	return p
}

func (s *Beq_beqContext) GetParser() antlr.Parser { return s.parser }

func (s *Beq_beqContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Beq_beqContext) PROTOCOL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserPROTOCOL, 0)
}

func (s *Beq_beqContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Beq_beqContext) BEQ() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserBEQ, 0)
}

func (s *Beq_beqContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Beq_beqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Beq_beqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Beq_beqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBeq_beq(s)
	}
}

func (s *Beq_beqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBeq_beq(s)
	}
}

func (s *Beq_beqContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBeq_beq(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Beq_beq() (localctx IBeq_beqContext) {
	localctx = NewBeq_beqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, tnsnamesParserRULE_beq_beq)

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
		p.SetState(567)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(568)
		p.Match(tnsnamesParserPROTOCOL)
	}
	{
		p.SetState(569)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(570)
		p.Match(tnsnamesParserBEQ)
	}
	{
		p.SetState(571)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IBeq_programContext is an interface to support dynamic dispatch.
type IBeq_programContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBeq_programContext differentiates from other interfaces.
	IsBeq_programContext()
}

type Beq_programContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeq_programContext() *Beq_programContext {
	var p = new(Beq_programContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_beq_program
	return p
}

func (*Beq_programContext) IsBeq_programContext() {}

func NewBeq_programContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Beq_programContext {
	var p = new(Beq_programContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_beq_program

	return p
}

func (s *Beq_programContext) GetParser() antlr.Parser { return s.parser }

func (s *Beq_programContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Beq_programContext) PROGRAM() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserPROGRAM, 0)
}

func (s *Beq_programContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Beq_programContext) ID() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, 0)
}

func (s *Beq_programContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Beq_programContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Beq_programContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Beq_programContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBeq_program(s)
	}
}

func (s *Beq_programContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBeq_program(s)
	}
}

func (s *Beq_programContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBeq_program(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Beq_program() (localctx IBeq_programContext) {
	localctx = NewBeq_programContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, tnsnamesParserRULE_beq_program)

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
		p.SetState(573)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(574)
		p.Match(tnsnamesParserPROGRAM)
	}
	{
		p.SetState(575)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(576)
		p.Match(tnsnamesParserID)
	}
	{
		p.SetState(577)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IBeq_argv0Context is an interface to support dynamic dispatch.
type IBeq_argv0Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBeq_argv0Context differentiates from other interfaces.
	IsBeq_argv0Context()
}

type Beq_argv0Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeq_argv0Context() *Beq_argv0Context {
	var p = new(Beq_argv0Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_beq_argv0
	return p
}

func (*Beq_argv0Context) IsBeq_argv0Context() {}

func NewBeq_argv0Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Beq_argv0Context {
	var p = new(Beq_argv0Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_beq_argv0

	return p
}

func (s *Beq_argv0Context) GetParser() antlr.Parser { return s.parser }

func (s *Beq_argv0Context) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Beq_argv0Context) ARGV0() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserARGV0, 0)
}

func (s *Beq_argv0Context) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Beq_argv0Context) ID() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, 0)
}

func (s *Beq_argv0Context) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Beq_argv0Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Beq_argv0Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Beq_argv0Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBeq_argv0(s)
	}
}

func (s *Beq_argv0Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBeq_argv0(s)
	}
}

func (s *Beq_argv0Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBeq_argv0(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Beq_argv0() (localctx IBeq_argv0Context) {
	localctx = NewBeq_argv0Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, tnsnamesParserRULE_beq_argv0)

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
		p.SetState(579)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(580)
		p.Match(tnsnamesParserARGV0)
	}
	{
		p.SetState(581)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(582)
		p.Match(tnsnamesParserID)
	}
	{
		p.SetState(583)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IBeq_argsContext is an interface to support dynamic dispatch.
type IBeq_argsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBeq_argsContext differentiates from other interfaces.
	IsBeq_argsContext()
}

type Beq_argsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBeq_argsContext() *Beq_argsContext {
	var p = new(Beq_argsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_beq_args
	return p
}

func (*Beq_argsContext) IsBeq_argsContext() {}

func NewBeq_argsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Beq_argsContext {
	var p = new(Beq_argsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_beq_args

	return p
}

func (s *Beq_argsContext) GetParser() antlr.Parser { return s.parser }

func (s *Beq_argsContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Beq_argsContext) ARGS() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserARGS, 0)
}

func (s *Beq_argsContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Beq_argsContext) Ba_parameter() IBa_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBa_parameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBa_parameterContext)
}

func (s *Beq_argsContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Beq_argsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Beq_argsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Beq_argsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBeq_args(s)
	}
}

func (s *Beq_argsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBeq_args(s)
	}
}

func (s *Beq_argsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBeq_args(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Beq_args() (localctx IBeq_argsContext) {
	localctx = NewBeq_argsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, tnsnamesParserRULE_beq_args)

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
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(586)
		p.Match(tnsnamesParserARGS)
	}
	{
		p.SetState(587)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(588)
		p.Ba_parameter()
	}
	{
		p.SetState(589)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IBa_parameterContext is an interface to support dynamic dispatch.
type IBa_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBa_parameterContext differentiates from other interfaces.
	IsBa_parameterContext()
}

type Ba_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBa_parameterContext() *Ba_parameterContext {
	var p = new(Ba_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_ba_parameter
	return p
}

func (*Ba_parameterContext) IsBa_parameterContext() {}

func NewBa_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ba_parameterContext {
	var p = new(Ba_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_ba_parameter

	return p
}

func (s *Ba_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Ba_parameterContext) AllS_QUOTE() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserS_QUOTE)
}

func (s *Ba_parameterContext) S_QUOTE(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserS_QUOTE, i)
}

func (s *Ba_parameterContext) Ba_description() IBa_descriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBa_descriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBa_descriptionContext)
}

func (s *Ba_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ba_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ba_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBa_parameter(s)
	}
}

func (s *Ba_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBa_parameter(s)
	}
}

func (s *Ba_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBa_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Ba_parameter() (localctx IBa_parameterContext) {
	localctx = NewBa_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, tnsnamesParserRULE_ba_parameter)

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
		p.SetState(591)
		p.Match(tnsnamesParserS_QUOTE)
	}
	{
		p.SetState(592)
		p.Ba_description()
	}
	{
		p.SetState(593)
		p.Match(tnsnamesParserS_QUOTE)
	}

	return localctx
}

// IBa_descriptionContext is an interface to support dynamic dispatch.
type IBa_descriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBa_descriptionContext differentiates from other interfaces.
	IsBa_descriptionContext()
}

type Ba_descriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBa_descriptionContext() *Ba_descriptionContext {
	var p = new(Ba_descriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_ba_description
	return p
}

func (*Ba_descriptionContext) IsBa_descriptionContext() {}

func NewBa_descriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ba_descriptionContext {
	var p = new(Ba_descriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_ba_description

	return p
}

func (s *Ba_descriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *Ba_descriptionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Ba_descriptionContext) DESCRIPTION() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDESCRIPTION, 0)
}

func (s *Ba_descriptionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Ba_descriptionContext) Bad_params() IBad_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBad_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBad_paramsContext)
}

func (s *Ba_descriptionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Ba_descriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ba_descriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ba_descriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBa_description(s)
	}
}

func (s *Ba_descriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBa_description(s)
	}
}

func (s *Ba_descriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBa_description(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Ba_description() (localctx IBa_descriptionContext) {
	localctx = NewBa_descriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, tnsnamesParserRULE_ba_description)

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
		p.SetState(595)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(596)
		p.Match(tnsnamesParserDESCRIPTION)
	}
	{
		p.SetState(597)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(598)
		p.Bad_params()
	}
	{
		p.SetState(599)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IBad_paramsContext is an interface to support dynamic dispatch.
type IBad_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBad_paramsContext differentiates from other interfaces.
	IsBad_paramsContext()
}

type Bad_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBad_paramsContext() *Bad_paramsContext {
	var p = new(Bad_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_bad_params
	return p
}

func (*Bad_paramsContext) IsBad_paramsContext() {}

func NewBad_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bad_paramsContext {
	var p = new(Bad_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_bad_params

	return p
}

func (s *Bad_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Bad_paramsContext) AllBad_parameter() []IBad_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBad_parameterContext)(nil)).Elem())
	var tst = make([]IBad_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBad_parameterContext)
		}
	}

	return tst
}

func (s *Bad_paramsContext) Bad_parameter(i int) IBad_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBad_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBad_parameterContext)
}

func (s *Bad_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bad_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bad_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBad_params(s)
	}
}

func (s *Bad_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBad_params(s)
	}
}

func (s *Bad_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBad_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Bad_params() (localctx IBad_paramsContext) {
	localctx = NewBad_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, tnsnamesParserRULE_bad_params)
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
	p.SetState(602)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tnsnamesParserL_PAREN {
		{
			p.SetState(601)
			p.Bad_parameter()
		}

		p.SetState(604)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBad_parameterContext is an interface to support dynamic dispatch.
type IBad_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBad_parameterContext differentiates from other interfaces.
	IsBad_parameterContext()
}

type Bad_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBad_parameterContext() *Bad_parameterContext {
	var p = new(Bad_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_bad_parameter
	return p
}

func (*Bad_parameterContext) IsBad_parameterContext() {}

func NewBad_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bad_parameterContext {
	var p = new(Bad_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_bad_parameter

	return p
}

func (s *Bad_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Bad_parameterContext) Bad_local() IBad_localContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBad_localContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBad_localContext)
}

func (s *Bad_parameterContext) Bad_address() IBad_addressContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBad_addressContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBad_addressContext)
}

func (s *Bad_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bad_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bad_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBad_parameter(s)
	}
}

func (s *Bad_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBad_parameter(s)
	}
}

func (s *Bad_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBad_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Bad_parameter() (localctx IBad_parameterContext) {
	localctx = NewBad_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, tnsnamesParserRULE_bad_parameter)

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

	p.SetState(608)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(606)
			p.Bad_local()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(607)
			p.Bad_address()
		}

	}

	return localctx
}

// IBad_localContext is an interface to support dynamic dispatch.
type IBad_localContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBad_localContext differentiates from other interfaces.
	IsBad_localContext()
}

type Bad_localContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBad_localContext() *Bad_localContext {
	var p = new(Bad_localContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_bad_local
	return p
}

func (*Bad_localContext) IsBad_localContext() {}

func NewBad_localContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bad_localContext {
	var p = new(Bad_localContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_bad_local

	return p
}

func (s *Bad_localContext) GetParser() antlr.Parser { return s.parser }

func (s *Bad_localContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Bad_localContext) LOCAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserLOCAL, 0)
}

func (s *Bad_localContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Bad_localContext) YES_NO() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserYES_NO, 0)
}

func (s *Bad_localContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Bad_localContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bad_localContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bad_localContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBad_local(s)
	}
}

func (s *Bad_localContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBad_local(s)
	}
}

func (s *Bad_localContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBad_local(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Bad_local() (localctx IBad_localContext) {
	localctx = NewBad_localContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, tnsnamesParserRULE_bad_local)

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
		p.SetState(610)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(611)
		p.Match(tnsnamesParserLOCAL)
	}
	{
		p.SetState(612)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(613)
		p.Match(tnsnamesParserYES_NO)
	}
	{
		p.SetState(614)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IBad_addressContext is an interface to support dynamic dispatch.
type IBad_addressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBad_addressContext differentiates from other interfaces.
	IsBad_addressContext()
}

type Bad_addressContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBad_addressContext() *Bad_addressContext {
	var p = new(Bad_addressContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_bad_address
	return p
}

func (*Bad_addressContext) IsBad_addressContext() {}

func NewBad_addressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bad_addressContext {
	var p = new(Bad_addressContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_bad_address

	return p
}

func (s *Bad_addressContext) GetParser() antlr.Parser { return s.parser }

func (s *Bad_addressContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Bad_addressContext) ADDRESS() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserADDRESS, 0)
}

func (s *Bad_addressContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Bad_addressContext) Beq_beq() IBeq_beqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBeq_beqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBeq_beqContext)
}

func (s *Bad_addressContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Bad_addressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bad_addressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bad_addressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterBad_address(s)
	}
}

func (s *Bad_addressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitBad_address(s)
	}
}

func (s *Bad_addressContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitBad_address(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Bad_address() (localctx IBad_addressContext) {
	localctx = NewBad_addressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, tnsnamesParserRULE_bad_address)

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
		p.SetState(616)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(617)
		p.Match(tnsnamesParserADDRESS)
	}
	{
		p.SetState(618)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(619)
		p.Beq_beq()
	}
	{
		p.SetState(620)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IConnect_dataContext is an interface to support dynamic dispatch.
type IConnect_dataContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConnect_dataContext differentiates from other interfaces.
	IsConnect_dataContext()
}

type Connect_dataContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnect_dataContext() *Connect_dataContext {
	var p = new(Connect_dataContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_connect_data
	return p
}

func (*Connect_dataContext) IsConnect_dataContext() {}

func NewConnect_dataContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Connect_dataContext {
	var p = new(Connect_dataContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_connect_data

	return p
}

func (s *Connect_dataContext) GetParser() antlr.Parser { return s.parser }

func (s *Connect_dataContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Connect_dataContext) CONNECT_DATA() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserCONNECT_DATA, 0)
}

func (s *Connect_dataContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Connect_dataContext) Cd_params() ICd_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICd_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICd_paramsContext)
}

func (s *Connect_dataContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Connect_dataContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Connect_dataContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Connect_dataContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterConnect_data(s)
	}
}

func (s *Connect_dataContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitConnect_data(s)
	}
}

func (s *Connect_dataContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitConnect_data(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Connect_data() (localctx IConnect_dataContext) {
	localctx = NewConnect_dataContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, tnsnamesParserRULE_connect_data)

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
		p.SetState(622)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(623)
		p.Match(tnsnamesParserCONNECT_DATA)
	}
	{
		p.SetState(624)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(625)
		p.Cd_params()
	}
	{
		p.SetState(626)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ICd_paramsContext is an interface to support dynamic dispatch.
type ICd_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCd_paramsContext differentiates from other interfaces.
	IsCd_paramsContext()
}

type Cd_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCd_paramsContext() *Cd_paramsContext {
	var p = new(Cd_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_cd_params
	return p
}

func (*Cd_paramsContext) IsCd_paramsContext() {}

func NewCd_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cd_paramsContext {
	var p = new(Cd_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_cd_params

	return p
}

func (s *Cd_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Cd_paramsContext) AllCd_parameter() []ICd_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICd_parameterContext)(nil)).Elem())
	var tst = make([]ICd_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICd_parameterContext)
		}
	}

	return tst
}

func (s *Cd_paramsContext) Cd_parameter(i int) ICd_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICd_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICd_parameterContext)
}

func (s *Cd_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cd_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cd_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterCd_params(s)
	}
}

func (s *Cd_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitCd_params(s)
	}
}

func (s *Cd_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitCd_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Cd_params() (localctx ICd_paramsContext) {
	localctx = NewCd_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, tnsnamesParserRULE_cd_params)
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
	p.SetState(629)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tnsnamesParserL_PAREN {
		{
			p.SetState(628)
			p.Cd_parameter()
		}

		p.SetState(631)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICd_parameterContext is an interface to support dynamic dispatch.
type ICd_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCd_parameterContext differentiates from other interfaces.
	IsCd_parameterContext()
}

type Cd_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCd_parameterContext() *Cd_parameterContext {
	var p = new(Cd_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_cd_parameter
	return p
}

func (*Cd_parameterContext) IsCd_parameterContext() {}

func NewCd_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cd_parameterContext {
	var p = new(Cd_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_cd_parameter

	return p
}

func (s *Cd_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Cd_parameterContext) Cd_service_name() ICd_service_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICd_service_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICd_service_nameContext)
}

func (s *Cd_parameterContext) Cd_sid() ICd_sidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICd_sidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICd_sidContext)
}

func (s *Cd_parameterContext) Cd_instance_name() ICd_instance_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICd_instance_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICd_instance_nameContext)
}

func (s *Cd_parameterContext) Cd_failover_mode() ICd_failover_modeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICd_failover_modeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICd_failover_modeContext)
}

func (s *Cd_parameterContext) Cd_global_name() ICd_global_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICd_global_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICd_global_nameContext)
}

func (s *Cd_parameterContext) Cd_hs() ICd_hsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICd_hsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICd_hsContext)
}

func (s *Cd_parameterContext) Cd_rdb_database() ICd_rdb_databaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICd_rdb_databaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICd_rdb_databaseContext)
}

func (s *Cd_parameterContext) Cd_server() ICd_serverContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICd_serverContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICd_serverContext)
}

func (s *Cd_parameterContext) Cd_ur() ICd_urContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICd_urContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICd_urContext)
}

func (s *Cd_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cd_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cd_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterCd_parameter(s)
	}
}

func (s *Cd_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitCd_parameter(s)
	}
}

func (s *Cd_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitCd_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Cd_parameter() (localctx ICd_parameterContext) {
	localctx = NewCd_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, tnsnamesParserRULE_cd_parameter)

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

	p.SetState(642)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(633)
			p.Cd_service_name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(634)
			p.Cd_sid()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(635)
			p.Cd_instance_name()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(636)
			p.Cd_failover_mode()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(637)
			p.Cd_global_name()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(638)
			p.Cd_hs()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(639)
			p.Cd_rdb_database()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(640)
			p.Cd_server()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(641)
			p.Cd_ur()
		}

	}

	return localctx
}

// ICd_service_nameContext is an interface to support dynamic dispatch.
type ICd_service_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCd_service_nameContext differentiates from other interfaces.
	IsCd_service_nameContext()
}

type Cd_service_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCd_service_nameContext() *Cd_service_nameContext {
	var p = new(Cd_service_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_cd_service_name
	return p
}

func (*Cd_service_nameContext) IsCd_service_nameContext() {}

func NewCd_service_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cd_service_nameContext {
	var p = new(Cd_service_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_cd_service_name

	return p
}

func (s *Cd_service_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Cd_service_nameContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Cd_service_nameContext) SERVICE_NAME() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSERVICE_NAME, 0)
}

func (s *Cd_service_nameContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Cd_service_nameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserID)
}

func (s *Cd_service_nameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, i)
}

func (s *Cd_service_nameContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Cd_service_nameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserDOT)
}

func (s *Cd_service_nameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDOT, i)
}

func (s *Cd_service_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cd_service_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cd_service_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterCd_service_name(s)
	}
}

func (s *Cd_service_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitCd_service_name(s)
	}
}

func (s *Cd_service_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitCd_service_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Cd_service_name() (localctx ICd_service_nameContext) {
	localctx = NewCd_service_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, tnsnamesParserRULE_cd_service_name)
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
		p.SetState(644)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(645)
		p.Match(tnsnamesParserSERVICE_NAME)
	}
	{
		p.SetState(646)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(647)
		p.Match(tnsnamesParserID)
	}
	p.SetState(652)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tnsnamesParserDOT {
		{
			p.SetState(648)
			p.Match(tnsnamesParserDOT)
		}
		{
			p.SetState(649)
			p.Match(tnsnamesParserID)
		}

		p.SetState(654)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(655)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ICd_sidContext is an interface to support dynamic dispatch.
type ICd_sidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCd_sidContext differentiates from other interfaces.
	IsCd_sidContext()
}

type Cd_sidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCd_sidContext() *Cd_sidContext {
	var p = new(Cd_sidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_cd_sid
	return p
}

func (*Cd_sidContext) IsCd_sidContext() {}

func NewCd_sidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cd_sidContext {
	var p = new(Cd_sidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_cd_sid

	return p
}

func (s *Cd_sidContext) GetParser() antlr.Parser { return s.parser }

func (s *Cd_sidContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Cd_sidContext) SID() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSID, 0)
}

func (s *Cd_sidContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Cd_sidContext) ID() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, 0)
}

func (s *Cd_sidContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Cd_sidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cd_sidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cd_sidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterCd_sid(s)
	}
}

func (s *Cd_sidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitCd_sid(s)
	}
}

func (s *Cd_sidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitCd_sid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Cd_sid() (localctx ICd_sidContext) {
	localctx = NewCd_sidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, tnsnamesParserRULE_cd_sid)

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
		p.SetState(657)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(658)
		p.Match(tnsnamesParserSID)
	}
	{
		p.SetState(659)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(660)
		p.Match(tnsnamesParserID)
	}
	{
		p.SetState(661)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ICd_instance_nameContext is an interface to support dynamic dispatch.
type ICd_instance_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCd_instance_nameContext differentiates from other interfaces.
	IsCd_instance_nameContext()
}

type Cd_instance_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCd_instance_nameContext() *Cd_instance_nameContext {
	var p = new(Cd_instance_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_cd_instance_name
	return p
}

func (*Cd_instance_nameContext) IsCd_instance_nameContext() {}

func NewCd_instance_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cd_instance_nameContext {
	var p = new(Cd_instance_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_cd_instance_name

	return p
}

func (s *Cd_instance_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Cd_instance_nameContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Cd_instance_nameContext) INSTANCE_NAME() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserINSTANCE_NAME, 0)
}

func (s *Cd_instance_nameContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Cd_instance_nameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserID)
}

func (s *Cd_instance_nameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, i)
}

func (s *Cd_instance_nameContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Cd_instance_nameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserDOT)
}

func (s *Cd_instance_nameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDOT, i)
}

func (s *Cd_instance_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cd_instance_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cd_instance_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterCd_instance_name(s)
	}
}

func (s *Cd_instance_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitCd_instance_name(s)
	}
}

func (s *Cd_instance_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitCd_instance_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Cd_instance_name() (localctx ICd_instance_nameContext) {
	localctx = NewCd_instance_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, tnsnamesParserRULE_cd_instance_name)
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
		p.SetState(663)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(664)
		p.Match(tnsnamesParserINSTANCE_NAME)
	}
	{
		p.SetState(665)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(666)
		p.Match(tnsnamesParserID)
	}
	p.SetState(671)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tnsnamesParserDOT {
		{
			p.SetState(667)
			p.Match(tnsnamesParserDOT)
		}
		{
			p.SetState(668)
			p.Match(tnsnamesParserID)
		}

		p.SetState(673)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(674)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ICd_failover_modeContext is an interface to support dynamic dispatch.
type ICd_failover_modeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCd_failover_modeContext differentiates from other interfaces.
	IsCd_failover_modeContext()
}

type Cd_failover_modeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCd_failover_modeContext() *Cd_failover_modeContext {
	var p = new(Cd_failover_modeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_cd_failover_mode
	return p
}

func (*Cd_failover_modeContext) IsCd_failover_modeContext() {}

func NewCd_failover_modeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cd_failover_modeContext {
	var p = new(Cd_failover_modeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_cd_failover_mode

	return p
}

func (s *Cd_failover_modeContext) GetParser() antlr.Parser { return s.parser }

func (s *Cd_failover_modeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Cd_failover_modeContext) FAILOVER_MODE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserFAILOVER_MODE, 0)
}

func (s *Cd_failover_modeContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Cd_failover_modeContext) Fo_params() IFo_paramsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFo_paramsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFo_paramsContext)
}

func (s *Cd_failover_modeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Cd_failover_modeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cd_failover_modeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cd_failover_modeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterCd_failover_mode(s)
	}
}

func (s *Cd_failover_modeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitCd_failover_mode(s)
	}
}

func (s *Cd_failover_modeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitCd_failover_mode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Cd_failover_mode() (localctx ICd_failover_modeContext) {
	localctx = NewCd_failover_modeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 152, tnsnamesParserRULE_cd_failover_mode)

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
		p.SetState(676)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(677)
		p.Match(tnsnamesParserFAILOVER_MODE)
	}
	{
		p.SetState(678)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(679)
		p.Fo_params()
	}
	{
		p.SetState(680)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ICd_global_nameContext is an interface to support dynamic dispatch.
type ICd_global_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCd_global_nameContext differentiates from other interfaces.
	IsCd_global_nameContext()
}

type Cd_global_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCd_global_nameContext() *Cd_global_nameContext {
	var p = new(Cd_global_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_cd_global_name
	return p
}

func (*Cd_global_nameContext) IsCd_global_nameContext() {}

func NewCd_global_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cd_global_nameContext {
	var p = new(Cd_global_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_cd_global_name

	return p
}

func (s *Cd_global_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Cd_global_nameContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Cd_global_nameContext) GLOBAL_NAME() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserGLOBAL_NAME, 0)
}

func (s *Cd_global_nameContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Cd_global_nameContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserID)
}

func (s *Cd_global_nameContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, i)
}

func (s *Cd_global_nameContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Cd_global_nameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserDOT)
}

func (s *Cd_global_nameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDOT, i)
}

func (s *Cd_global_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cd_global_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cd_global_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterCd_global_name(s)
	}
}

func (s *Cd_global_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitCd_global_name(s)
	}
}

func (s *Cd_global_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitCd_global_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Cd_global_name() (localctx ICd_global_nameContext) {
	localctx = NewCd_global_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 154, tnsnamesParserRULE_cd_global_name)
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
		p.SetState(682)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(683)
		p.Match(tnsnamesParserGLOBAL_NAME)
	}
	{
		p.SetState(684)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(685)
		p.Match(tnsnamesParserID)
	}
	p.SetState(690)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tnsnamesParserDOT {
		{
			p.SetState(686)
			p.Match(tnsnamesParserDOT)
		}
		{
			p.SetState(687)
			p.Match(tnsnamesParserID)
		}

		p.SetState(692)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(693)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ICd_hsContext is an interface to support dynamic dispatch.
type ICd_hsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCd_hsContext differentiates from other interfaces.
	IsCd_hsContext()
}

type Cd_hsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCd_hsContext() *Cd_hsContext {
	var p = new(Cd_hsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_cd_hs
	return p
}

func (*Cd_hsContext) IsCd_hsContext() {}

func NewCd_hsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cd_hsContext {
	var p = new(Cd_hsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_cd_hs

	return p
}

func (s *Cd_hsContext) GetParser() antlr.Parser { return s.parser }

func (s *Cd_hsContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Cd_hsContext) HS() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserHS, 0)
}

func (s *Cd_hsContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Cd_hsContext) OK() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserOK, 0)
}

func (s *Cd_hsContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Cd_hsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cd_hsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cd_hsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterCd_hs(s)
	}
}

func (s *Cd_hsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitCd_hs(s)
	}
}

func (s *Cd_hsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitCd_hs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Cd_hs() (localctx ICd_hsContext) {
	localctx = NewCd_hsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 156, tnsnamesParserRULE_cd_hs)

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
		p.SetState(695)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(696)
		p.Match(tnsnamesParserHS)
	}
	{
		p.SetState(697)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(698)
		p.Match(tnsnamesParserOK)
	}
	{
		p.SetState(699)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ICd_rdb_databaseContext is an interface to support dynamic dispatch.
type ICd_rdb_databaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCd_rdb_databaseContext differentiates from other interfaces.
	IsCd_rdb_databaseContext()
}

type Cd_rdb_databaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCd_rdb_databaseContext() *Cd_rdb_databaseContext {
	var p = new(Cd_rdb_databaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_cd_rdb_database
	return p
}

func (*Cd_rdb_databaseContext) IsCd_rdb_databaseContext() {}

func NewCd_rdb_databaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cd_rdb_databaseContext {
	var p = new(Cd_rdb_databaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_cd_rdb_database

	return p
}

func (s *Cd_rdb_databaseContext) GetParser() antlr.Parser { return s.parser }

func (s *Cd_rdb_databaseContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Cd_rdb_databaseContext) RDB_DATABASE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserRDB_DATABASE, 0)
}

func (s *Cd_rdb_databaseContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Cd_rdb_databaseContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserID)
}

func (s *Cd_rdb_databaseContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, i)
}

func (s *Cd_rdb_databaseContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Cd_rdb_databaseContext) L_SQUARE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_SQUARE, 0)
}

func (s *Cd_rdb_databaseContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserDOT)
}

func (s *Cd_rdb_databaseContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDOT, i)
}

func (s *Cd_rdb_databaseContext) R_SQUARE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_SQUARE, 0)
}

func (s *Cd_rdb_databaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cd_rdb_databaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cd_rdb_databaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterCd_rdb_database(s)
	}
}

func (s *Cd_rdb_databaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitCd_rdb_database(s)
	}
}

func (s *Cd_rdb_databaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitCd_rdb_database(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Cd_rdb_database() (localctx ICd_rdb_databaseContext) {
	localctx = NewCd_rdb_databaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 158, tnsnamesParserRULE_cd_rdb_database)
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
		p.SetState(701)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(702)
		p.Match(tnsnamesParserRDB_DATABASE)
	}
	{
		p.SetState(703)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(708)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == tnsnamesParserL_SQUARE {
		{
			p.SetState(704)
			p.Match(tnsnamesParserL_SQUARE)
		}
		{
			p.SetState(705)
			p.Match(tnsnamesParserDOT)
		}
		{
			p.SetState(706)
			p.Match(tnsnamesParserID)
		}
		{
			p.SetState(707)
			p.Match(tnsnamesParserR_SQUARE)
		}

	}
	{
		p.SetState(710)
		p.Match(tnsnamesParserID)
	}
	p.SetState(715)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tnsnamesParserDOT {
		{
			p.SetState(711)
			p.Match(tnsnamesParserDOT)
		}
		{
			p.SetState(712)
			p.Match(tnsnamesParserID)
		}

		p.SetState(717)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(718)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ICd_serverContext is an interface to support dynamic dispatch.
type ICd_serverContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCd_serverContext differentiates from other interfaces.
	IsCd_serverContext()
}

type Cd_serverContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCd_serverContext() *Cd_serverContext {
	var p = new(Cd_serverContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_cd_server
	return p
}

func (*Cd_serverContext) IsCd_serverContext() {}

func NewCd_serverContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cd_serverContext {
	var p = new(Cd_serverContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_cd_server

	return p
}

func (s *Cd_serverContext) GetParser() antlr.Parser { return s.parser }

func (s *Cd_serverContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Cd_serverContext) SERVER() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSERVER, 0)
}

func (s *Cd_serverContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Cd_serverContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Cd_serverContext) DEDICATED() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDEDICATED, 0)
}

func (s *Cd_serverContext) SHARED() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSHARED, 0)
}

func (s *Cd_serverContext) POOLED() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserPOOLED, 0)
}

func (s *Cd_serverContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cd_serverContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cd_serverContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterCd_server(s)
	}
}

func (s *Cd_serverContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitCd_server(s)
	}
}

func (s *Cd_serverContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitCd_server(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Cd_server() (localctx ICd_serverContext) {
	localctx = NewCd_serverContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 160, tnsnamesParserRULE_cd_server)
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
		p.SetState(720)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(721)
		p.Match(tnsnamesParserSERVER)
	}
	{
		p.SetState(722)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(723)
	_la = p.GetTokenStream().LA(1)

	if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<tnsnamesParserDEDICATED)|(1<<tnsnamesParserSHARED)|(1<<tnsnamesParserPOOLED))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(724)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// ICd_urContext is an interface to support dynamic dispatch.
type ICd_urContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCd_urContext differentiates from other interfaces.
	IsCd_urContext()
}

type Cd_urContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCd_urContext() *Cd_urContext {
	var p = new(Cd_urContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_cd_ur
	return p
}

func (*Cd_urContext) IsCd_urContext() {}

func NewCd_urContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cd_urContext {
	var p = new(Cd_urContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_cd_ur

	return p
}

func (s *Cd_urContext) GetParser() antlr.Parser { return s.parser }

func (s *Cd_urContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Cd_urContext) UR() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserUR, 0)
}

func (s *Cd_urContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Cd_urContext) UR_A() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserUR_A, 0)
}

func (s *Cd_urContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Cd_urContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cd_urContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cd_urContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterCd_ur(s)
	}
}

func (s *Cd_urContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitCd_ur(s)
	}
}

func (s *Cd_urContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitCd_ur(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Cd_ur() (localctx ICd_urContext) {
	localctx = NewCd_urContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 162, tnsnamesParserRULE_cd_ur)

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
		p.SetState(726)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(727)
		p.Match(tnsnamesParserUR)
	}
	{
		p.SetState(728)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(729)
		p.Match(tnsnamesParserUR_A)
	}
	{
		p.SetState(730)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IFo_paramsContext is an interface to support dynamic dispatch.
type IFo_paramsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFo_paramsContext differentiates from other interfaces.
	IsFo_paramsContext()
}

type Fo_paramsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFo_paramsContext() *Fo_paramsContext {
	var p = new(Fo_paramsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_fo_params
	return p
}

func (*Fo_paramsContext) IsFo_paramsContext() {}

func NewFo_paramsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fo_paramsContext {
	var p = new(Fo_paramsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_fo_params

	return p
}

func (s *Fo_paramsContext) GetParser() antlr.Parser { return s.parser }

func (s *Fo_paramsContext) AllFo_parameter() []IFo_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFo_parameterContext)(nil)).Elem())
	var tst = make([]IFo_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFo_parameterContext)
		}
	}

	return tst
}

func (s *Fo_paramsContext) Fo_parameter(i int) IFo_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFo_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFo_parameterContext)
}

func (s *Fo_paramsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fo_paramsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fo_paramsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterFo_params(s)
	}
}

func (s *Fo_paramsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitFo_params(s)
	}
}

func (s *Fo_paramsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitFo_params(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Fo_params() (localctx IFo_paramsContext) {
	localctx = NewFo_paramsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 164, tnsnamesParserRULE_fo_params)
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
	p.SetState(733)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == tnsnamesParserL_PAREN {
		{
			p.SetState(732)
			p.Fo_parameter()
		}

		p.SetState(735)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFo_parameterContext is an interface to support dynamic dispatch.
type IFo_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFo_parameterContext differentiates from other interfaces.
	IsFo_parameterContext()
}

type Fo_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFo_parameterContext() *Fo_parameterContext {
	var p = new(Fo_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_fo_parameter
	return p
}

func (*Fo_parameterContext) IsFo_parameterContext() {}

func NewFo_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fo_parameterContext {
	var p = new(Fo_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_fo_parameter

	return p
}

func (s *Fo_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Fo_parameterContext) Fo_type() IFo_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFo_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFo_typeContext)
}

func (s *Fo_parameterContext) Fo_backup() IFo_backupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFo_backupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFo_backupContext)
}

func (s *Fo_parameterContext) Fo_method() IFo_methodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFo_methodContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFo_methodContext)
}

func (s *Fo_parameterContext) Fo_retries() IFo_retriesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFo_retriesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFo_retriesContext)
}

func (s *Fo_parameterContext) Fo_delay() IFo_delayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFo_delayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFo_delayContext)
}

func (s *Fo_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fo_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fo_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterFo_parameter(s)
	}
}

func (s *Fo_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitFo_parameter(s)
	}
}

func (s *Fo_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitFo_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Fo_parameter() (localctx IFo_parameterContext) {
	localctx = NewFo_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 166, tnsnamesParserRULE_fo_parameter)

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

	p.SetState(742)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(737)
			p.Fo_type()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(738)
			p.Fo_backup()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(739)
			p.Fo_method()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(740)
			p.Fo_retries()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(741)
			p.Fo_delay()
		}

	}

	return localctx
}

// IFo_typeContext is an interface to support dynamic dispatch.
type IFo_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFo_typeContext differentiates from other interfaces.
	IsFo_typeContext()
}

type Fo_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFo_typeContext() *Fo_typeContext {
	var p = new(Fo_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_fo_type
	return p
}

func (*Fo_typeContext) IsFo_typeContext() {}

func NewFo_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fo_typeContext {
	var p = new(Fo_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_fo_type

	return p
}

func (s *Fo_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Fo_typeContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Fo_typeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserTYPE, 0)
}

func (s *Fo_typeContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Fo_typeContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Fo_typeContext) SESSION() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSESSION, 0)
}

func (s *Fo_typeContext) SELECT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserSELECT, 0)
}

func (s *Fo_typeContext) NONE() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserNONE, 0)
}

func (s *Fo_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fo_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fo_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterFo_type(s)
	}
}

func (s *Fo_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitFo_type(s)
	}
}

func (s *Fo_typeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitFo_type(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Fo_type() (localctx IFo_typeContext) {
	localctx = NewFo_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 168, tnsnamesParserRULE_fo_type)
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
		p.SetState(744)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(745)
		p.Match(tnsnamesParserTYPE)
	}
	{
		p.SetState(746)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(747)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-68)&-(0x1f+1)) == 0 && ((1<<uint((_la-68)))&((1<<(tnsnamesParserSESSION-68))|(1<<(tnsnamesParserSELECT-68))|(1<<(tnsnamesParserNONE-68)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(748)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IFo_backupContext is an interface to support dynamic dispatch.
type IFo_backupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFo_backupContext differentiates from other interfaces.
	IsFo_backupContext()
}

type Fo_backupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFo_backupContext() *Fo_backupContext {
	var p = new(Fo_backupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_fo_backup
	return p
}

func (*Fo_backupContext) IsFo_backupContext() {}

func NewFo_backupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fo_backupContext {
	var p = new(Fo_backupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_fo_backup

	return p
}

func (s *Fo_backupContext) GetParser() antlr.Parser { return s.parser }

func (s *Fo_backupContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Fo_backupContext) BACKUP() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserBACKUP, 0)
}

func (s *Fo_backupContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Fo_backupContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserID)
}

func (s *Fo_backupContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserID, i)
}

func (s *Fo_backupContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Fo_backupContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(tnsnamesParserDOT)
}

func (s *Fo_backupContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDOT, i)
}

func (s *Fo_backupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fo_backupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fo_backupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterFo_backup(s)
	}
}

func (s *Fo_backupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitFo_backup(s)
	}
}

func (s *Fo_backupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitFo_backup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Fo_backup() (localctx IFo_backupContext) {
	localctx = NewFo_backupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 170, tnsnamesParserRULE_fo_backup)
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
		p.SetState(750)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(751)
		p.Match(tnsnamesParserBACKUP)
	}
	{
		p.SetState(752)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(753)
		p.Match(tnsnamesParserID)
	}
	p.SetState(758)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == tnsnamesParserDOT {
		{
			p.SetState(754)
			p.Match(tnsnamesParserDOT)
		}
		{
			p.SetState(755)
			p.Match(tnsnamesParserID)
		}

		p.SetState(760)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(761)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IFo_methodContext is an interface to support dynamic dispatch.
type IFo_methodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFo_methodContext differentiates from other interfaces.
	IsFo_methodContext()
}

type Fo_methodContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFo_methodContext() *Fo_methodContext {
	var p = new(Fo_methodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_fo_method
	return p
}

func (*Fo_methodContext) IsFo_methodContext() {}

func NewFo_methodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fo_methodContext {
	var p = new(Fo_methodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_fo_method

	return p
}

func (s *Fo_methodContext) GetParser() antlr.Parser { return s.parser }

func (s *Fo_methodContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Fo_methodContext) METHOD() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserMETHOD, 0)
}

func (s *Fo_methodContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Fo_methodContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Fo_methodContext) BASIC() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserBASIC, 0)
}

func (s *Fo_methodContext) PRECONNECT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserPRECONNECT, 0)
}

func (s *Fo_methodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fo_methodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fo_methodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterFo_method(s)
	}
}

func (s *Fo_methodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitFo_method(s)
	}
}

func (s *Fo_methodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitFo_method(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Fo_method() (localctx IFo_methodContext) {
	localctx = NewFo_methodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 172, tnsnamesParserRULE_fo_method)
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
		p.SetState(763)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(764)
		p.Match(tnsnamesParserMETHOD)
	}
	{
		p.SetState(765)
		p.Match(tnsnamesParserEQUAL)
	}
	p.SetState(766)
	_la = p.GetTokenStream().LA(1)

	if !(_la == tnsnamesParserBASIC || _la == tnsnamesParserPRECONNECT) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	{
		p.SetState(767)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IFo_retriesContext is an interface to support dynamic dispatch.
type IFo_retriesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFo_retriesContext differentiates from other interfaces.
	IsFo_retriesContext()
}

type Fo_retriesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFo_retriesContext() *Fo_retriesContext {
	var p = new(Fo_retriesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_fo_retries
	return p
}

func (*Fo_retriesContext) IsFo_retriesContext() {}

func NewFo_retriesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fo_retriesContext {
	var p = new(Fo_retriesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_fo_retries

	return p
}

func (s *Fo_retriesContext) GetParser() antlr.Parser { return s.parser }

func (s *Fo_retriesContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Fo_retriesContext) RETRIES() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserRETRIES, 0)
}

func (s *Fo_retriesContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Fo_retriesContext) INT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserINT, 0)
}

func (s *Fo_retriesContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Fo_retriesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fo_retriesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fo_retriesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterFo_retries(s)
	}
}

func (s *Fo_retriesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitFo_retries(s)
	}
}

func (s *Fo_retriesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitFo_retries(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Fo_retries() (localctx IFo_retriesContext) {
	localctx = NewFo_retriesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 174, tnsnamesParserRULE_fo_retries)

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
		p.SetState(769)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(770)
		p.Match(tnsnamesParserRETRIES)
	}
	{
		p.SetState(771)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(772)
		p.Match(tnsnamesParserINT)
	}
	{
		p.SetState(773)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}

// IFo_delayContext is an interface to support dynamic dispatch.
type IFo_delayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFo_delayContext differentiates from other interfaces.
	IsFo_delayContext()
}

type Fo_delayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFo_delayContext() *Fo_delayContext {
	var p = new(Fo_delayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = tnsnamesParserRULE_fo_delay
	return p
}

func (*Fo_delayContext) IsFo_delayContext() {}

func NewFo_delayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fo_delayContext {
	var p = new(Fo_delayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = tnsnamesParserRULE_fo_delay

	return p
}

func (s *Fo_delayContext) GetParser() antlr.Parser { return s.parser }

func (s *Fo_delayContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserL_PAREN, 0)
}

func (s *Fo_delayContext) DELAY() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserDELAY, 0)
}

func (s *Fo_delayContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserEQUAL, 0)
}

func (s *Fo_delayContext) INT() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserINT, 0)
}

func (s *Fo_delayContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(tnsnamesParserR_PAREN, 0)
}

func (s *Fo_delayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fo_delayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fo_delayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.EnterFo_delay(s)
	}
}

func (s *Fo_delayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tnsnamesParserListener); ok {
		listenerT.ExitFo_delay(s)
	}
}

func (s *Fo_delayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tnsnamesParserVisitor:
		return t.VisitFo_delay(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tnsnamesParser) Fo_delay() (localctx IFo_delayContext) {
	localctx = NewFo_delayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 176, tnsnamesParserRULE_fo_delay)

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
		p.SetState(775)
		p.Match(tnsnamesParserL_PAREN)
	}
	{
		p.SetState(776)
		p.Match(tnsnamesParserDELAY)
	}
	{
		p.SetState(777)
		p.Match(tnsnamesParserEQUAL)
	}
	{
		p.SetState(778)
		p.Match(tnsnamesParserINT)
	}
	{
		p.SetState(779)
		p.Match(tnsnamesParserR_PAREN)
	}

	return localctx
}
