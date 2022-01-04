// Code generated from alloy.g4 by ANTLR 4.9.3. DO NOT EDIT.

package alloy // alloy
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 75, 445,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 3, 2, 5, 2, 62, 10, 2, 3, 2, 7, 2, 65, 10, 2,
	12, 2, 14, 2, 68, 11, 2, 3, 2, 6, 2, 71, 10, 2, 13, 2, 14, 2, 72, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 80, 10, 3, 12, 3, 14, 3, 83, 11, 3, 3, 4,
	3, 4, 3, 4, 3, 4, 7, 4, 89, 10, 4, 12, 4, 14, 4, 92, 11, 4, 3, 4, 3, 4,
	5, 4, 96, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 104, 10, 5,
	3, 6, 5, 6, 107, 10, 6, 3, 6, 5, 6, 110, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	7, 6, 116, 10, 6, 12, 6, 14, 6, 119, 11, 6, 3, 6, 5, 6, 122, 10, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 7, 6, 128, 10, 6, 12, 6, 14, 6, 131, 11, 6, 5, 6,
	133, 10, 6, 3, 6, 3, 6, 5, 6, 137, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 7, 7, 145, 10, 7, 12, 7, 14, 7, 148, 11, 7, 5, 7, 150, 10, 7, 3,
	8, 3, 8, 3, 9, 5, 9, 155, 10, 9, 3, 9, 3, 9, 3, 9, 7, 9, 160, 10, 9, 12,
	9, 14, 9, 163, 11, 9, 3, 9, 3, 9, 5, 9, 167, 10, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 5, 10, 173, 10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 5,
	11, 181, 10, 11, 3, 11, 3, 11, 5, 11, 185, 10, 11, 3, 11, 3, 11, 3, 12,
	3, 12, 3, 12, 3, 12, 5, 12, 193, 10, 12, 3, 12, 3, 12, 5, 12, 197, 10,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	7, 13, 209, 10, 13, 12, 13, 14, 13, 212, 11, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 7, 13, 220, 10, 13, 12, 13, 14, 13, 223, 11, 13, 3, 13,
	3, 13, 5, 13, 227, 10, 13, 3, 14, 3, 14, 5, 14, 231, 10, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 5, 15, 238, 10, 15, 3, 15, 5, 15, 241, 10, 15,
	3, 15, 3, 15, 5, 15, 245, 10, 15, 3, 15, 5, 15, 248, 10, 15, 3, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 256, 10, 16, 12, 16, 14, 16, 259,
	11, 16, 5, 16, 261, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 267, 10,
	16, 12, 16, 14, 16, 270, 11, 16, 5, 16, 272, 10, 16, 3, 17, 5, 17, 275,
	10, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 293, 10, 18, 12,
	18, 14, 18, 296, 11, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18,
	304, 10, 18, 12, 18, 14, 18, 307, 11, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 7, 18, 315, 10, 18, 12, 18, 14, 18, 318, 11, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 328, 10, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 340,
	10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 347, 10, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 6, 18, 357, 10, 18, 13, 18,
	14, 18, 358, 3, 18, 3, 18, 7, 18, 363, 10, 18, 12, 18, 14, 18, 366, 11,
	18, 3, 19, 5, 19, 369, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 375,
	10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	5, 20, 386, 10, 20, 3, 21, 3, 21, 3, 22, 3, 22, 5, 22, 392, 10, 22, 3,
	22, 3, 22, 3, 22, 5, 22, 397, 10, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 25, 3, 25, 7, 25, 407, 10, 25, 12, 25, 14, 25, 410, 11, 25, 3,
	25, 3, 25, 3, 26, 3, 26, 3, 26, 5, 26, 417, 10, 26, 3, 27, 3, 27, 3, 27,
	3, 27, 5, 27, 423, 10, 27, 3, 28, 5, 28, 426, 10, 28, 3, 28, 3, 28, 3,
	28, 7, 28, 431, 10, 28, 12, 28, 14, 28, 434, 11, 28, 3, 28, 3, 28, 3, 29,
	3, 29, 3, 30, 6, 30, 441, 10, 30, 13, 30, 14, 30, 442, 3, 30, 2, 3, 34,
	31, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 2, 8, 3, 2, 14, 16, 3, 2, 28,
	29, 3, 2, 35, 36, 3, 2, 37, 38, 7, 2, 13, 13, 21, 21, 37, 38, 41, 41, 51,
	60, 4, 2, 12, 12, 62, 66, 2, 500, 2, 61, 3, 2, 2, 2, 4, 74, 3, 2, 2, 2,
	6, 84, 3, 2, 2, 2, 8, 103, 3, 2, 2, 2, 10, 106, 3, 2, 2, 2, 12, 149, 3,
	2, 2, 2, 14, 151, 3, 2, 2, 2, 16, 154, 3, 2, 2, 2, 18, 170, 3, 2, 2, 2,
	20, 176, 3, 2, 2, 2, 22, 188, 3, 2, 2, 2, 24, 226, 3, 2, 2, 2, 26, 228,
	3, 2, 2, 2, 28, 237, 3, 2, 2, 2, 30, 271, 3, 2, 2, 2, 32, 274, 3, 2, 2,
	2, 34, 327, 3, 2, 2, 2, 36, 374, 3, 2, 2, 2, 38, 385, 3, 2, 2, 2, 40, 387,
	3, 2, 2, 2, 42, 391, 3, 2, 2, 2, 44, 398, 3, 2, 2, 2, 46, 400, 3, 2, 2,
	2, 48, 404, 3, 2, 2, 2, 50, 416, 3, 2, 2, 2, 52, 422, 3, 2, 2, 2, 54, 425,
	3, 2, 2, 2, 56, 437, 3, 2, 2, 2, 58, 440, 3, 2, 2, 2, 60, 62, 5, 4, 3,
	2, 61, 60, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 66, 3, 2, 2, 2, 63, 65,
	5, 6, 4, 2, 64, 63, 3, 2, 2, 2, 65, 68, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2,
	66, 67, 3, 2, 2, 2, 67, 70, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 71, 5,
	8, 5, 2, 70, 69, 3, 2, 2, 2, 71, 72, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 72,
	73, 3, 2, 2, 2, 73, 3, 3, 2, 2, 2, 74, 75, 7, 3, 2, 2, 75, 76, 5, 54, 28,
	2, 76, 81, 5, 56, 29, 2, 77, 78, 7, 4, 2, 2, 78, 80, 5, 56, 29, 2, 79,
	77, 3, 2, 2, 2, 80, 83, 3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 81, 82, 3, 2, 2,
	2, 82, 5, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 84, 85, 7, 5, 2, 2, 85, 90, 5,
	54, 28, 2, 86, 87, 7, 4, 2, 2, 87, 89, 5, 54, 28, 2, 88, 86, 3, 2, 2, 2,
	89, 92, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 95, 3,
	2, 2, 2, 92, 90, 3, 2, 2, 2, 93, 94, 7, 6, 2, 2, 94, 96, 5, 56, 29, 2,
	95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 7, 3, 2, 2, 2, 97, 104, 5,
	10, 6, 2, 98, 104, 5, 18, 10, 2, 99, 104, 5, 20, 11, 2, 100, 104, 5, 22,
	12, 2, 101, 104, 5, 26, 14, 2, 102, 104, 5, 28, 15, 2, 103, 97, 3, 2, 2,
	2, 103, 98, 3, 2, 2, 2, 103, 99, 3, 2, 2, 2, 103, 100, 3, 2, 2, 2, 103,
	101, 3, 2, 2, 2, 103, 102, 3, 2, 2, 2, 104, 9, 3, 2, 2, 2, 105, 107, 7,
	7, 2, 2, 106, 105, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 109, 3, 2, 2,
	2, 108, 110, 5, 14, 8, 2, 109, 108, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110,
	111, 3, 2, 2, 2, 111, 112, 7, 8, 2, 2, 112, 117, 5, 56, 29, 2, 113, 114,
	7, 4, 2, 2, 114, 116, 5, 56, 29, 2, 115, 113, 3, 2, 2, 2, 116, 119, 3,
	2, 2, 2, 117, 115, 3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 121, 3, 2, 2,
	2, 119, 117, 3, 2, 2, 2, 120, 122, 5, 12, 7, 2, 121, 120, 3, 2, 2, 2, 121,
	122, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 132, 7, 9, 2, 2, 124, 129,
	5, 16, 9, 2, 125, 126, 7, 4, 2, 2, 126, 128, 5, 16, 9, 2, 127, 125, 3,
	2, 2, 2, 128, 131, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130, 3, 2, 2,
	2, 130, 133, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 132, 124, 3, 2, 2, 2, 132,
	133, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 136, 7, 10, 2, 2, 135, 137,
	5, 48, 25, 2, 136, 135, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 11, 3, 2,
	2, 2, 138, 139, 7, 11, 2, 2, 139, 150, 5, 54, 28, 2, 140, 141, 7, 12, 2,
	2, 141, 146, 5, 54, 28, 2, 142, 143, 7, 13, 2, 2, 143, 145, 5, 54, 28,
	2, 144, 142, 3, 2, 2, 2, 145, 148, 3, 2, 2, 2, 146, 144, 3, 2, 2, 2, 146,
	147, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2, 148, 146, 3, 2, 2, 2, 149, 138,
	3, 2, 2, 2, 149, 140, 3, 2, 2, 2, 150, 13, 3, 2, 2, 2, 151, 152, 9, 2,
	2, 2, 152, 15, 3, 2, 2, 2, 153, 155, 7, 17, 2, 2, 154, 153, 3, 2, 2, 2,
	154, 155, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 161, 5, 56, 29, 2, 157,
	158, 7, 4, 2, 2, 158, 160, 5, 56, 29, 2, 159, 157, 3, 2, 2, 2, 160, 163,
	3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162, 164, 3, 2,
	2, 2, 163, 161, 3, 2, 2, 2, 164, 166, 7, 18, 2, 2, 165, 167, 7, 17, 2,
	2, 166, 165, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168,
	169, 5, 34, 18, 2, 169, 17, 3, 2, 2, 2, 170, 172, 7, 19, 2, 2, 171, 173,
	5, 56, 29, 2, 172, 171, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 174, 3,
	2, 2, 2, 174, 175, 5, 48, 25, 2, 175, 19, 3, 2, 2, 2, 176, 180, 7, 20,
	2, 2, 177, 178, 5, 54, 28, 2, 178, 179, 7, 21, 2, 2, 179, 181, 3, 2, 2,
	2, 180, 177, 3, 2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182,
	184, 5, 56, 29, 2, 183, 185, 5, 24, 13, 2, 184, 183, 3, 2, 2, 2, 184, 185,
	3, 2, 2, 2, 185, 186, 3, 2, 2, 2, 186, 187, 5, 48, 25, 2, 187, 21, 3, 2,
	2, 2, 188, 192, 7, 22, 2, 2, 189, 190, 5, 54, 28, 2, 190, 191, 7, 21, 2,
	2, 191, 193, 3, 2, 2, 2, 192, 189, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193,
	194, 3, 2, 2, 2, 194, 196, 5, 56, 29, 2, 195, 197, 5, 24, 13, 2, 196, 195,
	3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 198, 3, 2, 2, 2, 198, 199, 7, 18,
	2, 2, 199, 200, 5, 34, 18, 2, 200, 201, 7, 9, 2, 2, 201, 202, 5, 34, 18,
	2, 202, 203, 7, 10, 2, 2, 203, 23, 3, 2, 2, 2, 204, 205, 7, 23, 2, 2, 205,
	210, 5, 16, 9, 2, 206, 207, 7, 4, 2, 2, 207, 209, 5, 16, 9, 2, 208, 206,
	3, 2, 2, 2, 209, 212, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210, 211, 3, 2,
	2, 2, 211, 213, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 213, 214, 7, 24, 2, 2,
	214, 227, 3, 2, 2, 2, 215, 216, 7, 25, 2, 2, 216, 221, 5, 16, 9, 2, 217,
	218, 7, 4, 2, 2, 218, 220, 5, 16, 9, 2, 219, 217, 3, 2, 2, 2, 220, 223,
	3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221, 222, 3, 2, 2, 2, 222, 224, 3, 2,
	2, 2, 223, 221, 3, 2, 2, 2, 224, 225, 7, 26, 2, 2, 225, 227, 3, 2, 2, 2,
	226, 204, 3, 2, 2, 2, 226, 215, 3, 2, 2, 2, 227, 25, 3, 2, 2, 2, 228, 230,
	7, 27, 2, 2, 229, 231, 5, 56, 29, 2, 230, 229, 3, 2, 2, 2, 230, 231, 3,
	2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 233, 5, 48, 25, 2, 233, 27, 3, 2, 2,
	2, 234, 235, 5, 56, 29, 2, 235, 236, 7, 18, 2, 2, 236, 238, 3, 2, 2, 2,
	237, 234, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 240, 3, 2, 2, 2, 239,
	241, 9, 3, 2, 2, 240, 239, 3, 2, 2, 2, 240, 241, 3, 2, 2, 2, 241, 244,
	3, 2, 2, 2, 242, 245, 5, 54, 28, 2, 243, 245, 5, 48, 25, 2, 244, 242, 3,
	2, 2, 2, 244, 243, 3, 2, 2, 2, 245, 247, 3, 2, 2, 2, 246, 248, 5, 30, 16,
	2, 247, 246, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 29, 3, 2, 2, 2, 249,
	250, 7, 30, 2, 2, 250, 260, 5, 58, 30, 2, 251, 252, 7, 31, 2, 2, 252, 257,
	5, 32, 17, 2, 253, 254, 7, 4, 2, 2, 254, 256, 5, 32, 17, 2, 255, 253, 3,
	2, 2, 2, 256, 259, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 257, 258, 3, 2, 2,
	2, 258, 261, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 260, 251, 3, 2, 2, 2, 260,
	261, 3, 2, 2, 2, 261, 272, 3, 2, 2, 2, 262, 263, 7, 30, 2, 2, 263, 268,
	5, 32, 17, 2, 264, 265, 7, 4, 2, 2, 265, 267, 5, 32, 17, 2, 266, 264, 3,
	2, 2, 2, 267, 270, 3, 2, 2, 2, 268, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2,
	2, 269, 272, 3, 2, 2, 2, 270, 268, 3, 2, 2, 2, 271, 249, 3, 2, 2, 2, 271,
	262, 3, 2, 2, 2, 272, 31, 3, 2, 2, 2, 273, 275, 7, 32, 2, 2, 274, 273,
	3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 277, 5, 58,
	30, 2, 277, 278, 5, 54, 28, 2, 278, 33, 3, 2, 2, 2, 279, 280, 8, 18, 1,
	2, 280, 328, 5, 36, 19, 2, 281, 328, 5, 54, 28, 2, 282, 283, 7, 33, 2,
	2, 283, 328, 5, 56, 29, 2, 284, 328, 7, 34, 2, 2, 285, 286, 5, 38, 20,
	2, 286, 287, 5, 34, 18, 13, 287, 328, 3, 2, 2, 2, 288, 289, 7, 40, 2, 2,
	289, 294, 5, 46, 24, 2, 290, 291, 7, 4, 2, 2, 291, 293, 5, 46, 24, 2, 292,
	290, 3, 2, 2, 2, 293, 296, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2, 294, 295,
	3, 2, 2, 2, 295, 297, 3, 2, 2, 2, 296, 294, 3, 2, 2, 2, 297, 298, 5, 50,
	26, 2, 298, 328, 3, 2, 2, 2, 299, 300, 5, 52, 27, 2, 300, 305, 5, 16, 9,
	2, 301, 302, 7, 4, 2, 2, 302, 304, 5, 16, 9, 2, 303, 301, 3, 2, 2, 2, 304,
	307, 3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 308,
	3, 2, 2, 2, 307, 305, 3, 2, 2, 2, 308, 309, 5, 50, 26, 2, 309, 328, 3,
	2, 2, 2, 310, 311, 7, 9, 2, 2, 311, 316, 5, 16, 9, 2, 312, 313, 7, 4, 2,
	2, 313, 315, 5, 16, 9, 2, 314, 312, 3, 2, 2, 2, 315, 318, 3, 2, 2, 2, 316,
	314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 319, 3, 2, 2, 2, 318, 316,
	3, 2, 2, 2, 319, 320, 5, 50, 26, 2, 320, 321, 7, 10, 2, 2, 321, 328, 3,
	2, 2, 2, 322, 323, 7, 23, 2, 2, 323, 324, 5, 34, 18, 2, 324, 325, 7, 24,
	2, 2, 325, 328, 3, 2, 2, 2, 326, 328, 5, 48, 25, 2, 327, 279, 3, 2, 2,
	2, 327, 281, 3, 2, 2, 2, 327, 282, 3, 2, 2, 2, 327, 284, 3, 2, 2, 2, 327,
	285, 3, 2, 2, 2, 327, 288, 3, 2, 2, 2, 327, 299, 3, 2, 2, 2, 327, 310,
	3, 2, 2, 2, 327, 322, 3, 2, 2, 2, 327, 326, 3, 2, 2, 2, 328, 364, 3, 2,
	2, 2, 329, 330, 12, 12, 2, 2, 330, 331, 5, 40, 21, 2, 331, 332, 5, 34,
	18, 13, 332, 363, 3, 2, 2, 2, 333, 334, 12, 11, 2, 2, 334, 335, 5, 42,
	22, 2, 335, 336, 5, 34, 18, 12, 336, 363, 3, 2, 2, 2, 337, 339, 12, 9,
	2, 2, 338, 340, 9, 4, 2, 2, 339, 338, 3, 2, 2, 2, 339, 340, 3, 2, 2, 2,
	340, 341, 3, 2, 2, 2, 341, 342, 5, 44, 23, 2, 342, 343, 5, 34, 18, 10,
	343, 363, 3, 2, 2, 2, 344, 346, 12, 8, 2, 2, 345, 347, 9, 5, 2, 2, 346,
	345, 3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348, 349,
	5, 34, 18, 2, 349, 350, 7, 39, 2, 2, 350, 351, 5, 34, 18, 9, 351, 363,
	3, 2, 2, 2, 352, 353, 12, 10, 2, 2, 353, 356, 7, 25, 2, 2, 354, 355, 7,
	4, 2, 2, 355, 357, 5, 34, 18, 2, 356, 354, 3, 2, 2, 2, 357, 358, 3, 2,
	2, 2, 358, 356, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 360, 3, 2, 2, 2,
	360, 361, 7, 26, 2, 2, 361, 363, 3, 2, 2, 2, 362, 329, 3, 2, 2, 2, 362,
	333, 3, 2, 2, 2, 362, 337, 3, 2, 2, 2, 362, 344, 3, 2, 2, 2, 362, 352,
	3, 2, 2, 2, 363, 366, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 364, 365, 3, 2,
	2, 2, 365, 35, 3, 2, 2, 2, 366, 364, 3, 2, 2, 2, 367, 369, 7, 41, 2, 2,
	368, 367, 3, 2, 2, 2, 368, 369, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370,
	375, 5, 58, 30, 2, 371, 375, 7, 42, 2, 2, 372, 375, 7, 43, 2, 2, 373, 375,
	7, 44, 2, 2, 374, 368, 3, 2, 2, 2, 374, 371, 3, 2, 2, 2, 374, 372, 3, 2,
	2, 2, 374, 373, 3, 2, 2, 2, 375, 37, 3, 2, 2, 2, 376, 386, 7, 35, 2, 2,
	377, 386, 7, 36, 2, 2, 378, 386, 7, 45, 2, 2, 379, 386, 5, 14, 8, 2, 380,
	386, 7, 46, 2, 2, 381, 386, 7, 47, 2, 2, 382, 386, 7, 48, 2, 2, 383, 386,
	7, 49, 2, 2, 384, 386, 7, 50, 2, 2, 385, 376, 3, 2, 2, 2, 385, 377, 3,
	2, 2, 2, 385, 378, 3, 2, 2, 2, 385, 379, 3, 2, 2, 2, 385, 380, 3, 2, 2,
	2, 385, 381, 3, 2, 2, 2, 385, 382, 3, 2, 2, 2, 385, 383, 3, 2, 2, 2, 385,
	384, 3, 2, 2, 2, 386, 39, 3, 2, 2, 2, 387, 388, 9, 6, 2, 2, 388, 41, 3,
	2, 2, 2, 389, 392, 5, 14, 8, 2, 390, 392, 7, 46, 2, 2, 391, 389, 3, 2,
	2, 2, 391, 390, 3, 2, 2, 2, 391, 392, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2,
	393, 396, 7, 61, 2, 2, 394, 397, 5, 14, 8, 2, 395, 397, 7, 46, 2, 2, 396,
	394, 3, 2, 2, 2, 396, 395, 3, 2, 2, 2, 396, 397, 3, 2, 2, 2, 397, 43, 3,
	2, 2, 2, 398, 399, 9, 7, 2, 2, 399, 45, 3, 2, 2, 2, 400, 401, 5, 56, 29,
	2, 401, 402, 7, 62, 2, 2, 402, 403, 5, 34, 18, 2, 403, 47, 3, 2, 2, 2,
	404, 408, 7, 9, 2, 2, 405, 407, 5, 34, 18, 2, 406, 405, 3, 2, 2, 2, 407,
	410, 3, 2, 2, 2, 408, 406, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 411,
	3, 2, 2, 2, 410, 408, 3, 2, 2, 2, 411, 412, 7, 10, 2, 2, 412, 49, 3, 2,
	2, 2, 413, 417, 5, 48, 25, 2, 414, 415, 7, 71, 2, 2, 415, 417, 5, 34, 18,
	2, 416, 413, 3, 2, 2, 2, 416, 414, 3, 2, 2, 2, 417, 51, 3, 2, 2, 2, 418,
	423, 7, 67, 2, 2, 419, 423, 7, 45, 2, 2, 420, 423, 7, 68, 2, 2, 421, 423,
	5, 14, 8, 2, 422, 418, 3, 2, 2, 2, 422, 419, 3, 2, 2, 2, 422, 420, 3, 2,
	2, 2, 422, 421, 3, 2, 2, 2, 423, 53, 3, 2, 2, 2, 424, 426, 7, 69, 2, 2,
	425, 424, 3, 2, 2, 2, 425, 426, 3, 2, 2, 2, 426, 432, 3, 2, 2, 2, 427,
	428, 5, 56, 29, 2, 428, 429, 7, 70, 2, 2, 429, 431, 3, 2, 2, 2, 430, 427,
	3, 2, 2, 2, 431, 434, 3, 2, 2, 2, 432, 430, 3, 2, 2, 2, 432, 433, 3, 2,
	2, 2, 433, 435, 3, 2, 2, 2, 434, 432, 3, 2, 2, 2, 435, 436, 5, 56, 29,
	2, 436, 55, 3, 2, 2, 2, 437, 438, 7, 73, 2, 2, 438, 57, 3, 2, 2, 2, 439,
	441, 7, 72, 2, 2, 440, 439, 3, 2, 2, 2, 441, 442, 3, 2, 2, 2, 442, 440,
	3, 2, 2, 2, 442, 443, 3, 2, 2, 2, 443, 59, 3, 2, 2, 2, 59, 61, 66, 72,
	81, 90, 95, 103, 106, 109, 117, 121, 129, 132, 136, 146, 149, 154, 161,
	166, 172, 180, 184, 192, 196, 210, 221, 226, 230, 237, 240, 244, 247, 257,
	260, 268, 271, 274, 294, 305, 316, 327, 339, 346, 358, 362, 364, 368, 374,
	385, 391, 396, 408, 416, 422, 425, 432, 442,
}
var literalNames = []string{
	"", "'module'", "','", "'open'", "'as'", "'abstract'", "'sig'", "'{'",
	"'}'", "'extends'", "'in'", "'+'", "'lone'", "'some'", "'one'", "'disj'",
	"':'", "'fact'", "'pred'", "'.'", "'fun'", "'('", "')'", "'['", "']'",
	"'assert'", "'run'", "'check'", "'for'", "'but'", "'exactly'", "'@'", "'this'",
	"'!'", "'not'", "'=>'", "'implies'", "'else'", "'let'", "'-'", "'none'",
	"'univ'", "'iden'", "'no'", "'set'", "'#'", "'~'", "'*'", "'^'", "'||'",
	"'or'", "'&&'", "'and'", "'<=>'", "'iff'", "'&'", "'++'", "'<:'", "':>'",
	"'->'", "'='", "'<'", "'>'", "'=<'", "'>='", "'all'", "'sum'", "'this/'",
	"'/'", "'|'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "BAR", "DIGIT",
	"IDENTIFIER", "COMMENT", "WS",
}

var ruleNames = []string{
	"alloyModule", "moduleDecl", "import_", "paragraph", "sigDecl", "sigExt",
	"mult", "decl", "factDecl", "predDecl", "funDecl", "paraDecls", "assertDecl",
	"cmdDecl", "scope", "typescope", "expr", "const_", "unOp", "binOp", "arrowOp",
	"compareOp", "letDecl", "block", "blockOrBar", "quant", "qualName", "name",
	"number",
}

type alloyParser struct {
	*antlr.BaseParser
}

// NewalloyParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *alloyParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewalloyParser(input antlr.TokenStream) *alloyParser {
	this := new(alloyParser)
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
	this.GrammarFileName = "alloy.g4"

	return this
}

// alloyParser tokens.
const (
	alloyParserEOF        = antlr.TokenEOF
	alloyParserT__0       = 1
	alloyParserT__1       = 2
	alloyParserT__2       = 3
	alloyParserT__3       = 4
	alloyParserT__4       = 5
	alloyParserT__5       = 6
	alloyParserT__6       = 7
	alloyParserT__7       = 8
	alloyParserT__8       = 9
	alloyParserT__9       = 10
	alloyParserT__10      = 11
	alloyParserT__11      = 12
	alloyParserT__12      = 13
	alloyParserT__13      = 14
	alloyParserT__14      = 15
	alloyParserT__15      = 16
	alloyParserT__16      = 17
	alloyParserT__17      = 18
	alloyParserT__18      = 19
	alloyParserT__19      = 20
	alloyParserT__20      = 21
	alloyParserT__21      = 22
	alloyParserT__22      = 23
	alloyParserT__23      = 24
	alloyParserT__24      = 25
	alloyParserT__25      = 26
	alloyParserT__26      = 27
	alloyParserT__27      = 28
	alloyParserT__28      = 29
	alloyParserT__29      = 30
	alloyParserT__30      = 31
	alloyParserT__31      = 32
	alloyParserT__32      = 33
	alloyParserT__33      = 34
	alloyParserT__34      = 35
	alloyParserT__35      = 36
	alloyParserT__36      = 37
	alloyParserT__37      = 38
	alloyParserT__38      = 39
	alloyParserT__39      = 40
	alloyParserT__40      = 41
	alloyParserT__41      = 42
	alloyParserT__42      = 43
	alloyParserT__43      = 44
	alloyParserT__44      = 45
	alloyParserT__45      = 46
	alloyParserT__46      = 47
	alloyParserT__47      = 48
	alloyParserT__48      = 49
	alloyParserT__49      = 50
	alloyParserT__50      = 51
	alloyParserT__51      = 52
	alloyParserT__52      = 53
	alloyParserT__53      = 54
	alloyParserT__54      = 55
	alloyParserT__55      = 56
	alloyParserT__56      = 57
	alloyParserT__57      = 58
	alloyParserT__58      = 59
	alloyParserT__59      = 60
	alloyParserT__60      = 61
	alloyParserT__61      = 62
	alloyParserT__62      = 63
	alloyParserT__63      = 64
	alloyParserT__64      = 65
	alloyParserT__65      = 66
	alloyParserT__66      = 67
	alloyParserT__67      = 68
	alloyParserBAR        = 69
	alloyParserDIGIT      = 70
	alloyParserIDENTIFIER = 71
	alloyParserCOMMENT    = 72
	alloyParserWS         = 73
)

// alloyParser rules.
const (
	alloyParserRULE_alloyModule = 0
	alloyParserRULE_moduleDecl  = 1
	alloyParserRULE_import_     = 2
	alloyParserRULE_paragraph   = 3
	alloyParserRULE_sigDecl     = 4
	alloyParserRULE_sigExt      = 5
	alloyParserRULE_mult        = 6
	alloyParserRULE_decl        = 7
	alloyParserRULE_factDecl    = 8
	alloyParserRULE_predDecl    = 9
	alloyParserRULE_funDecl     = 10
	alloyParserRULE_paraDecls   = 11
	alloyParserRULE_assertDecl  = 12
	alloyParserRULE_cmdDecl     = 13
	alloyParserRULE_scope       = 14
	alloyParserRULE_typescope   = 15
	alloyParserRULE_expr        = 16
	alloyParserRULE_const_      = 17
	alloyParserRULE_unOp        = 18
	alloyParserRULE_binOp       = 19
	alloyParserRULE_arrowOp     = 20
	alloyParserRULE_compareOp   = 21
	alloyParserRULE_letDecl     = 22
	alloyParserRULE_block       = 23
	alloyParserRULE_blockOrBar  = 24
	alloyParserRULE_quant       = 25
	alloyParserRULE_qualName    = 26
	alloyParserRULE_name        = 27
	alloyParserRULE_number      = 28
)

// IAlloyModuleContext is an interface to support dynamic dispatch.
type IAlloyModuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlloyModuleContext differentiates from other interfaces.
	IsAlloyModuleContext()
}

type AlloyModuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlloyModuleContext() *AlloyModuleContext {
	var p = new(AlloyModuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_alloyModule
	return p
}

func (*AlloyModuleContext) IsAlloyModuleContext() {}

func NewAlloyModuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AlloyModuleContext {
	var p = new(AlloyModuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_alloyModule

	return p
}

func (s *AlloyModuleContext) GetParser() antlr.Parser { return s.parser }

func (s *AlloyModuleContext) ModuleDecl() IModuleDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleDeclContext)
}

func (s *AlloyModuleContext) AllImport_() []IImport_Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImport_Context)(nil)).Elem())
	var tst = make([]IImport_Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImport_Context)
		}
	}

	return tst
}

func (s *AlloyModuleContext) Import_(i int) IImport_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImport_Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImport_Context)
}

func (s *AlloyModuleContext) AllParagraph() []IParagraphContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParagraphContext)(nil)).Elem())
	var tst = make([]IParagraphContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParagraphContext)
		}
	}

	return tst
}

func (s *AlloyModuleContext) Paragraph(i int) IParagraphContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParagraphContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParagraphContext)
}

func (s *AlloyModuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AlloyModuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AlloyModuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterAlloyModule(s)
	}
}

func (s *AlloyModuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitAlloyModule(s)
	}
}

func (p *alloyParser) AlloyModule() (localctx IAlloyModuleContext) {
	this := p
	_ = this

	localctx = NewAlloyModuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, alloyParserRULE_alloyModule)
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
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__0 {
		{
			p.SetState(58)
			p.ModuleDecl()
		}

	}
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == alloyParserT__2 {
		{
			p.SetState(61)
			p.Import_()
		}

		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<alloyParserT__4)|(1<<alloyParserT__5)|(1<<alloyParserT__6)|(1<<alloyParserT__11)|(1<<alloyParserT__12)|(1<<alloyParserT__13)|(1<<alloyParserT__16)|(1<<alloyParserT__17)|(1<<alloyParserT__19)|(1<<alloyParserT__24)|(1<<alloyParserT__25)|(1<<alloyParserT__26))) != 0) || _la == alloyParserT__66 || _la == alloyParserIDENTIFIER {
		{
			p.SetState(67)
			p.Paragraph()
		}

		p.SetState(70)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IModuleDeclContext is an interface to support dynamic dispatch.
type IModuleDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsModuleDeclContext differentiates from other interfaces.
	IsModuleDeclContext()
}

type ModuleDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyModuleDeclContext() *ModuleDeclContext {
	var p = new(ModuleDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_moduleDecl
	return p
}

func (*ModuleDeclContext) IsModuleDeclContext() {}

func NewModuleDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDeclContext {
	var p = new(ModuleDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_moduleDecl

	return p
}

func (s *ModuleDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDeclContext) QualName() IQualNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualNameContext)
}

func (s *ModuleDeclContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *ModuleDeclContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ModuleDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterModuleDecl(s)
	}
}

func (s *ModuleDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitModuleDecl(s)
	}
}

func (p *alloyParser) ModuleDecl() (localctx IModuleDeclContext) {
	this := p
	_ = this

	localctx = NewModuleDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, alloyParserRULE_moduleDecl)
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
		p.SetState(72)
		p.Match(alloyParserT__0)
	}
	{
		p.SetState(73)
		p.QualName()
	}
	{
		p.SetState(74)
		p.Name()
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == alloyParserT__1 {
		{
			p.SetState(75)
			p.Match(alloyParserT__1)
		}
		{
			p.SetState(76)
			p.Name()
		}

		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IImport_Context is an interface to support dynamic dispatch.
type IImport_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImport_Context differentiates from other interfaces.
	IsImport_Context()
}

type Import_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImport_Context() *Import_Context {
	var p = new(Import_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_import_
	return p
}

func (*Import_Context) IsImport_Context() {}

func NewImport_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Import_Context {
	var p = new(Import_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_import_

	return p
}

func (s *Import_Context) GetParser() antlr.Parser { return s.parser }

func (s *Import_Context) AllQualName() []IQualNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQualNameContext)(nil)).Elem())
	var tst = make([]IQualNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQualNameContext)
		}
	}

	return tst
}

func (s *Import_Context) QualName(i int) IQualNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQualNameContext)
}

func (s *Import_Context) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Import_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Import_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Import_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterImport_(s)
	}
}

func (s *Import_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitImport_(s)
	}
}

func (p *alloyParser) Import_() (localctx IImport_Context) {
	this := p
	_ = this

	localctx = NewImport_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, alloyParserRULE_import_)
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
		p.Match(alloyParserT__2)
	}
	{
		p.SetState(83)
		p.QualName()
	}
	p.SetState(88)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == alloyParserT__1 {
		{
			p.SetState(84)
			p.Match(alloyParserT__1)
		}
		{
			p.SetState(85)
			p.QualName()
		}

		p.SetState(90)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__3 {
		{
			p.SetState(91)
			p.Match(alloyParserT__3)
		}
		{
			p.SetState(92)
			p.Name()
		}

	}

	return localctx
}

// IParagraphContext is an interface to support dynamic dispatch.
type IParagraphContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParagraphContext differentiates from other interfaces.
	IsParagraphContext()
}

type ParagraphContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParagraphContext() *ParagraphContext {
	var p = new(ParagraphContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_paragraph
	return p
}

func (*ParagraphContext) IsParagraphContext() {}

func NewParagraphContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParagraphContext {
	var p = new(ParagraphContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_paragraph

	return p
}

func (s *ParagraphContext) GetParser() antlr.Parser { return s.parser }

func (s *ParagraphContext) SigDecl() ISigDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISigDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISigDeclContext)
}

func (s *ParagraphContext) FactDecl() IFactDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFactDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFactDeclContext)
}

func (s *ParagraphContext) PredDecl() IPredDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredDeclContext)
}

func (s *ParagraphContext) FunDecl() IFunDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunDeclContext)
}

func (s *ParagraphContext) AssertDecl() IAssertDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssertDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssertDeclContext)
}

func (s *ParagraphContext) CmdDecl() ICmdDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICmdDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICmdDeclContext)
}

func (s *ParagraphContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParagraphContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParagraphContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterParagraph(s)
	}
}

func (s *ParagraphContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitParagraph(s)
	}
}

func (p *alloyParser) Paragraph() (localctx IParagraphContext) {
	this := p
	_ = this

	localctx = NewParagraphContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, alloyParserRULE_paragraph)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(101)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case alloyParserT__4, alloyParserT__5, alloyParserT__11, alloyParserT__12, alloyParserT__13:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.SigDecl()
		}

	case alloyParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.FactDecl()
		}

	case alloyParserT__17:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(97)
			p.PredDecl()
		}

	case alloyParserT__19:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)
			p.FunDecl()
		}

	case alloyParserT__24:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(99)
			p.AssertDecl()
		}

	case alloyParserT__6, alloyParserT__25, alloyParserT__26, alloyParserT__66, alloyParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(100)
			p.CmdDecl()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISigDeclContext is an interface to support dynamic dispatch.
type ISigDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSigDeclContext differentiates from other interfaces.
	IsSigDeclContext()
}

type SigDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySigDeclContext() *SigDeclContext {
	var p = new(SigDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_sigDecl
	return p
}

func (*SigDeclContext) IsSigDeclContext() {}

func NewSigDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SigDeclContext {
	var p = new(SigDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_sigDecl

	return p
}

func (s *SigDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *SigDeclContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *SigDeclContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *SigDeclContext) Mult() IMultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultContext)
}

func (s *SigDeclContext) SigExt() ISigExtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISigExtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISigExtContext)
}

func (s *SigDeclContext) AllDecl() []IDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclContext)(nil)).Elem())
	var tst = make([]IDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclContext)
		}
	}

	return tst
}

func (s *SigDeclContext) Decl(i int) IDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *SigDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *SigDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SigDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SigDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterSigDecl(s)
	}
}

func (s *SigDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitSigDecl(s)
	}
}

func (p *alloyParser) SigDecl() (localctx ISigDeclContext) {
	this := p
	_ = this

	localctx = NewSigDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, alloyParserRULE_sigDecl)
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
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__4 {
		{
			p.SetState(103)
			p.Match(alloyParserT__4)
		}

	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<alloyParserT__11)|(1<<alloyParserT__12)|(1<<alloyParserT__13))) != 0 {
		{
			p.SetState(106)
			p.Mult()
		}

	}
	{
		p.SetState(109)
		p.Match(alloyParserT__5)
	}
	{
		p.SetState(110)
		p.Name()
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == alloyParserT__1 {
		{
			p.SetState(111)
			p.Match(alloyParserT__1)
		}
		{
			p.SetState(112)
			p.Name()
		}

		p.SetState(117)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__8 || _la == alloyParserT__9 {
		{
			p.SetState(118)
			p.SigExt()
		}

	}
	{
		p.SetState(121)
		p.Match(alloyParserT__6)
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__14 || _la == alloyParserIDENTIFIER {
		{
			p.SetState(122)
			p.Decl()
		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == alloyParserT__1 {
			{
				p.SetState(123)
				p.Match(alloyParserT__1)
			}
			{
				p.SetState(124)
				p.Decl()
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(132)
		p.Match(alloyParserT__7)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(133)
			p.Block()
		}

	}

	return localctx
}

// ISigExtContext is an interface to support dynamic dispatch.
type ISigExtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSigExtContext differentiates from other interfaces.
	IsSigExtContext()
}

type SigExtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySigExtContext() *SigExtContext {
	var p = new(SigExtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_sigExt
	return p
}

func (*SigExtContext) IsSigExtContext() {}

func NewSigExtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SigExtContext {
	var p = new(SigExtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_sigExt

	return p
}

func (s *SigExtContext) GetParser() antlr.Parser { return s.parser }

func (s *SigExtContext) AllQualName() []IQualNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQualNameContext)(nil)).Elem())
	var tst = make([]IQualNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQualNameContext)
		}
	}

	return tst
}

func (s *SigExtContext) QualName(i int) IQualNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQualNameContext)
}

func (s *SigExtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SigExtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SigExtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterSigExt(s)
	}
}

func (s *SigExtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitSigExt(s)
	}
}

func (p *alloyParser) SigExt() (localctx ISigExtContext) {
	this := p
	_ = this

	localctx = NewSigExtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, alloyParserRULE_sigExt)
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

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case alloyParserT__8:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(136)
			p.Match(alloyParserT__8)
		}
		{
			p.SetState(137)
			p.QualName()
		}

	case alloyParserT__9:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.Match(alloyParserT__9)
		}
		{
			p.SetState(139)
			p.QualName()
		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == alloyParserT__10 {
			{
				p.SetState(140)
				p.Match(alloyParserT__10)
			}
			{
				p.SetState(141)
				p.QualName()
			}

			p.SetState(146)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMultContext is an interface to support dynamic dispatch.
type IMultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultContext differentiates from other interfaces.
	IsMultContext()
}

type MultContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultContext() *MultContext {
	var p = new(MultContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_mult
	return p
}

func (*MultContext) IsMultContext() {}

func NewMultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultContext {
	var p = new(MultContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_mult

	return p
}

func (s *MultContext) GetParser() antlr.Parser { return s.parser }
func (s *MultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterMult(s)
	}
}

func (s *MultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitMult(s)
	}
}

func (p *alloyParser) Mult() (localctx IMultContext) {
	this := p
	_ = this

	localctx = NewMultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, alloyParserRULE_mult)
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
		p.SetState(149)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<alloyParserT__11)|(1<<alloyParserT__12)|(1<<alloyParserT__13))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_decl
	return p
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *DeclContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DeclContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterDecl(s)
	}
}

func (s *DeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitDecl(s)
	}
}

func (p *alloyParser) Decl() (localctx IDeclContext) {
	this := p
	_ = this

	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, alloyParserRULE_decl)
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
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__14 {
		{
			p.SetState(151)
			p.Match(alloyParserT__14)
		}

	}
	{
		p.SetState(154)
		p.Name()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == alloyParserT__1 {
		{
			p.SetState(155)
			p.Match(alloyParserT__1)
		}
		{
			p.SetState(156)
			p.Name()
		}

		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(162)
		p.Match(alloyParserT__15)
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__14 {
		{
			p.SetState(163)
			p.Match(alloyParserT__14)
		}

	}
	{
		p.SetState(166)
		p.expr(0)
	}

	return localctx
}

// IFactDeclContext is an interface to support dynamic dispatch.
type IFactDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFactDeclContext differentiates from other interfaces.
	IsFactDeclContext()
}

type FactDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFactDeclContext() *FactDeclContext {
	var p = new(FactDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_factDecl
	return p
}

func (*FactDeclContext) IsFactDeclContext() {}

func NewFactDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FactDeclContext {
	var p = new(FactDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_factDecl

	return p
}

func (s *FactDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FactDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FactDeclContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FactDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FactDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FactDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterFactDecl(s)
	}
}

func (s *FactDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitFactDecl(s)
	}
}

func (p *alloyParser) FactDecl() (localctx IFactDeclContext) {
	this := p
	_ = this

	localctx = NewFactDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, alloyParserRULE_factDecl)
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
		p.SetState(168)
		p.Match(alloyParserT__16)
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserIDENTIFIER {
		{
			p.SetState(169)
			p.Name()
		}

	}
	{
		p.SetState(172)
		p.Block()
	}

	return localctx
}

// IPredDeclContext is an interface to support dynamic dispatch.
type IPredDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredDeclContext differentiates from other interfaces.
	IsPredDeclContext()
}

type PredDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredDeclContext() *PredDeclContext {
	var p = new(PredDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_predDecl
	return p
}

func (*PredDeclContext) IsPredDeclContext() {}

func NewPredDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredDeclContext {
	var p = new(PredDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_predDecl

	return p
}

func (s *PredDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PredDeclContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *PredDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *PredDeclContext) QualName() IQualNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualNameContext)
}

func (s *PredDeclContext) ParaDecls() IParaDeclsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParaDeclsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParaDeclsContext)
}

func (s *PredDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterPredDecl(s)
	}
}

func (s *PredDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitPredDecl(s)
	}
}

func (p *alloyParser) PredDecl() (localctx IPredDeclContext) {
	this := p
	_ = this

	localctx = NewPredDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, alloyParserRULE_predDecl)
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
		p.SetState(174)
		p.Match(alloyParserT__17)
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(175)
			p.QualName()
		}
		{
			p.SetState(176)
			p.Match(alloyParserT__18)
		}

	}
	{
		p.SetState(180)
		p.Name()
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__20 || _la == alloyParserT__22 {
		{
			p.SetState(181)
			p.ParaDecls()
		}

	}
	{
		p.SetState(184)
		p.Block()
	}

	return localctx
}

// IFunDeclContext is an interface to support dynamic dispatch.
type IFunDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunDeclContext differentiates from other interfaces.
	IsFunDeclContext()
}

type FunDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunDeclContext() *FunDeclContext {
	var p = new(FunDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_funDecl
	return p
}

func (*FunDeclContext) IsFunDeclContext() {}

func NewFunDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunDeclContext {
	var p = new(FunDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_funDecl

	return p
}

func (s *FunDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *FunDeclContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FunDeclContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *FunDeclContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FunDeclContext) QualName() IQualNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualNameContext)
}

func (s *FunDeclContext) ParaDecls() IParaDeclsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParaDeclsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParaDeclsContext)
}

func (s *FunDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterFunDecl(s)
	}
}

func (s *FunDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitFunDecl(s)
	}
}

func (p *alloyParser) FunDecl() (localctx IFunDeclContext) {
	this := p
	_ = this

	localctx = NewFunDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, alloyParserRULE_funDecl)
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
		p.SetState(186)
		p.Match(alloyParserT__19)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(187)
			p.QualName()
		}
		{
			p.SetState(188)
			p.Match(alloyParserT__18)
		}

	}
	{
		p.SetState(192)
		p.Name()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__20 || _la == alloyParserT__22 {
		{
			p.SetState(193)
			p.ParaDecls()
		}

	}
	{
		p.SetState(196)
		p.Match(alloyParserT__15)
	}
	{
		p.SetState(197)
		p.expr(0)
	}
	{
		p.SetState(198)
		p.Match(alloyParserT__6)
	}
	{
		p.SetState(199)
		p.expr(0)
	}
	{
		p.SetState(200)
		p.Match(alloyParserT__7)
	}

	return localctx
}

// IParaDeclsContext is an interface to support dynamic dispatch.
type IParaDeclsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParaDeclsContext differentiates from other interfaces.
	IsParaDeclsContext()
}

type ParaDeclsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParaDeclsContext() *ParaDeclsContext {
	var p = new(ParaDeclsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_paraDecls
	return p
}

func (*ParaDeclsContext) IsParaDeclsContext() {}

func NewParaDeclsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParaDeclsContext {
	var p = new(ParaDeclsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_paraDecls

	return p
}

func (s *ParaDeclsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParaDeclsContext) AllDecl() []IDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclContext)(nil)).Elem())
	var tst = make([]IDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclContext)
		}
	}

	return tst
}

func (s *ParaDeclsContext) Decl(i int) IDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ParaDeclsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParaDeclsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParaDeclsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterParaDecls(s)
	}
}

func (s *ParaDeclsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitParaDecls(s)
	}
}

func (p *alloyParser) ParaDecls() (localctx IParaDeclsContext) {
	this := p
	_ = this

	localctx = NewParaDeclsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, alloyParserRULE_paraDecls)
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

	p.SetState(224)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case alloyParserT__20:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Match(alloyParserT__20)
		}
		{
			p.SetState(203)
			p.Decl()
		}
		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == alloyParserT__1 {
			{
				p.SetState(204)
				p.Match(alloyParserT__1)
			}
			{
				p.SetState(205)
				p.Decl()
			}

			p.SetState(210)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(211)
			p.Match(alloyParserT__21)
		}

	case alloyParserT__22:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(213)
			p.Match(alloyParserT__22)
		}
		{
			p.SetState(214)
			p.Decl()
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == alloyParserT__1 {
			{
				p.SetState(215)
				p.Match(alloyParserT__1)
			}
			{
				p.SetState(216)
				p.Decl()
			}

			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(222)
			p.Match(alloyParserT__23)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAssertDeclContext is an interface to support dynamic dispatch.
type IAssertDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssertDeclContext differentiates from other interfaces.
	IsAssertDeclContext()
}

type AssertDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssertDeclContext() *AssertDeclContext {
	var p = new(AssertDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_assertDecl
	return p
}

func (*AssertDeclContext) IsAssertDeclContext() {}

func NewAssertDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssertDeclContext {
	var p = new(AssertDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_assertDecl

	return p
}

func (s *AssertDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *AssertDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *AssertDeclContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AssertDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssertDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterAssertDecl(s)
	}
}

func (s *AssertDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitAssertDecl(s)
	}
}

func (p *alloyParser) AssertDecl() (localctx IAssertDeclContext) {
	this := p
	_ = this

	localctx = NewAssertDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, alloyParserRULE_assertDecl)
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
		p.SetState(226)
		p.Match(alloyParserT__24)
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserIDENTIFIER {
		{
			p.SetState(227)
			p.Name()
		}

	}
	{
		p.SetState(230)
		p.Block()
	}

	return localctx
}

// ICmdDeclContext is an interface to support dynamic dispatch.
type ICmdDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCmdDeclContext differentiates from other interfaces.
	IsCmdDeclContext()
}

type CmdDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCmdDeclContext() *CmdDeclContext {
	var p = new(CmdDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_cmdDecl
	return p
}

func (*CmdDeclContext) IsCmdDeclContext() {}

func NewCmdDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CmdDeclContext {
	var p = new(CmdDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_cmdDecl

	return p
}

func (s *CmdDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *CmdDeclContext) QualName() IQualNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualNameContext)
}

func (s *CmdDeclContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *CmdDeclContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *CmdDeclContext) Scope() IScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopeContext)
}

func (s *CmdDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CmdDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CmdDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterCmdDecl(s)
	}
}

func (s *CmdDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitCmdDecl(s)
	}
}

func (p *alloyParser) CmdDecl() (localctx ICmdDeclContext) {
	this := p
	_ = this

	localctx = NewCmdDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, alloyParserRULE_cmdDecl)
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
	p.SetState(235)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(232)
			p.Name()
		}
		{
			p.SetState(233)
			p.Match(alloyParserT__15)
		}

	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__25 || _la == alloyParserT__26 {
		{
			p.SetState(237)
			_la = p.GetTokenStream().LA(1)

			if !(_la == alloyParserT__25 || _la == alloyParserT__26) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(242)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case alloyParserT__66, alloyParserIDENTIFIER:
		{
			p.SetState(240)
			p.QualName()
		}

	case alloyParserT__6:
		{
			p.SetState(241)
			p.Block()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__27 {
		{
			p.SetState(244)
			p.Scope()
		}

	}

	return localctx
}

// IScopeContext is an interface to support dynamic dispatch.
type IScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScopeContext differentiates from other interfaces.
	IsScopeContext()
}

type ScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeContext() *ScopeContext {
	var p = new(ScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_scope
	return p
}

func (*ScopeContext) IsScopeContext() {}

func NewScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeContext {
	var p = new(ScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_scope

	return p
}

func (s *ScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *ScopeContext) AllTypescope() []ITypescopeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypescopeContext)(nil)).Elem())
	var tst = make([]ITypescopeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypescopeContext)
		}
	}

	return tst
}

func (s *ScopeContext) Typescope(i int) ITypescopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypescopeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypescopeContext)
}

func (s *ScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterScope(s)
	}
}

func (s *ScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitScope(s)
	}
}

func (p *alloyParser) Scope() (localctx IScopeContext) {
	this := p
	_ = this

	localctx = NewScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, alloyParserRULE_scope)
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

	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(247)
			p.Match(alloyParserT__27)
		}
		{
			p.SetState(248)
			p.Number()
		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == alloyParserT__28 {
			{
				p.SetState(249)
				p.Match(alloyParserT__28)
			}
			{
				p.SetState(250)
				p.Typescope()
			}
			p.SetState(255)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == alloyParserT__1 {
				{
					p.SetState(251)
					p.Match(alloyParserT__1)
				}
				{
					p.SetState(252)
					p.Typescope()
				}

				p.SetState(257)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(260)
			p.Match(alloyParserT__27)
		}
		{
			p.SetState(261)
			p.Typescope()
		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == alloyParserT__1 {
			{
				p.SetState(262)
				p.Match(alloyParserT__1)
			}
			{
				p.SetState(263)
				p.Typescope()
			}

			p.SetState(268)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// ITypescopeContext is an interface to support dynamic dispatch.
type ITypescopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypescopeContext differentiates from other interfaces.
	IsTypescopeContext()
}

type TypescopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypescopeContext() *TypescopeContext {
	var p = new(TypescopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_typescope
	return p
}

func (*TypescopeContext) IsTypescopeContext() {}

func NewTypescopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypescopeContext {
	var p = new(TypescopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_typescope

	return p
}

func (s *TypescopeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypescopeContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *TypescopeContext) QualName() IQualNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualNameContext)
}

func (s *TypescopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypescopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypescopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterTypescope(s)
	}
}

func (s *TypescopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitTypescope(s)
	}
}

func (p *alloyParser) Typescope() (localctx ITypescopeContext) {
	this := p
	_ = this

	localctx = NewTypescopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, alloyParserRULE_typescope)
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
	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__29 {
		{
			p.SetState(271)
			p.Match(alloyParserT__29)
		}

	}
	{
		p.SetState(274)
		p.Number()
	}
	{
		p.SetState(275)
		p.QualName()
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
	p.RuleIndex = alloyParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) Const_() IConst_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConst_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConst_Context)
}

func (s *ExprContext) QualName() IQualNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualNameContext)
}

func (s *ExprContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ExprContext) UnOp() IUnOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnOpContext)
}

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

func (s *ExprContext) AllLetDecl() []ILetDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILetDeclContext)(nil)).Elem())
	var tst = make([]ILetDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILetDeclContext)
		}
	}

	return tst
}

func (s *ExprContext) LetDecl(i int) ILetDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILetDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILetDeclContext)
}

func (s *ExprContext) BlockOrBar() IBlockOrBarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockOrBarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockOrBarContext)
}

func (s *ExprContext) Quant() IQuantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuantContext)
}

func (s *ExprContext) AllDecl() []IDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclContext)(nil)).Elem())
	var tst = make([]IDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclContext)
		}
	}

	return tst
}

func (s *ExprContext) Decl(i int) IDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *ExprContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ExprContext) BinOp() IBinOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinOpContext)
}

func (s *ExprContext) ArrowOp() IArrowOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrowOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrowOpContext)
}

func (s *ExprContext) CompareOp() ICompareOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICompareOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICompareOpContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *alloyParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *alloyParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 32
	p.EnterRecursionRule(localctx, 32, alloyParserRULE_expr, _p)
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
	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(278)
			p.Const_()
		}

	case 2:
		{
			p.SetState(279)
			p.QualName()
		}

	case 3:
		{
			p.SetState(280)
			p.Match(alloyParserT__30)
		}
		{
			p.SetState(281)
			p.Name()
		}

	case 4:
		{
			p.SetState(282)
			p.Match(alloyParserT__31)
		}

	case 5:
		{
			p.SetState(283)
			p.UnOp()
		}
		{
			p.SetState(284)
			p.expr(11)
		}

	case 6:
		{
			p.SetState(286)
			p.Match(alloyParserT__37)
		}
		{
			p.SetState(287)
			p.LetDecl()
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == alloyParserT__1 {
			{
				p.SetState(288)
				p.Match(alloyParserT__1)
			}
			{
				p.SetState(289)
				p.LetDecl()
			}

			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(295)
			p.BlockOrBar()
		}

	case 7:
		{
			p.SetState(297)
			p.Quant()
		}
		{
			p.SetState(298)
			p.Decl()
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == alloyParserT__1 {
			{
				p.SetState(299)
				p.Match(alloyParserT__1)
			}
			{
				p.SetState(300)
				p.Decl()
			}

			p.SetState(305)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(306)
			p.BlockOrBar()
		}

	case 8:
		{
			p.SetState(308)
			p.Match(alloyParserT__6)
		}
		{
			p.SetState(309)
			p.Decl()
		}
		p.SetState(314)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == alloyParserT__1 {
			{
				p.SetState(310)
				p.Match(alloyParserT__1)
			}
			{
				p.SetState(311)
				p.Decl()
			}

			p.SetState(316)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(317)
			p.BlockOrBar()
		}
		{
			p.SetState(318)
			p.Match(alloyParserT__7)
		}

	case 9:
		{
			p.SetState(320)
			p.Match(alloyParserT__20)
		}
		{
			p.SetState(321)
			p.expr(0)
		}
		{
			p.SetState(322)
			p.Match(alloyParserT__21)
		}

	case 10:
		{
			p.SetState(324)
			p.Block()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(360)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, alloyParserRULE_expr)
				p.SetState(327)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(328)
					p.BinOp()
				}
				{
					p.SetState(329)
					p.expr(11)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, alloyParserRULE_expr)
				p.SetState(331)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(332)
					p.ArrowOp()
				}
				{
					p.SetState(333)
					p.expr(10)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, alloyParserRULE_expr)
				p.SetState(335)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				p.SetState(337)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == alloyParserT__32 || _la == alloyParserT__33 {
					{
						p.SetState(336)
						_la = p.GetTokenStream().LA(1)

						if !(_la == alloyParserT__32 || _la == alloyParserT__33) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				}
				{
					p.SetState(339)
					p.CompareOp()
				}
				{
					p.SetState(340)
					p.expr(8)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, alloyParserRULE_expr)
				p.SetState(342)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				p.SetState(344)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == alloyParserT__34 || _la == alloyParserT__35 {
					{
						p.SetState(343)
						_la = p.GetTokenStream().LA(1)

						if !(_la == alloyParserT__34 || _la == alloyParserT__35) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				}
				{
					p.SetState(346)
					p.expr(0)
				}
				{
					p.SetState(347)
					p.Match(alloyParserT__36)
				}
				{
					p.SetState(348)
					p.expr(7)
				}

			case 5:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, alloyParserRULE_expr)
				p.SetState(350)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(351)
					p.Match(alloyParserT__22)
				}
				p.SetState(354)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				for ok := true; ok; ok = _la == alloyParserT__1 {
					{
						p.SetState(352)
						p.Match(alloyParserT__1)
					}
					{
						p.SetState(353)
						p.expr(0)
					}

					p.SetState(356)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)
				}
				{
					p.SetState(358)
					p.Match(alloyParserT__23)
				}

			}

		}
		p.SetState(364)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext())
	}

	return localctx
}

// IConst_Context is an interface to support dynamic dispatch.
type IConst_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConst_Context differentiates from other interfaces.
	IsConst_Context()
}

type Const_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConst_Context() *Const_Context {
	var p = new(Const_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_const_
	return p
}

func (*Const_Context) IsConst_Context() {}

func NewConst_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Const_Context {
	var p = new(Const_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_const_

	return p
}

func (s *Const_Context) GetParser() antlr.Parser { return s.parser }

func (s *Const_Context) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *Const_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Const_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Const_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterConst_(s)
	}
}

func (s *Const_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitConst_(s)
	}
}

func (p *alloyParser) Const_() (localctx IConst_Context) {
	this := p
	_ = this

	localctx = NewConst_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, alloyParserRULE_const_)
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

	p.SetState(372)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case alloyParserT__38, alloyParserDIGIT:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(366)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == alloyParserT__38 {
			{
				p.SetState(365)
				p.Match(alloyParserT__38)
			}

		}
		{
			p.SetState(368)
			p.Number()
		}

	case alloyParserT__39:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(369)
			p.Match(alloyParserT__39)
		}

	case alloyParserT__40:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(370)
			p.Match(alloyParserT__40)
		}

	case alloyParserT__41:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(371)
			p.Match(alloyParserT__41)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IUnOpContext is an interface to support dynamic dispatch.
type IUnOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnOpContext differentiates from other interfaces.
	IsUnOpContext()
}

type UnOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnOpContext() *UnOpContext {
	var p = new(UnOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_unOp
	return p
}

func (*UnOpContext) IsUnOpContext() {}

func NewUnOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnOpContext {
	var p = new(UnOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_unOp

	return p
}

func (s *UnOpContext) GetParser() antlr.Parser { return s.parser }

func (s *UnOpContext) Mult() IMultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultContext)
}

func (s *UnOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterUnOp(s)
	}
}

func (s *UnOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitUnOp(s)
	}
}

func (p *alloyParser) UnOp() (localctx IUnOpContext) {
	this := p
	_ = this

	localctx = NewUnOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, alloyParserRULE_unOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(383)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case alloyParserT__32:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(374)
			p.Match(alloyParserT__32)
		}

	case alloyParserT__33:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(375)
			p.Match(alloyParserT__33)
		}

	case alloyParserT__42:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(376)
			p.Match(alloyParserT__42)
		}

	case alloyParserT__11, alloyParserT__12, alloyParserT__13:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(377)
			p.Mult()
		}

	case alloyParserT__43:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(378)
			p.Match(alloyParserT__43)
		}

	case alloyParserT__44:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(379)
			p.Match(alloyParserT__44)
		}

	case alloyParserT__45:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(380)
			p.Match(alloyParserT__45)
		}

	case alloyParserT__46:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(381)
			p.Match(alloyParserT__46)
		}

	case alloyParserT__47:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(382)
			p.Match(alloyParserT__47)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBinOpContext is an interface to support dynamic dispatch.
type IBinOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinOpContext differentiates from other interfaces.
	IsBinOpContext()
}

type BinOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinOpContext() *BinOpContext {
	var p = new(BinOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_binOp
	return p
}

func (*BinOpContext) IsBinOpContext() {}

func NewBinOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinOpContext {
	var p = new(BinOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_binOp

	return p
}

func (s *BinOpContext) GetParser() antlr.Parser { return s.parser }
func (s *BinOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterBinOp(s)
	}
}

func (s *BinOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitBinOp(s)
	}
}

func (p *alloyParser) BinOp() (localctx IBinOpContext) {
	this := p
	_ = this

	localctx = NewBinOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, alloyParserRULE_binOp)
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
		p.SetState(385)
		_la = p.GetTokenStream().LA(1)

		if !(_la == alloyParserT__10 || _la == alloyParserT__18 || (((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(alloyParserT__34-35))|(1<<(alloyParserT__35-35))|(1<<(alloyParserT__38-35))|(1<<(alloyParserT__48-35))|(1<<(alloyParserT__49-35))|(1<<(alloyParserT__50-35))|(1<<(alloyParserT__51-35))|(1<<(alloyParserT__52-35))|(1<<(alloyParserT__53-35))|(1<<(alloyParserT__54-35))|(1<<(alloyParserT__55-35))|(1<<(alloyParserT__56-35))|(1<<(alloyParserT__57-35)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IArrowOpContext is an interface to support dynamic dispatch.
type IArrowOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrowOpContext differentiates from other interfaces.
	IsArrowOpContext()
}

type ArrowOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrowOpContext() *ArrowOpContext {
	var p = new(ArrowOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_arrowOp
	return p
}

func (*ArrowOpContext) IsArrowOpContext() {}

func NewArrowOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrowOpContext {
	var p = new(ArrowOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_arrowOp

	return p
}

func (s *ArrowOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrowOpContext) AllMult() []IMultContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMultContext)(nil)).Elem())
	var tst = make([]IMultContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMultContext)
		}
	}

	return tst
}

func (s *ArrowOpContext) Mult(i int) IMultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMultContext)
}

func (s *ArrowOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrowOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrowOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterArrowOp(s)
	}
}

func (s *ArrowOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitArrowOp(s)
	}
}

func (p *alloyParser) ArrowOp() (localctx IArrowOpContext) {
	this := p
	_ = this

	localctx = NewArrowOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, alloyParserRULE_arrowOp)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	p.SetState(389)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case alloyParserT__11, alloyParserT__12, alloyParserT__13:
		{
			p.SetState(387)
			p.Mult()
		}

	case alloyParserT__43:
		{
			p.SetState(388)
			p.Match(alloyParserT__43)
		}

	case alloyParserT__58:

	default:
	}
	{
		p.SetState(391)
		p.Match(alloyParserT__58)
	}
	p.SetState(394)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(392)
			p.Mult()
		}

	} else if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) == 2 {
		{
			p.SetState(393)
			p.Match(alloyParserT__43)
		}

	}

	return localctx
}

// ICompareOpContext is an interface to support dynamic dispatch.
type ICompareOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCompareOpContext differentiates from other interfaces.
	IsCompareOpContext()
}

type CompareOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCompareOpContext() *CompareOpContext {
	var p = new(CompareOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_compareOp
	return p
}

func (*CompareOpContext) IsCompareOpContext() {}

func NewCompareOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompareOpContext {
	var p = new(CompareOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_compareOp

	return p
}

func (s *CompareOpContext) GetParser() antlr.Parser { return s.parser }
func (s *CompareOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompareOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterCompareOp(s)
	}
}

func (s *CompareOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitCompareOp(s)
	}
}

func (p *alloyParser) CompareOp() (localctx ICompareOpContext) {
	this := p
	_ = this

	localctx = NewCompareOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, alloyParserRULE_compareOp)
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
		p.SetState(396)
		_la = p.GetTokenStream().LA(1)

		if !(_la == alloyParserT__9 || (((_la-60)&-(0x1f+1)) == 0 && ((1<<uint((_la-60)))&((1<<(alloyParserT__59-60))|(1<<(alloyParserT__60-60))|(1<<(alloyParserT__61-60))|(1<<(alloyParserT__62-60))|(1<<(alloyParserT__63-60)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILetDeclContext is an interface to support dynamic dispatch.
type ILetDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLetDeclContext differentiates from other interfaces.
	IsLetDeclContext()
}

type LetDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLetDeclContext() *LetDeclContext {
	var p = new(LetDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_letDecl
	return p
}

func (*LetDeclContext) IsLetDeclContext() {}

func NewLetDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LetDeclContext {
	var p = new(LetDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_letDecl

	return p
}

func (s *LetDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *LetDeclContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *LetDeclContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LetDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LetDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LetDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterLetDecl(s)
	}
}

func (s *LetDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitLetDecl(s)
	}
}

func (p *alloyParser) LetDecl() (localctx ILetDeclContext) {
	this := p
	_ = this

	localctx = NewLetDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, alloyParserRULE_letDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(398)
		p.Name()
	}
	{
		p.SetState(399)
		p.Match(alloyParserT__59)
	}
	{
		p.SetState(400)
		p.expr(0)
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
	p.RuleIndex = alloyParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *BlockContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *alloyParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, alloyParserRULE_block)
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
		p.SetState(402)
		p.Match(alloyParserT__6)
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<alloyParserT__6)|(1<<alloyParserT__11)|(1<<alloyParserT__12)|(1<<alloyParserT__13)|(1<<alloyParserT__20)|(1<<alloyParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(alloyParserT__31-32))|(1<<(alloyParserT__32-32))|(1<<(alloyParserT__33-32))|(1<<(alloyParserT__37-32))|(1<<(alloyParserT__38-32))|(1<<(alloyParserT__39-32))|(1<<(alloyParserT__40-32))|(1<<(alloyParserT__41-32))|(1<<(alloyParserT__42-32))|(1<<(alloyParserT__43-32))|(1<<(alloyParserT__44-32))|(1<<(alloyParserT__45-32))|(1<<(alloyParserT__46-32))|(1<<(alloyParserT__47-32)))) != 0) || (((_la-65)&-(0x1f+1)) == 0 && ((1<<uint((_la-65)))&((1<<(alloyParserT__64-65))|(1<<(alloyParserT__65-65))|(1<<(alloyParserT__66-65))|(1<<(alloyParserDIGIT-65))|(1<<(alloyParserIDENTIFIER-65)))) != 0) {
		{
			p.SetState(403)
			p.expr(0)
		}

		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(409)
		p.Match(alloyParserT__7)
	}

	return localctx
}

// IBlockOrBarContext is an interface to support dynamic dispatch.
type IBlockOrBarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockOrBarContext differentiates from other interfaces.
	IsBlockOrBarContext()
}

type BlockOrBarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockOrBarContext() *BlockOrBarContext {
	var p = new(BlockOrBarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_blockOrBar
	return p
}

func (*BlockOrBarContext) IsBlockOrBarContext() {}

func NewBlockOrBarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockOrBarContext {
	var p = new(BlockOrBarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_blockOrBar

	return p
}

func (s *BlockOrBarContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockOrBarContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BlockOrBarContext) BAR() antlr.TerminalNode {
	return s.GetToken(alloyParserBAR, 0)
}

func (s *BlockOrBarContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BlockOrBarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockOrBarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockOrBarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterBlockOrBar(s)
	}
}

func (s *BlockOrBarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitBlockOrBar(s)
	}
}

func (p *alloyParser) BlockOrBar() (localctx IBlockOrBarContext) {
	this := p
	_ = this

	localctx = NewBlockOrBarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, alloyParserRULE_blockOrBar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(414)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case alloyParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(411)
			p.Block()
		}

	case alloyParserBAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(412)
			p.Match(alloyParserBAR)
		}
		{
			p.SetState(413)
			p.expr(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQuantContext is an interface to support dynamic dispatch.
type IQuantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQuantContext differentiates from other interfaces.
	IsQuantContext()
}

type QuantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuantContext() *QuantContext {
	var p = new(QuantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_quant
	return p
}

func (*QuantContext) IsQuantContext() {}

func NewQuantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuantContext {
	var p = new(QuantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_quant

	return p
}

func (s *QuantContext) GetParser() antlr.Parser { return s.parser }

func (s *QuantContext) Mult() IMultContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMultContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMultContext)
}

func (s *QuantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterQuant(s)
	}
}

func (s *QuantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitQuant(s)
	}
}

func (p *alloyParser) Quant() (localctx IQuantContext) {
	this := p
	_ = this

	localctx = NewQuantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, alloyParserRULE_quant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(420)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case alloyParserT__64:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(416)
			p.Match(alloyParserT__64)
		}

	case alloyParserT__42:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(417)
			p.Match(alloyParserT__42)
		}

	case alloyParserT__65:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(418)
			p.Match(alloyParserT__65)
		}

	case alloyParserT__11, alloyParserT__12, alloyParserT__13:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(419)
			p.Mult()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQualNameContext is an interface to support dynamic dispatch.
type IQualNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualNameContext differentiates from other interfaces.
	IsQualNameContext()
}

type QualNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualNameContext() *QualNameContext {
	var p = new(QualNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_qualName
	return p
}

func (*QualNameContext) IsQualNameContext() {}

func NewQualNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualNameContext {
	var p = new(QualNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_qualName

	return p
}

func (s *QualNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QualNameContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *QualNameContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *QualNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterQualName(s)
	}
}

func (s *QualNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitQualName(s)
	}
}

func (p *alloyParser) QualName() (localctx IQualNameContext) {
	this := p
	_ = this

	localctx = NewQualNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, alloyParserRULE_qualName)
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
	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == alloyParserT__66 {
		{
			p.SetState(422)
			p.Match(alloyParserT__66)
		}

	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(425)
				p.Name()
			}
			{
				p.SetState(426)
				p.Match(alloyParserT__67)
			}

		}
		p.SetState(432)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())
	}
	{
		p.SetState(433)
		p.Name()
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
	p.RuleIndex = alloyParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(alloyParserIDENTIFIER, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitName(s)
	}
}

func (p *alloyParser) Name() (localctx INameContext) {
	this := p
	_ = this

	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, alloyParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(435)
		p.Match(alloyParserIDENTIFIER)
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = alloyParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = alloyParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) AllDIGIT() []antlr.TerminalNode {
	return s.GetTokens(alloyParserDIGIT)
}

func (s *NumberContext) DIGIT(i int) antlr.TerminalNode {
	return s.GetToken(alloyParserDIGIT, i)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(alloyListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *alloyParser) Number() (localctx INumberContext) {
	this := p
	_ = this

	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, alloyParserRULE_number)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(437)
				p.Match(alloyParserDIGIT)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(440)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 56, p.GetParserRuleContext())
	}

	return localctx
}

func (p *alloyParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 16:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *alloyParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
