package simplehash

import "strconv"

// TODO maybe need to flexible
const total = 10

type SimpleHash struct {
}

func NewSimpleHash() *SimpleHash {
	return &SimpleHash{}
}

func (s *SimpleHash) CreateSuffixLabel(raw uint32) string {
	// 没范型很难提高代码重用写出规模代码库 golang2.0才提出范型草案
	digit := raw % total

	return strconv.Itoa(int(digit))
}
