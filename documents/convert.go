package documents

import (
	"fmt"
	"strings"
)

func IsParamedicDocument(name string) bool {
	return strings.HasPrefix(name, "paramedic-")
}

func ConvertToSSMName(name string) string {
	return fmt.Sprintf("paramedic-%s", name)
}

func ConvertFromSSMName(name string) string {
	return strings.TrimPrefix(name, "paramedic-")
}
