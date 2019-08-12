// Code generated from LessParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package less // LessParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 145, 371,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3, 2, 7,
	2, 78, 10, 2, 12, 2, 14, 2, 81, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 5, 3, 89, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 95, 10, 4, 3, 5, 6, 5,
	98, 10, 5, 13, 5, 14, 5, 99, 3, 5, 5, 5, 103, 10, 5, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 118,
	10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 129,
	10, 8, 3, 9, 3, 9, 3, 9, 5, 9, 134, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	10, 7, 10, 141, 10, 10, 12, 10, 14, 10, 144, 11, 10, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 155, 10, 11, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 5, 12, 162, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 173, 10, 14, 12, 14, 14, 14,
	176, 11, 14, 3, 14, 3, 14, 5, 14, 180, 10, 14, 3, 14, 3, 14, 5, 14, 184,
	10, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16,
	194, 10, 16, 3, 17, 3, 17, 3, 17, 7, 17, 199, 10, 17, 12, 17, 14, 17, 202,
	11, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	7, 19, 213, 10, 19, 12, 19, 14, 19, 216, 11, 19, 3, 19, 5, 19, 219, 10,
	19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 7, 20, 228, 10, 20,
	12, 20, 14, 20, 231, 11, 20, 5, 20, 233, 10, 20, 3, 20, 5, 20, 236, 10,
	20, 3, 20, 3, 20, 5, 20, 240, 10, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21,
	3, 22, 3, 22, 5, 22, 249, 10, 22, 3, 23, 3, 23, 3, 23, 5, 23, 254, 10,
	23, 3, 23, 3, 23, 5, 23, 258, 10, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24,
	7, 24, 265, 10, 24, 12, 24, 14, 24, 268, 11, 24, 3, 25, 6, 25, 271, 10,
	25, 13, 25, 14, 25, 272, 3, 25, 7, 25, 276, 10, 25, 12, 25, 14, 25, 279,
	11, 25, 3, 25, 5, 25, 282, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5,
	26, 289, 10, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 297,
	10, 27, 3, 27, 3, 27, 5, 27, 301, 10, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3,
	28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	5, 29, 318, 10, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 7, 32, 326,
	10, 32, 12, 32, 14, 32, 329, 11, 32, 3, 32, 3, 32, 3, 32, 3, 32, 7, 32,
	335, 10, 32, 12, 32, 14, 32, 338, 11, 32, 5, 32, 340, 10, 32, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 5, 33, 347, 10, 33, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 7, 36, 358, 10, 36, 12, 36, 14, 36,
	361, 11, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 5, 38, 369, 10,
	38, 3, 38, 2, 2, 39, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
	66, 68, 70, 72, 74, 2, 11, 3, 2, 26, 30, 4, 2, 19, 19, 54, 54, 6, 2, 14,
	14, 16, 16, 32, 33, 35, 35, 3, 2, 46, 51, 3, 2, 55, 56, 4, 2, 17, 17, 25,
	25, 4, 2, 14, 15, 26, 26, 3, 2, 35, 37, 4, 2, 55, 55, 145, 145, 2, 388,
	2, 79, 3, 2, 2, 2, 4, 88, 3, 2, 2, 2, 6, 94, 3, 2, 2, 2, 8, 97, 3, 2, 2,
	2, 10, 104, 3, 2, 2, 2, 12, 106, 3, 2, 2, 2, 14, 128, 3, 2, 2, 2, 16, 130,
	3, 2, 2, 2, 18, 137, 3, 2, 2, 2, 20, 154, 3, 2, 2, 2, 22, 161, 3, 2, 2,
	2, 24, 163, 3, 2, 2, 2, 26, 167, 3, 2, 2, 2, 28, 187, 3, 2, 2, 2, 30, 193,
	3, 2, 2, 2, 32, 195, 3, 2, 2, 2, 34, 203, 3, 2, 2, 2, 36, 206, 3, 2, 2,
	2, 38, 222, 3, 2, 2, 2, 40, 243, 3, 2, 2, 2, 42, 248, 3, 2, 2, 2, 44, 250,
	3, 2, 2, 2, 46, 261, 3, 2, 2, 2, 48, 270, 3, 2, 2, 2, 50, 283, 3, 2, 2,
	2, 52, 292, 3, 2, 2, 2, 54, 304, 3, 2, 2, 2, 56, 317, 3, 2, 2, 2, 58, 319,
	3, 2, 2, 2, 60, 321, 3, 2, 2, 2, 62, 339, 3, 2, 2, 2, 64, 346, 3, 2, 2,
	2, 66, 348, 3, 2, 2, 2, 68, 350, 3, 2, 2, 2, 70, 354, 3, 2, 2, 2, 72, 362,
	3, 2, 2, 2, 74, 366, 3, 2, 2, 2, 76, 78, 5, 4, 3, 2, 77, 76, 3, 2, 2, 2,
	78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 3, 3, 2,
	2, 2, 81, 79, 3, 2, 2, 2, 82, 89, 5, 26, 14, 2, 83, 89, 5, 34, 18, 2, 84,
	85, 5, 24, 13, 2, 85, 86, 7, 18, 2, 2, 86, 89, 3, 2, 2, 2, 87, 89, 5, 38,
	20, 2, 88, 82, 3, 2, 2, 2, 88, 83, 3, 2, 2, 2, 88, 84, 3, 2, 2, 2, 88,
	87, 3, 2, 2, 2, 89, 5, 3, 2, 2, 2, 90, 91, 7, 22, 2, 2, 91, 95, 5, 6, 4,
	2, 92, 93, 7, 22, 2, 2, 93, 95, 7, 55, 2, 2, 94, 90, 3, 2, 2, 2, 94, 92,
	3, 2, 2, 2, 95, 7, 3, 2, 2, 2, 96, 98, 5, 14, 8, 2, 97, 96, 3, 2, 2, 2,
	98, 99, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 102,
	3, 2, 2, 2, 101, 103, 5, 12, 7, 2, 102, 101, 3, 2, 2, 2, 102, 103, 3, 2,
	2, 2, 103, 9, 3, 2, 2, 2, 104, 105, 9, 2, 2, 2, 105, 11, 3, 2, 2, 2, 106,
	107, 5, 10, 6, 2, 107, 108, 5, 8, 5, 2, 108, 13, 3, 2, 2, 2, 109, 129,
	5, 74, 38, 2, 110, 111, 5, 62, 32, 2, 111, 112, 7, 43, 2, 2, 112, 129,
	3, 2, 2, 2, 113, 129, 5, 62, 32, 2, 114, 115, 5, 62, 32, 2, 115, 117, 7,
	8, 2, 2, 116, 118, 5, 70, 36, 2, 117, 116, 3, 2, 2, 2, 117, 118, 3, 2,
	2, 2, 118, 119, 3, 2, 2, 2, 119, 120, 7, 9, 2, 2, 120, 129, 3, 2, 2, 2,
	121, 129, 7, 58, 2, 2, 122, 129, 7, 56, 2, 2, 123, 129, 5, 72, 37, 2, 124,
	125, 5, 6, 4, 2, 125, 126, 7, 43, 2, 2, 126, 129, 3, 2, 2, 2, 127, 129,
	5, 6, 4, 2, 128, 109, 3, 2, 2, 2, 128, 110, 3, 2, 2, 2, 128, 113, 3, 2,
	2, 2, 128, 114, 3, 2, 2, 2, 128, 121, 3, 2, 2, 2, 128, 122, 3, 2, 2, 2,
	128, 123, 3, 2, 2, 2, 128, 124, 3, 2, 2, 2, 128, 127, 3, 2, 2, 2, 129,
	15, 3, 2, 2, 2, 130, 131, 7, 62, 2, 2, 131, 133, 7, 8, 2, 2, 132, 134,
	5, 70, 36, 2, 133, 132, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 3,
	2, 2, 2, 135, 136, 7, 9, 2, 2, 136, 17, 3, 2, 2, 2, 137, 142, 5, 20, 11,
	2, 138, 139, 9, 3, 2, 2, 139, 141, 5, 20, 11, 2, 140, 138, 3, 2, 2, 2,
	141, 144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143,
	19, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 146, 7, 8, 2, 2, 146, 147, 5,
	22, 12, 2, 147, 148, 7, 9, 2, 2, 148, 155, 3, 2, 2, 2, 149, 150, 7, 53,
	2, 2, 150, 151, 7, 8, 2, 2, 151, 152, 5, 22, 12, 2, 152, 153, 7, 9, 2,
	2, 153, 155, 3, 2, 2, 2, 154, 145, 3, 2, 2, 2, 154, 149, 3, 2, 2, 2, 155,
	21, 3, 2, 2, 2, 156, 157, 5, 8, 5, 2, 157, 158, 9, 4, 2, 2, 158, 159, 5,
	8, 5, 2, 159, 162, 3, 2, 2, 2, 160, 162, 5, 8, 5, 2, 161, 156, 3, 2, 2,
	2, 161, 160, 3, 2, 2, 2, 162, 23, 3, 2, 2, 2, 163, 164, 5, 6, 4, 2, 164,
	165, 7, 17, 2, 2, 165, 166, 5, 70, 36, 2, 166, 25, 3, 2, 2, 2, 167, 179,
	7, 40, 2, 2, 168, 169, 7, 8, 2, 2, 169, 174, 5, 28, 15, 2, 170, 171, 7,
	19, 2, 2, 171, 173, 5, 28, 15, 2, 172, 170, 3, 2, 2, 2, 173, 176, 3, 2,
	2, 2, 174, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 177, 3, 2, 2, 2,
	176, 174, 3, 2, 2, 2, 177, 178, 7, 9, 2, 2, 178, 180, 3, 2, 2, 2, 179,
	168, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 183,
	5, 30, 16, 2, 182, 184, 5, 32, 17, 2, 183, 182, 3, 2, 2, 2, 183, 184, 3,
	2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 186, 7, 18, 2, 2, 186, 27, 3, 2, 2,
	2, 187, 188, 9, 5, 2, 2, 188, 29, 3, 2, 2, 2, 189, 194, 7, 56, 2, 2, 190,
	191, 7, 39, 2, 2, 191, 192, 7, 142, 2, 2, 192, 194, 7, 141, 2, 2, 193,
	189, 3, 2, 2, 2, 193, 190, 3, 2, 2, 2, 194, 31, 3, 2, 2, 2, 195, 200, 7,
	55, 2, 2, 196, 197, 7, 19, 2, 2, 197, 199, 7, 55, 2, 2, 198, 196, 3, 2,
	2, 2, 199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2,
	201, 33, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203, 204, 5, 46, 24, 2, 204,
	205, 5, 36, 19, 2, 205, 35, 3, 2, 2, 2, 206, 214, 7, 10, 2, 2, 207, 208,
	5, 68, 35, 2, 208, 209, 7, 18, 2, 2, 209, 213, 3, 2, 2, 2, 210, 213, 5,
	4, 3, 2, 211, 213, 5, 44, 23, 2, 212, 207, 3, 2, 2, 2, 212, 210, 3, 2,
	2, 2, 212, 211, 3, 2, 2, 2, 213, 216, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2,
	214, 215, 3, 2, 2, 2, 215, 218, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 217,
	219, 5, 68, 35, 2, 218, 217, 3, 2, 2, 2, 218, 219, 3, 2, 2, 2, 219, 220,
	3, 2, 2, 2, 220, 221, 7, 11, 2, 2, 221, 37, 3, 2, 2, 2, 222, 223, 5, 46,
	24, 2, 223, 232, 7, 8, 2, 2, 224, 229, 5, 42, 22, 2, 225, 226, 7, 18, 2,
	2, 226, 228, 5, 42, 22, 2, 227, 225, 3, 2, 2, 2, 228, 231, 3, 2, 2, 2,
	229, 227, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 233, 3, 2, 2, 2, 231,
	229, 3, 2, 2, 2, 232, 224, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 235,
	3, 2, 2, 2, 234, 236, 7, 6, 2, 2, 235, 234, 3, 2, 2, 2, 235, 236, 3, 2,
	2, 2, 236, 237, 3, 2, 2, 2, 237, 239, 7, 9, 2, 2, 238, 240, 5, 40, 21,
	2, 239, 238, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241,
	242, 5, 36, 19, 2, 242, 39, 3, 2, 2, 2, 243, 244, 7, 52, 2, 2, 244, 245,
	5, 18, 10, 2, 245, 41, 3, 2, 2, 2, 246, 249, 5, 6, 4, 2, 247, 249, 5, 24,
	13, 2, 248, 246, 3, 2, 2, 2, 248, 247, 3, 2, 2, 2, 249, 43, 3, 2, 2, 2,
	250, 251, 5, 48, 25, 2, 251, 253, 7, 8, 2, 2, 252, 254, 5, 70, 36, 2, 253,
	252, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 257,
	7, 9, 2, 2, 256, 258, 7, 43, 2, 2, 257, 256, 3, 2, 2, 2, 257, 258, 3, 2,
	2, 2, 258, 259, 3, 2, 2, 2, 259, 260, 7, 18, 2, 2, 260, 45, 3, 2, 2, 2,
	261, 266, 5, 48, 25, 2, 262, 263, 7, 19, 2, 2, 263, 265, 5, 48, 25, 2,
	264, 262, 3, 2, 2, 2, 265, 268, 3, 2, 2, 2, 266, 264, 3, 2, 2, 2, 266,
	267, 3, 2, 2, 2, 267, 47, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 269, 271, 5,
	56, 29, 2, 270, 269, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 270, 3, 2,
	2, 2, 272, 273, 3, 2, 2, 2, 273, 277, 3, 2, 2, 2, 274, 276, 5, 50, 26,
	2, 275, 274, 3, 2, 2, 2, 276, 279, 3, 2, 2, 2, 277, 275, 3, 2, 2, 2, 277,
	278, 3, 2, 2, 2, 278, 281, 3, 2, 2, 2, 279, 277, 3, 2, 2, 2, 280, 282,
	5, 54, 28, 2, 281, 280, 3, 2, 2, 2, 281, 282, 3, 2, 2, 2, 282, 49, 3, 2,
	2, 2, 283, 284, 7, 12, 2, 2, 284, 288, 7, 55, 2, 2, 285, 286, 5, 60, 31,
	2, 286, 287, 9, 6, 2, 2, 287, 289, 3, 2, 2, 2, 288, 285, 3, 2, 2, 2, 288,
	289, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 291, 7, 13, 2, 2, 291, 51,
	3, 2, 2, 2, 292, 293, 7, 17, 2, 2, 293, 294, 7, 53, 2, 2, 294, 296, 7,
	8, 2, 2, 295, 297, 7, 12, 2, 2, 296, 295, 3, 2, 2, 2, 296, 297, 3, 2, 2,
	2, 297, 298, 3, 2, 2, 2, 298, 300, 5, 46, 24, 2, 299, 301, 7, 13, 2, 2,
	300, 299, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 302, 3, 2, 2, 2, 302,
	303, 7, 9, 2, 2, 303, 53, 3, 2, 2, 2, 304, 305, 9, 7, 2, 2, 305, 306, 7,
	55, 2, 2, 306, 55, 3, 2, 2, 2, 307, 308, 5, 58, 30, 2, 308, 309, 5, 62,
	32, 2, 309, 318, 3, 2, 2, 2, 310, 318, 5, 62, 32, 2, 311, 312, 7, 24, 2,
	2, 312, 318, 5, 62, 32, 2, 313, 318, 5, 54, 28, 2, 314, 318, 5, 52, 27,
	2, 315, 318, 7, 23, 2, 2, 316, 318, 7, 27, 2, 2, 317, 307, 3, 2, 2, 2,
	317, 310, 3, 2, 2, 2, 317, 311, 3, 2, 2, 2, 317, 313, 3, 2, 2, 2, 317,
	314, 3, 2, 2, 2, 317, 315, 3, 2, 2, 2, 317, 316, 3, 2, 2, 2, 318, 57, 3,
	2, 2, 2, 319, 320, 9, 8, 2, 2, 320, 59, 3, 2, 2, 2, 321, 322, 9, 9, 2,
	2, 322, 61, 3, 2, 2, 2, 323, 327, 7, 55, 2, 2, 324, 326, 5, 64, 33, 2,
	325, 324, 3, 2, 2, 2, 326, 329, 3, 2, 2, 2, 327, 325, 3, 2, 2, 2, 327,
	328, 3, 2, 2, 2, 328, 340, 3, 2, 2, 2, 329, 327, 3, 2, 2, 2, 330, 331,
	7, 7, 2, 2, 331, 332, 5, 66, 34, 2, 332, 336, 7, 11, 2, 2, 333, 335, 5,
	64, 33, 2, 334, 333, 3, 2, 2, 2, 335, 338, 3, 2, 2, 2, 336, 334, 3, 2,
	2, 2, 336, 337, 3, 2, 2, 2, 337, 340, 3, 2, 2, 2, 338, 336, 3, 2, 2, 2,
	339, 323, 3, 2, 2, 2, 339, 330, 3, 2, 2, 2, 340, 63, 3, 2, 2, 2, 341, 342,
	7, 144, 2, 2, 342, 343, 5, 66, 34, 2, 343, 344, 7, 11, 2, 2, 344, 347,
	3, 2, 2, 2, 345, 347, 7, 145, 2, 2, 346, 341, 3, 2, 2, 2, 346, 345, 3,
	2, 2, 2, 347, 65, 3, 2, 2, 2, 348, 349, 9, 10, 2, 2, 349, 67, 3, 2, 2,
	2, 350, 351, 5, 62, 32, 2, 351, 352, 7, 17, 2, 2, 352, 353, 5, 70, 36,
	2, 353, 69, 3, 2, 2, 2, 354, 359, 5, 8, 5, 2, 355, 356, 7, 19, 2, 2, 356,
	358, 5, 8, 5, 2, 357, 355, 3, 2, 2, 2, 358, 361, 3, 2, 2, 2, 359, 357,
	3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360, 71, 3, 2, 2, 2, 361, 359, 3, 2,
	2, 2, 362, 363, 7, 39, 2, 2, 363, 364, 7, 142, 2, 2, 364, 365, 7, 141,
	2, 2, 365, 73, 3, 2, 2, 2, 366, 368, 7, 57, 2, 2, 367, 369, 7, 5, 2, 2,
	368, 367, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 75, 3, 2, 2, 2, 42, 79,
	88, 94, 99, 102, 117, 128, 133, 142, 154, 161, 174, 179, 183, 193, 200,
	212, 214, 218, 229, 232, 235, 239, 248, 253, 257, 266, 272, 277, 281, 288,
	296, 300, 317, 327, 336, 339, 346, 359, 368,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'null'", "'in'", "", "'...'", "", "'('", "')'", "'{'", "'}'", "'['",
	"']'", "'>'", "'~'", "'<'", "':'", "';'", "','", "'.'", "'$'", "'@'", "'&'",
	"'#'", "'::'", "'+'", "'*'", "'/'", "'-'", "", "'=='", "'>='", "'<='",
	"'!='", "'='", "'|='", "'~='", "'url'", "", "'@import'", "'@media'", "':extend'",
	"'!important'", "'@arguments'", "'@rest'", "'reference'", "'inline'", "'less'",
	"'css'", "'once'", "'multiple'", "'when'", "'not'", "'and'", "", "", "",
	"", "", "", "", "", "'color'", "'convert'", "'data-uri'", "'default'",
	"'unit'", "'get-unit'", "'svg-gradient'", "'escape'", "'e'", "", "'replace'",
	"'length'", "'extract'", "'ceil'", "'floor'", "'percentage'", "'round'",
	"'sqrt'", "'abs'", "'sin'", "'asin'", "'cos'", "'acos'", "'tan'", "'atan'",
	"'pi'", "'pow'", "'mod'", "'min'", "'max'", "'isnumber'", "'isstring'",
	"'iscolor'", "'iskeyword'", "'isurl'", "'ispixel'", "'isem'", "'ispercentage'",
	"'isunit'", "'rgb'", "'rgba'", "'argb'", "'hsl'", "'hsla'", "'hsv'", "'hsva'",
	"'hue'", "'saturation'", "'lightness'", "'hsvhue'", "'hsvsaturation'",
	"'hsvvalue'", "'red'", "'green'", "'blue'", "'alpha'", "'luma'", "'luminance'",
	"'saturate'", "'desaturate'", "'lighten'", "'darken'", "'fadein'", "'fadeout'",
	"'fade'", "'spin'", "'mix'", "'greyscale'", "'contrast'", "'multiply'",
	"'screen'", "'overlay'", "'softlight'", "'hardlight'", "'difference'",
	"'exclusion'", "'average'", "'negation'",
}
var symbolicNames = []string{
	"", "NULL", "IN", "Unit", "Ellipsis", "InterpolationStart", "LPAREN", "RPAREN",
	"BlockStart", "BlockEnd", "LBRACK", "RBRACK", "GT", "TIL", "LT", "COLON",
	"SEMI", "COMMA", "DOT", "DOLLAR", "AT", "PARENTREF", "HASH", "COLONCOLON",
	"PLUS", "TIMES", "DIV", "MINUS", "PERC", "EQEQ", "GTEQ", "LTEQ", "NOTEQ",
	"EQ", "PIPE_EQ", "TILD_EQ", "URL", "UrlStart", "IMPORT", "MEDIA", "EXTEND",
	"IMPORTANT", "ARGUMENTS", "REST", "REFERENCE", "INLINE", "LESS", "CSS",
	"ONCE", "MULTIPLE", "WHEN", "NOT", "AND", "Identifier", "StringLiteral",
	"Number", "Color", "WS", "SL_COMMENT", "COMMENT", "FUNCTION_NAME", "COLOR",
	"CONVERT", "DATA_URI", "DEFAULT", "UNIT", "GET_UNIT", "SVG_GRADIENT", "ESCAPE",
	"E", "FORMAT", "REPLACE", "LENGTH", "EXTRACT", "CEIL", "FLOOR", "PERCENTAGE",
	"ROUND", "SQRT", "ABS", "SIN", "ASIN", "COS", "ACOS", "TAN", "ATAN", "PI",
	"POW", "MOD", "MIN", "MAX", "ISNUMBER", "ISSTRING", "ISCOLOR", "ISKEYWORD",
	"ISURL", "ISPIXEL", "ISEM", "ISPERCENTAGE", "ISUNIT", "RGB", "RGBA", "ARGB",
	"HSL", "HSLA", "HSV", "HSVA", "HUE", "SATURATION", "LIGHTNESS", "HSVHUE",
	"HSVSATURATION", "HSVVALUE", "RED", "GREEN", "BLUE", "ALPHA", "LUMA", "LUMINANCE",
	"SATURATE", "DESATURATE", "LIGHTEN", "DARKEN", "FADEIN", "FADEOUT", "FADE",
	"SPIN", "MIX", "GREYSCALE", "CONTRAST", "MULTIPLY", "SCREEN", "OVERLAY",
	"SOFTLIGHT", "HARDLIGHT", "DIFFERENCE", "EXCLUSION", "AVERAGE", "NEGATION",
	"UrlEnd", "Url", "SPACE", "InterpolationStartAfter", "IdentifierAfter",
}

var ruleNames = []string{
	"stylesheet", "statement", "variableName", "commandStatement", "mathCharacter",
	"mathStatement", "expression", "function", "conditions", "condition", "conditionStatement",
	"variableDeclaration", "importDeclaration", "importOption", "referenceUrl",
	"mediaTypes", "ruleset", "block", "mixinDefinition", "mixinGuard", "mixinDefinitionParam",
	"mixinReference", "selectors", "selector", "attrib", "negation", "pseudo",
	"element", "selectorPrefix", "attribRelate", "identifier", "identifierPart",
	"identifierVariableName", "property", "values", "url", "measurement",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type LessParser struct {
	*antlr.BaseParser
}

func NewLessParser(input antlr.TokenStream) *LessParser {
	this := new(LessParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "LessParser.g4"

	return this
}

// LessParser tokens.
const (
	LessParserEOF                     = antlr.TokenEOF
	LessParserNULL                    = 1
	LessParserIN                      = 2
	LessParserUnit                    = 3
	LessParserEllipsis                = 4
	LessParserInterpolationStart      = 5
	LessParserLPAREN                  = 6
	LessParserRPAREN                  = 7
	LessParserBlockStart              = 8
	LessParserBlockEnd                = 9
	LessParserLBRACK                  = 10
	LessParserRBRACK                  = 11
	LessParserGT                      = 12
	LessParserTIL                     = 13
	LessParserLT                      = 14
	LessParserCOLON                   = 15
	LessParserSEMI                    = 16
	LessParserCOMMA                   = 17
	LessParserDOT                     = 18
	LessParserDOLLAR                  = 19
	LessParserAT                      = 20
	LessParserPARENTREF               = 21
	LessParserHASH                    = 22
	LessParserCOLONCOLON              = 23
	LessParserPLUS                    = 24
	LessParserTIMES                   = 25
	LessParserDIV                     = 26
	LessParserMINUS                   = 27
	LessParserPERC                    = 28
	LessParserEQEQ                    = 29
	LessParserGTEQ                    = 30
	LessParserLTEQ                    = 31
	LessParserNOTEQ                   = 32
	LessParserEQ                      = 33
	LessParserPIPE_EQ                 = 34
	LessParserTILD_EQ                 = 35
	LessParserURL                     = 36
	LessParserUrlStart                = 37
	LessParserIMPORT                  = 38
	LessParserMEDIA                   = 39
	LessParserEXTEND                  = 40
	LessParserIMPORTANT               = 41
	LessParserARGUMENTS               = 42
	LessParserREST                    = 43
	LessParserREFERENCE               = 44
	LessParserINLINE                  = 45
	LessParserLESS                    = 46
	LessParserCSS                     = 47
	LessParserONCE                    = 48
	LessParserMULTIPLE                = 49
	LessParserWHEN                    = 50
	LessParserNOT                     = 51
	LessParserAND                     = 52
	LessParserIdentifier              = 53
	LessParserStringLiteral           = 54
	LessParserNumber                  = 55
	LessParserColor                   = 56
	LessParserWS                      = 57
	LessParserSL_COMMENT              = 58
	LessParserCOMMENT                 = 59
	LessParserFUNCTION_NAME           = 60
	LessParserCOLOR                   = 61
	LessParserCONVERT                 = 62
	LessParserDATA_URI                = 63
	LessParserDEFAULT                 = 64
	LessParserUNIT                    = 65
	LessParserGET_UNIT                = 66
	LessParserSVG_GRADIENT            = 67
	LessParserESCAPE                  = 68
	LessParserE                       = 69
	LessParserFORMAT                  = 70
	LessParserREPLACE                 = 71
	LessParserLENGTH                  = 72
	LessParserEXTRACT                 = 73
	LessParserCEIL                    = 74
	LessParserFLOOR                   = 75
	LessParserPERCENTAGE              = 76
	LessParserROUND                   = 77
	LessParserSQRT                    = 78
	LessParserABS                     = 79
	LessParserSIN                     = 80
	LessParserASIN                    = 81
	LessParserCOS                     = 82
	LessParserACOS                    = 83
	LessParserTAN                     = 84
	LessParserATAN                    = 85
	LessParserPI                      = 86
	LessParserPOW                     = 87
	LessParserMOD                     = 88
	LessParserMIN                     = 89
	LessParserMAX                     = 90
	LessParserISNUMBER                = 91
	LessParserISSTRING                = 92
	LessParserISCOLOR                 = 93
	LessParserISKEYWORD               = 94
	LessParserISURL                   = 95
	LessParserISPIXEL                 = 96
	LessParserISEM                    = 97
	LessParserISPERCENTAGE            = 98
	LessParserISUNIT                  = 99
	LessParserRGB                     = 100
	LessParserRGBA                    = 101
	LessParserARGB                    = 102
	LessParserHSL                     = 103
	LessParserHSLA                    = 104
	LessParserHSV                     = 105
	LessParserHSVA                    = 106
	LessParserHUE                     = 107
	LessParserSATURATION              = 108
	LessParserLIGHTNESS               = 109
	LessParserHSVHUE                  = 110
	LessParserHSVSATURATION           = 111
	LessParserHSVVALUE                = 112
	LessParserRED                     = 113
	LessParserGREEN                   = 114
	LessParserBLUE                    = 115
	LessParserALPHA                   = 116
	LessParserLUMA                    = 117
	LessParserLUMINANCE               = 118
	LessParserSATURATE                = 119
	LessParserDESATURATE              = 120
	LessParserLIGHTEN                 = 121
	LessParserDARKEN                  = 122
	LessParserFADEIN                  = 123
	LessParserFADEOUT                 = 124
	LessParserFADE                    = 125
	LessParserSPIN                    = 126
	LessParserMIX                     = 127
	LessParserGREYSCALE               = 128
	LessParserCONTRAST                = 129
	LessParserMULTIPLY                = 130
	LessParserSCREEN                  = 131
	LessParserOVERLAY                 = 132
	LessParserSOFTLIGHT               = 133
	LessParserHARDLIGHT               = 134
	LessParserDIFFERENCE              = 135
	LessParserEXCLUSION               = 136
	LessParserAVERAGE                 = 137
	LessParserNEGATION                = 138
	LessParserUrlEnd                  = 139
	LessParserUrl                     = 140
	LessParserSPACE                   = 141
	LessParserInterpolationStartAfter = 142
	LessParserIdentifierAfter         = 143
)

// LessParser rules.
const (
	LessParserRULE_stylesheet             = 0
	LessParserRULE_statement              = 1
	LessParserRULE_variableName           = 2
	LessParserRULE_commandStatement       = 3
	LessParserRULE_mathCharacter          = 4
	LessParserRULE_mathStatement          = 5
	LessParserRULE_expression             = 6
	LessParserRULE_function               = 7
	LessParserRULE_conditions             = 8
	LessParserRULE_condition              = 9
	LessParserRULE_conditionStatement     = 10
	LessParserRULE_variableDeclaration    = 11
	LessParserRULE_importDeclaration      = 12
	LessParserRULE_importOption           = 13
	LessParserRULE_referenceUrl           = 14
	LessParserRULE_mediaTypes             = 15
	LessParserRULE_ruleset                = 16
	LessParserRULE_block                  = 17
	LessParserRULE_mixinDefinition        = 18
	LessParserRULE_mixinGuard             = 19
	LessParserRULE_mixinDefinitionParam   = 20
	LessParserRULE_mixinReference         = 21
	LessParserRULE_selectors              = 22
	LessParserRULE_selector               = 23
	LessParserRULE_attrib                 = 24
	LessParserRULE_negation               = 25
	LessParserRULE_pseudo                 = 26
	LessParserRULE_element                = 27
	LessParserRULE_selectorPrefix         = 28
	LessParserRULE_attribRelate           = 29
	LessParserRULE_identifier             = 30
	LessParserRULE_identifierPart         = 31
	LessParserRULE_identifierVariableName = 32
	LessParserRULE_property               = 33
	LessParserRULE_values                 = 34
	LessParserRULE_url                    = 35
	LessParserRULE_measurement            = 36
)

// IStylesheetContext is an interface to support dynamic dispatch.
type IStylesheetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStylesheetContext differentiates from other interfaces.
	IsStylesheetContext()
}

type StylesheetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStylesheetContext() *StylesheetContext {
	var p = new(StylesheetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_stylesheet
	return p
}

func (*StylesheetContext) IsStylesheetContext() {}

func NewStylesheetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StylesheetContext {
	var p = new(StylesheetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_stylesheet

	return p
}

func (s *StylesheetContext) GetParser() antlr.Parser { return s.parser }

func (s *StylesheetContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StylesheetContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StylesheetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StylesheetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StylesheetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterStylesheet(s)
	}
}

func (s *StylesheetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitStylesheet(s)
	}
}

func (p *LessParser) Stylesheet() (localctx IStylesheetContext) {
	localctx = NewStylesheetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LessParserRULE_stylesheet)
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
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LessParserInterpolationStart)|(1<<LessParserGT)|(1<<LessParserTIL)|(1<<LessParserCOLON)|(1<<LessParserAT)|(1<<LessParserPARENTREF)|(1<<LessParserHASH)|(1<<LessParserCOLONCOLON)|(1<<LessParserPLUS)|(1<<LessParserTIMES))) != 0) || _la == LessParserIMPORT || _la == LessParserIdentifier {
		{
			p.SetState(74)
			p.Statement()
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = LessParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ImportDeclaration() IImportDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportDeclarationContext)
}

func (s *StatementContext) Ruleset() IRulesetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRulesetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRulesetContext)
}

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(LessParserSEMI, 0)
}

func (s *StatementContext) MixinDefinition() IMixinDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMixinDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMixinDefinitionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *LessParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LessParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(80)
			p.ImportDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(81)
			p.Ruleset()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(82)
			p.VariableDeclaration()
		}
		{
			p.SetState(83)
			p.Match(LessParserSEMI)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(85)
			p.MixinDefinition()
		}

	}

	return localctx
}

// IVariableNameContext is an interface to support dynamic dispatch.
type IVariableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableNameContext differentiates from other interfaces.
	IsVariableNameContext()
}

type VariableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableNameContext() *VariableNameContext {
	var p = new(VariableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_variableName
	return p
}

func (*VariableNameContext) IsVariableNameContext() {}

func NewVariableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableNameContext {
	var p = new(VariableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_variableName

	return p
}

func (s *VariableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableNameContext) AT() antlr.TerminalNode {
	return s.GetToken(LessParserAT, 0)
}

func (s *VariableNameContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *VariableNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LessParserIdentifier, 0)
}

func (s *VariableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterVariableName(s)
	}
}

func (s *VariableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitVariableName(s)
	}
}

func (p *LessParser) VariableName() (localctx IVariableNameContext) {
	localctx = NewVariableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LessParserRULE_variableName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(88)
			p.Match(LessParserAT)
		}
		{
			p.SetState(89)
			p.VariableName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(90)
			p.Match(LessParserAT)
		}
		{
			p.SetState(91)
			p.Match(LessParserIdentifier)
		}

	}

	return localctx
}

// ICommandStatementContext is an interface to support dynamic dispatch.
type ICommandStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandStatementContext differentiates from other interfaces.
	IsCommandStatementContext()
}

type CommandStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandStatementContext() *CommandStatementContext {
	var p = new(CommandStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_commandStatement
	return p
}

func (*CommandStatementContext) IsCommandStatementContext() {}

func NewCommandStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandStatementContext {
	var p = new(CommandStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_commandStatement

	return p
}

func (s *CommandStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandStatementContext) MathStatement() IMathStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathStatementContext)
}

func (s *CommandStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CommandStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CommandStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterCommandStatement(s)
	}
}

func (s *CommandStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitCommandStatement(s)
	}
}

func (p *LessParser) CommandStatement() (localctx ICommandStatementContext) {
	localctx = NewCommandStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LessParserRULE_commandStatement)
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LessParserInterpolationStart || _la == LessParserAT || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(LessParserUrlStart-37))|(1<<(LessParserIdentifier-37))|(1<<(LessParserStringLiteral-37))|(1<<(LessParserNumber-37))|(1<<(LessParserColor-37)))) != 0) {
		{
			p.SetState(94)
			p.Expression()
		}

		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LessParserPLUS)|(1<<LessParserTIMES)|(1<<LessParserDIV)|(1<<LessParserMINUS)|(1<<LessParserPERC))) != 0 {
		{
			p.SetState(99)
			p.MathStatement()
		}

	}

	return localctx
}

// IMathCharacterContext is an interface to support dynamic dispatch.
type IMathCharacterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathCharacterContext differentiates from other interfaces.
	IsMathCharacterContext()
}

type MathCharacterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathCharacterContext() *MathCharacterContext {
	var p = new(MathCharacterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mathCharacter
	return p
}

func (*MathCharacterContext) IsMathCharacterContext() {}

func NewMathCharacterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathCharacterContext {
	var p = new(MathCharacterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mathCharacter

	return p
}

func (s *MathCharacterContext) GetParser() antlr.Parser { return s.parser }

func (s *MathCharacterContext) TIMES() antlr.TerminalNode {
	return s.GetToken(LessParserTIMES, 0)
}

func (s *MathCharacterContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LessParserPLUS, 0)
}

func (s *MathCharacterContext) DIV() antlr.TerminalNode {
	return s.GetToken(LessParserDIV, 0)
}

func (s *MathCharacterContext) MINUS() antlr.TerminalNode {
	return s.GetToken(LessParserMINUS, 0)
}

func (s *MathCharacterContext) PERC() antlr.TerminalNode {
	return s.GetToken(LessParserPERC, 0)
}

func (s *MathCharacterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathCharacterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathCharacterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMathCharacter(s)
	}
}

func (s *MathCharacterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMathCharacter(s)
	}
}

func (p *LessParser) MathCharacter() (localctx IMathCharacterContext) {
	localctx = NewMathCharacterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LessParserRULE_mathCharacter)
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
		p.SetState(102)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LessParserPLUS)|(1<<LessParserTIMES)|(1<<LessParserDIV)|(1<<LessParserMINUS)|(1<<LessParserPERC))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMathStatementContext is an interface to support dynamic dispatch.
type IMathStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathStatementContext differentiates from other interfaces.
	IsMathStatementContext()
}

type MathStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathStatementContext() *MathStatementContext {
	var p = new(MathStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mathStatement
	return p
}

func (*MathStatementContext) IsMathStatementContext() {}

func NewMathStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathStatementContext {
	var p = new(MathStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mathStatement

	return p
}

func (s *MathStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *MathStatementContext) MathCharacter() IMathCharacterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathCharacterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathCharacterContext)
}

func (s *MathStatementContext) CommandStatement() ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *MathStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMathStatement(s)
	}
}

func (s *MathStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMathStatement(s)
	}
}

func (p *LessParser) MathStatement() (localctx IMathStatementContext) {
	localctx = NewMathStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LessParserRULE_mathStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(104)
		p.MathCharacter()
	}
	{
		p.SetState(105)
		p.CommandStatement()
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
	p.RuleIndex = LessParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Measurement() IMeasurementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMeasurementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMeasurementContext)
}

func (s *ExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ExpressionContext) IMPORTANT() antlr.TerminalNode {
	return s.GetToken(LessParserIMPORTANT, 0)
}

func (s *ExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserLPAREN, 0)
}

func (s *ExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserRPAREN, 0)
}

func (s *ExpressionContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *ExpressionContext) Color() antlr.TerminalNode {
	return s.GetToken(LessParserColor, 0)
}

func (s *ExpressionContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(LessParserStringLiteral, 0)
}

func (s *ExpressionContext) Url() IUrlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUrlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUrlContext)
}

func (s *ExpressionContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *LessParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LessParserRULE_expression)
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

	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.Measurement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.Identifier()
		}
		{
			p.SetState(109)
			p.Match(LessParserIMPORTANT)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.Identifier()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(112)
			p.Identifier()
		}
		{
			p.SetState(113)
			p.Match(LessParserLPAREN)
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LessParserInterpolationStart || _la == LessParserAT || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(LessParserUrlStart-37))|(1<<(LessParserIdentifier-37))|(1<<(LessParserStringLiteral-37))|(1<<(LessParserNumber-37))|(1<<(LessParserColor-37)))) != 0) {
			{
				p.SetState(114)
				p.Values()
			}

		}
		{
			p.SetState(117)
			p.Match(LessParserRPAREN)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(119)
			p.Match(LessParserColor)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(120)
			p.Match(LessParserStringLiteral)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(121)
			p.Url()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(122)
			p.VariableName()
		}
		{
			p.SetState(123)
			p.Match(LessParserIMPORTANT)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(125)
			p.VariableName()
		}

	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) FUNCTION_NAME() antlr.TerminalNode {
	return s.GetToken(LessParserFUNCTION_NAME, 0)
}

func (s *FunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserLPAREN, 0)
}

func (s *FunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserRPAREN, 0)
}

func (s *FunctionContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *LessParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LessParserRULE_function)
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
		p.Match(LessParserFUNCTION_NAME)
	}
	{
		p.SetState(129)
		p.Match(LessParserLPAREN)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserInterpolationStart || _la == LessParserAT || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(LessParserUrlStart-37))|(1<<(LessParserIdentifier-37))|(1<<(LessParserStringLiteral-37))|(1<<(LessParserNumber-37))|(1<<(LessParserColor-37)))) != 0) {
		{
			p.SetState(130)
			p.Values()
		}

	}
	{
		p.SetState(133)
		p.Match(LessParserRPAREN)
	}

	return localctx
}

// IConditionsContext is an interface to support dynamic dispatch.
type IConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionsContext differentiates from other interfaces.
	IsConditionsContext()
}

type ConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionsContext() *ConditionsContext {
	var p = new(ConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_conditions
	return p
}

func (*ConditionsContext) IsConditionsContext() {}

func NewConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsContext {
	var p = new(ConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_conditions

	return p
}

func (s *ConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsContext) AllCondition() []IConditionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConditionContext)(nil)).Elem())
	var tst = make([]IConditionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConditionContext)
		}
	}

	return tst
}

func (s *ConditionsContext) Condition(i int) IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ConditionsContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(LessParserAND)
}

func (s *ConditionsContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(LessParserAND, i)
}

func (s *ConditionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LessParserCOMMA)
}

func (s *ConditionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LessParserCOMMA, i)
}

func (s *ConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterConditions(s)
	}
}

func (s *ConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitConditions(s)
	}
}

func (p *LessParser) Conditions() (localctx IConditionsContext) {
	localctx = NewConditionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LessParserRULE_conditions)
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
		p.Condition()
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LessParserCOMMA || _la == LessParserAND {
		{
			p.SetState(136)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LessParserCOMMA || _la == LessParserAND) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(137)
			p.Condition()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserLPAREN, 0)
}

func (s *ConditionContext) ConditionStatement() IConditionStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionStatementContext)
}

func (s *ConditionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserRPAREN, 0)
}

func (s *ConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(LessParserNOT, 0)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (p *LessParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LessParserRULE_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LessParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.Match(LessParserLPAREN)
		}
		{
			p.SetState(144)
			p.ConditionStatement()
		}
		{
			p.SetState(145)
			p.Match(LessParserRPAREN)
		}

	case LessParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(147)
			p.Match(LessParserNOT)
		}
		{
			p.SetState(148)
			p.Match(LessParserLPAREN)
		}
		{
			p.SetState(149)
			p.ConditionStatement()
		}
		{
			p.SetState(150)
			p.Match(LessParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConditionStatementContext is an interface to support dynamic dispatch.
type IConditionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionStatementContext differentiates from other interfaces.
	IsConditionStatementContext()
}

type ConditionStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionStatementContext() *ConditionStatementContext {
	var p = new(ConditionStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_conditionStatement
	return p
}

func (*ConditionStatementContext) IsConditionStatementContext() {}

func NewConditionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionStatementContext {
	var p = new(ConditionStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_conditionStatement

	return p
}

func (s *ConditionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionStatementContext) AllCommandStatement() []ICommandStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem())
	var tst = make([]ICommandStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandStatementContext)
		}
	}

	return tst
}

func (s *ConditionStatementContext) CommandStatement(i int) ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *ConditionStatementContext) EQ() antlr.TerminalNode {
	return s.GetToken(LessParserEQ, 0)
}

func (s *ConditionStatementContext) LT() antlr.TerminalNode {
	return s.GetToken(LessParserLT, 0)
}

func (s *ConditionStatementContext) GT() antlr.TerminalNode {
	return s.GetToken(LessParserGT, 0)
}

func (s *ConditionStatementContext) GTEQ() antlr.TerminalNode {
	return s.GetToken(LessParserGTEQ, 0)
}

func (s *ConditionStatementContext) LTEQ() antlr.TerminalNode {
	return s.GetToken(LessParserLTEQ, 0)
}

func (s *ConditionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterConditionStatement(s)
	}
}

func (s *ConditionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitConditionStatement(s)
	}
}

func (p *LessParser) ConditionStatement() (localctx IConditionStatementContext) {
	localctx = NewConditionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LessParserRULE_conditionStatement)
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

	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(154)
			p.CommandStatement()
		}
		{
			p.SetState(155)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-12)&-(0x1f+1)) == 0 && ((1<<uint((_la-12)))&((1<<(LessParserGT-12))|(1<<(LessParserLT-12))|(1<<(LessParserGTEQ-12))|(1<<(LessParserLTEQ-12))|(1<<(LessParserEQ-12)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(156)
			p.CommandStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.CommandStatement()
		}

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
	p.RuleIndex = LessParserRULE_variableDeclaration
	return p
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *VariableDeclarationContext) COLON() antlr.TerminalNode {
	return s.GetToken(LessParserCOLON, 0)
}

func (s *VariableDeclarationContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (p *LessParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LessParserRULE_variableDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.VariableName()
	}
	{
		p.SetState(162)
		p.Match(LessParserCOLON)
	}
	{
		p.SetState(163)
		p.Values()
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
	p.RuleIndex = LessParserRULE_importDeclaration
	return p
}

func (*ImportDeclarationContext) IsImportDeclarationContext() {}

func NewImportDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclarationContext {
	var p = new(ImportDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_importDeclaration

	return p
}

func (s *ImportDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclarationContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(LessParserIMPORT, 0)
}

func (s *ImportDeclarationContext) ReferenceUrl() IReferenceUrlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferenceUrlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferenceUrlContext)
}

func (s *ImportDeclarationContext) SEMI() antlr.TerminalNode {
	return s.GetToken(LessParserSEMI, 0)
}

func (s *ImportDeclarationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserLPAREN, 0)
}

func (s *ImportDeclarationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserRPAREN, 0)
}

func (s *ImportDeclarationContext) MediaTypes() IMediaTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMediaTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMediaTypesContext)
}

func (s *ImportDeclarationContext) AllImportOption() []IImportOptionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportOptionContext)(nil)).Elem())
	var tst = make([]IImportOptionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportOptionContext)
		}
	}

	return tst
}

func (s *ImportDeclarationContext) ImportOption(i int) IImportOptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportOptionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportOptionContext)
}

func (s *ImportDeclarationContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LessParserCOMMA)
}

func (s *ImportDeclarationContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LessParserCOMMA, i)
}

func (s *ImportDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterImportDeclaration(s)
	}
}

func (s *ImportDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitImportDeclaration(s)
	}
}

func (p *LessParser) ImportDeclaration() (localctx IImportDeclarationContext) {
	localctx = NewImportDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LessParserRULE_importDeclaration)
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
		p.SetState(165)
		p.Match(LessParserIMPORT)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserLPAREN {
		{
			p.SetState(166)
			p.Match(LessParserLPAREN)
		}

		{
			p.SetState(167)
			p.ImportOption()
		}
		p.SetState(172)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LessParserCOMMA {
			{
				p.SetState(168)
				p.Match(LessParserCOMMA)
			}
			{
				p.SetState(169)
				p.ImportOption()
			}

			p.SetState(174)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		{
			p.SetState(175)
			p.Match(LessParserRPAREN)
		}

	}
	{
		p.SetState(179)
		p.ReferenceUrl()
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserIdentifier {
		{
			p.SetState(180)
			p.MediaTypes()
		}

	}
	{
		p.SetState(183)
		p.Match(LessParserSEMI)
	}

	return localctx
}

// IImportOptionContext is an interface to support dynamic dispatch.
type IImportOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportOptionContext differentiates from other interfaces.
	IsImportOptionContext()
}

type ImportOptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportOptionContext() *ImportOptionContext {
	var p = new(ImportOptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_importOption
	return p
}

func (*ImportOptionContext) IsImportOptionContext() {}

func NewImportOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportOptionContext {
	var p = new(ImportOptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_importOption

	return p
}

func (s *ImportOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportOptionContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(LessParserREFERENCE, 0)
}

func (s *ImportOptionContext) INLINE() antlr.TerminalNode {
	return s.GetToken(LessParserINLINE, 0)
}

func (s *ImportOptionContext) LESS() antlr.TerminalNode {
	return s.GetToken(LessParserLESS, 0)
}

func (s *ImportOptionContext) CSS() antlr.TerminalNode {
	return s.GetToken(LessParserCSS, 0)
}

func (s *ImportOptionContext) ONCE() antlr.TerminalNode {
	return s.GetToken(LessParserONCE, 0)
}

func (s *ImportOptionContext) MULTIPLE() antlr.TerminalNode {
	return s.GetToken(LessParserMULTIPLE, 0)
}

func (s *ImportOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportOptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterImportOption(s)
	}
}

func (s *ImportOptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitImportOption(s)
	}
}

func (p *LessParser) ImportOption() (localctx IImportOptionContext) {
	localctx = NewImportOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LessParserRULE_importOption)
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
		p.SetState(185)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(LessParserREFERENCE-44))|(1<<(LessParserINLINE-44))|(1<<(LessParserLESS-44))|(1<<(LessParserCSS-44))|(1<<(LessParserONCE-44))|(1<<(LessParserMULTIPLE-44)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IReferenceUrlContext is an interface to support dynamic dispatch.
type IReferenceUrlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReferenceUrlContext differentiates from other interfaces.
	IsReferenceUrlContext()
}

type ReferenceUrlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReferenceUrlContext() *ReferenceUrlContext {
	var p = new(ReferenceUrlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_referenceUrl
	return p
}

func (*ReferenceUrlContext) IsReferenceUrlContext() {}

func NewReferenceUrlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferenceUrlContext {
	var p = new(ReferenceUrlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_referenceUrl

	return p
}

func (s *ReferenceUrlContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferenceUrlContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(LessParserStringLiteral, 0)
}

func (s *ReferenceUrlContext) UrlStart() antlr.TerminalNode {
	return s.GetToken(LessParserUrlStart, 0)
}

func (s *ReferenceUrlContext) Url() antlr.TerminalNode {
	return s.GetToken(LessParserUrl, 0)
}

func (s *ReferenceUrlContext) UrlEnd() antlr.TerminalNode {
	return s.GetToken(LessParserUrlEnd, 0)
}

func (s *ReferenceUrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferenceUrlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferenceUrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterReferenceUrl(s)
	}
}

func (s *ReferenceUrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitReferenceUrl(s)
	}
}

func (p *LessParser) ReferenceUrl() (localctx IReferenceUrlContext) {
	localctx = NewReferenceUrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LessParserRULE_referenceUrl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LessParserStringLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Match(LessParserStringLiteral)
		}

	case LessParserUrlStart:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(188)
			p.Match(LessParserUrlStart)
		}
		{
			p.SetState(189)
			p.Match(LessParserUrl)
		}
		{
			p.SetState(190)
			p.Match(LessParserUrlEnd)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMediaTypesContext is an interface to support dynamic dispatch.
type IMediaTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMediaTypesContext differentiates from other interfaces.
	IsMediaTypesContext()
}

type MediaTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMediaTypesContext() *MediaTypesContext {
	var p = new(MediaTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mediaTypes
	return p
}

func (*MediaTypesContext) IsMediaTypesContext() {}

func NewMediaTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MediaTypesContext {
	var p = new(MediaTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mediaTypes

	return p
}

func (s *MediaTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *MediaTypesContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(LessParserIdentifier)
}

func (s *MediaTypesContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(LessParserIdentifier, i)
}

func (s *MediaTypesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LessParserCOMMA)
}

func (s *MediaTypesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LessParserCOMMA, i)
}

func (s *MediaTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MediaTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MediaTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMediaTypes(s)
	}
}

func (s *MediaTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMediaTypes(s)
	}
}

func (p *LessParser) MediaTypes() (localctx IMediaTypesContext) {
	localctx = NewMediaTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LessParserRULE_mediaTypes)
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
		p.SetState(193)
		p.Match(LessParserIdentifier)
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LessParserCOMMA {
		{
			p.SetState(194)
			p.Match(LessParserCOMMA)
		}
		{
			p.SetState(195)
			p.Match(LessParserIdentifier)
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRulesetContext is an interface to support dynamic dispatch.
type IRulesetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulesetContext differentiates from other interfaces.
	IsRulesetContext()
}

type RulesetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulesetContext() *RulesetContext {
	var p = new(RulesetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_ruleset
	return p
}

func (*RulesetContext) IsRulesetContext() {}

func NewRulesetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulesetContext {
	var p = new(RulesetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_ruleset

	return p
}

func (s *RulesetContext) GetParser() antlr.Parser { return s.parser }

func (s *RulesetContext) Selectors() ISelectorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorsContext)
}

func (s *RulesetContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RulesetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulesetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulesetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterRuleset(s)
	}
}

func (s *RulesetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitRuleset(s)
	}
}

func (p *LessParser) Ruleset() (localctx IRulesetContext) {
	localctx = NewRulesetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LessParserRULE_ruleset)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Selectors()
	}
	{
		p.SetState(202)
		p.Block()
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
	p.RuleIndex = LessParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) BlockStart() antlr.TerminalNode {
	return s.GetToken(LessParserBlockStart, 0)
}

func (s *BlockContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(LessParserBlockEnd, 0)
}

func (s *BlockContext) AllProperty() []IPropertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyContext)(nil)).Elem())
	var tst = make([]IPropertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyContext)
		}
	}

	return tst
}

func (s *BlockContext) Property(i int) IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *BlockContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(LessParserSEMI)
}

func (s *BlockContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(LessParserSEMI, i)
}

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

func (s *BlockContext) AllMixinReference() []IMixinReferenceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMixinReferenceContext)(nil)).Elem())
	var tst = make([]IMixinReferenceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMixinReferenceContext)
		}
	}

	return tst
}

func (s *BlockContext) MixinReference(i int) IMixinReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMixinReferenceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMixinReferenceContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *LessParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LessParserRULE_block)
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
		p.SetState(204)
		p.Match(LessParserBlockStart)
	}
	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(210)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(205)
					p.Property()
				}
				{
					p.SetState(206)
					p.Match(LessParserSEMI)
				}

			case 2:
				{
					p.SetState(208)
					p.Statement()
				}

			case 3:
				{
					p.SetState(209)
					p.MixinReference()
				}

			}

		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserInterpolationStart || _la == LessParserIdentifier {
		{
			p.SetState(215)
			p.Property()
		}

	}
	{
		p.SetState(218)
		p.Match(LessParserBlockEnd)
	}

	return localctx
}

// IMixinDefinitionContext is an interface to support dynamic dispatch.
type IMixinDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMixinDefinitionContext differentiates from other interfaces.
	IsMixinDefinitionContext()
}

type MixinDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixinDefinitionContext() *MixinDefinitionContext {
	var p = new(MixinDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mixinDefinition
	return p
}

func (*MixinDefinitionContext) IsMixinDefinitionContext() {}

func NewMixinDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MixinDefinitionContext {
	var p = new(MixinDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mixinDefinition

	return p
}

func (s *MixinDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MixinDefinitionContext) Selectors() ISelectorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorsContext)
}

func (s *MixinDefinitionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserLPAREN, 0)
}

func (s *MixinDefinitionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserRPAREN, 0)
}

func (s *MixinDefinitionContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *MixinDefinitionContext) AllMixinDefinitionParam() []IMixinDefinitionParamContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMixinDefinitionParamContext)(nil)).Elem())
	var tst = make([]IMixinDefinitionParamContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMixinDefinitionParamContext)
		}
	}

	return tst
}

func (s *MixinDefinitionContext) MixinDefinitionParam(i int) IMixinDefinitionParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMixinDefinitionParamContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMixinDefinitionParamContext)
}

func (s *MixinDefinitionContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(LessParserEllipsis, 0)
}

func (s *MixinDefinitionContext) MixinGuard() IMixinGuardContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMixinGuardContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMixinGuardContext)
}

func (s *MixinDefinitionContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(LessParserSEMI)
}

func (s *MixinDefinitionContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(LessParserSEMI, i)
}

func (s *MixinDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MixinDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MixinDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMixinDefinition(s)
	}
}

func (s *MixinDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMixinDefinition(s)
	}
}

func (p *LessParser) MixinDefinition() (localctx IMixinDefinitionContext) {
	localctx = NewMixinDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LessParserRULE_mixinDefinition)
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
		p.Selectors()
	}
	{
		p.SetState(221)
		p.Match(LessParserLPAREN)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserAT {
		{
			p.SetState(222)
			p.MixinDefinitionParam()
		}
		p.SetState(227)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LessParserSEMI {
			{
				p.SetState(223)
				p.Match(LessParserSEMI)
			}
			{
				p.SetState(224)
				p.MixinDefinitionParam()
			}

			p.SetState(229)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserEllipsis {
		{
			p.SetState(232)
			p.Match(LessParserEllipsis)
		}

	}
	{
		p.SetState(235)
		p.Match(LessParserRPAREN)
	}
	p.SetState(237)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserWHEN {
		{
			p.SetState(236)
			p.MixinGuard()
		}

	}
	{
		p.SetState(239)
		p.Block()
	}

	return localctx
}

// IMixinGuardContext is an interface to support dynamic dispatch.
type IMixinGuardContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMixinGuardContext differentiates from other interfaces.
	IsMixinGuardContext()
}

type MixinGuardContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixinGuardContext() *MixinGuardContext {
	var p = new(MixinGuardContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mixinGuard
	return p
}

func (*MixinGuardContext) IsMixinGuardContext() {}

func NewMixinGuardContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MixinGuardContext {
	var p = new(MixinGuardContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mixinGuard

	return p
}

func (s *MixinGuardContext) GetParser() antlr.Parser { return s.parser }

func (s *MixinGuardContext) WHEN() antlr.TerminalNode {
	return s.GetToken(LessParserWHEN, 0)
}

func (s *MixinGuardContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *MixinGuardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MixinGuardContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MixinGuardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMixinGuard(s)
	}
}

func (s *MixinGuardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMixinGuard(s)
	}
}

func (p *LessParser) MixinGuard() (localctx IMixinGuardContext) {
	localctx = NewMixinGuardContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LessParserRULE_mixinGuard)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(LessParserWHEN)
	}
	{
		p.SetState(242)
		p.Conditions()
	}

	return localctx
}

// IMixinDefinitionParamContext is an interface to support dynamic dispatch.
type IMixinDefinitionParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMixinDefinitionParamContext differentiates from other interfaces.
	IsMixinDefinitionParamContext()
}

type MixinDefinitionParamContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixinDefinitionParamContext() *MixinDefinitionParamContext {
	var p = new(MixinDefinitionParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mixinDefinitionParam
	return p
}

func (*MixinDefinitionParamContext) IsMixinDefinitionParamContext() {}

func NewMixinDefinitionParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MixinDefinitionParamContext {
	var p = new(MixinDefinitionParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mixinDefinitionParam

	return p
}

func (s *MixinDefinitionParamContext) GetParser() antlr.Parser { return s.parser }

func (s *MixinDefinitionParamContext) VariableName() IVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableNameContext)
}

func (s *MixinDefinitionParamContext) VariableDeclaration() IVariableDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *MixinDefinitionParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MixinDefinitionParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MixinDefinitionParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMixinDefinitionParam(s)
	}
}

func (s *MixinDefinitionParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMixinDefinitionParam(s)
	}
}

func (p *LessParser) MixinDefinitionParam() (localctx IMixinDefinitionParamContext) {
	localctx = NewMixinDefinitionParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LessParserRULE_mixinDefinitionParam)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(244)
			p.VariableName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(245)
			p.VariableDeclaration()
		}

	}

	return localctx
}

// IMixinReferenceContext is an interface to support dynamic dispatch.
type IMixinReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMixinReferenceContext differentiates from other interfaces.
	IsMixinReferenceContext()
}

type MixinReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMixinReferenceContext() *MixinReferenceContext {
	var p = new(MixinReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_mixinReference
	return p
}

func (*MixinReferenceContext) IsMixinReferenceContext() {}

func NewMixinReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MixinReferenceContext {
	var p = new(MixinReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_mixinReference

	return p
}

func (s *MixinReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *MixinReferenceContext) Selector() ISelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *MixinReferenceContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserLPAREN, 0)
}

func (s *MixinReferenceContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserRPAREN, 0)
}

func (s *MixinReferenceContext) SEMI() antlr.TerminalNode {
	return s.GetToken(LessParserSEMI, 0)
}

func (s *MixinReferenceContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *MixinReferenceContext) IMPORTANT() antlr.TerminalNode {
	return s.GetToken(LessParserIMPORTANT, 0)
}

func (s *MixinReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MixinReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MixinReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMixinReference(s)
	}
}

func (s *MixinReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMixinReference(s)
	}
}

func (p *LessParser) MixinReference() (localctx IMixinReferenceContext) {
	localctx = NewMixinReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LessParserRULE_mixinReference)
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
		p.Selector()
	}
	{
		p.SetState(249)
		p.Match(LessParserLPAREN)
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserInterpolationStart || _la == LessParserAT || (((_la-37)&-(0x1f+1)) == 0 && ((1<<uint((_la-37)))&((1<<(LessParserUrlStart-37))|(1<<(LessParserIdentifier-37))|(1<<(LessParserStringLiteral-37))|(1<<(LessParserNumber-37))|(1<<(LessParserColor-37)))) != 0) {
		{
			p.SetState(250)
			p.Values()
		}

	}
	{
		p.SetState(253)
		p.Match(LessParserRPAREN)
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserIMPORTANT {
		{
			p.SetState(254)
			p.Match(LessParserIMPORTANT)
		}

	}
	{
		p.SetState(257)
		p.Match(LessParserSEMI)
	}

	return localctx
}

// ISelectorsContext is an interface to support dynamic dispatch.
type ISelectorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorsContext differentiates from other interfaces.
	IsSelectorsContext()
}

type SelectorsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorsContext() *SelectorsContext {
	var p = new(SelectorsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_selectors
	return p
}

func (*SelectorsContext) IsSelectorsContext() {}

func NewSelectorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorsContext {
	var p = new(SelectorsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_selectors

	return p
}

func (s *SelectorsContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorsContext) AllSelector() []ISelectorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectorContext)(nil)).Elem())
	var tst = make([]ISelectorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectorContext)
		}
	}

	return tst
}

func (s *SelectorsContext) Selector(i int) ISelectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectorContext)
}

func (s *SelectorsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LessParserCOMMA)
}

func (s *SelectorsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LessParserCOMMA, i)
}

func (s *SelectorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterSelectors(s)
	}
}

func (s *SelectorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitSelectors(s)
	}
}

func (p *LessParser) Selectors() (localctx ISelectorsContext) {
	localctx = NewSelectorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LessParserRULE_selectors)
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
		p.SetState(259)
		p.Selector()
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LessParserCOMMA {
		{
			p.SetState(260)
			p.Match(LessParserCOMMA)
		}
		{
			p.SetState(261)
			p.Selector()
		}

		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectorContext is an interface to support dynamic dispatch.
type ISelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorContext differentiates from other interfaces.
	IsSelectorContext()
}

type SelectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorContext() *SelectorContext {
	var p = new(SelectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_selector
	return p
}

func (*SelectorContext) IsSelectorContext() {}

func NewSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorContext {
	var p = new(SelectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_selector

	return p
}

func (s *SelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorContext) AllElement() []IElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElementContext)(nil)).Elem())
	var tst = make([]IElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElementContext)
		}
	}

	return tst
}

func (s *SelectorContext) Element(i int) IElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElementContext)
}

func (s *SelectorContext) AllAttrib() []IAttribContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttribContext)(nil)).Elem())
	var tst = make([]IAttribContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttribContext)
		}
	}

	return tst
}

func (s *SelectorContext) Attrib(i int) IAttribContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttribContext)
}

func (s *SelectorContext) Pseudo() IPseudoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPseudoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPseudoContext)
}

func (s *SelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterSelector(s)
	}
}

func (s *SelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitSelector(s)
	}
}

func (p *LessParser) Selector() (localctx ISelectorContext) {
	localctx = NewSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LessParserRULE_selector)
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
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(267)
				p.Element()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext())
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LessParserLBRACK {
		{
			p.SetState(272)
			p.Attrib()
		}

		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserCOLON || _la == LessParserCOLONCOLON {
		{
			p.SetState(278)
			p.Pseudo()
		}

	}

	return localctx
}

// IAttribContext is an interface to support dynamic dispatch.
type IAttribContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttribContext differentiates from other interfaces.
	IsAttribContext()
}

type AttribContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribContext() *AttribContext {
	var p = new(AttribContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_attrib
	return p
}

func (*AttribContext) IsAttribContext() {}

func NewAttribContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttribContext {
	var p = new(AttribContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_attrib

	return p
}

func (s *AttribContext) GetParser() antlr.Parser { return s.parser }

func (s *AttribContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(LessParserLBRACK, 0)
}

func (s *AttribContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(LessParserIdentifier)
}

func (s *AttribContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(LessParserIdentifier, i)
}

func (s *AttribContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(LessParserRBRACK, 0)
}

func (s *AttribContext) AttribRelate() IAttribRelateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribRelateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttribRelateContext)
}

func (s *AttribContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(LessParserStringLiteral, 0)
}

func (s *AttribContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttribContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterAttrib(s)
	}
}

func (s *AttribContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitAttrib(s)
	}
}

func (p *LessParser) Attrib() (localctx IAttribContext) {
	localctx = NewAttribContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LessParserRULE_attrib)
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
		p.SetState(281)
		p.Match(LessParserLBRACK)
	}
	{
		p.SetState(282)
		p.Match(LessParserIdentifier)
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(LessParserEQ-33))|(1<<(LessParserPIPE_EQ-33))|(1<<(LessParserTILD_EQ-33)))) != 0 {
		{
			p.SetState(283)
			p.AttribRelate()
		}
		{
			p.SetState(284)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LessParserIdentifier || _la == LessParserStringLiteral) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(288)
		p.Match(LessParserRBRACK)
	}

	return localctx
}

// INegationContext is an interface to support dynamic dispatch.
type INegationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNegationContext differentiates from other interfaces.
	IsNegationContext()
}

type NegationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNegationContext() *NegationContext {
	var p = new(NegationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_negation
	return p
}

func (*NegationContext) IsNegationContext() {}

func NewNegationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NegationContext {
	var p = new(NegationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_negation

	return p
}

func (s *NegationContext) GetParser() antlr.Parser { return s.parser }

func (s *NegationContext) COLON() antlr.TerminalNode {
	return s.GetToken(LessParserCOLON, 0)
}

func (s *NegationContext) NOT() antlr.TerminalNode {
	return s.GetToken(LessParserNOT, 0)
}

func (s *NegationContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserLPAREN, 0)
}

func (s *NegationContext) Selectors() ISelectorsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorsContext)
}

func (s *NegationContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LessParserRPAREN, 0)
}

func (s *NegationContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(LessParserLBRACK, 0)
}

func (s *NegationContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(LessParserRBRACK, 0)
}

func (s *NegationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NegationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterNegation(s)
	}
}

func (s *NegationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitNegation(s)
	}
}

func (p *LessParser) Negation() (localctx INegationContext) {
	localctx = NewNegationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LessParserRULE_negation)
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
		p.SetState(290)
		p.Match(LessParserCOLON)
	}
	{
		p.SetState(291)
		p.Match(LessParserNOT)
	}
	{
		p.SetState(292)
		p.Match(LessParserLPAREN)
	}
	p.SetState(294)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserLBRACK {
		{
			p.SetState(293)
			p.Match(LessParserLBRACK)
		}

	}
	{
		p.SetState(296)
		p.Selectors()
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserRBRACK {
		{
			p.SetState(297)
			p.Match(LessParserRBRACK)
		}

	}
	{
		p.SetState(300)
		p.Match(LessParserRPAREN)
	}

	return localctx
}

// IPseudoContext is an interface to support dynamic dispatch.
type IPseudoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPseudoContext differentiates from other interfaces.
	IsPseudoContext()
}

type PseudoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPseudoContext() *PseudoContext {
	var p = new(PseudoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_pseudo
	return p
}

func (*PseudoContext) IsPseudoContext() {}

func NewPseudoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PseudoContext {
	var p = new(PseudoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_pseudo

	return p
}

func (s *PseudoContext) GetParser() antlr.Parser { return s.parser }

func (s *PseudoContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LessParserIdentifier, 0)
}

func (s *PseudoContext) COLON() antlr.TerminalNode {
	return s.GetToken(LessParserCOLON, 0)
}

func (s *PseudoContext) COLONCOLON() antlr.TerminalNode {
	return s.GetToken(LessParserCOLONCOLON, 0)
}

func (s *PseudoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PseudoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PseudoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterPseudo(s)
	}
}

func (s *PseudoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitPseudo(s)
	}
}

func (p *LessParser) Pseudo() (localctx IPseudoContext) {
	localctx = NewPseudoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LessParserRULE_pseudo)
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
		p.SetState(302)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LessParserCOLON || _la == LessParserCOLONCOLON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(303)
		p.Match(LessParserIdentifier)
	}

	return localctx
}

// IElementContext is an interface to support dynamic dispatch.
type IElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElementContext differentiates from other interfaces.
	IsElementContext()
}

type ElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElementContext() *ElementContext {
	var p = new(ElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_element
	return p
}

func (*ElementContext) IsElementContext() {}

func NewElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElementContext {
	var p = new(ElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_element

	return p
}

func (s *ElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ElementContext) SelectorPrefix() ISelectorPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectorPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectorPrefixContext)
}

func (s *ElementContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ElementContext) HASH() antlr.TerminalNode {
	return s.GetToken(LessParserHASH, 0)
}

func (s *ElementContext) Pseudo() IPseudoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPseudoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPseudoContext)
}

func (s *ElementContext) Negation() INegationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INegationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INegationContext)
}

func (s *ElementContext) PARENTREF() antlr.TerminalNode {
	return s.GetToken(LessParserPARENTREF, 0)
}

func (s *ElementContext) TIMES() antlr.TerminalNode {
	return s.GetToken(LessParserTIMES, 0)
}

func (s *ElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterElement(s)
	}
}

func (s *ElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitElement(s)
	}
}

func (p *LessParser) Element() (localctx IElementContext) {
	localctx = NewElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, LessParserRULE_element)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(305)
			p.SelectorPrefix()
		}
		{
			p.SetState(306)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(308)
			p.Identifier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(309)
			p.Match(LessParserHASH)
		}
		{
			p.SetState(310)
			p.Identifier()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(311)
			p.Pseudo()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(312)
			p.Negation()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(313)
			p.Match(LessParserPARENTREF)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(314)
			p.Match(LessParserTIMES)
		}

	}

	return localctx
}

// ISelectorPrefixContext is an interface to support dynamic dispatch.
type ISelectorPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectorPrefixContext differentiates from other interfaces.
	IsSelectorPrefixContext()
}

type SelectorPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectorPrefixContext() *SelectorPrefixContext {
	var p = new(SelectorPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_selectorPrefix
	return p
}

func (*SelectorPrefixContext) IsSelectorPrefixContext() {}

func NewSelectorPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectorPrefixContext {
	var p = new(SelectorPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_selectorPrefix

	return p
}

func (s *SelectorPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectorPrefixContext) GT() antlr.TerminalNode {
	return s.GetToken(LessParserGT, 0)
}

func (s *SelectorPrefixContext) PLUS() antlr.TerminalNode {
	return s.GetToken(LessParserPLUS, 0)
}

func (s *SelectorPrefixContext) TIL() antlr.TerminalNode {
	return s.GetToken(LessParserTIL, 0)
}

func (s *SelectorPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectorPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectorPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterSelectorPrefix(s)
	}
}

func (s *SelectorPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitSelectorPrefix(s)
	}
}

func (p *LessParser) SelectorPrefix() (localctx ISelectorPrefixContext) {
	localctx = NewSelectorPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, LessParserRULE_selectorPrefix)
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
		p.SetState(317)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LessParserGT)|(1<<LessParserTIL)|(1<<LessParserPLUS))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAttribRelateContext is an interface to support dynamic dispatch.
type IAttribRelateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttribRelateContext differentiates from other interfaces.
	IsAttribRelateContext()
}

type AttribRelateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribRelateContext() *AttribRelateContext {
	var p = new(AttribRelateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_attribRelate
	return p
}

func (*AttribRelateContext) IsAttribRelateContext() {}

func NewAttribRelateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttribRelateContext {
	var p = new(AttribRelateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_attribRelate

	return p
}

func (s *AttribRelateContext) GetParser() antlr.Parser { return s.parser }

func (s *AttribRelateContext) EQ() antlr.TerminalNode {
	return s.GetToken(LessParserEQ, 0)
}

func (s *AttribRelateContext) TILD_EQ() antlr.TerminalNode {
	return s.GetToken(LessParserTILD_EQ, 0)
}

func (s *AttribRelateContext) PIPE_EQ() antlr.TerminalNode {
	return s.GetToken(LessParserPIPE_EQ, 0)
}

func (s *AttribRelateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribRelateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttribRelateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterAttribRelate(s)
	}
}

func (s *AttribRelateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitAttribRelate(s)
	}
}

func (p *LessParser) AttribRelate() (localctx IAttribRelateContext) {
	localctx = NewAttribRelateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, LessParserRULE_attribRelate)
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
		p.SetState(319)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(LessParserEQ-33))|(1<<(LessParserPIPE_EQ-33))|(1<<(LessParserTILD_EQ-33)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = LessParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LessParserIdentifier, 0)
}

func (s *IdentifierContext) AllIdentifierPart() []IIdentifierPartContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierPartContext)(nil)).Elem())
	var tst = make([]IIdentifierPartContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierPartContext)
		}
	}

	return tst
}

func (s *IdentifierContext) IdentifierPart(i int) IIdentifierPartContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierPartContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierPartContext)
}

func (s *IdentifierContext) InterpolationStart() antlr.TerminalNode {
	return s.GetToken(LessParserInterpolationStart, 0)
}

func (s *IdentifierContext) IdentifierVariableName() IIdentifierVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierVariableNameContext)
}

func (s *IdentifierContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(LessParserBlockEnd, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *LessParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, LessParserRULE_identifier)
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

	p.SetState(337)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LessParserIdentifier:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.Match(LessParserIdentifier)
		}
		p.SetState(325)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LessParserInterpolationStartAfter || _la == LessParserIdentifierAfter {
			{
				p.SetState(322)
				p.IdentifierPart()
			}

			p.SetState(327)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case LessParserInterpolationStart:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(328)
			p.Match(LessParserInterpolationStart)
		}
		{
			p.SetState(329)
			p.IdentifierVariableName()
		}
		{
			p.SetState(330)
			p.Match(LessParserBlockEnd)
		}
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LessParserInterpolationStartAfter || _la == LessParserIdentifierAfter {
			{
				p.SetState(331)
				p.IdentifierPart()
			}

			p.SetState(336)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifierPartContext is an interface to support dynamic dispatch.
type IIdentifierPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierPartContext differentiates from other interfaces.
	IsIdentifierPartContext()
}

type IdentifierPartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierPartContext() *IdentifierPartContext {
	var p = new(IdentifierPartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_identifierPart
	return p
}

func (*IdentifierPartContext) IsIdentifierPartContext() {}

func NewIdentifierPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierPartContext {
	var p = new(IdentifierPartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_identifierPart

	return p
}

func (s *IdentifierPartContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierPartContext) InterpolationStartAfter() antlr.TerminalNode {
	return s.GetToken(LessParserInterpolationStartAfter, 0)
}

func (s *IdentifierPartContext) IdentifierVariableName() IIdentifierVariableNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierVariableNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierVariableNameContext)
}

func (s *IdentifierPartContext) BlockEnd() antlr.TerminalNode {
	return s.GetToken(LessParserBlockEnd, 0)
}

func (s *IdentifierPartContext) IdentifierAfter() antlr.TerminalNode {
	return s.GetToken(LessParserIdentifierAfter, 0)
}

func (s *IdentifierPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierPartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterIdentifierPart(s)
	}
}

func (s *IdentifierPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitIdentifierPart(s)
	}
}

func (p *LessParser) IdentifierPart() (localctx IIdentifierPartContext) {
	localctx = NewIdentifierPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, LessParserRULE_identifierPart)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(344)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LessParserInterpolationStartAfter:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(339)
			p.Match(LessParserInterpolationStartAfter)
		}
		{
			p.SetState(340)
			p.IdentifierVariableName()
		}
		{
			p.SetState(341)
			p.Match(LessParserBlockEnd)
		}

	case LessParserIdentifierAfter:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(343)
			p.Match(LessParserIdentifierAfter)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIdentifierVariableNameContext is an interface to support dynamic dispatch.
type IIdentifierVariableNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierVariableNameContext differentiates from other interfaces.
	IsIdentifierVariableNameContext()
}

type IdentifierVariableNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierVariableNameContext() *IdentifierVariableNameContext {
	var p = new(IdentifierVariableNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_identifierVariableName
	return p
}

func (*IdentifierVariableNameContext) IsIdentifierVariableNameContext() {}

func NewIdentifierVariableNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierVariableNameContext {
	var p = new(IdentifierVariableNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_identifierVariableName

	return p
}

func (s *IdentifierVariableNameContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierVariableNameContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LessParserIdentifier, 0)
}

func (s *IdentifierVariableNameContext) IdentifierAfter() antlr.TerminalNode {
	return s.GetToken(LessParserIdentifierAfter, 0)
}

func (s *IdentifierVariableNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierVariableNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierVariableNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterIdentifierVariableName(s)
	}
}

func (s *IdentifierVariableNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitIdentifierVariableName(s)
	}
}

func (p *LessParser) IdentifierVariableName() (localctx IIdentifierVariableNameContext) {
	localctx = NewIdentifierVariableNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, LessParserRULE_identifierVariableName)
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
		p.SetState(346)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LessParserIdentifier || _la == LessParserIdentifierAfter) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PropertyContext) COLON() antlr.TerminalNode {
	return s.GetToken(LessParserCOLON, 0)
}

func (s *PropertyContext) Values() IValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValuesContext)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (p *LessParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, LessParserRULE_property)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Identifier()
	}
	{
		p.SetState(349)
		p.Match(LessParserCOLON)
	}
	{
		p.SetState(350)
		p.Values()
	}

	return localctx
}

// IValuesContext is an interface to support dynamic dispatch.
type IValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValuesContext differentiates from other interfaces.
	IsValuesContext()
}

type ValuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValuesContext() *ValuesContext {
	var p = new(ValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_values
	return p
}

func (*ValuesContext) IsValuesContext() {}

func NewValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValuesContext {
	var p = new(ValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_values

	return p
}

func (s *ValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ValuesContext) AllCommandStatement() []ICommandStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem())
	var tst = make([]ICommandStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICommandStatementContext)
		}
	}

	return tst
}

func (s *ValuesContext) CommandStatement(i int) ICommandStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICommandStatementContext)
}

func (s *ValuesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LessParserCOMMA)
}

func (s *ValuesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LessParserCOMMA, i)
}

func (s *ValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterValues(s)
	}
}

func (s *ValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitValues(s)
	}
}

func (p *LessParser) Values() (localctx IValuesContext) {
	localctx = NewValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, LessParserRULE_values)
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
		p.SetState(352)
		p.CommandStatement()
	}
	p.SetState(357)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LessParserCOMMA {
		{
			p.SetState(353)
			p.Match(LessParserCOMMA)
		}
		{
			p.SetState(354)
			p.CommandStatement()
		}

		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUrlContext is an interface to support dynamic dispatch.
type IUrlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUrlContext differentiates from other interfaces.
	IsUrlContext()
}

type UrlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUrlContext() *UrlContext {
	var p = new(UrlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_url
	return p
}

func (*UrlContext) IsUrlContext() {}

func NewUrlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UrlContext {
	var p = new(UrlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_url

	return p
}

func (s *UrlContext) GetParser() antlr.Parser { return s.parser }

func (s *UrlContext) UrlStart() antlr.TerminalNode {
	return s.GetToken(LessParserUrlStart, 0)
}

func (s *UrlContext) Url() antlr.TerminalNode {
	return s.GetToken(LessParserUrl, 0)
}

func (s *UrlContext) UrlEnd() antlr.TerminalNode {
	return s.GetToken(LessParserUrlEnd, 0)
}

func (s *UrlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UrlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UrlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterUrl(s)
	}
}

func (s *UrlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitUrl(s)
	}
}

func (p *LessParser) Url() (localctx IUrlContext) {
	localctx = NewUrlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, LessParserRULE_url)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(LessParserUrlStart)
	}
	{
		p.SetState(361)
		p.Match(LessParserUrl)
	}
	{
		p.SetState(362)
		p.Match(LessParserUrlEnd)
	}

	return localctx
}

// IMeasurementContext is an interface to support dynamic dispatch.
type IMeasurementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMeasurementContext differentiates from other interfaces.
	IsMeasurementContext()
}

type MeasurementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMeasurementContext() *MeasurementContext {
	var p = new(MeasurementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LessParserRULE_measurement
	return p
}

func (*MeasurementContext) IsMeasurementContext() {}

func NewMeasurementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MeasurementContext {
	var p = new(MeasurementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LessParserRULE_measurement

	return p
}

func (s *MeasurementContext) GetParser() antlr.Parser { return s.parser }

func (s *MeasurementContext) Number() antlr.TerminalNode {
	return s.GetToken(LessParserNumber, 0)
}

func (s *MeasurementContext) Unit() antlr.TerminalNode {
	return s.GetToken(LessParserUnit, 0)
}

func (s *MeasurementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MeasurementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MeasurementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.EnterMeasurement(s)
	}
}

func (s *MeasurementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LessParserListener); ok {
		listenerT.ExitMeasurement(s)
	}
}

func (p *LessParser) Measurement() (localctx IMeasurementContext) {
	localctx = NewMeasurementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, LessParserRULE_measurement)
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
		p.Match(LessParserNumber)
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LessParserUnit {
		{
			p.SetState(365)
			p.Match(LessParserUnit)
		}

	}

	return localctx
}
