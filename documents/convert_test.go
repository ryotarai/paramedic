package documents

import "testing"

func TestIsParamedicDocument(t *testing.T) {
	if !IsParamedicDocument("paramedic-foo") {
		t.Error("IsParamedicDocuemnt(paramedic-foo) = false, want true")
	}
	if IsParamedicDocument("foo") {
		t.Error("IsParamedicDocuemnt(foo) = true, want false")
	}
}

func TestConvertToSSMName(t *testing.T) {
	want := "paramedic-foo"
	if got := ConvertToSSMName("foo"); got != want {
		t.Errorf("ConvertToSSMName() = %v, want %v", got, want)
	}
}

func TestConvertFromSSMName(t *testing.T) {
	want := "foo"
	if got := ConvertFromSSMName("paramedic-foo"); got != want {
		t.Errorf("ConvertFromSSMName() = %v, want %v", got, want)
	}
}
