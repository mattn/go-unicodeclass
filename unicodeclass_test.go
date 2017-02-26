package unicodeclass

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestIs(t *testing.T) {
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

func TestUnicodeSplit(t *testing.T) {
	scan := bufio.NewScanner(strings.NewReader("本日は晴天なり"))
	scan.Split(UnicodeSplit)
	var got []string
	for scan.Scan() {
		got = append(got, scan.Text())
	}
	want := []string{"本日", "は", "晴天", "なり"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v but %v", want, got)
	}
}
