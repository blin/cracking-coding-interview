package palindromes

import (
	"math/big"
	"unicode"
)

func IsPalindromePermutation(s string) bool {
	runes := map[rune]int{}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		runes[r]++
	}

	oddRunesCount := 0
	for _, c := range runes {
		if c%2 == 0 {
			continue
		}
		oddRunesCount++
	}

	return (oddRunesCount == 1 || oddRunesCount == 0)
}

// rune     alias for int32
// uint     either 32 or 64 bits
// int      same size as uint
//
// from the above it follows that int(rune) never truncates
//
var IntZero = &big.Int{}

func IsPalindromePermutationBitVec(s string) bool {
	bitVec := new(big.Int)
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		bitIndex := int(r)
		bitVec.SetBit(bitVec, bitIndex, bitVec.Bit(bitIndex)^1)
	}
	if bitVec.Cmp(IntZero) == 0 {
		return true
	}
	bitVecMinusOne := new(big.Int)
	bitVecMinusOne.Sub(bitVec, big.NewInt(1))
	return bitVec.And(bitVec, bitVecMinusOne).Cmp(IntZero) == 0
}
