package documents

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

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
		base := filepath.Base(file)
		name := strings.TrimSuffix(base, filepath.Ext(base))
		d.Name = name
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

func (d *Definition) DocumentContent(bucket, key string) (string, error) {
	allowedPattern := "[\\w-/\\.]+"

	j := map[string]interface{}{
		"schemaVersion": "2.2",
		"description":   d.Description,
		"parameters": map[string]map[string]string{
			"executionTimeout": map[string]string{
				"type":           "String",
				"default":        "3600",
				"description":    "(Optional) The time in seconds for a command to be completed before it is considered to have failed. Default is 3600 (1 hour). Maximum is 28800 (8 hours).",
				"allowedPattern": "([1-9][0-9]{0,3})|(1[0-9]{1,4})|(2[0-7][0-9]{1,3})|(28[0-7][0-9]{1,2})|(28800)",
			},
			"outputLogGroup": map[string]string{
				"type":           "String",
				"description":    "(Required) Log group name",
				"allowedPattern": allowedPattern,
			},
			"outputLogStreamPrefix": map[string]string{
				"type":           "String",
				"description":    "(Required) Log stream name prefix",
				"allowedPattern": allowedPattern,
			},
			"signalS3Bucket": map[string]string{
				"type":           "String",
				"description":    "(Required) S3 bucket the signal object is stored in",
				"allowedPattern": allowedPattern,
			},
			"signalS3Key": map[string]string{
				"type":           "String",
				"description":    "(Required) S3 object key the signal object is stored at",
				"allowedPattern": allowedPattern,
			},
		},
		"mainSteps": []interface{}{
			map[string]interface{}{
				"action": "aws:runShellScript",
				"name":   "script",
				"inputs": map[string]interface{}{
					"runCommand": []string{
						"export PARAMEDIC_OUTPUT_LOG_GROUP={{outputLogGroup}}",
						"export PARAMEDIC_OUTPUT_LOG_STREAM_PREFIX={{outputLogStreamPrefix}}",
						"export PARAMEDIC_SIGNAL_S3_BUCKET={{signalS3Bucket}}",
						"export PARAMEDIC_SIGNAL_S3_KEY={{signalS3Key}}",
						fmt.Sprintf("export PARAMEDIC_SCRIPT_S3_BUCKET=%s", bucket),
						fmt.Sprintf("export PARAMEDIC_SCRIPT_S3_KEY=%s", key),
						"exec paramedic-agent",
					},
					"timeoutSeconds": "{{ executionTimeout }}",
				},
			},
		},
	}

	b, err := json.Marshal(j)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
