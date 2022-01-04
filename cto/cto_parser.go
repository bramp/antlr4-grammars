// Code generated from CtoParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package cto // CtoParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 388,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 3, 2, 3, 2, 7, 2, 87, 10,
	2, 12, 2, 14, 2, 90, 11, 2, 3, 2, 7, 2, 93, 10, 2, 12, 2, 14, 2, 96, 11,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 107, 10,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 115, 10, 5, 3, 6, 3, 6, 5,
	6, 119, 10, 6, 3, 7, 7, 7, 122, 10, 7, 12, 7, 14, 7, 125, 11, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 7, 8, 133, 10, 8, 12, 8, 14, 8, 136, 11,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 142, 10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 7, 9, 150, 10, 9, 12, 9, 14, 9, 153, 11, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 7, 12, 165, 10, 12, 12,
	12, 14, 12, 168, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 7, 13,
	176, 10, 13, 12, 13, 14, 13, 179, 11, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	14, 3, 14, 3, 14, 5, 14, 188, 10, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16,
	7, 16, 195, 10, 16, 12, 16, 14, 16, 198, 11, 16, 3, 16, 3, 16, 3, 17, 3,
	17, 5, 17, 204, 10, 17, 3, 18, 3, 18, 3, 18, 5, 18, 209, 10, 18, 3, 18,
	5, 18, 212, 10, 18, 3, 18, 5, 18, 215, 10, 18, 3, 18, 3, 18, 3, 18, 5,
	18, 220, 10, 18, 3, 18, 5, 18, 223, 10, 18, 3, 18, 3, 18, 3, 18, 5, 18,
	228, 10, 18, 3, 18, 5, 18, 231, 10, 18, 3, 18, 5, 18, 234, 10, 18, 3, 18,
	3, 18, 3, 18, 5, 18, 239, 10, 18, 3, 18, 5, 18, 242, 10, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 250, 10, 18, 3, 19, 3, 19, 3, 19,
	7, 19, 255, 10, 19, 12, 19, 14, 19, 258, 11, 19, 3, 20, 3, 20, 3, 20, 7,
	20, 263, 10, 20, 12, 20, 14, 20, 266, 11, 20, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 22, 7, 22, 273, 10, 22, 12, 22, 14, 22, 276, 11, 22, 3, 23, 3, 23, 3,
	23, 7, 23, 281, 10, 23, 12, 23, 14, 23, 284, 11, 23, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 7, 26, 297, 10,
	26, 12, 26, 14, 26, 300, 11, 26, 3, 27, 3, 27, 3, 27, 7, 27, 305, 10, 27,
	12, 27, 14, 27, 308, 11, 27, 3, 28, 3, 28, 3, 28, 7, 28, 313, 10, 28, 12,
	28, 14, 28, 316, 11, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 5, 30, 338, 10, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3,
	35, 3, 35, 5, 35, 357, 10, 35, 3, 36, 3, 36, 5, 36, 361, 10, 36, 3, 37,
	3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3,
	40, 5, 40, 375, 10, 40, 3, 41, 3, 41, 3, 41, 7, 41, 380, 10, 41, 12, 41,
	14, 41, 383, 11, 41, 3, 42, 3, 42, 3, 42, 3, 42, 2, 2, 43, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
	46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
	82, 2, 6, 3, 2, 20, 22, 5, 2, 4, 4, 14, 14, 49, 49, 3, 2, 50, 51, 3, 2,
	40, 41, 2, 393, 2, 84, 3, 2, 2, 2, 4, 99, 3, 2, 2, 2, 6, 102, 3, 2, 2,
	2, 8, 114, 3, 2, 2, 2, 10, 118, 3, 2, 2, 2, 12, 123, 3, 2, 2, 2, 14, 134,
	3, 2, 2, 2, 16, 145, 3, 2, 2, 2, 18, 156, 3, 2, 2, 2, 20, 159, 3, 2, 2,
	2, 22, 166, 3, 2, 2, 2, 24, 177, 3, 2, 2, 2, 26, 187, 3, 2, 2, 2, 28, 189,
	3, 2, 2, 2, 30, 192, 3, 2, 2, 2, 32, 203, 3, 2, 2, 2, 34, 249, 3, 2, 2,
	2, 36, 251, 3, 2, 2, 2, 38, 259, 3, 2, 2, 2, 40, 267, 3, 2, 2, 2, 42, 269,
	3, 2, 2, 2, 44, 277, 3, 2, 2, 2, 46, 285, 3, 2, 2, 2, 48, 289, 3, 2, 2,
	2, 50, 293, 3, 2, 2, 2, 52, 301, 3, 2, 2, 2, 54, 309, 3, 2, 2, 2, 56, 317,
	3, 2, 2, 2, 58, 337, 3, 2, 2, 2, 60, 339, 3, 2, 2, 2, 62, 343, 3, 2, 2,
	2, 64, 347, 3, 2, 2, 2, 66, 351, 3, 2, 2, 2, 68, 356, 3, 2, 2, 2, 70, 360,
	3, 2, 2, 2, 72, 362, 3, 2, 2, 2, 74, 364, 3, 2, 2, 2, 76, 366, 3, 2, 2,
	2, 78, 368, 3, 2, 2, 2, 80, 376, 3, 2, 2, 2, 82, 384, 3, 2, 2, 2, 84, 88,
	5, 4, 3, 2, 85, 87, 5, 6, 4, 2, 86, 85, 3, 2, 2, 2, 87, 90, 3, 2, 2, 2,
	88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 94, 3, 2, 2, 2, 90, 88, 3,
	2, 2, 2, 91, 93, 5, 8, 5, 2, 92, 91, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94,
	92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 97, 3, 2, 2, 2, 96, 94, 3, 2, 2,
	2, 97, 98, 7, 2, 2, 3, 98, 3, 3, 2, 2, 2, 99, 100, 7, 12, 2, 2, 100, 101,
	5, 54, 28, 2, 101, 5, 3, 2, 2, 2, 102, 103, 7, 11, 2, 2, 103, 106, 5, 54,
	28, 2, 104, 105, 7, 32, 2, 2, 105, 107, 7, 35, 2, 2, 106, 104, 3, 2, 2,
	2, 106, 107, 3, 2, 2, 2, 107, 7, 3, 2, 2, 2, 108, 115, 5, 12, 7, 2, 109,
	115, 5, 14, 8, 2, 110, 115, 5, 16, 9, 2, 111, 115, 5, 22, 12, 2, 112, 115,
	5, 24, 13, 2, 113, 115, 5, 20, 11, 2, 114, 108, 3, 2, 2, 2, 114, 109, 3,
	2, 2, 2, 114, 110, 3, 2, 2, 2, 114, 111, 3, 2, 2, 2, 114, 112, 3, 2, 2,
	2, 114, 113, 3, 2, 2, 2, 115, 9, 3, 2, 2, 2, 116, 119, 5, 78, 40, 2, 117,
	119, 7, 3, 2, 2, 118, 116, 3, 2, 2, 2, 118, 117, 3, 2, 2, 2, 119, 11, 3,
	2, 2, 2, 120, 122, 5, 10, 6, 2, 121, 120, 3, 2, 2, 2, 122, 125, 3, 2, 2,
	2, 123, 121, 3, 2, 2, 2, 123, 124, 3, 2, 2, 2, 124, 126, 3, 2, 2, 2, 125,
	123, 3, 2, 2, 2, 126, 127, 7, 4, 2, 2, 127, 128, 7, 49, 2, 2, 128, 129,
	5, 26, 14, 2, 129, 130, 5, 30, 16, 2, 130, 13, 3, 2, 2, 2, 131, 133, 5,
	10, 6, 2, 132, 131, 3, 2, 2, 2, 133, 136, 3, 2, 2, 2, 134, 132, 3, 2, 2,
	2, 134, 135, 3, 2, 2, 2, 135, 137, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 137,
	138, 7, 5, 2, 2, 138, 141, 7, 49, 2, 2, 139, 140, 7, 9, 2, 2, 140, 142,
	7, 49, 2, 2, 141, 139, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 143, 3, 2,
	2, 2, 143, 144, 5, 30, 16, 2, 144, 15, 3, 2, 2, 2, 145, 146, 7, 7, 2, 2,
	146, 147, 7, 49, 2, 2, 147, 151, 7, 26, 2, 2, 148, 150, 5, 18, 10, 2, 149,
	148, 3, 2, 2, 2, 150, 153, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152,
	3, 2, 2, 2, 152, 154, 3, 2, 2, 2, 153, 151, 3, 2, 2, 2, 154, 155, 7, 27,
	2, 2, 155, 17, 3, 2, 2, 2, 156, 157, 7, 39, 2, 2, 157, 158, 7, 49, 2, 2,
	158, 19, 3, 2, 2, 2, 159, 160, 7, 8, 2, 2, 160, 161, 7, 49, 2, 2, 161,
	162, 5, 30, 16, 2, 162, 21, 3, 2, 2, 2, 163, 165, 5, 10, 6, 2, 164, 163,
	3, 2, 2, 2, 165, 168, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2, 166, 167, 3, 2,
	2, 2, 167, 169, 3, 2, 2, 2, 168, 166, 3, 2, 2, 2, 169, 170, 7, 14, 2, 2,
	170, 171, 7, 49, 2, 2, 171, 172, 5, 26, 14, 2, 172, 173, 5, 30, 16, 2,
	173, 23, 3, 2, 2, 2, 174, 176, 5, 10, 6, 2, 175, 174, 3, 2, 2, 2, 176,
	179, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 180,
	3, 2, 2, 2, 179, 177, 3, 2, 2, 2, 180, 181, 7, 17, 2, 2, 181, 182, 7, 49,
	2, 2, 182, 183, 5, 30, 16, 2, 183, 25, 3, 2, 2, 2, 184, 185, 7, 9, 2, 2,
	185, 188, 7, 49, 2, 2, 186, 188, 5, 28, 15, 2, 187, 184, 3, 2, 2, 2, 187,
	186, 3, 2, 2, 2, 188, 27, 3, 2, 2, 2, 189, 190, 7, 10, 2, 2, 190, 191,
	7, 49, 2, 2, 191, 29, 3, 2, 2, 2, 192, 196, 7, 26, 2, 2, 193, 195, 5, 32,
	17, 2, 194, 193, 3, 2, 2, 2, 195, 198, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2,
	196, 197, 3, 2, 2, 2, 197, 199, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 199,
	200, 7, 27, 2, 2, 200, 31, 3, 2, 2, 2, 201, 204, 7, 30, 2, 2, 202, 204,
	5, 34, 18, 2, 203, 201, 3, 2, 2, 2, 203, 202, 3, 2, 2, 2, 204, 33, 3, 2,
	2, 2, 205, 206, 5, 50, 26, 2, 206, 208, 5, 66, 34, 2, 207, 209, 5, 62,
	32, 2, 208, 207, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 211, 3, 2, 2, 2,
	210, 212, 5, 48, 25, 2, 211, 210, 3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212,
	214, 3, 2, 2, 2, 213, 215, 7, 13, 2, 2, 214, 213, 3, 2, 2, 2, 214, 215,
	3, 2, 2, 2, 215, 250, 3, 2, 2, 2, 216, 217, 5, 42, 22, 2, 217, 219, 5,
	66, 34, 2, 218, 220, 5, 60, 31, 2, 219, 218, 3, 2, 2, 2, 219, 220, 3, 2,
	2, 2, 220, 222, 3, 2, 2, 2, 221, 223, 7, 13, 2, 2, 222, 221, 3, 2, 2, 2,
	222, 223, 3, 2, 2, 2, 223, 250, 3, 2, 2, 2, 224, 225, 5, 38, 20, 2, 225,
	227, 5, 66, 34, 2, 226, 228, 5, 64, 33, 2, 227, 226, 3, 2, 2, 2, 227, 228,
	3, 2, 2, 2, 228, 230, 3, 2, 2, 2, 229, 231, 5, 56, 29, 2, 230, 229, 3,
	2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 233, 3, 2, 2, 2, 232, 234, 7, 13, 2,
	2, 233, 232, 3, 2, 2, 2, 233, 234, 3, 2, 2, 2, 234, 250, 3, 2, 2, 2, 235,
	236, 5, 44, 23, 2, 236, 238, 5, 66, 34, 2, 237, 239, 5, 46, 24, 2, 238,
	237, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239, 241, 3, 2, 2, 2, 240, 242,
	7, 13, 2, 2, 241, 240, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 250, 3, 2,
	2, 2, 243, 244, 5, 36, 19, 2, 244, 245, 5, 66, 34, 2, 245, 250, 3, 2, 2,
	2, 246, 247, 5, 52, 27, 2, 247, 248, 5, 66, 34, 2, 248, 250, 3, 2, 2, 2,
	249, 205, 3, 2, 2, 2, 249, 216, 3, 2, 2, 2, 249, 224, 3, 2, 2, 2, 249,
	235, 3, 2, 2, 2, 249, 243, 3, 2, 2, 2, 249, 246, 3, 2, 2, 2, 250, 35, 3,
	2, 2, 2, 251, 252, 7, 39, 2, 2, 252, 256, 7, 49, 2, 2, 253, 255, 5, 82,
	42, 2, 254, 253, 3, 2, 2, 2, 255, 258, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2,
	256, 257, 3, 2, 2, 2, 257, 37, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2, 259, 260,
	7, 39, 2, 2, 260, 264, 5, 40, 21, 2, 261, 263, 5, 82, 42, 2, 262, 261,
	3, 2, 2, 2, 263, 266, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 264, 265, 3, 2,
	2, 2, 265, 39, 3, 2, 2, 2, 266, 264, 3, 2, 2, 2, 267, 268, 9, 2, 2, 2,
	268, 41, 3, 2, 2, 2, 269, 270, 7, 39, 2, 2, 270, 274, 7, 18, 2, 2, 271,
	273, 5, 82, 42, 2, 272, 271, 3, 2, 2, 2, 273, 276, 3, 2, 2, 2, 274, 272,
	3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 43, 3, 2, 2, 2, 276, 274, 3, 2,
	2, 2, 277, 278, 7, 39, 2, 2, 278, 282, 7, 19, 2, 2, 279, 281, 5, 82, 42,
	2, 280, 279, 3, 2, 2, 2, 281, 284, 3, 2, 2, 2, 282, 280, 3, 2, 2, 2, 282,
	283, 3, 2, 2, 2, 283, 45, 3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 285, 286, 7,
	6, 2, 2, 286, 287, 7, 34, 2, 2, 287, 288, 7, 44, 2, 2, 288, 47, 3, 2, 2,
	2, 289, 290, 7, 16, 2, 2, 290, 291, 7, 34, 2, 2, 291, 292, 7, 48, 2, 2,
	292, 49, 3, 2, 2, 2, 293, 294, 7, 39, 2, 2, 294, 298, 7, 23, 2, 2, 295,
	297, 5, 82, 42, 2, 296, 295, 3, 2, 2, 2, 297, 300, 3, 2, 2, 2, 298, 296,
	3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 51, 3, 2, 2, 2, 300, 298, 3, 2,
	2, 2, 301, 302, 7, 38, 2, 2, 302, 306, 7, 49, 2, 2, 303, 305, 5, 82, 42,
	2, 304, 303, 3, 2, 2, 2, 305, 308, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306,
	307, 3, 2, 2, 2, 307, 53, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 309, 314, 7,
	49, 2, 2, 310, 311, 7, 32, 2, 2, 311, 313, 7, 49, 2, 2, 312, 310, 3, 2,
	2, 2, 313, 316, 3, 2, 2, 2, 314, 312, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2,
	315, 55, 3, 2, 2, 2, 316, 314, 3, 2, 2, 2, 317, 318, 7, 15, 2, 2, 318,
	319, 7, 34, 2, 2, 319, 320, 5, 58, 30, 2, 320, 57, 3, 2, 2, 2, 321, 322,
	7, 28, 2, 2, 322, 323, 5, 70, 36, 2, 323, 324, 7, 31, 2, 2, 324, 325, 7,
	29, 2, 2, 325, 338, 3, 2, 2, 2, 326, 327, 7, 28, 2, 2, 327, 328, 7, 31,
	2, 2, 328, 329, 5, 70, 36, 2, 329, 330, 7, 29, 2, 2, 330, 338, 3, 2, 2,
	2, 331, 332, 7, 28, 2, 2, 332, 333, 5, 70, 36, 2, 333, 334, 7, 31, 2, 2,
	334, 335, 5, 70, 36, 2, 335, 336, 7, 29, 2, 2, 336, 338, 3, 2, 2, 2, 337,
	321, 3, 2, 2, 2, 337, 326, 3, 2, 2, 2, 337, 331, 3, 2, 2, 2, 338, 59, 3,
	2, 2, 2, 339, 340, 7, 6, 2, 2, 340, 341, 7, 34, 2, 2, 341, 342, 7, 43,
	2, 2, 342, 61, 3, 2, 2, 2, 343, 344, 7, 6, 2, 2, 344, 345, 7, 34, 2, 2,
	345, 346, 5, 72, 37, 2, 346, 63, 3, 2, 2, 2, 347, 348, 7, 6, 2, 2, 348,
	349, 7, 34, 2, 2, 349, 350, 5, 70, 36, 2, 350, 65, 3, 2, 2, 2, 351, 352,
	9, 3, 2, 2, 352, 67, 3, 2, 2, 2, 353, 357, 5, 70, 36, 2, 354, 357, 5, 72,
	37, 2, 355, 357, 7, 43, 2, 2, 356, 353, 3, 2, 2, 2, 356, 354, 3, 2, 2,
	2, 356, 355, 3, 2, 2, 2, 357, 69, 3, 2, 2, 2, 358, 361, 5, 74, 38, 2, 359,
	361, 5, 76, 39, 2, 360, 358, 3, 2, 2, 2, 360, 359, 3, 2, 2, 2, 361, 71,
	3, 2, 2, 2, 362, 363, 9, 4, 2, 2, 363, 73, 3, 2, 2, 2, 364, 365, 9, 5,
	2, 2, 365, 75, 3, 2, 2, 2, 366, 367, 7, 42, 2, 2, 367, 77, 3, 2, 2, 2,
	368, 369, 7, 36, 2, 2, 369, 374, 5, 54, 28, 2, 370, 371, 7, 24, 2, 2, 371,
	372, 5, 80, 41, 2, 372, 373, 7, 25, 2, 2, 373, 375, 3, 2, 2, 2, 374, 370,
	3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375, 79, 3, 2, 2, 2, 376, 381, 5, 68,
	35, 2, 377, 378, 7, 31, 2, 2, 378, 380, 5, 68, 35, 2, 379, 377, 3, 2, 2,
	2, 380, 383, 3, 2, 2, 2, 381, 379, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382,
	81, 3, 2, 2, 2, 383, 381, 3, 2, 2, 2, 384, 385, 7, 28, 2, 2, 385, 386,
	7, 29, 2, 2, 386, 83, 3, 2, 2, 2, 39, 88, 94, 106, 114, 118, 123, 134,
	141, 151, 166, 177, 187, 196, 203, 208, 211, 214, 219, 222, 227, 230, 233,
	238, 241, 249, 256, 264, 274, 282, 298, 306, 314, 337, 356, 360, 374, 381,
}
var literalNames = []string{
	"", "'abstract'", "'asset'", "'concept'", "'default'", "'enum'", "'event'",
	"'extends'", "'identified by'", "'import'", "'namespace'", "'optional'",
	"'participant'", "'range'", "'regex'", "'transaction'", "'Boolean'", "'DateTime'",
	"'Double'", "'Integer'", "'Long'", "'String'", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "';'", "','", "'.'", "':'", "'='", "'*'", "'@'", "'...'",
	"'--> '", "'o '",
}
var symbolicNames = []string{
	"", "ABSTRACT", "ASSET", "CONCEPT", "DEFAULT", "ENUM", "EVENT", "EXTENDS",
	"IDENTIFIED", "IMPORT", "NAMESPACE", "OPTIONAL", "PARTICIPANT", "RANGE",
	"REGEX", "TRANSACTION", "BOOLEAN", "DATE_TIME", "DOUBLE", "INTEGER", "LONG",
	"STRING", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", "RBRACK", "SEMI",
	"COMMA", "DOT", "COLON", "ASSIGN", "MUL", "AT", "ELLIPSIS", "REF", "VAR",
	"DECIMAL_LITERAL", "OCT_LITERAL", "FLOAT_LITERAL", "BOOL_LITERAL", "DATE_TIME_LITERAL",
	"WS", "LINE_COMMENT", "COMMENT", "REGEX_EXPR", "IDENTIFIER", "CHAR_LITERAL",
	"STRING_LITERAL",
}

var ruleNames = []string{
	"modelUnit", "namespaceDeclaration", "importDeclaration", "typeDeclaration",
	"classModifier", "assetDeclaration", "conceptDeclaration", "enumDeclaration",
	"enumConstant", "eventDeclaration", "participantDeclaration", "transactionDeclaration",
	"extendsOrIdentified", "identified", "classBody", "classBodyDeclaration",
	"fieldDeclaration", "identifierField", "numericField", "numericPrimitive",
	"booleanField", "dateField", "defaultDate", "regexDeclaration", "stringField",
	"reference", "qualifiedName", "rangeValidation", "rangeDeclaration", "defaultBoolean",
	"defaultString", "defaultNumber", "identifier", "literal", "numberLiteral",
	"stringLiteral", "integerLiteral", "floatLiteral", "decorator", "elementValuePair",
	"square",
}

type CtoParser struct {
	*antlr.BaseParser
}

// NewCtoParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *CtoParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCtoParser(input antlr.TokenStream) *CtoParser {
	this := new(CtoParser)
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
	this.GrammarFileName = "CtoParser.g4"

	return this
}

// CtoParser tokens.
const (
	CtoParserEOF               = antlr.TokenEOF
	CtoParserABSTRACT          = 1
	CtoParserASSET             = 2
	CtoParserCONCEPT           = 3
	CtoParserDEFAULT           = 4
	CtoParserENUM              = 5
	CtoParserEVENT             = 6
	CtoParserEXTENDS           = 7
	CtoParserIDENTIFIED        = 8
	CtoParserIMPORT            = 9
	CtoParserNAMESPACE         = 10
	CtoParserOPTIONAL          = 11
	CtoParserPARTICIPANT       = 12
	CtoParserRANGE             = 13
	CtoParserREGEX             = 14
	CtoParserTRANSACTION       = 15
	CtoParserBOOLEAN           = 16
	CtoParserDATE_TIME         = 17
	CtoParserDOUBLE            = 18
	CtoParserINTEGER           = 19
	CtoParserLONG              = 20
	CtoParserSTRING            = 21
	CtoParserLPAREN            = 22
	CtoParserRPAREN            = 23
	CtoParserLBRACE            = 24
	CtoParserRBRACE            = 25
	CtoParserLBRACK            = 26
	CtoParserRBRACK            = 27
	CtoParserSEMI              = 28
	CtoParserCOMMA             = 29
	CtoParserDOT               = 30
	CtoParserCOLON             = 31
	CtoParserASSIGN            = 32
	CtoParserMUL               = 33
	CtoParserAT                = 34
	CtoParserELLIPSIS          = 35
	CtoParserREF               = 36
	CtoParserVAR               = 37
	CtoParserDECIMAL_LITERAL   = 38
	CtoParserOCT_LITERAL       = 39
	CtoParserFLOAT_LITERAL     = 40
	CtoParserBOOL_LITERAL      = 41
	CtoParserDATE_TIME_LITERAL = 42
	CtoParserWS                = 43
	CtoParserLINE_COMMENT      = 44
	CtoParserCOMMENT           = 45
	CtoParserREGEX_EXPR        = 46
	CtoParserIDENTIFIER        = 47
	CtoParserCHAR_LITERAL      = 48
	CtoParserSTRING_LITERAL    = 49
)

// CtoParser rules.
const (
	CtoParserRULE_modelUnit              = 0
	CtoParserRULE_namespaceDeclaration   = 1
	CtoParserRULE_importDeclaration      = 2
	CtoParserRULE_typeDeclaration        = 3
	CtoParserRULE_classModifier          = 4
	CtoParserRULE_assetDeclaration       = 5
	CtoParserRULE_conceptDeclaration     = 6
	CtoParserRULE_enumDeclaration        = 7
	CtoParserRULE_enumConstant           = 8
	CtoParserRULE_eventDeclaration       = 9
	CtoParserRULE_participantDeclaration = 10
	CtoParserRULE_transactionDeclaration = 11
	CtoParserRULE_extendsOrIdentified    = 12
	CtoParserRULE_identified             = 13
	CtoParserRULE_classBody              = 14
	CtoParserRULE_classBodyDeclaration   = 15
	CtoParserRULE_fieldDeclaration       = 16
	CtoParserRULE_identifierField        = 17
	CtoParserRULE_numericField           = 18
	CtoParserRULE_numericPrimitive       = 19
	CtoParserRULE_booleanField           = 20
	CtoParserRULE_dateField              = 21
	CtoParserRULE_defaultDate            = 22
	CtoParserRULE_regexDeclaration       = 23
	CtoParserRULE_stringField            = 24
	CtoParserRULE_reference              = 25
	CtoParserRULE_qualifiedName          = 26
	CtoParserRULE_rangeValidation        = 27
	CtoParserRULE_rangeDeclaration       = 28
	CtoParserRULE_defaultBoolean         = 29
	CtoParserRULE_defaultString          = 30
	CtoParserRULE_defaultNumber          = 31
	CtoParserRULE_identifier             = 32
	CtoParserRULE_literal                = 33
	CtoParserRULE_numberLiteral          = 34
	CtoParserRULE_stringLiteral          = 35
	CtoParserRULE_integerLiteral         = 36
	CtoParserRULE_floatLiteral           = 37
	CtoParserRULE_decorator              = 38
	CtoParserRULE_elementValuePair       = 39
	CtoParserRULE_square                 = 40
)

// IModelUnitContext is an interface to support dynamic dispatch.
type IModelUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModelUnitContext differentiates from other interfaces.
	IsModelUnitContext()
}

type ModelUnitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModelUnitContext() *ModelUnitContext {
	var p = new(ModelUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_modelUnit
	return p
}

func (*ModelUnitContext) IsModelUnitContext() {}

func NewModelUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModelUnitContext {
	var p = new(ModelUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_modelUnit

	return p
}

func (s *ModelUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *ModelUnitContext) NamespaceDeclaration() INamespaceDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespaceDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespaceDeclarationContext)
}

func (s *ModelUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(CtoParserEOF, 0)
}

func (s *ModelUnitContext) AllImportDeclaration() []IImportDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem())
	var tst = make([]IImportDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportDeclarationContext)
		}
	}

	return tst
}

func (s *ModelUnitContext) ImportDeclaration(i int) IImportDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportDeclarationContext)
}

func (s *ModelUnitContext) AllTypeDeclaration() []ITypeDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem())
	var tst = make([]ITypeDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeDeclarationContext)
		}
	}

	return tst
}

func (s *ModelUnitContext) TypeDeclaration(i int) ITypeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *ModelUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModelUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModelUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterModelUnit(s)
	}
}

func (s *ModelUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitModelUnit(s)
	}
}

func (p *CtoParser) ModelUnit() (localctx IModelUnitContext) {
	this := p
	_ = this

	localctx = NewModelUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CtoParserRULE_modelUnit)
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
		p.SetState(82)
		p.NamespaceDeclaration()
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CtoParserIMPORT {
		{
			p.SetState(83)
			p.ImportDeclaration()
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CtoParserABSTRACT)|(1<<CtoParserASSET)|(1<<CtoParserCONCEPT)|(1<<CtoParserENUM)|(1<<CtoParserEVENT)|(1<<CtoParserPARTICIPANT)|(1<<CtoParserTRANSACTION))) != 0) || _la == CtoParserAT {
		{
			p.SetState(89)
			p.TypeDeclaration()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(95)
		p.Match(CtoParserEOF)
	}

	return localctx
}

// INamespaceDeclarationContext is an interface to support dynamic dispatch.
type INamespaceDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespaceDeclarationContext differentiates from other interfaces.
	IsNamespaceDeclarationContext()
}

type NamespaceDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespaceDeclarationContext() *NamespaceDeclarationContext {
	var p = new(NamespaceDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_namespaceDeclaration
	return p
}

func (*NamespaceDeclarationContext) IsNamespaceDeclarationContext() {}

func NewNamespaceDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamespaceDeclarationContext {
	var p = new(NamespaceDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_namespaceDeclaration

	return p
}

func (s *NamespaceDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *NamespaceDeclarationContext) NAMESPACE() antlr.TerminalNode {
	return s.GetToken(CtoParserNAMESPACE, 0)
}

func (s *NamespaceDeclarationContext) QualifiedName() IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *NamespaceDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamespaceDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamespaceDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterNamespaceDeclaration(s)
	}
}

func (s *NamespaceDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitNamespaceDeclaration(s)
	}
}

func (p *CtoParser) NamespaceDeclaration() (localctx INamespaceDeclarationContext) {
	this := p
	_ = this

	localctx = NewNamespaceDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CtoParserRULE_namespaceDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(CtoParserNAMESPACE)
	}
	{
		p.SetState(98)
		p.QualifiedName()
	}

	return localctx
}

// IImportDeclarationContext is an interface to support dynamic dispatch.
type IImportDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclarationContext differentiates from other interfaces.
	IsImportDeclarationContext()
}

type ImportDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclarationContext() *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_importDeclaration
	return p
}

func (*ImportDeclarationContext) IsImportDeclarationContext() {}

func NewImportDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_importDeclaration

	return p
}

func (s *ImportDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclarationContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(CtoParserIMPORT, 0)
}

func (s *ImportDeclarationContext) QualifiedName() IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *ImportDeclarationContext) DOT() antlr.TerminalNode {
	return s.GetToken(CtoParserDOT, 0)
}

func (s *ImportDeclarationContext) MUL() antlr.TerminalNode {
	return s.GetToken(CtoParserMUL, 0)
}

func (s *ImportDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitImportDeclaration(s)
	}
}

func (p *CtoParser) ImportDeclaration() (localctx IImportDeclarationContext) {
	this := p
	_ = this

	localctx = NewImportDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CtoParserRULE_importDeclaration)
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
		p.SetState(100)
		p.Match(CtoParserIMPORT)
	}
	{
		p.SetState(101)
		p.QualifiedName()
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CtoParserDOT {
		{
			p.SetState(102)
			p.Match(CtoParserDOT)
		}
		{
			p.SetState(103)
			p.Match(CtoParserMUL)
		}

	}

	return localctx
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) AssetDeclaration() IAssetDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssetDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssetDeclarationContext)
}

func (s *TypeDeclarationContext) ConceptDeclaration() IConceptDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConceptDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConceptDeclarationContext)
}

func (s *TypeDeclarationContext) EnumDeclaration() IEnumDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumDeclarationContext)
}

func (s *TypeDeclarationContext) ParticipantDeclaration() IParticipantDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParticipantDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParticipantDeclarationContext)
}

func (s *TypeDeclarationContext) TransactionDeclaration() ITransactionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITransactionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITransactionDeclarationContext)
}

func (s *TypeDeclarationContext) EventDeclaration() IEventDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEventDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEventDeclarationContext)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (p *CtoParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	this := p
	_ = this

	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CtoParserRULE_typeDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(106)
			p.AssetDeclaration()
		}

	case 2:
		{
			p.SetState(107)
			p.ConceptDeclaration()
		}

	case 3:
		{
			p.SetState(108)
			p.EnumDeclaration()
		}

	case 4:
		{
			p.SetState(109)
			p.ParticipantDeclaration()
		}

	case 5:
		{
			p.SetState(110)
			p.TransactionDeclaration()
		}

	case 6:
		{
			p.SetState(111)
			p.EventDeclaration()
		}

	}

	return localctx
}

// IClassModifierContext is an interface to support dynamic dispatch.
type IClassModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassModifierContext differentiates from other interfaces.
	IsClassModifierContext()
}

type ClassModifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassModifierContext() *ClassModifierContext {
	var p = new(ClassModifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_classModifier
	return p
}

func (*ClassModifierContext) IsClassModifierContext() {}

func NewClassModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassModifierContext {
	var p = new(ClassModifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_classModifier

	return p
}

func (s *ClassModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassModifierContext) Decorator() IDecoratorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecoratorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecoratorContext)
}

func (s *ClassModifierContext) ABSTRACT() antlr.TerminalNode {
	return s.GetToken(CtoParserABSTRACT, 0)
}

func (s *ClassModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterClassModifier(s)
	}
}

func (s *ClassModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitClassModifier(s)
	}
}

func (p *CtoParser) ClassModifier() (localctx IClassModifierContext) {
	this := p
	_ = this

	localctx = NewClassModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CtoParserRULE_classModifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	case CtoParserAT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Decorator()
		}

	case CtoParserABSTRACT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(115)
			p.Match(CtoParserABSTRACT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssetDeclarationContext is an interface to support dynamic dispatch.
type IAssetDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssetDeclarationContext differentiates from other interfaces.
	IsAssetDeclarationContext()
}

type AssetDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssetDeclarationContext() *AssetDeclarationContext {
	var p = new(AssetDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_assetDeclaration
	return p
}

func (*AssetDeclarationContext) IsAssetDeclarationContext() {}

func NewAssetDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssetDeclarationContext {
	var p = new(AssetDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_assetDeclaration

	return p
}

func (s *AssetDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *AssetDeclarationContext) ASSET() antlr.TerminalNode {
	return s.GetToken(CtoParserASSET, 0)
}

func (s *AssetDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, 0)
}

func (s *AssetDeclarationContext) ExtendsOrIdentified() IExtendsOrIdentifiedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendsOrIdentifiedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtendsOrIdentifiedContext)
}

func (s *AssetDeclarationContext) ClassBody() IClassBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassBodyContext)
}

func (s *AssetDeclarationContext) AllClassModifier() []IClassModifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassModifierContext)(nil)).Elem())
	var tst = make([]IClassModifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassModifierContext)
		}
	}

	return tst
}

func (s *AssetDeclarationContext) ClassModifier(i int) IClassModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassModifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassModifierContext)
}

func (s *AssetDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssetDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssetDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterAssetDeclaration(s)
	}
}

func (s *AssetDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitAssetDeclaration(s)
	}
}

func (p *CtoParser) AssetDeclaration() (localctx IAssetDeclarationContext) {
	this := p
	_ = this

	localctx = NewAssetDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CtoParserRULE_assetDeclaration)
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

	for _la == CtoParserABSTRACT || _la == CtoParserAT {
		{
			p.SetState(118)
			p.ClassModifier()
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(124)
		p.Match(CtoParserASSET)
	}
	{
		p.SetState(125)
		p.Match(CtoParserIDENTIFIER)
	}
	{
		p.SetState(126)
		p.ExtendsOrIdentified()
	}
	{
		p.SetState(127)
		p.ClassBody()
	}

	return localctx
}

// IConceptDeclarationContext is an interface to support dynamic dispatch.
type IConceptDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConceptDeclarationContext differentiates from other interfaces.
	IsConceptDeclarationContext()
}

type ConceptDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConceptDeclarationContext() *ConceptDeclarationContext {
	var p = new(ConceptDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_conceptDeclaration
	return p
}

func (*ConceptDeclarationContext) IsConceptDeclarationContext() {}

func NewConceptDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConceptDeclarationContext {
	var p = new(ConceptDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_conceptDeclaration

	return p
}

func (s *ConceptDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ConceptDeclarationContext) CONCEPT() antlr.TerminalNode {
	return s.GetToken(CtoParserCONCEPT, 0)
}

func (s *ConceptDeclarationContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(CtoParserIDENTIFIER)
}

func (s *ConceptDeclarationContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, i)
}

func (s *ConceptDeclarationContext) ClassBody() IClassBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassBodyContext)
}

func (s *ConceptDeclarationContext) AllClassModifier() []IClassModifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassModifierContext)(nil)).Elem())
	var tst = make([]IClassModifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassModifierContext)
		}
	}

	return tst
}

func (s *ConceptDeclarationContext) ClassModifier(i int) IClassModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassModifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassModifierContext)
}

func (s *ConceptDeclarationContext) EXTENDS() antlr.TerminalNode {
	return s.GetToken(CtoParserEXTENDS, 0)
}

func (s *ConceptDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConceptDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConceptDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterConceptDeclaration(s)
	}
}

func (s *ConceptDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitConceptDeclaration(s)
	}
}

func (p *CtoParser) ConceptDeclaration() (localctx IConceptDeclarationContext) {
	this := p
	_ = this

	localctx = NewConceptDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, CtoParserRULE_conceptDeclaration)
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
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CtoParserABSTRACT || _la == CtoParserAT {
		{
			p.SetState(129)
			p.ClassModifier()
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(135)
		p.Match(CtoParserCONCEPT)
	}
	{
		p.SetState(136)
		p.Match(CtoParserIDENTIFIER)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CtoParserEXTENDS {
		{
			p.SetState(137)
			p.Match(CtoParserEXTENDS)
		}
		{
			p.SetState(138)
			p.Match(CtoParserIDENTIFIER)
		}

	}
	{
		p.SetState(141)
		p.ClassBody()
	}

	return localctx
}

// IEnumDeclarationContext is an interface to support dynamic dispatch.
type IEnumDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumDeclarationContext differentiates from other interfaces.
	IsEnumDeclarationContext()
}

type EnumDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDeclarationContext() *EnumDeclarationContext {
	var p = new(EnumDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_enumDeclaration
	return p
}

func (*EnumDeclarationContext) IsEnumDeclarationContext() {}

func NewEnumDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDeclarationContext {
	var p = new(EnumDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_enumDeclaration

	return p
}

func (s *EnumDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDeclarationContext) ENUM() antlr.TerminalNode {
	return s.GetToken(CtoParserENUM, 0)
}

func (s *EnumDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, 0)
}

func (s *EnumDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(CtoParserLBRACE, 0)
}

func (s *EnumDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(CtoParserRBRACE, 0)
}

func (s *EnumDeclarationContext) AllEnumConstant() []IEnumConstantContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumConstantContext)(nil)).Elem())
	var tst = make([]IEnumConstantContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumConstantContext)
		}
	}

	return tst
}

func (s *EnumDeclarationContext) EnumConstant(i int) IEnumConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumConstantContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumConstantContext)
}

func (s *EnumDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterEnumDeclaration(s)
	}
}

func (s *EnumDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitEnumDeclaration(s)
	}
}

func (p *CtoParser) EnumDeclaration() (localctx IEnumDeclarationContext) {
	this := p
	_ = this

	localctx = NewEnumDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, CtoParserRULE_enumDeclaration)
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
		p.SetState(143)
		p.Match(CtoParserENUM)
	}
	{
		p.SetState(144)
		p.Match(CtoParserIDENTIFIER)
	}
	{
		p.SetState(145)
		p.Match(CtoParserLBRACE)
	}
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CtoParserVAR {
		{
			p.SetState(146)
			p.EnumConstant()
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(152)
		p.Match(CtoParserRBRACE)
	}

	return localctx
}

// IEnumConstantContext is an interface to support dynamic dispatch.
type IEnumConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumConstantContext differentiates from other interfaces.
	IsEnumConstantContext()
}

type EnumConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumConstantContext() *EnumConstantContext {
	var p = new(EnumConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_enumConstant
	return p
}

func (*EnumConstantContext) IsEnumConstantContext() {}

func NewEnumConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumConstantContext {
	var p = new(EnumConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_enumConstant

	return p
}

func (s *EnumConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumConstantContext) VAR() antlr.TerminalNode {
	return s.GetToken(CtoParserVAR, 0)
}

func (s *EnumConstantContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, 0)
}

func (s *EnumConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterEnumConstant(s)
	}
}

func (s *EnumConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitEnumConstant(s)
	}
}

func (p *CtoParser) EnumConstant() (localctx IEnumConstantContext) {
	this := p
	_ = this

	localctx = NewEnumConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, CtoParserRULE_enumConstant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(CtoParserVAR)
	}
	{
		p.SetState(155)
		p.Match(CtoParserIDENTIFIER)
	}

	return localctx
}

// IEventDeclarationContext is an interface to support dynamic dispatch.
type IEventDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEventDeclarationContext differentiates from other interfaces.
	IsEventDeclarationContext()
}

type EventDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEventDeclarationContext() *EventDeclarationContext {
	var p = new(EventDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_eventDeclaration
	return p
}

func (*EventDeclarationContext) IsEventDeclarationContext() {}

func NewEventDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EventDeclarationContext {
	var p = new(EventDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_eventDeclaration

	return p
}

func (s *EventDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *EventDeclarationContext) EVENT() antlr.TerminalNode {
	return s.GetToken(CtoParserEVENT, 0)
}

func (s *EventDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, 0)
}

func (s *EventDeclarationContext) ClassBody() IClassBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassBodyContext)
}

func (s *EventDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EventDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EventDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterEventDeclaration(s)
	}
}

func (s *EventDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitEventDeclaration(s)
	}
}

func (p *CtoParser) EventDeclaration() (localctx IEventDeclarationContext) {
	this := p
	_ = this

	localctx = NewEventDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, CtoParserRULE_eventDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(157)
		p.Match(CtoParserEVENT)
	}
	{
		p.SetState(158)
		p.Match(CtoParserIDENTIFIER)
	}
	{
		p.SetState(159)
		p.ClassBody()
	}

	return localctx
}

// IParticipantDeclarationContext is an interface to support dynamic dispatch.
type IParticipantDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParticipantDeclarationContext differentiates from other interfaces.
	IsParticipantDeclarationContext()
}

type ParticipantDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParticipantDeclarationContext() *ParticipantDeclarationContext {
	var p = new(ParticipantDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_participantDeclaration
	return p
}

func (*ParticipantDeclarationContext) IsParticipantDeclarationContext() {}

func NewParticipantDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParticipantDeclarationContext {
	var p = new(ParticipantDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_participantDeclaration

	return p
}

func (s *ParticipantDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ParticipantDeclarationContext) PARTICIPANT() antlr.TerminalNode {
	return s.GetToken(CtoParserPARTICIPANT, 0)
}

func (s *ParticipantDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, 0)
}

func (s *ParticipantDeclarationContext) ExtendsOrIdentified() IExtendsOrIdentifiedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendsOrIdentifiedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtendsOrIdentifiedContext)
}

func (s *ParticipantDeclarationContext) ClassBody() IClassBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassBodyContext)
}

func (s *ParticipantDeclarationContext) AllClassModifier() []IClassModifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassModifierContext)(nil)).Elem())
	var tst = make([]IClassModifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassModifierContext)
		}
	}

	return tst
}

func (s *ParticipantDeclarationContext) ClassModifier(i int) IClassModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassModifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassModifierContext)
}

func (s *ParticipantDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParticipantDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParticipantDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterParticipantDeclaration(s)
	}
}

func (s *ParticipantDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitParticipantDeclaration(s)
	}
}

func (p *CtoParser) ParticipantDeclaration() (localctx IParticipantDeclarationContext) {
	this := p
	_ = this

	localctx = NewParticipantDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, CtoParserRULE_participantDeclaration)
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
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CtoParserABSTRACT || _la == CtoParserAT {
		{
			p.SetState(161)
			p.ClassModifier()
		}

		p.SetState(166)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(167)
		p.Match(CtoParserPARTICIPANT)
	}
	{
		p.SetState(168)
		p.Match(CtoParserIDENTIFIER)
	}
	{
		p.SetState(169)
		p.ExtendsOrIdentified()
	}
	{
		p.SetState(170)
		p.ClassBody()
	}

	return localctx
}

// ITransactionDeclarationContext is an interface to support dynamic dispatch.
type ITransactionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTransactionDeclarationContext differentiates from other interfaces.
	IsTransactionDeclarationContext()
}

type TransactionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTransactionDeclarationContext() *TransactionDeclarationContext {
	var p = new(TransactionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_transactionDeclaration
	return p
}

func (*TransactionDeclarationContext) IsTransactionDeclarationContext() {}

func NewTransactionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TransactionDeclarationContext {
	var p = new(TransactionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_transactionDeclaration

	return p
}

func (s *TransactionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TransactionDeclarationContext) TRANSACTION() antlr.TerminalNode {
	return s.GetToken(CtoParserTRANSACTION, 0)
}

func (s *TransactionDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, 0)
}

func (s *TransactionDeclarationContext) ClassBody() IClassBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClassBodyContext)
}

func (s *TransactionDeclarationContext) AllClassModifier() []IClassModifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassModifierContext)(nil)).Elem())
	var tst = make([]IClassModifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassModifierContext)
		}
	}

	return tst
}

func (s *TransactionDeclarationContext) ClassModifier(i int) IClassModifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassModifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassModifierContext)
}

func (s *TransactionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransactionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TransactionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterTransactionDeclaration(s)
	}
}

func (s *TransactionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitTransactionDeclaration(s)
	}
}

func (p *CtoParser) TransactionDeclaration() (localctx ITransactionDeclarationContext) {
	this := p
	_ = this

	localctx = NewTransactionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, CtoParserRULE_transactionDeclaration)
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
	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CtoParserABSTRACT || _la == CtoParserAT {
		{
			p.SetState(172)
			p.ClassModifier()
		}

		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(178)
		p.Match(CtoParserTRANSACTION)
	}
	{
		p.SetState(179)
		p.Match(CtoParserIDENTIFIER)
	}
	{
		p.SetState(180)
		p.ClassBody()
	}

	return localctx
}

// IExtendsOrIdentifiedContext is an interface to support dynamic dispatch.
type IExtendsOrIdentifiedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtendsOrIdentifiedContext differentiates from other interfaces.
	IsExtendsOrIdentifiedContext()
}

type ExtendsOrIdentifiedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendsOrIdentifiedContext() *ExtendsOrIdentifiedContext {
	var p = new(ExtendsOrIdentifiedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_extendsOrIdentified
	return p
}

func (*ExtendsOrIdentifiedContext) IsExtendsOrIdentifiedContext() {}

func NewExtendsOrIdentifiedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendsOrIdentifiedContext {
	var p = new(ExtendsOrIdentifiedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_extendsOrIdentified

	return p
}

func (s *ExtendsOrIdentifiedContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendsOrIdentifiedContext) Identified() IIdentifiedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifiedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifiedContext)
}

func (s *ExtendsOrIdentifiedContext) EXTENDS() antlr.TerminalNode {
	return s.GetToken(CtoParserEXTENDS, 0)
}

func (s *ExtendsOrIdentifiedContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, 0)
}

func (s *ExtendsOrIdentifiedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendsOrIdentifiedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendsOrIdentifiedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterExtendsOrIdentified(s)
	}
}

func (s *ExtendsOrIdentifiedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitExtendsOrIdentified(s)
	}
}

func (p *CtoParser) ExtendsOrIdentified() (localctx IExtendsOrIdentifiedContext) {
	this := p
	_ = this

	localctx = NewExtendsOrIdentifiedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, CtoParserRULE_extendsOrIdentified)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(185)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CtoParserEXTENDS:
		{
			p.SetState(182)
			p.Match(CtoParserEXTENDS)
		}
		{
			p.SetState(183)
			p.Match(CtoParserIDENTIFIER)
		}

	case CtoParserIDENTIFIED:
		{
			p.SetState(184)
			p.Identified()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifiedContext is an interface to support dynamic dispatch.
type IIdentifiedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifiedContext differentiates from other interfaces.
	IsIdentifiedContext()
}

type IdentifiedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifiedContext() *IdentifiedContext {
	var p = new(IdentifiedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_identified
	return p
}

func (*IdentifiedContext) IsIdentifiedContext() {}

func NewIdentifiedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifiedContext {
	var p = new(IdentifiedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_identified

	return p
}

func (s *IdentifiedContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifiedContext) IDENTIFIED() antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIED, 0)
}

func (s *IdentifiedContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, 0)
}

func (s *IdentifiedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifiedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifiedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterIdentified(s)
	}
}

func (s *IdentifiedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitIdentified(s)
	}
}

func (p *CtoParser) Identified() (localctx IIdentifiedContext) {
	this := p
	_ = this

	localctx = NewIdentifiedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, CtoParserRULE_identified)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(CtoParserIDENTIFIED)
	}
	{
		p.SetState(188)
		p.Match(CtoParserIDENTIFIER)
	}

	return localctx
}

// IClassBodyContext is an interface to support dynamic dispatch.
type IClassBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassBodyContext differentiates from other interfaces.
	IsClassBodyContext()
}

type ClassBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBodyContext() *ClassBodyContext {
	var p = new(ClassBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_classBody
	return p
}

func (*ClassBodyContext) IsClassBodyContext() {}

func NewClassBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBodyContext {
	var p = new(ClassBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_classBody

	return p
}

func (s *ClassBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(CtoParserLBRACE, 0)
}

func (s *ClassBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(CtoParserRBRACE, 0)
}

func (s *ClassBodyContext) AllClassBodyDeclaration() []IClassBodyDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IClassBodyDeclarationContext)(nil)).Elem())
	var tst = make([]IClassBodyDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IClassBodyDeclarationContext)
		}
	}

	return tst
}

func (s *ClassBodyContext) ClassBodyDeclaration(i int) IClassBodyDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClassBodyDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IClassBodyDeclarationContext)
}

func (s *ClassBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterClassBody(s)
	}
}

func (s *ClassBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitClassBody(s)
	}
}

func (p *CtoParser) ClassBody() (localctx IClassBodyContext) {
	this := p
	_ = this

	localctx = NewClassBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, CtoParserRULE_classBody)
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
		p.SetState(190)
		p.Match(CtoParserLBRACE)
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-28)&-(0x1f+1)) == 0 && ((1<<uint((_la-28)))&((1<<(CtoParserSEMI-28))|(1<<(CtoParserREF-28))|(1<<(CtoParserVAR-28)))) != 0 {
		{
			p.SetState(191)
			p.ClassBodyDeclaration()
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(197)
		p.Match(CtoParserRBRACE)
	}

	return localctx
}

// IClassBodyDeclarationContext is an interface to support dynamic dispatch.
type IClassBodyDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClassBodyDeclarationContext differentiates from other interfaces.
	IsClassBodyDeclarationContext()
}

type ClassBodyDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassBodyDeclarationContext() *ClassBodyDeclarationContext {
	var p = new(ClassBodyDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_classBodyDeclaration
	return p
}

func (*ClassBodyDeclarationContext) IsClassBodyDeclarationContext() {}

func NewClassBodyDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassBodyDeclarationContext {
	var p = new(ClassBodyDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_classBodyDeclaration

	return p
}

func (s *ClassBodyDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassBodyDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(CtoParserSEMI, 0)
}

func (s *ClassBodyDeclarationContext) FieldDeclaration() IFieldDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldDeclarationContext)
}

func (s *ClassBodyDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassBodyDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassBodyDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterClassBodyDeclaration(s)
	}
}

func (s *ClassBodyDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitClassBodyDeclaration(s)
	}
}

func (p *CtoParser) ClassBodyDeclaration() (localctx IClassBodyDeclarationContext) {
	this := p
	_ = this

	localctx = NewClassBodyDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, CtoParserRULE_classBodyDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(201)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CtoParserSEMI:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Match(CtoParserSEMI)
		}

	case CtoParserREF, CtoParserVAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(200)
			p.FieldDeclaration()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFieldDeclarationContext is an interface to support dynamic dispatch.
type IFieldDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDeclarationContext differentiates from other interfaces.
	IsFieldDeclarationContext()
}

type FieldDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDeclarationContext() *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_fieldDeclaration
	return p
}

func (*FieldDeclarationContext) IsFieldDeclarationContext() {}

func NewFieldDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDeclarationContext {
	var p = new(FieldDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_fieldDeclaration

	return p
}

func (s *FieldDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDeclarationContext) StringField() IStringFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringFieldContext)
}

func (s *FieldDeclarationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *FieldDeclarationContext) DefaultString() IDefaultStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultStringContext)
}

func (s *FieldDeclarationContext) RegexDeclaration() IRegexDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRegexDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRegexDeclarationContext)
}

func (s *FieldDeclarationContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(CtoParserOPTIONAL, 0)
}

func (s *FieldDeclarationContext) BooleanField() IBooleanFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanFieldContext)
}

func (s *FieldDeclarationContext) DefaultBoolean() IDefaultBooleanContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultBooleanContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultBooleanContext)
}

func (s *FieldDeclarationContext) NumericField() INumericFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericFieldContext)
}

func (s *FieldDeclarationContext) DefaultNumber() IDefaultNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultNumberContext)
}

func (s *FieldDeclarationContext) RangeValidation() IRangeValidationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeValidationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeValidationContext)
}

func (s *FieldDeclarationContext) DateField() IDateFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDateFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDateFieldContext)
}

func (s *FieldDeclarationContext) DefaultDate() IDefaultDateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultDateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultDateContext)
}

func (s *FieldDeclarationContext) IdentifierField() IIdentifierFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierFieldContext)
}

func (s *FieldDeclarationContext) Reference() IReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferenceContext)
}

func (s *FieldDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterFieldDeclaration(s)
	}
}

func (s *FieldDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitFieldDeclaration(s)
	}
}

func (p *CtoParser) FieldDeclaration() (localctx IFieldDeclarationContext) {
	this := p
	_ = this

	localctx = NewFieldDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, CtoParserRULE_fieldDeclaration)
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.StringField()
		}
		{
			p.SetState(204)
			p.Identifier()
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CtoParserDEFAULT {
			{
				p.SetState(205)
				p.DefaultString()
			}

		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CtoParserREGEX {
			{
				p.SetState(208)
				p.RegexDeclaration()
			}

		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CtoParserOPTIONAL {
			{
				p.SetState(211)
				p.Match(CtoParserOPTIONAL)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.BooleanField()
		}
		{
			p.SetState(215)
			p.Identifier()
		}
		p.SetState(217)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CtoParserDEFAULT {
			{
				p.SetState(216)
				p.DefaultBoolean()
			}

		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CtoParserOPTIONAL {
			{
				p.SetState(219)
				p.Match(CtoParserOPTIONAL)
			}

		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(222)
			p.NumericField()
		}
		{
			p.SetState(223)
			p.Identifier()
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CtoParserDEFAULT {
			{
				p.SetState(224)
				p.DefaultNumber()
			}

		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CtoParserRANGE {
			{
				p.SetState(227)
				p.RangeValidation()
			}

		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CtoParserOPTIONAL {
			{
				p.SetState(230)
				p.Match(CtoParserOPTIONAL)
			}

		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(233)
			p.DateField()
		}
		{
			p.SetState(234)
			p.Identifier()
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CtoParserDEFAULT {
			{
				p.SetState(235)
				p.DefaultDate()
			}

		}
		p.SetState(239)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == CtoParserOPTIONAL {
			{
				p.SetState(238)
				p.Match(CtoParserOPTIONAL)
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(241)
			p.IdentifierField()
		}
		{
			p.SetState(242)
			p.Identifier()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(244)
			p.Reference()
		}
		{
			p.SetState(245)
			p.Identifier()
		}

	}

	return localctx
}

// IIdentifierFieldContext is an interface to support dynamic dispatch.
type IIdentifierFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierFieldContext differentiates from other interfaces.
	IsIdentifierFieldContext()
}

type IdentifierFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierFieldContext() *IdentifierFieldContext {
	var p = new(IdentifierFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_identifierField
	return p
}

func (*IdentifierFieldContext) IsIdentifierFieldContext() {}

func NewIdentifierFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierFieldContext {
	var p = new(IdentifierFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_identifierField

	return p
}

func (s *IdentifierFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierFieldContext) VAR() antlr.TerminalNode {
	return s.GetToken(CtoParserVAR, 0)
}

func (s *IdentifierFieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, 0)
}

func (s *IdentifierFieldContext) AllSquare() []ISquareContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISquareContext)(nil)).Elem())
	var tst = make([]ISquareContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISquareContext)
		}
	}

	return tst
}

func (s *IdentifierFieldContext) Square(i int) ISquareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISquareContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISquareContext)
}

func (s *IdentifierFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterIdentifierField(s)
	}
}

func (s *IdentifierFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitIdentifierField(s)
	}
}

func (p *CtoParser) IdentifierField() (localctx IIdentifierFieldContext) {
	this := p
	_ = this

	localctx = NewIdentifierFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, CtoParserRULE_identifierField)
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
		p.SetState(249)
		p.Match(CtoParserVAR)
	}
	{
		p.SetState(250)
		p.Match(CtoParserIDENTIFIER)
	}
	p.SetState(254)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CtoParserLBRACK {
		{
			p.SetState(251)
			p.Square()
		}

		p.SetState(256)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INumericFieldContext is an interface to support dynamic dispatch.
type INumericFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericFieldContext differentiates from other interfaces.
	IsNumericFieldContext()
}

type NumericFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericFieldContext() *NumericFieldContext {
	var p = new(NumericFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_numericField
	return p
}

func (*NumericFieldContext) IsNumericFieldContext() {}

func NewNumericFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericFieldContext {
	var p = new(NumericFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_numericField

	return p
}

func (s *NumericFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericFieldContext) VAR() antlr.TerminalNode {
	return s.GetToken(CtoParserVAR, 0)
}

func (s *NumericFieldContext) NumericPrimitive() INumericPrimitiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericPrimitiveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericPrimitiveContext)
}

func (s *NumericFieldContext) AllSquare() []ISquareContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISquareContext)(nil)).Elem())
	var tst = make([]ISquareContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISquareContext)
		}
	}

	return tst
}

func (s *NumericFieldContext) Square(i int) ISquareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISquareContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISquareContext)
}

func (s *NumericFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterNumericField(s)
	}
}

func (s *NumericFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitNumericField(s)
	}
}

func (p *CtoParser) NumericField() (localctx INumericFieldContext) {
	this := p
	_ = this

	localctx = NewNumericFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, CtoParserRULE_numericField)
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
		p.SetState(257)
		p.Match(CtoParserVAR)
	}
	{
		p.SetState(258)
		p.NumericPrimitive()
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CtoParserLBRACK {
		{
			p.SetState(259)
			p.Square()
		}

		p.SetState(264)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INumericPrimitiveContext is an interface to support dynamic dispatch.
type INumericPrimitiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericPrimitiveContext differentiates from other interfaces.
	IsNumericPrimitiveContext()
}

type NumericPrimitiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericPrimitiveContext() *NumericPrimitiveContext {
	var p = new(NumericPrimitiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_numericPrimitive
	return p
}

func (*NumericPrimitiveContext) IsNumericPrimitiveContext() {}

func NewNumericPrimitiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericPrimitiveContext {
	var p = new(NumericPrimitiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_numericPrimitive

	return p
}

func (s *NumericPrimitiveContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericPrimitiveContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(CtoParserDOUBLE, 0)
}

func (s *NumericPrimitiveContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(CtoParserINTEGER, 0)
}

func (s *NumericPrimitiveContext) LONG() antlr.TerminalNode {
	return s.GetToken(CtoParserLONG, 0)
}

func (s *NumericPrimitiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericPrimitiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericPrimitiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterNumericPrimitive(s)
	}
}

func (s *NumericPrimitiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitNumericPrimitive(s)
	}
}

func (p *CtoParser) NumericPrimitive() (localctx INumericPrimitiveContext) {
	this := p
	_ = this

	localctx = NewNumericPrimitiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, CtoParserRULE_numericPrimitive)
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
		p.SetState(265)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<CtoParserDOUBLE)|(1<<CtoParserINTEGER)|(1<<CtoParserLONG))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBooleanFieldContext is an interface to support dynamic dispatch.
type IBooleanFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanFieldContext differentiates from other interfaces.
	IsBooleanFieldContext()
}

type BooleanFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanFieldContext() *BooleanFieldContext {
	var p = new(BooleanFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_booleanField
	return p
}

func (*BooleanFieldContext) IsBooleanFieldContext() {}

func NewBooleanFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanFieldContext {
	var p = new(BooleanFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_booleanField

	return p
}

func (s *BooleanFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanFieldContext) VAR() antlr.TerminalNode {
	return s.GetToken(CtoParserVAR, 0)
}

func (s *BooleanFieldContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(CtoParserBOOLEAN, 0)
}

func (s *BooleanFieldContext) AllSquare() []ISquareContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISquareContext)(nil)).Elem())
	var tst = make([]ISquareContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISquareContext)
		}
	}

	return tst
}

func (s *BooleanFieldContext) Square(i int) ISquareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISquareContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISquareContext)
}

func (s *BooleanFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterBooleanField(s)
	}
}

func (s *BooleanFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitBooleanField(s)
	}
}

func (p *CtoParser) BooleanField() (localctx IBooleanFieldContext) {
	this := p
	_ = this

	localctx = NewBooleanFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, CtoParserRULE_booleanField)
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
		p.SetState(267)
		p.Match(CtoParserVAR)
	}
	{
		p.SetState(268)
		p.Match(CtoParserBOOLEAN)
	}
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CtoParserLBRACK {
		{
			p.SetState(269)
			p.Square()
		}

		p.SetState(274)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDateFieldContext is an interface to support dynamic dispatch.
type IDateFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDateFieldContext differentiates from other interfaces.
	IsDateFieldContext()
}

type DateFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateFieldContext() *DateFieldContext {
	var p = new(DateFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_dateField
	return p
}

func (*DateFieldContext) IsDateFieldContext() {}

func NewDateFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateFieldContext {
	var p = new(DateFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_dateField

	return p
}

func (s *DateFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *DateFieldContext) VAR() antlr.TerminalNode {
	return s.GetToken(CtoParserVAR, 0)
}

func (s *DateFieldContext) DATE_TIME() antlr.TerminalNode {
	return s.GetToken(CtoParserDATE_TIME, 0)
}

func (s *DateFieldContext) AllSquare() []ISquareContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISquareContext)(nil)).Elem())
	var tst = make([]ISquareContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISquareContext)
		}
	}

	return tst
}

func (s *DateFieldContext) Square(i int) ISquareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISquareContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISquareContext)
}

func (s *DateFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterDateField(s)
	}
}

func (s *DateFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitDateField(s)
	}
}

func (p *CtoParser) DateField() (localctx IDateFieldContext) {
	this := p
	_ = this

	localctx = NewDateFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, CtoParserRULE_dateField)
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
		p.SetState(275)
		p.Match(CtoParserVAR)
	}
	{
		p.SetState(276)
		p.Match(CtoParserDATE_TIME)
	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CtoParserLBRACK {
		{
			p.SetState(277)
			p.Square()
		}

		p.SetState(282)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefaultDateContext is an interface to support dynamic dispatch.
type IDefaultDateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultDateContext differentiates from other interfaces.
	IsDefaultDateContext()
}

type DefaultDateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultDateContext() *DefaultDateContext {
	var p = new(DefaultDateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_defaultDate
	return p
}

func (*DefaultDateContext) IsDefaultDateContext() {}

func NewDefaultDateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultDateContext {
	var p = new(DefaultDateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_defaultDate

	return p
}

func (s *DefaultDateContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultDateContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(CtoParserDEFAULT, 0)
}

func (s *DefaultDateContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CtoParserASSIGN, 0)
}

func (s *DefaultDateContext) DATE_TIME_LITERAL() antlr.TerminalNode {
	return s.GetToken(CtoParserDATE_TIME_LITERAL, 0)
}

func (s *DefaultDateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultDateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultDateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterDefaultDate(s)
	}
}

func (s *DefaultDateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitDefaultDate(s)
	}
}

func (p *CtoParser) DefaultDate() (localctx IDefaultDateContext) {
	this := p
	_ = this

	localctx = NewDefaultDateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, CtoParserRULE_defaultDate)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(CtoParserDEFAULT)
	}
	{
		p.SetState(284)
		p.Match(CtoParserASSIGN)
	}
	{
		p.SetState(285)
		p.Match(CtoParserDATE_TIME_LITERAL)
	}

	return localctx
}

// IRegexDeclarationContext is an interface to support dynamic dispatch.
type IRegexDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRegexDeclarationContext differentiates from other interfaces.
	IsRegexDeclarationContext()
}

type RegexDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexDeclarationContext() *RegexDeclarationContext {
	var p = new(RegexDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_regexDeclaration
	return p
}

func (*RegexDeclarationContext) IsRegexDeclarationContext() {}

func NewRegexDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexDeclarationContext {
	var p = new(RegexDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_regexDeclaration

	return p
}

func (s *RegexDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexDeclarationContext) REGEX() antlr.TerminalNode {
	return s.GetToken(CtoParserREGEX, 0)
}

func (s *RegexDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CtoParserASSIGN, 0)
}

func (s *RegexDeclarationContext) REGEX_EXPR() antlr.TerminalNode {
	return s.GetToken(CtoParserREGEX_EXPR, 0)
}

func (s *RegexDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterRegexDeclaration(s)
	}
}

func (s *RegexDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitRegexDeclaration(s)
	}
}

func (p *CtoParser) RegexDeclaration() (localctx IRegexDeclarationContext) {
	this := p
	_ = this

	localctx = NewRegexDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, CtoParserRULE_regexDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(287)
		p.Match(CtoParserREGEX)
	}
	{
		p.SetState(288)
		p.Match(CtoParserASSIGN)
	}
	{
		p.SetState(289)
		p.Match(CtoParserREGEX_EXPR)
	}

	return localctx
}

// IStringFieldContext is an interface to support dynamic dispatch.
type IStringFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringFieldContext differentiates from other interfaces.
	IsStringFieldContext()
}

type StringFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringFieldContext() *StringFieldContext {
	var p = new(StringFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_stringField
	return p
}

func (*StringFieldContext) IsStringFieldContext() {}

func NewStringFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringFieldContext {
	var p = new(StringFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_stringField

	return p
}

func (s *StringFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *StringFieldContext) VAR() antlr.TerminalNode {
	return s.GetToken(CtoParserVAR, 0)
}

func (s *StringFieldContext) STRING() antlr.TerminalNode {
	return s.GetToken(CtoParserSTRING, 0)
}

func (s *StringFieldContext) AllSquare() []ISquareContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISquareContext)(nil)).Elem())
	var tst = make([]ISquareContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISquareContext)
		}
	}

	return tst
}

func (s *StringFieldContext) Square(i int) ISquareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISquareContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISquareContext)
}

func (s *StringFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterStringField(s)
	}
}

func (s *StringFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitStringField(s)
	}
}

func (p *CtoParser) StringField() (localctx IStringFieldContext) {
	this := p
	_ = this

	localctx = NewStringFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, CtoParserRULE_stringField)
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
		p.SetState(291)
		p.Match(CtoParserVAR)
	}
	{
		p.SetState(292)
		p.Match(CtoParserSTRING)
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CtoParserLBRACK {
		{
			p.SetState(293)
			p.Square()
		}

		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IReferenceContext is an interface to support dynamic dispatch.
type IReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferenceContext differentiates from other interfaces.
	IsReferenceContext()
}

type ReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceContext() *ReferenceContext {
	var p = new(ReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_reference
	return p
}

func (*ReferenceContext) IsReferenceContext() {}

func NewReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceContext {
	var p = new(ReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_reference

	return p
}

func (s *ReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceContext) REF() antlr.TerminalNode {
	return s.GetToken(CtoParserREF, 0)
}

func (s *ReferenceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, 0)
}

func (s *ReferenceContext) AllSquare() []ISquareContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISquareContext)(nil)).Elem())
	var tst = make([]ISquareContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISquareContext)
		}
	}

	return tst
}

func (s *ReferenceContext) Square(i int) ISquareContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISquareContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISquareContext)
}

func (s *ReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterReference(s)
	}
}

func (s *ReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitReference(s)
	}
}

func (p *CtoParser) Reference() (localctx IReferenceContext) {
	this := p
	_ = this

	localctx = NewReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, CtoParserRULE_reference)
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
		p.SetState(299)
		p.Match(CtoParserREF)
	}
	{
		p.SetState(300)
		p.Match(CtoParserIDENTIFIER)
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CtoParserLBRACK {
		{
			p.SetState(301)
			p.Square()
		}

		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IQualifiedNameContext is an interface to support dynamic dispatch.
type IQualifiedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedNameContext differentiates from other interfaces.
	IsQualifiedNameContext()
}

type QualifiedNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedNameContext() *QualifiedNameContext {
	var p = new(QualifiedNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_qualifiedName
	return p
}

func (*QualifiedNameContext) IsQualifiedNameContext() {}

func NewQualifiedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameContext {
	var p = new(QualifiedNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_qualifiedName

	return p
}

func (s *QualifiedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(CtoParserIDENTIFIER)
}

func (s *QualifiedNameContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, i)
}

func (s *QualifiedNameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(CtoParserDOT)
}

func (s *QualifiedNameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(CtoParserDOT, i)
}

func (s *QualifiedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterQualifiedName(s)
	}
}

func (s *QualifiedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitQualifiedName(s)
	}
}

func (p *CtoParser) QualifiedName() (localctx IQualifiedNameContext) {
	this := p
	_ = this

	localctx = NewQualifiedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, CtoParserRULE_qualifiedName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(307)
		p.Match(CtoParserIDENTIFIER)
	}
	p.SetState(312)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(308)
				p.Match(CtoParserDOT)
			}
			{
				p.SetState(309)
				p.Match(CtoParserIDENTIFIER)
			}

		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}

	return localctx
}

// IRangeValidationContext is an interface to support dynamic dispatch.
type IRangeValidationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeValidationContext differentiates from other interfaces.
	IsRangeValidationContext()
}

type RangeValidationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeValidationContext() *RangeValidationContext {
	var p = new(RangeValidationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_rangeValidation
	return p
}

func (*RangeValidationContext) IsRangeValidationContext() {}

func NewRangeValidationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeValidationContext {
	var p = new(RangeValidationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_rangeValidation

	return p
}

func (s *RangeValidationContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeValidationContext) RANGE() antlr.TerminalNode {
	return s.GetToken(CtoParserRANGE, 0)
}

func (s *RangeValidationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CtoParserASSIGN, 0)
}

func (s *RangeValidationContext) RangeDeclaration() IRangeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeDeclarationContext)
}

func (s *RangeValidationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeValidationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeValidationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterRangeValidation(s)
	}
}

func (s *RangeValidationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitRangeValidation(s)
	}
}

func (p *CtoParser) RangeValidation() (localctx IRangeValidationContext) {
	this := p
	_ = this

	localctx = NewRangeValidationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, CtoParserRULE_rangeValidation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Match(CtoParserRANGE)
	}
	{
		p.SetState(316)
		p.Match(CtoParserASSIGN)
	}
	{
		p.SetState(317)
		p.RangeDeclaration()
	}

	return localctx
}

// IRangeDeclarationContext is an interface to support dynamic dispatch.
type IRangeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeDeclarationContext differentiates from other interfaces.
	IsRangeDeclarationContext()
}

type RangeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeDeclarationContext() *RangeDeclarationContext {
	var p = new(RangeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_rangeDeclaration
	return p
}

func (*RangeDeclarationContext) IsRangeDeclarationContext() {}

func NewRangeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeDeclarationContext {
	var p = new(RangeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_rangeDeclaration

	return p
}

func (s *RangeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeDeclarationContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(CtoParserLBRACK, 0)
}

func (s *RangeDeclarationContext) AllNumberLiteral() []INumberLiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INumberLiteralContext)(nil)).Elem())
	var tst = make([]INumberLiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INumberLiteralContext)
		}
	}

	return tst
}

func (s *RangeDeclarationContext) NumberLiteral(i int) INumberLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INumberLiteralContext)
}

func (s *RangeDeclarationContext) COMMA() antlr.TerminalNode {
	return s.GetToken(CtoParserCOMMA, 0)
}

func (s *RangeDeclarationContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(CtoParserRBRACK, 0)
}

func (s *RangeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterRangeDeclaration(s)
	}
}

func (s *RangeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitRangeDeclaration(s)
	}
}

func (p *CtoParser) RangeDeclaration() (localctx IRangeDeclarationContext) {
	this := p
	_ = this

	localctx = NewRangeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, CtoParserRULE_rangeDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.Match(CtoParserLBRACK)
		}
		{
			p.SetState(320)
			p.NumberLiteral()
		}
		{
			p.SetState(321)
			p.Match(CtoParserCOMMA)
		}
		{
			p.SetState(322)
			p.Match(CtoParserRBRACK)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.Match(CtoParserLBRACK)
		}
		{
			p.SetState(325)
			p.Match(CtoParserCOMMA)
		}
		{
			p.SetState(326)
			p.NumberLiteral()
		}
		{
			p.SetState(327)
			p.Match(CtoParserRBRACK)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(329)
			p.Match(CtoParserLBRACK)
		}
		{
			p.SetState(330)
			p.NumberLiteral()
		}
		{
			p.SetState(331)
			p.Match(CtoParserCOMMA)
		}
		{
			p.SetState(332)
			p.NumberLiteral()
		}
		{
			p.SetState(333)
			p.Match(CtoParserRBRACK)
		}

	}

	return localctx
}

// IDefaultBooleanContext is an interface to support dynamic dispatch.
type IDefaultBooleanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultBooleanContext differentiates from other interfaces.
	IsDefaultBooleanContext()
}

type DefaultBooleanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultBooleanContext() *DefaultBooleanContext {
	var p = new(DefaultBooleanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_defaultBoolean
	return p
}

func (*DefaultBooleanContext) IsDefaultBooleanContext() {}

func NewDefaultBooleanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultBooleanContext {
	var p = new(DefaultBooleanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_defaultBoolean

	return p
}

func (s *DefaultBooleanContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultBooleanContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(CtoParserDEFAULT, 0)
}

func (s *DefaultBooleanContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CtoParserASSIGN, 0)
}

func (s *DefaultBooleanContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(CtoParserBOOL_LITERAL, 0)
}

func (s *DefaultBooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultBooleanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultBooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterDefaultBoolean(s)
	}
}

func (s *DefaultBooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitDefaultBoolean(s)
	}
}

func (p *CtoParser) DefaultBoolean() (localctx IDefaultBooleanContext) {
	this := p
	_ = this

	localctx = NewDefaultBooleanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, CtoParserRULE_defaultBoolean)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		p.Match(CtoParserDEFAULT)
	}
	{
		p.SetState(338)
		p.Match(CtoParserASSIGN)
	}
	{
		p.SetState(339)
		p.Match(CtoParserBOOL_LITERAL)
	}

	return localctx
}

// IDefaultStringContext is an interface to support dynamic dispatch.
type IDefaultStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultStringContext differentiates from other interfaces.
	IsDefaultStringContext()
}

type DefaultStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultStringContext() *DefaultStringContext {
	var p = new(DefaultStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_defaultString
	return p
}

func (*DefaultStringContext) IsDefaultStringContext() {}

func NewDefaultStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultStringContext {
	var p = new(DefaultStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_defaultString

	return p
}

func (s *DefaultStringContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultStringContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(CtoParserDEFAULT, 0)
}

func (s *DefaultStringContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CtoParserASSIGN, 0)
}

func (s *DefaultStringContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *DefaultStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterDefaultString(s)
	}
}

func (s *DefaultStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitDefaultString(s)
	}
}

func (p *CtoParser) DefaultString() (localctx IDefaultStringContext) {
	this := p
	_ = this

	localctx = NewDefaultStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, CtoParserRULE_defaultString)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(341)
		p.Match(CtoParserDEFAULT)
	}
	{
		p.SetState(342)
		p.Match(CtoParserASSIGN)
	}
	{
		p.SetState(343)
		p.StringLiteral()
	}

	return localctx
}

// IDefaultNumberContext is an interface to support dynamic dispatch.
type IDefaultNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultNumberContext differentiates from other interfaces.
	IsDefaultNumberContext()
}

type DefaultNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultNumberContext() *DefaultNumberContext {
	var p = new(DefaultNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_defaultNumber
	return p
}

func (*DefaultNumberContext) IsDefaultNumberContext() {}

func NewDefaultNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultNumberContext {
	var p = new(DefaultNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_defaultNumber

	return p
}

func (s *DefaultNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultNumberContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(CtoParserDEFAULT, 0)
}

func (s *DefaultNumberContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(CtoParserASSIGN, 0)
}

func (s *DefaultNumberContext) NumberLiteral() INumberLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLiteralContext)
}

func (s *DefaultNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterDefaultNumber(s)
	}
}

func (s *DefaultNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitDefaultNumber(s)
	}
}

func (p *CtoParser) DefaultNumber() (localctx IDefaultNumberContext) {
	this := p
	_ = this

	localctx = NewDefaultNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, CtoParserRULE_defaultNumber)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(CtoParserDEFAULT)
	}
	{
		p.SetState(346)
		p.Match(CtoParserASSIGN)
	}
	{
		p.SetState(347)
		p.NumberLiteral()
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(CtoParserIDENTIFIER, 0)
}

func (s *IdentifierContext) ASSET() antlr.TerminalNode {
	return s.GetToken(CtoParserASSET, 0)
}

func (s *IdentifierContext) PARTICIPANT() antlr.TerminalNode {
	return s.GetToken(CtoParserPARTICIPANT, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *CtoParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, CtoParserRULE_identifier)
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
		p.SetState(349)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CtoParserASSET || _la == CtoParserPARTICIPANT || _la == CtoParserIDENTIFIER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = CtoParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NumberLiteral() INumberLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberLiteralContext)
}

func (s *LiteralContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *LiteralContext) BOOL_LITERAL() antlr.TerminalNode {
	return s.GetToken(CtoParserBOOL_LITERAL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (p *CtoParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, CtoParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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

	switch p.GetTokenStream().LA(1) {
	case CtoParserDECIMAL_LITERAL, CtoParserOCT_LITERAL, CtoParserFLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(351)
			p.NumberLiteral()
		}

	case CtoParserCHAR_LITERAL, CtoParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(352)
			p.StringLiteral()
		}

	case CtoParserBOOL_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(353)
			p.Match(CtoParserBOOL_LITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumberLiteralContext is an interface to support dynamic dispatch.
type INumberLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberLiteralContext differentiates from other interfaces.
	IsNumberLiteralContext()
}

type NumberLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberLiteralContext() *NumberLiteralContext {
	var p = new(NumberLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_numberLiteral
	return p
}

func (*NumberLiteralContext) IsNumberLiteralContext() {}

func NewNumberLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_numberLiteral

	return p
}

func (s *NumberLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberLiteralContext) IntegerLiteral() IIntegerLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerLiteralContext)
}

func (s *NumberLiteralContext) FloatLiteral() IFloatLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatLiteralContext)
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitNumberLiteral(s)
	}
}

func (p *CtoParser) NumberLiteral() (localctx INumberLiteralContext) {
	this := p
	_ = this

	localctx = NewNumberLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, CtoParserRULE_numberLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(358)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CtoParserDECIMAL_LITERAL, CtoParserOCT_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(356)
			p.IntegerLiteral()
		}

	case CtoParserFLOAT_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(357)
			p.FloatLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) CHAR_LITERAL() antlr.TerminalNode {
	return s.GetToken(CtoParserCHAR_LITERAL, 0)
}

func (s *StringLiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(CtoParserSTRING_LITERAL, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (p *CtoParser) StringLiteral() (localctx IStringLiteralContext) {
	this := p
	_ = this

	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, CtoParserRULE_stringLiteral)
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
		p.SetState(360)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CtoParserCHAR_LITERAL || _la == CtoParserSTRING_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIntegerLiteralContext is an interface to support dynamic dispatch.
type IIntegerLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerLiteralContext differentiates from other interfaces.
	IsIntegerLiteralContext()
}

type IntegerLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerLiteralContext() *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_integerLiteral
	return p
}

func (*IntegerLiteralContext) IsIntegerLiteralContext() {}

func NewIntegerLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerLiteralContext {
	var p = new(IntegerLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_integerLiteral

	return p
}

func (s *IntegerLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(CtoParserDECIMAL_LITERAL, 0)
}

func (s *IntegerLiteralContext) OCT_LITERAL() antlr.TerminalNode {
	return s.GetToken(CtoParserOCT_LITERAL, 0)
}

func (s *IntegerLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterIntegerLiteral(s)
	}
}

func (s *IntegerLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitIntegerLiteral(s)
	}
}

func (p *CtoParser) IntegerLiteral() (localctx IIntegerLiteralContext) {
	this := p
	_ = this

	localctx = NewIntegerLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, CtoParserRULE_integerLiteral)
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
		p.SetState(362)
		_la = p.GetTokenStream().LA(1)

		if !(_la == CtoParserDECIMAL_LITERAL || _la == CtoParserOCT_LITERAL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFloatLiteralContext is an interface to support dynamic dispatch.
type IFloatLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatLiteralContext differentiates from other interfaces.
	IsFloatLiteralContext()
}

type FloatLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLiteralContext() *FloatLiteralContext {
	var p = new(FloatLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_floatLiteral
	return p
}

func (*FloatLiteralContext) IsFloatLiteralContext() {}

func NewFloatLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLiteralContext {
	var p = new(FloatLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_floatLiteral

	return p
}

func (s *FloatLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLiteralContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(CtoParserFLOAT_LITERAL, 0)
}

func (s *FloatLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterFloatLiteral(s)
	}
}

func (s *FloatLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitFloatLiteral(s)
	}
}

func (p *CtoParser) FloatLiteral() (localctx IFloatLiteralContext) {
	this := p
	_ = this

	localctx = NewFloatLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, CtoParserRULE_floatLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(CtoParserFLOAT_LITERAL)
	}

	return localctx
}

// IDecoratorContext is an interface to support dynamic dispatch.
type IDecoratorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecoratorContext differentiates from other interfaces.
	IsDecoratorContext()
}

type DecoratorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecoratorContext() *DecoratorContext {
	var p = new(DecoratorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_decorator
	return p
}

func (*DecoratorContext) IsDecoratorContext() {}

func NewDecoratorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecoratorContext {
	var p = new(DecoratorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_decorator

	return p
}

func (s *DecoratorContext) GetParser() antlr.Parser { return s.parser }

func (s *DecoratorContext) AT() antlr.TerminalNode {
	return s.GetToken(CtoParserAT, 0)
}

func (s *DecoratorContext) QualifiedName() IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *DecoratorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(CtoParserLPAREN, 0)
}

func (s *DecoratorContext) ElementValuePair() IElementValuePairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementValuePairContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElementValuePairContext)
}

func (s *DecoratorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(CtoParserRPAREN, 0)
}

func (s *DecoratorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecoratorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecoratorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterDecorator(s)
	}
}

func (s *DecoratorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitDecorator(s)
	}
}

func (p *CtoParser) Decorator() (localctx IDecoratorContext) {
	this := p
	_ = this

	localctx = NewDecoratorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, CtoParserRULE_decorator)
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
		p.SetState(366)
		p.Match(CtoParserAT)
	}
	{
		p.SetState(367)
		p.QualifiedName()
	}
	p.SetState(372)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == CtoParserLPAREN {
		{
			p.SetState(368)
			p.Match(CtoParserLPAREN)
		}
		{
			p.SetState(369)
			p.ElementValuePair()
		}
		{
			p.SetState(370)
			p.Match(CtoParserRPAREN)
		}

	}

	return localctx
}

// IElementValuePairContext is an interface to support dynamic dispatch.
type IElementValuePairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementValuePairContext differentiates from other interfaces.
	IsElementValuePairContext()
}

type ElementValuePairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementValuePairContext() *ElementValuePairContext {
	var p = new(ElementValuePairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_elementValuePair
	return p
}

func (*ElementValuePairContext) IsElementValuePairContext() {}

func NewElementValuePairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementValuePairContext {
	var p = new(ElementValuePairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_elementValuePair

	return p
}

func (s *ElementValuePairContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementValuePairContext) AllLiteral() []ILiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralContext)(nil)).Elem())
	var tst = make([]ILiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralContext)
		}
	}

	return tst
}

func (s *ElementValuePairContext) Literal(i int) ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ElementValuePairContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(CtoParserCOMMA)
}

func (s *ElementValuePairContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(CtoParserCOMMA, i)
}

func (s *ElementValuePairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementValuePairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementValuePairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterElementValuePair(s)
	}
}

func (s *ElementValuePairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitElementValuePair(s)
	}
}

func (p *CtoParser) ElementValuePair() (localctx IElementValuePairContext) {
	this := p
	_ = this

	localctx = NewElementValuePairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, CtoParserRULE_elementValuePair)
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
		p.SetState(374)
		p.Literal()
	}
	p.SetState(379)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == CtoParserCOMMA {
		{
			p.SetState(375)
			p.Match(CtoParserCOMMA)
		}
		{
			p.SetState(376)
			p.Literal()
		}

		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISquareContext is an interface to support dynamic dispatch.
type ISquareContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSquareContext differentiates from other interfaces.
	IsSquareContext()
}

type SquareContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySquareContext() *SquareContext {
	var p = new(SquareContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CtoParserRULE_square
	return p
}

func (*SquareContext) IsSquareContext() {}

func NewSquareContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SquareContext {
	var p = new(SquareContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CtoParserRULE_square

	return p
}

func (s *SquareContext) GetParser() antlr.Parser { return s.parser }

func (s *SquareContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(CtoParserLBRACK, 0)
}

func (s *SquareContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(CtoParserRBRACK, 0)
}

func (s *SquareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SquareContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SquareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.EnterSquare(s)
	}
}

func (s *SquareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CtoParserListener); ok {
		listenerT.ExitSquare(s)
	}
}

func (p *CtoParser) Square() (localctx ISquareContext) {
	this := p
	_ = this

	localctx = NewSquareContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, CtoParserRULE_square)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(382)
		p.Match(CtoParserLBRACK)
	}
	{
		p.SetState(383)
		p.Match(CtoParserRBRACK)
	}

	return localctx
}
