package bloomfilter

import (
	"testing"
)

func TestTrueForExistedElement(t *testing.T) {
	bf := New(3, 20)
	item := []byte("Hey BloomFilter")
	bf.Add(item)

	if !bf.MayContains(item) {
		t.Fatalf("expected bloomfilter contains %s\n", item)
	}
}

func TestFalseForNotExistedElement(t *testing.T) {
	bf := New(3, 20)
	item := []byte("Bye BloomFilter")

	if bf.MayContains(item) {
		t.Fatalf("expected bloomfilter contains %s\n", item)
	}
}
