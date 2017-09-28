package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func requireStringFlags(keys []string) error {
	missing := []string{}

	for _, k := range keys {
		if viper.GetString(k) == "" {
			missing = append(missing, k)
		}
	}

	if len(missing) == 0 {
		return nil
	}

	var verb string
	if len(missing) > 1 {
		verb = "are"
	} else {
		verb = "is"
	}

	return fmt.Errorf("%s %s required", strings.Join(missing, ", "), verb)
}
