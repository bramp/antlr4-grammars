// Code generated from wkt.g4 by ANTLR 4.9.3. DO NOT EDIT.

package wkt // wkt
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 26, 308,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 7, 2, 48, 10, 2, 12, 2, 14, 2, 51, 11, 2, 3, 2, 3, 2, 3, 2, 5, 2, 56,
	10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 5, 3, 72, 10, 3, 3, 4, 3, 4, 5, 4, 76, 10, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 5, 4, 83, 10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 96, 10, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 5, 7, 102, 10, 7, 7, 7, 104, 10, 7, 12, 7, 14, 7, 107, 11, 7, 3, 7,
	3, 7, 3, 7, 5, 7, 112, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 118, 10, 8,
	3, 8, 3, 8, 3, 8, 5, 8, 123, 10, 8, 7, 8, 125, 10, 8, 12, 8, 14, 8, 128,
	11, 8, 3, 8, 3, 8, 3, 8, 5, 8, 133, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	5, 9, 140, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 146, 10, 9, 7, 9, 148,
	10, 9, 12, 9, 14, 9, 151, 11, 9, 3, 9, 3, 9, 3, 9, 5, 9, 156, 10, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 5, 10, 162, 10, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	167, 10, 10, 7, 10, 169, 10, 10, 12, 10, 14, 10, 172, 11, 10, 3, 10, 3,
	10, 3, 10, 5, 10, 177, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11,
	184, 10, 11, 12, 11, 14, 11, 187, 11, 11, 3, 11, 3, 11, 3, 11, 5, 11, 192,
	10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 199, 10, 12, 12, 12,
	14, 12, 202, 11, 12, 3, 12, 3, 12, 3, 12, 5, 12, 207, 10, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 7, 13, 214, 10, 13, 12, 13, 14, 13, 217, 11, 13,
	3, 13, 3, 13, 3, 13, 5, 13, 222, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 7, 14, 229, 10, 14, 12, 14, 14, 14, 232, 11, 14, 3, 14, 3, 14, 3, 14,
	5, 14, 237, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 7, 15, 244, 10,
	15, 12, 15, 14, 15, 247, 11, 15, 3, 15, 3, 15, 3, 15, 5, 15, 252, 10, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 259, 10, 16, 12, 16, 14, 16,
	262, 11, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 271,
	10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 277, 10, 18, 12, 18, 14, 18,
	280, 11, 18, 3, 18, 3, 18, 3, 18, 5, 18, 285, 10, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 7, 19, 291, 10, 19, 12, 19, 14, 19, 294, 11, 19, 3, 19, 3, 19,
	3, 19, 5, 19, 299, 10, 19, 3, 20, 6, 20, 302, 10, 20, 13, 20, 14, 20, 303,
	3, 21, 3, 21, 3, 21, 2, 2, 22, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 2, 2, 2, 341, 2, 42, 3, 2, 2, 2, 4,
	71, 3, 2, 2, 2, 6, 73, 3, 2, 2, 2, 8, 84, 3, 2, 2, 2, 10, 87, 3, 2, 2,
	2, 12, 90, 3, 2, 2, 2, 14, 113, 3, 2, 2, 2, 16, 134, 3, 2, 2, 2, 18, 157,
	3, 2, 2, 2, 20, 178, 3, 2, 2, 2, 22, 193, 3, 2, 2, 2, 24, 208, 3, 2, 2,
	2, 26, 223, 3, 2, 2, 2, 28, 238, 3, 2, 2, 2, 30, 253, 3, 2, 2, 2, 32, 270,
	3, 2, 2, 2, 34, 284, 3, 2, 2, 2, 36, 298, 3, 2, 2, 2, 38, 301, 3, 2, 2,
	2, 40, 305, 3, 2, 2, 2, 42, 55, 7, 15, 2, 2, 43, 44, 7, 7, 2, 2, 44, 49,
	5, 4, 3, 2, 45, 46, 7, 6, 2, 2, 46, 48, 5, 4, 3, 2, 47, 45, 3, 2, 2, 2,
	48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50, 3, 2, 2, 2, 50, 52, 3,
	2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 53, 7, 8, 2, 2, 53, 56, 3, 2, 2, 2, 54,
	56, 7, 16, 2, 2, 55, 43, 3, 2, 2, 2, 55, 54, 3, 2, 2, 2, 56, 3, 3, 2, 2,
	2, 57, 72, 5, 10, 6, 2, 58, 72, 5, 8, 5, 2, 59, 72, 5, 6, 4, 2, 60, 72,
	5, 18, 10, 2, 61, 72, 5, 16, 9, 2, 62, 72, 5, 14, 8, 2, 63, 72, 5, 12,
	7, 2, 64, 72, 5, 20, 11, 2, 65, 72, 5, 22, 12, 2, 66, 72, 5, 24, 13, 2,
	67, 72, 5, 30, 16, 2, 68, 72, 5, 26, 14, 2, 69, 72, 5, 28, 15, 2, 70, 72,
	5, 2, 2, 2, 71, 57, 3, 2, 2, 2, 71, 58, 3, 2, 2, 2, 71, 59, 3, 2, 2, 2,
	71, 60, 3, 2, 2, 2, 71, 61, 3, 2, 2, 2, 71, 62, 3, 2, 2, 2, 71, 63, 3,
	2, 2, 2, 71, 64, 3, 2, 2, 2, 71, 65, 3, 2, 2, 2, 71, 66, 3, 2, 2, 2, 71,
	67, 3, 2, 2, 2, 71, 68, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 71, 70, 3, 2, 2,
	2, 72, 5, 3, 2, 2, 2, 73, 82, 7, 9, 2, 2, 74, 76, 5, 40, 21, 2, 75, 74,
	3, 2, 2, 2, 75, 76, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 7, 7, 2, 2,
	78, 79, 5, 38, 20, 2, 79, 80, 7, 8, 2, 2, 80, 83, 3, 2, 2, 2, 81, 83, 7,
	16, 2, 2, 82, 75, 3, 2, 2, 2, 82, 81, 3, 2, 2, 2, 83, 7, 3, 2, 2, 2, 84,
	85, 7, 10, 2, 2, 85, 86, 5, 36, 19, 2, 86, 9, 3, 2, 2, 2, 87, 88, 7, 11,
	2, 2, 88, 89, 5, 34, 18, 2, 89, 11, 3, 2, 2, 2, 90, 111, 7, 21, 2, 2, 91,
	95, 7, 7, 2, 2, 92, 96, 5, 36, 19, 2, 93, 96, 5, 30, 16, 2, 94, 96, 5,
	18, 10, 2, 95, 92, 3, 2, 2, 2, 95, 93, 3, 2, 2, 2, 95, 94, 3, 2, 2, 2,
	96, 105, 3, 2, 2, 2, 97, 101, 7, 6, 2, 2, 98, 102, 5, 30, 16, 2, 99, 102,
	5, 36, 19, 2, 100, 102, 5, 18, 10, 2, 101, 98, 3, 2, 2, 2, 101, 99, 3,
	2, 2, 2, 101, 100, 3, 2, 2, 2, 102, 104, 3, 2, 2, 2, 103, 97, 3, 2, 2,
	2, 104, 107, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106,
	108, 3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 108, 109, 7, 8, 2, 2, 109, 112,
	3, 2, 2, 2, 110, 112, 7, 16, 2, 2, 111, 91, 3, 2, 2, 2, 111, 110, 3, 2,
	2, 2, 112, 13, 3, 2, 2, 2, 113, 132, 7, 19, 2, 2, 114, 117, 7, 7, 2, 2,
	115, 118, 5, 34, 18, 2, 116, 118, 5, 16, 9, 2, 117, 115, 3, 2, 2, 2, 117,
	116, 3, 2, 2, 2, 118, 126, 3, 2, 2, 2, 119, 122, 7, 6, 2, 2, 120, 123,
	5, 34, 18, 2, 121, 123, 5, 16, 9, 2, 122, 120, 3, 2, 2, 2, 122, 121, 3,
	2, 2, 2, 123, 125, 3, 2, 2, 2, 124, 119, 3, 2, 2, 2, 125, 128, 3, 2, 2,
	2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 129, 3, 2, 2, 2, 128,
	126, 3, 2, 2, 2, 129, 130, 7, 8, 2, 2, 130, 133, 3, 2, 2, 2, 131, 133,
	7, 16, 2, 2, 132, 114, 3, 2, 2, 2, 132, 131, 3, 2, 2, 2, 133, 15, 3, 2,
	2, 2, 134, 155, 7, 20, 2, 2, 135, 139, 7, 7, 2, 2, 136, 140, 5, 36, 19,
	2, 137, 140, 5, 30, 16, 2, 138, 140, 5, 18, 10, 2, 139, 136, 3, 2, 2, 2,
	139, 137, 3, 2, 2, 2, 139, 138, 3, 2, 2, 2, 140, 149, 3, 2, 2, 2, 141,
	145, 7, 6, 2, 2, 142, 146, 5, 30, 16, 2, 143, 146, 5, 36, 19, 2, 144, 146,
	5, 18, 10, 2, 145, 142, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 145, 144, 3,
	2, 2, 2, 146, 148, 3, 2, 2, 2, 147, 141, 3, 2, 2, 2, 148, 151, 3, 2, 2,
	2, 149, 147, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 152, 3, 2, 2, 2, 151,
	149, 3, 2, 2, 2, 152, 153, 7, 8, 2, 2, 153, 156, 3, 2, 2, 2, 154, 156,
	7, 16, 2, 2, 155, 135, 3, 2, 2, 2, 155, 154, 3, 2, 2, 2, 156, 17, 3, 2,
	2, 2, 157, 176, 7, 18, 2, 2, 158, 161, 7, 7, 2, 2, 159, 162, 5, 36, 19,
	2, 160, 162, 5, 30, 16, 2, 161, 159, 3, 2, 2, 2, 161, 160, 3, 2, 2, 2,
	162, 170, 3, 2, 2, 2, 163, 166, 7, 6, 2, 2, 164, 167, 5, 30, 16, 2, 165,
	167, 5, 36, 19, 2, 166, 164, 3, 2, 2, 2, 166, 165, 3, 2, 2, 2, 167, 169,
	3, 2, 2, 2, 168, 163, 3, 2, 2, 2, 169, 172, 3, 2, 2, 2, 170, 168, 3, 2,
	2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2, 172, 170, 3, 2, 2, 2,
	173, 174, 7, 8, 2, 2, 174, 177, 3, 2, 2, 2, 175, 177, 7, 16, 2, 2, 176,
	158, 3, 2, 2, 2, 176, 175, 3, 2, 2, 2, 177, 19, 3, 2, 2, 2, 178, 191, 7,
	12, 2, 2, 179, 180, 7, 7, 2, 2, 180, 185, 5, 32, 17, 2, 181, 182, 7, 6,
	2, 2, 182, 184, 5, 32, 17, 2, 183, 181, 3, 2, 2, 2, 184, 187, 3, 2, 2,
	2, 185, 183, 3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 188, 3, 2, 2, 2, 187,
	185, 3, 2, 2, 2, 188, 189, 7, 8, 2, 2, 189, 192, 3, 2, 2, 2, 190, 192,
	7, 16, 2, 2, 191, 179, 3, 2, 2, 2, 191, 190, 3, 2, 2, 2, 192, 21, 3, 2,
	2, 2, 193, 206, 7, 13, 2, 2, 194, 195, 7, 7, 2, 2, 195, 200, 5, 36, 19,
	2, 196, 197, 7, 6, 2, 2, 197, 199, 5, 36, 19, 2, 198, 196, 3, 2, 2, 2,
	199, 202, 3, 2, 2, 2, 200, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201,
	203, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 203, 204, 7, 8, 2, 2, 204, 207,
	3, 2, 2, 2, 205, 207, 7, 16, 2, 2, 206, 194, 3, 2, 2, 2, 206, 205, 3, 2,
	2, 2, 207, 23, 3, 2, 2, 2, 208, 221, 7, 14, 2, 2, 209, 210, 7, 7, 2, 2,
	210, 215, 5, 34, 18, 2, 211, 212, 7, 6, 2, 2, 212, 214, 5, 34, 18, 2, 213,
	211, 3, 2, 2, 2, 214, 217, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 215, 216,
	3, 2, 2, 2, 216, 218, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 218, 219, 7, 8,
	2, 2, 219, 222, 3, 2, 2, 2, 220, 222, 7, 16, 2, 2, 221, 209, 3, 2, 2, 2,
	221, 220, 3, 2, 2, 2, 222, 25, 3, 2, 2, 2, 223, 236, 7, 24, 2, 2, 224,
	225, 7, 7, 2, 2, 225, 230, 5, 34, 18, 2, 226, 227, 7, 6, 2, 2, 227, 229,
	5, 34, 18, 2, 228, 226, 3, 2, 2, 2, 229, 232, 3, 2, 2, 2, 230, 228, 3,
	2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 233, 3, 2, 2, 2, 232, 230, 3, 2, 2,
	2, 233, 234, 7, 8, 2, 2, 234, 237, 3, 2, 2, 2, 235, 237, 7, 16, 2, 2, 236,
	224, 3, 2, 2, 2, 236, 235, 3, 2, 2, 2, 237, 27, 3, 2, 2, 2, 238, 251, 7,
	23, 2, 2, 239, 240, 7, 7, 2, 2, 240, 245, 5, 34, 18, 2, 241, 242, 7, 6,
	2, 2, 242, 244, 5, 34, 18, 2, 243, 241, 3, 2, 2, 2, 244, 247, 3, 2, 2,
	2, 245, 243, 3, 2, 2, 2, 245, 246, 3, 2, 2, 2, 246, 248, 3, 2, 2, 2, 247,
	245, 3, 2, 2, 2, 248, 249, 7, 8, 2, 2, 249, 252, 3, 2, 2, 2, 250, 252,
	7, 16, 2, 2, 251, 239, 3, 2, 2, 2, 251, 250, 3, 2, 2, 2, 252, 29, 3, 2,
	2, 2, 253, 254, 7, 17, 2, 2, 254, 255, 7, 7, 2, 2, 255, 260, 5, 38, 20,
	2, 256, 257, 7, 6, 2, 2, 257, 259, 5, 38, 20, 2, 258, 256, 3, 2, 2, 2,
	259, 262, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 260, 261, 3, 2, 2, 2, 261,
	263, 3, 2, 2, 2, 262, 260, 3, 2, 2, 2, 263, 264, 7, 8, 2, 2, 264, 31, 3,
	2, 2, 2, 265, 271, 5, 38, 20, 2, 266, 267, 7, 7, 2, 2, 267, 268, 5, 38,
	20, 2, 268, 269, 7, 8, 2, 2, 269, 271, 3, 2, 2, 2, 270, 265, 3, 2, 2, 2,
	270, 266, 3, 2, 2, 2, 271, 33, 3, 2, 2, 2, 272, 273, 7, 7, 2, 2, 273, 278,
	5, 36, 19, 2, 274, 275, 7, 6, 2, 2, 275, 277, 5, 36, 19, 2, 276, 274, 3,
	2, 2, 2, 277, 280, 3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 278, 279, 3, 2, 2,
	2, 279, 281, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2, 281, 282, 7, 8, 2, 2, 282,
	285, 3, 2, 2, 2, 283, 285, 7, 16, 2, 2, 284, 272, 3, 2, 2, 2, 284, 283,
	3, 2, 2, 2, 285, 35, 3, 2, 2, 2, 286, 287, 7, 7, 2, 2, 287, 292, 5, 38,
	20, 2, 288, 289, 7, 6, 2, 2, 289, 291, 5, 38, 20, 2, 290, 288, 3, 2, 2,
	2, 291, 294, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 292, 293, 3, 2, 2, 2, 293,
	295, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2, 295, 296, 7, 8, 2, 2, 296, 299,
	3, 2, 2, 2, 297, 299, 7, 16, 2, 2, 298, 286, 3, 2, 2, 2, 298, 297, 3, 2,
	2, 2, 299, 37, 3, 2, 2, 2, 300, 302, 7, 3, 2, 2, 301, 300, 3, 2, 2, 2,
	302, 303, 3, 2, 2, 2, 303, 301, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304,
	39, 3, 2, 2, 2, 305, 306, 7, 25, 2, 2, 306, 41, 3, 2, 2, 2, 40, 49, 55,
	71, 75, 82, 95, 101, 105, 111, 117, 122, 126, 132, 139, 145, 149, 155,
	161, 166, 170, 176, 185, 191, 200, 206, 215, 221, 230, 236, 245, 251, 260,
	270, 278, 284, 292, 298, 303,
}
var literalNames = []string{
	"", "", "", "", "','", "'('", "')'",
}
var symbolicNames = []string{
	"", "DECIMAL", "INTEGERPART", "DECIMALPART", "COMMA", "LPAR", "RPAR", "POINT",
	"LINESTRING", "POLYGON", "MULTIPOINT", "MULTILINESTRING", "MULTIPOLYGON",
	"GEOMETRYCOLLECTION", "EMPTY", "CIRCULARSTRING", "COMPOUNDCURVE", "MULTISURFACE",
	"CURVEPOLYGON", "MULTICURVE", "TRIANGLE", "TIN", "POLYHEDRALSURFACE", "STRING",
	"WS",
}

var ruleNames = []string{
	"geometryCollection", "geometry", "pointGeometry", "lineStringGeometry",
	"polygonGeometry", "multiCurveGeometry", "multiSurfaceGeometry", "curvePolygonGeometry",
	"compoundCurveGeometry", "multiPointGeometry", "multiLineStringGeometry",
	"multiPolygonGeometry", "multiPolyhedralSurfaceGeometry", "multiTinGeometry",
	"circularStringGeometry", "pointOrClosedPoint", "polygon", "lineString",
	"point", "name",
}

type wktParser struct {
	*antlr.BaseParser
}

// NewwktParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *wktParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewwktParser(input antlr.TokenStream) *wktParser {
	this := new(wktParser)
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
	this.GrammarFileName = "wkt.g4"

	return this
}

// wktParser tokens.
const (
	wktParserEOF                = antlr.TokenEOF
	wktParserDECIMAL            = 1
	wktParserINTEGERPART        = 2
	wktParserDECIMALPART        = 3
	wktParserCOMMA              = 4
	wktParserLPAR               = 5
	wktParserRPAR               = 6
	wktParserPOINT              = 7
	wktParserLINESTRING         = 8
	wktParserPOLYGON            = 9
	wktParserMULTIPOINT         = 10
	wktParserMULTILINESTRING    = 11
	wktParserMULTIPOLYGON       = 12
	wktParserGEOMETRYCOLLECTION = 13
	wktParserEMPTY              = 14
	wktParserCIRCULARSTRING     = 15
	wktParserCOMPOUNDCURVE      = 16
	wktParserMULTISURFACE       = 17
	wktParserCURVEPOLYGON       = 18
	wktParserMULTICURVE         = 19
	wktParserTRIANGLE           = 20
	wktParserTIN                = 21
	wktParserPOLYHEDRALSURFACE  = 22
	wktParserSTRING             = 23
	wktParserWS                 = 24
)

// wktParser rules.
const (
	wktParserRULE_geometryCollection             = 0
	wktParserRULE_geometry                       = 1
	wktParserRULE_pointGeometry                  = 2
	wktParserRULE_lineStringGeometry             = 3
	wktParserRULE_polygonGeometry                = 4
	wktParserRULE_multiCurveGeometry             = 5
	wktParserRULE_multiSurfaceGeometry           = 6
	wktParserRULE_curvePolygonGeometry           = 7
	wktParserRULE_compoundCurveGeometry          = 8
	wktParserRULE_multiPointGeometry             = 9
	wktParserRULE_multiLineStringGeometry        = 10
	wktParserRULE_multiPolygonGeometry           = 11
	wktParserRULE_multiPolyhedralSurfaceGeometry = 12
	wktParserRULE_multiTinGeometry               = 13
	wktParserRULE_circularStringGeometry         = 14
	wktParserRULE_pointOrClosedPoint             = 15
	wktParserRULE_polygon                        = 16
	wktParserRULE_lineString                     = 17
	wktParserRULE_point                          = 18
	wktParserRULE_name                           = 19
)

// IGeometryCollectionContext is an interface to support dynamic dispatch.
type IGeometryCollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeometryCollectionContext differentiates from other interfaces.
	IsGeometryCollectionContext()
}

type GeometryCollectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeometryCollectionContext() *GeometryCollectionContext {
	var p = new(GeometryCollectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_geometryCollection
	return p
}

func (*GeometryCollectionContext) IsGeometryCollectionContext() {}

func NewGeometryCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeometryCollectionContext {
	var p = new(GeometryCollectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_geometryCollection

	return p
}

func (s *GeometryCollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *GeometryCollectionContext) GEOMETRYCOLLECTION() antlr.TerminalNode {
	return s.GetToken(wktParserGEOMETRYCOLLECTION, 0)
}

func (s *GeometryCollectionContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *GeometryCollectionContext) AllGeometry() []IGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGeometryContext)(nil)).Elem())
	var tst = make([]IGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGeometryContext)
		}
	}

	return tst
}

func (s *GeometryCollectionContext) Geometry(i int) IGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGeometryContext)
}

func (s *GeometryCollectionContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *GeometryCollectionContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *GeometryCollectionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *GeometryCollectionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *GeometryCollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeometryCollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeometryCollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterGeometryCollection(s)
	}
}

func (s *GeometryCollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitGeometryCollection(s)
	}
}

func (p *wktParser) GeometryCollection() (localctx IGeometryCollectionContext) {
	this := p
	_ = this

	localctx = NewGeometryCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, wktParserRULE_geometryCollection)
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
		p.SetState(40)
		p.Match(wktParserGEOMETRYCOLLECTION)
	}
	p.SetState(53)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		{
			p.SetState(41)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(42)
			p.Geometry()
		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(43)
				p.Match(wktParserCOMMA)
			}
			{
				p.SetState(44)
				p.Geometry()
			}

			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(50)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(52)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGeometryContext is an interface to support dynamic dispatch.
type IGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGeometryContext differentiates from other interfaces.
	IsGeometryContext()
}

type GeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGeometryContext() *GeometryContext {
	var p = new(GeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_geometry
	return p
}

func (*GeometryContext) IsGeometryContext() {}

func NewGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GeometryContext {
	var p = new(GeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_geometry

	return p
}

func (s *GeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *GeometryContext) PolygonGeometry() IPolygonGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPolygonGeometryContext)
}

func (s *GeometryContext) LineStringGeometry() ILineStringGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineStringGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineStringGeometryContext)
}

func (s *GeometryContext) PointGeometry() IPointGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointGeometryContext)
}

func (s *GeometryContext) CompoundCurveGeometry() ICompoundCurveGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundCurveGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompoundCurveGeometryContext)
}

func (s *GeometryContext) CurvePolygonGeometry() ICurvePolygonGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurvePolygonGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICurvePolygonGeometryContext)
}

func (s *GeometryContext) MultiSurfaceGeometry() IMultiSurfaceGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiSurfaceGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiSurfaceGeometryContext)
}

func (s *GeometryContext) MultiCurveGeometry() IMultiCurveGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiCurveGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiCurveGeometryContext)
}

func (s *GeometryContext) MultiPointGeometry() IMultiPointGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiPointGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiPointGeometryContext)
}

func (s *GeometryContext) MultiLineStringGeometry() IMultiLineStringGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiLineStringGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiLineStringGeometryContext)
}

func (s *GeometryContext) MultiPolygonGeometry() IMultiPolygonGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiPolygonGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiPolygonGeometryContext)
}

func (s *GeometryContext) CircularStringGeometry() ICircularStringGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICircularStringGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICircularStringGeometryContext)
}

func (s *GeometryContext) MultiPolyhedralSurfaceGeometry() IMultiPolyhedralSurfaceGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiPolyhedralSurfaceGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiPolyhedralSurfaceGeometryContext)
}

func (s *GeometryContext) MultiTinGeometry() IMultiTinGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultiTinGeometryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultiTinGeometryContext)
}

func (s *GeometryContext) GeometryCollection() IGeometryCollectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGeometryCollectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGeometryCollectionContext)
}

func (s *GeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterGeometry(s)
	}
}

func (s *GeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitGeometry(s)
	}
}

func (p *wktParser) Geometry() (localctx IGeometryContext) {
	this := p
	_ = this

	localctx = NewGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, wktParserRULE_geometry)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserPOLYGON:
		{
			p.SetState(55)
			p.PolygonGeometry()
		}

	case wktParserLINESTRING:
		{
			p.SetState(56)
			p.LineStringGeometry()
		}

	case wktParserPOINT:
		{
			p.SetState(57)
			p.PointGeometry()
		}

	case wktParserCOMPOUNDCURVE:
		{
			p.SetState(58)
			p.CompoundCurveGeometry()
		}

	case wktParserCURVEPOLYGON:
		{
			p.SetState(59)
			p.CurvePolygonGeometry()
		}

	case wktParserMULTISURFACE:
		{
			p.SetState(60)
			p.MultiSurfaceGeometry()
		}

	case wktParserMULTICURVE:
		{
			p.SetState(61)
			p.MultiCurveGeometry()
		}

	case wktParserMULTIPOINT:
		{
			p.SetState(62)
			p.MultiPointGeometry()
		}

	case wktParserMULTILINESTRING:
		{
			p.SetState(63)
			p.MultiLineStringGeometry()
		}

	case wktParserMULTIPOLYGON:
		{
			p.SetState(64)
			p.MultiPolygonGeometry()
		}

	case wktParserCIRCULARSTRING:
		{
			p.SetState(65)
			p.CircularStringGeometry()
		}

	case wktParserPOLYHEDRALSURFACE:
		{
			p.SetState(66)
			p.MultiPolyhedralSurfaceGeometry()
		}

	case wktParserTIN:
		{
			p.SetState(67)
			p.MultiTinGeometry()
		}

	case wktParserGEOMETRYCOLLECTION:
		{
			p.SetState(68)
			p.GeometryCollection()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPointGeometryContext is an interface to support dynamic dispatch.
type IPointGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointGeometryContext differentiates from other interfaces.
	IsPointGeometryContext()
}

type PointGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointGeometryContext() *PointGeometryContext {
	var p = new(PointGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_pointGeometry
	return p
}

func (*PointGeometryContext) IsPointGeometryContext() {}

func NewPointGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointGeometryContext {
	var p = new(PointGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_pointGeometry

	return p
}

func (s *PointGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *PointGeometryContext) POINT() antlr.TerminalNode {
	return s.GetToken(wktParserPOINT, 0)
}

func (s *PointGeometryContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *PointGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *PointGeometryContext) Point() IPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointContext)
}

func (s *PointGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *PointGeometryContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *PointGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterPointGeometry(s)
	}
}

func (s *PointGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitPointGeometry(s)
	}
}

func (p *wktParser) PointGeometry() (localctx IPointGeometryContext) {
	this := p
	_ = this

	localctx = NewPointGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, wktParserRULE_pointGeometry)
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
		p.SetState(71)
		p.Match(wktParserPOINT)
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR, wktParserSTRING:
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == wktParserSTRING {
			{
				p.SetState(72)
				p.Name()
			}

		}
		{
			p.SetState(75)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(76)
			p.Point()
		}
		{
			p.SetState(77)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(79)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILineStringGeometryContext is an interface to support dynamic dispatch.
type ILineStringGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineStringGeometryContext differentiates from other interfaces.
	IsLineStringGeometryContext()
}

type LineStringGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineStringGeometryContext() *LineStringGeometryContext {
	var p = new(LineStringGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_lineStringGeometry
	return p
}

func (*LineStringGeometryContext) IsLineStringGeometryContext() {}

func NewLineStringGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineStringGeometryContext {
	var p = new(LineStringGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_lineStringGeometry

	return p
}

func (s *LineStringGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *LineStringGeometryContext) LINESTRING() antlr.TerminalNode {
	return s.GetToken(wktParserLINESTRING, 0)
}

func (s *LineStringGeometryContext) LineString() ILineStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineStringContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILineStringContext)
}

func (s *LineStringGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineStringGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineStringGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterLineStringGeometry(s)
	}
}

func (s *LineStringGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitLineStringGeometry(s)
	}
}

func (p *wktParser) LineStringGeometry() (localctx ILineStringGeometryContext) {
	this := p
	_ = this

	localctx = NewLineStringGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, wktParserRULE_lineStringGeometry)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
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
		p.Match(wktParserLINESTRING)
	}
	{
		p.SetState(83)
		p.LineString()
	}

	return localctx
}

// IPolygonGeometryContext is an interface to support dynamic dispatch.
type IPolygonGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPolygonGeometryContext differentiates from other interfaces.
	IsPolygonGeometryContext()
}

type PolygonGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPolygonGeometryContext() *PolygonGeometryContext {
	var p = new(PolygonGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_polygonGeometry
	return p
}

func (*PolygonGeometryContext) IsPolygonGeometryContext() {}

func NewPolygonGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolygonGeometryContext {
	var p = new(PolygonGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_polygonGeometry

	return p
}

func (s *PolygonGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *PolygonGeometryContext) POLYGON() antlr.TerminalNode {
	return s.GetToken(wktParserPOLYGON, 0)
}

func (s *PolygonGeometryContext) Polygon() IPolygonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPolygonContext)
}

func (s *PolygonGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolygonGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolygonGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterPolygonGeometry(s)
	}
}

func (s *PolygonGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitPolygonGeometry(s)
	}
}

func (p *wktParser) PolygonGeometry() (localctx IPolygonGeometryContext) {
	this := p
	_ = this

	localctx = NewPolygonGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, wktParserRULE_polygonGeometry)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.Match(wktParserPOLYGON)
	}
	{
		p.SetState(86)
		p.Polygon()
	}

	return localctx
}

// IMultiCurveGeometryContext is an interface to support dynamic dispatch.
type IMultiCurveGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiCurveGeometryContext differentiates from other interfaces.
	IsMultiCurveGeometryContext()
}

type MultiCurveGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiCurveGeometryContext() *MultiCurveGeometryContext {
	var p = new(MultiCurveGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_multiCurveGeometry
	return p
}

func (*MultiCurveGeometryContext) IsMultiCurveGeometryContext() {}

func NewMultiCurveGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiCurveGeometryContext {
	var p = new(MultiCurveGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_multiCurveGeometry

	return p
}

func (s *MultiCurveGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiCurveGeometryContext) MULTICURVE() antlr.TerminalNode {
	return s.GetToken(wktParserMULTICURVE, 0)
}

func (s *MultiCurveGeometryContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *MultiCurveGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *MultiCurveGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *MultiCurveGeometryContext) AllLineString() []ILineStringContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineStringContext)(nil)).Elem())
	var tst = make([]ILineStringContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineStringContext)
		}
	}

	return tst
}

func (s *MultiCurveGeometryContext) LineString(i int) ILineStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineStringContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineStringContext)
}

func (s *MultiCurveGeometryContext) AllCircularStringGeometry() []ICircularStringGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICircularStringGeometryContext)(nil)).Elem())
	var tst = make([]ICircularStringGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICircularStringGeometryContext)
		}
	}

	return tst
}

func (s *MultiCurveGeometryContext) CircularStringGeometry(i int) ICircularStringGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICircularStringGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICircularStringGeometryContext)
}

func (s *MultiCurveGeometryContext) AllCompoundCurveGeometry() []ICompoundCurveGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICompoundCurveGeometryContext)(nil)).Elem())
	var tst = make([]ICompoundCurveGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICompoundCurveGeometryContext)
		}
	}

	return tst
}

func (s *MultiCurveGeometryContext) CompoundCurveGeometry(i int) ICompoundCurveGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundCurveGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICompoundCurveGeometryContext)
}

func (s *MultiCurveGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *MultiCurveGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *MultiCurveGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiCurveGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiCurveGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterMultiCurveGeometry(s)
	}
}

func (s *MultiCurveGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitMultiCurveGeometry(s)
	}
}

func (p *wktParser) MultiCurveGeometry() (localctx IMultiCurveGeometryContext) {
	this := p
	_ = this

	localctx = NewMultiCurveGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, wktParserRULE_multiCurveGeometry)
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
		p.SetState(88)
		p.Match(wktParserMULTICURVE)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		{
			p.SetState(89)
			p.Match(wktParserLPAR)
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case wktParserLPAR, wktParserEMPTY:
			{
				p.SetState(90)
				p.LineString()
			}

		case wktParserCIRCULARSTRING:
			{
				p.SetState(91)
				p.CircularStringGeometry()
			}

		case wktParserCOMPOUNDCURVE:
			{
				p.SetState(92)
				p.CompoundCurveGeometry()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(95)
				p.Match(wktParserCOMMA)
			}
			p.SetState(99)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case wktParserCIRCULARSTRING:
				{
					p.SetState(96)
					p.CircularStringGeometry()
				}

			case wktParserLPAR, wktParserEMPTY:
				{
					p.SetState(97)
					p.LineString()
				}

			case wktParserCOMPOUNDCURVE:
				{
					p.SetState(98)
					p.CompoundCurveGeometry()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(105)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(106)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(108)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMultiSurfaceGeometryContext is an interface to support dynamic dispatch.
type IMultiSurfaceGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiSurfaceGeometryContext differentiates from other interfaces.
	IsMultiSurfaceGeometryContext()
}

type MultiSurfaceGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiSurfaceGeometryContext() *MultiSurfaceGeometryContext {
	var p = new(MultiSurfaceGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_multiSurfaceGeometry
	return p
}

func (*MultiSurfaceGeometryContext) IsMultiSurfaceGeometryContext() {}

func NewMultiSurfaceGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiSurfaceGeometryContext {
	var p = new(MultiSurfaceGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_multiSurfaceGeometry

	return p
}

func (s *MultiSurfaceGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiSurfaceGeometryContext) MULTISURFACE() antlr.TerminalNode {
	return s.GetToken(wktParserMULTISURFACE, 0)
}

func (s *MultiSurfaceGeometryContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *MultiSurfaceGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *MultiSurfaceGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *MultiSurfaceGeometryContext) AllPolygon() []IPolygonContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPolygonContext)(nil)).Elem())
	var tst = make([]IPolygonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPolygonContext)
		}
	}

	return tst
}

func (s *MultiSurfaceGeometryContext) Polygon(i int) IPolygonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPolygonContext)
}

func (s *MultiSurfaceGeometryContext) AllCurvePolygonGeometry() []ICurvePolygonGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICurvePolygonGeometryContext)(nil)).Elem())
	var tst = make([]ICurvePolygonGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICurvePolygonGeometryContext)
		}
	}

	return tst
}

func (s *MultiSurfaceGeometryContext) CurvePolygonGeometry(i int) ICurvePolygonGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICurvePolygonGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICurvePolygonGeometryContext)
}

func (s *MultiSurfaceGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *MultiSurfaceGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *MultiSurfaceGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiSurfaceGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiSurfaceGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterMultiSurfaceGeometry(s)
	}
}

func (s *MultiSurfaceGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitMultiSurfaceGeometry(s)
	}
}

func (p *wktParser) MultiSurfaceGeometry() (localctx IMultiSurfaceGeometryContext) {
	this := p
	_ = this

	localctx = NewMultiSurfaceGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, wktParserRULE_multiSurfaceGeometry)
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
		p.SetState(111)
		p.Match(wktParserMULTISURFACE)
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		{
			p.SetState(112)
			p.Match(wktParserLPAR)
		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case wktParserLPAR, wktParserEMPTY:
			{
				p.SetState(113)
				p.Polygon()
			}

		case wktParserCURVEPOLYGON:
			{
				p.SetState(114)
				p.CurvePolygonGeometry()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(117)
				p.Match(wktParserCOMMA)
			}
			p.SetState(120)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case wktParserLPAR, wktParserEMPTY:
				{
					p.SetState(118)
					p.Polygon()
				}

			case wktParserCURVEPOLYGON:
				{
					p.SetState(119)
					p.CurvePolygonGeometry()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(127)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(129)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICurvePolygonGeometryContext is an interface to support dynamic dispatch.
type ICurvePolygonGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCurvePolygonGeometryContext differentiates from other interfaces.
	IsCurvePolygonGeometryContext()
}

type CurvePolygonGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCurvePolygonGeometryContext() *CurvePolygonGeometryContext {
	var p = new(CurvePolygonGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_curvePolygonGeometry
	return p
}

func (*CurvePolygonGeometryContext) IsCurvePolygonGeometryContext() {}

func NewCurvePolygonGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CurvePolygonGeometryContext {
	var p = new(CurvePolygonGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_curvePolygonGeometry

	return p
}

func (s *CurvePolygonGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *CurvePolygonGeometryContext) CURVEPOLYGON() antlr.TerminalNode {
	return s.GetToken(wktParserCURVEPOLYGON, 0)
}

func (s *CurvePolygonGeometryContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *CurvePolygonGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *CurvePolygonGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *CurvePolygonGeometryContext) AllLineString() []ILineStringContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineStringContext)(nil)).Elem())
	var tst = make([]ILineStringContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineStringContext)
		}
	}

	return tst
}

func (s *CurvePolygonGeometryContext) LineString(i int) ILineStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineStringContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineStringContext)
}

func (s *CurvePolygonGeometryContext) AllCircularStringGeometry() []ICircularStringGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICircularStringGeometryContext)(nil)).Elem())
	var tst = make([]ICircularStringGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICircularStringGeometryContext)
		}
	}

	return tst
}

func (s *CurvePolygonGeometryContext) CircularStringGeometry(i int) ICircularStringGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICircularStringGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICircularStringGeometryContext)
}

func (s *CurvePolygonGeometryContext) AllCompoundCurveGeometry() []ICompoundCurveGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICompoundCurveGeometryContext)(nil)).Elem())
	var tst = make([]ICompoundCurveGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICompoundCurveGeometryContext)
		}
	}

	return tst
}

func (s *CurvePolygonGeometryContext) CompoundCurveGeometry(i int) ICompoundCurveGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompoundCurveGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICompoundCurveGeometryContext)
}

func (s *CurvePolygonGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *CurvePolygonGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *CurvePolygonGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CurvePolygonGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CurvePolygonGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterCurvePolygonGeometry(s)
	}
}

func (s *CurvePolygonGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitCurvePolygonGeometry(s)
	}
}

func (p *wktParser) CurvePolygonGeometry() (localctx ICurvePolygonGeometryContext) {
	this := p
	_ = this

	localctx = NewCurvePolygonGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, wktParserRULE_curvePolygonGeometry)
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
		p.SetState(132)
		p.Match(wktParserCURVEPOLYGON)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		{
			p.SetState(133)
			p.Match(wktParserLPAR)
		}
		p.SetState(137)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case wktParserLPAR, wktParserEMPTY:
			{
				p.SetState(134)
				p.LineString()
			}

		case wktParserCIRCULARSTRING:
			{
				p.SetState(135)
				p.CircularStringGeometry()
			}

		case wktParserCOMPOUNDCURVE:
			{
				p.SetState(136)
				p.CompoundCurveGeometry()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(139)
				p.Match(wktParserCOMMA)
			}
			p.SetState(143)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case wktParserCIRCULARSTRING:
				{
					p.SetState(140)
					p.CircularStringGeometry()
				}

			case wktParserLPAR, wktParserEMPTY:
				{
					p.SetState(141)
					p.LineString()
				}

			case wktParserCOMPOUNDCURVE:
				{
					p.SetState(142)
					p.CompoundCurveGeometry()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(149)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(150)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(152)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICompoundCurveGeometryContext is an interface to support dynamic dispatch.
type ICompoundCurveGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompoundCurveGeometryContext differentiates from other interfaces.
	IsCompoundCurveGeometryContext()
}

type CompoundCurveGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompoundCurveGeometryContext() *CompoundCurveGeometryContext {
	var p = new(CompoundCurveGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_compoundCurveGeometry
	return p
}

func (*CompoundCurveGeometryContext) IsCompoundCurveGeometryContext() {}

func NewCompoundCurveGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompoundCurveGeometryContext {
	var p = new(CompoundCurveGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_compoundCurveGeometry

	return p
}

func (s *CompoundCurveGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *CompoundCurveGeometryContext) COMPOUNDCURVE() antlr.TerminalNode {
	return s.GetToken(wktParserCOMPOUNDCURVE, 0)
}

func (s *CompoundCurveGeometryContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *CompoundCurveGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *CompoundCurveGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *CompoundCurveGeometryContext) AllLineString() []ILineStringContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineStringContext)(nil)).Elem())
	var tst = make([]ILineStringContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineStringContext)
		}
	}

	return tst
}

func (s *CompoundCurveGeometryContext) LineString(i int) ILineStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineStringContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineStringContext)
}

func (s *CompoundCurveGeometryContext) AllCircularStringGeometry() []ICircularStringGeometryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICircularStringGeometryContext)(nil)).Elem())
	var tst = make([]ICircularStringGeometryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICircularStringGeometryContext)
		}
	}

	return tst
}

func (s *CompoundCurveGeometryContext) CircularStringGeometry(i int) ICircularStringGeometryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICircularStringGeometryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICircularStringGeometryContext)
}

func (s *CompoundCurveGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *CompoundCurveGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *CompoundCurveGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompoundCurveGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompoundCurveGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterCompoundCurveGeometry(s)
	}
}

func (s *CompoundCurveGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitCompoundCurveGeometry(s)
	}
}

func (p *wktParser) CompoundCurveGeometry() (localctx ICompoundCurveGeometryContext) {
	this := p
	_ = this

	localctx = NewCompoundCurveGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, wktParserRULE_compoundCurveGeometry)
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
		p.SetState(155)
		p.Match(wktParserCOMPOUNDCURVE)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		{
			p.SetState(156)
			p.Match(wktParserLPAR)
		}
		p.SetState(159)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case wktParserLPAR, wktParserEMPTY:
			{
				p.SetState(157)
				p.LineString()
			}

		case wktParserCIRCULARSTRING:
			{
				p.SetState(158)
				p.CircularStringGeometry()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(161)
				p.Match(wktParserCOMMA)
			}
			p.SetState(164)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case wktParserCIRCULARSTRING:
				{
					p.SetState(162)
					p.CircularStringGeometry()
				}

			case wktParserLPAR, wktParserEMPTY:
				{
					p.SetState(163)
					p.LineString()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(171)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(173)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMultiPointGeometryContext is an interface to support dynamic dispatch.
type IMultiPointGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiPointGeometryContext differentiates from other interfaces.
	IsMultiPointGeometryContext()
}

type MultiPointGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiPointGeometryContext() *MultiPointGeometryContext {
	var p = new(MultiPointGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_multiPointGeometry
	return p
}

func (*MultiPointGeometryContext) IsMultiPointGeometryContext() {}

func NewMultiPointGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPointGeometryContext {
	var p = new(MultiPointGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_multiPointGeometry

	return p
}

func (s *MultiPointGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPointGeometryContext) MULTIPOINT() antlr.TerminalNode {
	return s.GetToken(wktParserMULTIPOINT, 0)
}

func (s *MultiPointGeometryContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *MultiPointGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *MultiPointGeometryContext) AllPointOrClosedPoint() []IPointOrClosedPointContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPointOrClosedPointContext)(nil)).Elem())
	var tst = make([]IPointOrClosedPointContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPointOrClosedPointContext)
		}
	}

	return tst
}

func (s *MultiPointGeometryContext) PointOrClosedPoint(i int) IPointOrClosedPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointOrClosedPointContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPointOrClosedPointContext)
}

func (s *MultiPointGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *MultiPointGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *MultiPointGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *MultiPointGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPointGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPointGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterMultiPointGeometry(s)
	}
}

func (s *MultiPointGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitMultiPointGeometry(s)
	}
}

func (p *wktParser) MultiPointGeometry() (localctx IMultiPointGeometryContext) {
	this := p
	_ = this

	localctx = NewMultiPointGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, wktParserRULE_multiPointGeometry)
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
		p.SetState(176)
		p.Match(wktParserMULTIPOINT)
	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		{
			p.SetState(177)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(178)
			p.PointOrClosedPoint()
		}
		p.SetState(183)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(179)
				p.Match(wktParserCOMMA)
			}
			{
				p.SetState(180)
				p.PointOrClosedPoint()
			}

			p.SetState(185)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(186)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(188)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMultiLineStringGeometryContext is an interface to support dynamic dispatch.
type IMultiLineStringGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiLineStringGeometryContext differentiates from other interfaces.
	IsMultiLineStringGeometryContext()
}

type MultiLineStringGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiLineStringGeometryContext() *MultiLineStringGeometryContext {
	var p = new(MultiLineStringGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_multiLineStringGeometry
	return p
}

func (*MultiLineStringGeometryContext) IsMultiLineStringGeometryContext() {}

func NewMultiLineStringGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiLineStringGeometryContext {
	var p = new(MultiLineStringGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_multiLineStringGeometry

	return p
}

func (s *MultiLineStringGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiLineStringGeometryContext) MULTILINESTRING() antlr.TerminalNode {
	return s.GetToken(wktParserMULTILINESTRING, 0)
}

func (s *MultiLineStringGeometryContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *MultiLineStringGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *MultiLineStringGeometryContext) AllLineString() []ILineStringContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineStringContext)(nil)).Elem())
	var tst = make([]ILineStringContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineStringContext)
		}
	}

	return tst
}

func (s *MultiLineStringGeometryContext) LineString(i int) ILineStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineStringContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineStringContext)
}

func (s *MultiLineStringGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *MultiLineStringGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *MultiLineStringGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *MultiLineStringGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiLineStringGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiLineStringGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterMultiLineStringGeometry(s)
	}
}

func (s *MultiLineStringGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitMultiLineStringGeometry(s)
	}
}

func (p *wktParser) MultiLineStringGeometry() (localctx IMultiLineStringGeometryContext) {
	this := p
	_ = this

	localctx = NewMultiLineStringGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, wktParserRULE_multiLineStringGeometry)
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
		p.SetState(191)
		p.Match(wktParserMULTILINESTRING)
	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		{
			p.SetState(192)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(193)
			p.LineString()
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(194)
				p.Match(wktParserCOMMA)
			}
			{
				p.SetState(195)
				p.LineString()
			}

			p.SetState(200)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(201)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(203)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMultiPolygonGeometryContext is an interface to support dynamic dispatch.
type IMultiPolygonGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiPolygonGeometryContext differentiates from other interfaces.
	IsMultiPolygonGeometryContext()
}

type MultiPolygonGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiPolygonGeometryContext() *MultiPolygonGeometryContext {
	var p = new(MultiPolygonGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_multiPolygonGeometry
	return p
}

func (*MultiPolygonGeometryContext) IsMultiPolygonGeometryContext() {}

func NewMultiPolygonGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPolygonGeometryContext {
	var p = new(MultiPolygonGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_multiPolygonGeometry

	return p
}

func (s *MultiPolygonGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPolygonGeometryContext) MULTIPOLYGON() antlr.TerminalNode {
	return s.GetToken(wktParserMULTIPOLYGON, 0)
}

func (s *MultiPolygonGeometryContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *MultiPolygonGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *MultiPolygonGeometryContext) AllPolygon() []IPolygonContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPolygonContext)(nil)).Elem())
	var tst = make([]IPolygonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPolygonContext)
		}
	}

	return tst
}

func (s *MultiPolygonGeometryContext) Polygon(i int) IPolygonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPolygonContext)
}

func (s *MultiPolygonGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *MultiPolygonGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *MultiPolygonGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *MultiPolygonGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPolygonGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPolygonGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterMultiPolygonGeometry(s)
	}
}

func (s *MultiPolygonGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitMultiPolygonGeometry(s)
	}
}

func (p *wktParser) MultiPolygonGeometry() (localctx IMultiPolygonGeometryContext) {
	this := p
	_ = this

	localctx = NewMultiPolygonGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, wktParserRULE_multiPolygonGeometry)
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
		p.SetState(206)
		p.Match(wktParserMULTIPOLYGON)
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		{
			p.SetState(207)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(208)
			p.Polygon()
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(209)
				p.Match(wktParserCOMMA)
			}
			{
				p.SetState(210)
				p.Polygon()
			}

			p.SetState(215)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(216)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(218)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMultiPolyhedralSurfaceGeometryContext is an interface to support dynamic dispatch.
type IMultiPolyhedralSurfaceGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiPolyhedralSurfaceGeometryContext differentiates from other interfaces.
	IsMultiPolyhedralSurfaceGeometryContext()
}

type MultiPolyhedralSurfaceGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiPolyhedralSurfaceGeometryContext() *MultiPolyhedralSurfaceGeometryContext {
	var p = new(MultiPolyhedralSurfaceGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_multiPolyhedralSurfaceGeometry
	return p
}

func (*MultiPolyhedralSurfaceGeometryContext) IsMultiPolyhedralSurfaceGeometryContext() {}

func NewMultiPolyhedralSurfaceGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiPolyhedralSurfaceGeometryContext {
	var p = new(MultiPolyhedralSurfaceGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_multiPolyhedralSurfaceGeometry

	return p
}

func (s *MultiPolyhedralSurfaceGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiPolyhedralSurfaceGeometryContext) POLYHEDRALSURFACE() antlr.TerminalNode {
	return s.GetToken(wktParserPOLYHEDRALSURFACE, 0)
}

func (s *MultiPolyhedralSurfaceGeometryContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *MultiPolyhedralSurfaceGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *MultiPolyhedralSurfaceGeometryContext) AllPolygon() []IPolygonContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPolygonContext)(nil)).Elem())
	var tst = make([]IPolygonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPolygonContext)
		}
	}

	return tst
}

func (s *MultiPolyhedralSurfaceGeometryContext) Polygon(i int) IPolygonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPolygonContext)
}

func (s *MultiPolyhedralSurfaceGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *MultiPolyhedralSurfaceGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *MultiPolyhedralSurfaceGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *MultiPolyhedralSurfaceGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiPolyhedralSurfaceGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiPolyhedralSurfaceGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterMultiPolyhedralSurfaceGeometry(s)
	}
}

func (s *MultiPolyhedralSurfaceGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitMultiPolyhedralSurfaceGeometry(s)
	}
}

func (p *wktParser) MultiPolyhedralSurfaceGeometry() (localctx IMultiPolyhedralSurfaceGeometryContext) {
	this := p
	_ = this

	localctx = NewMultiPolyhedralSurfaceGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, wktParserRULE_multiPolyhedralSurfaceGeometry)
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
		p.SetState(221)
		p.Match(wktParserPOLYHEDRALSURFACE)
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		{
			p.SetState(222)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(223)
			p.Polygon()
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(224)
				p.Match(wktParserCOMMA)
			}
			{
				p.SetState(225)
				p.Polygon()
			}

			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(231)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(233)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMultiTinGeometryContext is an interface to support dynamic dispatch.
type IMultiTinGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultiTinGeometryContext differentiates from other interfaces.
	IsMultiTinGeometryContext()
}

type MultiTinGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiTinGeometryContext() *MultiTinGeometryContext {
	var p = new(MultiTinGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_multiTinGeometry
	return p
}

func (*MultiTinGeometryContext) IsMultiTinGeometryContext() {}

func NewMultiTinGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiTinGeometryContext {
	var p = new(MultiTinGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_multiTinGeometry

	return p
}

func (s *MultiTinGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiTinGeometryContext) TIN() antlr.TerminalNode {
	return s.GetToken(wktParserTIN, 0)
}

func (s *MultiTinGeometryContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *MultiTinGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *MultiTinGeometryContext) AllPolygon() []IPolygonContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPolygonContext)(nil)).Elem())
	var tst = make([]IPolygonContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPolygonContext)
		}
	}

	return tst
}

func (s *MultiTinGeometryContext) Polygon(i int) IPolygonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPolygonContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPolygonContext)
}

func (s *MultiTinGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *MultiTinGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *MultiTinGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *MultiTinGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiTinGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiTinGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterMultiTinGeometry(s)
	}
}

func (s *MultiTinGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitMultiTinGeometry(s)
	}
}

func (p *wktParser) MultiTinGeometry() (localctx IMultiTinGeometryContext) {
	this := p
	_ = this

	localctx = NewMultiTinGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, wktParserRULE_multiTinGeometry)
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
		p.SetState(236)
		p.Match(wktParserTIN)
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		{
			p.SetState(237)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(238)
			p.Polygon()
		}
		p.SetState(243)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(239)
				p.Match(wktParserCOMMA)
			}
			{
				p.SetState(240)
				p.Polygon()
			}

			p.SetState(245)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(246)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		{
			p.SetState(248)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICircularStringGeometryContext is an interface to support dynamic dispatch.
type ICircularStringGeometryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCircularStringGeometryContext differentiates from other interfaces.
	IsCircularStringGeometryContext()
}

type CircularStringGeometryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCircularStringGeometryContext() *CircularStringGeometryContext {
	var p = new(CircularStringGeometryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_circularStringGeometry
	return p
}

func (*CircularStringGeometryContext) IsCircularStringGeometryContext() {}

func NewCircularStringGeometryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CircularStringGeometryContext {
	var p = new(CircularStringGeometryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_circularStringGeometry

	return p
}

func (s *CircularStringGeometryContext) GetParser() antlr.Parser { return s.parser }

func (s *CircularStringGeometryContext) CIRCULARSTRING() antlr.TerminalNode {
	return s.GetToken(wktParserCIRCULARSTRING, 0)
}

func (s *CircularStringGeometryContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *CircularStringGeometryContext) AllPoint() []IPointContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPointContext)(nil)).Elem())
	var tst = make([]IPointContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPointContext)
		}
	}

	return tst
}

func (s *CircularStringGeometryContext) Point(i int) IPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPointContext)
}

func (s *CircularStringGeometryContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *CircularStringGeometryContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *CircularStringGeometryContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *CircularStringGeometryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CircularStringGeometryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CircularStringGeometryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterCircularStringGeometry(s)
	}
}

func (s *CircularStringGeometryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitCircularStringGeometry(s)
	}
}

func (p *wktParser) CircularStringGeometry() (localctx ICircularStringGeometryContext) {
	this := p
	_ = this

	localctx = NewCircularStringGeometryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, wktParserRULE_circularStringGeometry)
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
		p.SetState(251)
		p.Match(wktParserCIRCULARSTRING)
	}
	{
		p.SetState(252)
		p.Match(wktParserLPAR)
	}
	{
		p.SetState(253)
		p.Point()
	}
	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == wktParserCOMMA {
		{
			p.SetState(254)
			p.Match(wktParserCOMMA)
		}
		{
			p.SetState(255)
			p.Point()
		}

		p.SetState(260)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(261)
		p.Match(wktParserRPAR)
	}

	return localctx
}

// IPointOrClosedPointContext is an interface to support dynamic dispatch.
type IPointOrClosedPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointOrClosedPointContext differentiates from other interfaces.
	IsPointOrClosedPointContext()
}

type PointOrClosedPointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointOrClosedPointContext() *PointOrClosedPointContext {
	var p = new(PointOrClosedPointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_pointOrClosedPoint
	return p
}

func (*PointOrClosedPointContext) IsPointOrClosedPointContext() {}

func NewPointOrClosedPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointOrClosedPointContext {
	var p = new(PointOrClosedPointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_pointOrClosedPoint

	return p
}

func (s *PointOrClosedPointContext) GetParser() antlr.Parser { return s.parser }

func (s *PointOrClosedPointContext) Point() IPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointContext)
}

func (s *PointOrClosedPointContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *PointOrClosedPointContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *PointOrClosedPointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointOrClosedPointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointOrClosedPointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterPointOrClosedPoint(s)
	}
}

func (s *PointOrClosedPointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitPointOrClosedPoint(s)
	}
}

func (p *wktParser) PointOrClosedPoint() (localctx IPointOrClosedPointContext) {
	this := p
	_ = this

	localctx = NewPointOrClosedPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, wktParserRULE_pointOrClosedPoint)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(268)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserDECIMAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Point()
		}

	case wktParserLPAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(265)
			p.Point()
		}
		{
			p.SetState(266)
			p.Match(wktParserRPAR)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPolygonContext is an interface to support dynamic dispatch.
type IPolygonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPolygonContext differentiates from other interfaces.
	IsPolygonContext()
}

type PolygonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPolygonContext() *PolygonContext {
	var p = new(PolygonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_polygon
	return p
}

func (*PolygonContext) IsPolygonContext() {}

func NewPolygonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PolygonContext {
	var p = new(PolygonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_polygon

	return p
}

func (s *PolygonContext) GetParser() antlr.Parser { return s.parser }

func (s *PolygonContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *PolygonContext) AllLineString() []ILineStringContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILineStringContext)(nil)).Elem())
	var tst = make([]ILineStringContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILineStringContext)
		}
	}

	return tst
}

func (s *PolygonContext) LineString(i int) ILineStringContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILineStringContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILineStringContext)
}

func (s *PolygonContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *PolygonContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *PolygonContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *PolygonContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *PolygonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PolygonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PolygonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterPolygon(s)
	}
}

func (s *PolygonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitPolygon(s)
	}
}

func (p *wktParser) Polygon() (localctx IPolygonContext) {
	this := p
	_ = this

	localctx = NewPolygonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, wktParserRULE_polygon)
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

	p.SetState(282)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(270)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(271)
			p.LineString()
		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(272)
				p.Match(wktParserCOMMA)
			}
			{
				p.SetState(273)
				p.LineString()
			}

			p.SetState(278)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(279)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(281)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILineStringContext is an interface to support dynamic dispatch.
type ILineStringContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLineStringContext differentiates from other interfaces.
	IsLineStringContext()
}

type LineStringContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineStringContext() *LineStringContext {
	var p = new(LineStringContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_lineString
	return p
}

func (*LineStringContext) IsLineStringContext() {}

func NewLineStringContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineStringContext {
	var p = new(LineStringContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_lineString

	return p
}

func (s *LineStringContext) GetParser() antlr.Parser { return s.parser }

func (s *LineStringContext) LPAR() antlr.TerminalNode {
	return s.GetToken(wktParserLPAR, 0)
}

func (s *LineStringContext) AllPoint() []IPointContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPointContext)(nil)).Elem())
	var tst = make([]IPointContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPointContext)
		}
	}

	return tst
}

func (s *LineStringContext) Point(i int) IPointContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPointContext)
}

func (s *LineStringContext) RPAR() antlr.TerminalNode {
	return s.GetToken(wktParserRPAR, 0)
}

func (s *LineStringContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(wktParserCOMMA)
}

func (s *LineStringContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(wktParserCOMMA, i)
}

func (s *LineStringContext) EMPTY() antlr.TerminalNode {
	return s.GetToken(wktParserEMPTY, 0)
}

func (s *LineStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineStringContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterLineString(s)
	}
}

func (s *LineStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitLineString(s)
	}
}

func (p *wktParser) LineString() (localctx ILineStringContext) {
	this := p
	_ = this

	localctx = NewLineStringContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, wktParserRULE_lineString)
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

	p.SetState(296)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case wktParserLPAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.Match(wktParserLPAR)
		}
		{
			p.SetState(285)
			p.Point()
		}
		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == wktParserCOMMA {
			{
				p.SetState(286)
				p.Match(wktParserCOMMA)
			}
			{
				p.SetState(287)
				p.Point()
			}

			p.SetState(292)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(293)
			p.Match(wktParserRPAR)
		}

	case wktParserEMPTY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.Match(wktParserEMPTY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPointContext is an interface to support dynamic dispatch.
type IPointContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointContext differentiates from other interfaces.
	IsPointContext()
}

type PointContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointContext() *PointContext {
	var p = new(PointContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = wktParserRULE_point
	return p
}

func (*PointContext) IsPointContext() {}

func NewPointContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointContext {
	var p = new(PointContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_point

	return p
}

func (s *PointContext) GetParser() antlr.Parser { return s.parser }

func (s *PointContext) AllDECIMAL() []antlr.TerminalNode {
	return s.GetTokens(wktParserDECIMAL)
}

func (s *PointContext) DECIMAL(i int) antlr.TerminalNode {
	return s.GetToken(wktParserDECIMAL, i)
}

func (s *PointContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterPoint(s)
	}
}

func (s *PointContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitPoint(s)
	}
}

func (p *wktParser) Point() (localctx IPointContext) {
	this := p
	_ = this

	localctx = NewPointContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, wktParserRULE_point)
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
	p.SetState(299)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == wktParserDECIMAL {
		{
			p.SetState(298)
			p.Match(wktParserDECIMAL)
		}

		p.SetState(301)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = wktParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = wktParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) STRING() antlr.TerminalNode {
	return s.GetToken(wktParserSTRING, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(wktListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *wktParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, wktParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(303)
		p.Match(wktParserSTRING)
	}

	return localctx
}
