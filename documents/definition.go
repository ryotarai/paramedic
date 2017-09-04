package documents

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"

// 	"github.com/ryotarai/paramedic/awsclient"
// 	"github.com/ryotarai/paramedic/shellwords"

// 	"gopkg.in/yaml.v2"
// )

type Definition struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Script      string `yaml:"script"`
	ScriptFile  string `yaml:"scriptFile"`
	Timeout     string `yaml:"timeout"`
}

func LoadDefinition(file string) (*Definition, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	d := &Definition{}
	err = yaml.Unmarshal(data, d)
	if err != nil {
		return nil, err
	}

	if d.Name == "" {
		return nil, errors.New("name is not set")
	}
	if d.Script == "" {
		if d.ScriptFile == "" {
			return nil, errors.New("neither script nor scriptFile is not set")
		}

		scriptFile := d.ScriptFile
		if !filepath.IsAbs(scriptFile) {
			scriptFile = filepath.Join(filepath.Dir(file), scriptFile)
		}

		b, err := ioutil.ReadFile(scriptFile)
		if err != nil {
			return nil, err
		}

		d.Script = string(b)
	}

	return d, nil
}

func (d *Definition) ScriptSha256() string {
	sum := sha256.Sum256([]byte(d.Script))
	return fmt.Sprintf("%x", sum)
}

// func newDefinitionFromFile(path string) (*definition, error) {
// }

// func (d *definition) documentContent() (string, error) {
// 	cmd := shellwords.Join(d.Command)

// 	i := []awsclient.SSMDocumentRunShellScript{
// 		{
// 			RunCommand: []string{
// 				"export PARAMEDIC_INVOCATION_ID={{ invocationID }}",
// 				fmt.Sprintf("exec %s", cmd),
// 			},
// 		},
// 	}

// 	doc := awsclient.SSMDocument{
// 		SchemaVersion: "2.2",
// 		Description:   "Document created by Paramedic",
// 		Parameters: map[string]map[string]string{
// 			"invocationID": map[string]string{
// 				"type":        "String",
// 				"description": "invocationID",
// 			},
// 		},
// 		MainSteps: []awsclient.SSMDocumentStep{
// 			{
// 				Action: "aws:runShellScript",
// 				Name:   "script",
// 				Inputs: []interface{}(i),
// 			},
// 		},
// 	}

// 	data, err := json.Marshal(doc)
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(data), nil
// }

// func (d *definition) documentName() string {
// 	return fmt.Sprintf("paramedic-%s", d.ID)
// }
