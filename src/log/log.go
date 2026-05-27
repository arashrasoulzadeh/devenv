package log

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func SetOutput(w io.Writer) {
	out = w
}

func Start() {}

func Print(msgs ...any) {
	fmt.Fprintln(out, msgs...)
}

func Info(msgs ...any) {
	fmt.Fprint(out, "[INFO] ")
	fmt.Fprintln(out, msgs...)
}

func Fatal(msgs ...any) {
	fmt.Fprint(out, "[ERROR] ")
	fmt.Fprintln(out, msgs...)
	os.Exit(1)
}
