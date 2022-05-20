package util

import (
	"fmt"
	"strings"
)

func marshall(str string) {
	str = strings.ReplaceAll(str, `"`, `'`)
	str = strings.ReplaceAll(str, "[", "{")
	str = strings.ReplaceAll(str, "]", "}")
	fmt.Println(str)
}
