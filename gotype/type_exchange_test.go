package gotype

import (
	"fmt"
	"testing"
)

func TestInt2StringByfmt(t *testing.T) {
	want := "123"
	generate := Int2StringByfmt(123)
	fmt.Println(want == generate)
}
func BenchmarkInt2StringByfmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int2StringByfmt(123)
	}
}

func TestInt2StringByStrconv(t *testing.T) {
	want := "123"
	generate := Int2StringByStrconv(123)
	fmt.Println(want == generate)
}

func BenchmarkInt2StringByStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int2StringByStrconv(123)
	}
}
