// Code generated from PowerBuilderDWParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

package powerbuilderdw // PowerBuilderDWParser
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 157, 357,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	3, 2, 5, 2, 48, 10, 2, 3, 2, 6, 2, 51, 10, 2, 13, 2, 14, 2, 52, 3, 2, 3,
	2, 3, 3, 7, 3, 58, 10, 3, 12, 3, 14, 3, 61, 11, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 66, 10, 3, 3, 4, 6, 4, 69, 10, 4, 13, 4, 14, 4, 70, 3, 5, 3, 5, 3, 5,
	7, 5, 76, 10, 5, 12, 5, 14, 5, 79, 11, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	6, 5, 86, 10, 5, 13, 5, 14, 5, 87, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5,
	95, 10, 5, 12, 5, 14, 5, 98, 11, 5, 3, 5, 5, 5, 101, 10, 5, 3, 6, 3, 6,
	5, 6, 105, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 6, 7, 117, 10, 7, 13, 7, 14, 7, 118, 3, 7, 3, 7, 3, 8, 3, 8, 5, 8,
	125, 10, 8, 3, 8, 5, 8, 128, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 5, 9, 139, 10, 9, 3, 9, 3, 9, 6, 9, 143, 10, 9, 13, 9,
	14, 9, 144, 3, 9, 3, 9, 5, 9, 149, 10, 9, 5, 9, 151, 10, 9, 3, 9, 5, 9,
	154, 10, 9, 3, 10, 3, 10, 3, 10, 5, 10, 159, 10, 10, 3, 10, 5, 10, 162,
	10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 170, 10, 10, 7,
	10, 172, 10, 10, 12, 10, 14, 10, 175, 11, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 5, 12, 182, 10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	5, 12, 201, 10, 12, 3, 12, 3, 12, 3, 12, 5, 12, 206, 10, 12, 5, 12, 208,
	10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 215, 10, 12, 3, 12, 3,
	12, 3, 12, 5, 12, 220, 10, 12, 5, 12, 222, 10, 12, 3, 12, 3, 12, 7, 12,
	226, 10, 12, 12, 12, 14, 12, 229, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 5, 12, 237, 10, 12, 5, 12, 239, 10, 12, 3, 13, 3, 13, 3, 14,
	3, 14, 3, 15, 3, 15, 5, 15, 247, 10, 15, 3, 16, 3, 16, 3, 16, 5, 16, 252,
	10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 257, 10, 16, 3, 16, 5, 16, 260, 10,
	16, 3, 16, 3, 16, 5, 16, 264, 10, 16, 3, 16, 3, 16, 3, 16, 5, 16, 269,
	10, 16, 3, 16, 5, 16, 272, 10, 16, 7, 16, 274, 10, 16, 12, 16, 14, 16,
	277, 11, 16, 5, 16, 279, 10, 16, 3, 16, 5, 16, 282, 10, 16, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 291, 10, 17, 3, 18, 5, 18, 294,
	10, 18, 3, 18, 3, 18, 3, 18, 5, 18, 299, 10, 18, 3, 18, 7, 18, 302, 10,
	18, 12, 18, 14, 18, 305, 11, 18, 3, 19, 3, 19, 5, 19, 309, 10, 19, 3, 19,
	3, 19, 5, 19, 313, 10, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 5, 20, 323, 10, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 337, 10, 20, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21,
	5, 21, 351, 10, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 2, 2, 24, 2, 4,
	6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
	44, 2, 6, 3, 2, 10, 32, 3, 2, 116, 117, 3, 2, 87, 88, 3, 2, 33, 34, 2,
	423, 2, 47, 3, 2, 2, 2, 4, 59, 3, 2, 2, 2, 6, 68, 3, 2, 2, 2, 8, 100, 3,
	2, 2, 2, 10, 104, 3, 2, 2, 2, 12, 106, 3, 2, 2, 2, 14, 122, 3, 2, 2, 2,
	16, 150, 3, 2, 2, 2, 18, 158, 3, 2, 2, 2, 20, 176, 3, 2, 2, 2, 22, 238,
	3, 2, 2, 2, 24, 240, 3, 2, 2, 2, 26, 242, 3, 2, 2, 2, 28, 246, 3, 2, 2,
	2, 30, 281, 3, 2, 2, 2, 32, 290, 3, 2, 2, 2, 34, 293, 3, 2, 2, 2, 36, 308,
	3, 2, 2, 2, 38, 336, 3, 2, 2, 2, 40, 350, 3, 2, 2, 2, 42, 352, 3, 2, 2,
	2, 44, 354, 3, 2, 2, 2, 46, 48, 5, 4, 3, 2, 47, 46, 3, 2, 2, 2, 47, 48,
	3, 2, 2, 2, 48, 50, 3, 2, 2, 2, 49, 51, 5, 6, 4, 2, 50, 49, 3, 2, 2, 2,
	51, 52, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 54, 3,
	2, 2, 2, 54, 55, 7, 2, 2, 3, 55, 3, 3, 2, 2, 2, 56, 58, 7, 148, 2, 2, 57,
	56, 3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2, 59, 60, 3, 2, 2,
	2, 60, 65, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 63, 7, 86, 2, 2, 63, 64,
	7, 143, 2, 2, 64, 66, 7, 135, 2, 2, 65, 62, 3, 2, 2, 2, 65, 66, 3, 2, 2,
	2, 66, 5, 3, 2, 2, 2, 67, 69, 5, 8, 5, 2, 68, 67, 3, 2, 2, 2, 69, 70, 3,
	2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 7, 3, 2, 2, 2, 72,
	73, 5, 18, 10, 2, 73, 77, 7, 136, 2, 2, 74, 76, 5, 16, 9, 2, 75, 74, 3,
	2, 2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78,
	80, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 81, 7, 137, 2, 2, 81, 101, 3, 2,
	2, 2, 82, 83, 7, 3, 2, 2, 83, 85, 7, 136, 2, 2, 84, 86, 5, 10, 6, 2, 85,
	84, 3, 2, 2, 2, 86, 87, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2, 2,
	2, 88, 89, 3, 2, 2, 2, 89, 90, 7, 137, 2, 2, 90, 101, 3, 2, 2, 2, 91, 92,
	7, 4, 2, 2, 92, 96, 7, 136, 2, 2, 93, 95, 5, 16, 9, 2, 94, 93, 3, 2, 2,
	2, 95, 98, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 99,
	3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99, 101, 7, 137, 2, 2, 100, 72, 3, 2, 2,
	2, 100, 82, 3, 2, 2, 2, 100, 91, 3, 2, 2, 2, 101, 9, 3, 2, 2, 2, 102, 105,
	5, 12, 7, 2, 103, 105, 5, 14, 8, 2, 104, 102, 3, 2, 2, 2, 104, 103, 3,
	2, 2, 2, 105, 11, 3, 2, 2, 2, 106, 107, 7, 4, 2, 2, 107, 108, 7, 110, 2,
	2, 108, 109, 7, 136, 2, 2, 109, 110, 7, 47, 2, 2, 110, 111, 7, 110, 2,
	2, 111, 112, 5, 26, 14, 2, 112, 113, 7, 136, 2, 2, 113, 114, 5, 24, 13,
	2, 114, 116, 7, 137, 2, 2, 115, 117, 5, 16, 9, 2, 116, 115, 3, 2, 2, 2,
	117, 118, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119,
	120, 3, 2, 2, 2, 120, 121, 7, 137, 2, 2, 121, 13, 3, 2, 2, 2, 122, 124,
	7, 151, 2, 2, 123, 125, 7, 152, 2, 2, 124, 123, 3, 2, 2, 2, 124, 125, 3,
	2, 2, 2, 125, 127, 3, 2, 2, 2, 126, 128, 7, 153, 2, 2, 127, 126, 3, 2,
	2, 2, 127, 128, 3, 2, 2, 2, 128, 15, 3, 2, 2, 2, 129, 151, 7, 52, 2, 2,
	130, 151, 5, 24, 13, 2, 131, 151, 7, 132, 2, 2, 132, 151, 7, 145, 2, 2,
	133, 151, 7, 146, 2, 2, 134, 135, 5, 18, 10, 2, 135, 148, 7, 110, 2, 2,
	136, 138, 5, 22, 12, 2, 137, 139, 5, 30, 16, 2, 138, 137, 3, 2, 2, 2, 138,
	139, 3, 2, 2, 2, 139, 149, 3, 2, 2, 2, 140, 142, 7, 136, 2, 2, 141, 143,
	5, 16, 9, 2, 142, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 142, 3, 2,
	2, 2, 144, 145, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 147, 7, 137, 2,
	2, 147, 149, 3, 2, 2, 2, 148, 136, 3, 2, 2, 2, 148, 140, 3, 2, 2, 2, 149,
	151, 3, 2, 2, 2, 150, 129, 3, 2, 2, 2, 150, 130, 3, 2, 2, 2, 150, 131,
	3, 2, 2, 2, 150, 132, 3, 2, 2, 2, 150, 133, 3, 2, 2, 2, 150, 134, 3, 2,
	2, 2, 151, 153, 3, 2, 2, 2, 152, 154, 7, 134, 2, 2, 153, 152, 3, 2, 2,
	2, 153, 154, 3, 2, 2, 2, 154, 17, 3, 2, 2, 2, 155, 159, 5, 20, 11, 2, 156,
	159, 7, 47, 2, 2, 157, 159, 7, 53, 2, 2, 158, 155, 3, 2, 2, 2, 158, 156,
	3, 2, 2, 2, 158, 157, 3, 2, 2, 2, 159, 161, 3, 2, 2, 2, 160, 162, 7, 143,
	2, 2, 161, 160, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 173, 3, 2, 2, 2,
	163, 169, 7, 144, 2, 2, 164, 170, 5, 20, 11, 2, 165, 170, 7, 54, 2, 2,
	166, 170, 7, 47, 2, 2, 167, 170, 7, 48, 2, 2, 168, 170, 7, 55, 2, 2, 169,
	164, 3, 2, 2, 2, 169, 165, 3, 2, 2, 2, 169, 166, 3, 2, 2, 2, 169, 167,
	3, 2, 2, 2, 169, 168, 3, 2, 2, 2, 170, 172, 3, 2, 2, 2, 171, 163, 3, 2,
	2, 2, 172, 175, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2,
	174, 19, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 176, 177, 7, 150, 2, 2, 177,
	21, 3, 2, 2, 2, 178, 239, 5, 36, 19, 2, 179, 239, 5, 42, 22, 2, 180, 182,
	7, 117, 2, 2, 181, 180, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 183, 3,
	2, 2, 2, 183, 239, 5, 24, 13, 2, 184, 239, 5, 44, 23, 2, 185, 239, 7, 149,
	2, 2, 186, 239, 7, 132, 2, 2, 187, 239, 7, 133, 2, 2, 188, 239, 7, 145,
	2, 2, 189, 239, 7, 146, 2, 2, 190, 239, 7, 47, 2, 2, 191, 239, 7, 49, 2,
	2, 192, 239, 7, 50, 2, 2, 193, 239, 7, 51, 2, 2, 194, 239, 7, 52, 2, 2,
	195, 239, 7, 64, 2, 2, 196, 197, 7, 136, 2, 2, 197, 200, 7, 136, 2, 2,
	198, 201, 5, 28, 15, 2, 199, 201, 5, 26, 14, 2, 200, 198, 3, 2, 2, 2, 200,
	199, 3, 2, 2, 2, 201, 207, 3, 2, 2, 2, 202, 205, 7, 134, 2, 2, 203, 206,
	5, 28, 15, 2, 204, 206, 5, 26, 14, 2, 205, 203, 3, 2, 2, 2, 205, 204, 3,
	2, 2, 2, 206, 208, 3, 2, 2, 2, 207, 202, 3, 2, 2, 2, 207, 208, 3, 2, 2,
	2, 208, 209, 3, 2, 2, 2, 209, 227, 7, 137, 2, 2, 210, 211, 7, 134, 2, 2,
	211, 214, 7, 136, 2, 2, 212, 215, 5, 28, 15, 2, 213, 215, 5, 26, 14, 2,
	214, 212, 3, 2, 2, 2, 214, 213, 3, 2, 2, 2, 215, 221, 3, 2, 2, 2, 216,
	219, 7, 134, 2, 2, 217, 220, 5, 28, 15, 2, 218, 220, 5, 26, 14, 2, 219,
	217, 3, 2, 2, 2, 219, 218, 3, 2, 2, 2, 220, 222, 3, 2, 2, 2, 221, 216,
	3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 224, 7, 137,
	2, 2, 224, 226, 3, 2, 2, 2, 225, 210, 3, 2, 2, 2, 226, 229, 3, 2, 2, 2,
	227, 225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 230, 3, 2, 2, 2, 229,
	227, 3, 2, 2, 2, 230, 231, 7, 137, 2, 2, 231, 239, 3, 2, 2, 2, 232, 236,
	5, 26, 14, 2, 233, 234, 7, 136, 2, 2, 234, 235, 7, 143, 2, 2, 235, 237,
	7, 137, 2, 2, 236, 233, 3, 2, 2, 2, 236, 237, 3, 2, 2, 2, 237, 239, 3,
	2, 2, 2, 238, 178, 3, 2, 2, 2, 238, 179, 3, 2, 2, 2, 238, 181, 3, 2, 2,
	2, 238, 184, 3, 2, 2, 2, 238, 185, 3, 2, 2, 2, 238, 186, 3, 2, 2, 2, 238,
	187, 3, 2, 2, 2, 238, 188, 3, 2, 2, 2, 238, 189, 3, 2, 2, 2, 238, 190,
	3, 2, 2, 2, 238, 191, 3, 2, 2, 2, 238, 192, 3, 2, 2, 2, 238, 193, 3, 2,
	2, 2, 238, 194, 3, 2, 2, 2, 238, 195, 3, 2, 2, 2, 238, 196, 3, 2, 2, 2,
	238, 232, 3, 2, 2, 2, 239, 23, 3, 2, 2, 2, 240, 241, 7, 143, 2, 2, 241,
	25, 3, 2, 2, 2, 242, 243, 9, 2, 2, 2, 243, 27, 3, 2, 2, 2, 244, 247, 5,
	32, 17, 2, 245, 247, 7, 126, 2, 2, 246, 244, 3, 2, 2, 2, 246, 245, 3, 2,
	2, 2, 247, 29, 3, 2, 2, 2, 248, 282, 7, 130, 2, 2, 249, 278, 7, 128, 2,
	2, 250, 252, 9, 3, 2, 2, 251, 250, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252,
	253, 3, 2, 2, 2, 253, 259, 7, 143, 2, 2, 254, 256, 7, 49, 2, 2, 255, 257,
	9, 3, 2, 2, 256, 255, 3, 2, 2, 2, 256, 257, 3, 2, 2, 2, 257, 258, 3, 2,
	2, 2, 258, 260, 7, 143, 2, 2, 259, 254, 3, 2, 2, 2, 259, 260, 3, 2, 2,
	2, 260, 275, 3, 2, 2, 2, 261, 263, 7, 134, 2, 2, 262, 264, 9, 3, 2, 2,
	263, 262, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265,
	271, 7, 143, 2, 2, 266, 268, 7, 49, 2, 2, 267, 269, 9, 3, 2, 2, 268, 267,
	3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 272, 7, 143,
	2, 2, 271, 266, 3, 2, 2, 2, 271, 272, 3, 2, 2, 2, 272, 274, 3, 2, 2, 2,
	273, 261, 3, 2, 2, 2, 274, 277, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2, 275,
	276, 3, 2, 2, 2, 276, 279, 3, 2, 2, 2, 277, 275, 3, 2, 2, 2, 278, 251,
	3, 2, 2, 2, 278, 279, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 282, 7, 129,
	2, 2, 281, 248, 3, 2, 2, 2, 281, 249, 3, 2, 2, 2, 282, 31, 3, 2, 2, 2,
	283, 284, 7, 78, 2, 2, 284, 285, 7, 136, 2, 2, 285, 286, 5, 34, 18, 2,
	286, 287, 7, 137, 2, 2, 287, 291, 3, 2, 2, 2, 288, 289, 7, 101, 2, 2, 289,
	291, 7, 78, 2, 2, 290, 283, 3, 2, 2, 2, 290, 288, 3, 2, 2, 2, 291, 33,
	3, 2, 2, 2, 292, 294, 7, 51, 2, 2, 293, 292, 3, 2, 2, 2, 293, 294, 3, 2,
	2, 2, 294, 295, 3, 2, 2, 2, 295, 303, 5, 28, 15, 2, 296, 298, 7, 134, 2,
	2, 297, 299, 7, 51, 2, 2, 298, 297, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299,
	300, 3, 2, 2, 2, 300, 302, 5, 28, 15, 2, 301, 296, 3, 2, 2, 2, 302, 305,
	3, 2, 2, 2, 303, 301, 3, 2, 2, 2, 303, 304, 3, 2, 2, 2, 304, 35, 3, 2,
	2, 2, 305, 303, 3, 2, 2, 2, 306, 309, 5, 38, 20, 2, 307, 309, 7, 95, 2,
	2, 308, 306, 3, 2, 2, 2, 308, 307, 3, 2, 2, 2, 309, 310, 3, 2, 2, 2, 310,
	312, 7, 136, 2, 2, 311, 313, 5, 34, 18, 2, 312, 311, 3, 2, 2, 2, 312, 313,
	3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 315, 7, 137, 2, 2, 315, 37, 3, 2,
	2, 2, 316, 337, 5, 20, 11, 2, 317, 318, 7, 102, 2, 2, 318, 322, 7, 120,
	2, 2, 319, 323, 7, 87, 2, 2, 320, 323, 7, 88, 2, 2, 321, 323, 5, 40, 21,
	2, 322, 319, 3, 2, 2, 2, 322, 320, 3, 2, 2, 2, 322, 321, 3, 2, 2, 2, 323,
	337, 3, 2, 2, 2, 324, 325, 5, 20, 11, 2, 325, 326, 7, 120, 2, 2, 326, 327,
	9, 4, 2, 2, 327, 337, 3, 2, 2, 2, 328, 329, 5, 20, 11, 2, 329, 330, 7,
	144, 2, 2, 330, 331, 9, 4, 2, 2, 331, 337, 3, 2, 2, 2, 332, 333, 5, 20,
	11, 2, 333, 334, 7, 120, 2, 2, 334, 335, 5, 40, 21, 2, 335, 337, 3, 2,
	2, 2, 336, 316, 3, 2, 2, 2, 336, 317, 3, 2, 2, 2, 336, 324, 3, 2, 2, 2,
	336, 328, 3, 2, 2, 2, 336, 332, 3, 2, 2, 2, 337, 39, 3, 2, 2, 2, 338, 351,
	5, 20, 11, 2, 339, 351, 7, 92, 2, 2, 340, 351, 7, 47, 2, 2, 341, 351, 7,
	53, 2, 2, 342, 351, 7, 93, 2, 2, 343, 351, 7, 64, 2, 2, 344, 351, 7, 78,
	2, 2, 345, 351, 7, 65, 2, 2, 346, 351, 7, 94, 2, 2, 347, 351, 7, 95, 2,
	2, 348, 351, 7, 146, 2, 2, 349, 351, 7, 45, 2, 2, 350, 338, 3, 2, 2, 2,
	350, 339, 3, 2, 2, 2, 350, 340, 3, 2, 2, 2, 350, 341, 3, 2, 2, 2, 350,
	342, 3, 2, 2, 2, 350, 343, 3, 2, 2, 2, 350, 344, 3, 2, 2, 2, 350, 345,
	3, 2, 2, 2, 350, 346, 3, 2, 2, 2, 350, 347, 3, 2, 2, 2, 350, 348, 3, 2,
	2, 2, 350, 349, 3, 2, 2, 2, 351, 41, 3, 2, 2, 2, 352, 353, 5, 38, 20, 2,
	353, 43, 3, 2, 2, 2, 354, 355, 9, 5, 2, 2, 355, 45, 3, 2, 2, 2, 53, 47,
	52, 59, 65, 70, 77, 87, 96, 100, 104, 118, 124, 127, 138, 144, 148, 150,
	153, 158, 161, 169, 173, 181, 200, 205, 207, 214, 219, 221, 227, 236, 238,
	246, 251, 256, 259, 263, 268, 271, 275, 278, 281, 290, 293, 298, 303, 308,
	312, 322, 336, 350,
}
var literalNames = []string{
	"", "'TABLE'", "'COLUMN'", "'RETRIEVE'", "'PBSELECT'", "'VERSION'", "'ARGUMENTS'",
	"'SORT'", "'ANY'", "'BLOB'", "'BOOLEAN'", "'BYTE'", "'CHARACTER'", "'CHAR'",
	"'DATE'", "'DATETIME'", "'DECIMAL'", "'DEC'", "'DOUBLE'", "'INTEGER'",
	"'INT'", "'LONG'", "'LONGLONG'", "'REAL'", "'STRING'", "'TIME'", "'UNSIGNEDINTEGER'",
	"'UINT'", "'UNSIGNEDLONG'", "'ULONG'", "'WINDOW'", "'TRUE'", "'FALSE'",
	"'GLOBAL'", "'SHARED'", "'END'", "'INDIRECT'", "'VARIABLES'", "'FORWARD'",
	"'PUBLIC'", "'PRIVATE'", "'FUNCTION'", "'SUBROUTINE'", "'READONLY'", "'PROTOTYPES'",
	"'TYPE'", "'ON'", "'TO'", "'FROM'", "'REF'", "'NULL'", "'UPDATE'", "'CASE'",
	"'DYNAMIC'", "'WITHIN'", "'PRIVATEWRITE'", "'PROTECTED'", "'PRIVATEREAD'",
	"'PROTECTEDREAD'", "'PROTECTEDWRITE'", "'LOCAL'", "'EVENT'", "'OPEN'",
	"'GOTO'", "'ELSE'", "'IF'", "'THEN'", "'ELSEIF'", "'TRY'", "'EXIT'", "'CHOOSE'",
	"'IS'", "'CONTINUE'", "'DO'", "'WHILE'", "'FOR'", "'CLOSE'", "'NEXT'",
	"'LOOP'", "'UNTIL'", "'STEP'", "'CATCH'", "'FINALLY'", "'THROW'", "'RELEASE'",
	"'CREATE'", "'DESTROY'", "'USING'", "'POST'", "'TRIGGER'", "'SELECT'",
	"'DELETE'", "'INSERT'", "'DESCRIBE'", "'RETURN'", "'OR'", "'AND'", "'NOT'",
	"'CALL'", "'HALT'", "'SUPER'", "'LIBRARY'", "'SYSTEM'", "'RPCFUNC'", "'ALIAS'",
	"'THROWS'", "'AUTOINSTANTIATE'", "'DESCRIPTOR'", "'='", "'>'", "'>='",
	"'<'", "'<='", "'<>'", "'+'", "'-'", "'+='", "'-='", "'::'", "'*'", "'/'",
	"'*='", "'/='", "'^'", "'{'", "'}'", "'['", "']'", "'[]'", "'`'", "", "",
	"','", "';'", "'('", "')'", "':'", "'\"'", "'???'", "'||'", "'...'", "",
	"'.'",
}
var symbolicNames = []string{
	"", "TABLE", "COLUMN", "RETRIEVE", "PBSELECT", "VERSION", "ARGUMENTS",
	"SORT", "ANY", "BLOB", "BOOLEAN", "BYTE", "CHARACTER", "CHAR", "DATE_TYPE",
	"DATETIME", "DECIMAL", "DEC", "DOUBLE", "INTEGER", "INT", "LONG", "LONGLONG",
	"REAL", "STRING", "TIME_TYPE", "UNSIGNEDINTEGER", "UINT", "UNSIGNEDLONG",
	"ULONG", "WINDOW", "TRUE", "FALSE", "GLOBAL", "SHARED", "END", "INDIRECT",
	"VARIABLES", "FORWARD", "PUBLIC", "PRIVATE", "FUNCTION", "SUBROUTINE",
	"READONLY", "PROTOTYPES", "TYPE", "ON", "TO", "FROM", "REF", "NULL_", "UPDATE",
	"CASE", "DYNAMIC", "WITHIN", "PRIVATEWRITE", "PROTECTED", "PRIVATEREAD",
	"PROTECTEDREAD", "PROTECTEDWRITE", "LOCAL", "EVENT", "OPEN", "GOTO", "ELSE",
	"IF", "THEN", "ELSEIF", "TRY", "EXIT", "CHOOSE", "IS", "CONTINUE", "DO",
	"WHILE", "FOR", "CLOSE", "NEXT", "LOOP", "UNTIL", "STEP", "CATCH", "FINALLY",
	"THROW", "RELEASE", "CREATE", "DESTROY", "USING", "POST", "TRIGGER", "SELECT",
	"DELETE", "INSERT", "DESCRIBE", "RETURN", "OR", "AND", "NOT", "CALL", "HALT",
	"SUPER", "LIBRARY", "SYSTEM", "RPCFUNC", "ALIAS", "THROWS", "AUTOINSTANTIATE",
	"DESCRIPTOR", "EQ", "GT", "GTE", "LT", "LTE", "GTLT", "PLUS", "MINUS",
	"PLUSEQ", "MINUSEQ", "COLONCOLON", "MULT", "DIV", "MULTEQ", "DIVEQ", "CARAT",
	"LCURLY", "RCURLY", "LBRACE", "RBRACE", "BRACES", "TICK", "DQUOTED_STRING",
	"QUOTED_STRING", "COMMA", "SEMI", "LPAREN", "RPAREN", "COLON", "DQUOTE",
	"TQ", "DOUBLE_PIPE", "DOTDOTDOT", "NUMBER", "DOT", "DATE", "TIME", "BINDPAR",
	"EXPORT_HEADER", "ENUM", "ID", "RET_LIT", "ARGS_LIT", "SORT_LIT", "LINE_CONTINUATION",
	"SL_COMMENT", "ML_COMMENT", "WS",
}

var ruleNames = []string{
	"start_rule", "header_rule", "datawindow_rule", "datawindow_property",
	"table_attribute", "column_attribute", "retrieve_attribute", "datawindow_property_attribute_sub",
	"attribute_name", "identifier_name", "attribute_value", "numeric_atom",
	"dataTypeSub", "expression", "array_decl_sub", "close_call_sub", "expression_list",
	"atom_sub_call1", "identifier", "identifier_name_ex", "atom_sub_member1",
	"boolean_atom",
}

type PowerBuilderDWParser struct {
	*antlr.BaseParser
}

// NewPowerBuilderDWParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *PowerBuilderDWParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPowerBuilderDWParser(input antlr.TokenStream) *PowerBuilderDWParser {
	this := new(PowerBuilderDWParser)
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
	this.GrammarFileName = "PowerBuilderDWParser.g4"

	return this
}

// PowerBuilderDWParser tokens.
const (
	PowerBuilderDWParserEOF               = antlr.TokenEOF
	PowerBuilderDWParserTABLE             = 1
	PowerBuilderDWParserCOLUMN            = 2
	PowerBuilderDWParserRETRIEVE          = 3
	PowerBuilderDWParserPBSELECT          = 4
	PowerBuilderDWParserVERSION           = 5
	PowerBuilderDWParserARGUMENTS         = 6
	PowerBuilderDWParserSORT              = 7
	PowerBuilderDWParserANY               = 8
	PowerBuilderDWParserBLOB              = 9
	PowerBuilderDWParserBOOLEAN           = 10
	PowerBuilderDWParserBYTE              = 11
	PowerBuilderDWParserCHARACTER         = 12
	PowerBuilderDWParserCHAR              = 13
	PowerBuilderDWParserDATE_TYPE         = 14
	PowerBuilderDWParserDATETIME          = 15
	PowerBuilderDWParserDECIMAL           = 16
	PowerBuilderDWParserDEC               = 17
	PowerBuilderDWParserDOUBLE            = 18
	PowerBuilderDWParserINTEGER           = 19
	PowerBuilderDWParserINT               = 20
	PowerBuilderDWParserLONG              = 21
	PowerBuilderDWParserLONGLONG          = 22
	PowerBuilderDWParserREAL              = 23
	PowerBuilderDWParserSTRING            = 24
	PowerBuilderDWParserTIME_TYPE         = 25
	PowerBuilderDWParserUNSIGNEDINTEGER   = 26
	PowerBuilderDWParserUINT              = 27
	PowerBuilderDWParserUNSIGNEDLONG      = 28
	PowerBuilderDWParserULONG             = 29
	PowerBuilderDWParserWINDOW            = 30
	PowerBuilderDWParserTRUE              = 31
	PowerBuilderDWParserFALSE             = 32
	PowerBuilderDWParserGLOBAL            = 33
	PowerBuilderDWParserSHARED            = 34
	PowerBuilderDWParserEND               = 35
	PowerBuilderDWParserINDIRECT          = 36
	PowerBuilderDWParserVARIABLES         = 37
	PowerBuilderDWParserFORWARD           = 38
	PowerBuilderDWParserPUBLIC            = 39
	PowerBuilderDWParserPRIVATE           = 40
	PowerBuilderDWParserFUNCTION          = 41
	PowerBuilderDWParserSUBROUTINE        = 42
	PowerBuilderDWParserREADONLY          = 43
	PowerBuilderDWParserPROTOTYPES        = 44
	PowerBuilderDWParserTYPE              = 45
	PowerBuilderDWParserON                = 46
	PowerBuilderDWParserTO                = 47
	PowerBuilderDWParserFROM              = 48
	PowerBuilderDWParserREF               = 49
	PowerBuilderDWParserNULL_             = 50
	PowerBuilderDWParserUPDATE            = 51
	PowerBuilderDWParserCASE              = 52
	PowerBuilderDWParserDYNAMIC           = 53
	PowerBuilderDWParserWITHIN            = 54
	PowerBuilderDWParserPRIVATEWRITE      = 55
	PowerBuilderDWParserPROTECTED         = 56
	PowerBuilderDWParserPRIVATEREAD       = 57
	PowerBuilderDWParserPROTECTEDREAD     = 58
	PowerBuilderDWParserPROTECTEDWRITE    = 59
	PowerBuilderDWParserLOCAL             = 60
	PowerBuilderDWParserEVENT             = 61
	PowerBuilderDWParserOPEN              = 62
	PowerBuilderDWParserGOTO              = 63
	PowerBuilderDWParserELSE              = 64
	PowerBuilderDWParserIF                = 65
	PowerBuilderDWParserTHEN              = 66
	PowerBuilderDWParserELSEIF            = 67
	PowerBuilderDWParserTRY               = 68
	PowerBuilderDWParserEXIT              = 69
	PowerBuilderDWParserCHOOSE            = 70
	PowerBuilderDWParserIS                = 71
	PowerBuilderDWParserCONTINUE          = 72
	PowerBuilderDWParserDO                = 73
	PowerBuilderDWParserWHILE             = 74
	PowerBuilderDWParserFOR               = 75
	PowerBuilderDWParserCLOSE             = 76
	PowerBuilderDWParserNEXT              = 77
	PowerBuilderDWParserLOOP              = 78
	PowerBuilderDWParserUNTIL             = 79
	PowerBuilderDWParserSTEP              = 80
	PowerBuilderDWParserCATCH             = 81
	PowerBuilderDWParserFINALLY           = 82
	PowerBuilderDWParserTHROW             = 83
	PowerBuilderDWParserRELEASE           = 84
	PowerBuilderDWParserCREATE            = 85
	PowerBuilderDWParserDESTROY           = 86
	PowerBuilderDWParserUSING             = 87
	PowerBuilderDWParserPOST              = 88
	PowerBuilderDWParserTRIGGER           = 89
	PowerBuilderDWParserSELECT            = 90
	PowerBuilderDWParserDELETE            = 91
	PowerBuilderDWParserINSERT            = 92
	PowerBuilderDWParserDESCRIBE          = 93
	PowerBuilderDWParserRETURN            = 94
	PowerBuilderDWParserOR                = 95
	PowerBuilderDWParserAND               = 96
	PowerBuilderDWParserNOT               = 97
	PowerBuilderDWParserCALL              = 98
	PowerBuilderDWParserHALT              = 99
	PowerBuilderDWParserSUPER             = 100
	PowerBuilderDWParserLIBRARY           = 101
	PowerBuilderDWParserSYSTEM            = 102
	PowerBuilderDWParserRPCFUNC           = 103
	PowerBuilderDWParserALIAS             = 104
	PowerBuilderDWParserTHROWS            = 105
	PowerBuilderDWParserAUTOINSTANTIATE   = 106
	PowerBuilderDWParserDESCRIPTOR        = 107
	PowerBuilderDWParserEQ                = 108
	PowerBuilderDWParserGT                = 109
	PowerBuilderDWParserGTE               = 110
	PowerBuilderDWParserLT                = 111
	PowerBuilderDWParserLTE               = 112
	PowerBuilderDWParserGTLT              = 113
	PowerBuilderDWParserPLUS              = 114
	PowerBuilderDWParserMINUS             = 115
	PowerBuilderDWParserPLUSEQ            = 116
	PowerBuilderDWParserMINUSEQ           = 117
	PowerBuilderDWParserCOLONCOLON        = 118
	PowerBuilderDWParserMULT              = 119
	PowerBuilderDWParserDIV               = 120
	PowerBuilderDWParserMULTEQ            = 121
	PowerBuilderDWParserDIVEQ             = 122
	PowerBuilderDWParserCARAT             = 123
	PowerBuilderDWParserLCURLY            = 124
	PowerBuilderDWParserRCURLY            = 125
	PowerBuilderDWParserLBRACE            = 126
	PowerBuilderDWParserRBRACE            = 127
	PowerBuilderDWParserBRACES            = 128
	PowerBuilderDWParserTICK              = 129
	PowerBuilderDWParserDQUOTED_STRING    = 130
	PowerBuilderDWParserQUOTED_STRING     = 131
	PowerBuilderDWParserCOMMA             = 132
	PowerBuilderDWParserSEMI              = 133
	PowerBuilderDWParserLPAREN            = 134
	PowerBuilderDWParserRPAREN            = 135
	PowerBuilderDWParserCOLON             = 136
	PowerBuilderDWParserDQUOTE            = 137
	PowerBuilderDWParserTQ                = 138
	PowerBuilderDWParserDOUBLE_PIPE       = 139
	PowerBuilderDWParserDOTDOTDOT         = 140
	PowerBuilderDWParserNUMBER            = 141
	PowerBuilderDWParserDOT               = 142
	PowerBuilderDWParserDATE              = 143
	PowerBuilderDWParserTIME              = 144
	PowerBuilderDWParserBINDPAR           = 145
	PowerBuilderDWParserEXPORT_HEADER     = 146
	PowerBuilderDWParserENUM              = 147
	PowerBuilderDWParserID                = 148
	PowerBuilderDWParserRET_LIT           = 149
	PowerBuilderDWParserARGS_LIT          = 150
	PowerBuilderDWParserSORT_LIT          = 151
	PowerBuilderDWParserLINE_CONTINUATION = 152
	PowerBuilderDWParserSL_COMMENT        = 153
	PowerBuilderDWParserML_COMMENT        = 154
	PowerBuilderDWParserWS                = 155
)

// PowerBuilderDWParser rules.
const (
	PowerBuilderDWParserRULE_start_rule                        = 0
	PowerBuilderDWParserRULE_header_rule                       = 1
	PowerBuilderDWParserRULE_datawindow_rule                   = 2
	PowerBuilderDWParserRULE_datawindow_property               = 3
	PowerBuilderDWParserRULE_table_attribute                   = 4
	PowerBuilderDWParserRULE_column_attribute                  = 5
	PowerBuilderDWParserRULE_retrieve_attribute                = 6
	PowerBuilderDWParserRULE_datawindow_property_attribute_sub = 7
	PowerBuilderDWParserRULE_attribute_name                    = 8
	PowerBuilderDWParserRULE_identifier_name                   = 9
	PowerBuilderDWParserRULE_attribute_value                   = 10
	PowerBuilderDWParserRULE_numeric_atom                      = 11
	PowerBuilderDWParserRULE_dataTypeSub                       = 12
	PowerBuilderDWParserRULE_expression                        = 13
	PowerBuilderDWParserRULE_array_decl_sub                    = 14
	PowerBuilderDWParserRULE_close_call_sub                    = 15
	PowerBuilderDWParserRULE_expression_list                   = 16
	PowerBuilderDWParserRULE_atom_sub_call1                    = 17
	PowerBuilderDWParserRULE_identifier                        = 18
	PowerBuilderDWParserRULE_identifier_name_ex                = 19
	PowerBuilderDWParserRULE_atom_sub_member1                  = 20
	PowerBuilderDWParserRULE_boolean_atom                      = 21
)

// IStart_ruleContext is an interface to support dynamic dispatch.
type IStart_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStart_ruleContext differentiates from other interfaces.
	IsStart_ruleContext()
}

type Start_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStart_ruleContext() *Start_ruleContext {
	var p = new(Start_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_start_rule
	return p
}

func (*Start_ruleContext) IsStart_ruleContext() {}

func NewStart_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Start_ruleContext {
	var p = new(Start_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_start_rule

	return p
}

func (s *Start_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Start_ruleContext) EOF() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserEOF, 0)
}

func (s *Start_ruleContext) Header_rule() IHeader_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHeader_ruleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHeader_ruleContext)
}

func (s *Start_ruleContext) AllDatawindow_rule() []IDatawindow_ruleContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatawindow_ruleContext)(nil)).Elem())
	var tst = make([]IDatawindow_ruleContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatawindow_ruleContext)
		}
	}

	return tst
}

func (s *Start_ruleContext) Datawindow_rule(i int) IDatawindow_ruleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatawindow_ruleContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatawindow_ruleContext)
}

func (s *Start_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Start_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Start_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterStart_rule(s)
	}
}

func (s *Start_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitStart_rule(s)
	}
}

func (p *PowerBuilderDWParser) Start_rule() (localctx IStart_ruleContext) {
	this := p
	_ = this

	localctx = NewStart_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PowerBuilderDWParserRULE_start_rule)
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
	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(44)
			p.Header_rule()
		}

	}
	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == PowerBuilderDWParserTABLE || _la == PowerBuilderDWParserCOLUMN || _la == PowerBuilderDWParserTYPE || _la == PowerBuilderDWParserUPDATE || _la == PowerBuilderDWParserID {
		{
			p.SetState(47)
			p.Datawindow_rule()
		}

		p.SetState(50)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
		p.Match(PowerBuilderDWParserEOF)
	}

	return localctx
}

// IHeader_ruleContext is an interface to support dynamic dispatch.
type IHeader_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHeader_ruleContext differentiates from other interfaces.
	IsHeader_ruleContext()
}

type Header_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHeader_ruleContext() *Header_ruleContext {
	var p = new(Header_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_header_rule
	return p
}

func (*Header_ruleContext) IsHeader_ruleContext() {}

func NewHeader_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Header_ruleContext {
	var p = new(Header_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_header_rule

	return p
}

func (s *Header_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Header_ruleContext) AllEXPORT_HEADER() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserEXPORT_HEADER)
}

func (s *Header_ruleContext) EXPORT_HEADER(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserEXPORT_HEADER, i)
}

func (s *Header_ruleContext) RELEASE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserRELEASE, 0)
}

func (s *Header_ruleContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserNUMBER, 0)
}

func (s *Header_ruleContext) SEMI() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserSEMI, 0)
}

func (s *Header_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Header_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Header_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterHeader_rule(s)
	}
}

func (s *Header_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitHeader_rule(s)
	}
}

func (p *PowerBuilderDWParser) Header_rule() (localctx IHeader_ruleContext) {
	this := p
	_ = this

	localctx = NewHeader_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PowerBuilderDWParserRULE_header_rule)
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
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PowerBuilderDWParserEXPORT_HEADER {
		{
			p.SetState(54)
			p.Match(PowerBuilderDWParserEXPORT_HEADER)
		}

		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerBuilderDWParserRELEASE {
		{
			p.SetState(60)
			p.Match(PowerBuilderDWParserRELEASE)
		}
		{
			p.SetState(61)
			p.Match(PowerBuilderDWParserNUMBER)
		}
		{
			p.SetState(62)
			p.Match(PowerBuilderDWParserSEMI)
		}

	}

	return localctx
}

// IDatawindow_ruleContext is an interface to support dynamic dispatch.
type IDatawindow_ruleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatawindow_ruleContext differentiates from other interfaces.
	IsDatawindow_ruleContext()
}

type Datawindow_ruleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatawindow_ruleContext() *Datawindow_ruleContext {
	var p = new(Datawindow_ruleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_datawindow_rule
	return p
}

func (*Datawindow_ruleContext) IsDatawindow_ruleContext() {}

func NewDatawindow_ruleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Datawindow_ruleContext {
	var p = new(Datawindow_ruleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_datawindow_rule

	return p
}

func (s *Datawindow_ruleContext) GetParser() antlr.Parser { return s.parser }

func (s *Datawindow_ruleContext) AllDatawindow_property() []IDatawindow_propertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatawindow_propertyContext)(nil)).Elem())
	var tst = make([]IDatawindow_propertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatawindow_propertyContext)
		}
	}

	return tst
}

func (s *Datawindow_ruleContext) Datawindow_property(i int) IDatawindow_propertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatawindow_propertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatawindow_propertyContext)
}

func (s *Datawindow_ruleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Datawindow_ruleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Datawindow_ruleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterDatawindow_rule(s)
	}
}

func (s *Datawindow_ruleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitDatawindow_rule(s)
	}
}

func (p *PowerBuilderDWParser) Datawindow_rule() (localctx IDatawindow_ruleContext) {
	this := p
	_ = this

	localctx = NewDatawindow_ruleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PowerBuilderDWParserRULE_datawindow_rule)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(65)
				p.Datawindow_property()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IDatawindow_propertyContext is an interface to support dynamic dispatch.
type IDatawindow_propertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatawindow_propertyContext differentiates from other interfaces.
	IsDatawindow_propertyContext()
}

type Datawindow_propertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatawindow_propertyContext() *Datawindow_propertyContext {
	var p = new(Datawindow_propertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_datawindow_property
	return p
}

func (*Datawindow_propertyContext) IsDatawindow_propertyContext() {}

func NewDatawindow_propertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Datawindow_propertyContext {
	var p = new(Datawindow_propertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_datawindow_property

	return p
}

func (s *Datawindow_propertyContext) GetParser() antlr.Parser { return s.parser }

func (s *Datawindow_propertyContext) Attribute_name() IAttribute_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribute_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttribute_nameContext)
}

func (s *Datawindow_propertyContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserLPAREN, 0)
}

func (s *Datawindow_propertyContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserRPAREN, 0)
}

func (s *Datawindow_propertyContext) AllDatawindow_property_attribute_sub() []IDatawindow_property_attribute_subContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatawindow_property_attribute_subContext)(nil)).Elem())
	var tst = make([]IDatawindow_property_attribute_subContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatawindow_property_attribute_subContext)
		}
	}

	return tst
}

func (s *Datawindow_propertyContext) Datawindow_property_attribute_sub(i int) IDatawindow_property_attribute_subContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatawindow_property_attribute_subContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatawindow_property_attribute_subContext)
}

func (s *Datawindow_propertyContext) TABLE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserTABLE, 0)
}

func (s *Datawindow_propertyContext) AllTable_attribute() []ITable_attributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITable_attributeContext)(nil)).Elem())
	var tst = make([]ITable_attributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITable_attributeContext)
		}
	}

	return tst
}

func (s *Datawindow_propertyContext) Table_attribute(i int) ITable_attributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITable_attributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITable_attributeContext)
}

func (s *Datawindow_propertyContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCOLUMN, 0)
}

func (s *Datawindow_propertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Datawindow_propertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Datawindow_propertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterDatawindow_property(s)
	}
}

func (s *Datawindow_propertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitDatawindow_property(s)
	}
}

func (p *PowerBuilderDWParser) Datawindow_property() (localctx IDatawindow_propertyContext) {
	this := p
	_ = this

	localctx = NewDatawindow_propertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, PowerBuilderDWParserRULE_datawindow_property)
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

	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PowerBuilderDWParserTYPE, PowerBuilderDWParserUPDATE, PowerBuilderDWParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(70)
			p.Attribute_name()
		}
		{
			p.SetState(71)
			p.Match(PowerBuilderDWParserLPAREN)
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(PowerBuilderDWParserTYPE-45))|(1<<(PowerBuilderDWParserNULL_-45))|(1<<(PowerBuilderDWParserUPDATE-45)))) != 0) || (((_la-130)&-(0x1f+1)) == 0 && ((1<<uint((_la-130)))&((1<<(PowerBuilderDWParserDQUOTED_STRING-130))|(1<<(PowerBuilderDWParserNUMBER-130))|(1<<(PowerBuilderDWParserDATE-130))|(1<<(PowerBuilderDWParserTIME-130))|(1<<(PowerBuilderDWParserID-130)))) != 0) {
			{
				p.SetState(72)
				p.Datawindow_property_attribute_sub()
			}

			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(78)
			p.Match(PowerBuilderDWParserRPAREN)
		}

	case PowerBuilderDWParserTABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.Match(PowerBuilderDWParserTABLE)
		}
		{
			p.SetState(81)
			p.Match(PowerBuilderDWParserLPAREN)
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == PowerBuilderDWParserCOLUMN || _la == PowerBuilderDWParserRET_LIT {
			{
				p.SetState(82)
				p.Table_attribute()
			}

			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(87)
			p.Match(PowerBuilderDWParserRPAREN)
		}

	case PowerBuilderDWParserCOLUMN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)
			p.Match(PowerBuilderDWParserCOLUMN)
		}
		{
			p.SetState(90)
			p.Match(PowerBuilderDWParserLPAREN)
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(PowerBuilderDWParserTYPE-45))|(1<<(PowerBuilderDWParserNULL_-45))|(1<<(PowerBuilderDWParserUPDATE-45)))) != 0) || (((_la-130)&-(0x1f+1)) == 0 && ((1<<uint((_la-130)))&((1<<(PowerBuilderDWParserDQUOTED_STRING-130))|(1<<(PowerBuilderDWParserNUMBER-130))|(1<<(PowerBuilderDWParserDATE-130))|(1<<(PowerBuilderDWParserTIME-130))|(1<<(PowerBuilderDWParserID-130)))) != 0) {
			{
				p.SetState(91)
				p.Datawindow_property_attribute_sub()
			}

			p.SetState(96)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)
			p.Match(PowerBuilderDWParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITable_attributeContext is an interface to support dynamic dispatch.
type ITable_attributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTable_attributeContext differentiates from other interfaces.
	IsTable_attributeContext()
}

type Table_attributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTable_attributeContext() *Table_attributeContext {
	var p = new(Table_attributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_table_attribute
	return p
}

func (*Table_attributeContext) IsTable_attributeContext() {}

func NewTable_attributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Table_attributeContext {
	var p = new(Table_attributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_table_attribute

	return p
}

func (s *Table_attributeContext) GetParser() antlr.Parser { return s.parser }

func (s *Table_attributeContext) Column_attribute() IColumn_attributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumn_attributeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumn_attributeContext)
}

func (s *Table_attributeContext) Retrieve_attribute() IRetrieve_attributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRetrieve_attributeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRetrieve_attributeContext)
}

func (s *Table_attributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Table_attributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Table_attributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterTable_attribute(s)
	}
}

func (s *Table_attributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitTable_attribute(s)
	}
}

func (p *PowerBuilderDWParser) Table_attribute() (localctx ITable_attributeContext) {
	this := p
	_ = this

	localctx = NewTable_attributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PowerBuilderDWParserRULE_table_attribute)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	case PowerBuilderDWParserCOLUMN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.Column_attribute()
		}

	case PowerBuilderDWParserRET_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.Retrieve_attribute()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IColumn_attributeContext is an interface to support dynamic dispatch.
type IColumn_attributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumn_attributeContext differentiates from other interfaces.
	IsColumn_attributeContext()
}

type Column_attributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumn_attributeContext() *Column_attributeContext {
	var p = new(Column_attributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_column_attribute
	return p
}

func (*Column_attributeContext) IsColumn_attributeContext() {}

func NewColumn_attributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Column_attributeContext {
	var p = new(Column_attributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_column_attribute

	return p
}

func (s *Column_attributeContext) GetParser() antlr.Parser { return s.parser }

func (s *Column_attributeContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCOLUMN, 0)
}

func (s *Column_attributeContext) AllEQ() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserEQ)
}

func (s *Column_attributeContext) EQ(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserEQ, i)
}

func (s *Column_attributeContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserLPAREN)
}

func (s *Column_attributeContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserLPAREN, i)
}

func (s *Column_attributeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserTYPE, 0)
}

func (s *Column_attributeContext) DataTypeSub() IDataTypeSubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeSubContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDataTypeSubContext)
}

func (s *Column_attributeContext) Numeric_atom() INumeric_atomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumeric_atomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumeric_atomContext)
}

func (s *Column_attributeContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserRPAREN)
}

func (s *Column_attributeContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserRPAREN, i)
}

func (s *Column_attributeContext) AllDatawindow_property_attribute_sub() []IDatawindow_property_attribute_subContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatawindow_property_attribute_subContext)(nil)).Elem())
	var tst = make([]IDatawindow_property_attribute_subContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatawindow_property_attribute_subContext)
		}
	}

	return tst
}

func (s *Column_attributeContext) Datawindow_property_attribute_sub(i int) IDatawindow_property_attribute_subContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatawindow_property_attribute_subContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatawindow_property_attribute_subContext)
}

func (s *Column_attributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Column_attributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Column_attributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterColumn_attribute(s)
	}
}

func (s *Column_attributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitColumn_attribute(s)
	}
}

func (p *PowerBuilderDWParser) Column_attribute() (localctx IColumn_attributeContext) {
	this := p
	_ = this

	localctx = NewColumn_attributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PowerBuilderDWParserRULE_column_attribute)
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
		p.SetState(104)
		p.Match(PowerBuilderDWParserCOLUMN)
	}
	{
		p.SetState(105)
		p.Match(PowerBuilderDWParserEQ)
	}
	{
		p.SetState(106)
		p.Match(PowerBuilderDWParserLPAREN)
	}
	{
		p.SetState(107)
		p.Match(PowerBuilderDWParserTYPE)
	}
	{
		p.SetState(108)
		p.Match(PowerBuilderDWParserEQ)
	}
	{
		p.SetState(109)
		p.DataTypeSub()
	}
	{
		p.SetState(110)
		p.Match(PowerBuilderDWParserLPAREN)
	}
	{
		p.SetState(111)
		p.Numeric_atom()
	}
	{
		p.SetState(112)
		p.Match(PowerBuilderDWParserRPAREN)
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(PowerBuilderDWParserTYPE-45))|(1<<(PowerBuilderDWParserNULL_-45))|(1<<(PowerBuilderDWParserUPDATE-45)))) != 0) || (((_la-130)&-(0x1f+1)) == 0 && ((1<<uint((_la-130)))&((1<<(PowerBuilderDWParserDQUOTED_STRING-130))|(1<<(PowerBuilderDWParserNUMBER-130))|(1<<(PowerBuilderDWParserDATE-130))|(1<<(PowerBuilderDWParserTIME-130))|(1<<(PowerBuilderDWParserID-130)))) != 0) {
		{
			p.SetState(113)
			p.Datawindow_property_attribute_sub()
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(118)
		p.Match(PowerBuilderDWParserRPAREN)
	}

	return localctx
}

// IRetrieve_attributeContext is an interface to support dynamic dispatch.
type IRetrieve_attributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRetrieve_attributeContext differentiates from other interfaces.
	IsRetrieve_attributeContext()
}

type Retrieve_attributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRetrieve_attributeContext() *Retrieve_attributeContext {
	var p = new(Retrieve_attributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_retrieve_attribute
	return p
}

func (*Retrieve_attributeContext) IsRetrieve_attributeContext() {}

func NewRetrieve_attributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Retrieve_attributeContext {
	var p = new(Retrieve_attributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_retrieve_attribute

	return p
}

func (s *Retrieve_attributeContext) GetParser() antlr.Parser { return s.parser }

func (s *Retrieve_attributeContext) RET_LIT() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserRET_LIT, 0)
}

func (s *Retrieve_attributeContext) ARGS_LIT() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserARGS_LIT, 0)
}

func (s *Retrieve_attributeContext) SORT_LIT() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserSORT_LIT, 0)
}

func (s *Retrieve_attributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Retrieve_attributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Retrieve_attributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterRetrieve_attribute(s)
	}
}

func (s *Retrieve_attributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitRetrieve_attribute(s)
	}
}

func (p *PowerBuilderDWParser) Retrieve_attribute() (localctx IRetrieve_attributeContext) {
	this := p
	_ = this

	localctx = NewRetrieve_attributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PowerBuilderDWParserRULE_retrieve_attribute)
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
		p.SetState(120)
		p.Match(PowerBuilderDWParserRET_LIT)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerBuilderDWParserARGS_LIT {
		{
			p.SetState(121)
			p.Match(PowerBuilderDWParserARGS_LIT)
		}

	}
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerBuilderDWParserSORT_LIT {
		{
			p.SetState(124)
			p.Match(PowerBuilderDWParserSORT_LIT)
		}

	}

	return localctx
}

// IDatawindow_property_attribute_subContext is an interface to support dynamic dispatch.
type IDatawindow_property_attribute_subContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetEq returns the eq token.
	GetEq() antlr.Token

	// SetEq sets the eq token.
	SetEq(antlr.Token)

	// IsDatawindow_property_attribute_subContext differentiates from other interfaces.
	IsDatawindow_property_attribute_subContext()
}

type Datawindow_property_attribute_subContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	eq     antlr.Token
}

func NewEmptyDatawindow_property_attribute_subContext() *Datawindow_property_attribute_subContext {
	var p = new(Datawindow_property_attribute_subContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_datawindow_property_attribute_sub
	return p
}

func (*Datawindow_property_attribute_subContext) IsDatawindow_property_attribute_subContext() {}

func NewDatawindow_property_attribute_subContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Datawindow_property_attribute_subContext {
	var p = new(Datawindow_property_attribute_subContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_datawindow_property_attribute_sub

	return p
}

func (s *Datawindow_property_attribute_subContext) GetParser() antlr.Parser { return s.parser }

func (s *Datawindow_property_attribute_subContext) GetEq() antlr.Token { return s.eq }

func (s *Datawindow_property_attribute_subContext) SetEq(v antlr.Token) { s.eq = v }

func (s *Datawindow_property_attribute_subContext) NULL_() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserNULL_, 0)
}

func (s *Datawindow_property_attribute_subContext) Numeric_atom() INumeric_atomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumeric_atomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumeric_atomContext)
}

func (s *Datawindow_property_attribute_subContext) DQUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDQUOTED_STRING, 0)
}

func (s *Datawindow_property_attribute_subContext) DATE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDATE, 0)
}

func (s *Datawindow_property_attribute_subContext) TIME() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserTIME, 0)
}

func (s *Datawindow_property_attribute_subContext) Attribute_name() IAttribute_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribute_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttribute_nameContext)
}

func (s *Datawindow_property_attribute_subContext) EQ() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserEQ, 0)
}

func (s *Datawindow_property_attribute_subContext) COMMA() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCOMMA, 0)
}

func (s *Datawindow_property_attribute_subContext) Attribute_value() IAttribute_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttribute_valueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAttribute_valueContext)
}

func (s *Datawindow_property_attribute_subContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserLPAREN, 0)
}

func (s *Datawindow_property_attribute_subContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserRPAREN, 0)
}

func (s *Datawindow_property_attribute_subContext) Array_decl_sub() IArray_decl_subContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArray_decl_subContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArray_decl_subContext)
}

func (s *Datawindow_property_attribute_subContext) AllDatawindow_property_attribute_sub() []IDatawindow_property_attribute_subContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDatawindow_property_attribute_subContext)(nil)).Elem())
	var tst = make([]IDatawindow_property_attribute_subContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDatawindow_property_attribute_subContext)
		}
	}

	return tst
}

func (s *Datawindow_property_attribute_subContext) Datawindow_property_attribute_sub(i int) IDatawindow_property_attribute_subContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatawindow_property_attribute_subContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDatawindow_property_attribute_subContext)
}

func (s *Datawindow_property_attribute_subContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Datawindow_property_attribute_subContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Datawindow_property_attribute_subContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterDatawindow_property_attribute_sub(s)
	}
}

func (s *Datawindow_property_attribute_subContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitDatawindow_property_attribute_sub(s)
	}
}

func (p *PowerBuilderDWParser) Datawindow_property_attribute_sub() (localctx IDatawindow_property_attribute_subContext) {
	this := p
	_ = this

	localctx = NewDatawindow_property_attribute_subContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PowerBuilderDWParserRULE_datawindow_property_attribute_sub)
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
	p.SetState(148)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PowerBuilderDWParserNULL_:
		{
			p.SetState(127)
			p.Match(PowerBuilderDWParserNULL_)
		}

	case PowerBuilderDWParserNUMBER:
		{
			p.SetState(128)
			p.Numeric_atom()
		}

	case PowerBuilderDWParserDQUOTED_STRING:
		{
			p.SetState(129)
			p.Match(PowerBuilderDWParserDQUOTED_STRING)
		}

	case PowerBuilderDWParserDATE:
		{
			p.SetState(130)
			p.Match(PowerBuilderDWParserDATE)
		}

	case PowerBuilderDWParserTIME:
		{
			p.SetState(131)
			p.Match(PowerBuilderDWParserTIME)
		}

	case PowerBuilderDWParserTYPE, PowerBuilderDWParserUPDATE, PowerBuilderDWParserID:
		{
			p.SetState(132)
			p.Attribute_name()
		}
		{
			p.SetState(133)

			var _m = p.Match(PowerBuilderDWParserEQ)

			localctx.(*Datawindow_property_attribute_subContext).eq = _m
		}
		p.SetState(146)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(134)
				p.Attribute_value()
			}
			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == PowerBuilderDWParserLBRACE || _la == PowerBuilderDWParserBRACES {
				{
					p.SetState(135)
					p.Array_decl_sub()
				}

			}

		case 2:
			{
				p.SetState(138)
				p.Match(PowerBuilderDWParserLPAREN)
			}
			p.SetState(140)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(PowerBuilderDWParserTYPE-45))|(1<<(PowerBuilderDWParserNULL_-45))|(1<<(PowerBuilderDWParserUPDATE-45)))) != 0) || (((_la-130)&-(0x1f+1)) == 0 && ((1<<uint((_la-130)))&((1<<(PowerBuilderDWParserDQUOTED_STRING-130))|(1<<(PowerBuilderDWParserNUMBER-130))|(1<<(PowerBuilderDWParserDATE-130))|(1<<(PowerBuilderDWParserTIME-130))|(1<<(PowerBuilderDWParserID-130)))) != 0) {
				{
					p.SetState(139)
					p.Datawindow_property_attribute_sub()
				}

				p.SetState(142)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(144)
				p.Match(PowerBuilderDWParserRPAREN)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerBuilderDWParserCOMMA {
		{
			p.SetState(150)
			p.Match(PowerBuilderDWParserCOMMA)
		}

	}

	return localctx
}

// IAttribute_nameContext is an interface to support dynamic dispatch.
type IAttribute_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttribute_nameContext differentiates from other interfaces.
	IsAttribute_nameContext()
}

type Attribute_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribute_nameContext() *Attribute_nameContext {
	var p = new(Attribute_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_attribute_name
	return p
}

func (*Attribute_nameContext) IsAttribute_nameContext() {}

func NewAttribute_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Attribute_nameContext {
	var p = new(Attribute_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_attribute_name

	return p
}

func (s *Attribute_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Attribute_nameContext) AllIdentifier_name() []IIdentifier_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifier_nameContext)(nil)).Elem())
	var tst = make([]IIdentifier_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifier_nameContext)
		}
	}

	return tst
}

func (s *Attribute_nameContext) Identifier_name(i int) IIdentifier_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifier_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifier_nameContext)
}

func (s *Attribute_nameContext) AllTYPE() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserTYPE)
}

func (s *Attribute_nameContext) TYPE(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserTYPE, i)
}

func (s *Attribute_nameContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserUPDATE, 0)
}

func (s *Attribute_nameContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserNUMBER, 0)
}

func (s *Attribute_nameContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserDOT)
}

func (s *Attribute_nameContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDOT, i)
}

func (s *Attribute_nameContext) AllCASE() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserCASE)
}

func (s *Attribute_nameContext) CASE(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCASE, i)
}

func (s *Attribute_nameContext) AllON() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserON)
}

func (s *Attribute_nameContext) ON(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserON, i)
}

func (s *Attribute_nameContext) AllDYNAMIC() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserDYNAMIC)
}

func (s *Attribute_nameContext) DYNAMIC(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDYNAMIC, i)
}

func (s *Attribute_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Attribute_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Attribute_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterAttribute_name(s)
	}
}

func (s *Attribute_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitAttribute_name(s)
	}
}

func (p *PowerBuilderDWParser) Attribute_name() (localctx IAttribute_nameContext) {
	this := p
	_ = this

	localctx = NewAttribute_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PowerBuilderDWParserRULE_attribute_name)
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
	p.SetState(156)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PowerBuilderDWParserID:
		{
			p.SetState(153)
			p.Identifier_name()
		}

	case PowerBuilderDWParserTYPE:
		{
			p.SetState(154)
			p.Match(PowerBuilderDWParserTYPE)
		}

	case PowerBuilderDWParserUPDATE:
		{
			p.SetState(155)
			p.Match(PowerBuilderDWParserUPDATE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerBuilderDWParserNUMBER {
		{
			p.SetState(158)
			p.Match(PowerBuilderDWParserNUMBER)
		}

	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PowerBuilderDWParserDOT {
		{
			p.SetState(161)
			p.Match(PowerBuilderDWParserDOT)
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case PowerBuilderDWParserID:
			{
				p.SetState(162)
				p.Identifier_name()
			}

		case PowerBuilderDWParserCASE:
			{
				p.SetState(163)
				p.Match(PowerBuilderDWParserCASE)
			}

		case PowerBuilderDWParserTYPE:
			{
				p.SetState(164)
				p.Match(PowerBuilderDWParserTYPE)
			}

		case PowerBuilderDWParserON:
			{
				p.SetState(165)
				p.Match(PowerBuilderDWParserON)
			}

		case PowerBuilderDWParserDYNAMIC:
			{
				p.SetState(166)
				p.Match(PowerBuilderDWParserDYNAMIC)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIdentifier_nameContext is an interface to support dynamic dispatch.
type IIdentifier_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifier_nameContext differentiates from other interfaces.
	IsIdentifier_nameContext()
}

type Identifier_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifier_nameContext() *Identifier_nameContext {
	var p = new(Identifier_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_identifier_name
	return p
}

func (*Identifier_nameContext) IsIdentifier_nameContext() {}

func NewIdentifier_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Identifier_nameContext {
	var p = new(Identifier_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_identifier_name

	return p
}

func (s *Identifier_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Identifier_nameContext) ID() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserID, 0)
}

func (s *Identifier_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Identifier_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Identifier_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterIdentifier_name(s)
	}
}

func (s *Identifier_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitIdentifier_name(s)
	}
}

func (p *PowerBuilderDWParser) Identifier_name() (localctx IIdentifier_nameContext) {
	this := p
	_ = this

	localctx = NewIdentifier_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PowerBuilderDWParserRULE_identifier_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(174)
		p.Match(PowerBuilderDWParserID)
	}

	return localctx
}

// IAttribute_valueContext is an interface to support dynamic dispatch.
type IAttribute_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAttribute_valueContext differentiates from other interfaces.
	IsAttribute_valueContext()
}

type Attribute_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttribute_valueContext() *Attribute_valueContext {
	var p = new(Attribute_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_attribute_value
	return p
}

func (*Attribute_valueContext) IsAttribute_valueContext() {}

func NewAttribute_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Attribute_valueContext {
	var p = new(Attribute_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_attribute_value

	return p
}

func (s *Attribute_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Attribute_valueContext) Atom_sub_call1() IAtom_sub_call1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtom_sub_call1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtom_sub_call1Context)
}

func (s *Attribute_valueContext) Atom_sub_member1() IAtom_sub_member1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtom_sub_member1Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtom_sub_member1Context)
}

func (s *Attribute_valueContext) Numeric_atom() INumeric_atomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumeric_atomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumeric_atomContext)
}

func (s *Attribute_valueContext) MINUS() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserMINUS, 0)
}

func (s *Attribute_valueContext) Boolean_atom() IBoolean_atomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolean_atomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolean_atomContext)
}

func (s *Attribute_valueContext) ENUM() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserENUM, 0)
}

func (s *Attribute_valueContext) DQUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDQUOTED_STRING, 0)
}

func (s *Attribute_valueContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserQUOTED_STRING, 0)
}

func (s *Attribute_valueContext) DATE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDATE, 0)
}

func (s *Attribute_valueContext) TIME() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserTIME, 0)
}

func (s *Attribute_valueContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserTYPE, 0)
}

func (s *Attribute_valueContext) TO() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserTO, 0)
}

func (s *Attribute_valueContext) FROM() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserFROM, 0)
}

func (s *Attribute_valueContext) REF() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserREF, 0)
}

func (s *Attribute_valueContext) NULL_() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserNULL_, 0)
}

func (s *Attribute_valueContext) OPEN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserOPEN, 0)
}

func (s *Attribute_valueContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserLPAREN)
}

func (s *Attribute_valueContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserLPAREN, i)
}

func (s *Attribute_valueContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserRPAREN)
}

func (s *Attribute_valueContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserRPAREN, i)
}

func (s *Attribute_valueContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Attribute_valueContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Attribute_valueContext) AllDataTypeSub() []IDataTypeSubContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDataTypeSubContext)(nil)).Elem())
	var tst = make([]IDataTypeSubContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDataTypeSubContext)
		}
	}

	return tst
}

func (s *Attribute_valueContext) DataTypeSub(i int) IDataTypeSubContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDataTypeSubContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDataTypeSubContext)
}

func (s *Attribute_valueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserCOMMA)
}

func (s *Attribute_valueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCOMMA, i)
}

func (s *Attribute_valueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserNUMBER, 0)
}

func (s *Attribute_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Attribute_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Attribute_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterAttribute_value(s)
	}
}

func (s *Attribute_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitAttribute_value(s)
	}
}

func (p *PowerBuilderDWParser) Attribute_value() (localctx IAttribute_valueContext) {
	this := p
	_ = this

	localctx = NewAttribute_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PowerBuilderDWParserRULE_attribute_value)
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

	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Atom_sub_call1()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Atom_sub_member1()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PowerBuilderDWParserMINUS {
			{
				p.SetState(178)
				p.Match(PowerBuilderDWParserMINUS)
			}

		}
		{
			p.SetState(181)
			p.Numeric_atom()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(182)
			p.Boolean_atom()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(183)
			p.Match(PowerBuilderDWParserENUM)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(184)
			p.Match(PowerBuilderDWParserDQUOTED_STRING)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(185)
			p.Match(PowerBuilderDWParserQUOTED_STRING)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(186)
			p.Match(PowerBuilderDWParserDATE)
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(187)
			p.Match(PowerBuilderDWParserTIME)
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(188)
			p.Match(PowerBuilderDWParserTYPE)
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(189)
			p.Match(PowerBuilderDWParserTO)
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(190)
			p.Match(PowerBuilderDWParserFROM)
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(191)
			p.Match(PowerBuilderDWParserREF)
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(192)
			p.Match(PowerBuilderDWParserNULL_)
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(193)
			p.Match(PowerBuilderDWParserOPEN)
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(194)
			p.Match(PowerBuilderDWParserLPAREN)
		}
		{
			p.SetState(195)
			p.Match(PowerBuilderDWParserLPAREN)
		}
		p.SetState(198)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case PowerBuilderDWParserCLOSE, PowerBuilderDWParserHALT, PowerBuilderDWParserLCURLY:
			{
				p.SetState(196)
				p.Expression()
			}

		case PowerBuilderDWParserANY, PowerBuilderDWParserBLOB, PowerBuilderDWParserBOOLEAN, PowerBuilderDWParserBYTE, PowerBuilderDWParserCHARACTER, PowerBuilderDWParserCHAR, PowerBuilderDWParserDATE_TYPE, PowerBuilderDWParserDATETIME, PowerBuilderDWParserDECIMAL, PowerBuilderDWParserDEC, PowerBuilderDWParserDOUBLE, PowerBuilderDWParserINTEGER, PowerBuilderDWParserINT, PowerBuilderDWParserLONG, PowerBuilderDWParserLONGLONG, PowerBuilderDWParserREAL, PowerBuilderDWParserSTRING, PowerBuilderDWParserTIME_TYPE, PowerBuilderDWParserUNSIGNEDINTEGER, PowerBuilderDWParserUINT, PowerBuilderDWParserUNSIGNEDLONG, PowerBuilderDWParserULONG, PowerBuilderDWParserWINDOW:
			{
				p.SetState(197)
				p.DataTypeSub()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PowerBuilderDWParserCOMMA {
			{
				p.SetState(200)
				p.Match(PowerBuilderDWParserCOMMA)
			}
			p.SetState(203)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case PowerBuilderDWParserCLOSE, PowerBuilderDWParserHALT, PowerBuilderDWParserLCURLY:
				{
					p.SetState(201)
					p.Expression()
				}

			case PowerBuilderDWParserANY, PowerBuilderDWParserBLOB, PowerBuilderDWParserBOOLEAN, PowerBuilderDWParserBYTE, PowerBuilderDWParserCHARACTER, PowerBuilderDWParserCHAR, PowerBuilderDWParserDATE_TYPE, PowerBuilderDWParserDATETIME, PowerBuilderDWParserDECIMAL, PowerBuilderDWParserDEC, PowerBuilderDWParserDOUBLE, PowerBuilderDWParserINTEGER, PowerBuilderDWParserINT, PowerBuilderDWParserLONG, PowerBuilderDWParserLONGLONG, PowerBuilderDWParserREAL, PowerBuilderDWParserSTRING, PowerBuilderDWParserTIME_TYPE, PowerBuilderDWParserUNSIGNEDINTEGER, PowerBuilderDWParserUINT, PowerBuilderDWParserUNSIGNEDLONG, PowerBuilderDWParserULONG, PowerBuilderDWParserWINDOW:
				{
					p.SetState(202)
					p.DataTypeSub()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		}
		{
			p.SetState(207)
			p.Match(PowerBuilderDWParserRPAREN)
		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == PowerBuilderDWParserCOMMA {
			{
				p.SetState(208)
				p.Match(PowerBuilderDWParserCOMMA)
			}
			{
				p.SetState(209)
				p.Match(PowerBuilderDWParserLPAREN)
			}
			p.SetState(212)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case PowerBuilderDWParserCLOSE, PowerBuilderDWParserHALT, PowerBuilderDWParserLCURLY:
				{
					p.SetState(210)
					p.Expression()
				}

			case PowerBuilderDWParserANY, PowerBuilderDWParserBLOB, PowerBuilderDWParserBOOLEAN, PowerBuilderDWParserBYTE, PowerBuilderDWParserCHARACTER, PowerBuilderDWParserCHAR, PowerBuilderDWParserDATE_TYPE, PowerBuilderDWParserDATETIME, PowerBuilderDWParserDECIMAL, PowerBuilderDWParserDEC, PowerBuilderDWParserDOUBLE, PowerBuilderDWParserINTEGER, PowerBuilderDWParserINT, PowerBuilderDWParserLONG, PowerBuilderDWParserLONGLONG, PowerBuilderDWParserREAL, PowerBuilderDWParserSTRING, PowerBuilderDWParserTIME_TYPE, PowerBuilderDWParserUNSIGNEDINTEGER, PowerBuilderDWParserUINT, PowerBuilderDWParserUNSIGNEDLONG, PowerBuilderDWParserULONG, PowerBuilderDWParserWINDOW:
				{
					p.SetState(211)
					p.DataTypeSub()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			p.SetState(219)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == PowerBuilderDWParserCOMMA {
				{
					p.SetState(214)
					p.Match(PowerBuilderDWParserCOMMA)
				}
				p.SetState(217)
				p.GetErrorHandler().Sync(p)

				switch p.GetTokenStream().LA(1) {
				case PowerBuilderDWParserCLOSE, PowerBuilderDWParserHALT, PowerBuilderDWParserLCURLY:
					{
						p.SetState(215)
						p.Expression()
					}

				case PowerBuilderDWParserANY, PowerBuilderDWParserBLOB, PowerBuilderDWParserBOOLEAN, PowerBuilderDWParserBYTE, PowerBuilderDWParserCHARACTER, PowerBuilderDWParserCHAR, PowerBuilderDWParserDATE_TYPE, PowerBuilderDWParserDATETIME, PowerBuilderDWParserDECIMAL, PowerBuilderDWParserDEC, PowerBuilderDWParserDOUBLE, PowerBuilderDWParserINTEGER, PowerBuilderDWParserINT, PowerBuilderDWParserLONG, PowerBuilderDWParserLONGLONG, PowerBuilderDWParserREAL, PowerBuilderDWParserSTRING, PowerBuilderDWParserTIME_TYPE, PowerBuilderDWParserUNSIGNEDINTEGER, PowerBuilderDWParserUINT, PowerBuilderDWParserUNSIGNEDLONG, PowerBuilderDWParserULONG, PowerBuilderDWParserWINDOW:
					{
						p.SetState(216)
						p.DataTypeSub()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

			}
			{
				p.SetState(221)
				p.Match(PowerBuilderDWParserRPAREN)
			}

			p.SetState(227)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(228)
			p.Match(PowerBuilderDWParserRPAREN)
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(230)
			p.DataTypeSub()
		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PowerBuilderDWParserLPAREN {
			{
				p.SetState(231)
				p.Match(PowerBuilderDWParserLPAREN)
			}
			{
				p.SetState(232)
				p.Match(PowerBuilderDWParserNUMBER)
			}
			{
				p.SetState(233)
				p.Match(PowerBuilderDWParserRPAREN)
			}

		}

	}

	return localctx
}

// INumeric_atomContext is an interface to support dynamic dispatch.
type INumeric_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumeric_atomContext differentiates from other interfaces.
	IsNumeric_atomContext()
}

type Numeric_atomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumeric_atomContext() *Numeric_atomContext {
	var p = new(Numeric_atomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_numeric_atom
	return p
}

func (*Numeric_atomContext) IsNumeric_atomContext() {}

func NewNumeric_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Numeric_atomContext {
	var p = new(Numeric_atomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_numeric_atom

	return p
}

func (s *Numeric_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Numeric_atomContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserNUMBER, 0)
}

func (s *Numeric_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Numeric_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Numeric_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterNumeric_atom(s)
	}
}

func (s *Numeric_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitNumeric_atom(s)
	}
}

func (p *PowerBuilderDWParser) Numeric_atom() (localctx INumeric_atomContext) {
	this := p
	_ = this

	localctx = NewNumeric_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PowerBuilderDWParserRULE_numeric_atom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(PowerBuilderDWParserNUMBER)
	}

	return localctx
}

// IDataTypeSubContext is an interface to support dynamic dispatch.
type IDataTypeSubContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataTypeSubContext differentiates from other interfaces.
	IsDataTypeSubContext()
}

type DataTypeSubContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataTypeSubContext() *DataTypeSubContext {
	var p = new(DataTypeSubContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_dataTypeSub
	return p
}

func (*DataTypeSubContext) IsDataTypeSubContext() {}

func NewDataTypeSubContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataTypeSubContext {
	var p = new(DataTypeSubContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_dataTypeSub

	return p
}

func (s *DataTypeSubContext) GetParser() antlr.Parser { return s.parser }

func (s *DataTypeSubContext) ANY() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserANY, 0)
}

func (s *DataTypeSubContext) BLOB() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserBLOB, 0)
}

func (s *DataTypeSubContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserBOOLEAN, 0)
}

func (s *DataTypeSubContext) BYTE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserBYTE, 0)
}

func (s *DataTypeSubContext) CHARACTER() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCHARACTER, 0)
}

func (s *DataTypeSubContext) CHAR() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCHAR, 0)
}

func (s *DataTypeSubContext) DATE_TYPE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDATE_TYPE, 0)
}

func (s *DataTypeSubContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDATETIME, 0)
}

func (s *DataTypeSubContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDECIMAL, 0)
}

func (s *DataTypeSubContext) DEC() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDEC, 0)
}

func (s *DataTypeSubContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDOUBLE, 0)
}

func (s *DataTypeSubContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserINTEGER, 0)
}

func (s *DataTypeSubContext) INT() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserINT, 0)
}

func (s *DataTypeSubContext) LONG() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserLONG, 0)
}

func (s *DataTypeSubContext) LONGLONG() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserLONGLONG, 0)
}

func (s *DataTypeSubContext) REAL() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserREAL, 0)
}

func (s *DataTypeSubContext) STRING() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserSTRING, 0)
}

func (s *DataTypeSubContext) TIME_TYPE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserTIME_TYPE, 0)
}

func (s *DataTypeSubContext) UNSIGNEDINTEGER() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserUNSIGNEDINTEGER, 0)
}

func (s *DataTypeSubContext) UINT() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserUINT, 0)
}

func (s *DataTypeSubContext) UNSIGNEDLONG() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserUNSIGNEDLONG, 0)
}

func (s *DataTypeSubContext) ULONG() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserULONG, 0)
}

func (s *DataTypeSubContext) WINDOW() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserWINDOW, 0)
}

func (s *DataTypeSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataTypeSubContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataTypeSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterDataTypeSub(s)
	}
}

func (s *DataTypeSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitDataTypeSub(s)
	}
}

func (p *PowerBuilderDWParser) DataTypeSub() (localctx IDataTypeSubContext) {
	this := p
	_ = this

	localctx = NewDataTypeSubContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PowerBuilderDWParserRULE_dataTypeSub)
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
		p.SetState(240)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<PowerBuilderDWParserANY)|(1<<PowerBuilderDWParserBLOB)|(1<<PowerBuilderDWParserBOOLEAN)|(1<<PowerBuilderDWParserBYTE)|(1<<PowerBuilderDWParserCHARACTER)|(1<<PowerBuilderDWParserCHAR)|(1<<PowerBuilderDWParserDATE_TYPE)|(1<<PowerBuilderDWParserDATETIME)|(1<<PowerBuilderDWParserDECIMAL)|(1<<PowerBuilderDWParserDEC)|(1<<PowerBuilderDWParserDOUBLE)|(1<<PowerBuilderDWParserINTEGER)|(1<<PowerBuilderDWParserINT)|(1<<PowerBuilderDWParserLONG)|(1<<PowerBuilderDWParserLONGLONG)|(1<<PowerBuilderDWParserREAL)|(1<<PowerBuilderDWParserSTRING)|(1<<PowerBuilderDWParserTIME_TYPE)|(1<<PowerBuilderDWParserUNSIGNEDINTEGER)|(1<<PowerBuilderDWParserUINT)|(1<<PowerBuilderDWParserUNSIGNEDLONG)|(1<<PowerBuilderDWParserULONG)|(1<<PowerBuilderDWParserWINDOW))) != 0) {
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
	p.RuleIndex = PowerBuilderDWParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Close_call_sub() IClose_call_subContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IClose_call_subContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IClose_call_subContext)
}

func (s *ExpressionContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserLCURLY, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PowerBuilderDWParser) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PowerBuilderDWParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(244)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PowerBuilderDWParserCLOSE, PowerBuilderDWParserHALT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.Close_call_sub()
		}

	case PowerBuilderDWParserLCURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.Match(PowerBuilderDWParserLCURLY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IArray_decl_subContext is an interface to support dynamic dispatch.
type IArray_decl_subContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArray_decl_subContext differentiates from other interfaces.
	IsArray_decl_subContext()
}

type Array_decl_subContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArray_decl_subContext() *Array_decl_subContext {
	var p = new(Array_decl_subContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_array_decl_sub
	return p
}

func (*Array_decl_subContext) IsArray_decl_subContext() {}

func NewArray_decl_subContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Array_decl_subContext {
	var p = new(Array_decl_subContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_array_decl_sub

	return p
}

func (s *Array_decl_subContext) GetParser() antlr.Parser { return s.parser }

func (s *Array_decl_subContext) BRACES() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserBRACES, 0)
}

func (s *Array_decl_subContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserLBRACE, 0)
}

func (s *Array_decl_subContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserRBRACE, 0)
}

func (s *Array_decl_subContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserNUMBER)
}

func (s *Array_decl_subContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserNUMBER, i)
}

func (s *Array_decl_subContext) AllTO() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserTO)
}

func (s *Array_decl_subContext) TO(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserTO, i)
}

func (s *Array_decl_subContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserCOMMA)
}

func (s *Array_decl_subContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCOMMA, i)
}

func (s *Array_decl_subContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserPLUS)
}

func (s *Array_decl_subContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserPLUS, i)
}

func (s *Array_decl_subContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserMINUS)
}

func (s *Array_decl_subContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserMINUS, i)
}

func (s *Array_decl_subContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Array_decl_subContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Array_decl_subContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterArray_decl_sub(s)
	}
}

func (s *Array_decl_subContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitArray_decl_sub(s)
	}
}

func (p *PowerBuilderDWParser) Array_decl_sub() (localctx IArray_decl_subContext) {
	this := p
	_ = this

	localctx = NewArray_decl_subContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PowerBuilderDWParserRULE_array_decl_sub)
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

	p.SetState(279)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PowerBuilderDWParserBRACES:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(246)
			p.Match(PowerBuilderDWParserBRACES)
		}

	case PowerBuilderDWParserLBRACE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(247)
			p.Match(PowerBuilderDWParserLBRACE)
		}
		p.SetState(276)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-114)&-(0x1f+1)) == 0 && ((1<<uint((_la-114)))&((1<<(PowerBuilderDWParserPLUS-114))|(1<<(PowerBuilderDWParserMINUS-114))|(1<<(PowerBuilderDWParserNUMBER-114)))) != 0 {
			p.SetState(249)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == PowerBuilderDWParserPLUS || _la == PowerBuilderDWParserMINUS {
				{
					p.SetState(248)
					_la = p.GetTokenStream().LA(1)

					if !(_la == PowerBuilderDWParserPLUS || _la == PowerBuilderDWParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			}
			{
				p.SetState(251)
				p.Match(PowerBuilderDWParserNUMBER)
			}
			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == PowerBuilderDWParserTO {
				{
					p.SetState(252)
					p.Match(PowerBuilderDWParserTO)
				}
				p.SetState(254)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == PowerBuilderDWParserPLUS || _la == PowerBuilderDWParserMINUS {
					{
						p.SetState(253)
						_la = p.GetTokenStream().LA(1)

						if !(_la == PowerBuilderDWParserPLUS || _la == PowerBuilderDWParserMINUS) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				}
				{
					p.SetState(256)
					p.Match(PowerBuilderDWParserNUMBER)
				}

			}
			p.SetState(273)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == PowerBuilderDWParserCOMMA {
				{
					p.SetState(259)
					p.Match(PowerBuilderDWParserCOMMA)
				}
				p.SetState(261)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == PowerBuilderDWParserPLUS || _la == PowerBuilderDWParserMINUS {
					{
						p.SetState(260)
						_la = p.GetTokenStream().LA(1)

						if !(_la == PowerBuilderDWParserPLUS || _la == PowerBuilderDWParserMINUS) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				}
				{
					p.SetState(263)
					p.Match(PowerBuilderDWParserNUMBER)
				}
				p.SetState(269)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == PowerBuilderDWParserTO {
					{
						p.SetState(264)
						p.Match(PowerBuilderDWParserTO)
					}
					p.SetState(266)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					if _la == PowerBuilderDWParserPLUS || _la == PowerBuilderDWParserMINUS {
						{
							p.SetState(265)
							_la = p.GetTokenStream().LA(1)

							if !(_la == PowerBuilderDWParserPLUS || _la == PowerBuilderDWParserMINUS) {
								p.GetErrorHandler().RecoverInline(p)
							} else {
								p.GetErrorHandler().ReportMatch(p)
								p.Consume()
							}
						}

					}
					{
						p.SetState(268)
						p.Match(PowerBuilderDWParserNUMBER)
					}

				}

				p.SetState(275)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(278)
			p.Match(PowerBuilderDWParserRBRACE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IClose_call_subContext is an interface to support dynamic dispatch.
type IClose_call_subContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsClose_call_subContext differentiates from other interfaces.
	IsClose_call_subContext()
}

type Close_call_subContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClose_call_subContext() *Close_call_subContext {
	var p = new(Close_call_subContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_close_call_sub
	return p
}

func (*Close_call_subContext) IsClose_call_subContext() {}

func NewClose_call_subContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Close_call_subContext {
	var p = new(Close_call_subContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_close_call_sub

	return p
}

func (s *Close_call_subContext) GetParser() antlr.Parser { return s.parser }

func (s *Close_call_subContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCLOSE, 0)
}

func (s *Close_call_subContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserLPAREN, 0)
}

func (s *Close_call_subContext) Expression_list() IExpression_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *Close_call_subContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserRPAREN, 0)
}

func (s *Close_call_subContext) HALT() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserHALT, 0)
}

func (s *Close_call_subContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Close_call_subContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Close_call_subContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterClose_call_sub(s)
	}
}

func (s *Close_call_subContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitClose_call_sub(s)
	}
}

func (p *PowerBuilderDWParser) Close_call_sub() (localctx IClose_call_subContext) {
	this := p
	_ = this

	localctx = NewClose_call_subContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PowerBuilderDWParserRULE_close_call_sub)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(288)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PowerBuilderDWParserCLOSE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Match(PowerBuilderDWParserCLOSE)
		}
		{
			p.SetState(282)
			p.Match(PowerBuilderDWParserLPAREN)
		}
		{
			p.SetState(283)
			p.Expression_list()
		}
		{
			p.SetState(284)
			p.Match(PowerBuilderDWParserRPAREN)
		}

	case PowerBuilderDWParserHALT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.Match(PowerBuilderDWParserHALT)
		}
		{
			p.SetState(287)
			p.Match(PowerBuilderDWParserCLOSE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpression_listContext is an interface to support dynamic dispatch.
type IExpression_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpression_listContext differentiates from other interfaces.
	IsExpression_listContext()
}

type Expression_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpression_listContext() *Expression_listContext {
	var p = new(Expression_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_expression_list
	return p
}

func (*Expression_listContext) IsExpression_listContext() {}

func NewExpression_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expression_listContext {
	var p = new(Expression_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_expression_list

	return p
}

func (s *Expression_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Expression_listContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Expression_listContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expression_listContext) AllREF() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserREF)
}

func (s *Expression_listContext) REF(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserREF, i)
}

func (s *Expression_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(PowerBuilderDWParserCOMMA)
}

func (s *Expression_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCOMMA, i)
}

func (s *Expression_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expression_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expression_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterExpression_list(s)
	}
}

func (s *Expression_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitExpression_list(s)
	}
}

func (p *PowerBuilderDWParser) Expression_list() (localctx IExpression_listContext) {
	this := p
	_ = this

	localctx = NewExpression_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PowerBuilderDWParserRULE_expression_list)
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
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerBuilderDWParserREF {
		{
			p.SetState(290)
			p.Match(PowerBuilderDWParserREF)
		}

	}
	{
		p.SetState(293)
		p.Expression()
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == PowerBuilderDWParserCOMMA {
		{
			p.SetState(294)
			p.Match(PowerBuilderDWParserCOMMA)
		}
		p.SetState(296)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == PowerBuilderDWParserREF {
			{
				p.SetState(295)
				p.Match(PowerBuilderDWParserREF)
			}

		}
		{
			p.SetState(298)
			p.Expression()
		}

		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IAtom_sub_call1Context is an interface to support dynamic dispatch.
type IAtom_sub_call1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtom_sub_call1Context differentiates from other interfaces.
	IsAtom_sub_call1Context()
}

type Atom_sub_call1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtom_sub_call1Context() *Atom_sub_call1Context {
	var p = new(Atom_sub_call1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_atom_sub_call1
	return p
}

func (*Atom_sub_call1Context) IsAtom_sub_call1Context() {}

func NewAtom_sub_call1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Atom_sub_call1Context {
	var p = new(Atom_sub_call1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_atom_sub_call1

	return p
}

func (s *Atom_sub_call1Context) GetParser() antlr.Parser { return s.parser }

func (s *Atom_sub_call1Context) LPAREN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserLPAREN, 0)
}

func (s *Atom_sub_call1Context) RPAREN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserRPAREN, 0)
}

func (s *Atom_sub_call1Context) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Atom_sub_call1Context) DESCRIBE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDESCRIBE, 0)
}

func (s *Atom_sub_call1Context) Expression_list() IExpression_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpression_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpression_listContext)
}

func (s *Atom_sub_call1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atom_sub_call1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Atom_sub_call1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterAtom_sub_call1(s)
	}
}

func (s *Atom_sub_call1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitAtom_sub_call1(s)
	}
}

func (p *PowerBuilderDWParser) Atom_sub_call1() (localctx IAtom_sub_call1Context) {
	this := p
	_ = this

	localctx = NewAtom_sub_call1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PowerBuilderDWParserRULE_atom_sub_call1)
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
	p.SetState(306)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case PowerBuilderDWParserSUPER, PowerBuilderDWParserID:
		{
			p.SetState(304)
			p.Identifier()
		}

	case PowerBuilderDWParserDESCRIBE:
		{
			p.SetState(305)
			p.Match(PowerBuilderDWParserDESCRIBE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(308)
		p.Match(PowerBuilderDWParserLPAREN)
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == PowerBuilderDWParserREF || _la == PowerBuilderDWParserCLOSE || _la == PowerBuilderDWParserHALT || _la == PowerBuilderDWParserLCURLY {
		{
			p.SetState(309)
			p.Expression_list()
		}

	}
	{
		p.SetState(312)
		p.Match(PowerBuilderDWParserRPAREN)
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
	p.RuleIndex = PowerBuilderDWParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) Identifier_name() IIdentifier_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifier_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifier_nameContext)
}

func (s *IdentifierContext) SUPER() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserSUPER, 0)
}

func (s *IdentifierContext) COLONCOLON() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCOLONCOLON, 0)
}

func (s *IdentifierContext) CREATE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCREATE, 0)
}

func (s *IdentifierContext) DESTROY() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDESTROY, 0)
}

func (s *IdentifierContext) Identifier_name_ex() IIdentifier_name_exContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifier_name_exContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifier_name_exContext)
}

func (s *IdentifierContext) DOT() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDOT, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *PowerBuilderDWParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PowerBuilderDWParserRULE_identifier)
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

	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(314)
			p.Identifier_name()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)
			p.Match(PowerBuilderDWParserSUPER)
		}
		{
			p.SetState(316)
			p.Match(PowerBuilderDWParserCOLONCOLON)
		}
		p.SetState(320)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case PowerBuilderDWParserCREATE:
			{
				p.SetState(317)
				p.Match(PowerBuilderDWParserCREATE)
			}

		case PowerBuilderDWParserDESTROY:
			{
				p.SetState(318)
				p.Match(PowerBuilderDWParserDESTROY)
			}

		case PowerBuilderDWParserREADONLY, PowerBuilderDWParserTYPE, PowerBuilderDWParserUPDATE, PowerBuilderDWParserOPEN, PowerBuilderDWParserGOTO, PowerBuilderDWParserCLOSE, PowerBuilderDWParserSELECT, PowerBuilderDWParserDELETE, PowerBuilderDWParserINSERT, PowerBuilderDWParserDESCRIBE, PowerBuilderDWParserTIME, PowerBuilderDWParserID:
			{
				p.SetState(319)
				p.Identifier_name_ex()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(322)
			p.Identifier_name()
		}
		{
			p.SetState(323)
			p.Match(PowerBuilderDWParserCOLONCOLON)
		}
		{
			p.SetState(324)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PowerBuilderDWParserCREATE || _la == PowerBuilderDWParserDESTROY) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(326)
			p.Identifier_name()
		}
		{
			p.SetState(327)
			p.Match(PowerBuilderDWParserDOT)
		}
		{
			p.SetState(328)
			_la = p.GetTokenStream().LA(1)

			if !(_la == PowerBuilderDWParserCREATE || _la == PowerBuilderDWParserDESTROY) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(330)
			p.Identifier_name()
		}
		{
			p.SetState(331)
			p.Match(PowerBuilderDWParserCOLONCOLON)
		}
		{
			p.SetState(332)
			p.Identifier_name_ex()
		}

	}

	return localctx
}

// IIdentifier_name_exContext is an interface to support dynamic dispatch.
type IIdentifier_name_exContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifier_name_exContext differentiates from other interfaces.
	IsIdentifier_name_exContext()
}

type Identifier_name_exContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifier_name_exContext() *Identifier_name_exContext {
	var p = new(Identifier_name_exContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_identifier_name_ex
	return p
}

func (*Identifier_name_exContext) IsIdentifier_name_exContext() {}

func NewIdentifier_name_exContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Identifier_name_exContext {
	var p = new(Identifier_name_exContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_identifier_name_ex

	return p
}

func (s *Identifier_name_exContext) GetParser() antlr.Parser { return s.parser }

func (s *Identifier_name_exContext) Identifier_name() IIdentifier_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifier_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifier_nameContext)
}

func (s *Identifier_name_exContext) SELECT() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserSELECT, 0)
}

func (s *Identifier_name_exContext) TYPE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserTYPE, 0)
}

func (s *Identifier_name_exContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserUPDATE, 0)
}

func (s *Identifier_name_exContext) DELETE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDELETE, 0)
}

func (s *Identifier_name_exContext) OPEN() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserOPEN, 0)
}

func (s *Identifier_name_exContext) CLOSE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserCLOSE, 0)
}

func (s *Identifier_name_exContext) GOTO() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserGOTO, 0)
}

func (s *Identifier_name_exContext) INSERT() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserINSERT, 0)
}

func (s *Identifier_name_exContext) DESCRIBE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserDESCRIBE, 0)
}

func (s *Identifier_name_exContext) TIME() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserTIME, 0)
}

func (s *Identifier_name_exContext) READONLY() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserREADONLY, 0)
}

func (s *Identifier_name_exContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Identifier_name_exContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Identifier_name_exContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterIdentifier_name_ex(s)
	}
}

func (s *Identifier_name_exContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitIdentifier_name_ex(s)
	}
}

func (p *PowerBuilderDWParser) Identifier_name_ex() (localctx IIdentifier_name_exContext) {
	this := p
	_ = this

	localctx = NewIdentifier_name_exContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PowerBuilderDWParserRULE_identifier_name_ex)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	case PowerBuilderDWParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(336)
			p.Identifier_name()
		}

	case PowerBuilderDWParserSELECT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(337)
			p.Match(PowerBuilderDWParserSELECT)
		}

	case PowerBuilderDWParserTYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(338)
			p.Match(PowerBuilderDWParserTYPE)
		}

	case PowerBuilderDWParserUPDATE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(339)
			p.Match(PowerBuilderDWParserUPDATE)
		}

	case PowerBuilderDWParserDELETE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(340)
			p.Match(PowerBuilderDWParserDELETE)
		}

	case PowerBuilderDWParserOPEN:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(341)
			p.Match(PowerBuilderDWParserOPEN)
		}

	case PowerBuilderDWParserCLOSE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(342)
			p.Match(PowerBuilderDWParserCLOSE)
		}

	case PowerBuilderDWParserGOTO:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(343)
			p.Match(PowerBuilderDWParserGOTO)
		}

	case PowerBuilderDWParserINSERT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(344)
			p.Match(PowerBuilderDWParserINSERT)
		}

	case PowerBuilderDWParserDESCRIBE:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(345)
			p.Match(PowerBuilderDWParserDESCRIBE)
		}

	case PowerBuilderDWParserTIME:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(346)
			p.Match(PowerBuilderDWParserTIME)
		}

	case PowerBuilderDWParserREADONLY:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(347)
			p.Match(PowerBuilderDWParserREADONLY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAtom_sub_member1Context is an interface to support dynamic dispatch.
type IAtom_sub_member1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtom_sub_member1Context differentiates from other interfaces.
	IsAtom_sub_member1Context()
}

type Atom_sub_member1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtom_sub_member1Context() *Atom_sub_member1Context {
	var p = new(Atom_sub_member1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_atom_sub_member1
	return p
}

func (*Atom_sub_member1Context) IsAtom_sub_member1Context() {}

func NewAtom_sub_member1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Atom_sub_member1Context {
	var p = new(Atom_sub_member1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_atom_sub_member1

	return p
}

func (s *Atom_sub_member1Context) GetParser() antlr.Parser { return s.parser }

func (s *Atom_sub_member1Context) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *Atom_sub_member1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atom_sub_member1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Atom_sub_member1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterAtom_sub_member1(s)
	}
}

func (s *Atom_sub_member1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitAtom_sub_member1(s)
	}
}

func (p *PowerBuilderDWParser) Atom_sub_member1() (localctx IAtom_sub_member1Context) {
	this := p
	_ = this

	localctx = NewAtom_sub_member1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PowerBuilderDWParserRULE_atom_sub_member1)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(350)
		p.Identifier()
	}

	return localctx
}

// IBoolean_atomContext is an interface to support dynamic dispatch.
type IBoolean_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolean_atomContext differentiates from other interfaces.
	IsBoolean_atomContext()
}

type Boolean_atomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_atomContext() *Boolean_atomContext {
	var p = new(Boolean_atomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = PowerBuilderDWParserRULE_boolean_atom
	return p
}

func (*Boolean_atomContext) IsBoolean_atomContext() {}

func NewBoolean_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_atomContext {
	var p = new(Boolean_atomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = PowerBuilderDWParserRULE_boolean_atom

	return p
}

func (s *Boolean_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_atomContext) TRUE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserTRUE, 0)
}

func (s *Boolean_atomContext) FALSE() antlr.TerminalNode {
	return s.GetToken(PowerBuilderDWParserFALSE, 0)
}

func (s *Boolean_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.EnterBoolean_atom(s)
	}
}

func (s *Boolean_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PowerBuilderDWParserListener); ok {
		listenerT.ExitBoolean_atom(s)
	}
}

func (p *PowerBuilderDWParser) Boolean_atom() (localctx IBoolean_atomContext) {
	this := p
	_ = this

	localctx = NewBoolean_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, PowerBuilderDWParserRULE_boolean_atom)
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
		_la = p.GetTokenStream().LA(1)

		if !(_la == PowerBuilderDWParserTRUE || _la == PowerBuilderDWParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
