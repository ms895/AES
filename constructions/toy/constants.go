package toy

import (
	"github.com/OpenWhiteBox/primitives/matrix"
)

var subBytesConst = matrix.Row{
	0x63, 0x63, 0x63, 0x63, 0x63, 0x63, 0x63, 0x63, 0x63, 0x63, 0x63, 0x63, 0x63, 0x63, 0x63, 0x63,
}

var unSubBytesConst = matrix.Row{
	0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05, 0x05,
}

var round = matrix.Matrix{
	matrix.Row{0xf8, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf1},
	matrix.Row{0x09, 0x00, 0x00, 0x00, 0x00, 0xea, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0xe3},
	matrix.Row{0xe3, 0x00, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xc7},
	matrix.Row{0x3f, 0x00, 0x00, 0x00, 0x00, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x8f},
	matrix.Row{0x77, 0x00, 0x00, 0x00, 0x00, 0x68, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x1f},
	matrix.Row{0x1f, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x3e},
	matrix.Row{0x3e, 0x00, 0x00, 0x00, 0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x7c},
	matrix.Row{0x7c, 0x00, 0x00, 0x00, 0x00, 0x84, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0xf8},
	matrix.Row{0xf1, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0xf1},
	matrix.Row{0xe3, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0xea, 0x00, 0x00, 0x00, 0x00, 0xe3},
	matrix.Row{0xc7, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x00, 0xc7},
	matrix.Row{0x8f, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x8f},
	matrix.Row{0x1f, 0x00, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0x00, 0x68, 0x00, 0x00, 0x00, 0x00, 0x1f},
	matrix.Row{0x3e, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00, 0x3e},
	matrix.Row{0x7c, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x7c},
	matrix.Row{0xf8, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x84, 0x00, 0x00, 0x00, 0x00, 0xf8},
	matrix.Row{0xf1, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x09},
	matrix.Row{0xe3, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0xea},
	matrix.Row{0xc7, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x24},
	matrix.Row{0x8f, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0xb0},
	matrix.Row{0x1f, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0x00, 0x68},
	matrix.Row{0x3e, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x21},
	matrix.Row{0x7c, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x42},
	matrix.Row{0xf8, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x84},
	matrix.Row{0x09, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf8},
	matrix.Row{0xea, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x09},
	matrix.Row{0x24, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xe3},
	matrix.Row{0xb0, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x3f},
	matrix.Row{0x68, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x77},
	matrix.Row{0x21, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x1f},
	matrix.Row{0x42, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x3e},
	matrix.Row{0x84, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x7c},
	matrix.Row{0x00, 0x00, 0x00, 0xf1, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xe3, 0x09, 0x00, 0x00, 0x00, 0x00, 0xea, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xc7, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x8f, 0x3f, 0x00, 0x00, 0x00, 0x00, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x1f, 0x77, 0x00, 0x00, 0x00, 0x00, 0x68, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x3e, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x7c, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xf8, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x84, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xf1, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xe3, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0xea, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xc7, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x24, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x8f, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0xb0, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x1f, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0x00, 0x68, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x3e, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x7c, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x42, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xf8, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x84, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x09, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xea, 0xe3, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x24, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xb0, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x68, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x77, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x21, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x42, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x84, 0xf8, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xf8, 0x09, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x09, 0xea, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xe3, 0x24, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x3f, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x77, 0x68, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x1f, 0x21, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x3e, 0x42, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x7c, 0x84, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00},
	matrix.Row{0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf1, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x09, 0x00, 0x00, 0x00, 0x00, 0xea, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xc7, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x3f, 0x00, 0x00, 0x00, 0x00, 0xb0, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x77, 0x00, 0x00, 0x00, 0x00, 0x68, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x42, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x84, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0xf1, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xea, 0x00, 0x00, 0x00, 0x00, 0xe3, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x00, 0xc7, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x68, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x84, 0x00, 0x00, 0x00, 0x00, 0xf8, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x09, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0xea, 0xe3, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x24, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0xb0, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0x00, 0x68, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x21, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x42, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x84, 0xf8, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x09, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x09, 0xea, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x24, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x3f, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x77, 0x68, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x21, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x42, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x84, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00},
	matrix.Row{0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf1, 0xf8, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xea, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x09, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x24, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xc7, 0xe3, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x3f, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x68, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x77, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x21, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x1f, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x3e, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x84, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x7c, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0xf1, 0xf1, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0xea, 0x00, 0x00, 0x00, 0x00, 0xe3, 0xe3, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x00, 0xc7, 0xc7, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0xb0, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x8f, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x77, 0x00, 0x00, 0x00, 0x00, 0x68, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x1f, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x21, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x3e, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x42, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x7c, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x84, 0x00, 0x00, 0x00, 0x00, 0xf8, 0xf8, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x09, 0xf1, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x00, 0x00, 0xea, 0xe3, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x24, 0xc7, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x3f, 0x00, 0x00, 0x00, 0x00, 0xb0, 0x8f, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x77, 0x00, 0x00, 0x00, 0x00, 0x68, 0x1f, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x21, 0x3e, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x42, 0x7c, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x84, 0xf8, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x09, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x09, 0xea, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x24, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x3f, 0xb0, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x77, 0x68, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x21, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x42, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x84, 0x00, 0x00, 0x00},
}

var unRound, _ = round.Invert()

var lastRound = matrix.Matrix{
	matrix.Row{0xf1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0xe3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0xc7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x8f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x3e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf1},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe3},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc7},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8f},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7c},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf1, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xe3, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc7, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1f, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x3e, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x7c, 0x00, 0x00, 0x00, 0x00},
	matrix.Row{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf8, 0x00, 0x00, 0x00, 0x00},
}

var firstRound, _ = lastRound.Invert()