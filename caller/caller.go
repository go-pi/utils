package caller

import (
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var sourceDir string

func init() {
	_, file, _, _ := runtime.Caller(0)
	sourceDir = regexp.MustCompile(`utils.caller\.go`).ReplaceAllString(file, "")
}

func FileWithLineNum() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.HasPrefix(file, sourceDir) || strings.HasSuffix(file, "_test.go")) {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}
	return ""
}
