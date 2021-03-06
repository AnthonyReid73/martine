package rle

import (
	"testing"
)

func TestRleEncode(t *testing.T) {
	in := []byte{0x10, 0x10, 0x10,
		0x20, 0x20,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x40,
		0x50}
	out := Encode(in)
	if len(out) != 10 {
		t.Fatalf("Expected 10 elements and gets %d\n", len(out))
	}
	t.Log(out)

	in = []byte{0x10, 0x10, 0x10,
		0x20, 0x20,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x40,
		0x50,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30}
	out = Encode(in)
	if len(out) != 12 {
		t.Fatalf("Expected 12 elements and gets %d\n", len(out))
	}
	t.Log(out)

	inRepeat := make([]byte, 0)
	for i := 0; i < 510; i++ {
		inRepeat = append(inRepeat, 0x10)
	}
	out = Encode(inRepeat)
	if len(out) != 4 {
		t.Fatalf("Expected 4 elements and gets %d\n", len(out))
	}
	t.Log(out)
}
