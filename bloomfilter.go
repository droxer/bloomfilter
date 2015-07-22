package bloomfilter

import (
	"hash"
	"hash/fnv"
)

type BloomFilter struct {
	bitVector []bool
	k         int
	n         int
	m         int
	hashFn    hash.Hash64
}

func New(numHashFn, bfSize int) *BloomFilter {
	return &BloomFilter{
		bitVector: make([]bool, bfSize),
		k:         numHashFn,
		m:         bfSize,
		n:         0,
		hashFn:    fnv.New64(),
	}
}

func (bf *BloomFilter) Add(e []byte) {
	h1, h2 := bf.getHash(e)
	for i := 0; i < bf.k; i++ {
		ind := (h1 + uint32(i)*h2) % uint32(bf.m)
		bf.bitVector[ind] = true
	}
	bf.n++
}

func (bf *BloomFilter) MayContains(e []byte) bool {
	var result = true
	h1, h2 := bf.getHash(e)
	for i := 0; i < bf.k; i++ {
		ind := (h1 + uint32(i)*h2) % uint32(bf.m)
		result = result && bf.bitVector[ind]
	}
	return result
}

func (bf *BloomFilter) getHash(b []byte) (uint32, uint32) {
	bf.hashFn.Reset()
	bf.hashFn.Write(b)
	hash64 := bf.hashFn.Sum64()

	h1 := uint32(hash64 & ((1 << 32) - 1))
	h2 := uint32(hash64 >> 32)
	return h1, h2
}
