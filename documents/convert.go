package documents

import (
	"fmt"
	"strings"
)

func ConvertToSSMName(name string) string {
	return fmt.Sprintf("paramedic-%s", name)
}

func ConvertFromSSMName(name string) string {
	return strings.TrimPrefix(name, "paramedic-")
}
