package fn

import (
	"strings"
)

func StringEqual(first string, second string) (bool, error) {
	return strings.EqualFold(first, second), nil
}
