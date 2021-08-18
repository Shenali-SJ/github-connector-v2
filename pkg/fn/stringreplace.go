package fn

import (
	"strings"
)

func StringReplace(replaceString string, old string, new string) (string, error) {
	return strings.Replace(replaceString, old, new, -1), nil
}
