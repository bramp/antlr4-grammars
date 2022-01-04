// Code generated from smiles.g4 by ANTLR 4.9.3. DO NOT EDIT.

package smiles // smiles
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 147, 468,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 5, 2, 43, 10, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 5, 3, 49, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 63, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6,
	5, 6, 69, 10, 6, 3, 6, 3, 6, 5, 6, 73, 10, 6, 3, 6, 5, 6, 76, 10, 6, 3,
	6, 5, 6, 79, 10, 6, 3, 6, 5, 6, 82, 10, 6, 3, 6, 3, 6, 3, 7, 6, 7, 87,
	10, 7, 13, 7, 14, 7, 88, 3, 8, 3, 8, 3, 8, 5, 8, 94, 10, 8, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 310, 10, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 321, 10,
	10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	5, 11, 396, 10, 11, 3, 12, 3, 12, 5, 12, 400, 10, 12, 3, 13, 3, 13, 5,
	13, 404, 10, 13, 3, 13, 5, 13, 407, 10, 13, 3, 13, 3, 13, 5, 13, 411, 10,
	13, 3, 14, 3, 14, 6, 14, 415, 10, 14, 13, 14, 14, 14, 416, 3, 15, 3, 15,
	3, 16, 5, 16, 422, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 428, 10,
	16, 3, 17, 3, 17, 7, 17, 432, 10, 17, 12, 17, 14, 17, 435, 11, 17, 3, 17,
	7, 17, 438, 10, 17, 12, 17, 14, 17, 441, 11, 17, 3, 18, 3, 18, 3, 18, 5,
	18, 446, 10, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 5, 19, 454,
	10, 19, 3, 19, 7, 19, 457, 10, 19, 12, 19, 14, 19, 460, 11, 19, 3, 20,
	3, 20, 3, 20, 3, 20, 5, 20, 466, 10, 20, 3, 20, 2, 2, 21, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 2, 6, 5, 2, 91,
	92, 103, 105, 108, 108, 3, 2, 75, 76, 3, 2, 77, 78, 4, 2, 78, 78, 81, 86,
	2, 675, 2, 40, 3, 2, 2, 2, 4, 48, 3, 2, 2, 2, 6, 62, 3, 2, 2, 2, 8, 64,
	3, 2, 2, 2, 10, 66, 3, 2, 2, 2, 12, 86, 3, 2, 2, 2, 14, 93, 3, 2, 2, 2,
	16, 309, 3, 2, 2, 2, 18, 320, 3, 2, 2, 2, 20, 395, 3, 2, 2, 2, 22, 397,
	3, 2, 2, 2, 24, 410, 3, 2, 2, 2, 26, 412, 3, 2, 2, 2, 28, 418, 3, 2, 2,
	2, 30, 421, 3, 2, 2, 2, 32, 429, 3, 2, 2, 2, 34, 442, 3, 2, 2, 2, 36, 450,
	3, 2, 2, 2, 38, 465, 3, 2, 2, 2, 40, 42, 5, 36, 19, 2, 41, 43, 5, 38, 20,
	2, 42, 41, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 3, 3, 2, 2, 2, 44, 49, 5,
	10, 6, 2, 45, 49, 5, 6, 4, 2, 46, 49, 5, 8, 5, 2, 47, 49, 7, 3, 2, 2, 48,
	44, 3, 2, 2, 2, 48, 45, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 47, 3, 2, 2,
	2, 49, 5, 3, 2, 2, 2, 50, 63, 7, 117, 2, 2, 51, 63, 7, 118, 2, 2, 52, 63,
	7, 129, 2, 2, 53, 63, 7, 130, 2, 2, 54, 63, 7, 134, 2, 2, 55, 63, 7, 131,
	2, 2, 56, 63, 7, 121, 2, 2, 57, 58, 7, 118, 2, 2, 58, 63, 7, 101, 2, 2,
	59, 60, 7, 117, 2, 2, 60, 63, 7, 107, 2, 2, 61, 63, 7, 124, 2, 2, 62, 50,
	3, 2, 2, 2, 62, 51, 3, 2, 2, 2, 62, 52, 3, 2, 2, 2, 62, 53, 3, 2, 2, 2,
	62, 54, 3, 2, 2, 2, 62, 55, 3, 2, 2, 2, 62, 56, 3, 2, 2, 2, 62, 57, 3,
	2, 2, 2, 62, 59, 3, 2, 2, 2, 62, 61, 3, 2, 2, 2, 63, 7, 3, 2, 2, 2, 64,
	65, 9, 2, 2, 2, 65, 9, 3, 2, 2, 2, 66, 68, 7, 4, 2, 2, 67, 69, 5, 12, 7,
	2, 68, 67, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 72,
	5, 14, 8, 2, 71, 73, 5, 20, 11, 2, 72, 71, 3, 2, 2, 2, 72, 73, 3, 2, 2,
	2, 73, 75, 3, 2, 2, 2, 74, 76, 5, 22, 12, 2, 75, 74, 3, 2, 2, 2, 75, 76,
	3, 2, 2, 2, 76, 78, 3, 2, 2, 2, 77, 79, 5, 24, 13, 2, 78, 77, 3, 2, 2,
	2, 78, 79, 3, 2, 2, 2, 79, 81, 3, 2, 2, 2, 80, 82, 5, 26, 14, 2, 81, 80,
	3, 2, 2, 2, 81, 82, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 84, 7, 5, 2, 2,
	84, 11, 3, 2, 2, 2, 85, 87, 7, 146, 2, 2, 86, 85, 3, 2, 2, 2, 87, 88, 3,
	2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 13, 3, 2, 2, 2, 90,
	94, 5, 16, 9, 2, 91, 94, 5, 18, 10, 2, 92, 94, 7, 3, 2, 2, 93, 90, 3, 2,
	2, 2, 93, 91, 3, 2, 2, 2, 93, 92, 3, 2, 2, 2, 94, 15, 3, 2, 2, 2, 95, 310,
	7, 123, 2, 2, 96, 97, 7, 123, 2, 2, 97, 310, 7, 94, 2, 2, 98, 99, 7, 127,
	2, 2, 99, 310, 7, 98, 2, 2, 100, 101, 7, 117, 2, 2, 101, 310, 7, 94, 2,
	2, 102, 310, 7, 117, 2, 2, 103, 310, 7, 118, 2, 2, 104, 310, 7, 129, 2,
	2, 105, 310, 7, 130, 2, 2, 106, 310, 7, 121, 2, 2, 107, 108, 7, 129, 2,
	2, 108, 310, 7, 94, 2, 2, 109, 110, 7, 129, 2, 2, 110, 310, 7, 90, 2, 2,
	111, 112, 7, 128, 2, 2, 112, 310, 7, 96, 2, 2, 113, 114, 7, 116, 2, 2,
	114, 310, 7, 101, 2, 2, 115, 116, 7, 134, 2, 2, 116, 310, 7, 98, 2, 2,
	117, 310, 7, 131, 2, 2, 118, 310, 7, 134, 2, 2, 119, 120, 7, 118, 2, 2,
	120, 310, 7, 101, 2, 2, 121, 122, 7, 116, 2, 2, 122, 310, 7, 107, 2, 2,
	123, 310, 7, 126, 2, 2, 124, 125, 7, 118, 2, 2, 125, 310, 7, 90, 2, 2,
	126, 127, 7, 134, 2, 2, 127, 310, 7, 92, 2, 2, 128, 129, 7, 135, 2, 2,
	129, 310, 7, 98, 2, 2, 130, 310, 7, 137, 2, 2, 131, 132, 7, 118, 2, 2,
	132, 310, 7, 107, 2, 2, 133, 134, 7, 128, 2, 2, 134, 310, 7, 103, 2, 2,
	135, 136, 7, 121, 2, 2, 136, 310, 7, 94, 2, 2, 137, 138, 7, 118, 2, 2,
	138, 310, 7, 104, 2, 2, 139, 140, 7, 129, 2, 2, 140, 310, 7, 98, 2, 2,
	141, 142, 7, 118, 2, 2, 142, 310, 7, 110, 2, 2, 143, 144, 7, 141, 2, 2,
	144, 310, 7, 103, 2, 2, 145, 146, 7, 122, 2, 2, 146, 310, 7, 90, 2, 2,
	147, 148, 7, 122, 2, 2, 148, 310, 7, 94, 2, 2, 149, 150, 7, 116, 2, 2,
	150, 310, 7, 108, 2, 2, 151, 152, 7, 134, 2, 2, 152, 310, 7, 94, 2, 2,
	153, 154, 7, 117, 2, 2, 154, 310, 7, 107, 2, 2, 155, 156, 7, 126, 2, 2,
	156, 310, 7, 107, 2, 2, 157, 158, 7, 133, 2, 2, 158, 310, 7, 91, 2, 2,
	159, 160, 7, 134, 2, 2, 160, 310, 7, 107, 2, 2, 161, 310, 7, 140, 2, 2,
	162, 163, 7, 141, 2, 2, 163, 310, 7, 107, 2, 2, 164, 165, 7, 129, 2, 2,
	165, 310, 7, 91, 2, 2, 166, 167, 7, 128, 2, 2, 167, 310, 7, 104, 2, 2,
	168, 169, 7, 135, 2, 2, 169, 310, 7, 92, 2, 2, 170, 171, 7, 133, 2, 2,
	171, 310, 7, 110, 2, 2, 172, 173, 7, 133, 2, 2, 173, 310, 7, 97, 2, 2,
	174, 175, 7, 131, 2, 2, 175, 310, 7, 93, 2, 2, 176, 177, 7, 116, 2, 2,
	177, 310, 7, 96, 2, 2, 178, 179, 7, 118, 2, 2, 179, 310, 7, 93, 2, 2, 180,
	181, 7, 124, 2, 2, 181, 310, 7, 103, 2, 2, 182, 183, 7, 134, 2, 2, 183,
	310, 7, 103, 2, 2, 184, 185, 7, 134, 2, 2, 185, 310, 7, 91, 2, 2, 186,
	187, 7, 135, 2, 2, 187, 310, 7, 94, 2, 2, 188, 310, 7, 124, 2, 2, 189,
	190, 7, 139, 2, 2, 190, 310, 7, 94, 2, 2, 191, 192, 7, 118, 2, 2, 192,
	310, 7, 108, 2, 2, 193, 194, 7, 117, 2, 2, 194, 310, 7, 90, 2, 2, 195,
	196, 7, 123, 2, 2, 196, 310, 7, 95, 2, 2, 197, 198, 7, 135, 2, 2, 198,
	310, 7, 90, 2, 2, 199, 310, 7, 138, 2, 2, 200, 201, 7, 133, 2, 2, 201,
	310, 7, 94, 2, 2, 202, 203, 7, 130, 2, 2, 203, 310, 7, 108, 2, 2, 204,
	205, 7, 124, 2, 2, 205, 310, 7, 107, 2, 2, 206, 207, 7, 131, 2, 2, 207,
	310, 7, 109, 2, 2, 208, 209, 7, 116, 2, 2, 209, 310, 7, 110, 2, 2, 210,
	211, 7, 123, 2, 2, 211, 310, 7, 96, 2, 2, 212, 213, 7, 135, 2, 2, 213,
	310, 7, 101, 2, 2, 214, 215, 7, 131, 2, 2, 215, 310, 7, 91, 2, 2, 216,
	217, 7, 117, 2, 2, 217, 310, 7, 98, 2, 2, 218, 219, 7, 131, 2, 2, 219,
	310, 7, 104, 2, 2, 220, 221, 7, 116, 2, 2, 221, 310, 7, 109, 2, 2, 222,
	223, 7, 133, 2, 2, 223, 310, 7, 103, 2, 2, 224, 225, 7, 121, 2, 2, 225,
	310, 7, 107, 2, 2, 226, 227, 7, 133, 2, 2, 227, 310, 7, 90, 2, 2, 228,
	229, 7, 133, 2, 2, 229, 310, 7, 95, 2, 2, 230, 231, 7, 119, 2, 2, 231,
	310, 7, 91, 2, 2, 232, 233, 7, 134, 2, 2, 233, 310, 7, 96, 2, 2, 234, 235,
	7, 117, 2, 2, 235, 310, 7, 97, 2, 2, 236, 237, 7, 123, 2, 2, 237, 310,
	7, 108, 2, 2, 238, 239, 7, 128, 2, 2, 239, 310, 7, 109, 2, 2, 240, 241,
	7, 119, 2, 2, 241, 310, 7, 108, 2, 2, 242, 243, 7, 133, 2, 2, 243, 310,
	7, 96, 2, 2, 244, 245, 7, 118, 2, 2, 245, 310, 7, 103, 2, 2, 246, 247,
	7, 121, 2, 2, 247, 310, 7, 101, 2, 2, 248, 249, 7, 127, 2, 2, 249, 310,
	7, 111, 2, 2, 250, 251, 7, 127, 2, 2, 251, 310, 7, 90, 2, 2, 252, 253,
	7, 118, 2, 2, 253, 310, 7, 94, 2, 2, 254, 255, 7, 131, 2, 2, 255, 310,
	7, 107, 2, 2, 256, 257, 7, 129, 2, 2, 257, 310, 7, 93, 2, 2, 258, 259,
	7, 131, 2, 2, 259, 310, 7, 102, 2, 2, 260, 261, 7, 134, 2, 2, 261, 310,
	7, 102, 2, 2, 262, 263, 7, 120, 2, 2, 263, 310, 7, 110, 2, 2, 264, 265,
	7, 122, 2, 2, 265, 310, 7, 93, 2, 2, 266, 267, 7, 135, 2, 2, 267, 310,
	7, 91, 2, 2, 268, 269, 7, 119, 2, 2, 269, 310, 7, 114, 2, 2, 270, 271,
	7, 123, 2, 2, 271, 310, 7, 104, 2, 2, 272, 273, 7, 120, 2, 2, 273, 310,
	7, 107, 2, 2, 274, 275, 7, 135, 2, 2, 275, 310, 7, 102, 2, 2, 276, 277,
	7, 140, 2, 2, 277, 310, 7, 91, 2, 2, 278, 279, 7, 127, 2, 2, 279, 310,
	7, 110, 2, 2, 280, 281, 7, 116, 2, 2, 281, 310, 7, 92, 2, 2, 282, 283,
	7, 135, 2, 2, 283, 310, 7, 97, 2, 2, 284, 285, 7, 131, 2, 2, 285, 310,
	7, 90, 2, 2, 286, 310, 7, 136, 2, 2, 287, 288, 7, 129, 2, 2, 288, 310,
	7, 105, 2, 2, 289, 290, 7, 131, 2, 2, 290, 310, 7, 110, 2, 2, 291, 292,
	7, 116, 2, 2, 292, 310, 7, 102, 2, 2, 293, 294, 7, 118, 2, 2, 294, 310,
	7, 102, 2, 2, 295, 296, 7, 117, 2, 2, 296, 310, 7, 100, 2, 2, 297, 298,
	7, 118, 2, 2, 298, 310, 7, 95, 2, 2, 299, 300, 7, 120, 2, 2, 300, 310,
	7, 108, 2, 2, 301, 302, 7, 121, 2, 2, 302, 310, 7, 102, 2, 2, 303, 304,
	7, 128, 2, 2, 304, 310, 7, 93, 2, 2, 305, 306, 7, 129, 2, 2, 306, 310,
	7, 104, 2, 2, 307, 308, 7, 127, 2, 2, 308, 310, 7, 107, 2, 2, 309, 95,
	3, 2, 2, 2, 309, 96, 3, 2, 2, 2, 309, 98, 3, 2, 2, 2, 309, 100, 3, 2, 2,
	2, 309, 102, 3, 2, 2, 2, 309, 103, 3, 2, 2, 2, 309, 104, 3, 2, 2, 2, 309,
	105, 3, 2, 2, 2, 309, 106, 3, 2, 2, 2, 309, 107, 3, 2, 2, 2, 309, 109,
	3, 2, 2, 2, 309, 111, 3, 2, 2, 2, 309, 113, 3, 2, 2, 2, 309, 115, 3, 2,
	2, 2, 309, 117, 3, 2, 2, 2, 309, 118, 3, 2, 2, 2, 309, 119, 3, 2, 2, 2,
	309, 121, 3, 2, 2, 2, 309, 123, 3, 2, 2, 2, 309, 124, 3, 2, 2, 2, 309,
	126, 3, 2, 2, 2, 309, 128, 3, 2, 2, 2, 309, 130, 3, 2, 2, 2, 309, 131,
	3, 2, 2, 2, 309, 133, 3, 2, 2, 2, 309, 135, 3, 2, 2, 2, 309, 137, 3, 2,
	2, 2, 309, 139, 3, 2, 2, 2, 309, 141, 3, 2, 2, 2, 309, 143, 3, 2, 2, 2,
	309, 145, 3, 2, 2, 2, 309, 147, 3, 2, 2, 2, 309, 149, 3, 2, 2, 2, 309,
	151, 3, 2, 2, 2, 309, 153, 3, 2, 2, 2, 309, 155, 3, 2, 2, 2, 309, 157,
	3, 2, 2, 2, 309, 159, 3, 2, 2, 2, 309, 161, 3, 2, 2, 2, 309, 162, 3, 2,
	2, 2, 309, 164, 3, 2, 2, 2, 309, 166, 3, 2, 2, 2, 309, 168, 3, 2, 2, 2,
	309, 170, 3, 2, 2, 2, 309, 172, 3, 2, 2, 2, 309, 174, 3, 2, 2, 2, 309,
	176, 3, 2, 2, 2, 309, 178, 3, 2, 2, 2, 309, 180, 3, 2, 2, 2, 309, 182,
	3, 2, 2, 2, 309, 184, 3, 2, 2, 2, 309, 186, 3, 2, 2, 2, 309, 188, 3, 2,
	2, 2, 309, 189, 3, 2, 2, 2, 309, 191, 3, 2, 2, 2, 309, 193, 3, 2, 2, 2,
	309, 195, 3, 2, 2, 2, 309, 197, 3, 2, 2, 2, 309, 199, 3, 2, 2, 2, 309,
	200, 3, 2, 2, 2, 309, 202, 3, 2, 2, 2, 309, 204, 3, 2, 2, 2, 309, 206,
	3, 2, 2, 2, 309, 208, 3, 2, 2, 2, 309, 210, 3, 2, 2, 2, 309, 212, 3, 2,
	2, 2, 309, 214, 3, 2, 2, 2, 309, 216, 3, 2, 2, 2, 309, 218, 3, 2, 2, 2,
	309, 220, 3, 2, 2, 2, 309, 222, 3, 2, 2, 2, 309, 224, 3, 2, 2, 2, 309,
	226, 3, 2, 2, 2, 309, 228, 3, 2, 2, 2, 309, 230, 3, 2, 2, 2, 309, 232,
	3, 2, 2, 2, 309, 234, 3, 2, 2, 2, 309, 236, 3, 2, 2, 2, 309, 238, 3, 2,
	2, 2, 309, 240, 3, 2, 2, 2, 309, 242, 3, 2, 2, 2, 309, 244, 3, 2, 2, 2,
	309, 246, 3, 2, 2, 2, 309, 248, 3, 2, 2, 2, 309, 250, 3, 2, 2, 2, 309,
	252, 3, 2, 2, 2, 309, 254, 3, 2, 2, 2, 309, 256, 3, 2, 2, 2, 309, 258,
	3, 2, 2, 2, 309, 260, 3, 2, 2, 2, 309, 262, 3, 2, 2, 2, 309, 264, 3, 2,
	2, 2, 309, 266, 3, 2, 2, 2, 309, 268, 3, 2, 2, 2, 309, 270, 3, 2, 2, 2,
	309, 272, 3, 2, 2, 2, 309, 274, 3, 2, 2, 2, 309, 276, 3, 2, 2, 2, 309,
	278, 3, 2, 2, 2, 309, 280, 3, 2, 2, 2, 309, 282, 3, 2, 2, 2, 309, 284,
	3, 2, 2, 2, 309, 286, 3, 2, 2, 2, 309, 287, 3, 2, 2, 2, 309, 289, 3, 2,
	2, 2, 309, 291, 3, 2, 2, 2, 309, 293, 3, 2, 2, 2, 309, 295, 3, 2, 2, 2,
	309, 297, 3, 2, 2, 2, 309, 299, 3, 2, 2, 2, 309, 301, 3, 2, 2, 2, 309,
	303, 3, 2, 2, 2, 309, 305, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2, 310, 17, 3,
	2, 2, 2, 311, 321, 7, 92, 2, 2, 312, 321, 7, 103, 2, 2, 313, 321, 7, 104,
	2, 2, 314, 321, 7, 105, 2, 2, 315, 321, 7, 108, 2, 2, 316, 317, 7, 108,
	2, 2, 317, 321, 7, 94, 2, 2, 318, 319, 7, 90, 2, 2, 319, 321, 7, 108, 2,
	2, 320, 311, 3, 2, 2, 2, 320, 312, 3, 2, 2, 2, 320, 313, 3, 2, 2, 2, 320,
	314, 3, 2, 2, 2, 320, 315, 3, 2, 2, 2, 320, 316, 3, 2, 2, 2, 320, 318,
	3, 2, 2, 2, 321, 19, 3, 2, 2, 2, 322, 396, 7, 6, 2, 2, 323, 396, 7, 7,
	2, 2, 324, 396, 7, 8, 2, 2, 325, 396, 7, 9, 2, 2, 326, 396, 7, 10, 2, 2,
	327, 396, 7, 11, 2, 2, 328, 396, 7, 12, 2, 2, 329, 396, 7, 13, 2, 2, 330,
	396, 7, 14, 2, 2, 331, 396, 7, 15, 2, 2, 332, 396, 7, 16, 2, 2, 333, 396,
	7, 17, 2, 2, 334, 396, 7, 17, 2, 2, 335, 396, 7, 18, 2, 2, 336, 396, 7,
	19, 2, 2, 337, 396, 7, 20, 2, 2, 338, 396, 7, 21, 2, 2, 339, 396, 7, 22,
	2, 2, 340, 396, 7, 23, 2, 2, 341, 396, 7, 24, 2, 2, 342, 396, 7, 25, 2,
	2, 343, 396, 7, 26, 2, 2, 344, 396, 7, 27, 2, 2, 345, 396, 7, 28, 2, 2,
	346, 396, 7, 29, 2, 2, 347, 396, 7, 30, 2, 2, 348, 396, 7, 31, 2, 2, 349,
	396, 7, 32, 2, 2, 350, 396, 7, 33, 2, 2, 351, 396, 7, 34, 2, 2, 352, 396,
	7, 35, 2, 2, 353, 396, 7, 36, 2, 2, 354, 396, 7, 37, 2, 2, 355, 396, 7,
	38, 2, 2, 356, 396, 7, 39, 2, 2, 357, 396, 7, 40, 2, 2, 358, 396, 7, 41,
	2, 2, 359, 396, 7, 42, 2, 2, 360, 396, 7, 43, 2, 2, 361, 396, 7, 44, 2,
	2, 362, 396, 7, 45, 2, 2, 363, 396, 7, 46, 2, 2, 364, 396, 7, 47, 2, 2,
	365, 396, 7, 48, 2, 2, 366, 396, 7, 49, 2, 2, 367, 396, 7, 50, 2, 2, 368,
	396, 7, 51, 2, 2, 369, 396, 7, 52, 2, 2, 370, 396, 7, 53, 2, 2, 371, 396,
	7, 54, 2, 2, 372, 396, 7, 55, 2, 2, 373, 396, 7, 56, 2, 2, 374, 396, 7,
	57, 2, 2, 375, 396, 7, 58, 2, 2, 376, 396, 7, 59, 2, 2, 377, 396, 7, 60,
	2, 2, 378, 396, 7, 61, 2, 2, 379, 396, 7, 62, 2, 2, 380, 396, 7, 63, 2,
	2, 381, 396, 7, 64, 2, 2, 382, 396, 7, 65, 2, 2, 383, 396, 7, 66, 2, 2,
	384, 396, 7, 67, 2, 2, 385, 396, 7, 68, 2, 2, 386, 396, 7, 69, 2, 2, 387,
	396, 7, 70, 2, 2, 388, 396, 7, 71, 2, 2, 389, 396, 7, 72, 2, 2, 390, 396,
	7, 73, 2, 2, 391, 396, 7, 74, 2, 2, 392, 393, 9, 3, 2, 2, 393, 394, 7,
	146, 2, 2, 394, 396, 7, 146, 2, 2, 395, 322, 3, 2, 2, 2, 395, 323, 3, 2,
	2, 2, 395, 324, 3, 2, 2, 2, 395, 325, 3, 2, 2, 2, 395, 326, 3, 2, 2, 2,
	395, 327, 3, 2, 2, 2, 395, 328, 3, 2, 2, 2, 395, 329, 3, 2, 2, 2, 395,
	330, 3, 2, 2, 2, 395, 331, 3, 2, 2, 2, 395, 332, 3, 2, 2, 2, 395, 333,
	3, 2, 2, 2, 395, 334, 3, 2, 2, 2, 395, 335, 3, 2, 2, 2, 395, 336, 3, 2,
	2, 2, 395, 337, 3, 2, 2, 2, 395, 338, 3, 2, 2, 2, 395, 339, 3, 2, 2, 2,
	395, 340, 3, 2, 2, 2, 395, 341, 3, 2, 2, 2, 395, 342, 3, 2, 2, 2, 395,
	343, 3, 2, 2, 2, 395, 344, 3, 2, 2, 2, 395, 345, 3, 2, 2, 2, 395, 346,
	3, 2, 2, 2, 395, 347, 3, 2, 2, 2, 395, 348, 3, 2, 2, 2, 395, 349, 3, 2,
	2, 2, 395, 350, 3, 2, 2, 2, 395, 351, 3, 2, 2, 2, 395, 352, 3, 2, 2, 2,
	395, 353, 3, 2, 2, 2, 395, 354, 3, 2, 2, 2, 395, 355, 3, 2, 2, 2, 395,
	356, 3, 2, 2, 2, 395, 357, 3, 2, 2, 2, 395, 358, 3, 2, 2, 2, 395, 359,
	3, 2, 2, 2, 395, 360, 3, 2, 2, 2, 395, 361, 3, 2, 2, 2, 395, 362, 3, 2,
	2, 2, 395, 363, 3, 2, 2, 2, 395, 364, 3, 2, 2, 2, 395, 365, 3, 2, 2, 2,
	395, 366, 3, 2, 2, 2, 395, 367, 3, 2, 2, 2, 395, 368, 3, 2, 2, 2, 395,
	369, 3, 2, 2, 2, 395, 370, 3, 2, 2, 2, 395, 371, 3, 2, 2, 2, 395, 372,
	3, 2, 2, 2, 395, 373, 3, 2, 2, 2, 395, 374, 3, 2, 2, 2, 395, 375, 3, 2,
	2, 2, 395, 376, 3, 2, 2, 2, 395, 377, 3, 2, 2, 2, 395, 378, 3, 2, 2, 2,
	395, 379, 3, 2, 2, 2, 395, 380, 3, 2, 2, 2, 395, 381, 3, 2, 2, 2, 395,
	382, 3, 2, 2, 2, 395, 383, 3, 2, 2, 2, 395, 384, 3, 2, 2, 2, 395, 385,
	3, 2, 2, 2, 395, 386, 3, 2, 2, 2, 395, 387, 3, 2, 2, 2, 395, 388, 3, 2,
	2, 2, 395, 389, 3, 2, 2, 2, 395, 390, 3, 2, 2, 2, 395, 391, 3, 2, 2, 2,
	395, 392, 3, 2, 2, 2, 396, 21, 3, 2, 2, 2, 397, 399, 7, 123, 2, 2, 398,
	400, 7, 146, 2, 2, 399, 398, 3, 2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 23,
	3, 2, 2, 2, 401, 403, 9, 4, 2, 2, 402, 404, 7, 146, 2, 2, 403, 402, 3,
	2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 406, 3, 2, 2, 2, 405, 407, 7, 146,
	2, 2, 406, 405, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407, 411, 3, 2, 2, 2,
	408, 411, 7, 79, 2, 2, 409, 411, 7, 80, 2, 2, 410, 401, 3, 2, 2, 2, 410,
	408, 3, 2, 2, 2, 410, 409, 3, 2, 2, 2, 411, 25, 3, 2, 2, 2, 412, 414, 7,
	81, 2, 2, 413, 415, 7, 146, 2, 2, 414, 413, 3, 2, 2, 2, 415, 416, 3, 2,
	2, 2, 416, 414, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 27, 3, 2, 2, 2,
	418, 419, 9, 5, 2, 2, 419, 29, 3, 2, 2, 2, 420, 422, 5, 28, 15, 2, 421,
	420, 3, 2, 2, 2, 421, 422, 3, 2, 2, 2, 422, 427, 3, 2, 2, 2, 423, 428,
	7, 146, 2, 2, 424, 425, 7, 87, 2, 2, 425, 426, 7, 146, 2, 2, 426, 428,
	7, 146, 2, 2, 427, 423, 3, 2, 2, 2, 427, 424, 3, 2, 2, 2, 428, 31, 3, 2,
	2, 2, 429, 433, 5, 4, 3, 2, 430, 432, 5, 30, 16, 2, 431, 430, 3, 2, 2,
	2, 432, 435, 3, 2, 2, 2, 433, 431, 3, 2, 2, 2, 433, 434, 3, 2, 2, 2, 434,
	439, 3, 2, 2, 2, 435, 433, 3, 2, 2, 2, 436, 438, 5, 34, 18, 2, 437, 436,
	3, 2, 2, 2, 438, 441, 3, 2, 2, 2, 439, 437, 3, 2, 2, 2, 439, 440, 3, 2,
	2, 2, 440, 33, 3, 2, 2, 2, 441, 439, 3, 2, 2, 2, 442, 445, 7, 88, 2, 2,
	443, 446, 5, 28, 15, 2, 444, 446, 7, 142, 2, 2, 445, 443, 3, 2, 2, 2, 445,
	444, 3, 2, 2, 2, 445, 446, 3, 2, 2, 2, 446, 447, 3, 2, 2, 2, 447, 448,
	5, 36, 19, 2, 448, 449, 7, 89, 2, 2, 449, 35, 3, 2, 2, 2, 450, 458, 5,
	32, 17, 2, 451, 454, 5, 28, 15, 2, 452, 454, 7, 142, 2, 2, 453, 451, 3,
	2, 2, 2, 453, 452, 3, 2, 2, 2, 453, 454, 3, 2, 2, 2, 454, 455, 3, 2, 2,
	2, 455, 457, 5, 32, 17, 2, 456, 453, 3, 2, 2, 2, 457, 460, 3, 2, 2, 2,
	458, 456, 3, 2, 2, 2, 458, 459, 3, 2, 2, 2, 459, 37, 3, 2, 2, 2, 460, 458,
	3, 2, 2, 2, 461, 462, 7, 145, 2, 2, 462, 466, 7, 147, 2, 2, 463, 466, 7,
	143, 2, 2, 464, 466, 7, 144, 2, 2, 465, 461, 3, 2, 2, 2, 465, 463, 3, 2,
	2, 2, 465, 464, 3, 2, 2, 2, 466, 39, 3, 2, 2, 2, 28, 42, 48, 62, 68, 72,
	75, 78, 81, 88, 93, 309, 320, 395, 399, 403, 406, 410, 416, 421, 427, 433,
	439, 445, 453, 458, 465,
}
var literalNames = []string{
	"", "'*'", "'['", "']'", "'@'", "'@@'", "'@TH1'", "'@TH2'", "'@AL1'", "'@AL2'",
	"'@SP1'", "'@SP2'", "'@SP3'", "'@TB1'", "'@TB2'", "'@TB3'", "'@TB4'", "'@TB5'",
	"'@TB6'", "'@TB7'", "'@TB8'", "'@TB9'", "'@TB10'", "'@TB11'", "'@TB12'",
	"'@TB13'", "'@TB14'", "'@TB15'", "'@TB16'", "'@TB17'", "'@TB18'", "'@TB19'",
	"'@TB20'", "'@TB21'", "'@TB22'", "'@TB23'", "'@TB24'", "'@TB25'", "'@TB26'",
	"'@TB27'", "'@TB28'", "'@TB29'", "'@TB30'", "'@OH1'", "'@OH2'", "'@OH3'",
	"'@OH4'", "'@OH5'", "'@OH6'", "'@OH7'", "'@OH8'", "'@OH9'", "'@OH10'",
	"'@OH11'", "'@OH12'", "'@OH13'", "'@OH14'", "'@OH15'", "'@OH16'", "'@OH17'",
	"'@OH18'", "'@OH19'", "'@OH20'", "'@OH21'", "'@OH22'", "'@OH23'", "'@OH24'",
	"'@OH25'", "'@OH26'", "'@OH27'", "'@OH28'", "'@OH29'", "'@OH30'", "'@TB'",
	"'@OH'", "'+'", "'-'", "'--'", "'++'", "':'", "'='", "'#'", "'$'", "'/'",
	"'\\'", "'%'", "'('", "')'", "'a'", "'b'", "'c'", "'d'", "'e'", "'f'",
	"'g'", "'h'", "'i'", "'j'", "'k'", "'l'", "'m'", "'n'", "'o'", "'p'", "'q'",
	"'r'", "'s'", "'t'", "'u'", "'v'", "'w'", "'x'", "'y'", "'z'", "'A'", "'B'",
	"'C'", "'D'", "'E'", "'F'", "'G'", "'H'", "'I'", "'J'", "'K'", "'L'", "'M'",
	"'N'", "'O'", "'P'", "'Q'", "'R'", "'S'", "'T'", "'U'", "'V'", "'W'", "'X'",
	"'Y'", "'Z'", "'.'", "'\r'", "'\n'", "' '", "", "'\t'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "LA", "LB",
	"LC", "LD", "LE", "LF", "LG", "LH", "LI", "LJ", "LK", "LL", "LM", "LN",
	"LO", "LP", "LQ", "LR", "LS", "LT", "LU", "LV", "LW", "LX", "LY", "LZ",
	"UA", "UB", "UC", "UD", "UE", "UF", "UG", "UH", "UI", "UJ", "UK", "UL",
	"UM", "UN", "UO", "UP", "UQ", "UR", "US", "UT", "UU", "UV", "UW", "UX",
	"UY", "UZ", "DOT", "LINEFEED", "CARRIAGE_RETURN", "SPACE", "DIGIT", "TAB",
}

var ruleNames = []string{
	"smiles", "atom", "aliphatic_organic", "aromatic_organic", "bracket_atom",
	"isotope", "symbol", "element_symbols", "aromatic_symbol", "chiral", "hcount",
	"charge", "class_", "bond", "ringbond", "branched_atom", "branch", "chain",
	"terminator",
}

type smilesParser struct {
	*antlr.BaseParser
}

// NewsmilesParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *smilesParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewsmilesParser(input antlr.TokenStream) *smilesParser {
	this := new(smilesParser)
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
	this.GrammarFileName = "smiles.g4"

	return this
}

// smilesParser tokens.
const (
	smilesParserEOF             = antlr.TokenEOF
	smilesParserT__0            = 1
	smilesParserT__1            = 2
	smilesParserT__2            = 3
	smilesParserT__3            = 4
	smilesParserT__4            = 5
	smilesParserT__5            = 6
	smilesParserT__6            = 7
	smilesParserT__7            = 8
	smilesParserT__8            = 9
	smilesParserT__9            = 10
	smilesParserT__10           = 11
	smilesParserT__11           = 12
	smilesParserT__12           = 13
	smilesParserT__13           = 14
	smilesParserT__14           = 15
	smilesParserT__15           = 16
	smilesParserT__16           = 17
	smilesParserT__17           = 18
	smilesParserT__18           = 19
	smilesParserT__19           = 20
	smilesParserT__20           = 21
	smilesParserT__21           = 22
	smilesParserT__22           = 23
	smilesParserT__23           = 24
	smilesParserT__24           = 25
	smilesParserT__25           = 26
	smilesParserT__26           = 27
	smilesParserT__27           = 28
	smilesParserT__28           = 29
	smilesParserT__29           = 30
	smilesParserT__30           = 31
	smilesParserT__31           = 32
	smilesParserT__32           = 33
	smilesParserT__33           = 34
	smilesParserT__34           = 35
	smilesParserT__35           = 36
	smilesParserT__36           = 37
	smilesParserT__37           = 38
	smilesParserT__38           = 39
	smilesParserT__39           = 40
	smilesParserT__40           = 41
	smilesParserT__41           = 42
	smilesParserT__42           = 43
	smilesParserT__43           = 44
	smilesParserT__44           = 45
	smilesParserT__45           = 46
	smilesParserT__46           = 47
	smilesParserT__47           = 48
	smilesParserT__48           = 49
	smilesParserT__49           = 50
	smilesParserT__50           = 51
	smilesParserT__51           = 52
	smilesParserT__52           = 53
	smilesParserT__53           = 54
	smilesParserT__54           = 55
	smilesParserT__55           = 56
	smilesParserT__56           = 57
	smilesParserT__57           = 58
	smilesParserT__58           = 59
	smilesParserT__59           = 60
	smilesParserT__60           = 61
	smilesParserT__61           = 62
	smilesParserT__62           = 63
	smilesParserT__63           = 64
	smilesParserT__64           = 65
	smilesParserT__65           = 66
	smilesParserT__66           = 67
	smilesParserT__67           = 68
	smilesParserT__68           = 69
	smilesParserT__69           = 70
	smilesParserT__70           = 71
	smilesParserT__71           = 72
	smilesParserT__72           = 73
	smilesParserT__73           = 74
	smilesParserT__74           = 75
	smilesParserT__75           = 76
	smilesParserT__76           = 77
	smilesParserT__77           = 78
	smilesParserT__78           = 79
	smilesParserT__79           = 80
	smilesParserT__80           = 81
	smilesParserT__81           = 82
	smilesParserT__82           = 83
	smilesParserT__83           = 84
	smilesParserT__84           = 85
	smilesParserT__85           = 86
	smilesParserT__86           = 87
	smilesParserLA              = 88
	smilesParserLB              = 89
	smilesParserLC              = 90
	smilesParserLD              = 91
	smilesParserLE              = 92
	smilesParserLF              = 93
	smilesParserLG              = 94
	smilesParserLH              = 95
	smilesParserLI              = 96
	smilesParserLJ              = 97
	smilesParserLK              = 98
	smilesParserLL              = 99
	smilesParserLM              = 100
	smilesParserLN              = 101
	smilesParserLO              = 102
	smilesParserLP              = 103
	smilesParserLQ              = 104
	smilesParserLR              = 105
	smilesParserLS              = 106
	smilesParserLT              = 107
	smilesParserLU              = 108
	smilesParserLV              = 109
	smilesParserLW              = 110
	smilesParserLX              = 111
	smilesParserLY              = 112
	smilesParserLZ              = 113
	smilesParserUA              = 114
	smilesParserUB              = 115
	smilesParserUC              = 116
	smilesParserUD              = 117
	smilesParserUE              = 118
	smilesParserUF              = 119
	smilesParserUG              = 120
	smilesParserUH              = 121
	smilesParserUI              = 122
	smilesParserUJ              = 123
	smilesParserUK              = 124
	smilesParserUL              = 125
	smilesParserUM              = 126
	smilesParserUN              = 127
	smilesParserUO              = 128
	smilesParserUP              = 129
	smilesParserUQ              = 130
	smilesParserUR              = 131
	smilesParserUS              = 132
	smilesParserUT              = 133
	smilesParserUU              = 134
	smilesParserUV              = 135
	smilesParserUW              = 136
	smilesParserUX              = 137
	smilesParserUY              = 138
	smilesParserUZ              = 139
	smilesParserDOT             = 140
	smilesParserLINEFEED        = 141
	smilesParserCARRIAGE_RETURN = 142
	smilesParserSPACE           = 143
	smilesParserDIGIT           = 144
	smilesParserTAB             = 145
)

// smilesParser rules.
const (
	smilesParserRULE_smiles            = 0
	smilesParserRULE_atom              = 1
	smilesParserRULE_aliphatic_organic = 2
	smilesParserRULE_aromatic_organic  = 3
	smilesParserRULE_bracket_atom      = 4
	smilesParserRULE_isotope           = 5
	smilesParserRULE_symbol            = 6
	smilesParserRULE_element_symbols   = 7
	smilesParserRULE_aromatic_symbol   = 8
	smilesParserRULE_chiral            = 9
	smilesParserRULE_hcount            = 10
	smilesParserRULE_charge            = 11
	smilesParserRULE_class_            = 12
	smilesParserRULE_bond              = 13
	smilesParserRULE_ringbond          = 14
	smilesParserRULE_branched_atom     = 15
	smilesParserRULE_branch            = 16
	smilesParserRULE_chain             = 17
	smilesParserRULE_terminator        = 18
)

// ISmilesContext is an interface to support dynamic dispatch.
type ISmilesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSmilesContext differentiates from other interfaces.
	IsSmilesContext()
}

type SmilesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySmilesContext() *SmilesContext {
	var p = new(SmilesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_smiles
	return p
}

func (*SmilesContext) IsSmilesContext() {}

func NewSmilesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SmilesContext {
	var p = new(SmilesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_smiles

	return p
}

func (s *SmilesContext) GetParser() antlr.Parser { return s.parser }

func (s *SmilesContext) Chain() IChainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChainContext)
}

func (s *SmilesContext) Terminator() ITerminatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITerminatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITerminatorContext)
}

func (s *SmilesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SmilesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SmilesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterSmiles(s)
	}
}

func (s *SmilesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitSmiles(s)
	}
}

func (p *smilesParser) Smiles() (localctx ISmilesContext) {
	this := p
	_ = this

	localctx = NewSmilesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, smilesParserRULE_smiles)
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
		p.SetState(38)
		p.Chain()
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-141)&-(0x1f+1)) == 0 && ((1<<uint((_la-141)))&((1<<(smilesParserLINEFEED-141))|(1<<(smilesParserCARRIAGE_RETURN-141))|(1<<(smilesParserSPACE-141)))) != 0 {
		{
			p.SetState(39)
			p.Terminator()
		}

	}

	return localctx
}

// IAtomContext is an interface to support dynamic dispatch.
type IAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomContext differentiates from other interfaces.
	IsAtomContext()
}

type AtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomContext() *AtomContext {
	var p = new(AtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) Bracket_atom() IBracket_atomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBracket_atomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBracket_atomContext)
}

func (s *AtomContext) Aliphatic_organic() IAliphatic_organicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliphatic_organicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliphatic_organicContext)
}

func (s *AtomContext) Aromatic_organic() IAromatic_organicContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAromatic_organicContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAromatic_organicContext)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *smilesParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, smilesParserRULE_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(46)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case smilesParserT__1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Bracket_atom()
		}

	case smilesParserUB, smilesParserUC, smilesParserUF, smilesParserUI, smilesParserUN, smilesParserUO, smilesParserUP, smilesParserUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(43)
			p.Aliphatic_organic()
		}

	case smilesParserLB, smilesParserLC, smilesParserLN, smilesParserLO, smilesParserLP, smilesParserLS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(44)
			p.Aromatic_organic()
		}

	case smilesParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(45)
			p.Match(smilesParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAliphatic_organicContext is an interface to support dynamic dispatch.
type IAliphatic_organicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliphatic_organicContext differentiates from other interfaces.
	IsAliphatic_organicContext()
}

type Aliphatic_organicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliphatic_organicContext() *Aliphatic_organicContext {
	var p = new(Aliphatic_organicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_aliphatic_organic
	return p
}

func (*Aliphatic_organicContext) IsAliphatic_organicContext() {}

func NewAliphatic_organicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aliphatic_organicContext {
	var p = new(Aliphatic_organicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_aliphatic_organic

	return p
}

func (s *Aliphatic_organicContext) GetParser() antlr.Parser { return s.parser }

func (s *Aliphatic_organicContext) UB() antlr.TerminalNode {
	return s.GetToken(smilesParserUB, 0)
}

func (s *Aliphatic_organicContext) UC() antlr.TerminalNode {
	return s.GetToken(smilesParserUC, 0)
}

func (s *Aliphatic_organicContext) UN() antlr.TerminalNode {
	return s.GetToken(smilesParserUN, 0)
}

func (s *Aliphatic_organicContext) UO() antlr.TerminalNode {
	return s.GetToken(smilesParserUO, 0)
}

func (s *Aliphatic_organicContext) US() antlr.TerminalNode {
	return s.GetToken(smilesParserUS, 0)
}

func (s *Aliphatic_organicContext) UP() antlr.TerminalNode {
	return s.GetToken(smilesParserUP, 0)
}

func (s *Aliphatic_organicContext) UF() antlr.TerminalNode {
	return s.GetToken(smilesParserUF, 0)
}

func (s *Aliphatic_organicContext) LL() antlr.TerminalNode {
	return s.GetToken(smilesParserLL, 0)
}

func (s *Aliphatic_organicContext) LR() antlr.TerminalNode {
	return s.GetToken(smilesParserLR, 0)
}

func (s *Aliphatic_organicContext) UI() antlr.TerminalNode {
	return s.GetToken(smilesParserUI, 0)
}

func (s *Aliphatic_organicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aliphatic_organicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Aliphatic_organicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterAliphatic_organic(s)
	}
}

func (s *Aliphatic_organicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitAliphatic_organic(s)
	}
}

func (p *smilesParser) Aliphatic_organic() (localctx IAliphatic_organicContext) {
	this := p
	_ = this

	localctx = NewAliphatic_organicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, smilesParserRULE_aliphatic_organic)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Match(smilesParserUB)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(49)
			p.Match(smilesParserUC)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Match(smilesParserUN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(51)
			p.Match(smilesParserUO)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(52)
			p.Match(smilesParserUS)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(53)
			p.Match(smilesParserUP)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(54)
			p.Match(smilesParserUF)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(55)
			p.Match(smilesParserUC)
		}
		{
			p.SetState(56)
			p.Match(smilesParserLL)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(57)
			p.Match(smilesParserUB)
		}
		{
			p.SetState(58)
			p.Match(smilesParserLR)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(59)
			p.Match(smilesParserUI)
		}

	}

	return localctx
}

// IAromatic_organicContext is an interface to support dynamic dispatch.
type IAromatic_organicContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAromatic_organicContext differentiates from other interfaces.
	IsAromatic_organicContext()
}

type Aromatic_organicContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAromatic_organicContext() *Aromatic_organicContext {
	var p = new(Aromatic_organicContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_aromatic_organic
	return p
}

func (*Aromatic_organicContext) IsAromatic_organicContext() {}

func NewAromatic_organicContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aromatic_organicContext {
	var p = new(Aromatic_organicContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_aromatic_organic

	return p
}

func (s *Aromatic_organicContext) GetParser() antlr.Parser { return s.parser }

func (s *Aromatic_organicContext) LB() antlr.TerminalNode {
	return s.GetToken(smilesParserLB, 0)
}

func (s *Aromatic_organicContext) LC() antlr.TerminalNode {
	return s.GetToken(smilesParserLC, 0)
}

func (s *Aromatic_organicContext) LN() antlr.TerminalNode {
	return s.GetToken(smilesParserLN, 0)
}

func (s *Aromatic_organicContext) LO() antlr.TerminalNode {
	return s.GetToken(smilesParserLO, 0)
}

func (s *Aromatic_organicContext) LS() antlr.TerminalNode {
	return s.GetToken(smilesParserLS, 0)
}

func (s *Aromatic_organicContext) LP() antlr.TerminalNode {
	return s.GetToken(smilesParserLP, 0)
}

func (s *Aromatic_organicContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aromatic_organicContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Aromatic_organicContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterAromatic_organic(s)
	}
}

func (s *Aromatic_organicContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitAromatic_organic(s)
	}
}

func (p *smilesParser) Aromatic_organic() (localctx IAromatic_organicContext) {
	this := p
	_ = this

	localctx = NewAromatic_organicContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, smilesParserRULE_aromatic_organic)
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
		p.SetState(62)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-89)&-(0x1f+1)) == 0 && ((1<<uint((_la-89)))&((1<<(smilesParserLB-89))|(1<<(smilesParserLC-89))|(1<<(smilesParserLN-89))|(1<<(smilesParserLO-89))|(1<<(smilesParserLP-89))|(1<<(smilesParserLS-89)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBracket_atomContext is an interface to support dynamic dispatch.
type IBracket_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBracket_atomContext differentiates from other interfaces.
	IsBracket_atomContext()
}

type Bracket_atomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBracket_atomContext() *Bracket_atomContext {
	var p = new(Bracket_atomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_bracket_atom
	return p
}

func (*Bracket_atomContext) IsBracket_atomContext() {}

func NewBracket_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bracket_atomContext {
	var p = new(Bracket_atomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_bracket_atom

	return p
}

func (s *Bracket_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Bracket_atomContext) Symbol() ISymbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISymbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISymbolContext)
}

func (s *Bracket_atomContext) Isotope() IIsotopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIsotopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIsotopeContext)
}

func (s *Bracket_atomContext) Chiral() IChiralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChiralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChiralContext)
}

func (s *Bracket_atomContext) Hcount() IHcountContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHcountContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHcountContext)
}

func (s *Bracket_atomContext) Charge() IChargeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChargeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChargeContext)
}

func (s *Bracket_atomContext) Class_() IClass_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClass_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClass_Context)
}

func (s *Bracket_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bracket_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bracket_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterBracket_atom(s)
	}
}

func (s *Bracket_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitBracket_atom(s)
	}
}

func (p *smilesParser) Bracket_atom() (localctx IBracket_atomContext) {
	this := p
	_ = this

	localctx = NewBracket_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, smilesParserRULE_bracket_atom)
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
		p.SetState(64)
		p.Match(smilesParserT__1)
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == smilesParserDIGIT {
		{
			p.SetState(65)
			p.Isotope()
		}

	}
	{
		p.SetState(68)
		p.Symbol()
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<smilesParserT__3)|(1<<smilesParserT__4)|(1<<smilesParserT__5)|(1<<smilesParserT__6)|(1<<smilesParserT__7)|(1<<smilesParserT__8)|(1<<smilesParserT__9)|(1<<smilesParserT__10)|(1<<smilesParserT__11)|(1<<smilesParserT__12)|(1<<smilesParserT__13)|(1<<smilesParserT__14)|(1<<smilesParserT__15)|(1<<smilesParserT__16)|(1<<smilesParserT__17)|(1<<smilesParserT__18)|(1<<smilesParserT__19)|(1<<smilesParserT__20)|(1<<smilesParserT__21)|(1<<smilesParserT__22)|(1<<smilesParserT__23)|(1<<smilesParserT__24)|(1<<smilesParserT__25)|(1<<smilesParserT__26)|(1<<smilesParserT__27)|(1<<smilesParserT__28)|(1<<smilesParserT__29)|(1<<smilesParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(smilesParserT__31-32))|(1<<(smilesParserT__32-32))|(1<<(smilesParserT__33-32))|(1<<(smilesParserT__34-32))|(1<<(smilesParserT__35-32))|(1<<(smilesParserT__36-32))|(1<<(smilesParserT__37-32))|(1<<(smilesParserT__38-32))|(1<<(smilesParserT__39-32))|(1<<(smilesParserT__40-32))|(1<<(smilesParserT__41-32))|(1<<(smilesParserT__42-32))|(1<<(smilesParserT__43-32))|(1<<(smilesParserT__44-32))|(1<<(smilesParserT__45-32))|(1<<(smilesParserT__46-32))|(1<<(smilesParserT__47-32))|(1<<(smilesParserT__48-32))|(1<<(smilesParserT__49-32))|(1<<(smilesParserT__50-32))|(1<<(smilesParserT__51-32))|(1<<(smilesParserT__52-32))|(1<<(smilesParserT__53-32))|(1<<(smilesParserT__54-32))|(1<<(smilesParserT__55-32))|(1<<(smilesParserT__56-32))|(1<<(smilesParserT__57-32))|(1<<(smilesParserT__58-32))|(1<<(smilesParserT__59-32))|(1<<(smilesParserT__60-32))|(1<<(smilesParserT__61-32))|(1<<(smilesParserT__62-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(smilesParserT__63-64))|(1<<(smilesParserT__64-64))|(1<<(smilesParserT__65-64))|(1<<(smilesParserT__66-64))|(1<<(smilesParserT__67-64))|(1<<(smilesParserT__68-64))|(1<<(smilesParserT__69-64))|(1<<(smilesParserT__70-64))|(1<<(smilesParserT__71-64))|(1<<(smilesParserT__72-64))|(1<<(smilesParserT__73-64)))) != 0) {
		{
			p.SetState(69)
			p.Chiral()
		}

	}
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == smilesParserUH {
		{
			p.SetState(72)
			p.Hcount()
		}

	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-75)&-(0x1f+1)) == 0 && ((1<<uint((_la-75)))&((1<<(smilesParserT__74-75))|(1<<(smilesParserT__75-75))|(1<<(smilesParserT__76-75))|(1<<(smilesParserT__77-75)))) != 0 {
		{
			p.SetState(75)
			p.Charge()
		}

	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == smilesParserT__78 {
		{
			p.SetState(78)
			p.Class_()
		}

	}
	{
		p.SetState(81)
		p.Match(smilesParserT__2)
	}

	return localctx
}

// IIsotopeContext is an interface to support dynamic dispatch.
type IIsotopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIsotopeContext differentiates from other interfaces.
	IsIsotopeContext()
}

type IsotopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIsotopeContext() *IsotopeContext {
	var p = new(IsotopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_isotope
	return p
}

func (*IsotopeContext) IsIsotopeContext() {}

func NewIsotopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IsotopeContext {
	var p = new(IsotopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_isotope

	return p
}

func (s *IsotopeContext) GetParser() antlr.Parser { return s.parser }

func (s *IsotopeContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(smilesParserDIGIT)
}

func (s *IsotopeContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(smilesParserDIGIT, i)
}

func (s *IsotopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsotopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IsotopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterIsotope(s)
	}
}

func (s *IsotopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitIsotope(s)
	}
}

func (p *smilesParser) Isotope() (localctx IIsotopeContext) {
	this := p
	_ = this

	localctx = NewIsotopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, smilesParserRULE_isotope)
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
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == smilesParserDIGIT {
		{
			p.SetState(83)
			p.Match(smilesParserDIGIT)
		}

		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISymbolContext is an interface to support dynamic dispatch.
type ISymbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSymbolContext differentiates from other interfaces.
	IsSymbolContext()
}

type SymbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySymbolContext() *SymbolContext {
	var p = new(SymbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_symbol
	return p
}

func (*SymbolContext) IsSymbolContext() {}

func NewSymbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SymbolContext {
	var p = new(SymbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_symbol

	return p
}

func (s *SymbolContext) GetParser() antlr.Parser { return s.parser }

func (s *SymbolContext) Element_symbols() IElement_symbolsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElement_symbolsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElement_symbolsContext)
}

func (s *SymbolContext) Aromatic_symbol() IAromatic_symbolContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAromatic_symbolContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAromatic_symbolContext)
}

func (s *SymbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SymbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SymbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterSymbol(s)
	}
}

func (s *SymbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitSymbol(s)
	}
}

func (p *smilesParser) Symbol() (localctx ISymbolContext) {
	this := p
	_ = this

	localctx = NewSymbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, smilesParserRULE_symbol)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(91)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case smilesParserUA, smilesParserUB, smilesParserUC, smilesParserUD, smilesParserUE, smilesParserUF, smilesParserUG, smilesParserUH, smilesParserUI, smilesParserUK, smilesParserUL, smilesParserUM, smilesParserUN, smilesParserUO, smilesParserUP, smilesParserUR, smilesParserUS, smilesParserUT, smilesParserUU, smilesParserUV, smilesParserUW, smilesParserUX, smilesParserUY, smilesParserUZ:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Element_symbols()
		}

	case smilesParserLA, smilesParserLC, smilesParserLN, smilesParserLO, smilesParserLP, smilesParserLS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Aromatic_symbol()
		}

	case smilesParserT__0:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(90)
			p.Match(smilesParserT__0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IElement_symbolsContext is an interface to support dynamic dispatch.
type IElement_symbolsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElement_symbolsContext differentiates from other interfaces.
	IsElement_symbolsContext()
}

type Element_symbolsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElement_symbolsContext() *Element_symbolsContext {
	var p = new(Element_symbolsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_element_symbols
	return p
}

func (*Element_symbolsContext) IsElement_symbolsContext() {}

func NewElement_symbolsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Element_symbolsContext {
	var p = new(Element_symbolsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_element_symbols

	return p
}

func (s *Element_symbolsContext) GetParser() antlr.Parser { return s.parser }

func (s *Element_symbolsContext) UH() antlr.TerminalNode {
	return s.GetToken(smilesParserUH, 0)
}

func (s *Element_symbolsContext) LE() antlr.TerminalNode {
	return s.GetToken(smilesParserLE, 0)
}

func (s *Element_symbolsContext) UL() antlr.TerminalNode {
	return s.GetToken(smilesParserUL, 0)
}

func (s *Element_symbolsContext) LI() antlr.TerminalNode {
	return s.GetToken(smilesParserLI, 0)
}

func (s *Element_symbolsContext) UB() antlr.TerminalNode {
	return s.GetToken(smilesParserUB, 0)
}

func (s *Element_symbolsContext) UC() antlr.TerminalNode {
	return s.GetToken(smilesParserUC, 0)
}

func (s *Element_symbolsContext) UN() antlr.TerminalNode {
	return s.GetToken(smilesParserUN, 0)
}

func (s *Element_symbolsContext) UO() antlr.TerminalNode {
	return s.GetToken(smilesParserUO, 0)
}

func (s *Element_symbolsContext) UF() antlr.TerminalNode {
	return s.GetToken(smilesParserUF, 0)
}

func (s *Element_symbolsContext) LA() antlr.TerminalNode {
	return s.GetToken(smilesParserLA, 0)
}

func (s *Element_symbolsContext) UM() antlr.TerminalNode {
	return s.GetToken(smilesParserUM, 0)
}

func (s *Element_symbolsContext) LG() antlr.TerminalNode {
	return s.GetToken(smilesParserLG, 0)
}

func (s *Element_symbolsContext) UA() antlr.TerminalNode {
	return s.GetToken(smilesParserUA, 0)
}

func (s *Element_symbolsContext) LL() antlr.TerminalNode {
	return s.GetToken(smilesParserLL, 0)
}

func (s *Element_symbolsContext) US() antlr.TerminalNode {
	return s.GetToken(smilesParserUS, 0)
}

func (s *Element_symbolsContext) UP() antlr.TerminalNode {
	return s.GetToken(smilesParserUP, 0)
}

func (s *Element_symbolsContext) LR() antlr.TerminalNode {
	return s.GetToken(smilesParserLR, 0)
}

func (s *Element_symbolsContext) UK() antlr.TerminalNode {
	return s.GetToken(smilesParserUK, 0)
}

func (s *Element_symbolsContext) LC() antlr.TerminalNode {
	return s.GetToken(smilesParserLC, 0)
}

func (s *Element_symbolsContext) UT() antlr.TerminalNode {
	return s.GetToken(smilesParserUT, 0)
}

func (s *Element_symbolsContext) UV() antlr.TerminalNode {
	return s.GetToken(smilesParserUV, 0)
}

func (s *Element_symbolsContext) LN() antlr.TerminalNode {
	return s.GetToken(smilesParserLN, 0)
}

func (s *Element_symbolsContext) LO() antlr.TerminalNode {
	return s.GetToken(smilesParserLO, 0)
}

func (s *Element_symbolsContext) LU() antlr.TerminalNode {
	return s.GetToken(smilesParserLU, 0)
}

func (s *Element_symbolsContext) UZ() antlr.TerminalNode {
	return s.GetToken(smilesParserUZ, 0)
}

func (s *Element_symbolsContext) UG() antlr.TerminalNode {
	return s.GetToken(smilesParserUG, 0)
}

func (s *Element_symbolsContext) LS() antlr.TerminalNode {
	return s.GetToken(smilesParserLS, 0)
}

func (s *Element_symbolsContext) UR() antlr.TerminalNode {
	return s.GetToken(smilesParserUR, 0)
}

func (s *Element_symbolsContext) LB() antlr.TerminalNode {
	return s.GetToken(smilesParserLB, 0)
}

func (s *Element_symbolsContext) UY() antlr.TerminalNode {
	return s.GetToken(smilesParserUY, 0)
}

func (s *Element_symbolsContext) LH() antlr.TerminalNode {
	return s.GetToken(smilesParserLH, 0)
}

func (s *Element_symbolsContext) LD() antlr.TerminalNode {
	return s.GetToken(smilesParserLD, 0)
}

func (s *Element_symbolsContext) UI() antlr.TerminalNode {
	return s.GetToken(smilesParserUI, 0)
}

func (s *Element_symbolsContext) UX() antlr.TerminalNode {
	return s.GetToken(smilesParserUX, 0)
}

func (s *Element_symbolsContext) LF() antlr.TerminalNode {
	return s.GetToken(smilesParserLF, 0)
}

func (s *Element_symbolsContext) UW() antlr.TerminalNode {
	return s.GetToken(smilesParserUW, 0)
}

func (s *Element_symbolsContext) LT() antlr.TerminalNode {
	return s.GetToken(smilesParserLT, 0)
}

func (s *Element_symbolsContext) UD() antlr.TerminalNode {
	return s.GetToken(smilesParserUD, 0)
}

func (s *Element_symbolsContext) LV() antlr.TerminalNode {
	return s.GetToken(smilesParserLV, 0)
}

func (s *Element_symbolsContext) LM() antlr.TerminalNode {
	return s.GetToken(smilesParserLM, 0)
}

func (s *Element_symbolsContext) UE() antlr.TerminalNode {
	return s.GetToken(smilesParserUE, 0)
}

func (s *Element_symbolsContext) LY() antlr.TerminalNode {
	return s.GetToken(smilesParserLY, 0)
}

func (s *Element_symbolsContext) UU() antlr.TerminalNode {
	return s.GetToken(smilesParserUU, 0)
}

func (s *Element_symbolsContext) LP() antlr.TerminalNode {
	return s.GetToken(smilesParserLP, 0)
}

func (s *Element_symbolsContext) LK() antlr.TerminalNode {
	return s.GetToken(smilesParserLK, 0)
}

func (s *Element_symbolsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Element_symbolsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Element_symbolsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterElement_symbols(s)
	}
}

func (s *Element_symbolsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitElement_symbols(s)
	}
}

func (p *smilesParser) Element_symbols() (localctx IElement_symbolsContext) {
	this := p
	_ = this

	localctx = NewElement_symbolsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, smilesParserRULE_element_symbols)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(93)
			p.Match(smilesParserUH)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)
			p.Match(smilesParserUH)
		}
		{
			p.SetState(95)
			p.Match(smilesParserLE)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.Match(smilesParserUL)
		}
		{
			p.SetState(97)
			p.Match(smilesParserLI)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)
			p.Match(smilesParserUB)
		}
		{
			p.SetState(99)
			p.Match(smilesParserLE)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(100)
			p.Match(smilesParserUB)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(101)
			p.Match(smilesParserUC)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(102)
			p.Match(smilesParserUN)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(103)
			p.Match(smilesParserUO)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(104)
			p.Match(smilesParserUF)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(105)
			p.Match(smilesParserUN)
		}
		{
			p.SetState(106)
			p.Match(smilesParserLE)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(107)
			p.Match(smilesParserUN)
		}
		{
			p.SetState(108)
			p.Match(smilesParserLA)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(109)
			p.Match(smilesParserUM)
		}
		{
			p.SetState(110)
			p.Match(smilesParserLG)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(111)
			p.Match(smilesParserUA)
		}
		{
			p.SetState(112)
			p.Match(smilesParserLL)
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(113)
			p.Match(smilesParserUS)
		}
		{
			p.SetState(114)
			p.Match(smilesParserLI)
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(115)
			p.Match(smilesParserUP)
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(116)
			p.Match(smilesParserUS)
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(117)
			p.Match(smilesParserUC)
		}
		{
			p.SetState(118)
			p.Match(smilesParserLL)
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(119)
			p.Match(smilesParserUA)
		}
		{
			p.SetState(120)
			p.Match(smilesParserLR)
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(121)
			p.Match(smilesParserUK)
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(122)
			p.Match(smilesParserUC)
		}
		{
			p.SetState(123)
			p.Match(smilesParserLA)
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(124)
			p.Match(smilesParserUS)
		}
		{
			p.SetState(125)
			p.Match(smilesParserLC)
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(126)
			p.Match(smilesParserUT)
		}
		{
			p.SetState(127)
			p.Match(smilesParserLI)
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(128)
			p.Match(smilesParserUV)
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(129)
			p.Match(smilesParserUC)
		}
		{
			p.SetState(130)
			p.Match(smilesParserLR)
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(131)
			p.Match(smilesParserUM)
		}
		{
			p.SetState(132)
			p.Match(smilesParserLN)
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(133)
			p.Match(smilesParserUF)
		}
		{
			p.SetState(134)
			p.Match(smilesParserLE)
		}

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(135)
			p.Match(smilesParserUC)
		}
		{
			p.SetState(136)
			p.Match(smilesParserLO)
		}

	case 28:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(137)
			p.Match(smilesParserUN)
		}
		{
			p.SetState(138)
			p.Match(smilesParserLI)
		}

	case 29:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(139)
			p.Match(smilesParserUC)
		}
		{
			p.SetState(140)
			p.Match(smilesParserLU)
		}

	case 30:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(141)
			p.Match(smilesParserUZ)
		}
		{
			p.SetState(142)
			p.Match(smilesParserLN)
		}

	case 31:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(143)
			p.Match(smilesParserUG)
		}
		{
			p.SetState(144)
			p.Match(smilesParserLA)
		}

	case 32:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(145)
			p.Match(smilesParserUG)
		}
		{
			p.SetState(146)
			p.Match(smilesParserLE)
		}

	case 33:
		p.EnterOuterAlt(localctx, 33)
		{
			p.SetState(147)
			p.Match(smilesParserUA)
		}
		{
			p.SetState(148)
			p.Match(smilesParserLS)
		}

	case 34:
		p.EnterOuterAlt(localctx, 34)
		{
			p.SetState(149)
			p.Match(smilesParserUS)
		}
		{
			p.SetState(150)
			p.Match(smilesParserLE)
		}

	case 35:
		p.EnterOuterAlt(localctx, 35)
		{
			p.SetState(151)
			p.Match(smilesParserUB)
		}
		{
			p.SetState(152)
			p.Match(smilesParserLR)
		}

	case 36:
		p.EnterOuterAlt(localctx, 36)
		{
			p.SetState(153)
			p.Match(smilesParserUK)
		}
		{
			p.SetState(154)
			p.Match(smilesParserLR)
		}

	case 37:
		p.EnterOuterAlt(localctx, 37)
		{
			p.SetState(155)
			p.Match(smilesParserUR)
		}
		{
			p.SetState(156)
			p.Match(smilesParserLB)
		}

	case 38:
		p.EnterOuterAlt(localctx, 38)
		{
			p.SetState(157)
			p.Match(smilesParserUS)
		}
		{
			p.SetState(158)
			p.Match(smilesParserLR)
		}

	case 39:
		p.EnterOuterAlt(localctx, 39)
		{
			p.SetState(159)
			p.Match(smilesParserUY)
		}

	case 40:
		p.EnterOuterAlt(localctx, 40)
		{
			p.SetState(160)
			p.Match(smilesParserUZ)
		}
		{
			p.SetState(161)
			p.Match(smilesParserLR)
		}

	case 41:
		p.EnterOuterAlt(localctx, 41)
		{
			p.SetState(162)
			p.Match(smilesParserUN)
		}
		{
			p.SetState(163)
			p.Match(smilesParserLB)
		}

	case 42:
		p.EnterOuterAlt(localctx, 42)
		{
			p.SetState(164)
			p.Match(smilesParserUM)
		}
		{
			p.SetState(165)
			p.Match(smilesParserLO)
		}

	case 43:
		p.EnterOuterAlt(localctx, 43)
		{
			p.SetState(166)
			p.Match(smilesParserUT)
		}
		{
			p.SetState(167)
			p.Match(smilesParserLC)
		}

	case 44:
		p.EnterOuterAlt(localctx, 44)
		{
			p.SetState(168)
			p.Match(smilesParserUR)
		}
		{
			p.SetState(169)
			p.Match(smilesParserLU)
		}

	case 45:
		p.EnterOuterAlt(localctx, 45)
		{
			p.SetState(170)
			p.Match(smilesParserUR)
		}
		{
			p.SetState(171)
			p.Match(smilesParserLH)
		}

	case 46:
		p.EnterOuterAlt(localctx, 46)
		{
			p.SetState(172)
			p.Match(smilesParserUP)
		}
		{
			p.SetState(173)
			p.Match(smilesParserLD)
		}

	case 47:
		p.EnterOuterAlt(localctx, 47)
		{
			p.SetState(174)
			p.Match(smilesParserUA)
		}
		{
			p.SetState(175)
			p.Match(smilesParserLG)
		}

	case 48:
		p.EnterOuterAlt(localctx, 48)
		{
			p.SetState(176)
			p.Match(smilesParserUC)
		}
		{
			p.SetState(177)
			p.Match(smilesParserLD)
		}

	case 49:
		p.EnterOuterAlt(localctx, 49)
		{
			p.SetState(178)
			p.Match(smilesParserUI)
		}
		{
			p.SetState(179)
			p.Match(smilesParserLN)
		}

	case 50:
		p.EnterOuterAlt(localctx, 50)
		{
			p.SetState(180)
			p.Match(smilesParserUS)
		}
		{
			p.SetState(181)
			p.Match(smilesParserLN)
		}

	case 51:
		p.EnterOuterAlt(localctx, 51)
		{
			p.SetState(182)
			p.Match(smilesParserUS)
		}
		{
			p.SetState(183)
			p.Match(smilesParserLB)
		}

	case 52:
		p.EnterOuterAlt(localctx, 52)
		{
			p.SetState(184)
			p.Match(smilesParserUT)
		}
		{
			p.SetState(185)
			p.Match(smilesParserLE)
		}

	case 53:
		p.EnterOuterAlt(localctx, 53)
		{
			p.SetState(186)
			p.Match(smilesParserUI)
		}

	case 54:
		p.EnterOuterAlt(localctx, 54)
		{
			p.SetState(187)
			p.Match(smilesParserUX)
		}
		{
			p.SetState(188)
			p.Match(smilesParserLE)
		}

	case 55:
		p.EnterOuterAlt(localctx, 55)
		{
			p.SetState(189)
			p.Match(smilesParserUC)
		}
		{
			p.SetState(190)
			p.Match(smilesParserLS)
		}

	case 56:
		p.EnterOuterAlt(localctx, 56)
		{
			p.SetState(191)
			p.Match(smilesParserUB)
		}
		{
			p.SetState(192)
			p.Match(smilesParserLA)
		}

	case 57:
		p.EnterOuterAlt(localctx, 57)
		{
			p.SetState(193)
			p.Match(smilesParserUH)
		}
		{
			p.SetState(194)
			p.Match(smilesParserLF)
		}

	case 58:
		p.EnterOuterAlt(localctx, 58)
		{
			p.SetState(195)
			p.Match(smilesParserUT)
		}
		{
			p.SetState(196)
			p.Match(smilesParserLA)
		}

	case 59:
		p.EnterOuterAlt(localctx, 59)
		{
			p.SetState(197)
			p.Match(smilesParserUW)
		}

	case 60:
		p.EnterOuterAlt(localctx, 60)
		{
			p.SetState(198)
			p.Match(smilesParserUR)
		}
		{
			p.SetState(199)
			p.Match(smilesParserLE)
		}

	case 61:
		p.EnterOuterAlt(localctx, 61)
		{
			p.SetState(200)
			p.Match(smilesParserUO)
		}
		{
			p.SetState(201)
			p.Match(smilesParserLS)
		}

	case 62:
		p.EnterOuterAlt(localctx, 62)
		{
			p.SetState(202)
			p.Match(smilesParserUI)
		}
		{
			p.SetState(203)
			p.Match(smilesParserLR)
		}

	case 63:
		p.EnterOuterAlt(localctx, 63)
		{
			p.SetState(204)
			p.Match(smilesParserUP)
		}
		{
			p.SetState(205)
			p.Match(smilesParserLT)
		}

	case 64:
		p.EnterOuterAlt(localctx, 64)
		{
			p.SetState(206)
			p.Match(smilesParserUA)
		}
		{
			p.SetState(207)
			p.Match(smilesParserLU)
		}

	case 65:
		p.EnterOuterAlt(localctx, 65)
		{
			p.SetState(208)
			p.Match(smilesParserUH)
		}
		{
			p.SetState(209)
			p.Match(smilesParserLG)
		}

	case 66:
		p.EnterOuterAlt(localctx, 66)
		{
			p.SetState(210)
			p.Match(smilesParserUT)
		}
		{
			p.SetState(211)
			p.Match(smilesParserLL)
		}

	case 67:
		p.EnterOuterAlt(localctx, 67)
		{
			p.SetState(212)
			p.Match(smilesParserUP)
		}
		{
			p.SetState(213)
			p.Match(smilesParserLB)
		}

	case 68:
		p.EnterOuterAlt(localctx, 68)
		{
			p.SetState(214)
			p.Match(smilesParserUB)
		}
		{
			p.SetState(215)
			p.Match(smilesParserLI)
		}

	case 69:
		p.EnterOuterAlt(localctx, 69)
		{
			p.SetState(216)
			p.Match(smilesParserUP)
		}
		{
			p.SetState(217)
			p.Match(smilesParserLO)
		}

	case 70:
		p.EnterOuterAlt(localctx, 70)
		{
			p.SetState(218)
			p.Match(smilesParserUA)
		}
		{
			p.SetState(219)
			p.Match(smilesParserLT)
		}

	case 71:
		p.EnterOuterAlt(localctx, 71)
		{
			p.SetState(220)
			p.Match(smilesParserUR)
		}
		{
			p.SetState(221)
			p.Match(smilesParserLN)
		}

	case 72:
		p.EnterOuterAlt(localctx, 72)
		{
			p.SetState(222)
			p.Match(smilesParserUF)
		}
		{
			p.SetState(223)
			p.Match(smilesParserLR)
		}

	case 73:
		p.EnterOuterAlt(localctx, 73)
		{
			p.SetState(224)
			p.Match(smilesParserUR)
		}
		{
			p.SetState(225)
			p.Match(smilesParserLA)
		}

	case 74:
		p.EnterOuterAlt(localctx, 74)
		{
			p.SetState(226)
			p.Match(smilesParserUR)
		}
		{
			p.SetState(227)
			p.Match(smilesParserLF)
		}

	case 75:
		p.EnterOuterAlt(localctx, 75)
		{
			p.SetState(228)
			p.Match(smilesParserUD)
		}
		{
			p.SetState(229)
			p.Match(smilesParserLB)
		}

	case 76:
		p.EnterOuterAlt(localctx, 76)
		{
			p.SetState(230)
			p.Match(smilesParserUS)
		}
		{
			p.SetState(231)
			p.Match(smilesParserLG)
		}

	case 77:
		p.EnterOuterAlt(localctx, 77)
		{
			p.SetState(232)
			p.Match(smilesParserUB)
		}
		{
			p.SetState(233)
			p.Match(smilesParserLH)
		}

	case 78:
		p.EnterOuterAlt(localctx, 78)
		{
			p.SetState(234)
			p.Match(smilesParserUH)
		}
		{
			p.SetState(235)
			p.Match(smilesParserLS)
		}

	case 79:
		p.EnterOuterAlt(localctx, 79)
		{
			p.SetState(236)
			p.Match(smilesParserUM)
		}
		{
			p.SetState(237)
			p.Match(smilesParserLT)
		}

	case 80:
		p.EnterOuterAlt(localctx, 80)
		{
			p.SetState(238)
			p.Match(smilesParserUD)
		}
		{
			p.SetState(239)
			p.Match(smilesParserLS)
		}

	case 81:
		p.EnterOuterAlt(localctx, 81)
		{
			p.SetState(240)
			p.Match(smilesParserUR)
		}
		{
			p.SetState(241)
			p.Match(smilesParserLG)
		}

	case 82:
		p.EnterOuterAlt(localctx, 82)
		{
			p.SetState(242)
			p.Match(smilesParserUC)
		}
		{
			p.SetState(243)
			p.Match(smilesParserLN)
		}

	case 83:
		p.EnterOuterAlt(localctx, 83)
		{
			p.SetState(244)
			p.Match(smilesParserUF)
		}
		{
			p.SetState(245)
			p.Match(smilesParserLL)
		}

	case 84:
		p.EnterOuterAlt(localctx, 84)
		{
			p.SetState(246)
			p.Match(smilesParserUL)
		}
		{
			p.SetState(247)
			p.Match(smilesParserLV)
		}

	case 85:
		p.EnterOuterAlt(localctx, 85)
		{
			p.SetState(248)
			p.Match(smilesParserUL)
		}
		{
			p.SetState(249)
			p.Match(smilesParserLA)
		}

	case 86:
		p.EnterOuterAlt(localctx, 86)
		{
			p.SetState(250)
			p.Match(smilesParserUC)
		}
		{
			p.SetState(251)
			p.Match(smilesParserLE)
		}

	case 87:
		p.EnterOuterAlt(localctx, 87)
		{
			p.SetState(252)
			p.Match(smilesParserUP)
		}
		{
			p.SetState(253)
			p.Match(smilesParserLR)
		}

	case 88:
		p.EnterOuterAlt(localctx, 88)
		{
			p.SetState(254)
			p.Match(smilesParserUN)
		}
		{
			p.SetState(255)
			p.Match(smilesParserLD)
		}

	case 89:
		p.EnterOuterAlt(localctx, 89)
		{
			p.SetState(256)
			p.Match(smilesParserUP)
		}
		{
			p.SetState(257)
			p.Match(smilesParserLM)
		}

	case 90:
		p.EnterOuterAlt(localctx, 90)
		{
			p.SetState(258)
			p.Match(smilesParserUS)
		}
		{
			p.SetState(259)
			p.Match(smilesParserLM)
		}

	case 91:
		p.EnterOuterAlt(localctx, 91)
		{
			p.SetState(260)
			p.Match(smilesParserUE)
		}
		{
			p.SetState(261)
			p.Match(smilesParserLU)
		}

	case 92:
		p.EnterOuterAlt(localctx, 92)
		{
			p.SetState(262)
			p.Match(smilesParserUG)
		}
		{
			p.SetState(263)
			p.Match(smilesParserLD)
		}

	case 93:
		p.EnterOuterAlt(localctx, 93)
		{
			p.SetState(264)
			p.Match(smilesParserUT)
		}
		{
			p.SetState(265)
			p.Match(smilesParserLB)
		}

	case 94:
		p.EnterOuterAlt(localctx, 94)
		{
			p.SetState(266)
			p.Match(smilesParserUD)
		}
		{
			p.SetState(267)
			p.Match(smilesParserLY)
		}

	case 95:
		p.EnterOuterAlt(localctx, 95)
		{
			p.SetState(268)
			p.Match(smilesParserUH)
		}
		{
			p.SetState(269)
			p.Match(smilesParserLO)
		}

	case 96:
		p.EnterOuterAlt(localctx, 96)
		{
			p.SetState(270)
			p.Match(smilesParserUE)
		}
		{
			p.SetState(271)
			p.Match(smilesParserLR)
		}

	case 97:
		p.EnterOuterAlt(localctx, 97)
		{
			p.SetState(272)
			p.Match(smilesParserUT)
		}
		{
			p.SetState(273)
			p.Match(smilesParserLM)
		}

	case 98:
		p.EnterOuterAlt(localctx, 98)
		{
			p.SetState(274)
			p.Match(smilesParserUY)
		}
		{
			p.SetState(275)
			p.Match(smilesParserLB)
		}

	case 99:
		p.EnterOuterAlt(localctx, 99)
		{
			p.SetState(276)
			p.Match(smilesParserUL)
		}
		{
			p.SetState(277)
			p.Match(smilesParserLU)
		}

	case 100:
		p.EnterOuterAlt(localctx, 100)
		{
			p.SetState(278)
			p.Match(smilesParserUA)
		}
		{
			p.SetState(279)
			p.Match(smilesParserLC)
		}

	case 101:
		p.EnterOuterAlt(localctx, 101)
		{
			p.SetState(280)
			p.Match(smilesParserUT)
		}
		{
			p.SetState(281)
			p.Match(smilesParserLH)
		}

	case 102:
		p.EnterOuterAlt(localctx, 102)
		{
			p.SetState(282)
			p.Match(smilesParserUP)
		}
		{
			p.SetState(283)
			p.Match(smilesParserLA)
		}

	case 103:
		p.EnterOuterAlt(localctx, 103)
		{
			p.SetState(284)
			p.Match(smilesParserUU)
		}

	case 104:
		p.EnterOuterAlt(localctx, 104)
		{
			p.SetState(285)
			p.Match(smilesParserUN)
		}
		{
			p.SetState(286)
			p.Match(smilesParserLP)
		}

	case 105:
		p.EnterOuterAlt(localctx, 105)
		{
			p.SetState(287)
			p.Match(smilesParserUP)
		}
		{
			p.SetState(288)
			p.Match(smilesParserLU)
		}

	case 106:
		p.EnterOuterAlt(localctx, 106)
		{
			p.SetState(289)
			p.Match(smilesParserUA)
		}
		{
			p.SetState(290)
			p.Match(smilesParserLM)
		}

	case 107:
		p.EnterOuterAlt(localctx, 107)
		{
			p.SetState(291)
			p.Match(smilesParserUC)
		}
		{
			p.SetState(292)
			p.Match(smilesParserLM)
		}

	case 108:
		p.EnterOuterAlt(localctx, 108)
		{
			p.SetState(293)
			p.Match(smilesParserUB)
		}
		{
			p.SetState(294)
			p.Match(smilesParserLK)
		}

	case 109:
		p.EnterOuterAlt(localctx, 109)
		{
			p.SetState(295)
			p.Match(smilesParserUC)
		}
		{
			p.SetState(296)
			p.Match(smilesParserLF)
		}

	case 110:
		p.EnterOuterAlt(localctx, 110)
		{
			p.SetState(297)
			p.Match(smilesParserUE)
		}
		{
			p.SetState(298)
			p.Match(smilesParserLS)
		}

	case 111:
		p.EnterOuterAlt(localctx, 111)
		{
			p.SetState(299)
			p.Match(smilesParserUF)
		}
		{
			p.SetState(300)
			p.Match(smilesParserLM)
		}

	case 112:
		p.EnterOuterAlt(localctx, 112)
		{
			p.SetState(301)
			p.Match(smilesParserUM)
		}
		{
			p.SetState(302)
			p.Match(smilesParserLD)
		}

	case 113:
		p.EnterOuterAlt(localctx, 113)
		{
			p.SetState(303)
			p.Match(smilesParserUN)
		}
		{
			p.SetState(304)
			p.Match(smilesParserLO)
		}

	case 114:
		p.EnterOuterAlt(localctx, 114)
		{
			p.SetState(305)
			p.Match(smilesParserUL)
		}
		{
			p.SetState(306)
			p.Match(smilesParserLR)
		}

	}

	return localctx
}

// IAromatic_symbolContext is an interface to support dynamic dispatch.
type IAromatic_symbolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAromatic_symbolContext differentiates from other interfaces.
	IsAromatic_symbolContext()
}

type Aromatic_symbolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAromatic_symbolContext() *Aromatic_symbolContext {
	var p = new(Aromatic_symbolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_aromatic_symbol
	return p
}

func (*Aromatic_symbolContext) IsAromatic_symbolContext() {}

func NewAromatic_symbolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Aromatic_symbolContext {
	var p = new(Aromatic_symbolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_aromatic_symbol

	return p
}

func (s *Aromatic_symbolContext) GetParser() antlr.Parser { return s.parser }

func (s *Aromatic_symbolContext) LC() antlr.TerminalNode {
	return s.GetToken(smilesParserLC, 0)
}

func (s *Aromatic_symbolContext) LN() antlr.TerminalNode {
	return s.GetToken(smilesParserLN, 0)
}

func (s *Aromatic_symbolContext) LO() antlr.TerminalNode {
	return s.GetToken(smilesParserLO, 0)
}

func (s *Aromatic_symbolContext) LP() antlr.TerminalNode {
	return s.GetToken(smilesParserLP, 0)
}

func (s *Aromatic_symbolContext) LS() antlr.TerminalNode {
	return s.GetToken(smilesParserLS, 0)
}

func (s *Aromatic_symbolContext) LE() antlr.TerminalNode {
	return s.GetToken(smilesParserLE, 0)
}

func (s *Aromatic_symbolContext) LA() antlr.TerminalNode {
	return s.GetToken(smilesParserLA, 0)
}

func (s *Aromatic_symbolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Aromatic_symbolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Aromatic_symbolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterAromatic_symbol(s)
	}
}

func (s *Aromatic_symbolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitAromatic_symbol(s)
	}
}

func (p *smilesParser) Aromatic_symbol() (localctx IAromatic_symbolContext) {
	this := p
	_ = this

	localctx = NewAromatic_symbolContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, smilesParserRULE_aromatic_symbol)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.Match(smilesParserLC)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(310)
			p.Match(smilesParserLN)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(311)
			p.Match(smilesParserLO)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(312)
			p.Match(smilesParserLP)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(313)
			p.Match(smilesParserLS)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(314)
			p.Match(smilesParserLS)
		}
		{
			p.SetState(315)
			p.Match(smilesParserLE)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(316)
			p.Match(smilesParserLA)
		}
		{
			p.SetState(317)
			p.Match(smilesParserLS)
		}

	}

	return localctx
}

// IChiralContext is an interface to support dynamic dispatch.
type IChiralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChiralContext differentiates from other interfaces.
	IsChiralContext()
}

type ChiralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChiralContext() *ChiralContext {
	var p = new(ChiralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_chiral
	return p
}

func (*ChiralContext) IsChiralContext() {}

func NewChiralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChiralContext {
	var p = new(ChiralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_chiral

	return p
}

func (s *ChiralContext) GetParser() antlr.Parser { return s.parser }

func (s *ChiralContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(smilesParserDIGIT)
}

func (s *ChiralContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(smilesParserDIGIT, i)
}

func (s *ChiralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChiralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChiralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterChiral(s)
	}
}

func (s *ChiralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitChiral(s)
	}
}

func (p *smilesParser) Chiral() (localctx IChiralContext) {
	this := p
	_ = this

	localctx = NewChiralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, smilesParserRULE_chiral)
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

	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(320)
			p.Match(smilesParserT__3)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(321)
			p.Match(smilesParserT__4)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(322)
			p.Match(smilesParserT__5)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(323)
			p.Match(smilesParserT__6)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(324)
			p.Match(smilesParserT__7)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(325)
			p.Match(smilesParserT__8)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(326)
			p.Match(smilesParserT__9)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(327)
			p.Match(smilesParserT__10)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(328)
			p.Match(smilesParserT__11)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(329)
			p.Match(smilesParserT__12)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(330)
			p.Match(smilesParserT__13)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(331)
			p.Match(smilesParserT__14)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(332)
			p.Match(smilesParserT__14)
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(333)
			p.Match(smilesParserT__15)
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(334)
			p.Match(smilesParserT__16)
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(335)
			p.Match(smilesParserT__17)
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(336)
			p.Match(smilesParserT__18)
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(337)
			p.Match(smilesParserT__19)
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(338)
			p.Match(smilesParserT__20)
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(339)
			p.Match(smilesParserT__21)
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(340)
			p.Match(smilesParserT__22)
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(341)
			p.Match(smilesParserT__23)
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(342)
			p.Match(smilesParserT__24)
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(343)
			p.Match(smilesParserT__25)
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(344)
			p.Match(smilesParserT__26)
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(345)
			p.Match(smilesParserT__27)
		}

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(346)
			p.Match(smilesParserT__28)
		}

	case 28:
		p.EnterOuterAlt(localctx, 28)
		{
			p.SetState(347)
			p.Match(smilesParserT__29)
		}

	case 29:
		p.EnterOuterAlt(localctx, 29)
		{
			p.SetState(348)
			p.Match(smilesParserT__30)
		}

	case 30:
		p.EnterOuterAlt(localctx, 30)
		{
			p.SetState(349)
			p.Match(smilesParserT__31)
		}

	case 31:
		p.EnterOuterAlt(localctx, 31)
		{
			p.SetState(350)
			p.Match(smilesParserT__32)
		}

	case 32:
		p.EnterOuterAlt(localctx, 32)
		{
			p.SetState(351)
			p.Match(smilesParserT__33)
		}

	case 33:
		p.EnterOuterAlt(localctx, 33)
		{
			p.SetState(352)
			p.Match(smilesParserT__34)
		}

	case 34:
		p.EnterOuterAlt(localctx, 34)
		{
			p.SetState(353)
			p.Match(smilesParserT__35)
		}

	case 35:
		p.EnterOuterAlt(localctx, 35)
		{
			p.SetState(354)
			p.Match(smilesParserT__36)
		}

	case 36:
		p.EnterOuterAlt(localctx, 36)
		{
			p.SetState(355)
			p.Match(smilesParserT__37)
		}

	case 37:
		p.EnterOuterAlt(localctx, 37)
		{
			p.SetState(356)
			p.Match(smilesParserT__38)
		}

	case 38:
		p.EnterOuterAlt(localctx, 38)
		{
			p.SetState(357)
			p.Match(smilesParserT__39)
		}

	case 39:
		p.EnterOuterAlt(localctx, 39)
		{
			p.SetState(358)
			p.Match(smilesParserT__40)
		}

	case 40:
		p.EnterOuterAlt(localctx, 40)
		{
			p.SetState(359)
			p.Match(smilesParserT__41)
		}

	case 41:
		p.EnterOuterAlt(localctx, 41)
		{
			p.SetState(360)
			p.Match(smilesParserT__42)
		}

	case 42:
		p.EnterOuterAlt(localctx, 42)
		{
			p.SetState(361)
			p.Match(smilesParserT__43)
		}

	case 43:
		p.EnterOuterAlt(localctx, 43)
		{
			p.SetState(362)
			p.Match(smilesParserT__44)
		}

	case 44:
		p.EnterOuterAlt(localctx, 44)
		{
			p.SetState(363)
			p.Match(smilesParserT__45)
		}

	case 45:
		p.EnterOuterAlt(localctx, 45)
		{
			p.SetState(364)
			p.Match(smilesParserT__46)
		}

	case 46:
		p.EnterOuterAlt(localctx, 46)
		{
			p.SetState(365)
			p.Match(smilesParserT__47)
		}

	case 47:
		p.EnterOuterAlt(localctx, 47)
		{
			p.SetState(366)
			p.Match(smilesParserT__48)
		}

	case 48:
		p.EnterOuterAlt(localctx, 48)
		{
			p.SetState(367)
			p.Match(smilesParserT__49)
		}

	case 49:
		p.EnterOuterAlt(localctx, 49)
		{
			p.SetState(368)
			p.Match(smilesParserT__50)
		}

	case 50:
		p.EnterOuterAlt(localctx, 50)
		{
			p.SetState(369)
			p.Match(smilesParserT__51)
		}

	case 51:
		p.EnterOuterAlt(localctx, 51)
		{
			p.SetState(370)
			p.Match(smilesParserT__52)
		}

	case 52:
		p.EnterOuterAlt(localctx, 52)
		{
			p.SetState(371)
			p.Match(smilesParserT__53)
		}

	case 53:
		p.EnterOuterAlt(localctx, 53)
		{
			p.SetState(372)
			p.Match(smilesParserT__54)
		}

	case 54:
		p.EnterOuterAlt(localctx, 54)
		{
			p.SetState(373)
			p.Match(smilesParserT__55)
		}

	case 55:
		p.EnterOuterAlt(localctx, 55)
		{
			p.SetState(374)
			p.Match(smilesParserT__56)
		}

	case 56:
		p.EnterOuterAlt(localctx, 56)
		{
			p.SetState(375)
			p.Match(smilesParserT__57)
		}

	case 57:
		p.EnterOuterAlt(localctx, 57)
		{
			p.SetState(376)
			p.Match(smilesParserT__58)
		}

	case 58:
		p.EnterOuterAlt(localctx, 58)
		{
			p.SetState(377)
			p.Match(smilesParserT__59)
		}

	case 59:
		p.EnterOuterAlt(localctx, 59)
		{
			p.SetState(378)
			p.Match(smilesParserT__60)
		}

	case 60:
		p.EnterOuterAlt(localctx, 60)
		{
			p.SetState(379)
			p.Match(smilesParserT__61)
		}

	case 61:
		p.EnterOuterAlt(localctx, 61)
		{
			p.SetState(380)
			p.Match(smilesParserT__62)
		}

	case 62:
		p.EnterOuterAlt(localctx, 62)
		{
			p.SetState(381)
			p.Match(smilesParserT__63)
		}

	case 63:
		p.EnterOuterAlt(localctx, 63)
		{
			p.SetState(382)
			p.Match(smilesParserT__64)
		}

	case 64:
		p.EnterOuterAlt(localctx, 64)
		{
			p.SetState(383)
			p.Match(smilesParserT__65)
		}

	case 65:
		p.EnterOuterAlt(localctx, 65)
		{
			p.SetState(384)
			p.Match(smilesParserT__66)
		}

	case 66:
		p.EnterOuterAlt(localctx, 66)
		{
			p.SetState(385)
			p.Match(smilesParserT__67)
		}

	case 67:
		p.EnterOuterAlt(localctx, 67)
		{
			p.SetState(386)
			p.Match(smilesParserT__68)
		}

	case 68:
		p.EnterOuterAlt(localctx, 68)
		{
			p.SetState(387)
			p.Match(smilesParserT__69)
		}

	case 69:
		p.EnterOuterAlt(localctx, 69)
		{
			p.SetState(388)
			p.Match(smilesParserT__70)
		}

	case 70:
		p.EnterOuterAlt(localctx, 70)
		{
			p.SetState(389)
			p.Match(smilesParserT__71)
		}

	case 71:
		p.EnterOuterAlt(localctx, 71)
		{
			p.SetState(390)
			_la = p.GetTokenStream().LA(1)

			if !(_la == smilesParserT__72 || _la == smilesParserT__73) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(391)
			p.Match(smilesParserDIGIT)
		}
		{
			p.SetState(392)
			p.Match(smilesParserDIGIT)
		}

	}

	return localctx
}

// IHcountContext is an interface to support dynamic dispatch.
type IHcountContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHcountContext differentiates from other interfaces.
	IsHcountContext()
}

type HcountContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHcountContext() *HcountContext {
	var p = new(HcountContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_hcount
	return p
}

func (*HcountContext) IsHcountContext() {}

func NewHcountContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HcountContext {
	var p = new(HcountContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_hcount

	return p
}

func (s *HcountContext) GetParser() antlr.Parser { return s.parser }

func (s *HcountContext) UH() antlr.TerminalNode {
	return s.GetToken(smilesParserUH, 0)
}

func (s *HcountContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(smilesParserDIGIT, 0)
}

func (s *HcountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HcountContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HcountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterHcount(s)
	}
}

func (s *HcountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitHcount(s)
	}
}

func (p *smilesParser) Hcount() (localctx IHcountContext) {
	this := p
	_ = this

	localctx = NewHcountContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, smilesParserRULE_hcount)
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
		p.SetState(395)
		p.Match(smilesParserUH)
	}
	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == smilesParserDIGIT {
		{
			p.SetState(396)
			p.Match(smilesParserDIGIT)
		}

	}

	return localctx
}

// IChargeContext is an interface to support dynamic dispatch.
type IChargeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChargeContext differentiates from other interfaces.
	IsChargeContext()
}

type ChargeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChargeContext() *ChargeContext {
	var p = new(ChargeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_charge
	return p
}

func (*ChargeContext) IsChargeContext() {}

func NewChargeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChargeContext {
	var p = new(ChargeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_charge

	return p
}

func (s *ChargeContext) GetParser() antlr.Parser { return s.parser }

func (s *ChargeContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(smilesParserDIGIT)
}

func (s *ChargeContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(smilesParserDIGIT, i)
}

func (s *ChargeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChargeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChargeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterCharge(s)
	}
}

func (s *ChargeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitCharge(s)
	}
}

func (p *smilesParser) Charge() (localctx IChargeContext) {
	this := p
	_ = this

	localctx = NewChargeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, smilesParserRULE_charge)
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

	p.SetState(408)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case smilesParserT__74, smilesParserT__75:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(399)
			_la = p.GetTokenStream().LA(1)

			if !(_la == smilesParserT__74 || _la == smilesParserT__75) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(401)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(400)
				p.Match(smilesParserDIGIT)
			}

		}
		p.SetState(404)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == smilesParserDIGIT {
			{
				p.SetState(403)
				p.Match(smilesParserDIGIT)
			}

		}

	case smilesParserT__76:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(406)
			p.Match(smilesParserT__76)
		}

	case smilesParserT__77:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(407)
			p.Match(smilesParserT__77)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IClass_Context is an interface to support dynamic dispatch.
type IClass_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClass_Context differentiates from other interfaces.
	IsClass_Context()
}

type Class_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClass_Context() *Class_Context {
	var p = new(Class_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_class_
	return p
}

func (*Class_Context) IsClass_Context() {}

func NewClass_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Class_Context {
	var p = new(Class_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_class_

	return p
}

func (s *Class_Context) GetParser() antlr.Parser { return s.parser }

func (s *Class_Context) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(smilesParserDIGIT)
}

func (s *Class_Context) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(smilesParserDIGIT, i)
}

func (s *Class_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Class_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Class_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterClass_(s)
	}
}

func (s *Class_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitClass_(s)
	}
}

func (p *smilesParser) Class_() (localctx IClass_Context) {
	this := p
	_ = this

	localctx = NewClass_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, smilesParserRULE_class_)
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
		p.Match(smilesParserT__78)
	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == smilesParserDIGIT {
		{
			p.SetState(411)
			p.Match(smilesParserDIGIT)
		}

		p.SetState(414)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBondContext is an interface to support dynamic dispatch.
type IBondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBondContext differentiates from other interfaces.
	IsBondContext()
}

type BondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBondContext() *BondContext {
	var p = new(BondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_bond
	return p
}

func (*BondContext) IsBondContext() {}

func NewBondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BondContext {
	var p = new(BondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_bond

	return p
}

func (s *BondContext) GetParser() antlr.Parser { return s.parser }
func (s *BondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterBond(s)
	}
}

func (s *BondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitBond(s)
	}
}

func (p *smilesParser) Bond() (localctx IBondContext) {
	this := p
	_ = this

	localctx = NewBondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, smilesParserRULE_bond)
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
		p.SetState(416)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(smilesParserT__75-76))|(1<<(smilesParserT__78-76))|(1<<(smilesParserT__79-76))|(1<<(smilesParserT__80-76))|(1<<(smilesParserT__81-76))|(1<<(smilesParserT__82-76))|(1<<(smilesParserT__83-76)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRingbondContext is an interface to support dynamic dispatch.
type IRingbondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRingbondContext differentiates from other interfaces.
	IsRingbondContext()
}

type RingbondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRingbondContext() *RingbondContext {
	var p = new(RingbondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_ringbond
	return p
}

func (*RingbondContext) IsRingbondContext() {}

func NewRingbondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RingbondContext {
	var p = new(RingbondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_ringbond

	return p
}

func (s *RingbondContext) GetParser() antlr.Parser { return s.parser }

func (s *RingbondContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(smilesParserDIGIT)
}

func (s *RingbondContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(smilesParserDIGIT, i)
}

func (s *RingbondContext) Bond() IBondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBondContext)
}

func (s *RingbondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RingbondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RingbondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterRingbond(s)
	}
}

func (s *RingbondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitRingbond(s)
	}
}

func (p *smilesParser) Ringbond() (localctx IRingbondContext) {
	this := p
	_ = this

	localctx = NewRingbondContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, smilesParserRULE_ringbond)
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
	p.SetState(419)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(smilesParserT__75-76))|(1<<(smilesParserT__78-76))|(1<<(smilesParserT__79-76))|(1<<(smilesParserT__80-76))|(1<<(smilesParserT__81-76))|(1<<(smilesParserT__82-76))|(1<<(smilesParserT__83-76)))) != 0 {
		{
			p.SetState(418)
			p.Bond()
		}

	}
	p.SetState(425)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case smilesParserDIGIT:
		{
			p.SetState(421)
			p.Match(smilesParserDIGIT)
		}

	case smilesParserT__84:
		{
			p.SetState(422)
			p.Match(smilesParserT__84)
		}
		{
			p.SetState(423)
			p.Match(smilesParserDIGIT)
		}
		{
			p.SetState(424)
			p.Match(smilesParserDIGIT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBranched_atomContext is an interface to support dynamic dispatch.
type IBranched_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBranched_atomContext differentiates from other interfaces.
	IsBranched_atomContext()
}

type Branched_atomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranched_atomContext() *Branched_atomContext {
	var p = new(Branched_atomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_branched_atom
	return p
}

func (*Branched_atomContext) IsBranched_atomContext() {}

func NewBranched_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Branched_atomContext {
	var p = new(Branched_atomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_branched_atom

	return p
}

func (s *Branched_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Branched_atomContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *Branched_atomContext) AllRingbond() []IRingbondContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRingbondContext)(nil)).Elem())
	var tst = make([]IRingbondContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRingbondContext)
		}
	}

	return tst
}

func (s *Branched_atomContext) Ringbond(i int) IRingbondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRingbondContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRingbondContext)
}

func (s *Branched_atomContext) AllBranch() []IBranchContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBranchContext)(nil)).Elem())
	var tst = make([]IBranchContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBranchContext)
		}
	}

	return tst
}

func (s *Branched_atomContext) Branch(i int) IBranchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranchContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBranchContext)
}

func (s *Branched_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Branched_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Branched_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterBranched_atom(s)
	}
}

func (s *Branched_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitBranched_atom(s)
	}
}

func (p *smilesParser) Branched_atom() (localctx IBranched_atomContext) {
	this := p
	_ = this

	localctx = NewBranched_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, smilesParserRULE_branched_atom)
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
		p.SetState(427)
		p.Atom()
	}
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(428)
				p.Ringbond()
			}

		}
		p.SetState(433)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == smilesParserT__85 {
		{
			p.SetState(434)
			p.Branch()
		}

		p.SetState(439)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IBranchContext is an interface to support dynamic dispatch.
type IBranchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBranchContext differentiates from other interfaces.
	IsBranchContext()
}

type BranchContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBranchContext() *BranchContext {
	var p = new(BranchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_branch
	return p
}

func (*BranchContext) IsBranchContext() {}

func NewBranchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BranchContext {
	var p = new(BranchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_branch

	return p
}

func (s *BranchContext) GetParser() antlr.Parser { return s.parser }

func (s *BranchContext) Chain() IChainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChainContext)
}

func (s *BranchContext) Bond() IBondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBondContext)
}

func (s *BranchContext) DOT() antlr.TerminalNode {
	return s.GetToken(smilesParserDOT, 0)
}

func (s *BranchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BranchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BranchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterBranch(s)
	}
}

func (s *BranchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitBranch(s)
	}
}

func (p *smilesParser) Branch() (localctx IBranchContext) {
	this := p
	_ = this

	localctx = NewBranchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, smilesParserRULE_branch)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(smilesParserT__85)
	}
	p.SetState(443)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case smilesParserT__75, smilesParserT__78, smilesParserT__79, smilesParserT__80, smilesParserT__81, smilesParserT__82, smilesParserT__83:
		{
			p.SetState(441)
			p.Bond()
		}

	case smilesParserDOT:
		{
			p.SetState(442)
			p.Match(smilesParserDOT)
		}

	case smilesParserT__0, smilesParserT__1, smilesParserLB, smilesParserLC, smilesParserLN, smilesParserLO, smilesParserLP, smilesParserLS, smilesParserUB, smilesParserUC, smilesParserUF, smilesParserUI, smilesParserUN, smilesParserUO, smilesParserUP, smilesParserUS:

	default:
	}
	{
		p.SetState(445)
		p.Chain()
	}
	{
		p.SetState(446)
		p.Match(smilesParserT__86)
	}

	return localctx
}

// IChainContext is an interface to support dynamic dispatch.
type IChainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChainContext differentiates from other interfaces.
	IsChainContext()
}

type ChainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChainContext() *ChainContext {
	var p = new(ChainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_chain
	return p
}

func (*ChainContext) IsChainContext() {}

func NewChainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChainContext {
	var p = new(ChainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_chain

	return p
}

func (s *ChainContext) GetParser() antlr.Parser { return s.parser }

func (s *ChainContext) AllBranched_atom() []IBranched_atomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBranched_atomContext)(nil)).Elem())
	var tst = make([]IBranched_atomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBranched_atomContext)
		}
	}

	return tst
}

func (s *ChainContext) Branched_atom(i int) IBranched_atomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBranched_atomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBranched_atomContext)
}

func (s *ChainContext) AllBond() []IBondContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBondContext)(nil)).Elem())
	var tst = make([]IBondContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBondContext)
		}
	}

	return tst
}

func (s *ChainContext) Bond(i int) IBondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBondContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBondContext)
}

func (s *ChainContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(smilesParserDOT)
}

func (s *ChainContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(smilesParserDOT, i)
}

func (s *ChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterChain(s)
	}
}

func (s *ChainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitChain(s)
	}
}

func (p *smilesParser) Chain() (localctx IChainContext) {
	this := p
	_ = this

	localctx = NewChainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, smilesParserRULE_chain)
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
		p.Branched_atom()
	}
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == smilesParserT__0 || _la == smilesParserT__1 || (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(smilesParserT__75-76))|(1<<(smilesParserT__78-76))|(1<<(smilesParserT__79-76))|(1<<(smilesParserT__80-76))|(1<<(smilesParserT__81-76))|(1<<(smilesParserT__82-76))|(1<<(smilesParserT__83-76))|(1<<(smilesParserLB-76))|(1<<(smilesParserLC-76))|(1<<(smilesParserLN-76))|(1<<(smilesParserLO-76))|(1<<(smilesParserLP-76))|(1<<(smilesParserLS-76)))) != 0) || (((_la-115)&-(0x1f+1)) == 0 && ((1<<uint((_la-115)))&((1<<(smilesParserUB-115))|(1<<(smilesParserUC-115))|(1<<(smilesParserUF-115))|(1<<(smilesParserUI-115))|(1<<(smilesParserUN-115))|(1<<(smilesParserUO-115))|(1<<(smilesParserUP-115))|(1<<(smilesParserUS-115))|(1<<(smilesParserDOT-115)))) != 0) {
		p.SetState(451)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case smilesParserT__75, smilesParserT__78, smilesParserT__79, smilesParserT__80, smilesParserT__81, smilesParserT__82, smilesParserT__83:
			{
				p.SetState(449)
				p.Bond()
			}

		case smilesParserDOT:
			{
				p.SetState(450)
				p.Match(smilesParserDOT)
			}

		case smilesParserT__0, smilesParserT__1, smilesParserLB, smilesParserLC, smilesParserLN, smilesParserLO, smilesParserLP, smilesParserLS, smilesParserUB, smilesParserUC, smilesParserUF, smilesParserUI, smilesParserUN, smilesParserUO, smilesParserUP, smilesParserUS:

		default:
		}
		{
			p.SetState(453)
			p.Branched_atom()
		}

		p.SetState(458)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITerminatorContext is an interface to support dynamic dispatch.
type ITerminatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTerminatorContext differentiates from other interfaces.
	IsTerminatorContext()
}

type TerminatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTerminatorContext() *TerminatorContext {
	var p = new(TerminatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = smilesParserRULE_terminator
	return p
}

func (*TerminatorContext) IsTerminatorContext() {}

func NewTerminatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TerminatorContext {
	var p = new(TerminatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = smilesParserRULE_terminator

	return p
}

func (s *TerminatorContext) GetParser() antlr.Parser { return s.parser }

func (s *TerminatorContext) SPACE() antlr.TerminalNode {
	return s.GetToken(smilesParserSPACE, 0)
}

func (s *TerminatorContext) TAB() antlr.TerminalNode {
	return s.GetToken(smilesParserTAB, 0)
}

func (s *TerminatorContext) LINEFEED() antlr.TerminalNode {
	return s.GetToken(smilesParserLINEFEED, 0)
}

func (s *TerminatorContext) CARRIAGE_RETURN() antlr.TerminalNode {
	return s.GetToken(smilesParserCARRIAGE_RETURN, 0)
}

func (s *TerminatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TerminatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TerminatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.EnterTerminator(s)
	}
}

func (s *TerminatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(smilesListener); ok {
		listenerT.ExitTerminator(s)
	}
}

func (p *smilesParser) Terminator() (localctx ITerminatorContext) {
	this := p
	_ = this

	localctx = NewTerminatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, smilesParserRULE_terminator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(463)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case smilesParserSPACE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(459)
			p.Match(smilesParserSPACE)
		}
		{
			p.SetState(460)
			p.Match(smilesParserTAB)
		}

	case smilesParserLINEFEED:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(461)
			p.Match(smilesParserLINEFEED)
		}

	case smilesParserCARRIAGE_RETURN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(462)
			p.Match(smilesParserCARRIAGE_RETURN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
