// Code generated from ISL.g4 by ANTLR 4.9.3. DO NOT EDIT.

package isl // ISL
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 444,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 6, 2, 24, 10,
	2, 13, 2, 14, 2, 25, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 32, 10, 3, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 6, 4, 39, 10, 4, 13, 4, 14, 4, 40, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 7, 4, 58, 10, 4, 12, 4, 14, 4, 61, 11, 4, 3, 4, 3, 4, 3, 4, 5, 4, 66,
	10, 4, 3, 5, 3, 5, 3, 5, 6, 5, 71, 10, 5, 13, 5, 14, 5, 72, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 6, 5, 80, 10, 5, 13, 5, 14, 5, 81, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 7, 5, 101, 10, 5, 12, 5, 14, 5, 104, 11, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 114, 10, 5, 12, 5, 14, 5, 117, 11,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 127, 10, 5, 12,
	5, 14, 5, 130, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 7, 5, 144, 10, 5, 12, 5, 14, 5, 147, 11, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5,
	161, 10, 5, 12, 5, 14, 5, 164, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 178, 10, 5, 12, 5, 14, 5,
	181, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 7, 5, 195, 10, 5, 12, 5, 14, 5, 198, 11, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 213,
	10, 5, 12, 5, 14, 5, 216, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 7, 5, 225, 10, 5, 12, 5, 14, 5, 228, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 239, 10, 5, 13, 5, 14, 5, 240, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 252, 10, 5, 12, 5,
	14, 5, 255, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 275, 10, 5,
	13, 5, 14, 5, 276, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 286,
	10, 5, 13, 5, 14, 5, 287, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 308,
	10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 315, 10, 6, 12, 6, 14, 6, 318,
	11, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 329,
	10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 337, 10, 7, 12, 7, 14,
	7, 340, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5,
	7, 351, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 6, 8, 376, 10, 8, 13, 8, 14, 8, 377, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 392, 10, 8, 3, 8, 3,
	8, 5, 8, 396, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 6, 9, 414, 10, 9, 13, 9, 14,
	9, 415, 3, 9, 5, 9, 419, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 433, 10, 9, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 2, 2, 12, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 2, 3, 3, 2, 40, 41, 2, 511, 2, 23, 3, 2, 2, 2,
	4, 31, 3, 2, 2, 2, 6, 65, 3, 2, 2, 2, 8, 307, 3, 2, 2, 2, 10, 328, 3, 2,
	2, 2, 12, 350, 3, 2, 2, 2, 14, 395, 3, 2, 2, 2, 16, 432, 3, 2, 2, 2, 18,
	434, 3, 2, 2, 2, 20, 441, 3, 2, 2, 2, 22, 24, 5, 4, 3, 2, 23, 22, 3, 2,
	2, 2, 24, 25, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 3,
	3, 2, 2, 2, 27, 32, 5, 6, 4, 2, 28, 32, 5, 8, 5, 2, 29, 32, 5, 14, 8, 2,
	30, 32, 5, 16, 9, 2, 31, 27, 3, 2, 2, 2, 31, 28, 3, 2, 2, 2, 31, 29, 3,
	2, 2, 2, 31, 30, 3, 2, 2, 2, 32, 5, 3, 2, 2, 2, 33, 34, 7, 3, 2, 2, 34,
	35, 7, 4, 2, 2, 35, 36, 7, 3, 2, 2, 36, 38, 5, 20, 11, 2, 37, 39, 7, 41,
	2, 2, 38, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 38, 3, 2, 2, 2, 40, 41,
	3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 7, 5, 2, 2, 43, 44, 5, 8, 5, 2,
	44, 45, 7, 5, 2, 2, 45, 66, 3, 2, 2, 2, 46, 47, 7, 3, 2, 2, 47, 48, 7,
	4, 2, 2, 48, 49, 5, 20, 11, 2, 49, 50, 5, 8, 5, 2, 50, 51, 7, 5, 2, 2,
	51, 66, 3, 2, 2, 2, 52, 53, 7, 3, 2, 2, 53, 54, 7, 6, 2, 2, 54, 55, 5,
	20, 11, 2, 55, 59, 7, 3, 2, 2, 56, 58, 5, 20, 11, 2, 57, 56, 3, 2, 2, 2,
	58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 62, 3,
	2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 63, 7, 5, 2, 2, 63, 64, 7, 5, 2, 2, 64,
	66, 3, 2, 2, 2, 65, 33, 3, 2, 2, 2, 65, 46, 3, 2, 2, 2, 65, 52, 3, 2, 2,
	2, 66, 7, 3, 2, 2, 2, 67, 68, 7, 3, 2, 2, 68, 70, 7, 7, 2, 2, 69, 71, 5,
	8, 5, 2, 70, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72,
	73, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 75, 7, 5, 2, 2, 75, 308, 3, 2,
	2, 2, 76, 77, 7, 3, 2, 2, 77, 79, 7, 8, 2, 2, 78, 80, 5, 8, 5, 2, 79, 78,
	3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2, 2,
	82, 83, 3, 2, 2, 2, 83, 84, 7, 5, 2, 2, 84, 308, 3, 2, 2, 2, 85, 86, 7,
	3, 2, 2, 86, 87, 7, 9, 2, 2, 87, 88, 7, 41, 2, 2, 88, 89, 5, 8, 5, 2, 89,
	90, 7, 5, 2, 2, 90, 308, 3, 2, 2, 2, 91, 92, 7, 3, 2, 2, 92, 93, 7, 10,
	2, 2, 93, 94, 5, 8, 5, 2, 94, 95, 7, 5, 2, 2, 95, 308, 3, 2, 2, 2, 96,
	97, 7, 3, 2, 2, 97, 98, 7, 11, 2, 2, 98, 102, 7, 3, 2, 2, 99, 101, 7, 41,
	2, 2, 100, 99, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2,
	102, 103, 3, 2, 2, 2, 103, 105, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105,
	106, 7, 5, 2, 2, 106, 107, 5, 8, 5, 2, 107, 108, 7, 5, 2, 2, 108, 308,
	3, 2, 2, 2, 109, 110, 7, 3, 2, 2, 110, 111, 7, 12, 2, 2, 111, 115, 7, 3,
	2, 2, 112, 114, 7, 41, 2, 2, 113, 112, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2,
	115, 113, 3, 2, 2, 2, 115, 116, 3, 2, 2, 2, 116, 118, 3, 2, 2, 2, 117,
	115, 3, 2, 2, 2, 118, 119, 7, 5, 2, 2, 119, 120, 5, 8, 5, 2, 120, 121,
	7, 5, 2, 2, 121, 308, 3, 2, 2, 2, 122, 123, 7, 3, 2, 2, 123, 124, 7, 13,
	2, 2, 124, 128, 7, 14, 2, 2, 125, 127, 5, 6, 4, 2, 126, 125, 3, 2, 2, 2,
	127, 130, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 128, 129, 3, 2, 2, 2, 129,
	131, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 131, 132, 7, 15, 2, 2, 132, 133,
	5, 8, 5, 2, 133, 134, 7, 5, 2, 2, 134, 308, 3, 2, 2, 2, 135, 136, 7, 3,
	2, 2, 136, 137, 7, 16, 2, 2, 137, 145, 7, 3, 2, 2, 138, 139, 7, 14, 2,
	2, 139, 140, 5, 20, 11, 2, 140, 141, 5, 8, 5, 2, 141, 142, 7, 15, 2, 2,
	142, 144, 3, 2, 2, 2, 143, 138, 3, 2, 2, 2, 144, 147, 3, 2, 2, 2, 145,
	143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 148, 3, 2, 2, 2, 147, 145,
	3, 2, 2, 2, 148, 149, 7, 5, 2, 2, 149, 150, 5, 8, 5, 2, 150, 151, 7, 5,
	2, 2, 151, 308, 3, 2, 2, 2, 152, 153, 7, 3, 2, 2, 153, 154, 7, 17, 2, 2,
	154, 162, 7, 3, 2, 2, 155, 156, 7, 14, 2, 2, 156, 157, 5, 20, 11, 2, 157,
	158, 5, 8, 5, 2, 158, 159, 7, 15, 2, 2, 159, 161, 3, 2, 2, 2, 160, 155,
	3, 2, 2, 2, 161, 164, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 162, 163, 3, 2,
	2, 2, 163, 165, 3, 2, 2, 2, 164, 162, 3, 2, 2, 2, 165, 166, 7, 5, 2, 2,
	166, 167, 5, 8, 5, 2, 167, 168, 7, 5, 2, 2, 168, 308, 3, 2, 2, 2, 169,
	170, 7, 3, 2, 2, 170, 171, 7, 18, 2, 2, 171, 179, 7, 3, 2, 2, 172, 173,
	7, 14, 2, 2, 173, 174, 5, 20, 11, 2, 174, 175, 5, 8, 5, 2, 175, 176, 7,
	15, 2, 2, 176, 178, 3, 2, 2, 2, 177, 172, 3, 2, 2, 2, 178, 181, 3, 2, 2,
	2, 179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 182, 3, 2, 2, 2, 181,
	179, 3, 2, 2, 2, 182, 183, 7, 5, 2, 2, 183, 184, 5, 8, 5, 2, 184, 185,
	7, 5, 2, 2, 185, 308, 3, 2, 2, 2, 186, 187, 7, 3, 2, 2, 187, 188, 7, 19,
	2, 2, 188, 196, 7, 3, 2, 2, 189, 190, 7, 14, 2, 2, 190, 191, 5, 20, 11,
	2, 191, 192, 5, 8, 5, 2, 192, 193, 7, 15, 2, 2, 193, 195, 3, 2, 2, 2, 194,
	189, 3, 2, 2, 2, 195, 198, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 196, 197,
	3, 2, 2, 2, 197, 199, 3, 2, 2, 2, 198, 196, 3, 2, 2, 2, 199, 200, 7, 5,
	2, 2, 200, 201, 5, 8, 5, 2, 201, 202, 7, 5, 2, 2, 202, 308, 3, 2, 2, 2,
	203, 204, 7, 3, 2, 2, 204, 205, 7, 20, 2, 2, 205, 206, 5, 20, 11, 2, 206,
	214, 7, 3, 2, 2, 207, 208, 7, 14, 2, 2, 208, 209, 5, 20, 11, 2, 209, 210,
	5, 8, 5, 2, 210, 211, 7, 15, 2, 2, 211, 213, 3, 2, 2, 2, 212, 207, 3, 2,
	2, 2, 213, 216, 3, 2, 2, 2, 214, 212, 3, 2, 2, 2, 214, 215, 3, 2, 2, 2,
	215, 217, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 217, 218, 7, 5, 2, 2, 218,
	219, 5, 8, 5, 2, 219, 220, 7, 5, 2, 2, 220, 308, 3, 2, 2, 2, 221, 222,
	7, 3, 2, 2, 222, 226, 5, 8, 5, 2, 223, 225, 5, 8, 5, 2, 224, 223, 3, 2,
	2, 2, 225, 228, 3, 2, 2, 2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2,
	227, 229, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2, 229, 230, 7, 5, 2, 2, 230,
	308, 3, 2, 2, 2, 231, 232, 7, 3, 2, 2, 232, 238, 7, 21, 2, 2, 233, 234,
	7, 14, 2, 2, 234, 235, 5, 8, 5, 2, 235, 236, 5, 8, 5, 2, 236, 237, 7, 15,
	2, 2, 237, 239, 3, 2, 2, 2, 238, 233, 3, 2, 2, 2, 239, 240, 3, 2, 2, 2,
	240, 238, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242,
	243, 7, 5, 2, 2, 243, 308, 3, 2, 2, 2, 244, 245, 7, 3, 2, 2, 245, 253,
	7, 21, 2, 2, 246, 247, 7, 14, 2, 2, 247, 248, 5, 8, 5, 2, 248, 249, 5,
	8, 5, 2, 249, 250, 7, 15, 2, 2, 250, 252, 3, 2, 2, 2, 251, 246, 3, 2, 2,
	2, 252, 255, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254,
	256, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 256, 257, 7, 14, 2, 2, 257, 258,
	7, 22, 2, 2, 258, 259, 5, 8, 5, 2, 259, 260, 7, 15, 2, 2, 260, 261, 7,
	5, 2, 2, 261, 308, 3, 2, 2, 2, 262, 263, 7, 3, 2, 2, 263, 264, 7, 23, 2,
	2, 264, 265, 5, 8, 5, 2, 265, 266, 5, 8, 5, 2, 266, 267, 5, 8, 5, 2, 267,
	268, 7, 5, 2, 2, 268, 308, 3, 2, 2, 2, 269, 270, 7, 3, 2, 2, 270, 271,
	7, 24, 2, 2, 271, 272, 5, 8, 5, 2, 272, 274, 5, 8, 5, 2, 273, 275, 5, 8,
	5, 2, 274, 273, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2,
	276, 277, 3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 279, 7, 5, 2, 2, 279,
	308, 3, 2, 2, 2, 280, 281, 7, 3, 2, 2, 281, 282, 7, 25, 2, 2, 282, 283,
	5, 8, 5, 2, 283, 285, 5, 8, 5, 2, 284, 286, 5, 8, 5, 2, 285, 284, 3, 2,
	2, 2, 286, 287, 3, 2, 2, 2, 287, 285, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2,
	288, 289, 3, 2, 2, 2, 289, 290, 7, 5, 2, 2, 290, 308, 3, 2, 2, 2, 291,
	292, 7, 3, 2, 2, 292, 293, 7, 26, 2, 2, 293, 294, 5, 8, 5, 2, 294, 295,
	7, 5, 2, 2, 295, 308, 3, 2, 2, 2, 296, 308, 5, 20, 11, 2, 297, 298, 7,
	27, 2, 2, 298, 308, 5, 10, 6, 2, 299, 300, 7, 28, 2, 2, 300, 308, 5, 12,
	7, 2, 301, 302, 7, 29, 2, 2, 302, 308, 7, 30, 2, 2, 303, 308, 7, 42, 2,
	2, 304, 308, 7, 44, 2, 2, 305, 308, 7, 45, 2, 2, 306, 308, 7, 46, 2, 2,
	307, 67, 3, 2, 2, 2, 307, 76, 3, 2, 2, 2, 307, 85, 3, 2, 2, 2, 307, 91,
	3, 2, 2, 2, 307, 96, 3, 2, 2, 2, 307, 109, 3, 2, 2, 2, 307, 122, 3, 2,
	2, 2, 307, 135, 3, 2, 2, 2, 307, 152, 3, 2, 2, 2, 307, 169, 3, 2, 2, 2,
	307, 186, 3, 2, 2, 2, 307, 203, 3, 2, 2, 2, 307, 221, 3, 2, 2, 2, 307,
	231, 3, 2, 2, 2, 307, 244, 3, 2, 2, 2, 307, 262, 3, 2, 2, 2, 307, 269,
	3, 2, 2, 2, 307, 280, 3, 2, 2, 2, 307, 291, 3, 2, 2, 2, 307, 296, 3, 2,
	2, 2, 307, 297, 3, 2, 2, 2, 307, 299, 3, 2, 2, 2, 307, 301, 3, 2, 2, 2,
	307, 303, 3, 2, 2, 2, 307, 304, 3, 2, 2, 2, 307, 305, 3, 2, 2, 2, 307,
	306, 3, 2, 2, 2, 308, 9, 3, 2, 2, 2, 309, 329, 7, 41, 2, 2, 310, 329, 7,
	45, 2, 2, 311, 329, 7, 46, 2, 2, 312, 316, 7, 3, 2, 2, 313, 315, 5, 10,
	6, 2, 314, 313, 3, 2, 2, 2, 315, 318, 3, 2, 2, 2, 316, 314, 3, 2, 2, 2,
	316, 317, 3, 2, 2, 2, 317, 319, 3, 2, 2, 2, 318, 316, 3, 2, 2, 2, 319,
	329, 7, 5, 2, 2, 320, 321, 7, 27, 2, 2, 321, 329, 5, 10, 6, 2, 322, 323,
	7, 28, 2, 2, 323, 329, 5, 10, 6, 2, 324, 325, 7, 31, 2, 2, 325, 329, 5,
	10, 6, 2, 326, 327, 7, 32, 2, 2, 327, 329, 5, 10, 6, 2, 328, 309, 3, 2,
	2, 2, 328, 310, 3, 2, 2, 2, 328, 311, 3, 2, 2, 2, 328, 312, 3, 2, 2, 2,
	328, 320, 3, 2, 2, 2, 328, 322, 3, 2, 2, 2, 328, 324, 3, 2, 2, 2, 328,
	326, 3, 2, 2, 2, 329, 11, 3, 2, 2, 2, 330, 351, 7, 41, 2, 2, 331, 351,
	7, 42, 2, 2, 332, 351, 7, 45, 2, 2, 333, 351, 7, 46, 2, 2, 334, 338, 7,
	3, 2, 2, 335, 337, 5, 12, 7, 2, 336, 335, 3, 2, 2, 2, 337, 340, 3, 2, 2,
	2, 338, 336, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 341, 3, 2, 2, 2, 340,
	338, 3, 2, 2, 2, 341, 351, 7, 5, 2, 2, 342, 343, 7, 27, 2, 2, 343, 351,
	5, 12, 7, 2, 344, 345, 7, 28, 2, 2, 345, 351, 5, 12, 7, 2, 346, 347, 7,
	31, 2, 2, 347, 351, 5, 8, 5, 2, 348, 349, 7, 32, 2, 2, 349, 351, 5, 8,
	5, 2, 350, 330, 3, 2, 2, 2, 350, 331, 3, 2, 2, 2, 350, 332, 3, 2, 2, 2,
	350, 333, 3, 2, 2, 2, 350, 334, 3, 2, 2, 2, 350, 342, 3, 2, 2, 2, 350,
	344, 3, 2, 2, 2, 350, 346, 3, 2, 2, 2, 350, 348, 3, 2, 2, 2, 351, 13, 3,
	2, 2, 2, 352, 353, 7, 3, 2, 2, 353, 354, 7, 33, 2, 2, 354, 355, 5, 8, 5,
	2, 355, 356, 5, 8, 5, 2, 356, 357, 7, 5, 2, 2, 357, 396, 3, 2, 2, 2, 358,
	359, 7, 3, 2, 2, 359, 360, 7, 34, 2, 2, 360, 361, 5, 8, 5, 2, 361, 362,
	5, 8, 5, 2, 362, 363, 7, 5, 2, 2, 363, 396, 3, 2, 2, 2, 364, 365, 7, 3,
	2, 2, 365, 366, 7, 35, 2, 2, 366, 367, 5, 8, 5, 2, 367, 368, 5, 8, 5, 2,
	368, 369, 5, 8, 5, 2, 369, 370, 7, 5, 2, 2, 370, 396, 3, 2, 2, 2, 371,
	372, 7, 3, 2, 2, 372, 373, 7, 36, 2, 2, 373, 375, 5, 8, 5, 2, 374, 376,
	5, 8, 5, 2, 375, 374, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 375, 3, 2,
	2, 2, 377, 378, 3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 380, 7, 5, 2, 2,
	380, 396, 3, 2, 2, 2, 381, 382, 7, 3, 2, 2, 382, 383, 7, 37, 2, 2, 383,
	384, 5, 8, 5, 2, 384, 385, 5, 20, 11, 2, 385, 386, 7, 5, 2, 2, 386, 396,
	3, 2, 2, 2, 387, 388, 7, 3, 2, 2, 388, 389, 7, 38, 2, 2, 389, 391, 5, 8,
	5, 2, 390, 392, 5, 8, 5, 2, 391, 390, 3, 2, 2, 2, 391, 392, 3, 2, 2, 2,
	392, 393, 3, 2, 2, 2, 393, 394, 7, 5, 2, 2, 394, 396, 3, 2, 2, 2, 395,
	352, 3, 2, 2, 2, 395, 358, 3, 2, 2, 2, 395, 364, 3, 2, 2, 2, 395, 371,
	3, 2, 2, 2, 395, 381, 3, 2, 2, 2, 395, 387, 3, 2, 2, 2, 396, 15, 3, 2,
	2, 2, 397, 398, 7, 3, 2, 2, 398, 399, 7, 39, 2, 2, 399, 400, 7, 45, 2,
	2, 400, 433, 7, 5, 2, 2, 401, 402, 7, 3, 2, 2, 402, 403, 7, 39, 2, 2, 403,
	404, 5, 20, 11, 2, 404, 405, 7, 5, 2, 2, 405, 433, 3, 2, 2, 2, 406, 407,
	7, 3, 2, 2, 407, 408, 7, 39, 2, 2, 408, 409, 7, 3, 2, 2, 409, 410, 5, 20,
	11, 2, 410, 418, 7, 45, 2, 2, 411, 413, 7, 3, 2, 2, 412, 414, 7, 45, 2,
	2, 413, 412, 3, 2, 2, 2, 414, 415, 3, 2, 2, 2, 415, 413, 3, 2, 2, 2, 415,
	416, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 419, 7, 5, 2, 2, 418, 411,
	3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420, 421, 7, 5,
	2, 2, 421, 422, 7, 5, 2, 2, 422, 433, 3, 2, 2, 2, 423, 424, 7, 3, 2, 2,
	424, 425, 7, 39, 2, 2, 425, 426, 7, 3, 2, 2, 426, 427, 5, 20, 11, 2, 427,
	428, 7, 45, 2, 2, 428, 429, 5, 18, 10, 2, 429, 430, 7, 5, 2, 2, 430, 431,
	7, 5, 2, 2, 431, 433, 3, 2, 2, 2, 432, 397, 3, 2, 2, 2, 432, 401, 3, 2,
	2, 2, 432, 406, 3, 2, 2, 2, 432, 423, 3, 2, 2, 2, 433, 17, 3, 2, 2, 2,
	434, 435, 7, 3, 2, 2, 435, 436, 7, 45, 2, 2, 436, 437, 7, 45, 2, 2, 437,
	438, 7, 42, 2, 2, 438, 439, 7, 42, 2, 2, 439, 440, 7, 5, 2, 2, 440, 19,
	3, 2, 2, 2, 441, 442, 9, 2, 2, 2, 442, 21, 3, 2, 2, 2, 33, 25, 31, 40,
	59, 65, 72, 81, 102, 115, 128, 145, 162, 179, 196, 214, 226, 240, 253,
	276, 287, 307, 316, 328, 338, 350, 377, 391, 395, 415, 418, 432,
}
var literalNames = []string{
	"", "'('", "'define'", "')'", "'define-struct'", "'begin'", "'begin0'",
	"'set!'", "'delay'", "'lambda'", "'\u03BB'", "'local'", "'['", "']'", "'letrec'",
	"'shared'", "'let'", "'let*'", "'recur'", "'cond'", "'else '", "'if'",
	"'and'", "'or'", "'time'", "'\u2019'", "'\u2018'", "'\u0027'", "'()'",
	"','", "',@'", "'check-expect'", "'check-random'", "'check-within'", "'check-member-of'",
	"'check-satisfied'", "'check-error'", "'require'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "SYMBOL", "NAME", "NUMBER", "INT", "BOOLEAN", "STRING", "CHARACTER",
	"LANG", "COMMENT", "WS",
}

var ruleNames = []string{
	"program", "defOrExpr", "definition", "expr", "quoted", "quasiQuoted",
	"testCase", "libraryRequire", "pkg", "name",
}

type ISLParser struct {
	*antlr.BaseParser
}

// NewISLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ISLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewISLParser(input antlr.TokenStream) *ISLParser {
	this := new(ISLParser)
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
	this.GrammarFileName = "ISL.g4"

	return this
}

// ISLParser tokens.
const (
	ISLParserEOF       = antlr.TokenEOF
	ISLParserT__0      = 1
	ISLParserT__1      = 2
	ISLParserT__2      = 3
	ISLParserT__3      = 4
	ISLParserT__4      = 5
	ISLParserT__5      = 6
	ISLParserT__6      = 7
	ISLParserT__7      = 8
	ISLParserT__8      = 9
	ISLParserT__9      = 10
	ISLParserT__10     = 11
	ISLParserT__11     = 12
	ISLParserT__12     = 13
	ISLParserT__13     = 14
	ISLParserT__14     = 15
	ISLParserT__15     = 16
	ISLParserT__16     = 17
	ISLParserT__17     = 18
	ISLParserT__18     = 19
	ISLParserT__19     = 20
	ISLParserT__20     = 21
	ISLParserT__21     = 22
	ISLParserT__22     = 23
	ISLParserT__23     = 24
	ISLParserT__24     = 25
	ISLParserT__25     = 26
	ISLParserT__26     = 27
	ISLParserT__27     = 28
	ISLParserT__28     = 29
	ISLParserT__29     = 30
	ISLParserT__30     = 31
	ISLParserT__31     = 32
	ISLParserT__32     = 33
	ISLParserT__33     = 34
	ISLParserT__34     = 35
	ISLParserT__35     = 36
	ISLParserT__36     = 37
	ISLParserSYMBOL    = 38
	ISLParserNAME      = 39
	ISLParserNUMBER    = 40
	ISLParserINT       = 41
	ISLParserBOOLEAN   = 42
	ISLParserSTRING    = 43
	ISLParserCHARACTER = 44
	ISLParserLANG      = 45
	ISLParserCOMMENT   = 46
	ISLParserWS        = 47
)

// ISLParser rules.
const (
	ISLParserRULE_program        = 0
	ISLParserRULE_defOrExpr      = 1
	ISLParserRULE_definition     = 2
	ISLParserRULE_expr           = 3
	ISLParserRULE_quoted         = 4
	ISLParserRULE_quasiQuoted    = 5
	ISLParserRULE_testCase       = 6
	ISLParserRULE_libraryRequire = 7
	ISLParserRULE_pkg            = 8
	ISLParserRULE_name           = 9
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
	p.RuleIndex = ISLParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ISLParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllDefOrExpr() []IDefOrExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefOrExprContext)(nil)).Elem())
	var tst = make([]IDefOrExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefOrExprContext)
		}
	}

	return tst
}

func (s *ProgramContext) DefOrExpr(i int) IDefOrExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefOrExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefOrExprContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *ISLParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ISLParserRULE_program)
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
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ISLParserT__0)|(1<<ISLParserT__24)|(1<<ISLParserT__25)|(1<<ISLParserT__26))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ISLParserSYMBOL-38))|(1<<(ISLParserNAME-38))|(1<<(ISLParserNUMBER-38))|(1<<(ISLParserBOOLEAN-38))|(1<<(ISLParserSTRING-38))|(1<<(ISLParserCHARACTER-38)))) != 0) {
		{
			p.SetState(20)
			p.DefOrExpr()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDefOrExprContext is an interface to support dynamic dispatch.
type IDefOrExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefOrExprContext differentiates from other interfaces.
	IsDefOrExprContext()
}

type DefOrExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefOrExprContext() *DefOrExprContext {
	var p = new(DefOrExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ISLParserRULE_defOrExpr
	return p
}

func (*DefOrExprContext) IsDefOrExprContext() {}

func NewDefOrExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefOrExprContext {
	var p = new(DefOrExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ISLParserRULE_defOrExpr

	return p
}

func (s *DefOrExprContext) GetParser() antlr.Parser { return s.parser }

func (s *DefOrExprContext) Definition() IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DefOrExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DefOrExprContext) TestCase() ITestCaseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITestCaseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITestCaseContext)
}

func (s *DefOrExprContext) LibraryRequire() ILibraryRequireContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILibraryRequireContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILibraryRequireContext)
}

func (s *DefOrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefOrExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefOrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.EnterDefOrExpr(s)
	}
}

func (s *DefOrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.ExitDefOrExpr(s)
	}
}

func (p *ISLParser) DefOrExpr() (localctx IDefOrExprContext) {
	this := p
	_ = this

	localctx = NewDefOrExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ISLParserRULE_defOrExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(25)
			p.Definition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)
			p.Expr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(27)
			p.TestCase()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(28)
			p.LibraryRequire()
		}

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
	p.RuleIndex = ISLParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ISLParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *DefinitionContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DefinitionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DefinitionContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(ISLParserNAME)
}

func (s *DefinitionContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(ISLParserNAME, i)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (p *ISLParser) Definition() (localctx IDefinitionContext) {
	this := p
	_ = this

	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ISLParserRULE_definition)
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

	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(31)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(32)
			p.Match(ISLParserT__1)
		}
		{
			p.SetState(33)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(34)
			p.Name()
		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ISLParserNAME {
			{
				p.SetState(35)
				p.Match(ISLParserNAME)
			}

			p.SetState(38)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(40)
			p.Match(ISLParserT__2)
		}
		{
			p.SetState(41)
			p.Expr()
		}
		{
			p.SetState(42)
			p.Match(ISLParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(45)
			p.Match(ISLParserT__1)
		}
		{
			p.SetState(46)
			p.Name()
		}
		{
			p.SetState(47)
			p.Expr()
		}
		{
			p.SetState(48)
			p.Match(ISLParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(50)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(51)
			p.Match(ISLParserT__3)
		}
		{
			p.SetState(52)
			p.Name()
		}
		{
			p.SetState(53)
			p.Match(ISLParserT__0)
		}
		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ISLParserSYMBOL || _la == ISLParserNAME {
			{
				p.SetState(54)
				p.Name()
			}

			p.SetState(59)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(60)
			p.Match(ISLParserT__2)
		}
		{
			p.SetState(61)
			p.Match(ISLParserT__2)
		}

	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ISLParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ISLParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) AllNAME() []antlr.TerminalNode {
	return s.GetTokens(ISLParserNAME)
}

func (s *ExprContext) NAME(i int) antlr.TerminalNode {
	return s.GetToken(ISLParserNAME, i)
}

func (s *ExprContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *ExprContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *ExprContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *ExprContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ExprContext) Quoted() IQuotedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuotedContext)
}

func (s *ExprContext) QuasiQuoted() IQuasiQuotedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuasiQuotedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuasiQuotedContext)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ISLParserNUMBER, 0)
}

func (s *ExprContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(ISLParserBOOLEAN, 0)
}

func (s *ExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(ISLParserSTRING, 0)
}

func (s *ExprContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(ISLParserCHARACTER, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *ISLParser) Expr() (localctx IExprContext) {
	this := p
	_ = this

	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ISLParserRULE_expr)
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

	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(66)
			p.Match(ISLParserT__4)
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ISLParserT__0)|(1<<ISLParserT__24)|(1<<ISLParserT__25)|(1<<ISLParserT__26))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ISLParserSYMBOL-38))|(1<<(ISLParserNAME-38))|(1<<(ISLParserNUMBER-38))|(1<<(ISLParserBOOLEAN-38))|(1<<(ISLParserSTRING-38))|(1<<(ISLParserCHARACTER-38)))) != 0) {
			{
				p.SetState(67)
				p.Expr()
			}

			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(72)
			p.Match(ISLParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(74)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(75)
			p.Match(ISLParserT__5)
		}
		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ISLParserT__0)|(1<<ISLParserT__24)|(1<<ISLParserT__25)|(1<<ISLParserT__26))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ISLParserSYMBOL-38))|(1<<(ISLParserNAME-38))|(1<<(ISLParserNUMBER-38))|(1<<(ISLParserBOOLEAN-38))|(1<<(ISLParserSTRING-38))|(1<<(ISLParserCHARACTER-38)))) != 0) {
			{
				p.SetState(76)
				p.Expr()
			}

			p.SetState(79)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(81)
			p.Match(ISLParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(83)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(84)
			p.Match(ISLParserT__6)
		}
		{
			p.SetState(85)
			p.Match(ISLParserNAME)
		}
		{
			p.SetState(86)
			p.Expr()
		}
		{
			p.SetState(87)
			p.Match(ISLParserT__2)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(90)
			p.Match(ISLParserT__7)
		}
		{
			p.SetState(91)
			p.Expr()
		}
		{
			p.SetState(92)
			p.Match(ISLParserT__2)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(94)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(95)
			p.Match(ISLParserT__8)
		}
		{
			p.SetState(96)
			p.Match(ISLParserT__0)
		}
		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ISLParserNAME {
			{
				p.SetState(97)
				p.Match(ISLParserNAME)
			}

			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(103)
			p.Match(ISLParserT__2)
		}
		{
			p.SetState(104)
			p.Expr()
		}
		{
			p.SetState(105)
			p.Match(ISLParserT__2)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(107)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(108)
			p.Match(ISLParserT__9)
		}
		{
			p.SetState(109)
			p.Match(ISLParserT__0)
		}
		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ISLParserNAME {
			{
				p.SetState(110)
				p.Match(ISLParserNAME)
			}

			p.SetState(115)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(116)
			p.Match(ISLParserT__2)
		}
		{
			p.SetState(117)
			p.Expr()
		}
		{
			p.SetState(118)
			p.Match(ISLParserT__2)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(120)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(121)
			p.Match(ISLParserT__10)
		}
		{
			p.SetState(122)
			p.Match(ISLParserT__11)
		}
		p.SetState(126)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ISLParserT__0 {
			{
				p.SetState(123)
				p.Definition()
			}

			p.SetState(128)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(129)
			p.Match(ISLParserT__12)
		}
		{
			p.SetState(130)
			p.Expr()
		}
		{
			p.SetState(131)
			p.Match(ISLParserT__2)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(133)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(134)
			p.Match(ISLParserT__13)
		}
		{
			p.SetState(135)
			p.Match(ISLParserT__0)
		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ISLParserT__11 {
			{
				p.SetState(136)
				p.Match(ISLParserT__11)
			}
			{
				p.SetState(137)
				p.Name()
			}
			{
				p.SetState(138)
				p.Expr()
			}
			{
				p.SetState(139)
				p.Match(ISLParserT__12)
			}

			p.SetState(145)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(146)
			p.Match(ISLParserT__2)
		}
		{
			p.SetState(147)
			p.Expr()
		}
		{
			p.SetState(148)
			p.Match(ISLParserT__2)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(150)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(151)
			p.Match(ISLParserT__14)
		}
		{
			p.SetState(152)
			p.Match(ISLParserT__0)
		}
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ISLParserT__11 {
			{
				p.SetState(153)
				p.Match(ISLParserT__11)
			}
			{
				p.SetState(154)
				p.Name()
			}
			{
				p.SetState(155)
				p.Expr()
			}
			{
				p.SetState(156)
				p.Match(ISLParserT__12)
			}

			p.SetState(162)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(163)
			p.Match(ISLParserT__2)
		}
		{
			p.SetState(164)
			p.Expr()
		}
		{
			p.SetState(165)
			p.Match(ISLParserT__2)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(167)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(168)
			p.Match(ISLParserT__15)
		}
		{
			p.SetState(169)
			p.Match(ISLParserT__0)
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ISLParserT__11 {
			{
				p.SetState(170)
				p.Match(ISLParserT__11)
			}
			{
				p.SetState(171)
				p.Name()
			}
			{
				p.SetState(172)
				p.Expr()
			}
			{
				p.SetState(173)
				p.Match(ISLParserT__12)
			}

			p.SetState(179)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(180)
			p.Match(ISLParserT__2)
		}
		{
			p.SetState(181)
			p.Expr()
		}
		{
			p.SetState(182)
			p.Match(ISLParserT__2)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(184)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(185)
			p.Match(ISLParserT__16)
		}
		{
			p.SetState(186)
			p.Match(ISLParserT__0)
		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ISLParserT__11 {
			{
				p.SetState(187)
				p.Match(ISLParserT__11)
			}
			{
				p.SetState(188)
				p.Name()
			}
			{
				p.SetState(189)
				p.Expr()
			}
			{
				p.SetState(190)
				p.Match(ISLParserT__12)
			}

			p.SetState(196)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(197)
			p.Match(ISLParserT__2)
		}
		{
			p.SetState(198)
			p.Expr()
		}
		{
			p.SetState(199)
			p.Match(ISLParserT__2)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(201)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(202)
			p.Match(ISLParserT__17)
		}
		{
			p.SetState(203)
			p.Name()
		}
		{
			p.SetState(204)
			p.Match(ISLParserT__0)
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == ISLParserT__11 {
			{
				p.SetState(205)
				p.Match(ISLParserT__11)
			}
			{
				p.SetState(206)
				p.Name()
			}
			{
				p.SetState(207)
				p.Expr()
			}
			{
				p.SetState(208)
				p.Match(ISLParserT__12)
			}

			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(215)
			p.Match(ISLParserT__2)
		}
		{
			p.SetState(216)
			p.Expr()
		}
		{
			p.SetState(217)
			p.Match(ISLParserT__2)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(219)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(220)
			p.Expr()
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ISLParserT__0)|(1<<ISLParserT__24)|(1<<ISLParserT__25)|(1<<ISLParserT__26))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ISLParserSYMBOL-38))|(1<<(ISLParserNAME-38))|(1<<(ISLParserNUMBER-38))|(1<<(ISLParserBOOLEAN-38))|(1<<(ISLParserSTRING-38))|(1<<(ISLParserCHARACTER-38)))) != 0) {
			{
				p.SetState(221)
				p.Expr()
			}

			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(227)
			p.Match(ISLParserT__2)
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(229)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(230)
			p.Match(ISLParserT__18)
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ISLParserT__11 {
			{
				p.SetState(231)
				p.Match(ISLParserT__11)
			}
			{
				p.SetState(232)
				p.Expr()
			}
			{
				p.SetState(233)
				p.Expr()
			}
			{
				p.SetState(234)
				p.Match(ISLParserT__12)
			}

			p.SetState(238)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(240)
			p.Match(ISLParserT__2)
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(242)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(243)
			p.Match(ISLParserT__18)
		}
		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(244)
					p.Match(ISLParserT__11)
				}
				{
					p.SetState(245)
					p.Expr()
				}
				{
					p.SetState(246)
					p.Expr()
				}
				{
					p.SetState(247)
					p.Match(ISLParserT__12)
				}

			}
			p.SetState(253)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
		}
		{
			p.SetState(254)
			p.Match(ISLParserT__11)
		}
		{
			p.SetState(255)
			p.Match(ISLParserT__19)
		}
		{
			p.SetState(256)
			p.Expr()
		}
		{
			p.SetState(257)
			p.Match(ISLParserT__12)
		}
		{
			p.SetState(258)
			p.Match(ISLParserT__2)
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(260)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(261)
			p.Match(ISLParserT__20)
		}
		{
			p.SetState(262)
			p.Expr()
		}
		{
			p.SetState(263)
			p.Expr()
		}
		{
			p.SetState(264)
			p.Expr()
		}
		{
			p.SetState(265)
			p.Match(ISLParserT__2)
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(267)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(268)
			p.Match(ISLParserT__21)
		}
		{
			p.SetState(269)
			p.Expr()
		}
		{
			p.SetState(270)
			p.Expr()
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ISLParserT__0)|(1<<ISLParserT__24)|(1<<ISLParserT__25)|(1<<ISLParserT__26))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ISLParserSYMBOL-38))|(1<<(ISLParserNAME-38))|(1<<(ISLParserNUMBER-38))|(1<<(ISLParserBOOLEAN-38))|(1<<(ISLParserSTRING-38))|(1<<(ISLParserCHARACTER-38)))) != 0) {
			{
				p.SetState(271)
				p.Expr()
			}

			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(276)
			p.Match(ISLParserT__2)
		}

	case 18:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(278)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(279)
			p.Match(ISLParserT__22)
		}
		{
			p.SetState(280)
			p.Expr()
		}
		{
			p.SetState(281)
			p.Expr()
		}
		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ISLParserT__0)|(1<<ISLParserT__24)|(1<<ISLParserT__25)|(1<<ISLParserT__26))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ISLParserSYMBOL-38))|(1<<(ISLParserNAME-38))|(1<<(ISLParserNUMBER-38))|(1<<(ISLParserBOOLEAN-38))|(1<<(ISLParserSTRING-38))|(1<<(ISLParserCHARACTER-38)))) != 0) {
			{
				p.SetState(282)
				p.Expr()
			}

			p.SetState(285)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(287)
			p.Match(ISLParserT__2)
		}

	case 19:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(289)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(290)
			p.Match(ISLParserT__23)
		}
		{
			p.SetState(291)
			p.Expr()
		}
		{
			p.SetState(292)
			p.Match(ISLParserT__2)
		}

	case 20:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(294)
			p.Name()
		}

	case 21:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(295)
			p.Match(ISLParserT__24)
		}
		{
			p.SetState(296)
			p.Quoted()
		}

	case 22:
		p.EnterOuterAlt(localctx, 22)
		{
			p.SetState(297)
			p.Match(ISLParserT__25)
		}
		{
			p.SetState(298)
			p.QuasiQuoted()
		}

	case 23:
		p.EnterOuterAlt(localctx, 23)
		{
			p.SetState(299)
			p.Match(ISLParserT__26)
		}
		{
			p.SetState(300)
			p.Match(ISLParserT__27)
		}

	case 24:
		p.EnterOuterAlt(localctx, 24)
		{
			p.SetState(301)
			p.Match(ISLParserNUMBER)
		}

	case 25:
		p.EnterOuterAlt(localctx, 25)
		{
			p.SetState(302)
			p.Match(ISLParserBOOLEAN)
		}

	case 26:
		p.EnterOuterAlt(localctx, 26)
		{
			p.SetState(303)
			p.Match(ISLParserSTRING)
		}

	case 27:
		p.EnterOuterAlt(localctx, 27)
		{
			p.SetState(304)
			p.Match(ISLParserCHARACTER)
		}

	}

	return localctx
}

// IQuotedContext is an interface to support dynamic dispatch.
type IQuotedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuotedContext differentiates from other interfaces.
	IsQuotedContext()
}

type QuotedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedContext() *QuotedContext {
	var p = new(QuotedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ISLParserRULE_quoted
	return p
}

func (*QuotedContext) IsQuotedContext() {}

func NewQuotedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedContext {
	var p = new(QuotedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ISLParserRULE_quoted

	return p
}

func (s *QuotedContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedContext) NAME() antlr.TerminalNode {
	return s.GetToken(ISLParserNAME, 0)
}

func (s *QuotedContext) STRING() antlr.TerminalNode {
	return s.GetToken(ISLParserSTRING, 0)
}

func (s *QuotedContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(ISLParserCHARACTER, 0)
}

func (s *QuotedContext) AllQuoted() []IQuotedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuotedContext)(nil)).Elem())
	var tst = make([]IQuotedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuotedContext)
		}
	}

	return tst
}

func (s *QuotedContext) Quoted(i int) IQuotedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuotedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuotedContext)
}

func (s *QuotedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.EnterQuoted(s)
	}
}

func (s *QuotedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.ExitQuoted(s)
	}
}

func (p *ISLParser) Quoted() (localctx IQuotedContext) {
	this := p
	_ = this

	localctx = NewQuotedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ISLParserRULE_quoted)
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

	p.SetState(326)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ISLParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(307)
			p.Match(ISLParserNAME)
		}

	case ISLParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(308)
			p.Match(ISLParserSTRING)
		}

	case ISLParserCHARACTER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(309)
			p.Match(ISLParserCHARACTER)
		}

	case ISLParserT__0:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(310)
			p.Match(ISLParserT__0)
		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ISLParserT__0)|(1<<ISLParserT__24)|(1<<ISLParserT__25)|(1<<ISLParserT__28)|(1<<ISLParserT__29))) != 0) || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(ISLParserNAME-39))|(1<<(ISLParserSTRING-39))|(1<<(ISLParserCHARACTER-39)))) != 0) {
			{
				p.SetState(311)
				p.Quoted()
			}

			p.SetState(316)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(317)
			p.Match(ISLParserT__2)
		}

	case ISLParserT__24:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(318)
			p.Match(ISLParserT__24)
		}
		{
			p.SetState(319)
			p.Quoted()
		}

	case ISLParserT__25:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(320)
			p.Match(ISLParserT__25)
		}
		{
			p.SetState(321)
			p.Quoted()
		}

	case ISLParserT__28:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(322)
			p.Match(ISLParserT__28)
		}
		{
			p.SetState(323)
			p.Quoted()
		}

	case ISLParserT__29:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(324)
			p.Match(ISLParserT__29)
		}
		{
			p.SetState(325)
			p.Quoted()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQuasiQuotedContext is an interface to support dynamic dispatch.
type IQuasiQuotedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuasiQuotedContext differentiates from other interfaces.
	IsQuasiQuotedContext()
}

type QuasiQuotedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuasiQuotedContext() *QuasiQuotedContext {
	var p = new(QuasiQuotedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ISLParserRULE_quasiQuoted
	return p
}

func (*QuasiQuotedContext) IsQuasiQuotedContext() {}

func NewQuasiQuotedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuasiQuotedContext {
	var p = new(QuasiQuotedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ISLParserRULE_quasiQuoted

	return p
}

func (s *QuasiQuotedContext) GetParser() antlr.Parser { return s.parser }

func (s *QuasiQuotedContext) NAME() antlr.TerminalNode {
	return s.GetToken(ISLParserNAME, 0)
}

func (s *QuasiQuotedContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ISLParserNUMBER, 0)
}

func (s *QuasiQuotedContext) STRING() antlr.TerminalNode {
	return s.GetToken(ISLParserSTRING, 0)
}

func (s *QuasiQuotedContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(ISLParserCHARACTER, 0)
}

func (s *QuasiQuotedContext) AllQuasiQuoted() []IQuasiQuotedContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQuasiQuotedContext)(nil)).Elem())
	var tst = make([]IQuasiQuotedContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQuasiQuotedContext)
		}
	}

	return tst
}

func (s *QuasiQuotedContext) QuasiQuoted(i int) IQuasiQuotedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuasiQuotedContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQuasiQuotedContext)
}

func (s *QuasiQuotedContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *QuasiQuotedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuasiQuotedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuasiQuotedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.EnterQuasiQuoted(s)
	}
}

func (s *QuasiQuotedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.ExitQuasiQuoted(s)
	}
}

func (p *ISLParser) QuasiQuoted() (localctx IQuasiQuotedContext) {
	this := p
	_ = this

	localctx = NewQuasiQuotedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ISLParserRULE_quasiQuoted)
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

	p.SetState(348)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ISLParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(328)
			p.Match(ISLParserNAME)
		}

	case ISLParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.Match(ISLParserNUMBER)
		}

	case ISLParserSTRING:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(330)
			p.Match(ISLParserSTRING)
		}

	case ISLParserCHARACTER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(331)
			p.Match(ISLParserCHARACTER)
		}

	case ISLParserT__0:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(332)
			p.Match(ISLParserT__0)
		}
		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ISLParserT__0)|(1<<ISLParserT__24)|(1<<ISLParserT__25)|(1<<ISLParserT__28)|(1<<ISLParserT__29))) != 0) || (((_la-39)&-(0x1f+1)) == 0 && ((1<<uint((_la-39)))&((1<<(ISLParserNAME-39))|(1<<(ISLParserNUMBER-39))|(1<<(ISLParserSTRING-39))|(1<<(ISLParserCHARACTER-39)))) != 0) {
			{
				p.SetState(333)
				p.QuasiQuoted()
			}

			p.SetState(338)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(339)
			p.Match(ISLParserT__2)
		}

	case ISLParserT__24:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(340)
			p.Match(ISLParserT__24)
		}
		{
			p.SetState(341)
			p.QuasiQuoted()
		}

	case ISLParserT__25:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(342)
			p.Match(ISLParserT__25)
		}
		{
			p.SetState(343)
			p.QuasiQuoted()
		}

	case ISLParserT__28:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(344)
			p.Match(ISLParserT__28)
		}
		{
			p.SetState(345)
			p.Expr()
		}

	case ISLParserT__29:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(346)
			p.Match(ISLParserT__29)
		}
		{
			p.SetState(347)
			p.Expr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITestCaseContext is an interface to support dynamic dispatch.
type ITestCaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTestCaseContext differentiates from other interfaces.
	IsTestCaseContext()
}

type TestCaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTestCaseContext() *TestCaseContext {
	var p = new(TestCaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ISLParserRULE_testCase
	return p
}

func (*TestCaseContext) IsTestCaseContext() {}

func NewTestCaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TestCaseContext {
	var p = new(TestCaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ISLParserRULE_testCase

	return p
}

func (s *TestCaseContext) GetParser() antlr.Parser { return s.parser }

func (s *TestCaseContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *TestCaseContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TestCaseContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *TestCaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TestCaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TestCaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.EnterTestCase(s)
	}
}

func (s *TestCaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.ExitTestCase(s)
	}
}

func (p *ISLParser) TestCase() (localctx ITestCaseContext) {
	this := p
	_ = this

	localctx = NewTestCaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ISLParserRULE_testCase)
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(350)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(351)
			p.Match(ISLParserT__30)
		}
		{
			p.SetState(352)
			p.Expr()
		}
		{
			p.SetState(353)
			p.Expr()
		}
		{
			p.SetState(354)
			p.Match(ISLParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(356)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(357)
			p.Match(ISLParserT__31)
		}
		{
			p.SetState(358)
			p.Expr()
		}
		{
			p.SetState(359)
			p.Expr()
		}
		{
			p.SetState(360)
			p.Match(ISLParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(362)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(363)
			p.Match(ISLParserT__32)
		}
		{
			p.SetState(364)
			p.Expr()
		}
		{
			p.SetState(365)
			p.Expr()
		}
		{
			p.SetState(366)
			p.Expr()
		}
		{
			p.SetState(367)
			p.Match(ISLParserT__2)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(369)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(370)
			p.Match(ISLParserT__33)
		}
		{
			p.SetState(371)
			p.Expr()
		}
		p.SetState(373)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ISLParserT__0)|(1<<ISLParserT__24)|(1<<ISLParserT__25)|(1<<ISLParserT__26))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ISLParserSYMBOL-38))|(1<<(ISLParserNAME-38))|(1<<(ISLParserNUMBER-38))|(1<<(ISLParserBOOLEAN-38))|(1<<(ISLParserSTRING-38))|(1<<(ISLParserCHARACTER-38)))) != 0) {
			{
				p.SetState(372)
				p.Expr()
			}

			p.SetState(375)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(377)
			p.Match(ISLParserT__2)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(379)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(380)
			p.Match(ISLParserT__34)
		}
		{
			p.SetState(381)
			p.Expr()
		}
		{
			p.SetState(382)
			p.Name()
		}
		{
			p.SetState(383)
			p.Match(ISLParserT__2)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(385)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(386)
			p.Match(ISLParserT__35)
		}
		{
			p.SetState(387)
			p.Expr()
		}
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ISLParserT__0)|(1<<ISLParserT__24)|(1<<ISLParserT__25)|(1<<ISLParserT__26))) != 0) || (((_la-38)&-(0x1f+1)) == 0 && ((1<<uint((_la-38)))&((1<<(ISLParserSYMBOL-38))|(1<<(ISLParserNAME-38))|(1<<(ISLParserNUMBER-38))|(1<<(ISLParserBOOLEAN-38))|(1<<(ISLParserSTRING-38))|(1<<(ISLParserCHARACTER-38)))) != 0) {
			{
				p.SetState(388)
				p.Expr()
			}

		}
		{
			p.SetState(391)
			p.Match(ISLParserT__2)
		}

	}

	return localctx
}

// ILibraryRequireContext is an interface to support dynamic dispatch.
type ILibraryRequireContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLibraryRequireContext differentiates from other interfaces.
	IsLibraryRequireContext()
}

type LibraryRequireContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLibraryRequireContext() *LibraryRequireContext {
	var p = new(LibraryRequireContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ISLParserRULE_libraryRequire
	return p
}

func (*LibraryRequireContext) IsLibraryRequireContext() {}

func NewLibraryRequireContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LibraryRequireContext {
	var p = new(LibraryRequireContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ISLParserRULE_libraryRequire

	return p
}

func (s *LibraryRequireContext) GetParser() antlr.Parser { return s.parser }

func (s *LibraryRequireContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ISLParserSTRING)
}

func (s *LibraryRequireContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ISLParserSTRING, i)
}

func (s *LibraryRequireContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LibraryRequireContext) Pkg() IPkgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPkgContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPkgContext)
}

func (s *LibraryRequireContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LibraryRequireContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LibraryRequireContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.EnterLibraryRequire(s)
	}
}

func (s *LibraryRequireContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.ExitLibraryRequire(s)
	}
}

func (p *ISLParser) LibraryRequire() (localctx ILibraryRequireContext) {
	this := p
	_ = this

	localctx = NewLibraryRequireContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ISLParserRULE_libraryRequire)
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

	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(395)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(396)
			p.Match(ISLParserT__36)
		}
		{
			p.SetState(397)
			p.Match(ISLParserSTRING)
		}
		{
			p.SetState(398)
			p.Match(ISLParserT__2)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(399)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(400)
			p.Match(ISLParserT__36)
		}
		{
			p.SetState(401)
			p.Name()
		}
		{
			p.SetState(402)
			p.Match(ISLParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(404)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(405)
			p.Match(ISLParserT__36)
		}
		{
			p.SetState(406)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(407)
			p.Name()
		}
		{
			p.SetState(408)
			p.Match(ISLParserSTRING)
		}
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == ISLParserT__0 {
			{
				p.SetState(409)
				p.Match(ISLParserT__0)
			}
			p.SetState(411)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == ISLParserSTRING {
				{
					p.SetState(410)
					p.Match(ISLParserSTRING)
				}

				p.SetState(413)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(415)
				p.Match(ISLParserT__2)
			}

		}
		{
			p.SetState(418)
			p.Match(ISLParserT__2)
		}
		{
			p.SetState(419)
			p.Match(ISLParserT__2)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(421)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(422)
			p.Match(ISLParserT__36)
		}
		{
			p.SetState(423)
			p.Match(ISLParserT__0)
		}
		{
			p.SetState(424)
			p.Name()
		}
		{
			p.SetState(425)
			p.Match(ISLParserSTRING)
		}
		{
			p.SetState(426)
			p.Pkg()
		}
		{
			p.SetState(427)
			p.Match(ISLParserT__2)
		}
		{
			p.SetState(428)
			p.Match(ISLParserT__2)
		}

	}

	return localctx
}

// IPkgContext is an interface to support dynamic dispatch.
type IPkgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPkgContext differentiates from other interfaces.
	IsPkgContext()
}

type PkgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPkgContext() *PkgContext {
	var p = new(PkgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ISLParserRULE_pkg
	return p
}

func (*PkgContext) IsPkgContext() {}

func NewPkgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PkgContext {
	var p = new(PkgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ISLParserRULE_pkg

	return p
}

func (s *PkgContext) GetParser() antlr.Parser { return s.parser }

func (s *PkgContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(ISLParserSTRING)
}

func (s *PkgContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(ISLParserSTRING, i)
}

func (s *PkgContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(ISLParserNUMBER)
}

func (s *PkgContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(ISLParserNUMBER, i)
}

func (s *PkgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PkgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PkgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.EnterPkg(s)
	}
}

func (s *PkgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.ExitPkg(s)
	}
}

func (p *ISLParser) Pkg() (localctx IPkgContext) {
	this := p
	_ = this

	localctx = NewPkgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ISLParserRULE_pkg)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.Match(ISLParserT__0)
	}
	{
		p.SetState(433)
		p.Match(ISLParserSTRING)
	}
	{
		p.SetState(434)
		p.Match(ISLParserSTRING)
	}
	{
		p.SetState(435)
		p.Match(ISLParserNUMBER)
	}
	{
		p.SetState(436)
		p.Match(ISLParserNUMBER)
	}
	{
		p.SetState(437)
		p.Match(ISLParserT__2)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ISLParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ISLParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) SYMBOL() antlr.TerminalNode {
	return s.GetToken(ISLParserSYMBOL, 0)
}

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(ISLParserNAME, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ISLListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *ISLParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ISLParserRULE_name)
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
		p.SetState(439)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ISLParserSYMBOL || _la == ISLParserNAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
