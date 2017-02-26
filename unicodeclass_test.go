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
	scan.Split(SplitClass)
	var got []string
	for scan.Scan() {
		got = append(got, scan.Text())
	}
	want := []string{"本日", "は", "晴天", "なり"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v but %v", want, got)
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		s    string
		want []string
	}{
		{"本日は晴天なり", []string{"本日", "は", "晴天", "なり"}},
		{"佐藤B作", []string{"佐藤", "B", "作"}},
	}
	for _, test := range tests {
		got := Split(test.s)
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("want %v but %v for %q", test.want, got, test.s)
		}
	}
}
