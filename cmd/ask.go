package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func askContinue(msg string) (bool, error) {
	fmt.Printf("%s (y/N): ", msg)

	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		return false, err
	}
	return strings.HasPrefix(line, "y"), nil
}
