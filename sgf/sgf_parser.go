// Code generated from sgf.g4 by ANTLR 4.9.3. DO NOT EDIT.

package sgf // sgf
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 80, 320,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 6, 2, 38, 10, 2, 13, 2, 14, 2, 39, 3, 3, 3, 3, 3, 3, 7, 3, 45,
	10, 3, 12, 3, 14, 3, 48, 11, 3, 3, 3, 3, 3, 3, 4, 6, 4, 53, 10, 4, 13,
	4, 14, 4, 54, 3, 5, 3, 5, 7, 5, 59, 10, 5, 12, 5, 14, 5, 62, 11, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5,
	6, 76, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 84, 10, 7, 3, 8,
	3, 8, 6, 8, 88, 10, 8, 13, 8, 14, 8, 89, 3, 8, 3, 8, 6, 8, 94, 10, 8, 13,
	8, 14, 8, 95, 3, 8, 3, 8, 6, 8, 100, 10, 8, 13, 8, 14, 8, 101, 3, 8, 3,
	8, 5, 8, 106, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 124, 10, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 134, 10, 10, 3, 11,
	3, 11, 6, 11, 138, 10, 11, 13, 11, 14, 11, 139, 3, 11, 3, 11, 6, 11, 144,
	10, 11, 13, 11, 14, 11, 145, 3, 11, 3, 11, 3, 11, 6, 11, 151, 10, 11, 13,
	11, 14, 11, 152, 5, 11, 155, 10, 11, 3, 11, 3, 11, 6, 11, 159, 10, 11,
	13, 11, 14, 11, 160, 3, 11, 3, 11, 6, 11, 165, 10, 11, 13, 11, 14, 11,
	166, 3, 11, 3, 11, 6, 11, 171, 10, 11, 13, 11, 14, 11, 172, 3, 11, 3, 11,
	6, 11, 177, 10, 11, 13, 11, 14, 11, 178, 3, 11, 3, 11, 6, 11, 183, 10,
	11, 13, 11, 14, 11, 184, 3, 11, 3, 11, 6, 11, 189, 10, 11, 13, 11, 14,
	11, 190, 5, 11, 193, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 207, 10, 12, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 251,
	10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14,
	261, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 6, 15, 269, 10,
	15, 13, 15, 14, 15, 270, 5, 15, 273, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 285, 10, 16, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 6, 17, 294, 10, 17, 13, 17, 14,
	17, 295, 5, 17, 298, 10, 17, 3, 17, 3, 17, 3, 17, 6, 17, 303, 10, 17, 13,
	17, 14, 17, 304, 5, 17, 307, 10, 17, 5, 17, 309, 10, 17, 3, 18, 3, 18,
	3, 18, 6, 18, 314, 10, 18, 13, 18, 14, 18, 315, 5, 18, 318, 10, 18, 3,
	18, 2, 2, 19, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32,
	34, 2, 3, 3, 2, 78, 79, 2, 397, 2, 37, 3, 2, 2, 2, 4, 41, 3, 2, 2, 2, 6,
	52, 3, 2, 2, 2, 8, 56, 3, 2, 2, 2, 10, 75, 3, 2, 2, 2, 12, 83, 3, 2, 2,
	2, 14, 105, 3, 2, 2, 2, 16, 123, 3, 2, 2, 2, 18, 133, 3, 2, 2, 2, 20, 192,
	3, 2, 2, 2, 22, 206, 3, 2, 2, 2, 24, 250, 3, 2, 2, 2, 26, 260, 3, 2, 2,
	2, 28, 272, 3, 2, 2, 2, 30, 284, 3, 2, 2, 2, 32, 308, 3, 2, 2, 2, 34, 310,
	3, 2, 2, 2, 36, 38, 5, 4, 3, 2, 37, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2,
	39, 37, 3, 2, 2, 2, 39, 40, 3, 2, 2, 2, 40, 3, 3, 2, 2, 2, 41, 42, 7, 3,
	2, 2, 42, 46, 5, 6, 4, 2, 43, 45, 5, 4, 3, 2, 44, 43, 3, 2, 2, 2, 45, 48,
	3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 49, 3, 2, 2, 2,
	48, 46, 3, 2, 2, 2, 49, 50, 7, 4, 2, 2, 50, 5, 3, 2, 2, 2, 51, 53, 5, 8,
	5, 2, 52, 51, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 54, 55,
	3, 2, 2, 2, 55, 7, 3, 2, 2, 2, 56, 60, 7, 5, 2, 2, 57, 59, 5, 10, 6, 2,
	58, 57, 3, 2, 2, 2, 59, 62, 3, 2, 2, 2, 60, 58, 3, 2, 2, 2, 60, 61, 3,
	2, 2, 2, 61, 9, 3, 2, 2, 2, 62, 60, 3, 2, 2, 2, 63, 76, 5, 12, 7, 2, 64,
	76, 5, 14, 8, 2, 65, 76, 5, 16, 9, 2, 66, 76, 5, 18, 10, 2, 67, 76, 5,
	20, 11, 2, 68, 76, 5, 22, 12, 2, 69, 76, 5, 24, 13, 2, 70, 76, 5, 26, 14,
	2, 71, 76, 5, 28, 15, 2, 72, 76, 5, 30, 16, 2, 73, 76, 5, 32, 17, 2, 74,
	76, 5, 34, 18, 2, 75, 63, 3, 2, 2, 2, 75, 64, 3, 2, 2, 2, 75, 65, 3, 2,
	2, 2, 75, 66, 3, 2, 2, 2, 75, 67, 3, 2, 2, 2, 75, 68, 3, 2, 2, 2, 75, 69,
	3, 2, 2, 2, 75, 70, 3, 2, 2, 2, 75, 71, 3, 2, 2, 2, 75, 72, 3, 2, 2, 2,
	75, 73, 3, 2, 2, 2, 75, 74, 3, 2, 2, 2, 76, 11, 3, 2, 2, 2, 77, 78, 7,
	76, 2, 2, 78, 84, 9, 2, 2, 2, 79, 80, 7, 6, 2, 2, 80, 84, 7, 78, 2, 2,
	81, 82, 7, 7, 2, 2, 82, 84, 7, 79, 2, 2, 83, 77, 3, 2, 2, 2, 83, 79, 3,
	2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 13, 3, 2, 2, 2, 85, 87, 7, 8, 2, 2, 86,
	88, 7, 79, 2, 2, 87, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 87, 3, 2,
	2, 2, 89, 90, 3, 2, 2, 2, 90, 106, 3, 2, 2, 2, 91, 93, 7, 9, 2, 2, 92,
	94, 7, 79, 2, 2, 93, 92, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 93, 3, 2,
	2, 2, 95, 96, 3, 2, 2, 2, 96, 106, 3, 2, 2, 2, 97, 99, 7, 10, 2, 2, 98,
	100, 7, 79, 2, 2, 99, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 99, 3,
	2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 106, 3, 2, 2, 2, 103, 104, 7, 11, 2,
	2, 104, 106, 7, 79, 2, 2, 105, 85, 3, 2, 2, 2, 105, 91, 3, 2, 2, 2, 105,
	97, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 106, 15, 3, 2, 2, 2, 107, 108, 7,
	12, 2, 2, 108, 124, 7, 79, 2, 2, 109, 110, 7, 13, 2, 2, 110, 124, 7, 79,
	2, 2, 111, 112, 7, 14, 2, 2, 112, 124, 7, 79, 2, 2, 113, 114, 7, 15, 2,
	2, 114, 124, 7, 79, 2, 2, 115, 116, 7, 16, 2, 2, 116, 124, 7, 79, 2, 2,
	117, 118, 7, 17, 2, 2, 118, 124, 7, 79, 2, 2, 119, 120, 7, 18, 2, 2, 120,
	124, 7, 79, 2, 2, 121, 122, 7, 19, 2, 2, 122, 124, 7, 79, 2, 2, 123, 107,
	3, 2, 2, 2, 123, 109, 3, 2, 2, 2, 123, 111, 3, 2, 2, 2, 123, 113, 3, 2,
	2, 2, 123, 115, 3, 2, 2, 2, 123, 117, 3, 2, 2, 2, 123, 119, 3, 2, 2, 2,
	123, 121, 3, 2, 2, 2, 124, 17, 3, 2, 2, 2, 125, 126, 7, 20, 2, 2, 126,
	134, 7, 79, 2, 2, 127, 128, 7, 21, 2, 2, 128, 134, 7, 78, 2, 2, 129, 130,
	7, 22, 2, 2, 130, 134, 7, 78, 2, 2, 131, 132, 7, 23, 2, 2, 132, 134, 7,
	79, 2, 2, 133, 125, 3, 2, 2, 2, 133, 127, 3, 2, 2, 2, 133, 129, 3, 2, 2,
	2, 133, 131, 3, 2, 2, 2, 134, 19, 3, 2, 2, 2, 135, 137, 7, 24, 2, 2, 136,
	138, 7, 79, 2, 2, 137, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 137,
	3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 193, 3, 2, 2, 2, 141, 143, 7, 25,
	2, 2, 142, 144, 7, 79, 2, 2, 143, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2,
	145, 143, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 193, 3, 2, 2, 2, 147,
	154, 7, 26, 2, 2, 148, 155, 7, 78, 2, 2, 149, 151, 7, 79, 2, 2, 150, 149,
	3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152, 153, 3, 2,
	2, 2, 153, 155, 3, 2, 2, 2, 154, 148, 3, 2, 2, 2, 154, 150, 3, 2, 2, 2,
	155, 193, 3, 2, 2, 2, 156, 158, 7, 27, 2, 2, 157, 159, 7, 79, 2, 2, 158,
	157, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160, 158, 3, 2, 2, 2, 160, 161,
	3, 2, 2, 2, 161, 193, 3, 2, 2, 2, 162, 164, 7, 28, 2, 2, 163, 165, 7, 79,
	2, 2, 164, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166, 164, 3, 2, 2, 2,
	166, 167, 3, 2, 2, 2, 167, 193, 3, 2, 2, 2, 168, 170, 7, 29, 2, 2, 169,
	171, 7, 79, 2, 2, 170, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172, 170,
	3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 193, 3, 2, 2, 2, 174, 176, 7, 30,
	2, 2, 175, 177, 7, 79, 2, 2, 176, 175, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2,
	178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 193, 3, 2, 2, 2, 180,
	182, 7, 31, 2, 2, 181, 183, 7, 79, 2, 2, 182, 181, 3, 2, 2, 2, 183, 184,
	3, 2, 2, 2, 184, 182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 193, 3, 2,
	2, 2, 186, 188, 7, 32, 2, 2, 187, 189, 7, 79, 2, 2, 188, 187, 3, 2, 2,
	2, 189, 190, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191,
	193, 3, 2, 2, 2, 192, 135, 3, 2, 2, 2, 192, 141, 3, 2, 2, 2, 192, 147,
	3, 2, 2, 2, 192, 156, 3, 2, 2, 2, 192, 162, 3, 2, 2, 2, 192, 168, 3, 2,
	2, 2, 192, 174, 3, 2, 2, 2, 192, 180, 3, 2, 2, 2, 192, 186, 3, 2, 2, 2,
	193, 21, 3, 2, 2, 2, 194, 195, 7, 33, 2, 2, 195, 207, 7, 79, 2, 2, 196,
	197, 7, 34, 2, 2, 197, 207, 7, 79, 2, 2, 198, 199, 7, 35, 2, 2, 199, 207,
	7, 79, 2, 2, 200, 201, 7, 36, 2, 2, 201, 207, 7, 79, 2, 2, 202, 203, 7,
	37, 2, 2, 203, 207, 7, 79, 2, 2, 204, 205, 7, 38, 2, 2, 205, 207, 7, 79,
	2, 2, 206, 194, 3, 2, 2, 2, 206, 196, 3, 2, 2, 2, 206, 198, 3, 2, 2, 2,
	206, 200, 3, 2, 2, 2, 206, 202, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 207,
	23, 3, 2, 2, 2, 208, 209, 7, 39, 2, 2, 209, 251, 7, 79, 2, 2, 210, 211,
	7, 40, 2, 2, 211, 251, 7, 79, 2, 2, 212, 213, 7, 41, 2, 2, 213, 251, 7,
	79, 2, 2, 214, 215, 7, 42, 2, 2, 215, 251, 7, 79, 2, 2, 216, 217, 7, 43,
	2, 2, 217, 251, 7, 79, 2, 2, 218, 219, 7, 44, 2, 2, 219, 251, 7, 79, 2,
	2, 220, 221, 7, 45, 2, 2, 221, 251, 7, 79, 2, 2, 222, 223, 7, 46, 2, 2,
	223, 251, 7, 79, 2, 2, 224, 225, 7, 47, 2, 2, 225, 251, 7, 79, 2, 2, 226,
	227, 7, 48, 2, 2, 227, 251, 7, 79, 2, 2, 228, 229, 7, 49, 2, 2, 229, 251,
	7, 79, 2, 2, 230, 231, 7, 50, 2, 2, 231, 251, 7, 79, 2, 2, 232, 233, 7,
	51, 2, 2, 233, 251, 7, 79, 2, 2, 234, 235, 7, 52, 2, 2, 235, 251, 7, 79,
	2, 2, 236, 237, 7, 53, 2, 2, 237, 251, 7, 79, 2, 2, 238, 239, 7, 54, 2,
	2, 239, 251, 7, 79, 2, 2, 240, 241, 7, 55, 2, 2, 241, 251, 7, 79, 2, 2,
	242, 243, 7, 56, 2, 2, 243, 251, 7, 79, 2, 2, 244, 245, 7, 57, 2, 2, 245,
	251, 7, 79, 2, 2, 246, 247, 7, 58, 2, 2, 247, 251, 7, 79, 2, 2, 248, 249,
	7, 59, 2, 2, 249, 251, 7, 79, 2, 2, 250, 208, 3, 2, 2, 2, 250, 210, 3,
	2, 2, 2, 250, 212, 3, 2, 2, 2, 250, 214, 3, 2, 2, 2, 250, 216, 3, 2, 2,
	2, 250, 218, 3, 2, 2, 2, 250, 220, 3, 2, 2, 2, 250, 222, 3, 2, 2, 2, 250,
	224, 3, 2, 2, 2, 250, 226, 3, 2, 2, 2, 250, 228, 3, 2, 2, 2, 250, 230,
	3, 2, 2, 2, 250, 232, 3, 2, 2, 2, 250, 234, 3, 2, 2, 2, 250, 236, 3, 2,
	2, 2, 250, 238, 3, 2, 2, 2, 250, 240, 3, 2, 2, 2, 250, 242, 3, 2, 2, 2,
	250, 244, 3, 2, 2, 2, 250, 246, 3, 2, 2, 2, 250, 248, 3, 2, 2, 2, 251,
	25, 3, 2, 2, 2, 252, 253, 7, 60, 2, 2, 253, 261, 7, 79, 2, 2, 254, 255,
	7, 61, 2, 2, 255, 261, 7, 79, 2, 2, 256, 257, 7, 62, 2, 2, 257, 261, 7,
	79, 2, 2, 258, 259, 7, 63, 2, 2, 259, 261, 7, 79, 2, 2, 260, 252, 3, 2,
	2, 2, 260, 254, 3, 2, 2, 2, 260, 256, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2,
	261, 27, 3, 2, 2, 2, 262, 263, 7, 64, 2, 2, 263, 273, 9, 2, 2, 2, 264,
	265, 7, 65, 2, 2, 265, 273, 7, 79, 2, 2, 266, 268, 7, 66, 2, 2, 267, 269,
	7, 79, 2, 2, 268, 267, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 268, 3, 2,
	2, 2, 270, 271, 3, 2, 2, 2, 271, 273, 3, 2, 2, 2, 272, 262, 3, 2, 2, 2,
	272, 264, 3, 2, 2, 2, 272, 266, 3, 2, 2, 2, 273, 29, 3, 2, 2, 2, 274, 275,
	7, 67, 2, 2, 275, 285, 7, 79, 2, 2, 276, 277, 7, 68, 2, 2, 277, 285, 7,
	79, 2, 2, 278, 279, 7, 69, 2, 2, 279, 285, 7, 79, 2, 2, 280, 281, 7, 70,
	2, 2, 281, 285, 7, 79, 2, 2, 282, 283, 7, 71, 2, 2, 283, 285, 7, 79, 2,
	2, 284, 274, 3, 2, 2, 2, 284, 276, 3, 2, 2, 2, 284, 278, 3, 2, 2, 2, 284,
	280, 3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 285, 31, 3, 2, 2, 2, 286, 287, 7,
	72, 2, 2, 287, 309, 7, 79, 2, 2, 288, 289, 7, 73, 2, 2, 289, 309, 7, 79,
	2, 2, 290, 297, 7, 74, 2, 2, 291, 298, 7, 78, 2, 2, 292, 294, 7, 79, 2,
	2, 293, 292, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 293, 3, 2, 2, 2, 295,
	296, 3, 2, 2, 2, 296, 298, 3, 2, 2, 2, 297, 291, 3, 2, 2, 2, 297, 293,
	3, 2, 2, 2, 298, 309, 3, 2, 2, 2, 299, 306, 7, 75, 2, 2, 300, 307, 7, 78,
	2, 2, 301, 303, 7, 79, 2, 2, 302, 301, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2,
	304, 302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 307, 3, 2, 2, 2, 306,
	300, 3, 2, 2, 2, 306, 302, 3, 2, 2, 2, 307, 309, 3, 2, 2, 2, 308, 286,
	3, 2, 2, 2, 308, 288, 3, 2, 2, 2, 308, 290, 3, 2, 2, 2, 308, 299, 3, 2,
	2, 2, 309, 33, 3, 2, 2, 2, 310, 317, 7, 77, 2, 2, 311, 318, 7, 78, 2, 2,
	312, 314, 7, 79, 2, 2, 313, 312, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315,
	313, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 318, 3, 2, 2, 2, 317, 311,
	3, 2, 2, 2, 317, 313, 3, 2, 2, 2, 318, 35, 3, 2, 2, 2, 38, 39, 46, 54,
	60, 75, 83, 89, 95, 101, 105, 123, 133, 139, 145, 152, 154, 160, 166, 172,
	178, 184, 190, 192, 206, 250, 260, 270, 272, 284, 295, 297, 304, 306, 308,
	315, 317,
}
var literalNames = []string{
	"", "'('", "')'", "';'", "'KO'", "'MN'", "'AB'", "'AE'", "'AW'", "'PL'",
	"'C'", "'DM'", "'GB'", "'GW'", "'HO'", "'N'", "'UC'", "'V'", "'BM'", "'DO'",
	"'IT'", "'TE'", "'AR'", "'CR'", "'DD'", "'LB'", "'LN'", "'MA'", "'SL'",
	"'SQ'", "'TR'", "'AP'", "'CA'", "'FF'", "'GM'", "'ST'", "'SZ'", "'AN'",
	"'BR'", "'BT'", "'CP'", "'DT'", "'EV'", "'GN'", "'GC'", "'ON'", "'OT'",
	"'PB'", "'PC'", "'PW'", "'RE'", "'RO'", "'RU'", "'SO'", "'TM'", "'US'",
	"'WR'", "'WT'", "'BL'", "'OB'", "'OW'", "'WL'", "'FG'", "'PM'", "'VW'",
	"'AS'", "'IP'", "'IY'", "'SE'", "'SU'", "'HA'", "'KM'", "'TB'", "'TW'",
	"", "", "'[]'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "COLOR", "UCLETTER", "NONE", "TEXT", "WS",
}

var ruleNames = []string{
	"collection", "gameTree", "sequence", "node", "property_", "move", "setup",
	"nodeAnnotation", "moveAnnotation", "markup", "root", "gameInfo", "timing",
	"misc", "loa", "go_", "privateProp",
}

type sgfParser struct {
	*antlr.BaseParser
}

// NewsgfParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *sgfParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewsgfParser(input antlr.TokenStream) *sgfParser {
	this := new(sgfParser)
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
	this.GrammarFileName = "sgf.g4"

	return this
}

// sgfParser tokens.
const (
	sgfParserEOF      = antlr.TokenEOF
	sgfParserT__0     = 1
	sgfParserT__1     = 2
	sgfParserT__2     = 3
	sgfParserT__3     = 4
	sgfParserT__4     = 5
	sgfParserT__5     = 6
	sgfParserT__6     = 7
	sgfParserT__7     = 8
	sgfParserT__8     = 9
	sgfParserT__9     = 10
	sgfParserT__10    = 11
	sgfParserT__11    = 12
	sgfParserT__12    = 13
	sgfParserT__13    = 14
	sgfParserT__14    = 15
	sgfParserT__15    = 16
	sgfParserT__16    = 17
	sgfParserT__17    = 18
	sgfParserT__18    = 19
	sgfParserT__19    = 20
	sgfParserT__20    = 21
	sgfParserT__21    = 22
	sgfParserT__22    = 23
	sgfParserT__23    = 24
	sgfParserT__24    = 25
	sgfParserT__25    = 26
	sgfParserT__26    = 27
	sgfParserT__27    = 28
	sgfParserT__28    = 29
	sgfParserT__29    = 30
	sgfParserT__30    = 31
	sgfParserT__31    = 32
	sgfParserT__32    = 33
	sgfParserT__33    = 34
	sgfParserT__34    = 35
	sgfParserT__35    = 36
	sgfParserT__36    = 37
	sgfParserT__37    = 38
	sgfParserT__38    = 39
	sgfParserT__39    = 40
	sgfParserT__40    = 41
	sgfParserT__41    = 42
	sgfParserT__42    = 43
	sgfParserT__43    = 44
	sgfParserT__44    = 45
	sgfParserT__45    = 46
	sgfParserT__46    = 47
	sgfParserT__47    = 48
	sgfParserT__48    = 49
	sgfParserT__49    = 50
	sgfParserT__50    = 51
	sgfParserT__51    = 52
	sgfParserT__52    = 53
	sgfParserT__53    = 54
	sgfParserT__54    = 55
	sgfParserT__55    = 56
	sgfParserT__56    = 57
	sgfParserT__57    = 58
	sgfParserT__58    = 59
	sgfParserT__59    = 60
	sgfParserT__60    = 61
	sgfParserT__61    = 62
	sgfParserT__62    = 63
	sgfParserT__63    = 64
	sgfParserT__64    = 65
	sgfParserT__65    = 66
	sgfParserT__66    = 67
	sgfParserT__67    = 68
	sgfParserT__68    = 69
	sgfParserT__69    = 70
	sgfParserT__70    = 71
	sgfParserT__71    = 72
	sgfParserT__72    = 73
	sgfParserCOLOR    = 74
	sgfParserUCLETTER = 75
	sgfParserNONE     = 76
	sgfParserTEXT     = 77
	sgfParserWS       = 78
)

// sgfParser rules.
const (
	sgfParserRULE_collection     = 0
	sgfParserRULE_gameTree       = 1
	sgfParserRULE_sequence       = 2
	sgfParserRULE_node           = 3
	sgfParserRULE_property_      = 4
	sgfParserRULE_move           = 5
	sgfParserRULE_setup          = 6
	sgfParserRULE_nodeAnnotation = 7
	sgfParserRULE_moveAnnotation = 8
	sgfParserRULE_markup         = 9
	sgfParserRULE_root           = 10
	sgfParserRULE_gameInfo       = 11
	sgfParserRULE_timing         = 12
	sgfParserRULE_misc           = 13
	sgfParserRULE_loa            = 14
	sgfParserRULE_go_            = 15
	sgfParserRULE_privateProp    = 16
)

// ICollectionContext is an interface to support dynamic dispatch.
type ICollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCollectionContext differentiates from other interfaces.
	IsCollectionContext()
}

type CollectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionContext() *CollectionContext {
	var p = new(CollectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_collection
	return p
}

func (*CollectionContext) IsCollectionContext() {}

func NewCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionContext {
	var p = new(CollectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_collection

	return p
}

func (s *CollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionContext) AllGameTree() []IGameTreeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGameTreeContext)(nil)).Elem())
	var tst = make([]IGameTreeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGameTreeContext)
		}
	}

	return tst
}

func (s *CollectionContext) GameTree(i int) IGameTreeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGameTreeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGameTreeContext)
}

func (s *CollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterCollection(s)
	}
}

func (s *CollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitCollection(s)
	}
}

func (p *sgfParser) Collection() (localctx ICollectionContext) {
	this := p
	_ = this

	localctx = NewCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, sgfParserRULE_collection)
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
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == sgfParserT__0 {
		{
			p.SetState(34)
			p.GameTree()
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGameTreeContext is an interface to support dynamic dispatch.
type IGameTreeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGameTreeContext differentiates from other interfaces.
	IsGameTreeContext()
}

type GameTreeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGameTreeContext() *GameTreeContext {
	var p = new(GameTreeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_gameTree
	return p
}

func (*GameTreeContext) IsGameTreeContext() {}

func NewGameTreeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GameTreeContext {
	var p = new(GameTreeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_gameTree

	return p
}

func (s *GameTreeContext) GetParser() antlr.Parser { return s.parser }

func (s *GameTreeContext) Sequence() ISequenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISequenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISequenceContext)
}

func (s *GameTreeContext) AllGameTree() []IGameTreeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGameTreeContext)(nil)).Elem())
	var tst = make([]IGameTreeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGameTreeContext)
		}
	}

	return tst
}

func (s *GameTreeContext) GameTree(i int) IGameTreeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGameTreeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGameTreeContext)
}

func (s *GameTreeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GameTreeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GameTreeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterGameTree(s)
	}
}

func (s *GameTreeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitGameTree(s)
	}
}

func (p *sgfParser) GameTree() (localctx IGameTreeContext) {
	this := p
	_ = this

	localctx = NewGameTreeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, sgfParserRULE_gameTree)
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
		p.SetState(39)
		p.Match(sgfParserT__0)
	}
	{
		p.SetState(40)
		p.Sequence()
	}
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == sgfParserT__0 {
		{
			p.SetState(41)
			p.GameTree()
		}

		p.SetState(46)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(47)
		p.Match(sgfParserT__1)
	}

	return localctx
}

// ISequenceContext is an interface to support dynamic dispatch.
type ISequenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSequenceContext differentiates from other interfaces.
	IsSequenceContext()
}

type SequenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySequenceContext() *SequenceContext {
	var p = new(SequenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_sequence
	return p
}

func (*SequenceContext) IsSequenceContext() {}

func NewSequenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SequenceContext {
	var p = new(SequenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_sequence

	return p
}

func (s *SequenceContext) GetParser() antlr.Parser { return s.parser }

func (s *SequenceContext) AllNode() []INodeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INodeContext)(nil)).Elem())
	var tst = make([]INodeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INodeContext)
		}
	}

	return tst
}

func (s *SequenceContext) Node(i int) INodeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INodeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INodeContext)
}

func (s *SequenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SequenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SequenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterSequence(s)
	}
}

func (s *SequenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitSequence(s)
	}
}

func (p *sgfParser) Sequence() (localctx ISequenceContext) {
	this := p
	_ = this

	localctx = NewSequenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, sgfParserRULE_sequence)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == sgfParserT__2 {
		{
			p.SetState(49)
			p.Node()
		}

		p.SetState(52)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// INodeContext is an interface to support dynamic dispatch.
type INodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNodeContext differentiates from other interfaces.
	IsNodeContext()
}

type NodeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeContext() *NodeContext {
	var p = new(NodeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_node
	return p
}

func (*NodeContext) IsNodeContext() {}

func NewNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeContext {
	var p = new(NodeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_node

	return p
}

func (s *NodeContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeContext) AllProperty_() []IProperty_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProperty_Context)(nil)).Elem())
	var tst = make([]IProperty_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProperty_Context)
		}
	}

	return tst
}

func (s *NodeContext) Property_(i int) IProperty_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProperty_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProperty_Context)
}

func (s *NodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterNode(s)
	}
}

func (s *NodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitNode(s)
	}
}

func (p *sgfParser) Node() (localctx INodeContext) {
	this := p
	_ = this

	localctx = NewNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, sgfParserRULE_node)
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
		p.SetState(54)
		p.Match(sgfParserT__2)
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<sgfParserT__3)|(1<<sgfParserT__4)|(1<<sgfParserT__5)|(1<<sgfParserT__6)|(1<<sgfParserT__7)|(1<<sgfParserT__8)|(1<<sgfParserT__9)|(1<<sgfParserT__10)|(1<<sgfParserT__11)|(1<<sgfParserT__12)|(1<<sgfParserT__13)|(1<<sgfParserT__14)|(1<<sgfParserT__15)|(1<<sgfParserT__16)|(1<<sgfParserT__17)|(1<<sgfParserT__18)|(1<<sgfParserT__19)|(1<<sgfParserT__20)|(1<<sgfParserT__21)|(1<<sgfParserT__22)|(1<<sgfParserT__23)|(1<<sgfParserT__24)|(1<<sgfParserT__25)|(1<<sgfParserT__26)|(1<<sgfParserT__27)|(1<<sgfParserT__28)|(1<<sgfParserT__29)|(1<<sgfParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(sgfParserT__31-32))|(1<<(sgfParserT__32-32))|(1<<(sgfParserT__33-32))|(1<<(sgfParserT__34-32))|(1<<(sgfParserT__35-32))|(1<<(sgfParserT__36-32))|(1<<(sgfParserT__37-32))|(1<<(sgfParserT__38-32))|(1<<(sgfParserT__39-32))|(1<<(sgfParserT__40-32))|(1<<(sgfParserT__41-32))|(1<<(sgfParserT__42-32))|(1<<(sgfParserT__43-32))|(1<<(sgfParserT__44-32))|(1<<(sgfParserT__45-32))|(1<<(sgfParserT__46-32))|(1<<(sgfParserT__47-32))|(1<<(sgfParserT__48-32))|(1<<(sgfParserT__49-32))|(1<<(sgfParserT__50-32))|(1<<(sgfParserT__51-32))|(1<<(sgfParserT__52-32))|(1<<(sgfParserT__53-32))|(1<<(sgfParserT__54-32))|(1<<(sgfParserT__55-32))|(1<<(sgfParserT__56-32))|(1<<(sgfParserT__57-32))|(1<<(sgfParserT__58-32))|(1<<(sgfParserT__59-32))|(1<<(sgfParserT__60-32))|(1<<(sgfParserT__61-32))|(1<<(sgfParserT__62-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(sgfParserT__63-64))|(1<<(sgfParserT__64-64))|(1<<(sgfParserT__65-64))|(1<<(sgfParserT__66-64))|(1<<(sgfParserT__67-64))|(1<<(sgfParserT__68-64))|(1<<(sgfParserT__69-64))|(1<<(sgfParserT__70-64))|(1<<(sgfParserT__71-64))|(1<<(sgfParserT__72-64))|(1<<(sgfParserCOLOR-64))|(1<<(sgfParserUCLETTER-64)))) != 0) {
		{
			p.SetState(55)
			p.Property_()
		}

		p.SetState(60)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IProperty_Context is an interface to support dynamic dispatch.
type IProperty_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProperty_Context differentiates from other interfaces.
	IsProperty_Context()
}

type Property_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProperty_Context() *Property_Context {
	var p = new(Property_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_property_
	return p
}

func (*Property_Context) IsProperty_Context() {}

func NewProperty_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Property_Context {
	var p = new(Property_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_property_

	return p
}

func (s *Property_Context) GetParser() antlr.Parser { return s.parser }

func (s *Property_Context) Move() IMoveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMoveContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMoveContext)
}

func (s *Property_Context) Setup() ISetupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetupContext)
}

func (s *Property_Context) NodeAnnotation() INodeAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INodeAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INodeAnnotationContext)
}

func (s *Property_Context) MoveAnnotation() IMoveAnnotationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMoveAnnotationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMoveAnnotationContext)
}

func (s *Property_Context) Markup() IMarkupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarkupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMarkupContext)
}

func (s *Property_Context) Root() IRootContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRootContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRootContext)
}

func (s *Property_Context) GameInfo() IGameInfoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGameInfoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGameInfoContext)
}

func (s *Property_Context) Timing() ITimingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimingContext)
}

func (s *Property_Context) Misc() IMiscContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMiscContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMiscContext)
}

func (s *Property_Context) Loa() ILoaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoaContext)
}

func (s *Property_Context) Go_() IGo_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGo_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGo_Context)
}

func (s *Property_Context) PrivateProp() IPrivatePropContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrivatePropContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrivatePropContext)
}

func (s *Property_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Property_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Property_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterProperty_(s)
	}
}

func (s *Property_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitProperty_(s)
	}
}

func (p *sgfParser) Property_() (localctx IProperty_Context) {
	this := p
	_ = this

	localctx = NewProperty_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, sgfParserRULE_property_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(73)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sgfParserT__3, sgfParserT__4, sgfParserCOLOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Move()
		}

	case sgfParserT__5, sgfParserT__6, sgfParserT__7, sgfParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.Setup()
		}

	case sgfParserT__9, sgfParserT__10, sgfParserT__11, sgfParserT__12, sgfParserT__13, sgfParserT__14, sgfParserT__15, sgfParserT__16:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.NodeAnnotation()
		}

	case sgfParserT__17, sgfParserT__18, sgfParserT__19, sgfParserT__20:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(64)
			p.MoveAnnotation()
		}

	case sgfParserT__21, sgfParserT__22, sgfParserT__23, sgfParserT__24, sgfParserT__25, sgfParserT__26, sgfParserT__27, sgfParserT__28, sgfParserT__29:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(65)
			p.Markup()
		}

	case sgfParserT__30, sgfParserT__31, sgfParserT__32, sgfParserT__33, sgfParserT__34, sgfParserT__35:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(66)
			p.Root()
		}

	case sgfParserT__36, sgfParserT__37, sgfParserT__38, sgfParserT__39, sgfParserT__40, sgfParserT__41, sgfParserT__42, sgfParserT__43, sgfParserT__44, sgfParserT__45, sgfParserT__46, sgfParserT__47, sgfParserT__48, sgfParserT__49, sgfParserT__50, sgfParserT__51, sgfParserT__52, sgfParserT__53, sgfParserT__54, sgfParserT__55, sgfParserT__56:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(67)
			p.GameInfo()
		}

	case sgfParserT__57, sgfParserT__58, sgfParserT__59, sgfParserT__60:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(68)
			p.Timing()
		}

	case sgfParserT__61, sgfParserT__62, sgfParserT__63:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(69)
			p.Misc()
		}

	case sgfParserT__64, sgfParserT__65, sgfParserT__66, sgfParserT__67, sgfParserT__68:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(70)
			p.Loa()
		}

	case sgfParserT__69, sgfParserT__70, sgfParserT__71, sgfParserT__72:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(71)
			p.Go_()
		}

	case sgfParserUCLETTER:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(72)
			p.PrivateProp()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMoveContext is an interface to support dynamic dispatch.
type IMoveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoveContext differentiates from other interfaces.
	IsMoveContext()
}

type MoveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoveContext() *MoveContext {
	var p = new(MoveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_move
	return p
}

func (*MoveContext) IsMoveContext() {}

func NewMoveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoveContext {
	var p = new(MoveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_move

	return p
}

func (s *MoveContext) GetParser() antlr.Parser { return s.parser }

func (s *MoveContext) COLOR() antlr.TerminalNode {
	return s.GetToken(sgfParserCOLOR, 0)
}

func (s *MoveContext) NONE() antlr.TerminalNode {
	return s.GetToken(sgfParserNONE, 0)
}

func (s *MoveContext) TEXT() antlr.TerminalNode {
	return s.GetToken(sgfParserTEXT, 0)
}

func (s *MoveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterMove(s)
	}
}

func (s *MoveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitMove(s)
	}
}

func (p *sgfParser) Move() (localctx IMoveContext) {
	this := p
	_ = this

	localctx = NewMoveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, sgfParserRULE_move)
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

	p.SetState(81)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sgfParserCOLOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Match(sgfParserCOLOR)
		}
		{
			p.SetState(76)
			_la = p.GetTokenStream().LA(1)

			if !(_la == sgfParserNONE || _la == sgfParserTEXT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case sgfParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(77)
			p.Match(sgfParserT__3)
		}
		{
			p.SetState(78)
			p.Match(sgfParserNONE)
		}

	case sgfParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(79)
			p.Match(sgfParserT__4)
		}
		{
			p.SetState(80)
			p.Match(sgfParserTEXT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISetupContext is an interface to support dynamic dispatch.
type ISetupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetupContext differentiates from other interfaces.
	IsSetupContext()
}

type SetupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetupContext() *SetupContext {
	var p = new(SetupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_setup
	return p
}

func (*SetupContext) IsSetupContext() {}

func NewSetupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetupContext {
	var p = new(SetupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_setup

	return p
}

func (s *SetupContext) GetParser() antlr.Parser { return s.parser }

func (s *SetupContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(sgfParserTEXT)
}

func (s *SetupContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(sgfParserTEXT, i)
}

func (s *SetupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterSetup(s)
	}
}

func (s *SetupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitSetup(s)
	}
}

func (p *sgfParser) Setup() (localctx ISetupContext) {
	this := p
	_ = this

	localctx = NewSetupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, sgfParserRULE_setup)
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

	p.SetState(103)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sgfParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.Match(sgfParserT__5)
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(84)
				p.Match(sgfParserTEXT)
			}

			p.SetState(87)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case sgfParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(89)
			p.Match(sgfParserT__6)
		}
		p.SetState(91)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(90)
				p.Match(sgfParserTEXT)
			}

			p.SetState(93)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case sgfParserT__7:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(95)
			p.Match(sgfParserT__7)
		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(96)
				p.Match(sgfParserTEXT)
			}

			p.SetState(99)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case sgfParserT__8:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(101)
			p.Match(sgfParserT__8)
		}
		{
			p.SetState(102)
			p.Match(sgfParserTEXT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INodeAnnotationContext is an interface to support dynamic dispatch.
type INodeAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNodeAnnotationContext differentiates from other interfaces.
	IsNodeAnnotationContext()
}

type NodeAnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNodeAnnotationContext() *NodeAnnotationContext {
	var p = new(NodeAnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_nodeAnnotation
	return p
}

func (*NodeAnnotationContext) IsNodeAnnotationContext() {}

func NewNodeAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NodeAnnotationContext {
	var p = new(NodeAnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_nodeAnnotation

	return p
}

func (s *NodeAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *NodeAnnotationContext) TEXT() antlr.TerminalNode {
	return s.GetToken(sgfParserTEXT, 0)
}

func (s *NodeAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NodeAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NodeAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterNodeAnnotation(s)
	}
}

func (s *NodeAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitNodeAnnotation(s)
	}
}

func (p *sgfParser) NodeAnnotation() (localctx INodeAnnotationContext) {
	this := p
	_ = this

	localctx = NewNodeAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, sgfParserRULE_nodeAnnotation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(121)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sgfParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.Match(sgfParserT__9)
		}
		{
			p.SetState(106)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(107)
			p.Match(sgfParserT__10)
		}
		{
			p.SetState(108)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__11:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.Match(sgfParserT__11)
		}
		{
			p.SetState(110)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__12:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.Match(sgfParserT__12)
		}
		{
			p.SetState(112)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__13:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(113)
			p.Match(sgfParserT__13)
		}
		{
			p.SetState(114)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__14:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(115)
			p.Match(sgfParserT__14)
		}
		{
			p.SetState(116)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__15:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(117)
			p.Match(sgfParserT__15)
		}
		{
			p.SetState(118)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__16:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(119)
			p.Match(sgfParserT__16)
		}
		{
			p.SetState(120)
			p.Match(sgfParserTEXT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMoveAnnotationContext is an interface to support dynamic dispatch.
type IMoveAnnotationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMoveAnnotationContext differentiates from other interfaces.
	IsMoveAnnotationContext()
}

type MoveAnnotationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMoveAnnotationContext() *MoveAnnotationContext {
	var p = new(MoveAnnotationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_moveAnnotation
	return p
}

func (*MoveAnnotationContext) IsMoveAnnotationContext() {}

func NewMoveAnnotationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MoveAnnotationContext {
	var p = new(MoveAnnotationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_moveAnnotation

	return p
}

func (s *MoveAnnotationContext) GetParser() antlr.Parser { return s.parser }

func (s *MoveAnnotationContext) TEXT() antlr.TerminalNode {
	return s.GetToken(sgfParserTEXT, 0)
}

func (s *MoveAnnotationContext) NONE() antlr.TerminalNode {
	return s.GetToken(sgfParserNONE, 0)
}

func (s *MoveAnnotationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MoveAnnotationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MoveAnnotationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterMoveAnnotation(s)
	}
}

func (s *MoveAnnotationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitMoveAnnotation(s)
	}
}

func (p *sgfParser) MoveAnnotation() (localctx IMoveAnnotationContext) {
	this := p
	_ = this

	localctx = NewMoveAnnotationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, sgfParserRULE_moveAnnotation)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sgfParserT__17:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Match(sgfParserT__17)
		}
		{
			p.SetState(124)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__18:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Match(sgfParserT__18)
		}
		{
			p.SetState(126)
			p.Match(sgfParserNONE)
		}

	case sgfParserT__19:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(127)
			p.Match(sgfParserT__19)
		}
		{
			p.SetState(128)
			p.Match(sgfParserNONE)
		}

	case sgfParserT__20:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(129)
			p.Match(sgfParserT__20)
		}
		{
			p.SetState(130)
			p.Match(sgfParserTEXT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMarkupContext is an interface to support dynamic dispatch.
type IMarkupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMarkupContext differentiates from other interfaces.
	IsMarkupContext()
}

type MarkupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMarkupContext() *MarkupContext {
	var p = new(MarkupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_markup
	return p
}

func (*MarkupContext) IsMarkupContext() {}

func NewMarkupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MarkupContext {
	var p = new(MarkupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_markup

	return p
}

func (s *MarkupContext) GetParser() antlr.Parser { return s.parser }

func (s *MarkupContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(sgfParserTEXT)
}

func (s *MarkupContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(sgfParserTEXT, i)
}

func (s *MarkupContext) NONE() antlr.TerminalNode {
	return s.GetToken(sgfParserNONE, 0)
}

func (s *MarkupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarkupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MarkupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterMarkup(s)
	}
}

func (s *MarkupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitMarkup(s)
	}
}

func (p *sgfParser) Markup() (localctx IMarkupContext) {
	this := p
	_ = this

	localctx = NewMarkupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, sgfParserRULE_markup)
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

	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sgfParserT__21:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(133)
			p.Match(sgfParserT__21)
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(134)
				p.Match(sgfParserTEXT)
			}

			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case sgfParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.Match(sgfParserT__22)
		}
		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(140)
				p.Match(sgfParserTEXT)
			}

			p.SetState(143)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case sgfParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(145)
			p.Match(sgfParserT__23)
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case sgfParserNONE:
			{
				p.SetState(146)
				p.Match(sgfParserNONE)
			}

		case sgfParserTEXT:
			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == sgfParserTEXT {
				{
					p.SetState(147)
					p.Match(sgfParserTEXT)
				}

				p.SetState(150)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case sgfParserT__24:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(154)
			p.Match(sgfParserT__24)
		}
		p.SetState(156)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(155)
				p.Match(sgfParserTEXT)
			}

			p.SetState(158)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case sgfParserT__25:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(160)
			p.Match(sgfParserT__25)
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(161)
				p.Match(sgfParserTEXT)
			}

			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case sgfParserT__26:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(166)
			p.Match(sgfParserT__26)
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(167)
				p.Match(sgfParserTEXT)
			}

			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case sgfParserT__27:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(172)
			p.Match(sgfParserT__27)
		}
		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(173)
				p.Match(sgfParserTEXT)
			}

			p.SetState(176)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case sgfParserT__28:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(178)
			p.Match(sgfParserT__28)
		}
		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(179)
				p.Match(sgfParserTEXT)
			}

			p.SetState(182)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case sgfParserT__29:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(184)
			p.Match(sgfParserT__29)
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(185)
				p.Match(sgfParserTEXT)
			}

			p.SetState(188)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) TEXT() antlr.TerminalNode {
	return s.GetToken(sgfParserTEXT, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *sgfParser) Root() (localctx IRootContext) {
	this := p
	_ = this

	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, sgfParserRULE_root)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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

	switch p.GetTokenStream().LA(1) {
	case sgfParserT__30:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(192)
			p.Match(sgfParserT__30)
		}
		{
			p.SetState(193)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__31:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.Match(sgfParserT__31)
		}
		{
			p.SetState(195)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__32:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(196)
			p.Match(sgfParserT__32)
		}
		{
			p.SetState(197)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__33:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(198)
			p.Match(sgfParserT__33)
		}
		{
			p.SetState(199)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__34:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(200)
			p.Match(sgfParserT__34)
		}
		{
			p.SetState(201)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__35:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(202)
			p.Match(sgfParserT__35)
		}
		{
			p.SetState(203)
			p.Match(sgfParserTEXT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGameInfoContext is an interface to support dynamic dispatch.
type IGameInfoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGameInfoContext differentiates from other interfaces.
	IsGameInfoContext()
}

type GameInfoContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGameInfoContext() *GameInfoContext {
	var p = new(GameInfoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_gameInfo
	return p
}

func (*GameInfoContext) IsGameInfoContext() {}

func NewGameInfoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GameInfoContext {
	var p = new(GameInfoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_gameInfo

	return p
}

func (s *GameInfoContext) GetParser() antlr.Parser { return s.parser }

func (s *GameInfoContext) TEXT() antlr.TerminalNode {
	return s.GetToken(sgfParserTEXT, 0)
}

func (s *GameInfoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GameInfoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GameInfoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterGameInfo(s)
	}
}

func (s *GameInfoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitGameInfo(s)
	}
}

func (p *sgfParser) GameInfo() (localctx IGameInfoContext) {
	this := p
	_ = this

	localctx = NewGameInfoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, sgfParserRULE_gameInfo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(248)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sgfParserT__36:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.Match(sgfParserT__36)
		}
		{
			p.SetState(207)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__37:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.Match(sgfParserT__37)
		}
		{
			p.SetState(209)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__38:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(210)
			p.Match(sgfParserT__38)
		}
		{
			p.SetState(211)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__39:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(212)
			p.Match(sgfParserT__39)
		}
		{
			p.SetState(213)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__40:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(214)
			p.Match(sgfParserT__40)
		}
		{
			p.SetState(215)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__41:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(216)
			p.Match(sgfParserT__41)
		}
		{
			p.SetState(217)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__42:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(218)
			p.Match(sgfParserT__42)
		}
		{
			p.SetState(219)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__43:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(220)
			p.Match(sgfParserT__43)
		}
		{
			p.SetState(221)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__44:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(222)
			p.Match(sgfParserT__44)
		}
		{
			p.SetState(223)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__45:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(224)
			p.Match(sgfParserT__45)
		}
		{
			p.SetState(225)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__46:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(226)
			p.Match(sgfParserT__46)
		}
		{
			p.SetState(227)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__47:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(228)
			p.Match(sgfParserT__47)
		}
		{
			p.SetState(229)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__48:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(230)
			p.Match(sgfParserT__48)
		}
		{
			p.SetState(231)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__49:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(232)
			p.Match(sgfParserT__49)
		}
		{
			p.SetState(233)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__50:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(234)
			p.Match(sgfParserT__50)
		}
		{
			p.SetState(235)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__51:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(236)
			p.Match(sgfParserT__51)
		}
		{
			p.SetState(237)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__52:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(238)
			p.Match(sgfParserT__52)
		}
		{
			p.SetState(239)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__53:
		p.EnterOuterAlt(localctx, 18)
		{
			p.SetState(240)
			p.Match(sgfParserT__53)
		}
		{
			p.SetState(241)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__54:
		p.EnterOuterAlt(localctx, 19)
		{
			p.SetState(242)
			p.Match(sgfParserT__54)
		}
		{
			p.SetState(243)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__55:
		p.EnterOuterAlt(localctx, 20)
		{
			p.SetState(244)
			p.Match(sgfParserT__55)
		}
		{
			p.SetState(245)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__56:
		p.EnterOuterAlt(localctx, 21)
		{
			p.SetState(246)
			p.Match(sgfParserT__56)
		}
		{
			p.SetState(247)
			p.Match(sgfParserTEXT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITimingContext is an interface to support dynamic dispatch.
type ITimingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimingContext differentiates from other interfaces.
	IsTimingContext()
}

type TimingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimingContext() *TimingContext {
	var p = new(TimingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_timing
	return p
}

func (*TimingContext) IsTimingContext() {}

func NewTimingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimingContext {
	var p = new(TimingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_timing

	return p
}

func (s *TimingContext) GetParser() antlr.Parser { return s.parser }

func (s *TimingContext) TEXT() antlr.TerminalNode {
	return s.GetToken(sgfParserTEXT, 0)
}

func (s *TimingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterTiming(s)
	}
}

func (s *TimingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitTiming(s)
	}
}

func (p *sgfParser) Timing() (localctx ITimingContext) {
	this := p
	_ = this

	localctx = NewTimingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, sgfParserRULE_timing)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(258)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sgfParserT__57:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(250)
			p.Match(sgfParserT__57)
		}
		{
			p.SetState(251)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__58:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(252)
			p.Match(sgfParserT__58)
		}
		{
			p.SetState(253)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__59:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(254)
			p.Match(sgfParserT__59)
		}
		{
			p.SetState(255)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__60:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(256)
			p.Match(sgfParserT__60)
		}
		{
			p.SetState(257)
			p.Match(sgfParserTEXT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMiscContext is an interface to support dynamic dispatch.
type IMiscContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMiscContext differentiates from other interfaces.
	IsMiscContext()
}

type MiscContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMiscContext() *MiscContext {
	var p = new(MiscContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_misc
	return p
}

func (*MiscContext) IsMiscContext() {}

func NewMiscContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MiscContext {
	var p = new(MiscContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_misc

	return p
}

func (s *MiscContext) GetParser() antlr.Parser { return s.parser }

func (s *MiscContext) NONE() antlr.TerminalNode {
	return s.GetToken(sgfParserNONE, 0)
}

func (s *MiscContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(sgfParserTEXT)
}

func (s *MiscContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(sgfParserTEXT, i)
}

func (s *MiscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MiscContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MiscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterMisc(s)
	}
}

func (s *MiscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitMisc(s)
	}
}

func (p *sgfParser) Misc() (localctx IMiscContext) {
	this := p
	_ = this

	localctx = NewMiscContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, sgfParserRULE_misc)
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

	p.SetState(270)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sgfParserT__61:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(260)
			p.Match(sgfParserT__61)
		}
		{
			p.SetState(261)
			_la = p.GetTokenStream().LA(1)

			if !(_la == sgfParserNONE || _la == sgfParserTEXT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case sgfParserT__62:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(262)
			p.Match(sgfParserT__62)
		}
		{
			p.SetState(263)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__63:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(264)
			p.Match(sgfParserT__63)
		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(265)
				p.Match(sgfParserTEXT)
			}

			p.SetState(268)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILoaContext is an interface to support dynamic dispatch.
type ILoaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoaContext differentiates from other interfaces.
	IsLoaContext()
}

type LoaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoaContext() *LoaContext {
	var p = new(LoaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_loa
	return p
}

func (*LoaContext) IsLoaContext() {}

func NewLoaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoaContext {
	var p = new(LoaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_loa

	return p
}

func (s *LoaContext) GetParser() antlr.Parser { return s.parser }

func (s *LoaContext) TEXT() antlr.TerminalNode {
	return s.GetToken(sgfParserTEXT, 0)
}

func (s *LoaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterLoa(s)
	}
}

func (s *LoaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitLoa(s)
	}
}

func (p *sgfParser) Loa() (localctx ILoaContext) {
	this := p
	_ = this

	localctx = NewLoaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, sgfParserRULE_loa)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	case sgfParserT__64:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(272)
			p.Match(sgfParserT__64)
		}
		{
			p.SetState(273)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__65:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(274)
			p.Match(sgfParserT__65)
		}
		{
			p.SetState(275)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__66:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(276)
			p.Match(sgfParserT__66)
		}
		{
			p.SetState(277)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__67:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(278)
			p.Match(sgfParserT__67)
		}
		{
			p.SetState(279)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__68:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(280)
			p.Match(sgfParserT__68)
		}
		{
			p.SetState(281)
			p.Match(sgfParserTEXT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGo_Context is an interface to support dynamic dispatch.
type IGo_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGo_Context differentiates from other interfaces.
	IsGo_Context()
}

type Go_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGo_Context() *Go_Context {
	var p = new(Go_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_go_
	return p
}

func (*Go_Context) IsGo_Context() {}

func NewGo_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Go_Context {
	var p = new(Go_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_go_

	return p
}

func (s *Go_Context) GetParser() antlr.Parser { return s.parser }

func (s *Go_Context) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(sgfParserTEXT)
}

func (s *Go_Context) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(sgfParserTEXT, i)
}

func (s *Go_Context) NONE() antlr.TerminalNode {
	return s.GetToken(sgfParserNONE, 0)
}

func (s *Go_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Go_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Go_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterGo_(s)
	}
}

func (s *Go_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitGo_(s)
	}
}

func (p *sgfParser) Go_() (localctx IGo_Context) {
	this := p
	_ = this

	localctx = NewGo_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, sgfParserRULE_go_)
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

	p.SetState(306)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sgfParserT__69:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(284)
			p.Match(sgfParserT__69)
		}
		{
			p.SetState(285)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__70:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.Match(sgfParserT__70)
		}
		{
			p.SetState(287)
			p.Match(sgfParserTEXT)
		}

	case sgfParserT__71:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(288)
			p.Match(sgfParserT__71)
		}
		p.SetState(295)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case sgfParserNONE:
			{
				p.SetState(289)
				p.Match(sgfParserNONE)
			}

		case sgfParserTEXT:
			p.SetState(291)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == sgfParserTEXT {
				{
					p.SetState(290)
					p.Match(sgfParserTEXT)
				}

				p.SetState(293)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case sgfParserT__72:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(297)
			p.Match(sgfParserT__72)
		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case sgfParserNONE:
			{
				p.SetState(298)
				p.Match(sgfParserNONE)
			}

		case sgfParserTEXT:
			p.SetState(300)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == sgfParserTEXT {
				{
					p.SetState(299)
					p.Match(sgfParserTEXT)
				}

				p.SetState(302)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPrivatePropContext is an interface to support dynamic dispatch.
type IPrivatePropContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrivatePropContext differentiates from other interfaces.
	IsPrivatePropContext()
}

type PrivatePropContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrivatePropContext() *PrivatePropContext {
	var p = new(PrivatePropContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sgfParserRULE_privateProp
	return p
}

func (*PrivatePropContext) IsPrivatePropContext() {}

func NewPrivatePropContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrivatePropContext {
	var p = new(PrivatePropContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sgfParserRULE_privateProp

	return p
}

func (s *PrivatePropContext) GetParser() antlr.Parser { return s.parser }

func (s *PrivatePropContext) UCLETTER() antlr.TerminalNode {
	return s.GetToken(sgfParserUCLETTER, 0)
}

func (s *PrivatePropContext) NONE() antlr.TerminalNode {
	return s.GetToken(sgfParserNONE, 0)
}

func (s *PrivatePropContext) AllTEXT() []antlr.TerminalNode {
	return s.GetTokens(sgfParserTEXT)
}

func (s *PrivatePropContext) TEXT(i int) antlr.TerminalNode {
	return s.GetToken(sgfParserTEXT, i)
}

func (s *PrivatePropContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrivatePropContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrivatePropContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.EnterPrivateProp(s)
	}
}

func (s *PrivatePropContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sgfListener); ok {
		listenerT.ExitPrivateProp(s)
	}
}

func (p *sgfParser) PrivateProp() (localctx IPrivatePropContext) {
	this := p
	_ = this

	localctx = NewPrivatePropContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, sgfParserRULE_privateProp)
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
		p.SetState(308)
		p.Match(sgfParserUCLETTER)
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sgfParserNONE:
		{
			p.SetState(309)
			p.Match(sgfParserNONE)
		}

	case sgfParserTEXT:
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == sgfParserTEXT {
			{
				p.SetState(310)
				p.Match(sgfParserTEXT)
			}

			p.SetState(313)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
