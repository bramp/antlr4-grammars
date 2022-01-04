// Code generated from Thrift.g4 by ANTLR 4.9.3. DO NOT EDIT.

package thrift // Thrift
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 55, 408,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 3, 2, 7, 2, 82, 10, 2, 12, 2, 14, 2, 85, 11, 2, 3, 2,
	7, 2, 88, 10, 2, 12, 2, 14, 2, 91, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	5, 3, 98, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 113, 10, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 126, 10, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 5, 8, 133, 10, 8, 3, 8, 5, 8, 136, 10, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 5, 9, 142, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 148, 10, 10,
	12, 10, 14, 10, 151, 11, 10, 3, 10, 3, 10, 5, 10, 155, 10, 10, 3, 11, 3,
	11, 3, 11, 5, 11, 160, 10, 11, 3, 11, 5, 11, 163, 10, 11, 3, 11, 5, 11,
	166, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 173, 10, 12, 7,
	12, 175, 10, 12, 12, 12, 14, 12, 178, 11, 12, 3, 12, 3, 12, 5, 12, 182,
	10, 12, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 188, 10, 13, 12, 13, 14, 13,
	191, 11, 13, 3, 13, 3, 13, 5, 13, 195, 10, 13, 3, 14, 3, 14, 3, 14, 3,
	14, 7, 14, 201, 10, 14, 12, 14, 14, 14, 204, 11, 14, 3, 14, 3, 14, 5, 14,
	208, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 214, 10, 15, 12, 15, 14,
	15, 217, 11, 15, 3, 15, 3, 15, 5, 15, 221, 10, 15, 3, 16, 3, 16, 3, 16,
	3, 16, 5, 16, 227, 10, 16, 3, 16, 3, 16, 7, 16, 231, 10, 16, 12, 16, 14,
	16, 234, 11, 16, 3, 16, 3, 16, 5, 16, 238, 10, 16, 3, 17, 5, 17, 241, 10,
	17, 3, 17, 5, 17, 244, 10, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 250,
	10, 17, 3, 17, 5, 17, 253, 10, 17, 3, 17, 5, 17, 256, 10, 17, 3, 18, 3,
	18, 3, 18, 3, 19, 3, 19, 3, 20, 5, 20, 264, 10, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 7, 20, 270, 10, 20, 12, 20, 14, 20, 273, 11, 20, 3, 20, 3, 20, 5,
	20, 277, 10, 20, 3, 20, 5, 20, 280, 10, 20, 3, 20, 5, 20, 283, 10, 20,
	3, 21, 3, 21, 3, 22, 3, 22, 5, 22, 289, 10, 22, 3, 23, 3, 23, 3, 23, 7,
	23, 294, 10, 23, 12, 23, 14, 23, 297, 11, 23, 3, 23, 3, 23, 3, 24, 3, 24,
	7, 24, 303, 10, 24, 12, 24, 14, 24, 306, 11, 24, 3, 24, 3, 24, 3, 25, 3,
	25, 3, 25, 5, 25, 313, 10, 25, 3, 25, 5, 25, 316, 10, 25, 3, 26, 3, 26,
	5, 26, 320, 10, 26, 3, 27, 3, 27, 3, 27, 5, 27, 325, 10, 27, 3, 28, 3,
	28, 5, 28, 329, 10, 28, 3, 29, 3, 29, 3, 29, 5, 29, 334, 10, 29, 3, 29,
	5, 29, 337, 10, 29, 3, 30, 3, 30, 5, 30, 341, 10, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 5, 31, 351, 10, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 362, 10, 32, 3,
	33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 373,
	10, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 5, 36, 380, 10, 36, 7, 36, 382,
	10, 36, 12, 36, 14, 36, 385, 11, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37,
	3, 37, 5, 37, 393, 10, 37, 3, 38, 3, 38, 7, 38, 397, 10, 38, 12, 38, 14,
	38, 400, 11, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 2, 2,
	41, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
	74, 76, 78, 2, 8, 3, 2, 50, 51, 3, 2, 22, 23, 3, 2, 26, 27, 3, 2, 39, 40,
	4, 2, 38, 38, 52, 52, 3, 2, 42, 49, 2, 436, 2, 83, 3, 2, 2, 2, 4, 97, 3,
	2, 2, 2, 6, 99, 3, 2, 2, 2, 8, 112, 3, 2, 2, 2, 10, 114, 3, 2, 2, 2, 12,
	125, 3, 2, 2, 2, 14, 127, 3, 2, 2, 2, 16, 137, 3, 2, 2, 2, 18, 143, 3,
	2, 2, 2, 20, 156, 3, 2, 2, 2, 22, 167, 3, 2, 2, 2, 24, 183, 3, 2, 2, 2,
	26, 196, 3, 2, 2, 2, 28, 209, 3, 2, 2, 2, 30, 222, 3, 2, 2, 2, 32, 240,
	3, 2, 2, 2, 34, 257, 3, 2, 2, 2, 36, 260, 3, 2, 2, 2, 38, 263, 3, 2, 2,
	2, 40, 284, 3, 2, 2, 2, 42, 288, 3, 2, 2, 2, 44, 290, 3, 2, 2, 2, 46, 300,
	3, 2, 2, 2, 48, 309, 3, 2, 2, 2, 50, 319, 3, 2, 2, 2, 52, 324, 3, 2, 2,
	2, 54, 326, 3, 2, 2, 2, 56, 333, 3, 2, 2, 2, 58, 338, 3, 2, 2, 2, 60, 348,
	3, 2, 2, 2, 62, 356, 3, 2, 2, 2, 64, 363, 3, 2, 2, 2, 66, 372, 3, 2, 2,
	2, 68, 374, 3, 2, 2, 2, 70, 376, 3, 2, 2, 2, 72, 388, 3, 2, 2, 2, 74, 394,
	3, 2, 2, 2, 76, 403, 3, 2, 2, 2, 78, 405, 3, 2, 2, 2, 80, 82, 5, 4, 3,
	2, 81, 80, 3, 2, 2, 2, 82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84,
	3, 2, 2, 2, 84, 89, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 88, 5, 12, 7, 2,
	87, 86, 3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3,
	2, 2, 2, 90, 92, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 7, 2, 2, 3, 93,
	3, 3, 2, 2, 2, 94, 98, 5, 6, 4, 2, 95, 98, 5, 8, 5, 2, 96, 98, 5, 10, 6,
	2, 97, 94, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 96, 3, 2, 2, 2, 98, 5, 3,
	2, 2, 2, 99, 100, 7, 3, 2, 2, 100, 101, 7, 50, 2, 2, 101, 7, 3, 2, 2, 2,
	102, 103, 7, 4, 2, 2, 103, 104, 7, 5, 2, 2, 104, 113, 9, 2, 2, 2, 105,
	106, 7, 4, 2, 2, 106, 107, 7, 51, 2, 2, 107, 113, 9, 2, 2, 2, 108, 109,
	7, 6, 2, 2, 109, 113, 7, 51, 2, 2, 110, 111, 7, 7, 2, 2, 111, 113, 7, 51,
	2, 2, 112, 102, 3, 2, 2, 2, 112, 105, 3, 2, 2, 2, 112, 108, 3, 2, 2, 2,
	112, 110, 3, 2, 2, 2, 113, 9, 3, 2, 2, 2, 114, 115, 7, 8, 2, 2, 115, 116,
	7, 50, 2, 2, 116, 11, 3, 2, 2, 2, 117, 126, 5, 14, 8, 2, 118, 126, 5, 16,
	9, 2, 119, 126, 5, 18, 10, 2, 120, 126, 5, 22, 12, 2, 121, 126, 5, 24,
	13, 2, 122, 126, 5, 26, 14, 2, 123, 126, 5, 28, 15, 2, 124, 126, 5, 30,
	16, 2, 125, 117, 3, 2, 2, 2, 125, 118, 3, 2, 2, 2, 125, 119, 3, 2, 2, 2,
	125, 120, 3, 2, 2, 2, 125, 121, 3, 2, 2, 2, 125, 122, 3, 2, 2, 2, 125,
	123, 3, 2, 2, 2, 125, 124, 3, 2, 2, 2, 126, 13, 3, 2, 2, 2, 127, 128, 7,
	9, 2, 2, 128, 129, 5, 52, 27, 2, 129, 132, 7, 51, 2, 2, 130, 131, 7, 10,
	2, 2, 131, 133, 5, 66, 34, 2, 132, 130, 3, 2, 2, 2, 132, 133, 3, 2, 2,
	2, 133, 135, 3, 2, 2, 2, 134, 136, 5, 76, 39, 2, 135, 134, 3, 2, 2, 2,
	135, 136, 3, 2, 2, 2, 136, 15, 3, 2, 2, 2, 137, 138, 7, 11, 2, 2, 138,
	139, 5, 52, 27, 2, 139, 141, 7, 51, 2, 2, 140, 142, 5, 46, 24, 2, 141,
	140, 3, 2, 2, 2, 141, 142, 3, 2, 2, 2, 142, 17, 3, 2, 2, 2, 143, 144, 7,
	12, 2, 2, 144, 145, 7, 51, 2, 2, 145, 149, 7, 13, 2, 2, 146, 148, 5, 20,
	11, 2, 147, 146, 3, 2, 2, 2, 148, 151, 3, 2, 2, 2, 149, 147, 3, 2, 2, 2,
	149, 150, 3, 2, 2, 2, 150, 152, 3, 2, 2, 2, 151, 149, 3, 2, 2, 2, 152,
	154, 7, 14, 2, 2, 153, 155, 5, 46, 24, 2, 154, 153, 3, 2, 2, 2, 154, 155,
	3, 2, 2, 2, 155, 19, 3, 2, 2, 2, 156, 159, 7, 51, 2, 2, 157, 158, 7, 10,
	2, 2, 158, 160, 5, 68, 35, 2, 159, 157, 3, 2, 2, 2, 159, 160, 3, 2, 2,
	2, 160, 162, 3, 2, 2, 2, 161, 163, 5, 46, 24, 2, 162, 161, 3, 2, 2, 2,
	162, 163, 3, 2, 2, 2, 163, 165, 3, 2, 2, 2, 164, 166, 5, 76, 39, 2, 165,
	164, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 21, 3, 2, 2, 2, 167, 168, 7,
	15, 2, 2, 168, 169, 7, 51, 2, 2, 169, 176, 7, 13, 2, 2, 170, 172, 7, 50,
	2, 2, 171, 173, 5, 76, 39, 2, 172, 171, 3, 2, 2, 2, 172, 173, 3, 2, 2,
	2, 173, 175, 3, 2, 2, 2, 174, 170, 3, 2, 2, 2, 175, 178, 3, 2, 2, 2, 176,
	174, 3, 2, 2, 2, 176, 177, 3, 2, 2, 2, 177, 179, 3, 2, 2, 2, 178, 176,
	3, 2, 2, 2, 179, 181, 7, 14, 2, 2, 180, 182, 5, 46, 24, 2, 181, 180, 3,
	2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 23, 3, 2, 2, 2, 183, 184, 7, 16, 2,
	2, 184, 185, 7, 51, 2, 2, 185, 189, 7, 13, 2, 2, 186, 188, 5, 32, 17, 2,
	187, 186, 3, 2, 2, 2, 188, 191, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 189,
	190, 3, 2, 2, 2, 190, 192, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 192, 194,
	7, 14, 2, 2, 193, 195, 5, 46, 24, 2, 194, 193, 3, 2, 2, 2, 194, 195, 3,
	2, 2, 2, 195, 25, 3, 2, 2, 2, 196, 197, 7, 17, 2, 2, 197, 198, 7, 51, 2,
	2, 198, 202, 7, 13, 2, 2, 199, 201, 5, 32, 17, 2, 200, 199, 3, 2, 2, 2,
	201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203,
	205, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 205, 207, 7, 14, 2, 2, 206, 208,
	5, 46, 24, 2, 207, 206, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 27, 3, 2,
	2, 2, 209, 210, 7, 18, 2, 2, 210, 211, 7, 51, 2, 2, 211, 215, 7, 13, 2,
	2, 212, 214, 5, 32, 17, 2, 213, 212, 3, 2, 2, 2, 214, 217, 3, 2, 2, 2,
	215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 218, 3, 2, 2, 2, 217,
	215, 3, 2, 2, 2, 218, 220, 7, 14, 2, 2, 219, 221, 5, 46, 24, 2, 220, 219,
	3, 2, 2, 2, 220, 221, 3, 2, 2, 2, 221, 29, 3, 2, 2, 2, 222, 223, 7, 19,
	2, 2, 223, 226, 7, 51, 2, 2, 224, 225, 7, 20, 2, 2, 225, 227, 7, 51, 2,
	2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228,
	232, 7, 13, 2, 2, 229, 231, 5, 38, 20, 2, 230, 229, 3, 2, 2, 2, 231, 234,
	3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 235, 3, 2,
	2, 2, 234, 232, 3, 2, 2, 2, 235, 237, 7, 14, 2, 2, 236, 238, 5, 46, 24,
	2, 237, 236, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 31, 3, 2, 2, 2, 239,
	241, 5, 34, 18, 2, 240, 239, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 243,
	3, 2, 2, 2, 242, 244, 5, 36, 19, 2, 243, 242, 3, 2, 2, 2, 243, 244, 3,
	2, 2, 2, 244, 245, 3, 2, 2, 2, 245, 246, 5, 52, 27, 2, 246, 249, 7, 51,
	2, 2, 247, 248, 7, 10, 2, 2, 248, 250, 5, 66, 34, 2, 249, 247, 3, 2, 2,
	2, 249, 250, 3, 2, 2, 2, 250, 252, 3, 2, 2, 2, 251, 253, 5, 46, 24, 2,
	252, 251, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 255, 3, 2, 2, 2, 254,
	256, 5, 76, 39, 2, 255, 254, 3, 2, 2, 2, 255, 256, 3, 2, 2, 2, 256, 33,
	3, 2, 2, 2, 257, 258, 5, 68, 35, 2, 258, 259, 7, 21, 2, 2, 259, 35, 3,
	2, 2, 2, 260, 261, 9, 3, 2, 2, 261, 37, 3, 2, 2, 2, 262, 264, 5, 40, 21,
	2, 263, 262, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265,
	266, 5, 42, 22, 2, 266, 267, 7, 51, 2, 2, 267, 271, 7, 24, 2, 2, 268, 270,
	5, 32, 17, 2, 269, 268, 3, 2, 2, 2, 270, 273, 3, 2, 2, 2, 271, 269, 3,
	2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 274, 3, 2, 2, 2, 273, 271, 3, 2, 2,
	2, 274, 276, 7, 25, 2, 2, 275, 277, 5, 44, 23, 2, 276, 275, 3, 2, 2, 2,
	276, 277, 3, 2, 2, 2, 277, 279, 3, 2, 2, 2, 278, 280, 5, 46, 24, 2, 279,
	278, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 282, 3, 2, 2, 2, 281, 283,
	5, 76, 39, 2, 282, 281, 3, 2, 2, 2, 282, 283, 3, 2, 2, 2, 283, 39, 3, 2,
	2, 2, 284, 285, 9, 4, 2, 2, 285, 41, 3, 2, 2, 2, 286, 289, 5, 52, 27, 2,
	287, 289, 7, 28, 2, 2, 288, 286, 3, 2, 2, 2, 288, 287, 3, 2, 2, 2, 289,
	43, 3, 2, 2, 2, 290, 291, 7, 29, 2, 2, 291, 295, 7, 24, 2, 2, 292, 294,
	5, 32, 17, 2, 293, 292, 3, 2, 2, 2, 294, 297, 3, 2, 2, 2, 295, 293, 3,
	2, 2, 2, 295, 296, 3, 2, 2, 2, 296, 298, 3, 2, 2, 2, 297, 295, 3, 2, 2,
	2, 298, 299, 7, 25, 2, 2, 299, 45, 3, 2, 2, 2, 300, 304, 7, 24, 2, 2, 301,
	303, 5, 48, 25, 2, 302, 301, 3, 2, 2, 2, 303, 306, 3, 2, 2, 2, 304, 302,
	3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 307, 3, 2, 2, 2, 306, 304, 3, 2,
	2, 2, 307, 308, 7, 25, 2, 2, 308, 47, 3, 2, 2, 2, 309, 312, 7, 51, 2, 2,
	310, 311, 7, 10, 2, 2, 311, 313, 5, 50, 26, 2, 312, 310, 3, 2, 2, 2, 312,
	313, 3, 2, 2, 2, 313, 315, 3, 2, 2, 2, 314, 316, 5, 76, 39, 2, 315, 314,
	3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 49, 3, 2, 2, 2, 317, 320, 5, 68,
	35, 2, 318, 320, 7, 50, 2, 2, 319, 317, 3, 2, 2, 2, 319, 318, 3, 2, 2,
	2, 320, 51, 3, 2, 2, 2, 321, 325, 5, 54, 28, 2, 322, 325, 7, 51, 2, 2,
	323, 325, 5, 56, 29, 2, 324, 321, 3, 2, 2, 2, 324, 322, 3, 2, 2, 2, 324,
	323, 3, 2, 2, 2, 325, 53, 3, 2, 2, 2, 326, 328, 5, 78, 40, 2, 327, 329,
	5, 46, 24, 2, 328, 327, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 55, 3, 2,
	2, 2, 330, 334, 5, 58, 30, 2, 331, 334, 5, 60, 31, 2, 332, 334, 5, 62,
	32, 2, 333, 330, 3, 2, 2, 2, 333, 331, 3, 2, 2, 2, 333, 332, 3, 2, 2, 2,
	334, 336, 3, 2, 2, 2, 335, 337, 5, 46, 24, 2, 336, 335, 3, 2, 2, 2, 336,
	337, 3, 2, 2, 2, 337, 57, 3, 2, 2, 2, 338, 340, 7, 30, 2, 2, 339, 341,
	5, 64, 33, 2, 340, 339, 3, 2, 2, 2, 340, 341, 3, 2, 2, 2, 341, 342, 3,
	2, 2, 2, 342, 343, 7, 31, 2, 2, 343, 344, 5, 52, 27, 2, 344, 345, 7, 52,
	2, 2, 345, 346, 5, 52, 27, 2, 346, 347, 7, 32, 2, 2, 347, 59, 3, 2, 2,
	2, 348, 350, 7, 33, 2, 2, 349, 351, 5, 64, 33, 2, 350, 349, 3, 2, 2, 2,
	350, 351, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2, 352, 353, 7, 31, 2, 2, 353,
	354, 5, 52, 27, 2, 354, 355, 7, 32, 2, 2, 355, 61, 3, 2, 2, 2, 356, 357,
	7, 34, 2, 2, 357, 358, 7, 31, 2, 2, 358, 359, 5, 52, 27, 2, 359, 361, 7,
	32, 2, 2, 360, 362, 5, 64, 33, 2, 361, 360, 3, 2, 2, 2, 361, 362, 3, 2,
	2, 2, 362, 63, 3, 2, 2, 2, 363, 364, 7, 35, 2, 2, 364, 365, 7, 50, 2, 2,
	365, 65, 3, 2, 2, 2, 366, 373, 5, 68, 35, 2, 367, 373, 7, 41, 2, 2, 368,
	373, 7, 50, 2, 2, 369, 373, 7, 51, 2, 2, 370, 373, 5, 70, 36, 2, 371, 373,
	5, 74, 38, 2, 372, 366, 3, 2, 2, 2, 372, 367, 3, 2, 2, 2, 372, 368, 3,
	2, 2, 2, 372, 369, 3, 2, 2, 2, 372, 370, 3, 2, 2, 2, 372, 371, 3, 2, 2,
	2, 373, 67, 3, 2, 2, 2, 374, 375, 9, 5, 2, 2, 375, 69, 3, 2, 2, 2, 376,
	383, 7, 36, 2, 2, 377, 379, 5, 66, 34, 2, 378, 380, 5, 76, 39, 2, 379,
	378, 3, 2, 2, 2, 379, 380, 3, 2, 2, 2, 380, 382, 3, 2, 2, 2, 381, 377,
	3, 2, 2, 2, 382, 385, 3, 2, 2, 2, 383, 381, 3, 2, 2, 2, 383, 384, 3, 2,
	2, 2, 384, 386, 3, 2, 2, 2, 385, 383, 3, 2, 2, 2, 386, 387, 7, 37, 2, 2,
	387, 71, 3, 2, 2, 2, 388, 389, 5, 66, 34, 2, 389, 390, 7, 21, 2, 2, 390,
	392, 5, 66, 34, 2, 391, 393, 5, 76, 39, 2, 392, 391, 3, 2, 2, 2, 392, 393,
	3, 2, 2, 2, 393, 73, 3, 2, 2, 2, 394, 398, 7, 13, 2, 2, 395, 397, 5, 72,
	37, 2, 396, 395, 3, 2, 2, 2, 397, 400, 3, 2, 2, 2, 398, 396, 3, 2, 2, 2,
	398, 399, 3, 2, 2, 2, 399, 401, 3, 2, 2, 2, 400, 398, 3, 2, 2, 2, 401,
	402, 7, 14, 2, 2, 402, 75, 3, 2, 2, 2, 403, 404, 9, 6, 2, 2, 404, 77, 3,
	2, 2, 2, 405, 406, 9, 7, 2, 2, 406, 79, 3, 2, 2, 2, 55, 83, 89, 97, 112,
	125, 132, 135, 141, 149, 154, 159, 162, 165, 172, 176, 181, 189, 194, 202,
	207, 215, 220, 226, 232, 237, 240, 243, 249, 252, 255, 263, 271, 276, 279,
	282, 288, 295, 304, 312, 315, 319, 324, 328, 333, 336, 340, 350, 361, 372,
	379, 383, 392, 398,
}
var literalNames = []string{
	"", "'include'", "'namespace'", "'*'", "'cpp_namespace'", "'php_namespace'",
	"'cpp_include'", "'const'", "'='", "'typedef'", "'enum'", "'{'", "'}'",
	"'senum'", "'struct'", "'union'", "'exception'", "'service'", "'extends'",
	"':'", "'required'", "'optional'", "'('", "')'", "'oneway'", "'async'",
	"'void'", "'throws'", "'map'", "'<'", "'>'", "'set'", "'list'", "'cpp_type'",
	"'['", "']'", "';'", "", "", "", "'bool'", "'byte'", "'i16'", "'i32'",
	"'i64'", "'double'", "'string'", "'binary'", "", "", "','",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "INTEGER", "HEX_INTEGER", "DOUBLE", "TYPE_BOOL", "TYPE_BYTE", "TYPE_I16",
	"TYPE_I32", "TYPE_I64", "TYPE_DOUBLE", "TYPE_STRING", "TYPE_BINARY", "LITERAL",
	"IDENTIFIER", "COMMA", "WS", "SL_COMMENT", "ML_COMMENT",
}

var ruleNames = []string{
	"document", "header", "include_", "namespace_", "cpp_include", "definition",
	"const_rule", "typedef_", "enum_rule", "enum_field", "senum", "struct_",
	"union_", "exception", "service", "field", "field_id", "field_req", "function_",
	"oneway", "function_type", "throws_list", "type_annotations", "type_annotation",
	"annotation_value", "field_type", "base_type", "container_type", "map_type",
	"set_type", "list_type", "cpp_type", "const_value", "integer", "const_list",
	"const_map_entry", "const_map", "list_separator", "real_base_type",
}

type ThriftParser struct {
	*antlr.BaseParser
}

// NewThriftParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ThriftParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewThriftParser(input antlr.TokenStream) *ThriftParser {
	this := new(ThriftParser)
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
	this.GrammarFileName = "Thrift.g4"

	return this
}

// ThriftParser tokens.
const (
	ThriftParserEOF         = antlr.TokenEOF
	ThriftParserT__0        = 1
	ThriftParserT__1        = 2
	ThriftParserT__2        = 3
	ThriftParserT__3        = 4
	ThriftParserT__4        = 5
	ThriftParserT__5        = 6
	ThriftParserT__6        = 7
	ThriftParserT__7        = 8
	ThriftParserT__8        = 9
	ThriftParserT__9        = 10
	ThriftParserT__10       = 11
	ThriftParserT__11       = 12
	ThriftParserT__12       = 13
	ThriftParserT__13       = 14
	ThriftParserT__14       = 15
	ThriftParserT__15       = 16
	ThriftParserT__16       = 17
	ThriftParserT__17       = 18
	ThriftParserT__18       = 19
	ThriftParserT__19       = 20
	ThriftParserT__20       = 21
	ThriftParserT__21       = 22
	ThriftParserT__22       = 23
	ThriftParserT__23       = 24
	ThriftParserT__24       = 25
	ThriftParserT__25       = 26
	ThriftParserT__26       = 27
	ThriftParserT__27       = 28
	ThriftParserT__28       = 29
	ThriftParserT__29       = 30
	ThriftParserT__30       = 31
	ThriftParserT__31       = 32
	ThriftParserT__32       = 33
	ThriftParserT__33       = 34
	ThriftParserT__34       = 35
	ThriftParserT__35       = 36
	ThriftParserINTEGER     = 37
	ThriftParserHEX_INTEGER = 38
	ThriftParserDOUBLE      = 39
	ThriftParserTYPE_BOOL   = 40
	ThriftParserTYPE_BYTE   = 41
	ThriftParserTYPE_I16    = 42
	ThriftParserTYPE_I32    = 43
	ThriftParserTYPE_I64    = 44
	ThriftParserTYPE_DOUBLE = 45
	ThriftParserTYPE_STRING = 46
	ThriftParserTYPE_BINARY = 47
	ThriftParserLITERAL     = 48
	ThriftParserIDENTIFIER  = 49
	ThriftParserCOMMA       = 50
	ThriftParserWS          = 51
	ThriftParserSL_COMMENT  = 52
	ThriftParserML_COMMENT  = 53
)

// ThriftParser rules.
const (
	ThriftParserRULE_document         = 0
	ThriftParserRULE_header           = 1
	ThriftParserRULE_include_         = 2
	ThriftParserRULE_namespace_       = 3
	ThriftParserRULE_cpp_include      = 4
	ThriftParserRULE_definition       = 5
	ThriftParserRULE_const_rule       = 6
	ThriftParserRULE_typedef_         = 7
	ThriftParserRULE_enum_rule        = 8
	ThriftParserRULE_enum_field       = 9
	ThriftParserRULE_senum            = 10
	ThriftParserRULE_struct_          = 11
	ThriftParserRULE_union_           = 12
	ThriftParserRULE_exception        = 13
	ThriftParserRULE_service          = 14
	ThriftParserRULE_field            = 15
	ThriftParserRULE_field_id         = 16
	ThriftParserRULE_field_req        = 17
	ThriftParserRULE_function_        = 18
	ThriftParserRULE_oneway           = 19
	ThriftParserRULE_function_type    = 20
	ThriftParserRULE_throws_list      = 21
	ThriftParserRULE_type_annotations = 22
	ThriftParserRULE_type_annotation  = 23
	ThriftParserRULE_annotation_value = 24
	ThriftParserRULE_field_type       = 25
	ThriftParserRULE_base_type        = 26
	ThriftParserRULE_container_type   = 27
	ThriftParserRULE_map_type         = 28
	ThriftParserRULE_set_type         = 29
	ThriftParserRULE_list_type        = 30
	ThriftParserRULE_cpp_type         = 31
	ThriftParserRULE_const_value      = 32
	ThriftParserRULE_integer          = 33
	ThriftParserRULE_const_list       = 34
	ThriftParserRULE_const_map_entry  = 35
	ThriftParserRULE_const_map        = 36
	ThriftParserRULE_list_separator   = 37
	ThriftParserRULE_real_base_type   = 38
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
	p.RuleIndex = ThriftParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(ThriftParserEOF, 0)
}

func (s *DocumentContext) AllHeader() []IHeaderContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHeaderContext)(nil)).Elem())
	var tst = make([]IHeaderContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHeaderContext)
		}
	}

	return tst
}

func (s *DocumentContext) Header(i int) IHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeaderContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHeaderContext)
}

func (s *DocumentContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *DocumentContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (p *ThriftParser) Document() (localctx IDocumentContext) {
	this := p
	_ = this

	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ThriftParserRULE_document)
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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThriftParserT__0)|(1<<ThriftParserT__1)|(1<<ThriftParserT__3)|(1<<ThriftParserT__4)|(1<<ThriftParserT__5))) != 0 {
		{
			p.SetState(78)
			p.Header()
		}

		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ThriftParserT__6)|(1<<ThriftParserT__8)|(1<<ThriftParserT__9)|(1<<ThriftParserT__12)|(1<<ThriftParserT__13)|(1<<ThriftParserT__14)|(1<<ThriftParserT__15)|(1<<ThriftParserT__16))) != 0 {
		{
			p.SetState(84)
			p.Definition()
		}

		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(90)
		p.Match(ThriftParserEOF)
	}

	return localctx
}

// IHeaderContext is an interface to support dynamic dispatch.
type IHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeaderContext differentiates from other interfaces.
	IsHeaderContext()
}

type HeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeaderContext() *HeaderContext {
	var p = new(HeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_header
	return p
}

func (*HeaderContext) IsHeaderContext() {}

func NewHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HeaderContext {
	var p = new(HeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_header

	return p
}

func (s *HeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *HeaderContext) Include_() IInclude_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInclude_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInclude_Context)
}

func (s *HeaderContext) Namespace_() INamespace_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamespace_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamespace_Context)
}

func (s *HeaderContext) Cpp_include() ICpp_includeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICpp_includeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICpp_includeContext)
}

func (s *HeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterHeader(s)
	}
}

func (s *HeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitHeader(s)
	}
}

func (p *ThriftParser) Header() (localctx IHeaderContext) {
	this := p
	_ = this

	localctx = NewHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ThriftParserRULE_header)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(95)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(92)
			p.Include_()
		}

	case ThriftParserT__1, ThriftParserT__3, ThriftParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Namespace_()
		}

	case ThriftParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Cpp_include()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInclude_Context is an interface to support dynamic dispatch.
type IInclude_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInclude_Context differentiates from other interfaces.
	IsInclude_Context()
}

type Include_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInclude_Context() *Include_Context {
	var p = new(Include_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_include_
	return p
}

func (*Include_Context) IsInclude_Context() {}

func NewInclude_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Include_Context {
	var p = new(Include_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_include_

	return p
}

func (s *Include_Context) GetParser() antlr.Parser { return s.parser }

func (s *Include_Context) LITERAL() antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, 0)
}

func (s *Include_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Include_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Include_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterInclude_(s)
	}
}

func (s *Include_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitInclude_(s)
	}
}

func (p *ThriftParser) Include_() (localctx IInclude_Context) {
	this := p
	_ = this

	localctx = NewInclude_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ThriftParserRULE_include_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(ThriftParserT__0)
	}
	{
		p.SetState(98)
		p.Match(ThriftParserLITERAL)
	}

	return localctx
}

// INamespace_Context is an interface to support dynamic dispatch.
type INamespace_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamespace_Context differentiates from other interfaces.
	IsNamespace_Context()
}

type Namespace_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamespace_Context() *Namespace_Context {
	var p = new(Namespace_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_namespace_
	return p
}

func (*Namespace_Context) IsNamespace_Context() {}

func NewNamespace_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Namespace_Context {
	var p = new(Namespace_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_namespace_

	return p
}

func (s *Namespace_Context) GetParser() antlr.Parser { return s.parser }

func (s *Namespace_Context) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ThriftParserIDENTIFIER)
}

func (s *Namespace_Context) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, i)
}

func (s *Namespace_Context) LITERAL() antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, 0)
}

func (s *Namespace_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Namespace_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Namespace_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterNamespace_(s)
	}
}

func (s *Namespace_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitNamespace_(s)
	}
}

func (p *ThriftParser) Namespace_() (localctx INamespace_Context) {
	this := p
	_ = this

	localctx = NewNamespace_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ThriftParserRULE_namespace_)
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

	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Match(ThriftParserT__1)
		}
		{
			p.SetState(101)
			p.Match(ThriftParserT__2)
		}
		{
			p.SetState(102)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ThriftParserLITERAL || _la == ThriftParserIDENTIFIER) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(103)
			p.Match(ThriftParserT__1)
		}
		{
			p.SetState(104)
			p.Match(ThriftParserIDENTIFIER)
		}
		{
			p.SetState(105)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ThriftParserLITERAL || _la == ThriftParserIDENTIFIER) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(106)
			p.Match(ThriftParserT__3)
		}
		{
			p.SetState(107)
			p.Match(ThriftParserIDENTIFIER)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(108)
			p.Match(ThriftParserT__4)
		}
		{
			p.SetState(109)
			p.Match(ThriftParserIDENTIFIER)
		}

	}

	return localctx
}

// ICpp_includeContext is an interface to support dynamic dispatch.
type ICpp_includeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCpp_includeContext differentiates from other interfaces.
	IsCpp_includeContext()
}

type Cpp_includeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCpp_includeContext() *Cpp_includeContext {
	var p = new(Cpp_includeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_cpp_include
	return p
}

func (*Cpp_includeContext) IsCpp_includeContext() {}

func NewCpp_includeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cpp_includeContext {
	var p = new(Cpp_includeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_cpp_include

	return p
}

func (s *Cpp_includeContext) GetParser() antlr.Parser { return s.parser }

func (s *Cpp_includeContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, 0)
}

func (s *Cpp_includeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cpp_includeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cpp_includeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterCpp_include(s)
	}
}

func (s *Cpp_includeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitCpp_include(s)
	}
}

func (p *ThriftParser) Cpp_include() (localctx ICpp_includeContext) {
	this := p
	_ = this

	localctx = NewCpp_includeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ThriftParserRULE_cpp_include)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(ThriftParserT__5)
	}
	{
		p.SetState(113)
		p.Match(ThriftParserLITERAL)
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
	p.RuleIndex = ThriftParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) Const_rule() IConst_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_ruleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConst_ruleContext)
}

func (s *DefinitionContext) Typedef_() ITypedef_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypedef_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypedef_Context)
}

func (s *DefinitionContext) Enum_rule() IEnum_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnum_ruleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnum_ruleContext)
}

func (s *DefinitionContext) Senum() ISenumContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISenumContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISenumContext)
}

func (s *DefinitionContext) Struct_() IStruct_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStruct_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStruct_Context)
}

func (s *DefinitionContext) Union_() IUnion_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnion_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnion_Context)
}

func (s *DefinitionContext) Exception() IExceptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExceptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExceptionContext)
}

func (s *DefinitionContext) Service() IServiceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IServiceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IServiceContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *ThriftParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ThriftParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.Const_rule()
		}

	case ThriftParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
			p.Typedef_()
		}

	case ThriftParserT__9:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(117)
			p.Enum_rule()
		}

	case ThriftParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(118)
			p.Senum()
		}

	case ThriftParserT__13:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(119)
			p.Struct_()
		}

	case ThriftParserT__14:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(120)
			p.Union_()
		}

	case ThriftParserT__15:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(121)
			p.Exception()
		}

	case ThriftParserT__16:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(122)
			p.Service()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConst_ruleContext is an interface to support dynamic dispatch.
type IConst_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConst_ruleContext differentiates from other interfaces.
	IsConst_ruleContext()
}

type Const_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_ruleContext() *Const_ruleContext {
	var p = new(Const_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_const_rule
	return p
}

func (*Const_ruleContext) IsConst_ruleContext() {}

func NewConst_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_ruleContext {
	var p = new(Const_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_const_rule

	return p
}

func (s *Const_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_ruleContext) Field_type() IField_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Const_ruleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Const_ruleContext) Const_value() IConst_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *Const_ruleContext) List_separator() IList_separatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_separatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Const_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterConst_rule(s)
	}
}

func (s *Const_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitConst_rule(s)
	}
}

func (p *ThriftParser) Const_rule() (localctx IConst_ruleContext) {
	this := p
	_ = this

	localctx = NewConst_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ThriftParserRULE_const_rule)
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
		p.SetState(125)
		p.Match(ThriftParserT__6)
	}
	{
		p.SetState(126)
		p.Field_type()
	}
	{
		p.SetState(127)
		p.Match(ThriftParserIDENTIFIER)
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__7 {
		{
			p.SetState(128)
			p.Match(ThriftParserT__7)
		}
		{
			p.SetState(129)
			p.Const_value()
		}

	}
	p.SetState(133)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
		{
			p.SetState(132)
			p.List_separator()
		}

	}

	return localctx
}

// ITypedef_Context is an interface to support dynamic dispatch.
type ITypedef_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypedef_Context differentiates from other interfaces.
	IsTypedef_Context()
}

type Typedef_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypedef_Context() *Typedef_Context {
	var p = new(Typedef_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_typedef_
	return p
}

func (*Typedef_Context) IsTypedef_Context() {}

func NewTypedef_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Typedef_Context {
	var p = new(Typedef_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_typedef_

	return p
}

func (s *Typedef_Context) GetParser() antlr.Parser { return s.parser }

func (s *Typedef_Context) Field_type() IField_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Typedef_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Typedef_Context) Type_annotations() IType_annotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Typedef_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Typedef_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Typedef_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterTypedef_(s)
	}
}

func (s *Typedef_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitTypedef_(s)
	}
}

func (p *ThriftParser) Typedef_() (localctx ITypedef_Context) {
	this := p
	_ = this

	localctx = NewTypedef_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ThriftParserRULE_typedef_)
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
		p.Match(ThriftParserT__8)
	}
	{
		p.SetState(136)
		p.Field_type()
	}
	{
		p.SetState(137)
		p.Match(ThriftParserIDENTIFIER)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(138)
			p.Type_annotations()
		}

	}

	return localctx
}

// IEnum_ruleContext is an interface to support dynamic dispatch.
type IEnum_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnum_ruleContext differentiates from other interfaces.
	IsEnum_ruleContext()
}

type Enum_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnum_ruleContext() *Enum_ruleContext {
	var p = new(Enum_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_enum_rule
	return p
}

func (*Enum_ruleContext) IsEnum_ruleContext() {}

func NewEnum_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_ruleContext {
	var p = new(Enum_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_enum_rule

	return p
}

func (s *Enum_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Enum_ruleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Enum_ruleContext) AllEnum_field() []IEnum_fieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnum_fieldContext)(nil)).Elem())
	var tst = make([]IEnum_fieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnum_fieldContext)
		}
	}

	return tst
}

func (s *Enum_ruleContext) Enum_field(i int) IEnum_fieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnum_fieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnum_fieldContext)
}

func (s *Enum_ruleContext) Type_annotations() IType_annotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Enum_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterEnum_rule(s)
	}
}

func (s *Enum_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitEnum_rule(s)
	}
}

func (p *ThriftParser) Enum_rule() (localctx IEnum_ruleContext) {
	this := p
	_ = this

	localctx = NewEnum_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ThriftParserRULE_enum_rule)
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
		p.SetState(141)
		p.Match(ThriftParserT__9)
	}
	{
		p.SetState(142)
		p.Match(ThriftParserIDENTIFIER)
	}
	{
		p.SetState(143)
		p.Match(ThriftParserT__10)
	}
	p.SetState(147)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ThriftParserIDENTIFIER {
		{
			p.SetState(144)
			p.Enum_field()
		}

		p.SetState(149)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(150)
		p.Match(ThriftParserT__11)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(151)
			p.Type_annotations()
		}

	}

	return localctx
}

// IEnum_fieldContext is an interface to support dynamic dispatch.
type IEnum_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnum_fieldContext differentiates from other interfaces.
	IsEnum_fieldContext()
}

type Enum_fieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnum_fieldContext() *Enum_fieldContext {
	var p = new(Enum_fieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_enum_field
	return p
}

func (*Enum_fieldContext) IsEnum_fieldContext() {}

func NewEnum_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Enum_fieldContext {
	var p = new(Enum_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_enum_field

	return p
}

func (s *Enum_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Enum_fieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Enum_fieldContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *Enum_fieldContext) Type_annotations() IType_annotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Enum_fieldContext) List_separator() IList_separatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_separatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Enum_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Enum_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Enum_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterEnum_field(s)
	}
}

func (s *Enum_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitEnum_field(s)
	}
}

func (p *ThriftParser) Enum_field() (localctx IEnum_fieldContext) {
	this := p
	_ = this

	localctx = NewEnum_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ThriftParserRULE_enum_field)
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
		p.Match(ThriftParserIDENTIFIER)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__7 {
		{
			p.SetState(155)
			p.Match(ThriftParserT__7)
		}
		{
			p.SetState(156)
			p.Integer()
		}

	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(159)
			p.Type_annotations()
		}

	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
		{
			p.SetState(162)
			p.List_separator()
		}

	}

	return localctx
}

// ISenumContext is an interface to support dynamic dispatch.
type ISenumContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSenumContext differentiates from other interfaces.
	IsSenumContext()
}

type SenumContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySenumContext() *SenumContext {
	var p = new(SenumContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_senum
	return p
}

func (*SenumContext) IsSenumContext() {}

func NewSenumContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SenumContext {
	var p = new(SenumContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_senum

	return p
}

func (s *SenumContext) GetParser() antlr.Parser { return s.parser }

func (s *SenumContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *SenumContext) AllLITERAL() []antlr.TerminalNode {
	return s.GetTokens(ThriftParserLITERAL)
}

func (s *SenumContext) LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, i)
}

func (s *SenumContext) Type_annotations() IType_annotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *SenumContext) AllList_separator() []IList_separatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IList_separatorContext)(nil)).Elem())
	var tst = make([]IList_separatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IList_separatorContext)
		}
	}

	return tst
}

func (s *SenumContext) List_separator(i int) IList_separatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_separatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *SenumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SenumContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SenumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterSenum(s)
	}
}

func (s *SenumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitSenum(s)
	}
}

func (p *ThriftParser) Senum() (localctx ISenumContext) {
	this := p
	_ = this

	localctx = NewSenumContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ThriftParserRULE_senum)
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
		p.Match(ThriftParserT__12)
	}
	{
		p.SetState(166)
		p.Match(ThriftParserIDENTIFIER)
	}
	{
		p.SetState(167)
		p.Match(ThriftParserT__10)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ThriftParserLITERAL {
		{
			p.SetState(168)
			p.Match(ThriftParserLITERAL)
		}
		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
			{
				p.SetState(169)
				p.List_separator()
			}

		}

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(177)
		p.Match(ThriftParserT__11)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(178)
			p.Type_annotations()
		}

	}

	return localctx
}

// IStruct_Context is an interface to support dynamic dispatch.
type IStruct_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStruct_Context differentiates from other interfaces.
	IsStruct_Context()
}

type Struct_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStruct_Context() *Struct_Context {
	var p = new(Struct_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_struct_
	return p
}

func (*Struct_Context) IsStruct_Context() {}

func NewStruct_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Struct_Context {
	var p = new(Struct_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_struct_

	return p
}

func (s *Struct_Context) GetParser() antlr.Parser { return s.parser }

func (s *Struct_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Struct_Context) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *Struct_Context) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Struct_Context) Type_annotations() IType_annotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Struct_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Struct_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Struct_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterStruct_(s)
	}
}

func (s *Struct_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitStruct_(s)
	}
}

func (p *ThriftParser) Struct_() (localctx IStruct_Context) {
	this := p
	_ = this

	localctx = NewStruct_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ThriftParserRULE_struct_)
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
		p.SetState(181)
		p.Match(ThriftParserT__13)
	}
	{
		p.SetState(182)
		p.Match(ThriftParserIDENTIFIER)
	}
	{
		p.SetState(183)
		p.Match(ThriftParserT__10)
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(ThriftParserT__19-20))|(1<<(ThriftParserT__20-20))|(1<<(ThriftParserT__27-20))|(1<<(ThriftParserT__30-20))|(1<<(ThriftParserT__31-20))|(1<<(ThriftParserINTEGER-20))|(1<<(ThriftParserHEX_INTEGER-20))|(1<<(ThriftParserTYPE_BOOL-20))|(1<<(ThriftParserTYPE_BYTE-20))|(1<<(ThriftParserTYPE_I16-20))|(1<<(ThriftParserTYPE_I32-20))|(1<<(ThriftParserTYPE_I64-20))|(1<<(ThriftParserTYPE_DOUBLE-20))|(1<<(ThriftParserTYPE_STRING-20))|(1<<(ThriftParserTYPE_BINARY-20))|(1<<(ThriftParserIDENTIFIER-20)))) != 0 {
		{
			p.SetState(184)
			p.Field()
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(190)
		p.Match(ThriftParserT__11)
	}
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(191)
			p.Type_annotations()
		}

	}

	return localctx
}

// IUnion_Context is an interface to support dynamic dispatch.
type IUnion_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnion_Context differentiates from other interfaces.
	IsUnion_Context()
}

type Union_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnion_Context() *Union_Context {
	var p = new(Union_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_union_
	return p
}

func (*Union_Context) IsUnion_Context() {}

func NewUnion_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Union_Context {
	var p = new(Union_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_union_

	return p
}

func (s *Union_Context) GetParser() antlr.Parser { return s.parser }

func (s *Union_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Union_Context) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *Union_Context) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Union_Context) Type_annotations() IType_annotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Union_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Union_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Union_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterUnion_(s)
	}
}

func (s *Union_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitUnion_(s)
	}
}

func (p *ThriftParser) Union_() (localctx IUnion_Context) {
	this := p
	_ = this

	localctx = NewUnion_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ThriftParserRULE_union_)
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
		p.SetState(194)
		p.Match(ThriftParserT__14)
	}
	{
		p.SetState(195)
		p.Match(ThriftParserIDENTIFIER)
	}
	{
		p.SetState(196)
		p.Match(ThriftParserT__10)
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(ThriftParserT__19-20))|(1<<(ThriftParserT__20-20))|(1<<(ThriftParserT__27-20))|(1<<(ThriftParserT__30-20))|(1<<(ThriftParserT__31-20))|(1<<(ThriftParserINTEGER-20))|(1<<(ThriftParserHEX_INTEGER-20))|(1<<(ThriftParserTYPE_BOOL-20))|(1<<(ThriftParserTYPE_BYTE-20))|(1<<(ThriftParserTYPE_I16-20))|(1<<(ThriftParserTYPE_I32-20))|(1<<(ThriftParserTYPE_I64-20))|(1<<(ThriftParserTYPE_DOUBLE-20))|(1<<(ThriftParserTYPE_STRING-20))|(1<<(ThriftParserTYPE_BINARY-20))|(1<<(ThriftParserIDENTIFIER-20)))) != 0 {
		{
			p.SetState(197)
			p.Field()
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(203)
		p.Match(ThriftParserT__11)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(204)
			p.Type_annotations()
		}

	}

	return localctx
}

// IExceptionContext is an interface to support dynamic dispatch.
type IExceptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExceptionContext differentiates from other interfaces.
	IsExceptionContext()
}

type ExceptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExceptionContext() *ExceptionContext {
	var p = new(ExceptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_exception
	return p
}

func (*ExceptionContext) IsExceptionContext() {}

func NewExceptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExceptionContext {
	var p = new(ExceptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_exception

	return p
}

func (s *ExceptionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExceptionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *ExceptionContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *ExceptionContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ExceptionContext) Type_annotations() IType_annotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *ExceptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExceptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExceptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterException(s)
	}
}

func (s *ExceptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitException(s)
	}
}

func (p *ThriftParser) Exception() (localctx IExceptionContext) {
	this := p
	_ = this

	localctx = NewExceptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ThriftParserRULE_exception)
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
		p.Match(ThriftParserT__15)
	}
	{
		p.SetState(208)
		p.Match(ThriftParserIDENTIFIER)
	}
	{
		p.SetState(209)
		p.Match(ThriftParserT__10)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(ThriftParserT__19-20))|(1<<(ThriftParserT__20-20))|(1<<(ThriftParserT__27-20))|(1<<(ThriftParserT__30-20))|(1<<(ThriftParserT__31-20))|(1<<(ThriftParserINTEGER-20))|(1<<(ThriftParserHEX_INTEGER-20))|(1<<(ThriftParserTYPE_BOOL-20))|(1<<(ThriftParserTYPE_BYTE-20))|(1<<(ThriftParserTYPE_I16-20))|(1<<(ThriftParserTYPE_I32-20))|(1<<(ThriftParserTYPE_I64-20))|(1<<(ThriftParserTYPE_DOUBLE-20))|(1<<(ThriftParserTYPE_STRING-20))|(1<<(ThriftParserTYPE_BINARY-20))|(1<<(ThriftParserIDENTIFIER-20)))) != 0 {
		{
			p.SetState(210)
			p.Field()
		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(216)
		p.Match(ThriftParserT__11)
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(217)
			p.Type_annotations()
		}

	}

	return localctx
}

// IServiceContext is an interface to support dynamic dispatch.
type IServiceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsServiceContext differentiates from other interfaces.
	IsServiceContext()
}

type ServiceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceContext() *ServiceContext {
	var p = new(ServiceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_service
	return p
}

func (*ServiceContext) IsServiceContext() {}

func NewServiceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceContext {
	var p = new(ServiceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_service

	return p
}

func (s *ServiceContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ThriftParserIDENTIFIER)
}

func (s *ServiceContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, i)
}

func (s *ServiceContext) AllFunction_() []IFunction_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunction_Context)(nil)).Elem())
	var tst = make([]IFunction_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunction_Context)
		}
	}

	return tst
}

func (s *ServiceContext) Function_(i int) IFunction_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunction_Context)
}

func (s *ServiceContext) Type_annotations() IType_annotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *ServiceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterService(s)
	}
}

func (s *ServiceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitService(s)
	}
}

func (p *ThriftParser) Service() (localctx IServiceContext) {
	this := p
	_ = this

	localctx = NewServiceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ThriftParserRULE_service)
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
		p.Match(ThriftParserT__16)
	}
	{
		p.SetState(221)
		p.Match(ThriftParserIDENTIFIER)
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__17 {
		{
			p.SetState(222)
			p.Match(ThriftParserT__17)
		}
		{
			p.SetState(223)
			p.Match(ThriftParserIDENTIFIER)
		}

	}
	{
		p.SetState(226)
		p.Match(ThriftParserT__10)
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-24)&-(0x1f+1)) == 0 && ((1<<uint((_la-24)))&((1<<(ThriftParserT__23-24))|(1<<(ThriftParserT__24-24))|(1<<(ThriftParserT__25-24))|(1<<(ThriftParserT__27-24))|(1<<(ThriftParserT__30-24))|(1<<(ThriftParserT__31-24))|(1<<(ThriftParserTYPE_BOOL-24))|(1<<(ThriftParserTYPE_BYTE-24))|(1<<(ThriftParserTYPE_I16-24))|(1<<(ThriftParserTYPE_I32-24))|(1<<(ThriftParserTYPE_I64-24))|(1<<(ThriftParserTYPE_DOUBLE-24))|(1<<(ThriftParserTYPE_STRING-24))|(1<<(ThriftParserTYPE_BINARY-24))|(1<<(ThriftParserIDENTIFIER-24)))) != 0 {
		{
			p.SetState(227)
			p.Function_()
		}

		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(233)
		p.Match(ThriftParserT__11)
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(234)
			p.Type_annotations()
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
	p.RuleIndex = ThriftParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Field_type() IField_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *FieldContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *FieldContext) Field_id() IField_idContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_idContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_idContext)
}

func (s *FieldContext) Field_req() IField_reqContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_reqContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_reqContext)
}

func (s *FieldContext) Const_value() IConst_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *FieldContext) Type_annotations() IType_annotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *FieldContext) List_separator() IList_separatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_separatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitField(s)
	}
}

func (p *ThriftParser) Field() (localctx IFieldContext) {
	this := p
	_ = this

	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ThriftParserRULE_field)
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
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserINTEGER || _la == ThriftParserHEX_INTEGER {
		{
			p.SetState(237)
			p.Field_id()
		}

	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__19 || _la == ThriftParserT__20 {
		{
			p.SetState(240)
			p.Field_req()
		}

	}
	{
		p.SetState(243)
		p.Field_type()
	}
	{
		p.SetState(244)
		p.Match(ThriftParserIDENTIFIER)
	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__7 {
		{
			p.SetState(245)
			p.Match(ThriftParserT__7)
		}
		{
			p.SetState(246)
			p.Const_value()
		}

	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(249)
			p.Type_annotations()
		}

	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
		{
			p.SetState(252)
			p.List_separator()
		}

	}

	return localctx
}

// IField_idContext is an interface to support dynamic dispatch.
type IField_idContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_idContext differentiates from other interfaces.
	IsField_idContext()
}

type Field_idContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_idContext() *Field_idContext {
	var p = new(Field_idContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_field_id
	return p
}

func (*Field_idContext) IsField_idContext() {}

func NewField_idContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_idContext {
	var p = new(Field_idContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_field_id

	return p
}

func (s *Field_idContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_idContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *Field_idContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_idContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_idContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterField_id(s)
	}
}

func (s *Field_idContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitField_id(s)
	}
}

func (p *ThriftParser) Field_id() (localctx IField_idContext) {
	this := p
	_ = this

	localctx = NewField_idContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ThriftParserRULE_field_id)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Integer()
	}
	{
		p.SetState(256)
		p.Match(ThriftParserT__18)
	}

	return localctx
}

// IField_reqContext is an interface to support dynamic dispatch.
type IField_reqContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_reqContext differentiates from other interfaces.
	IsField_reqContext()
}

type Field_reqContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_reqContext() *Field_reqContext {
	var p = new(Field_reqContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_field_req
	return p
}

func (*Field_reqContext) IsField_reqContext() {}

func NewField_reqContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_reqContext {
	var p = new(Field_reqContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_field_req

	return p
}

func (s *Field_reqContext) GetParser() antlr.Parser { return s.parser }
func (s *Field_reqContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_reqContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_reqContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterField_req(s)
	}
}

func (s *Field_reqContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitField_req(s)
	}
}

func (p *ThriftParser) Field_req() (localctx IField_reqContext) {
	this := p
	_ = this

	localctx = NewField_reqContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ThriftParserRULE_field_req)
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
		p.SetState(258)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ThriftParserT__19 || _la == ThriftParserT__20) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunction_Context is an interface to support dynamic dispatch.
type IFunction_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_Context differentiates from other interfaces.
	IsFunction_Context()
}

type Function_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_Context() *Function_Context {
	var p = new(Function_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_function_
	return p
}

func (*Function_Context) IsFunction_Context() {}

func NewFunction_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_Context {
	var p = new(Function_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_function_

	return p
}

func (s *Function_Context) GetParser() antlr.Parser { return s.parser }

func (s *Function_Context) Function_type() IFunction_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_typeContext)
}

func (s *Function_Context) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Function_Context) Oneway() IOnewayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOnewayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOnewayContext)
}

func (s *Function_Context) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *Function_Context) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Function_Context) Throws_list() IThrows_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThrows_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThrows_listContext)
}

func (s *Function_Context) Type_annotations() IType_annotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Function_Context) List_separator() IList_separatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_separatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Function_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterFunction_(s)
	}
}

func (s *Function_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitFunction_(s)
	}
}

func (p *ThriftParser) Function_() (localctx IFunction_Context) {
	this := p
	_ = this

	localctx = NewFunction_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ThriftParserRULE_function_)
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
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__23 || _la == ThriftParserT__24 {
		{
			p.SetState(260)
			p.Oneway()
		}

	}
	{
		p.SetState(263)
		p.Function_type()
	}
	{
		p.SetState(264)
		p.Match(ThriftParserIDENTIFIER)
	}
	{
		p.SetState(265)
		p.Match(ThriftParserT__21)
	}
	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(ThriftParserT__19-20))|(1<<(ThriftParserT__20-20))|(1<<(ThriftParserT__27-20))|(1<<(ThriftParserT__30-20))|(1<<(ThriftParserT__31-20))|(1<<(ThriftParserINTEGER-20))|(1<<(ThriftParserHEX_INTEGER-20))|(1<<(ThriftParserTYPE_BOOL-20))|(1<<(ThriftParserTYPE_BYTE-20))|(1<<(ThriftParserTYPE_I16-20))|(1<<(ThriftParserTYPE_I32-20))|(1<<(ThriftParserTYPE_I64-20))|(1<<(ThriftParserTYPE_DOUBLE-20))|(1<<(ThriftParserTYPE_STRING-20))|(1<<(ThriftParserTYPE_BINARY-20))|(1<<(ThriftParserIDENTIFIER-20)))) != 0 {
		{
			p.SetState(266)
			p.Field()
		}

		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(272)
		p.Match(ThriftParserT__22)
	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__26 {
		{
			p.SetState(273)
			p.Throws_list()
		}

	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(276)
			p.Type_annotations()
		}

	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
		{
			p.SetState(279)
			p.List_separator()
		}

	}

	return localctx
}

// IOnewayContext is an interface to support dynamic dispatch.
type IOnewayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOnewayContext differentiates from other interfaces.
	IsOnewayContext()
}

type OnewayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOnewayContext() *OnewayContext {
	var p = new(OnewayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_oneway
	return p
}

func (*OnewayContext) IsOnewayContext() {}

func NewOnewayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OnewayContext {
	var p = new(OnewayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_oneway

	return p
}

func (s *OnewayContext) GetParser() antlr.Parser { return s.parser }
func (s *OnewayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OnewayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OnewayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterOneway(s)
	}
}

func (s *OnewayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitOneway(s)
	}
}

func (p *ThriftParser) Oneway() (localctx IOnewayContext) {
	this := p
	_ = this

	localctx = NewOnewayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ThriftParserRULE_oneway)
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
		p.SetState(282)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ThriftParserT__23 || _la == ThriftParserT__24) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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
	p.RuleIndex = ThriftParserRULE_function_type
	return p
}

func (*Function_typeContext) IsFunction_typeContext() {}

func NewFunction_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_typeContext {
	var p = new(Function_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_function_type

	return p
}

func (s *Function_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_typeContext) Field_type() IField_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Function_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterFunction_type(s)
	}
}

func (s *Function_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitFunction_type(s)
	}
}

func (p *ThriftParser) Function_type() (localctx IFunction_typeContext) {
	this := p
	_ = this

	localctx = NewFunction_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ThriftParserRULE_function_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	case ThriftParserT__27, ThriftParserT__30, ThriftParserT__31, ThriftParserTYPE_BOOL, ThriftParserTYPE_BYTE, ThriftParserTYPE_I16, ThriftParserTYPE_I32, ThriftParserTYPE_I64, ThriftParserTYPE_DOUBLE, ThriftParserTYPE_STRING, ThriftParserTYPE_BINARY, ThriftParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.Field_type()
		}

	case ThriftParserT__25:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.Match(ThriftParserT__25)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IThrows_listContext is an interface to support dynamic dispatch.
type IThrows_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThrows_listContext differentiates from other interfaces.
	IsThrows_listContext()
}

type Throws_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThrows_listContext() *Throws_listContext {
	var p = new(Throws_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_throws_list
	return p
}

func (*Throws_listContext) IsThrows_listContext() {}

func NewThrows_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Throws_listContext {
	var p = new(Throws_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_throws_list

	return p
}

func (s *Throws_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Throws_listContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *Throws_listContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Throws_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Throws_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Throws_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterThrows_list(s)
	}
}

func (s *Throws_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitThrows_list(s)
	}
}

func (p *ThriftParser) Throws_list() (localctx IThrows_listContext) {
	this := p
	_ = this

	localctx = NewThrows_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ThriftParserRULE_throws_list)
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
		p.Match(ThriftParserT__26)
	}
	{
		p.SetState(289)
		p.Match(ThriftParserT__21)
	}
	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-20)&-(0x1f+1)) == 0 && ((1<<uint((_la-20)))&((1<<(ThriftParserT__19-20))|(1<<(ThriftParserT__20-20))|(1<<(ThriftParserT__27-20))|(1<<(ThriftParserT__30-20))|(1<<(ThriftParserT__31-20))|(1<<(ThriftParserINTEGER-20))|(1<<(ThriftParserHEX_INTEGER-20))|(1<<(ThriftParserTYPE_BOOL-20))|(1<<(ThriftParserTYPE_BYTE-20))|(1<<(ThriftParserTYPE_I16-20))|(1<<(ThriftParserTYPE_I32-20))|(1<<(ThriftParserTYPE_I64-20))|(1<<(ThriftParserTYPE_DOUBLE-20))|(1<<(ThriftParserTYPE_STRING-20))|(1<<(ThriftParserTYPE_BINARY-20))|(1<<(ThriftParserIDENTIFIER-20)))) != 0 {
		{
			p.SetState(290)
			p.Field()
		}

		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(296)
		p.Match(ThriftParserT__22)
	}

	return localctx
}

// IType_annotationsContext is an interface to support dynamic dispatch.
type IType_annotationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_annotationsContext differentiates from other interfaces.
	IsType_annotationsContext()
}

type Type_annotationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_annotationsContext() *Type_annotationsContext {
	var p = new(Type_annotationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_type_annotations
	return p
}

func (*Type_annotationsContext) IsType_annotationsContext() {}

func NewType_annotationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_annotationsContext {
	var p = new(Type_annotationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_type_annotations

	return p
}

func (s *Type_annotationsContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_annotationsContext) AllType_annotation() []IType_annotationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IType_annotationContext)(nil)).Elem())
	var tst = make([]IType_annotationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IType_annotationContext)
		}
	}

	return tst
}

func (s *Type_annotationsContext) Type_annotation(i int) IType_annotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IType_annotationContext)
}

func (s *Type_annotationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_annotationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_annotationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterType_annotations(s)
	}
}

func (s *Type_annotationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitType_annotations(s)
	}
}

func (p *ThriftParser) Type_annotations() (localctx IType_annotationsContext) {
	this := p
	_ = this

	localctx = NewType_annotationsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ThriftParserRULE_type_annotations)
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
		p.SetState(298)
		p.Match(ThriftParserT__21)
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ThriftParserIDENTIFIER {
		{
			p.SetState(299)
			p.Type_annotation()
		}

		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(305)
		p.Match(ThriftParserT__22)
	}

	return localctx
}

// IType_annotationContext is an interface to support dynamic dispatch.
type IType_annotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_annotationContext differentiates from other interfaces.
	IsType_annotationContext()
}

type Type_annotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_annotationContext() *Type_annotationContext {
	var p = new(Type_annotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_type_annotation
	return p
}

func (*Type_annotationContext) IsType_annotationContext() {}

func NewType_annotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_annotationContext {
	var p = new(Type_annotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_type_annotation

	return p
}

func (s *Type_annotationContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_annotationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Type_annotationContext) Annotation_value() IAnnotation_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnnotation_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnnotation_valueContext)
}

func (s *Type_annotationContext) List_separator() IList_separatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_separatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Type_annotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_annotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_annotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterType_annotation(s)
	}
}

func (s *Type_annotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitType_annotation(s)
	}
}

func (p *ThriftParser) Type_annotation() (localctx IType_annotationContext) {
	this := p
	_ = this

	localctx = NewType_annotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ThriftParserRULE_type_annotation)
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
		p.SetState(307)
		p.Match(ThriftParserIDENTIFIER)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__7 {
		{
			p.SetState(308)
			p.Match(ThriftParserT__7)
		}
		{
			p.SetState(309)
			p.Annotation_value()
		}

	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
		{
			p.SetState(312)
			p.List_separator()
		}

	}

	return localctx
}

// IAnnotation_valueContext is an interface to support dynamic dispatch.
type IAnnotation_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnnotation_valueContext differentiates from other interfaces.
	IsAnnotation_valueContext()
}

type Annotation_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnnotation_valueContext() *Annotation_valueContext {
	var p = new(Annotation_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_annotation_value
	return p
}

func (*Annotation_valueContext) IsAnnotation_valueContext() {}

func NewAnnotation_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Annotation_valueContext {
	var p = new(Annotation_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_annotation_value

	return p
}

func (s *Annotation_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Annotation_valueContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *Annotation_valueContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, 0)
}

func (s *Annotation_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Annotation_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Annotation_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterAnnotation_value(s)
	}
}

func (s *Annotation_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitAnnotation_value(s)
	}
}

func (p *ThriftParser) Annotation_value() (localctx IAnnotation_valueContext) {
	this := p
	_ = this

	localctx = NewAnnotation_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ThriftParserRULE_annotation_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(317)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserINTEGER, ThriftParserHEX_INTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(315)
			p.Integer()
		}

	case ThriftParserLITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(316)
			p.Match(ThriftParserLITERAL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = ThriftParserRULE_field_type
	return p
}

func (*Field_typeContext) IsField_typeContext() {}

func NewField_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_typeContext {
	var p = new(Field_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_field_type

	return p
}

func (s *Field_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_typeContext) Base_type() IBase_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBase_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBase_typeContext)
}

func (s *Field_typeContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Field_typeContext) Container_type() IContainer_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContainer_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContainer_typeContext)
}

func (s *Field_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterField_type(s)
	}
}

func (s *Field_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitField_type(s)
	}
}

func (p *ThriftParser) Field_type() (localctx IField_typeContext) {
	this := p
	_ = this

	localctx = NewField_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ThriftParserRULE_field_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(322)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserTYPE_BOOL, ThriftParserTYPE_BYTE, ThriftParserTYPE_I16, ThriftParserTYPE_I32, ThriftParserTYPE_I64, ThriftParserTYPE_DOUBLE, ThriftParserTYPE_STRING, ThriftParserTYPE_BINARY:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(319)
			p.Base_type()
		}

	case ThriftParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(320)
			p.Match(ThriftParserIDENTIFIER)
		}

	case ThriftParserT__27, ThriftParserT__30, ThriftParserT__31:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(321)
			p.Container_type()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBase_typeContext is an interface to support dynamic dispatch.
type IBase_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBase_typeContext differentiates from other interfaces.
	IsBase_typeContext()
}

type Base_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBase_typeContext() *Base_typeContext {
	var p = new(Base_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_base_type
	return p
}

func (*Base_typeContext) IsBase_typeContext() {}

func NewBase_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Base_typeContext {
	var p = new(Base_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_base_type

	return p
}

func (s *Base_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Base_typeContext) Real_base_type() IReal_base_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReal_base_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReal_base_typeContext)
}

func (s *Base_typeContext) Type_annotations() IType_annotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Base_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Base_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Base_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterBase_type(s)
	}
}

func (s *Base_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitBase_type(s)
	}
}

func (p *ThriftParser) Base_type() (localctx IBase_typeContext) {
	this := p
	_ = this

	localctx = NewBase_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ThriftParserRULE_base_type)
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
		p.SetState(324)
		p.Real_base_type()
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(325)
			p.Type_annotations()
		}

	}

	return localctx
}

// IContainer_typeContext is an interface to support dynamic dispatch.
type IContainer_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContainer_typeContext differentiates from other interfaces.
	IsContainer_typeContext()
}

type Container_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContainer_typeContext() *Container_typeContext {
	var p = new(Container_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_container_type
	return p
}

func (*Container_typeContext) IsContainer_typeContext() {}

func NewContainer_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Container_typeContext {
	var p = new(Container_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_container_type

	return p
}

func (s *Container_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Container_typeContext) Map_type() IMap_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMap_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMap_typeContext)
}

func (s *Container_typeContext) Set_type() ISet_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_typeContext)
}

func (s *Container_typeContext) List_type() IList_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_typeContext)
}

func (s *Container_typeContext) Type_annotations() IType_annotationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_annotationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_annotationsContext)
}

func (s *Container_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Container_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Container_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterContainer_type(s)
	}
}

func (s *Container_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitContainer_type(s)
	}
}

func (p *ThriftParser) Container_type() (localctx IContainer_typeContext) {
	this := p
	_ = this

	localctx = NewContainer_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ThriftParserRULE_container_type)
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
	p.SetState(331)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserT__27:
		{
			p.SetState(328)
			p.Map_type()
		}

	case ThriftParserT__30:
		{
			p.SetState(329)
			p.Set_type()
		}

	case ThriftParserT__31:
		{
			p.SetState(330)
			p.List_type()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__21 {
		{
			p.SetState(333)
			p.Type_annotations()
		}

	}

	return localctx
}

// IMap_typeContext is an interface to support dynamic dispatch.
type IMap_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMap_typeContext differentiates from other interfaces.
	IsMap_typeContext()
}

type Map_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMap_typeContext() *Map_typeContext {
	var p = new(Map_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_map_type
	return p
}

func (*Map_typeContext) IsMap_typeContext() {}

func NewMap_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Map_typeContext {
	var p = new(Map_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_map_type

	return p
}

func (s *Map_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Map_typeContext) AllField_type() []IField_typeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IField_typeContext)(nil)).Elem())
	var tst = make([]IField_typeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IField_typeContext)
		}
	}

	return tst
}

func (s *Map_typeContext) Field_type(i int) IField_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_typeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Map_typeContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ThriftParserCOMMA, 0)
}

func (s *Map_typeContext) Cpp_type() ICpp_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICpp_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICpp_typeContext)
}

func (s *Map_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Map_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Map_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterMap_type(s)
	}
}

func (s *Map_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitMap_type(s)
	}
}

func (p *ThriftParser) Map_type() (localctx IMap_typeContext) {
	this := p
	_ = this

	localctx = NewMap_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ThriftParserRULE_map_type)
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
		p.SetState(336)
		p.Match(ThriftParserT__27)
	}
	p.SetState(338)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__32 {
		{
			p.SetState(337)
			p.Cpp_type()
		}

	}
	{
		p.SetState(340)
		p.Match(ThriftParserT__28)
	}
	{
		p.SetState(341)
		p.Field_type()
	}
	{
		p.SetState(342)
		p.Match(ThriftParserCOMMA)
	}
	{
		p.SetState(343)
		p.Field_type()
	}
	{
		p.SetState(344)
		p.Match(ThriftParserT__29)
	}

	return localctx
}

// ISet_typeContext is an interface to support dynamic dispatch.
type ISet_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_typeContext differentiates from other interfaces.
	IsSet_typeContext()
}

type Set_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_typeContext() *Set_typeContext {
	var p = new(Set_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_set_type
	return p
}

func (*Set_typeContext) IsSet_typeContext() {}

func NewSet_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_typeContext {
	var p = new(Set_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_set_type

	return p
}

func (s *Set_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_typeContext) Field_type() IField_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *Set_typeContext) Cpp_type() ICpp_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICpp_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICpp_typeContext)
}

func (s *Set_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterSet_type(s)
	}
}

func (s *Set_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitSet_type(s)
	}
}

func (p *ThriftParser) Set_type() (localctx ISet_typeContext) {
	this := p
	_ = this

	localctx = NewSet_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ThriftParserRULE_set_type)
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
		p.Match(ThriftParserT__30)
	}
	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__32 {
		{
			p.SetState(347)
			p.Cpp_type()
		}

	}
	{
		p.SetState(350)
		p.Match(ThriftParserT__28)
	}
	{
		p.SetState(351)
		p.Field_type()
	}
	{
		p.SetState(352)
		p.Match(ThriftParserT__29)
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
	p.RuleIndex = ThriftParserRULE_list_type
	return p
}

func (*List_typeContext) IsList_typeContext() {}

func NewList_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_typeContext {
	var p = new(List_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_list_type

	return p
}

func (s *List_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *List_typeContext) Field_type() IField_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_typeContext)
}

func (s *List_typeContext) Cpp_type() ICpp_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICpp_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICpp_typeContext)
}

func (s *List_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterList_type(s)
	}
}

func (s *List_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitList_type(s)
	}
}

func (p *ThriftParser) List_type() (localctx IList_typeContext) {
	this := p
	_ = this

	localctx = NewList_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ThriftParserRULE_list_type)
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
		p.SetState(354)
		p.Match(ThriftParserT__31)
	}
	{
		p.SetState(355)
		p.Match(ThriftParserT__28)
	}
	{
		p.SetState(356)
		p.Field_type()
	}
	{
		p.SetState(357)
		p.Match(ThriftParserT__29)
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__32 {
		{
			p.SetState(358)
			p.Cpp_type()
		}

	}

	return localctx
}

// ICpp_typeContext is an interface to support dynamic dispatch.
type ICpp_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCpp_typeContext differentiates from other interfaces.
	IsCpp_typeContext()
}

type Cpp_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCpp_typeContext() *Cpp_typeContext {
	var p = new(Cpp_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_cpp_type
	return p
}

func (*Cpp_typeContext) IsCpp_typeContext() {}

func NewCpp_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Cpp_typeContext {
	var p = new(Cpp_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_cpp_type

	return p
}

func (s *Cpp_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Cpp_typeContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, 0)
}

func (s *Cpp_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cpp_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Cpp_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterCpp_type(s)
	}
}

func (s *Cpp_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitCpp_type(s)
	}
}

func (p *ThriftParser) Cpp_type() (localctx ICpp_typeContext) {
	this := p
	_ = this

	localctx = NewCpp_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ThriftParserRULE_cpp_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(361)
		p.Match(ThriftParserT__32)
	}
	{
		p.SetState(362)
		p.Match(ThriftParserLITERAL)
	}

	return localctx
}

// IConst_valueContext is an interface to support dynamic dispatch.
type IConst_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConst_valueContext differentiates from other interfaces.
	IsConst_valueContext()
}

type Const_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_valueContext() *Const_valueContext {
	var p = new(Const_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_const_value
	return p
}

func (*Const_valueContext) IsConst_valueContext() {}

func NewConst_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_valueContext {
	var p = new(Const_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_const_value

	return p
}

func (s *Const_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_valueContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *Const_valueContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(ThriftParserDOUBLE, 0)
}

func (s *Const_valueContext) LITERAL() antlr.TerminalNode {
	return s.GetToken(ThriftParserLITERAL, 0)
}

func (s *Const_valueContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ThriftParserIDENTIFIER, 0)
}

func (s *Const_valueContext) Const_list() IConst_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConst_listContext)
}

func (s *Const_valueContext) Const_map() IConst_mapContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_mapContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConst_mapContext)
}

func (s *Const_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterConst_value(s)
	}
}

func (s *Const_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitConst_value(s)
	}
}

func (p *ThriftParser) Const_value() (localctx IConst_valueContext) {
	this := p
	_ = this

	localctx = NewConst_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ThriftParserRULE_const_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(370)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ThriftParserINTEGER, ThriftParserHEX_INTEGER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(364)
			p.Integer()
		}

	case ThriftParserDOUBLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(365)
			p.Match(ThriftParserDOUBLE)
		}

	case ThriftParserLITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(366)
			p.Match(ThriftParserLITERAL)
		}

	case ThriftParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(367)
			p.Match(ThriftParserIDENTIFIER)
		}

	case ThriftParserT__33:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(368)
			p.Const_list()
		}

	case ThriftParserT__10:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(369)
			p.Const_map()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = ThriftParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(ThriftParserINTEGER, 0)
}

func (s *IntegerContext) HEX_INTEGER() antlr.TerminalNode {
	return s.GetToken(ThriftParserHEX_INTEGER, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (p *ThriftParser) Integer() (localctx IIntegerContext) {
	this := p
	_ = this

	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ThriftParserRULE_integer)
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
		p.SetState(372)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ThriftParserINTEGER || _la == ThriftParserHEX_INTEGER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConst_listContext is an interface to support dynamic dispatch.
type IConst_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConst_listContext differentiates from other interfaces.
	IsConst_listContext()
}

type Const_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_listContext() *Const_listContext {
	var p = new(Const_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_const_list
	return p
}

func (*Const_listContext) IsConst_listContext() {}

func NewConst_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_listContext {
	var p = new(Const_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_const_list

	return p
}

func (s *Const_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_listContext) AllConst_value() []IConst_valueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConst_valueContext)(nil)).Elem())
	var tst = make([]IConst_valueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConst_valueContext)
		}
	}

	return tst
}

func (s *Const_listContext) Const_value(i int) IConst_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_valueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *Const_listContext) AllList_separator() []IList_separatorContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IList_separatorContext)(nil)).Elem())
	var tst = make([]IList_separatorContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IList_separatorContext)
		}
	}

	return tst
}

func (s *Const_listContext) List_separator(i int) IList_separatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_separatorContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Const_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterConst_list(s)
	}
}

func (s *Const_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitConst_list(s)
	}
}

func (p *ThriftParser) Const_list() (localctx IConst_listContext) {
	this := p
	_ = this

	localctx = NewConst_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ThriftParserRULE_const_list)
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
		p.Match(ThriftParserT__33)
	}
	p.SetState(381)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ThriftParserT__10 || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(ThriftParserT__33-34))|(1<<(ThriftParserINTEGER-34))|(1<<(ThriftParserHEX_INTEGER-34))|(1<<(ThriftParserDOUBLE-34))|(1<<(ThriftParserLITERAL-34))|(1<<(ThriftParserIDENTIFIER-34)))) != 0) {
		{
			p.SetState(375)
			p.Const_value()
		}
		p.SetState(377)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
			{
				p.SetState(376)
				p.List_separator()
			}

		}

		p.SetState(383)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(384)
		p.Match(ThriftParserT__34)
	}

	return localctx
}

// IConst_map_entryContext is an interface to support dynamic dispatch.
type IConst_map_entryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConst_map_entryContext differentiates from other interfaces.
	IsConst_map_entryContext()
}

type Const_map_entryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_map_entryContext() *Const_map_entryContext {
	var p = new(Const_map_entryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_const_map_entry
	return p
}

func (*Const_map_entryContext) IsConst_map_entryContext() {}

func NewConst_map_entryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_map_entryContext {
	var p = new(Const_map_entryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_const_map_entry

	return p
}

func (s *Const_map_entryContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_map_entryContext) AllConst_value() []IConst_valueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConst_valueContext)(nil)).Elem())
	var tst = make([]IConst_valueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConst_valueContext)
		}
	}

	return tst
}

func (s *Const_map_entryContext) Const_value(i int) IConst_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_valueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConst_valueContext)
}

func (s *Const_map_entryContext) List_separator() IList_separatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_separatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_separatorContext)
}

func (s *Const_map_entryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_map_entryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_map_entryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterConst_map_entry(s)
	}
}

func (s *Const_map_entryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitConst_map_entry(s)
	}
}

func (p *ThriftParser) Const_map_entry() (localctx IConst_map_entryContext) {
	this := p
	_ = this

	localctx = NewConst_map_entryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ThriftParserRULE_const_map_entry)
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
		p.SetState(386)
		p.Const_value()
	}
	{
		p.SetState(387)
		p.Match(ThriftParserT__18)
	}
	{
		p.SetState(388)
		p.Const_value()
	}
	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == ThriftParserT__35 || _la == ThriftParserCOMMA {
		{
			p.SetState(389)
			p.List_separator()
		}

	}

	return localctx
}

// IConst_mapContext is an interface to support dynamic dispatch.
type IConst_mapContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConst_mapContext differentiates from other interfaces.
	IsConst_mapContext()
}

type Const_mapContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_mapContext() *Const_mapContext {
	var p = new(Const_mapContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_const_map
	return p
}

func (*Const_mapContext) IsConst_mapContext() {}

func NewConst_mapContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_mapContext {
	var p = new(Const_mapContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_const_map

	return p
}

func (s *Const_mapContext) GetParser() antlr.Parser { return s.parser }

func (s *Const_mapContext) AllConst_map_entry() []IConst_map_entryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConst_map_entryContext)(nil)).Elem())
	var tst = make([]IConst_map_entryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConst_map_entryContext)
		}
	}

	return tst
}

func (s *Const_mapContext) Const_map_entry(i int) IConst_map_entryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_map_entryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConst_map_entryContext)
}

func (s *Const_mapContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_mapContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_mapContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterConst_map(s)
	}
}

func (s *Const_mapContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitConst_map(s)
	}
}

func (p *ThriftParser) Const_map() (localctx IConst_mapContext) {
	this := p
	_ = this

	localctx = NewConst_mapContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ThriftParserRULE_const_map)
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
		p.Match(ThriftParserT__10)
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ThriftParserT__10 || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(ThriftParserT__33-34))|(1<<(ThriftParserINTEGER-34))|(1<<(ThriftParserHEX_INTEGER-34))|(1<<(ThriftParserDOUBLE-34))|(1<<(ThriftParserLITERAL-34))|(1<<(ThriftParserIDENTIFIER-34)))) != 0) {
		{
			p.SetState(393)
			p.Const_map_entry()
		}

		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(399)
		p.Match(ThriftParserT__11)
	}

	return localctx
}

// IList_separatorContext is an interface to support dynamic dispatch.
type IList_separatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_separatorContext differentiates from other interfaces.
	IsList_separatorContext()
}

type List_separatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_separatorContext() *List_separatorContext {
	var p = new(List_separatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_list_separator
	return p
}

func (*List_separatorContext) IsList_separatorContext() {}

func NewList_separatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_separatorContext {
	var p = new(List_separatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_list_separator

	return p
}

func (s *List_separatorContext) GetParser() antlr.Parser { return s.parser }

func (s *List_separatorContext) COMMA() antlr.TerminalNode {
	return s.GetToken(ThriftParserCOMMA, 0)
}

func (s *List_separatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_separatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_separatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterList_separator(s)
	}
}

func (s *List_separatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitList_separator(s)
	}
}

func (p *ThriftParser) List_separator() (localctx IList_separatorContext) {
	this := p
	_ = this

	localctx = NewList_separatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ThriftParserRULE_list_separator)
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
		p.SetState(401)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ThriftParserT__35 || _la == ThriftParserCOMMA) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IReal_base_typeContext is an interface to support dynamic dispatch.
type IReal_base_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReal_base_typeContext differentiates from other interfaces.
	IsReal_base_typeContext()
}

type Real_base_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReal_base_typeContext() *Real_base_typeContext {
	var p = new(Real_base_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ThriftParserRULE_real_base_type
	return p
}

func (*Real_base_typeContext) IsReal_base_typeContext() {}

func NewReal_base_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Real_base_typeContext {
	var p = new(Real_base_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ThriftParserRULE_real_base_type

	return p
}

func (s *Real_base_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Real_base_typeContext) TYPE_BOOL() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_BOOL, 0)
}

func (s *Real_base_typeContext) TYPE_BYTE() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_BYTE, 0)
}

func (s *Real_base_typeContext) TYPE_I16() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_I16, 0)
}

func (s *Real_base_typeContext) TYPE_I32() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_I32, 0)
}

func (s *Real_base_typeContext) TYPE_I64() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_I64, 0)
}

func (s *Real_base_typeContext) TYPE_DOUBLE() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_DOUBLE, 0)
}

func (s *Real_base_typeContext) TYPE_STRING() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_STRING, 0)
}

func (s *Real_base_typeContext) TYPE_BINARY() antlr.TerminalNode {
	return s.GetToken(ThriftParserTYPE_BINARY, 0)
}

func (s *Real_base_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Real_base_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Real_base_typeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.EnterReal_base_type(s)
	}
}

func (s *Real_base_typeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ThriftListener); ok {
		listenerT.ExitReal_base_type(s)
	}
}

func (p *ThriftParser) Real_base_type() (localctx IReal_base_typeContext) {
	this := p
	_ = this

	localctx = NewReal_base_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ThriftParserRULE_real_base_type)
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
		p.SetState(403)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-40)&-(0x1f+1)) == 0 && ((1<<uint((_la-40)))&((1<<(ThriftParserTYPE_BOOL-40))|(1<<(ThriftParserTYPE_BYTE-40))|(1<<(ThriftParserTYPE_I16-40))|(1<<(ThriftParserTYPE_I32-40))|(1<<(ThriftParserTYPE_I64-40))|(1<<(ThriftParserTYPE_DOUBLE-40))|(1<<(ThriftParserTYPE_STRING-40))|(1<<(ThriftParserTYPE_BINARY-40)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
