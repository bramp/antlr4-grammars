// Code generated from PowerQueryParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package powerquery // PowerQueryParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 73, 762,
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
	92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4, 96, 9, 96, 4, 97,
	9, 97, 4, 98, 9, 98, 4, 99, 9, 99, 4, 100, 9, 100, 4, 101, 9, 101, 4, 102,
	9, 102, 3, 2, 3, 2, 5, 2, 207, 10, 2, 3, 3, 3, 3, 3, 4, 5, 4, 212, 10,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 218, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 5,
	6, 224, 10, 6, 3, 7, 5, 7, 227, 10, 7, 3, 7, 5, 7, 230, 10, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 249, 10, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 5, 11, 256, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	7, 12, 264, 10, 12, 12, 12, 14, 12, 267, 11, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 7, 13, 275, 10, 13, 12, 13, 14, 13, 278, 11, 13, 3, 14,
	5, 14, 281, 10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 7, 15, 291, 10, 15, 12, 15, 14, 15, 294, 11, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 305, 10, 16, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 324, 10, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	5, 18, 339, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 5, 19, 350, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20,
	357, 10, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 366,
	10, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 5, 22, 378, 10, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 5, 22, 400, 10, 22, 3, 22, 7, 22, 403, 10, 22, 12,
	22, 14, 22, 406, 11, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 5, 25,
	414, 10, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 5, 31, 436, 10, 31, 3, 32, 3, 32, 5, 32, 440, 10, 32, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 449, 10, 33, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 5, 34, 456, 10, 34, 3, 35, 3, 35, 5, 35, 460, 10,
	35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 469, 10, 36,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 5,
	40, 481, 10, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3,
	45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 5, 46, 508, 10, 46, 3, 47,
	3, 47, 5, 47, 512, 10, 47, 3, 48, 3, 48, 5, 48, 516, 10, 48, 3, 48, 3,
	48, 5, 48, 520, 10, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50,
	3, 50, 3, 50, 3, 50, 5, 50, 532, 10, 50, 3, 51, 3, 51, 3, 51, 3, 51, 3,
	51, 5, 51, 539, 10, 51, 3, 52, 3, 52, 5, 52, 543, 10, 52, 3, 53, 3, 53,
	3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3,
	57, 3, 57, 5, 57, 559, 10, 57, 3, 58, 3, 58, 3, 58, 3, 59, 3, 59, 3, 59,
	3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3,
	62, 3, 62, 5, 62, 579, 10, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64,
	3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3,
	67, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 5, 69, 603, 10, 69, 3, 70, 3, 70,
	5, 70, 607, 10, 70, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 3, 71, 5, 71, 615,
	10, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 5, 73,
	625, 10, 73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 3, 73, 5, 73, 634,
	10, 73, 3, 74, 3, 74, 3, 74, 3, 74, 3, 74, 5, 74, 641, 10, 74, 3, 75, 5,
	75, 644, 10, 75, 3, 75, 3, 75, 5, 75, 648, 10, 75, 3, 76, 3, 76, 3, 76,
	3, 77, 3, 77, 3, 78, 3, 78, 3, 79, 3, 79, 3, 79, 3, 79, 3, 80, 3, 80, 3,
	81, 3, 81, 5, 81, 665, 10, 81, 3, 81, 3, 81, 3, 81, 3, 82, 3, 82, 3, 82,
	3, 82, 3, 82, 3, 82, 5, 82, 676, 10, 82, 3, 83, 3, 83, 3, 83, 3, 83, 3,
	83, 5, 83, 683, 10, 83, 3, 84, 3, 84, 3, 85, 3, 85, 3, 85, 3, 85, 3, 85,
	5, 85, 692, 10, 85, 3, 86, 3, 86, 3, 86, 3, 87, 3, 87, 3, 87, 3, 88, 3,
	88, 3, 88, 3, 89, 3, 89, 3, 89, 3, 89, 3, 90, 3, 90, 3, 90, 3, 91, 3, 91,
	3, 91, 3, 92, 3, 92, 3, 92, 5, 92, 716, 10, 92, 3, 93, 3, 93, 3, 94, 3,
	94, 3, 94, 3, 95, 3, 95, 3, 96, 3, 96, 3, 97, 3, 97, 5, 97, 729, 10, 97,
	3, 97, 3, 97, 3, 98, 3, 98, 3, 98, 3, 98, 3, 98, 5, 98, 738, 10, 98, 3,
	99, 3, 99, 3, 99, 3, 99, 3, 100, 3, 100, 5, 100, 746, 10, 100, 3, 100,
	3, 100, 3, 101, 3, 101, 3, 101, 3, 101, 3, 101, 5, 101, 755, 10, 101, 3,
	102, 3, 102, 3, 102, 5, 102, 760, 10, 102, 3, 102, 2, 6, 22, 24, 28, 42,
	103, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
	74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106,
	108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136,
	138, 140, 142, 144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166,
	168, 170, 172, 174, 176, 178, 180, 182, 184, 186, 188, 190, 192, 194, 196,
	198, 200, 202, 2, 3, 6, 2, 17, 17, 29, 29, 36, 49, 68, 68, 2, 750, 2, 206,
	3, 2, 2, 2, 4, 208, 3, 2, 2, 2, 6, 211, 3, 2, 2, 2, 8, 219, 3, 2, 2, 2,
	10, 221, 3, 2, 2, 2, 12, 226, 3, 2, 2, 2, 14, 236, 3, 2, 2, 2, 16, 238,
	3, 2, 2, 2, 18, 248, 3, 2, 2, 2, 20, 255, 3, 2, 2, 2, 22, 257, 3, 2, 2,
	2, 24, 268, 3, 2, 2, 2, 26, 280, 3, 2, 2, 2, 28, 284, 3, 2, 2, 2, 30, 304,
	3, 2, 2, 2, 32, 323, 3, 2, 2, 2, 34, 338, 3, 2, 2, 2, 36, 349, 3, 2, 2,
	2, 38, 356, 3, 2, 2, 2, 40, 365, 3, 2, 2, 2, 42, 377, 3, 2, 2, 2, 44, 407,
	3, 2, 2, 2, 46, 409, 3, 2, 2, 2, 48, 413, 3, 2, 2, 2, 50, 415, 3, 2, 2,
	2, 52, 417, 3, 2, 2, 2, 54, 420, 3, 2, 2, 2, 56, 424, 3, 2, 2, 2, 58, 428,
	3, 2, 2, 2, 60, 435, 3, 2, 2, 2, 62, 437, 3, 2, 2, 2, 64, 448, 3, 2, 2,
	2, 66, 455, 3, 2, 2, 2, 68, 457, 3, 2, 2, 2, 70, 468, 3, 2, 2, 2, 72, 470,
	3, 2, 2, 2, 74, 474, 3, 2, 2, 2, 76, 476, 3, 2, 2, 2, 78, 480, 3, 2, 2,
	2, 80, 482, 3, 2, 2, 2, 82, 486, 3, 2, 2, 2, 84, 491, 3, 2, 2, 2, 86, 493,
	3, 2, 2, 2, 88, 497, 3, 2, 2, 2, 90, 507, 3, 2, 2, 2, 92, 511, 3, 2, 2,
	2, 94, 513, 3, 2, 2, 2, 96, 524, 3, 2, 2, 2, 98, 531, 3, 2, 2, 2, 100,
	538, 3, 2, 2, 2, 102, 540, 3, 2, 2, 2, 104, 544, 3, 2, 2, 2, 106, 546,
	3, 2, 2, 2, 108, 548, 3, 2, 2, 2, 110, 550, 3, 2, 2, 2, 112, 558, 3, 2,
	2, 2, 114, 560, 3, 2, 2, 2, 116, 563, 3, 2, 2, 2, 118, 566, 3, 2, 2, 2,
	120, 568, 3, 2, 2, 2, 122, 578, 3, 2, 2, 2, 124, 580, 3, 2, 2, 2, 126,
	584, 3, 2, 2, 2, 128, 586, 3, 2, 2, 2, 130, 593, 3, 2, 2, 2, 132, 595,
	3, 2, 2, 2, 134, 597, 3, 2, 2, 2, 136, 602, 3, 2, 2, 2, 138, 606, 3, 2,
	2, 2, 140, 614, 3, 2, 2, 2, 142, 616, 3, 2, 2, 2, 144, 633, 3, 2, 2, 2,
	146, 640, 3, 2, 2, 2, 148, 643, 3, 2, 2, 2, 150, 649, 3, 2, 2, 2, 152,
	652, 3, 2, 2, 2, 154, 654, 3, 2, 2, 2, 156, 656, 3, 2, 2, 2, 158, 660,
	3, 2, 2, 2, 160, 662, 3, 2, 2, 2, 162, 675, 3, 2, 2, 2, 164, 682, 3, 2,
	2, 2, 166, 684, 3, 2, 2, 2, 168, 691, 3, 2, 2, 2, 170, 693, 3, 2, 2, 2,
	172, 696, 3, 2, 2, 2, 174, 699, 3, 2, 2, 2, 176, 702, 3, 2, 2, 2, 178,
	706, 3, 2, 2, 2, 180, 709, 3, 2, 2, 2, 182, 712, 3, 2, 2, 2, 184, 717,
	3, 2, 2, 2, 186, 719, 3, 2, 2, 2, 188, 722, 3, 2, 2, 2, 190, 724, 3, 2,
	2, 2, 192, 726, 3, 2, 2, 2, 194, 737, 3, 2, 2, 2, 196, 739, 3, 2, 2, 2,
	198, 743, 3, 2, 2, 2, 200, 754, 3, 2, 2, 2, 202, 759, 3, 2, 2, 2, 204,
	207, 5, 4, 3, 2, 205, 207, 5, 16, 9, 2, 206, 204, 3, 2, 2, 2, 206, 205,
	3, 2, 2, 2, 207, 3, 3, 2, 2, 2, 208, 209, 5, 6, 4, 2, 209, 5, 3, 2, 2,
	2, 210, 212, 5, 190, 96, 2, 211, 210, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2,
	212, 213, 3, 2, 2, 2, 213, 214, 7, 20, 2, 2, 214, 215, 5, 8, 5, 2, 215,
	217, 7, 19, 2, 2, 216, 218, 5, 10, 6, 2, 217, 216, 3, 2, 2, 2, 217, 218,
	3, 2, 2, 2, 218, 7, 3, 2, 2, 2, 219, 220, 7, 70, 2, 2, 220, 9, 3, 2, 2,
	2, 221, 223, 5, 12, 7, 2, 222, 224, 5, 10, 6, 2, 223, 222, 3, 2, 2, 2,
	223, 224, 3, 2, 2, 2, 224, 11, 3, 2, 2, 2, 225, 227, 5, 190, 96, 2, 226,
	225, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 229, 3, 2, 2, 2, 228, 230,
	7, 21, 2, 2, 229, 228, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 231, 3, 2,
	2, 2, 231, 232, 5, 14, 8, 2, 232, 233, 7, 7, 2, 2, 233, 234, 5, 18, 10,
	2, 234, 235, 7, 19, 2, 2, 235, 13, 3, 2, 2, 2, 236, 237, 7, 70, 2, 2, 237,
	15, 3, 2, 2, 2, 238, 239, 5, 18, 10, 2, 239, 17, 3, 2, 2, 2, 240, 249,
	5, 20, 11, 2, 241, 249, 5, 116, 59, 2, 242, 249, 5, 94, 48, 2, 243, 249,
	5, 120, 61, 2, 244, 249, 5, 128, 65, 2, 245, 249, 5, 120, 61, 2, 246, 249,
	5, 180, 91, 2, 247, 249, 5, 182, 92, 2, 248, 240, 3, 2, 2, 2, 248, 241,
	3, 2, 2, 2, 248, 242, 3, 2, 2, 2, 248, 243, 3, 2, 2, 2, 248, 244, 3, 2,
	2, 2, 248, 245, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 248, 247, 3, 2, 2, 2,
	249, 19, 3, 2, 2, 2, 250, 256, 5, 22, 12, 2, 251, 252, 5, 22, 12, 2, 252,
	253, 7, 23, 2, 2, 253, 254, 5, 20, 11, 2, 254, 256, 3, 2, 2, 2, 255, 250,
	3, 2, 2, 2, 255, 251, 3, 2, 2, 2, 256, 21, 3, 2, 2, 2, 257, 258, 8, 12,
	1, 2, 258, 259, 5, 24, 13, 2, 259, 265, 3, 2, 2, 2, 260, 261, 12, 3, 2,
	2, 261, 262, 7, 22, 2, 2, 262, 264, 5, 24, 13, 2, 263, 260, 3, 2, 2, 2,
	264, 267, 3, 2, 2, 2, 265, 263, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266,
	23, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 268, 269, 8, 13, 1, 2, 269, 270,
	5, 28, 15, 2, 270, 276, 3, 2, 2, 2, 271, 272, 12, 3, 2, 2, 272, 273, 7,
	59, 2, 2, 273, 275, 5, 26, 14, 2, 274, 271, 3, 2, 2, 2, 275, 278, 3, 2,
	2, 2, 276, 274, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 25, 3, 2, 2, 2,
	278, 276, 3, 2, 2, 2, 279, 281, 7, 18, 2, 2, 280, 279, 3, 2, 2, 2, 280,
	281, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 283, 5, 142, 72, 2, 283, 27,
	3, 2, 2, 2, 284, 285, 8, 15, 1, 2, 285, 286, 5, 30, 16, 2, 286, 292, 3,
	2, 2, 2, 287, 288, 12, 3, 2, 2, 288, 289, 7, 51, 2, 2, 289, 291, 5, 26,
	14, 2, 290, 287, 3, 2, 2, 2, 291, 294, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2,
	292, 293, 3, 2, 2, 2, 293, 29, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2, 295, 305,
	5, 32, 17, 2, 296, 297, 5, 32, 17, 2, 297, 298, 7, 7, 2, 2, 298, 299, 5,
	30, 16, 2, 299, 305, 3, 2, 2, 2, 300, 301, 5, 32, 17, 2, 301, 302, 7, 60,
	2, 2, 302, 303, 5, 30, 16, 2, 303, 305, 3, 2, 2, 2, 304, 295, 3, 2, 2,
	2, 304, 296, 3, 2, 2, 2, 304, 300, 3, 2, 2, 2, 305, 31, 3, 2, 2, 2, 306,
	324, 5, 34, 18, 2, 307, 308, 5, 34, 18, 2, 308, 309, 7, 62, 2, 2, 309,
	310, 5, 32, 17, 2, 310, 324, 3, 2, 2, 2, 311, 312, 5, 34, 18, 2, 312, 313,
	7, 61, 2, 2, 313, 314, 5, 32, 17, 2, 314, 324, 3, 2, 2, 2, 315, 316, 5,
	34, 18, 2, 316, 317, 7, 66, 2, 2, 317, 318, 5, 34, 18, 2, 318, 324, 3,
	2, 2, 2, 319, 320, 5, 34, 18, 2, 320, 321, 7, 67, 2, 2, 321, 322, 5, 32,
	17, 2, 322, 324, 3, 2, 2, 2, 323, 306, 3, 2, 2, 2, 323, 307, 3, 2, 2, 2,
	323, 311, 3, 2, 2, 2, 323, 315, 3, 2, 2, 2, 323, 319, 3, 2, 2, 2, 324,
	33, 3, 2, 2, 2, 325, 339, 5, 36, 19, 2, 326, 327, 5, 36, 19, 2, 327, 328,
	7, 56, 2, 2, 328, 329, 5, 34, 18, 2, 329, 339, 3, 2, 2, 2, 330, 331, 5,
	36, 19, 2, 331, 332, 7, 57, 2, 2, 332, 333, 5, 34, 18, 2, 333, 339, 3,
	2, 2, 2, 334, 335, 5, 36, 19, 2, 335, 336, 7, 65, 2, 2, 336, 337, 5, 34,
	18, 2, 337, 339, 3, 2, 2, 2, 338, 325, 3, 2, 2, 2, 338, 326, 3, 2, 2, 2,
	338, 330, 3, 2, 2, 2, 338, 334, 3, 2, 2, 2, 339, 35, 3, 2, 2, 2, 340, 350,
	5, 38, 20, 2, 341, 342, 5, 38, 20, 2, 342, 343, 7, 64, 2, 2, 343, 344,
	5, 36, 19, 2, 344, 350, 3, 2, 2, 2, 345, 346, 5, 38, 20, 2, 346, 347, 7,
	63, 2, 2, 347, 348, 5, 36, 19, 2, 348, 350, 3, 2, 2, 2, 349, 340, 3, 2,
	2, 2, 349, 341, 3, 2, 2, 2, 349, 345, 3, 2, 2, 2, 350, 37, 3, 2, 2, 2,
	351, 357, 5, 40, 21, 2, 352, 353, 5, 40, 21, 2, 353, 354, 7, 58, 2, 2,
	354, 355, 5, 40, 21, 2, 355, 357, 3, 2, 2, 2, 356, 351, 3, 2, 2, 2, 356,
	352, 3, 2, 2, 2, 357, 39, 3, 2, 2, 2, 358, 366, 5, 136, 69, 2, 359, 360,
	7, 56, 2, 2, 360, 366, 5, 40, 21, 2, 361, 362, 7, 57, 2, 2, 362, 366, 5,
	40, 21, 2, 363, 364, 7, 55, 2, 2, 364, 366, 5, 40, 21, 2, 365, 358, 3,
	2, 2, 2, 365, 359, 3, 2, 2, 2, 365, 361, 3, 2, 2, 2, 365, 363, 3, 2, 2,
	2, 366, 41, 3, 2, 2, 2, 367, 368, 8, 22, 1, 2, 368, 378, 5, 44, 23, 2,
	369, 378, 5, 62, 32, 2, 370, 378, 5, 68, 35, 2, 371, 378, 5, 46, 24, 2,
	372, 378, 5, 54, 28, 2, 373, 378, 5, 56, 29, 2, 374, 378, 5, 84, 43, 2,
	375, 378, 5, 92, 47, 2, 376, 378, 5, 58, 30, 2, 377, 367, 3, 2, 2, 2, 377,
	369, 3, 2, 2, 2, 377, 370, 3, 2, 2, 2, 377, 371, 3, 2, 2, 2, 377, 372,
	3, 2, 2, 2, 377, 373, 3, 2, 2, 2, 377, 374, 3, 2, 2, 2, 377, 375, 3, 2,
	2, 2, 377, 376, 3, 2, 2, 2, 378, 404, 3, 2, 2, 2, 379, 380, 12, 11, 2,
	2, 380, 403, 5, 78, 40, 2, 381, 382, 12, 9, 2, 2, 382, 403, 5, 86, 44,
	2, 383, 384, 12, 8, 2, 2, 384, 403, 5, 88, 45, 2, 385, 386, 12, 6, 2, 2,
	386, 387, 7, 11, 2, 2, 387, 388, 5, 76, 39, 2, 388, 389, 7, 12, 2, 2, 389,
	403, 3, 2, 2, 2, 390, 391, 12, 5, 2, 2, 391, 392, 7, 11, 2, 2, 392, 393,
	5, 76, 39, 2, 393, 394, 7, 12, 2, 2, 394, 395, 7, 15, 2, 2, 395, 403, 3,
	2, 2, 2, 396, 397, 12, 4, 2, 2, 397, 399, 7, 13, 2, 2, 398, 400, 5, 60,
	31, 2, 399, 398, 3, 2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 401, 3, 2, 2, 2,
	401, 403, 7, 14, 2, 2, 402, 379, 3, 2, 2, 2, 402, 381, 3, 2, 2, 2, 402,
	383, 3, 2, 2, 2, 402, 385, 3, 2, 2, 2, 402, 390, 3, 2, 2, 2, 402, 396,
	3, 2, 2, 2, 403, 406, 3, 2, 2, 2, 404, 402, 3, 2, 2, 2, 404, 405, 3, 2,
	2, 2, 405, 43, 3, 2, 2, 2, 406, 404, 3, 2, 2, 2, 407, 408, 7, 68, 2, 2,
	408, 45, 3, 2, 2, 2, 409, 410, 5, 48, 25, 2, 410, 47, 3, 2, 2, 2, 411,
	414, 5, 50, 26, 2, 412, 414, 5, 52, 27, 2, 413, 411, 3, 2, 2, 2, 413, 412,
	3, 2, 2, 2, 414, 49, 3, 2, 2, 2, 415, 416, 7, 70, 2, 2, 416, 51, 3, 2,
	2, 2, 417, 418, 7, 50, 2, 2, 418, 419, 7, 70, 2, 2, 419, 53, 3, 2, 2, 2,
	420, 421, 7, 70, 2, 2, 421, 422, 7, 54, 2, 2, 422, 423, 7, 70, 2, 2, 423,
	55, 3, 2, 2, 2, 424, 425, 7, 13, 2, 2, 425, 426, 5, 18, 10, 2, 426, 427,
	7, 14, 2, 2, 427, 57, 3, 2, 2, 2, 428, 429, 7, 28, 2, 2, 429, 59, 3, 2,
	2, 2, 430, 436, 5, 18, 10, 2, 431, 432, 5, 18, 10, 2, 432, 433, 7, 8, 2,
	2, 433, 434, 5, 60, 31, 2, 434, 436, 3, 2, 2, 2, 435, 430, 3, 2, 2, 2,
	435, 431, 3, 2, 2, 2, 436, 61, 3, 2, 2, 2, 437, 439, 7, 11, 2, 2, 438,
	440, 5, 64, 33, 2, 439, 438, 3, 2, 2, 2, 439, 440, 3, 2, 2, 2, 440, 441,
	3, 2, 2, 2, 441, 442, 7, 12, 2, 2, 442, 63, 3, 2, 2, 2, 443, 449, 5, 66,
	34, 2, 444, 445, 5, 66, 34, 2, 445, 446, 7, 8, 2, 2, 446, 447, 5, 64, 33,
	2, 447, 449, 3, 2, 2, 2, 448, 443, 3, 2, 2, 2, 448, 444, 3, 2, 2, 2, 449,
	65, 3, 2, 2, 2, 450, 456, 5, 18, 10, 2, 451, 452, 5, 18, 10, 2, 452, 453,
	7, 53, 2, 2, 453, 454, 5, 18, 10, 2, 454, 456, 3, 2, 2, 2, 455, 450, 3,
	2, 2, 2, 455, 451, 3, 2, 2, 2, 456, 67, 3, 2, 2, 2, 457, 459, 7, 9, 2,
	2, 458, 460, 5, 70, 36, 2, 459, 458, 3, 2, 2, 2, 459, 460, 3, 2, 2, 2,
	460, 461, 3, 2, 2, 2, 461, 462, 7, 10, 2, 2, 462, 69, 3, 2, 2, 2, 463,
	469, 5, 72, 37, 2, 464, 465, 5, 72, 37, 2, 465, 466, 7, 8, 2, 2, 466, 467,
	5, 70, 36, 2, 467, 469, 3, 2, 2, 2, 468, 463, 3, 2, 2, 2, 468, 464, 3,
	2, 2, 2, 469, 71, 3, 2, 2, 2, 470, 471, 5, 74, 38, 2, 471, 472, 7, 7, 2,
	2, 472, 473, 5, 18, 10, 2, 473, 73, 3, 2, 2, 2, 474, 475, 7, 70, 2, 2,
	475, 75, 3, 2, 2, 2, 476, 477, 5, 18, 10, 2, 477, 77, 3, 2, 2, 2, 478,
	481, 5, 80, 41, 2, 479, 481, 5, 82, 42, 2, 480, 478, 3, 2, 2, 2, 480, 479,
	3, 2, 2, 2, 481, 79, 3, 2, 2, 2, 482, 483, 7, 9, 2, 2, 483, 484, 5, 74,
	38, 2, 484, 485, 7, 10, 2, 2, 485, 81, 3, 2, 2, 2, 486, 487, 7, 9, 2, 2,
	487, 488, 5, 74, 38, 2, 488, 489, 7, 10, 2, 2, 489, 490, 7, 15, 2, 2, 490,
	83, 3, 2, 2, 2, 491, 492, 5, 78, 40, 2, 492, 85, 3, 2, 2, 2, 493, 494,
	7, 9, 2, 2, 494, 495, 5, 90, 46, 2, 495, 496, 7, 10, 2, 2, 496, 87, 3,
	2, 2, 2, 497, 498, 7, 9, 2, 2, 498, 499, 5, 90, 46, 2, 499, 500, 7, 10,
	2, 2, 500, 501, 7, 15, 2, 2, 501, 89, 3, 2, 2, 2, 502, 508, 5, 80, 41,
	2, 503, 504, 5, 80, 41, 2, 504, 505, 7, 8, 2, 2, 505, 506, 5, 90, 46, 2,
	506, 508, 3, 2, 2, 2, 507, 502, 3, 2, 2, 2, 507, 503, 3, 2, 2, 2, 508,
	91, 3, 2, 2, 2, 509, 512, 5, 86, 44, 2, 510, 512, 5, 88, 45, 2, 511, 509,
	3, 2, 2, 2, 511, 510, 3, 2, 2, 2, 512, 93, 3, 2, 2, 2, 513, 515, 7, 13,
	2, 2, 514, 516, 5, 98, 50, 2, 515, 514, 3, 2, 2, 2, 515, 516, 3, 2, 2,
	2, 516, 517, 3, 2, 2, 2, 517, 519, 7, 14, 2, 2, 518, 520, 5, 108, 55, 2,
	519, 518, 3, 2, 2, 2, 519, 520, 3, 2, 2, 2, 520, 521, 3, 2, 2, 2, 521,
	522, 7, 52, 2, 2, 522, 523, 5, 96, 49, 2, 523, 95, 3, 2, 2, 2, 524, 525,
	5, 18, 10, 2, 525, 97, 3, 2, 2, 2, 526, 532, 5, 100, 51, 2, 527, 528, 5,
	100, 51, 2, 528, 529, 7, 8, 2, 2, 529, 530, 5, 112, 57, 2, 530, 532, 3,
	2, 2, 2, 531, 526, 3, 2, 2, 2, 531, 527, 3, 2, 2, 2, 532, 99, 3, 2, 2,
	2, 533, 539, 5, 102, 52, 2, 534, 535, 5, 102, 52, 2, 535, 536, 7, 8, 2,
	2, 536, 537, 5, 100, 51, 2, 537, 539, 3, 2, 2, 2, 538, 533, 3, 2, 2, 2,
	538, 534, 3, 2, 2, 2, 539, 101, 3, 2, 2, 2, 540, 542, 5, 104, 53, 2, 541,
	543, 5, 106, 54, 2, 542, 541, 3, 2, 2, 2, 542, 543, 3, 2, 2, 2, 543, 103,
	3, 2, 2, 2, 544, 545, 7, 70, 2, 2, 545, 105, 3, 2, 2, 2, 546, 547, 5, 110,
	56, 2, 547, 107, 3, 2, 2, 2, 548, 549, 5, 110, 56, 2, 549, 109, 3, 2, 2,
	2, 550, 551, 7, 51, 2, 2, 551, 552, 5, 26, 14, 2, 552, 111, 3, 2, 2, 2,
	553, 559, 5, 114, 58, 2, 554, 555, 5, 114, 58, 2, 555, 556, 7, 8, 2, 2,
	556, 557, 5, 112, 57, 2, 557, 559, 3, 2, 2, 2, 558, 553, 3, 2, 2, 2, 558,
	554, 3, 2, 2, 2, 559, 113, 3, 2, 2, 2, 560, 561, 7, 16, 2, 2, 561, 562,
	5, 102, 52, 2, 562, 115, 3, 2, 2, 2, 563, 564, 7, 30, 2, 2, 564, 565, 5,
	118, 60, 2, 565, 117, 3, 2, 2, 2, 566, 567, 5, 96, 49, 2, 567, 119, 3,
	2, 2, 2, 568, 569, 7, 31, 2, 2, 569, 570, 5, 122, 62, 2, 570, 571, 7, 32,
	2, 2, 571, 572, 5, 18, 10, 2, 572, 121, 3, 2, 2, 2, 573, 579, 5, 124, 63,
	2, 574, 575, 5, 124, 63, 2, 575, 576, 7, 8, 2, 2, 576, 577, 5, 122, 62,
	2, 577, 579, 3, 2, 2, 2, 578, 573, 3, 2, 2, 2, 578, 574, 3, 2, 2, 2, 579,
	123, 3, 2, 2, 2, 580, 581, 5, 126, 64, 2, 581, 582, 7, 7, 2, 2, 582, 583,
	5, 18, 10, 2, 583, 125, 3, 2, 2, 2, 584, 585, 7, 70, 2, 2, 585, 127, 3,
	2, 2, 2, 586, 587, 7, 33, 2, 2, 587, 588, 5, 130, 66, 2, 588, 589, 7, 34,
	2, 2, 589, 590, 5, 132, 67, 2, 590, 591, 7, 35, 2, 2, 591, 592, 5, 134,
	68, 2, 592, 129, 3, 2, 2, 2, 593, 594, 5, 18, 10, 2, 594, 131, 3, 2, 2,
	2, 595, 596, 5, 18, 10, 2, 596, 133, 3, 2, 2, 2, 597, 598, 5, 18, 10, 2,
	598, 135, 3, 2, 2, 2, 599, 603, 5, 42, 22, 2, 600, 601, 7, 29, 2, 2, 601,
	603, 5, 140, 71, 2, 602, 599, 3, 2, 2, 2, 602, 600, 3, 2, 2, 2, 603, 137,
	3, 2, 2, 2, 604, 607, 5, 56, 29, 2, 605, 607, 5, 140, 71, 2, 606, 604,
	3, 2, 2, 2, 606, 605, 3, 2, 2, 2, 607, 139, 3, 2, 2, 2, 608, 615, 5, 142,
	72, 2, 609, 615, 5, 144, 73, 2, 610, 615, 5, 156, 79, 2, 611, 615, 5, 160,
	81, 2, 612, 615, 5, 174, 88, 2, 613, 615, 5, 178, 90, 2, 614, 608, 3, 2,
	2, 2, 614, 609, 3, 2, 2, 2, 614, 610, 3, 2, 2, 2, 614, 611, 3, 2, 2, 2,
	614, 612, 3, 2, 2, 2, 614, 613, 3, 2, 2, 2, 615, 141, 3, 2, 2, 2, 616,
	617, 9, 2, 2, 2, 617, 143, 3, 2, 2, 2, 618, 619, 7, 9, 2, 2, 619, 620,
	5, 154, 78, 2, 620, 621, 7, 10, 2, 2, 621, 634, 3, 2, 2, 2, 622, 624, 7,
	9, 2, 2, 623, 625, 5, 146, 74, 2, 624, 623, 3, 2, 2, 2, 624, 625, 3, 2,
	2, 2, 625, 626, 3, 2, 2, 2, 626, 634, 7, 10, 2, 2, 627, 628, 7, 9, 2, 2,
	628, 629, 5, 146, 74, 2, 629, 630, 7, 8, 2, 2, 630, 631, 5, 154, 78, 2,
	631, 632, 7, 10, 2, 2, 632, 634, 3, 2, 2, 2, 633, 618, 3, 2, 2, 2, 633,
	622, 3, 2, 2, 2, 633, 627, 3, 2, 2, 2, 634, 145, 3, 2, 2, 2, 635, 641,
	5, 148, 75, 2, 636, 637, 5, 148, 75, 2, 637, 638, 7, 8, 2, 2, 638, 639,
	5, 146, 74, 2, 639, 641, 3, 2, 2, 2, 640, 635, 3, 2, 2, 2, 640, 636, 3,
	2, 2, 2, 641, 147, 3, 2, 2, 2, 642, 644, 7, 16, 2, 2, 643, 642, 3, 2, 2,
	2, 643, 644, 3, 2, 2, 2, 644, 645, 3, 2, 2, 2, 645, 647, 5, 74, 38, 2,
	646, 648, 5, 150, 76, 2, 647, 646, 3, 2, 2, 2, 647, 648, 3, 2, 2, 2, 648,
	149, 3, 2, 2, 2, 649, 650, 7, 7, 2, 2, 650, 651, 5, 152, 77, 2, 651, 151,
	3, 2, 2, 2, 652, 653, 5, 138, 70, 2, 653, 153, 3, 2, 2, 2, 654, 655, 7,
	28, 2, 2, 655, 155, 3, 2, 2, 2, 656, 657, 7, 11, 2, 2, 657, 658, 5, 158,
	80, 2, 658, 659, 7, 12, 2, 2, 659, 157, 3, 2, 2, 2, 660, 661, 5, 138, 70,
	2, 661, 159, 3, 2, 2, 2, 662, 664, 7, 27, 2, 2, 663, 665, 5, 162, 82, 2,
	664, 663, 3, 2, 2, 2, 664, 665, 3, 2, 2, 2, 665, 666, 3, 2, 2, 2, 666,
	667, 7, 14, 2, 2, 667, 668, 5, 108, 55, 2, 668, 161, 3, 2, 2, 2, 669, 676,
	5, 164, 83, 2, 670, 671, 5, 164, 83, 2, 671, 672, 7, 8, 2, 2, 672, 673,
	5, 168, 85, 2, 673, 676, 3, 2, 2, 2, 674, 676, 5, 168, 85, 2, 675, 669,
	3, 2, 2, 2, 675, 670, 3, 2, 2, 2, 675, 674, 3, 2, 2, 2, 676, 163, 3, 2,
	2, 2, 677, 683, 5, 166, 84, 2, 678, 679, 5, 166, 84, 2, 679, 680, 7, 8,
	2, 2, 680, 681, 5, 164, 83, 2, 681, 683, 3, 2, 2, 2, 682, 677, 3, 2, 2,
	2, 682, 678, 3, 2, 2, 2, 683, 165, 3, 2, 2, 2, 684, 685, 5, 172, 87, 2,
	685, 167, 3, 2, 2, 2, 686, 692, 5, 170, 86, 2, 687, 688, 5, 170, 86, 2,
	688, 689, 7, 8, 2, 2, 689, 690, 5, 168, 85, 2, 690, 692, 3, 2, 2, 2, 691,
	686, 3, 2, 2, 2, 691, 687, 3, 2, 2, 2, 692, 169, 3, 2, 2, 2, 693, 694,
	7, 16, 2, 2, 694, 695, 5, 172, 87, 2, 695, 171, 3, 2, 2, 2, 696, 697, 5,
	104, 53, 2, 697, 698, 5, 106, 54, 2, 698, 173, 3, 2, 2, 2, 699, 700, 7,
	17, 2, 2, 700, 701, 5, 176, 89, 2, 701, 175, 3, 2, 2, 2, 702, 703, 7, 9,
	2, 2, 703, 704, 5, 146, 74, 2, 704, 705, 7, 10, 2, 2, 705, 177, 3, 2, 2,
	2, 706, 707, 7, 18, 2, 2, 707, 708, 5, 138, 70, 2, 708, 179, 3, 2, 2, 2,
	709, 710, 7, 26, 2, 2, 710, 711, 5, 18, 10, 2, 711, 181, 3, 2, 2, 2, 712,
	713, 7, 25, 2, 2, 713, 715, 5, 184, 93, 2, 714, 716, 5, 186, 94, 2, 715,
	714, 3, 2, 2, 2, 715, 716, 3, 2, 2, 2, 716, 183, 3, 2, 2, 2, 717, 718,
	5, 18, 10, 2, 718, 185, 3, 2, 2, 2, 719, 720, 7, 24, 2, 2, 720, 721, 5,
	188, 95, 2, 721, 187, 3, 2, 2, 2, 722, 723, 5, 18, 10, 2, 723, 189, 3,
	2, 2, 2, 724, 725, 5, 192, 97, 2, 725, 191, 3, 2, 2, 2, 726, 728, 7, 9,
	2, 2, 727, 729, 5, 194, 98, 2, 728, 727, 3, 2, 2, 2, 728, 729, 3, 2, 2,
	2, 729, 730, 3, 2, 2, 2, 730, 731, 7, 10, 2, 2, 731, 193, 3, 2, 2, 2, 732,
	738, 5, 196, 99, 2, 733, 734, 5, 196, 99, 2, 734, 735, 7, 8, 2, 2, 735,
	736, 5, 194, 98, 2, 736, 738, 3, 2, 2, 2, 737, 732, 3, 2, 2, 2, 737, 733,
	3, 2, 2, 2, 738, 195, 3, 2, 2, 2, 739, 740, 5, 74, 38, 2, 740, 741, 7,
	7, 2, 2, 741, 742, 5, 202, 102, 2, 742, 197, 3, 2, 2, 2, 743, 745, 7, 11,
	2, 2, 744, 746, 5, 200, 101, 2, 745, 744, 3, 2, 2, 2, 745, 746, 3, 2, 2,
	2, 746, 747, 3, 2, 2, 2, 747, 748, 7, 12, 2, 2, 748, 199, 3, 2, 2, 2, 749,
	755, 5, 202, 102, 2, 750, 751, 5, 202, 102, 2, 751, 752, 7, 8, 2, 2, 752,
	753, 5, 200, 101, 2, 753, 755, 3, 2, 2, 2, 754, 749, 3, 2, 2, 2, 754, 750,
	3, 2, 2, 2, 755, 201, 3, 2, 2, 2, 756, 760, 5, 192, 97, 2, 757, 760, 5,
	198, 100, 2, 758, 760, 7, 68, 2, 2, 759, 756, 3, 2, 2, 2, 759, 757, 3,
	2, 2, 2, 759, 758, 3, 2, 2, 2, 760, 203, 3, 2, 2, 2, 59, 206, 211, 217,
	223, 226, 229, 248, 255, 265, 276, 280, 292, 304, 323, 338, 349, 356, 365,
	377, 399, 402, 404, 413, 435, 439, 448, 455, 459, 468, 480, 507, 511, 515,
	519, 531, 538, 542, 558, 578, 602, 606, 614, 624, 633, 640, 643, 647, 664,
	675, 682, 691, 715, 728, 737, 745, 754, 759,
}
var literalNames = []string{
	"", "", "", "", "", "'='", "','", "'['", "']'", "'{'", "'}'", "'('", "')'",
	"'?'", "'optional'", "'table'", "'nullable'", "';'", "'section'", "'shared'",
	"'and'", "'or'", "'otherwise'", "'try'", "'error'", "'function ('", "'...'",
	"'type'", "'each'", "'let'", "'in'", "'if'", "'then'", "'else'", "'text'",
	"'record'", "'number'", "'none'", "'logical'", "'list'", "'fuction'", "'duration'",
	"'datetimezone'", "'any'", "'anynonnull'", "'binary'", "'date'", "'datetime'",
	"'@'", "'as'", "'=>'", "'..'", "'!'", "'not'", "'+'", "'-'", "'meta'",
	"'is'", "'<>'", "'>'", "'<'", "'/'", "'*'", "'&'", "'<='", "'>='",
}
var symbolicNames = []string{
	"", "WHITESPACE", "NEW_LINE_CHAR", "COMMENT", "CHARACHTER_ESCAPE_SEQUENCE",
	"EQUALS", "COMMA", "OPEN_BRACKET", "CLOSE_BRACKET", "OPEN_BRACE", "CLOSE_BRACE",
	"OPEN_PAREN", "CLOSE_PAREN", "OPTIONAL", "OPTIONAL_TEXT", "TABLE", "NULLABLE",
	"SEMICOLON", "SECTION", "SHARED", "AND", "OR", "OTHERWISE", "TRY", "ERROR",
	"FUNCTION_START", "ELLIPSES", "TYPE", "EACH", "LET", "IN", "IF", "THEN",
	"ELSE", "TEXT", "RECORD", "NUMBER", "NONE", "LOGICAL", "LIST", "FUNCTION",
	"DURATION", "DATETIMEZONE", "ANY", "ANYNONNULL", "BINARY", "DATE", "DATETIME",
	"AT", "AS", "ARROW", "DOTDOT", "BANG", "NOT", "PLUS", "MINUS", "META",
	"IS", "NEQ", "GE", "LE", "SLASH", "STAR", "AMP", "LEQ", "GEQ", "LITERAL",
	"TEXT_LITERAL", "IDENTIFIER", "REGULAR_IDENTIFIER", "AVAILABLE_IDENTIFIER",
	"KEYWORD_OR_IDENTIFIER",
}

var ruleNames = []string{
	"document", "section_document", "section", "section_name", "section_members",
	"section_member", "section_member_name", "expression_document", "expression",
	"logical_or_expression", "logical_and_expression", "is_expression", "nullable_primitive_type",
	"as_expression", "equality_expression", "relational_expression", "additive_expression",
	"multiplicative_expression", "metadata_expression", "unary_expression",
	"primary_expression", "literal_expression", "identifier_expression", "identifier_reference",
	"exclusive_identifier_reference", "inclusive_identifier_reference", "section_access_expression",
	"parenthesized_expression", "not_implemented_expression", "argument_list",
	"list_expression", "item_list", "item", "record_expression", "field_list",
	"field", "field_name", "item_selector", "field_selector", "required_field_selector",
	"optional_field_selector", "implicit_target_field_selection", "required_projection",
	"optional_projection", "required_selector_list", "implicit_target_projection",
	"function_expression", "function_body", "parameter_list", "fixed_parameter_list",
	"parameter", "parameter_name", "parameter_type", "return_type", "assertion",
	"optional_parameter_list", "optional_parameter", "each_expression", "each_expression_body",
	"let_expression", "variable_list", "variable", "variable_name", "if_expression",
	"if_condition", "true_expression", "false_expression", "type_expression",
	"type_expr", "primary_type", "primitive_type", "record_type", "field_specification_list",
	"field_specification", "field_type_specification", "field_type", "open_record_marker",
	"list_type", "item_type", "function_type", "parameter_specification_list",
	"required_parameter_specification_list", "required_parameter_specification",
	"optional_parameter_specification_list", "optional_parameter_specification",
	"parameter_specification", "table_type", "row_type", "nullable_type", "error_raising_expression",
	"error_handling_expression", "protected_expression", "otherwise_clause",
	"default_expression", "literal_attribs", "record_literal", "literal_field_list",
	"literal_field", "list_literal", "literal_item_list", "any_literal",
}

type PowerQueryParser struct {
	*antlr.BaseParser
}

// NewPowerQueryParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *PowerQueryParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPowerQueryParser(input antlr.TokenStream) *PowerQueryParser {
	this := new(PowerQueryParser)
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
	this.GrammarFileName = "PowerQueryParser.g4"

	return this
}

// PowerQueryParser tokens.
const (
	PowerQueryParserEOF                        = antlr.TokenEOF
	PowerQueryParserWHITESPACE                 = 1
	PowerQueryParserNEW_LINE_CHAR              = 2
	PowerQueryParserCOMMENT                    = 3
	PowerQueryParserCHARACHTER_ESCAPE_SEQUENCE = 4
	PowerQueryParserEQUALS                     = 5
	PowerQueryParserCOMMA                      = 6
	PowerQueryParserOPEN_BRACKET               = 7
	PowerQueryParserCLOSE_BRACKET              = 8
	PowerQueryParserOPEN_BRACE                 = 9
	PowerQueryParserCLOSE_BRACE                = 10
	PowerQueryParserOPEN_PAREN                 = 11
	PowerQueryParserCLOSE_PAREN                = 12
	PowerQueryParserOPTIONAL                   = 13
	PowerQueryParserOPTIONAL_TEXT              = 14
	PowerQueryParserTABLE                      = 15
	PowerQueryParserNULLABLE                   = 16
	PowerQueryParserSEMICOLON                  = 17
	PowerQueryParserSECTION                    = 18
	PowerQueryParserSHARED                     = 19
	PowerQueryParserAND                        = 20
	PowerQueryParserOR                         = 21
	PowerQueryParserOTHERWISE                  = 22
	PowerQueryParserTRY                        = 23
	PowerQueryParserERROR                      = 24
	PowerQueryParserFUNCTION_START             = 25
	PowerQueryParserELLIPSES                   = 26
	PowerQueryParserTYPE                       = 27
	PowerQueryParserEACH                       = 28
	PowerQueryParserLET                        = 29
	PowerQueryParserIN                         = 30
	PowerQueryParserIF                         = 31
	PowerQueryParserTHEN                       = 32
	PowerQueryParserELSE                       = 33
	PowerQueryParserTEXT                       = 34
	PowerQueryParserRECORD                     = 35
	PowerQueryParserNUMBER                     = 36
	PowerQueryParserNONE                       = 37
	PowerQueryParserLOGICAL                    = 38
	PowerQueryParserLIST                       = 39
	PowerQueryParserFUNCTION                   = 40
	PowerQueryParserDURATION                   = 41
	PowerQueryParserDATETIMEZONE               = 42
	PowerQueryParserANY                        = 43
	PowerQueryParserANYNONNULL                 = 44
	PowerQueryParserBINARY                     = 45
	PowerQueryParserDATE                       = 46
	PowerQueryParserDATETIME                   = 47
	PowerQueryParserAT                         = 48
	PowerQueryParserAS                         = 49
	PowerQueryParserARROW                      = 50
	PowerQueryParserDOTDOT                     = 51
	PowerQueryParserBANG                       = 52
	PowerQueryParserNOT                        = 53
	PowerQueryParserPLUS                       = 54
	PowerQueryParserMINUS                      = 55
	PowerQueryParserMETA                       = 56
	PowerQueryParserIS                         = 57
	PowerQueryParserNEQ                        = 58
	PowerQueryParserGE                         = 59
	PowerQueryParserLE                         = 60
	PowerQueryParserSLASH                      = 61
	PowerQueryParserSTAR                       = 62
	PowerQueryParserAMP                        = 63
	PowerQueryParserLEQ                        = 64
	PowerQueryParserGEQ                        = 65
	PowerQueryParserLITERAL                    = 66
	PowerQueryParserTEXT_LITERAL               = 67
	PowerQueryParserIDENTIFIER                 = 68
	PowerQueryParserREGULAR_IDENTIFIER         = 69
	PowerQueryParserAVAILABLE_IDENTIFIER       = 70
	PowerQueryParserKEYWORD_OR_IDENTIFIER      = 71
)

// PowerQueryParser rules.
const (
	PowerQueryParserRULE_document                              = 0
	PowerQueryParserRULE_section_document                      = 1
	PowerQueryParserRULE_section                               = 2
	PowerQueryParserRULE_section_name                          = 3
	PowerQueryParserRULE_section_members                       = 4
	PowerQueryParserRULE_section_member                        = 5
	PowerQueryParserRULE_section_member_name                   = 6
	PowerQueryParserRULE_expression_document                   = 7
	PowerQueryParserRULE_expression                            = 8
	PowerQueryParserRULE_logical_or_expression                 = 9
	PowerQueryParserRULE_logical_and_expression                = 10
	PowerQueryParserRULE_is_expression                         = 11
	PowerQueryParserRULE_nullable_primitive_type               = 12
	PowerQueryParserRULE_as_expression                         = 13
	PowerQueryParserRULE_equality_expression                   = 14
	PowerQueryParserRULE_relational_expression                 = 15
	PowerQueryParserRULE_additive_expression                   = 16
	PowerQueryParserRULE_multiplicative_expression             = 17
	PowerQueryParserRULE_metadata_expression                   = 18
	PowerQueryParserRULE_unary_expression                      = 19
	PowerQueryParserRULE_primary_expression                    = 20
	PowerQueryParserRULE_literal_expression                    = 21
	PowerQueryParserRULE_identifier_expression                 = 22
	PowerQueryParserRULE_identifier_reference                  = 23
	PowerQueryParserRULE_exclusive_identifier_reference        = 24
	PowerQueryParserRULE_inclusive_identifier_reference        = 25
	PowerQueryParserRULE_section_access_expression             = 26
	PowerQueryParserRULE_parenthesized_expression              = 27
	PowerQueryParserRULE_not_implemented_expression            = 28
	PowerQueryParserRULE_argument_list                         = 29
	PowerQueryParserRULE_list_expression                       = 30
	PowerQueryParserRULE_item_list                             = 31
	PowerQueryParserRULE_item                                  = 32
	PowerQueryParserRULE_record_expression                     = 33
	PowerQueryParserRULE_field_list                            = 34
	PowerQueryParserRULE_field                                 = 35
	PowerQueryParserRULE_field_name                            = 36
	PowerQueryParserRULE_item_selector                         = 37
	PowerQueryParserRULE_field_selector                        = 38
	PowerQueryParserRULE_required_field_selector               = 39
	PowerQueryParserRULE_optional_field_selector               = 40
	PowerQueryParserRULE_implicit_target_field_selection       = 41
	PowerQueryParserRULE_required_projection                   = 42
	PowerQueryParserRULE_optional_projection                   = 43
	PowerQueryParserRULE_required_selector_list                = 44
	PowerQueryParserRULE_implicit_target_projection            = 45
	PowerQueryParserRULE_function_expression                   = 46
	PowerQueryParserRULE_function_body                         = 47
	PowerQueryParserRULE_parameter_list                        = 48
	PowerQueryParserRULE_fixed_parameter_list                  = 49
	PowerQueryParserRULE_parameter                             = 50
	PowerQueryParserRULE_parameter_name                        = 51
	PowerQueryParserRULE_parameter_type                        = 52
	PowerQueryParserRULE_return_type                           = 53
	PowerQueryParserRULE_assertion                             = 54
	PowerQueryParserRULE_optional_parameter_list               = 55
	PowerQueryParserRULE_optional_parameter                    = 56
	PowerQueryParserRULE_each_expression                       = 57
	PowerQueryParserRULE_each_expression_body                  = 58
	PowerQueryParserRULE_let_expression                        = 59
	PowerQueryParserRULE_variable_list                         = 60
	PowerQueryParserRULE_variable                              = 61
	PowerQueryParserRULE_variable_name                         = 62
	PowerQueryParserRULE_if_expression                         = 63
	PowerQueryParserRULE_if_condition                          = 64
	PowerQueryParserRULE_true_expression                       = 65
	PowerQueryParserRULE_false_expression                      = 66
	PowerQueryParserRULE_type_expression                       = 67
	PowerQueryParserRULE_type_expr                             = 68
	PowerQueryParserRULE_primary_type                          = 69
	PowerQueryParserRULE_primitive_type                        = 70
	PowerQueryParserRULE_record_type                           = 71
	PowerQueryParserRULE_field_specification_list              = 72
	PowerQueryParserRULE_field_specification                   = 73
	PowerQueryParserRULE_field_type_specification              = 74
	PowerQueryParserRULE_field_type                            = 75
	PowerQueryParserRULE_open_record_marker                    = 76
	PowerQueryParserRULE_list_type                             = 77
	PowerQueryParserRULE_item_type                             = 78
	PowerQueryParserRULE_function_type                         = 79
	PowerQueryParserRULE_parameter_specification_list          = 80
	PowerQueryParserRULE_required_parameter_specification_list = 81
	PowerQueryParserRULE_required_parameter_specification      = 82
	PowerQueryParserRULE_optional_parameter_specification_list = 83
	PowerQueryParserRULE_optional_parameter_specification      = 84
	PowerQueryParserRULE_parameter_specification               = 85
	PowerQueryParserRULE_table_type                            = 86
	PowerQueryParserRULE_row_type                              = 87
	PowerQueryParserRULE_nullable_type                         = 88
	PowerQueryParserRULE_error_raising_expression              = 89
	PowerQueryParserRULE_error_handling_expression             = 90
	PowerQueryParserRULE_protected_expression                  = 91
	PowerQueryParserRULE_otherwise_clause                      = 92
	PowerQueryParserRULE_default_expression                    = 93
	PowerQueryParserRULE_literal_attribs                       = 94
	PowerQueryParserRULE_record_literal                        = 95
	PowerQueryParserRULE_literal_field_list                    = 96
	PowerQueryParserRULE_literal_field                         = 97
	PowerQueryParserRULE_list_literal                          = 98
	PowerQueryParserRULE_literal_item_list                     = 99
	PowerQueryParserRULE_any_literal                           = 100
)

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) Section_document() ISection_documentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISection_documentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISection_documentContext)
}

func (s *DocumentContext) Expression_document() IExpression_documentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_documentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_documentContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *PowerQueryParser) Document() (localctx IDocumentContext) {
	this := p
	_ = this

	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PowerQueryParserRULE_document)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Section_document()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)
			p.Expression_document()
		}

	}

	return localctx
}

// ISection_documentContext is an interface to support dynamic dispatch.
type ISection_documentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSection_documentContext differentiates from other interfaces.
	IsSection_documentContext()
}

type Section_documentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySection_documentContext() *Section_documentContext {
	var p = new(Section_documentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_section_document
	return p
}

func (*Section_documentContext) IsSection_documentContext() {}

func NewSection_documentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Section_documentContext {
	var p = new(Section_documentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_section_document

	return p
}

func (s *Section_documentContext) GetParser() antlr.Parser { return s.parser }

func (s *Section_documentContext) Section() ISectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISectionContext)
}

func (s *Section_documentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Section_documentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Section_documentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterSection_document(s)
	}
}

func (s *Section_documentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitSection_document(s)
	}
}

func (p *PowerQueryParser) Section_document() (localctx ISection_documentContext) {
	this := p
	_ = this

	localctx = NewSection_documentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PowerQueryParserRULE_section_document)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(206)
		p.Section()
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
	p.RuleIndex = PowerQueryParserRULE_section
	return p
}

func (*SectionContext) IsSectionContext() {}

func NewSectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionContext {
	var p = new(SectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_section

	return p
}

func (s *SectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionContext) SECTION() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserSECTION, 0)
}

func (s *SectionContext) Section_name() ISection_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISection_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISection_nameContext)
}

func (s *SectionContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserSEMICOLON, 0)
}

func (s *SectionContext) Literal_attribs() ILiteral_attribsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteral_attribsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteral_attribsContext)
}

func (s *SectionContext) Section_members() ISection_membersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISection_membersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISection_membersContext)
}

func (s *SectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterSection(s)
	}
}

func (s *SectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitSection(s)
	}
}

func (p *PowerQueryParser) Section() (localctx ISectionContext) {
	this := p
	_ = this

	localctx = NewSectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PowerQueryParserRULE_section)
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
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserOPEN_BRACKET {
		{
			p.SetState(208)
			p.Literal_attribs()
		}

	}
	{
		p.SetState(211)
		p.Match(PowerQueryParserSECTION)
	}
	{
		p.SetState(212)
		p.Section_name()
	}
	{
		p.SetState(213)
		p.Match(PowerQueryParserSEMICOLON)
	}
	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserOPEN_BRACKET || _la == PowerQueryParserSHARED || _la == PowerQueryParserIDENTIFIER {
		{
			p.SetState(214)
			p.Section_members()
		}

	}

	return localctx
}

// ISection_nameContext is an interface to support dynamic dispatch.
type ISection_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSection_nameContext differentiates from other interfaces.
	IsSection_nameContext()
}

type Section_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySection_nameContext() *Section_nameContext {
	var p = new(Section_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_section_name
	return p
}

func (*Section_nameContext) IsSection_nameContext() {}

func NewSection_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Section_nameContext {
	var p = new(Section_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_section_name

	return p
}

func (s *Section_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Section_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserIDENTIFIER, 0)
}

func (s *Section_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Section_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Section_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterSection_name(s)
	}
}

func (s *Section_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitSection_name(s)
	}
}

func (p *PowerQueryParser) Section_name() (localctx ISection_nameContext) {
	this := p
	_ = this

	localctx = NewSection_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PowerQueryParserRULE_section_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PowerQueryParserIDENTIFIER)
	}

	return localctx
}

// ISection_membersContext is an interface to support dynamic dispatch.
type ISection_membersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSection_membersContext differentiates from other interfaces.
	IsSection_membersContext()
}

type Section_membersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySection_membersContext() *Section_membersContext {
	var p = new(Section_membersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_section_members
	return p
}

func (*Section_membersContext) IsSection_membersContext() {}

func NewSection_membersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Section_membersContext {
	var p = new(Section_membersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_section_members

	return p
}

func (s *Section_membersContext) GetParser() antlr.Parser { return s.parser }

func (s *Section_membersContext) Section_member() ISection_memberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISection_memberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISection_memberContext)
}

func (s *Section_membersContext) Section_members() ISection_membersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISection_membersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISection_membersContext)
}

func (s *Section_membersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Section_membersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Section_membersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterSection_members(s)
	}
}

func (s *Section_membersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitSection_members(s)
	}
}

func (p *PowerQueryParser) Section_members() (localctx ISection_membersContext) {
	this := p
	_ = this

	localctx = NewSection_membersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PowerQueryParserRULE_section_members)
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
		p.SetState(219)
		p.Section_member()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserOPEN_BRACKET || _la == PowerQueryParserSHARED || _la == PowerQueryParserIDENTIFIER {
		{
			p.SetState(220)
			p.Section_members()
		}

	}

	return localctx
}

// ISection_memberContext is an interface to support dynamic dispatch.
type ISection_memberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSection_memberContext differentiates from other interfaces.
	IsSection_memberContext()
}

type Section_memberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySection_memberContext() *Section_memberContext {
	var p = new(Section_memberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_section_member
	return p
}

func (*Section_memberContext) IsSection_memberContext() {}

func NewSection_memberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Section_memberContext {
	var p = new(Section_memberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_section_member

	return p
}

func (s *Section_memberContext) GetParser() antlr.Parser { return s.parser }

func (s *Section_memberContext) Section_member_name() ISection_member_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISection_member_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISection_member_nameContext)
}

func (s *Section_memberContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserEQUALS, 0)
}

func (s *Section_memberContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Section_memberContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserSEMICOLON, 0)
}

func (s *Section_memberContext) Literal_attribs() ILiteral_attribsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteral_attribsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteral_attribsContext)
}

func (s *Section_memberContext) SHARED() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserSHARED, 0)
}

func (s *Section_memberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Section_memberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Section_memberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterSection_member(s)
	}
}

func (s *Section_memberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitSection_member(s)
	}
}

func (p *PowerQueryParser) Section_member() (localctx ISection_memberContext) {
	this := p
	_ = this

	localctx = NewSection_memberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PowerQueryParserRULE_section_member)
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
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserOPEN_BRACKET {
		{
			p.SetState(223)
			p.Literal_attribs()
		}

	}
	p.SetState(227)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserSHARED {
		{
			p.SetState(226)
			p.Match(PowerQueryParserSHARED)
		}

	}
	{
		p.SetState(229)
		p.Section_member_name()
	}
	{
		p.SetState(230)
		p.Match(PowerQueryParserEQUALS)
	}
	{
		p.SetState(231)
		p.Expression()
	}
	{
		p.SetState(232)
		p.Match(PowerQueryParserSEMICOLON)
	}

	return localctx
}

// ISection_member_nameContext is an interface to support dynamic dispatch.
type ISection_member_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSection_member_nameContext differentiates from other interfaces.
	IsSection_member_nameContext()
}

type Section_member_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySection_member_nameContext() *Section_member_nameContext {
	var p = new(Section_member_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_section_member_name
	return p
}

func (*Section_member_nameContext) IsSection_member_nameContext() {}

func NewSection_member_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Section_member_nameContext {
	var p = new(Section_member_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_section_member_name

	return p
}

func (s *Section_member_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Section_member_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserIDENTIFIER, 0)
}

func (s *Section_member_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Section_member_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Section_member_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterSection_member_name(s)
	}
}

func (s *Section_member_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitSection_member_name(s)
	}
}

func (p *PowerQueryParser) Section_member_name() (localctx ISection_member_nameContext) {
	this := p
	_ = this

	localctx = NewSection_member_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PowerQueryParserRULE_section_member_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(234)
		p.Match(PowerQueryParserIDENTIFIER)
	}

	return localctx
}

// IExpression_documentContext is an interface to support dynamic dispatch.
type IExpression_documentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression_documentContext differentiates from other interfaces.
	IsExpression_documentContext()
}

type Expression_documentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_documentContext() *Expression_documentContext {
	var p = new(Expression_documentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_expression_document
	return p
}

func (*Expression_documentContext) IsExpression_documentContext() {}

func NewExpression_documentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_documentContext {
	var p = new(Expression_documentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_expression_document

	return p
}

func (s *Expression_documentContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_documentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expression_documentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_documentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_documentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterExpression_document(s)
	}
}

func (s *Expression_documentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitExpression_document(s)
	}
}

func (p *PowerQueryParser) Expression_document() (localctx IExpression_documentContext) {
	this := p
	_ = this

	localctx = NewExpression_documentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PowerQueryParserRULE_expression_document)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(236)
		p.Expression()
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
	p.RuleIndex = PowerQueryParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Logical_or_expression() ILogical_or_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_or_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_or_expressionContext)
}

func (s *ExpressionContext) Each_expression() IEach_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEach_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEach_expressionContext)
}

func (s *ExpressionContext) Function_expression() IFunction_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_expressionContext)
}

func (s *ExpressionContext) Let_expression() ILet_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILet_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILet_expressionContext)
}

func (s *ExpressionContext) If_expression() IIf_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_expressionContext)
}

func (s *ExpressionContext) Error_raising_expression() IError_raising_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IError_raising_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IError_raising_expressionContext)
}

func (s *ExpressionContext) Error_handling_expression() IError_handling_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IError_handling_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IError_handling_expressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PowerQueryParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PowerQueryParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.Logical_or_expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
			p.Each_expression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(240)
			p.Function_expression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(241)
			p.Let_expression()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(242)
			p.If_expression()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(243)
			p.Let_expression()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(244)
			p.Error_raising_expression()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(245)
			p.Error_handling_expression()
		}

	}

	return localctx
}

// ILogical_or_expressionContext is an interface to support dynamic dispatch.
type ILogical_or_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogical_or_expressionContext differentiates from other interfaces.
	IsLogical_or_expressionContext()
}

type Logical_or_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogical_or_expressionContext() *Logical_or_expressionContext {
	var p = new(Logical_or_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_logical_or_expression
	return p
}

func (*Logical_or_expressionContext) IsLogical_or_expressionContext() {}

func NewLogical_or_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_or_expressionContext {
	var p = new(Logical_or_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_logical_or_expression

	return p
}

func (s *Logical_or_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_or_expressionContext) Logical_and_expression() ILogical_and_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_and_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_and_expressionContext)
}

func (s *Logical_or_expressionContext) OR() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOR, 0)
}

func (s *Logical_or_expressionContext) Logical_or_expression() ILogical_or_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_or_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_or_expressionContext)
}

func (s *Logical_or_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_or_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_or_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterLogical_or_expression(s)
	}
}

func (s *Logical_or_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitLogical_or_expression(s)
	}
}

func (p *PowerQueryParser) Logical_or_expression() (localctx ILogical_or_expressionContext) {
	this := p
	_ = this

	localctx = NewLogical_or_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PowerQueryParserRULE_logical_or_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.logical_and_expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(249)
			p.logical_and_expression(0)
		}
		{
			p.SetState(250)
			p.Match(PowerQueryParserOR)
		}
		{
			p.SetState(251)
			p.Logical_or_expression()
		}

	}

	return localctx
}

// ILogical_and_expressionContext is an interface to support dynamic dispatch.
type ILogical_and_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogical_and_expressionContext differentiates from other interfaces.
	IsLogical_and_expressionContext()
}

type Logical_and_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogical_and_expressionContext() *Logical_and_expressionContext {
	var p = new(Logical_and_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_logical_and_expression
	return p
}

func (*Logical_and_expressionContext) IsLogical_and_expressionContext() {}

func NewLogical_and_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Logical_and_expressionContext {
	var p = new(Logical_and_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_logical_and_expression

	return p
}

func (s *Logical_and_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Logical_and_expressionContext) Is_expression() IIs_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIs_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIs_expressionContext)
}

func (s *Logical_and_expressionContext) Logical_and_expression() ILogical_and_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogical_and_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogical_and_expressionContext)
}

func (s *Logical_and_expressionContext) AND() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserAND, 0)
}

func (s *Logical_and_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Logical_and_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Logical_and_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterLogical_and_expression(s)
	}
}

func (s *Logical_and_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitLogical_and_expression(s)
	}
}

func (p *PowerQueryParser) Logical_and_expression() (localctx ILogical_and_expressionContext) {
	return p.logical_and_expression(0)
}

func (p *PowerQueryParser) logical_and_expression(_p int) (localctx ILogical_and_expressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLogical_and_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILogical_and_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, PowerQueryParserRULE_logical_and_expression, _p)

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
	{
		p.SetState(256)
		p.is_expression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogical_and_expressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, PowerQueryParserRULE_logical_and_expression)
			p.SetState(258)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(259)
				p.Match(PowerQueryParserAND)
			}
			{
				p.SetState(260)
				p.is_expression(0)
			}

		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// IIs_expressionContext is an interface to support dynamic dispatch.
type IIs_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIs_expressionContext differentiates from other interfaces.
	IsIs_expressionContext()
}

type Is_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIs_expressionContext() *Is_expressionContext {
	var p = new(Is_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_is_expression
	return p
}

func (*Is_expressionContext) IsIs_expressionContext() {}

func NewIs_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Is_expressionContext {
	var p = new(Is_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_is_expression

	return p
}

func (s *Is_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Is_expressionContext) As_expression() IAs_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAs_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAs_expressionContext)
}

func (s *Is_expressionContext) Is_expression() IIs_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIs_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIs_expressionContext)
}

func (s *Is_expressionContext) IS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserIS, 0)
}

func (s *Is_expressionContext) Nullable_primitive_type() INullable_primitive_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullable_primitive_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullable_primitive_typeContext)
}

func (s *Is_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Is_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Is_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterIs_expression(s)
	}
}

func (s *Is_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitIs_expression(s)
	}
}

func (p *PowerQueryParser) Is_expression() (localctx IIs_expressionContext) {
	return p.is_expression(0)
}

func (p *PowerQueryParser) is_expression(_p int) (localctx IIs_expressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewIs_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IIs_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 22
	p.EnterRecursionRule(localctx, 22, PowerQueryParserRULE_is_expression, _p)

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
	{
		p.SetState(267)
		p.as_expression(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewIs_expressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, PowerQueryParserRULE_is_expression)
			p.SetState(269)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(270)
				p.Match(PowerQueryParserIS)
			}
			{
				p.SetState(271)
				p.Nullable_primitive_type()
			}

		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}

	return localctx
}

// INullable_primitive_typeContext is an interface to support dynamic dispatch.
type INullable_primitive_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullable_primitive_typeContext differentiates from other interfaces.
	IsNullable_primitive_typeContext()
}

type Nullable_primitive_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullable_primitive_typeContext() *Nullable_primitive_typeContext {
	var p = new(Nullable_primitive_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_nullable_primitive_type
	return p
}

func (*Nullable_primitive_typeContext) IsNullable_primitive_typeContext() {}

func NewNullable_primitive_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nullable_primitive_typeContext {
	var p = new(Nullable_primitive_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_nullable_primitive_type

	return p
}

func (s *Nullable_primitive_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Nullable_primitive_typeContext) Primitive_type() IPrimitive_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitive_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitive_typeContext)
}

func (s *Nullable_primitive_typeContext) NULLABLE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserNULLABLE, 0)
}

func (s *Nullable_primitive_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nullable_primitive_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nullable_primitive_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterNullable_primitive_type(s)
	}
}

func (s *Nullable_primitive_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitNullable_primitive_type(s)
	}
}

func (p *PowerQueryParser) Nullable_primitive_type() (localctx INullable_primitive_typeContext) {
	this := p
	_ = this

	localctx = NewNullable_primitive_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PowerQueryParserRULE_nullable_primitive_type)
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
	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserNULLABLE {
		{
			p.SetState(277)
			p.Match(PowerQueryParserNULLABLE)
		}

	}
	{
		p.SetState(280)
		p.Primitive_type()
	}

	return localctx
}

// IAs_expressionContext is an interface to support dynamic dispatch.
type IAs_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAs_expressionContext differentiates from other interfaces.
	IsAs_expressionContext()
}

type As_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAs_expressionContext() *As_expressionContext {
	var p = new(As_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_as_expression
	return p
}

func (*As_expressionContext) IsAs_expressionContext() {}

func NewAs_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *As_expressionContext {
	var p = new(As_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_as_expression

	return p
}

func (s *As_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *As_expressionContext) Equality_expression() IEquality_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEquality_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEquality_expressionContext)
}

func (s *As_expressionContext) As_expression() IAs_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAs_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAs_expressionContext)
}

func (s *As_expressionContext) AS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserAS, 0)
}

func (s *As_expressionContext) Nullable_primitive_type() INullable_primitive_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullable_primitive_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullable_primitive_typeContext)
}

func (s *As_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *As_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *As_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterAs_expression(s)
	}
}

func (s *As_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitAs_expression(s)
	}
}

func (p *PowerQueryParser) As_expression() (localctx IAs_expressionContext) {
	return p.as_expression(0)
}

func (p *PowerQueryParser) as_expression(_p int) (localctx IAs_expressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAs_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAs_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 26
	p.EnterRecursionRule(localctx, 26, PowerQueryParserRULE_as_expression, _p)

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
	{
		p.SetState(283)
		p.Equality_expression()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAs_expressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, PowerQueryParserRULE_as_expression)
			p.SetState(285)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(286)
				p.Match(PowerQueryParserAS)
			}
			{
				p.SetState(287)
				p.Nullable_primitive_type()
			}

		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IEquality_expressionContext is an interface to support dynamic dispatch.
type IEquality_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEquality_expressionContext differentiates from other interfaces.
	IsEquality_expressionContext()
}

type Equality_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEquality_expressionContext() *Equality_expressionContext {
	var p = new(Equality_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_equality_expression
	return p
}

func (*Equality_expressionContext) IsEquality_expressionContext() {}

func NewEquality_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Equality_expressionContext {
	var p = new(Equality_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_equality_expression

	return p
}

func (s *Equality_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Equality_expressionContext) Relational_expression() IRelational_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelational_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelational_expressionContext)
}

func (s *Equality_expressionContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserEQUALS, 0)
}

func (s *Equality_expressionContext) Equality_expression() IEquality_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEquality_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEquality_expressionContext)
}

func (s *Equality_expressionContext) NEQ() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserNEQ, 0)
}

func (s *Equality_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Equality_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Equality_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterEquality_expression(s)
	}
}

func (s *Equality_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitEquality_expression(s)
	}
}

func (p *PowerQueryParser) Equality_expression() (localctx IEquality_expressionContext) {
	this := p
	_ = this

	localctx = NewEquality_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PowerQueryParserRULE_equality_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.Relational_expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.Relational_expression()
		}
		{
			p.SetState(295)
			p.Match(PowerQueryParserEQUALS)
		}
		{
			p.SetState(296)
			p.Equality_expression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(298)
			p.Relational_expression()
		}
		{
			p.SetState(299)
			p.Match(PowerQueryParserNEQ)
		}
		{
			p.SetState(300)
			p.Equality_expression()
		}

	}

	return localctx
}

// IRelational_expressionContext is an interface to support dynamic dispatch.
type IRelational_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelational_expressionContext differentiates from other interfaces.
	IsRelational_expressionContext()
}

type Relational_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelational_expressionContext() *Relational_expressionContext {
	var p = new(Relational_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_relational_expression
	return p
}

func (*Relational_expressionContext) IsRelational_expressionContext() {}

func NewRelational_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Relational_expressionContext {
	var p = new(Relational_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_relational_expression

	return p
}

func (s *Relational_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Relational_expressionContext) AllAdditive_expression() []IAdditive_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAdditive_expressionContext)(nil)).Elem())
	var tst = make([]IAdditive_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAdditive_expressionContext)
		}
	}

	return tst
}

func (s *Relational_expressionContext) Additive_expression(i int) IAdditive_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditive_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAdditive_expressionContext)
}

func (s *Relational_expressionContext) LE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserLE, 0)
}

func (s *Relational_expressionContext) Relational_expression() IRelational_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelational_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelational_expressionContext)
}

func (s *Relational_expressionContext) GE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserGE, 0)
}

func (s *Relational_expressionContext) LEQ() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserLEQ, 0)
}

func (s *Relational_expressionContext) GEQ() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserGEQ, 0)
}

func (s *Relational_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Relational_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Relational_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterRelational_expression(s)
	}
}

func (s *Relational_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitRelational_expression(s)
	}
}

func (p *PowerQueryParser) Relational_expression() (localctx IRelational_expressionContext) {
	this := p
	_ = this

	localctx = NewRelational_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PowerQueryParserRULE_relational_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(321)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(304)
			p.Additive_expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.Additive_expression()
		}
		{
			p.SetState(306)
			p.Match(PowerQueryParserLE)
		}
		{
			p.SetState(307)
			p.Relational_expression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(309)
			p.Additive_expression()
		}
		{
			p.SetState(310)
			p.Match(PowerQueryParserGE)
		}
		{
			p.SetState(311)
			p.Relational_expression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(313)
			p.Additive_expression()
		}
		{
			p.SetState(314)
			p.Match(PowerQueryParserLEQ)
		}
		{
			p.SetState(315)
			p.Additive_expression()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(317)
			p.Additive_expression()
		}
		{
			p.SetState(318)
			p.Match(PowerQueryParserGEQ)
		}
		{
			p.SetState(319)
			p.Relational_expression()
		}

	}

	return localctx
}

// IAdditive_expressionContext is an interface to support dynamic dispatch.
type IAdditive_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAdditive_expressionContext differentiates from other interfaces.
	IsAdditive_expressionContext()
}

type Additive_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditive_expressionContext() *Additive_expressionContext {
	var p = new(Additive_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_additive_expression
	return p
}

func (*Additive_expressionContext) IsAdditive_expressionContext() {}

func NewAdditive_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Additive_expressionContext {
	var p = new(Additive_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_additive_expression

	return p
}

func (s *Additive_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Additive_expressionContext) Multiplicative_expression() IMultiplicative_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicative_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicative_expressionContext)
}

func (s *Additive_expressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserPLUS, 0)
}

func (s *Additive_expressionContext) Additive_expression() IAdditive_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAdditive_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAdditive_expressionContext)
}

func (s *Additive_expressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserMINUS, 0)
}

func (s *Additive_expressionContext) AMP() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserAMP, 0)
}

func (s *Additive_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Additive_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Additive_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterAdditive_expression(s)
	}
}

func (s *Additive_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitAdditive_expression(s)
	}
}

func (p *PowerQueryParser) Additive_expression() (localctx IAdditive_expressionContext) {
	this := p
	_ = this

	localctx = NewAdditive_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PowerQueryParserRULE_additive_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(323)
			p.Multiplicative_expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.Multiplicative_expression()
		}
		{
			p.SetState(325)
			p.Match(PowerQueryParserPLUS)
		}
		{
			p.SetState(326)
			p.Additive_expression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(328)
			p.Multiplicative_expression()
		}
		{
			p.SetState(329)
			p.Match(PowerQueryParserMINUS)
		}
		{
			p.SetState(330)
			p.Additive_expression()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(332)
			p.Multiplicative_expression()
		}
		{
			p.SetState(333)
			p.Match(PowerQueryParserAMP)
		}
		{
			p.SetState(334)
			p.Additive_expression()
		}

	}

	return localctx
}

// IMultiplicative_expressionContext is an interface to support dynamic dispatch.
type IMultiplicative_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplicative_expressionContext differentiates from other interfaces.
	IsMultiplicative_expressionContext()
}

type Multiplicative_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicative_expressionContext() *Multiplicative_expressionContext {
	var p = new(Multiplicative_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_multiplicative_expression
	return p
}

func (*Multiplicative_expressionContext) IsMultiplicative_expressionContext() {}

func NewMultiplicative_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Multiplicative_expressionContext {
	var p = new(Multiplicative_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_multiplicative_expression

	return p
}

func (s *Multiplicative_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Multiplicative_expressionContext) Metadata_expression() IMetadata_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadata_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetadata_expressionContext)
}

func (s *Multiplicative_expressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserSTAR, 0)
}

func (s *Multiplicative_expressionContext) Multiplicative_expression() IMultiplicative_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplicative_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiplicative_expressionContext)
}

func (s *Multiplicative_expressionContext) SLASH() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserSLASH, 0)
}

func (s *Multiplicative_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Multiplicative_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Multiplicative_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterMultiplicative_expression(s)
	}
}

func (s *Multiplicative_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitMultiplicative_expression(s)
	}
}

func (p *PowerQueryParser) Multiplicative_expression() (localctx IMultiplicative_expressionContext) {
	this := p
	_ = this

	localctx = NewMultiplicative_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PowerQueryParserRULE_multiplicative_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(347)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)
			p.Metadata_expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(339)
			p.Metadata_expression()
		}
		{
			p.SetState(340)
			p.Match(PowerQueryParserSTAR)
		}
		{
			p.SetState(341)
			p.Multiplicative_expression()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(343)
			p.Metadata_expression()
		}
		{
			p.SetState(344)
			p.Match(PowerQueryParserSLASH)
		}
		{
			p.SetState(345)
			p.Multiplicative_expression()
		}

	}

	return localctx
}

// IMetadata_expressionContext is an interface to support dynamic dispatch.
type IMetadata_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMetadata_expressionContext differentiates from other interfaces.
	IsMetadata_expressionContext()
}

type Metadata_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetadata_expressionContext() *Metadata_expressionContext {
	var p = new(Metadata_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_metadata_expression
	return p
}

func (*Metadata_expressionContext) IsMetadata_expressionContext() {}

func NewMetadata_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Metadata_expressionContext {
	var p = new(Metadata_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_metadata_expression

	return p
}

func (s *Metadata_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Metadata_expressionContext) AllUnary_expression() []IUnary_expressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUnary_expressionContext)(nil)).Elem())
	var tst = make([]IUnary_expressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUnary_expressionContext)
		}
	}

	return tst
}

func (s *Metadata_expressionContext) Unary_expression(i int) IUnary_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnary_expressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUnary_expressionContext)
}

func (s *Metadata_expressionContext) META() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserMETA, 0)
}

func (s *Metadata_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Metadata_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Metadata_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterMetadata_expression(s)
	}
}

func (s *Metadata_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitMetadata_expression(s)
	}
}

func (p *PowerQueryParser) Metadata_expression() (localctx IMetadata_expressionContext) {
	this := p
	_ = this

	localctx = NewMetadata_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PowerQueryParserRULE_metadata_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(354)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(349)
			p.Unary_expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(350)
			p.Unary_expression()
		}
		{
			p.SetState(351)
			p.Match(PowerQueryParserMETA)
		}
		{
			p.SetState(352)
			p.Unary_expression()
		}

	}

	return localctx
}

// IUnary_expressionContext is an interface to support dynamic dispatch.
type IUnary_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnary_expressionContext differentiates from other interfaces.
	IsUnary_expressionContext()
}

type Unary_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnary_expressionContext() *Unary_expressionContext {
	var p = new(Unary_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_unary_expression
	return p
}

func (*Unary_expressionContext) IsUnary_expressionContext() {}

func NewUnary_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Unary_expressionContext {
	var p = new(Unary_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_unary_expression

	return p
}

func (s *Unary_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Unary_expressionContext) Type_expression() IType_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_expressionContext)
}

func (s *Unary_expressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserPLUS, 0)
}

func (s *Unary_expressionContext) Unary_expression() IUnary_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnary_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnary_expressionContext)
}

func (s *Unary_expressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserMINUS, 0)
}

func (s *Unary_expressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserNOT, 0)
}

func (s *Unary_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Unary_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Unary_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterUnary_expression(s)
	}
}

func (s *Unary_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitUnary_expression(s)
	}
}

func (p *PowerQueryParser) Unary_expression() (localctx IUnary_expressionContext) {
	this := p
	_ = this

	localctx = NewUnary_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PowerQueryParserRULE_unary_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(363)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PowerQueryParserOPEN_BRACKET, PowerQueryParserOPEN_BRACE, PowerQueryParserOPEN_PAREN, PowerQueryParserELLIPSES, PowerQueryParserTYPE, PowerQueryParserAT, PowerQueryParserLITERAL, PowerQueryParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(356)
			p.Type_expression()
		}

	case PowerQueryParserPLUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.Match(PowerQueryParserPLUS)
		}
		{
			p.SetState(358)
			p.Unary_expression()
		}

	case PowerQueryParserMINUS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(359)
			p.Match(PowerQueryParserMINUS)
		}
		{
			p.SetState(360)
			p.Unary_expression()
		}

	case PowerQueryParserNOT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(361)
			p.Match(PowerQueryParserNOT)
		}
		{
			p.SetState(362)
			p.Unary_expression()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimary_expressionContext is an interface to support dynamic dispatch.
type IPrimary_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimary_expressionContext differentiates from other interfaces.
	IsPrimary_expressionContext()
}

type Primary_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimary_expressionContext() *Primary_expressionContext {
	var p = new(Primary_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_primary_expression
	return p
}

func (*Primary_expressionContext) IsPrimary_expressionContext() {}

func NewPrimary_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primary_expressionContext {
	var p = new(Primary_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_primary_expression

	return p
}

func (s *Primary_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Primary_expressionContext) Literal_expression() ILiteral_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteral_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteral_expressionContext)
}

func (s *Primary_expressionContext) List_expression() IList_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_expressionContext)
}

func (s *Primary_expressionContext) Record_expression() IRecord_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecord_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecord_expressionContext)
}

func (s *Primary_expressionContext) Identifier_expression() IIdentifier_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifier_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifier_expressionContext)
}

func (s *Primary_expressionContext) Section_access_expression() ISection_access_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISection_access_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISection_access_expressionContext)
}

func (s *Primary_expressionContext) Parenthesized_expression() IParenthesized_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParenthesized_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParenthesized_expressionContext)
}

func (s *Primary_expressionContext) Implicit_target_field_selection() IImplicit_target_field_selectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplicit_target_field_selectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplicit_target_field_selectionContext)
}

func (s *Primary_expressionContext) Implicit_target_projection() IImplicit_target_projectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplicit_target_projectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplicit_target_projectionContext)
}

func (s *Primary_expressionContext) Not_implemented_expression() INot_implemented_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INot_implemented_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INot_implemented_expressionContext)
}

func (s *Primary_expressionContext) Primary_expression() IPrimary_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimary_expressionContext)
}

func (s *Primary_expressionContext) Field_selector() IField_selectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_selectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_selectorContext)
}

func (s *Primary_expressionContext) Required_projection() IRequired_projectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequired_projectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequired_projectionContext)
}

func (s *Primary_expressionContext) Optional_projection() IOptional_projectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptional_projectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptional_projectionContext)
}

func (s *Primary_expressionContext) OPEN_BRACE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_BRACE, 0)
}

func (s *Primary_expressionContext) Item_selector() IItem_selectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItem_selectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItem_selectorContext)
}

func (s *Primary_expressionContext) CLOSE_BRACE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_BRACE, 0)
}

func (s *Primary_expressionContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPTIONAL, 0)
}

func (s *Primary_expressionContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_PAREN, 0)
}

func (s *Primary_expressionContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_PAREN, 0)
}

func (s *Primary_expressionContext) Argument_list() IArgument_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgument_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgument_listContext)
}

func (s *Primary_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primary_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterPrimary_expression(s)
	}
}

func (s *Primary_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitPrimary_expression(s)
	}
}

func (p *PowerQueryParser) Primary_expression() (localctx IPrimary_expressionContext) {
	return p.primary_expression(0)
}

func (p *PowerQueryParser) primary_expression(_p int) (localctx IPrimary_expressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimary_expressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimary_expressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 40
	p.EnterRecursionRule(localctx, 40, PowerQueryParserRULE_primary_expression, _p)
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
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(366)
			p.Literal_expression()
		}

	case 2:
		{
			p.SetState(367)
			p.List_expression()
		}

	case 3:
		{
			p.SetState(368)
			p.Record_expression()
		}

	case 4:
		{
			p.SetState(369)
			p.Identifier_expression()
		}

	case 5:
		{
			p.SetState(370)
			p.Section_access_expression()
		}

	case 6:
		{
			p.SetState(371)
			p.Parenthesized_expression()
		}

	case 7:
		{
			p.SetState(372)
			p.Implicit_target_field_selection()
		}

	case 8:
		{
			p.SetState(373)
			p.Implicit_target_projection()
		}

	case 9:
		{
			p.SetState(374)
			p.Not_implemented_expression()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(400)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
			case 1:
				localctx = NewPrimary_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PowerQueryParserRULE_primary_expression)
				p.SetState(377)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(378)
					p.Field_selector()
				}

			case 2:
				localctx = NewPrimary_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PowerQueryParserRULE_primary_expression)
				p.SetState(379)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(380)
					p.Required_projection()
				}

			case 3:
				localctx = NewPrimary_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PowerQueryParserRULE_primary_expression)
				p.SetState(381)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(382)
					p.Optional_projection()
				}

			case 4:
				localctx = NewPrimary_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PowerQueryParserRULE_primary_expression)
				p.SetState(383)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(384)
					p.Match(PowerQueryParserOPEN_BRACE)
				}
				{
					p.SetState(385)
					p.Item_selector()
				}
				{
					p.SetState(386)
					p.Match(PowerQueryParserCLOSE_BRACE)
				}

			case 5:
				localctx = NewPrimary_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PowerQueryParserRULE_primary_expression)
				p.SetState(388)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(389)
					p.Match(PowerQueryParserOPEN_BRACE)
				}
				{
					p.SetState(390)
					p.Item_selector()
				}
				{
					p.SetState(391)
					p.Match(PowerQueryParserCLOSE_BRACE)
				}
				{
					p.SetState(392)
					p.Match(PowerQueryParserOPTIONAL)
				}

			case 6:
				localctx = NewPrimary_expressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, PowerQueryParserRULE_primary_expression)
				p.SetState(394)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(395)
					p.Match(PowerQueryParserOPEN_PAREN)
				}
				p.SetState(397)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PowerQueryParserOPEN_BRACKET)|(1<<PowerQueryParserOPEN_BRACE)|(1<<PowerQueryParserOPEN_PAREN)|(1<<PowerQueryParserTRY)|(1<<PowerQueryParserERROR)|(1<<PowerQueryParserELLIPSES)|(1<<PowerQueryParserTYPE)|(1<<PowerQueryParserEACH)|(1<<PowerQueryParserLET)|(1<<PowerQueryParserIF))) != 0) || (((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(PowerQueryParserAT-48))|(1<<(PowerQueryParserNOT-48))|(1<<(PowerQueryParserPLUS-48))|(1<<(PowerQueryParserMINUS-48))|(1<<(PowerQueryParserLITERAL-48))|(1<<(PowerQueryParserIDENTIFIER-48)))) != 0) {
					{
						p.SetState(396)
						p.Argument_list()
					}

				}
				{
					p.SetState(399)
					p.Match(PowerQueryParserCLOSE_PAREN)
				}

			}

		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
	}

	return localctx
}

// ILiteral_expressionContext is an interface to support dynamic dispatch.
type ILiteral_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteral_expressionContext differentiates from other interfaces.
	IsLiteral_expressionContext()
}

type Literal_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_expressionContext() *Literal_expressionContext {
	var p = new(Literal_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_literal_expression
	return p
}

func (*Literal_expressionContext) IsLiteral_expressionContext() {}

func NewLiteral_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_expressionContext {
	var p = new(Literal_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_literal_expression

	return p
}

func (s *Literal_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_expressionContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserLITERAL, 0)
}

func (s *Literal_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterLiteral_expression(s)
	}
}

func (s *Literal_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitLiteral_expression(s)
	}
}

func (p *PowerQueryParser) Literal_expression() (localctx ILiteral_expressionContext) {
	this := p
	_ = this

	localctx = NewLiteral_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PowerQueryParserRULE_literal_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PowerQueryParserLITERAL)
	}

	return localctx
}

// IIdentifier_expressionContext is an interface to support dynamic dispatch.
type IIdentifier_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifier_expressionContext differentiates from other interfaces.
	IsIdentifier_expressionContext()
}

type Identifier_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifier_expressionContext() *Identifier_expressionContext {
	var p = new(Identifier_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_identifier_expression
	return p
}

func (*Identifier_expressionContext) IsIdentifier_expressionContext() {}

func NewIdentifier_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Identifier_expressionContext {
	var p = new(Identifier_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_identifier_expression

	return p
}

func (s *Identifier_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Identifier_expressionContext) Identifier_reference() IIdentifier_referenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifier_referenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifier_referenceContext)
}

func (s *Identifier_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Identifier_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Identifier_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterIdentifier_expression(s)
	}
}

func (s *Identifier_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitIdentifier_expression(s)
	}
}

func (p *PowerQueryParser) Identifier_expression() (localctx IIdentifier_expressionContext) {
	this := p
	_ = this

	localctx = NewIdentifier_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, PowerQueryParserRULE_identifier_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(407)
		p.Identifier_reference()
	}

	return localctx
}

// IIdentifier_referenceContext is an interface to support dynamic dispatch.
type IIdentifier_referenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifier_referenceContext differentiates from other interfaces.
	IsIdentifier_referenceContext()
}

type Identifier_referenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifier_referenceContext() *Identifier_referenceContext {
	var p = new(Identifier_referenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_identifier_reference
	return p
}

func (*Identifier_referenceContext) IsIdentifier_referenceContext() {}

func NewIdentifier_referenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Identifier_referenceContext {
	var p = new(Identifier_referenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_identifier_reference

	return p
}

func (s *Identifier_referenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Identifier_referenceContext) Exclusive_identifier_reference() IExclusive_identifier_referenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExclusive_identifier_referenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExclusive_identifier_referenceContext)
}

func (s *Identifier_referenceContext) Inclusive_identifier_reference() IInclusive_identifier_referenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInclusive_identifier_referenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInclusive_identifier_referenceContext)
}

func (s *Identifier_referenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Identifier_referenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Identifier_referenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterIdentifier_reference(s)
	}
}

func (s *Identifier_referenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitIdentifier_reference(s)
	}
}

func (p *PowerQueryParser) Identifier_reference() (localctx IIdentifier_referenceContext) {
	this := p
	_ = this

	localctx = NewIdentifier_referenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, PowerQueryParserRULE_identifier_reference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	case PowerQueryParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.Exclusive_identifier_reference()
		}

	case PowerQueryParserAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(410)
			p.Inclusive_identifier_reference()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExclusive_identifier_referenceContext is an interface to support dynamic dispatch.
type IExclusive_identifier_referenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExclusive_identifier_referenceContext differentiates from other interfaces.
	IsExclusive_identifier_referenceContext()
}

type Exclusive_identifier_referenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExclusive_identifier_referenceContext() *Exclusive_identifier_referenceContext {
	var p = new(Exclusive_identifier_referenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_exclusive_identifier_reference
	return p
}

func (*Exclusive_identifier_referenceContext) IsExclusive_identifier_referenceContext() {}

func NewExclusive_identifier_referenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Exclusive_identifier_referenceContext {
	var p = new(Exclusive_identifier_referenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_exclusive_identifier_reference

	return p
}

func (s *Exclusive_identifier_referenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Exclusive_identifier_referenceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserIDENTIFIER, 0)
}

func (s *Exclusive_identifier_referenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Exclusive_identifier_referenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Exclusive_identifier_referenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterExclusive_identifier_reference(s)
	}
}

func (s *Exclusive_identifier_referenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitExclusive_identifier_reference(s)
	}
}

func (p *PowerQueryParser) Exclusive_identifier_reference() (localctx IExclusive_identifier_referenceContext) {
	this := p
	_ = this

	localctx = NewExclusive_identifier_referenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, PowerQueryParserRULE_exclusive_identifier_reference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PowerQueryParserIDENTIFIER)
	}

	return localctx
}

// IInclusive_identifier_referenceContext is an interface to support dynamic dispatch.
type IInclusive_identifier_referenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInclusive_identifier_referenceContext differentiates from other interfaces.
	IsInclusive_identifier_referenceContext()
}

type Inclusive_identifier_referenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInclusive_identifier_referenceContext() *Inclusive_identifier_referenceContext {
	var p = new(Inclusive_identifier_referenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_inclusive_identifier_reference
	return p
}

func (*Inclusive_identifier_referenceContext) IsInclusive_identifier_referenceContext() {}

func NewInclusive_identifier_referenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inclusive_identifier_referenceContext {
	var p = new(Inclusive_identifier_referenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_inclusive_identifier_reference

	return p
}

func (s *Inclusive_identifier_referenceContext) GetParser() antlr.Parser { return s.parser }

func (s *Inclusive_identifier_referenceContext) AT() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserAT, 0)
}

func (s *Inclusive_identifier_referenceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserIDENTIFIER, 0)
}

func (s *Inclusive_identifier_referenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inclusive_identifier_referenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Inclusive_identifier_referenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterInclusive_identifier_reference(s)
	}
}

func (s *Inclusive_identifier_referenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitInclusive_identifier_reference(s)
	}
}

func (p *PowerQueryParser) Inclusive_identifier_reference() (localctx IInclusive_identifier_referenceContext) {
	this := p
	_ = this

	localctx = NewInclusive_identifier_referenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, PowerQueryParserRULE_inclusive_identifier_reference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(415)
		p.Match(PowerQueryParserAT)
	}
	{
		p.SetState(416)
		p.Match(PowerQueryParserIDENTIFIER)
	}

	return localctx
}

// ISection_access_expressionContext is an interface to support dynamic dispatch.
type ISection_access_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSection_access_expressionContext differentiates from other interfaces.
	IsSection_access_expressionContext()
}

type Section_access_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySection_access_expressionContext() *Section_access_expressionContext {
	var p = new(Section_access_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_section_access_expression
	return p
}

func (*Section_access_expressionContext) IsSection_access_expressionContext() {}

func NewSection_access_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Section_access_expressionContext {
	var p = new(Section_access_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_section_access_expression

	return p
}

func (s *Section_access_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Section_access_expressionContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PowerQueryParserIDENTIFIER)
}

func (s *Section_access_expressionContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PowerQueryParserIDENTIFIER, i)
}

func (s *Section_access_expressionContext) BANG() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserBANG, 0)
}

func (s *Section_access_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Section_access_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Section_access_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterSection_access_expression(s)
	}
}

func (s *Section_access_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitSection_access_expression(s)
	}
}

func (p *PowerQueryParser) Section_access_expression() (localctx ISection_access_expressionContext) {
	this := p
	_ = this

	localctx = NewSection_access_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, PowerQueryParserRULE_section_access_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(418)
		p.Match(PowerQueryParserIDENTIFIER)
	}
	{
		p.SetState(419)
		p.Match(PowerQueryParserBANG)
	}
	{
		p.SetState(420)
		p.Match(PowerQueryParserIDENTIFIER)
	}

	return localctx
}

// IParenthesized_expressionContext is an interface to support dynamic dispatch.
type IParenthesized_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParenthesized_expressionContext differentiates from other interfaces.
	IsParenthesized_expressionContext()
}

type Parenthesized_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParenthesized_expressionContext() *Parenthesized_expressionContext {
	var p = new(Parenthesized_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_parenthesized_expression
	return p
}

func (*Parenthesized_expressionContext) IsParenthesized_expressionContext() {}

func NewParenthesized_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parenthesized_expressionContext {
	var p = new(Parenthesized_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_parenthesized_expression

	return p
}

func (s *Parenthesized_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Parenthesized_expressionContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_PAREN, 0)
}

func (s *Parenthesized_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Parenthesized_expressionContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_PAREN, 0)
}

func (s *Parenthesized_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parenthesized_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parenthesized_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterParenthesized_expression(s)
	}
}

func (s *Parenthesized_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitParenthesized_expression(s)
	}
}

func (p *PowerQueryParser) Parenthesized_expression() (localctx IParenthesized_expressionContext) {
	this := p
	_ = this

	localctx = NewParenthesized_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, PowerQueryParserRULE_parenthesized_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(422)
		p.Match(PowerQueryParserOPEN_PAREN)
	}
	{
		p.SetState(423)
		p.Expression()
	}
	{
		p.SetState(424)
		p.Match(PowerQueryParserCLOSE_PAREN)
	}

	return localctx
}

// INot_implemented_expressionContext is an interface to support dynamic dispatch.
type INot_implemented_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNot_implemented_expressionContext differentiates from other interfaces.
	IsNot_implemented_expressionContext()
}

type Not_implemented_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNot_implemented_expressionContext() *Not_implemented_expressionContext {
	var p = new(Not_implemented_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_not_implemented_expression
	return p
}

func (*Not_implemented_expressionContext) IsNot_implemented_expressionContext() {}

func NewNot_implemented_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Not_implemented_expressionContext {
	var p = new(Not_implemented_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_not_implemented_expression

	return p
}

func (s *Not_implemented_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Not_implemented_expressionContext) ELLIPSES() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserELLIPSES, 0)
}

func (s *Not_implemented_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Not_implemented_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Not_implemented_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterNot_implemented_expression(s)
	}
}

func (s *Not_implemented_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitNot_implemented_expression(s)
	}
}

func (p *PowerQueryParser) Not_implemented_expression() (localctx INot_implemented_expressionContext) {
	this := p
	_ = this

	localctx = NewNot_implemented_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, PowerQueryParserRULE_not_implemented_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PowerQueryParserELLIPSES)
	}

	return localctx
}

// IArgument_listContext is an interface to support dynamic dispatch.
type IArgument_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgument_listContext differentiates from other interfaces.
	IsArgument_listContext()
}

type Argument_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgument_listContext() *Argument_listContext {
	var p = new(Argument_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_argument_list
	return p
}

func (*Argument_listContext) IsArgument_listContext() {}

func NewArgument_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Argument_listContext {
	var p = new(Argument_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_argument_list

	return p
}

func (s *Argument_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Argument_listContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Argument_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Argument_listContext) Argument_list() IArgument_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgument_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgument_listContext)
}

func (s *Argument_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Argument_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Argument_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterArgument_list(s)
	}
}

func (s *Argument_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitArgument_list(s)
	}
}

func (p *PowerQueryParser) Argument_list() (localctx IArgument_listContext) {
	this := p
	_ = this

	localctx = NewArgument_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, PowerQueryParserRULE_argument_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(428)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(429)
			p.Expression()
		}
		{
			p.SetState(430)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(431)
			p.Argument_list()
		}

	}

	return localctx
}

// IList_expressionContext is an interface to support dynamic dispatch.
type IList_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_expressionContext differentiates from other interfaces.
	IsList_expressionContext()
}

type List_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_expressionContext() *List_expressionContext {
	var p = new(List_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_list_expression
	return p
}

func (*List_expressionContext) IsList_expressionContext() {}

func NewList_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_expressionContext {
	var p = new(List_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_list_expression

	return p
}

func (s *List_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *List_expressionContext) OPEN_BRACE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_BRACE, 0)
}

func (s *List_expressionContext) CLOSE_BRACE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_BRACE, 0)
}

func (s *List_expressionContext) Item_list() IItem_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItem_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItem_listContext)
}

func (s *List_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterList_expression(s)
	}
}

func (s *List_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitList_expression(s)
	}
}

func (p *PowerQueryParser) List_expression() (localctx IList_expressionContext) {
	this := p
	_ = this

	localctx = NewList_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, PowerQueryParserRULE_list_expression)
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
		p.SetState(435)
		p.Match(PowerQueryParserOPEN_BRACE)
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PowerQueryParserOPEN_BRACKET)|(1<<PowerQueryParserOPEN_BRACE)|(1<<PowerQueryParserOPEN_PAREN)|(1<<PowerQueryParserTRY)|(1<<PowerQueryParserERROR)|(1<<PowerQueryParserELLIPSES)|(1<<PowerQueryParserTYPE)|(1<<PowerQueryParserEACH)|(1<<PowerQueryParserLET)|(1<<PowerQueryParserIF))) != 0) || (((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(PowerQueryParserAT-48))|(1<<(PowerQueryParserNOT-48))|(1<<(PowerQueryParserPLUS-48))|(1<<(PowerQueryParserMINUS-48))|(1<<(PowerQueryParserLITERAL-48))|(1<<(PowerQueryParserIDENTIFIER-48)))) != 0) {
		{
			p.SetState(436)
			p.Item_list()
		}

	}
	{
		p.SetState(439)
		p.Match(PowerQueryParserCLOSE_BRACE)
	}

	return localctx
}

// IItem_listContext is an interface to support dynamic dispatch.
type IItem_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItem_listContext differentiates from other interfaces.
	IsItem_listContext()
}

type Item_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItem_listContext() *Item_listContext {
	var p = new(Item_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_item_list
	return p
}

func (*Item_listContext) IsItem_listContext() {}

func NewItem_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Item_listContext {
	var p = new(Item_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_item_list

	return p
}

func (s *Item_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Item_listContext) Item() IItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItemContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItemContext)
}

func (s *Item_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Item_listContext) Item_list() IItem_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItem_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItem_listContext)
}

func (s *Item_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Item_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Item_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterItem_list(s)
	}
}

func (s *Item_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitItem_list(s)
	}
}

func (p *PowerQueryParser) Item_list() (localctx IItem_listContext) {
	this := p
	_ = this

	localctx = NewItem_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, PowerQueryParserRULE_item_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(441)
			p.Item()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(442)
			p.Item()
		}
		{
			p.SetState(443)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(444)
			p.Item_list()
		}

	}

	return localctx
}

// IItemContext is an interface to support dynamic dispatch.
type IItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItemContext differentiates from other interfaces.
	IsItemContext()
}

type ItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItemContext() *ItemContext {
	var p = new(ItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_item
	return p
}

func (*ItemContext) IsItemContext() {}

func NewItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ItemContext {
	var p = new(ItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_item

	return p
}

func (s *ItemContext) GetParser() antlr.Parser { return s.parser }

func (s *ItemContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ItemContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ItemContext) DOTDOT() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserDOTDOT, 0)
}

func (s *ItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterItem(s)
	}
}

func (s *ItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitItem(s)
	}
}

func (p *PowerQueryParser) Item() (localctx IItemContext) {
	this := p
	_ = this

	localctx = NewItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, PowerQueryParserRULE_item)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(453)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(448)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(449)
			p.Expression()
		}
		{
			p.SetState(450)
			p.Match(PowerQueryParserDOTDOT)
		}
		{
			p.SetState(451)
			p.Expression()
		}

	}

	return localctx
}

// IRecord_expressionContext is an interface to support dynamic dispatch.
type IRecord_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecord_expressionContext differentiates from other interfaces.
	IsRecord_expressionContext()
}

type Record_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecord_expressionContext() *Record_expressionContext {
	var p = new(Record_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_record_expression
	return p
}

func (*Record_expressionContext) IsRecord_expressionContext() {}

func NewRecord_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Record_expressionContext {
	var p = new(Record_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_record_expression

	return p
}

func (s *Record_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Record_expressionContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_BRACKET, 0)
}

func (s *Record_expressionContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_BRACKET, 0)
}

func (s *Record_expressionContext) Field_list() IField_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_listContext)
}

func (s *Record_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Record_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Record_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterRecord_expression(s)
	}
}

func (s *Record_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitRecord_expression(s)
	}
}

func (p *PowerQueryParser) Record_expression() (localctx IRecord_expressionContext) {
	this := p
	_ = this

	localctx = NewRecord_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, PowerQueryParserRULE_record_expression)
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
		p.SetState(455)
		p.Match(PowerQueryParserOPEN_BRACKET)
	}
	p.SetState(457)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserIDENTIFIER {
		{
			p.SetState(456)
			p.Field_list()
		}

	}
	{
		p.SetState(459)
		p.Match(PowerQueryParserCLOSE_BRACKET)
	}

	return localctx
}

// IField_listContext is an interface to support dynamic dispatch.
type IField_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_listContext differentiates from other interfaces.
	IsField_listContext()
}

type Field_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_listContext() *Field_listContext {
	var p = new(Field_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_field_list
	return p
}

func (*Field_listContext) IsField_listContext() {}

func NewField_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_listContext {
	var p = new(Field_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_field_list

	return p
}

func (s *Field_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_listContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Field_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Field_listContext) Field_list() IField_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_listContext)
}

func (s *Field_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterField_list(s)
	}
}

func (s *Field_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitField_list(s)
	}
}

func (p *PowerQueryParser) Field_list() (localctx IField_listContext) {
	this := p
	_ = this

	localctx = NewField_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, PowerQueryParserRULE_field_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(466)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(461)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(462)
			p.Field()
		}
		{
			p.SetState(463)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(464)
			p.Field_list()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *FieldContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserEQUALS, 0)
}

func (s *FieldContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *PowerQueryParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, PowerQueryParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Field_name()
	}
	{
		p.SetState(469)
		p.Match(PowerQueryParserEQUALS)
	}
	{
		p.SetState(470)
		p.Expression()
	}

	return localctx
}

// IField_nameContext is an interface to support dynamic dispatch.
type IField_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_nameContext differentiates from other interfaces.
	IsField_nameContext()
}

type Field_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_nameContext() *Field_nameContext {
	var p = new(Field_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_field_name
	return p
}

func (*Field_nameContext) IsField_nameContext() {}

func NewField_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_nameContext {
	var p = new(Field_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_field_name

	return p
}

func (s *Field_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserIDENTIFIER, 0)
}

func (s *Field_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterField_name(s)
	}
}

func (s *Field_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitField_name(s)
	}
}

func (p *PowerQueryParser) Field_name() (localctx IField_nameContext) {
	this := p
	_ = this

	localctx = NewField_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, PowerQueryParserRULE_field_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PowerQueryParserIDENTIFIER)
	}

	return localctx
}

// IItem_selectorContext is an interface to support dynamic dispatch.
type IItem_selectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItem_selectorContext differentiates from other interfaces.
	IsItem_selectorContext()
}

type Item_selectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItem_selectorContext() *Item_selectorContext {
	var p = new(Item_selectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_item_selector
	return p
}

func (*Item_selectorContext) IsItem_selectorContext() {}

func NewItem_selectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Item_selectorContext {
	var p = new(Item_selectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_item_selector

	return p
}

func (s *Item_selectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Item_selectorContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Item_selectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Item_selectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Item_selectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterItem_selector(s)
	}
}

func (s *Item_selectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitItem_selector(s)
	}
}

func (p *PowerQueryParser) Item_selector() (localctx IItem_selectorContext) {
	this := p
	_ = this

	localctx = NewItem_selectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, PowerQueryParserRULE_item_selector)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(474)
		p.Expression()
	}

	return localctx
}

// IField_selectorContext is an interface to support dynamic dispatch.
type IField_selectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_selectorContext differentiates from other interfaces.
	IsField_selectorContext()
}

type Field_selectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_selectorContext() *Field_selectorContext {
	var p = new(Field_selectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_field_selector
	return p
}

func (*Field_selectorContext) IsField_selectorContext() {}

func NewField_selectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_selectorContext {
	var p = new(Field_selectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_field_selector

	return p
}

func (s *Field_selectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_selectorContext) Required_field_selector() IRequired_field_selectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequired_field_selectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequired_field_selectorContext)
}

func (s *Field_selectorContext) Optional_field_selector() IOptional_field_selectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptional_field_selectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptional_field_selectorContext)
}

func (s *Field_selectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_selectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_selectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterField_selector(s)
	}
}

func (s *Field_selectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitField_selector(s)
	}
}

func (p *PowerQueryParser) Field_selector() (localctx IField_selectorContext) {
	this := p
	_ = this

	localctx = NewField_selectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, PowerQueryParserRULE_field_selector)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(478)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(476)
			p.Required_field_selector()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(477)
			p.Optional_field_selector()
		}

	}

	return localctx
}

// IRequired_field_selectorContext is an interface to support dynamic dispatch.
type IRequired_field_selectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequired_field_selectorContext differentiates from other interfaces.
	IsRequired_field_selectorContext()
}

type Required_field_selectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequired_field_selectorContext() *Required_field_selectorContext {
	var p = new(Required_field_selectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_required_field_selector
	return p
}

func (*Required_field_selectorContext) IsRequired_field_selectorContext() {}

func NewRequired_field_selectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Required_field_selectorContext {
	var p = new(Required_field_selectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_required_field_selector

	return p
}

func (s *Required_field_selectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Required_field_selectorContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_BRACKET, 0)
}

func (s *Required_field_selectorContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *Required_field_selectorContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_BRACKET, 0)
}

func (s *Required_field_selectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Required_field_selectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Required_field_selectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterRequired_field_selector(s)
	}
}

func (s *Required_field_selectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitRequired_field_selector(s)
	}
}

func (p *PowerQueryParser) Required_field_selector() (localctx IRequired_field_selectorContext) {
	this := p
	_ = this

	localctx = NewRequired_field_selectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, PowerQueryParserRULE_required_field_selector)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PowerQueryParserOPEN_BRACKET)
	}
	{
		p.SetState(481)
		p.Field_name()
	}
	{
		p.SetState(482)
		p.Match(PowerQueryParserCLOSE_BRACKET)
	}

	return localctx
}

// IOptional_field_selectorContext is an interface to support dynamic dispatch.
type IOptional_field_selectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptional_field_selectorContext differentiates from other interfaces.
	IsOptional_field_selectorContext()
}

type Optional_field_selectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptional_field_selectorContext() *Optional_field_selectorContext {
	var p = new(Optional_field_selectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_optional_field_selector
	return p
}

func (*Optional_field_selectorContext) IsOptional_field_selectorContext() {}

func NewOptional_field_selectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Optional_field_selectorContext {
	var p = new(Optional_field_selectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_optional_field_selector

	return p
}

func (s *Optional_field_selectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Optional_field_selectorContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_BRACKET, 0)
}

func (s *Optional_field_selectorContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *Optional_field_selectorContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_BRACKET, 0)
}

func (s *Optional_field_selectorContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPTIONAL, 0)
}

func (s *Optional_field_selectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Optional_field_selectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Optional_field_selectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterOptional_field_selector(s)
	}
}

func (s *Optional_field_selectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitOptional_field_selector(s)
	}
}

func (p *PowerQueryParser) Optional_field_selector() (localctx IOptional_field_selectorContext) {
	this := p
	_ = this

	localctx = NewOptional_field_selectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, PowerQueryParserRULE_optional_field_selector)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PowerQueryParserOPEN_BRACKET)
	}
	{
		p.SetState(485)
		p.Field_name()
	}
	{
		p.SetState(486)
		p.Match(PowerQueryParserCLOSE_BRACKET)
	}
	{
		p.SetState(487)
		p.Match(PowerQueryParserOPTIONAL)
	}

	return localctx
}

// IImplicit_target_field_selectionContext is an interface to support dynamic dispatch.
type IImplicit_target_field_selectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImplicit_target_field_selectionContext differentiates from other interfaces.
	IsImplicit_target_field_selectionContext()
}

type Implicit_target_field_selectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImplicit_target_field_selectionContext() *Implicit_target_field_selectionContext {
	var p = new(Implicit_target_field_selectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_implicit_target_field_selection
	return p
}

func (*Implicit_target_field_selectionContext) IsImplicit_target_field_selectionContext() {}

func NewImplicit_target_field_selectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Implicit_target_field_selectionContext {
	var p = new(Implicit_target_field_selectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_implicit_target_field_selection

	return p
}

func (s *Implicit_target_field_selectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Implicit_target_field_selectionContext) Field_selector() IField_selectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_selectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_selectorContext)
}

func (s *Implicit_target_field_selectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Implicit_target_field_selectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Implicit_target_field_selectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterImplicit_target_field_selection(s)
	}
}

func (s *Implicit_target_field_selectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitImplicit_target_field_selection(s)
	}
}

func (p *PowerQueryParser) Implicit_target_field_selection() (localctx IImplicit_target_field_selectionContext) {
	this := p
	_ = this

	localctx = NewImplicit_target_field_selectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, PowerQueryParserRULE_implicit_target_field_selection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Field_selector()
	}

	return localctx
}

// IRequired_projectionContext is an interface to support dynamic dispatch.
type IRequired_projectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequired_projectionContext differentiates from other interfaces.
	IsRequired_projectionContext()
}

type Required_projectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequired_projectionContext() *Required_projectionContext {
	var p = new(Required_projectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_required_projection
	return p
}

func (*Required_projectionContext) IsRequired_projectionContext() {}

func NewRequired_projectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Required_projectionContext {
	var p = new(Required_projectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_required_projection

	return p
}

func (s *Required_projectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Required_projectionContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_BRACKET, 0)
}

func (s *Required_projectionContext) Required_selector_list() IRequired_selector_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequired_selector_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequired_selector_listContext)
}

func (s *Required_projectionContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_BRACKET, 0)
}

func (s *Required_projectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Required_projectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Required_projectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterRequired_projection(s)
	}
}

func (s *Required_projectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitRequired_projection(s)
	}
}

func (p *PowerQueryParser) Required_projection() (localctx IRequired_projectionContext) {
	this := p
	_ = this

	localctx = NewRequired_projectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, PowerQueryParserRULE_required_projection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(491)
		p.Match(PowerQueryParserOPEN_BRACKET)
	}
	{
		p.SetState(492)
		p.Required_selector_list()
	}
	{
		p.SetState(493)
		p.Match(PowerQueryParserCLOSE_BRACKET)
	}

	return localctx
}

// IOptional_projectionContext is an interface to support dynamic dispatch.
type IOptional_projectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptional_projectionContext differentiates from other interfaces.
	IsOptional_projectionContext()
}

type Optional_projectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptional_projectionContext() *Optional_projectionContext {
	var p = new(Optional_projectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_optional_projection
	return p
}

func (*Optional_projectionContext) IsOptional_projectionContext() {}

func NewOptional_projectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Optional_projectionContext {
	var p = new(Optional_projectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_optional_projection

	return p
}

func (s *Optional_projectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Optional_projectionContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_BRACKET, 0)
}

func (s *Optional_projectionContext) Required_selector_list() IRequired_selector_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequired_selector_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequired_selector_listContext)
}

func (s *Optional_projectionContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_BRACKET, 0)
}

func (s *Optional_projectionContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPTIONAL, 0)
}

func (s *Optional_projectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Optional_projectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Optional_projectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterOptional_projection(s)
	}
}

func (s *Optional_projectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitOptional_projection(s)
	}
}

func (p *PowerQueryParser) Optional_projection() (localctx IOptional_projectionContext) {
	this := p
	_ = this

	localctx = NewOptional_projectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, PowerQueryParserRULE_optional_projection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PowerQueryParserOPEN_BRACKET)
	}
	{
		p.SetState(496)
		p.Required_selector_list()
	}
	{
		p.SetState(497)
		p.Match(PowerQueryParserCLOSE_BRACKET)
	}
	{
		p.SetState(498)
		p.Match(PowerQueryParserOPTIONAL)
	}

	return localctx
}

// IRequired_selector_listContext is an interface to support dynamic dispatch.
type IRequired_selector_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequired_selector_listContext differentiates from other interfaces.
	IsRequired_selector_listContext()
}

type Required_selector_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequired_selector_listContext() *Required_selector_listContext {
	var p = new(Required_selector_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_required_selector_list
	return p
}

func (*Required_selector_listContext) IsRequired_selector_listContext() {}

func NewRequired_selector_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Required_selector_listContext {
	var p = new(Required_selector_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_required_selector_list

	return p
}

func (s *Required_selector_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Required_selector_listContext) Required_field_selector() IRequired_field_selectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequired_field_selectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequired_field_selectorContext)
}

func (s *Required_selector_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Required_selector_listContext) Required_selector_list() IRequired_selector_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequired_selector_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequired_selector_listContext)
}

func (s *Required_selector_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Required_selector_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Required_selector_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterRequired_selector_list(s)
	}
}

func (s *Required_selector_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitRequired_selector_list(s)
	}
}

func (p *PowerQueryParser) Required_selector_list() (localctx IRequired_selector_listContext) {
	this := p
	_ = this

	localctx = NewRequired_selector_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, PowerQueryParserRULE_required_selector_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(500)
			p.Required_field_selector()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(501)
			p.Required_field_selector()
		}
		{
			p.SetState(502)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(503)
			p.Required_selector_list()
		}

	}

	return localctx
}

// IImplicit_target_projectionContext is an interface to support dynamic dispatch.
type IImplicit_target_projectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImplicit_target_projectionContext differentiates from other interfaces.
	IsImplicit_target_projectionContext()
}

type Implicit_target_projectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImplicit_target_projectionContext() *Implicit_target_projectionContext {
	var p = new(Implicit_target_projectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_implicit_target_projection
	return p
}

func (*Implicit_target_projectionContext) IsImplicit_target_projectionContext() {}

func NewImplicit_target_projectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Implicit_target_projectionContext {
	var p = new(Implicit_target_projectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_implicit_target_projection

	return p
}

func (s *Implicit_target_projectionContext) GetParser() antlr.Parser { return s.parser }

func (s *Implicit_target_projectionContext) Required_projection() IRequired_projectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequired_projectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequired_projectionContext)
}

func (s *Implicit_target_projectionContext) Optional_projection() IOptional_projectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptional_projectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptional_projectionContext)
}

func (s *Implicit_target_projectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Implicit_target_projectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Implicit_target_projectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterImplicit_target_projection(s)
	}
}

func (s *Implicit_target_projectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitImplicit_target_projection(s)
	}
}

func (p *PowerQueryParser) Implicit_target_projection() (localctx IImplicit_target_projectionContext) {
	this := p
	_ = this

	localctx = NewImplicit_target_projectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, PowerQueryParserRULE_implicit_target_projection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(509)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(507)
			p.Required_projection()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(508)
			p.Optional_projection()
		}

	}

	return localctx
}

// IFunction_expressionContext is an interface to support dynamic dispatch.
type IFunction_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_expressionContext differentiates from other interfaces.
	IsFunction_expressionContext()
}

type Function_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_expressionContext() *Function_expressionContext {
	var p = new(Function_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_function_expression
	return p
}

func (*Function_expressionContext) IsFunction_expressionContext() {}

func NewFunction_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_expressionContext {
	var p = new(Function_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_function_expression

	return p
}

func (s *Function_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_expressionContext) OPEN_PAREN() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_PAREN, 0)
}

func (s *Function_expressionContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_PAREN, 0)
}

func (s *Function_expressionContext) ARROW() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserARROW, 0)
}

func (s *Function_expressionContext) Function_body() IFunction_bodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_bodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_bodyContext)
}

func (s *Function_expressionContext) Parameter_list() IParameter_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameter_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameter_listContext)
}

func (s *Function_expressionContext) Return_type() IReturn_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_typeContext)
}

func (s *Function_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterFunction_expression(s)
	}
}

func (s *Function_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitFunction_expression(s)
	}
}

func (p *PowerQueryParser) Function_expression() (localctx IFunction_expressionContext) {
	this := p
	_ = this

	localctx = NewFunction_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, PowerQueryParserRULE_function_expression)
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
		p.SetState(511)
		p.Match(PowerQueryParserOPEN_PAREN)
	}
	p.SetState(513)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserIDENTIFIER {
		{
			p.SetState(512)
			p.Parameter_list()
		}

	}
	{
		p.SetState(515)
		p.Match(PowerQueryParserCLOSE_PAREN)
	}
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserAS {
		{
			p.SetState(516)
			p.Return_type()
		}

	}
	{
		p.SetState(519)
		p.Match(PowerQueryParserARROW)
	}
	{
		p.SetState(520)
		p.Function_body()
	}

	return localctx
}

// IFunction_bodyContext is an interface to support dynamic dispatch.
type IFunction_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_bodyContext differentiates from other interfaces.
	IsFunction_bodyContext()
}

type Function_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_bodyContext() *Function_bodyContext {
	var p = new(Function_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_function_body
	return p
}

func (*Function_bodyContext) IsFunction_bodyContext() {}

func NewFunction_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_bodyContext {
	var p = new(Function_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_function_body

	return p
}

func (s *Function_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_bodyContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Function_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterFunction_body(s)
	}
}

func (s *Function_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitFunction_body(s)
	}
}

func (p *PowerQueryParser) Function_body() (localctx IFunction_bodyContext) {
	this := p
	_ = this

	localctx = NewFunction_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, PowerQueryParserRULE_function_body)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(522)
		p.Expression()
	}

	return localctx
}

// IParameter_listContext is an interface to support dynamic dispatch.
type IParameter_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameter_listContext differentiates from other interfaces.
	IsParameter_listContext()
}

type Parameter_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameter_listContext() *Parameter_listContext {
	var p = new(Parameter_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_parameter_list
	return p
}

func (*Parameter_listContext) IsParameter_listContext() {}

func NewParameter_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parameter_listContext {
	var p = new(Parameter_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_parameter_list

	return p
}

func (s *Parameter_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Parameter_listContext) Fixed_parameter_list() IFixed_parameter_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFixed_parameter_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFixed_parameter_listContext)
}

func (s *Parameter_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Parameter_listContext) Optional_parameter_list() IOptional_parameter_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptional_parameter_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptional_parameter_listContext)
}

func (s *Parameter_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parameter_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parameter_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterParameter_list(s)
	}
}

func (s *Parameter_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitParameter_list(s)
	}
}

func (p *PowerQueryParser) Parameter_list() (localctx IParameter_listContext) {
	this := p
	_ = this

	localctx = NewParameter_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, PowerQueryParserRULE_parameter_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(529)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(524)
			p.Fixed_parameter_list()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(525)
			p.Fixed_parameter_list()
		}
		{
			p.SetState(526)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(527)
			p.Optional_parameter_list()
		}

	}

	return localctx
}

// IFixed_parameter_listContext is an interface to support dynamic dispatch.
type IFixed_parameter_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFixed_parameter_listContext differentiates from other interfaces.
	IsFixed_parameter_listContext()
}

type Fixed_parameter_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFixed_parameter_listContext() *Fixed_parameter_listContext {
	var p = new(Fixed_parameter_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_fixed_parameter_list
	return p
}

func (*Fixed_parameter_listContext) IsFixed_parameter_listContext() {}

func NewFixed_parameter_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Fixed_parameter_listContext {
	var p = new(Fixed_parameter_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_fixed_parameter_list

	return p
}

func (s *Fixed_parameter_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Fixed_parameter_listContext) Parameter() IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *Fixed_parameter_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Fixed_parameter_listContext) Fixed_parameter_list() IFixed_parameter_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFixed_parameter_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFixed_parameter_listContext)
}

func (s *Fixed_parameter_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Fixed_parameter_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Fixed_parameter_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterFixed_parameter_list(s)
	}
}

func (s *Fixed_parameter_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitFixed_parameter_list(s)
	}
}

func (p *PowerQueryParser) Fixed_parameter_list() (localctx IFixed_parameter_listContext) {
	this := p
	_ = this

	localctx = NewFixed_parameter_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, PowerQueryParserRULE_fixed_parameter_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(536)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(531)
			p.Parameter()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(532)
			p.Parameter()
		}
		{
			p.SetState(533)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(534)
			p.Fixed_parameter_list()
		}

	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) Parameter_name() IParameter_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameter_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameter_nameContext)
}

func (s *ParameterContext) Parameter_type() IParameter_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameter_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameter_typeContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *PowerQueryParser) Parameter() (localctx IParameterContext) {
	this := p
	_ = this

	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, PowerQueryParserRULE_parameter)
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
		p.SetState(538)
		p.Parameter_name()
	}
	p.SetState(540)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserAS {
		{
			p.SetState(539)
			p.Parameter_type()
		}

	}

	return localctx
}

// IParameter_nameContext is an interface to support dynamic dispatch.
type IParameter_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameter_nameContext differentiates from other interfaces.
	IsParameter_nameContext()
}

type Parameter_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameter_nameContext() *Parameter_nameContext {
	var p = new(Parameter_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_parameter_name
	return p
}

func (*Parameter_nameContext) IsParameter_nameContext() {}

func NewParameter_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parameter_nameContext {
	var p = new(Parameter_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_parameter_name

	return p
}

func (s *Parameter_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Parameter_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserIDENTIFIER, 0)
}

func (s *Parameter_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parameter_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parameter_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterParameter_name(s)
	}
}

func (s *Parameter_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitParameter_name(s)
	}
}

func (p *PowerQueryParser) Parameter_name() (localctx IParameter_nameContext) {
	this := p
	_ = this

	localctx = NewParameter_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, PowerQueryParserRULE_parameter_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PowerQueryParserIDENTIFIER)
	}

	return localctx
}

// IParameter_typeContext is an interface to support dynamic dispatch.
type IParameter_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameter_typeContext differentiates from other interfaces.
	IsParameter_typeContext()
}

type Parameter_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameter_typeContext() *Parameter_typeContext {
	var p = new(Parameter_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_parameter_type
	return p
}

func (*Parameter_typeContext) IsParameter_typeContext() {}

func NewParameter_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parameter_typeContext {
	var p = new(Parameter_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_parameter_type

	return p
}

func (s *Parameter_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Parameter_typeContext) Assertion() IAssertionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssertionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssertionContext)
}

func (s *Parameter_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parameter_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parameter_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterParameter_type(s)
	}
}

func (s *Parameter_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitParameter_type(s)
	}
}

func (p *PowerQueryParser) Parameter_type() (localctx IParameter_typeContext) {
	this := p
	_ = this

	localctx = NewParameter_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, PowerQueryParserRULE_parameter_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(544)
		p.Assertion()
	}

	return localctx
}

// IReturn_typeContext is an interface to support dynamic dispatch.
type IReturn_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturn_typeContext differentiates from other interfaces.
	IsReturn_typeContext()
}

type Return_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturn_typeContext() *Return_typeContext {
	var p = new(Return_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_return_type
	return p
}

func (*Return_typeContext) IsReturn_typeContext() {}

func NewReturn_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Return_typeContext {
	var p = new(Return_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_return_type

	return p
}

func (s *Return_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Return_typeContext) Assertion() IAssertionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssertionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssertionContext)
}

func (s *Return_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Return_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Return_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterReturn_type(s)
	}
}

func (s *Return_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitReturn_type(s)
	}
}

func (p *PowerQueryParser) Return_type() (localctx IReturn_typeContext) {
	this := p
	_ = this

	localctx = NewReturn_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, PowerQueryParserRULE_return_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(546)
		p.Assertion()
	}

	return localctx
}

// IAssertionContext is an interface to support dynamic dispatch.
type IAssertionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssertionContext differentiates from other interfaces.
	IsAssertionContext()
}

type AssertionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssertionContext() *AssertionContext {
	var p = new(AssertionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_assertion
	return p
}

func (*AssertionContext) IsAssertionContext() {}

func NewAssertionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssertionContext {
	var p = new(AssertionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_assertion

	return p
}

func (s *AssertionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssertionContext) AS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserAS, 0)
}

func (s *AssertionContext) Nullable_primitive_type() INullable_primitive_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullable_primitive_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullable_primitive_typeContext)
}

func (s *AssertionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssertionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterAssertion(s)
	}
}

func (s *AssertionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitAssertion(s)
	}
}

func (p *PowerQueryParser) Assertion() (localctx IAssertionContext) {
	this := p
	_ = this

	localctx = NewAssertionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, PowerQueryParserRULE_assertion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(PowerQueryParserAS)
	}
	{
		p.SetState(549)
		p.Nullable_primitive_type()
	}

	return localctx
}

// IOptional_parameter_listContext is an interface to support dynamic dispatch.
type IOptional_parameter_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptional_parameter_listContext differentiates from other interfaces.
	IsOptional_parameter_listContext()
}

type Optional_parameter_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptional_parameter_listContext() *Optional_parameter_listContext {
	var p = new(Optional_parameter_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_optional_parameter_list
	return p
}

func (*Optional_parameter_listContext) IsOptional_parameter_listContext() {}

func NewOptional_parameter_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Optional_parameter_listContext {
	var p = new(Optional_parameter_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_optional_parameter_list

	return p
}

func (s *Optional_parameter_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Optional_parameter_listContext) Optional_parameter() IOptional_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptional_parameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptional_parameterContext)
}

func (s *Optional_parameter_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Optional_parameter_listContext) Optional_parameter_list() IOptional_parameter_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptional_parameter_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptional_parameter_listContext)
}

func (s *Optional_parameter_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Optional_parameter_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Optional_parameter_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterOptional_parameter_list(s)
	}
}

func (s *Optional_parameter_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitOptional_parameter_list(s)
	}
}

func (p *PowerQueryParser) Optional_parameter_list() (localctx IOptional_parameter_listContext) {
	this := p
	_ = this

	localctx = NewOptional_parameter_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, PowerQueryParserRULE_optional_parameter_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(556)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(551)
			p.Optional_parameter()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(552)
			p.Optional_parameter()
		}
		{
			p.SetState(553)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(554)
			p.Optional_parameter_list()
		}

	}

	return localctx
}

// IOptional_parameterContext is an interface to support dynamic dispatch.
type IOptional_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptional_parameterContext differentiates from other interfaces.
	IsOptional_parameterContext()
}

type Optional_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptional_parameterContext() *Optional_parameterContext {
	var p = new(Optional_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_optional_parameter
	return p
}

func (*Optional_parameterContext) IsOptional_parameterContext() {}

func NewOptional_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Optional_parameterContext {
	var p = new(Optional_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_optional_parameter

	return p
}

func (s *Optional_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Optional_parameterContext) OPTIONAL_TEXT() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPTIONAL_TEXT, 0)
}

func (s *Optional_parameterContext) Parameter() IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *Optional_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Optional_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Optional_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterOptional_parameter(s)
	}
}

func (s *Optional_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitOptional_parameter(s)
	}
}

func (p *PowerQueryParser) Optional_parameter() (localctx IOptional_parameterContext) {
	this := p
	_ = this

	localctx = NewOptional_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, PowerQueryParserRULE_optional_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(558)
		p.Match(PowerQueryParserOPTIONAL_TEXT)
	}
	{
		p.SetState(559)
		p.Parameter()
	}

	return localctx
}

// IEach_expressionContext is an interface to support dynamic dispatch.
type IEach_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEach_expressionContext differentiates from other interfaces.
	IsEach_expressionContext()
}

type Each_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEach_expressionContext() *Each_expressionContext {
	var p = new(Each_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_each_expression
	return p
}

func (*Each_expressionContext) IsEach_expressionContext() {}

func NewEach_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Each_expressionContext {
	var p = new(Each_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_each_expression

	return p
}

func (s *Each_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Each_expressionContext) EACH() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserEACH, 0)
}

func (s *Each_expressionContext) Each_expression_body() IEach_expression_bodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEach_expression_bodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEach_expression_bodyContext)
}

func (s *Each_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Each_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Each_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterEach_expression(s)
	}
}

func (s *Each_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitEach_expression(s)
	}
}

func (p *PowerQueryParser) Each_expression() (localctx IEach_expressionContext) {
	this := p
	_ = this

	localctx = NewEach_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, PowerQueryParserRULE_each_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(561)
		p.Match(PowerQueryParserEACH)
	}
	{
		p.SetState(562)
		p.Each_expression_body()
	}

	return localctx
}

// IEach_expression_bodyContext is an interface to support dynamic dispatch.
type IEach_expression_bodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEach_expression_bodyContext differentiates from other interfaces.
	IsEach_expression_bodyContext()
}

type Each_expression_bodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEach_expression_bodyContext() *Each_expression_bodyContext {
	var p = new(Each_expression_bodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_each_expression_body
	return p
}

func (*Each_expression_bodyContext) IsEach_expression_bodyContext() {}

func NewEach_expression_bodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Each_expression_bodyContext {
	var p = new(Each_expression_bodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_each_expression_body

	return p
}

func (s *Each_expression_bodyContext) GetParser() antlr.Parser { return s.parser }

func (s *Each_expression_bodyContext) Function_body() IFunction_bodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_bodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_bodyContext)
}

func (s *Each_expression_bodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Each_expression_bodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Each_expression_bodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterEach_expression_body(s)
	}
}

func (s *Each_expression_bodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitEach_expression_body(s)
	}
}

func (p *PowerQueryParser) Each_expression_body() (localctx IEach_expression_bodyContext) {
	this := p
	_ = this

	localctx = NewEach_expression_bodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, PowerQueryParserRULE_each_expression_body)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(564)
		p.Function_body()
	}

	return localctx
}

// ILet_expressionContext is an interface to support dynamic dispatch.
type ILet_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLet_expressionContext differentiates from other interfaces.
	IsLet_expressionContext()
}

type Let_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLet_expressionContext() *Let_expressionContext {
	var p = new(Let_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_let_expression
	return p
}

func (*Let_expressionContext) IsLet_expressionContext() {}

func NewLet_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Let_expressionContext {
	var p = new(Let_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_let_expression

	return p
}

func (s *Let_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Let_expressionContext) LET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserLET, 0)
}

func (s *Let_expressionContext) Variable_list() IVariable_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariable_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariable_listContext)
}

func (s *Let_expressionContext) IN() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserIN, 0)
}

func (s *Let_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Let_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Let_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Let_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterLet_expression(s)
	}
}

func (s *Let_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitLet_expression(s)
	}
}

func (p *PowerQueryParser) Let_expression() (localctx ILet_expressionContext) {
	this := p
	_ = this

	localctx = NewLet_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, PowerQueryParserRULE_let_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(566)
		p.Match(PowerQueryParserLET)
	}
	{
		p.SetState(567)
		p.Variable_list()
	}
	{
		p.SetState(568)
		p.Match(PowerQueryParserIN)
	}
	{
		p.SetState(569)
		p.Expression()
	}

	return localctx
}

// IVariable_listContext is an interface to support dynamic dispatch.
type IVariable_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariable_listContext differentiates from other interfaces.
	IsVariable_listContext()
}

type Variable_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariable_listContext() *Variable_listContext {
	var p = new(Variable_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_variable_list
	return p
}

func (*Variable_listContext) IsVariable_listContext() {}

func NewVariable_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_listContext {
	var p = new(Variable_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_variable_list

	return p
}

func (s *Variable_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_listContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *Variable_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Variable_listContext) Variable_list() IVariable_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariable_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariable_listContext)
}

func (s *Variable_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Variable_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterVariable_list(s)
	}
}

func (s *Variable_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitVariable_list(s)
	}
}

func (p *PowerQueryParser) Variable_list() (localctx IVariable_listContext) {
	this := p
	_ = this

	localctx = NewVariable_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, PowerQueryParserRULE_variable_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(576)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(571)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(572)
			p.Variable()
		}
		{
			p.SetState(573)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(574)
			p.Variable_list()
		}

	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Variable_name() IVariable_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariable_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariable_nameContext)
}

func (s *VariableContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserEQUALS, 0)
}

func (s *VariableContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (p *PowerQueryParser) Variable() (localctx IVariableContext) {
	this := p
	_ = this

	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, PowerQueryParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(578)
		p.Variable_name()
	}
	{
		p.SetState(579)
		p.Match(PowerQueryParserEQUALS)
	}
	{
		p.SetState(580)
		p.Expression()
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
	p.RuleIndex = PowerQueryParserRULE_variable_name
	return p
}

func (*Variable_nameContext) IsVariable_nameContext() {}

func NewVariable_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Variable_nameContext {
	var p = new(Variable_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_variable_name

	return p
}

func (s *Variable_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Variable_nameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserIDENTIFIER, 0)
}

func (s *Variable_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Variable_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Variable_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterVariable_name(s)
	}
}

func (s *Variable_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitVariable_name(s)
	}
}

func (p *PowerQueryParser) Variable_name() (localctx IVariable_nameContext) {
	this := p
	_ = this

	localctx = NewVariable_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, PowerQueryParserRULE_variable_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(582)
		p.Match(PowerQueryParserIDENTIFIER)
	}

	return localctx
}

// IIf_expressionContext is an interface to support dynamic dispatch.
type IIf_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_expressionContext differentiates from other interfaces.
	IsIf_expressionContext()
}

type If_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_expressionContext() *If_expressionContext {
	var p = new(If_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_if_expression
	return p
}

func (*If_expressionContext) IsIf_expressionContext() {}

func NewIf_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_expressionContext {
	var p = new(If_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_if_expression

	return p
}

func (s *If_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *If_expressionContext) IF() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserIF, 0)
}

func (s *If_expressionContext) If_condition() IIf_conditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIf_conditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIf_conditionContext)
}

func (s *If_expressionContext) THEN() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserTHEN, 0)
}

func (s *If_expressionContext) True_expression() ITrue_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrue_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrue_expressionContext)
}

func (s *If_expressionContext) ELSE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserELSE, 0)
}

func (s *If_expressionContext) False_expression() IFalse_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFalse_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFalse_expressionContext)
}

func (s *If_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterIf_expression(s)
	}
}

func (s *If_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitIf_expression(s)
	}
}

func (p *PowerQueryParser) If_expression() (localctx IIf_expressionContext) {
	this := p
	_ = this

	localctx = NewIf_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, PowerQueryParserRULE_if_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(584)
		p.Match(PowerQueryParserIF)
	}
	{
		p.SetState(585)
		p.If_condition()
	}
	{
		p.SetState(586)
		p.Match(PowerQueryParserTHEN)
	}
	{
		p.SetState(587)
		p.True_expression()
	}
	{
		p.SetState(588)
		p.Match(PowerQueryParserELSE)
	}
	{
		p.SetState(589)
		p.False_expression()
	}

	return localctx
}

// IIf_conditionContext is an interface to support dynamic dispatch.
type IIf_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIf_conditionContext differentiates from other interfaces.
	IsIf_conditionContext()
}

type If_conditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_conditionContext() *If_conditionContext {
	var p = new(If_conditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_if_condition
	return p
}

func (*If_conditionContext) IsIf_conditionContext() {}

func NewIf_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_conditionContext {
	var p = new(If_conditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_if_condition

	return p
}

func (s *If_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *If_conditionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *If_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *If_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterIf_condition(s)
	}
}

func (s *If_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitIf_condition(s)
	}
}

func (p *PowerQueryParser) If_condition() (localctx IIf_conditionContext) {
	this := p
	_ = this

	localctx = NewIf_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, PowerQueryParserRULE_if_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Expression()
	}

	return localctx
}

// ITrue_expressionContext is an interface to support dynamic dispatch.
type ITrue_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTrue_expressionContext differentiates from other interfaces.
	IsTrue_expressionContext()
}

type True_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrue_expressionContext() *True_expressionContext {
	var p = new(True_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_true_expression
	return p
}

func (*True_expressionContext) IsTrue_expressionContext() {}

func NewTrue_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *True_expressionContext {
	var p = new(True_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_true_expression

	return p
}

func (s *True_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *True_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *True_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *True_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *True_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterTrue_expression(s)
	}
}

func (s *True_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitTrue_expression(s)
	}
}

func (p *PowerQueryParser) True_expression() (localctx ITrue_expressionContext) {
	this := p
	_ = this

	localctx = NewTrue_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, PowerQueryParserRULE_true_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(593)
		p.Expression()
	}

	return localctx
}

// IFalse_expressionContext is an interface to support dynamic dispatch.
type IFalse_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFalse_expressionContext differentiates from other interfaces.
	IsFalse_expressionContext()
}

type False_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFalse_expressionContext() *False_expressionContext {
	var p = new(False_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_false_expression
	return p
}

func (*False_expressionContext) IsFalse_expressionContext() {}

func NewFalse_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *False_expressionContext {
	var p = new(False_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_false_expression

	return p
}

func (s *False_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *False_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *False_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *False_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *False_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterFalse_expression(s)
	}
}

func (s *False_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitFalse_expression(s)
	}
}

func (p *PowerQueryParser) False_expression() (localctx IFalse_expressionContext) {
	this := p
	_ = this

	localctx = NewFalse_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, PowerQueryParserRULE_false_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Expression()
	}

	return localctx
}

// IType_expressionContext is an interface to support dynamic dispatch.
type IType_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_expressionContext differentiates from other interfaces.
	IsType_expressionContext()
}

type Type_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_expressionContext() *Type_expressionContext {
	var p = new(Type_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_type_expression
	return p
}

func (*Type_expressionContext) IsType_expressionContext() {}

func NewType_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_expressionContext {
	var p = new(Type_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_type_expression

	return p
}

func (s *Type_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_expressionContext) Primary_expression() IPrimary_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimary_expressionContext)
}

func (s *Type_expressionContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserTYPE, 0)
}

func (s *Type_expressionContext) Primary_type() IPrimary_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimary_typeContext)
}

func (s *Type_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterType_expression(s)
	}
}

func (s *Type_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitType_expression(s)
	}
}

func (p *PowerQueryParser) Type_expression() (localctx IType_expressionContext) {
	this := p
	_ = this

	localctx = NewType_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, PowerQueryParserRULE_type_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(600)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PowerQueryParserOPEN_BRACKET, PowerQueryParserOPEN_BRACE, PowerQueryParserOPEN_PAREN, PowerQueryParserELLIPSES, PowerQueryParserAT, PowerQueryParserLITERAL, PowerQueryParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(597)
			p.primary_expression(0)
		}

	case PowerQueryParserTYPE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(598)
			p.Match(PowerQueryParserTYPE)
		}
		{
			p.SetState(599)
			p.Primary_type()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IType_exprContext is an interface to support dynamic dispatch.
type IType_exprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_exprContext differentiates from other interfaces.
	IsType_exprContext()
}

type Type_exprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_exprContext() *Type_exprContext {
	var p = new(Type_exprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_type_expr
	return p
}

func (*Type_exprContext) IsType_exprContext() {}

func NewType_exprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_exprContext {
	var p = new(Type_exprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_type_expr

	return p
}

func (s *Type_exprContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_exprContext) Parenthesized_expression() IParenthesized_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParenthesized_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParenthesized_expressionContext)
}

func (s *Type_exprContext) Primary_type() IPrimary_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimary_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimary_typeContext)
}

func (s *Type_exprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_exprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_exprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterType_expr(s)
	}
}

func (s *Type_exprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitType_expr(s)
	}
}

func (p *PowerQueryParser) Type_expr() (localctx IType_exprContext) {
	this := p
	_ = this

	localctx = NewType_exprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, PowerQueryParserRULE_type_expr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(604)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PowerQueryParserOPEN_PAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(602)
			p.Parenthesized_expression()
		}

	case PowerQueryParserOPEN_BRACKET, PowerQueryParserOPEN_BRACE, PowerQueryParserTABLE, PowerQueryParserNULLABLE, PowerQueryParserFUNCTION_START, PowerQueryParserTYPE, PowerQueryParserTEXT, PowerQueryParserRECORD, PowerQueryParserNUMBER, PowerQueryParserNONE, PowerQueryParserLOGICAL, PowerQueryParserLIST, PowerQueryParserFUNCTION, PowerQueryParserDURATION, PowerQueryParserDATETIMEZONE, PowerQueryParserANY, PowerQueryParserANYNONNULL, PowerQueryParserBINARY, PowerQueryParserDATE, PowerQueryParserDATETIME, PowerQueryParserLITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(603)
			p.Primary_type()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrimary_typeContext is an interface to support dynamic dispatch.
type IPrimary_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimary_typeContext differentiates from other interfaces.
	IsPrimary_typeContext()
}

type Primary_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimary_typeContext() *Primary_typeContext {
	var p = new(Primary_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_primary_type
	return p
}

func (*Primary_typeContext) IsPrimary_typeContext() {}

func NewPrimary_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primary_typeContext {
	var p = new(Primary_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_primary_type

	return p
}

func (s *Primary_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Primary_typeContext) Primitive_type() IPrimitive_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitive_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitive_typeContext)
}

func (s *Primary_typeContext) Record_type() IRecord_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecord_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecord_typeContext)
}

func (s *Primary_typeContext) List_type() IList_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_typeContext)
}

func (s *Primary_typeContext) Function_type() IFunction_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_typeContext)
}

func (s *Primary_typeContext) Table_type() ITable_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITable_typeContext)
}

func (s *Primary_typeContext) Nullable_type() INullable_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullable_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullable_typeContext)
}

func (s *Primary_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primary_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primary_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterPrimary_type(s)
	}
}

func (s *Primary_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitPrimary_type(s)
	}
}

func (p *PowerQueryParser) Primary_type() (localctx IPrimary_typeContext) {
	this := p
	_ = this

	localctx = NewPrimary_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, PowerQueryParserRULE_primary_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(612)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(606)
			p.Primitive_type()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(607)
			p.Record_type()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(608)
			p.List_type()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(609)
			p.Function_type()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(610)
			p.Table_type()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(611)
			p.Nullable_type()
		}

	}

	return localctx
}

// IPrimitive_typeContext is an interface to support dynamic dispatch.
type IPrimitive_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimitive_typeContext differentiates from other interfaces.
	IsPrimitive_typeContext()
}

type Primitive_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimitive_typeContext() *Primitive_typeContext {
	var p = new(Primitive_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_primitive_type
	return p
}

func (*Primitive_typeContext) IsPrimitive_typeContext() {}

func NewPrimitive_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Primitive_typeContext {
	var p = new(Primitive_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_primitive_type

	return p
}

func (s *Primitive_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Primitive_typeContext) ANY() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserANY, 0)
}

func (s *Primitive_typeContext) ANYNONNULL() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserANYNONNULL, 0)
}

func (s *Primitive_typeContext) BINARY() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserBINARY, 0)
}

func (s *Primitive_typeContext) DATE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserDATE, 0)
}

func (s *Primitive_typeContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserDATETIME, 0)
}

func (s *Primitive_typeContext) DATETIMEZONE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserDATETIMEZONE, 0)
}

func (s *Primitive_typeContext) DURATION() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserDURATION, 0)
}

func (s *Primitive_typeContext) FUNCTION() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserFUNCTION, 0)
}

func (s *Primitive_typeContext) LIST() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserLIST, 0)
}

func (s *Primitive_typeContext) LOGICAL() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserLOGICAL, 0)
}

func (s *Primitive_typeContext) NONE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserNONE, 0)
}

func (s *Primitive_typeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserNUMBER, 0)
}

func (s *Primitive_typeContext) RECORD() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserRECORD, 0)
}

func (s *Primitive_typeContext) TABLE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserTABLE, 0)
}

func (s *Primitive_typeContext) TEXT() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserTEXT, 0)
}

func (s *Primitive_typeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserTYPE, 0)
}

func (s *Primitive_typeContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserLITERAL, 0)
}

func (s *Primitive_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Primitive_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Primitive_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterPrimitive_type(s)
	}
}

func (s *Primitive_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitPrimitive_type(s)
	}
}

func (p *PowerQueryParser) Primitive_type() (localctx IPrimitive_typeContext) {
	this := p
	_ = this

	localctx = NewPrimitive_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, PowerQueryParserRULE_primitive_type)
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
		p.SetState(614)
		_la = p.GetTokenStream().LA(1)

		if !((((_la-15)&-(0x1f+1)) == 0 && ((1<<uint((_la-15)))&((1<<(PowerQueryParserTABLE-15))|(1<<(PowerQueryParserTYPE-15))|(1<<(PowerQueryParserTEXT-15))|(1<<(PowerQueryParserRECORD-15))|(1<<(PowerQueryParserNUMBER-15))|(1<<(PowerQueryParserNONE-15))|(1<<(PowerQueryParserLOGICAL-15))|(1<<(PowerQueryParserLIST-15))|(1<<(PowerQueryParserFUNCTION-15))|(1<<(PowerQueryParserDURATION-15))|(1<<(PowerQueryParserDATETIMEZONE-15))|(1<<(PowerQueryParserANY-15))|(1<<(PowerQueryParserANYNONNULL-15))|(1<<(PowerQueryParserBINARY-15))|(1<<(PowerQueryParserDATE-15)))) != 0) || _la == PowerQueryParserDATETIME || _la == PowerQueryParserLITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRecord_typeContext is an interface to support dynamic dispatch.
type IRecord_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecord_typeContext differentiates from other interfaces.
	IsRecord_typeContext()
}

type Record_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecord_typeContext() *Record_typeContext {
	var p = new(Record_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_record_type
	return p
}

func (*Record_typeContext) IsRecord_typeContext() {}

func NewRecord_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Record_typeContext {
	var p = new(Record_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_record_type

	return p
}

func (s *Record_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Record_typeContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_BRACKET, 0)
}

func (s *Record_typeContext) Open_record_marker() IOpen_record_markerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpen_record_markerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpen_record_markerContext)
}

func (s *Record_typeContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_BRACKET, 0)
}

func (s *Record_typeContext) Field_specification_list() IField_specification_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_specification_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_specification_listContext)
}

func (s *Record_typeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Record_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Record_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Record_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterRecord_type(s)
	}
}

func (s *Record_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitRecord_type(s)
	}
}

func (p *PowerQueryParser) Record_type() (localctx IRecord_typeContext) {
	this := p
	_ = this

	localctx = NewRecord_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, PowerQueryParserRULE_record_type)
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

	p.SetState(631)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(616)
			p.Match(PowerQueryParserOPEN_BRACKET)
		}
		{
			p.SetState(617)
			p.Open_record_marker()
		}
		{
			p.SetState(618)
			p.Match(PowerQueryParserCLOSE_BRACKET)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(620)
			p.Match(PowerQueryParserOPEN_BRACKET)
		}
		p.SetState(622)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PowerQueryParserOPTIONAL_TEXT || _la == PowerQueryParserIDENTIFIER {
			{
				p.SetState(621)
				p.Field_specification_list()
			}

		}
		{
			p.SetState(624)
			p.Match(PowerQueryParserCLOSE_BRACKET)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(625)
			p.Match(PowerQueryParserOPEN_BRACKET)
		}
		{
			p.SetState(626)
			p.Field_specification_list()
		}
		{
			p.SetState(627)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(628)
			p.Open_record_marker()
		}
		{
			p.SetState(629)
			p.Match(PowerQueryParserCLOSE_BRACKET)
		}

	}

	return localctx
}

// IField_specification_listContext is an interface to support dynamic dispatch.
type IField_specification_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_specification_listContext differentiates from other interfaces.
	IsField_specification_listContext()
}

type Field_specification_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_specification_listContext() *Field_specification_listContext {
	var p = new(Field_specification_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_field_specification_list
	return p
}

func (*Field_specification_listContext) IsField_specification_listContext() {}

func NewField_specification_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_specification_listContext {
	var p = new(Field_specification_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_field_specification_list

	return p
}

func (s *Field_specification_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_specification_listContext) Field_specification() IField_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_specificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_specificationContext)
}

func (s *Field_specification_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Field_specification_listContext) Field_specification_list() IField_specification_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_specification_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_specification_listContext)
}

func (s *Field_specification_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_specification_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_specification_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterField_specification_list(s)
	}
}

func (s *Field_specification_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitField_specification_list(s)
	}
}

func (p *PowerQueryParser) Field_specification_list() (localctx IField_specification_listContext) {
	this := p
	_ = this

	localctx = NewField_specification_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, PowerQueryParserRULE_field_specification_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(638)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(633)
			p.Field_specification()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(634)
			p.Field_specification()
		}
		{
			p.SetState(635)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(636)
			p.Field_specification_list()
		}

	}

	return localctx
}

// IField_specificationContext is an interface to support dynamic dispatch.
type IField_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_specificationContext differentiates from other interfaces.
	IsField_specificationContext()
}

type Field_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_specificationContext() *Field_specificationContext {
	var p = new(Field_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_field_specification
	return p
}

func (*Field_specificationContext) IsField_specificationContext() {}

func NewField_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_specificationContext {
	var p = new(Field_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_field_specification

	return p
}

func (s *Field_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_specificationContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *Field_specificationContext) OPTIONAL_TEXT() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPTIONAL_TEXT, 0)
}

func (s *Field_specificationContext) Field_type_specification() IField_type_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_type_specificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_type_specificationContext)
}

func (s *Field_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterField_specification(s)
	}
}

func (s *Field_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitField_specification(s)
	}
}

func (p *PowerQueryParser) Field_specification() (localctx IField_specificationContext) {
	this := p
	_ = this

	localctx = NewField_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, PowerQueryParserRULE_field_specification)
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
	p.SetState(641)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserOPTIONAL_TEXT {
		{
			p.SetState(640)
			p.Match(PowerQueryParserOPTIONAL_TEXT)
		}

	}
	{
		p.SetState(643)
		p.Field_name()
	}
	p.SetState(645)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserEQUALS {
		{
			p.SetState(644)
			p.Field_type_specification()
		}

	}

	return localctx
}

// IField_type_specificationContext is an interface to support dynamic dispatch.
type IField_type_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_type_specificationContext differentiates from other interfaces.
	IsField_type_specificationContext()
}

type Field_type_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_type_specificationContext() *Field_type_specificationContext {
	var p = new(Field_type_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_field_type_specification
	return p
}

func (*Field_type_specificationContext) IsField_type_specificationContext() {}

func NewField_type_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_type_specificationContext {
	var p = new(Field_type_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_field_type_specification

	return p
}

func (s *Field_type_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_type_specificationContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserEQUALS, 0)
}

func (s *Field_type_specificationContext) Field_type() IField_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Field_type_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_type_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_type_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterField_type_specification(s)
	}
}

func (s *Field_type_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitField_type_specification(s)
	}
}

func (p *PowerQueryParser) Field_type_specification() (localctx IField_type_specificationContext) {
	this := p
	_ = this

	localctx = NewField_type_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, PowerQueryParserRULE_field_type_specification)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(647)
		p.Match(PowerQueryParserEQUALS)
	}
	{
		p.SetState(648)
		p.Field_type()
	}

	return localctx
}

// IField_typeContext is an interface to support dynamic dispatch.
type IField_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_typeContext differentiates from other interfaces.
	IsField_typeContext()
}

type Field_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_typeContext() *Field_typeContext {
	var p = new(Field_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_field_type
	return p
}

func (*Field_typeContext) IsField_typeContext() {}

func NewField_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_typeContext {
	var p = new(Field_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_field_type

	return p
}

func (s *Field_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_typeContext) Type_expr() IType_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Field_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterField_type(s)
	}
}

func (s *Field_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitField_type(s)
	}
}

func (p *PowerQueryParser) Field_type() (localctx IField_typeContext) {
	this := p
	_ = this

	localctx = NewField_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, PowerQueryParserRULE_field_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(650)
		p.Type_expr()
	}

	return localctx
}

// IOpen_record_markerContext is an interface to support dynamic dispatch.
type IOpen_record_markerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpen_record_markerContext differentiates from other interfaces.
	IsOpen_record_markerContext()
}

type Open_record_markerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpen_record_markerContext() *Open_record_markerContext {
	var p = new(Open_record_markerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_open_record_marker
	return p
}

func (*Open_record_markerContext) IsOpen_record_markerContext() {}

func NewOpen_record_markerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Open_record_markerContext {
	var p = new(Open_record_markerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_open_record_marker

	return p
}

func (s *Open_record_markerContext) GetParser() antlr.Parser { return s.parser }

func (s *Open_record_markerContext) ELLIPSES() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserELLIPSES, 0)
}

func (s *Open_record_markerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Open_record_markerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Open_record_markerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterOpen_record_marker(s)
	}
}

func (s *Open_record_markerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitOpen_record_marker(s)
	}
}

func (p *PowerQueryParser) Open_record_marker() (localctx IOpen_record_markerContext) {
	this := p
	_ = this

	localctx = NewOpen_record_markerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 152, PowerQueryParserRULE_open_record_marker)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(652)
		p.Match(PowerQueryParserELLIPSES)
	}

	return localctx
}

// IList_typeContext is an interface to support dynamic dispatch.
type IList_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_typeContext differentiates from other interfaces.
	IsList_typeContext()
}

type List_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_typeContext() *List_typeContext {
	var p = new(List_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_list_type
	return p
}

func (*List_typeContext) IsList_typeContext() {}

func NewList_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_typeContext {
	var p = new(List_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_list_type

	return p
}

func (s *List_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *List_typeContext) OPEN_BRACE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_BRACE, 0)
}

func (s *List_typeContext) Item_type() IItem_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IItem_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IItem_typeContext)
}

func (s *List_typeContext) CLOSE_BRACE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_BRACE, 0)
}

func (s *List_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterList_type(s)
	}
}

func (s *List_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitList_type(s)
	}
}

func (p *PowerQueryParser) List_type() (localctx IList_typeContext) {
	this := p
	_ = this

	localctx = NewList_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 154, PowerQueryParserRULE_list_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(654)
		p.Match(PowerQueryParserOPEN_BRACE)
	}
	{
		p.SetState(655)
		p.Item_type()
	}
	{
		p.SetState(656)
		p.Match(PowerQueryParserCLOSE_BRACE)
	}

	return localctx
}

// IItem_typeContext is an interface to support dynamic dispatch.
type IItem_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsItem_typeContext differentiates from other interfaces.
	IsItem_typeContext()
}

type Item_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyItem_typeContext() *Item_typeContext {
	var p = new(Item_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_item_type
	return p
}

func (*Item_typeContext) IsItem_typeContext() {}

func NewItem_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Item_typeContext {
	var p = new(Item_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_item_type

	return p
}

func (s *Item_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Item_typeContext) Type_expr() IType_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Item_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Item_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Item_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterItem_type(s)
	}
}

func (s *Item_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitItem_type(s)
	}
}

func (p *PowerQueryParser) Item_type() (localctx IItem_typeContext) {
	this := p
	_ = this

	localctx = NewItem_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 156, PowerQueryParserRULE_item_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(658)
		p.Type_expr()
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
	p.RuleIndex = PowerQueryParserRULE_function_type
	return p
}

func (*Function_typeContext) IsFunction_typeContext() {}

func NewFunction_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_typeContext {
	var p = new(Function_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_function_type

	return p
}

func (s *Function_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_typeContext) FUNCTION_START() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserFUNCTION_START, 0)
}

func (s *Function_typeContext) CLOSE_PAREN() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_PAREN, 0)
}

func (s *Function_typeContext) Return_type() IReturn_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturn_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturn_typeContext)
}

func (s *Function_typeContext) Parameter_specification_list() IParameter_specification_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameter_specification_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameter_specification_listContext)
}

func (s *Function_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterFunction_type(s)
	}
}

func (s *Function_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitFunction_type(s)
	}
}

func (p *PowerQueryParser) Function_type() (localctx IFunction_typeContext) {
	this := p
	_ = this

	localctx = NewFunction_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 158, PowerQueryParserRULE_function_type)
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
		p.SetState(660)
		p.Match(PowerQueryParserFUNCTION_START)
	}
	p.SetState(662)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserOPTIONAL_TEXT || _la == PowerQueryParserIDENTIFIER {
		{
			p.SetState(661)
			p.Parameter_specification_list()
		}

	}
	{
		p.SetState(664)
		p.Match(PowerQueryParserCLOSE_PAREN)
	}
	{
		p.SetState(665)
		p.Return_type()
	}

	return localctx
}

// IParameter_specification_listContext is an interface to support dynamic dispatch.
type IParameter_specification_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameter_specification_listContext differentiates from other interfaces.
	IsParameter_specification_listContext()
}

type Parameter_specification_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameter_specification_listContext() *Parameter_specification_listContext {
	var p = new(Parameter_specification_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_parameter_specification_list
	return p
}

func (*Parameter_specification_listContext) IsParameter_specification_listContext() {}

func NewParameter_specification_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parameter_specification_listContext {
	var p = new(Parameter_specification_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_parameter_specification_list

	return p
}

func (s *Parameter_specification_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Parameter_specification_listContext) Required_parameter_specification_list() IRequired_parameter_specification_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequired_parameter_specification_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequired_parameter_specification_listContext)
}

func (s *Parameter_specification_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Parameter_specification_listContext) Optional_parameter_specification_list() IOptional_parameter_specification_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptional_parameter_specification_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptional_parameter_specification_listContext)
}

func (s *Parameter_specification_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parameter_specification_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parameter_specification_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterParameter_specification_list(s)
	}
}

func (s *Parameter_specification_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitParameter_specification_list(s)
	}
}

func (p *PowerQueryParser) Parameter_specification_list() (localctx IParameter_specification_listContext) {
	this := p
	_ = this

	localctx = NewParameter_specification_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 160, PowerQueryParserRULE_parameter_specification_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(667)
			p.Required_parameter_specification_list()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(668)
			p.Required_parameter_specification_list()
		}
		{
			p.SetState(669)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(670)
			p.Optional_parameter_specification_list()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(672)
			p.Optional_parameter_specification_list()
		}

	}

	return localctx
}

// IRequired_parameter_specification_listContext is an interface to support dynamic dispatch.
type IRequired_parameter_specification_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequired_parameter_specification_listContext differentiates from other interfaces.
	IsRequired_parameter_specification_listContext()
}

type Required_parameter_specification_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequired_parameter_specification_listContext() *Required_parameter_specification_listContext {
	var p = new(Required_parameter_specification_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_required_parameter_specification_list
	return p
}

func (*Required_parameter_specification_listContext) IsRequired_parameter_specification_listContext() {
}

func NewRequired_parameter_specification_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Required_parameter_specification_listContext {
	var p = new(Required_parameter_specification_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_required_parameter_specification_list

	return p
}

func (s *Required_parameter_specification_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Required_parameter_specification_listContext) Required_parameter_specification() IRequired_parameter_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequired_parameter_specificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequired_parameter_specificationContext)
}

func (s *Required_parameter_specification_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Required_parameter_specification_listContext) Required_parameter_specification_list() IRequired_parameter_specification_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRequired_parameter_specification_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRequired_parameter_specification_listContext)
}

func (s *Required_parameter_specification_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Required_parameter_specification_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Required_parameter_specification_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterRequired_parameter_specification_list(s)
	}
}

func (s *Required_parameter_specification_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitRequired_parameter_specification_list(s)
	}
}

func (p *PowerQueryParser) Required_parameter_specification_list() (localctx IRequired_parameter_specification_listContext) {
	this := p
	_ = this

	localctx = NewRequired_parameter_specification_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 162, PowerQueryParserRULE_required_parameter_specification_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(680)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(675)
			p.Required_parameter_specification()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(676)
			p.Required_parameter_specification()
		}
		{
			p.SetState(677)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(678)
			p.Required_parameter_specification_list()
		}

	}

	return localctx
}

// IRequired_parameter_specificationContext is an interface to support dynamic dispatch.
type IRequired_parameter_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRequired_parameter_specificationContext differentiates from other interfaces.
	IsRequired_parameter_specificationContext()
}

type Required_parameter_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequired_parameter_specificationContext() *Required_parameter_specificationContext {
	var p = new(Required_parameter_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_required_parameter_specification
	return p
}

func (*Required_parameter_specificationContext) IsRequired_parameter_specificationContext() {}

func NewRequired_parameter_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Required_parameter_specificationContext {
	var p = new(Required_parameter_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_required_parameter_specification

	return p
}

func (s *Required_parameter_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Required_parameter_specificationContext) Parameter_specification() IParameter_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameter_specificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameter_specificationContext)
}

func (s *Required_parameter_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Required_parameter_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Required_parameter_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterRequired_parameter_specification(s)
	}
}

func (s *Required_parameter_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitRequired_parameter_specification(s)
	}
}

func (p *PowerQueryParser) Required_parameter_specification() (localctx IRequired_parameter_specificationContext) {
	this := p
	_ = this

	localctx = NewRequired_parameter_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 164, PowerQueryParserRULE_required_parameter_specification)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Parameter_specification()
	}

	return localctx
}

// IOptional_parameter_specification_listContext is an interface to support dynamic dispatch.
type IOptional_parameter_specification_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptional_parameter_specification_listContext differentiates from other interfaces.
	IsOptional_parameter_specification_listContext()
}

type Optional_parameter_specification_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptional_parameter_specification_listContext() *Optional_parameter_specification_listContext {
	var p = new(Optional_parameter_specification_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_optional_parameter_specification_list
	return p
}

func (*Optional_parameter_specification_listContext) IsOptional_parameter_specification_listContext() {
}

func NewOptional_parameter_specification_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Optional_parameter_specification_listContext {
	var p = new(Optional_parameter_specification_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_optional_parameter_specification_list

	return p
}

func (s *Optional_parameter_specification_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Optional_parameter_specification_listContext) Optional_parameter_specification() IOptional_parameter_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptional_parameter_specificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptional_parameter_specificationContext)
}

func (s *Optional_parameter_specification_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Optional_parameter_specification_listContext) Optional_parameter_specification_list() IOptional_parameter_specification_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptional_parameter_specification_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptional_parameter_specification_listContext)
}

func (s *Optional_parameter_specification_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Optional_parameter_specification_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Optional_parameter_specification_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterOptional_parameter_specification_list(s)
	}
}

func (s *Optional_parameter_specification_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitOptional_parameter_specification_list(s)
	}
}

func (p *PowerQueryParser) Optional_parameter_specification_list() (localctx IOptional_parameter_specification_listContext) {
	this := p
	_ = this

	localctx = NewOptional_parameter_specification_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 166, PowerQueryParserRULE_optional_parameter_specification_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(689)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(684)
			p.Optional_parameter_specification()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(685)
			p.Optional_parameter_specification()
		}
		{
			p.SetState(686)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(687)
			p.Optional_parameter_specification_list()
		}

	}

	return localctx
}

// IOptional_parameter_specificationContext is an interface to support dynamic dispatch.
type IOptional_parameter_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptional_parameter_specificationContext differentiates from other interfaces.
	IsOptional_parameter_specificationContext()
}

type Optional_parameter_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptional_parameter_specificationContext() *Optional_parameter_specificationContext {
	var p = new(Optional_parameter_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_optional_parameter_specification
	return p
}

func (*Optional_parameter_specificationContext) IsOptional_parameter_specificationContext() {}

func NewOptional_parameter_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Optional_parameter_specificationContext {
	var p = new(Optional_parameter_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_optional_parameter_specification

	return p
}

func (s *Optional_parameter_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Optional_parameter_specificationContext) OPTIONAL_TEXT() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPTIONAL_TEXT, 0)
}

func (s *Optional_parameter_specificationContext) Parameter_specification() IParameter_specificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameter_specificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameter_specificationContext)
}

func (s *Optional_parameter_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Optional_parameter_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Optional_parameter_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterOptional_parameter_specification(s)
	}
}

func (s *Optional_parameter_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitOptional_parameter_specification(s)
	}
}

func (p *PowerQueryParser) Optional_parameter_specification() (localctx IOptional_parameter_specificationContext) {
	this := p
	_ = this

	localctx = NewOptional_parameter_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 168, PowerQueryParserRULE_optional_parameter_specification)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(691)
		p.Match(PowerQueryParserOPTIONAL_TEXT)
	}
	{
		p.SetState(692)
		p.Parameter_specification()
	}

	return localctx
}

// IParameter_specificationContext is an interface to support dynamic dispatch.
type IParameter_specificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameter_specificationContext differentiates from other interfaces.
	IsParameter_specificationContext()
}

type Parameter_specificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameter_specificationContext() *Parameter_specificationContext {
	var p = new(Parameter_specificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_parameter_specification
	return p
}

func (*Parameter_specificationContext) IsParameter_specificationContext() {}

func NewParameter_specificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Parameter_specificationContext {
	var p = new(Parameter_specificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_parameter_specification

	return p
}

func (s *Parameter_specificationContext) GetParser() antlr.Parser { return s.parser }

func (s *Parameter_specificationContext) Parameter_name() IParameter_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameter_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameter_nameContext)
}

func (s *Parameter_specificationContext) Parameter_type() IParameter_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameter_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParameter_typeContext)
}

func (s *Parameter_specificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parameter_specificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Parameter_specificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterParameter_specification(s)
	}
}

func (s *Parameter_specificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitParameter_specification(s)
	}
}

func (p *PowerQueryParser) Parameter_specification() (localctx IParameter_specificationContext) {
	this := p
	_ = this

	localctx = NewParameter_specificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 170, PowerQueryParserRULE_parameter_specification)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(694)
		p.Parameter_name()
	}
	{
		p.SetState(695)
		p.Parameter_type()
	}

	return localctx
}

// ITable_typeContext is an interface to support dynamic dispatch.
type ITable_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_typeContext differentiates from other interfaces.
	IsTable_typeContext()
}

type Table_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_typeContext() *Table_typeContext {
	var p = new(Table_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_table_type
	return p
}

func (*Table_typeContext) IsTable_typeContext() {}

func NewTable_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_typeContext {
	var p = new(Table_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_table_type

	return p
}

func (s *Table_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_typeContext) TABLE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserTABLE, 0)
}

func (s *Table_typeContext) Row_type() IRow_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRow_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRow_typeContext)
}

func (s *Table_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterTable_type(s)
	}
}

func (s *Table_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitTable_type(s)
	}
}

func (p *PowerQueryParser) Table_type() (localctx ITable_typeContext) {
	this := p
	_ = this

	localctx = NewTable_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 172, PowerQueryParserRULE_table_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(697)
		p.Match(PowerQueryParserTABLE)
	}
	{
		p.SetState(698)
		p.Row_type()
	}

	return localctx
}

// IRow_typeContext is an interface to support dynamic dispatch.
type IRow_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRow_typeContext differentiates from other interfaces.
	IsRow_typeContext()
}

type Row_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRow_typeContext() *Row_typeContext {
	var p = new(Row_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_row_type
	return p
}

func (*Row_typeContext) IsRow_typeContext() {}

func NewRow_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Row_typeContext {
	var p = new(Row_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_row_type

	return p
}

func (s *Row_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Row_typeContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_BRACKET, 0)
}

func (s *Row_typeContext) Field_specification_list() IField_specification_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_specification_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_specification_listContext)
}

func (s *Row_typeContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_BRACKET, 0)
}

func (s *Row_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Row_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Row_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterRow_type(s)
	}
}

func (s *Row_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitRow_type(s)
	}
}

func (p *PowerQueryParser) Row_type() (localctx IRow_typeContext) {
	this := p
	_ = this

	localctx = NewRow_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 174, PowerQueryParserRULE_row_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(700)
		p.Match(PowerQueryParserOPEN_BRACKET)
	}
	{
		p.SetState(701)
		p.Field_specification_list()
	}
	{
		p.SetState(702)
		p.Match(PowerQueryParserCLOSE_BRACKET)
	}

	return localctx
}

// INullable_typeContext is an interface to support dynamic dispatch.
type INullable_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullable_typeContext differentiates from other interfaces.
	IsNullable_typeContext()
}

type Nullable_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullable_typeContext() *Nullable_typeContext {
	var p = new(Nullable_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_nullable_type
	return p
}

func (*Nullable_typeContext) IsNullable_typeContext() {}

func NewNullable_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Nullable_typeContext {
	var p = new(Nullable_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_nullable_type

	return p
}

func (s *Nullable_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Nullable_typeContext) NULLABLE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserNULLABLE, 0)
}

func (s *Nullable_typeContext) Type_expr() IType_exprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_exprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_exprContext)
}

func (s *Nullable_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Nullable_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Nullable_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterNullable_type(s)
	}
}

func (s *Nullable_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitNullable_type(s)
	}
}

func (p *PowerQueryParser) Nullable_type() (localctx INullable_typeContext) {
	this := p
	_ = this

	localctx = NewNullable_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 176, PowerQueryParserRULE_nullable_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(704)
		p.Match(PowerQueryParserNULLABLE)
	}
	{
		p.SetState(705)
		p.Type_expr()
	}

	return localctx
}

// IError_raising_expressionContext is an interface to support dynamic dispatch.
type IError_raising_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsError_raising_expressionContext differentiates from other interfaces.
	IsError_raising_expressionContext()
}

type Error_raising_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyError_raising_expressionContext() *Error_raising_expressionContext {
	var p = new(Error_raising_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_error_raising_expression
	return p
}

func (*Error_raising_expressionContext) IsError_raising_expressionContext() {}

func NewError_raising_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Error_raising_expressionContext {
	var p = new(Error_raising_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_error_raising_expression

	return p
}

func (s *Error_raising_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Error_raising_expressionContext) ERROR() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserERROR, 0)
}

func (s *Error_raising_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Error_raising_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Error_raising_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Error_raising_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterError_raising_expression(s)
	}
}

func (s *Error_raising_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitError_raising_expression(s)
	}
}

func (p *PowerQueryParser) Error_raising_expression() (localctx IError_raising_expressionContext) {
	this := p
	_ = this

	localctx = NewError_raising_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 178, PowerQueryParserRULE_error_raising_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(707)
		p.Match(PowerQueryParserERROR)
	}
	{
		p.SetState(708)
		p.Expression()
	}

	return localctx
}

// IError_handling_expressionContext is an interface to support dynamic dispatch.
type IError_handling_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsError_handling_expressionContext differentiates from other interfaces.
	IsError_handling_expressionContext()
}

type Error_handling_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyError_handling_expressionContext() *Error_handling_expressionContext {
	var p = new(Error_handling_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_error_handling_expression
	return p
}

func (*Error_handling_expressionContext) IsError_handling_expressionContext() {}

func NewError_handling_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Error_handling_expressionContext {
	var p = new(Error_handling_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_error_handling_expression

	return p
}

func (s *Error_handling_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Error_handling_expressionContext) TRY() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserTRY, 0)
}

func (s *Error_handling_expressionContext) Protected_expression() IProtected_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProtected_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IProtected_expressionContext)
}

func (s *Error_handling_expressionContext) Otherwise_clause() IOtherwise_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOtherwise_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOtherwise_clauseContext)
}

func (s *Error_handling_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Error_handling_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Error_handling_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterError_handling_expression(s)
	}
}

func (s *Error_handling_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitError_handling_expression(s)
	}
}

func (p *PowerQueryParser) Error_handling_expression() (localctx IError_handling_expressionContext) {
	this := p
	_ = this

	localctx = NewError_handling_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 180, PowerQueryParserRULE_error_handling_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(710)
		p.Match(PowerQueryParserTRY)
	}
	{
		p.SetState(711)
		p.Protected_expression()
	}
	p.SetState(713)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(712)
			p.Otherwise_clause()
		}

	}

	return localctx
}

// IProtected_expressionContext is an interface to support dynamic dispatch.
type IProtected_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProtected_expressionContext differentiates from other interfaces.
	IsProtected_expressionContext()
}

type Protected_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtected_expressionContext() *Protected_expressionContext {
	var p = new(Protected_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_protected_expression
	return p
}

func (*Protected_expressionContext) IsProtected_expressionContext() {}

func NewProtected_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Protected_expressionContext {
	var p = new(Protected_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_protected_expression

	return p
}

func (s *Protected_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Protected_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Protected_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Protected_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Protected_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterProtected_expression(s)
	}
}

func (s *Protected_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitProtected_expression(s)
	}
}

func (p *PowerQueryParser) Protected_expression() (localctx IProtected_expressionContext) {
	this := p
	_ = this

	localctx = NewProtected_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 182, PowerQueryParserRULE_protected_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(715)
		p.Expression()
	}

	return localctx
}

// IOtherwise_clauseContext is an interface to support dynamic dispatch.
type IOtherwise_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOtherwise_clauseContext differentiates from other interfaces.
	IsOtherwise_clauseContext()
}

type Otherwise_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOtherwise_clauseContext() *Otherwise_clauseContext {
	var p = new(Otherwise_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_otherwise_clause
	return p
}

func (*Otherwise_clauseContext) IsOtherwise_clauseContext() {}

func NewOtherwise_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Otherwise_clauseContext {
	var p = new(Otherwise_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_otherwise_clause

	return p
}

func (s *Otherwise_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Otherwise_clauseContext) OTHERWISE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOTHERWISE, 0)
}

func (s *Otherwise_clauseContext) Default_expression() IDefault_expressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefault_expressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefault_expressionContext)
}

func (s *Otherwise_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Otherwise_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Otherwise_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterOtherwise_clause(s)
	}
}

func (s *Otherwise_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitOtherwise_clause(s)
	}
}

func (p *PowerQueryParser) Otherwise_clause() (localctx IOtherwise_clauseContext) {
	this := p
	_ = this

	localctx = NewOtherwise_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 184, PowerQueryParserRULE_otherwise_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(717)
		p.Match(PowerQueryParserOTHERWISE)
	}
	{
		p.SetState(718)
		p.Default_expression()
	}

	return localctx
}

// IDefault_expressionContext is an interface to support dynamic dispatch.
type IDefault_expressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefault_expressionContext differentiates from other interfaces.
	IsDefault_expressionContext()
}

type Default_expressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefault_expressionContext() *Default_expressionContext {
	var p = new(Default_expressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_default_expression
	return p
}

func (*Default_expressionContext) IsDefault_expressionContext() {}

func NewDefault_expressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Default_expressionContext {
	var p = new(Default_expressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_default_expression

	return p
}

func (s *Default_expressionContext) GetParser() antlr.Parser { return s.parser }

func (s *Default_expressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Default_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Default_expressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Default_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterDefault_expression(s)
	}
}

func (s *Default_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitDefault_expression(s)
	}
}

func (p *PowerQueryParser) Default_expression() (localctx IDefault_expressionContext) {
	this := p
	_ = this

	localctx = NewDefault_expressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 186, PowerQueryParserRULE_default_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Expression()
	}

	return localctx
}

// ILiteral_attribsContext is an interface to support dynamic dispatch.
type ILiteral_attribsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteral_attribsContext differentiates from other interfaces.
	IsLiteral_attribsContext()
}

type Literal_attribsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_attribsContext() *Literal_attribsContext {
	var p = new(Literal_attribsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_literal_attribs
	return p
}

func (*Literal_attribsContext) IsLiteral_attribsContext() {}

func NewLiteral_attribsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_attribsContext {
	var p = new(Literal_attribsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_literal_attribs

	return p
}

func (s *Literal_attribsContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_attribsContext) Record_literal() IRecord_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecord_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecord_literalContext)
}

func (s *Literal_attribsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_attribsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_attribsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterLiteral_attribs(s)
	}
}

func (s *Literal_attribsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitLiteral_attribs(s)
	}
}

func (p *PowerQueryParser) Literal_attribs() (localctx ILiteral_attribsContext) {
	this := p
	_ = this

	localctx = NewLiteral_attribsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 188, PowerQueryParserRULE_literal_attribs)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(722)
		p.Record_literal()
	}

	return localctx
}

// IRecord_literalContext is an interface to support dynamic dispatch.
type IRecord_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRecord_literalContext differentiates from other interfaces.
	IsRecord_literalContext()
}

type Record_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRecord_literalContext() *Record_literalContext {
	var p = new(Record_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_record_literal
	return p
}

func (*Record_literalContext) IsRecord_literalContext() {}

func NewRecord_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Record_literalContext {
	var p = new(Record_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_record_literal

	return p
}

func (s *Record_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Record_literalContext) OPEN_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_BRACKET, 0)
}

func (s *Record_literalContext) CLOSE_BRACKET() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_BRACKET, 0)
}

func (s *Record_literalContext) Literal_field_list() ILiteral_field_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteral_field_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteral_field_listContext)
}

func (s *Record_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Record_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Record_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterRecord_literal(s)
	}
}

func (s *Record_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitRecord_literal(s)
	}
}

func (p *PowerQueryParser) Record_literal() (localctx IRecord_literalContext) {
	this := p
	_ = this

	localctx = NewRecord_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 190, PowerQueryParserRULE_record_literal)
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
		p.SetState(724)
		p.Match(PowerQueryParserOPEN_BRACKET)
	}
	p.SetState(726)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserIDENTIFIER {
		{
			p.SetState(725)
			p.Literal_field_list()
		}

	}
	{
		p.SetState(728)
		p.Match(PowerQueryParserCLOSE_BRACKET)
	}

	return localctx
}

// ILiteral_field_listContext is an interface to support dynamic dispatch.
type ILiteral_field_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteral_field_listContext differentiates from other interfaces.
	IsLiteral_field_listContext()
}

type Literal_field_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_field_listContext() *Literal_field_listContext {
	var p = new(Literal_field_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_literal_field_list
	return p
}

func (*Literal_field_listContext) IsLiteral_field_listContext() {}

func NewLiteral_field_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_field_listContext {
	var p = new(Literal_field_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_literal_field_list

	return p
}

func (s *Literal_field_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_field_listContext) Literal_field() ILiteral_fieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteral_fieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteral_fieldContext)
}

func (s *Literal_field_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Literal_field_listContext) Literal_field_list() ILiteral_field_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteral_field_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteral_field_listContext)
}

func (s *Literal_field_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_field_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_field_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterLiteral_field_list(s)
	}
}

func (s *Literal_field_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitLiteral_field_list(s)
	}
}

func (p *PowerQueryParser) Literal_field_list() (localctx ILiteral_field_listContext) {
	this := p
	_ = this

	localctx = NewLiteral_field_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 192, PowerQueryParserRULE_literal_field_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(735)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(730)
			p.Literal_field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(731)
			p.Literal_field()
		}
		{
			p.SetState(732)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(733)
			p.Literal_field_list()
		}

	}

	return localctx
}

// ILiteral_fieldContext is an interface to support dynamic dispatch.
type ILiteral_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteral_fieldContext differentiates from other interfaces.
	IsLiteral_fieldContext()
}

type Literal_fieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_fieldContext() *Literal_fieldContext {
	var p = new(Literal_fieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_literal_field
	return p
}

func (*Literal_fieldContext) IsLiteral_fieldContext() {}

func NewLiteral_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_fieldContext {
	var p = new(Literal_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_literal_field

	return p
}

func (s *Literal_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_fieldContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *Literal_fieldContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserEQUALS, 0)
}

func (s *Literal_fieldContext) Any_literal() IAny_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAny_literalContext)
}

func (s *Literal_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterLiteral_field(s)
	}
}

func (s *Literal_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitLiteral_field(s)
	}
}

func (p *PowerQueryParser) Literal_field() (localctx ILiteral_fieldContext) {
	this := p
	_ = this

	localctx = NewLiteral_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 194, PowerQueryParserRULE_literal_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(737)
		p.Field_name()
	}
	{
		p.SetState(738)
		p.Match(PowerQueryParserEQUALS)
	}
	{
		p.SetState(739)
		p.Any_literal()
	}

	return localctx
}

// IList_literalContext is an interface to support dynamic dispatch.
type IList_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_literalContext differentiates from other interfaces.
	IsList_literalContext()
}

type List_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_literalContext() *List_literalContext {
	var p = new(List_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_list_literal
	return p
}

func (*List_literalContext) IsList_literalContext() {}

func NewList_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_literalContext {
	var p = new(List_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_list_literal

	return p
}

func (s *List_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *List_literalContext) OPEN_BRACE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserOPEN_BRACE, 0)
}

func (s *List_literalContext) CLOSE_BRACE() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCLOSE_BRACE, 0)
}

func (s *List_literalContext) Literal_item_list() ILiteral_item_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteral_item_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteral_item_listContext)
}

func (s *List_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterList_literal(s)
	}
}

func (s *List_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitList_literal(s)
	}
}

func (p *PowerQueryParser) List_literal() (localctx IList_literalContext) {
	this := p
	_ = this

	localctx = NewList_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 196, PowerQueryParserRULE_list_literal)
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
		p.SetState(741)
		p.Match(PowerQueryParserOPEN_BRACE)
	}
	p.SetState(743)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerQueryParserOPEN_BRACKET || _la == PowerQueryParserOPEN_BRACE || _la == PowerQueryParserLITERAL {
		{
			p.SetState(742)
			p.Literal_item_list()
		}

	}
	{
		p.SetState(745)
		p.Match(PowerQueryParserCLOSE_BRACE)
	}

	return localctx
}

// ILiteral_item_listContext is an interface to support dynamic dispatch.
type ILiteral_item_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteral_item_listContext differentiates from other interfaces.
	IsLiteral_item_listContext()
}

type Literal_item_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteral_item_listContext() *Literal_item_listContext {
	var p = new(Literal_item_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_literal_item_list
	return p
}

func (*Literal_item_listContext) IsLiteral_item_listContext() {}

func NewLiteral_item_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Literal_item_listContext {
	var p = new(Literal_item_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_literal_item_list

	return p
}

func (s *Literal_item_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Literal_item_listContext) Any_literal() IAny_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAny_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAny_literalContext)
}

func (s *Literal_item_listContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserCOMMA, 0)
}

func (s *Literal_item_listContext) Literal_item_list() ILiteral_item_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteral_item_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteral_item_listContext)
}

func (s *Literal_item_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Literal_item_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Literal_item_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterLiteral_item_list(s)
	}
}

func (s *Literal_item_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitLiteral_item_list(s)
	}
}

func (p *PowerQueryParser) Literal_item_list() (localctx ILiteral_item_listContext) {
	this := p
	_ = this

	localctx = NewLiteral_item_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 198, PowerQueryParserRULE_literal_item_list)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(752)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(747)
			p.Any_literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(748)
			p.Any_literal()
		}
		{
			p.SetState(749)
			p.Match(PowerQueryParserCOMMA)
		}
		{
			p.SetState(750)
			p.Literal_item_list()
		}

	}

	return localctx
}

// IAny_literalContext is an interface to support dynamic dispatch.
type IAny_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAny_literalContext differentiates from other interfaces.
	IsAny_literalContext()
}

type Any_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAny_literalContext() *Any_literalContext {
	var p = new(Any_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerQueryParserRULE_any_literal
	return p
}

func (*Any_literalContext) IsAny_literalContext() {}

func NewAny_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Any_literalContext {
	var p = new(Any_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerQueryParserRULE_any_literal

	return p
}

func (s *Any_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Any_literalContext) Record_literal() IRecord_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRecord_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRecord_literalContext)
}

func (s *Any_literalContext) List_literal() IList_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_literalContext)
}

func (s *Any_literalContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(PowerQueryParserLITERAL, 0)
}

func (s *Any_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Any_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Any_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.EnterAny_literal(s)
	}
}

func (s *Any_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerQueryParserListener); ok {
		listenerT.ExitAny_literal(s)
	}
}

func (p *PowerQueryParser) Any_literal() (localctx IAny_literalContext) {
	this := p
	_ = this

	localctx = NewAny_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 200, PowerQueryParserRULE_any_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(757)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PowerQueryParserOPEN_BRACKET:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(754)
			p.Record_literal()
		}

	case PowerQueryParserOPEN_BRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(755)
			p.List_literal()
		}

	case PowerQueryParserLITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(756)
			p.Match(PowerQueryParserLITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *PowerQueryParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *Logical_and_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Logical_and_expressionContext)
		}
		return p.Logical_and_expression_Sempred(t, predIndex)

	case 11:
		var t *Is_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Is_expressionContext)
		}
		return p.Is_expression_Sempred(t, predIndex)

	case 13:
		var t *As_expressionContext = nil
		if localctx != nil {
			t = localctx.(*As_expressionContext)
		}
		return p.As_expression_Sempred(t, predIndex)

	case 20:
		var t *Primary_expressionContext = nil
		if localctx != nil {
			t = localctx.(*Primary_expressionContext)
		}
		return p.Primary_expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PowerQueryParser) Logical_and_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PowerQueryParser) Is_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PowerQueryParser) As_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *PowerQueryParser) Primary_expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
