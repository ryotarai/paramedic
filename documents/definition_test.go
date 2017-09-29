package documents

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestLoadDefinition(t *testing.T) {
	scriptFile, err := ioutil.TempFile("", "paramedic-test")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(scriptFile.Name())

	scriptFile.WriteString("baz")
	scriptFile.Close()

	defFile, err := ioutil.TempFile("", "paramedic-test")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(defFile.Name())

	fmt.Fprintf(defFile, `name: 'foo'
description: 'bar'
scriptFile: '%s'`, scriptFile.Name())
	defFile.Close()

	d, err := LoadDefinition(defFile.Name())
	if err != nil {
		t.Error(err)
	}

	want := &Definition{
		Name:        "foo",
		Description: "bar",
		Script:      "baz",
		ScriptFile:  scriptFile.Name(),
	}
	if !reflect.DeepEqual(d, want) {
		t.Errorf("got %+v, want %+v", d, want)
	}
}
