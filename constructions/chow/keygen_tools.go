// Contains tools for key generation that don't fit anywhere else.
package chow

import (
	"../../primitives/encoding"
	"../../primitives/table"
	"crypto/aes"
	"crypto/cipher"
	"io"
)

// Two Expand->Squash rounds comprise one AES round.  The Inside Surface is the output of the first E->S round (or the
// input of the second), and the Outside Surface is the output of the second E->S round (the whole AES round's output,
// feeding into the next round).
type Surface int

const (
	Inside Surface = iota
	Outside
)

// In a Squash step, we take four words and XOR them together into one word with 3 pairwise XORs: (a ^ b) ^ (c ^ d)
// This requires two top-level XOR tables, a Left computing (a ^ b) and a Right computing (c ^ d).  A bottom-level XOR
// takes the output of a Left and a Right gate and computes (a ^ b) ^ (c ^ d).
type Side int

const (
	Left Side = iota
	Right
)

// Index in, index out.  Example: shiftRows(5) = 1 because ShiftRows(block) returns [16]byte{block[0], block[5], ...
func shiftRows(i int) int {
	return []int{0, 13, 10, 7, 4, 1, 14, 11, 8, 5, 2, 15, 12, 9, 6, 3}[i]
}

// generateStream takes a (private) seed and a (possibly public) label and produces an io.Reader giving random bytes,
// useful for deterministically generating random matrices/encodings, in place of (crypto/rand).Reader.
//
// It does this by using the seed as an AES key and the label as the IV in CTR mode.  The io.Reader is providing the
// AES-CTR encryption of /dev/null.
type devNull struct{}

func (dn devNull) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = 0
	}

	return len(p), nil
}

func generateStream(seed, label [16]byte) io.Reader {
	// Generate sub-key
	subKey := [16]byte{}
	c, _ := aes.NewCipher(seed[:])
	c.Encrypt(subKey[:], label[:])

	// Create pseudo-random byte stream keyed by sub-key.
	block, _ := aes.NewCipher(subKey[:])
	stream := cipher.StreamReader{
		cipher.NewCTR(block, label[:]),
		devNull{},
	}

	return stream
}

func getShuffle(seed, label [16]byte) encoding.Shuffle {
	key := [32]byte{}
	copy(key[0:16], seed[:])
	copy(key[16:32], label[:])

	cached, ok := encodingCache[key]

	if ok {
		return cached
	} else {
		encodingCache[key] = encoding.GenerateShuffle(generateStream(seed, label))
		return encodingCache[key]
	}
}

// Generate the XOR Tables for squashing the result of a Input/Output Mask.
func blockXORTables(seed [16]byte, surface Surface) (out [32][15]table.Nibble) {
	for pos := 0; pos < 32; pos++ {
		out[pos][0] = encoding.NibbleTable{
			encoding.ConcatenatedByte{MaskEncoding(seed, 0, pos, surface), MaskEncoding(seed, 1, pos, surface)},
			XOREncoding(seed, 9, pos, surface, Left),
			XORTable{},
		}

		for i := 1; i < 14; i++ {
			out[pos][i] = encoding.NibbleTable{
				encoding.ConcatenatedByte{XOREncoding(seed, i+8, pos, surface, Left), MaskEncoding(seed, i+1, pos, surface)},
				XOREncoding(seed, i+9, pos, surface, Left),
				XORTable{},
			}
		}

		var outEnc encoding.Nibble
		if surface == Inside {
			outEnc = RoundEncoding(seed, -1, 2*shiftRows(pos/2)+pos%2, Outside)
		} else {
			outEnc = encoding.IdentityByte{}
		}

		out[pos][14] = encoding.NibbleTable{
			encoding.ConcatenatedByte{XOREncoding(seed, 22, pos, surface, Left), MaskEncoding(seed, 15, pos, surface)},
			outEnc,
			XORTable{},
		}
	}

	return
}

// Generate the XOR Tables for squashing the result of a Tyi Table or MB^(-1) Table.
func xorTables(seed [16]byte, surface Surface) (out [9][32][3]table.Nibble) {
	var outPos func(int) int
	if surface == Inside {
		outPos = func(pos int) int { return pos }
	} else {
		outPos = func(pos int) int { return 2*shiftRows(pos/2) + pos%2 }
	}

	for round := 0; round < 9; round++ {
		for pos := 0; pos < 32; pos++ {
			out[round][pos][0] = topLevelXORTable(seed, round, pos, surface, Left)
			out[round][pos][1] = topLevelXORTable(seed, round, pos, surface, Right)
			out[round][pos][2] = encoding.NibbleTable{
				encoding.ConcatenatedByte{
					XOREncoding(seed, round, pos, surface, Left),
					XOREncoding(seed, round, pos, surface, Right),
				},
				RoundEncoding(seed, round, outPos(pos), surface),
				XORTable{},
			}
		}
	}

	return
}

// Generate a top-level XOR (one that takes new input, rather than one that takes semi-squashed input)
func topLevelXORTable(seed [16]byte, round, pos int, surface Surface, side Side) table.Nibble {
	return encoding.NibbleTable{
		encoding.ConcatenatedByte{
			StepEncoding(seed, round, pos/8*4+2*int(side)+0, pos%8, surface),
			StepEncoding(seed, round, pos/8*4+2*int(side)+1, pos%8, surface),
		},
		XOREncoding(seed, round, pos, surface, side),
		XORTable{},
	}
}