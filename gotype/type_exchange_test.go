package gotype

import (
	"fmt"
	"testing"
)

func TestInt2StringByfmt(t *testing.T) {
	want := "256"
	generate := Int2StringByfmt(256)
	fmt.Println(want == generate)
}
func BenchmarkInt2StringByfmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int2StringByfmt(256)
	}
}

func TestInt2StringByStrconv(t *testing.T) {
	want := "256"
	generate := Int2StringByStrconv(256)
	fmt.Println(want == generate)
}

func BenchmarkInt2StringByStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int2StringByStrconv(256)
	}
}

func TestInt642StringByfmt(t *testing.T) {
	want := "256"
	generate := Int642StringByStrconv(123)
	fmt.Println(want == generate)
}

func TestInt642StringByStrconv(t *testing.T) {
	want := "256"
	generate := Int642StringByfmt(256)
	fmt.Println(want == generate)
}

func BenchmarkInt642StringByStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int642StringByfmt(256)
	}
}
func BenchmarkInt642StringByfmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Int642StringByStrconv(256)
	}
}

func TestString2Rune(t *testing.T) {
	s := "test"
	if s != string(String2Rune(s)) {
		fmt.Errorf("TestString2Rune fail,test string case: %s", s)
	}
}

func TestChangeStringByByte(t *testing.T) {
	test := "*&advb"
	if test == ChangeStringByByte(test) {
		fmt.Errorf("test change string fail,string case: %s", test)
	}
}

func TestChangeStringByRune(t *testing.T) {
	test := "*&advb"
	if test == ChangeStringByByte(test) {
		fmt.Errorf("test change string fail,string case: %s", test)
	}
}

func BenchmarkChangeStringByRune(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChangeStringByRune("*&advb")
	}
}
func BenchmarkChangeStringByByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ChangeStringByByte("*&advb")
	}
}
