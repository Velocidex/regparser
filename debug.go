package regparser

import (
	"fmt"
	"os"
	"strings"
)

func DebugPrint(fmt_str string, v ...interface{}) {
	for _, x := range os.Environ() {
		if strings.HasPrefix(x, "REGPARSER_DEBUG=") {
			fmt.Printf(fmt_str, v...)
			return
		}
	}
}
