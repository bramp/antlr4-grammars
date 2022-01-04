// Code generated from parkingsign.g4 by ANTLR 4.9.3. DO NOT EDIT.

package file. // parkingsign
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 55, 297, 
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13, 
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9, 
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23, 
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4, 
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34, 
	9, 34, 4, 35, 9, 35, 3, 2, 7, 2, 72, 10, 2, 12, 2, 14, 2, 75, 11, 2, 3, 
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 85, 10, 3, 3, 4, 5, 
	4, 88, 10, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 5, 5, 95, 10, 5, 3, 5, 3, 5, 
	5, 5, 99, 10, 5, 3, 5, 5, 5, 102, 10, 5, 3, 5, 5, 5, 105, 10, 5, 3, 6, 
	5, 6, 108, 10, 6, 3, 6, 3, 6, 5, 6, 112, 10, 6, 3, 6, 6, 6, 115, 10, 6, 
	13, 6, 14, 6, 116, 3, 6, 5, 6, 120, 10, 6, 3, 6, 5, 6, 123, 10, 6, 3, 7, 
	5, 7, 126, 10, 7, 3, 7, 3, 7, 6, 7, 130, 10, 7, 13, 7, 14, 7, 131, 3, 7, 
	5, 7, 135, 10, 7, 3, 7, 5, 7, 138, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 
	5, 8, 145, 10, 8, 3, 8, 5, 8, 148, 10, 8, 3, 8, 5, 8, 151, 10, 8, 3, 9, 
	3, 9, 3, 9, 3, 9, 5, 9, 157, 10, 9, 3, 9, 3, 9, 3, 9, 5, 9, 162, 10, 9, 
	3, 9, 5, 9, 165, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 180, 10, 10, 3, 11, 3, 
	11, 5, 11, 184, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 
	15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 220, 10, 18, 3, 19, 3, 
	19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 241, 10, 22, 3, 
	23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 5, 26, 251, 10, 26, 
	3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 5, 29, 262, 
	10, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 5, 31, 269, 10, 31, 3, 31, 3, 
	31, 5, 31, 273, 10, 31, 3, 31, 3, 31, 5, 31, 277, 10, 31, 3, 32, 3, 32, 
	3, 32, 5, 32, 282, 10, 32, 3, 33, 3, 33, 3, 33, 5, 33, 287, 10, 33, 3, 
	34, 3, 34, 5, 34, 291, 10, 34, 3, 35, 3, 35, 5, 35, 295, 10, 35, 3, 35, 
	2, 2, 36, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 2, 
	9, 3, 2, 50, 51, 3, 2, 18, 19, 3, 2, 11, 13, 3, 2, 23, 24, 3, 2, 37, 38, 
	3, 2, 34, 35, 3, 2, 42, 48, 2, 310, 2, 73, 3, 2, 2, 2, 4, 84, 3, 2, 2, 
	2, 6, 87, 3, 2, 2, 2, 8, 94, 3, 2, 2, 2, 10, 107, 3, 2, 2, 2, 12, 125, 
	3, 2, 2, 2, 14, 139, 3, 2, 2, 2, 16, 156, 3, 2, 2, 2, 18, 166, 3, 2, 2, 
	2, 20, 183, 3, 2, 2, 2, 22, 191, 3, 2, 2, 2, 24, 194, 3, 2, 2, 2, 26, 197, 
	3, 2, 2, 2, 28, 200, 3, 2, 2, 2, 30, 204, 3, 2, 2, 2, 32, 208, 3, 2, 2, 
	2, 34, 219, 3, 2, 2, 2, 36, 221, 3, 2, 2, 2, 38, 223, 3, 2, 2, 2, 40, 227, 
	3, 2, 2, 2, 42, 240, 3, 2, 2, 2, 44, 242, 3, 2, 2, 2, 46, 244, 3, 2, 2, 
	2, 48, 246, 3, 2, 2, 2, 50, 248, 3, 2, 2, 2, 52, 254, 3, 2, 2, 2, 54, 256, 
	3, 2, 2, 2, 56, 261, 3, 2, 2, 2, 58, 263, 3, 2, 2, 2, 60, 276, 3, 2, 2, 
	2, 62, 281, 3, 2, 2, 2, 64, 286, 3, 2, 2, 2, 66, 290, 3, 2, 2, 2, 68, 294, 
	3, 2, 2, 2, 70, 72, 5, 4, 3, 2, 71, 70, 3, 2, 2, 2, 72, 75, 3, 2, 2, 2, 
	73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 3, 3, 2, 2, 2, 75, 73, 3, 2, 
	2, 2, 76, 85, 5, 6, 4, 2, 77, 85, 5, 8, 5, 2, 78, 85, 5, 10, 6, 2, 79, 
	85, 5, 12, 7, 2, 80, 85, 5, 16, 9, 2, 81, 85, 5, 18, 10, 2, 82, 85, 5, 
	14, 8, 2, 83, 85, 5, 20, 11, 2, 84, 76, 3, 2, 2, 2, 84, 77, 3, 2, 2, 2, 
	84, 78, 3, 2, 2, 2, 84, 79, 3, 2, 2, 2, 84, 80, 3, 2, 2, 2, 84, 81, 3, 
	2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 5, 3, 2, 2, 2, 86, 
	88, 5, 24, 13, 2, 87, 86, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 89, 3, 2, 
	2, 2, 89, 90, 5, 34, 18, 2, 90, 91, 5, 58, 30, 2, 91, 92, 5, 22, 12, 2, 
	92, 7, 3, 2, 2, 2, 93, 95, 5, 24, 13, 2, 94, 93, 3, 2, 2, 2, 94, 95, 3, 
	2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 99, 5, 56, 29, 2, 97, 99, 5, 34, 18, 2, 
	98, 96, 3, 2, 2, 2, 98, 97, 3, 2, 2, 2, 99, 101, 3, 2, 2, 2, 100, 102, 
	7, 17, 2, 2, 101, 100, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 104, 3, 2, 
	2, 2, 103, 105, 5, 42, 22, 2, 104, 103, 3, 2, 2, 2, 104, 105, 3, 2, 2, 
	2, 105, 9, 3, 2, 2, 2, 106, 108, 5, 50, 26, 2, 107, 106, 3, 2, 2, 2, 107, 
	108, 3, 2, 2, 2, 108, 111, 3, 2, 2, 2, 109, 112, 5, 26, 14, 2, 110, 112, 
	5, 28, 15, 2, 111, 109, 3, 2, 2, 2, 111, 110, 3, 2, 2, 2, 112, 114, 3, 
	2, 2, 2, 113, 115, 5, 34, 18, 2, 114, 113, 3, 2, 2, 2, 115, 116, 3, 2, 
	2, 2, 116, 114, 3, 2, 2, 2, 116, 117, 3, 2, 2, 2, 117, 119, 3, 2, 2, 2, 
	118, 120, 7, 17, 2, 2, 119, 118, 3, 2, 2, 2, 119, 120, 3, 2, 2, 2, 120, 
	122, 3, 2, 2, 2, 121, 123, 5, 42, 22, 2, 122, 121, 3, 2, 2, 2, 122, 123, 
	3, 2, 2, 2, 123, 11, 3, 2, 2, 2, 124, 126, 5, 50, 26, 2, 125, 124, 3, 2, 
	2, 2, 125, 126, 3, 2, 2, 2, 126, 127, 3, 2, 2, 2, 127, 129, 5, 30, 16, 
	2, 128, 130, 5, 34, 18, 2, 129, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 
	131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 134, 3, 2, 2, 2, 133, 
	135, 7, 17, 2, 2, 134, 133, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 137, 
	3, 2, 2, 2, 136, 138, 5, 42, 22, 2, 137, 136, 3, 2, 2, 2, 137, 138, 3, 
	2, 2, 2, 138, 13, 3, 2, 2, 2, 139, 140, 5, 50, 26, 2, 140, 141, 7, 39, 
	2, 2, 141, 144, 5, 24, 13, 2, 142, 145, 5, 56, 29, 2, 143, 145, 5, 34, 
	18, 2, 144, 142, 3, 2, 2, 2, 144, 143, 3, 2, 2, 2, 145, 147, 3, 2, 2, 2, 
	146, 148, 7, 17, 2, 2, 147, 146, 3, 2, 2, 2, 147, 148, 3, 2, 2, 2, 148, 
	150, 3, 2, 2, 2, 149, 151, 5, 42, 22, 2, 150, 149, 3, 2, 2, 2, 150, 151, 
	3, 2, 2, 2, 151, 15, 3, 2, 2, 2, 152, 153, 7, 54, 2, 2, 153, 157, 7, 36, 
	2, 2, 154, 155, 7, 54, 2, 2, 155, 157, 5, 52, 27, 2, 156, 152, 3, 2, 2, 
	2, 156, 154, 3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 159, 7, 10, 2, 2, 159, 
	161, 5, 34, 18, 2, 160, 162, 7, 17, 2, 2, 161, 160, 3, 2, 2, 2, 161, 162, 
	3, 2, 2, 2, 162, 164, 3, 2, 2, 2, 163, 165, 5, 42, 22, 2, 164, 163, 3, 
	2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 17, 3, 2, 2, 2, 166, 167, 7, 54, 2, 
	2, 167, 168, 7, 36, 2, 2, 168, 169, 7, 10, 2, 2, 169, 170, 5, 42, 22, 2, 
	170, 171, 5, 42, 22, 2, 171, 172, 5, 60, 31, 2, 172, 173, 5, 60, 31, 2, 
	173, 174, 7, 11, 2, 2, 174, 175, 7, 11, 2, 2, 175, 176, 5, 60, 31, 2, 176, 
	179, 5, 60, 31, 2, 177, 178, 7, 17, 2, 2, 178, 180, 7, 22, 2, 2, 179, 177, 
	3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 19, 3, 2, 2, 2, 181, 182, 7, 30, 
	2, 2, 182, 184, 7, 31, 2, 2, 183, 181, 3, 2, 2, 2, 183, 184, 3, 2, 2, 2, 
	184, 185, 3, 2, 2, 2, 185, 186, 7, 32, 2, 2, 186, 187, 7, 9, 2, 2, 187, 
	188, 7, 54, 2, 2, 188, 189, 7, 33, 2, 2, 189, 190, 5, 54, 28, 2, 190, 21, 
	3, 2, 2, 2, 191, 192, 7, 49, 2, 2, 192, 193, 9, 2, 2, 2, 193, 23, 3, 2, 
	2, 2, 194, 195, 7, 9, 2, 2, 195, 196, 7, 10, 2, 2, 196, 25, 3, 2, 2, 2, 
	197, 198, 7, 9, 2, 2, 198, 199, 7, 27, 2, 2, 199, 27, 3, 2, 2, 2, 200, 
	201, 7, 28, 2, 2, 201, 202, 7, 10, 2, 2, 202, 203, 7, 29, 2, 2, 203, 29, 
	3, 2, 2, 2, 204, 205, 7, 40, 2, 2, 205, 206, 7, 41, 2, 2, 206, 207, 7, 
	29, 2, 2, 207, 31, 3, 2, 2, 2, 208, 209, 7, 20, 2, 2, 209, 210, 7, 21, 
	2, 2, 210, 33, 3, 2, 2, 2, 211, 212, 5, 60, 31, 2, 212, 213, 5, 46, 24, 
	2, 213, 214, 5, 60, 31, 2, 214, 220, 3, 2, 2, 2, 215, 216, 7, 54, 2, 2, 
	216, 217, 5, 46, 24, 2, 217, 218, 5, 60, 31, 2, 218, 220, 3, 2, 2, 2, 219, 
	211, 3, 2, 2, 2, 219, 215, 3, 2, 2, 2, 220, 35, 3, 2, 2, 2, 221, 222, 9, 
	3, 2, 2, 222, 37, 3, 2, 2, 2, 223, 224, 5, 58, 30, 2, 224, 225, 5, 46, 
	24, 2, 225, 226, 5, 58, 30, 2, 226, 39, 3, 2, 2, 2, 227, 228, 5, 58, 30, 
	2, 228, 229, 5, 48, 25, 2, 229, 230, 5, 58, 30, 2, 230, 41, 3, 2, 2, 2, 
	231, 241, 5, 36, 19, 2, 232, 241, 5, 32, 17, 2, 233, 241, 7, 22, 2, 2, 
	234, 241, 5, 40, 21, 2, 235, 241, 5, 38, 20, 2, 236, 237, 5, 58, 30, 2, 
	237, 238, 7, 29, 2, 2, 238, 241, 3, 2, 2, 2, 239, 241, 5, 58, 30, 2, 240, 
	231, 3, 2, 2, 2, 240, 232, 3, 2, 2, 2, 240, 233, 3, 2, 2, 2, 240, 234, 
	3, 2, 2, 2, 240, 235, 3, 2, 2, 2, 240, 236, 3, 2, 2, 2, 240, 239, 3, 2, 
	2, 2, 241, 43, 3, 2, 2, 2, 242, 243, 5, 42, 22, 2, 243, 45, 3, 2, 2, 2, 
	244, 245, 9, 4, 2, 2, 245, 47, 3, 2, 2, 2, 246, 247, 9, 5, 2, 2, 247, 49, 
	3, 2, 2, 2, 248, 250, 7, 25, 2, 2, 249, 251, 7, 13, 2, 2, 250, 249, 3, 
	2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 253, 7, 26, 2, 
	2, 253, 51, 3, 2, 2, 2, 254, 255, 9, 6, 2, 2, 255, 53, 3, 2, 2, 2, 256, 
	257, 9, 7, 2, 2, 257, 55, 3, 2, 2, 2, 258, 262, 7, 14, 2, 2, 259, 260, 
	7, 15, 2, 2, 260, 262, 7, 16, 2, 2, 261, 258, 3, 2, 2, 2, 261, 259, 3, 
	2, 2, 2, 262, 57, 3, 2, 2, 2, 263, 264, 9, 8, 2, 2, 264, 59, 3, 2, 2, 2, 
	265, 268, 7, 54, 2, 2, 266, 267, 7, 3, 2, 2, 267, 269, 7, 54, 2, 2, 268, 
	266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 272, 3, 2, 2, 2, 270, 273, 
	5, 66, 34, 2, 271, 273, 5, 68, 35, 2, 272, 270, 3, 2, 2, 2, 272, 271, 3, 
	2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 277, 3, 2, 2, 2, 274, 277, 5, 62, 32, 
	2, 275, 277, 5, 64, 33, 2, 276, 265, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2, 
	276, 275, 3, 2, 2, 2, 277, 61, 3, 2, 2, 2, 278, 282, 7, 52, 2, 2, 279, 
	280, 7, 4, 2, 2, 280, 282, 7, 52, 2, 2, 281, 278, 3, 2, 2, 2, 281, 279, 
	3, 2, 2, 2, 282, 63, 3, 2, 2, 2, 283, 287, 7, 53, 2, 2, 284, 285, 7, 4, 
	2, 2, 285, 287, 7, 53, 2, 2, 286, 283, 3, 2, 2, 2, 286, 284, 3, 2, 2, 2, 
	287, 65, 3, 2, 2, 2, 288, 291, 7, 5, 2, 2, 289, 291, 7, 6, 2, 2, 290, 288, 
	3, 2, 2, 2, 290, 289, 3, 2, 2, 2, 291, 67, 3, 2, 2, 2, 292, 295, 7, 7, 
	2, 2, 293, 295, 7, 8, 2, 2, 294, 292, 3, 2, 2, 2, 294, 293, 3, 2, 2, 2, 
	295, 69, 3, 2, 2, 2, 37, 73, 84, 87, 94, 98, 101, 104, 107, 111, 116, 119, 
	122, 125, 131, 134, 137, 144, 147, 150, 156, 161, 164, 179, 183, 219, 240, 
	250, 261, 268, 272, 276, 281, 286, 290, 294,
}
var literalNames = []string{
	"", "':'", "'12'", "'AM'", "'A.M.'", "'PM'", "'P.M.'", "'NO'", "'PARKING'", 
	"'TO'", "'THRU'", "'-'", "'ANYTIME'", "'ANY'", "'TIME'", "'EXCEPT'", "'DAILY'", 
	"'NIGHTLY'", "'SCHOOL'", "'DAYS'", "'HOLIDAYS'", "'AND'", "'&'", "'TOW'", 
	"'AWAY'", "'STOPPING'", "'VALET'", "'ONLY'", "'VEHICLES'", "'WITH'", "'DISTRICT'", 
	"'PERMITS'", "'EXEMPTED'", "'EXEMPT'", "'HOUR'", "'MINUTE'", "'MIN'", "'TEMPORARY'", 
	"'PASSENGER'", "'LOADING'", "", "", "", "", "", "", "", "'STREET'", "'SWEEPING'", 
	"'CLEANING'", "'NOON'", "'MIDNIGHT'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "NO", "PARKING", "TO", "THRU", "DASH", "ANYTIME", 
	"ANY", "TIME", "EXCEPT", "DAILY", "NIGHTLY", "SCHOOL", "DAYS", "HOLIDAYS", 
	"AND", "AMPERSAND", "TOW", "AWAY", "STOPPING", "VALET", "ONLY", "VEHICLES", 
	"WITH", "DISTRICT", "PERMITS", "EXEMPTED", "EXEMPT", "HOUR", "MINUTE", 
	"MIN", "TEMPORARY", "PASSENGER", "LOADING", "MON", "TUE", "WED", "THU", 
	"FRI", "SAT", "SUN", "STREET", "SWEEPING", "CLEANING", "NOON", "MIDNIGHT", 
	"INT", "WS",
}

var ruleNames = []string{
	"parkingSigns", "parkingSign", "streetSweepingSign", "noParkingSign", "noStoppingSign", 
	"passengerLoadingSign", "temporaryNoParkingSign", "singleTimeLimitSign", 
	"doubleTimeLimitSign", "permitSign", "streetSweeping", "noParking", "noStopping", 
	"valetOnly", "loadingOnly", "schoolDays", "timeRange", "everyDay", "dayToDay", 
	"dayAndDay", "dayRange", "dayRangePlus", "to", "and_", "towAway", "minute", 
	"exempt", "anyTime", "day", "time", "twelveNoon", "twelveMidnight", "am", 
	"pm",
}
type parkingsignParser struct {
	*antlr.BaseParser
}

// NewparkingsignParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *parkingsignParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewparkingsignParser(input antlr.TokenStream) *parkingsignParser {
	this := new(parkingsignParser)
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
	this.GrammarFileName = "parkingsign.g4"

	return this
}


// parkingsignParser tokens.
const (
	parkingsignParserEOF = antlr.TokenEOF
	parkingsignParserT__0 = 1
	parkingsignParserT__1 = 2
	parkingsignParserT__2 = 3
	parkingsignParserT__3 = 4
	parkingsignParserT__4 = 5
	parkingsignParserT__5 = 6
	parkingsignParserNO = 7
	parkingsignParserPARKING = 8
	parkingsignParserTO = 9
	parkingsignParserTHRU = 10
	parkingsignParserDASH = 11
	parkingsignParserANYTIME = 12
	parkingsignParserANY = 13
	parkingsignParserTIME = 14
	parkingsignParserEXCEPT = 15
	parkingsignParserDAILY = 16
	parkingsignParserNIGHTLY = 17
	parkingsignParserSCHOOL = 18
	parkingsignParserDAYS = 19
	parkingsignParserHOLIDAYS = 20
	parkingsignParserAND = 21
	parkingsignParserAMPERSAND = 22
	parkingsignParserTOW = 23
	parkingsignParserAWAY = 24
	parkingsignParserSTOPPING = 25
	parkingsignParserVALET = 26
	parkingsignParserONLY = 27
	parkingsignParserVEHICLES = 28
	parkingsignParserWITH = 29
	parkingsignParserDISTRICT = 30
	parkingsignParserPERMITS = 31
	parkingsignParserEXEMPTED = 32
	parkingsignParserEXEMPT = 33
	parkingsignParserHOUR = 34
	parkingsignParserMINUTE = 35
	parkingsignParserMIN = 36
	parkingsignParserTEMPORARY = 37
	parkingsignParserPASSENGER = 38
	parkingsignParserLOADING = 39
	parkingsignParserMON = 40
	parkingsignParserTUE = 41
	parkingsignParserWED = 42
	parkingsignParserTHU = 43
	parkingsignParserFRI = 44
	parkingsignParserSAT = 45
	parkingsignParserSUN = 46
	parkingsignParserSTREET = 47
	parkingsignParserSWEEPING = 48
	parkingsignParserCLEANING = 49
	parkingsignParserNOON = 50
	parkingsignParserMIDNIGHT = 51
	parkingsignParserINT = 52
	parkingsignParserWS = 53
)

// parkingsignParser rules.
const (
	parkingsignParserRULE_parkingSigns = 0
	parkingsignParserRULE_parkingSign = 1
	parkingsignParserRULE_streetSweepingSign = 2
	parkingsignParserRULE_noParkingSign = 3
	parkingsignParserRULE_noStoppingSign = 4
	parkingsignParserRULE_passengerLoadingSign = 5
	parkingsignParserRULE_temporaryNoParkingSign = 6
	parkingsignParserRULE_singleTimeLimitSign = 7
	parkingsignParserRULE_doubleTimeLimitSign = 8
	parkingsignParserRULE_permitSign = 9
	parkingsignParserRULE_streetSweeping = 10
	parkingsignParserRULE_noParking = 11
	parkingsignParserRULE_noStopping = 12
	parkingsignParserRULE_valetOnly = 13
	parkingsignParserRULE_loadingOnly = 14
	parkingsignParserRULE_schoolDays = 15
	parkingsignParserRULE_timeRange = 16
	parkingsignParserRULE_everyDay = 17
	parkingsignParserRULE_dayToDay = 18
	parkingsignParserRULE_dayAndDay = 19
	parkingsignParserRULE_dayRange = 20
	parkingsignParserRULE_dayRangePlus = 21
	parkingsignParserRULE_to = 22
	parkingsignParserRULE_and_ = 23
	parkingsignParserRULE_towAway = 24
	parkingsignParserRULE_minute = 25
	parkingsignParserRULE_exempt = 26
	parkingsignParserRULE_anyTime = 27
	parkingsignParserRULE_day = 28
	parkingsignParserRULE_time = 29
	parkingsignParserRULE_twelveNoon = 30
	parkingsignParserRULE_twelveMidnight = 31
	parkingsignParserRULE_am = 32
	parkingsignParserRULE_pm = 33
)

// IParkingSignsContext is an interface to support dynamic dispatch.
type IParkingSignsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParkingSignsContext differentiates from other interfaces.
	IsParkingSignsContext()
}

type ParkingSignsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParkingSignsContext() *ParkingSignsContext {
	var p = new(ParkingSignsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_parkingSigns
	return p
}

func (*ParkingSignsContext) IsParkingSignsContext() {}

func NewParkingSignsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParkingSignsContext {
	var p = new(ParkingSignsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_parkingSigns

	return p
}

func (s *ParkingSignsContext) GetParser() antlr.Parser { return s.parser }

func (s *ParkingSignsContext) AllParkingSign() []IParkingSignContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParkingSignContext)(nil)).Elem())
	var tst = make([]IParkingSignContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParkingSignContext)
		}
	}

	return tst
}

func (s *ParkingSignsContext) ParkingSign(i int) IParkingSignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParkingSignContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParkingSignContext)
}

func (s *ParkingSignsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParkingSignsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParkingSignsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterParkingSigns(s)
	}
}

func (s *ParkingSignsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitParkingSigns(s)
	}
}




func (p *parkingsignParser) ParkingSigns() (localctx IParkingSignsContext) {
	this := p
	_ = this

	localctx = NewParkingSignsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, parkingsignParserRULE_parkingSigns)
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
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	for (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << parkingsignParserT__1) | (1 << parkingsignParserNO) | (1 << parkingsignParserANYTIME) | (1 << parkingsignParserANY) | (1 << parkingsignParserTOW) | (1 << parkingsignParserVALET) | (1 << parkingsignParserVEHICLES) | (1 << parkingsignParserDISTRICT))) != 0) || ((((_la - 38)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 38))) & ((1 << (parkingsignParserPASSENGER - 38)) | (1 << (parkingsignParserNOON - 38)) | (1 << (parkingsignParserMIDNIGHT - 38)) | (1 << (parkingsignParserINT - 38)))) != 0) {
		{
			p.SetState(68)
			p.ParkingSign()
		}


		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}



	return localctx
}


// IParkingSignContext is an interface to support dynamic dispatch.
type IParkingSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParkingSignContext differentiates from other interfaces.
	IsParkingSignContext()
}

type ParkingSignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParkingSignContext() *ParkingSignContext {
	var p = new(ParkingSignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_parkingSign
	return p
}

func (*ParkingSignContext) IsParkingSignContext() {}

func NewParkingSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParkingSignContext {
	var p = new(ParkingSignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_parkingSign

	return p
}

func (s *ParkingSignContext) GetParser() antlr.Parser { return s.parser }

func (s *ParkingSignContext) StreetSweepingSign() IStreetSweepingSignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStreetSweepingSignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStreetSweepingSignContext)
}

func (s *ParkingSignContext) NoParkingSign() INoParkingSignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoParkingSignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoParkingSignContext)
}

func (s *ParkingSignContext) NoStoppingSign() INoStoppingSignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoStoppingSignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoStoppingSignContext)
}

func (s *ParkingSignContext) PassengerLoadingSign() IPassengerLoadingSignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPassengerLoadingSignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPassengerLoadingSignContext)
}

func (s *ParkingSignContext) SingleTimeLimitSign() ISingleTimeLimitSignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingleTimeLimitSignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingleTimeLimitSignContext)
}

func (s *ParkingSignContext) DoubleTimeLimitSign() IDoubleTimeLimitSignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDoubleTimeLimitSignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDoubleTimeLimitSignContext)
}

func (s *ParkingSignContext) TemporaryNoParkingSign() ITemporaryNoParkingSignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemporaryNoParkingSignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITemporaryNoParkingSignContext)
}

func (s *ParkingSignContext) PermitSign() IPermitSignContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPermitSignContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPermitSignContext)
}

func (s *ParkingSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParkingSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ParkingSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterParkingSign(s)
	}
}

func (s *ParkingSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitParkingSign(s)
	}
}




func (p *parkingsignParser) ParkingSign() (localctx IParkingSignContext) {
	this := p
	_ = this

	localctx = NewParkingSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, parkingsignParserRULE_parkingSign)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.StreetSweepingSign()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.NoParkingSign()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(76)
			p.NoStoppingSign()
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(77)
			p.PassengerLoadingSign()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(78)
			p.SingleTimeLimitSign()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(79)
			p.DoubleTimeLimitSign()
		}


	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(80)
			p.TemporaryNoParkingSign()
		}


	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(81)
			p.PermitSign()
		}

	}


	return localctx
}


// IStreetSweepingSignContext is an interface to support dynamic dispatch.
type IStreetSweepingSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStreetSweepingSignContext differentiates from other interfaces.
	IsStreetSweepingSignContext()
}

type StreetSweepingSignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStreetSweepingSignContext() *StreetSweepingSignContext {
	var p = new(StreetSweepingSignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_streetSweepingSign
	return p
}

func (*StreetSweepingSignContext) IsStreetSweepingSignContext() {}

func NewStreetSweepingSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StreetSweepingSignContext {
	var p = new(StreetSweepingSignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_streetSweepingSign

	return p
}

func (s *StreetSweepingSignContext) GetParser() antlr.Parser { return s.parser }

func (s *StreetSweepingSignContext) TimeRange() ITimeRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeRangeContext)
}

func (s *StreetSweepingSignContext) Day() IDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *StreetSweepingSignContext) StreetSweeping() IStreetSweepingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStreetSweepingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStreetSweepingContext)
}

func (s *StreetSweepingSignContext) NoParking() INoParkingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoParkingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoParkingContext)
}

func (s *StreetSweepingSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StreetSweepingSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StreetSweepingSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterStreetSweepingSign(s)
	}
}

func (s *StreetSweepingSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitStreetSweepingSign(s)
	}
}




func (p *parkingsignParser) StreetSweepingSign() (localctx IStreetSweepingSignContext) {
	this := p
	_ = this

	localctx = NewStreetSweepingSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, parkingsignParserRULE_streetSweepingSign)
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


	if _la == parkingsignParserNO {
		{
			p.SetState(84)
			p.NoParking()
		}

	}
	{
		p.SetState(87)
		p.TimeRange()
	}
	{
		p.SetState(88)
		p.Day()
	}
	{
		p.SetState(89)
		p.StreetSweeping()
	}



	return localctx
}


// INoParkingSignContext is an interface to support dynamic dispatch.
type INoParkingSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoParkingSignContext differentiates from other interfaces.
	IsNoParkingSignContext()
}

type NoParkingSignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoParkingSignContext() *NoParkingSignContext {
	var p = new(NoParkingSignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_noParkingSign
	return p
}

func (*NoParkingSignContext) IsNoParkingSignContext() {}

func NewNoParkingSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoParkingSignContext {
	var p = new(NoParkingSignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_noParkingSign

	return p
}

func (s *NoParkingSignContext) GetParser() antlr.Parser { return s.parser }

func (s *NoParkingSignContext) AnyTime() IAnyTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnyTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnyTimeContext)
}

func (s *NoParkingSignContext) TimeRange() ITimeRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeRangeContext)
}

func (s *NoParkingSignContext) NoParking() INoParkingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoParkingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoParkingContext)
}

func (s *NoParkingSignContext) EXCEPT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserEXCEPT, 0)
}

func (s *NoParkingSignContext) DayRange() IDayRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayRangeContext)
}

func (s *NoParkingSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoParkingSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NoParkingSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterNoParkingSign(s)
	}
}

func (s *NoParkingSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitNoParkingSign(s)
	}
}




func (p *parkingsignParser) NoParkingSign() (localctx INoParkingSignContext) {
	this := p
	_ = this

	localctx = NewNoParkingSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, parkingsignParserRULE_noParkingSign)
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
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == parkingsignParserNO {
		{
			p.SetState(91)
			p.NoParking()
		}

	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case parkingsignParserANYTIME, parkingsignParserANY:
		{
			p.SetState(94)
			p.AnyTime()
		}


	case parkingsignParserT__1, parkingsignParserNOON, parkingsignParserMIDNIGHT, parkingsignParserINT:
		{
			p.SetState(95)
			p.TimeRange()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == parkingsignParserEXCEPT {
		{
			p.SetState(98)
			p.Match(parkingsignParserEXCEPT)
		}

	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if ((((_la - 16)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 16))) & ((1 << (parkingsignParserDAILY - 16)) | (1 << (parkingsignParserNIGHTLY - 16)) | (1 << (parkingsignParserSCHOOL - 16)) | (1 << (parkingsignParserHOLIDAYS - 16)) | (1 << (parkingsignParserMON - 16)) | (1 << (parkingsignParserTUE - 16)) | (1 << (parkingsignParserWED - 16)) | (1 << (parkingsignParserTHU - 16)) | (1 << (parkingsignParserFRI - 16)) | (1 << (parkingsignParserSAT - 16)) | (1 << (parkingsignParserSUN - 16)))) != 0) {
		{
			p.SetState(101)
			p.DayRange()
		}

	}



	return localctx
}


// INoStoppingSignContext is an interface to support dynamic dispatch.
type INoStoppingSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoStoppingSignContext differentiates from other interfaces.
	IsNoStoppingSignContext()
}

type NoStoppingSignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoStoppingSignContext() *NoStoppingSignContext {
	var p = new(NoStoppingSignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_noStoppingSign
	return p
}

func (*NoStoppingSignContext) IsNoStoppingSignContext() {}

func NewNoStoppingSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoStoppingSignContext {
	var p = new(NoStoppingSignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_noStoppingSign

	return p
}

func (s *NoStoppingSignContext) GetParser() antlr.Parser { return s.parser }

func (s *NoStoppingSignContext) NoStopping() INoStoppingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoStoppingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoStoppingContext)
}

func (s *NoStoppingSignContext) ValetOnly() IValetOnlyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValetOnlyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValetOnlyContext)
}

func (s *NoStoppingSignContext) TowAway() ITowAwayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITowAwayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITowAwayContext)
}

func (s *NoStoppingSignContext) AllTimeRange() []ITimeRangeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITimeRangeContext)(nil)).Elem())
	var tst = make([]ITimeRangeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITimeRangeContext)
		}
	}

	return tst
}

func (s *NoStoppingSignContext) TimeRange(i int) ITimeRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeRangeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITimeRangeContext)
}

func (s *NoStoppingSignContext) EXCEPT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserEXCEPT, 0)
}

func (s *NoStoppingSignContext) DayRange() IDayRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayRangeContext)
}

func (s *NoStoppingSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoStoppingSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NoStoppingSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterNoStoppingSign(s)
	}
}

func (s *NoStoppingSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitNoStoppingSign(s)
	}
}




func (p *parkingsignParser) NoStoppingSign() (localctx INoStoppingSignContext) {
	this := p
	_ = this

	localctx = NewNoStoppingSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, parkingsignParserRULE_noStoppingSign)
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
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == parkingsignParserTOW {
		{
			p.SetState(104)
			p.TowAway()
		}

	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case parkingsignParserNO:
		{
			p.SetState(107)
			p.NoStopping()
		}


	case parkingsignParserVALET:
		{
			p.SetState(108)
			p.ValetOnly()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(111)
					p.TimeRange()
				}




		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext())
	}
	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == parkingsignParserEXCEPT {
		{
			p.SetState(116)
			p.Match(parkingsignParserEXCEPT)
		}

	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if ((((_la - 16)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 16))) & ((1 << (parkingsignParserDAILY - 16)) | (1 << (parkingsignParserNIGHTLY - 16)) | (1 << (parkingsignParserSCHOOL - 16)) | (1 << (parkingsignParserHOLIDAYS - 16)) | (1 << (parkingsignParserMON - 16)) | (1 << (parkingsignParserTUE - 16)) | (1 << (parkingsignParserWED - 16)) | (1 << (parkingsignParserTHU - 16)) | (1 << (parkingsignParserFRI - 16)) | (1 << (parkingsignParserSAT - 16)) | (1 << (parkingsignParserSUN - 16)))) != 0) {
		{
			p.SetState(119)
			p.DayRange()
		}

	}



	return localctx
}


// IPassengerLoadingSignContext is an interface to support dynamic dispatch.
type IPassengerLoadingSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPassengerLoadingSignContext differentiates from other interfaces.
	IsPassengerLoadingSignContext()
}

type PassengerLoadingSignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassengerLoadingSignContext() *PassengerLoadingSignContext {
	var p = new(PassengerLoadingSignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_passengerLoadingSign
	return p
}

func (*PassengerLoadingSignContext) IsPassengerLoadingSignContext() {}

func NewPassengerLoadingSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassengerLoadingSignContext {
	var p = new(PassengerLoadingSignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_passengerLoadingSign

	return p
}

func (s *PassengerLoadingSignContext) GetParser() antlr.Parser { return s.parser }

func (s *PassengerLoadingSignContext) LoadingOnly() ILoadingOnlyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoadingOnlyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoadingOnlyContext)
}

func (s *PassengerLoadingSignContext) TowAway() ITowAwayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITowAwayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITowAwayContext)
}

func (s *PassengerLoadingSignContext) AllTimeRange() []ITimeRangeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITimeRangeContext)(nil)).Elem())
	var tst = make([]ITimeRangeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITimeRangeContext)
		}
	}

	return tst
}

func (s *PassengerLoadingSignContext) TimeRange(i int) ITimeRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeRangeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITimeRangeContext)
}

func (s *PassengerLoadingSignContext) EXCEPT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserEXCEPT, 0)
}

func (s *PassengerLoadingSignContext) DayRange() IDayRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayRangeContext)
}

func (s *PassengerLoadingSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassengerLoadingSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PassengerLoadingSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterPassengerLoadingSign(s)
	}
}

func (s *PassengerLoadingSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitPassengerLoadingSign(s)
	}
}




func (p *parkingsignParser) PassengerLoadingSign() (localctx IPassengerLoadingSignContext) {
	this := p
	_ = this

	localctx = NewPassengerLoadingSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, parkingsignParserRULE_passengerLoadingSign)
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
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == parkingsignParserTOW {
		{
			p.SetState(122)
			p.TowAway()
		}

	}
	{
		p.SetState(125)
		p.LoadingOnly()
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
				{
					p.SetState(126)
					p.TimeRange()
				}




		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == parkingsignParserEXCEPT {
		{
			p.SetState(131)
			p.Match(parkingsignParserEXCEPT)
		}

	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if ((((_la - 16)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 16))) & ((1 << (parkingsignParserDAILY - 16)) | (1 << (parkingsignParserNIGHTLY - 16)) | (1 << (parkingsignParserSCHOOL - 16)) | (1 << (parkingsignParserHOLIDAYS - 16)) | (1 << (parkingsignParserMON - 16)) | (1 << (parkingsignParserTUE - 16)) | (1 << (parkingsignParserWED - 16)) | (1 << (parkingsignParserTHU - 16)) | (1 << (parkingsignParserFRI - 16)) | (1 << (parkingsignParserSAT - 16)) | (1 << (parkingsignParserSUN - 16)))) != 0) {
		{
			p.SetState(134)
			p.DayRange()
		}

	}



	return localctx
}


// ITemporaryNoParkingSignContext is an interface to support dynamic dispatch.
type ITemporaryNoParkingSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemporaryNoParkingSignContext differentiates from other interfaces.
	IsTemporaryNoParkingSignContext()
}

type TemporaryNoParkingSignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemporaryNoParkingSignContext() *TemporaryNoParkingSignContext {
	var p = new(TemporaryNoParkingSignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_temporaryNoParkingSign
	return p
}

func (*TemporaryNoParkingSignContext) IsTemporaryNoParkingSignContext() {}

func NewTemporaryNoParkingSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemporaryNoParkingSignContext {
	var p = new(TemporaryNoParkingSignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_temporaryNoParkingSign

	return p
}

func (s *TemporaryNoParkingSignContext) GetParser() antlr.Parser { return s.parser }

func (s *TemporaryNoParkingSignContext) TowAway() ITowAwayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITowAwayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITowAwayContext)
}

func (s *TemporaryNoParkingSignContext) TEMPORARY() antlr.TerminalNode {
	return s.GetToken(parkingsignParserTEMPORARY, 0)
}

func (s *TemporaryNoParkingSignContext) NoParking() INoParkingContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INoParkingContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INoParkingContext)
}

func (s *TemporaryNoParkingSignContext) AnyTime() IAnyTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnyTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnyTimeContext)
}

func (s *TemporaryNoParkingSignContext) TimeRange() ITimeRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeRangeContext)
}

func (s *TemporaryNoParkingSignContext) EXCEPT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserEXCEPT, 0)
}

func (s *TemporaryNoParkingSignContext) DayRange() IDayRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayRangeContext)
}

func (s *TemporaryNoParkingSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemporaryNoParkingSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TemporaryNoParkingSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterTemporaryNoParkingSign(s)
	}
}

func (s *TemporaryNoParkingSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitTemporaryNoParkingSign(s)
	}
}




func (p *parkingsignParser) TemporaryNoParkingSign() (localctx ITemporaryNoParkingSignContext) {
	this := p
	_ = this

	localctx = NewTemporaryNoParkingSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, parkingsignParserRULE_temporaryNoParkingSign)
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
		p.SetState(137)
		p.TowAway()
	}
	{
		p.SetState(138)
		p.Match(parkingsignParserTEMPORARY)
	}
	{
		p.SetState(139)
		p.NoParking()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case parkingsignParserANYTIME, parkingsignParserANY:
		{
			p.SetState(140)
			p.AnyTime()
		}


	case parkingsignParserT__1, parkingsignParserNOON, parkingsignParserMIDNIGHT, parkingsignParserINT:
		{
			p.SetState(141)
			p.TimeRange()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == parkingsignParserEXCEPT {
		{
			p.SetState(144)
			p.Match(parkingsignParserEXCEPT)
		}

	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if ((((_la - 16)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 16))) & ((1 << (parkingsignParserDAILY - 16)) | (1 << (parkingsignParserNIGHTLY - 16)) | (1 << (parkingsignParserSCHOOL - 16)) | (1 << (parkingsignParserHOLIDAYS - 16)) | (1 << (parkingsignParserMON - 16)) | (1 << (parkingsignParserTUE - 16)) | (1 << (parkingsignParserWED - 16)) | (1 << (parkingsignParserTHU - 16)) | (1 << (parkingsignParserFRI - 16)) | (1 << (parkingsignParserSAT - 16)) | (1 << (parkingsignParserSUN - 16)))) != 0) {
		{
			p.SetState(147)
			p.DayRange()
		}

	}



	return localctx
}


// ISingleTimeLimitSignContext is an interface to support dynamic dispatch.
type ISingleTimeLimitSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleTimeLimitSignContext differentiates from other interfaces.
	IsSingleTimeLimitSignContext()
}

type SingleTimeLimitSignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleTimeLimitSignContext() *SingleTimeLimitSignContext {
	var p = new(SingleTimeLimitSignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_singleTimeLimitSign
	return p
}

func (*SingleTimeLimitSignContext) IsSingleTimeLimitSignContext() {}

func NewSingleTimeLimitSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleTimeLimitSignContext {
	var p = new(SingleTimeLimitSignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_singleTimeLimitSign

	return p
}

func (s *SingleTimeLimitSignContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleTimeLimitSignContext) PARKING() antlr.TerminalNode {
	return s.GetToken(parkingsignParserPARKING, 0)
}

func (s *SingleTimeLimitSignContext) TimeRange() ITimeRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeRangeContext)
}

func (s *SingleTimeLimitSignContext) EXCEPT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserEXCEPT, 0)
}

func (s *SingleTimeLimitSignContext) DayRange() IDayRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayRangeContext)
}

func (s *SingleTimeLimitSignContext) INT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserINT, 0)
}

func (s *SingleTimeLimitSignContext) HOUR() antlr.TerminalNode {
	return s.GetToken(parkingsignParserHOUR, 0)
}

func (s *SingleTimeLimitSignContext) Minute() IMinuteContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMinuteContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMinuteContext)
}

func (s *SingleTimeLimitSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleTimeLimitSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SingleTimeLimitSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterSingleTimeLimitSign(s)
	}
}

func (s *SingleTimeLimitSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitSingleTimeLimitSign(s)
	}
}




func (p *parkingsignParser) SingleTimeLimitSign() (localctx ISingleTimeLimitSignContext) {
	this := p
	_ = this

	localctx = NewSingleTimeLimitSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, parkingsignParserRULE_singleTimeLimitSign)
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
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(150)
			p.Match(parkingsignParserINT)
		}
		{
			p.SetState(151)
			p.Match(parkingsignParserHOUR)
		}



	case 2:
		{
			p.SetState(152)
			p.Match(parkingsignParserINT)
		}
		{
			p.SetState(153)
			p.Minute()
		}


	}
	{
		p.SetState(156)
		p.Match(parkingsignParserPARKING)
	}
	{
		p.SetState(157)
		p.TimeRange()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == parkingsignParserEXCEPT {
		{
			p.SetState(158)
			p.Match(parkingsignParserEXCEPT)
		}

	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if ((((_la - 16)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 16))) & ((1 << (parkingsignParserDAILY - 16)) | (1 << (parkingsignParserNIGHTLY - 16)) | (1 << (parkingsignParserSCHOOL - 16)) | (1 << (parkingsignParserHOLIDAYS - 16)) | (1 << (parkingsignParserMON - 16)) | (1 << (parkingsignParserTUE - 16)) | (1 << (parkingsignParserWED - 16)) | (1 << (parkingsignParserTHU - 16)) | (1 << (parkingsignParserFRI - 16)) | (1 << (parkingsignParserSAT - 16)) | (1 << (parkingsignParserSUN - 16)))) != 0) {
		{
			p.SetState(161)
			p.DayRange()
		}

	}



	return localctx
}


// IDoubleTimeLimitSignContext is an interface to support dynamic dispatch.
type IDoubleTimeLimitSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDoubleTimeLimitSignContext differentiates from other interfaces.
	IsDoubleTimeLimitSignContext()
}

type DoubleTimeLimitSignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoubleTimeLimitSignContext() *DoubleTimeLimitSignContext {
	var p = new(DoubleTimeLimitSignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_doubleTimeLimitSign
	return p
}

func (*DoubleTimeLimitSignContext) IsDoubleTimeLimitSignContext() {}

func NewDoubleTimeLimitSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoubleTimeLimitSignContext {
	var p = new(DoubleTimeLimitSignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_doubleTimeLimitSign

	return p
}

func (s *DoubleTimeLimitSignContext) GetParser() antlr.Parser { return s.parser }

func (s *DoubleTimeLimitSignContext) INT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserINT, 0)
}

func (s *DoubleTimeLimitSignContext) HOUR() antlr.TerminalNode {
	return s.GetToken(parkingsignParserHOUR, 0)
}

func (s *DoubleTimeLimitSignContext) PARKING() antlr.TerminalNode {
	return s.GetToken(parkingsignParserPARKING, 0)
}

func (s *DoubleTimeLimitSignContext) AllDayRange() []IDayRangeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDayRangeContext)(nil)).Elem())
	var tst = make([]IDayRangeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDayRangeContext)
		}
	}

	return tst
}

func (s *DoubleTimeLimitSignContext) DayRange(i int) IDayRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayRangeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDayRangeContext)
}

func (s *DoubleTimeLimitSignContext) AllTime() []ITimeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITimeContext)(nil)).Elem())
	var tst = make([]ITimeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITimeContext)
		}
	}

	return tst
}

func (s *DoubleTimeLimitSignContext) Time(i int) ITimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITimeContext)
}

func (s *DoubleTimeLimitSignContext) AllTO() []antlr.TerminalNode {
	return s.GetTokens(parkingsignParserTO)
}

func (s *DoubleTimeLimitSignContext) TO(i int) antlr.TerminalNode {
	return s.GetToken(parkingsignParserTO, i)
}

func (s *DoubleTimeLimitSignContext) EXCEPT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserEXCEPT, 0)
}

func (s *DoubleTimeLimitSignContext) HOLIDAYS() antlr.TerminalNode {
	return s.GetToken(parkingsignParserHOLIDAYS, 0)
}

func (s *DoubleTimeLimitSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleTimeLimitSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DoubleTimeLimitSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterDoubleTimeLimitSign(s)
	}
}

func (s *DoubleTimeLimitSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitDoubleTimeLimitSign(s)
	}
}




func (p *parkingsignParser) DoubleTimeLimitSign() (localctx IDoubleTimeLimitSignContext) {
	this := p
	_ = this

	localctx = NewDoubleTimeLimitSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, parkingsignParserRULE_doubleTimeLimitSign)
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
		p.SetState(164)
		p.Match(parkingsignParserINT)
	}
	{
		p.SetState(165)
		p.Match(parkingsignParserHOUR)
	}
	{
		p.SetState(166)
		p.Match(parkingsignParserPARKING)
	}
	{
		p.SetState(167)
		p.DayRange()
	}
	{
		p.SetState(168)
		p.DayRange()
	}
	{
		p.SetState(169)
		p.Time()
	}
	{
		p.SetState(170)
		p.Time()
	}
	{
		p.SetState(171)
		p.Match(parkingsignParserTO)
	}
	{
		p.SetState(172)
		p.Match(parkingsignParserTO)
	}
	{
		p.SetState(173)
		p.Time()
	}
	{
		p.SetState(174)
		p.Time()
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == parkingsignParserEXCEPT {
		{
			p.SetState(175)
			p.Match(parkingsignParserEXCEPT)
		}
		{
			p.SetState(176)
			p.Match(parkingsignParserHOLIDAYS)
		}

	}



	return localctx
}


// IPermitSignContext is an interface to support dynamic dispatch.
type IPermitSignContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPermitSignContext differentiates from other interfaces.
	IsPermitSignContext()
}

type PermitSignContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPermitSignContext() *PermitSignContext {
	var p = new(PermitSignContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_permitSign
	return p
}

func (*PermitSignContext) IsPermitSignContext() {}

func NewPermitSignContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PermitSignContext {
	var p = new(PermitSignContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_permitSign

	return p
}

func (s *PermitSignContext) GetParser() antlr.Parser { return s.parser }

func (s *PermitSignContext) DISTRICT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserDISTRICT, 0)
}

func (s *PermitSignContext) NO() antlr.TerminalNode {
	return s.GetToken(parkingsignParserNO, 0)
}

func (s *PermitSignContext) INT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserINT, 0)
}

func (s *PermitSignContext) PERMITS() antlr.TerminalNode {
	return s.GetToken(parkingsignParserPERMITS, 0)
}

func (s *PermitSignContext) Exempt() IExemptContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExemptContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExemptContext)
}

func (s *PermitSignContext) VEHICLES() antlr.TerminalNode {
	return s.GetToken(parkingsignParserVEHICLES, 0)
}

func (s *PermitSignContext) WITH() antlr.TerminalNode {
	return s.GetToken(parkingsignParserWITH, 0)
}

func (s *PermitSignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PermitSignContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PermitSignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterPermitSign(s)
	}
}

func (s *PermitSignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitPermitSign(s)
	}
}




func (p *parkingsignParser) PermitSign() (localctx IPermitSignContext) {
	this := p
	_ = this

	localctx = NewPermitSignContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, parkingsignParserRULE_permitSign)
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
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == parkingsignParserVEHICLES {
		{
			p.SetState(179)
			p.Match(parkingsignParserVEHICLES)
		}
		{
			p.SetState(180)
			p.Match(parkingsignParserWITH)
		}

	}
	{
		p.SetState(183)
		p.Match(parkingsignParserDISTRICT)
	}
	{
		p.SetState(184)
		p.Match(parkingsignParserNO)
	}
	{
		p.SetState(185)
		p.Match(parkingsignParserINT)
	}
	{
		p.SetState(186)
		p.Match(parkingsignParserPERMITS)
	}
	{
		p.SetState(187)
		p.Exempt()
	}



	return localctx
}


// IStreetSweepingContext is an interface to support dynamic dispatch.
type IStreetSweepingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStreetSweepingContext differentiates from other interfaces.
	IsStreetSweepingContext()
}

type StreetSweepingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStreetSweepingContext() *StreetSweepingContext {
	var p = new(StreetSweepingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_streetSweeping
	return p
}

func (*StreetSweepingContext) IsStreetSweepingContext() {}

func NewStreetSweepingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StreetSweepingContext {
	var p = new(StreetSweepingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_streetSweeping

	return p
}

func (s *StreetSweepingContext) GetParser() antlr.Parser { return s.parser }

func (s *StreetSweepingContext) STREET() antlr.TerminalNode {
	return s.GetToken(parkingsignParserSTREET, 0)
}

func (s *StreetSweepingContext) SWEEPING() antlr.TerminalNode {
	return s.GetToken(parkingsignParserSWEEPING, 0)
}

func (s *StreetSweepingContext) CLEANING() antlr.TerminalNode {
	return s.GetToken(parkingsignParserCLEANING, 0)
}

func (s *StreetSweepingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StreetSweepingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *StreetSweepingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterStreetSweeping(s)
	}
}

func (s *StreetSweepingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitStreetSweeping(s)
	}
}




func (p *parkingsignParser) StreetSweeping() (localctx IStreetSweepingContext) {
	this := p
	_ = this

	localctx = NewStreetSweepingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, parkingsignParserRULE_streetSweeping)
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
		p.SetState(189)
		p.Match(parkingsignParserSTREET)
	}
	{
		p.SetState(190)
		_la = p.GetTokenStream().LA(1)

		if !(_la == parkingsignParserSWEEPING || _la == parkingsignParserCLEANING) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// INoParkingContext is an interface to support dynamic dispatch.
type INoParkingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoParkingContext differentiates from other interfaces.
	IsNoParkingContext()
}

type NoParkingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoParkingContext() *NoParkingContext {
	var p = new(NoParkingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_noParking
	return p
}

func (*NoParkingContext) IsNoParkingContext() {}

func NewNoParkingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoParkingContext {
	var p = new(NoParkingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_noParking

	return p
}

func (s *NoParkingContext) GetParser() antlr.Parser { return s.parser }

func (s *NoParkingContext) NO() antlr.TerminalNode {
	return s.GetToken(parkingsignParserNO, 0)
}

func (s *NoParkingContext) PARKING() antlr.TerminalNode {
	return s.GetToken(parkingsignParserPARKING, 0)
}

func (s *NoParkingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoParkingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NoParkingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterNoParking(s)
	}
}

func (s *NoParkingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitNoParking(s)
	}
}




func (p *parkingsignParser) NoParking() (localctx INoParkingContext) {
	this := p
	_ = this

	localctx = NewNoParkingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, parkingsignParserRULE_noParking)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(192)
		p.Match(parkingsignParserNO)
	}
	{
		p.SetState(193)
		p.Match(parkingsignParserPARKING)
	}



	return localctx
}


// INoStoppingContext is an interface to support dynamic dispatch.
type INoStoppingContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNoStoppingContext differentiates from other interfaces.
	IsNoStoppingContext()
}

type NoStoppingContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNoStoppingContext() *NoStoppingContext {
	var p = new(NoStoppingContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_noStopping
	return p
}

func (*NoStoppingContext) IsNoStoppingContext() {}

func NewNoStoppingContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NoStoppingContext {
	var p = new(NoStoppingContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_noStopping

	return p
}

func (s *NoStoppingContext) GetParser() antlr.Parser { return s.parser }

func (s *NoStoppingContext) NO() antlr.TerminalNode {
	return s.GetToken(parkingsignParserNO, 0)
}

func (s *NoStoppingContext) STOPPING() antlr.TerminalNode {
	return s.GetToken(parkingsignParserSTOPPING, 0)
}

func (s *NoStoppingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NoStoppingContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *NoStoppingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterNoStopping(s)
	}
}

func (s *NoStoppingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitNoStopping(s)
	}
}




func (p *parkingsignParser) NoStopping() (localctx INoStoppingContext) {
	this := p
	_ = this

	localctx = NewNoStoppingContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, parkingsignParserRULE_noStopping)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(195)
		p.Match(parkingsignParserNO)
	}
	{
		p.SetState(196)
		p.Match(parkingsignParserSTOPPING)
	}



	return localctx
}


// IValetOnlyContext is an interface to support dynamic dispatch.
type IValetOnlyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValetOnlyContext differentiates from other interfaces.
	IsValetOnlyContext()
}

type ValetOnlyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValetOnlyContext() *ValetOnlyContext {
	var p = new(ValetOnlyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_valetOnly
	return p
}

func (*ValetOnlyContext) IsValetOnlyContext() {}

func NewValetOnlyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValetOnlyContext {
	var p = new(ValetOnlyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_valetOnly

	return p
}

func (s *ValetOnlyContext) GetParser() antlr.Parser { return s.parser }

func (s *ValetOnlyContext) VALET() antlr.TerminalNode {
	return s.GetToken(parkingsignParserVALET, 0)
}

func (s *ValetOnlyContext) PARKING() antlr.TerminalNode {
	return s.GetToken(parkingsignParserPARKING, 0)
}

func (s *ValetOnlyContext) ONLY() antlr.TerminalNode {
	return s.GetToken(parkingsignParserONLY, 0)
}

func (s *ValetOnlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValetOnlyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ValetOnlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterValetOnly(s)
	}
}

func (s *ValetOnlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitValetOnly(s)
	}
}




func (p *parkingsignParser) ValetOnly() (localctx IValetOnlyContext) {
	this := p
	_ = this

	localctx = NewValetOnlyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, parkingsignParserRULE_valetOnly)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(198)
		p.Match(parkingsignParserVALET)
	}
	{
		p.SetState(199)
		p.Match(parkingsignParserPARKING)
	}
	{
		p.SetState(200)
		p.Match(parkingsignParserONLY)
	}



	return localctx
}


// ILoadingOnlyContext is an interface to support dynamic dispatch.
type ILoadingOnlyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLoadingOnlyContext differentiates from other interfaces.
	IsLoadingOnlyContext()
}

type LoadingOnlyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoadingOnlyContext() *LoadingOnlyContext {
	var p = new(LoadingOnlyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_loadingOnly
	return p
}

func (*LoadingOnlyContext) IsLoadingOnlyContext() {}

func NewLoadingOnlyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoadingOnlyContext {
	var p = new(LoadingOnlyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_loadingOnly

	return p
}

func (s *LoadingOnlyContext) GetParser() antlr.Parser { return s.parser }

func (s *LoadingOnlyContext) PASSENGER() antlr.TerminalNode {
	return s.GetToken(parkingsignParserPASSENGER, 0)
}

func (s *LoadingOnlyContext) LOADING() antlr.TerminalNode {
	return s.GetToken(parkingsignParserLOADING, 0)
}

func (s *LoadingOnlyContext) ONLY() antlr.TerminalNode {
	return s.GetToken(parkingsignParserONLY, 0)
}

func (s *LoadingOnlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoadingOnlyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *LoadingOnlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterLoadingOnly(s)
	}
}

func (s *LoadingOnlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitLoadingOnly(s)
	}
}




func (p *parkingsignParser) LoadingOnly() (localctx ILoadingOnlyContext) {
	this := p
	_ = this

	localctx = NewLoadingOnlyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, parkingsignParserRULE_loadingOnly)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(202)
		p.Match(parkingsignParserPASSENGER)
	}
	{
		p.SetState(203)
		p.Match(parkingsignParserLOADING)
	}
	{
		p.SetState(204)
		p.Match(parkingsignParserONLY)
	}



	return localctx
}


// ISchoolDaysContext is an interface to support dynamic dispatch.
type ISchoolDaysContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchoolDaysContext differentiates from other interfaces.
	IsSchoolDaysContext()
}

type SchoolDaysContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchoolDaysContext() *SchoolDaysContext {
	var p = new(SchoolDaysContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_schoolDays
	return p
}

func (*SchoolDaysContext) IsSchoolDaysContext() {}

func NewSchoolDaysContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchoolDaysContext {
	var p = new(SchoolDaysContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_schoolDays

	return p
}

func (s *SchoolDaysContext) GetParser() antlr.Parser { return s.parser }

func (s *SchoolDaysContext) SCHOOL() antlr.TerminalNode {
	return s.GetToken(parkingsignParserSCHOOL, 0)
}

func (s *SchoolDaysContext) DAYS() antlr.TerminalNode {
	return s.GetToken(parkingsignParserDAYS, 0)
}

func (s *SchoolDaysContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchoolDaysContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *SchoolDaysContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterSchoolDays(s)
	}
}

func (s *SchoolDaysContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitSchoolDays(s)
	}
}




func (p *parkingsignParser) SchoolDays() (localctx ISchoolDaysContext) {
	this := p
	_ = this

	localctx = NewSchoolDaysContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, parkingsignParserRULE_schoolDays)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Match(parkingsignParserSCHOOL)
	}
	{
		p.SetState(207)
		p.Match(parkingsignParserDAYS)
	}



	return localctx
}


// ITimeRangeContext is an interface to support dynamic dispatch.
type ITimeRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeRangeContext differentiates from other interfaces.
	IsTimeRangeContext()
}

type TimeRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeRangeContext() *TimeRangeContext {
	var p = new(TimeRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_timeRange
	return p
}

func (*TimeRangeContext) IsTimeRangeContext() {}

func NewTimeRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeRangeContext {
	var p = new(TimeRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_timeRange

	return p
}

func (s *TimeRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeRangeContext) AllTime() []ITimeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITimeContext)(nil)).Elem())
	var tst = make([]ITimeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITimeContext)
		}
	}

	return tst
}

func (s *TimeRangeContext) Time(i int) ITimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITimeContext)
}

func (s *TimeRangeContext) To() IToContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IToContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IToContext)
}

func (s *TimeRangeContext) INT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserINT, 0)
}

func (s *TimeRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TimeRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterTimeRange(s)
	}
}

func (s *TimeRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitTimeRange(s)
	}
}




func (p *parkingsignParser) TimeRange() (localctx ITimeRangeContext) {
	this := p
	_ = this

	localctx = NewTimeRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, parkingsignParserRULE_timeRange)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(217)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.Time()
		}
		{
			p.SetState(210)
			p.To()
		}
		{
			p.SetState(211)
			p.Time()
		}



	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(213)
			p.Match(parkingsignParserINT)
		}
		{
			p.SetState(214)
			p.To()
		}
		{
			p.SetState(215)
			p.Time()
		}


	}


	return localctx
}


// IEveryDayContext is an interface to support dynamic dispatch.
type IEveryDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEveryDayContext differentiates from other interfaces.
	IsEveryDayContext()
}

type EveryDayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEveryDayContext() *EveryDayContext {
	var p = new(EveryDayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_everyDay
	return p
}

func (*EveryDayContext) IsEveryDayContext() {}

func NewEveryDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EveryDayContext {
	var p = new(EveryDayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_everyDay

	return p
}

func (s *EveryDayContext) GetParser() antlr.Parser { return s.parser }

func (s *EveryDayContext) DAILY() antlr.TerminalNode {
	return s.GetToken(parkingsignParserDAILY, 0)
}

func (s *EveryDayContext) NIGHTLY() antlr.TerminalNode {
	return s.GetToken(parkingsignParserNIGHTLY, 0)
}

func (s *EveryDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EveryDayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *EveryDayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterEveryDay(s)
	}
}

func (s *EveryDayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitEveryDay(s)
	}
}




func (p *parkingsignParser) EveryDay() (localctx IEveryDayContext) {
	this := p
	_ = this

	localctx = NewEveryDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, parkingsignParserRULE_everyDay)
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
		p.SetState(219)
		_la = p.GetTokenStream().LA(1)

		if !(_la == parkingsignParserDAILY || _la == parkingsignParserNIGHTLY) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IDayToDayContext is an interface to support dynamic dispatch.
type IDayToDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDayToDayContext differentiates from other interfaces.
	IsDayToDayContext()
}

type DayToDayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDayToDayContext() *DayToDayContext {
	var p = new(DayToDayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_dayToDay
	return p
}

func (*DayToDayContext) IsDayToDayContext() {}

func NewDayToDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DayToDayContext {
	var p = new(DayToDayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_dayToDay

	return p
}

func (s *DayToDayContext) GetParser() antlr.Parser { return s.parser }

func (s *DayToDayContext) AllDay() []IDayContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDayContext)(nil)).Elem())
	var tst = make([]IDayContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDayContext)
		}
	}

	return tst
}

func (s *DayToDayContext) Day(i int) IDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *DayToDayContext) To() IToContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IToContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IToContext)
}

func (s *DayToDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayToDayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DayToDayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterDayToDay(s)
	}
}

func (s *DayToDayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitDayToDay(s)
	}
}




func (p *parkingsignParser) DayToDay() (localctx IDayToDayContext) {
	this := p
	_ = this

	localctx = NewDayToDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, parkingsignParserRULE_dayToDay)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.Day()
	}
	{
		p.SetState(222)
		p.To()
	}
	{
		p.SetState(223)
		p.Day()
	}



	return localctx
}


// IDayAndDayContext is an interface to support dynamic dispatch.
type IDayAndDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDayAndDayContext differentiates from other interfaces.
	IsDayAndDayContext()
}

type DayAndDayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDayAndDayContext() *DayAndDayContext {
	var p = new(DayAndDayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_dayAndDay
	return p
}

func (*DayAndDayContext) IsDayAndDayContext() {}

func NewDayAndDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DayAndDayContext {
	var p = new(DayAndDayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_dayAndDay

	return p
}

func (s *DayAndDayContext) GetParser() antlr.Parser { return s.parser }

func (s *DayAndDayContext) AllDay() []IDayContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDayContext)(nil)).Elem())
	var tst = make([]IDayContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDayContext)
		}
	}

	return tst
}

func (s *DayAndDayContext) Day(i int) IDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *DayAndDayContext) And_() IAnd_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAnd_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAnd_Context)
}

func (s *DayAndDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayAndDayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DayAndDayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterDayAndDay(s)
	}
}

func (s *DayAndDayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitDayAndDay(s)
	}
}




func (p *parkingsignParser) DayAndDay() (localctx IDayAndDayContext) {
	this := p
	_ = this

	localctx = NewDayAndDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, parkingsignParserRULE_dayAndDay)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.SetState(225)
		p.Day()
	}
	{
		p.SetState(226)
		p.And_()
	}
	{
		p.SetState(227)
		p.Day()
	}



	return localctx
}


// IDayRangeContext is an interface to support dynamic dispatch.
type IDayRangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDayRangeContext differentiates from other interfaces.
	IsDayRangeContext()
}

type DayRangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDayRangeContext() *DayRangeContext {
	var p = new(DayRangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_dayRange
	return p
}

func (*DayRangeContext) IsDayRangeContext() {}

func NewDayRangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DayRangeContext {
	var p = new(DayRangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_dayRange

	return p
}

func (s *DayRangeContext) GetParser() antlr.Parser { return s.parser }

func (s *DayRangeContext) EveryDay() IEveryDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEveryDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEveryDayContext)
}

func (s *DayRangeContext) SchoolDays() ISchoolDaysContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchoolDaysContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchoolDaysContext)
}

func (s *DayRangeContext) HOLIDAYS() antlr.TerminalNode {
	return s.GetToken(parkingsignParserHOLIDAYS, 0)
}

func (s *DayRangeContext) DayAndDay() IDayAndDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayAndDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayAndDayContext)
}

func (s *DayRangeContext) DayToDay() IDayToDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayToDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayToDayContext)
}

func (s *DayRangeContext) Day() IDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayContext)
}

func (s *DayRangeContext) ONLY() antlr.TerminalNode {
	return s.GetToken(parkingsignParserONLY, 0)
}

func (s *DayRangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayRangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DayRangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterDayRange(s)
	}
}

func (s *DayRangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitDayRange(s)
	}
}




func (p *parkingsignParser) DayRange() (localctx IDayRangeContext) {
	this := p
	_ = this

	localctx = NewDayRangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, parkingsignParserRULE_dayRange)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.EveryDay()
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.SchoolDays()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(231)
			p.Match(parkingsignParserHOLIDAYS)
		}


	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(232)
			p.DayAndDay()
		}


	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(233)
			p.DayToDay()
		}


	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(234)
			p.Day()
		}
		{
			p.SetState(235)
			p.Match(parkingsignParserONLY)
		}


	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(237)
			p.Day()
		}

	}


	return localctx
}


// IDayRangePlusContext is an interface to support dynamic dispatch.
type IDayRangePlusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDayRangePlusContext differentiates from other interfaces.
	IsDayRangePlusContext()
}

type DayRangePlusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDayRangePlusContext() *DayRangePlusContext {
	var p = new(DayRangePlusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_dayRangePlus
	return p
}

func (*DayRangePlusContext) IsDayRangePlusContext() {}

func NewDayRangePlusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DayRangePlusContext {
	var p = new(DayRangePlusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_dayRangePlus

	return p
}

func (s *DayRangePlusContext) GetParser() antlr.Parser { return s.parser }

func (s *DayRangePlusContext) DayRange() IDayRangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDayRangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDayRangeContext)
}

func (s *DayRangePlusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayRangePlusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DayRangePlusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterDayRangePlus(s)
	}
}

func (s *DayRangePlusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitDayRangePlus(s)
	}
}




func (p *parkingsignParser) DayRangePlus() (localctx IDayRangePlusContext) {
	this := p
	_ = this

	localctx = NewDayRangePlusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, parkingsignParserRULE_dayRangePlus)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
		p.DayRange()
	}



	return localctx
}


// IToContext is an interface to support dynamic dispatch.
type IToContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsToContext differentiates from other interfaces.
	IsToContext()
}

type ToContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyToContext() *ToContext {
	var p = new(ToContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_to
	return p
}

func (*ToContext) IsToContext() {}

func NewToContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ToContext {
	var p = new(ToContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_to

	return p
}

func (s *ToContext) GetParser() antlr.Parser { return s.parser }

func (s *ToContext) TO() antlr.TerminalNode {
	return s.GetToken(parkingsignParserTO, 0)
}

func (s *ToContext) DASH() antlr.TerminalNode {
	return s.GetToken(parkingsignParserDASH, 0)
}

func (s *ToContext) THRU() antlr.TerminalNode {
	return s.GetToken(parkingsignParserTHRU, 0)
}

func (s *ToContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ToContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ToContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterTo(s)
	}
}

func (s *ToContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitTo(s)
	}
}




func (p *parkingsignParser) To() (localctx IToContext) {
	this := p
	_ = this

	localctx = NewToContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, parkingsignParserRULE_to)
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
		p.SetState(242)
		_la = p.GetTokenStream().LA(1)

		if !((((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << parkingsignParserTO) | (1 << parkingsignParserTHRU) | (1 << parkingsignParserDASH))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IAnd_Context is an interface to support dynamic dispatch.
type IAnd_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnd_Context differentiates from other interfaces.
	IsAnd_Context()
}

type And_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnd_Context() *And_Context {
	var p = new(And_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_and_
	return p
}

func (*And_Context) IsAnd_Context() {}

func NewAnd_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *And_Context {
	var p = new(And_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_and_

	return p
}

func (s *And_Context) GetParser() antlr.Parser { return s.parser }

func (s *And_Context) AND() antlr.TerminalNode {
	return s.GetToken(parkingsignParserAND, 0)
}

func (s *And_Context) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(parkingsignParserAMPERSAND, 0)
}

func (s *And_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *And_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterAnd_(s)
	}
}

func (s *And_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitAnd_(s)
	}
}




func (p *parkingsignParser) And_() (localctx IAnd_Context) {
	this := p
	_ = this

	localctx = NewAnd_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, parkingsignParserRULE_and_)
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
		p.SetState(244)
		_la = p.GetTokenStream().LA(1)

		if !(_la == parkingsignParserAND || _la == parkingsignParserAMPERSAND) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// ITowAwayContext is an interface to support dynamic dispatch.
type ITowAwayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTowAwayContext differentiates from other interfaces.
	IsTowAwayContext()
}

type TowAwayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTowAwayContext() *TowAwayContext {
	var p = new(TowAwayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_towAway
	return p
}

func (*TowAwayContext) IsTowAwayContext() {}

func NewTowAwayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TowAwayContext {
	var p = new(TowAwayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_towAway

	return p
}

func (s *TowAwayContext) GetParser() antlr.Parser { return s.parser }

func (s *TowAwayContext) TOW() antlr.TerminalNode {
	return s.GetToken(parkingsignParserTOW, 0)
}

func (s *TowAwayContext) AWAY() antlr.TerminalNode {
	return s.GetToken(parkingsignParserAWAY, 0)
}

func (s *TowAwayContext) DASH() antlr.TerminalNode {
	return s.GetToken(parkingsignParserDASH, 0)
}

func (s *TowAwayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TowAwayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TowAwayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterTowAway(s)
	}
}

func (s *TowAwayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitTowAway(s)
	}
}




func (p *parkingsignParser) TowAway() (localctx ITowAwayContext) {
	this := p
	_ = this

	localctx = NewTowAwayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, parkingsignParserRULE_towAway)
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
		p.SetState(246)
		p.Match(parkingsignParserTOW)
	}
	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == parkingsignParserDASH {
		{
			p.SetState(247)
			p.Match(parkingsignParserDASH)
		}

	}
	{
		p.SetState(250)
		p.Match(parkingsignParserAWAY)
	}



	return localctx
}


// IMinuteContext is an interface to support dynamic dispatch.
type IMinuteContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMinuteContext differentiates from other interfaces.
	IsMinuteContext()
}

type MinuteContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMinuteContext() *MinuteContext {
	var p = new(MinuteContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_minute
	return p
}

func (*MinuteContext) IsMinuteContext() {}

func NewMinuteContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MinuteContext {
	var p = new(MinuteContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_minute

	return p
}

func (s *MinuteContext) GetParser() antlr.Parser { return s.parser }

func (s *MinuteContext) MIN() antlr.TerminalNode {
	return s.GetToken(parkingsignParserMIN, 0)
}

func (s *MinuteContext) MINUTE() antlr.TerminalNode {
	return s.GetToken(parkingsignParserMINUTE, 0)
}

func (s *MinuteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinuteContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *MinuteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterMinute(s)
	}
}

func (s *MinuteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitMinute(s)
	}
}




func (p *parkingsignParser) Minute() (localctx IMinuteContext) {
	this := p
	_ = this

	localctx = NewMinuteContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, parkingsignParserRULE_minute)
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
		p.SetState(252)
		_la = p.GetTokenStream().LA(1)

		if !(_la == parkingsignParserMINUTE || _la == parkingsignParserMIN) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IExemptContext is an interface to support dynamic dispatch.
type IExemptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExemptContext differentiates from other interfaces.
	IsExemptContext()
}

type ExemptContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExemptContext() *ExemptContext {
	var p = new(ExemptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_exempt
	return p
}

func (*ExemptContext) IsExemptContext() {}

func NewExemptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExemptContext {
	var p = new(ExemptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_exempt

	return p
}

func (s *ExemptContext) GetParser() antlr.Parser { return s.parser }

func (s *ExemptContext) EXEMPT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserEXEMPT, 0)
}

func (s *ExemptContext) EXEMPTED() antlr.TerminalNode {
	return s.GetToken(parkingsignParserEXEMPTED, 0)
}

func (s *ExemptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExemptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *ExemptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterExempt(s)
	}
}

func (s *ExemptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitExempt(s)
	}
}




func (p *parkingsignParser) Exempt() (localctx IExemptContext) {
	this := p
	_ = this

	localctx = NewExemptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, parkingsignParserRULE_exempt)
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
		p.SetState(254)
		_la = p.GetTokenStream().LA(1)

		if !(_la == parkingsignParserEXEMPTED || _la == parkingsignParserEXEMPT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// IAnyTimeContext is an interface to support dynamic dispatch.
type IAnyTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAnyTimeContext differentiates from other interfaces.
	IsAnyTimeContext()
}

type AnyTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAnyTimeContext() *AnyTimeContext {
	var p = new(AnyTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_anyTime
	return p
}

func (*AnyTimeContext) IsAnyTimeContext() {}

func NewAnyTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AnyTimeContext {
	var p = new(AnyTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_anyTime

	return p
}

func (s *AnyTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *AnyTimeContext) ANYTIME() antlr.TerminalNode {
	return s.GetToken(parkingsignParserANYTIME, 0)
}

func (s *AnyTimeContext) ANY() antlr.TerminalNode {
	return s.GetToken(parkingsignParserANY, 0)
}

func (s *AnyTimeContext) TIME() antlr.TerminalNode {
	return s.GetToken(parkingsignParserTIME, 0)
}

func (s *AnyTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AnyTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AnyTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterAnyTime(s)
	}
}

func (s *AnyTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitAnyTime(s)
	}
}




func (p *parkingsignParser) AnyTime() (localctx IAnyTimeContext) {
	this := p
	_ = this

	localctx = NewAnyTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, parkingsignParserRULE_anyTime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case parkingsignParserANYTIME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Match(parkingsignParserANYTIME)
		}


	case parkingsignParserANY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(257)
			p.Match(parkingsignParserANY)
		}
		{
			p.SetState(258)
			p.Match(parkingsignParserTIME)
		}




	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IDayContext is an interface to support dynamic dispatch.
type IDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDayContext differentiates from other interfaces.
	IsDayContext()
}

type DayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDayContext() *DayContext {
	var p = new(DayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_day
	return p
}

func (*DayContext) IsDayContext() {}

func NewDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DayContext {
	var p = new(DayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_day

	return p
}

func (s *DayContext) GetParser() antlr.Parser { return s.parser }

func (s *DayContext) MON() antlr.TerminalNode {
	return s.GetToken(parkingsignParserMON, 0)
}

func (s *DayContext) TUE() antlr.TerminalNode {
	return s.GetToken(parkingsignParserTUE, 0)
}

func (s *DayContext) WED() antlr.TerminalNode {
	return s.GetToken(parkingsignParserWED, 0)
}

func (s *DayContext) THU() antlr.TerminalNode {
	return s.GetToken(parkingsignParserTHU, 0)
}

func (s *DayContext) FRI() antlr.TerminalNode {
	return s.GetToken(parkingsignParserFRI, 0)
}

func (s *DayContext) SAT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserSAT, 0)
}

func (s *DayContext) SUN() antlr.TerminalNode {
	return s.GetToken(parkingsignParserSUN, 0)
}

func (s *DayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *DayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterDay(s)
	}
}

func (s *DayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitDay(s)
	}
}




func (p *parkingsignParser) Day() (localctx IDayContext) {
	this := p
	_ = this

	localctx = NewDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, parkingsignParserRULE_day)
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
		p.SetState(261)
		_la = p.GetTokenStream().LA(1)

		if !(((((_la - 40)) & -(0x1f+1)) == 0 && ((1 << uint((_la - 40))) & ((1 << (parkingsignParserMON - 40)) | (1 << (parkingsignParserTUE - 40)) | (1 << (parkingsignParserWED - 40)) | (1 << (parkingsignParserTHU - 40)) | (1 << (parkingsignParserFRI - 40)) | (1 << (parkingsignParserSAT - 40)) | (1 << (parkingsignParserSUN - 40)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}



	return localctx
}


// ITimeContext is an interface to support dynamic dispatch.
type ITimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeContext differentiates from other interfaces.
	IsTimeContext()
}

type TimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeContext() *TimeContext {
	var p = new(TimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_time
	return p
}

func (*TimeContext) IsTimeContext() {}

func NewTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeContext {
	var p = new(TimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_time

	return p
}

func (s *TimeContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeContext) AllINT() []antlr.TerminalNode {
	return s.GetTokens(parkingsignParserINT)
}

func (s *TimeContext) INT(i int) antlr.TerminalNode {
	return s.GetToken(parkingsignParserINT, i)
}

func (s *TimeContext) Am() IAmContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAmContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAmContext)
}

func (s *TimeContext) Pm() IPmContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPmContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPmContext)
}

func (s *TimeContext) TwelveNoon() ITwelveNoonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITwelveNoonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITwelveNoonContext)
}

func (s *TimeContext) TwelveMidnight() ITwelveMidnightContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITwelveMidnightContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITwelveMidnightContext)
}

func (s *TimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterTime(s)
	}
}

func (s *TimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitTime(s)
	}
}




func (p *parkingsignParser) Time() (localctx ITimeContext) {
	this := p
	_ = this

	localctx = NewTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, parkingsignParserRULE_time)
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

	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.Match(parkingsignParserINT)
		}
		p.SetState(266)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)


		if _la == parkingsignParserT__0 {
			{
				p.SetState(264)
				p.Match(parkingsignParserT__0)
			}
			{
				p.SetState(265)
				p.Match(parkingsignParserINT)
			}

		}
		p.SetState(270)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case parkingsignParserT__2, parkingsignParserT__3:
			{
				p.SetState(268)
				p.Am()
			}


		case parkingsignParserT__4, parkingsignParserT__5:
			{
				p.SetState(269)
				p.Pm()
			}


		case parkingsignParserEOF, parkingsignParserT__1, parkingsignParserNO, parkingsignParserTO, parkingsignParserTHRU, parkingsignParserDASH, parkingsignParserANYTIME, parkingsignParserANY, parkingsignParserEXCEPT, parkingsignParserDAILY, parkingsignParserNIGHTLY, parkingsignParserSCHOOL, parkingsignParserHOLIDAYS, parkingsignParserTOW, parkingsignParserVALET, parkingsignParserVEHICLES, parkingsignParserDISTRICT, parkingsignParserPASSENGER, parkingsignParserMON, parkingsignParserTUE, parkingsignParserWED, parkingsignParserTHU, parkingsignParserFRI, parkingsignParserSAT, parkingsignParserSUN, parkingsignParserNOON, parkingsignParserMIDNIGHT, parkingsignParserINT:



		default:
		}


	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(272)
			p.TwelveNoon()
		}


	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(273)
			p.TwelveMidnight()
		}

	}


	return localctx
}


// ITwelveNoonContext is an interface to support dynamic dispatch.
type ITwelveNoonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTwelveNoonContext differentiates from other interfaces.
	IsTwelveNoonContext()
}

type TwelveNoonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTwelveNoonContext() *TwelveNoonContext {
	var p = new(TwelveNoonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_twelveNoon
	return p
}

func (*TwelveNoonContext) IsTwelveNoonContext() {}

func NewTwelveNoonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TwelveNoonContext {
	var p = new(TwelveNoonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_twelveNoon

	return p
}

func (s *TwelveNoonContext) GetParser() antlr.Parser { return s.parser }

func (s *TwelveNoonContext) NOON() antlr.TerminalNode {
	return s.GetToken(parkingsignParserNOON, 0)
}

func (s *TwelveNoonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TwelveNoonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TwelveNoonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterTwelveNoon(s)
	}
}

func (s *TwelveNoonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitTwelveNoon(s)
	}
}




func (p *parkingsignParser) TwelveNoon() (localctx ITwelveNoonContext) {
	this := p
	_ = this

	localctx = NewTwelveNoonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, parkingsignParserRULE_twelveNoon)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	case parkingsignParserNOON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(276)
			p.Match(parkingsignParserNOON)
		}


	case parkingsignParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.Match(parkingsignParserT__1)
		}
		{
			p.SetState(278)
			p.Match(parkingsignParserNOON)
		}




	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// ITwelveMidnightContext is an interface to support dynamic dispatch.
type ITwelveMidnightContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTwelveMidnightContext differentiates from other interfaces.
	IsTwelveMidnightContext()
}

type TwelveMidnightContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTwelveMidnightContext() *TwelveMidnightContext {
	var p = new(TwelveMidnightContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_twelveMidnight
	return p
}

func (*TwelveMidnightContext) IsTwelveMidnightContext() {}

func NewTwelveMidnightContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TwelveMidnightContext {
	var p = new(TwelveMidnightContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_twelveMidnight

	return p
}

func (s *TwelveMidnightContext) GetParser() antlr.Parser { return s.parser }

func (s *TwelveMidnightContext) MIDNIGHT() antlr.TerminalNode {
	return s.GetToken(parkingsignParserMIDNIGHT, 0)
}

func (s *TwelveMidnightContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TwelveMidnightContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *TwelveMidnightContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterTwelveMidnight(s)
	}
}

func (s *TwelveMidnightContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitTwelveMidnight(s)
	}
}




func (p *parkingsignParser) TwelveMidnight() (localctx ITwelveMidnightContext) {
	this := p
	_ = this

	localctx = NewTwelveMidnightContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, parkingsignParserRULE_twelveMidnight)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(284)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case parkingsignParserMIDNIGHT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(281)
			p.Match(parkingsignParserMIDNIGHT)
		}


	case parkingsignParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(282)
			p.Match(parkingsignParserT__1)
		}
		{
			p.SetState(283)
			p.Match(parkingsignParserMIDNIGHT)
		}




	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IAmContext is an interface to support dynamic dispatch.
type IAmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAmContext differentiates from other interfaces.
	IsAmContext()
}

type AmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAmContext() *AmContext {
	var p = new(AmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_am
	return p
}

func (*AmContext) IsAmContext() {}

func NewAmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AmContext {
	var p = new(AmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_am

	return p
}

func (s *AmContext) GetParser() antlr.Parser { return s.parser }
func (s *AmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *AmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterAm(s)
	}
}

func (s *AmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitAm(s)
	}
}




func (p *parkingsignParser) Am() (localctx IAmContext) {
	this := p
	_ = this

	localctx = NewAmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, parkingsignParserRULE_am)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
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
	case parkingsignParserT__2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(286)
			p.Match(parkingsignParserT__2)
		}


	case parkingsignParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(287)
			p.Match(parkingsignParserT__3)
		}




	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


// IPmContext is an interface to support dynamic dispatch.
type IPmContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPmContext differentiates from other interfaces.
	IsPmContext()
}

type PmContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPmContext() *PmContext {
	var p = new(PmContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = parkingsignParserRULE_pm
	return p
}

func (*PmContext) IsPmContext() {}

func NewPmContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PmContext {
	var p = new(PmContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = parkingsignParserRULE_pm

	return p
}

func (s *PmContext) GetParser() antlr.Parser { return s.parser }
func (s *PmContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PmContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}


func (s *PmContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.EnterPm(s)
	}
}

func (s *PmContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(parkingsignListener); ok {
		listenerT.ExitPm(s)
	}
}




func (p *parkingsignParser) Pm() (localctx IPmContext) {
	this := p
	_ = this

	localctx = NewPmContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, parkingsignParserRULE_pm)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(292)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case parkingsignParserT__4:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(290)
			p.Match(parkingsignParserT__4)
		}


	case parkingsignParserT__5:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(291)
			p.Match(parkingsignParserT__5)
		}




	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


