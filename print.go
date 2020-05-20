package utilsgo

import (
	"fmt"
	"strings"
)

func PrintFailure(text string) {
	fmt.Println("\u274c ", strings.ToUpper(text))
}

func PrintSuccess(text string) {
	fmt.Println("\xE2\x9C\x94 ", strings.ToUpper(text))
}
