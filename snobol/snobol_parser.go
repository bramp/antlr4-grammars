// Code generated from snobol.g4 by ANTLR 4.9.3. DO NOT EDIT.

package snobol // snobol
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 57, 378,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 3, 2, 6, 2,
	88, 10, 2, 13, 2, 14, 2, 89, 3, 3, 5, 3, 93, 10, 3, 3, 3, 3, 3, 3, 4, 5,
	4, 98, 10, 4, 3, 4, 3, 4, 5, 4, 102, 10, 4, 3, 4, 3, 4, 6, 4, 106, 10,
	4, 13, 4, 14, 4, 107, 5, 4, 110, 10, 4, 3, 4, 3, 4, 5, 4, 114, 10, 4, 3,
	4, 3, 4, 3, 4, 5, 4, 119, 10, 4, 3, 5, 3, 5, 3, 6, 5, 6, 124, 10, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 131, 10, 6, 12, 6, 14, 6, 134, 11, 6,
	3, 6, 5, 6, 137, 10, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 7, 8, 144, 10, 8,
	12, 8, 14, 8, 147, 11, 8, 3, 9, 3, 9, 3, 9, 7, 9, 152, 10, 9, 12, 9, 14,
	9, 155, 11, 9, 3, 10, 3, 10, 3, 10, 5, 10, 160, 10, 10, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 169, 10, 11, 12, 11, 14, 11, 172,
	11, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 180, 10, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 209, 10, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21,
	3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34,
	3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 5, 42, 366, 10, 42, 3, 42, 3,
	42, 3, 42, 5, 42, 371, 10, 42, 3, 42, 5, 42, 374, 10, 42, 3, 43, 3, 43,
	3, 43, 2, 2, 44, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
	68, 70, 72, 74, 76, 78, 80, 82, 84, 2, 8, 4, 2, 49, 49, 55, 55, 3, 2, 50,
	51, 3, 2, 42, 43, 3, 2, 44, 45, 4, 2, 50, 51, 53, 53, 3, 2, 34, 37, 2,
	384, 2, 87, 3, 2, 2, 2, 4, 92, 3, 2, 2, 2, 6, 118, 3, 2, 2, 2, 8, 120,
	3, 2, 2, 2, 10, 123, 3, 2, 2, 2, 12, 138, 3, 2, 2, 2, 14, 140, 3, 2, 2,
	2, 16, 148, 3, 2, 2, 2, 18, 156, 3, 2, 2, 2, 20, 179, 3, 2, 2, 2, 22, 208,
	3, 2, 2, 2, 24, 210, 3, 2, 2, 2, 26, 215, 3, 2, 2, 2, 28, 220, 3, 2, 2,
	2, 30, 225, 3, 2, 2, 2, 32, 230, 3, 2, 2, 2, 34, 235, 3, 2, 2, 2, 36, 240,
	3, 2, 2, 2, 38, 245, 3, 2, 2, 2, 40, 250, 3, 2, 2, 2, 42, 255, 3, 2, 2,
	2, 44, 260, 3, 2, 2, 2, 46, 265, 3, 2, 2, 2, 48, 270, 3, 2, 2, 2, 50, 275,
	3, 2, 2, 2, 52, 280, 3, 2, 2, 2, 54, 285, 3, 2, 2, 2, 56, 290, 3, 2, 2,
	2, 58, 295, 3, 2, 2, 2, 60, 300, 3, 2, 2, 2, 62, 307, 3, 2, 2, 2, 64, 312,
	3, 2, 2, 2, 66, 316, 3, 2, 2, 2, 68, 325, 3, 2, 2, 2, 70, 330, 3, 2, 2,
	2, 72, 335, 3, 2, 2, 2, 74, 342, 3, 2, 2, 2, 76, 349, 3, 2, 2, 2, 78, 354,
	3, 2, 2, 2, 80, 359, 3, 2, 2, 2, 82, 373, 3, 2, 2, 2, 84, 375, 3, 2, 2,
	2, 86, 88, 5, 4, 3, 2, 87, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 87,
	3, 2, 2, 2, 89, 90, 3, 2, 2, 2, 90, 3, 3, 2, 2, 2, 91, 93, 5, 6, 4, 2,
	92, 91, 3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 95, 7,
	56, 2, 2, 95, 5, 3, 2, 2, 2, 96, 98, 5, 8, 5, 2, 97, 96, 3, 2, 2, 2, 97,
	98, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 101, 5, 10, 6, 2, 100, 102, 5,
	12, 7, 2, 101, 100, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 109, 3, 2, 2,
	2, 103, 105, 7, 47, 2, 2, 104, 106, 5, 14, 8, 2, 105, 104, 3, 2, 2, 2,
	106, 107, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108,
	110, 3, 2, 2, 2, 109, 103, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 113,
	3, 2, 2, 2, 111, 112, 7, 48, 2, 2, 112, 114, 5, 82, 42, 2, 113, 111, 3,
	2, 2, 2, 113, 114, 3, 2, 2, 2, 114, 119, 3, 2, 2, 2, 115, 116, 7, 48, 2,
	2, 116, 119, 5, 82, 42, 2, 117, 119, 9, 2, 2, 2, 118, 97, 3, 2, 2, 2, 118,
	115, 3, 2, 2, 2, 118, 117, 3, 2, 2, 2, 119, 7, 3, 2, 2, 2, 120, 121, 7,
	52, 2, 2, 121, 9, 3, 2, 2, 2, 122, 124, 7, 41, 2, 2, 123, 122, 3, 2, 2,
	2, 123, 124, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 136, 7, 52, 2, 2, 126,
	127, 7, 3, 2, 2, 127, 132, 7, 52, 2, 2, 128, 129, 7, 38, 2, 2, 129, 131,
	7, 52, 2, 2, 130, 128, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2,
	2, 2, 132, 133, 3, 2, 2, 2, 133, 135, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2,
	135, 137, 7, 4, 2, 2, 136, 126, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137,
	11, 3, 2, 2, 2, 138, 139, 9, 3, 2, 2, 139, 13, 3, 2, 2, 2, 140, 145, 5,
	16, 9, 2, 141, 142, 9, 4, 2, 2, 142, 144, 5, 16, 9, 2, 143, 141, 3, 2,
	2, 2, 144, 147, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2,
	146, 15, 3, 2, 2, 2, 147, 145, 3, 2, 2, 2, 148, 153, 5, 18, 10, 2, 149,
	150, 9, 5, 2, 2, 150, 152, 5, 18, 10, 2, 151, 149, 3, 2, 2, 2, 152, 155,
	3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 17, 3, 2,
	2, 2, 155, 153, 3, 2, 2, 2, 156, 159, 5, 20, 11, 2, 157, 158, 7, 46, 2,
	2, 158, 160, 5, 14, 8, 2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160,
	19, 3, 2, 2, 2, 161, 180, 9, 6, 2, 2, 162, 180, 5, 10, 6, 2, 163, 180,
	5, 22, 12, 2, 164, 165, 7, 3, 2, 2, 165, 170, 5, 14, 8, 2, 166, 167, 7,
	38, 2, 2, 167, 169, 5, 14, 8, 2, 168, 166, 3, 2, 2, 2, 169, 172, 3, 2,
	2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2,
	172, 170, 3, 2, 2, 2, 173, 174, 7, 4, 2, 2, 174, 180, 3, 2, 2, 2, 175,
	176, 7, 39, 2, 2, 176, 177, 5, 14, 8, 2, 177, 178, 7, 40, 2, 2, 178, 180,
	3, 2, 2, 2, 179, 161, 3, 2, 2, 2, 179, 162, 3, 2, 2, 2, 179, 163, 3, 2,
	2, 2, 179, 164, 3, 2, 2, 2, 179, 175, 3, 2, 2, 2, 180, 21, 3, 2, 2, 2,
	181, 209, 5, 24, 13, 2, 182, 209, 5, 26, 14, 2, 183, 209, 5, 28, 15, 2,
	184, 209, 5, 30, 16, 2, 185, 209, 5, 32, 17, 2, 186, 209, 5, 36, 19, 2,
	187, 209, 5, 38, 20, 2, 188, 209, 5, 40, 21, 2, 189, 209, 5, 42, 22, 2,
	190, 209, 5, 44, 23, 2, 191, 209, 5, 46, 24, 2, 192, 209, 5, 48, 25, 2,
	193, 209, 5, 50, 26, 2, 194, 209, 5, 52, 27, 2, 195, 209, 5, 54, 28, 2,
	196, 209, 5, 56, 29, 2, 197, 209, 5, 58, 30, 2, 198, 209, 5, 64, 33, 2,
	199, 209, 5, 60, 31, 2, 200, 209, 5, 62, 32, 2, 201, 209, 5, 66, 34, 2,
	202, 209, 5, 68, 35, 2, 203, 209, 5, 70, 36, 2, 204, 209, 5, 72, 37, 2,
	205, 209, 5, 78, 40, 2, 206, 209, 5, 76, 39, 2, 207, 209, 5, 80, 41, 2,
	208, 181, 3, 2, 2, 2, 208, 182, 3, 2, 2, 2, 208, 183, 3, 2, 2, 2, 208,
	184, 3, 2, 2, 2, 208, 185, 3, 2, 2, 2, 208, 186, 3, 2, 2, 2, 208, 187,
	3, 2, 2, 2, 208, 188, 3, 2, 2, 2, 208, 189, 3, 2, 2, 2, 208, 190, 3, 2,
	2, 2, 208, 191, 3, 2, 2, 2, 208, 192, 3, 2, 2, 2, 208, 193, 3, 2, 2, 2,
	208, 194, 3, 2, 2, 2, 208, 195, 3, 2, 2, 2, 208, 196, 3, 2, 2, 2, 208,
	197, 3, 2, 2, 2, 208, 198, 3, 2, 2, 2, 208, 199, 3, 2, 2, 2, 208, 200,
	3, 2, 2, 2, 208, 201, 3, 2, 2, 2, 208, 202, 3, 2, 2, 2, 208, 203, 3, 2,
	2, 2, 208, 204, 3, 2, 2, 2, 208, 205, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2,
	208, 207, 3, 2, 2, 2, 209, 23, 3, 2, 2, 2, 210, 211, 7, 5, 2, 2, 211, 212,
	7, 39, 2, 2, 212, 213, 5, 14, 8, 2, 213, 214, 7, 40, 2, 2, 214, 25, 3,
	2, 2, 2, 215, 216, 7, 6, 2, 2, 216, 217, 7, 39, 2, 2, 217, 218, 5, 14,
	8, 2, 218, 219, 7, 40, 2, 2, 219, 27, 3, 2, 2, 2, 220, 221, 7, 7, 2, 2,
	221, 222, 7, 39, 2, 2, 222, 223, 5, 14, 8, 2, 223, 224, 7, 40, 2, 2, 224,
	29, 3, 2, 2, 2, 225, 226, 7, 8, 2, 2, 226, 227, 7, 39, 2, 2, 227, 228,
	5, 14, 8, 2, 228, 229, 7, 40, 2, 2, 229, 31, 3, 2, 2, 2, 230, 231, 7, 9,
	2, 2, 231, 232, 7, 39, 2, 2, 232, 233, 5, 14, 8, 2, 233, 234, 7, 40, 2,
	2, 234, 33, 3, 2, 2, 2, 235, 236, 7, 10, 2, 2, 236, 237, 7, 39, 2, 2, 237,
	238, 5, 14, 8, 2, 238, 239, 7, 40, 2, 2, 239, 35, 3, 2, 2, 2, 240, 241,
	7, 11, 2, 2, 241, 242, 7, 39, 2, 2, 242, 243, 5, 14, 8, 2, 243, 244, 7,
	40, 2, 2, 244, 37, 3, 2, 2, 2, 245, 246, 7, 12, 2, 2, 246, 247, 7, 39,
	2, 2, 247, 248, 5, 14, 8, 2, 248, 249, 7, 40, 2, 2, 249, 39, 3, 2, 2, 2,
	250, 251, 7, 13, 2, 2, 251, 252, 7, 39, 2, 2, 252, 253, 5, 14, 8, 2, 253,
	254, 7, 40, 2, 2, 254, 41, 3, 2, 2, 2, 255, 256, 7, 14, 2, 2, 256, 257,
	7, 39, 2, 2, 257, 258, 5, 14, 8, 2, 258, 259, 7, 40, 2, 2, 259, 43, 3,
	2, 2, 2, 260, 261, 7, 15, 2, 2, 261, 262, 7, 39, 2, 2, 262, 263, 5, 14,
	8, 2, 263, 264, 7, 40, 2, 2, 264, 45, 3, 2, 2, 2, 265, 266, 7, 16, 2, 2,
	266, 267, 7, 39, 2, 2, 267, 268, 5, 14, 8, 2, 268, 269, 7, 40, 2, 2, 269,
	47, 3, 2, 2, 2, 270, 271, 7, 17, 2, 2, 271, 272, 7, 39, 2, 2, 272, 273,
	5, 14, 8, 2, 273, 274, 7, 40, 2, 2, 274, 49, 3, 2, 2, 2, 275, 276, 7, 18,
	2, 2, 276, 277, 7, 39, 2, 2, 277, 278, 5, 14, 8, 2, 278, 279, 7, 40, 2,
	2, 279, 51, 3, 2, 2, 2, 280, 281, 7, 19, 2, 2, 281, 282, 7, 39, 2, 2, 282,
	283, 5, 14, 8, 2, 283, 284, 7, 40, 2, 2, 284, 53, 3, 2, 2, 2, 285, 286,
	7, 20, 2, 2, 286, 287, 7, 39, 2, 2, 287, 288, 5, 14, 8, 2, 288, 289, 7,
	40, 2, 2, 289, 55, 3, 2, 2, 2, 290, 291, 7, 21, 2, 2, 291, 292, 7, 39,
	2, 2, 292, 293, 5, 14, 8, 2, 293, 294, 7, 40, 2, 2, 294, 57, 3, 2, 2, 2,
	295, 296, 7, 22, 2, 2, 296, 297, 7, 39, 2, 2, 297, 298, 5, 14, 8, 2, 298,
	299, 7, 40, 2, 2, 299, 59, 3, 2, 2, 2, 300, 301, 7, 23, 2, 2, 301, 302,
	7, 39, 2, 2, 302, 303, 5, 14, 8, 2, 303, 304, 7, 38, 2, 2, 304, 305, 5,
	14, 8, 2, 305, 306, 7, 40, 2, 2, 306, 61, 3, 2, 2, 2, 307, 308, 7, 24,
	2, 2, 308, 309, 7, 39, 2, 2, 309, 310, 5, 14, 8, 2, 310, 311, 7, 40, 2,
	2, 311, 63, 3, 2, 2, 2, 312, 313, 7, 25, 2, 2, 313, 314, 7, 39, 2, 2, 314,
	315, 7, 40, 2, 2, 315, 65, 3, 2, 2, 2, 316, 317, 7, 26, 2, 2, 317, 318,
	7, 39, 2, 2, 318, 319, 5, 14, 8, 2, 319, 320, 7, 38, 2, 2, 320, 321, 5,
	14, 8, 2, 321, 322, 7, 38, 2, 2, 322, 323, 5, 14, 8, 2, 323, 324, 7, 40,
	2, 2, 324, 67, 3, 2, 2, 2, 325, 326, 7, 27, 2, 2, 326, 327, 7, 39, 2, 2,
	327, 328, 5, 14, 8, 2, 328, 329, 7, 40, 2, 2, 329, 69, 3, 2, 2, 2, 330,
	331, 7, 28, 2, 2, 331, 332, 7, 39, 2, 2, 332, 333, 5, 14, 8, 2, 333, 334,
	7, 40, 2, 2, 334, 71, 3, 2, 2, 2, 335, 336, 7, 29, 2, 2, 336, 337, 7, 39,
	2, 2, 337, 338, 5, 14, 8, 2, 338, 339, 7, 38, 2, 2, 339, 340, 5, 14, 8,
	2, 340, 341, 7, 40, 2, 2, 341, 73, 3, 2, 2, 2, 342, 343, 7, 30, 2, 2, 343,
	344, 7, 39, 2, 2, 344, 345, 5, 14, 8, 2, 345, 346, 7, 38, 2, 2, 346, 347,
	5, 14, 8, 2, 347, 348, 7, 40, 2, 2, 348, 75, 3, 2, 2, 2, 349, 350, 7, 31,
	2, 2, 350, 351, 7, 39, 2, 2, 351, 352, 5, 14, 8, 2, 352, 353, 7, 40, 2,
	2, 353, 77, 3, 2, 2, 2, 354, 355, 7, 32, 2, 2, 355, 356, 7, 39, 2, 2, 356,
	357, 5, 14, 8, 2, 357, 358, 7, 40, 2, 2, 358, 79, 3, 2, 2, 2, 359, 360,
	7, 33, 2, 2, 360, 361, 7, 39, 2, 2, 361, 362, 5, 14, 8, 2, 362, 363, 7,
	40, 2, 2, 363, 81, 3, 2, 2, 2, 364, 366, 5, 84, 43, 2, 365, 364, 3, 2,
	2, 2, 365, 366, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 370, 7, 39, 2, 2,
	368, 371, 5, 8, 5, 2, 369, 371, 7, 49, 2, 2, 370, 368, 3, 2, 2, 2, 370,
	369, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 374, 7, 40, 2, 2, 373, 365,
	3, 2, 2, 2, 373, 374, 3, 2, 2, 2, 374, 83, 3, 2, 2, 2, 375, 376, 9, 7,
	2, 2, 376, 85, 3, 2, 2, 2, 22, 89, 92, 97, 101, 107, 109, 113, 118, 123,
	132, 136, 145, 153, 159, 170, 179, 208, 365, 370, 373,
}
var literalNames = []string{
	"", "'['", "']'", "'ident'", "'differ'", "'eq'", "'ne'", "'ge'", "'gt'",
	"'le'", "'lt'", "'integer'", "'lgt'", "'atan'", "'chop'", "'cos'", "'exp'",
	"'ln'", "'remdr'", "'sin'", "'tan'", "'dupl'", "'reverse'", "'date'", "'replace'",
	"'size'", "'trim'", "'array'", "'convert'", "'table'", "'sort'", "'break'",
	"'f'", "'F'", "'s'", "'S'", "','", "'('", "')'", "'&'", "'+'", "'-'", "'*'",
	"'/'", "'^'", "'='", "':'", "'END'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"COMMA", "LPAREN", "RPAREN", "AMP", "PLUS", "MINUS", "TIMES", "DIV", "POW",
	"EQ", "COLON", "END", "STRINGLITERAL1", "STRINGLITERAL2", "STRING", "INTEGER",
	"REAL", "COMMENT", "EOL", "WS",
}

var ruleNames = []string{
	"prog", "lin", "line", "label", "subject", "pattern", "expression", "multiplyingExpression",
	"powExpression", "atom", "command", "ident", "differ", "eq", "ne", "ge",
	"gt", "le", "lt", "integer", "lgt", "atan", "chop", "cos", "exp", "ln",
	"remdr", "sin", "tan", "dupl", "reverse", "date", "replace", "size", "trim",
	"array_", "convert", "table", "sort", "break_", "transfer", "transferpre",
}

type snobolParser struct {
	*antlr.BaseParser
}

// NewsnobolParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *snobolParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewsnobolParser(input antlr.TokenStream) *snobolParser {
	this := new(snobolParser)
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
	this.GrammarFileName = "snobol.g4"

	return this
}

// snobolParser tokens.
const (
	snobolParserEOF            = antlr.TokenEOF
	snobolParserT__0           = 1
	snobolParserT__1           = 2
	snobolParserT__2           = 3
	snobolParserT__3           = 4
	snobolParserT__4           = 5
	snobolParserT__5           = 6
	snobolParserT__6           = 7
	snobolParserT__7           = 8
	snobolParserT__8           = 9
	snobolParserT__9           = 10
	snobolParserT__10          = 11
	snobolParserT__11          = 12
	snobolParserT__12          = 13
	snobolParserT__13          = 14
	snobolParserT__14          = 15
	snobolParserT__15          = 16
	snobolParserT__16          = 17
	snobolParserT__17          = 18
	snobolParserT__18          = 19
	snobolParserT__19          = 20
	snobolParserT__20          = 21
	snobolParserT__21          = 22
	snobolParserT__22          = 23
	snobolParserT__23          = 24
	snobolParserT__24          = 25
	snobolParserT__25          = 26
	snobolParserT__26          = 27
	snobolParserT__27          = 28
	snobolParserT__28          = 29
	snobolParserT__29          = 30
	snobolParserT__30          = 31
	snobolParserT__31          = 32
	snobolParserT__32          = 33
	snobolParserT__33          = 34
	snobolParserT__34          = 35
	snobolParserCOMMA          = 36
	snobolParserLPAREN         = 37
	snobolParserRPAREN         = 38
	snobolParserAMP            = 39
	snobolParserPLUS           = 40
	snobolParserMINUS          = 41
	snobolParserTIMES          = 42
	snobolParserDIV            = 43
	snobolParserPOW            = 44
	snobolParserEQ             = 45
	snobolParserCOLON          = 46
	snobolParserEND            = 47
	snobolParserSTRINGLITERAL1 = 48
	snobolParserSTRINGLITERAL2 = 49
	snobolParserSTRING         = 50
	snobolParserINTEGER        = 51
	snobolParserREAL           = 52
	snobolParserCOMMENT        = 53
	snobolParserEOL            = 54
	snobolParserWS             = 55
)

// snobolParser rules.
const (
	snobolParserRULE_prog                  = 0
	snobolParserRULE_lin                   = 1
	snobolParserRULE_line                  = 2
	snobolParserRULE_label                 = 3
	snobolParserRULE_subject               = 4
	snobolParserRULE_pattern               = 5
	snobolParserRULE_expression            = 6
	snobolParserRULE_multiplyingExpression = 7
	snobolParserRULE_powExpression         = 8
	snobolParserRULE_atom                  = 9
	snobolParserRULE_command               = 10
	snobolParserRULE_ident                 = 11
	snobolParserRULE_differ                = 12
	snobolParserRULE_eq                    = 13
	snobolParserRULE_ne                    = 14
	snobolParserRULE_ge                    = 15
	snobolParserRULE_gt                    = 16
	snobolParserRULE_le                    = 17
	snobolParserRULE_lt                    = 18
	snobolParserRULE_integer               = 19
	snobolParserRULE_lgt                   = 20
	snobolParserRULE_atan                  = 21
	snobolParserRULE_chop                  = 22
	snobolParserRULE_cos                   = 23
	snobolParserRULE_exp                   = 24
	snobolParserRULE_ln                    = 25
	snobolParserRULE_remdr                 = 26
	snobolParserRULE_sin                   = 27
	snobolParserRULE_tan                   = 28
	snobolParserRULE_dupl                  = 29
	snobolParserRULE_reverse               = 30
	snobolParserRULE_date                  = 31
	snobolParserRULE_replace               = 32
	snobolParserRULE_size                  = 33
	snobolParserRULE_trim                  = 34
	snobolParserRULE_array_                = 35
	snobolParserRULE_convert               = 36
	snobolParserRULE_table                 = 37
	snobolParserRULE_sort                  = 38
	snobolParserRULE_break_                = 39
	snobolParserRULE_transfer              = 40
	snobolParserRULE_transferpre           = 41
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllLin() []ILinContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILinContext)(nil)).Elem())
	var tst = make([]ILinContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILinContext)
		}
	}

	return tst
}

func (s *ProgContext) Lin(i int) ILinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILinContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILinContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *snobolParser) Prog() (localctx IProgContext) {
	this := p
	_ = this

	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, snobolParserRULE_prog)
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
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(snobolParserAMP-39))|(1<<(snobolParserCOLON-39))|(1<<(snobolParserEND-39))|(1<<(snobolParserSTRING-39))|(1<<(snobolParserCOMMENT-39))|(1<<(snobolParserEOL-39)))) != 0) {
		{
			p.SetState(84)
			p.Lin()
		}

		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILinContext is an interface to support dynamic dispatch.
type ILinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLinContext differentiates from other interfaces.
	IsLinContext()
}

type LinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLinContext() *LinContext {
	var p = new(LinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_lin
	return p
}

func (*LinContext) IsLinContext() {}

func NewLinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LinContext {
	var p = new(LinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_lin

	return p
}

func (s *LinContext) GetParser() antlr.Parser { return s.parser }

func (s *LinContext) EOL() antlr.TerminalNode {
	return s.GetToken(snobolParserEOL, 0)
}

func (s *LinContext) Line() ILineContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineContext)
}

func (s *LinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterLin(s)
	}
}

func (s *LinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitLin(s)
	}
}

func (p *snobolParser) Lin() (localctx ILinContext) {
	this := p
	_ = this

	localctx = NewLinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, snobolParserRULE_lin)
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
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(snobolParserAMP-39))|(1<<(snobolParserCOLON-39))|(1<<(snobolParserEND-39))|(1<<(snobolParserSTRING-39))|(1<<(snobolParserCOMMENT-39)))) != 0 {
		{
			p.SetState(89)
			p.Line()
		}

	}
	{
		p.SetState(92)
		p.Match(snobolParserEOL)
	}

	return localctx
}

// ILineContext is an interface to support dynamic dispatch.
type ILineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineContext differentiates from other interfaces.
	IsLineContext()
}

type LineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineContext() *LineContext {
	var p = new(LineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_line
	return p
}

func (*LineContext) IsLineContext() {}

func NewLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineContext {
	var p = new(LineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_line

	return p
}

func (s *LineContext) GetParser() antlr.Parser { return s.parser }

func (s *LineContext) Subject() ISubjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubjectContext)
}

func (s *LineContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *LineContext) Pattern() IPatternContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPatternContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPatternContext)
}

func (s *LineContext) EQ() antlr.TerminalNode {
	return s.GetToken(snobolParserEQ, 0)
}

func (s *LineContext) COLON() antlr.TerminalNode {
	return s.GetToken(snobolParserCOLON, 0)
}

func (s *LineContext) Transfer() ITransferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransferContext)
}

func (s *LineContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LineContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LineContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(snobolParserCOMMENT, 0)
}

func (s *LineContext) END() antlr.TerminalNode {
	return s.GetToken(snobolParserEND, 0)
}

func (s *LineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterLine(s)
	}
}

func (s *LineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitLine(s)
	}
}

func (p *snobolParser) Line() (localctx ILineContext) {
	this := p
	_ = this

	localctx = NewLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, snobolParserRULE_line)
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

	p.SetState(116)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case snobolParserAMP, snobolParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(95)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(94)
				p.Label()
			}

		}
		{
			p.SetState(97)
			p.Subject()
		}
		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == snobolParserSTRINGLITERAL1 || _la == snobolParserSTRINGLITERAL2 {
			{
				p.SetState(98)
				p.Pattern()
			}

		}
		p.SetState(107)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == snobolParserEQ {
			{
				p.SetState(101)
				p.Match(snobolParserEQ)
			}
			p.SetState(103)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<snobolParserT__0)|(1<<snobolParserT__2)|(1<<snobolParserT__3)|(1<<snobolParserT__4)|(1<<snobolParserT__5)|(1<<snobolParserT__6)|(1<<snobolParserT__8)|(1<<snobolParserT__9)|(1<<snobolParserT__10)|(1<<snobolParserT__11)|(1<<snobolParserT__12)|(1<<snobolParserT__13)|(1<<snobolParserT__14)|(1<<snobolParserT__15)|(1<<snobolParserT__16)|(1<<snobolParserT__17)|(1<<snobolParserT__18)|(1<<snobolParserT__19)|(1<<snobolParserT__20)|(1<<snobolParserT__21)|(1<<snobolParserT__22)|(1<<snobolParserT__23)|(1<<snobolParserT__24)|(1<<snobolParserT__25)|(1<<snobolParserT__26)|(1<<snobolParserT__28)|(1<<snobolParserT__29)|(1<<snobolParserT__30))) != 0) || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(snobolParserLPAREN-37))|(1<<(snobolParserAMP-37))|(1<<(snobolParserSTRINGLITERAL1-37))|(1<<(snobolParserSTRINGLITERAL2-37))|(1<<(snobolParserSTRING-37))|(1<<(snobolParserINTEGER-37)))) != 0) {
				{
					p.SetState(102)
					p.Expression()
				}

				p.SetState(105)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		p.SetState(111)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == snobolParserCOLON {
			{
				p.SetState(109)
				p.Match(snobolParserCOLON)
			}
			{
				p.SetState(110)
				p.Transfer()
			}

		}

	case snobolParserCOLON:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(113)
			p.Match(snobolParserCOLON)
		}
		{
			p.SetState(114)
			p.Transfer()
		}

	case snobolParserEND, snobolParserCOMMENT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(115)
			_la = p.GetTokenStream().LA(1)

			if !(_la == snobolParserEND || _la == snobolParserCOMMENT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILabelContext is an interface to support dynamic dispatch.
type ILabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLabelContext differentiates from other interfaces.
	IsLabelContext()
}

type LabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLabelContext() *LabelContext {
	var p = new(LabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_label
	return p
}

func (*LabelContext) IsLabelContext() {}

func NewLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LabelContext {
	var p = new(LabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_label

	return p
}

func (s *LabelContext) GetParser() antlr.Parser { return s.parser }

func (s *LabelContext) STRING() antlr.TerminalNode {
	return s.GetToken(snobolParserSTRING, 0)
}

func (s *LabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterLabel(s)
	}
}

func (s *LabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitLabel(s)
	}
}

func (p *snobolParser) Label() (localctx ILabelContext) {
	this := p
	_ = this

	localctx = NewLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, snobolParserRULE_label)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Match(snobolParserSTRING)
	}

	return localctx
}

// ISubjectContext is an interface to support dynamic dispatch.
type ISubjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubjectContext differentiates from other interfaces.
	IsSubjectContext()
}

type SubjectContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubjectContext() *SubjectContext {
	var p = new(SubjectContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_subject
	return p
}

func (*SubjectContext) IsSubjectContext() {}

func NewSubjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SubjectContext {
	var p = new(SubjectContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_subject

	return p
}

func (s *SubjectContext) GetParser() antlr.Parser { return s.parser }

func (s *SubjectContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(snobolParserSTRING)
}

func (s *SubjectContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(snobolParserSTRING, i)
}

func (s *SubjectContext) AMP() antlr.TerminalNode {
	return s.GetToken(snobolParserAMP, 0)
}

func (s *SubjectContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(snobolParserCOMMA)
}

func (s *SubjectContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(snobolParserCOMMA, i)
}

func (s *SubjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SubjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterSubject(s)
	}
}

func (s *SubjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitSubject(s)
	}
}

func (p *snobolParser) Subject() (localctx ISubjectContext) {
	this := p
	_ = this

	localctx = NewSubjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, snobolParserRULE_subject)
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
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == snobolParserAMP {
		{
			p.SetState(120)
			p.Match(snobolParserAMP)
		}

	}
	{
		p.SetState(123)
		p.Match(snobolParserSTRING)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(124)
			p.Match(snobolParserT__0)
		}
		{
			p.SetState(125)
			p.Match(snobolParserSTRING)
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == snobolParserCOMMA {
			{
				p.SetState(126)
				p.Match(snobolParserCOMMA)
			}
			{
				p.SetState(127)
				p.Match(snobolParserSTRING)
			}

			p.SetState(132)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(133)
			p.Match(snobolParserT__1)
		}

	}

	return localctx
}

// IPatternContext is an interface to support dynamic dispatch.
type IPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPatternContext differentiates from other interfaces.
	IsPatternContext()
}

type PatternContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPatternContext() *PatternContext {
	var p = new(PatternContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_pattern
	return p
}

func (*PatternContext) IsPatternContext() {}

func NewPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PatternContext {
	var p = new(PatternContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_pattern

	return p
}

func (s *PatternContext) GetParser() antlr.Parser { return s.parser }

func (s *PatternContext) STRINGLITERAL1() antlr.TerminalNode {
	return s.GetToken(snobolParserSTRINGLITERAL1, 0)
}

func (s *PatternContext) STRINGLITERAL2() antlr.TerminalNode {
	return s.GetToken(snobolParserSTRINGLITERAL2, 0)
}

func (s *PatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterPattern(s)
	}
}

func (s *PatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitPattern(s)
	}
}

func (p *snobolParser) Pattern() (localctx IPatternContext) {
	this := p
	_ = this

	localctx = NewPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, snobolParserRULE_pattern)
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
		p.SetState(136)
		_la = p.GetTokenStream().LA(1)

		if !(_la == snobolParserSTRINGLITERAL1 || _la == snobolParserSTRINGLITERAL2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = snobolParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllMultiplyingExpression() []IMultiplyingExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem())
	var tst = make([]IMultiplyingExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultiplyingExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) MultiplyingExpression(i int) IMultiplyingExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiplyingExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultiplyingExpressionContext)
}

func (s *ExpressionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(snobolParserPLUS)
}

func (s *ExpressionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(snobolParserPLUS, i)
}

func (s *ExpressionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(snobolParserMINUS)
}

func (s *ExpressionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(snobolParserMINUS, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *snobolParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, snobolParserRULE_expression)
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
		p.SetState(138)
		p.MultiplyingExpression()
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(139)
				_la = p.GetTokenStream().LA(1)

				if !(_la == snobolParserPLUS || _la == snobolParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(140)
				p.MultiplyingExpression()
			}

		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext())
	}

	return localctx
}

// IMultiplyingExpressionContext is an interface to support dynamic dispatch.
type IMultiplyingExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiplyingExpressionContext differentiates from other interfaces.
	IsMultiplyingExpressionContext()
}

type MultiplyingExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplyingExpressionContext() *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_multiplyingExpression
	return p
}

func (*MultiplyingExpressionContext) IsMultiplyingExpressionContext() {}

func NewMultiplyingExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplyingExpressionContext {
	var p = new(MultiplyingExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_multiplyingExpression

	return p
}

func (s *MultiplyingExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplyingExpressionContext) AllPowExpression() []IPowExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPowExpressionContext)(nil)).Elem())
	var tst = make([]IPowExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPowExpressionContext)
		}
	}

	return tst
}

func (s *MultiplyingExpressionContext) PowExpression(i int) IPowExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPowExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPowExpressionContext)
}

func (s *MultiplyingExpressionContext) AllTIMES() []antlr.TerminalNode {
	return s.GetTokens(snobolParserTIMES)
}

func (s *MultiplyingExpressionContext) TIMES(i int) antlr.TerminalNode {
	return s.GetToken(snobolParserTIMES, i)
}

func (s *MultiplyingExpressionContext) AllDIV() []antlr.TerminalNode {
	return s.GetTokens(snobolParserDIV)
}

func (s *MultiplyingExpressionContext) DIV(i int) antlr.TerminalNode {
	return s.GetToken(snobolParserDIV, i)
}

func (s *MultiplyingExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyingExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplyingExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterMultiplyingExpression(s)
	}
}

func (s *MultiplyingExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitMultiplyingExpression(s)
	}
}

func (p *snobolParser) MultiplyingExpression() (localctx IMultiplyingExpressionContext) {
	this := p
	_ = this

	localctx = NewMultiplyingExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, snobolParserRULE_multiplyingExpression)
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
		p.SetState(146)
		p.PowExpression()
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(147)
				_la = p.GetTokenStream().LA(1)

				if !(_la == snobolParserTIMES || _la == snobolParserDIV) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(148)
				p.PowExpression()
			}

		}
		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IPowExpressionContext is an interface to support dynamic dispatch.
type IPowExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPowExpressionContext differentiates from other interfaces.
	IsPowExpressionContext()
}

type PowExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPowExpressionContext() *PowExpressionContext {
	var p = new(PowExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_powExpression
	return p
}

func (*PowExpressionContext) IsPowExpressionContext() {}

func NewPowExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PowExpressionContext {
	var p = new(PowExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_powExpression

	return p
}

func (s *PowExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PowExpressionContext) Atom() IAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomContext)
}

func (s *PowExpressionContext) POW() antlr.TerminalNode {
	return s.GetToken(snobolParserPOW, 0)
}

func (s *PowExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PowExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PowExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterPowExpression(s)
	}
}

func (s *PowExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitPowExpression(s)
	}
}

func (p *snobolParser) PowExpression() (localctx IPowExpressionContext) {
	this := p
	_ = this

	localctx = NewPowExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, snobolParserRULE_powExpression)
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
		p.SetState(154)
		p.Atom()
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == snobolParserPOW {
		{
			p.SetState(155)
			p.Match(snobolParserPOW)
		}
		{
			p.SetState(156)
			p.Expression()
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
	p.RuleIndex = snobolParserRULE_atom
	return p
}

func (*AtomContext) IsAtomContext() {}

func NewAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomContext {
	var p = new(AtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_atom

	return p
}

func (s *AtomContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomContext) STRINGLITERAL1() antlr.TerminalNode {
	return s.GetToken(snobolParserSTRINGLITERAL1, 0)
}

func (s *AtomContext) STRINGLITERAL2() antlr.TerminalNode {
	return s.GetToken(snobolParserSTRINGLITERAL2, 0)
}

func (s *AtomContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(snobolParserINTEGER, 0)
}

func (s *AtomContext) Subject() ISubjectContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubjectContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubjectContext)
}

func (s *AtomContext) Command() ICommandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandContext)
}

func (s *AtomContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AtomContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AtomContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(snobolParserCOMMA)
}

func (s *AtomContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(snobolParserCOMMA, i)
}

func (s *AtomContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *AtomContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *AtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterAtom(s)
	}
}

func (s *AtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitAtom(s)
	}
}

func (p *snobolParser) Atom() (localctx IAtomContext) {
	this := p
	_ = this

	localctx = NewAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, snobolParserRULE_atom)
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

	p.SetState(177)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case snobolParserSTRINGLITERAL1, snobolParserSTRINGLITERAL2, snobolParserINTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(159)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-48)&-(0x1f+1)) == 0 && ((1<<uint((_la-48)))&((1<<(snobolParserSTRINGLITERAL1-48))|(1<<(snobolParserSTRINGLITERAL2-48))|(1<<(snobolParserINTEGER-48)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case snobolParserAMP, snobolParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(160)
			p.Subject()
		}

	case snobolParserT__2, snobolParserT__3, snobolParserT__4, snobolParserT__5, snobolParserT__6, snobolParserT__8, snobolParserT__9, snobolParserT__10, snobolParserT__11, snobolParserT__12, snobolParserT__13, snobolParserT__14, snobolParserT__15, snobolParserT__16, snobolParserT__17, snobolParserT__18, snobolParserT__19, snobolParserT__20, snobolParserT__21, snobolParserT__22, snobolParserT__23, snobolParserT__24, snobolParserT__25, snobolParserT__26, snobolParserT__28, snobolParserT__29, snobolParserT__30:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(161)
			p.Command()
		}

	case snobolParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(162)
			p.Match(snobolParserT__0)
		}
		{
			p.SetState(163)
			p.Expression()
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == snobolParserCOMMA {
			{
				p.SetState(164)
				p.Match(snobolParserCOMMA)
			}
			{
				p.SetState(165)
				p.Expression()
			}

			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(171)
			p.Match(snobolParserT__1)
		}

	case snobolParserLPAREN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(173)
			p.Match(snobolParserLPAREN)
		}
		{
			p.SetState(174)
			p.Expression()
		}
		{
			p.SetState(175)
			p.Match(snobolParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICommandContext is an interface to support dynamic dispatch.
type ICommandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandContext differentiates from other interfaces.
	IsCommandContext()
}

type CommandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandContext() *CommandContext {
	var p = new(CommandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_command
	return p
}

func (*CommandContext) IsCommandContext() {}

func NewCommandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandContext {
	var p = new(CommandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_command

	return p
}

func (s *CommandContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandContext) Ident() IIdentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *CommandContext) Differ() IDifferContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDifferContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDifferContext)
}

func (s *CommandContext) Eq() IEqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEqContext)
}

func (s *CommandContext) Ne() INeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INeContext)
}

func (s *CommandContext) Ge() IGeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeContext)
}

func (s *CommandContext) Le() ILeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILeContext)
}

func (s *CommandContext) Lt() ILtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILtContext)
}

func (s *CommandContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *CommandContext) Lgt() ILgtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILgtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILgtContext)
}

func (s *CommandContext) Atan() IAtanContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtanContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtanContext)
}

func (s *CommandContext) Chop() IChopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IChopContext)
}

func (s *CommandContext) Cos() ICosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICosContext)
}

func (s *CommandContext) Exp() IExpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpContext)
}

func (s *CommandContext) Ln() ILnContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILnContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILnContext)
}

func (s *CommandContext) Remdr() IRemdrContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRemdrContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRemdrContext)
}

func (s *CommandContext) Sin() ISinContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISinContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISinContext)
}

func (s *CommandContext) Tan() ITanContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITanContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITanContext)
}

func (s *CommandContext) Date() IDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateContext)
}

func (s *CommandContext) Dupl() IDuplContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDuplContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDuplContext)
}

func (s *CommandContext) Reverse() IReverseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReverseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReverseContext)
}

func (s *CommandContext) Replace() IReplaceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReplaceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReplaceContext)
}

func (s *CommandContext) Size() ISizeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISizeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISizeContext)
}

func (s *CommandContext) Trim() ITrimContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITrimContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITrimContext)
}

func (s *CommandContext) Array_() IArray_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_Context)
}

func (s *CommandContext) Sort() ISortContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISortContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISortContext)
}

func (s *CommandContext) Table() ITableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITableContext)
}

func (s *CommandContext) Break_() IBreak_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreak_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreak_Context)
}

func (s *CommandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterCommand(s)
	}
}

func (s *CommandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitCommand(s)
	}
}

func (p *snobolParser) Command() (localctx ICommandContext) {
	this := p
	_ = this

	localctx = NewCommandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, snobolParserRULE_command)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(206)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case snobolParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(179)
			p.Ident()
		}

	case snobolParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(180)
			p.Differ()
		}

	case snobolParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(181)
			p.Eq()
		}

	case snobolParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(182)
			p.Ne()
		}

	case snobolParserT__6:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(183)
			p.Ge()
		}

	case snobolParserT__8:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(184)
			p.Le()
		}

	case snobolParserT__9:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(185)
			p.Lt()
		}

	case snobolParserT__10:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(186)
			p.Integer()
		}

	case snobolParserT__11:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(187)
			p.Lgt()
		}

	case snobolParserT__12:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(188)
			p.Atan()
		}

	case snobolParserT__13:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(189)
			p.Chop()
		}

	case snobolParserT__14:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(190)
			p.Cos()
		}

	case snobolParserT__15:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(191)
			p.Exp()
		}

	case snobolParserT__16:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(192)
			p.Ln()
		}

	case snobolParserT__17:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(193)
			p.Remdr()
		}

	case snobolParserT__18:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(194)
			p.Sin()
		}

	case snobolParserT__19:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(195)
			p.Tan()
		}

	case snobolParserT__22:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(196)
			p.Date()
		}

	case snobolParserT__20:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(197)
			p.Dupl()
		}

	case snobolParserT__21:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(198)
			p.Reverse()
		}

	case snobolParserT__23:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(199)
			p.Replace()
		}

	case snobolParserT__24:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(200)
			p.Size()
		}

	case snobolParserT__25:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(201)
			p.Trim()
		}

	case snobolParserT__26:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(202)
			p.Array_()
		}

	case snobolParserT__29:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(203)
			p.Sort()
		}

	case snobolParserT__28:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(204)
			p.Table()
		}

	case snobolParserT__30:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(205)
			p.Break_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_ident
	return p
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *IdentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IdentContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterIdent(s)
	}
}

func (s *IdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitIdent(s)
	}
}

func (p *snobolParser) Ident() (localctx IIdentContext) {
	this := p
	_ = this

	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, snobolParserRULE_ident)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(snobolParserT__2)
	}
	{
		p.SetState(209)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(210)
		p.Expression()
	}
	{
		p.SetState(211)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IDifferContext is an interface to support dynamic dispatch.
type IDifferContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDifferContext differentiates from other interfaces.
	IsDifferContext()
}

type DifferContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDifferContext() *DifferContext {
	var p = new(DifferContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_differ
	return p
}

func (*DifferContext) IsDifferContext() {}

func NewDifferContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DifferContext {
	var p = new(DifferContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_differ

	return p
}

func (s *DifferContext) GetParser() antlr.Parser { return s.parser }

func (s *DifferContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *DifferContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DifferContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *DifferContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DifferContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DifferContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterDiffer(s)
	}
}

func (s *DifferContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitDiffer(s)
	}
}

func (p *snobolParser) Differ() (localctx IDifferContext) {
	this := p
	_ = this

	localctx = NewDifferContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, snobolParserRULE_differ)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(snobolParserT__3)
	}
	{
		p.SetState(214)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(215)
		p.Expression()
	}
	{
		p.SetState(216)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IEqContext is an interface to support dynamic dispatch.
type IEqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEqContext differentiates from other interfaces.
	IsEqContext()
}

type EqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqContext() *EqContext {
	var p = new(EqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_eq
	return p
}

func (*EqContext) IsEqContext() {}

func NewEqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqContext {
	var p = new(EqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_eq

	return p
}

func (s *EqContext) GetParser() antlr.Parser { return s.parser }

func (s *EqContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *EqContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EqContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *EqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterEq(s)
	}
}

func (s *EqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitEq(s)
	}
}

func (p *snobolParser) Eq() (localctx IEqContext) {
	this := p
	_ = this

	localctx = NewEqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, snobolParserRULE_eq)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(218)
		p.Match(snobolParserT__4)
	}
	{
		p.SetState(219)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(220)
		p.Expression()
	}
	{
		p.SetState(221)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// INeContext is an interface to support dynamic dispatch.
type INeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNeContext differentiates from other interfaces.
	IsNeContext()
}

type NeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNeContext() *NeContext {
	var p = new(NeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_ne
	return p
}

func (*NeContext) IsNeContext() {}

func NewNeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NeContext {
	var p = new(NeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_ne

	return p
}

func (s *NeContext) GetParser() antlr.Parser { return s.parser }

func (s *NeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *NeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *NeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *NeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterNe(s)
	}
}

func (s *NeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitNe(s)
	}
}

func (p *snobolParser) Ne() (localctx INeContext) {
	this := p
	_ = this

	localctx = NewNeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, snobolParserRULE_ne)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(snobolParserT__5)
	}
	{
		p.SetState(224)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(225)
		p.Expression()
	}
	{
		p.SetState(226)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IGeContext is an interface to support dynamic dispatch.
type IGeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeContext differentiates from other interfaces.
	IsGeContext()
}

type GeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeContext() *GeContext {
	var p = new(GeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_ge
	return p
}

func (*GeContext) IsGeContext() {}

func NewGeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeContext {
	var p = new(GeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_ge

	return p
}

func (s *GeContext) GetParser() antlr.Parser { return s.parser }

func (s *GeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *GeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *GeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterGe(s)
	}
}

func (s *GeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitGe(s)
	}
}

func (p *snobolParser) Ge() (localctx IGeContext) {
	this := p
	_ = this

	localctx = NewGeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, snobolParserRULE_ge)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(snobolParserT__6)
	}
	{
		p.SetState(229)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(230)
		p.Expression()
	}
	{
		p.SetState(231)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IGtContext is an interface to support dynamic dispatch.
type IGtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGtContext differentiates from other interfaces.
	IsGtContext()
}

type GtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGtContext() *GtContext {
	var p = new(GtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_gt
	return p
}

func (*GtContext) IsGtContext() {}

func NewGtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GtContext {
	var p = new(GtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_gt

	return p
}

func (s *GtContext) GetParser() antlr.Parser { return s.parser }

func (s *GtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *GtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *GtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterGt(s)
	}
}

func (s *GtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitGt(s)
	}
}

func (p *snobolParser) Gt() (localctx IGtContext) {
	this := p
	_ = this

	localctx = NewGtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, snobolParserRULE_gt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(233)
		p.Match(snobolParserT__7)
	}
	{
		p.SetState(234)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(235)
		p.Expression()
	}
	{
		p.SetState(236)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// ILeContext is an interface to support dynamic dispatch.
type ILeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLeContext differentiates from other interfaces.
	IsLeContext()
}

type LeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLeContext() *LeContext {
	var p = new(LeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_le
	return p
}

func (*LeContext) IsLeContext() {}

func NewLeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LeContext {
	var p = new(LeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_le

	return p
}

func (s *LeContext) GetParser() antlr.Parser { return s.parser }

func (s *LeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *LeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *LeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterLe(s)
	}
}

func (s *LeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitLe(s)
	}
}

func (p *snobolParser) Le() (localctx ILeContext) {
	this := p
	_ = this

	localctx = NewLeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, snobolParserRULE_le)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		p.Match(snobolParserT__8)
	}
	{
		p.SetState(239)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(240)
		p.Expression()
	}
	{
		p.SetState(241)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// ILtContext is an interface to support dynamic dispatch.
type ILtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLtContext differentiates from other interfaces.
	IsLtContext()
}

type LtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLtContext() *LtContext {
	var p = new(LtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_lt
	return p
}

func (*LtContext) IsLtContext() {}

func NewLtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LtContext {
	var p = new(LtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_lt

	return p
}

func (s *LtContext) GetParser() antlr.Parser { return s.parser }

func (s *LtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *LtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *LtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterLt(s)
	}
}

func (s *LtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitLt(s)
	}
}

func (p *snobolParser) Lt() (localctx ILtContext) {
	this := p
	_ = this

	localctx = NewLtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, snobolParserRULE_lt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(snobolParserT__9)
	}
	{
		p.SetState(244)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(245)
		p.Expression()
	}
	{
		p.SetState(246)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *IntegerContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IntegerContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *snobolParser) Integer() (localctx IIntegerContext) {
	this := p
	_ = this

	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, snobolParserRULE_integer)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(snobolParserT__10)
	}
	{
		p.SetState(249)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(250)
		p.Expression()
	}
	{
		p.SetState(251)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// ILgtContext is an interface to support dynamic dispatch.
type ILgtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLgtContext differentiates from other interfaces.
	IsLgtContext()
}

type LgtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLgtContext() *LgtContext {
	var p = new(LgtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_lgt
	return p
}

func (*LgtContext) IsLgtContext() {}

func NewLgtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LgtContext {
	var p = new(LgtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_lgt

	return p
}

func (s *LgtContext) GetParser() antlr.Parser { return s.parser }

func (s *LgtContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *LgtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LgtContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *LgtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LgtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LgtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterLgt(s)
	}
}

func (s *LgtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitLgt(s)
	}
}

func (p *snobolParser) Lgt() (localctx ILgtContext) {
	this := p
	_ = this

	localctx = NewLgtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, snobolParserRULE_lgt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(253)
		p.Match(snobolParserT__11)
	}
	{
		p.SetState(254)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(255)
		p.Expression()
	}
	{
		p.SetState(256)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IAtanContext is an interface to support dynamic dispatch.
type IAtanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtanContext differentiates from other interfaces.
	IsAtanContext()
}

type AtanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtanContext() *AtanContext {
	var p = new(AtanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_atan
	return p
}

func (*AtanContext) IsAtanContext() {}

func NewAtanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtanContext {
	var p = new(AtanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_atan

	return p
}

func (s *AtanContext) GetParser() antlr.Parser { return s.parser }

func (s *AtanContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *AtanContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AtanContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *AtanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterAtan(s)
	}
}

func (s *AtanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitAtan(s)
	}
}

func (p *snobolParser) Atan() (localctx IAtanContext) {
	this := p
	_ = this

	localctx = NewAtanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, snobolParserRULE_atan)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(snobolParserT__12)
	}
	{
		p.SetState(259)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(260)
		p.Expression()
	}
	{
		p.SetState(261)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IChopContext is an interface to support dynamic dispatch.
type IChopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChopContext differentiates from other interfaces.
	IsChopContext()
}

type ChopContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChopContext() *ChopContext {
	var p = new(ChopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_chop
	return p
}

func (*ChopContext) IsChopContext() {}

func NewChopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChopContext {
	var p = new(ChopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_chop

	return p
}

func (s *ChopContext) GetParser() antlr.Parser { return s.parser }

func (s *ChopContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *ChopContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ChopContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *ChopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterChop(s)
	}
}

func (s *ChopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitChop(s)
	}
}

func (p *snobolParser) Chop() (localctx IChopContext) {
	this := p
	_ = this

	localctx = NewChopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, snobolParserRULE_chop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(snobolParserT__13)
	}
	{
		p.SetState(264)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(265)
		p.Expression()
	}
	{
		p.SetState(266)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// ICosContext is an interface to support dynamic dispatch.
type ICosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCosContext differentiates from other interfaces.
	IsCosContext()
}

type CosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCosContext() *CosContext {
	var p = new(CosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_cos
	return p
}

func (*CosContext) IsCosContext() {}

func NewCosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CosContext {
	var p = new(CosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_cos

	return p
}

func (s *CosContext) GetParser() antlr.Parser { return s.parser }

func (s *CosContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *CosContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CosContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *CosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterCos(s)
	}
}

func (s *CosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitCos(s)
	}
}

func (p *snobolParser) Cos() (localctx ICosContext) {
	this := p
	_ = this

	localctx = NewCosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, snobolParserRULE_cos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(snobolParserT__14)
	}
	{
		p.SetState(269)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(270)
		p.Expression()
	}
	{
		p.SetState(271)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IExpContext is an interface to support dynamic dispatch.
type IExpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpContext differentiates from other interfaces.
	IsExpContext()
}

type ExpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpContext() *ExpContext {
	var p = new(ExpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_exp
	return p
}

func (*ExpContext) IsExpContext() {}

func NewExpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpContext {
	var p = new(ExpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_exp

	return p
}

func (s *ExpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *ExpContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *ExpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterExp(s)
	}
}

func (s *ExpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitExp(s)
	}
}

func (p *snobolParser) Exp() (localctx IExpContext) {
	this := p
	_ = this

	localctx = NewExpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, snobolParserRULE_exp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(snobolParserT__15)
	}
	{
		p.SetState(274)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(275)
		p.Expression()
	}
	{
		p.SetState(276)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// ILnContext is an interface to support dynamic dispatch.
type ILnContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLnContext differentiates from other interfaces.
	IsLnContext()
}

type LnContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLnContext() *LnContext {
	var p = new(LnContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_ln
	return p
}

func (*LnContext) IsLnContext() {}

func NewLnContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LnContext {
	var p = new(LnContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_ln

	return p
}

func (s *LnContext) GetParser() antlr.Parser { return s.parser }

func (s *LnContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *LnContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LnContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *LnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LnContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterLn(s)
	}
}

func (s *LnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitLn(s)
	}
}

func (p *snobolParser) Ln() (localctx ILnContext) {
	this := p
	_ = this

	localctx = NewLnContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, snobolParserRULE_ln)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(snobolParserT__16)
	}
	{
		p.SetState(279)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(280)
		p.Expression()
	}
	{
		p.SetState(281)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IRemdrContext is an interface to support dynamic dispatch.
type IRemdrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRemdrContext differentiates from other interfaces.
	IsRemdrContext()
}

type RemdrContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRemdrContext() *RemdrContext {
	var p = new(RemdrContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_remdr
	return p
}

func (*RemdrContext) IsRemdrContext() {}

func NewRemdrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RemdrContext {
	var p = new(RemdrContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_remdr

	return p
}

func (s *RemdrContext) GetParser() antlr.Parser { return s.parser }

func (s *RemdrContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *RemdrContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RemdrContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *RemdrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RemdrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RemdrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterRemdr(s)
	}
}

func (s *RemdrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitRemdr(s)
	}
}

func (p *snobolParser) Remdr() (localctx IRemdrContext) {
	this := p
	_ = this

	localctx = NewRemdrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, snobolParserRULE_remdr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Match(snobolParserT__17)
	}
	{
		p.SetState(284)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(285)
		p.Expression()
	}
	{
		p.SetState(286)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// ISinContext is an interface to support dynamic dispatch.
type ISinContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSinContext differentiates from other interfaces.
	IsSinContext()
}

type SinContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySinContext() *SinContext {
	var p = new(SinContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_sin
	return p
}

func (*SinContext) IsSinContext() {}

func NewSinContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SinContext {
	var p = new(SinContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_sin

	return p
}

func (s *SinContext) GetParser() antlr.Parser { return s.parser }

func (s *SinContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *SinContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SinContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *SinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterSin(s)
	}
}

func (s *SinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitSin(s)
	}
}

func (p *snobolParser) Sin() (localctx ISinContext) {
	this := p
	_ = this

	localctx = NewSinContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, snobolParserRULE_sin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(snobolParserT__18)
	}
	{
		p.SetState(289)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(290)
		p.Expression()
	}
	{
		p.SetState(291)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// ITanContext is an interface to support dynamic dispatch.
type ITanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTanContext differentiates from other interfaces.
	IsTanContext()
}

type TanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTanContext() *TanContext {
	var p = new(TanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_tan
	return p
}

func (*TanContext) IsTanContext() {}

func NewTanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TanContext {
	var p = new(TanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_tan

	return p
}

func (s *TanContext) GetParser() antlr.Parser { return s.parser }

func (s *TanContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *TanContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TanContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *TanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterTan(s)
	}
}

func (s *TanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitTan(s)
	}
}

func (p *snobolParser) Tan() (localctx ITanContext) {
	this := p
	_ = this

	localctx = NewTanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, snobolParserRULE_tan)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(snobolParserT__19)
	}
	{
		p.SetState(294)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(295)
		p.Expression()
	}
	{
		p.SetState(296)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IDuplContext is an interface to support dynamic dispatch.
type IDuplContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDuplContext differentiates from other interfaces.
	IsDuplContext()
}

type DuplContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDuplContext() *DuplContext {
	var p = new(DuplContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_dupl
	return p
}

func (*DuplContext) IsDuplContext() {}

func NewDuplContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DuplContext {
	var p = new(DuplContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_dupl

	return p
}

func (s *DuplContext) GetParser() antlr.Parser { return s.parser }

func (s *DuplContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *DuplContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DuplContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DuplContext) COMMA() antlr.TerminalNode {
	return s.GetToken(snobolParserCOMMA, 0)
}

func (s *DuplContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *DuplContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DuplContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DuplContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterDupl(s)
	}
}

func (s *DuplContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitDupl(s)
	}
}

func (p *snobolParser) Dupl() (localctx IDuplContext) {
	this := p
	_ = this

	localctx = NewDuplContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, snobolParserRULE_dupl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(298)
		p.Match(snobolParserT__20)
	}
	{
		p.SetState(299)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(300)
		p.Expression()
	}
	{
		p.SetState(301)
		p.Match(snobolParserCOMMA)
	}
	{
		p.SetState(302)
		p.Expression()
	}
	{
		p.SetState(303)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IReverseContext is an interface to support dynamic dispatch.
type IReverseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReverseContext differentiates from other interfaces.
	IsReverseContext()
}

type ReverseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReverseContext() *ReverseContext {
	var p = new(ReverseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_reverse
	return p
}

func (*ReverseContext) IsReverseContext() {}

func NewReverseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReverseContext {
	var p = new(ReverseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_reverse

	return p
}

func (s *ReverseContext) GetParser() antlr.Parser { return s.parser }

func (s *ReverseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *ReverseContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReverseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *ReverseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReverseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReverseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterReverse(s)
	}
}

func (s *ReverseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitReverse(s)
	}
}

func (p *snobolParser) Reverse() (localctx IReverseContext) {
	this := p
	_ = this

	localctx = NewReverseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, snobolParserRULE_reverse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(305)
		p.Match(snobolParserT__21)
	}
	{
		p.SetState(306)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(307)
		p.Expression()
	}
	{
		p.SetState(308)
		p.Match(snobolParserRPAREN)
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
	p.RuleIndex = snobolParserRULE_date
	return p
}

func (*DateContext) IsDateContext() {}

func NewDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateContext {
	var p = new(DateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_date

	return p
}

func (s *DateContext) GetParser() antlr.Parser { return s.parser }

func (s *DateContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *DateContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *DateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterDate(s)
	}
}

func (s *DateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitDate(s)
	}
}

func (p *snobolParser) Date() (localctx IDateContext) {
	this := p
	_ = this

	localctx = NewDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, snobolParserRULE_date)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(310)
		p.Match(snobolParserT__22)
	}
	{
		p.SetState(311)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(312)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IReplaceContext is an interface to support dynamic dispatch.
type IReplaceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReplaceContext differentiates from other interfaces.
	IsReplaceContext()
}

type ReplaceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReplaceContext() *ReplaceContext {
	var p = new(ReplaceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_replace
	return p
}

func (*ReplaceContext) IsReplaceContext() {}

func NewReplaceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReplaceContext {
	var p = new(ReplaceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_replace

	return p
}

func (s *ReplaceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReplaceContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *ReplaceContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ReplaceContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReplaceContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(snobolParserCOMMA)
}

func (s *ReplaceContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(snobolParserCOMMA, i)
}

func (s *ReplaceContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *ReplaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplaceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReplaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterReplace(s)
	}
}

func (s *ReplaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitReplace(s)
	}
}

func (p *snobolParser) Replace() (localctx IReplaceContext) {
	this := p
	_ = this

	localctx = NewReplaceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, snobolParserRULE_replace)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(314)
		p.Match(snobolParserT__23)
	}
	{
		p.SetState(315)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(316)
		p.Expression()
	}
	{
		p.SetState(317)
		p.Match(snobolParserCOMMA)
	}
	{
		p.SetState(318)
		p.Expression()
	}
	{
		p.SetState(319)
		p.Match(snobolParserCOMMA)
	}
	{
		p.SetState(320)
		p.Expression()
	}
	{
		p.SetState(321)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// ISizeContext is an interface to support dynamic dispatch.
type ISizeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSizeContext differentiates from other interfaces.
	IsSizeContext()
}

type SizeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySizeContext() *SizeContext {
	var p = new(SizeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_size
	return p
}

func (*SizeContext) IsSizeContext() {}

func NewSizeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SizeContext {
	var p = new(SizeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_size

	return p
}

func (s *SizeContext) GetParser() antlr.Parser { return s.parser }

func (s *SizeContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *SizeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SizeContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *SizeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SizeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SizeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterSize(s)
	}
}

func (s *SizeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitSize(s)
	}
}

func (p *snobolParser) Size() (localctx ISizeContext) {
	this := p
	_ = this

	localctx = NewSizeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, snobolParserRULE_size)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.Match(snobolParserT__24)
	}
	{
		p.SetState(324)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(325)
		p.Expression()
	}
	{
		p.SetState(326)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// ITrimContext is an interface to support dynamic dispatch.
type ITrimContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTrimContext differentiates from other interfaces.
	IsTrimContext()
}

type TrimContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrimContext() *TrimContext {
	var p = new(TrimContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_trim
	return p
}

func (*TrimContext) IsTrimContext() {}

func NewTrimContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TrimContext {
	var p = new(TrimContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_trim

	return p
}

func (s *TrimContext) GetParser() antlr.Parser { return s.parser }

func (s *TrimContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *TrimContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TrimContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *TrimContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrimContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TrimContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterTrim(s)
	}
}

func (s *TrimContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitTrim(s)
	}
}

func (p *snobolParser) Trim() (localctx ITrimContext) {
	this := p
	_ = this

	localctx = NewTrimContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, snobolParserRULE_trim)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(snobolParserT__25)
	}
	{
		p.SetState(329)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(330)
		p.Expression()
	}
	{
		p.SetState(331)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IArray_Context is an interface to support dynamic dispatch.
type IArray_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_Context differentiates from other interfaces.
	IsArray_Context()
}

type Array_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_Context() *Array_Context {
	var p = new(Array_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_array_
	return p
}

func (*Array_Context) IsArray_Context() {}

func NewArray_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_Context {
	var p = new(Array_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_array_

	return p
}

func (s *Array_Context) GetParser() antlr.Parser { return s.parser }

func (s *Array_Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *Array_Context) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Array_Context) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Array_Context) COMMA() antlr.TerminalNode {
	return s.GetToken(snobolParserCOMMA, 0)
}

func (s *Array_Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *Array_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterArray_(s)
	}
}

func (s *Array_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitArray_(s)
	}
}

func (p *snobolParser) Array_() (localctx IArray_Context) {
	this := p
	_ = this

	localctx = NewArray_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, snobolParserRULE_array_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(333)
		p.Match(snobolParserT__26)
	}
	{
		p.SetState(334)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(335)
		p.Expression()
	}
	{
		p.SetState(336)
		p.Match(snobolParserCOMMA)
	}
	{
		p.SetState(337)
		p.Expression()
	}
	{
		p.SetState(338)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IConvertContext is an interface to support dynamic dispatch.
type IConvertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConvertContext differentiates from other interfaces.
	IsConvertContext()
}

type ConvertContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConvertContext() *ConvertContext {
	var p = new(ConvertContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_convert
	return p
}

func (*ConvertContext) IsConvertContext() {}

func NewConvertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConvertContext {
	var p = new(ConvertContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_convert

	return p
}

func (s *ConvertContext) GetParser() antlr.Parser { return s.parser }

func (s *ConvertContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *ConvertContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ConvertContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConvertContext) COMMA() antlr.TerminalNode {
	return s.GetToken(snobolParserCOMMA, 0)
}

func (s *ConvertContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *ConvertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConvertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConvertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterConvert(s)
	}
}

func (s *ConvertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitConvert(s)
	}
}

func (p *snobolParser) Convert() (localctx IConvertContext) {
	this := p
	_ = this

	localctx = NewConvertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, snobolParserRULE_convert)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(340)
		p.Match(snobolParserT__27)
	}
	{
		p.SetState(341)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(342)
		p.Expression()
	}
	{
		p.SetState(343)
		p.Match(snobolParserCOMMA)
	}
	{
		p.SetState(344)
		p.Expression()
	}
	{
		p.SetState(345)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// ITableContext is an interface to support dynamic dispatch.
type ITableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableContext differentiates from other interfaces.
	IsTableContext()
}

type TableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableContext() *TableContext {
	var p = new(TableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_table
	return p
}

func (*TableContext) IsTableContext() {}

func NewTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableContext {
	var p = new(TableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_table

	return p
}

func (s *TableContext) GetParser() antlr.Parser { return s.parser }

func (s *TableContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *TableContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TableContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *TableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterTable(s)
	}
}

func (s *TableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitTable(s)
	}
}

func (p *snobolParser) Table() (localctx ITableContext) {
	this := p
	_ = this

	localctx = NewTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, snobolParserRULE_table)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.Match(snobolParserT__28)
	}
	{
		p.SetState(348)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(349)
		p.Expression()
	}
	{
		p.SetState(350)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// ISortContext is an interface to support dynamic dispatch.
type ISortContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSortContext differentiates from other interfaces.
	IsSortContext()
}

type SortContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySortContext() *SortContext {
	var p = new(SortContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_sort
	return p
}

func (*SortContext) IsSortContext() {}

func NewSortContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortContext {
	var p = new(SortContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_sort

	return p
}

func (s *SortContext) GetParser() antlr.Parser { return s.parser }

func (s *SortContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *SortContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SortContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *SortContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterSort(s)
	}
}

func (s *SortContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitSort(s)
	}
}

func (p *snobolParser) Sort() (localctx ISortContext) {
	this := p
	_ = this

	localctx = NewSortContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, snobolParserRULE_sort)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(352)
		p.Match(snobolParserT__29)
	}
	{
		p.SetState(353)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(354)
		p.Expression()
	}
	{
		p.SetState(355)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// IBreak_Context is an interface to support dynamic dispatch.
type IBreak_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreak_Context differentiates from other interfaces.
	IsBreak_Context()
}

type Break_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreak_Context() *Break_Context {
	var p = new(Break_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_break_
	return p
}

func (*Break_Context) IsBreak_Context() {}

func NewBreak_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Break_Context {
	var p = new(Break_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_break_

	return p
}

func (s *Break_Context) GetParser() antlr.Parser { return s.parser }

func (s *Break_Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *Break_Context) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Break_Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *Break_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Break_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Break_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterBreak_(s)
	}
}

func (s *Break_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitBreak_(s)
	}
}

func (p *snobolParser) Break_() (localctx IBreak_Context) {
	this := p
	_ = this

	localctx = NewBreak_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, snobolParserRULE_break_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(357)
		p.Match(snobolParserT__30)
	}
	{
		p.SetState(358)
		p.Match(snobolParserLPAREN)
	}
	{
		p.SetState(359)
		p.Expression()
	}
	{
		p.SetState(360)
		p.Match(snobolParserRPAREN)
	}

	return localctx
}

// ITransferContext is an interface to support dynamic dispatch.
type ITransferContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransferContext differentiates from other interfaces.
	IsTransferContext()
}

type TransferContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransferContext() *TransferContext {
	var p = new(TransferContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_transfer
	return p
}

func (*TransferContext) IsTransferContext() {}

func NewTransferContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransferContext {
	var p = new(TransferContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_transfer

	return p
}

func (s *TransferContext) GetParser() antlr.Parser { return s.parser }

func (s *TransferContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserLPAREN, 0)
}

func (s *TransferContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(snobolParserRPAREN, 0)
}

func (s *TransferContext) Label() ILabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILabelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILabelContext)
}

func (s *TransferContext) END() antlr.TerminalNode {
	return s.GetToken(snobolParserEND, 0)
}

func (s *TransferContext) Transferpre() ITransferpreContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransferpreContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransferpreContext)
}

func (s *TransferContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransferContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterTransfer(s)
	}
}

func (s *TransferContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitTransfer(s)
	}
}

func (p *snobolParser) Transfer() (localctx ITransferContext) {
	this := p
	_ = this

	localctx = NewTransferContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, snobolParserRULE_transfer)
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
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(snobolParserT__31-32))|(1<<(snobolParserT__32-32))|(1<<(snobolParserT__33-32))|(1<<(snobolParserT__34-32))|(1<<(snobolParserLPAREN-32)))) != 0 {
		p.SetState(363)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(snobolParserT__31-32))|(1<<(snobolParserT__32-32))|(1<<(snobolParserT__33-32))|(1<<(snobolParserT__34-32)))) != 0 {
			{
				p.SetState(362)
				p.Transferpre()
			}

		}
		{
			p.SetState(365)
			p.Match(snobolParserLPAREN)
		}
		p.SetState(368)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case snobolParserSTRING:
			{
				p.SetState(366)
				p.Label()
			}

		case snobolParserEND:
			{
				p.SetState(367)
				p.Match(snobolParserEND)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(370)
			p.Match(snobolParserRPAREN)
		}

	}

	return localctx
}

// ITransferpreContext is an interface to support dynamic dispatch.
type ITransferpreContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransferpreContext differentiates from other interfaces.
	IsTransferpreContext()
}

type TransferpreContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransferpreContext() *TransferpreContext {
	var p = new(TransferpreContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snobolParserRULE_transferpre
	return p
}

func (*TransferpreContext) IsTransferpreContext() {}

func NewTransferpreContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransferpreContext {
	var p = new(TransferpreContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snobolParserRULE_transferpre

	return p
}

func (s *TransferpreContext) GetParser() antlr.Parser { return s.parser }
func (s *TransferpreContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferpreContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransferpreContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.EnterTransferpre(s)
	}
}

func (s *TransferpreContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snobolListener); ok {
		listenerT.ExitTransferpre(s)
	}
}

func (p *snobolParser) Transferpre() (localctx ITransferpreContext) {
	this := p
	_ = this

	localctx = NewTransferpreContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, snobolParserRULE_transferpre)
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
		p.SetState(373)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(snobolParserT__31-32))|(1<<(snobolParserT__32-32))|(1<<(snobolParserT__33-32))|(1<<(snobolParserT__34-32)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
