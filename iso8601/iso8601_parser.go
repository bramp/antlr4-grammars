// Code generated from iso8601.g4 by ANTLR 4.9.3. DO NOT EDIT.

package iso8601 // iso8601
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 19, 703,
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
	4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9, 91, 4,
	92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 3, 2, 6, 2, 192, 10,
	2, 13, 2, 14, 2, 193, 3, 3, 6, 3, 197, 10, 3, 13, 3, 14, 3, 198, 3, 3,
	5, 3, 202, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 6, 7, 219, 10, 7, 13, 7, 14, 7, 220,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 6, 8, 228, 10, 8, 13, 8, 14, 8, 229, 3, 9,
	3, 9, 3, 9, 5, 9, 235, 10, 9, 3, 10, 3, 10, 5, 10, 239, 10, 10, 3, 11,
	3, 11, 5, 11, 243, 10, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 5, 23, 269, 10, 23, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 5, 29, 291,
	10, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32,
	5, 32, 302, 10, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 5, 35, 318, 10, 35, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 5,
	38, 331, 10, 38, 3, 39, 3, 39, 3, 39, 5, 39, 336, 10, 39, 3, 40, 3, 40,
	3, 40, 5, 40, 341, 10, 40, 3, 41, 3, 41, 5, 41, 345, 10, 41, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 5, 42, 352, 10, 42, 3, 43, 3, 43, 3, 43, 3, 43,
	3, 43, 5, 43, 359, 10, 43, 3, 44, 3, 44, 5, 44, 363, 10, 44, 3, 45, 3,
	45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47,
	5, 47, 377, 10, 47, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3,
	50, 3, 50, 3, 51, 3, 51, 5, 51, 390, 10, 51, 3, 52, 3, 52, 3, 52, 5, 52,
	395, 10, 52, 3, 53, 3, 53, 3, 53, 5, 53, 400, 10, 53, 3, 54, 3, 54, 3,
	55, 3, 55, 5, 55, 406, 10, 55, 3, 56, 3, 56, 3, 56, 5, 56, 411, 10, 56,
	3, 56, 5, 56, 414, 10, 56, 3, 57, 3, 57, 3, 57, 3, 57, 5, 57, 420, 10,
	57, 3, 57, 5, 57, 423, 10, 57, 3, 58, 3, 58, 5, 58, 427, 10, 58, 3, 59,
	3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 61, 5, 61, 436, 10, 61, 3, 61, 3,
	61, 5, 61, 440, 10, 61, 3, 62, 3, 62, 5, 62, 444, 10, 62, 3, 63, 3, 63,
	3, 63, 3, 63, 5, 63, 450, 10, 63, 3, 64, 3, 64, 3, 64, 3, 64, 5, 64, 456,
	10, 64, 3, 65, 3, 65, 5, 65, 460, 10, 65, 3, 66, 3, 66, 3, 66, 3, 66, 5,
	66, 466, 10, 66, 3, 67, 3, 67, 3, 67, 3, 67, 5, 67, 472, 10, 67, 3, 68,
	3, 68, 5, 68, 476, 10, 68, 3, 69, 3, 69, 5, 69, 480, 10, 69, 3, 70, 3,
	70, 5, 70, 484, 10, 70, 3, 71, 3, 71, 3, 71, 3, 71, 5, 71, 490, 10, 71,
	3, 71, 3, 71, 3, 71, 5, 71, 495, 10, 71, 3, 71, 3, 71, 3, 71, 5, 71, 500,
	10, 71, 3, 71, 3, 71, 3, 71, 3, 71, 5, 71, 506, 10, 71, 3, 71, 3, 71, 3,
	71, 5, 71, 511, 10, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71,
	5, 71, 520, 10, 71, 3, 71, 3, 71, 3, 71, 5, 71, 525, 10, 71, 3, 71, 3,
	71, 3, 71, 5, 71, 530, 10, 71, 3, 71, 3, 71, 3, 71, 3, 71, 5, 71, 536,
	10, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 5, 71, 545, 10,
	71, 3, 71, 3, 71, 3, 71, 5, 71, 550, 10, 71, 3, 71, 3, 71, 3, 71, 5, 71,
	555, 10, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 5,
	71, 565, 10, 71, 3, 71, 3, 71, 3, 71, 5, 71, 570, 10, 71, 3, 71, 3, 71,
	3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 5, 71, 579, 10, 71, 3, 71, 3, 71, 3,
	71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 5, 71, 592,
	10, 71, 3, 72, 3, 72, 3, 72, 3, 73, 3, 73, 3, 73, 3, 73, 3, 74, 3, 74,
	3, 74, 3, 74, 3, 74, 5, 74, 606, 10, 74, 3, 75, 3, 75, 3, 75, 3, 75, 3,
	75, 3, 75, 5, 75, 614, 10, 75, 3, 76, 3, 76, 5, 76, 618, 10, 76, 3, 77,
	3, 77, 3, 77, 3, 78, 3, 78, 3, 78, 3, 79, 3, 79, 5, 79, 628, 10, 79, 3,
	80, 3, 80, 3, 80, 3, 80, 3, 81, 3, 81, 3, 81, 3, 81, 3, 82, 3, 82, 5, 82,
	640, 10, 82, 3, 83, 3, 83, 3, 83, 3, 83, 3, 84, 3, 84, 3, 84, 3, 84, 3,
	85, 3, 85, 5, 85, 652, 10, 85, 3, 86, 3, 86, 3, 86, 3, 86, 3, 87, 3, 87,
	3, 87, 3, 87, 3, 88, 3, 88, 5, 88, 664, 10, 88, 3, 89, 3, 89, 3, 89, 3,
	89, 5, 89, 670, 10, 89, 3, 90, 3, 90, 3, 90, 3, 90, 5, 90, 676, 10, 90,
	3, 91, 3, 91, 5, 91, 680, 10, 91, 3, 92, 3, 92, 5, 92, 684, 10, 92, 3,
	93, 3, 93, 3, 93, 3, 93, 3, 94, 3, 94, 3, 94, 3, 94, 3, 95, 3, 95, 3, 95,
	3, 95, 3, 95, 5, 95, 699, 10, 95, 3, 95, 3, 95, 3, 95, 2, 2, 96, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78,
	80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112,
	114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136, 138, 140, 142,
	144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166, 168, 170, 172,
	174, 176, 178, 180, 182, 184, 186, 188, 2, 3, 3, 2, 3, 4, 2, 705, 2, 191,
	3, 2, 2, 2, 4, 196, 3, 2, 2, 2, 6, 203, 3, 2, 2, 2, 8, 206, 3, 2, 2, 2,
	10, 210, 3, 2, 2, 2, 12, 215, 3, 2, 2, 2, 14, 222, 3, 2, 2, 2, 16, 231,
	3, 2, 2, 2, 18, 238, 3, 2, 2, 2, 20, 242, 3, 2, 2, 2, 22, 244, 3, 2, 2,
	2, 24, 246, 3, 2, 2, 2, 26, 248, 3, 2, 2, 2, 28, 250, 3, 2, 2, 2, 30, 252,
	3, 2, 2, 2, 32, 254, 3, 2, 2, 2, 34, 256, 3, 2, 2, 2, 36, 258, 3, 2, 2,
	2, 38, 260, 3, 2, 2, 2, 40, 262, 3, 2, 2, 2, 42, 264, 3, 2, 2, 2, 44, 268,
	3, 2, 2, 2, 46, 270, 3, 2, 2, 2, 48, 274, 3, 2, 2, 2, 50, 280, 3, 2, 2,
	2, 52, 284, 3, 2, 2, 2, 54, 286, 3, 2, 2, 2, 56, 290, 3, 2, 2, 2, 58, 292,
	3, 2, 2, 2, 60, 295, 3, 2, 2, 2, 62, 301, 3, 2, 2, 2, 64, 303, 3, 2, 2,
	2, 66, 308, 3, 2, 2, 2, 68, 317, 3, 2, 2, 2, 70, 319, 3, 2, 2, 2, 72, 323,
	3, 2, 2, 2, 74, 330, 3, 2, 2, 2, 76, 335, 3, 2, 2, 2, 78, 340, 3, 2, 2,
	2, 80, 344, 3, 2, 2, 2, 82, 351, 3, 2, 2, 2, 84, 358, 3, 2, 2, 2, 86, 362,
	3, 2, 2, 2, 88, 364, 3, 2, 2, 2, 90, 368, 3, 2, 2, 2, 92, 376, 3, 2, 2,
	2, 94, 378, 3, 2, 2, 2, 96, 381, 3, 2, 2, 2, 98, 385, 3, 2, 2, 2, 100,
	389, 3, 2, 2, 2, 102, 394, 3, 2, 2, 2, 104, 399, 3, 2, 2, 2, 106, 401,
	3, 2, 2, 2, 108, 405, 3, 2, 2, 2, 110, 413, 3, 2, 2, 2, 112, 422, 3, 2,
	2, 2, 114, 426, 3, 2, 2, 2, 116, 428, 3, 2, 2, 2, 118, 431, 3, 2, 2, 2,
	120, 435, 3, 2, 2, 2, 122, 443, 3, 2, 2, 2, 124, 445, 3, 2, 2, 2, 126,
	451, 3, 2, 2, 2, 128, 459, 3, 2, 2, 2, 130, 461, 3, 2, 2, 2, 132, 467,
	3, 2, 2, 2, 134, 475, 3, 2, 2, 2, 136, 479, 3, 2, 2, 2, 138, 483, 3, 2,
	2, 2, 140, 591, 3, 2, 2, 2, 142, 593, 3, 2, 2, 2, 144, 596, 3, 2, 2, 2,
	146, 605, 3, 2, 2, 2, 148, 613, 3, 2, 2, 2, 150, 617, 3, 2, 2, 2, 152,
	619, 3, 2, 2, 2, 154, 622, 3, 2, 2, 2, 156, 627, 3, 2, 2, 2, 158, 629,
	3, 2, 2, 2, 160, 633, 3, 2, 2, 2, 162, 639, 3, 2, 2, 2, 164, 641, 3, 2,
	2, 2, 166, 645, 3, 2, 2, 2, 168, 651, 3, 2, 2, 2, 170, 653, 3, 2, 2, 2,
	172, 657, 3, 2, 2, 2, 174, 663, 3, 2, 2, 2, 176, 669, 3, 2, 2, 2, 178,
	675, 3, 2, 2, 2, 180, 677, 3, 2, 2, 2, 182, 683, 3, 2, 2, 2, 184, 685,
	3, 2, 2, 2, 186, 689, 3, 2, 2, 2, 188, 698, 3, 2, 2, 2, 190, 192, 7, 18,
	2, 2, 191, 190, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2,
	193, 194, 3, 2, 2, 2, 194, 3, 3, 2, 2, 2, 195, 197, 7, 18, 2, 2, 196, 195,
	3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 198, 199, 3, 2,
	2, 2, 199, 201, 3, 2, 2, 2, 200, 202, 7, 19, 2, 2, 201, 200, 3, 2, 2, 2,
	201, 202, 3, 2, 2, 2, 202, 5, 3, 2, 2, 2, 203, 204, 7, 18, 2, 2, 204, 205,
	7, 18, 2, 2, 205, 7, 3, 2, 2, 2, 206, 207, 7, 18, 2, 2, 207, 208, 7, 18,
	2, 2, 208, 209, 7, 18, 2, 2, 209, 9, 3, 2, 2, 2, 210, 211, 7, 18, 2, 2,
	211, 212, 7, 18, 2, 2, 212, 213, 7, 18, 2, 2, 213, 214, 7, 18, 2, 2, 214,
	11, 3, 2, 2, 2, 215, 216, 9, 2, 2, 2, 216, 218, 7, 18, 2, 2, 217, 219,
	7, 18, 2, 2, 218, 217, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 218, 3, 2,
	2, 2, 220, 221, 3, 2, 2, 2, 221, 13, 3, 2, 2, 2, 222, 223, 9, 2, 2, 2,
	223, 224, 7, 18, 2, 2, 224, 225, 7, 18, 2, 2, 225, 227, 7, 18, 2, 2, 226,
	228, 7, 18, 2, 2, 227, 226, 3, 2, 2, 2, 228, 229, 3, 2, 2, 2, 229, 227,
	3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 15, 3, 2, 2, 2, 231, 232, 7, 18,
	2, 2, 232, 234, 7, 18, 2, 2, 233, 235, 7, 19, 2, 2, 234, 233, 3, 2, 2,
	2, 234, 235, 3, 2, 2, 2, 235, 17, 3, 2, 2, 2, 236, 239, 5, 6, 4, 2, 237,
	239, 5, 12, 7, 2, 238, 236, 3, 2, 2, 2, 238, 237, 3, 2, 2, 2, 239, 19,
	3, 2, 2, 2, 240, 243, 5, 10, 6, 2, 241, 243, 5, 14, 8, 2, 242, 240, 3,
	2, 2, 2, 242, 241, 3, 2, 2, 2, 243, 21, 3, 2, 2, 2, 244, 245, 5, 6, 4,
	2, 245, 23, 3, 2, 2, 2, 246, 247, 5, 6, 4, 2, 247, 25, 3, 2, 2, 2, 248,
	249, 5, 8, 5, 2, 249, 27, 3, 2, 2, 2, 250, 251, 5, 6, 4, 2, 251, 29, 3,
	2, 2, 2, 252, 253, 7, 18, 2, 2, 253, 31, 3, 2, 2, 2, 254, 255, 5, 6, 4,
	2, 255, 33, 3, 2, 2, 2, 256, 257, 5, 6, 4, 2, 257, 35, 3, 2, 2, 2, 258,
	259, 5, 6, 4, 2, 259, 37, 3, 2, 2, 2, 260, 261, 5, 16, 9, 2, 261, 39, 3,
	2, 2, 2, 262, 263, 5, 16, 9, 2, 263, 41, 3, 2, 2, 2, 264, 265, 5, 16, 9,
	2, 265, 43, 3, 2, 2, 2, 266, 269, 5, 46, 24, 2, 267, 269, 5, 48, 25, 2,
	268, 266, 3, 2, 2, 2, 268, 267, 3, 2, 2, 2, 269, 45, 3, 2, 2, 2, 270, 271,
	5, 20, 11, 2, 271, 272, 5, 22, 12, 2, 272, 273, 5, 24, 13, 2, 273, 47,
	3, 2, 2, 2, 274, 275, 5, 20, 11, 2, 275, 276, 7, 4, 2, 2, 276, 277, 5,
	22, 12, 2, 277, 278, 7, 4, 2, 2, 278, 279, 5, 24, 13, 2, 279, 49, 3, 2,
	2, 2, 280, 281, 5, 20, 11, 2, 281, 282, 7, 4, 2, 2, 282, 283, 5, 22, 12,
	2, 283, 51, 3, 2, 2, 2, 284, 285, 5, 20, 11, 2, 285, 53, 3, 2, 2, 2, 286,
	287, 5, 18, 10, 2, 287, 55, 3, 2, 2, 2, 288, 291, 5, 58, 30, 2, 289, 291,
	5, 60, 31, 2, 290, 288, 3, 2, 2, 2, 290, 289, 3, 2, 2, 2, 291, 57, 3, 2,
	2, 2, 292, 293, 5, 20, 11, 2, 293, 294, 5, 26, 14, 2, 294, 59, 3, 2, 2,
	2, 295, 296, 5, 20, 11, 2, 296, 297, 7, 4, 2, 2, 297, 298, 5, 26, 14, 2,
	298, 61, 3, 2, 2, 2, 299, 302, 5, 64, 33, 2, 300, 302, 5, 66, 34, 2, 301,
	299, 3, 2, 2, 2, 301, 300, 3, 2, 2, 2, 302, 63, 3, 2, 2, 2, 303, 304, 5,
	20, 11, 2, 304, 305, 7, 10, 2, 2, 305, 306, 5, 28, 15, 2, 306, 307, 5,
	30, 16, 2, 307, 65, 3, 2, 2, 2, 308, 309, 5, 20, 11, 2, 309, 310, 7, 4,
	2, 2, 310, 311, 7, 10, 2, 2, 311, 312, 5, 28, 15, 2, 312, 313, 7, 4, 2,
	2, 313, 314, 5, 30, 16, 2, 314, 67, 3, 2, 2, 2, 315, 318, 5, 70, 36, 2,
	316, 318, 5, 72, 37, 2, 317, 315, 3, 2, 2, 2, 317, 316, 3, 2, 2, 2, 318,
	69, 3, 2, 2, 2, 319, 320, 5, 20, 11, 2, 320, 321, 7, 10, 2, 2, 321, 322,
	5, 28, 15, 2, 322, 71, 3, 2, 2, 2, 323, 324, 5, 20, 11, 2, 324, 325, 7,
	4, 2, 2, 325, 326, 7, 10, 2, 2, 326, 327, 5, 28, 15, 2, 327, 73, 3, 2,
	2, 2, 328, 331, 5, 76, 39, 2, 329, 331, 5, 78, 40, 2, 330, 328, 3, 2, 2,
	2, 330, 329, 3, 2, 2, 2, 331, 75, 3, 2, 2, 2, 332, 336, 5, 46, 24, 2, 333,
	336, 5, 58, 30, 2, 334, 336, 5, 64, 33, 2, 335, 332, 3, 2, 2, 2, 335, 333,
	3, 2, 2, 2, 335, 334, 3, 2, 2, 2, 336, 77, 3, 2, 2, 2, 337, 341, 5, 48,
	25, 2, 338, 341, 5, 60, 31, 2, 339, 341, 5, 66, 34, 2, 340, 337, 3, 2,
	2, 2, 340, 338, 3, 2, 2, 2, 340, 339, 3, 2, 2, 2, 341, 79, 3, 2, 2, 2,
	342, 345, 5, 82, 42, 2, 343, 345, 5, 84, 43, 2, 344, 342, 3, 2, 2, 2, 344,
	343, 3, 2, 2, 2, 345, 81, 3, 2, 2, 2, 346, 352, 5, 76, 39, 2, 347, 352,
	5, 50, 26, 2, 348, 352, 5, 52, 27, 2, 349, 352, 5, 54, 28, 2, 350, 352,
	5, 70, 36, 2, 351, 346, 3, 2, 2, 2, 351, 347, 3, 2, 2, 2, 351, 348, 3,
	2, 2, 2, 351, 349, 3, 2, 2, 2, 351, 350, 3, 2, 2, 2, 352, 83, 3, 2, 2,
	2, 353, 359, 5, 78, 40, 2, 354, 359, 5, 50, 26, 2, 355, 359, 5, 52, 27,
	2, 356, 359, 5, 54, 28, 2, 357, 359, 5, 72, 37, 2, 358, 353, 3, 2, 2, 2,
	358, 354, 3, 2, 2, 2, 358, 355, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 358,
	357, 3, 2, 2, 2, 359, 85, 3, 2, 2, 2, 360, 363, 5, 88, 45, 2, 361, 363,
	5, 90, 46, 2, 362, 360, 3, 2, 2, 2, 362, 361, 3, 2, 2, 2, 363, 87, 3, 2,
	2, 2, 364, 365, 5, 32, 17, 2, 365, 366, 5, 34, 18, 2, 366, 367, 5, 42,
	22, 2, 367, 89, 3, 2, 2, 2, 368, 369, 5, 32, 17, 2, 369, 370, 7, 5, 2,
	2, 370, 371, 5, 34, 18, 2, 371, 372, 7, 5, 2, 2, 372, 373, 5, 42, 22, 2,
	373, 91, 3, 2, 2, 2, 374, 377, 5, 94, 48, 2, 375, 377, 5, 96, 49, 2, 376,
	374, 3, 2, 2, 2, 376, 375, 3, 2, 2, 2, 377, 93, 3, 2, 2, 2, 378, 379, 5,
	32, 17, 2, 379, 380, 5, 40, 21, 2, 380, 95, 3, 2, 2, 2, 381, 382, 5, 32,
	17, 2, 382, 383, 7, 5, 2, 2, 383, 384, 5, 40, 21, 2, 384, 97, 3, 2, 2,
	2, 385, 386, 5, 38, 20, 2, 386, 99, 3, 2, 2, 2, 387, 390, 5, 102, 52, 2,
	388, 390, 5, 104, 53, 2, 389, 387, 3, 2, 2, 2, 389, 388, 3, 2, 2, 2, 390,
	101, 3, 2, 2, 2, 391, 395, 5, 88, 45, 2, 392, 395, 5, 94, 48, 2, 393, 395,
	5, 98, 50, 2, 394, 391, 3, 2, 2, 2, 394, 392, 3, 2, 2, 2, 394, 393, 3,
	2, 2, 2, 395, 103, 3, 2, 2, 2, 396, 400, 5, 90, 46, 2, 397, 400, 5, 96,
	49, 2, 398, 400, 5, 98, 50, 2, 399, 396, 3, 2, 2, 2, 399, 397, 3, 2, 2,
	2, 399, 398, 3, 2, 2, 2, 400, 105, 3, 2, 2, 2, 401, 402, 7, 9, 2, 2, 402,
	107, 3, 2, 2, 2, 403, 406, 5, 110, 56, 2, 404, 406, 5, 112, 57, 2, 405,
	403, 3, 2, 2, 2, 405, 404, 3, 2, 2, 2, 406, 109, 3, 2, 2, 2, 407, 408,
	9, 2, 2, 2, 408, 410, 5, 32, 17, 2, 409, 411, 5, 34, 18, 2, 410, 409, 3,
	2, 2, 2, 410, 411, 3, 2, 2, 2, 411, 414, 3, 2, 2, 2, 412, 414, 5, 106,
	54, 2, 413, 407, 3, 2, 2, 2, 413, 412, 3, 2, 2, 2, 414, 111, 3, 2, 2, 2,
	415, 416, 9, 2, 2, 2, 416, 419, 5, 32, 17, 2, 417, 418, 7, 5, 2, 2, 418,
	420, 5, 34, 18, 2, 419, 417, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420, 423,
	3, 2, 2, 2, 421, 423, 5, 106, 54, 2, 422, 415, 3, 2, 2, 2, 422, 421, 3,
	2, 2, 2, 423, 113, 3, 2, 2, 2, 424, 427, 5, 116, 59, 2, 425, 427, 5, 118,
	60, 2, 426, 424, 3, 2, 2, 2, 426, 425, 3, 2, 2, 2, 427, 115, 3, 2, 2, 2,
	428, 429, 5, 102, 52, 2, 429, 430, 5, 110, 56, 2, 430, 117, 3, 2, 2, 2,
	431, 432, 5, 104, 53, 2, 432, 433, 5, 112, 57, 2, 433, 119, 3, 2, 2, 2,
	434, 436, 7, 8, 2, 2, 435, 434, 3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436,
	437, 3, 2, 2, 2, 437, 439, 5, 100, 51, 2, 438, 440, 5, 108, 55, 2, 439,
	438, 3, 2, 2, 2, 439, 440, 3, 2, 2, 2, 440, 121, 3, 2, 2, 2, 441, 444,
	5, 124, 63, 2, 442, 444, 5, 126, 64, 2, 443, 441, 3, 2, 2, 2, 443, 442,
	3, 2, 2, 2, 444, 123, 3, 2, 2, 2, 445, 446, 5, 46, 24, 2, 446, 447, 7,
	8, 2, 2, 447, 449, 5, 88, 45, 2, 448, 450, 5, 110, 56, 2, 449, 448, 3,
	2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 125, 3, 2, 2, 2, 451, 452, 5, 48, 25,
	2, 452, 453, 7, 8, 2, 2, 453, 455, 5, 90, 46, 2, 454, 456, 5, 112, 57,
	2, 455, 454, 3, 2, 2, 2, 455, 456, 3, 2, 2, 2, 456, 127, 3, 2, 2, 2, 457,
	460, 5, 130, 66, 2, 458, 460, 5, 132, 67, 2, 459, 457, 3, 2, 2, 2, 459,
	458, 3, 2, 2, 2, 460, 129, 3, 2, 2, 2, 461, 462, 5, 76, 39, 2, 462, 463,
	7, 8, 2, 2, 463, 465, 5, 102, 52, 2, 464, 466, 5, 110, 56, 2, 465, 464,
	3, 2, 2, 2, 465, 466, 3, 2, 2, 2, 466, 131, 3, 2, 2, 2, 467, 468, 5, 78,
	40, 2, 468, 469, 7, 8, 2, 2, 469, 471, 5, 104, 53, 2, 470, 472, 5, 112,
	57, 2, 471, 470, 3, 2, 2, 2, 471, 472, 3, 2, 2, 2, 472, 133, 3, 2, 2, 2,
	473, 476, 5, 140, 71, 2, 474, 476, 5, 150, 76, 2, 475, 473, 3, 2, 2, 2,
	475, 474, 3, 2, 2, 2, 476, 135, 3, 2, 2, 2, 477, 480, 5, 140, 71, 2, 478,
	480, 5, 152, 77, 2, 479, 477, 3, 2, 2, 2, 479, 478, 3, 2, 2, 2, 480, 137,
	3, 2, 2, 2, 481, 484, 5, 140, 71, 2, 482, 484, 5, 154, 78, 2, 483, 481,
	3, 2, 2, 2, 483, 482, 3, 2, 2, 2, 484, 139, 3, 2, 2, 2, 485, 489, 7, 11,
	2, 2, 486, 487, 5, 2, 2, 2, 487, 488, 7, 12, 2, 2, 488, 490, 3, 2, 2, 2,
	489, 486, 3, 2, 2, 2, 489, 490, 3, 2, 2, 2, 490, 494, 3, 2, 2, 2, 491,
	492, 5, 2, 2, 2, 492, 493, 7, 13, 2, 2, 493, 495, 3, 2, 2, 2, 494, 491,
	3, 2, 2, 2, 494, 495, 3, 2, 2, 2, 495, 499, 3, 2, 2, 2, 496, 497, 5, 2,
	2, 2, 497, 498, 7, 14, 2, 2, 498, 500, 3, 2, 2, 2, 499, 496, 3, 2, 2, 2,
	499, 500, 3, 2, 2, 2, 500, 501, 3, 2, 2, 2, 501, 505, 7, 8, 2, 2, 502,
	503, 5, 2, 2, 2, 503, 504, 7, 15, 2, 2, 504, 506, 3, 2, 2, 2, 505, 502,
	3, 2, 2, 2, 505, 506, 3, 2, 2, 2, 506, 510, 3, 2, 2, 2, 507, 508, 5, 2,
	2, 2, 508, 509, 7, 13, 2, 2, 509, 511, 3, 2, 2, 2, 510, 507, 3, 2, 2, 2,
	510, 511, 3, 2, 2, 2, 511, 512, 3, 2, 2, 2, 512, 513, 5, 4, 3, 2, 513,
	514, 7, 16, 2, 2, 514, 592, 3, 2, 2, 2, 515, 519, 7, 11, 2, 2, 516, 517,
	5, 2, 2, 2, 517, 518, 7, 12, 2, 2, 518, 520, 3, 2, 2, 2, 519, 516, 3, 2,
	2, 2, 519, 520, 3, 2, 2, 2, 520, 524, 3, 2, 2, 2, 521, 522, 5, 2, 2, 2,
	522, 523, 7, 13, 2, 2, 523, 525, 3, 2, 2, 2, 524, 521, 3, 2, 2, 2, 524,
	525, 3, 2, 2, 2, 525, 529, 3, 2, 2, 2, 526, 527, 5, 2, 2, 2, 527, 528,
	7, 14, 2, 2, 528, 530, 3, 2, 2, 2, 529, 526, 3, 2, 2, 2, 529, 530, 3, 2,
	2, 2, 530, 531, 3, 2, 2, 2, 531, 535, 7, 8, 2, 2, 532, 533, 5, 2, 2, 2,
	533, 534, 7, 15, 2, 2, 534, 536, 3, 2, 2, 2, 535, 532, 3, 2, 2, 2, 535,
	536, 3, 2, 2, 2, 536, 537, 3, 2, 2, 2, 537, 538, 5, 4, 3, 2, 538, 539,
	7, 13, 2, 2, 539, 592, 3, 2, 2, 2, 540, 544, 7, 11, 2, 2, 541, 542, 5,
	2, 2, 2, 542, 543, 7, 12, 2, 2, 543, 545, 3, 2, 2, 2, 544, 541, 3, 2, 2,
	2, 544, 545, 3, 2, 2, 2, 545, 549, 3, 2, 2, 2, 546, 547, 5, 2, 2, 2, 547,
	548, 7, 13, 2, 2, 548, 550, 3, 2, 2, 2, 549, 546, 3, 2, 2, 2, 549, 550,
	3, 2, 2, 2, 550, 554, 3, 2, 2, 2, 551, 552, 5, 2, 2, 2, 552, 553, 7, 14,
	2, 2, 553, 555, 3, 2, 2, 2, 554, 551, 3, 2, 2, 2, 554, 555, 3, 2, 2, 2,
	555, 556, 3, 2, 2, 2, 556, 557, 7, 8, 2, 2, 557, 558, 5, 4, 3, 2, 558,
	559, 7, 15, 2, 2, 559, 592, 3, 2, 2, 2, 560, 564, 7, 11, 2, 2, 561, 562,
	5, 2, 2, 2, 562, 563, 7, 12, 2, 2, 563, 565, 3, 2, 2, 2, 564, 561, 3, 2,
	2, 2, 564, 565, 3, 2, 2, 2, 565, 569, 3, 2, 2, 2, 566, 567, 5, 2, 2, 2,
	567, 568, 7, 13, 2, 2, 568, 570, 3, 2, 2, 2, 569, 566, 3, 2, 2, 2, 569,
	570, 3, 2, 2, 2, 570, 571, 3, 2, 2, 2, 571, 572, 5, 4, 3, 2, 572, 573,
	7, 14, 2, 2, 573, 592, 3, 2, 2, 2, 574, 578, 7, 11, 2, 2, 575, 576, 5,
	2, 2, 2, 576, 577, 7, 12, 2, 2, 577, 579, 3, 2, 2, 2, 578, 575, 3, 2, 2,
	2, 578, 579, 3, 2, 2, 2, 579, 580, 3, 2, 2, 2, 580, 581, 5, 4, 3, 2, 581,
	582, 7, 13, 2, 2, 582, 592, 3, 2, 2, 2, 583, 584, 7, 11, 2, 2, 584, 585,
	5, 4, 3, 2, 585, 586, 7, 12, 2, 2, 586, 592, 3, 2, 2, 2, 587, 588, 7, 11,
	2, 2, 588, 589, 5, 4, 3, 2, 589, 590, 7, 10, 2, 2, 590, 592, 3, 2, 2, 2,
	591, 485, 3, 2, 2, 2, 591, 515, 3, 2, 2, 2, 591, 540, 3, 2, 2, 2, 591,
	560, 3, 2, 2, 2, 591, 574, 3, 2, 2, 2, 591, 583, 3, 2, 2, 2, 591, 587,
	3, 2, 2, 2, 592, 141, 3, 2, 2, 2, 593, 594, 5, 22, 12, 2, 594, 595, 5,
	24, 13, 2, 595, 143, 3, 2, 2, 2, 596, 597, 5, 22, 12, 2, 597, 598, 7, 4,
	2, 2, 598, 599, 5, 24, 13, 2, 599, 145, 3, 2, 2, 2, 600, 606, 5, 142, 72,
	2, 601, 606, 5, 24, 13, 2, 602, 606, 5, 130, 66, 2, 603, 606, 5, 82, 42,
	2, 604, 606, 5, 102, 52, 2, 605, 600, 3, 2, 2, 2, 605, 601, 3, 2, 2, 2,
	605, 602, 3, 2, 2, 2, 605, 603, 3, 2, 2, 2, 605, 604, 3, 2, 2, 2, 606,
	147, 3, 2, 2, 2, 607, 614, 5, 144, 73, 2, 608, 614, 5, 24, 13, 2, 609,
	614, 5, 132, 67, 2, 610, 614, 5, 84, 43, 2, 611, 614, 5, 104, 53, 2, 612,
	614, 5, 144, 73, 2, 613, 607, 3, 2, 2, 2, 613, 608, 3, 2, 2, 2, 613, 609,
	3, 2, 2, 2, 613, 610, 3, 2, 2, 2, 613, 611, 3, 2, 2, 2, 613, 612, 3, 2,
	2, 2, 614, 149, 3, 2, 2, 2, 615, 618, 5, 152, 77, 2, 616, 618, 5, 154,
	78, 2, 617, 615, 3, 2, 2, 2, 617, 616, 3, 2, 2, 2, 618, 151, 3, 2, 2, 2,
	619, 620, 7, 11, 2, 2, 620, 621, 5, 146, 74, 2, 621, 153, 3, 2, 2, 2, 622,
	623, 7, 11, 2, 2, 623, 624, 5, 148, 75, 2, 624, 155, 3, 2, 2, 2, 625, 628,
	5, 158, 80, 2, 626, 628, 5, 160, 81, 2, 627, 625, 3, 2, 2, 2, 627, 626,
	3, 2, 2, 2, 628, 157, 3, 2, 2, 2, 629, 630, 5, 146, 74, 2, 630, 631, 7,
	6, 2, 2, 631, 632, 5, 146, 74, 2, 632, 159, 3, 2, 2, 2, 633, 634, 5, 148,
	75, 2, 634, 635, 7, 6, 2, 2, 635, 636, 5, 148, 75, 2, 636, 161, 3, 2, 2,
	2, 637, 640, 5, 164, 83, 2, 638, 640, 5, 166, 84, 2, 639, 637, 3, 2, 2,
	2, 639, 638, 3, 2, 2, 2, 640, 163, 3, 2, 2, 2, 641, 642, 5, 130, 66, 2,
	642, 643, 7, 6, 2, 2, 643, 644, 5, 136, 69, 2, 644, 165, 3, 2, 2, 2, 645,
	646, 5, 132, 67, 2, 646, 647, 7, 6, 2, 2, 647, 648, 5, 138, 70, 2, 648,
	167, 3, 2, 2, 2, 649, 652, 5, 170, 86, 2, 650, 652, 5, 172, 87, 2, 651,
	649, 3, 2, 2, 2, 651, 650, 3, 2, 2, 2, 652, 169, 3, 2, 2, 2, 653, 654,
	5, 136, 69, 2, 654, 655, 7, 6, 2, 2, 655, 656, 5, 130, 66, 2, 656, 171,
	3, 2, 2, 2, 657, 658, 5, 138, 70, 2, 658, 659, 7, 6, 2, 2, 659, 660, 5,
	132, 67, 2, 660, 173, 3, 2, 2, 2, 661, 664, 5, 176, 89, 2, 662, 664, 5,
	178, 90, 2, 663, 661, 3, 2, 2, 2, 663, 662, 3, 2, 2, 2, 664, 175, 3, 2,
	2, 2, 665, 670, 5, 158, 80, 2, 666, 670, 5, 164, 83, 2, 667, 670, 5, 170,
	86, 2, 668, 670, 5, 136, 69, 2, 669, 665, 3, 2, 2, 2, 669, 666, 3, 2, 2,
	2, 669, 667, 3, 2, 2, 2, 669, 668, 3, 2, 2, 2, 670, 177, 3, 2, 2, 2, 671,
	676, 5, 160, 81, 2, 672, 676, 5, 166, 84, 2, 673, 676, 5, 172, 87, 2, 674,
	676, 5, 138, 70, 2, 675, 671, 3, 2, 2, 2, 675, 672, 3, 2, 2, 2, 675, 673,
	3, 2, 2, 2, 675, 674, 3, 2, 2, 2, 676, 179, 3, 2, 2, 2, 677, 679, 7, 17,
	2, 2, 678, 680, 5, 2, 2, 2, 679, 678, 3, 2, 2, 2, 679, 680, 3, 2, 2, 2,
	680, 181, 3, 2, 2, 2, 681, 684, 5, 184, 93, 2, 682, 684, 5, 186, 94, 2,
	683, 681, 3, 2, 2, 2, 683, 682, 3, 2, 2, 2, 684, 183, 3, 2, 2, 2, 685,
	686, 5, 180, 91, 2, 686, 687, 7, 6, 2, 2, 687, 688, 5, 176, 89, 2, 688,
	185, 3, 2, 2, 2, 689, 690, 5, 180, 91, 2, 690, 691, 7, 6, 2, 2, 691, 692,
	5, 178, 90, 2, 692, 187, 3, 2, 2, 2, 693, 699, 5, 120, 61, 2, 694, 699,
	5, 80, 41, 2, 695, 699, 5, 128, 65, 2, 696, 699, 5, 174, 88, 2, 697, 699,
	5, 182, 92, 2, 698, 693, 3, 2, 2, 2, 698, 694, 3, 2, 2, 2, 698, 695, 3,
	2, 2, 2, 698, 696, 3, 2, 2, 2, 698, 697, 3, 2, 2, 2, 699, 700, 3, 2, 2,
	2, 700, 701, 7, 2, 2, 3, 701, 189, 3, 2, 2, 2, 70, 193, 198, 201, 220,
	229, 234, 238, 242, 268, 290, 301, 317, 330, 335, 340, 344, 351, 358, 362,
	376, 389, 394, 399, 405, 410, 413, 419, 422, 426, 435, 439, 443, 449, 455,
	459, 465, 471, 475, 479, 483, 489, 494, 499, 505, 510, 519, 524, 529, 535,
	544, 549, 554, 564, 569, 578, 591, 605, 613, 617, 627, 639, 651, 663, 669,
	675, 679, 683, 698,
}
var literalNames = []string{
	"", "'+'", "'-'", "':'", "'/'",
}
var symbolicNames = []string{
	"", "", "", "", "", "Newline", "T", "Z", "W", "P", "Y", "M", "D", "H",
	"S", "R", "Digit", "Fraction",
}

var ruleNames = []string{
	"int_", "dec", "int2", "int3", "int4", "sint2p", "sint4p", "dec2", "century",
	"year", "month", "day", "ordinalDay", "week", "weekDay", "hour", "minute",
	"second", "hourFraction", "minuteFraction", "secondFraction", "calendarDate",
	"calendarDateBasic", "calendarDateExtended", "specificMonthCalendarDate",
	"specificYearCalendarDate", "specificCenturyCalendarDate", "ordinalDate",
	"ordinalDateBasic", "ordinalDateExtended", "weekDate", "weekDateBasic",
	"weekDateExtended", "specificWeekWeekDate", "specificWeekWeekDateBasic",
	"specificWeekWeekDateExtended", "datePrecise", "datePreciseBasic", "datePreciseExtended",
	"date", "dateBasic", "dateExtended", "localTimePrecise", "localTimePreciseBasic",
	"localTimePreciseExtended", "specificMinuteLocalTime", "specificMinuteLocalTimeBasic",
	"specificMinuteLocalTimeExtended", "specificHourLocalTime", "localTime",
	"localTimeBasic", "localTimeExtended", "timeZoneUtc", "timeZone", "timeZoneBasic",
	"timeZoneExtended", "localTimeAndTimeZone", "localTimeAndTimeZoneBasic",
	"localTimeAndTimeZoneExtended", "time", "datetimePrecise", "datetimePreciseBasic",
	"datetimePreciseExtended", "datetime", "datetimeBasic", "datetimeExtended",
	"interval", "intervalBasic", "intervalExtended", "intervalT", "monthDateBasic",
	"monthDateExtended", "intervalYMDTimeBasic", "intervalYMDTimeExtended",
	"intervalYMD", "intervalYMDBasic", "intervalYMDExtended", "timeBeginEnd",
	"timeBeginEndBasic", "timeBeginEndExtended", "timeBeginInterval", "timeBeginIntervalBasic",
	"timeBeginIntervalExtended", "timeIntervalEnd", "timeIntervalEndBasic",
	"timeIntervalEndExtended", "duration", "durationBasic", "durationExtended",
	"recurringCount", "recurring", "recurringBasic", "recurringExtended", "iso",
}

type iso8601Parser struct {
	*antlr.BaseParser
}

// Newiso8601Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *iso8601Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func Newiso8601Parser(input antlr.TokenStream) *iso8601Parser {
	this := new(iso8601Parser)
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
	this.GrammarFileName = "iso8601.g4"

	return this
}

// iso8601Parser tokens.
const (
	iso8601ParserEOF      = antlr.TokenEOF
	iso8601ParserT__0     = 1
	iso8601ParserT__1     = 2
	iso8601ParserT__2     = 3
	iso8601ParserT__3     = 4
	iso8601ParserNewline  = 5
	iso8601ParserT        = 6
	iso8601ParserZ        = 7
	iso8601ParserW        = 8
	iso8601ParserP        = 9
	iso8601ParserY        = 10
	iso8601ParserM        = 11
	iso8601ParserD        = 12
	iso8601ParserH        = 13
	iso8601ParserS        = 14
	iso8601ParserR        = 15
	iso8601ParserDigit    = 16
	iso8601ParserFraction = 17
)

// iso8601Parser rules.
const (
	iso8601ParserRULE_int_                            = 0
	iso8601ParserRULE_dec                             = 1
	iso8601ParserRULE_int2                            = 2
	iso8601ParserRULE_int3                            = 3
	iso8601ParserRULE_int4                            = 4
	iso8601ParserRULE_sint2p                          = 5
	iso8601ParserRULE_sint4p                          = 6
	iso8601ParserRULE_dec2                            = 7
	iso8601ParserRULE_century                         = 8
	iso8601ParserRULE_year                            = 9
	iso8601ParserRULE_month                           = 10
	iso8601ParserRULE_day                             = 11
	iso8601ParserRULE_ordinalDay                      = 12
	iso8601ParserRULE_week                            = 13
	iso8601ParserRULE_weekDay                         = 14
	iso8601ParserRULE_hour                            = 15
	iso8601ParserRULE_minute                          = 16
	iso8601ParserRULE_second                          = 17
	iso8601ParserRULE_hourFraction                    = 18
	iso8601ParserRULE_minuteFraction                  = 19
	iso8601ParserRULE_secondFraction                  = 20
	iso8601ParserRULE_calendarDate                    = 21
	iso8601ParserRULE_calendarDateBasic               = 22
	iso8601ParserRULE_calendarDateExtended            = 23
	iso8601ParserRULE_specificMonthCalendarDate       = 24
	iso8601ParserRULE_specificYearCalendarDate        = 25
	iso8601ParserRULE_specificCenturyCalendarDate     = 26
	iso8601ParserRULE_ordinalDate                     = 27
	iso8601ParserRULE_ordinalDateBasic                = 28
	iso8601ParserRULE_ordinalDateExtended             = 29
	iso8601ParserRULE_weekDate                        = 30
	iso8601ParserRULE_weekDateBasic                   = 31
	iso8601ParserRULE_weekDateExtended                = 32
	iso8601ParserRULE_specificWeekWeekDate            = 33
	iso8601ParserRULE_specificWeekWeekDateBasic       = 34
	iso8601ParserRULE_specificWeekWeekDateExtended    = 35
	iso8601ParserRULE_datePrecise                     = 36
	iso8601ParserRULE_datePreciseBasic                = 37
	iso8601ParserRULE_datePreciseExtended             = 38
	iso8601ParserRULE_date                            = 39
	iso8601ParserRULE_dateBasic                       = 40
	iso8601ParserRULE_dateExtended                    = 41
	iso8601ParserRULE_localTimePrecise                = 42
	iso8601ParserRULE_localTimePreciseBasic           = 43
	iso8601ParserRULE_localTimePreciseExtended        = 44
	iso8601ParserRULE_specificMinuteLocalTime         = 45
	iso8601ParserRULE_specificMinuteLocalTimeBasic    = 46
	iso8601ParserRULE_specificMinuteLocalTimeExtended = 47
	iso8601ParserRULE_specificHourLocalTime           = 48
	iso8601ParserRULE_localTime                       = 49
	iso8601ParserRULE_localTimeBasic                  = 50
	iso8601ParserRULE_localTimeExtended               = 51
	iso8601ParserRULE_timeZoneUtc                     = 52
	iso8601ParserRULE_timeZone                        = 53
	iso8601ParserRULE_timeZoneBasic                   = 54
	iso8601ParserRULE_timeZoneExtended                = 55
	iso8601ParserRULE_localTimeAndTimeZone            = 56
	iso8601ParserRULE_localTimeAndTimeZoneBasic       = 57
	iso8601ParserRULE_localTimeAndTimeZoneExtended    = 58
	iso8601ParserRULE_time                            = 59
	iso8601ParserRULE_datetimePrecise                 = 60
	iso8601ParserRULE_datetimePreciseBasic            = 61
	iso8601ParserRULE_datetimePreciseExtended         = 62
	iso8601ParserRULE_datetime                        = 63
	iso8601ParserRULE_datetimeBasic                   = 64
	iso8601ParserRULE_datetimeExtended                = 65
	iso8601ParserRULE_interval                        = 66
	iso8601ParserRULE_intervalBasic                   = 67
	iso8601ParserRULE_intervalExtended                = 68
	iso8601ParserRULE_intervalT                       = 69
	iso8601ParserRULE_monthDateBasic                  = 70
	iso8601ParserRULE_monthDateExtended               = 71
	iso8601ParserRULE_intervalYMDTimeBasic            = 72
	iso8601ParserRULE_intervalYMDTimeExtended         = 73
	iso8601ParserRULE_intervalYMD                     = 74
	iso8601ParserRULE_intervalYMDBasic                = 75
	iso8601ParserRULE_intervalYMDExtended             = 76
	iso8601ParserRULE_timeBeginEnd                    = 77
	iso8601ParserRULE_timeBeginEndBasic               = 78
	iso8601ParserRULE_timeBeginEndExtended            = 79
	iso8601ParserRULE_timeBeginInterval               = 80
	iso8601ParserRULE_timeBeginIntervalBasic          = 81
	iso8601ParserRULE_timeBeginIntervalExtended       = 82
	iso8601ParserRULE_timeIntervalEnd                 = 83
	iso8601ParserRULE_timeIntervalEndBasic            = 84
	iso8601ParserRULE_timeIntervalEndExtended         = 85
	iso8601ParserRULE_duration                        = 86
	iso8601ParserRULE_durationBasic                   = 87
	iso8601ParserRULE_durationExtended                = 88
	iso8601ParserRULE_recurringCount                  = 89
	iso8601ParserRULE_recurring                       = 90
	iso8601ParserRULE_recurringBasic                  = 91
	iso8601ParserRULE_recurringExtended               = 92
	iso8601ParserRULE_iso                             = 93
)

// IInt_Context is an interface to support dynamic dispatch.
type IInt_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInt_Context differentiates from other interfaces.
	IsInt_Context()
}

type Int_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInt_Context() *Int_Context {
	var p = new(Int_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_int_
	return p
}

func (*Int_Context) IsInt_Context() {}

func NewInt_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Int_Context {
	var p = new(Int_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_int_

	return p
}

func (s *Int_Context) GetParser() antlr.Parser { return s.parser }

func (s *Int_Context) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(iso8601ParserDigit)
}

func (s *Int_Context) Digit(i int) antlr.TerminalNode {
	return s.GetToken(iso8601ParserDigit, i)
}

func (s *Int_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Int_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Int_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterInt_(s)
	}
}

func (s *Int_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitInt_(s)
	}
}

func (p *iso8601Parser) Int_() (localctx IInt_Context) {
	this := p
	_ = this

	localctx = NewInt_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, iso8601ParserRULE_int_)
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
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == iso8601ParserDigit {
		{
			p.SetState(188)
			p.Match(iso8601ParserDigit)
		}

		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDecContext is an interface to support dynamic dispatch.
type IDecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecContext differentiates from other interfaces.
	IsDecContext()
}

type DecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecContext() *DecContext {
	var p = new(DecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_dec
	return p
}

func (*DecContext) IsDecContext() {}

func NewDecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecContext {
	var p = new(DecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_dec

	return p
}

func (s *DecContext) GetParser() antlr.Parser { return s.parser }

func (s *DecContext) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(iso8601ParserDigit)
}

func (s *DecContext) Digit(i int) antlr.TerminalNode {
	return s.GetToken(iso8601ParserDigit, i)
}

func (s *DecContext) Fraction() antlr.TerminalNode {
	return s.GetToken(iso8601ParserFraction, 0)
}

func (s *DecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDec(s)
	}
}

func (s *DecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDec(s)
	}
}

func (p *iso8601Parser) Dec() (localctx IDecContext) {
	this := p
	_ = this

	localctx = NewDecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, iso8601ParserRULE_dec)
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

	for ok := true; ok; ok = _la == iso8601ParserDigit {
		{
			p.SetState(193)
			p.Match(iso8601ParserDigit)
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == iso8601ParserFraction {
		{
			p.SetState(198)
			p.Match(iso8601ParserFraction)
		}

	}

	return localctx
}

// IInt2Context is an interface to support dynamic dispatch.
type IInt2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInt2Context differentiates from other interfaces.
	IsInt2Context()
}

type Int2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInt2Context() *Int2Context {
	var p = new(Int2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_int2
	return p
}

func (*Int2Context) IsInt2Context() {}

func NewInt2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Int2Context {
	var p = new(Int2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_int2

	return p
}

func (s *Int2Context) GetParser() antlr.Parser { return s.parser }

func (s *Int2Context) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(iso8601ParserDigit)
}

func (s *Int2Context) Digit(i int) antlr.TerminalNode {
	return s.GetToken(iso8601ParserDigit, i)
}

func (s *Int2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Int2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Int2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterInt2(s)
	}
}

func (s *Int2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitInt2(s)
	}
}

func (p *iso8601Parser) Int2() (localctx IInt2Context) {
	this := p
	_ = this

	localctx = NewInt2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, iso8601ParserRULE_int2)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(201)
		p.Match(iso8601ParserDigit)
	}
	{
		p.SetState(202)
		p.Match(iso8601ParserDigit)
	}

	return localctx
}

// IInt3Context is an interface to support dynamic dispatch.
type IInt3Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInt3Context differentiates from other interfaces.
	IsInt3Context()
}

type Int3Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInt3Context() *Int3Context {
	var p = new(Int3Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_int3
	return p
}

func (*Int3Context) IsInt3Context() {}

func NewInt3Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Int3Context {
	var p = new(Int3Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_int3

	return p
}

func (s *Int3Context) GetParser() antlr.Parser { return s.parser }

func (s *Int3Context) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(iso8601ParserDigit)
}

func (s *Int3Context) Digit(i int) antlr.TerminalNode {
	return s.GetToken(iso8601ParserDigit, i)
}

func (s *Int3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Int3Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Int3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterInt3(s)
	}
}

func (s *Int3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitInt3(s)
	}
}

func (p *iso8601Parser) Int3() (localctx IInt3Context) {
	this := p
	_ = this

	localctx = NewInt3Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, iso8601ParserRULE_int3)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(204)
		p.Match(iso8601ParserDigit)
	}
	{
		p.SetState(205)
		p.Match(iso8601ParserDigit)
	}
	{
		p.SetState(206)
		p.Match(iso8601ParserDigit)
	}

	return localctx
}

// IInt4Context is an interface to support dynamic dispatch.
type IInt4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInt4Context differentiates from other interfaces.
	IsInt4Context()
}

type Int4Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInt4Context() *Int4Context {
	var p = new(Int4Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_int4
	return p
}

func (*Int4Context) IsInt4Context() {}

func NewInt4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Int4Context {
	var p = new(Int4Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_int4

	return p
}

func (s *Int4Context) GetParser() antlr.Parser { return s.parser }

func (s *Int4Context) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(iso8601ParserDigit)
}

func (s *Int4Context) Digit(i int) antlr.TerminalNode {
	return s.GetToken(iso8601ParserDigit, i)
}

func (s *Int4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Int4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Int4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterInt4(s)
	}
}

func (s *Int4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitInt4(s)
	}
}

func (p *iso8601Parser) Int4() (localctx IInt4Context) {
	this := p
	_ = this

	localctx = NewInt4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, iso8601ParserRULE_int4)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(208)
		p.Match(iso8601ParserDigit)
	}
	{
		p.SetState(209)
		p.Match(iso8601ParserDigit)
	}
	{
		p.SetState(210)
		p.Match(iso8601ParserDigit)
	}
	{
		p.SetState(211)
		p.Match(iso8601ParserDigit)
	}

	return localctx
}

// ISint2pContext is an interface to support dynamic dispatch.
type ISint2pContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSint2pContext differentiates from other interfaces.
	IsSint2pContext()
}

type Sint2pContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySint2pContext() *Sint2pContext {
	var p = new(Sint2pContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_sint2p
	return p
}

func (*Sint2pContext) IsSint2pContext() {}

func NewSint2pContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sint2pContext {
	var p = new(Sint2pContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_sint2p

	return p
}

func (s *Sint2pContext) GetParser() antlr.Parser { return s.parser }

func (s *Sint2pContext) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(iso8601ParserDigit)
}

func (s *Sint2pContext) Digit(i int) antlr.TerminalNode {
	return s.GetToken(iso8601ParserDigit, i)
}

func (s *Sint2pContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sint2pContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sint2pContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSint2p(s)
	}
}

func (s *Sint2pContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSint2p(s)
	}
}

func (p *iso8601Parser) Sint2p() (localctx ISint2pContext) {
	this := p
	_ = this

	localctx = NewSint2pContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, iso8601ParserRULE_sint2p)
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
		p.SetState(213)
		_la = p.GetTokenStream().LA(1)

		if !(_la == iso8601ParserT__0 || _la == iso8601ParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(214)
		p.Match(iso8601ParserDigit)
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == iso8601ParserDigit {
		{
			p.SetState(215)
			p.Match(iso8601ParserDigit)
		}

		p.SetState(218)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISint4pContext is an interface to support dynamic dispatch.
type ISint4pContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSint4pContext differentiates from other interfaces.
	IsSint4pContext()
}

type Sint4pContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySint4pContext() *Sint4pContext {
	var p = new(Sint4pContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_sint4p
	return p
}

func (*Sint4pContext) IsSint4pContext() {}

func NewSint4pContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Sint4pContext {
	var p = new(Sint4pContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_sint4p

	return p
}

func (s *Sint4pContext) GetParser() antlr.Parser { return s.parser }

func (s *Sint4pContext) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(iso8601ParserDigit)
}

func (s *Sint4pContext) Digit(i int) antlr.TerminalNode {
	return s.GetToken(iso8601ParserDigit, i)
}

func (s *Sint4pContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Sint4pContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Sint4pContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSint4p(s)
	}
}

func (s *Sint4pContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSint4p(s)
	}
}

func (p *iso8601Parser) Sint4p() (localctx ISint4pContext) {
	this := p
	_ = this

	localctx = NewSint4pContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, iso8601ParserRULE_sint4p)
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
		p.SetState(220)
		_la = p.GetTokenStream().LA(1)

		if !(_la == iso8601ParserT__0 || _la == iso8601ParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(221)
		p.Match(iso8601ParserDigit)
	}
	{
		p.SetState(222)
		p.Match(iso8601ParserDigit)
	}
	{
		p.SetState(223)
		p.Match(iso8601ParserDigit)
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(224)
				p.Match(iso8601ParserDigit)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IDec2Context is an interface to support dynamic dispatch.
type IDec2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDec2Context differentiates from other interfaces.
	IsDec2Context()
}

type Dec2Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDec2Context() *Dec2Context {
	var p = new(Dec2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_dec2
	return p
}

func (*Dec2Context) IsDec2Context() {}

func NewDec2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dec2Context {
	var p = new(Dec2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_dec2

	return p
}

func (s *Dec2Context) GetParser() antlr.Parser { return s.parser }

func (s *Dec2Context) AllDigit() []antlr.TerminalNode {
	return s.GetTokens(iso8601ParserDigit)
}

func (s *Dec2Context) Digit(i int) antlr.TerminalNode {
	return s.GetToken(iso8601ParserDigit, i)
}

func (s *Dec2Context) Fraction() antlr.TerminalNode {
	return s.GetToken(iso8601ParserFraction, 0)
}

func (s *Dec2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Dec2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDec2(s)
	}
}

func (s *Dec2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDec2(s)
	}
}

func (p *iso8601Parser) Dec2() (localctx IDec2Context) {
	this := p
	_ = this

	localctx = NewDec2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, iso8601ParserRULE_dec2)
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
		p.SetState(229)
		p.Match(iso8601ParserDigit)
	}
	{
		p.SetState(230)
		p.Match(iso8601ParserDigit)
	}
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == iso8601ParserFraction {
		{
			p.SetState(231)
			p.Match(iso8601ParserFraction)
		}

	}

	return localctx
}

// ICenturyContext is an interface to support dynamic dispatch.
type ICenturyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCenturyContext differentiates from other interfaces.
	IsCenturyContext()
}

type CenturyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCenturyContext() *CenturyContext {
	var p = new(CenturyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_century
	return p
}

func (*CenturyContext) IsCenturyContext() {}

func NewCenturyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CenturyContext {
	var p = new(CenturyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_century

	return p
}

func (s *CenturyContext) GetParser() antlr.Parser { return s.parser }

func (s *CenturyContext) CopyFrom(ctx *CenturyContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CenturyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CenturyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompleteCenturyContext struct {
	*CenturyContext
}

func NewCompleteCenturyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompleteCenturyContext {
	var p = new(CompleteCenturyContext)

	p.CenturyContext = NewEmptyCenturyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CenturyContext))

	return p
}

func (s *CompleteCenturyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompleteCenturyContext) Int2() IInt2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInt2Context)
}

func (s *CompleteCenturyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterCompleteCentury(s)
	}
}

func (s *CompleteCenturyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitCompleteCentury(s)
	}
}

type ExpandedCenturyContext struct {
	*CenturyContext
}

func NewExpandedCenturyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpandedCenturyContext {
	var p = new(ExpandedCenturyContext)

	p.CenturyContext = NewEmptyCenturyContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CenturyContext))

	return p
}

func (s *ExpandedCenturyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpandedCenturyContext) Sint2p() ISint2pContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISint2pContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISint2pContext)
}

func (s *ExpandedCenturyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterExpandedCentury(s)
	}
}

func (s *ExpandedCenturyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitExpandedCentury(s)
	}
}

func (p *iso8601Parser) Century() (localctx ICenturyContext) {
	this := p
	_ = this

	localctx = NewCenturyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, iso8601ParserRULE_century)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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

	switch p.GetTokenStream().LA(1) {
	case iso8601ParserDigit:
		localctx = NewCompleteCenturyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Int2()
		}

	case iso8601ParserT__0, iso8601ParserT__1:
		localctx = NewExpandedCenturyContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.Sint2p()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IYearContext is an interface to support dynamic dispatch.
type IYearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsYearContext differentiates from other interfaces.
	IsYearContext()
}

type YearContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyYearContext() *YearContext {
	var p = new(YearContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_year
	return p
}

func (*YearContext) IsYearContext() {}

func NewYearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *YearContext {
	var p = new(YearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_year

	return p
}

func (s *YearContext) GetParser() antlr.Parser { return s.parser }

func (s *YearContext) CopyFrom(ctx *YearContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *YearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *YearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CompleteYearContext struct {
	*YearContext
}

func NewCompleteYearContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompleteYearContext {
	var p = new(CompleteYearContext)

	p.YearContext = NewEmptyYearContext()
	p.parser = parser
	p.CopyFrom(ctx.(*YearContext))

	return p
}

func (s *CompleteYearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompleteYearContext) Int4() IInt4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInt4Context)
}

func (s *CompleteYearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterCompleteYear(s)
	}
}

func (s *CompleteYearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitCompleteYear(s)
	}
}

type ExpandedYearContext struct {
	*YearContext
}

func NewExpandedYearContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpandedYearContext {
	var p = new(ExpandedYearContext)

	p.YearContext = NewEmptyYearContext()
	p.parser = parser
	p.CopyFrom(ctx.(*YearContext))

	return p
}

func (s *ExpandedYearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpandedYearContext) Sint4p() ISint4pContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISint4pContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISint4pContext)
}

func (s *ExpandedYearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterExpandedYear(s)
	}
}

func (s *ExpandedYearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitExpandedYear(s)
	}
}

func (p *iso8601Parser) Year() (localctx IYearContext) {
	this := p
	_ = this

	localctx = NewYearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, iso8601ParserRULE_year)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(240)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case iso8601ParserDigit:
		localctx = NewCompleteYearContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.Int4()
		}

	case iso8601ParserT__0, iso8601ParserT__1:
		localctx = NewExpandedYearContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
			p.Sint4p()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMonthContext is an interface to support dynamic dispatch.
type IMonthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonthContext differentiates from other interfaces.
	IsMonthContext()
}

type MonthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonthContext() *MonthContext {
	var p = new(MonthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_month
	return p
}

func (*MonthContext) IsMonthContext() {}

func NewMonthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonthContext {
	var p = new(MonthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_month

	return p
}

func (s *MonthContext) GetParser() antlr.Parser { return s.parser }

func (s *MonthContext) Int2() IInt2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInt2Context)
}

func (s *MonthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterMonth(s)
	}
}

func (s *MonthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitMonth(s)
	}
}

func (p *iso8601Parser) Month() (localctx IMonthContext) {
	this := p
	_ = this

	localctx = NewMonthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, iso8601ParserRULE_month)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(242)
		p.Int2()
	}

	return localctx
}

// IDayContext is an interface to support dynamic dispatch.
type IDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDayContext differentiates from other interfaces.
	IsDayContext()
}

type DayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDayContext() *DayContext {
	var p = new(DayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_day
	return p
}

func (*DayContext) IsDayContext() {}

func NewDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DayContext {
	var p = new(DayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_day

	return p
}

func (s *DayContext) GetParser() antlr.Parser { return s.parser }

func (s *DayContext) Int2() IInt2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInt2Context)
}

func (s *DayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDay(s)
	}
}

func (s *DayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDay(s)
	}
}

func (p *iso8601Parser) Day() (localctx IDayContext) {
	this := p
	_ = this

	localctx = NewDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, iso8601ParserRULE_day)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Int2()
	}

	return localctx
}

// IOrdinalDayContext is an interface to support dynamic dispatch.
type IOrdinalDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrdinalDayContext differentiates from other interfaces.
	IsOrdinalDayContext()
}

type OrdinalDayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrdinalDayContext() *OrdinalDayContext {
	var p = new(OrdinalDayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_ordinalDay
	return p
}

func (*OrdinalDayContext) IsOrdinalDayContext() {}

func NewOrdinalDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrdinalDayContext {
	var p = new(OrdinalDayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_ordinalDay

	return p
}

func (s *OrdinalDayContext) GetParser() antlr.Parser { return s.parser }

func (s *OrdinalDayContext) Int3() IInt3Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt3Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInt3Context)
}

func (s *OrdinalDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrdinalDayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrdinalDayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterOrdinalDay(s)
	}
}

func (s *OrdinalDayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitOrdinalDay(s)
	}
}

func (p *iso8601Parser) OrdinalDay() (localctx IOrdinalDayContext) {
	this := p
	_ = this

	localctx = NewOrdinalDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, iso8601ParserRULE_ordinalDay)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(246)
		p.Int3()
	}

	return localctx
}

// IWeekContext is an interface to support dynamic dispatch.
type IWeekContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWeekContext differentiates from other interfaces.
	IsWeekContext()
}

type WeekContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWeekContext() *WeekContext {
	var p = new(WeekContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_week
	return p
}

func (*WeekContext) IsWeekContext() {}

func NewWeekContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WeekContext {
	var p = new(WeekContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_week

	return p
}

func (s *WeekContext) GetParser() antlr.Parser { return s.parser }

func (s *WeekContext) Int2() IInt2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInt2Context)
}

func (s *WeekContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WeekContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WeekContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterWeek(s)
	}
}

func (s *WeekContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitWeek(s)
	}
}

func (p *iso8601Parser) Week() (localctx IWeekContext) {
	this := p
	_ = this

	localctx = NewWeekContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, iso8601ParserRULE_week)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Int2()
	}

	return localctx
}

// IWeekDayContext is an interface to support dynamic dispatch.
type IWeekDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWeekDayContext differentiates from other interfaces.
	IsWeekDayContext()
}

type WeekDayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWeekDayContext() *WeekDayContext {
	var p = new(WeekDayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_weekDay
	return p
}

func (*WeekDayContext) IsWeekDayContext() {}

func NewWeekDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WeekDayContext {
	var p = new(WeekDayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_weekDay

	return p
}

func (s *WeekDayContext) GetParser() antlr.Parser { return s.parser }

func (s *WeekDayContext) Digit() antlr.TerminalNode {
	return s.GetToken(iso8601ParserDigit, 0)
}

func (s *WeekDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WeekDayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WeekDayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterWeekDay(s)
	}
}

func (s *WeekDayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitWeekDay(s)
	}
}

func (p *iso8601Parser) WeekDay() (localctx IWeekDayContext) {
	this := p
	_ = this

	localctx = NewWeekDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, iso8601ParserRULE_weekDay)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(250)
		p.Match(iso8601ParserDigit)
	}

	return localctx
}

// IHourContext is an interface to support dynamic dispatch.
type IHourContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHourContext differentiates from other interfaces.
	IsHourContext()
}

type HourContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHourContext() *HourContext {
	var p = new(HourContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_hour
	return p
}

func (*HourContext) IsHourContext() {}

func NewHourContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HourContext {
	var p = new(HourContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_hour

	return p
}

func (s *HourContext) GetParser() antlr.Parser { return s.parser }

func (s *HourContext) Int2() IInt2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInt2Context)
}

func (s *HourContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HourContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HourContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterHour(s)
	}
}

func (s *HourContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitHour(s)
	}
}

func (p *iso8601Parser) Hour() (localctx IHourContext) {
	this := p
	_ = this

	localctx = NewHourContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, iso8601ParserRULE_hour)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(252)
		p.Int2()
	}

	return localctx
}

// IMinuteContext is an interface to support dynamic dispatch.
type IMinuteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMinuteContext differentiates from other interfaces.
	IsMinuteContext()
}

type MinuteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMinuteContext() *MinuteContext {
	var p = new(MinuteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_minute
	return p
}

func (*MinuteContext) IsMinuteContext() {}

func NewMinuteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MinuteContext {
	var p = new(MinuteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_minute

	return p
}

func (s *MinuteContext) GetParser() antlr.Parser { return s.parser }

func (s *MinuteContext) Int2() IInt2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInt2Context)
}

func (s *MinuteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinuteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MinuteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterMinute(s)
	}
}

func (s *MinuteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitMinute(s)
	}
}

func (p *iso8601Parser) Minute() (localctx IMinuteContext) {
	this := p
	_ = this

	localctx = NewMinuteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, iso8601ParserRULE_minute)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(254)
		p.Int2()
	}

	return localctx
}

// ISecondContext is an interface to support dynamic dispatch.
type ISecondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSecondContext differentiates from other interfaces.
	IsSecondContext()
}

type SecondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySecondContext() *SecondContext {
	var p = new(SecondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_second
	return p
}

func (*SecondContext) IsSecondContext() {}

func NewSecondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SecondContext {
	var p = new(SecondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_second

	return p
}

func (s *SecondContext) GetParser() antlr.Parser { return s.parser }

func (s *SecondContext) Int2() IInt2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInt2Context)
}

func (s *SecondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SecondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSecond(s)
	}
}

func (s *SecondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSecond(s)
	}
}

func (p *iso8601Parser) Second() (localctx ISecondContext) {
	this := p
	_ = this

	localctx = NewSecondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, iso8601ParserRULE_second)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(256)
		p.Int2()
	}

	return localctx
}

// IHourFractionContext is an interface to support dynamic dispatch.
type IHourFractionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHourFractionContext differentiates from other interfaces.
	IsHourFractionContext()
}

type HourFractionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHourFractionContext() *HourFractionContext {
	var p = new(HourFractionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_hourFraction
	return p
}

func (*HourFractionContext) IsHourFractionContext() {}

func NewHourFractionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HourFractionContext {
	var p = new(HourFractionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_hourFraction

	return p
}

func (s *HourFractionContext) GetParser() antlr.Parser { return s.parser }

func (s *HourFractionContext) Dec2() IDec2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDec2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDec2Context)
}

func (s *HourFractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HourFractionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HourFractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterHourFraction(s)
	}
}

func (s *HourFractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitHourFraction(s)
	}
}

func (p *iso8601Parser) HourFraction() (localctx IHourFractionContext) {
	this := p
	_ = this

	localctx = NewHourFractionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, iso8601ParserRULE_hourFraction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(258)
		p.Dec2()
	}

	return localctx
}

// IMinuteFractionContext is an interface to support dynamic dispatch.
type IMinuteFractionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMinuteFractionContext differentiates from other interfaces.
	IsMinuteFractionContext()
}

type MinuteFractionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMinuteFractionContext() *MinuteFractionContext {
	var p = new(MinuteFractionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_minuteFraction
	return p
}

func (*MinuteFractionContext) IsMinuteFractionContext() {}

func NewMinuteFractionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MinuteFractionContext {
	var p = new(MinuteFractionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_minuteFraction

	return p
}

func (s *MinuteFractionContext) GetParser() antlr.Parser { return s.parser }

func (s *MinuteFractionContext) Dec2() IDec2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDec2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDec2Context)
}

func (s *MinuteFractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinuteFractionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MinuteFractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterMinuteFraction(s)
	}
}

func (s *MinuteFractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitMinuteFraction(s)
	}
}

func (p *iso8601Parser) MinuteFraction() (localctx IMinuteFractionContext) {
	this := p
	_ = this

	localctx = NewMinuteFractionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, iso8601ParserRULE_minuteFraction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(260)
		p.Dec2()
	}

	return localctx
}

// ISecondFractionContext is an interface to support dynamic dispatch.
type ISecondFractionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSecondFractionContext differentiates from other interfaces.
	IsSecondFractionContext()
}

type SecondFractionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySecondFractionContext() *SecondFractionContext {
	var p = new(SecondFractionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_secondFraction
	return p
}

func (*SecondFractionContext) IsSecondFractionContext() {}

func NewSecondFractionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SecondFractionContext {
	var p = new(SecondFractionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_secondFraction

	return p
}

func (s *SecondFractionContext) GetParser() antlr.Parser { return s.parser }

func (s *SecondFractionContext) Dec2() IDec2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDec2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDec2Context)
}

func (s *SecondFractionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecondFractionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SecondFractionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSecondFraction(s)
	}
}

func (s *SecondFractionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSecondFraction(s)
	}
}

func (p *iso8601Parser) SecondFraction() (localctx ISecondFractionContext) {
	this := p
	_ = this

	localctx = NewSecondFractionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, iso8601ParserRULE_secondFraction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(262)
		p.Dec2()
	}

	return localctx
}

// ICalendarDateContext is an interface to support dynamic dispatch.
type ICalendarDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalendarDateContext differentiates from other interfaces.
	IsCalendarDateContext()
}

type CalendarDateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalendarDateContext() *CalendarDateContext {
	var p = new(CalendarDateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_calendarDate
	return p
}

func (*CalendarDateContext) IsCalendarDateContext() {}

func NewCalendarDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalendarDateContext {
	var p = new(CalendarDateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_calendarDate

	return p
}

func (s *CalendarDateContext) GetParser() antlr.Parser { return s.parser }

func (s *CalendarDateContext) CalendarDateBasic() ICalendarDateBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalendarDateBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalendarDateBasicContext)
}

func (s *CalendarDateContext) CalendarDateExtended() ICalendarDateExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalendarDateExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalendarDateExtendedContext)
}

func (s *CalendarDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalendarDateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalendarDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterCalendarDate(s)
	}
}

func (s *CalendarDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitCalendarDate(s)
	}
}

func (p *iso8601Parser) CalendarDate() (localctx ICalendarDateContext) {
	this := p
	_ = this

	localctx = NewCalendarDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, iso8601ParserRULE_calendarDate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(266)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.CalendarDateBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)
			p.CalendarDateExtended()
		}

	}

	return localctx
}

// ICalendarDateBasicContext is an interface to support dynamic dispatch.
type ICalendarDateBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalendarDateBasicContext differentiates from other interfaces.
	IsCalendarDateBasicContext()
}

type CalendarDateBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalendarDateBasicContext() *CalendarDateBasicContext {
	var p = new(CalendarDateBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_calendarDateBasic
	return p
}

func (*CalendarDateBasicContext) IsCalendarDateBasicContext() {}

func NewCalendarDateBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalendarDateBasicContext {
	var p = new(CalendarDateBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_calendarDateBasic

	return p
}

func (s *CalendarDateBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *CalendarDateBasicContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *CalendarDateBasicContext) Month() IMonthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthContext)
}

func (s *CalendarDateBasicContext) Day() IDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *CalendarDateBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalendarDateBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalendarDateBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterCalendarDateBasic(s)
	}
}

func (s *CalendarDateBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitCalendarDateBasic(s)
	}
}

func (p *iso8601Parser) CalendarDateBasic() (localctx ICalendarDateBasicContext) {
	this := p
	_ = this

	localctx = NewCalendarDateBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, iso8601ParserRULE_calendarDateBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(268)
		p.Year()
	}
	{
		p.SetState(269)
		p.Month()
	}
	{
		p.SetState(270)
		p.Day()
	}

	return localctx
}

// ICalendarDateExtendedContext is an interface to support dynamic dispatch.
type ICalendarDateExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCalendarDateExtendedContext differentiates from other interfaces.
	IsCalendarDateExtendedContext()
}

type CalendarDateExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalendarDateExtendedContext() *CalendarDateExtendedContext {
	var p = new(CalendarDateExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_calendarDateExtended
	return p
}

func (*CalendarDateExtendedContext) IsCalendarDateExtendedContext() {}

func NewCalendarDateExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalendarDateExtendedContext {
	var p = new(CalendarDateExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_calendarDateExtended

	return p
}

func (s *CalendarDateExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *CalendarDateExtendedContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *CalendarDateExtendedContext) Month() IMonthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthContext)
}

func (s *CalendarDateExtendedContext) Day() IDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *CalendarDateExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalendarDateExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalendarDateExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterCalendarDateExtended(s)
	}
}

func (s *CalendarDateExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitCalendarDateExtended(s)
	}
}

func (p *iso8601Parser) CalendarDateExtended() (localctx ICalendarDateExtendedContext) {
	this := p
	_ = this

	localctx = NewCalendarDateExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, iso8601ParserRULE_calendarDateExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(272)
		p.Year()
	}
	{
		p.SetState(273)
		p.Match(iso8601ParserT__1)
	}
	{
		p.SetState(274)
		p.Month()
	}
	{
		p.SetState(275)
		p.Match(iso8601ParserT__1)
	}
	{
		p.SetState(276)
		p.Day()
	}

	return localctx
}

// ISpecificMonthCalendarDateContext is an interface to support dynamic dispatch.
type ISpecificMonthCalendarDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecificMonthCalendarDateContext differentiates from other interfaces.
	IsSpecificMonthCalendarDateContext()
}

type SpecificMonthCalendarDateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecificMonthCalendarDateContext() *SpecificMonthCalendarDateContext {
	var p = new(SpecificMonthCalendarDateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_specificMonthCalendarDate
	return p
}

func (*SpecificMonthCalendarDateContext) IsSpecificMonthCalendarDateContext() {}

func NewSpecificMonthCalendarDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificMonthCalendarDateContext {
	var p = new(SpecificMonthCalendarDateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_specificMonthCalendarDate

	return p
}

func (s *SpecificMonthCalendarDateContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificMonthCalendarDateContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *SpecificMonthCalendarDateContext) Month() IMonthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthContext)
}

func (s *SpecificMonthCalendarDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificMonthCalendarDateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecificMonthCalendarDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSpecificMonthCalendarDate(s)
	}
}

func (s *SpecificMonthCalendarDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSpecificMonthCalendarDate(s)
	}
}

func (p *iso8601Parser) SpecificMonthCalendarDate() (localctx ISpecificMonthCalendarDateContext) {
	this := p
	_ = this

	localctx = NewSpecificMonthCalendarDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, iso8601ParserRULE_specificMonthCalendarDate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(278)
		p.Year()
	}
	{
		p.SetState(279)
		p.Match(iso8601ParserT__1)
	}
	{
		p.SetState(280)
		p.Month()
	}

	return localctx
}

// ISpecificYearCalendarDateContext is an interface to support dynamic dispatch.
type ISpecificYearCalendarDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecificYearCalendarDateContext differentiates from other interfaces.
	IsSpecificYearCalendarDateContext()
}

type SpecificYearCalendarDateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecificYearCalendarDateContext() *SpecificYearCalendarDateContext {
	var p = new(SpecificYearCalendarDateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_specificYearCalendarDate
	return p
}

func (*SpecificYearCalendarDateContext) IsSpecificYearCalendarDateContext() {}

func NewSpecificYearCalendarDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificYearCalendarDateContext {
	var p = new(SpecificYearCalendarDateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_specificYearCalendarDate

	return p
}

func (s *SpecificYearCalendarDateContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificYearCalendarDateContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *SpecificYearCalendarDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificYearCalendarDateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecificYearCalendarDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSpecificYearCalendarDate(s)
	}
}

func (s *SpecificYearCalendarDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSpecificYearCalendarDate(s)
	}
}

func (p *iso8601Parser) SpecificYearCalendarDate() (localctx ISpecificYearCalendarDateContext) {
	this := p
	_ = this

	localctx = NewSpecificYearCalendarDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, iso8601ParserRULE_specificYearCalendarDate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(282)
		p.Year()
	}

	return localctx
}

// ISpecificCenturyCalendarDateContext is an interface to support dynamic dispatch.
type ISpecificCenturyCalendarDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecificCenturyCalendarDateContext differentiates from other interfaces.
	IsSpecificCenturyCalendarDateContext()
}

type SpecificCenturyCalendarDateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecificCenturyCalendarDateContext() *SpecificCenturyCalendarDateContext {
	var p = new(SpecificCenturyCalendarDateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_specificCenturyCalendarDate
	return p
}

func (*SpecificCenturyCalendarDateContext) IsSpecificCenturyCalendarDateContext() {}

func NewSpecificCenturyCalendarDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificCenturyCalendarDateContext {
	var p = new(SpecificCenturyCalendarDateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_specificCenturyCalendarDate

	return p
}

func (s *SpecificCenturyCalendarDateContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificCenturyCalendarDateContext) Century() ICenturyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICenturyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICenturyContext)
}

func (s *SpecificCenturyCalendarDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificCenturyCalendarDateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecificCenturyCalendarDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSpecificCenturyCalendarDate(s)
	}
}

func (s *SpecificCenturyCalendarDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSpecificCenturyCalendarDate(s)
	}
}

func (p *iso8601Parser) SpecificCenturyCalendarDate() (localctx ISpecificCenturyCalendarDateContext) {
	this := p
	_ = this

	localctx = NewSpecificCenturyCalendarDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, iso8601ParserRULE_specificCenturyCalendarDate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(284)
		p.Century()
	}

	return localctx
}

// IOrdinalDateContext is an interface to support dynamic dispatch.
type IOrdinalDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrdinalDateContext differentiates from other interfaces.
	IsOrdinalDateContext()
}

type OrdinalDateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrdinalDateContext() *OrdinalDateContext {
	var p = new(OrdinalDateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_ordinalDate
	return p
}

func (*OrdinalDateContext) IsOrdinalDateContext() {}

func NewOrdinalDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrdinalDateContext {
	var p = new(OrdinalDateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_ordinalDate

	return p
}

func (s *OrdinalDateContext) GetParser() antlr.Parser { return s.parser }

func (s *OrdinalDateContext) OrdinalDateBasic() IOrdinalDateBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrdinalDateBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrdinalDateBasicContext)
}

func (s *OrdinalDateContext) OrdinalDateExtended() IOrdinalDateExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrdinalDateExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrdinalDateExtendedContext)
}

func (s *OrdinalDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrdinalDateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrdinalDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterOrdinalDate(s)
	}
}

func (s *OrdinalDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitOrdinalDate(s)
	}
}

func (p *iso8601Parser) OrdinalDate() (localctx IOrdinalDateContext) {
	this := p
	_ = this

	localctx = NewOrdinalDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, iso8601ParserRULE_ordinalDate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(286)
			p.OrdinalDateBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.OrdinalDateExtended()
		}

	}

	return localctx
}

// IOrdinalDateBasicContext is an interface to support dynamic dispatch.
type IOrdinalDateBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrdinalDateBasicContext differentiates from other interfaces.
	IsOrdinalDateBasicContext()
}

type OrdinalDateBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrdinalDateBasicContext() *OrdinalDateBasicContext {
	var p = new(OrdinalDateBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_ordinalDateBasic
	return p
}

func (*OrdinalDateBasicContext) IsOrdinalDateBasicContext() {}

func NewOrdinalDateBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrdinalDateBasicContext {
	var p = new(OrdinalDateBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_ordinalDateBasic

	return p
}

func (s *OrdinalDateBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *OrdinalDateBasicContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *OrdinalDateBasicContext) OrdinalDay() IOrdinalDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrdinalDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrdinalDayContext)
}

func (s *OrdinalDateBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrdinalDateBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrdinalDateBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterOrdinalDateBasic(s)
	}
}

func (s *OrdinalDateBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitOrdinalDateBasic(s)
	}
}

func (p *iso8601Parser) OrdinalDateBasic() (localctx IOrdinalDateBasicContext) {
	this := p
	_ = this

	localctx = NewOrdinalDateBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, iso8601ParserRULE_ordinalDateBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(290)
		p.Year()
	}
	{
		p.SetState(291)
		p.OrdinalDay()
	}

	return localctx
}

// IOrdinalDateExtendedContext is an interface to support dynamic dispatch.
type IOrdinalDateExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrdinalDateExtendedContext differentiates from other interfaces.
	IsOrdinalDateExtendedContext()
}

type OrdinalDateExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrdinalDateExtendedContext() *OrdinalDateExtendedContext {
	var p = new(OrdinalDateExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_ordinalDateExtended
	return p
}

func (*OrdinalDateExtendedContext) IsOrdinalDateExtendedContext() {}

func NewOrdinalDateExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrdinalDateExtendedContext {
	var p = new(OrdinalDateExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_ordinalDateExtended

	return p
}

func (s *OrdinalDateExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *OrdinalDateExtendedContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *OrdinalDateExtendedContext) OrdinalDay() IOrdinalDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrdinalDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrdinalDayContext)
}

func (s *OrdinalDateExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrdinalDateExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrdinalDateExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterOrdinalDateExtended(s)
	}
}

func (s *OrdinalDateExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitOrdinalDateExtended(s)
	}
}

func (p *iso8601Parser) OrdinalDateExtended() (localctx IOrdinalDateExtendedContext) {
	this := p
	_ = this

	localctx = NewOrdinalDateExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, iso8601ParserRULE_ordinalDateExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Year()
	}
	{
		p.SetState(294)
		p.Match(iso8601ParserT__1)
	}
	{
		p.SetState(295)
		p.OrdinalDay()
	}

	return localctx
}

// IWeekDateContext is an interface to support dynamic dispatch.
type IWeekDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWeekDateContext differentiates from other interfaces.
	IsWeekDateContext()
}

type WeekDateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWeekDateContext() *WeekDateContext {
	var p = new(WeekDateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_weekDate
	return p
}

func (*WeekDateContext) IsWeekDateContext() {}

func NewWeekDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WeekDateContext {
	var p = new(WeekDateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_weekDate

	return p
}

func (s *WeekDateContext) GetParser() antlr.Parser { return s.parser }

func (s *WeekDateContext) WeekDateBasic() IWeekDateBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWeekDateBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWeekDateBasicContext)
}

func (s *WeekDateContext) WeekDateExtended() IWeekDateExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWeekDateExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWeekDateExtendedContext)
}

func (s *WeekDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WeekDateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WeekDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterWeekDate(s)
	}
}

func (s *WeekDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitWeekDate(s)
	}
}

func (p *iso8601Parser) WeekDate() (localctx IWeekDateContext) {
	this := p
	_ = this

	localctx = NewWeekDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, iso8601ParserRULE_weekDate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.WeekDateBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(298)
			p.WeekDateExtended()
		}

	}

	return localctx
}

// IWeekDateBasicContext is an interface to support dynamic dispatch.
type IWeekDateBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWeekDateBasicContext differentiates from other interfaces.
	IsWeekDateBasicContext()
}

type WeekDateBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWeekDateBasicContext() *WeekDateBasicContext {
	var p = new(WeekDateBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_weekDateBasic
	return p
}

func (*WeekDateBasicContext) IsWeekDateBasicContext() {}

func NewWeekDateBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WeekDateBasicContext {
	var p = new(WeekDateBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_weekDateBasic

	return p
}

func (s *WeekDateBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *WeekDateBasicContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *WeekDateBasicContext) W() antlr.TerminalNode {
	return s.GetToken(iso8601ParserW, 0)
}

func (s *WeekDateBasicContext) Week() IWeekContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWeekContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWeekContext)
}

func (s *WeekDateBasicContext) WeekDay() IWeekDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWeekDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWeekDayContext)
}

func (s *WeekDateBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WeekDateBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WeekDateBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterWeekDateBasic(s)
	}
}

func (s *WeekDateBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitWeekDateBasic(s)
	}
}

func (p *iso8601Parser) WeekDateBasic() (localctx IWeekDateBasicContext) {
	this := p
	_ = this

	localctx = NewWeekDateBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, iso8601ParserRULE_weekDateBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(301)
		p.Year()
	}
	{
		p.SetState(302)
		p.Match(iso8601ParserW)
	}
	{
		p.SetState(303)
		p.Week()
	}
	{
		p.SetState(304)
		p.WeekDay()
	}

	return localctx
}

// IWeekDateExtendedContext is an interface to support dynamic dispatch.
type IWeekDateExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWeekDateExtendedContext differentiates from other interfaces.
	IsWeekDateExtendedContext()
}

type WeekDateExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWeekDateExtendedContext() *WeekDateExtendedContext {
	var p = new(WeekDateExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_weekDateExtended
	return p
}

func (*WeekDateExtendedContext) IsWeekDateExtendedContext() {}

func NewWeekDateExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WeekDateExtendedContext {
	var p = new(WeekDateExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_weekDateExtended

	return p
}

func (s *WeekDateExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *WeekDateExtendedContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *WeekDateExtendedContext) W() antlr.TerminalNode {
	return s.GetToken(iso8601ParserW, 0)
}

func (s *WeekDateExtendedContext) Week() IWeekContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWeekContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWeekContext)
}

func (s *WeekDateExtendedContext) WeekDay() IWeekDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWeekDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWeekDayContext)
}

func (s *WeekDateExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WeekDateExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WeekDateExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterWeekDateExtended(s)
	}
}

func (s *WeekDateExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitWeekDateExtended(s)
	}
}

func (p *iso8601Parser) WeekDateExtended() (localctx IWeekDateExtendedContext) {
	this := p
	_ = this

	localctx = NewWeekDateExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, iso8601ParserRULE_weekDateExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Year()
	}
	{
		p.SetState(307)
		p.Match(iso8601ParserT__1)
	}
	{
		p.SetState(308)
		p.Match(iso8601ParserW)
	}
	{
		p.SetState(309)
		p.Week()
	}
	{
		p.SetState(310)
		p.Match(iso8601ParserT__1)
	}
	{
		p.SetState(311)
		p.WeekDay()
	}

	return localctx
}

// ISpecificWeekWeekDateContext is an interface to support dynamic dispatch.
type ISpecificWeekWeekDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecificWeekWeekDateContext differentiates from other interfaces.
	IsSpecificWeekWeekDateContext()
}

type SpecificWeekWeekDateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecificWeekWeekDateContext() *SpecificWeekWeekDateContext {
	var p = new(SpecificWeekWeekDateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_specificWeekWeekDate
	return p
}

func (*SpecificWeekWeekDateContext) IsSpecificWeekWeekDateContext() {}

func NewSpecificWeekWeekDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificWeekWeekDateContext {
	var p = new(SpecificWeekWeekDateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_specificWeekWeekDate

	return p
}

func (s *SpecificWeekWeekDateContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificWeekWeekDateContext) SpecificWeekWeekDateBasic() ISpecificWeekWeekDateBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificWeekWeekDateBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificWeekWeekDateBasicContext)
}

func (s *SpecificWeekWeekDateContext) SpecificWeekWeekDateExtended() ISpecificWeekWeekDateExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificWeekWeekDateExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificWeekWeekDateExtendedContext)
}

func (s *SpecificWeekWeekDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificWeekWeekDateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecificWeekWeekDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSpecificWeekWeekDate(s)
	}
}

func (s *SpecificWeekWeekDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSpecificWeekWeekDate(s)
	}
}

func (p *iso8601Parser) SpecificWeekWeekDate() (localctx ISpecificWeekWeekDateContext) {
	this := p
	_ = this

	localctx = NewSpecificWeekWeekDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, iso8601ParserRULE_specificWeekWeekDate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(313)
			p.SpecificWeekWeekDateBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(314)
			p.SpecificWeekWeekDateExtended()
		}

	}

	return localctx
}

// ISpecificWeekWeekDateBasicContext is an interface to support dynamic dispatch.
type ISpecificWeekWeekDateBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecificWeekWeekDateBasicContext differentiates from other interfaces.
	IsSpecificWeekWeekDateBasicContext()
}

type SpecificWeekWeekDateBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecificWeekWeekDateBasicContext() *SpecificWeekWeekDateBasicContext {
	var p = new(SpecificWeekWeekDateBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_specificWeekWeekDateBasic
	return p
}

func (*SpecificWeekWeekDateBasicContext) IsSpecificWeekWeekDateBasicContext() {}

func NewSpecificWeekWeekDateBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificWeekWeekDateBasicContext {
	var p = new(SpecificWeekWeekDateBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_specificWeekWeekDateBasic

	return p
}

func (s *SpecificWeekWeekDateBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificWeekWeekDateBasicContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *SpecificWeekWeekDateBasicContext) W() antlr.TerminalNode {
	return s.GetToken(iso8601ParserW, 0)
}

func (s *SpecificWeekWeekDateBasicContext) Week() IWeekContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWeekContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWeekContext)
}

func (s *SpecificWeekWeekDateBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificWeekWeekDateBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecificWeekWeekDateBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSpecificWeekWeekDateBasic(s)
	}
}

func (s *SpecificWeekWeekDateBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSpecificWeekWeekDateBasic(s)
	}
}

func (p *iso8601Parser) SpecificWeekWeekDateBasic() (localctx ISpecificWeekWeekDateBasicContext) {
	this := p
	_ = this

	localctx = NewSpecificWeekWeekDateBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, iso8601ParserRULE_specificWeekWeekDateBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(317)
		p.Year()
	}
	{
		p.SetState(318)
		p.Match(iso8601ParserW)
	}
	{
		p.SetState(319)
		p.Week()
	}

	return localctx
}

// ISpecificWeekWeekDateExtendedContext is an interface to support dynamic dispatch.
type ISpecificWeekWeekDateExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecificWeekWeekDateExtendedContext differentiates from other interfaces.
	IsSpecificWeekWeekDateExtendedContext()
}

type SpecificWeekWeekDateExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecificWeekWeekDateExtendedContext() *SpecificWeekWeekDateExtendedContext {
	var p = new(SpecificWeekWeekDateExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_specificWeekWeekDateExtended
	return p
}

func (*SpecificWeekWeekDateExtendedContext) IsSpecificWeekWeekDateExtendedContext() {}

func NewSpecificWeekWeekDateExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificWeekWeekDateExtendedContext {
	var p = new(SpecificWeekWeekDateExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_specificWeekWeekDateExtended

	return p
}

func (s *SpecificWeekWeekDateExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificWeekWeekDateExtendedContext) Year() IYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IYearContext)
}

func (s *SpecificWeekWeekDateExtendedContext) W() antlr.TerminalNode {
	return s.GetToken(iso8601ParserW, 0)
}

func (s *SpecificWeekWeekDateExtendedContext) Week() IWeekContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWeekContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWeekContext)
}

func (s *SpecificWeekWeekDateExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificWeekWeekDateExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecificWeekWeekDateExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSpecificWeekWeekDateExtended(s)
	}
}

func (s *SpecificWeekWeekDateExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSpecificWeekWeekDateExtended(s)
	}
}

func (p *iso8601Parser) SpecificWeekWeekDateExtended() (localctx ISpecificWeekWeekDateExtendedContext) {
	this := p
	_ = this

	localctx = NewSpecificWeekWeekDateExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, iso8601ParserRULE_specificWeekWeekDateExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(321)
		p.Year()
	}
	{
		p.SetState(322)
		p.Match(iso8601ParserT__1)
	}
	{
		p.SetState(323)
		p.Match(iso8601ParserW)
	}
	{
		p.SetState(324)
		p.Week()
	}

	return localctx
}

// IDatePreciseContext is an interface to support dynamic dispatch.
type IDatePreciseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatePreciseContext differentiates from other interfaces.
	IsDatePreciseContext()
}

type DatePreciseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatePreciseContext() *DatePreciseContext {
	var p = new(DatePreciseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_datePrecise
	return p
}

func (*DatePreciseContext) IsDatePreciseContext() {}

func NewDatePreciseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatePreciseContext {
	var p = new(DatePreciseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_datePrecise

	return p
}

func (s *DatePreciseContext) GetParser() antlr.Parser { return s.parser }

func (s *DatePreciseContext) DatePreciseBasic() IDatePreciseBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatePreciseBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatePreciseBasicContext)
}

func (s *DatePreciseContext) DatePreciseExtended() IDatePreciseExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatePreciseExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatePreciseExtendedContext)
}

func (s *DatePreciseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatePreciseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatePreciseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDatePrecise(s)
	}
}

func (s *DatePreciseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDatePrecise(s)
	}
}

func (p *iso8601Parser) DatePrecise() (localctx IDatePreciseContext) {
	this := p
	_ = this

	localctx = NewDatePreciseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, iso8601ParserRULE_datePrecise)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)
			p.DatePreciseBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(327)
			p.DatePreciseExtended()
		}

	}

	return localctx
}

// IDatePreciseBasicContext is an interface to support dynamic dispatch.
type IDatePreciseBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatePreciseBasicContext differentiates from other interfaces.
	IsDatePreciseBasicContext()
}

type DatePreciseBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatePreciseBasicContext() *DatePreciseBasicContext {
	var p = new(DatePreciseBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_datePreciseBasic
	return p
}

func (*DatePreciseBasicContext) IsDatePreciseBasicContext() {}

func NewDatePreciseBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatePreciseBasicContext {
	var p = new(DatePreciseBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_datePreciseBasic

	return p
}

func (s *DatePreciseBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *DatePreciseBasicContext) CalendarDateBasic() ICalendarDateBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalendarDateBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalendarDateBasicContext)
}

func (s *DatePreciseBasicContext) OrdinalDateBasic() IOrdinalDateBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrdinalDateBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrdinalDateBasicContext)
}

func (s *DatePreciseBasicContext) WeekDateBasic() IWeekDateBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWeekDateBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWeekDateBasicContext)
}

func (s *DatePreciseBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatePreciseBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatePreciseBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDatePreciseBasic(s)
	}
}

func (s *DatePreciseBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDatePreciseBasic(s)
	}
}

func (p *iso8601Parser) DatePreciseBasic() (localctx IDatePreciseBasicContext) {
	this := p
	_ = this

	localctx = NewDatePreciseBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, iso8601ParserRULE_datePreciseBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(333)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(330)
			p.CalendarDateBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)
			p.OrdinalDateBasic()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(332)
			p.WeekDateBasic()
		}

	}

	return localctx
}

// IDatePreciseExtendedContext is an interface to support dynamic dispatch.
type IDatePreciseExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatePreciseExtendedContext differentiates from other interfaces.
	IsDatePreciseExtendedContext()
}

type DatePreciseExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatePreciseExtendedContext() *DatePreciseExtendedContext {
	var p = new(DatePreciseExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_datePreciseExtended
	return p
}

func (*DatePreciseExtendedContext) IsDatePreciseExtendedContext() {}

func NewDatePreciseExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatePreciseExtendedContext {
	var p = new(DatePreciseExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_datePreciseExtended

	return p
}

func (s *DatePreciseExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *DatePreciseExtendedContext) CalendarDateExtended() ICalendarDateExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalendarDateExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalendarDateExtendedContext)
}

func (s *DatePreciseExtendedContext) OrdinalDateExtended() IOrdinalDateExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrdinalDateExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrdinalDateExtendedContext)
}

func (s *DatePreciseExtendedContext) WeekDateExtended() IWeekDateExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWeekDateExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWeekDateExtendedContext)
}

func (s *DatePreciseExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatePreciseExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatePreciseExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDatePreciseExtended(s)
	}
}

func (s *DatePreciseExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDatePreciseExtended(s)
	}
}

func (p *iso8601Parser) DatePreciseExtended() (localctx IDatePreciseExtendedContext) {
	this := p
	_ = this

	localctx = NewDatePreciseExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, iso8601ParserRULE_datePreciseExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(335)
			p.CalendarDateExtended()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(336)
			p.OrdinalDateExtended()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(337)
			p.WeekDateExtended()
		}

	}

	return localctx
}

// IDateContext is an interface to support dynamic dispatch.
type IDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateContext differentiates from other interfaces.
	IsDateContext()
}

type DateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateContext() *DateContext {
	var p = new(DateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_date
	return p
}

func (*DateContext) IsDateContext() {}

func NewDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateContext {
	var p = new(DateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_date

	return p
}

func (s *DateContext) GetParser() antlr.Parser { return s.parser }

func (s *DateContext) DateBasic() IDateBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateBasicContext)
}

func (s *DateContext) DateExtended() IDateExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateExtendedContext)
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDate(s)
	}
}

func (p *iso8601Parser) Date() (localctx IDateContext) {
	this := p
	_ = this

	localctx = NewDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, iso8601ParserRULE_date)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(340)
			p.DateBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(341)
			p.DateExtended()
		}

	}

	return localctx
}

// IDateBasicContext is an interface to support dynamic dispatch.
type IDateBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateBasicContext differentiates from other interfaces.
	IsDateBasicContext()
}

type DateBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateBasicContext() *DateBasicContext {
	var p = new(DateBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_dateBasic
	return p
}

func (*DateBasicContext) IsDateBasicContext() {}

func NewDateBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateBasicContext {
	var p = new(DateBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_dateBasic

	return p
}

func (s *DateBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *DateBasicContext) DatePreciseBasic() IDatePreciseBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatePreciseBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatePreciseBasicContext)
}

func (s *DateBasicContext) SpecificMonthCalendarDate() ISpecificMonthCalendarDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificMonthCalendarDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificMonthCalendarDateContext)
}

func (s *DateBasicContext) SpecificYearCalendarDate() ISpecificYearCalendarDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificYearCalendarDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificYearCalendarDateContext)
}

func (s *DateBasicContext) SpecificCenturyCalendarDate() ISpecificCenturyCalendarDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificCenturyCalendarDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificCenturyCalendarDateContext)
}

func (s *DateBasicContext) SpecificWeekWeekDateBasic() ISpecificWeekWeekDateBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificWeekWeekDateBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificWeekWeekDateBasicContext)
}

func (s *DateBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDateBasic(s)
	}
}

func (s *DateBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDateBasic(s)
	}
}

func (p *iso8601Parser) DateBasic() (localctx IDateBasicContext) {
	this := p
	_ = this

	localctx = NewDateBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, iso8601ParserRULE_dateBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(344)
			p.DatePreciseBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)
			p.SpecificMonthCalendarDate()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(346)
			p.SpecificYearCalendarDate()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(347)
			p.SpecificCenturyCalendarDate()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(348)
			p.SpecificWeekWeekDateBasic()
		}

	}

	return localctx
}

// IDateExtendedContext is an interface to support dynamic dispatch.
type IDateExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateExtendedContext differentiates from other interfaces.
	IsDateExtendedContext()
}

type DateExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateExtendedContext() *DateExtendedContext {
	var p = new(DateExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_dateExtended
	return p
}

func (*DateExtendedContext) IsDateExtendedContext() {}

func NewDateExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateExtendedContext {
	var p = new(DateExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_dateExtended

	return p
}

func (s *DateExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *DateExtendedContext) DatePreciseExtended() IDatePreciseExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatePreciseExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatePreciseExtendedContext)
}

func (s *DateExtendedContext) SpecificMonthCalendarDate() ISpecificMonthCalendarDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificMonthCalendarDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificMonthCalendarDateContext)
}

func (s *DateExtendedContext) SpecificYearCalendarDate() ISpecificYearCalendarDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificYearCalendarDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificYearCalendarDateContext)
}

func (s *DateExtendedContext) SpecificCenturyCalendarDate() ISpecificCenturyCalendarDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificCenturyCalendarDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificCenturyCalendarDateContext)
}

func (s *DateExtendedContext) SpecificWeekWeekDateExtended() ISpecificWeekWeekDateExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificWeekWeekDateExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificWeekWeekDateExtendedContext)
}

func (s *DateExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDateExtended(s)
	}
}

func (s *DateExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDateExtended(s)
	}
}

func (p *iso8601Parser) DateExtended() (localctx IDateExtendedContext) {
	this := p
	_ = this

	localctx = NewDateExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, iso8601ParserRULE_dateExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(351)
			p.DatePreciseExtended()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(352)
			p.SpecificMonthCalendarDate()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(353)
			p.SpecificYearCalendarDate()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(354)
			p.SpecificCenturyCalendarDate()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(355)
			p.SpecificWeekWeekDateExtended()
		}

	}

	return localctx
}

// ILocalTimePreciseContext is an interface to support dynamic dispatch.
type ILocalTimePreciseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalTimePreciseContext differentiates from other interfaces.
	IsLocalTimePreciseContext()
}

type LocalTimePreciseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalTimePreciseContext() *LocalTimePreciseContext {
	var p = new(LocalTimePreciseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_localTimePrecise
	return p
}

func (*LocalTimePreciseContext) IsLocalTimePreciseContext() {}

func NewLocalTimePreciseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalTimePreciseContext {
	var p = new(LocalTimePreciseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_localTimePrecise

	return p
}

func (s *LocalTimePreciseContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalTimePreciseContext) LocalTimePreciseBasic() ILocalTimePreciseBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimePreciseBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimePreciseBasicContext)
}

func (s *LocalTimePreciseContext) LocalTimePreciseExtended() ILocalTimePreciseExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimePreciseExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimePreciseExtendedContext)
}

func (s *LocalTimePreciseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalTimePreciseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalTimePreciseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterLocalTimePrecise(s)
	}
}

func (s *LocalTimePreciseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitLocalTimePrecise(s)
	}
}

func (p *iso8601Parser) LocalTimePrecise() (localctx ILocalTimePreciseContext) {
	this := p
	_ = this

	localctx = NewLocalTimePreciseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, iso8601ParserRULE_localTimePrecise)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(358)
			p.LocalTimePreciseBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(359)
			p.LocalTimePreciseExtended()
		}

	}

	return localctx
}

// ILocalTimePreciseBasicContext is an interface to support dynamic dispatch.
type ILocalTimePreciseBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalTimePreciseBasicContext differentiates from other interfaces.
	IsLocalTimePreciseBasicContext()
}

type LocalTimePreciseBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalTimePreciseBasicContext() *LocalTimePreciseBasicContext {
	var p = new(LocalTimePreciseBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_localTimePreciseBasic
	return p
}

func (*LocalTimePreciseBasicContext) IsLocalTimePreciseBasicContext() {}

func NewLocalTimePreciseBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalTimePreciseBasicContext {
	var p = new(LocalTimePreciseBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_localTimePreciseBasic

	return p
}

func (s *LocalTimePreciseBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalTimePreciseBasicContext) Hour() IHourContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHourContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHourContext)
}

func (s *LocalTimePreciseBasicContext) Minute() IMinuteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMinuteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMinuteContext)
}

func (s *LocalTimePreciseBasicContext) SecondFraction() ISecondFractionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISecondFractionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISecondFractionContext)
}

func (s *LocalTimePreciseBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalTimePreciseBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalTimePreciseBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterLocalTimePreciseBasic(s)
	}
}

func (s *LocalTimePreciseBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitLocalTimePreciseBasic(s)
	}
}

func (p *iso8601Parser) LocalTimePreciseBasic() (localctx ILocalTimePreciseBasicContext) {
	this := p
	_ = this

	localctx = NewLocalTimePreciseBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, iso8601ParserRULE_localTimePreciseBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(362)
		p.Hour()
	}
	{
		p.SetState(363)
		p.Minute()
	}
	{
		p.SetState(364)
		p.SecondFraction()
	}

	return localctx
}

// ILocalTimePreciseExtendedContext is an interface to support dynamic dispatch.
type ILocalTimePreciseExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalTimePreciseExtendedContext differentiates from other interfaces.
	IsLocalTimePreciseExtendedContext()
}

type LocalTimePreciseExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalTimePreciseExtendedContext() *LocalTimePreciseExtendedContext {
	var p = new(LocalTimePreciseExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_localTimePreciseExtended
	return p
}

func (*LocalTimePreciseExtendedContext) IsLocalTimePreciseExtendedContext() {}

func NewLocalTimePreciseExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalTimePreciseExtendedContext {
	var p = new(LocalTimePreciseExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_localTimePreciseExtended

	return p
}

func (s *LocalTimePreciseExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalTimePreciseExtendedContext) Hour() IHourContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHourContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHourContext)
}

func (s *LocalTimePreciseExtendedContext) Minute() IMinuteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMinuteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMinuteContext)
}

func (s *LocalTimePreciseExtendedContext) SecondFraction() ISecondFractionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISecondFractionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISecondFractionContext)
}

func (s *LocalTimePreciseExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalTimePreciseExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalTimePreciseExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterLocalTimePreciseExtended(s)
	}
}

func (s *LocalTimePreciseExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitLocalTimePreciseExtended(s)
	}
}

func (p *iso8601Parser) LocalTimePreciseExtended() (localctx ILocalTimePreciseExtendedContext) {
	this := p
	_ = this

	localctx = NewLocalTimePreciseExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, iso8601ParserRULE_localTimePreciseExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(366)
		p.Hour()
	}
	{
		p.SetState(367)
		p.Match(iso8601ParserT__2)
	}
	{
		p.SetState(368)
		p.Minute()
	}
	{
		p.SetState(369)
		p.Match(iso8601ParserT__2)
	}
	{
		p.SetState(370)
		p.SecondFraction()
	}

	return localctx
}

// ISpecificMinuteLocalTimeContext is an interface to support dynamic dispatch.
type ISpecificMinuteLocalTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecificMinuteLocalTimeContext differentiates from other interfaces.
	IsSpecificMinuteLocalTimeContext()
}

type SpecificMinuteLocalTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecificMinuteLocalTimeContext() *SpecificMinuteLocalTimeContext {
	var p = new(SpecificMinuteLocalTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_specificMinuteLocalTime
	return p
}

func (*SpecificMinuteLocalTimeContext) IsSpecificMinuteLocalTimeContext() {}

func NewSpecificMinuteLocalTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificMinuteLocalTimeContext {
	var p = new(SpecificMinuteLocalTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_specificMinuteLocalTime

	return p
}

func (s *SpecificMinuteLocalTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificMinuteLocalTimeContext) SpecificMinuteLocalTimeBasic() ISpecificMinuteLocalTimeBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificMinuteLocalTimeBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificMinuteLocalTimeBasicContext)
}

func (s *SpecificMinuteLocalTimeContext) SpecificMinuteLocalTimeExtended() ISpecificMinuteLocalTimeExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificMinuteLocalTimeExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificMinuteLocalTimeExtendedContext)
}

func (s *SpecificMinuteLocalTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificMinuteLocalTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecificMinuteLocalTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSpecificMinuteLocalTime(s)
	}
}

func (s *SpecificMinuteLocalTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSpecificMinuteLocalTime(s)
	}
}

func (p *iso8601Parser) SpecificMinuteLocalTime() (localctx ISpecificMinuteLocalTimeContext) {
	this := p
	_ = this

	localctx = NewSpecificMinuteLocalTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, iso8601ParserRULE_specificMinuteLocalTime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(374)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(372)
			p.SpecificMinuteLocalTimeBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(373)
			p.SpecificMinuteLocalTimeExtended()
		}

	}

	return localctx
}

// ISpecificMinuteLocalTimeBasicContext is an interface to support dynamic dispatch.
type ISpecificMinuteLocalTimeBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecificMinuteLocalTimeBasicContext differentiates from other interfaces.
	IsSpecificMinuteLocalTimeBasicContext()
}

type SpecificMinuteLocalTimeBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecificMinuteLocalTimeBasicContext() *SpecificMinuteLocalTimeBasicContext {
	var p = new(SpecificMinuteLocalTimeBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_specificMinuteLocalTimeBasic
	return p
}

func (*SpecificMinuteLocalTimeBasicContext) IsSpecificMinuteLocalTimeBasicContext() {}

func NewSpecificMinuteLocalTimeBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificMinuteLocalTimeBasicContext {
	var p = new(SpecificMinuteLocalTimeBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_specificMinuteLocalTimeBasic

	return p
}

func (s *SpecificMinuteLocalTimeBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificMinuteLocalTimeBasicContext) Hour() IHourContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHourContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHourContext)
}

func (s *SpecificMinuteLocalTimeBasicContext) MinuteFraction() IMinuteFractionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMinuteFractionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMinuteFractionContext)
}

func (s *SpecificMinuteLocalTimeBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificMinuteLocalTimeBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecificMinuteLocalTimeBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSpecificMinuteLocalTimeBasic(s)
	}
}

func (s *SpecificMinuteLocalTimeBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSpecificMinuteLocalTimeBasic(s)
	}
}

func (p *iso8601Parser) SpecificMinuteLocalTimeBasic() (localctx ISpecificMinuteLocalTimeBasicContext) {
	this := p
	_ = this

	localctx = NewSpecificMinuteLocalTimeBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, iso8601ParserRULE_specificMinuteLocalTimeBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(376)
		p.Hour()
	}
	{
		p.SetState(377)
		p.MinuteFraction()
	}

	return localctx
}

// ISpecificMinuteLocalTimeExtendedContext is an interface to support dynamic dispatch.
type ISpecificMinuteLocalTimeExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecificMinuteLocalTimeExtendedContext differentiates from other interfaces.
	IsSpecificMinuteLocalTimeExtendedContext()
}

type SpecificMinuteLocalTimeExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecificMinuteLocalTimeExtendedContext() *SpecificMinuteLocalTimeExtendedContext {
	var p = new(SpecificMinuteLocalTimeExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_specificMinuteLocalTimeExtended
	return p
}

func (*SpecificMinuteLocalTimeExtendedContext) IsSpecificMinuteLocalTimeExtendedContext() {}

func NewSpecificMinuteLocalTimeExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificMinuteLocalTimeExtendedContext {
	var p = new(SpecificMinuteLocalTimeExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_specificMinuteLocalTimeExtended

	return p
}

func (s *SpecificMinuteLocalTimeExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificMinuteLocalTimeExtendedContext) Hour() IHourContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHourContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHourContext)
}

func (s *SpecificMinuteLocalTimeExtendedContext) MinuteFraction() IMinuteFractionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMinuteFractionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMinuteFractionContext)
}

func (s *SpecificMinuteLocalTimeExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificMinuteLocalTimeExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecificMinuteLocalTimeExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSpecificMinuteLocalTimeExtended(s)
	}
}

func (s *SpecificMinuteLocalTimeExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSpecificMinuteLocalTimeExtended(s)
	}
}

func (p *iso8601Parser) SpecificMinuteLocalTimeExtended() (localctx ISpecificMinuteLocalTimeExtendedContext) {
	this := p
	_ = this

	localctx = NewSpecificMinuteLocalTimeExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, iso8601ParserRULE_specificMinuteLocalTimeExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(379)
		p.Hour()
	}
	{
		p.SetState(380)
		p.Match(iso8601ParserT__2)
	}
	{
		p.SetState(381)
		p.MinuteFraction()
	}

	return localctx
}

// ISpecificHourLocalTimeContext is an interface to support dynamic dispatch.
type ISpecificHourLocalTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecificHourLocalTimeContext differentiates from other interfaces.
	IsSpecificHourLocalTimeContext()
}

type SpecificHourLocalTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecificHourLocalTimeContext() *SpecificHourLocalTimeContext {
	var p = new(SpecificHourLocalTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_specificHourLocalTime
	return p
}

func (*SpecificHourLocalTimeContext) IsSpecificHourLocalTimeContext() {}

func NewSpecificHourLocalTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecificHourLocalTimeContext {
	var p = new(SpecificHourLocalTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_specificHourLocalTime

	return p
}

func (s *SpecificHourLocalTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecificHourLocalTimeContext) HourFraction() IHourFractionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHourFractionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHourFractionContext)
}

func (s *SpecificHourLocalTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecificHourLocalTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecificHourLocalTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterSpecificHourLocalTime(s)
	}
}

func (s *SpecificHourLocalTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitSpecificHourLocalTime(s)
	}
}

func (p *iso8601Parser) SpecificHourLocalTime() (localctx ISpecificHourLocalTimeContext) {
	this := p
	_ = this

	localctx = NewSpecificHourLocalTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, iso8601ParserRULE_specificHourLocalTime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(383)
		p.HourFraction()
	}

	return localctx
}

// ILocalTimeContext is an interface to support dynamic dispatch.
type ILocalTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalTimeContext differentiates from other interfaces.
	IsLocalTimeContext()
}

type LocalTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalTimeContext() *LocalTimeContext {
	var p = new(LocalTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_localTime
	return p
}

func (*LocalTimeContext) IsLocalTimeContext() {}

func NewLocalTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalTimeContext {
	var p = new(LocalTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_localTime

	return p
}

func (s *LocalTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalTimeContext) LocalTimeBasic() ILocalTimeBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimeBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimeBasicContext)
}

func (s *LocalTimeContext) LocalTimeExtended() ILocalTimeExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimeExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimeExtendedContext)
}

func (s *LocalTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterLocalTime(s)
	}
}

func (s *LocalTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitLocalTime(s)
	}
}

func (p *iso8601Parser) LocalTime() (localctx ILocalTimeContext) {
	this := p
	_ = this

	localctx = NewLocalTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, iso8601ParserRULE_localTime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(385)
			p.LocalTimeBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(386)
			p.LocalTimeExtended()
		}

	}

	return localctx
}

// ILocalTimeBasicContext is an interface to support dynamic dispatch.
type ILocalTimeBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalTimeBasicContext differentiates from other interfaces.
	IsLocalTimeBasicContext()
}

type LocalTimeBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalTimeBasicContext() *LocalTimeBasicContext {
	var p = new(LocalTimeBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_localTimeBasic
	return p
}

func (*LocalTimeBasicContext) IsLocalTimeBasicContext() {}

func NewLocalTimeBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalTimeBasicContext {
	var p = new(LocalTimeBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_localTimeBasic

	return p
}

func (s *LocalTimeBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalTimeBasicContext) LocalTimePreciseBasic() ILocalTimePreciseBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimePreciseBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimePreciseBasicContext)
}

func (s *LocalTimeBasicContext) SpecificMinuteLocalTimeBasic() ISpecificMinuteLocalTimeBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificMinuteLocalTimeBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificMinuteLocalTimeBasicContext)
}

func (s *LocalTimeBasicContext) SpecificHourLocalTime() ISpecificHourLocalTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificHourLocalTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificHourLocalTimeContext)
}

func (s *LocalTimeBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalTimeBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalTimeBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterLocalTimeBasic(s)
	}
}

func (s *LocalTimeBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitLocalTimeBasic(s)
	}
}

func (p *iso8601Parser) LocalTimeBasic() (localctx ILocalTimeBasicContext) {
	this := p
	_ = this

	localctx = NewLocalTimeBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, iso8601ParserRULE_localTimeBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(392)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(389)
			p.LocalTimePreciseBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(390)
			p.SpecificMinuteLocalTimeBasic()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(391)
			p.SpecificHourLocalTime()
		}

	}

	return localctx
}

// ILocalTimeExtendedContext is an interface to support dynamic dispatch.
type ILocalTimeExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalTimeExtendedContext differentiates from other interfaces.
	IsLocalTimeExtendedContext()
}

type LocalTimeExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalTimeExtendedContext() *LocalTimeExtendedContext {
	var p = new(LocalTimeExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_localTimeExtended
	return p
}

func (*LocalTimeExtendedContext) IsLocalTimeExtendedContext() {}

func NewLocalTimeExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalTimeExtendedContext {
	var p = new(LocalTimeExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_localTimeExtended

	return p
}

func (s *LocalTimeExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalTimeExtendedContext) LocalTimePreciseExtended() ILocalTimePreciseExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimePreciseExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimePreciseExtendedContext)
}

func (s *LocalTimeExtendedContext) SpecificMinuteLocalTimeExtended() ISpecificMinuteLocalTimeExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificMinuteLocalTimeExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificMinuteLocalTimeExtendedContext)
}

func (s *LocalTimeExtendedContext) SpecificHourLocalTime() ISpecificHourLocalTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecificHourLocalTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecificHourLocalTimeContext)
}

func (s *LocalTimeExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalTimeExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalTimeExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterLocalTimeExtended(s)
	}
}

func (s *LocalTimeExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitLocalTimeExtended(s)
	}
}

func (p *iso8601Parser) LocalTimeExtended() (localctx ILocalTimeExtendedContext) {
	this := p
	_ = this

	localctx = NewLocalTimeExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, iso8601ParserRULE_localTimeExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(394)
			p.LocalTimePreciseExtended()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(395)
			p.SpecificMinuteLocalTimeExtended()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(396)
			p.SpecificHourLocalTime()
		}

	}

	return localctx
}

// ITimeZoneUtcContext is an interface to support dynamic dispatch.
type ITimeZoneUtcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeZoneUtcContext differentiates from other interfaces.
	IsTimeZoneUtcContext()
}

type TimeZoneUtcContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeZoneUtcContext() *TimeZoneUtcContext {
	var p = new(TimeZoneUtcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeZoneUtc
	return p
}

func (*TimeZoneUtcContext) IsTimeZoneUtcContext() {}

func NewTimeZoneUtcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeZoneUtcContext {
	var p = new(TimeZoneUtcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeZoneUtc

	return p
}

func (s *TimeZoneUtcContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeZoneUtcContext) Z() antlr.TerminalNode {
	return s.GetToken(iso8601ParserZ, 0)
}

func (s *TimeZoneUtcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeZoneUtcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeZoneUtcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeZoneUtc(s)
	}
}

func (s *TimeZoneUtcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeZoneUtc(s)
	}
}

func (p *iso8601Parser) TimeZoneUtc() (localctx ITimeZoneUtcContext) {
	this := p
	_ = this

	localctx = NewTimeZoneUtcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, iso8601ParserRULE_timeZoneUtc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(399)
		p.Match(iso8601ParserZ)
	}

	return localctx
}

// ITimeZoneContext is an interface to support dynamic dispatch.
type ITimeZoneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeZoneContext differentiates from other interfaces.
	IsTimeZoneContext()
}

type TimeZoneContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeZoneContext() *TimeZoneContext {
	var p = new(TimeZoneContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeZone
	return p
}

func (*TimeZoneContext) IsTimeZoneContext() {}

func NewTimeZoneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeZoneContext {
	var p = new(TimeZoneContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeZone

	return p
}

func (s *TimeZoneContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeZoneContext) TimeZoneBasic() ITimeZoneBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeZoneBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeZoneBasicContext)
}

func (s *TimeZoneContext) TimeZoneExtended() ITimeZoneExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeZoneExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeZoneExtendedContext)
}

func (s *TimeZoneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeZoneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeZoneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeZone(s)
	}
}

func (s *TimeZoneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeZone(s)
	}
}

func (p *iso8601Parser) TimeZone() (localctx ITimeZoneContext) {
	this := p
	_ = this

	localctx = NewTimeZoneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, iso8601ParserRULE_timeZone)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(401)
			p.TimeZoneBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(402)
			p.TimeZoneExtended()
		}

	}

	return localctx
}

// ITimeZoneBasicContext is an interface to support dynamic dispatch.
type ITimeZoneBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeZoneBasicContext differentiates from other interfaces.
	IsTimeZoneBasicContext()
}

type TimeZoneBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeZoneBasicContext() *TimeZoneBasicContext {
	var p = new(TimeZoneBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeZoneBasic
	return p
}

func (*TimeZoneBasicContext) IsTimeZoneBasicContext() {}

func NewTimeZoneBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeZoneBasicContext {
	var p = new(TimeZoneBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeZoneBasic

	return p
}

func (s *TimeZoneBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeZoneBasicContext) Hour() IHourContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHourContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHourContext)
}

func (s *TimeZoneBasicContext) Minute() IMinuteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMinuteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMinuteContext)
}

func (s *TimeZoneBasicContext) TimeZoneUtc() ITimeZoneUtcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeZoneUtcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeZoneUtcContext)
}

func (s *TimeZoneBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeZoneBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeZoneBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeZoneBasic(s)
	}
}

func (s *TimeZoneBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeZoneBasic(s)
	}
}

func (p *iso8601Parser) TimeZoneBasic() (localctx ITimeZoneBasicContext) {
	this := p
	_ = this

	localctx = NewTimeZoneBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, iso8601ParserRULE_timeZoneBasic)
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

	p.SetState(411)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case iso8601ParserT__0, iso8601ParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			_la = p.GetTokenStream().LA(1)

			if !(_la == iso8601ParserT__0 || _la == iso8601ParserT__1) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(406)
			p.Hour()
		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == iso8601ParserDigit {
			{
				p.SetState(407)
				p.Minute()
			}

		}

	case iso8601ParserZ:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.TimeZoneUtc()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITimeZoneExtendedContext is an interface to support dynamic dispatch.
type ITimeZoneExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeZoneExtendedContext differentiates from other interfaces.
	IsTimeZoneExtendedContext()
}

type TimeZoneExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeZoneExtendedContext() *TimeZoneExtendedContext {
	var p = new(TimeZoneExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeZoneExtended
	return p
}

func (*TimeZoneExtendedContext) IsTimeZoneExtendedContext() {}

func NewTimeZoneExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeZoneExtendedContext {
	var p = new(TimeZoneExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeZoneExtended

	return p
}

func (s *TimeZoneExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeZoneExtendedContext) Hour() IHourContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHourContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHourContext)
}

func (s *TimeZoneExtendedContext) Minute() IMinuteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMinuteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMinuteContext)
}

func (s *TimeZoneExtendedContext) TimeZoneUtc() ITimeZoneUtcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeZoneUtcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeZoneUtcContext)
}

func (s *TimeZoneExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeZoneExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeZoneExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeZoneExtended(s)
	}
}

func (s *TimeZoneExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeZoneExtended(s)
	}
}

func (p *iso8601Parser) TimeZoneExtended() (localctx ITimeZoneExtendedContext) {
	this := p
	_ = this

	localctx = NewTimeZoneExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, iso8601ParserRULE_timeZoneExtended)
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

	p.SetState(420)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case iso8601ParserT__0, iso8601ParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(413)
			_la = p.GetTokenStream().LA(1)

			if !(_la == iso8601ParserT__0 || _la == iso8601ParserT__1) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(414)
			p.Hour()
		}
		p.SetState(417)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == iso8601ParserT__2 {
			{
				p.SetState(415)
				p.Match(iso8601ParserT__2)
			}
			{
				p.SetState(416)
				p.Minute()
			}

		}

	case iso8601ParserZ:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(419)
			p.TimeZoneUtc()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILocalTimeAndTimeZoneContext is an interface to support dynamic dispatch.
type ILocalTimeAndTimeZoneContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalTimeAndTimeZoneContext differentiates from other interfaces.
	IsLocalTimeAndTimeZoneContext()
}

type LocalTimeAndTimeZoneContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalTimeAndTimeZoneContext() *LocalTimeAndTimeZoneContext {
	var p = new(LocalTimeAndTimeZoneContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_localTimeAndTimeZone
	return p
}

func (*LocalTimeAndTimeZoneContext) IsLocalTimeAndTimeZoneContext() {}

func NewLocalTimeAndTimeZoneContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalTimeAndTimeZoneContext {
	var p = new(LocalTimeAndTimeZoneContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_localTimeAndTimeZone

	return p
}

func (s *LocalTimeAndTimeZoneContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalTimeAndTimeZoneContext) LocalTimeAndTimeZoneBasic() ILocalTimeAndTimeZoneBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimeAndTimeZoneBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimeAndTimeZoneBasicContext)
}

func (s *LocalTimeAndTimeZoneContext) LocalTimeAndTimeZoneExtended() ILocalTimeAndTimeZoneExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimeAndTimeZoneExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimeAndTimeZoneExtendedContext)
}

func (s *LocalTimeAndTimeZoneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalTimeAndTimeZoneContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalTimeAndTimeZoneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterLocalTimeAndTimeZone(s)
	}
}

func (s *LocalTimeAndTimeZoneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitLocalTimeAndTimeZone(s)
	}
}

func (p *iso8601Parser) LocalTimeAndTimeZone() (localctx ILocalTimeAndTimeZoneContext) {
	this := p
	_ = this

	localctx = NewLocalTimeAndTimeZoneContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, iso8601ParserRULE_localTimeAndTimeZone)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(422)
			p.LocalTimeAndTimeZoneBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(423)
			p.LocalTimeAndTimeZoneExtended()
		}

	}

	return localctx
}

// ILocalTimeAndTimeZoneBasicContext is an interface to support dynamic dispatch.
type ILocalTimeAndTimeZoneBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalTimeAndTimeZoneBasicContext differentiates from other interfaces.
	IsLocalTimeAndTimeZoneBasicContext()
}

type LocalTimeAndTimeZoneBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalTimeAndTimeZoneBasicContext() *LocalTimeAndTimeZoneBasicContext {
	var p = new(LocalTimeAndTimeZoneBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_localTimeAndTimeZoneBasic
	return p
}

func (*LocalTimeAndTimeZoneBasicContext) IsLocalTimeAndTimeZoneBasicContext() {}

func NewLocalTimeAndTimeZoneBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalTimeAndTimeZoneBasicContext {
	var p = new(LocalTimeAndTimeZoneBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_localTimeAndTimeZoneBasic

	return p
}

func (s *LocalTimeAndTimeZoneBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalTimeAndTimeZoneBasicContext) LocalTimeBasic() ILocalTimeBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimeBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimeBasicContext)
}

func (s *LocalTimeAndTimeZoneBasicContext) TimeZoneBasic() ITimeZoneBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeZoneBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeZoneBasicContext)
}

func (s *LocalTimeAndTimeZoneBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalTimeAndTimeZoneBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalTimeAndTimeZoneBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterLocalTimeAndTimeZoneBasic(s)
	}
}

func (s *LocalTimeAndTimeZoneBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitLocalTimeAndTimeZoneBasic(s)
	}
}

func (p *iso8601Parser) LocalTimeAndTimeZoneBasic() (localctx ILocalTimeAndTimeZoneBasicContext) {
	this := p
	_ = this

	localctx = NewLocalTimeAndTimeZoneBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, iso8601ParserRULE_localTimeAndTimeZoneBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(426)
		p.LocalTimeBasic()
	}
	{
		p.SetState(427)
		p.TimeZoneBasic()
	}

	return localctx
}

// ILocalTimeAndTimeZoneExtendedContext is an interface to support dynamic dispatch.
type ILocalTimeAndTimeZoneExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLocalTimeAndTimeZoneExtendedContext differentiates from other interfaces.
	IsLocalTimeAndTimeZoneExtendedContext()
}

type LocalTimeAndTimeZoneExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLocalTimeAndTimeZoneExtendedContext() *LocalTimeAndTimeZoneExtendedContext {
	var p = new(LocalTimeAndTimeZoneExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_localTimeAndTimeZoneExtended
	return p
}

func (*LocalTimeAndTimeZoneExtendedContext) IsLocalTimeAndTimeZoneExtendedContext() {}

func NewLocalTimeAndTimeZoneExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LocalTimeAndTimeZoneExtendedContext {
	var p = new(LocalTimeAndTimeZoneExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_localTimeAndTimeZoneExtended

	return p
}

func (s *LocalTimeAndTimeZoneExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *LocalTimeAndTimeZoneExtendedContext) LocalTimeExtended() ILocalTimeExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimeExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimeExtendedContext)
}

func (s *LocalTimeAndTimeZoneExtendedContext) TimeZoneExtended() ITimeZoneExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeZoneExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeZoneExtendedContext)
}

func (s *LocalTimeAndTimeZoneExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LocalTimeAndTimeZoneExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LocalTimeAndTimeZoneExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterLocalTimeAndTimeZoneExtended(s)
	}
}

func (s *LocalTimeAndTimeZoneExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitLocalTimeAndTimeZoneExtended(s)
	}
}

func (p *iso8601Parser) LocalTimeAndTimeZoneExtended() (localctx ILocalTimeAndTimeZoneExtendedContext) {
	this := p
	_ = this

	localctx = NewLocalTimeAndTimeZoneExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, iso8601ParserRULE_localTimeAndTimeZoneExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(429)
		p.LocalTimeExtended()
	}
	{
		p.SetState(430)
		p.TimeZoneExtended()
	}

	return localctx
}

// ITimeContext is an interface to support dynamic dispatch.
type ITimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeContext differentiates from other interfaces.
	IsTimeContext()
}

type TimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeContext() *TimeContext {
	var p = new(TimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_time
	return p
}

func (*TimeContext) IsTimeContext() {}

func NewTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeContext {
	var p = new(TimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_time

	return p
}

func (s *TimeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeContext) LocalTime() ILocalTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimeContext)
}

func (s *TimeContext) T() antlr.TerminalNode {
	return s.GetToken(iso8601ParserT, 0)
}

func (s *TimeContext) TimeZone() ITimeZoneContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeZoneContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeZoneContext)
}

func (s *TimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTime(s)
	}
}

func (s *TimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTime(s)
	}
}

func (p *iso8601Parser) Time() (localctx ITimeContext) {
	this := p
	_ = this

	localctx = NewTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, iso8601ParserRULE_time)
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
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == iso8601ParserT {
		{
			p.SetState(432)
			p.Match(iso8601ParserT)
		}

	}
	{
		p.SetState(435)
		p.LocalTime()
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<iso8601ParserT__0)|(1<<iso8601ParserT__1)|(1<<iso8601ParserZ))) != 0 {
		{
			p.SetState(436)
			p.TimeZone()
		}

	}

	return localctx
}

// IDatetimePreciseContext is an interface to support dynamic dispatch.
type IDatetimePreciseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatetimePreciseContext differentiates from other interfaces.
	IsDatetimePreciseContext()
}

type DatetimePreciseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetimePreciseContext() *DatetimePreciseContext {
	var p = new(DatetimePreciseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_datetimePrecise
	return p
}

func (*DatetimePreciseContext) IsDatetimePreciseContext() {}

func NewDatetimePreciseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatetimePreciseContext {
	var p = new(DatetimePreciseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_datetimePrecise

	return p
}

func (s *DatetimePreciseContext) GetParser() antlr.Parser { return s.parser }

func (s *DatetimePreciseContext) DatetimePreciseBasic() IDatetimePreciseBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimePreciseBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimePreciseBasicContext)
}

func (s *DatetimePreciseContext) DatetimePreciseExtended() IDatetimePreciseExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimePreciseExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimePreciseExtendedContext)
}

func (s *DatetimePreciseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimePreciseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatetimePreciseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDatetimePrecise(s)
	}
}

func (s *DatetimePreciseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDatetimePrecise(s)
	}
}

func (p *iso8601Parser) DatetimePrecise() (localctx IDatetimePreciseContext) {
	this := p
	_ = this

	localctx = NewDatetimePreciseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, iso8601ParserRULE_datetimePrecise)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(439)
			p.DatetimePreciseBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(440)
			p.DatetimePreciseExtended()
		}

	}

	return localctx
}

// IDatetimePreciseBasicContext is an interface to support dynamic dispatch.
type IDatetimePreciseBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatetimePreciseBasicContext differentiates from other interfaces.
	IsDatetimePreciseBasicContext()
}

type DatetimePreciseBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetimePreciseBasicContext() *DatetimePreciseBasicContext {
	var p = new(DatetimePreciseBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_datetimePreciseBasic
	return p
}

func (*DatetimePreciseBasicContext) IsDatetimePreciseBasicContext() {}

func NewDatetimePreciseBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatetimePreciseBasicContext {
	var p = new(DatetimePreciseBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_datetimePreciseBasic

	return p
}

func (s *DatetimePreciseBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *DatetimePreciseBasicContext) CalendarDateBasic() ICalendarDateBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalendarDateBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalendarDateBasicContext)
}

func (s *DatetimePreciseBasicContext) T() antlr.TerminalNode {
	return s.GetToken(iso8601ParserT, 0)
}

func (s *DatetimePreciseBasicContext) LocalTimePreciseBasic() ILocalTimePreciseBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimePreciseBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimePreciseBasicContext)
}

func (s *DatetimePreciseBasicContext) TimeZoneBasic() ITimeZoneBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeZoneBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeZoneBasicContext)
}

func (s *DatetimePreciseBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimePreciseBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatetimePreciseBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDatetimePreciseBasic(s)
	}
}

func (s *DatetimePreciseBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDatetimePreciseBasic(s)
	}
}

func (p *iso8601Parser) DatetimePreciseBasic() (localctx IDatetimePreciseBasicContext) {
	this := p
	_ = this

	localctx = NewDatetimePreciseBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, iso8601ParserRULE_datetimePreciseBasic)
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
		p.CalendarDateBasic()
	}
	{
		p.SetState(444)
		p.Match(iso8601ParserT)
	}
	{
		p.SetState(445)
		p.LocalTimePreciseBasic()
	}
	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<iso8601ParserT__0)|(1<<iso8601ParserT__1)|(1<<iso8601ParserZ))) != 0 {
		{
			p.SetState(446)
			p.TimeZoneBasic()
		}

	}

	return localctx
}

// IDatetimePreciseExtendedContext is an interface to support dynamic dispatch.
type IDatetimePreciseExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatetimePreciseExtendedContext differentiates from other interfaces.
	IsDatetimePreciseExtendedContext()
}

type DatetimePreciseExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetimePreciseExtendedContext() *DatetimePreciseExtendedContext {
	var p = new(DatetimePreciseExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_datetimePreciseExtended
	return p
}

func (*DatetimePreciseExtendedContext) IsDatetimePreciseExtendedContext() {}

func NewDatetimePreciseExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatetimePreciseExtendedContext {
	var p = new(DatetimePreciseExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_datetimePreciseExtended

	return p
}

func (s *DatetimePreciseExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *DatetimePreciseExtendedContext) CalendarDateExtended() ICalendarDateExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICalendarDateExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICalendarDateExtendedContext)
}

func (s *DatetimePreciseExtendedContext) T() antlr.TerminalNode {
	return s.GetToken(iso8601ParserT, 0)
}

func (s *DatetimePreciseExtendedContext) LocalTimePreciseExtended() ILocalTimePreciseExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimePreciseExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimePreciseExtendedContext)
}

func (s *DatetimePreciseExtendedContext) TimeZoneExtended() ITimeZoneExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeZoneExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeZoneExtendedContext)
}

func (s *DatetimePreciseExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimePreciseExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatetimePreciseExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDatetimePreciseExtended(s)
	}
}

func (s *DatetimePreciseExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDatetimePreciseExtended(s)
	}
}

func (p *iso8601Parser) DatetimePreciseExtended() (localctx IDatetimePreciseExtendedContext) {
	this := p
	_ = this

	localctx = NewDatetimePreciseExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, iso8601ParserRULE_datetimePreciseExtended)
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
		p.SetState(449)
		p.CalendarDateExtended()
	}
	{
		p.SetState(450)
		p.Match(iso8601ParserT)
	}
	{
		p.SetState(451)
		p.LocalTimePreciseExtended()
	}
	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<iso8601ParserT__0)|(1<<iso8601ParserT__1)|(1<<iso8601ParserZ))) != 0 {
		{
			p.SetState(452)
			p.TimeZoneExtended()
		}

	}

	return localctx
}

// IDatetimeContext is an interface to support dynamic dispatch.
type IDatetimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatetimeContext differentiates from other interfaces.
	IsDatetimeContext()
}

type DatetimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetimeContext() *DatetimeContext {
	var p = new(DatetimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_datetime
	return p
}

func (*DatetimeContext) IsDatetimeContext() {}

func NewDatetimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatetimeContext {
	var p = new(DatetimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_datetime

	return p
}

func (s *DatetimeContext) GetParser() antlr.Parser { return s.parser }

func (s *DatetimeContext) DatetimeBasic() IDatetimeBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeBasicContext)
}

func (s *DatetimeContext) DatetimeExtended() IDatetimeExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeExtendedContext)
}

func (s *DatetimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatetimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDatetime(s)
	}
}

func (s *DatetimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDatetime(s)
	}
}

func (p *iso8601Parser) Datetime() (localctx IDatetimeContext) {
	this := p
	_ = this

	localctx = NewDatetimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, iso8601ParserRULE_datetime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(455)
			p.DatetimeBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(456)
			p.DatetimeExtended()
		}

	}

	return localctx
}

// IDatetimeBasicContext is an interface to support dynamic dispatch.
type IDatetimeBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatetimeBasicContext differentiates from other interfaces.
	IsDatetimeBasicContext()
}

type DatetimeBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetimeBasicContext() *DatetimeBasicContext {
	var p = new(DatetimeBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_datetimeBasic
	return p
}

func (*DatetimeBasicContext) IsDatetimeBasicContext() {}

func NewDatetimeBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatetimeBasicContext {
	var p = new(DatetimeBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_datetimeBasic

	return p
}

func (s *DatetimeBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *DatetimeBasicContext) DatePreciseBasic() IDatePreciseBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatePreciseBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatePreciseBasicContext)
}

func (s *DatetimeBasicContext) T() antlr.TerminalNode {
	return s.GetToken(iso8601ParserT, 0)
}

func (s *DatetimeBasicContext) LocalTimeBasic() ILocalTimeBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimeBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimeBasicContext)
}

func (s *DatetimeBasicContext) TimeZoneBasic() ITimeZoneBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeZoneBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeZoneBasicContext)
}

func (s *DatetimeBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimeBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatetimeBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDatetimeBasic(s)
	}
}

func (s *DatetimeBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDatetimeBasic(s)
	}
}

func (p *iso8601Parser) DatetimeBasic() (localctx IDatetimeBasicContext) {
	this := p
	_ = this

	localctx = NewDatetimeBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, iso8601ParserRULE_datetimeBasic)
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
		p.SetState(459)
		p.DatePreciseBasic()
	}
	{
		p.SetState(460)
		p.Match(iso8601ParserT)
	}
	{
		p.SetState(461)
		p.LocalTimeBasic()
	}
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<iso8601ParserT__0)|(1<<iso8601ParserT__1)|(1<<iso8601ParserZ))) != 0 {
		{
			p.SetState(462)
			p.TimeZoneBasic()
		}

	}

	return localctx
}

// IDatetimeExtendedContext is an interface to support dynamic dispatch.
type IDatetimeExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatetimeExtendedContext differentiates from other interfaces.
	IsDatetimeExtendedContext()
}

type DatetimeExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetimeExtendedContext() *DatetimeExtendedContext {
	var p = new(DatetimeExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_datetimeExtended
	return p
}

func (*DatetimeExtendedContext) IsDatetimeExtendedContext() {}

func NewDatetimeExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatetimeExtendedContext {
	var p = new(DatetimeExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_datetimeExtended

	return p
}

func (s *DatetimeExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *DatetimeExtendedContext) DatePreciseExtended() IDatePreciseExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatePreciseExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatePreciseExtendedContext)
}

func (s *DatetimeExtendedContext) T() antlr.TerminalNode {
	return s.GetToken(iso8601ParserT, 0)
}

func (s *DatetimeExtendedContext) LocalTimeExtended() ILocalTimeExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimeExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimeExtendedContext)
}

func (s *DatetimeExtendedContext) TimeZoneExtended() ITimeZoneExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeZoneExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeZoneExtendedContext)
}

func (s *DatetimeExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimeExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatetimeExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDatetimeExtended(s)
	}
}

func (s *DatetimeExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDatetimeExtended(s)
	}
}

func (p *iso8601Parser) DatetimeExtended() (localctx IDatetimeExtendedContext) {
	this := p
	_ = this

	localctx = NewDatetimeExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, iso8601ParserRULE_datetimeExtended)
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
		p.SetState(465)
		p.DatePreciseExtended()
	}
	{
		p.SetState(466)
		p.Match(iso8601ParserT)
	}
	{
		p.SetState(467)
		p.LocalTimeExtended()
	}
	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<iso8601ParserT__0)|(1<<iso8601ParserT__1)|(1<<iso8601ParserZ))) != 0 {
		{
			p.SetState(468)
			p.TimeZoneExtended()
		}

	}

	return localctx
}

// IIntervalContext is an interface to support dynamic dispatch.
type IIntervalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntervalContext differentiates from other interfaces.
	IsIntervalContext()
}

type IntervalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalContext() *IntervalContext {
	var p = new(IntervalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_interval
	return p
}

func (*IntervalContext) IsIntervalContext() {}

func NewIntervalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalContext {
	var p = new(IntervalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_interval

	return p
}

func (s *IntervalContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalContext) IntervalT() IIntervalTContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalTContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalTContext)
}

func (s *IntervalContext) IntervalYMD() IIntervalYMDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalYMDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalYMDContext)
}

func (s *IntervalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterInterval(s)
	}
}

func (s *IntervalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitInterval(s)
	}
}

func (p *iso8601Parser) Interval() (localctx IIntervalContext) {
	this := p
	_ = this

	localctx = NewIntervalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, iso8601ParserRULE_interval)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(471)
			p.IntervalT()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(472)
			p.IntervalYMD()
		}

	}

	return localctx
}

// IIntervalBasicContext is an interface to support dynamic dispatch.
type IIntervalBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntervalBasicContext differentiates from other interfaces.
	IsIntervalBasicContext()
}

type IntervalBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalBasicContext() *IntervalBasicContext {
	var p = new(IntervalBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_intervalBasic
	return p
}

func (*IntervalBasicContext) IsIntervalBasicContext() {}

func NewIntervalBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalBasicContext {
	var p = new(IntervalBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_intervalBasic

	return p
}

func (s *IntervalBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalBasicContext) IntervalT() IIntervalTContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalTContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalTContext)
}

func (s *IntervalBasicContext) IntervalYMDBasic() IIntervalYMDBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalYMDBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalYMDBasicContext)
}

func (s *IntervalBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterIntervalBasic(s)
	}
}

func (s *IntervalBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitIntervalBasic(s)
	}
}

func (p *iso8601Parser) IntervalBasic() (localctx IIntervalBasicContext) {
	this := p
	_ = this

	localctx = NewIntervalBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, iso8601ParserRULE_intervalBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(475)
			p.IntervalT()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(476)
			p.IntervalYMDBasic()
		}

	}

	return localctx
}

// IIntervalExtendedContext is an interface to support dynamic dispatch.
type IIntervalExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntervalExtendedContext differentiates from other interfaces.
	IsIntervalExtendedContext()
}

type IntervalExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalExtendedContext() *IntervalExtendedContext {
	var p = new(IntervalExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_intervalExtended
	return p
}

func (*IntervalExtendedContext) IsIntervalExtendedContext() {}

func NewIntervalExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalExtendedContext {
	var p = new(IntervalExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_intervalExtended

	return p
}

func (s *IntervalExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalExtendedContext) IntervalT() IIntervalTContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalTContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalTContext)
}

func (s *IntervalExtendedContext) IntervalYMDExtended() IIntervalYMDExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalYMDExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalYMDExtendedContext)
}

func (s *IntervalExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterIntervalExtended(s)
	}
}

func (s *IntervalExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitIntervalExtended(s)
	}
}

func (p *iso8601Parser) IntervalExtended() (localctx IIntervalExtendedContext) {
	this := p
	_ = this

	localctx = NewIntervalExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, iso8601ParserRULE_intervalExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(479)
			p.IntervalT()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(480)
			p.IntervalYMDExtended()
		}

	}

	return localctx
}

// IIntervalTContext is an interface to support dynamic dispatch.
type IIntervalTContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntervalTContext differentiates from other interfaces.
	IsIntervalTContext()
}

type IntervalTContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalTContext() *IntervalTContext {
	var p = new(IntervalTContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_intervalT
	return p
}

func (*IntervalTContext) IsIntervalTContext() {}

func NewIntervalTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalTContext {
	var p = new(IntervalTContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_intervalT

	return p
}

func (s *IntervalTContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalTContext) P() antlr.TerminalNode {
	return s.GetToken(iso8601ParserP, 0)
}

func (s *IntervalTContext) T() antlr.TerminalNode {
	return s.GetToken(iso8601ParserT, 0)
}

func (s *IntervalTContext) Dec() IDecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecContext)
}

func (s *IntervalTContext) S() antlr.TerminalNode {
	return s.GetToken(iso8601ParserS, 0)
}

func (s *IntervalTContext) AllInt_() []IInt_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInt_Context)(nil)).Elem())
	var tst = make([]IInt_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInt_Context)
		}
	}

	return tst
}

func (s *IntervalTContext) Int_(i int) IInt_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInt_Context)
}

func (s *IntervalTContext) Y() antlr.TerminalNode {
	return s.GetToken(iso8601ParserY, 0)
}

func (s *IntervalTContext) AllM() []antlr.TerminalNode {
	return s.GetTokens(iso8601ParserM)
}

func (s *IntervalTContext) M(i int) antlr.TerminalNode {
	return s.GetToken(iso8601ParserM, i)
}

func (s *IntervalTContext) D() antlr.TerminalNode {
	return s.GetToken(iso8601ParserD, 0)
}

func (s *IntervalTContext) H() antlr.TerminalNode {
	return s.GetToken(iso8601ParserH, 0)
}

func (s *IntervalTContext) W() antlr.TerminalNode {
	return s.GetToken(iso8601ParserW, 0)
}

func (s *IntervalTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalTContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterIntervalT(s)
	}
}

func (s *IntervalTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitIntervalT(s)
	}
}

func (p *iso8601Parser) IntervalT() (localctx IIntervalTContext) {
	this := p
	_ = this

	localctx = NewIntervalTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, iso8601ParserRULE_intervalT)
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

	p.SetState(589)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(483)
			p.Match(iso8601ParserP)
		}
		p.SetState(487)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(484)
				p.Int_()
			}
			{
				p.SetState(485)
				p.Match(iso8601ParserY)
			}

		}
		p.SetState(492)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(489)
				p.Int_()
			}
			{
				p.SetState(490)
				p.Match(iso8601ParserM)
			}

		}
		p.SetState(497)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == iso8601ParserDigit {
			{
				p.SetState(494)
				p.Int_()
			}
			{
				p.SetState(495)
				p.Match(iso8601ParserD)
			}

		}
		{
			p.SetState(499)
			p.Match(iso8601ParserT)
		}
		p.SetState(503)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(500)
				p.Int_()
			}
			{
				p.SetState(501)
				p.Match(iso8601ParserH)
			}

		}
		p.SetState(508)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(505)
				p.Int_()
			}
			{
				p.SetState(506)
				p.Match(iso8601ParserM)
			}

		}
		{
			p.SetState(510)
			p.Dec()
		}
		{
			p.SetState(511)
			p.Match(iso8601ParserS)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(513)
			p.Match(iso8601ParserP)
		}
		p.SetState(517)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(514)
				p.Int_()
			}
			{
				p.SetState(515)
				p.Match(iso8601ParserY)
			}

		}
		p.SetState(522)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(519)
				p.Int_()
			}
			{
				p.SetState(520)
				p.Match(iso8601ParserM)
			}

		}
		p.SetState(527)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == iso8601ParserDigit {
			{
				p.SetState(524)
				p.Int_()
			}
			{
				p.SetState(525)
				p.Match(iso8601ParserD)
			}

		}
		{
			p.SetState(529)
			p.Match(iso8601ParserT)
		}
		p.SetState(533)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(530)
				p.Int_()
			}
			{
				p.SetState(531)
				p.Match(iso8601ParserH)
			}

		}
		{
			p.SetState(535)
			p.Dec()
		}
		{
			p.SetState(536)
			p.Match(iso8601ParserM)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(538)
			p.Match(iso8601ParserP)
		}
		p.SetState(542)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(539)
				p.Int_()
			}
			{
				p.SetState(540)
				p.Match(iso8601ParserY)
			}

		}
		p.SetState(547)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(544)
				p.Int_()
			}
			{
				p.SetState(545)
				p.Match(iso8601ParserM)
			}

		}
		p.SetState(552)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == iso8601ParserDigit {
			{
				p.SetState(549)
				p.Int_()
			}
			{
				p.SetState(550)
				p.Match(iso8601ParserD)
			}

		}
		{
			p.SetState(554)
			p.Match(iso8601ParserT)
		}
		{
			p.SetState(555)
			p.Dec()
		}
		{
			p.SetState(556)
			p.Match(iso8601ParserH)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(558)
			p.Match(iso8601ParserP)
		}
		p.SetState(562)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(559)
				p.Int_()
			}
			{
				p.SetState(560)
				p.Match(iso8601ParserY)
			}

		}
		p.SetState(567)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(564)
				p.Int_()
			}
			{
				p.SetState(565)
				p.Match(iso8601ParserM)
			}

		}
		{
			p.SetState(569)
			p.Dec()
		}
		{
			p.SetState(570)
			p.Match(iso8601ParserD)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(572)
			p.Match(iso8601ParserP)
		}
		p.SetState(576)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(573)
				p.Int_()
			}
			{
				p.SetState(574)
				p.Match(iso8601ParserY)
			}

		}
		{
			p.SetState(578)
			p.Dec()
		}
		{
			p.SetState(579)
			p.Match(iso8601ParserM)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(581)
			p.Match(iso8601ParserP)
		}
		{
			p.SetState(582)
			p.Dec()
		}
		{
			p.SetState(583)
			p.Match(iso8601ParserY)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(585)
			p.Match(iso8601ParserP)
		}
		{
			p.SetState(586)
			p.Dec()
		}
		{
			p.SetState(587)
			p.Match(iso8601ParserW)
		}

	}

	return localctx
}

// IMonthDateBasicContext is an interface to support dynamic dispatch.
type IMonthDateBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonthDateBasicContext differentiates from other interfaces.
	IsMonthDateBasicContext()
}

type MonthDateBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonthDateBasicContext() *MonthDateBasicContext {
	var p = new(MonthDateBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_monthDateBasic
	return p
}

func (*MonthDateBasicContext) IsMonthDateBasicContext() {}

func NewMonthDateBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonthDateBasicContext {
	var p = new(MonthDateBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_monthDateBasic

	return p
}

func (s *MonthDateBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *MonthDateBasicContext) Month() IMonthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthContext)
}

func (s *MonthDateBasicContext) Day() IDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *MonthDateBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonthDateBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonthDateBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterMonthDateBasic(s)
	}
}

func (s *MonthDateBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitMonthDateBasic(s)
	}
}

func (p *iso8601Parser) MonthDateBasic() (localctx IMonthDateBasicContext) {
	this := p
	_ = this

	localctx = NewMonthDateBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, iso8601ParserRULE_monthDateBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Month()
	}
	{
		p.SetState(592)
		p.Day()
	}

	return localctx
}

// IMonthDateExtendedContext is an interface to support dynamic dispatch.
type IMonthDateExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonthDateExtendedContext differentiates from other interfaces.
	IsMonthDateExtendedContext()
}

type MonthDateExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonthDateExtendedContext() *MonthDateExtendedContext {
	var p = new(MonthDateExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_monthDateExtended
	return p
}

func (*MonthDateExtendedContext) IsMonthDateExtendedContext() {}

func NewMonthDateExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonthDateExtendedContext {
	var p = new(MonthDateExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_monthDateExtended

	return p
}

func (s *MonthDateExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *MonthDateExtendedContext) Month() IMonthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthContext)
}

func (s *MonthDateExtendedContext) Day() IDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *MonthDateExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonthDateExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonthDateExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterMonthDateExtended(s)
	}
}

func (s *MonthDateExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitMonthDateExtended(s)
	}
}

func (p *iso8601Parser) MonthDateExtended() (localctx IMonthDateExtendedContext) {
	this := p
	_ = this

	localctx = NewMonthDateExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, iso8601ParserRULE_monthDateExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(594)
		p.Month()
	}
	{
		p.SetState(595)
		p.Match(iso8601ParserT__1)
	}
	{
		p.SetState(596)
		p.Day()
	}

	return localctx
}

// IIntervalYMDTimeBasicContext is an interface to support dynamic dispatch.
type IIntervalYMDTimeBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntervalYMDTimeBasicContext differentiates from other interfaces.
	IsIntervalYMDTimeBasicContext()
}

type IntervalYMDTimeBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalYMDTimeBasicContext() *IntervalYMDTimeBasicContext {
	var p = new(IntervalYMDTimeBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_intervalYMDTimeBasic
	return p
}

func (*IntervalYMDTimeBasicContext) IsIntervalYMDTimeBasicContext() {}

func NewIntervalYMDTimeBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalYMDTimeBasicContext {
	var p = new(IntervalYMDTimeBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_intervalYMDTimeBasic

	return p
}

func (s *IntervalYMDTimeBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalYMDTimeBasicContext) MonthDateBasic() IMonthDateBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthDateBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthDateBasicContext)
}

func (s *IntervalYMDTimeBasicContext) Day() IDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *IntervalYMDTimeBasicContext) DatetimeBasic() IDatetimeBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeBasicContext)
}

func (s *IntervalYMDTimeBasicContext) DateBasic() IDateBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateBasicContext)
}

func (s *IntervalYMDTimeBasicContext) LocalTimeBasic() ILocalTimeBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimeBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimeBasicContext)
}

func (s *IntervalYMDTimeBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalYMDTimeBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalYMDTimeBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterIntervalYMDTimeBasic(s)
	}
}

func (s *IntervalYMDTimeBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitIntervalYMDTimeBasic(s)
	}
}

func (p *iso8601Parser) IntervalYMDTimeBasic() (localctx IIntervalYMDTimeBasicContext) {
	this := p
	_ = this

	localctx = NewIntervalYMDTimeBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, iso8601ParserRULE_intervalYMDTimeBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(603)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(598)
			p.MonthDateBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(599)
			p.Day()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(600)
			p.DatetimeBasic()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(601)
			p.DateBasic()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(602)
			p.LocalTimeBasic()
		}

	}

	return localctx
}

// IIntervalYMDTimeExtendedContext is an interface to support dynamic dispatch.
type IIntervalYMDTimeExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntervalYMDTimeExtendedContext differentiates from other interfaces.
	IsIntervalYMDTimeExtendedContext()
}

type IntervalYMDTimeExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalYMDTimeExtendedContext() *IntervalYMDTimeExtendedContext {
	var p = new(IntervalYMDTimeExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_intervalYMDTimeExtended
	return p
}

func (*IntervalYMDTimeExtendedContext) IsIntervalYMDTimeExtendedContext() {}

func NewIntervalYMDTimeExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalYMDTimeExtendedContext {
	var p = new(IntervalYMDTimeExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_intervalYMDTimeExtended

	return p
}

func (s *IntervalYMDTimeExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalYMDTimeExtendedContext) MonthDateExtended() IMonthDateExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthDateExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthDateExtendedContext)
}

func (s *IntervalYMDTimeExtendedContext) Day() IDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *IntervalYMDTimeExtendedContext) DatetimeExtended() IDatetimeExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeExtendedContext)
}

func (s *IntervalYMDTimeExtendedContext) DateExtended() IDateExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateExtendedContext)
}

func (s *IntervalYMDTimeExtendedContext) LocalTimeExtended() ILocalTimeExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILocalTimeExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILocalTimeExtendedContext)
}

func (s *IntervalYMDTimeExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalYMDTimeExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalYMDTimeExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterIntervalYMDTimeExtended(s)
	}
}

func (s *IntervalYMDTimeExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitIntervalYMDTimeExtended(s)
	}
}

func (p *iso8601Parser) IntervalYMDTimeExtended() (localctx IIntervalYMDTimeExtendedContext) {
	this := p
	_ = this

	localctx = NewIntervalYMDTimeExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, iso8601ParserRULE_intervalYMDTimeExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(611)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(605)
			p.MonthDateExtended()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(606)
			p.Day()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(607)
			p.DatetimeExtended()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(608)
			p.DateExtended()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(609)
			p.LocalTimeExtended()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(610)
			p.MonthDateExtended()
		}

	}

	return localctx
}

// IIntervalYMDContext is an interface to support dynamic dispatch.
type IIntervalYMDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntervalYMDContext differentiates from other interfaces.
	IsIntervalYMDContext()
}

type IntervalYMDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalYMDContext() *IntervalYMDContext {
	var p = new(IntervalYMDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_intervalYMD
	return p
}

func (*IntervalYMDContext) IsIntervalYMDContext() {}

func NewIntervalYMDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalYMDContext {
	var p = new(IntervalYMDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_intervalYMD

	return p
}

func (s *IntervalYMDContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalYMDContext) IntervalYMDBasic() IIntervalYMDBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalYMDBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalYMDBasicContext)
}

func (s *IntervalYMDContext) IntervalYMDExtended() IIntervalYMDExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalYMDExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalYMDExtendedContext)
}

func (s *IntervalYMDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalYMDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalYMDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterIntervalYMD(s)
	}
}

func (s *IntervalYMDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitIntervalYMD(s)
	}
}

func (p *iso8601Parser) IntervalYMD() (localctx IIntervalYMDContext) {
	this := p
	_ = this

	localctx = NewIntervalYMDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, iso8601ParserRULE_intervalYMD)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(613)
			p.IntervalYMDBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(614)
			p.IntervalYMDExtended()
		}

	}

	return localctx
}

// IIntervalYMDBasicContext is an interface to support dynamic dispatch.
type IIntervalYMDBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntervalYMDBasicContext differentiates from other interfaces.
	IsIntervalYMDBasicContext()
}

type IntervalYMDBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalYMDBasicContext() *IntervalYMDBasicContext {
	var p = new(IntervalYMDBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_intervalYMDBasic
	return p
}

func (*IntervalYMDBasicContext) IsIntervalYMDBasicContext() {}

func NewIntervalYMDBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalYMDBasicContext {
	var p = new(IntervalYMDBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_intervalYMDBasic

	return p
}

func (s *IntervalYMDBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalYMDBasicContext) P() antlr.TerminalNode {
	return s.GetToken(iso8601ParserP, 0)
}

func (s *IntervalYMDBasicContext) IntervalYMDTimeBasic() IIntervalYMDTimeBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalYMDTimeBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalYMDTimeBasicContext)
}

func (s *IntervalYMDBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalYMDBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalYMDBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterIntervalYMDBasic(s)
	}
}

func (s *IntervalYMDBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitIntervalYMDBasic(s)
	}
}

func (p *iso8601Parser) IntervalYMDBasic() (localctx IIntervalYMDBasicContext) {
	this := p
	_ = this

	localctx = NewIntervalYMDBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, iso8601ParserRULE_intervalYMDBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(iso8601ParserP)
	}
	{
		p.SetState(618)
		p.IntervalYMDTimeBasic()
	}

	return localctx
}

// IIntervalYMDExtendedContext is an interface to support dynamic dispatch.
type IIntervalYMDExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntervalYMDExtendedContext differentiates from other interfaces.
	IsIntervalYMDExtendedContext()
}

type IntervalYMDExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalYMDExtendedContext() *IntervalYMDExtendedContext {
	var p = new(IntervalYMDExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_intervalYMDExtended
	return p
}

func (*IntervalYMDExtendedContext) IsIntervalYMDExtendedContext() {}

func NewIntervalYMDExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalYMDExtendedContext {
	var p = new(IntervalYMDExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_intervalYMDExtended

	return p
}

func (s *IntervalYMDExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalYMDExtendedContext) P() antlr.TerminalNode {
	return s.GetToken(iso8601ParserP, 0)
}

func (s *IntervalYMDExtendedContext) IntervalYMDTimeExtended() IIntervalYMDTimeExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalYMDTimeExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalYMDTimeExtendedContext)
}

func (s *IntervalYMDExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalYMDExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalYMDExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterIntervalYMDExtended(s)
	}
}

func (s *IntervalYMDExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitIntervalYMDExtended(s)
	}
}

func (p *iso8601Parser) IntervalYMDExtended() (localctx IIntervalYMDExtendedContext) {
	this := p
	_ = this

	localctx = NewIntervalYMDExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 152, iso8601ParserRULE_intervalYMDExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(620)
		p.Match(iso8601ParserP)
	}
	{
		p.SetState(621)
		p.IntervalYMDTimeExtended()
	}

	return localctx
}

// ITimeBeginEndContext is an interface to support dynamic dispatch.
type ITimeBeginEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeBeginEndContext differentiates from other interfaces.
	IsTimeBeginEndContext()
}

type TimeBeginEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeBeginEndContext() *TimeBeginEndContext {
	var p = new(TimeBeginEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeBeginEnd
	return p
}

func (*TimeBeginEndContext) IsTimeBeginEndContext() {}

func NewTimeBeginEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeBeginEndContext {
	var p = new(TimeBeginEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeBeginEnd

	return p
}

func (s *TimeBeginEndContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeBeginEndContext) TimeBeginEndBasic() ITimeBeginEndBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeBeginEndBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeBeginEndBasicContext)
}

func (s *TimeBeginEndContext) TimeBeginEndExtended() ITimeBeginEndExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeBeginEndExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeBeginEndExtendedContext)
}

func (s *TimeBeginEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeBeginEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeBeginEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeBeginEnd(s)
	}
}

func (s *TimeBeginEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeBeginEnd(s)
	}
}

func (p *iso8601Parser) TimeBeginEnd() (localctx ITimeBeginEndContext) {
	this := p
	_ = this

	localctx = NewTimeBeginEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 154, iso8601ParserRULE_timeBeginEnd)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(625)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(623)
			p.TimeBeginEndBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(624)
			p.TimeBeginEndExtended()
		}

	}

	return localctx
}

// ITimeBeginEndBasicContext is an interface to support dynamic dispatch.
type ITimeBeginEndBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeBeginEndBasicContext differentiates from other interfaces.
	IsTimeBeginEndBasicContext()
}

type TimeBeginEndBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeBeginEndBasicContext() *TimeBeginEndBasicContext {
	var p = new(TimeBeginEndBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeBeginEndBasic
	return p
}

func (*TimeBeginEndBasicContext) IsTimeBeginEndBasicContext() {}

func NewTimeBeginEndBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeBeginEndBasicContext {
	var p = new(TimeBeginEndBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeBeginEndBasic

	return p
}

func (s *TimeBeginEndBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeBeginEndBasicContext) AllIntervalYMDTimeBasic() []IIntervalYMDTimeBasicContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntervalYMDTimeBasicContext)(nil)).Elem())
	var tst = make([]IIntervalYMDTimeBasicContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntervalYMDTimeBasicContext)
		}
	}

	return tst
}

func (s *TimeBeginEndBasicContext) IntervalYMDTimeBasic(i int) IIntervalYMDTimeBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalYMDTimeBasicContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntervalYMDTimeBasicContext)
}

func (s *TimeBeginEndBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeBeginEndBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeBeginEndBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeBeginEndBasic(s)
	}
}

func (s *TimeBeginEndBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeBeginEndBasic(s)
	}
}

func (p *iso8601Parser) TimeBeginEndBasic() (localctx ITimeBeginEndBasicContext) {
	this := p
	_ = this

	localctx = NewTimeBeginEndBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 156, iso8601ParserRULE_timeBeginEndBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(627)
		p.IntervalYMDTimeBasic()
	}
	{
		p.SetState(628)
		p.Match(iso8601ParserT__3)
	}
	{
		p.SetState(629)
		p.IntervalYMDTimeBasic()
	}

	return localctx
}

// ITimeBeginEndExtendedContext is an interface to support dynamic dispatch.
type ITimeBeginEndExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeBeginEndExtendedContext differentiates from other interfaces.
	IsTimeBeginEndExtendedContext()
}

type TimeBeginEndExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeBeginEndExtendedContext() *TimeBeginEndExtendedContext {
	var p = new(TimeBeginEndExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeBeginEndExtended
	return p
}

func (*TimeBeginEndExtendedContext) IsTimeBeginEndExtendedContext() {}

func NewTimeBeginEndExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeBeginEndExtendedContext {
	var p = new(TimeBeginEndExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeBeginEndExtended

	return p
}

func (s *TimeBeginEndExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeBeginEndExtendedContext) AllIntervalYMDTimeExtended() []IIntervalYMDTimeExtendedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIntervalYMDTimeExtendedContext)(nil)).Elem())
	var tst = make([]IIntervalYMDTimeExtendedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIntervalYMDTimeExtendedContext)
		}
	}

	return tst
}

func (s *TimeBeginEndExtendedContext) IntervalYMDTimeExtended(i int) IIntervalYMDTimeExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalYMDTimeExtendedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIntervalYMDTimeExtendedContext)
}

func (s *TimeBeginEndExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeBeginEndExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeBeginEndExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeBeginEndExtended(s)
	}
}

func (s *TimeBeginEndExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeBeginEndExtended(s)
	}
}

func (p *iso8601Parser) TimeBeginEndExtended() (localctx ITimeBeginEndExtendedContext) {
	this := p
	_ = this

	localctx = NewTimeBeginEndExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 158, iso8601ParserRULE_timeBeginEndExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(631)
		p.IntervalYMDTimeExtended()
	}
	{
		p.SetState(632)
		p.Match(iso8601ParserT__3)
	}
	{
		p.SetState(633)
		p.IntervalYMDTimeExtended()
	}

	return localctx
}

// ITimeBeginIntervalContext is an interface to support dynamic dispatch.
type ITimeBeginIntervalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeBeginIntervalContext differentiates from other interfaces.
	IsTimeBeginIntervalContext()
}

type TimeBeginIntervalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeBeginIntervalContext() *TimeBeginIntervalContext {
	var p = new(TimeBeginIntervalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeBeginInterval
	return p
}

func (*TimeBeginIntervalContext) IsTimeBeginIntervalContext() {}

func NewTimeBeginIntervalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeBeginIntervalContext {
	var p = new(TimeBeginIntervalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeBeginInterval

	return p
}

func (s *TimeBeginIntervalContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeBeginIntervalContext) TimeBeginIntervalBasic() ITimeBeginIntervalBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeBeginIntervalBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeBeginIntervalBasicContext)
}

func (s *TimeBeginIntervalContext) TimeBeginIntervalExtended() ITimeBeginIntervalExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeBeginIntervalExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeBeginIntervalExtendedContext)
}

func (s *TimeBeginIntervalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeBeginIntervalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeBeginIntervalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeBeginInterval(s)
	}
}

func (s *TimeBeginIntervalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeBeginInterval(s)
	}
}

func (p *iso8601Parser) TimeBeginInterval() (localctx ITimeBeginIntervalContext) {
	this := p
	_ = this

	localctx = NewTimeBeginIntervalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 160, iso8601ParserRULE_timeBeginInterval)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(637)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 60, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(635)
			p.TimeBeginIntervalBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(636)
			p.TimeBeginIntervalExtended()
		}

	}

	return localctx
}

// ITimeBeginIntervalBasicContext is an interface to support dynamic dispatch.
type ITimeBeginIntervalBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeBeginIntervalBasicContext differentiates from other interfaces.
	IsTimeBeginIntervalBasicContext()
}

type TimeBeginIntervalBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeBeginIntervalBasicContext() *TimeBeginIntervalBasicContext {
	var p = new(TimeBeginIntervalBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeBeginIntervalBasic
	return p
}

func (*TimeBeginIntervalBasicContext) IsTimeBeginIntervalBasicContext() {}

func NewTimeBeginIntervalBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeBeginIntervalBasicContext {
	var p = new(TimeBeginIntervalBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeBeginIntervalBasic

	return p
}

func (s *TimeBeginIntervalBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeBeginIntervalBasicContext) DatetimeBasic() IDatetimeBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeBasicContext)
}

func (s *TimeBeginIntervalBasicContext) IntervalBasic() IIntervalBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalBasicContext)
}

func (s *TimeBeginIntervalBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeBeginIntervalBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeBeginIntervalBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeBeginIntervalBasic(s)
	}
}

func (s *TimeBeginIntervalBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeBeginIntervalBasic(s)
	}
}

func (p *iso8601Parser) TimeBeginIntervalBasic() (localctx ITimeBeginIntervalBasicContext) {
	this := p
	_ = this

	localctx = NewTimeBeginIntervalBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 162, iso8601ParserRULE_timeBeginIntervalBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(639)
		p.DatetimeBasic()
	}
	{
		p.SetState(640)
		p.Match(iso8601ParserT__3)
	}
	{
		p.SetState(641)
		p.IntervalBasic()
	}

	return localctx
}

// ITimeBeginIntervalExtendedContext is an interface to support dynamic dispatch.
type ITimeBeginIntervalExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeBeginIntervalExtendedContext differentiates from other interfaces.
	IsTimeBeginIntervalExtendedContext()
}

type TimeBeginIntervalExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeBeginIntervalExtendedContext() *TimeBeginIntervalExtendedContext {
	var p = new(TimeBeginIntervalExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeBeginIntervalExtended
	return p
}

func (*TimeBeginIntervalExtendedContext) IsTimeBeginIntervalExtendedContext() {}

func NewTimeBeginIntervalExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeBeginIntervalExtendedContext {
	var p = new(TimeBeginIntervalExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeBeginIntervalExtended

	return p
}

func (s *TimeBeginIntervalExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeBeginIntervalExtendedContext) DatetimeExtended() IDatetimeExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeExtendedContext)
}

func (s *TimeBeginIntervalExtendedContext) IntervalExtended() IIntervalExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalExtendedContext)
}

func (s *TimeBeginIntervalExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeBeginIntervalExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeBeginIntervalExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeBeginIntervalExtended(s)
	}
}

func (s *TimeBeginIntervalExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeBeginIntervalExtended(s)
	}
}

func (p *iso8601Parser) TimeBeginIntervalExtended() (localctx ITimeBeginIntervalExtendedContext) {
	this := p
	_ = this

	localctx = NewTimeBeginIntervalExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 164, iso8601ParserRULE_timeBeginIntervalExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(643)
		p.DatetimeExtended()
	}
	{
		p.SetState(644)
		p.Match(iso8601ParserT__3)
	}
	{
		p.SetState(645)
		p.IntervalExtended()
	}

	return localctx
}

// ITimeIntervalEndContext is an interface to support dynamic dispatch.
type ITimeIntervalEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeIntervalEndContext differentiates from other interfaces.
	IsTimeIntervalEndContext()
}

type TimeIntervalEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeIntervalEndContext() *TimeIntervalEndContext {
	var p = new(TimeIntervalEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeIntervalEnd
	return p
}

func (*TimeIntervalEndContext) IsTimeIntervalEndContext() {}

func NewTimeIntervalEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeIntervalEndContext {
	var p = new(TimeIntervalEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeIntervalEnd

	return p
}

func (s *TimeIntervalEndContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeIntervalEndContext) TimeIntervalEndBasic() ITimeIntervalEndBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeIntervalEndBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeIntervalEndBasicContext)
}

func (s *TimeIntervalEndContext) TimeIntervalEndExtended() ITimeIntervalEndExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeIntervalEndExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeIntervalEndExtendedContext)
}

func (s *TimeIntervalEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeIntervalEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeIntervalEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeIntervalEnd(s)
	}
}

func (s *TimeIntervalEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeIntervalEnd(s)
	}
}

func (p *iso8601Parser) TimeIntervalEnd() (localctx ITimeIntervalEndContext) {
	this := p
	_ = this

	localctx = NewTimeIntervalEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 166, iso8601ParserRULE_timeIntervalEnd)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(649)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(647)
			p.TimeIntervalEndBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(648)
			p.TimeIntervalEndExtended()
		}

	}

	return localctx
}

// ITimeIntervalEndBasicContext is an interface to support dynamic dispatch.
type ITimeIntervalEndBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeIntervalEndBasicContext differentiates from other interfaces.
	IsTimeIntervalEndBasicContext()
}

type TimeIntervalEndBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeIntervalEndBasicContext() *TimeIntervalEndBasicContext {
	var p = new(TimeIntervalEndBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeIntervalEndBasic
	return p
}

func (*TimeIntervalEndBasicContext) IsTimeIntervalEndBasicContext() {}

func NewTimeIntervalEndBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeIntervalEndBasicContext {
	var p = new(TimeIntervalEndBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeIntervalEndBasic

	return p
}

func (s *TimeIntervalEndBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeIntervalEndBasicContext) IntervalBasic() IIntervalBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalBasicContext)
}

func (s *TimeIntervalEndBasicContext) DatetimeBasic() IDatetimeBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeBasicContext)
}

func (s *TimeIntervalEndBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeIntervalEndBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeIntervalEndBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeIntervalEndBasic(s)
	}
}

func (s *TimeIntervalEndBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeIntervalEndBasic(s)
	}
}

func (p *iso8601Parser) TimeIntervalEndBasic() (localctx ITimeIntervalEndBasicContext) {
	this := p
	_ = this

	localctx = NewTimeIntervalEndBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 168, iso8601ParserRULE_timeIntervalEndBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(651)
		p.IntervalBasic()
	}
	{
		p.SetState(652)
		p.Match(iso8601ParserT__3)
	}
	{
		p.SetState(653)
		p.DatetimeBasic()
	}

	return localctx
}

// ITimeIntervalEndExtendedContext is an interface to support dynamic dispatch.
type ITimeIntervalEndExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeIntervalEndExtendedContext differentiates from other interfaces.
	IsTimeIntervalEndExtendedContext()
}

type TimeIntervalEndExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeIntervalEndExtendedContext() *TimeIntervalEndExtendedContext {
	var p = new(TimeIntervalEndExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_timeIntervalEndExtended
	return p
}

func (*TimeIntervalEndExtendedContext) IsTimeIntervalEndExtendedContext() {}

func NewTimeIntervalEndExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeIntervalEndExtendedContext {
	var p = new(TimeIntervalEndExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_timeIntervalEndExtended

	return p
}

func (s *TimeIntervalEndExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeIntervalEndExtendedContext) IntervalExtended() IIntervalExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalExtendedContext)
}

func (s *TimeIntervalEndExtendedContext) DatetimeExtended() IDatetimeExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeExtendedContext)
}

func (s *TimeIntervalEndExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeIntervalEndExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeIntervalEndExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterTimeIntervalEndExtended(s)
	}
}

func (s *TimeIntervalEndExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitTimeIntervalEndExtended(s)
	}
}

func (p *iso8601Parser) TimeIntervalEndExtended() (localctx ITimeIntervalEndExtendedContext) {
	this := p
	_ = this

	localctx = NewTimeIntervalEndExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 170, iso8601ParserRULE_timeIntervalEndExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(655)
		p.IntervalExtended()
	}
	{
		p.SetState(656)
		p.Match(iso8601ParserT__3)
	}
	{
		p.SetState(657)
		p.DatetimeExtended()
	}

	return localctx
}

// IDurationContext is an interface to support dynamic dispatch.
type IDurationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDurationContext differentiates from other interfaces.
	IsDurationContext()
}

type DurationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationContext() *DurationContext {
	var p = new(DurationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_duration
	return p
}

func (*DurationContext) IsDurationContext() {}

func NewDurationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationContext {
	var p = new(DurationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_duration

	return p
}

func (s *DurationContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationContext) DurationBasic() IDurationBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDurationBasicContext)
}

func (s *DurationContext) DurationExtended() IDurationExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDurationExtendedContext)
}

func (s *DurationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDuration(s)
	}
}

func (s *DurationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDuration(s)
	}
}

func (p *iso8601Parser) Duration() (localctx IDurationContext) {
	this := p
	_ = this

	localctx = NewDurationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 172, iso8601ParserRULE_duration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(661)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(659)
			p.DurationBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(660)
			p.DurationExtended()
		}

	}

	return localctx
}

// IDurationBasicContext is an interface to support dynamic dispatch.
type IDurationBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDurationBasicContext differentiates from other interfaces.
	IsDurationBasicContext()
}

type DurationBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationBasicContext() *DurationBasicContext {
	var p = new(DurationBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_durationBasic
	return p
}

func (*DurationBasicContext) IsDurationBasicContext() {}

func NewDurationBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationBasicContext {
	var p = new(DurationBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_durationBasic

	return p
}

func (s *DurationBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationBasicContext) TimeBeginEndBasic() ITimeBeginEndBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeBeginEndBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeBeginEndBasicContext)
}

func (s *DurationBasicContext) TimeBeginIntervalBasic() ITimeBeginIntervalBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeBeginIntervalBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeBeginIntervalBasicContext)
}

func (s *DurationBasicContext) TimeIntervalEndBasic() ITimeIntervalEndBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeIntervalEndBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeIntervalEndBasicContext)
}

func (s *DurationBasicContext) IntervalBasic() IIntervalBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalBasicContext)
}

func (s *DurationBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDurationBasic(s)
	}
}

func (s *DurationBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDurationBasic(s)
	}
}

func (p *iso8601Parser) DurationBasic() (localctx IDurationBasicContext) {
	this := p
	_ = this

	localctx = NewDurationBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 174, iso8601ParserRULE_durationBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(667)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(663)
			p.TimeBeginEndBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(664)
			p.TimeBeginIntervalBasic()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(665)
			p.TimeIntervalEndBasic()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(666)
			p.IntervalBasic()
		}

	}

	return localctx
}

// IDurationExtendedContext is an interface to support dynamic dispatch.
type IDurationExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDurationExtendedContext differentiates from other interfaces.
	IsDurationExtendedContext()
}

type DurationExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDurationExtendedContext() *DurationExtendedContext {
	var p = new(DurationExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_durationExtended
	return p
}

func (*DurationExtendedContext) IsDurationExtendedContext() {}

func NewDurationExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DurationExtendedContext {
	var p = new(DurationExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_durationExtended

	return p
}

func (s *DurationExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *DurationExtendedContext) TimeBeginEndExtended() ITimeBeginEndExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeBeginEndExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeBeginEndExtendedContext)
}

func (s *DurationExtendedContext) TimeBeginIntervalExtended() ITimeBeginIntervalExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeBeginIntervalExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeBeginIntervalExtendedContext)
}

func (s *DurationExtendedContext) TimeIntervalEndExtended() ITimeIntervalEndExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeIntervalEndExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeIntervalEndExtendedContext)
}

func (s *DurationExtendedContext) IntervalExtended() IIntervalExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntervalExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntervalExtendedContext)
}

func (s *DurationExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DurationExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DurationExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterDurationExtended(s)
	}
}

func (s *DurationExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitDurationExtended(s)
	}
}

func (p *iso8601Parser) DurationExtended() (localctx IDurationExtendedContext) {
	this := p
	_ = this

	localctx = NewDurationExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 176, iso8601ParserRULE_durationExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(673)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(669)
			p.TimeBeginEndExtended()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(670)
			p.TimeBeginIntervalExtended()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(671)
			p.TimeIntervalEndExtended()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(672)
			p.IntervalExtended()
		}

	}

	return localctx
}

// IRecurringCountContext is an interface to support dynamic dispatch.
type IRecurringCountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecurringCountContext differentiates from other interfaces.
	IsRecurringCountContext()
}

type RecurringCountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecurringCountContext() *RecurringCountContext {
	var p = new(RecurringCountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_recurringCount
	return p
}

func (*RecurringCountContext) IsRecurringCountContext() {}

func NewRecurringCountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecurringCountContext {
	var p = new(RecurringCountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_recurringCount

	return p
}

func (s *RecurringCountContext) GetParser() antlr.Parser { return s.parser }

func (s *RecurringCountContext) R() antlr.TerminalNode {
	return s.GetToken(iso8601ParserR, 0)
}

func (s *RecurringCountContext) Int_() IInt_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInt_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInt_Context)
}

func (s *RecurringCountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecurringCountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecurringCountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterRecurringCount(s)
	}
}

func (s *RecurringCountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitRecurringCount(s)
	}
}

func (p *iso8601Parser) RecurringCount() (localctx IRecurringCountContext) {
	this := p
	_ = this

	localctx = NewRecurringCountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 178, iso8601ParserRULE_recurringCount)
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
		p.SetState(675)
		p.Match(iso8601ParserR)
	}
	p.SetState(677)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == iso8601ParserDigit {
		{
			p.SetState(676)
			p.Int_()
		}

	}

	return localctx
}

// IRecurringContext is an interface to support dynamic dispatch.
type IRecurringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecurringContext differentiates from other interfaces.
	IsRecurringContext()
}

type RecurringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecurringContext() *RecurringContext {
	var p = new(RecurringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_recurring
	return p
}

func (*RecurringContext) IsRecurringContext() {}

func NewRecurringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecurringContext {
	var p = new(RecurringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_recurring

	return p
}

func (s *RecurringContext) GetParser() antlr.Parser { return s.parser }

func (s *RecurringContext) RecurringBasic() IRecurringBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecurringBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecurringBasicContext)
}

func (s *RecurringContext) RecurringExtended() IRecurringExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecurringExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecurringExtendedContext)
}

func (s *RecurringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecurringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecurringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterRecurring(s)
	}
}

func (s *RecurringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitRecurring(s)
	}
}

func (p *iso8601Parser) Recurring() (localctx IRecurringContext) {
	this := p
	_ = this

	localctx = NewRecurringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 180, iso8601ParserRULE_recurring)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(681)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(679)
			p.RecurringBasic()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(680)
			p.RecurringExtended()
		}

	}

	return localctx
}

// IRecurringBasicContext is an interface to support dynamic dispatch.
type IRecurringBasicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecurringBasicContext differentiates from other interfaces.
	IsRecurringBasicContext()
}

type RecurringBasicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecurringBasicContext() *RecurringBasicContext {
	var p = new(RecurringBasicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_recurringBasic
	return p
}

func (*RecurringBasicContext) IsRecurringBasicContext() {}

func NewRecurringBasicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecurringBasicContext {
	var p = new(RecurringBasicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_recurringBasic

	return p
}

func (s *RecurringBasicContext) GetParser() antlr.Parser { return s.parser }

func (s *RecurringBasicContext) RecurringCount() IRecurringCountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecurringCountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecurringCountContext)
}

func (s *RecurringBasicContext) DurationBasic() IDurationBasicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationBasicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDurationBasicContext)
}

func (s *RecurringBasicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecurringBasicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecurringBasicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterRecurringBasic(s)
	}
}

func (s *RecurringBasicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitRecurringBasic(s)
	}
}

func (p *iso8601Parser) RecurringBasic() (localctx IRecurringBasicContext) {
	this := p
	_ = this

	localctx = NewRecurringBasicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 182, iso8601ParserRULE_recurringBasic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(683)
		p.RecurringCount()
	}
	{
		p.SetState(684)
		p.Match(iso8601ParserT__3)
	}
	{
		p.SetState(685)
		p.DurationBasic()
	}

	return localctx
}

// IRecurringExtendedContext is an interface to support dynamic dispatch.
type IRecurringExtendedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecurringExtendedContext differentiates from other interfaces.
	IsRecurringExtendedContext()
}

type RecurringExtendedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecurringExtendedContext() *RecurringExtendedContext {
	var p = new(RecurringExtendedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_recurringExtended
	return p
}

func (*RecurringExtendedContext) IsRecurringExtendedContext() {}

func NewRecurringExtendedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RecurringExtendedContext {
	var p = new(RecurringExtendedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_recurringExtended

	return p
}

func (s *RecurringExtendedContext) GetParser() antlr.Parser { return s.parser }

func (s *RecurringExtendedContext) RecurringCount() IRecurringCountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecurringCountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecurringCountContext)
}

func (s *RecurringExtendedContext) DurationExtended() IDurationExtendedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationExtendedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDurationExtendedContext)
}

func (s *RecurringExtendedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RecurringExtendedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RecurringExtendedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterRecurringExtended(s)
	}
}

func (s *RecurringExtendedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitRecurringExtended(s)
	}
}

func (p *iso8601Parser) RecurringExtended() (localctx IRecurringExtendedContext) {
	this := p
	_ = this

	localctx = NewRecurringExtendedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 184, iso8601ParserRULE_recurringExtended)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(687)
		p.RecurringCount()
	}
	{
		p.SetState(688)
		p.Match(iso8601ParserT__3)
	}
	{
		p.SetState(689)
		p.DurationExtended()
	}

	return localctx
}

// IIsoContext is an interface to support dynamic dispatch.
type IIsoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsoContext differentiates from other interfaces.
	IsIsoContext()
}

type IsoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsoContext() *IsoContext {
	var p = new(IsoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = iso8601ParserRULE_iso
	return p
}

func (*IsoContext) IsIsoContext() {}

func NewIsoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsoContext {
	var p = new(IsoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = iso8601ParserRULE_iso

	return p
}

func (s *IsoContext) GetParser() antlr.Parser { return s.parser }

func (s *IsoContext) EOF() antlr.TerminalNode {
	return s.GetToken(iso8601ParserEOF, 0)
}

func (s *IsoContext) Time() ITimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeContext)
}

func (s *IsoContext) Date() IDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *IsoContext) Datetime() IDatetimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetimeContext)
}

func (s *IsoContext) Duration() IDurationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDurationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDurationContext)
}

func (s *IsoContext) Recurring() IRecurringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecurringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecurringContext)
}

func (s *IsoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.EnterIso(s)
	}
}

func (s *IsoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(iso8601Listener); ok {
		listenerT.ExitIso(s)
	}
}

func (p *iso8601Parser) Iso() (localctx IIsoContext) {
	this := p
	_ = this

	localctx = NewIsoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 186, iso8601ParserRULE_iso)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	p.SetState(696)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(691)
			p.Time()
		}

	case 2:
		{
			p.SetState(692)
			p.Date()
		}

	case 3:
		{
			p.SetState(693)
			p.Datetime()
		}

	case 4:
		{
			p.SetState(694)
			p.Duration()
		}

	case 5:
		{
			p.SetState(695)
			p.Recurring()
		}

	}
	{
		p.SetState(698)
		p.Match(iso8601ParserEOF)
	}

	return localctx
}
