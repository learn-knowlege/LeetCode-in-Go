package problem0587

import (
	"sort"
	"testing"

	"github.com/aQuaYi/LeetCode-in-Go/kit"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	points [][]int
	ans    [][]int
}{

	{
		[][]int{
			{0, 0},
			{0, 4},
			{0, 19},
			{0, 22},
			{0, 25},
			{0, 26},
			{0, 32},
			{0, 39},
			{0, 40},
			{0, 49},
			{0, 54},
			{0, 55},
			{0, 63},
			{0, 67},
			{0, 72},
			{0, 85},
			{0, 89},
			{0, 96},
			{0, 99},
			{1, 9},
			{1, 13},
			{1, 21},
			{1, 23},
			{1, 25},
			{1, 28},
			{1, 31},
			{1, 32},
			{1, 35},
			{1, 41},
			{1, 50},
			{1, 56},
			{1, 59},
			{1, 61},
			{1, 63},
			{1, 65},
			{1, 66},
			{1, 68},
			{1, 89},
			{1, 97},
			{2, 1},
			{2, 6},
			{2, 7},
			{2, 16},
			{2, 26},
			{2, 38},
			{2, 41},
			{2, 43},
			{2, 46},
			{2, 52},
			{2, 56},
			{2, 59},
			{2, 68},
			{2, 87},
			{2, 88},
			{2, 91},
			{2, 92},
			{2, 95},
			{2, 99},
			{3, 7},
			{3, 25},
			{3, 27},
			{3, 28},
			{3, 29},
			{3, 32},
			{3, 40},
			{3, 45},
			{3, 46},
			{3, 48},
			{3, 51},
			{3, 52},
			{3, 63},
			{3, 64},
			{3, 66},
			{3, 70},
			{3, 71},
			{3, 75},
			{3, 81},
			{3, 88},
			{3, 97},
			{4, 2},
			{4, 3},
			{4, 7},
			{4, 24},
			{4, 28},
			{4, 30},
			{4, 34},
			{4, 43},
			{4, 45},
			{4, 46},
			{4, 56},
			{4, 59},
			{4, 61},
			{4, 63},
			{4, 70},
			{4, 71},
			{4, 78},
			{4, 83},
			{4, 91},
			{4, 92},
			{4, 93},
			{4, 95},
			{4, 98},
			{5, 2},
			{5, 3},
			{5, 21},
			{5, 26},
			{5, 36},
			{5, 44},
			{5, 49},
			{5, 54},
			{5, 64},
			{5, 72},
			{5, 73},
			{5, 80},
			{5, 86},
			{5, 87},
			{5, 91},
			{5, 98},
			{6, 0},
			{6, 8},
			{6, 20},
			{6, 22},
			{6, 23},
			{6, 29},
			{6, 37},
			{6, 46},
			{6, 50},
			{6, 57},
			{6, 58},
			{6, 59},
			{6, 65},
			{6, 68},
			{6, 70},
			{6, 73},
			{6, 77},
			{6, 79},
			{6, 82},
			{6, 84},
			{6, 87},
			{6, 98},
			{7, 0},
			{7, 6},
			{7, 8},
			{7, 13},
			{7, 15},
			{7, 18},
			{7, 19},
			{7, 28},
			{7, 35},
			{7, 36},
			{7, 38},
			{7, 44},
			{7, 46},
			{7, 59},
			{7, 60},
			{7, 61},
			{7, 67},
			{7, 89},
			{7, 92},
			{7, 96},
			{7, 99},
			{8, 0},
			{8, 1},
			{8, 4},
			{8, 9},
			{8, 15},
			{8, 28},
			{8, 35},
			{8, 45},
			{8, 54},
			{8, 57},
			{8, 58},
			{8, 59},
			{8, 72},
			{8, 73},
			{8, 81},
			{8, 83},
			{8, 84},
			{8, 85},
			{8, 89},
			{8, 93},
			{8, 96},
			{9, 4},
			{9, 8},
			{9, 15},
			{9, 24},
			{9, 30},
			{9, 36},
			{9, 71},
			{9, 75},
			{9, 76},
			{9, 77},
			{9, 79},
			{9, 80},
			{9, 81},
			{9, 84},
			{9, 85},
			{9, 86},
			{9, 87},
			{9, 89},
			{9, 94},
			{10, 0},
			{10, 6},
			{10, 7},
			{10, 13},
			{10, 16},
			{10, 17},
			{10, 19},
			{10, 22},
			{10, 53},
			{10, 54},
			{10, 56},
			{10, 58},
			{10, 64},
			{10, 66},
			{10, 67},
			{10, 74},
			{10, 84},
			{10, 87},
			{10, 91},
			{10, 92},
			{11, 0},
			{11, 3},
			{11, 4},
			{11, 7},
			{11, 18},
			{11, 25},
			{11, 27},
			{11, 28},
			{11, 30},
			{11, 32},
			{11, 41},
			{11, 44},
			{11, 52},
			{11, 57},
			{11, 62},
			{11, 63},
			{11, 66},
			{11, 72},
			{11, 78},
			{11, 80},
			{11, 96},
			{12, 10},
			{12, 12},
			{12, 22},
			{12, 25},
			{12, 29},
			{12, 34},
			{12, 38},
			{12, 45},
			{12, 55},
			{12, 57},
			{12, 64},
			{12, 65},
			{12, 68},
			{12, 71},
			{12, 79},
			{12, 82},
			{12, 91},
			{12, 98},
			{13, 0},
			{13, 7},
			{13, 8},
			{13, 10},
			{13, 13},
			{13, 15},
			{13, 33},
			{13, 34},
			{13, 36},
			{13, 38},
			{13, 40},
			{13, 42},
			{13, 63},
			{13, 66},
			{13, 70},
			{13, 73},
			{13, 81},
			{13, 84},
			{13, 88},
			{14, 10},
			{14, 15},
			{14, 21},
			{14, 22},
			{14, 24},
			{14, 41},
			{14, 46},
			{14, 47},
			{14, 52},
			{14, 55},
			{14, 61},
			{14, 67},
			{14, 73},
			{14, 88},
			{14, 92},
			{14, 97},
			{15, 0},
			{15, 5},
			{15, 12},
			{15, 32},
			{15, 45},
			{15, 48},
			{15, 49},
			{15, 51},
			{15, 54},
			{15, 55},
			{15, 62},
			{15, 64},
			{15, 65},
			{15, 66},
			{15, 74},
			{15, 76},
			{15, 84},
			{15, 97},
			{16, 8},
			{16, 9},
			{16, 14},
			{16, 16},
			{16, 18},
			{16, 23},
			{16, 28},
			{16, 38},
			{16, 41},
			{16, 43},
			{16, 48},
			{16, 49},
			{16, 53},
			{16, 61},
			{16, 63},
			{16, 64},
			{16, 70},
			{16, 73},
			{16, 84},
			{16, 88},
			{16, 92},
			{16, 93},
			{16, 97},
			{17, 0},
			{17, 1},
			{17, 15},
			{17, 21},
			{17, 27},
			{17, 30},
			{17, 34},
			{17, 39},
			{17, 42},
			{17, 43},
			{17, 50},
			{17, 52},
			{17, 57},
			{17, 58},
			{17, 62},
			{17, 63},
			{17, 64},
			{17, 65},
			{17, 66},
			{17, 77},
			{17, 79},
			{17, 81},
			{17, 82},
			{17, 83},
			{17, 86},
			{17, 90},
			{17, 93},
			{17, 96},
			{17, 98},
			{18, 0},
			{18, 4},
			{18, 5},
			{18, 8},
			{18, 14},
			{18, 16},
			{18, 17},
			{18, 21},
			{18, 29},
			{18, 34},
			{18, 37},
			{18, 46},
			{18, 47},
			{18, 63},
			{18, 65},
			{18, 67},
			{18, 72},
			{18, 73},
			{18, 77},
			{18, 79},
			{18, 82},
			{18, 84},
			{18, 85},
			{18, 91},
			{18, 98},
			{19, 3},
			{19, 4},
			{19, 8},
			{19, 13},
			{19, 17},
			{19, 26},
			{19, 36},
			{19, 38},
			{19, 48},
			{19, 49},
			{19, 52},
			{19, 57},
			{19, 61},
			{19, 64},
			{19, 69},
			{19, 71},
			{19, 76},
			{19, 80},
			{19, 82},
			{19, 83},
			{19, 84},
			{19, 86},
			{19, 94},
			{20, 4},
			{20, 6},
			{20, 13},
			{20, 23},
			{20, 36},
			{20, 45},
			{20, 46},
			{20, 47},
			{20, 55},
			{20, 57},
			{20, 62},
			{20, 65},
			{20, 67},
			{20, 71},
			{20, 77},
			{20, 78},
			{20, 82},
			{20, 86},
			{20, 87},
			{20, 89},
			{20, 94},
			{20, 97},
			{21, 2},
			{21, 12},
			{21, 16},
			{21, 20},
			{21, 27},
			{21, 30},
			{21, 32},
			{21, 34},
			{21, 52},
			{21, 63},
			{21, 65},
			{21, 69},
			{21, 75},
			{21, 88},
			{21, 91},
			{21, 93},
			{21, 94},
			{21, 95},
			{22, 2},
			{22, 4},
			{22, 6},
			{22, 7},
			{22, 10},
			{22, 12},
			{22, 38},
			{22, 42},
			{22, 49},
			{22, 50},
			{22, 70},
			{22, 72},
			{22, 76},
			{22, 79},
			{22, 80},
			{22, 88},
			{22, 89},
			{22, 90},
			{23, 9},
			{23, 11},
			{23, 21},
			{23, 23},
			{23, 24},
			{23, 35},
			{23, 38},
			{23, 41},
			{23, 43},
			{23, 48},
			{23, 51},
			{23, 61},
			{23, 62},
			{23, 68},
			{23, 70},
			{23, 71},
			{23, 72},
			{23, 84},
			{23, 85},
			{23, 98},
			{23, 99},
			{24, 1},
			{24, 2},
			{24, 6},
			{24, 19},
			{24, 41},
			{24, 49},
			{24, 51},
			{24, 68},
			{24, 72},
			{24, 77},
			{24, 79},
			{24, 82},
			{24, 83},
			{24, 84},
			{24, 99},
			{25, 8},
			{25, 15},
			{25, 24},
			{25, 28},
			{25, 41},
			{25, 46},
			{25, 47},
			{25, 48},
			{25, 60},
			{25, 61},
			{25, 62},
			{25, 65},
			{25, 69},
			{25, 72},
			{25, 76},
			{25, 78},
			{25, 80},
			{25, 98},
			{26, 11},
			{26, 31},
			{26, 32},
			{26, 35},
			{26, 42},
			{26, 51},
			{26, 56},
			{26, 73},
			{26, 74},
			{26, 75},
			{26, 80},
			{26, 82},
			{26, 90},
			{26, 99},
			{27, 0},
			{27, 1},
			{27, 9},
			{27, 14},
			{27, 22},
			{27, 23},
			{27, 26},
			{27, 30},
			{27, 43},
			{27, 47},
			{27, 58},
			{27, 60},
			{27, 61},
			{27, 67},
			{27, 70},
			{27, 79},
			{27, 83},
			{27, 85},
			{27, 89},
			{27, 91},
			{27, 94},
			{28, 3},
			{28, 22},
			{28, 26},
			{28, 31},
			{28, 32},
			{28, 43},
			{28, 48},
			{28, 59},
			{28, 62},
			{28, 64},
			{28, 67},
			{28, 72},
			{28, 89},
			{28, 94},
			{28, 97},
			{29, 5},
			{29, 8},
			{29, 22},
			{29, 23},
			{29, 29},
			{29, 31},
			{29, 38},
			{29, 48},
			{29, 58},
			{29, 72},
			{29, 73},
			{29, 80},
			{29, 82},
			{29, 83},
			{29, 84},
			{29, 87},
			{29, 88},
			{29, 92},
			{29, 95},
			{30, 0},
			{30, 6},
			{30, 7},
			{30, 10},
			{30, 12},
			{30, 23},
			{30, 38},
			{30, 44},
			{30, 49},
			{30, 50},
			{30, 62},
			{30, 69},
			{30, 70},
			{30, 77},
			{30, 78},
			{30, 81},
			{30, 87},
			{30, 89},
			{30, 93},
			{30, 97},
			{31, 3},
			{31, 4},
			{31, 20},
			{31, 22},
			{31, 23},
			{31, 28},
			{31, 29},
			{31, 38},
			{31, 40},
			{31, 50},
			{31, 58},
			{31, 67},
			{31, 70},
			{31, 73},
			{31, 81},
			{31, 89},
			{31, 90},
			{32, 1},
			{32, 10},
			{32, 12},
			{32, 17},
			{32, 18},
			{32, 29},
			{32, 39},
			{32, 41},
			{32, 43},
			{32, 44},
			{32, 54},
			{32, 55},
			{32, 56},
			{32, 58},
			{32, 66},
			{32, 68},
			{32, 70},
			{32, 73},
			{32, 82},
			{32, 84},
			{32, 93},
			{32, 96},
			{32, 97},
			{33, 4},
			{33, 5},
			{33, 6},
			{33, 8},
			{33, 13},
			{33, 26},
			{33, 31},
			{33, 47},
			{33, 51},
			{33, 59},
			{33, 60},
			{33, 66},
			{33, 67},
			{33, 71},
			{33, 86},
			{33, 94},
			{33, 95},
			{33, 98},
			{34, 11},
			{34, 13},
			{34, 29},
			{34, 36},
			{34, 38},
			{34, 47},
			{34, 50},
			{34, 59},
			{34, 62},
			{34, 63},
			{34, 67},
			{34, 73},
			{34, 77},
			{34, 86},
			{34, 88},
			{34, 94},
			{34, 96},
			{34, 97},
			{34, 98},
			{35, 16},
			{35, 17},
			{35, 28},
			{35, 37},
			{35, 38},
			{35, 39},
			{35, 42},
			{35, 45},
			{35, 53},
			{35, 55},
			{35, 64},
			{35, 68},
			{35, 71},
			{36, 5},
			{36, 14},
			{36, 22},
			{36, 28},
			{36, 29},
			{36, 31},
			{36, 35},
			{36, 45},
			{36, 50},
			{36, 61},
			{36, 70},
			{36, 71},
			{36, 77},
			{36, 84},
			{36, 85},
			{36, 86},
			{36, 96},
			{36, 99},
			{37, 7},
			{37, 21},
			{37, 23},
			{37, 45},
			{37, 50},
			{37, 52},
			{37, 53},
			{37, 57},
			{37, 59},
			{37, 63},
			{37, 65},
			{37, 68},
			{37, 72},
			{37, 73},
			{37, 74},
			{37, 80},
			{37, 84},
			{37, 96},
			{37, 98},
			{38, 1},
			{38, 2},
			{38, 5},
			{38, 10},
			{38, 11},
			{38, 19},
			{38, 36},
			{38, 45},
			{38, 48},
			{38, 49},
			{38, 51},
			{38, 52},
			{38, 55},
			{38, 56},
			{38, 64},
			{38, 78},
			{38, 80},
			{38, 82},
			{38, 88},
			{38, 93},
			{38, 95},
			{38, 96},
			{39, 0},
			{39, 1},
			{39, 5},
			{39, 15},
			{39, 17},
			{39, 22},
			{39, 23},
			{39, 30},
			{39, 42},
			{39, 48},
			{39, 54},
			{39, 65},
			{39, 68},
			{39, 72},
			{39, 90},
			{39, 92},
			{39, 94},
			{39, 97},
			{40, 5},
			{40, 8},
			{40, 10},
			{40, 13},
			{40, 19},
			{40, 26},
			{40, 27},
			{40, 42},
			{40, 45},
			{40, 55},
			{40, 56},
			{40, 62},
			{40, 63},
			{40, 77},
			{40, 83},
			{40, 85},
			{40, 88},
			{40, 92},
			{40, 93},
			{41, 3},
			{41, 8},
			{41, 11},
			{41, 15},
			{41, 24},
			{41, 28},
			{41, 29},
			{41, 40},
			{41, 48},
			{41, 55},
			{41, 59},
			{41, 63},
			{41, 79},
			{41, 80},
			{41, 87},
			{41, 91},
			{41, 99},
			{42, 2},
			{42, 12},
			{42, 14},
			{42, 19},
			{42, 25},
			{42, 42},
			{42, 47},
			{42, 52},
			{42, 62},
			{42, 64},
			{42, 73},
			{42, 74},
			{42, 90},
			{42, 91},
			{42, 95},
			{42, 96},
			{42, 97},
			{43, 2},
			{43, 16},
			{43, 17},
			{43, 21},
			{43, 33},
			{43, 39},
			{43, 40},
			{43, 43},
			{43, 49},
			{43, 57},
			{43, 70},
			{43, 77},
			{43, 79},
			{43, 86},
			{43, 89},
			{43, 92},
			{43, 96},
			{44, 6},
			{44, 8},
			{44, 10},
			{44, 27},
			{44, 30},
			{44, 34},
			{44, 35},
			{44, 39},
			{44, 41},
			{44, 45},
			{44, 46},
			{44, 48},
			{44, 54},
			{44, 56},
			{44, 65},
			{44, 69},
			{44, 71},
			{44, 72},
			{44, 74},
			{44, 80},
			{44, 85},
			{44, 86},
			{45, 1},
			{45, 3},
			{45, 5},
			{45, 21},
			{45, 22},
			{45, 34},
			{45, 36},
			{45, 46},
			{45, 47},
			{45, 48},
			{45, 55},
			{45, 60},
			{45, 61},
			{45, 63},
			{45, 64},
			{45, 68},
			{45, 72},
			{45, 77},
			{45, 78},
			{45, 82},
			{45, 84},
			{45, 88},
			{45, 90},
			{45, 92},
			{45, 94},
			{46, 13},
			{46, 15},
			{46, 19},
			{46, 20},
			{46, 25},
			{46, 32},
			{46, 37},
			{46, 40},
			{46, 41},
			{46, 46},
			{46, 50},
			{46, 51},
			{46, 55},
			{46, 59},
			{46, 69},
			{46, 74},
			{46, 77},
			{46, 86},
			{46, 97},
			{46, 98},
			{47, 2},
			{47, 3},
			{47, 5},
			{47, 7},
			{47, 8},
			{47, 12},
			{47, 14},
			{47, 23},
			{47, 26},
			{47, 28},
			{47, 33},
			{47, 45},
			{47, 46},
			{47, 49},
			{47, 59},
			{47, 62},
			{47, 73},
			{47, 76},
			{47, 79},
			{47, 89},
			{47, 95},
			{47, 96},
			{47, 99},
			{48, 2},
			{48, 6},
			{48, 9},
			{48, 10},
			{48, 12},
			{48, 16},
			{48, 17},
			{48, 21},
			{48, 24},
			{48, 26},
			{48, 28},
			{48, 30},
			{48, 32},
			{48, 35},
			{48, 40},
			{48, 42},
			{48, 44},
			{48, 48},
			{48, 53},
			{48, 65},
			{48, 67},
			{48, 70},
			{48, 72},
			{48, 79},
			{48, 87},
			{48, 91},
			{48, 94},
			{48, 99},
			{49, 0},
			{49, 2},
			{49, 13},
			{49, 15},
			{49, 19},
			{49, 25},
			{49, 28},
			{49, 44},
			{49, 49},
			{49, 52},
			{49, 65},
			{49, 68},
			{49, 69},
			{50, 6},
			{50, 20},
			{50, 21},
			{50, 31},
			{50, 37},
			{50, 39},
			{50, 52},
			{50, 60},
			{50, 61},
			{50, 63},
			{50, 66},
			{50, 74},
			{50, 85},
			{50, 93},
			{50, 97},
			{50, 99},
			{51, 1},
			{51, 11},
			{51, 14},
			{51, 22},
			{51, 24},
			{51, 25},
			{51, 27},
			{51, 30},
			{51, 38},
			{51, 46},
			{51, 49},
			{51, 50},
			{51, 54},
			{51, 60},
			{51, 61},
			{51, 62},
			{51, 75},
			{51, 85},
			{51, 86},
			{51, 88},
			{51, 90},
			{51, 91},
			{51, 93},
			{51, 95},
			{51, 96},
			{51, 97},
			{52, 2},
			{52, 3},
			{52, 7},
			{52, 14},
			{52, 16},
			{52, 19},
			{52, 21},
			{52, 22},
			{52, 27},
			{52, 28},
			{52, 29},
			{52, 30},
			{52, 32},
			{52, 39},
			{52, 42},
			{52, 49},
			{52, 54},
			{52, 55},
			{52, 61},
			{52, 65},
			{52, 66},
			{52, 68},
			{52, 71},
			{52, 72},
			{52, 76},
			{52, 78},
			{52, 80},
			{52, 85},
			{52, 86},
			{52, 90},
			{52, 99},
			{53, 2},
			{53, 5},
			{53, 10},
			{53, 14},
			{53, 25},
			{53, 27},
			{53, 30},
			{53, 41},
			{53, 44},
			{53, 49},
			{53, 73},
			{53, 74},
			{53, 78},
			{53, 80},
			{53, 86},
			{53, 88},
			{53, 90},
			{53, 97},
			{54, 3},
			{54, 9},
			{54, 12},
			{54, 14},
			{54, 20},
			{54, 23},
			{54, 25},
			{54, 38},
			{54, 41},
			{54, 42},
			{54, 43},
			{54, 45},
			{54, 48},
			{54, 52},
			{54, 54},
			{54, 57},
			{54, 65},
			{54, 67},
			{54, 70},
			{54, 83},
			{54, 90},
			{54, 97},
			{54, 98},
			{55, 0},
			{55, 1},
			{55, 8},
			{55, 9},
			{55, 19},
			{55, 28},
			{55, 29},
			{55, 32},
			{55, 33},
			{55, 42},
			{55, 44},
			{55, 48},
			{55, 58},
			{55, 72},
			{55, 74},
			{55, 88},
			{55, 89},
			{55, 92},
			{55, 94},
			{55, 97},
			{56, 1},
			{56, 2},
			{56, 6},
			{56, 7},
			{56, 9},
			{56, 13},
			{56, 14},
			{56, 15},
			{56, 18},
			{56, 23},
			{56, 26},
			{56, 37},
			{56, 44},
			{56, 45},
			{56, 46},
			{56, 48},
			{56, 50},
			{56, 54},
			{56, 56},
			{56, 58},
			{56, 59},
			{56, 78},
			{56, 83},
			{56, 86},
			{56, 96},
			{57, 6},
			{57, 9},
			{57, 10},
			{57, 13},
			{57, 25},
			{57, 27},
			{57, 36},
			{57, 39},
			{57, 40},
			{57, 41},
			{57, 53},
			{57, 69},
			{57, 75},
			{57, 81},
			{57, 87},
			{57, 89},
			{57, 94},
			{57, 97},
			{57, 99},
			{58, 2},
			{58, 7},
			{58, 8},
			{58, 12},
			{58, 18},
			{58, 19},
			{58, 22},
			{58, 24},
			{58, 25},
			{58, 47},
			{58, 59},
			{58, 60},
			{58, 78},
			{58, 85},
			{58, 92},
			{58, 94},
			{59, 4},
			{59, 6},
			{59, 8},
			{59, 13},
			{59, 15},
			{59, 16},
			{59, 17},
			{59, 18},
			{59, 25},
			{59, 31},
			{59, 32},
			{59, 34},
			{59, 39},
			{59, 43},
			{59, 63},
			{59, 68},
			{59, 72},
			{59, 85},
			{59, 87},
			{59, 88},
			{59, 99},
			{60, 1},
			{60, 3},
			{60, 11},
			{60, 14},
			{60, 24},
			{60, 41},
			{60, 43},
			{60, 45},
			{60, 52},
			{60, 57},
			{60, 58},
			{60, 60},
			{60, 61},
			{60, 65},
			{60, 69},
			{60, 72},
			{60, 74},
			{60, 79},
			{60, 80},
			{60, 82},
			{60, 90},
			{60, 97},
			{60, 99},
			{61, 11},
			{61, 18},
			{61, 21},
			{61, 31},
			{61, 39},
			{61, 41},
			{61, 47},
			{61, 48},
			{61, 55},
			{61, 57},
			{61, 65},
			{61, 70},
			{61, 74},
			{61, 77},
			{61, 82},
			{61, 87},
			{61, 89},
			{61, 95},
			{61, 98},
			{62, 6},
			{62, 17},
			{62, 20},
			{62, 23},
			{62, 25},
			{62, 36},
			{62, 38},
			{62, 46},
			{62, 51},
			{62, 55},
			{62, 57},
			{62, 60},
			{62, 70},
			{62, 73},
			{62, 76},
			{62, 85},
			{62, 90},
			{62, 97},
			{63, 2},
			{63, 7},
			{63, 11},
			{63, 20},
			{63, 21},
			{63, 23},
			{63, 38},
			{63, 47},
			{63, 54},
			{63, 59},
			{63, 62},
			{63, 65},
			{63, 70},
			{63, 81},
			{63, 87},
			{63, 93},
			{64, 0},
			{64, 3},
			{64, 11},
			{64, 15},
			{64, 18},
			{64, 25},
			{64, 26},
			{64, 34},
			{64, 48},
			{64, 49},
			{64, 50},
			{64, 51},
			{64, 53},
			{64, 54},
			{64, 55},
			{64, 56},
			{64, 58},
			{64, 59},
			{64, 65},
			{64, 75},
			{64, 78},
			{64, 79},
			{64, 83},
			{64, 93},
			{65, 4},
			{65, 11},
			{65, 15},
			{65, 24},
			{65, 26},
			{65, 43},
			{65, 52},
			{65, 61},
			{65, 64},
			{65, 81},
			{65, 84},
			{65, 86},
			{65, 92},
			{66, 2},
			{66, 7},
			{66, 13},
			{66, 31},
			{66, 34},
			{66, 35},
			{66, 42},
			{66, 44},
			{66, 46},
			{66, 48},
			{66, 49},
			{66, 58},
			{66, 59},
			{66, 64},
			{66, 79},
			{66, 88},
			{66, 90},
			{66, 92},
			{66, 96},
			{66, 98},
			{66, 99},
			{67, 1},
			{67, 3},
			{67, 4},
			{67, 9},
			{67, 11},
			{67, 14},
			{67, 18},
			{67, 20},
			{67, 21},
			{67, 24},
			{67, 27},
			{67, 29},
			{67, 40},
			{67, 41},
			{67, 47},
			{67, 53},
			{67, 55},
			{67, 58},
			{67, 59},
			{67, 67},
			{67, 72},
			{67, 75},
			{67, 76},
			{67, 81},
			{67, 84},
			{67, 93},
			{67, 95},
			{67, 97},
			{67, 98},
			{68, 3},
			{68, 12},
			{68, 18},
			{68, 25},
			{68, 27},
			{68, 33},
			{68, 34},
			{68, 35},
			{68, 40},
			{68, 42},
			{68, 43},
			{68, 50},
			{68, 51},
			{68, 53},
			{68, 60},
			{68, 74},
			{68, 81},
			{68, 82},
			{68, 96},
			{68, 98},
			{69, 3},
			{69, 8},
			{69, 20},
			{69, 24},
			{69, 28},
			{69, 30},
			{69, 35},
			{69, 44},
			{69, 59},
			{69, 60},
			{69, 67},
			{69, 75},
			{69, 76},
			{69, 85},
			{69, 86},
			{69, 88},
			{69, 95},
			{70, 4},
			{70, 6},
			{70, 16},
			{70, 18},
			{70, 20},
			{70, 22},
			{70, 24},
			{70, 29},
			{70, 31},
			{70, 36},
			{70, 37},
			{70, 39},
			{70, 53},
			{70, 54},
			{70, 55},
			{70, 63},
			{70, 71},
			{70, 83},
			{70, 85},
			{70, 90},
			{70, 91},
			{70, 94},
			{70, 96},
			{71, 14},
			{71, 18},
			{71, 25},
			{71, 29},
			{71, 30},
			{71, 32},
			{71, 37},
			{71, 38},
			{71, 41},
			{71, 43},
			{71, 46},
			{71, 50},
			{71, 51},
			{71, 54},
			{71, 70},
			{71, 78},
			{71, 86},
			{71, 88},
			{71, 90},
			{71, 91},
			{71, 94},
			{72, 3},
			{72, 13},
			{72, 21},
			{72, 34},
			{72, 43},
			{72, 45},
			{72, 48},
			{72, 52},
			{72, 55},
			{72, 58},
			{72, 61},
			{72, 71},
			{72, 77},
			{72, 78},
			{72, 92},
			{72, 93},
			{72, 94},
			{72, 96},
			{73, 1},
			{73, 5},
			{73, 22},
			{73, 25},
			{73, 50},
			{73, 63},
			{73, 69},
			{73, 77},
			{73, 79},
			{73, 82},
			{73, 86},
			{73, 88},
			{73, 94},
			{74, 1},
			{74, 5},
			{74, 8},
			{74, 11},
			{74, 17},
			{74, 19},
			{74, 20},
			{74, 25},
			{74, 32},
			{74, 36},
			{74, 43},
			{74, 44},
			{74, 50},
			{74, 54},
			{74, 55},
			{74, 58},
			{74, 72},
			{74, 74},
			{74, 75},
			{74, 78},
			{74, 87},
			{74, 91},
			{74, 98},
			{74, 99},
			{75, 4},
			{75, 14},
			{75, 32},
			{75, 33},
			{75, 36},
			{75, 43},
			{75, 45},
			{75, 47},
			{75, 60},
			{75, 61},
			{75, 62},
			{75, 66},
			{75, 78},
			{75, 79},
			{75, 82},
			{75, 92},
			{75, 93},
			{75, 99},
			{76, 1},
			{76, 10},
			{76, 19},
			{76, 20},
			{76, 37},
			{76, 38},
			{76, 43},
			{76, 44},
			{76, 48},
			{76, 51},
			{76, 52},
			{76, 56},
			{76, 61},
			{76, 69},
			{76, 71},
			{76, 80},
			{76, 84},
			{76, 89},
			{76, 92},
			{76, 93},
			{76, 94},
			{76, 97},
			{77, 12},
			{77, 21},
			{77, 23},
			{77, 24},
			{77, 25},
			{77, 31},
			{77, 33},
			{77, 40},
			{77, 41},
			{77, 45},
			{77, 47},
			{77, 55},
			{77, 57},
			{77, 58},
			{77, 61},
			{77, 84},
			{77, 95},
			{78, 3},
			{78, 5},
			{78, 7},
			{78, 10},
			{78, 11},
			{78, 15},
			{78, 16},
			{78, 17},
			{78, 55},
			{78, 59},
			{78, 70},
			{78, 71},
			{78, 72},
			{78, 75},
			{78, 81},
			{78, 86},
			{78, 88},
			{78, 91},
			{78, 94},
			{78, 99},
			{79, 1},
			{79, 9},
			{79, 13},
			{79, 20},
			{79, 23},
			{79, 34},
			{79, 36},
			{79, 37},
			{79, 44},
			{79, 45},
			{79, 46},
			{79, 48},
			{79, 54},
			{79, 57},
			{79, 59},
			{79, 60},
			{79, 92},
			{79, 95},
			{79, 97},
			{79, 98},
			{80, 3},
			{80, 17},
			{80, 20},
			{80, 21},
			{80, 23},
			{80, 26},
			{80, 34},
			{80, 38},
			{80, 39},
			{80, 44},
			{80, 48},
			{80, 55},
			{80, 63},
			{80, 66},
			{80, 67},
			{80, 71},
			{80, 78},
			{80, 84},
			{80, 93},
			{81, 2},
			{81, 5},
			{81, 27},
			{81, 28},
			{81, 30},
			{81, 37},
			{81, 42},
			{81, 44},
			{81, 46},
			{81, 47},
			{81, 48},
			{81, 50},
			{81, 55},
			{81, 57},
			{81, 61},
			{81, 67},
			{81, 68},
			{81, 77},
			{81, 78},
			{81, 85},
			{81, 86},
			{81, 90},
			{81, 96},
			{81, 98},
			{82, 6},
			{82, 9},
			{82, 12},
			{82, 15},
			{82, 19},
			{82, 31},
			{82, 35},
			{82, 41},
			{82, 46},
			{82, 51},
			{82, 54},
			{82, 58},
			{82, 69},
			{82, 78},
			{82, 80},
			{82, 82},
			{82, 89},
			{82, 94},
			{82, 96},
			{82, 97},
			{83, 11},
			{83, 14},
			{83, 26},
			{83, 30},
			{83, 32},
			{83, 36},
			{83, 37},
			{83, 44},
			{83, 47},
			{83, 59},
			{83, 60},
			{83, 64},
			{83, 77},
			{83, 79},
			{83, 93},
			{84, 1},
			{84, 5},
			{84, 11},
			{84, 19},
			{84, 20},
			{84, 22},
			{84, 27},
			{84, 33},
			{84, 38},
			{84, 42},
			{84, 44},
			{84, 48},
			{84, 54},
			{84, 67},
			{84, 72},
			{84, 73},
			{84, 75},
			{84, 76},
			{84, 84},
			{84, 89},
			{84, 95},
			{84, 99},
			{85, 12},
			{85, 15},
			{85, 16},
			{85, 20},
			{85, 21},
			{85, 23},
			{85, 24},
			{85, 35},
			{85, 37},
			{85, 44},
			{85, 54},
			{85, 60},
			{85, 62},
			{85, 63},
			{85, 64},
			{85, 66},
			{85, 67},
			{85, 68},
			{85, 69},
			{85, 71},
			{85, 81},
			{85, 88},
			{86, 0},
			{86, 3},
			{86, 5},
			{86, 6},
			{86, 7},
			{86, 8},
			{86, 20},
			{86, 32},
			{86, 34},
			{86, 35},
			{86, 38},
			{86, 39},
			{86, 42},
			{86, 43},
			{86, 44},
			{86, 62},
			{86, 67},
			{86, 74},
			{86, 79},
			{86, 83},
			{86, 94},
			{86, 98},
			{87, 8},
			{87, 12},
			{87, 19},
			{87, 26},
			{87, 28},
			{87, 31},
			{87, 39},
			{87, 41},
			{87, 43},
			{87, 47},
			{87, 48},
			{87, 50},
			{87, 51},
			{87, 55},
			{87, 57},
			{87, 63},
			{87, 67},
			{87, 68},
			{87, 69},
			{87, 72},
			{87, 81},
			{87, 83},
			{87, 92},
			{88, 0},
			{88, 5},
			{88, 6},
			{88, 14},
			{88, 17},
			{88, 21},
			{88, 28},
			{88, 29},
			{88, 30},
			{88, 31},
			{88, 50},
			{88, 51},
			{88, 54},
			{88, 58},
			{88, 63},
			{88, 65},
			{88, 75},
			{88, 80},
			{88, 83},
			{88, 86},
			{88, 88},
			{88, 92},
			{88, 93},
			{89, 1},
			{89, 2},
			{89, 7},
			{89, 11},
			{89, 15},
			{89, 16},
			{89, 24},
			{89, 27},
			{89, 33},
			{89, 36},
			{89, 39},
			{89, 40},
			{89, 41},
			{89, 42},
			{89, 44},
			{89, 45},
			{89, 46},
			{89, 50},
			{89, 53},
			{89, 54},
			{89, 57},
			{89, 58},
			{89, 60},
			{89, 61},
			{89, 65},
			{89, 84},
			{89, 85},
			{89, 86},
			{89, 93},
			{89, 97},
			{89, 98},
			{90, 5},
			{90, 6},
			{90, 8},
			{90, 9},
			{90, 10},
			{90, 11},
			{90, 13},
			{90, 17},
			{90, 19},
			{90, 26},
			{90, 29},
			{90, 30},
			{90, 34},
			{90, 44},
			{90, 50},
			{90, 61},
			{90, 63},
			{90, 64},
			{90, 75},
			{90, 78},
			{90, 85},
			{90, 88},
			{90, 89},
			{90, 92},
			{90, 94},
			{90, 96},
			{90, 98},
			{91, 0},
			{91, 17},
			{91, 23},
			{91, 24},
			{91, 34},
			{91, 40},
			{91, 44},
			{91, 45},
			{91, 49},
			{91, 52},
			{91, 57},
			{91, 58},
			{91, 61},
			{91, 73},
			{91, 96},
			{92, 0},
			{92, 16},
			{92, 22},
			{92, 29},
			{92, 36},
			{92, 39},
			{92, 42},
			{92, 44},
			{92, 46},
			{92, 54},
			{92, 60},
			{92, 65},
			{92, 73},
			{92, 78},
			{92, 86},
			{92, 88},
			{92, 91},
			{92, 94},
			{93, 11},
			{93, 12},
			{93, 14},
			{93, 16},
			{93, 19},
			{93, 38},
			{93, 43},
			{93, 48},
			{93, 53},
			{93, 57},
			{93, 61},
			{93, 62},
			{93, 64},
			{93, 78},
			{93, 89},
			{94, 0},
			{94, 1},
			{94, 2},
			{94, 3},
			{94, 9},
			{94, 12},
			{94, 22},
			{94, 27},
			{94, 35},
			{94, 42},
			{94, 62},
			{94, 68},
			{94, 73},
			{94, 75},
			{94, 83},
			{94, 85},
			{94, 87},
			{94, 88},
			{94, 92},
			{94, 93},
			{94, 96},
			{95, 0},
			{95, 15},
			{95, 17},
			{95, 20},
			{95, 28},
			{95, 36},
			{95, 38},
			{95, 41},
			{95, 45},
			{95, 47},
			{95, 53},
			{95, 56},
			{95, 79},
			{95, 82},
			{95, 84},
			{95, 89},
			{95, 97},
			{96, 14},
			{96, 16},
			{96, 22},
			{96, 23},
			{96, 39},
			{96, 44},
			{96, 45},
			{96, 46},
			{96, 59},
			{96, 61},
			{96, 67},
			{96, 68},
			{96, 69},
			{96, 70},
			{96, 90},
			{96, 92},
			{96, 95},
			{96, 99},
			{97, 8},
			{97, 13},
			{97, 17},
			{97, 29},
			{97, 42},
			{97, 43},
			{97, 44},
			{97, 47},
			{97, 50},
			{97, 55},
			{97, 58},
			{97, 67},
			{97, 71},
			{97, 78},
			{97, 80},
			{97, 87},
			{97, 92},
			{97, 94},
			{97, 96},
			{98, 3},
			{98, 4},
			{98, 5},
			{98, 8},
			{98, 9},
			{98, 17},
			{98, 19},
			{98, 22},
			{98, 26},
			{98, 30},
			{98, 32},
			{98, 37},
			{98, 38},
			{98, 39},
			{98, 46},
			{98, 53},
			{98, 56},
			{98, 59},
			{98, 66},
			{98, 77},
			{98, 85},
			{98, 89},
			{98, 96},
			{98, 97},
			{99, 8},
			{99, 9},
			{99, 10},
			{99, 13},
			{99, 15},
			{99, 21},
			{99, 24},
			{99, 29},
			{99, 35},
			{99, 61},
			{99, 73},
			{99, 81},
			{99, 84},
			{99, 89},
			{99, 90},
			{99, 91},
			{99, 97},
		},
		[][]int{
			{0, 0},
			{6, 0},
			{7, 0},
			{8, 0},
			{10, 0},
			{11, 0},
			{13, 0},
			{15, 0},
			{17, 0},
			{18, 0},
			{27, 0},
			{30, 0},
			{39, 0},
			{49, 0},
			{55, 0},
			{64, 0},
			{86, 0},
			{88, 0},
			{91, 0},
			{92, 0},
			{94, 0},
			{95, 0},
			{98, 3},
			{99, 8},
			{99, 9},
			{99, 10},
			{99, 13},
			{99, 15},
			{99, 21},
			{99, 24},
			{99, 29},
			{99, 35},
			{99, 61},
			{99, 73},
			{99, 81},
			{99, 84},
			{99, 89},
			{99, 90},
			{99, 91},
			{99, 97},
			{96, 99},
			{84, 99},
			{78, 99},
			{75, 99},
			{74, 99},
			{66, 99},
			{60, 99},
			{59, 99},
			{57, 99},
			{52, 99},
			{50, 99},
			{48, 99},
			{47, 99},
			{41, 99},
			{36, 99},
			{26, 99},
			{24, 99},
			{23, 99},
			{7, 99},
			{2, 99},
			{0, 99},
			{0, 96},
			{0, 89},
			{0, 85},
			{0, 72},
			{0, 67},
			{0, 63},
			{0, 55},
			{0, 54},
			{0, 49},
			{0, 40},
			{0, 39},
			{0, 32},
			{0, 26},
			{0, 25},
			{0, 22},
			{0, 19},
			{0, 4},
		},
	},

	{
		[][]int{{0, 0}, {1, 0}, {1, 3}, {1, 8}, {1, 9}, {2, 0}, {2, 6}, {3, 0}, {3, 1}, {3, 4}, {3, 6}, {4, 2}, {4, 6}, {5, 7}, {5, 8}, {6, 2}, {6, 4}, {7, 7}, {7, 9}, {8, 0}, {8, 1}, {8, 3}, {8, 5}, {9, 6}},
		[][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {8, 0}, {9, 6}, {7, 9}, {1, 9}},
	},

	{
		[][]int{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {2, 2}, {3, 2}, {3, 1}, {3, 0}, {2, 0}, {1, 0}, {1, 1}, {4, 3}},
		[][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 3}, {0, 2}, {0, 1}},
	},

	{
		[][]int{{0, 0}, {0, 100}, {100, 100}, {100, 0}, {50, 50}},
		[][]int{{0, 0}, {0, 100}, {100, 100}, {100, 0}},
	},

	{
		[][]int{{4, 2}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {1, 1}},
		[][]int{{2, 0}, {4, 2}, {3, 3}, {2, 4}, {1, 1}},
	},

	{
		[][]int{{1, 2}, {2, 2}, {4, 2}},
		[][]int{{1, 2}, {2, 2}, {4, 2}},
	},

	{
		[][]int{{2, 2}, {4, 2}, {1, 2}},
		[][]int{{1, 2}, {2, 2}, {4, 2}},
	},

	// 可以有多个 testcase
}

func Test_outerTrees(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		// fmt.Printf("~~%v~~\n", tc)
		points := kit.Intss2Points(tc.points)
		expected := kit.Points2Intss(outerTrees(points))
		sort.Sort(intss(expected))
		sort.Sort(intss(tc.ans))
		ast.Equal(tc.ans, expected)
	}
}

func Benchmark_outerTrees(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			points := kit.Intss2Points(tc.points)
			outerTrees(points)
		}
	}
}

// intss 实现了 sort.Interface 接口
type intss [][]int

func (iss intss) Len() int { return len(iss) }

func (iss intss) Less(i, j int) bool {
	if iss[i][0] == iss[j][0] {
		return iss[i][1] < iss[j][1]
	}
	return iss[i][0] < iss[j][0]
}

func (iss intss) Swap(i, j int) { iss[i], iss[j] = iss[j], iss[i] }
