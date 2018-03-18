package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/cihub/seelog"
)

var workingDir = "/"

func init() {
	wd, err := os.Getwd()
	if err == nil {
		workingDir = filepath.ToSlash(wd) + "/"
	}
}

type callerInfo struct {
	FullPath  string
	ShortPath string
	FuncName  string
	Line      int
}

func mkCallerInfo(skip int) (info callerInfo, err error) {
	pc, fp, ln, ok := runtime.Caller(skip)
	if !ok {
		err = fmt.Errorf("error during runtime.Caller")
		return
	}
	info.Line = ln
	info.FullPath = fp
	if strings.HasPrefix(fp, workingDir) {
		info.ShortPath = fp[len(workingDir):]
	} else {
		info.ShortPath = fp
	}
	info.FuncName = runtime.FuncForPC(pc).Name()
	if strings.HasPrefix(info.FuncName, workingDir) {
		info.FuncName = info.FuncName[len(workingDir):]
	}
	return
}

// LogIfPanic log panic info with stack info
func LogIfPanic(err interface{}) {
	if err == nil {
		return
	}
	c, _ := mkCallerInfo(2)
	seelog.Criticalf("Panic err : %v, in %s at %s[%d] \r\nStack : %s",
		err, c.FuncName, c.ShortPath, c.Line, string(debug.Stack()))
	seelog.Flush()
}
