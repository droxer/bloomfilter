package bloomfilter_test

import (
	"github.com/droxer/bloomfilter"
	"testing"
)

func TestTrueForExistedElement(t *testing.T) {
	bf := bloomfilter.New(3, 20)
	item := []byte("Hey BloomFilter")
	bf.Add(item)

	if !bf.MayContains(item) {
		t.Errorf("expected bloomfilter contains %s\n", item)
	}
}

func TestFalseForNotExistedElement(t *testing.T) {
	bf := bloomfilter.New(3, 20)
	item := []byte("Bye BloomFilter")

	if bf.MayContains(item) {
		t.Errorf("expected bloomfilter contains %s\n", item)
	}
}
