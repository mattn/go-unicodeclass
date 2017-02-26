package unicodeclass

import (
	"testing"
)

func TestSimple(t *testing.T) {
	tests := []struct {
		r    rune
		want Class
	}{
		{'1', Word},
		{'a', Word},
		{'世', CJKIdeographs},
		{',', Punctation},
		{' ', Blank},
	}
	for _, test := range tests {
		got := Is(test.r)
		if got != test.want {
			t.Fatalf("want %v but %v: %q(%d)", test.want, got, string(test.r), test.r)
		}
	}
}
