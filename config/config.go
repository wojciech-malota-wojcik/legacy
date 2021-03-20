package config

// RequiredToDecrypt specifies how many successors have to load their keys to decrypt data
const RequiredToDecrypt = 3

// SeedSize is the byte size of generated seed
const SeedSize = 1000000

// SeedToKeySteps specifies how many hashing steps is made to convert initial seed to the one used to generate decryption key
const SeedToKeySteps = 1000

// AESKeySize specifies size of AES key
const AESKeySize = 32

var (
	// Successors store public part of keys stored on YubiKeys owned by successors
	Successors = [][]byte{
		{0x30, 0x82, 0x1, 0xa, 0x2, 0x82, 0x1, 0x1, 0x0, 0xbd, 0xc8, 0x75, 0x71, 0x2, 0x6b, 0xc4, 0xa7, 0x14, 0x16, 0x61, 0xa0, 0x8d, 0x24, 0x85, 0xdd, 0xf8, 0x34, 0xf6, 0x21, 0x8b, 0xbe, 0x17, 0xce, 0xc2, 0xdf, 0x42, 0x32, 0x51, 0xb8, 0xc5, 0x4, 0xe0, 0x6c, 0x7d, 0x63, 0x4a, 0xb9, 0xad, 0xd2, 0xcf, 0x34, 0x81, 0xfd, 0xfc, 0xee, 0xe4, 0xe0, 0x33, 0xeb, 0x5a, 0x6c, 0x40, 0x12, 0x3d, 0x7c, 0x13, 0x6e, 0x93, 0x6b, 0xe, 0x98, 0x90, 0x7a, 0x91, 0x40, 0xbb, 0x35, 0xd9, 0x1, 0x8f, 0x6b, 0x85, 0x56, 0xc7, 0xf7, 0x50, 0x1d, 0xee, 0x20, 0x4d, 0xdc, 0xa5, 0x97, 0x97, 0xeb, 0x81, 0x21, 0x51, 0xc, 0x71, 0xb1, 0x6c, 0x90, 0x46, 0x21, 0x9f, 0xf4, 0xa4, 0xd5, 0xe7, 0x77, 0x10, 0x9a, 0xab, 0x92, 0x6a, 0x40, 0x11, 0xd4, 0x1d, 0x48, 0xa1, 0x74, 0x73, 0xed, 0xad, 0x19, 0x91, 0x56, 0x18, 0xed, 0xb, 0x6c, 0xca, 0x27, 0xef, 0x32, 0x7d, 0xf, 0x95, 0x58, 0xc9, 0xce, 0xee, 0x71, 0xbb, 0x18, 0xff, 0x6d, 0xb6, 0xf0, 0xb8, 0x6a, 0x50, 0x4, 0xde, 0x5, 0xb0, 0xc, 0xe9, 0x83, 0x60, 0xfe, 0x2, 0x84, 0xf6, 0x44, 0xed, 0xc1, 0xc9, 0xdc, 0x9c, 0xa4, 0x53, 0xa0, 0xd3, 0xaf, 0x4a, 0xe3, 0x24, 0x93, 0xef, 0x73, 0xab, 0x14, 0x76, 0x4a, 0xda, 0x98, 0xcb, 0xea, 0x4a, 0x7f, 0x4e, 0xf1, 0x94, 0x56, 0x77, 0xcd, 0x1b, 0x71, 0x13, 0x4f, 0xb6, 0x80, 0x1b, 0xf, 0x41, 0xcd, 0x82, 0xb9, 0x15, 0x51, 0x98, 0xc7, 0xa5, 0xbd, 0x3a, 0xe7, 0xf4, 0xe4, 0x56, 0xf5, 0x0, 0x30, 0x3b, 0xdd, 0xf6, 0xdc, 0xa4, 0x10, 0x81, 0xf, 0x8f, 0xc3, 0xeb, 0xed, 0xe2, 0xe6, 0xfe, 0xe7, 0xd7, 0x2a, 0xf5, 0x23, 0xc8, 0x14, 0xf0, 0xc7, 0xa4, 0x67, 0x9f, 0xe0, 0x49, 0x66, 0xcc, 0xc6, 0xb2, 0xa1, 0x34, 0x36, 0x52, 0xc4, 0xb9, 0x81, 0x2, 0x3, 0x1, 0x0, 0x1},
		{0x30, 0x82, 0x1, 0xa, 0x2, 0x82, 0x1, 0x1, 0x0, 0xbd, 0xc8, 0x75, 0x71, 0x2, 0x6b, 0xc4, 0xa7, 0x14, 0x16, 0x61, 0xa0, 0x8d, 0x24, 0x85, 0xdd, 0xf8, 0x34, 0xf6, 0x21, 0x8b, 0xbe, 0x17, 0xce, 0xc2, 0xdf, 0x42, 0x32, 0x51, 0xb8, 0xc5, 0x4, 0xe0, 0x6c, 0x7d, 0x63, 0x4a, 0xb9, 0xad, 0xd2, 0xcf, 0x34, 0x81, 0xfd, 0xfc, 0xee, 0xe4, 0xe0, 0x33, 0xeb, 0x5a, 0x6c, 0x40, 0x12, 0x3d, 0x7c, 0x13, 0x6e, 0x93, 0x6b, 0xe, 0x98, 0x90, 0x7a, 0x91, 0x40, 0xbb, 0x35, 0xd9, 0x1, 0x8f, 0x6b, 0x85, 0x56, 0xc7, 0xf7, 0x50, 0x1d, 0xee, 0x20, 0x4d, 0xdc, 0xa5, 0x97, 0x97, 0xeb, 0x81, 0x21, 0x51, 0xc, 0x71, 0xb1, 0x6c, 0x90, 0x46, 0x21, 0x9f, 0xf4, 0xa4, 0xd5, 0xe7, 0x77, 0x10, 0x9a, 0xab, 0x92, 0x6a, 0x40, 0x11, 0xd4, 0x1d, 0x48, 0xa1, 0x74, 0x73, 0xed, 0xad, 0x19, 0x91, 0x56, 0x18, 0xed, 0xb, 0x6c, 0xca, 0x27, 0xef, 0x32, 0x7d, 0xf, 0x95, 0x58, 0xc9, 0xce, 0xee, 0x71, 0xbb, 0x18, 0xff, 0x6d, 0xb6, 0xf0, 0xb8, 0x6a, 0x50, 0x4, 0xde, 0x5, 0xb0, 0xc, 0xe9, 0x83, 0x60, 0xfe, 0x2, 0x84, 0xf6, 0x44, 0xed, 0xc1, 0xc9, 0xdc, 0x9c, 0xa4, 0x53, 0xa0, 0xd3, 0xaf, 0x4a, 0xe3, 0x24, 0x93, 0xef, 0x73, 0xab, 0x14, 0x76, 0x4a, 0xda, 0x98, 0xcb, 0xea, 0x4a, 0x7f, 0x4e, 0xf1, 0x94, 0x56, 0x77, 0xcd, 0x1b, 0x71, 0x13, 0x4f, 0xb6, 0x80, 0x1b, 0xf, 0x41, 0xcd, 0x82, 0xb9, 0x15, 0x51, 0x98, 0xc7, 0xa5, 0xbd, 0x3a, 0xe7, 0xf4, 0xe4, 0x56, 0xf5, 0x0, 0x30, 0x3b, 0xdd, 0xf6, 0xdc, 0xa4, 0x10, 0x81, 0xf, 0x8f, 0xc3, 0xeb, 0xed, 0xe2, 0xe6, 0xfe, 0xe7, 0xd7, 0x2a, 0xf5, 0x23, 0xc8, 0x14, 0xf0, 0xc7, 0xa4, 0x67, 0x9f, 0xe0, 0x49, 0x66, 0xcc, 0xc6, 0xb2, 0xa1, 0x34, 0x36, 0x52, 0xc4, 0xb9, 0x81, 0x2, 0x3, 0x1, 0x0, 0x1},
		{0x30, 0x82, 0x1, 0xa, 0x2, 0x82, 0x1, 0x1, 0x0, 0xbd, 0xc8, 0x75, 0x71, 0x2, 0x6b, 0xc4, 0xa7, 0x14, 0x16, 0x61, 0xa0, 0x8d, 0x24, 0x85, 0xdd, 0xf8, 0x34, 0xf6, 0x21, 0x8b, 0xbe, 0x17, 0xce, 0xc2, 0xdf, 0x42, 0x32, 0x51, 0xb8, 0xc5, 0x4, 0xe0, 0x6c, 0x7d, 0x63, 0x4a, 0xb9, 0xad, 0xd2, 0xcf, 0x34, 0x81, 0xfd, 0xfc, 0xee, 0xe4, 0xe0, 0x33, 0xeb, 0x5a, 0x6c, 0x40, 0x12, 0x3d, 0x7c, 0x13, 0x6e, 0x93, 0x6b, 0xe, 0x98, 0x90, 0x7a, 0x91, 0x40, 0xbb, 0x35, 0xd9, 0x1, 0x8f, 0x6b, 0x85, 0x56, 0xc7, 0xf7, 0x50, 0x1d, 0xee, 0x20, 0x4d, 0xdc, 0xa5, 0x97, 0x97, 0xeb, 0x81, 0x21, 0x51, 0xc, 0x71, 0xb1, 0x6c, 0x90, 0x46, 0x21, 0x9f, 0xf4, 0xa4, 0xd5, 0xe7, 0x77, 0x10, 0x9a, 0xab, 0x92, 0x6a, 0x40, 0x11, 0xd4, 0x1d, 0x48, 0xa1, 0x74, 0x73, 0xed, 0xad, 0x19, 0x91, 0x56, 0x18, 0xed, 0xb, 0x6c, 0xca, 0x27, 0xef, 0x32, 0x7d, 0xf, 0x95, 0x58, 0xc9, 0xce, 0xee, 0x71, 0xbb, 0x18, 0xff, 0x6d, 0xb6, 0xf0, 0xb8, 0x6a, 0x50, 0x4, 0xde, 0x5, 0xb0, 0xc, 0xe9, 0x83, 0x60, 0xfe, 0x2, 0x84, 0xf6, 0x44, 0xed, 0xc1, 0xc9, 0xdc, 0x9c, 0xa4, 0x53, 0xa0, 0xd3, 0xaf, 0x4a, 0xe3, 0x24, 0x93, 0xef, 0x73, 0xab, 0x14, 0x76, 0x4a, 0xda, 0x98, 0xcb, 0xea, 0x4a, 0x7f, 0x4e, 0xf1, 0x94, 0x56, 0x77, 0xcd, 0x1b, 0x71, 0x13, 0x4f, 0xb6, 0x80, 0x1b, 0xf, 0x41, 0xcd, 0x82, 0xb9, 0x15, 0x51, 0x98, 0xc7, 0xa5, 0xbd, 0x3a, 0xe7, 0xf4, 0xe4, 0x56, 0xf5, 0x0, 0x30, 0x3b, 0xdd, 0xf6, 0xdc, 0xa4, 0x10, 0x81, 0xf, 0x8f, 0xc3, 0xeb, 0xed, 0xe2, 0xe6, 0xfe, 0xe7, 0xd7, 0x2a, 0xf5, 0x23, 0xc8, 0x14, 0xf0, 0xc7, 0xa4, 0x67, 0x9f, 0xe0, 0x49, 0x66, 0xcc, 0xc6, 0xb2, 0xa1, 0x34, 0x36, 0x52, 0xc4, 0xb9, 0x81, 0x2, 0x3, 0x1, 0x0, 0x1},
		{0x30, 0x82, 0x1, 0xa, 0x2, 0x82, 0x1, 0x1, 0x0, 0xbd, 0xc8, 0x75, 0x71, 0x2, 0x6b, 0xc4, 0xa7, 0x14, 0x16, 0x61, 0xa0, 0x8d, 0x24, 0x85, 0xdd, 0xf8, 0x34, 0xf6, 0x21, 0x8b, 0xbe, 0x17, 0xce, 0xc2, 0xdf, 0x42, 0x32, 0x51, 0xb8, 0xc5, 0x4, 0xe0, 0x6c, 0x7d, 0x63, 0x4a, 0xb9, 0xad, 0xd2, 0xcf, 0x34, 0x81, 0xfd, 0xfc, 0xee, 0xe4, 0xe0, 0x33, 0xeb, 0x5a, 0x6c, 0x40, 0x12, 0x3d, 0x7c, 0x13, 0x6e, 0x93, 0x6b, 0xe, 0x98, 0x90, 0x7a, 0x91, 0x40, 0xbb, 0x35, 0xd9, 0x1, 0x8f, 0x6b, 0x85, 0x56, 0xc7, 0xf7, 0x50, 0x1d, 0xee, 0x20, 0x4d, 0xdc, 0xa5, 0x97, 0x97, 0xeb, 0x81, 0x21, 0x51, 0xc, 0x71, 0xb1, 0x6c, 0x90, 0x46, 0x21, 0x9f, 0xf4, 0xa4, 0xd5, 0xe7, 0x77, 0x10, 0x9a, 0xab, 0x92, 0x6a, 0x40, 0x11, 0xd4, 0x1d, 0x48, 0xa1, 0x74, 0x73, 0xed, 0xad, 0x19, 0x91, 0x56, 0x18, 0xed, 0xb, 0x6c, 0xca, 0x27, 0xef, 0x32, 0x7d, 0xf, 0x95, 0x58, 0xc9, 0xce, 0xee, 0x71, 0xbb, 0x18, 0xff, 0x6d, 0xb6, 0xf0, 0xb8, 0x6a, 0x50, 0x4, 0xde, 0x5, 0xb0, 0xc, 0xe9, 0x83, 0x60, 0xfe, 0x2, 0x84, 0xf6, 0x44, 0xed, 0xc1, 0xc9, 0xdc, 0x9c, 0xa4, 0x53, 0xa0, 0xd3, 0xaf, 0x4a, 0xe3, 0x24, 0x93, 0xef, 0x73, 0xab, 0x14, 0x76, 0x4a, 0xda, 0x98, 0xcb, 0xea, 0x4a, 0x7f, 0x4e, 0xf1, 0x94, 0x56, 0x77, 0xcd, 0x1b, 0x71, 0x13, 0x4f, 0xb6, 0x80, 0x1b, 0xf, 0x41, 0xcd, 0x82, 0xb9, 0x15, 0x51, 0x98, 0xc7, 0xa5, 0xbd, 0x3a, 0xe7, 0xf4, 0xe4, 0x56, 0xf5, 0x0, 0x30, 0x3b, 0xdd, 0xf6, 0xdc, 0xa4, 0x10, 0x81, 0xf, 0x8f, 0xc3, 0xeb, 0xed, 0xe2, 0xe6, 0xfe, 0xe7, 0xd7, 0x2a, 0xf5, 0x23, 0xc8, 0x14, 0xf0, 0xc7, 0xa4, 0x67, 0x9f, 0xe0, 0x49, 0x66, 0xcc, 0xc6, 0xb2, 0xa1, 0x34, 0x36, 0x52, 0xc4, 0xb9, 0x81, 0x2, 0x3, 0x1, 0x0, 0x1},
		{0x30, 0x82, 0x1, 0xa, 0x2, 0x82, 0x1, 0x1, 0x0, 0xbd, 0xc8, 0x75, 0x71, 0x2, 0x6b, 0xc4, 0xa7, 0x14, 0x16, 0x61, 0xa0, 0x8d, 0x24, 0x85, 0xdd, 0xf8, 0x34, 0xf6, 0x21, 0x8b, 0xbe, 0x17, 0xce, 0xc2, 0xdf, 0x42, 0x32, 0x51, 0xb8, 0xc5, 0x4, 0xe0, 0x6c, 0x7d, 0x63, 0x4a, 0xb9, 0xad, 0xd2, 0xcf, 0x34, 0x81, 0xfd, 0xfc, 0xee, 0xe4, 0xe0, 0x33, 0xeb, 0x5a, 0x6c, 0x40, 0x12, 0x3d, 0x7c, 0x13, 0x6e, 0x93, 0x6b, 0xe, 0x98, 0x90, 0x7a, 0x91, 0x40, 0xbb, 0x35, 0xd9, 0x1, 0x8f, 0x6b, 0x85, 0x56, 0xc7, 0xf7, 0x50, 0x1d, 0xee, 0x20, 0x4d, 0xdc, 0xa5, 0x97, 0x97, 0xeb, 0x81, 0x21, 0x51, 0xc, 0x71, 0xb1, 0x6c, 0x90, 0x46, 0x21, 0x9f, 0xf4, 0xa4, 0xd5, 0xe7, 0x77, 0x10, 0x9a, 0xab, 0x92, 0x6a, 0x40, 0x11, 0xd4, 0x1d, 0x48, 0xa1, 0x74, 0x73, 0xed, 0xad, 0x19, 0x91, 0x56, 0x18, 0xed, 0xb, 0x6c, 0xca, 0x27, 0xef, 0x32, 0x7d, 0xf, 0x95, 0x58, 0xc9, 0xce, 0xee, 0x71, 0xbb, 0x18, 0xff, 0x6d, 0xb6, 0xf0, 0xb8, 0x6a, 0x50, 0x4, 0xde, 0x5, 0xb0, 0xc, 0xe9, 0x83, 0x60, 0xfe, 0x2, 0x84, 0xf6, 0x44, 0xed, 0xc1, 0xc9, 0xdc, 0x9c, 0xa4, 0x53, 0xa0, 0xd3, 0xaf, 0x4a, 0xe3, 0x24, 0x93, 0xef, 0x73, 0xab, 0x14, 0x76, 0x4a, 0xda, 0x98, 0xcb, 0xea, 0x4a, 0x7f, 0x4e, 0xf1, 0x94, 0x56, 0x77, 0xcd, 0x1b, 0x71, 0x13, 0x4f, 0xb6, 0x80, 0x1b, 0xf, 0x41, 0xcd, 0x82, 0xb9, 0x15, 0x51, 0x98, 0xc7, 0xa5, 0xbd, 0x3a, 0xe7, 0xf4, 0xe4, 0x56, 0xf5, 0x0, 0x30, 0x3b, 0xdd, 0xf6, 0xdc, 0xa4, 0x10, 0x81, 0xf, 0x8f, 0xc3, 0xeb, 0xed, 0xe2, 0xe6, 0xfe, 0xe7, 0xd7, 0x2a, 0xf5, 0x23, 0xc8, 0x14, 0xf0, 0xc7, 0xa4, 0x67, 0x9f, 0xe0, 0x49, 0x66, 0xcc, 0xc6, 0xb2, 0xa1, 0x34, 0x36, 0x52, 0xc4, 0xb9, 0x81, 0x2, 0x3, 0x1, 0x0, 0x1},
		{0x30, 0x82, 0x1, 0xa, 0x2, 0x82, 0x1, 0x1, 0x0, 0xbd, 0xc8, 0x75, 0x71, 0x2, 0x6b, 0xc4, 0xa7, 0x14, 0x16, 0x61, 0xa0, 0x8d, 0x24, 0x85, 0xdd, 0xf8, 0x34, 0xf6, 0x21, 0x8b, 0xbe, 0x17, 0xce, 0xc2, 0xdf, 0x42, 0x32, 0x51, 0xb8, 0xc5, 0x4, 0xe0, 0x6c, 0x7d, 0x63, 0x4a, 0xb9, 0xad, 0xd2, 0xcf, 0x34, 0x81, 0xfd, 0xfc, 0xee, 0xe4, 0xe0, 0x33, 0xeb, 0x5a, 0x6c, 0x40, 0x12, 0x3d, 0x7c, 0x13, 0x6e, 0x93, 0x6b, 0xe, 0x98, 0x90, 0x7a, 0x91, 0x40, 0xbb, 0x35, 0xd9, 0x1, 0x8f, 0x6b, 0x85, 0x56, 0xc7, 0xf7, 0x50, 0x1d, 0xee, 0x20, 0x4d, 0xdc, 0xa5, 0x97, 0x97, 0xeb, 0x81, 0x21, 0x51, 0xc, 0x71, 0xb1, 0x6c, 0x90, 0x46, 0x21, 0x9f, 0xf4, 0xa4, 0xd5, 0xe7, 0x77, 0x10, 0x9a, 0xab, 0x92, 0x6a, 0x40, 0x11, 0xd4, 0x1d, 0x48, 0xa1, 0x74, 0x73, 0xed, 0xad, 0x19, 0x91, 0x56, 0x18, 0xed, 0xb, 0x6c, 0xca, 0x27, 0xef, 0x32, 0x7d, 0xf, 0x95, 0x58, 0xc9, 0xce, 0xee, 0x71, 0xbb, 0x18, 0xff, 0x6d, 0xb6, 0xf0, 0xb8, 0x6a, 0x50, 0x4, 0xde, 0x5, 0xb0, 0xc, 0xe9, 0x83, 0x60, 0xfe, 0x2, 0x84, 0xf6, 0x44, 0xed, 0xc1, 0xc9, 0xdc, 0x9c, 0xa4, 0x53, 0xa0, 0xd3, 0xaf, 0x4a, 0xe3, 0x24, 0x93, 0xef, 0x73, 0xab, 0x14, 0x76, 0x4a, 0xda, 0x98, 0xcb, 0xea, 0x4a, 0x7f, 0x4e, 0xf1, 0x94, 0x56, 0x77, 0xcd, 0x1b, 0x71, 0x13, 0x4f, 0xb6, 0x80, 0x1b, 0xf, 0x41, 0xcd, 0x82, 0xb9, 0x15, 0x51, 0x98, 0xc7, 0xa5, 0xbd, 0x3a, 0xe7, 0xf4, 0xe4, 0x56, 0xf5, 0x0, 0x30, 0x3b, 0xdd, 0xf6, 0xdc, 0xa4, 0x10, 0x81, 0xf, 0x8f, 0xc3, 0xeb, 0xed, 0xe2, 0xe6, 0xfe, 0xe7, 0xd7, 0x2a, 0xf5, 0x23, 0xc8, 0x14, 0xf0, 0xc7, 0xa4, 0x67, 0x9f, 0xe0, 0x49, 0x66, 0xcc, 0xc6, 0xb2, 0xa1, 0x34, 0x36, 0x52, 0xc4, 0xb9, 0x81, 0x2, 0x3, 0x1, 0x0, 0x1},
		{0x30, 0x82, 0x1, 0xa, 0x2, 0x82, 0x1, 0x1, 0x0, 0xbd, 0xc8, 0x75, 0x71, 0x2, 0x6b, 0xc4, 0xa7, 0x14, 0x16, 0x61, 0xa0, 0x8d, 0x24, 0x85, 0xdd, 0xf8, 0x34, 0xf6, 0x21, 0x8b, 0xbe, 0x17, 0xce, 0xc2, 0xdf, 0x42, 0x32, 0x51, 0xb8, 0xc5, 0x4, 0xe0, 0x6c, 0x7d, 0x63, 0x4a, 0xb9, 0xad, 0xd2, 0xcf, 0x34, 0x81, 0xfd, 0xfc, 0xee, 0xe4, 0xe0, 0x33, 0xeb, 0x5a, 0x6c, 0x40, 0x12, 0x3d, 0x7c, 0x13, 0x6e, 0x93, 0x6b, 0xe, 0x98, 0x90, 0x7a, 0x91, 0x40, 0xbb, 0x35, 0xd9, 0x1, 0x8f, 0x6b, 0x85, 0x56, 0xc7, 0xf7, 0x50, 0x1d, 0xee, 0x20, 0x4d, 0xdc, 0xa5, 0x97, 0x97, 0xeb, 0x81, 0x21, 0x51, 0xc, 0x71, 0xb1, 0x6c, 0x90, 0x46, 0x21, 0x9f, 0xf4, 0xa4, 0xd5, 0xe7, 0x77, 0x10, 0x9a, 0xab, 0x92, 0x6a, 0x40, 0x11, 0xd4, 0x1d, 0x48, 0xa1, 0x74, 0x73, 0xed, 0xad, 0x19, 0x91, 0x56, 0x18, 0xed, 0xb, 0x6c, 0xca, 0x27, 0xef, 0x32, 0x7d, 0xf, 0x95, 0x58, 0xc9, 0xce, 0xee, 0x71, 0xbb, 0x18, 0xff, 0x6d, 0xb6, 0xf0, 0xb8, 0x6a, 0x50, 0x4, 0xde, 0x5, 0xb0, 0xc, 0xe9, 0x83, 0x60, 0xfe, 0x2, 0x84, 0xf6, 0x44, 0xed, 0xc1, 0xc9, 0xdc, 0x9c, 0xa4, 0x53, 0xa0, 0xd3, 0xaf, 0x4a, 0xe3, 0x24, 0x93, 0xef, 0x73, 0xab, 0x14, 0x76, 0x4a, 0xda, 0x98, 0xcb, 0xea, 0x4a, 0x7f, 0x4e, 0xf1, 0x94, 0x56, 0x77, 0xcd, 0x1b, 0x71, 0x13, 0x4f, 0xb6, 0x80, 0x1b, 0xf, 0x41, 0xcd, 0x82, 0xb9, 0x15, 0x51, 0x98, 0xc7, 0xa5, 0xbd, 0x3a, 0xe7, 0xf4, 0xe4, 0x56, 0xf5, 0x0, 0x30, 0x3b, 0xdd, 0xf6, 0xdc, 0xa4, 0x10, 0x81, 0xf, 0x8f, 0xc3, 0xeb, 0xed, 0xe2, 0xe6, 0xfe, 0xe7, 0xd7, 0x2a, 0xf5, 0x23, 0xc8, 0x14, 0xf0, 0xc7, 0xa4, 0x67, 0x9f, 0xe0, 0x49, 0x66, 0xcc, 0xc6, 0xb2, 0xa1, 0x34, 0x36, 0x52, 0xc4, 0xb9, 0x81, 0x2, 0x3, 0x1, 0x0, 0x1},
		{0x30, 0x82, 0x1, 0xa, 0x2, 0x82, 0x1, 0x1, 0x0, 0xbd, 0xc8, 0x75, 0x71, 0x2, 0x6b, 0xc4, 0xa7, 0x14, 0x16, 0x61, 0xa0, 0x8d, 0x24, 0x85, 0xdd, 0xf8, 0x34, 0xf6, 0x21, 0x8b, 0xbe, 0x17, 0xce, 0xc2, 0xdf, 0x42, 0x32, 0x51, 0xb8, 0xc5, 0x4, 0xe0, 0x6c, 0x7d, 0x63, 0x4a, 0xb9, 0xad, 0xd2, 0xcf, 0x34, 0x81, 0xfd, 0xfc, 0xee, 0xe4, 0xe0, 0x33, 0xeb, 0x5a, 0x6c, 0x40, 0x12, 0x3d, 0x7c, 0x13, 0x6e, 0x93, 0x6b, 0xe, 0x98, 0x90, 0x7a, 0x91, 0x40, 0xbb, 0x35, 0xd9, 0x1, 0x8f, 0x6b, 0x85, 0x56, 0xc7, 0xf7, 0x50, 0x1d, 0xee, 0x20, 0x4d, 0xdc, 0xa5, 0x97, 0x97, 0xeb, 0x81, 0x21, 0x51, 0xc, 0x71, 0xb1, 0x6c, 0x90, 0x46, 0x21, 0x9f, 0xf4, 0xa4, 0xd5, 0xe7, 0x77, 0x10, 0x9a, 0xab, 0x92, 0x6a, 0x40, 0x11, 0xd4, 0x1d, 0x48, 0xa1, 0x74, 0x73, 0xed, 0xad, 0x19, 0x91, 0x56, 0x18, 0xed, 0xb, 0x6c, 0xca, 0x27, 0xef, 0x32, 0x7d, 0xf, 0x95, 0x58, 0xc9, 0xce, 0xee, 0x71, 0xbb, 0x18, 0xff, 0x6d, 0xb6, 0xf0, 0xb8, 0x6a, 0x50, 0x4, 0xde, 0x5, 0xb0, 0xc, 0xe9, 0x83, 0x60, 0xfe, 0x2, 0x84, 0xf6, 0x44, 0xed, 0xc1, 0xc9, 0xdc, 0x9c, 0xa4, 0x53, 0xa0, 0xd3, 0xaf, 0x4a, 0xe3, 0x24, 0x93, 0xef, 0x73, 0xab, 0x14, 0x76, 0x4a, 0xda, 0x98, 0xcb, 0xea, 0x4a, 0x7f, 0x4e, 0xf1, 0x94, 0x56, 0x77, 0xcd, 0x1b, 0x71, 0x13, 0x4f, 0xb6, 0x80, 0x1b, 0xf, 0x41, 0xcd, 0x82, 0xb9, 0x15, 0x51, 0x98, 0xc7, 0xa5, 0xbd, 0x3a, 0xe7, 0xf4, 0xe4, 0x56, 0xf5, 0x0, 0x30, 0x3b, 0xdd, 0xf6, 0xdc, 0xa4, 0x10, 0x81, 0xf, 0x8f, 0xc3, 0xeb, 0xed, 0xe2, 0xe6, 0xfe, 0xe7, 0xd7, 0x2a, 0xf5, 0x23, 0xc8, 0x14, 0xf0, 0xc7, 0xa4, 0x67, 0x9f, 0xe0, 0x49, 0x66, 0xcc, 0xc6, 0xb2, 0xa1, 0x34, 0x36, 0x52, 0xc4, 0xb9, 0x81, 0x2, 0x3, 0x1, 0x0, 0x1},
	}
)
