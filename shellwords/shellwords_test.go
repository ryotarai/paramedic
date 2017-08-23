package shellwords

import "testing"

func TestEscape(t *testing.T) {
	actual := Escape("a b$c碁")
	expected := "a\\ b\\$c\\碁"
	if actual != expected {
		t.Errorf("expected '%v' but got '%v'", expected, actual)
	}
}

func TestJoin(t *testing.T) {
	actual := Join([]string{"a b", "c d"})
	expected := "a\\ b c\\ d"
	if actual != expected {
		t.Errorf("expected '%v' but got '%v'", expected, actual)
	}
}
