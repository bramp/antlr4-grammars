// Code generated from snowball.g4 by ANTLR 4.9.3. DO NOT EDIT.

package snowball // snowball
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 79, 349,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 7,
	2, 36, 10, 2, 12, 2, 14, 2, 39, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 49, 10, 3, 3, 4, 3, 4, 3, 4, 7, 4, 54, 10, 4, 12,
	4, 14, 4, 57, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 63, 10, 4, 12, 4, 14,
	4, 66, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 72, 10, 4, 12, 4, 14, 4, 75,
	11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 81, 10, 4, 12, 4, 14, 4, 84, 11, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 90, 10, 4, 12, 4, 14, 4, 93, 11, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 7, 4, 99, 10, 4, 12, 4, 14, 4, 102, 11, 4, 3, 4, 5, 4,
	105, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	7, 6, 117, 10, 6, 12, 6, 14, 6, 120, 11, 6, 3, 7, 3, 7, 3, 7, 7, 7, 125,
	10, 7, 12, 7, 14, 7, 128, 11, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 5, 7, 198, 10, 7, 3, 7, 7, 7, 201, 10, 7, 12, 7, 14, 7, 204, 11,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 216,
	10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 222, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 7, 7, 230, 10, 7, 12, 7, 14, 7, 233, 11, 7, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 290, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 5, 10, 298, 10, 10, 3, 11, 3, 11, 5, 11, 302, 10, 11, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 330, 10, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 344,
	10, 17, 12, 17, 14, 17, 347, 11, 17, 3, 17, 2, 4, 12, 32, 18, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 2, 2, 423, 2, 37,
	3, 2, 2, 2, 4, 48, 3, 2, 2, 2, 6, 104, 3, 2, 2, 2, 8, 106, 3, 2, 2, 2,
	10, 111, 3, 2, 2, 2, 12, 221, 3, 2, 2, 2, 14, 289, 3, 2, 2, 2, 16, 291,
	3, 2, 2, 2, 18, 297, 3, 2, 2, 2, 20, 301, 3, 2, 2, 2, 22, 303, 3, 2, 2,
	2, 24, 305, 3, 2, 2, 2, 26, 307, 3, 2, 2, 2, 28, 309, 3, 2, 2, 2, 30, 311,
	3, 2, 2, 2, 32, 329, 3, 2, 2, 2, 34, 36, 5, 4, 3, 2, 35, 34, 3, 2, 2, 2,
	36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 3, 3, 2,
	2, 2, 39, 37, 3, 2, 2, 2, 40, 49, 5, 6, 4, 2, 41, 49, 5, 8, 5, 2, 42, 49,
	5, 10, 6, 2, 43, 44, 7, 3, 2, 2, 44, 45, 7, 4, 2, 2, 45, 46, 5, 4, 3, 2,
	46, 47, 7, 5, 2, 2, 47, 49, 3, 2, 2, 2, 48, 40, 3, 2, 2, 2, 48, 41, 3,
	2, 2, 2, 48, 42, 3, 2, 2, 2, 48, 43, 3, 2, 2, 2, 49, 5, 3, 2, 2, 2, 50,
	51, 7, 6, 2, 2, 51, 55, 7, 4, 2, 2, 52, 54, 5, 22, 12, 2, 53, 52, 3, 2,
	2, 2, 54, 57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 58,
	3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 58, 105, 7, 5, 2, 2, 59, 60, 7, 7, 2, 2,
	60, 64, 7, 4, 2, 2, 61, 63, 5, 24, 13, 2, 62, 61, 3, 2, 2, 2, 63, 66, 3,
	2, 2, 2, 64, 62, 3, 2, 2, 2, 64, 65, 3, 2, 2, 2, 65, 67, 3, 2, 2, 2, 66,
	64, 3, 2, 2, 2, 67, 105, 7, 5, 2, 2, 68, 69, 7, 8, 2, 2, 69, 73, 7, 4,
	2, 2, 70, 72, 5, 26, 14, 2, 71, 70, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 73,
	71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 73, 3, 2, 2,
	2, 76, 105, 7, 5, 2, 2, 77, 78, 7, 9, 2, 2, 78, 82, 7, 4, 2, 2, 79, 81,
	5, 28, 15, 2, 80, 79, 3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80, 3, 2, 2,
	2, 82, 83, 3, 2, 2, 2, 83, 85, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 85, 105,
	7, 5, 2, 2, 86, 87, 7, 10, 2, 2, 87, 91, 7, 4, 2, 2, 88, 90, 5, 28, 15,
	2, 89, 88, 3, 2, 2, 2, 90, 93, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 91, 92,
	3, 2, 2, 2, 92, 94, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 94, 105, 7, 5, 2, 2,
	95, 96, 7, 11, 2, 2, 96, 100, 7, 4, 2, 2, 97, 99, 5, 30, 16, 2, 98, 97,
	3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2,
	2, 101, 103, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 105, 7, 5, 2, 2, 104,
	50, 3, 2, 2, 2, 104, 59, 3, 2, 2, 2, 104, 68, 3, 2, 2, 2, 104, 77, 3, 2,
	2, 2, 104, 86, 3, 2, 2, 2, 104, 95, 3, 2, 2, 2, 105, 7, 3, 2, 2, 2, 106,
	107, 7, 12, 2, 2, 107, 108, 5, 28, 15, 2, 108, 109, 7, 13, 2, 2, 109, 110,
	5, 12, 7, 2, 110, 9, 3, 2, 2, 2, 111, 112, 7, 12, 2, 2, 112, 113, 5, 30,
	16, 2, 113, 118, 5, 20, 11, 2, 114, 115, 7, 77, 2, 2, 115, 117, 5, 20,
	11, 2, 116, 114, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2,
	118, 119, 3, 2, 2, 2, 119, 11, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 122,
	8, 7, 1, 2, 122, 126, 7, 4, 2, 2, 123, 125, 5, 12, 7, 2, 124, 123, 3, 2,
	2, 2, 125, 128, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2,
	127, 129, 3, 2, 2, 2, 128, 126, 3, 2, 2, 2, 129, 222, 7, 5, 2, 2, 130,
	222, 5, 14, 8, 2, 131, 222, 5, 16, 9, 2, 132, 133, 7, 16, 2, 2, 133, 222,
	5, 12, 7, 43, 134, 135, 7, 17, 2, 2, 135, 222, 5, 12, 7, 42, 136, 137,
	7, 18, 2, 2, 137, 222, 5, 12, 7, 41, 138, 139, 7, 19, 2, 2, 139, 222, 5,
	12, 7, 40, 140, 141, 7, 20, 2, 2, 141, 222, 5, 12, 7, 39, 142, 143, 7,
	21, 2, 2, 143, 222, 5, 12, 7, 38, 144, 145, 7, 22, 2, 2, 145, 222, 5, 12,
	7, 37, 146, 147, 7, 23, 2, 2, 147, 222, 5, 12, 7, 36, 148, 149, 7, 24,
	2, 2, 149, 150, 5, 32, 17, 2, 150, 151, 5, 12, 7, 35, 151, 222, 3, 2, 2,
	2, 152, 153, 7, 25, 2, 2, 153, 154, 5, 32, 17, 2, 154, 155, 5, 12, 7, 34,
	155, 222, 3, 2, 2, 2, 156, 222, 5, 18, 10, 2, 157, 158, 7, 26, 2, 2, 158,
	222, 5, 18, 10, 2, 159, 160, 7, 27, 2, 2, 160, 222, 5, 18, 10, 2, 161,
	162, 7, 28, 2, 2, 162, 222, 5, 18, 10, 2, 163, 164, 7, 29, 2, 2, 164, 222,
	5, 18, 10, 2, 165, 222, 7, 30, 2, 2, 166, 167, 7, 31, 2, 2, 167, 222, 5,
	32, 17, 2, 168, 222, 7, 32, 2, 2, 169, 170, 7, 33, 2, 2, 170, 222, 5, 22,
	12, 2, 171, 222, 7, 34, 2, 2, 172, 222, 7, 35, 2, 2, 173, 174, 7, 36, 2,
	2, 174, 222, 5, 22, 12, 2, 175, 176, 7, 37, 2, 2, 176, 222, 5, 24, 13,
	2, 177, 178, 7, 38, 2, 2, 178, 222, 5, 32, 17, 2, 179, 180, 7, 39, 2, 2,
	180, 222, 5, 32, 17, 2, 181, 222, 7, 40, 2, 2, 182, 222, 7, 41, 2, 2, 183,
	184, 7, 42, 2, 2, 184, 185, 5, 12, 7, 2, 185, 186, 7, 43, 2, 2, 186, 187,
	5, 12, 7, 16, 187, 222, 3, 2, 2, 2, 188, 189, 7, 44, 2, 2, 189, 222, 5,
	12, 7, 15, 190, 191, 7, 45, 2, 2, 191, 222, 5, 12, 7, 14, 192, 222, 7,
	46, 2, 2, 193, 194, 7, 47, 2, 2, 194, 202, 7, 4, 2, 2, 195, 197, 7, 75,
	2, 2, 196, 198, 5, 28, 15, 2, 197, 196, 3, 2, 2, 2, 197, 198, 3, 2, 2,
	2, 198, 201, 3, 2, 2, 2, 199, 201, 5, 12, 7, 2, 200, 195, 3, 2, 2, 2, 200,
	199, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 202, 203,
	3, 2, 2, 2, 203, 205, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 205, 222, 7, 5,
	2, 2, 206, 207, 7, 48, 2, 2, 207, 222, 5, 26, 14, 2, 208, 209, 7, 49, 2,
	2, 209, 222, 5, 26, 14, 2, 210, 222, 5, 26, 14, 2, 211, 222, 5, 28, 15,
	2, 212, 222, 5, 30, 16, 2, 213, 215, 7, 50, 2, 2, 214, 216, 7, 51, 2, 2,
	215, 214, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217,
	222, 5, 30, 16, 2, 218, 222, 7, 52, 2, 2, 219, 222, 7, 53, 2, 2, 220, 222,
	7, 54, 2, 2, 221, 121, 3, 2, 2, 2, 221, 130, 3, 2, 2, 2, 221, 131, 3, 2,
	2, 2, 221, 132, 3, 2, 2, 2, 221, 134, 3, 2, 2, 2, 221, 136, 3, 2, 2, 2,
	221, 138, 3, 2, 2, 2, 221, 140, 3, 2, 2, 2, 221, 142, 3, 2, 2, 2, 221,
	144, 3, 2, 2, 2, 221, 146, 3, 2, 2, 2, 221, 148, 3, 2, 2, 2, 221, 152,
	3, 2, 2, 2, 221, 156, 3, 2, 2, 2, 221, 157, 3, 2, 2, 2, 221, 159, 3, 2,
	2, 2, 221, 161, 3, 2, 2, 2, 221, 163, 3, 2, 2, 2, 221, 165, 3, 2, 2, 2,
	221, 166, 3, 2, 2, 2, 221, 168, 3, 2, 2, 2, 221, 169, 3, 2, 2, 2, 221,
	171, 3, 2, 2, 2, 221, 172, 3, 2, 2, 2, 221, 173, 3, 2, 2, 2, 221, 175,
	3, 2, 2, 2, 221, 177, 3, 2, 2, 2, 221, 179, 3, 2, 2, 2, 221, 181, 3, 2,
	2, 2, 221, 182, 3, 2, 2, 2, 221, 183, 3, 2, 2, 2, 221, 188, 3, 2, 2, 2,
	221, 190, 3, 2, 2, 2, 221, 192, 3, 2, 2, 2, 221, 193, 3, 2, 2, 2, 221,
	206, 3, 2, 2, 2, 221, 208, 3, 2, 2, 2, 221, 210, 3, 2, 2, 2, 221, 211,
	3, 2, 2, 2, 221, 212, 3, 2, 2, 2, 221, 213, 3, 2, 2, 2, 221, 218, 3, 2,
	2, 2, 221, 219, 3, 2, 2, 2, 221, 220, 3, 2, 2, 2, 222, 231, 3, 2, 2, 2,
	223, 224, 12, 45, 2, 2, 224, 225, 7, 14, 2, 2, 225, 230, 5, 12, 7, 46,
	226, 227, 12, 44, 2, 2, 227, 228, 7, 15, 2, 2, 228, 230, 5, 12, 7, 45,
	229, 223, 3, 2, 2, 2, 229, 226, 3, 2, 2, 2, 230, 233, 3, 2, 2, 2, 231,
	229, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 13, 3, 2, 2, 2, 233, 231, 3,
	2, 2, 2, 234, 235, 7, 55, 2, 2, 235, 236, 5, 24, 13, 2, 236, 237, 7, 26,
	2, 2, 237, 238, 5, 32, 17, 2, 238, 290, 3, 2, 2, 2, 239, 240, 7, 55, 2,
	2, 240, 241, 5, 24, 13, 2, 241, 242, 7, 56, 2, 2, 242, 243, 5, 32, 17,
	2, 243, 290, 3, 2, 2, 2, 244, 245, 7, 55, 2, 2, 245, 246, 5, 24, 13, 2,
	246, 247, 7, 57, 2, 2, 247, 248, 5, 32, 17, 2, 248, 290, 3, 2, 2, 2, 249,
	250, 7, 55, 2, 2, 250, 251, 5, 24, 13, 2, 251, 252, 7, 58, 2, 2, 252, 253,
	5, 32, 17, 2, 253, 290, 3, 2, 2, 2, 254, 255, 7, 55, 2, 2, 255, 256, 5,
	24, 13, 2, 256, 257, 7, 59, 2, 2, 257, 258, 5, 32, 17, 2, 258, 290, 3,
	2, 2, 2, 259, 260, 7, 55, 2, 2, 260, 261, 5, 24, 13, 2, 261, 262, 7, 60,
	2, 2, 262, 263, 5, 32, 17, 2, 263, 290, 3, 2, 2, 2, 264, 265, 7, 55, 2,
	2, 265, 266, 5, 24, 13, 2, 266, 267, 7, 61, 2, 2, 267, 268, 5, 32, 17,
	2, 268, 290, 3, 2, 2, 2, 269, 270, 7, 55, 2, 2, 270, 271, 5, 24, 13, 2,
	271, 272, 7, 62, 2, 2, 272, 273, 5, 32, 17, 2, 273, 290, 3, 2, 2, 2, 274,
	275, 7, 55, 2, 2, 275, 276, 5, 24, 13, 2, 276, 277, 7, 63, 2, 2, 277, 278,
	5, 32, 17, 2, 278, 290, 3, 2, 2, 2, 279, 280, 7, 55, 2, 2, 280, 281, 5,
	24, 13, 2, 281, 282, 7, 64, 2, 2, 282, 283, 5, 32, 17, 2, 283, 290, 3,
	2, 2, 2, 284, 285, 7, 55, 2, 2, 285, 286, 5, 24, 13, 2, 286, 287, 7, 65,
	2, 2, 287, 288, 5, 32, 17, 2, 288, 290, 3, 2, 2, 2, 289, 234, 3, 2, 2,
	2, 289, 239, 3, 2, 2, 2, 289, 244, 3, 2, 2, 2, 289, 249, 3, 2, 2, 2, 289,
	254, 3, 2, 2, 2, 289, 259, 3, 2, 2, 2, 289, 264, 3, 2, 2, 2, 289, 269,
	3, 2, 2, 2, 289, 274, 3, 2, 2, 2, 289, 279, 3, 2, 2, 2, 289, 284, 3, 2,
	2, 2, 290, 15, 3, 2, 2, 2, 291, 292, 7, 55, 2, 2, 292, 293, 5, 22, 12,
	2, 293, 294, 5, 12, 7, 2, 294, 17, 3, 2, 2, 2, 295, 298, 5, 22, 12, 2,
	296, 298, 7, 75, 2, 2, 297, 295, 3, 2, 2, 2, 297, 296, 3, 2, 2, 2, 298,
	19, 3, 2, 2, 2, 299, 302, 5, 30, 16, 2, 300, 302, 7, 75, 2, 2, 301, 299,
	3, 2, 2, 2, 301, 300, 3, 2, 2, 2, 302, 21, 3, 2, 2, 2, 303, 304, 7, 78,
	2, 2, 304, 23, 3, 2, 2, 2, 305, 306, 7, 78, 2, 2, 306, 25, 3, 2, 2, 2,
	307, 308, 7, 78, 2, 2, 308, 27, 3, 2, 2, 2, 309, 310, 7, 78, 2, 2, 310,
	29, 3, 2, 2, 2, 311, 312, 7, 78, 2, 2, 312, 31, 3, 2, 2, 2, 313, 314, 8,
	17, 1, 2, 314, 315, 7, 4, 2, 2, 315, 316, 5, 32, 17, 2, 316, 317, 7, 5,
	2, 2, 317, 330, 3, 2, 2, 2, 318, 319, 7, 51, 2, 2, 319, 330, 5, 32, 17,
	11, 320, 330, 7, 69, 2, 2, 321, 330, 7, 70, 2, 2, 322, 330, 7, 71, 2, 2,
	323, 330, 7, 72, 2, 2, 324, 330, 7, 73, 2, 2, 325, 326, 7, 74, 2, 2, 326,
	330, 5, 22, 12, 2, 327, 330, 5, 24, 13, 2, 328, 330, 7, 76, 2, 2, 329,
	313, 3, 2, 2, 2, 329, 318, 3, 2, 2, 2, 329, 320, 3, 2, 2, 2, 329, 321,
	3, 2, 2, 2, 329, 322, 3, 2, 2, 2, 329, 323, 3, 2, 2, 2, 329, 324, 3, 2,
	2, 2, 329, 325, 3, 2, 2, 2, 329, 327, 3, 2, 2, 2, 329, 328, 3, 2, 2, 2,
	330, 345, 3, 2, 2, 2, 331, 332, 12, 15, 2, 2, 332, 333, 7, 66, 2, 2, 333,
	344, 5, 32, 17, 16, 334, 335, 12, 14, 2, 2, 335, 336, 7, 51, 2, 2, 336,
	344, 5, 32, 17, 15, 337, 338, 12, 13, 2, 2, 338, 339, 7, 67, 2, 2, 339,
	344, 5, 32, 17, 14, 340, 341, 12, 12, 2, 2, 341, 342, 7, 68, 2, 2, 342,
	344, 5, 32, 17, 13, 343, 331, 3, 2, 2, 2, 343, 334, 3, 2, 2, 2, 343, 337,
	3, 2, 2, 2, 343, 340, 3, 2, 2, 2, 344, 347, 3, 2, 2, 2, 345, 343, 3, 2,
	2, 2, 345, 346, 3, 2, 2, 2, 346, 33, 3, 2, 2, 2, 347, 345, 3, 2, 2, 2,
	26, 37, 48, 55, 64, 73, 82, 91, 100, 104, 118, 126, 197, 200, 202, 215,
	221, 229, 231, 289, 297, 301, 329, 343, 345,
}
var literalNames = []string{
	"", "'backwardmode'", "'('", "')'", "'strings'", "'integers'", "'booleans'",
	"'routines'", "'externals'", "'groupings'", "'define'", "'as'", "'or'",
	"'and'", "'not'", "'test'", "'try'", "'do'", "'fail'", "'goto'", "'gopast'",
	"'repeat'", "'loop'", "'atleast'", "'='", "'insert'", "'attach'", "'<-'",
	"'delete'", "'hop'", "'next'", "'=>'", "'['", "']'", "'->'", "'setmark'",
	"'tomark'", "'atmark'", "'tolimit'", "'atlimit'", "'setlimit'", "'for'",
	"'backwards'", "'reverse'", "'substring'", "'among'", "'set'", "'unset'",
	"'non'", "'-'", "'true'", "'false'", "'?'", "'$'", "'+='", "'-='", "'*='",
	"'/='", "'=='", "'!='", "'>'", "'>='", "'<'", "'<='", "'+'", "'*'", "'/'",
	"'maxint'", "'minint'", "'cursor'", "'limit'", "'size'", "'sizeof'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "LITERAL_STRING", "NUMBER", "PLUS_OR_MINUS", "NAME", "WS",
}

var ruleNames = []string{
	"program", "p", "declaration", "r_definition", "g_definition", "c", "i_command",
	"s_command", "s", "g", "s_name", "i_name", "b_name", "r_name", "g_name",
	"ae",
}

type snowballParser struct {
	*antlr.BaseParser
}

// NewsnowballParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *snowballParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewsnowballParser(input antlr.TokenStream) *snowballParser {
	this := new(snowballParser)
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
	this.GrammarFileName = "snowball.g4"

	return this
}

// snowballParser tokens.
const (
	snowballParserEOF            = antlr.TokenEOF
	snowballParserT__0           = 1
	snowballParserT__1           = 2
	snowballParserT__2           = 3
	snowballParserT__3           = 4
	snowballParserT__4           = 5
	snowballParserT__5           = 6
	snowballParserT__6           = 7
	snowballParserT__7           = 8
	snowballParserT__8           = 9
	snowballParserT__9           = 10
	snowballParserT__10          = 11
	snowballParserT__11          = 12
	snowballParserT__12          = 13
	snowballParserT__13          = 14
	snowballParserT__14          = 15
	snowballParserT__15          = 16
	snowballParserT__16          = 17
	snowballParserT__17          = 18
	snowballParserT__18          = 19
	snowballParserT__19          = 20
	snowballParserT__20          = 21
	snowballParserT__21          = 22
	snowballParserT__22          = 23
	snowballParserT__23          = 24
	snowballParserT__24          = 25
	snowballParserT__25          = 26
	snowballParserT__26          = 27
	snowballParserT__27          = 28
	snowballParserT__28          = 29
	snowballParserT__29          = 30
	snowballParserT__30          = 31
	snowballParserT__31          = 32
	snowballParserT__32          = 33
	snowballParserT__33          = 34
	snowballParserT__34          = 35
	snowballParserT__35          = 36
	snowballParserT__36          = 37
	snowballParserT__37          = 38
	snowballParserT__38          = 39
	snowballParserT__39          = 40
	snowballParserT__40          = 41
	snowballParserT__41          = 42
	snowballParserT__42          = 43
	snowballParserT__43          = 44
	snowballParserT__44          = 45
	snowballParserT__45          = 46
	snowballParserT__46          = 47
	snowballParserT__47          = 48
	snowballParserT__48          = 49
	snowballParserT__49          = 50
	snowballParserT__50          = 51
	snowballParserT__51          = 52
	snowballParserT__52          = 53
	snowballParserT__53          = 54
	snowballParserT__54          = 55
	snowballParserT__55          = 56
	snowballParserT__56          = 57
	snowballParserT__57          = 58
	snowballParserT__58          = 59
	snowballParserT__59          = 60
	snowballParserT__60          = 61
	snowballParserT__61          = 62
	snowballParserT__62          = 63
	snowballParserT__63          = 64
	snowballParserT__64          = 65
	snowballParserT__65          = 66
	snowballParserT__66          = 67
	snowballParserT__67          = 68
	snowballParserT__68          = 69
	snowballParserT__69          = 70
	snowballParserT__70          = 71
	snowballParserT__71          = 72
	snowballParserLITERAL_STRING = 73
	snowballParserNUMBER         = 74
	snowballParserPLUS_OR_MINUS  = 75
	snowballParserNAME           = 76
	snowballParserWS             = 77
)

// snowballParser rules.
const (
	snowballParserRULE_program      = 0
	snowballParserRULE_p            = 1
	snowballParserRULE_declaration  = 2
	snowballParserRULE_r_definition = 3
	snowballParserRULE_g_definition = 4
	snowballParserRULE_c            = 5
	snowballParserRULE_i_command    = 6
	snowballParserRULE_s_command    = 7
	snowballParserRULE_s            = 8
	snowballParserRULE_g            = 9
	snowballParserRULE_s_name       = 10
	snowballParserRULE_i_name       = 11
	snowballParserRULE_b_name       = 12
	snowballParserRULE_r_name       = 13
	snowballParserRULE_g_name       = 14
	snowballParserRULE_ae           = 15
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
	p.RuleIndex = snowballParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllP() []IPContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPContext)(nil)).Elem())
	var tst = make([]IPContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPContext)
		}
	}

	return tst
}

func (s *ProgramContext) P(i int) IPContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *snowballParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, snowballParserRULE_program)
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

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<snowballParserT__0)|(1<<snowballParserT__3)|(1<<snowballParserT__4)|(1<<snowballParserT__5)|(1<<snowballParserT__6)|(1<<snowballParserT__7)|(1<<snowballParserT__8)|(1<<snowballParserT__9))) != 0 {
		{
			p.SetState(32)
			p.P()
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPContext is an interface to support dynamic dispatch.
type IPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPContext differentiates from other interfaces.
	IsPContext()
}

type PContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPContext() *PContext {
	var p = new(PContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_p
	return p
}

func (*PContext) IsPContext() {}

func NewPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PContext {
	var p = new(PContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_p

	return p
}

func (s *PContext) GetParser() antlr.Parser { return s.parser }

func (s *PContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *PContext) R_definition() IR_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IR_definitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IR_definitionContext)
}

func (s *PContext) G_definition() IG_definitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IG_definitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IG_definitionContext)
}

func (s *PContext) P() IPContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPContext)
}

func (s *PContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterP(s)
	}
}

func (s *PContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitP(s)
	}
}

func (p *snowballParser) P() (localctx IPContext) {
	this := p
	_ = this

	localctx = NewPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, snowballParserRULE_p)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.R_definition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.G_definition()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(41)
			p.Match(snowballParserT__0)
		}
		{
			p.SetState(42)
			p.Match(snowballParserT__1)
		}
		{
			p.SetState(43)
			p.P()
		}
		{
			p.SetState(44)
			p.Match(snowballParserT__2)
		}

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) AllS_name() []IS_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IS_nameContext)(nil)).Elem())
	var tst = make([]IS_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IS_nameContext)
		}
	}

	return tst
}

func (s *DeclarationContext) S_name(i int) IS_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IS_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IS_nameContext)
}

func (s *DeclarationContext) AllI_name() []II_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*II_nameContext)(nil)).Elem())
	var tst = make([]II_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(II_nameContext)
		}
	}

	return tst
}

func (s *DeclarationContext) I_name(i int) II_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*II_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(II_nameContext)
}

func (s *DeclarationContext) AllB_name() []IB_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IB_nameContext)(nil)).Elem())
	var tst = make([]IB_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IB_nameContext)
		}
	}

	return tst
}

func (s *DeclarationContext) B_name(i int) IB_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IB_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IB_nameContext)
}

func (s *DeclarationContext) AllR_name() []IR_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IR_nameContext)(nil)).Elem())
	var tst = make([]IR_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IR_nameContext)
		}
	}

	return tst
}

func (s *DeclarationContext) R_name(i int) IR_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IR_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IR_nameContext)
}

func (s *DeclarationContext) AllG_name() []IG_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IG_nameContext)(nil)).Elem())
	var tst = make([]IG_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IG_nameContext)
		}
	}

	return tst
}

func (s *DeclarationContext) G_name(i int) IG_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IG_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IG_nameContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *snowballParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, snowballParserRULE_declaration)
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

	p.SetState(102)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case snowballParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(48)
			p.Match(snowballParserT__3)
		}
		{
			p.SetState(49)
			p.Match(snowballParserT__1)
		}
		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == snowballParserNAME {
			{
				p.SetState(50)
				p.S_name()
			}

			p.SetState(55)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(56)
			p.Match(snowballParserT__2)
		}

	case snowballParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(57)
			p.Match(snowballParserT__4)
		}
		{
			p.SetState(58)
			p.Match(snowballParserT__1)
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == snowballParserNAME {
			{
				p.SetState(59)
				p.I_name()
			}

			p.SetState(64)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(65)
			p.Match(snowballParserT__2)
		}

	case snowballParserT__5:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(66)
			p.Match(snowballParserT__5)
		}
		{
			p.SetState(67)
			p.Match(snowballParserT__1)
		}
		p.SetState(71)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == snowballParserNAME {
			{
				p.SetState(68)
				p.B_name()
			}

			p.SetState(73)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(74)
			p.Match(snowballParserT__2)
		}

	case snowballParserT__6:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(75)
			p.Match(snowballParserT__6)
		}
		{
			p.SetState(76)
			p.Match(snowballParserT__1)
		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == snowballParserNAME {
			{
				p.SetState(77)
				p.R_name()
			}

			p.SetState(82)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(83)
			p.Match(snowballParserT__2)
		}

	case snowballParserT__7:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(84)
			p.Match(snowballParserT__7)
		}
		{
			p.SetState(85)
			p.Match(snowballParserT__1)
		}
		p.SetState(89)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == snowballParserNAME {
			{
				p.SetState(86)
				p.R_name()
			}

			p.SetState(91)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(92)
			p.Match(snowballParserT__2)
		}

	case snowballParserT__8:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(93)
			p.Match(snowballParserT__8)
		}
		{
			p.SetState(94)
			p.Match(snowballParserT__1)
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == snowballParserNAME {
			{
				p.SetState(95)
				p.G_name()
			}

			p.SetState(100)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(101)
			p.Match(snowballParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IR_definitionContext is an interface to support dynamic dispatch.
type IR_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsR_definitionContext differentiates from other interfaces.
	IsR_definitionContext()
}

type R_definitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyR_definitionContext() *R_definitionContext {
	var p = new(R_definitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_r_definition
	return p
}

func (*R_definitionContext) IsR_definitionContext() {}

func NewR_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *R_definitionContext {
	var p = new(R_definitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_r_definition

	return p
}

func (s *R_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *R_definitionContext) R_name() IR_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IR_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IR_nameContext)
}

func (s *R_definitionContext) C() ICContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICContext)
}

func (s *R_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *R_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *R_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterR_definition(s)
	}
}

func (s *R_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitR_definition(s)
	}
}

func (p *snowballParser) R_definition() (localctx IR_definitionContext) {
	this := p
	_ = this

	localctx = NewR_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, snowballParserRULE_r_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(snowballParserT__9)
	}
	{
		p.SetState(105)
		p.R_name()
	}
	{
		p.SetState(106)
		p.Match(snowballParserT__10)
	}
	{
		p.SetState(107)
		p.c(0)
	}

	return localctx
}

// IG_definitionContext is an interface to support dynamic dispatch.
type IG_definitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsG_definitionContext differentiates from other interfaces.
	IsG_definitionContext()
}

type G_definitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyG_definitionContext() *G_definitionContext {
	var p = new(G_definitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_g_definition
	return p
}

func (*G_definitionContext) IsG_definitionContext() {}

func NewG_definitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *G_definitionContext {
	var p = new(G_definitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_g_definition

	return p
}

func (s *G_definitionContext) GetParser() antlr.Parser { return s.parser }

func (s *G_definitionContext) G_name() IG_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IG_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IG_nameContext)
}

func (s *G_definitionContext) AllG() []IGContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGContext)(nil)).Elem())
	var tst = make([]IGContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGContext)
		}
	}

	return tst
}

func (s *G_definitionContext) G(i int) IGContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGContext)
}

func (s *G_definitionContext) AllPLUS_OR_MINUS() []antlr.TerminalNode {
	return s.GetTokens(snowballParserPLUS_OR_MINUS)
}

func (s *G_definitionContext) PLUS_OR_MINUS(i int) antlr.TerminalNode {
	return s.GetToken(snowballParserPLUS_OR_MINUS, i)
}

func (s *G_definitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *G_definitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *G_definitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterG_definition(s)
	}
}

func (s *G_definitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitG_definition(s)
	}
}

func (p *snowballParser) G_definition() (localctx IG_definitionContext) {
	this := p
	_ = this

	localctx = NewG_definitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, snowballParserRULE_g_definition)
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
		p.SetState(109)
		p.Match(snowballParserT__9)
	}
	{
		p.SetState(110)
		p.G_name()
	}
	{
		p.SetState(111)
		p.G()
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == snowballParserPLUS_OR_MINUS {
		{
			p.SetState(112)
			p.Match(snowballParserPLUS_OR_MINUS)
		}
		{
			p.SetState(113)
			p.G()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICContext is an interface to support dynamic dispatch.
type ICContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCContext differentiates from other interfaces.
	IsCContext()
}

type CContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCContext() *CContext {
	var p = new(CContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_c
	return p
}

func (*CContext) IsCContext() {}

func NewCContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CContext {
	var p = new(CContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_c

	return p
}

func (s *CContext) GetParser() antlr.Parser { return s.parser }

func (s *CContext) AllC() []ICContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICContext)(nil)).Elem())
	var tst = make([]ICContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICContext)
		}
	}

	return tst
}

func (s *CContext) C(i int) ICContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICContext)
}

func (s *CContext) I_command() II_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*II_commandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(II_commandContext)
}

func (s *CContext) S_command() IS_commandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IS_commandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IS_commandContext)
}

func (s *CContext) Ae() IAeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAeContext)
}

func (s *CContext) S() ISContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISContext)
}

func (s *CContext) S_name() IS_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IS_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IS_nameContext)
}

func (s *CContext) I_name() II_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*II_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(II_nameContext)
}

func (s *CContext) AllLITERAL_STRING() []antlr.TerminalNode {
	return s.GetTokens(snowballParserLITERAL_STRING)
}

func (s *CContext) LITERAL_STRING(i int) antlr.TerminalNode {
	return s.GetToken(snowballParserLITERAL_STRING, i)
}

func (s *CContext) AllR_name() []IR_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IR_nameContext)(nil)).Elem())
	var tst = make([]IR_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IR_nameContext)
		}
	}

	return tst
}

func (s *CContext) R_name(i int) IR_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IR_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IR_nameContext)
}

func (s *CContext) B_name() IB_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IB_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IB_nameContext)
}

func (s *CContext) G_name() IG_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IG_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IG_nameContext)
}

func (s *CContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterC(s)
	}
}

func (s *CContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitC(s)
	}
}

func (p *snowballParser) C() (localctx ICContext) {
	return p.c(0)
}

func (p *snowballParser) c(_p int) (localctx ICContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, snowballParserRULE_c, _p)
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
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(120)
			p.Match(snowballParserT__1)
		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<snowballParserT__1)|(1<<snowballParserT__13)|(1<<snowballParserT__14)|(1<<snowballParserT__15)|(1<<snowballParserT__16)|(1<<snowballParserT__17)|(1<<snowballParserT__18)|(1<<snowballParserT__19)|(1<<snowballParserT__20)|(1<<snowballParserT__21)|(1<<snowballParserT__22)|(1<<snowballParserT__23)|(1<<snowballParserT__24)|(1<<snowballParserT__25)|(1<<snowballParserT__26)|(1<<snowballParserT__27)|(1<<snowballParserT__28)|(1<<snowballParserT__29)|(1<<snowballParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(snowballParserT__31-32))|(1<<(snowballParserT__32-32))|(1<<(snowballParserT__33-32))|(1<<(snowballParserT__34-32))|(1<<(snowballParserT__35-32))|(1<<(snowballParserT__36-32))|(1<<(snowballParserT__37-32))|(1<<(snowballParserT__38-32))|(1<<(snowballParserT__39-32))|(1<<(snowballParserT__41-32))|(1<<(snowballParserT__42-32))|(1<<(snowballParserT__43-32))|(1<<(snowballParserT__44-32))|(1<<(snowballParserT__45-32))|(1<<(snowballParserT__46-32))|(1<<(snowballParserT__47-32))|(1<<(snowballParserT__49-32))|(1<<(snowballParserT__50-32))|(1<<(snowballParserT__51-32))|(1<<(snowballParserT__52-32)))) != 0) || _la == snowballParserLITERAL_STRING || _la == snowballParserNAME {
			{
				p.SetState(121)
				p.c(0)
			}

			p.SetState(126)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(127)
			p.Match(snowballParserT__2)
		}

	case 2:
		{
			p.SetState(128)
			p.I_command()
		}

	case 3:
		{
			p.SetState(129)
			p.S_command()
		}

	case 4:
		{
			p.SetState(130)
			p.Match(snowballParserT__13)
		}
		{
			p.SetState(131)
			p.c(41)
		}

	case 5:
		{
			p.SetState(132)
			p.Match(snowballParserT__14)
		}
		{
			p.SetState(133)
			p.c(40)
		}

	case 6:
		{
			p.SetState(134)
			p.Match(snowballParserT__15)
		}
		{
			p.SetState(135)
			p.c(39)
		}

	case 7:
		{
			p.SetState(136)
			p.Match(snowballParserT__16)
		}
		{
			p.SetState(137)
			p.c(38)
		}

	case 8:
		{
			p.SetState(138)
			p.Match(snowballParserT__17)
		}
		{
			p.SetState(139)
			p.c(37)
		}

	case 9:
		{
			p.SetState(140)
			p.Match(snowballParserT__18)
		}
		{
			p.SetState(141)
			p.c(36)
		}

	case 10:
		{
			p.SetState(142)
			p.Match(snowballParserT__19)
		}
		{
			p.SetState(143)
			p.c(35)
		}

	case 11:
		{
			p.SetState(144)
			p.Match(snowballParserT__20)
		}
		{
			p.SetState(145)
			p.c(34)
		}

	case 12:
		{
			p.SetState(146)
			p.Match(snowballParserT__21)
		}
		{
			p.SetState(147)
			p.ae(0)
		}
		{
			p.SetState(148)
			p.c(33)
		}

	case 13:
		{
			p.SetState(150)
			p.Match(snowballParserT__22)
		}
		{
			p.SetState(151)
			p.ae(0)
		}
		{
			p.SetState(152)
			p.c(32)
		}

	case 14:
		{
			p.SetState(154)
			p.S()
		}

	case 15:
		{
			p.SetState(155)
			p.Match(snowballParserT__23)
		}
		{
			p.SetState(156)
			p.S()
		}

	case 16:
		{
			p.SetState(157)
			p.Match(snowballParserT__24)
		}
		{
			p.SetState(158)
			p.S()
		}

	case 17:
		{
			p.SetState(159)
			p.Match(snowballParserT__25)
		}
		{
			p.SetState(160)
			p.S()
		}

	case 18:
		{
			p.SetState(161)
			p.Match(snowballParserT__26)
		}
		{
			p.SetState(162)
			p.S()
		}

	case 19:
		{
			p.SetState(163)
			p.Match(snowballParserT__27)
		}

	case 20:
		{
			p.SetState(164)
			p.Match(snowballParserT__28)
		}
		{
			p.SetState(165)
			p.ae(0)
		}

	case 21:
		{
			p.SetState(166)
			p.Match(snowballParserT__29)
		}

	case 22:
		{
			p.SetState(167)
			p.Match(snowballParserT__30)
		}
		{
			p.SetState(168)
			p.S_name()
		}

	case 23:
		{
			p.SetState(169)
			p.Match(snowballParserT__31)
		}

	case 24:
		{
			p.SetState(170)
			p.Match(snowballParserT__32)
		}

	case 25:
		{
			p.SetState(171)
			p.Match(snowballParserT__33)
		}
		{
			p.SetState(172)
			p.S_name()
		}

	case 26:
		{
			p.SetState(173)
			p.Match(snowballParserT__34)
		}
		{
			p.SetState(174)
			p.I_name()
		}

	case 27:
		{
			p.SetState(175)
			p.Match(snowballParserT__35)
		}
		{
			p.SetState(176)
			p.ae(0)
		}

	case 28:
		{
			p.SetState(177)
			p.Match(snowballParserT__36)
		}
		{
			p.SetState(178)
			p.ae(0)
		}

	case 29:
		{
			p.SetState(179)
			p.Match(snowballParserT__37)
		}

	case 30:
		{
			p.SetState(180)
			p.Match(snowballParserT__38)
		}

	case 31:
		{
			p.SetState(181)
			p.Match(snowballParserT__39)
		}
		{
			p.SetState(182)
			p.c(0)
		}
		{
			p.SetState(183)
			p.Match(snowballParserT__40)
		}
		{
			p.SetState(184)
			p.c(14)
		}

	case 32:
		{
			p.SetState(186)
			p.Match(snowballParserT__41)
		}
		{
			p.SetState(187)
			p.c(13)
		}

	case 33:
		{
			p.SetState(188)
			p.Match(snowballParserT__42)
		}
		{
			p.SetState(189)
			p.c(12)
		}

	case 34:
		{
			p.SetState(190)
			p.Match(snowballParserT__43)
		}

	case 35:
		{
			p.SetState(191)
			p.Match(snowballParserT__44)
		}
		{
			p.SetState(192)
			p.Match(snowballParserT__1)
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<snowballParserT__1)|(1<<snowballParserT__13)|(1<<snowballParserT__14)|(1<<snowballParserT__15)|(1<<snowballParserT__16)|(1<<snowballParserT__17)|(1<<snowballParserT__18)|(1<<snowballParserT__19)|(1<<snowballParserT__20)|(1<<snowballParserT__21)|(1<<snowballParserT__22)|(1<<snowballParserT__23)|(1<<snowballParserT__24)|(1<<snowballParserT__25)|(1<<snowballParserT__26)|(1<<snowballParserT__27)|(1<<snowballParserT__28)|(1<<snowballParserT__29)|(1<<snowballParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(snowballParserT__31-32))|(1<<(snowballParserT__32-32))|(1<<(snowballParserT__33-32))|(1<<(snowballParserT__34-32))|(1<<(snowballParserT__35-32))|(1<<(snowballParserT__36-32))|(1<<(snowballParserT__37-32))|(1<<(snowballParserT__38-32))|(1<<(snowballParserT__39-32))|(1<<(snowballParserT__41-32))|(1<<(snowballParserT__42-32))|(1<<(snowballParserT__43-32))|(1<<(snowballParserT__44-32))|(1<<(snowballParserT__45-32))|(1<<(snowballParserT__46-32))|(1<<(snowballParserT__47-32))|(1<<(snowballParserT__49-32))|(1<<(snowballParserT__50-32))|(1<<(snowballParserT__51-32))|(1<<(snowballParserT__52-32)))) != 0) || _la == snowballParserLITERAL_STRING || _la == snowballParserNAME {
			p.SetState(198)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
			case 1:
				{
					p.SetState(193)
					p.Match(snowballParserLITERAL_STRING)
				}
				p.SetState(195)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(194)
						p.R_name()
					}

				}

			case 2:
				{
					p.SetState(197)
					p.c(0)
				}

			}

			p.SetState(202)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(203)
			p.Match(snowballParserT__2)
		}

	case 36:
		{
			p.SetState(204)
			p.Match(snowballParserT__45)
		}
		{
			p.SetState(205)
			p.B_name()
		}

	case 37:
		{
			p.SetState(206)
			p.Match(snowballParserT__46)
		}
		{
			p.SetState(207)
			p.B_name()
		}

	case 38:
		{
			p.SetState(208)
			p.B_name()
		}

	case 39:
		{
			p.SetState(209)
			p.R_name()
		}

	case 40:
		{
			p.SetState(210)
			p.G_name()
		}

	case 41:
		{
			p.SetState(211)
			p.Match(snowballParserT__47)
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == snowballParserT__48 {
			{
				p.SetState(212)
				p.Match(snowballParserT__48)
			}

		}
		{
			p.SetState(215)
			p.G_name()
		}

	case 42:
		{
			p.SetState(216)
			p.Match(snowballParserT__49)
		}

	case 43:
		{
			p.SetState(217)
			p.Match(snowballParserT__50)
		}

	case 44:
		{
			p.SetState(218)
			p.Match(snowballParserT__51)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(227)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, snowballParserRULE_c)
				p.SetState(221)

				if !(p.Precpred(p.GetParserRuleContext(), 43)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 43)", ""))
				}
				{
					p.SetState(222)
					p.Match(snowballParserT__11)
				}
				{
					p.SetState(223)
					p.c(44)
				}

			case 2:
				localctx = NewCContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, snowballParserRULE_c)
				p.SetState(224)

				if !(p.Precpred(p.GetParserRuleContext(), 42)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 42)", ""))
				}
				{
					p.SetState(225)
					p.Match(snowballParserT__12)
				}
				{
					p.SetState(226)
					p.c(43)
				}

			}

		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// II_commandContext is an interface to support dynamic dispatch.
type II_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsI_commandContext differentiates from other interfaces.
	IsI_commandContext()
}

type I_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyI_commandContext() *I_commandContext {
	var p = new(I_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_i_command
	return p
}

func (*I_commandContext) IsI_commandContext() {}

func NewI_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *I_commandContext {
	var p = new(I_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_i_command

	return p
}

func (s *I_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *I_commandContext) I_name() II_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*II_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(II_nameContext)
}

func (s *I_commandContext) Ae() IAeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAeContext)
}

func (s *I_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *I_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *I_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterI_command(s)
	}
}

func (s *I_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitI_command(s)
	}
}

func (p *snowballParser) I_command() (localctx II_commandContext) {
	this := p
	_ = this

	localctx = NewI_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, snowballParserRULE_i_command)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(232)
			p.Match(snowballParserT__52)
		}
		{
			p.SetState(233)
			p.I_name()
		}
		{
			p.SetState(234)
			p.Match(snowballParserT__23)
		}
		{
			p.SetState(235)
			p.ae(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(237)
			p.Match(snowballParserT__52)
		}
		{
			p.SetState(238)
			p.I_name()
		}
		{
			p.SetState(239)
			p.Match(snowballParserT__53)
		}
		{
			p.SetState(240)
			p.ae(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
			p.Match(snowballParserT__52)
		}
		{
			p.SetState(243)
			p.I_name()
		}
		{
			p.SetState(244)
			p.Match(snowballParserT__54)
		}
		{
			p.SetState(245)
			p.ae(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(247)
			p.Match(snowballParserT__52)
		}
		{
			p.SetState(248)
			p.I_name()
		}
		{
			p.SetState(249)
			p.Match(snowballParserT__55)
		}
		{
			p.SetState(250)
			p.ae(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(252)
			p.Match(snowballParserT__52)
		}
		{
			p.SetState(253)
			p.I_name()
		}
		{
			p.SetState(254)
			p.Match(snowballParserT__56)
		}
		{
			p.SetState(255)
			p.ae(0)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(257)
			p.Match(snowballParserT__52)
		}
		{
			p.SetState(258)
			p.I_name()
		}
		{
			p.SetState(259)
			p.Match(snowballParserT__57)
		}
		{
			p.SetState(260)
			p.ae(0)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(262)
			p.Match(snowballParserT__52)
		}
		{
			p.SetState(263)
			p.I_name()
		}
		{
			p.SetState(264)
			p.Match(snowballParserT__58)
		}
		{
			p.SetState(265)
			p.ae(0)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(267)
			p.Match(snowballParserT__52)
		}
		{
			p.SetState(268)
			p.I_name()
		}
		{
			p.SetState(269)
			p.Match(snowballParserT__59)
		}
		{
			p.SetState(270)
			p.ae(0)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(272)
			p.Match(snowballParserT__52)
		}
		{
			p.SetState(273)
			p.I_name()
		}
		{
			p.SetState(274)
			p.Match(snowballParserT__60)
		}
		{
			p.SetState(275)
			p.ae(0)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(277)
			p.Match(snowballParserT__52)
		}
		{
			p.SetState(278)
			p.I_name()
		}
		{
			p.SetState(279)
			p.Match(snowballParserT__61)
		}
		{
			p.SetState(280)
			p.ae(0)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(282)
			p.Match(snowballParserT__52)
		}
		{
			p.SetState(283)
			p.I_name()
		}
		{
			p.SetState(284)
			p.Match(snowballParserT__62)
		}
		{
			p.SetState(285)
			p.ae(0)
		}

	}

	return localctx
}

// IS_commandContext is an interface to support dynamic dispatch.
type IS_commandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsS_commandContext differentiates from other interfaces.
	IsS_commandContext()
}

type S_commandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyS_commandContext() *S_commandContext {
	var p = new(S_commandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_s_command
	return p
}

func (*S_commandContext) IsS_commandContext() {}

func NewS_commandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *S_commandContext {
	var p = new(S_commandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_s_command

	return p
}

func (s *S_commandContext) GetParser() antlr.Parser { return s.parser }

func (s *S_commandContext) S_name() IS_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IS_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IS_nameContext)
}

func (s *S_commandContext) C() ICContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICContext)
}

func (s *S_commandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_commandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *S_commandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterS_command(s)
	}
}

func (s *S_commandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitS_command(s)
	}
}

func (p *snowballParser) S_command() (localctx IS_commandContext) {
	this := p
	_ = this

	localctx = NewS_commandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, snowballParserRULE_s_command)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(289)
		p.Match(snowballParserT__52)
	}
	{
		p.SetState(290)
		p.S_name()
	}
	{
		p.SetState(291)
		p.c(0)
	}

	return localctx
}

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_s
	return p
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) S_name() IS_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IS_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IS_nameContext)
}

func (s *SContext) LITERAL_STRING() antlr.TerminalNode {
	return s.GetToken(snowballParserLITERAL_STRING, 0)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitS(s)
	}
}

func (p *snowballParser) S() (localctx ISContext) {
	this := p
	_ = this

	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, snowballParserRULE_s)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(295)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case snowballParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.S_name()
		}

	case snowballParserLITERAL_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(294)
			p.Match(snowballParserLITERAL_STRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IGContext is an interface to support dynamic dispatch.
type IGContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGContext differentiates from other interfaces.
	IsGContext()
}

type GContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGContext() *GContext {
	var p = new(GContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_g
	return p
}

func (*GContext) IsGContext() {}

func NewGContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GContext {
	var p = new(GContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_g

	return p
}

func (s *GContext) GetParser() antlr.Parser { return s.parser }

func (s *GContext) G_name() IG_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IG_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IG_nameContext)
}

func (s *GContext) LITERAL_STRING() antlr.TerminalNode {
	return s.GetToken(snowballParserLITERAL_STRING, 0)
}

func (s *GContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterG(s)
	}
}

func (s *GContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitG(s)
	}
}

func (p *snowballParser) G() (localctx IGContext) {
	this := p
	_ = this

	localctx = NewGContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, snowballParserRULE_g)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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

	switch p.GetTokenStream().LA(1) {
	case snowballParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.G_name()
		}

	case snowballParserLITERAL_STRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(298)
			p.Match(snowballParserLITERAL_STRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IS_nameContext is an interface to support dynamic dispatch.
type IS_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsS_nameContext differentiates from other interfaces.
	IsS_nameContext()
}

type S_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyS_nameContext() *S_nameContext {
	var p = new(S_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_s_name
	return p
}

func (*S_nameContext) IsS_nameContext() {}

func NewS_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *S_nameContext {
	var p = new(S_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_s_name

	return p
}

func (s *S_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *S_nameContext) NAME() antlr.TerminalNode {
	return s.GetToken(snowballParserNAME, 0)
}

func (s *S_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *S_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterS_name(s)
	}
}

func (s *S_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitS_name(s)
	}
}

func (p *snowballParser) S_name() (localctx IS_nameContext) {
	this := p
	_ = this

	localctx = NewS_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, snowballParserRULE_s_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(snowballParserNAME)
	}

	return localctx
}

// II_nameContext is an interface to support dynamic dispatch.
type II_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsI_nameContext differentiates from other interfaces.
	IsI_nameContext()
}

type I_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyI_nameContext() *I_nameContext {
	var p = new(I_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_i_name
	return p
}

func (*I_nameContext) IsI_nameContext() {}

func NewI_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *I_nameContext {
	var p = new(I_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_i_name

	return p
}

func (s *I_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *I_nameContext) NAME() antlr.TerminalNode {
	return s.GetToken(snowballParserNAME, 0)
}

func (s *I_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *I_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *I_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterI_name(s)
	}
}

func (s *I_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitI_name(s)
	}
}

func (p *snowballParser) I_name() (localctx II_nameContext) {
	this := p
	_ = this

	localctx = NewI_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, snowballParserRULE_i_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(snowballParserNAME)
	}

	return localctx
}

// IB_nameContext is an interface to support dynamic dispatch.
type IB_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsB_nameContext differentiates from other interfaces.
	IsB_nameContext()
}

type B_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyB_nameContext() *B_nameContext {
	var p = new(B_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_b_name
	return p
}

func (*B_nameContext) IsB_nameContext() {}

func NewB_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *B_nameContext {
	var p = new(B_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_b_name

	return p
}

func (s *B_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *B_nameContext) NAME() antlr.TerminalNode {
	return s.GetToken(snowballParserNAME, 0)
}

func (s *B_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *B_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *B_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterB_name(s)
	}
}

func (s *B_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitB_name(s)
	}
}

func (p *snowballParser) B_name() (localctx IB_nameContext) {
	this := p
	_ = this

	localctx = NewB_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, snowballParserRULE_b_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(snowballParserNAME)
	}

	return localctx
}

// IR_nameContext is an interface to support dynamic dispatch.
type IR_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsR_nameContext differentiates from other interfaces.
	IsR_nameContext()
}

type R_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyR_nameContext() *R_nameContext {
	var p = new(R_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_r_name
	return p
}

func (*R_nameContext) IsR_nameContext() {}

func NewR_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *R_nameContext {
	var p = new(R_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_r_name

	return p
}

func (s *R_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *R_nameContext) NAME() antlr.TerminalNode {
	return s.GetToken(snowballParserNAME, 0)
}

func (s *R_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *R_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *R_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterR_name(s)
	}
}

func (s *R_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitR_name(s)
	}
}

func (p *snowballParser) R_name() (localctx IR_nameContext) {
	this := p
	_ = this

	localctx = NewR_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, snowballParserRULE_r_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.Match(snowballParserNAME)
	}

	return localctx
}

// IG_nameContext is an interface to support dynamic dispatch.
type IG_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsG_nameContext differentiates from other interfaces.
	IsG_nameContext()
}

type G_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyG_nameContext() *G_nameContext {
	var p = new(G_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_g_name
	return p
}

func (*G_nameContext) IsG_nameContext() {}

func NewG_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *G_nameContext {
	var p = new(G_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_g_name

	return p
}

func (s *G_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *G_nameContext) NAME() antlr.TerminalNode {
	return s.GetToken(snowballParserNAME, 0)
}

func (s *G_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *G_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *G_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterG_name(s)
	}
}

func (s *G_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitG_name(s)
	}
}

func (p *snowballParser) G_name() (localctx IG_nameContext) {
	this := p
	_ = this

	localctx = NewG_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, snowballParserRULE_g_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
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
		p.SetState(309)
		p.Match(snowballParserNAME)
	}

	return localctx
}

// IAeContext is an interface to support dynamic dispatch.
type IAeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAeContext differentiates from other interfaces.
	IsAeContext()
}

type AeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAeContext() *AeContext {
	var p = new(AeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = snowballParserRULE_ae
	return p
}

func (*AeContext) IsAeContext() {}

func NewAeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AeContext {
	var p = new(AeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = snowballParserRULE_ae

	return p
}

func (s *AeContext) GetParser() antlr.Parser { return s.parser }

func (s *AeContext) AllAe() []IAeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAeContext)(nil)).Elem())
	var tst = make([]IAeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAeContext)
		}
	}

	return tst
}

func (s *AeContext) Ae(i int) IAeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAeContext)
}

func (s *AeContext) S_name() IS_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IS_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IS_nameContext)
}

func (s *AeContext) I_name() II_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*II_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(II_nameContext)
}

func (s *AeContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(snowballParserNUMBER, 0)
}

func (s *AeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.EnterAe(s)
	}
}

func (s *AeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(snowballListener); ok {
		listenerT.ExitAe(s)
	}
}

func (p *snowballParser) Ae() (localctx IAeContext) {
	return p.ae(0)
}

func (p *snowballParser) ae(_p int) (localctx IAeContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewAeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IAeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, snowballParserRULE_ae, _p)

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
	p.SetState(327)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case snowballParserT__1:
		{
			p.SetState(312)
			p.Match(snowballParserT__1)
		}
		{
			p.SetState(313)
			p.ae(0)
		}
		{
			p.SetState(314)
			p.Match(snowballParserT__2)
		}

	case snowballParserT__48:
		{
			p.SetState(316)
			p.Match(snowballParserT__48)
		}
		{
			p.SetState(317)
			p.ae(9)
		}

	case snowballParserT__66:
		{
			p.SetState(318)
			p.Match(snowballParserT__66)
		}

	case snowballParserT__67:
		{
			p.SetState(319)
			p.Match(snowballParserT__67)
		}

	case snowballParserT__68:
		{
			p.SetState(320)
			p.Match(snowballParserT__68)
		}

	case snowballParserT__69:
		{
			p.SetState(321)
			p.Match(snowballParserT__69)
		}

	case snowballParserT__70:
		{
			p.SetState(322)
			p.Match(snowballParserT__70)
		}

	case snowballParserT__71:
		{
			p.SetState(323)
			p.Match(snowballParserT__71)
		}
		{
			p.SetState(324)
			p.S_name()
		}

	case snowballParserNAME:
		{
			p.SetState(325)
			p.I_name()
		}

	case snowballParserNUMBER:
		{
			p.SetState(326)
			p.Match(snowballParserNUMBER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(341)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
			case 1:
				localctx = NewAeContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, snowballParserRULE_ae)
				p.SetState(329)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(330)
					p.Match(snowballParserT__63)
				}
				{
					p.SetState(331)
					p.ae(14)
				}

			case 2:
				localctx = NewAeContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, snowballParserRULE_ae)
				p.SetState(332)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(333)
					p.Match(snowballParserT__48)
				}
				{
					p.SetState(334)
					p.ae(13)
				}

			case 3:
				localctx = NewAeContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, snowballParserRULE_ae)
				p.SetState(335)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(336)
					p.Match(snowballParserT__64)
				}
				{
					p.SetState(337)
					p.ae(12)
				}

			case 4:
				localctx = NewAeContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, snowballParserRULE_ae)
				p.SetState(338)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(339)
					p.Match(snowballParserT__65)
				}
				{
					p.SetState(340)
					p.ae(11)
				}

			}

		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}

	return localctx
}

func (p *snowballParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *CContext = nil
		if localctx != nil {
			t = localctx.(*CContext)
		}
		return p.C_Sempred(t, predIndex)

	case 15:
		var t *AeContext = nil
		if localctx != nil {
			t = localctx.(*AeContext)
		}
		return p.Ae_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *snowballParser) C_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 43)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 42)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *snowballParser) Ae_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
