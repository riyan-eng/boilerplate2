package pkg

import (
	"strings"

	"github.com/lib/pq"
)

func PqErrGenerate(err error) (val string) {
	errString := err.(*pq.Error).Detail
	val = strings.ReplaceAll(errString, "Key ", "")
	val = strings.ReplaceAll(val, "(", "")
	val = strings.ReplaceAll(val, ")", "")
	val = strings.ReplaceAll(val, "=", " ")
	return
}
